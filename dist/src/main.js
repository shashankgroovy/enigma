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

router.beforeResolve((to, from, next) => {
  // If this isn't an initial page load.
  if (to.name) {
      // Start the route progress bar.
      NProgress.start()
  }
  next()
})

router.afterEach((to, from) => {
  // Complete the animation of the route progress bar.
  NProgress.done()
})

axios.interceptors.request.use(config => {
  NProgress.start()
  return config
})

// before a response is returned stop nprogress
axios.interceptors.response.use(response => {
  NProgress.done()
  return response
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
