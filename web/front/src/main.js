import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify';
import day from 'dayjs'

Vue.filter('dateformat', function(indate, outdate) {
  return day(indate).format(outdate)
})


import './plugins/http'
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
