import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import vueLib from '@starport/vue'
import VueToast from 'vue-toast-notification'
import 'vue-toast-notification/dist/theme-sugar.css'

const app = createApp(App)
app.use(VueToast)
app.config.globalProperties._depsLoaded = true
app.use(store).use(router).use(vueLib).mount('#app')

app.use(VueToast, {
	// One of the options
	position: 'top'
})

// app.$toast.success('Order placed.', {
// 	// override the global option
// 	position: 'top'
// })
//Vue.$toast.open({/* options */});
// let instance = app.$toast.open('You did it!')

// Force dismiss specific toast
// instance.dismiss()
// Dismiss all opened toast immediately
// app.$toast.clear()
