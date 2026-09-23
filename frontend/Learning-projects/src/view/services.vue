<template>
  <div>
    <div class="page-header">
      <div>
        <h1 class="page-title">Services</h1>
        <p class="page-subtitle">Manage the services your business offers.</p>
      </div>
      <button class="btn-primary" @click="openAddModal">
        <span class="plus">+</span> Add Service
      </button>
    </div>

    <div class="toolbar">
      <input
        v-model="searchQuery"
        type="text"
        class="search-input"
        placeholder="Search services by name…"
      />
    </div>

    <!-- Loading -->
    <div v-if="serviceStore.isLoading" class="state-box">Loading services…</div>

    <!-- Error -->
    <div v-else-if="serviceStore.error" class="state-box error">
      {{ serviceStore.error }}
      <button class="retry-btn" @click="serviceStore.fetchServices()">Retry</button>
    </div>

    <!-- Empty -->
    <div v-else-if="filteredServices.length === 0" class="state-box">
      <p v-if="searchQuery">No services match "{{ searchQuery }}".</p>
      <template v-else>
        <p>No services yet.</p>
        <button class="btn-primary" @click="openAddModal">+ Add your first service</button>
      </template>
    </div>

    <!-- Table -->
    <div v-else class="table-card">
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Duration</th>
            <th>Price</th>
            <th>Description</th>
            <th class="actions-col">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="service in filteredServices" :key="service.id">
            <td class="cell-name">{{ service.name }}</td>
            <td>{{ service.duration_minutes }} min</td>
            <td>{{ formatPrice(service.price) }}</td>
            <td class="cell-desc">{{ service.description || '—' }}</td>
            <td class="actions-col">
              <button class="icon-btn" aria-label="Edit" @click="openEditModal(service)">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M14.7 6.3a4 4 0 0 1-5.4 5.4L4 17v3h3l5.3-5.3a4 4 0 0 1 5.4-5.4Z"/></svg>
              </button>
              <button class="icon-btn danger" aria-label="Delete" @click="confirmDelete(service)">
                <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8"><path d="M4 7h16M9 7V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2m-8 0 1 13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1l1-13"/></svg>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Add/Edit modal -->
    <ServiceFormModal
      v-if="isModalOpen"
      :service="editingService"
      @close="isModalOpen = false"
      @saved="onSaved"
    />

    <!-- Delete confirm -->
    <div v-if="deleteTarget" class="modal-overlay" @click.self="deleteTarget = null">
      <div class="confirm-card">
        <h2>Delete "{{ deleteTarget.name }}"?</h2>
        <p>This can't be undone.</p>
        <div class="confirm-actions">
          <button class="btn-secondary" @click="deleteTarget = null">Cancel</button>
          <button class="btn-danger" :disabled="isDeleting" @click="handleDelete">
            {{ isDeleting ? 'Deleting…' : 'Delete' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useServiceStore, type Service } from '@/store/service'
import ServiceFormModal from '@/components/serviceformmodel.vue'

const serviceStore = useServiceStore()

const searchQuery = ref('')
const isModalOpen = ref(false)
const editingService = ref<Service | null>(null)
const deleteTarget = ref<Service | null>(null)
const isDeleting = ref(false)

const filteredServices = computed(() => {
  const q = searchQuery.value.trim().toLowerCase()
  if (!q) return serviceStore.services
  return serviceStore.services.filter((s) => s.name.toLowerCase().includes(q))
})

const formatPrice = (price: number) => `$${Number(price).toFixed(2)}`

const openAddModal = () => {
  editingService.value = null
  isModalOpen.value = true
}

const openEditModal = (service: Service) => {
  editingService.value = service
  isModalOpen.value = true
}

const onSaved = () => {
  // store already updates its own state on add/update, nothing else needed here
}

const confirmDelete = (service: Service) => {
  deleteTarget.value = service
}

const handleDelete = async () => {
  if (!deleteTarget.value) return
  isDeleting.value = true
  try {
    await serviceStore.deleteService(deleteTarget.value.id)
    deleteTarget.value = null
  } catch (err) {
    // keep the confirm dialog open so the user sees something went wrong
  } finally {
    isDeleting.value = false
  }
}

onMounted(() => {
  serviceStore.fetchServices()
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,600&family=Inter:wght@400;500;600&display=swap');

* { box-sizing: border-box; }

.page-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
  margin-bottom: 1.5rem;
  font-family: 'Inter', system-ui, sans-serif;
}

.page-title {
  font-family: 'Fraunces', serif;
  font-size: 1.9rem;
  font-weight: 600;
  color: #14213d;
  margin: 0 0 0.3rem;
}

.page-subtitle { color: #5b6472; margin: 0; font-size: 0.95rem; }

.toolbar { margin-bottom: 1.25rem; font-family: 'Inter', system-ui, sans-serif; }

.search-input {
  width: 100%;
  max-width: 320px;
  padding: 0.6rem 0.85rem;
  border: 1px solid #e4e2dc;
  border-radius: 8px;
  font-size: 0.9rem;
  font-family: inherit;
}

.search-input:focus {
  outline: none;
  border-color: #14213d;
  box-shadow: 0 0 0 3px rgba(20, 33, 61, 0.08);
}

.state-box {
  background: #fff;
  border: 1px solid #e4e2dc;
  border-radius: 10px;
  padding: 2.5rem;
  text-align: center;
  color: #5b6472;
  font-family: 'Inter', system-ui, sans-serif;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.9rem;
}

.state-box.error { color: #c0392b; }

.retry-btn {
  background: none;
  border: 1px solid currentColor;
  border-radius: 6px;
  padding: 0.35rem 0.8rem;
  cursor: pointer;
  color: inherit;
  font-size: 0.85rem;
}

.table-card {
  background: #fff;
  border: 1px solid #e4e2dc;
  border-radius: 10px;
  overflow: hidden;
  font-family: 'Inter', system-ui, sans-serif;
}

table { width: 100%; border-collapse: collapse; }

thead { background: #faf9f7; }

th {
  text-align: left;
  font-size: 0.78rem;
  text-transform: uppercase;
  letter-spacing: 0.03em;
  color: #5b6472;
  padding: 0.85rem 1.1rem;
  font-weight: 600;
  border-bottom: 1px solid #e4e2dc;
}

td {
  padding: 0.85rem 1.1rem;
  font-size: 0.9rem;
  color: #14213d;
  border-bottom: 1px solid #f0efeb;
}

tbody tr:last-child td { border-bottom: none; }
tbody tr:hover { background: #faf9f7; }

.cell-name { font-weight: 500; }

.cell-desc {
  color: #5b6472;
  max-width: 240px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.actions-col { text-align: right; white-space: nowrap; }

.icon-btn {
  background: none;
  border: 1px solid #e4e2dc;
  border-radius: 6px;
  padding: 0.4rem;
  cursor: pointer;
  color: #5b6472;
  margin-left: 0.4rem;
  display: inline-flex;
}

.icon-btn:hover { background: #faf9f7; color: #14213d; }
.icon-btn.danger:hover { color: #c0392b; border-color: #c0392b; }

.btn-primary {
  background: #14213d;
  color: #f4f2ee;
  border: none;
  border-radius: 8px;
  padding: 0.65rem 1.1rem;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  font-family: 'Inter', system-ui, sans-serif;
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  white-space: nowrap;
}

.btn-primary:hover { background: #1c2d52; }

.plus { font-size: 1rem; }

/* Delete confirm modal */
.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(20, 33, 61, 0.45);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1.5rem;
  z-index: 50;
}

.confirm-card {
  background: #fff;
  border-radius: 12px;
  padding: 1.75rem;
  width: 100%;
  max-width: 360px;
  font-family: 'Inter', system-ui, sans-serif;
}

.confirm-card h2 {
  font-family: 'Fraunces', serif;
  font-size: 1.15rem;
  color: #14213d;
  margin: 0 0 0.4rem;
}

.confirm-card p { color: #5b6472; font-size: 0.9rem; margin: 0 0 1.4rem; }

.confirm-actions { display: flex; justify-content: flex-end; gap: 0.6rem; }

.btn-secondary {
  background: #fff;
  border: 1px solid #e4e2dc;
  color: #14213d;
  border-radius: 8px;
  padding: 0.55rem 1rem;
  font-size: 0.88rem;
  font-weight: 600;
  cursor: pointer;
  font-family: inherit;
}

.btn-secondary:hover { background: #faf9f7; }

.btn-danger {
  background: #c0392b;
  color: #fff;
  border: none;
  border-radius: 8px;
  padding: 0.55rem 1rem;
  font-size: 0.88rem;
  font-weight: 600;
  cursor: pointer;
  font-family: inherit;
}

.btn-danger:hover:not(:disabled) { background: #a5301f; }
.btn-danger:disabled { opacity: 0.6; cursor: not-allowed; }

@media (max-width: 640px) {
  .page-header { flex-direction: column; align-items: stretch; }
  .cell-desc { display: none; }
  th:nth-child(4) { display: none; }
}
</style>