type LoginData = { email: string; password: string };
type RegisterData = { name: string; email: string; password: string };

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api'; // Update with your backend URL

export const register = async (data: RegisterData): Promise<{ success: boolean; message: string }> => {
  const response = await fetch(`${API_URL}/user/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(data),
  });
  return response.json();
};

export const login = async (data: LoginData): Promise<{ success: boolean; message: string }> => {
  const response = await fetch(`${API_URL}/user/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(data),
  });
  return response.json();
};
