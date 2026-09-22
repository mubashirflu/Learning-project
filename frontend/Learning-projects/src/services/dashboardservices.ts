import api from "./axios";

const getServices = async () => {
  const response = await api.get("/services");

  return response.data;
};

const getCustomers = async () => {
  const response = await api.get("/customers");

  return response.data;
};

const getAppointmentsByCustomer = async (customerId: number) => {
  const response = await api.get(
    `/appointments/customer/${customerId}`,
  );

  return response.data;
};

const getAppointmentsByService = async (
  serviceId: number,
  date: string,
) => {
  const response = await api.get(
    `/appointments/service/${serviceId}?date=${date}`,
  );

  return response.data;
};
const getAppointmentsByDate = async (date: string) => {
  const response = await api.get(
    `/appointments/date/${date}`,
  );

  return response.data;
};

export {
  getServices,
  getCustomers,
  getAppointmentsByCustomer,
  getAppointmentsByService,
  getAppointmentsByDate
  
};

