import Vue from "vue"
import VueRouter from "vue-router"
import Home from "../views/Home.vue"

Vue.use(VueRouter)

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/hostlist/:id",
    name: "hostlist",
    component: () =>
      import(/* webpackChunkName: "hostlist" */ "../views/Hostlist.vue"),
  },
  {
    path: "/manual/:id",
    name: "manual",
    component: () =>
      import(/* webpackChunkName: "manual" */ "../views/Manual.vue"),
  },
]

const router = new VueRouter({
  routes,
})

export default router
