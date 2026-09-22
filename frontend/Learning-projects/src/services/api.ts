import api from './axios'
import type { Service, ServicePayload } from '@/store/service'

const BASE = '/services'

export const serviceApi = {
  getAll() {
    return api.get<Service[]>(BASE)
  },
  create(payload: ServicePayload) {
    return api.post<Service>(BASE, payload)
  },
  update(id: Service['id'], payload: Partial<ServicePayload>) {
    return api.put<Service>(`${BASE}/${id}`, payload)
  },
  remove(id: Service['id']) {
    return api.delete<void>(`${BASE}/${id}`)
  }
}