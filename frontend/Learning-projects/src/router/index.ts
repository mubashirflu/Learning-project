import { createRouter, createWebHistory } from "vue-router";

import Login from "@/view/login.vue";
import Register from "@/view/register.vue";
import Dashboard from "@/view/dashboard.vue";
import Services from "@/view/services.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),

  routes: [
    {
      path: "/",
      name: "home",
      component: Login,
    },
    {
      path: "/login",
      name: "login",
      component: Login,
    },
    {
      path: "/register",
      name: "register",
      component: Register,
    },
    {
      path: "/dashboard",
      name: "dashboard",
      component: Dashboard,
      meta: {
        requiresAuth: true,
      },
    },
        {
      path: "/services",
      name: "services",
      component: Services,
    },
  ],
});

router.beforeEach((to) => {
  const token = localStorage.getItem("token");

  if (to.meta.requiresAuth && !token) {
    return "/login";
  }

  return true;
});

export default router;
