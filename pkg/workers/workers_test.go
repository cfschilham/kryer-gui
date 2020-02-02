package workers_test

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/cfschilham/kryer/pkg/workers"
	"golang.org/x/crypto/bcrypt"
)

func TestPool(t *testing.T) {
	var (
		poolSize    = 8
		taskQty     = 16
		taskContent = "test"
	)

	// Set the task and its parameters.
	c, wg := make(chan string, taskQty), &sync.WaitGroup{}
	task := workers.Task{
		Fn: func(params []interface{}) {
			var (
				taskContent = params[0].(string)
				c           = params[1].(chan string)
				wg          = params[2].(*sync.WaitGroup)
			)

			c <- taskContent
			wg.Done()
		},
		Params: []interface{}{taskContent, c, wg},
	}

	// Create the pool.
	pool, err := workers.NewPool(poolSize)
	if err != nil {
		t.Errorf("failed to create new pool")
		return
	}

	// Enqueue all tasks to the pools queue.
	for i := 0; i < taskQty; i++ {
		wg.Add(1)
		pool.QueueTask(task)
	}
	pool.Start()
	defer pool.Close()
	wg.Wait()

	// Check results
	if len(c) != taskQty {
		t.Errorf("did not receive enough results, want: %d got: %d", taskQty, len(c))
		return
	}

	for i := 0; i < len(c); i++ {
		want, got := taskContent, <-c
		if got != want {
			t.Errorf("unexpected result, want: %s got: %s", want, got)
			return
		}
	}
}

func TestPoolGoroutines(t *testing.T) {
	before := runtime.NumGoroutine()
	t.Run("", TestPool)

	after := time.After(time.Millisecond * 50)
	for runtime.NumGoroutine() > before {
		select {
		case <-after:
			t.Errorf("failed to close all of the pools goroutines within 50ms, before: %d now: %d", before, runtime.NumGoroutine())
			return
		default:
		}
	}

}

func BenchmarkPool(b *testing.B) {
	var (
		poolSize = runtime.NumCPU()
		taskQty  = runtime.NumCPU() * 2 * b.N
	)

	// Set the task and its parameters.
	wg := &sync.WaitGroup{}
	task := workers.Task{
		Fn: func(params []interface{}) {
			var (
				wg = params[0].(*sync.WaitGroup)
			)

			bcrypt.GenerateFromPassword([]byte("benchmark"), 12)
			wg.Done()
		},
		Params: []interface{}{wg},
	}

	// Create the pool.
	pool, err := workers.NewPool(poolSize)
	if err != nil {
		b.Errorf("failed to create new pool")
		return
	}

	// Enqueue all tasks to the pools queue.
	for i := 0; i < taskQty; i++ {
		wg.Add(1)
		pool.QueueTask(task)
	}
	pool.Start()
	defer pool.Close()
	wg.Wait()
}
