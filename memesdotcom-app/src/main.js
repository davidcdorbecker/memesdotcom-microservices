import { createApp } from 'vue'
import App from './App.vue'
import Spinner from '@/components/Spinner'
import router from './router'
import './assets/index.css'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import './icons'

const app = createApp(App)

app.component('Icon', FontAwesomeIcon)
app.component('Spinner', Spinner)

app.use(router).mount('#app')
