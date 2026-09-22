<template>
  <div class="layout">
    <!-- Mobile top bar -->
    <header class="mobile-bar">
      <button class="hamburger" aria-label="Toggle menu" @click="isSidebarOpen = !isSidebarOpen">
        <span /><span /><span />
      </button>
      <span class="mobile-brand">Schedula</span>
    </header> 

    <!-- Sidebar -->
    <aside class="sidebar" :class="{ open: isSidebarOpen }">
      <div class="brand-mark">Schedula</div>

      <nav class="nav-list">
        <RouterLink to="/dashboard" class="nav-item" active-class="active">
          <IconDashboard />
          <span>Dashboard</span>
        </RouterLink>
        <RouterLink to="/services" class="nav-item" active-class="active">
          <IconServices />
          <span>Services</span>
        </RouterLink>
        <RouterLink to="/customers" class="nav-item" active-class="active">
          <IconCustomers />
          <span>Customers</span>
        </RouterLink>
        <RouterLink to="/appointments" class="nav-item" active-class="active">
          <IconAppointments />
          <span>Appointments</span>
        </RouterLink>
      </nav>

      <button class="logout-btn" @click="handleLogout">
        <IconLogout />
        <span>Logout</span>
      </button>
    </aside>

    <!-- Overlay for mobile drawer -->
    <div v-if="isSidebarOpen" class="overlay" @click="isSidebarOpen = false" />

    <!-- Main content -->
    <main class="content">
      <h1 class="page-title">Dashboard</h1>
      <p class="page-subtitle">Here's how things stand today.</p>

      <section class="summary-grid">
        <article class="summary-card">
          <div class="card-icon services"><IconServices /></div>
          <div class="card-text">
            <span class="card-label">Total Services</span>
            <span class="card-value">{{ formatValue(totalServices) }}</span>
          </div>
        </article>

        <article class="summary-card">
          <div class="card-icon customers"><IconCustomers /></div>
          <div class="card-text">
            <span class="card-label">Total Customers</span>
            <span class="card-value">{{ formatValue(totalCustomers) }}</span>
          </div>
        </article>

        <article class="summary-card">
          <div class="card-icon today"><IconCalendar /></div>
          <div class="card-text">
            <span class="card-label">Today's Appointments</span>
            <span class="card-value">{{ formatValue(todaysAppointments) }}</span>
          </div>
        </article>

        <article class="summary-card">
          <div class="card-icon booked"><IconCheck /></div>
          <div class="card-text">
            <span class="card-label">Booked Appointments</span>
            <span class="card-value">{{ formatValue(bookedAppointments) }}</span>
          </div>
        </article>
      </section>
    </main>
  </div>
</template>

<!-- <script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useDashboardStore } from '@/store/dashboard'
const totalServices = ref(null)
const totalCustomers = ref(null)
const todaysAppointments = ref(null)
const bookedAppointments = ref(null)
const dashboard=useDashboardStore()

const formatValue = (val) => (val === null || val === undefined ? '—' : val)

const router = useRouter()
const authStore = useAuthStore()
const isSidebarOpen = ref(false)

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script> -->
<!-- <script setup>
import { ref } from 'vue'
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useDashboardStore } from '@/store/dashboard'

const router = useRouter()
const authStore = useAuthStore()
const dashboardStore = useDashboardStore()

const {
  totalServices,
  totalCustomers,
  todaysAppointments,
  bookedAppointments
} = storeToRefs(dashboardStore)

const formatValue = (val) =>
  val === null || val === undefined ? '—' : val

const isSidebarOpen = ref(false)

// Load dashboard data when page opens
onMounted(() => {
  dashboardStore.fetchDashboardData()
})

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script> -->

<!-- Inline icon components (dependency-free) -->
<!-- <script>
const IconDashboard = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3" y="3" width="7" height="9" rx="1.5"/><rect x="14" y="3" width="7" height="5" rx="1.5"/><rect x="14" y="12" width="7" height="9" rx="1.5"/><rect x="3" y="16" width="7" height="5" rx="1.5"/></svg>`
}

const IconServices = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M14.7 6.3a4 4 0 0 1-5.4 5.4L4 17v3h3l5.3-5.3a4 4 0 0 1 5.4-5.4Z"/></svg>`
}

const IconCustomers = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="9" cy="8" r="3.2"/><path d="M3.5 20c0-3.3 2.5-5.5 5.5-5.5s5.5 2.2 5.5 5.5"/><circle cx="17" cy="8.5" r="2.4"/><path d="M15.2 14.8c2.6.3 4.3 2.4 4.3 5.2"/></svg>`
}

const IconAppointments = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3.5" y="4.5" width="17" height="16" rx="2"/><path d="M3.5 9.5h17M8 3v3M16 3v3"/></svg>`
}

const IconCalendar = IconAppointments

const IconCheck = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="12" cy="12" r="9"/><path d="M8 12.5l2.5 2.5L16 9.5"/></svg>`
}

const IconLogout = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M9 4H5.5A1.5 1.5 0 0 0 4 5.5v13A1.5 1.5 0 0 0 5.5 20H9"/><path d="M14 15l4-3-4-3M18 12H9"/></svg>`
}

export default {
  components: {
    IconDashboard,
    IconServices,
    IconCustomers,
    IconAppointments,
    IconCalendar,
    IconCheck,
    IconLogout
  }
}
</script> -->


<!-- <script>
const IconDashboard = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3" y="3" width="7" height="9" rx="1.5"/><rect x="14" y="3" width="7" height="5" rx="1.5"/><rect x="14" y="12" width="7" height="9" rx="1.5"/><rect x="3" y="16" width="7" height="5" rx="1.5"/></svg>`
}
const IconServices = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M14.7 6.3a4 4 0 0 1-5.4 5.4L4 17v3h3l5.3-5.3a4 4 0 0 1 5.4-5.4Z"/></svg>`
}
const IconCustomers = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="9" cy="8" r="3.2"/><path d="M3.5 20c0-3.3 2.5-5.5 5.5-5.5s5.5 2.2 5.5 5.5"/><circle cx="17" cy="8.5" r="2.4"/><path d="M15.2 14.8c2.6.3 4.3 2.4 4.3 5.2"/></svg>`
}
const IconAppointments = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3.5" y="4.5" width="17" height="16" rx="2"/><path d="M3.5 9.5h17M8 3v3M16 3v3"/></svg>`
}
const IconCalendar = IconAppointments
const IconCheck = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="12" cy="12" r="9"/><path d="M8 12.5l2.5 2.5L16 9.5"/></svg>`
}
const IconLogout = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M9 4H5.5A1.5 1.5 0 0 0 4 5.5v13A1.5 1.5 0 0 0 5.5 20H9"/><path d="M14 15l4-3-4-3M18 12H9"/></svg>`
}

export default {
  components: {
    IconDashboard,
    IconServices,
    IconCustomers,
    IconAppointments,
    IconCalendar,
    IconCheck,
    IconLogout
  }
} -->
<!-- </script> -->

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'
import { useServiceStore } from '@/store/service'
// import { useCustomerStore } from '@/store/customer'         // TODO: once built
// import { useAppointmentStore } from '@/store/appointment'   // TODO: once built

const router = useRouter()
const authStore = useAuthStore()
const serviceStore = useServiceStore()

// Real value, comes straight from the services list already being fetched
const totalServices = computed(() => serviceStore.totalServices)

// Placeholders until Customer/Appointment stores + backend routes exist
const totalCustomers = ref(null)
const todaysAppointments = ref(null)
const bookedAppointments = ref(null)

const formatValue = (val) => (val === null || val === undefined ? '—' : val)

const isSidebarOpen = ref(false)

onMounted(() => {
  serviceStore.fetchServices()
})

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}
</script>
<!-- Inline icon components (dependency-free) -->
<script>
const IconDashboard = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3" y="3" width="7" height="9" rx="1.5"/><rect x="14" y="3" width="7" height="5" rx="1.5"/><rect x="14" y="12" width="7" height="9" rx="1.5"/><rect x="3" y="16" width="7" height="5" rx="1.5"/></svg>`
}

const IconServices = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M14.7 6.3a4 4 0 0 1-5.4 5.4L4 17v3h3l5.3-5.3a4 4 0 0 1 5.4-5.4Z"/></svg>`
}

const IconCustomers = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="9" cy="8" r="3.2"/><path d="M3.5 20c0-3.3 2.5-5.5 5.5-5.5s5.5 2.2 5.5 5.5"/><circle cx="17" cy="8.5" r="2.4"/><path d="M15.2 14.8c2.6.3 4.3 2.4 4.3 5.2"/></svg>`
}

const IconAppointments = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><rect x="3.5" y="4.5" width="17" height="16" rx="2"/><path d="M3.5 9.5h17M8 3v3M16 3v3"/></svg>`
}

const IconCalendar = IconAppointments

const IconCheck = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><circle cx="12" cy="12" r="9"/><path d="M8 12.5l2.5 2.5L16 9.5"/></svg>`
}

const IconLogout = {
  template: `<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M9 4H5.5A1.5 1.5 0 0 0 4 5.5v13A1.5 1.5 0 0 0 5.5 20H9"/><path d="M14 15l4-3-4-3M18 12H9"/></svg>`
}

export default {
  components: {
    IconDashboard,
    IconServices,
    IconCustomers,
    IconAppointments,
    IconCalendar,
    IconCheck,
    IconLogout
  }
}
</script>


<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,500;9..144,600&family=Inter:wght@400;500;600&display=swap');

* { box-sizing: border-box; }

.layout {
  min-height: 100vh;
  display: flex;
  font-family: 'Inter', system-ui, sans-serif;
  background: #faf9f7;
}

/* Sidebar */
.sidebar {
  width: 240px;
  flex-shrink: 0;
  background: #14213d;
  color: #f4f2ee;
  display: flex;
  flex-direction: column;
  padding: 1.75rem 1.25rem;
  position: sticky;
  top: 0;
  height: 100vh;
}

.brand-mark {
  font-family: 'Fraunces', serif;
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 2.25rem;
  padding: 0 0.5rem;
}

.nav-list {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  flex: 1;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.65rem 0.75rem;
  border-radius: 8px;
  color: rgba(244, 242, 238, 0.75);
  text-decoration: none;
  font-size: 0.92rem;
  font-weight: 500;
  transition: background 0.15s ease, color 0.15s ease;
}

.nav-item:hover { background: rgba(244, 242, 238, 0.08); color: #f4f2ee; }

.nav-item.active {
  background: rgba(232, 163, 61, 0.16);
  color: #e8a33d;
}

.nav-item svg { flex-shrink: 0; }

.logout-btn {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.65rem 0.75rem;
  border-radius: 8px;
  background: none;
  border: 1px solid rgba(244, 242, 238, 0.15);
  color: rgba(244, 242, 238, 0.85);
  font-family: inherit;
  font-size: 0.9rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.15s ease, border-color 0.15s ease;
}

.logout-btn:hover { background: rgba(244, 242, 238, 0.08); border-color: rgba(244, 242, 238, 0.3); }

/* Main content */
.content {
  flex: 1;
  padding: 2.75rem 3rem;
  max-width: 1080px;
}

.page-title {
  font-family: 'Fraunces', serif;
  font-size: 1.9rem;
  font-weight: 600;
  color: #14213d;
  margin: 0 0 0.35rem;
}

.page-subtitle {
  color: #5b6472;
  margin: 0 0 2rem;
  font-size: 0.95rem;
}

.summary-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1rem;
}

.summary-card {
  background: #fff;
  border: 1px solid #e4e2dc;
  border-radius: 10px;
  padding: 1.25rem;
  display: flex;
  align-items: center;
  gap: 0.9rem;
}

.card-icon {
  width: 42px;
  height: 42px;
  border-radius: 9px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.card-icon.services { background: rgba(20, 33, 61, 0.08); color: #14213d; }
.card-icon.customers { background: rgba(91, 100, 114, 0.12); color: #5b6472; }
.card-icon.today { background: rgba(232, 163, 61, 0.16); color: #e8a33d; }
.card-icon.booked { background: rgba(46, 139, 87, 0.12); color: #2e8b57; }

.card-text { display: flex; flex-direction: column; gap: 0.15rem; }

.card-label { font-size: 0.8rem; color: #5b6472; }

.card-value {
  font-size: 1.5rem;
  font-weight: 600;
  color: #14213d;
  font-variant-numeric: tabular-nums;
}

/* Mobile bar + drawer */
.mobile-bar { display: none; }
.overlay { display: none; }

@media (max-width: 900px) {
  .summary-grid { grid-template-columns: repeat(2, 1fr); }
}

@media (max-width: 720px) {
  .layout { flex-direction: column; }

  .mobile-bar {
    display: flex;
    align-items: center;
    gap: 1rem;
    background: #14213d;
    color: #f4f2ee;
    padding: 1rem 1.25rem;
    position: sticky;
    top: 0;
    z-index: 20;
  }

  .mobile-brand { font-family: 'Fraunces', serif; font-weight: 600; }

  .hamburger {
    background: none;
    border: none;
    display: flex;
    flex-direction: column;
    gap: 4px;
    cursor: pointer;
    padding: 0.25rem;
  }

  .hamburger span {
    width: 20px;
    height: 2px;
    background: #f4f2ee;
    border-radius: 2px;
  }

  .sidebar {
    position: fixed;
    left: 0;
    top: 0;
    height: 100vh;
    z-index: 30;
    transform: translateX(-100%);
    transition: transform 0.2s ease;
  }

  .sidebar.open { transform: translateX(0); }

  .overlay {
    display: block;
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.4);
    z-index: 25;
  }

  .content { padding: 1.5rem; }
  .summary-grid { grid-template-columns: 1fr; }
}

@media (prefers-reduced-motion: reduce) {
  .sidebar { transition: none; }
}
</style>