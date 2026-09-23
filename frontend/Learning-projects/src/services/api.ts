// import api from './axios'
// import type { Service, ServicePayload } from '@/store/service'

// const BASE = '/services'

// export const serviceApi = {
//   getAll() {
//     return api.get<Service[]>(BASE)
//   },
//   create(payload: ServicePayload) {
//     return api.post<Service>(BASE, payload)
//   },
//   update(id: Service['id'], payload: Partial<ServicePayload>) {
//     return api.put<Service>(`${BASE}/${id}`, payload)
//   },
//   remove(id: Service['id']) {
//     return api.delete<void>(`${BASE}/${id}`)
//   }
// }
import api from './axios'
import type { Service, ServicePayload } from '@/store/service'

const BASE = '/services'

// Response shapes match what the Go backend actually sends (wrapped, not raw)
interface ServiceListResponse {
  services: Service[]
}

interface ServiceItemResponse {
  service: Service
}

export const serviceApi = {
  getAll() {
    return api.get<ServiceListResponse>(BASE)
  },
  create(payload: ServicePayload) {
    return api.post<ServiceItemResponse>(BASE, payload)
  },
  update(id: Service['id'], payload: Partial<ServicePayload>) {
    return api.put<ServiceItemResponse>(`${BASE}/${id}`, payload)
  },
  remove(id: Service['id']) {
    return api.delete<void>(`${BASE}/${id}`)
  }
}