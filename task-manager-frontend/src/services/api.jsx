import axios from 'axios'

const API = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
})


// Attach token to every request (if exists)
API.interceptors.request.use((req) => {
  const token = localStorage.getItem('token')
  if (token) {
    req.headers.Authorization = `Bearer ${token}`
  }
  req.headers['Content-Type'] = 'application/json'  // ✅ CRITICAL
  return req
})

export default API
