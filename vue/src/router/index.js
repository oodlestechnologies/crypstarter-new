import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/Index.vue'
import Types from '@/views/Types.vue'
import Staking from '@/views/Staking.vue'
import Unstaking from '@/views/Unstaking.vue'
import Withdraw from '@/views/Withdraw.vue'
// import Footer from '@/views/Footer.vue'

const routerHistory = createWebHistory()
const routes = [
	{
		path: '/',
		component: Index
	},
	{ path: '/types', component: Types },
	{ path: '/Staking', component: Staking },
	{ path: '/Unstaking', component:Unstaking},
	{ path: '/Withdraw', component:Withdraw}
]

const router = createRouter({
	history: routerHistory,
	routes
})

export default router
