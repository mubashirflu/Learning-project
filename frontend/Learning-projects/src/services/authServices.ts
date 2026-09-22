import api from "./axios";

interface RegisterData {
  name: string;
  email: string;
  password: string;
}

interface LoginData {
  email: string;
  password: string;
}

const register = async (data: RegisterData) => {
  const response = await api.post("/../auth/register", data);

  return response.data;
};

const login = async (data: LoginData) => {
  const response = await api.post("/../auth/login", data);

  return response.data;
};

export { register, login };
