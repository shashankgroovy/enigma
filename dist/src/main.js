import App from './components/App.js';
import ShareSecret from './components/ShareSecret.js'
import RevealSecret from './components/RevealSecret.js'

Vue.use(VueRouter)

const routes = [
  { path: '/', component: ShareSecret, name: 'share' },
  { path: '/secret/:slug', component: RevealSecret, name: 'reveal' }
]

const router = new VueRouter({
  routes,
  mode: 'history'
})

new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  data: {
    currentRoute: window.location.pathname,
  },
}).$mount('#app')
