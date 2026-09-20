
import { defineStore } from "pinia";
import { login as loginRequest, register as registerRequest } from "../services/authserrvice";

interface User {
  id: number;
  name: string;
  email: string;
  created_at?: string;
}

interface RegisterData {
  name: string;
  email: string;
  password: string;
}

interface LoginData {
  email: string;
  password: string;
}

interface LoginResponse {
  message: string;
  token: string;
  user: User;
}

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: localStorage.getItem("token") as string | null,
    user: null as User | null,
    isLoading: false,
    error: null as string | null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
  },

  actions: {
    async register(data: RegisterData) {
      this.isLoading = true;
      this.error = null;

      try {
        const response = await registerRequest(data);

        return response;
      } catch (error: any) {
        this.error =
          error?.response?.data?.error ||
          error?.response?.data?.message ||
          "Registration failed";

        throw error;
      } finally {
        this.isLoading = false;
      }
    },

    async login(data: LoginData) {
      this.isLoading = true;
      this.error = null;

      try {
        const response: LoginResponse = await loginRequest(data);

        this.token = response.token;
        this.user = response.user;

        localStorage.setItem("token", response.token);

        return response;
      } catch (error: any) {
        this.error =
          error?.response?.data?.error ||
          error?.response?.data?.message ||
          "Login failed";

        throw error;
      } finally {
        this.isLoading = false;
      }
    },

    logout() {
      this.token = null;
      this.user = null;
      this.error = null;

      localStorage.removeItem("token");
    },

    clearError() {
      this.error = null;
    },
  },
});

