import Vue from 'vue'
import App from './App.vue'
import vuescroll from 'vuescroll';
import ProgressBar from 'vuejs-progress-bar'

Vue.use(ProgressBar);

Vue.config.productionTip = false

Vue.use(vuescroll);

new Vue({
  render: h => h(App),
  ProgressBar,
  vuescroll
}).$mount('#app')