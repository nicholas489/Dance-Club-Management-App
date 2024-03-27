import {createRouter, createWebHistory} from 'vue-router'
import LoginView from '../views/LoginView.vue'
import MemberDashboardView from "@/views/MemberDashboardView.vue";
import CoachDashboardView from "@/views/CoachDashboardView.vue";


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
            name: 'MemberDashboardView',
            component: MemberDashboardView, // Your component for member dashboard
        },
        {
            path: '/coach/dashboard',
            name: 'CoachDashboardView',
            component: CoachDashboardView, // Your component for coach dashboard
        }
    ]
})

export default router
