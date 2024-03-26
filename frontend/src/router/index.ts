import {createRouter, createWebHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import MemberDashboardView from "@/views/MemberDashboardView.vue";


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: LoginView
        },
        {
            path: '/member/dashboard',
            name: 'memberDashboard',
            component: MemberDashboardView
        }
    ]
})

export default router
