// import { defineStore } from 'pinia'
// import { serviceApi } from '@/services/api'

// export interface Service {
//   id: string | number
//   name: string
//   duration_minutes: number
//   price: number
//   description?: string
// }

// export type ServicePayload = Omit<Service, 'id'>

// interface ServiceState {
//   services: Service[]
//   isLoading: boolean
//   error: string | null
// }

// function extractErrorMessage(err: unknown, fallback: string): string {
//   if (err && typeof err === 'object' && 'response' in err) {
//     const response = (err as { response?: { data?: { message?: string } } }).response
//     if (response?.data?.message) return response.data.message
//   }
//   return fallback
// }

// export const useServiceStore = defineStore('service', {
//   state: (): ServiceState => ({
//     services: [],
//     isLoading: false,
//     error: null
//   }),

//   getters: {
//     totalServices: (state): number => state.services.length
//   },

//   actions: {
//     async fetchServices() {
//       this.isLoading = true
//       this.error = null
//       try {
//         const { data } = await serviceApi.getAll()
//         this.services = data
//       } catch (err) {
//         this.error = extractErrorMessage(err, 'Could not load services.')
//         throw err
//       } finally {
//         this.isLoading = false
//       }
//     },

//     async addService(payload: ServicePayload) {
//       const { data } = await serviceApi.create(payload)
//       this.services.push(data.service)
//       return data.service
//     },

//     async updateService(id: Service['id'], payload: Partial<ServicePayload>) {
//       const { data } = await serviceApi.update(id, payload)
//       const index = this.services.findIndex((s) => s.id === id)
//       if (index !== -1) this.services[index] = data
//       return data
//     },

//     async deleteService(id: Service['id']) {
//       await serviceApi.remove(id)
//       this.services = this.services.filter((s) => s.id !== id)
//     }
//   }
// })

import { defineStore } from 'pinia'
import { serviceApi } from '@/services/api'

export interface Service {
  id: string | number
  name: string
  duration_minutes: number
  price: number
  description?: string
}

export type ServicePayload = Omit<Service, 'id'>

interface ServiceState {
  services: Service[]
  isLoading: boolean
  error: string | null
}

function extractErrorMessage(err: unknown, fallback: string): string {
  if (err && typeof err === 'object' && 'response' in err) {
    const response = (err as { response?: { data?: { message?: string } } }).response
    if (response?.data?.message) return response.data.message
  }
  return fallback
}

export const useServiceStore = defineStore('service', {
  state: (): ServiceState => ({
    services: [],
    isLoading: false,
    error: null
  }),

  getters: {
    totalServices: (state): number => state.services.length
  },

  actions: {
    async fetchServices() {
      this.isLoading = true
      this.error = null
      try {
        const { data } = await serviceApi.getAll()
        // Backend wraps the array: { "services": [...] }
        this.services = data.services
      } catch (err) {
        this.error = extractErrorMessage(err, 'Could not load services.')
        throw err
      } finally {
        this.isLoading = false
      }
    },

    async addService(payload: ServicePayload) {
      const { data } = await serviceApi.create(payload)
      // Backend wraps the object: { "service": {...} }
      const created = data.service
      this.services.push(created)
      return created
    },

    async updateService(id: Service['id'], payload: Partial<ServicePayload>) {
      const { data } = await serviceApi.update(id, payload)
      const updated = data.service
      const index = this.services.findIndex((s) => s.id === id)
      if (index !== -1) this.services[index] = updated
      return updated
    },

    async deleteService(id: Service['id']) {
      await serviceApi.remove(id)
      this.services = this.services.filter((s) => s.id !== id)
    }
  }
})