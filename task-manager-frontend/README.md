# ✅ Task Manager (Go + React + JWT)

A full-stack Task Manager app built with **Golang** and **React** featuring:

- 🔐 JWT authentication
- 🧾 Task CRUD (Create, Toggle, Delete)
- 🌐 RESTful API
- 💻 Vite + Tailwind frontend
- 🧠 In-memory backend (map store)
- ☁️ Deployable to Railway + Netlify

---

## 🛠 Tech Stack

- **Frontend:** React (Vite), Tailwind CSS
- **Backend:** Golang, net/http
- **Auth:** JWT
- **State:** In-memory map
- **Deployment:** Railway (Go), Netlify (React)

---

## 🧪 Features

- ✅ Sign Up / Log In with JWT
- ✅ Create Task
- ✅ Toggle Complete
- ✅ Delete Task
- 🔐 Logout
- 🛡 Auth-guarded routes
- ⚙️ CORS & preflight-safe

---

## 📦 Running Locally

### Backend (Go)

```bash
cd task-manager
go run main.go

Runs at: http://localhost:8080

Frontend (React)

cd task-manager-frontend
npm install
npm run dev

Folder Structure

task-manager/
├── handlers/
├── models/
├── routes/
└── main.go

task-manager-frontend/
├── src/
│   ├── pages/
│   ├── services/
│   └── App.jsx
🚀 Deployment Plan
Railway:

Push Go code

Add JWT_SECRET as env var

Netlify:

Push frontend repo

Set base as /task-manager-frontend

Add VITE_API_URL=http://<railway-url> in environment

🙋‍♀️ Author
Made with 🧠 by Jagriti 💻
GitHub