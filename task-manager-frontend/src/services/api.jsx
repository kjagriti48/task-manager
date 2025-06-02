import axios from 'axios'

const API = axios.create({
  baseURL: 'http://localhost:8080',
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
