import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'redirector',
      component: () => import('../views/RedirectorView.vue')
    },
    {
      path: '/auth',
      name: 'auth',
      component: () => import('../views/AuthView.vue')
    },
    {
      path: '/restore',
      name: 'restore',
      component: () => import('../views/RestoreView.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue')
    },
    {
      path: '/static/howto',
      name: 'howto',
      component: () => import('../views/static/HowToView.vue')
    },
    {
      path: '/static/precautions',
      name: 'precautions',
      component: () => import('../views/static/PrecautionsView.vue')
    },
    {
      path: '/donations',
      name: 'donations',
      component: () => import('../views/authorized/donations/DonationsView.vue')
    },
    {
      path: '/donation/create',
      name: 'createDonation',
      component: () => import('../views/authorized/donations/CreateDonationView.vue')
    },
    {
      path: '/donation/plan',
      name: 'planDonation',
      component: () => import('../views/authorized/donations/PlanDonationView.vue')
    },
    {
      path: '/profile/setup',
      name: 'setupProfile',
      component: () => import('../views/authorized/user/SetupProfileView.vue')
    },
    {
      path: '/stations',
      name: 'stations',
      component: () => import('../views/authorized/stations/BloodStationsView.vue')
    },
    {
      path: '/stations/map',
      name: 'stationsMap',
      component: () => import('../views/authorized/stations/BloodStationsMapView.vue')
    },
    {
      path: '/bonuses',
      name: 'bonuses',
      component: () => import('../views/authorized/bonuses/BonusesView.vue')
    },
    {
      path: '/events',
      name: 'events',
      component: () => import('../views/authorized/events/EventsView.vue')
    },
    {
      path: '/top',
      name: 'top',
      component: () => import('../views/authorized/user/DonationsTopView.vue')
    },
  ]
})

export default router
