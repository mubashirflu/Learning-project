<template>
  <div class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-card" role="dialog" aria-modal="true" :aria-label="isEditMode ? 'Edit service' : 'Add service'">
      <div class="modal-header">
        <h2>{{ isEditMode ? 'Edit Service' : 'Add Service' }}</h2>
        <button class="close-btn" aria-label="Close" @click="$emit('close')">&times;</button>
      </div>

      <form @submit.prevent="handleSubmit" novalidate>
        <div class="field">
          <label for="svc-name">Service name</label>
          <input
            id="svc-name"
            v-model="form.name"
            type="text"
            placeholder="e.g. Haircut"
            :class="{ invalid: touched.name && !isNameValid }"
            @blur="touched.name = true"
          />
          <span v-if="touched.name && !isNameValid" class="field-error">Name is required.</span>
        </div>

        <div class="field-row-2">
          <div class="field">
            <label for="svc-duration">Duration (mins)</label>
            <input
              id="svc-duration"
              v-model.number="form.duration"
              type="number"
              min="1"
              placeholder="30"
              :class="{ invalid: touched.duration && !isDurationValid }"
              @blur="touched.duration = true"
            />
            <span v-if="touched.duration && !isDurationValid" class="field-error">Enter a valid duration.</span>
          </div>

          <div class="field">
            <label for="svc-price">Price</label>
            <input
              id="svc-price"
              v-model.number="form.price"
              type="number"
              min="0"
              step="0.01"
              placeholder="0.00"
              :class="{ invalid: touched.price && !isPriceValid }"
              @blur="touched.price = true"
            />
            <span v-if="touched.price && !isPriceValid" class="field-error">Enter a valid price.</span>
          </div>
        </div>

        <div class="field">
          <label for="svc-desc">Description <span class="optional">(optional)</span></label>
          <textarea id="svc-desc" v-model="form.description" rows="3" placeholder="Short note about this service" />
        </div>

        <p v-if="errorMessage" class="form-error" role="alert">{{ errorMessage }}</p>

        <div class="modal-actions">
          <button type="button" class="btn-secondary" @click="$emit('close')">Cancel</button>
          <button type="submit" class="btn-primary" :disabled="isSubmitting || !canSubmit">
            <span v-if="isSubmitting" class="spinner" aria-hidden="true"></span>
            {{ isSubmitting ? 'Saving…' : isEditMode ? 'Save changes' : 'Add service' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, computed, ref } from 'vue'
import { useServiceStore, type Service } from '@/store/service'

const props = defineProps<{
  service?: Service | null // pass an existing service to edit, omit/null to add
}>()

const emit = defineEmits<{
  close: []
  saved: []
}>()

const serviceStore = useServiceStore()
const isEditMode = computed(() => !!props.service)

const form = reactive({
  name: props.service?.name ?? '',
  duration: props.service?.duration_minutes ?? (null as number | null),
  price: props.service?.price ?? (null as number | null),
  description: props.service?.description ?? ''
})

const touched = reactive({ name: false, duration: false, price: false })
const isSubmitting = ref(false)
const errorMessage = ref('')

const isNameValid = computed(() => form.name.trim().length > 0)
const isDurationValid = computed(() => typeof form.duration === 'number' && form.duration > 0)
const isPriceValid = computed(() => typeof form.price === 'number' && form.price >= 0)
const canSubmit = computed(() => isNameValid.value && isDurationValid.value && isPriceValid.value)

const handleSubmit = async () => {
  touched.name = true
  touched.duration = true
  touched.price = true
  errorMessage.value = ''

  if (!canSubmit.value || form.duration === null || form.price === null) return

  const payload = {
    name: form.name,
    duration_minutes: form.duration,
    price: form.price,
    description: form.description
  }

  isSubmitting.value = true
  try {
    if (isEditMode.value && props.service) {
      await serviceStore.updateService(props.service.id, payload)
    } else {
      await serviceStore.addService(payload)
    }
    emit('saved')
    emit('close')
  } catch (err: any) {
    errorMessage.value = err?.response?.data?.message || 'Could not save this service. Please try again.'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Fraunces:opsz,wght@9..144,600&family=Inter:wght@400;500;600&display=swap');

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

.modal-card {
  background: #fff;
  border-radius: 12px;
  padding: 1.75rem;
  width: 100%;
  max-width: 440px;
  font-family: 'Inter', system-ui, sans-serif;
  box-shadow: 0 20px 48px rgba(20, 33, 61, 0.25);
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 1.25rem;
}

.modal-header h2 {
  font-family: 'Fraunces', serif;
  font-size: 1.3rem;
  font-weight: 600;
  color: #14213d;
  margin: 0;
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.4rem;
  line-height: 1;
  color: #5b6472;
  cursor: pointer;
  padding: 0.25rem;
}

.close-btn:hover { color: #14213d; }

.field { margin-bottom: 1.1rem; }

.field-row-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.85rem;
}

label {
  display: block;
  font-size: 0.85rem;
  font-weight: 500;
  color: #14213d;
  margin-bottom: 0.4rem;
}

.optional { font-weight: 400; color: #5b6472; }

input,
textarea {
  width: 100%;
  padding: 0.65rem 0.8rem;
  border: 1px solid #e4e2dc;
  border-radius: 8px;
  font-size: 0.92rem;
  font-family: inherit;
  background: #fff;
  box-sizing: border-box;
  resize: vertical;
}

input:focus,
textarea:focus {
  outline: none;
  border-color: #14213d;
  box-shadow: 0 0 0 3px rgba(20, 33, 61, 0.08);
}

input.invalid { border-color: #c0392b; }

.field-error {
  display: block;
  color: #c0392b;
  font-size: 0.78rem;
  margin-top: 0.35rem;
}

.form-error {
  background: #fdecea;
  color: #c0392b;
  padding: 0.6rem 0.8rem;
  border-radius: 8px;
  font-size: 0.85rem;
  margin: 0 0 1rem;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.6rem;
  margin-top: 1.4rem;
}

.btn-secondary,
.btn-primary {
  padding: 0.6rem 1.1rem;
  border-radius: 8px;
  font-size: 0.9rem;
  font-weight: 600;
  cursor: pointer;
  font-family: inherit;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.btn-secondary {
  background: #fff;
  border: 1px solid #e4e2dc;
  color: #14213d;
}

.btn-secondary:hover { background: #faf9f7; }

.btn-primary {
  background: #14213d;
  border: none;
  color: #f4f2ee;
}

.btn-primary:hover:not(:disabled) { background: #1c2d52; }
.btn-primary:disabled { opacity: 0.55; cursor: not-allowed; }

.spinner {
  width: 13px;
  height: 13px;
  border: 2px solid rgba(244, 242, 238, 0.4);
  border-top-color: #f4f2ee;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin { to { transform: rotate(360deg); } }

@media (max-width: 480px) {
  .field-row-2 { grid-template-columns: 1fr; }
}
</style>