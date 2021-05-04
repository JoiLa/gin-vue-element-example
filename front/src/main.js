import Vue from 'vue'

import Cookies from 'js-cookie'
import 'normalize.css/normalize.css' // A modern alternative to CSS resets

import ElementUI from 'element-ui'
import i18n      from './lang'
import 'element-ui/lib/theme-chalk/index.css'

import '@/styles/index.scss' // global css

import App    from './App'
import store  from './store'
import router from './router'

import '@/icons' // icon
import '@/permission' // permission control



// set ElementUI lang to EN
Vue.use(ElementUI, {
  size: Cookies.get('size') || 'medium', // set element-ui default size
  i18n: (key, value) => i18n.t(key, value)
})
// 如果想要中文版 element-ui，按如下方式声明
// Vue.use(ElementUI)

Vue.config.productionTip = false

new Vue({
  el    : '#app',
  i18n,
  router,
  store,
  render: h => h(App)
})
