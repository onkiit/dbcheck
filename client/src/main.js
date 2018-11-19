import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from './store'
import http from './http'

Vue.config.productionTip = false

new Vue({
  router,
  store,
  http,
  render: h => h(App)
}).$mount('#app')
