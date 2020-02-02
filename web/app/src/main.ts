import Vue from "vue"
import UniqueId from "vue-unique-id"
import router from "./router"
import App from "./App.vue"

Vue.config.productionTip = false
Vue.config.devtools = true

Vue.use(UniqueId)

new Vue({
  router,
  render: h => h(App),
}).$mount("#app")
