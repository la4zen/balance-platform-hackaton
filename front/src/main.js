import Vue from 'vue'
import App from './App.vue'
import vuescroll from 'vuescroll';

Vue.use(vuescroll, {
  ops: {
    vuescroll: {},
    scrollPanel: {},
    rail: {},
    bar: {}
  },
  name: 'myScroll'
});

// Vue.prototype.$vuescrollConfig = {
//   bar: {
//     background: '#5ED3BB'
//   }
// };

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')