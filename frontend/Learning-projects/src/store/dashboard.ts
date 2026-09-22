import { defineStore } from "pinia";
import {
  getServices,
  getCustomers,
  getAppointmentsByDate,
} from "@/services/dashboardservices";

export const useDashboardStore = defineStore("dashboard", {
  state: () => ({
    totalServices: 0,
    totalCustomers: 0,
    todaysAppointments: 0,
    bookedAppointments: 0,

    isLoading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchDashboardData() {
      this.isLoading = true;
      this.error = null;

      try {
        // Get today's date in local time
        const today = new Date();

        const year = today.getFullYear();
        const month = String(today.getMonth() + 1).padStart(2, "0");
        const day = String(today.getDate()).padStart(2, "0");

        const date = `${year}-${month}-${day}`;

        const [services, customers, appointments] =
          await Promise.all([
            getServices(),
            getCustomers(),
            getAppointmentsByDate(date),
          ]);

        // Services
        this.totalServices = services.services?.length ?? 0;

        // Customers
        this.totalCustomers = customers.customers?.length ?? 0;

        // Appointments
        const appointmentList =
          appointments.appointments ?? [];

        this.todaysAppointments = appointmentList.length;

        this.bookedAppointments =
          appointmentList.filter(
            (appointment: any) =>
              appointment.status === "booked"
          ).length;
      } catch (error: any) {
        this.error =
          error?.response?.data?.error ||
          error?.response?.data?.message ||
          "Failed to load dashboard data";
      } finally {
        this.isLoading = false;
      }
    },
  },
});