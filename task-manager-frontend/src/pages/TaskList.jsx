import { useEffect, useState } from 'react'
import API from '../services/api'
import { useNavigate } from 'react-router-dom'

function TaskList() {
  const [tasks, setTasks] = useState([])
  const [title, setTitle] = useState('')
  const navigate = useNavigate()

  // Fetch tasks
  const fetchTasks = async () => {
    try {
      const res = await API.get('/api/tasks/')
      console.log('📥 Tasks:', res.data)
      setTasks(res.data || []) // ⬅️ fallback if backend sends null
    } catch (err) {
      alert('Unauthorized or error fetching tasks')
      navigate('/login')
    }
  }
  

  

  useEffect(() => {
    fetchTasks()
  }, [])

  // Add task
  const handleAddTask = async (e) => {
    e.preventDefault()
  
    const task = {
      title: title,
      completed: false,
    }
  
    console.log("📤 Sending payload:", task)
  
    try {
      const res = await fetch("http://localhost:8080/api/tasks", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        body: JSON.stringify(task),
      })
  
      const data = await res.json()
      console.log("✅ Task added:", data)
  
      setTitle("")
      fetchTasks()
    } catch (err) {
      console.error("❌ Error adding task:", err)
      alert("Failed to add task")
    }
  }
  
  
  

  return (
    
    <div className="min-h-screen p-6 bg-gray-900 text-white">
      <h1 className="text-2xl font-bold mb-6">Your Tasks</h1>

      <form onSubmit={handleAddTask} className="mb-6 flex gap-4">
        <input
          type="text"
          placeholder="New task"
          className="p-2 rounded bg-gray-800 border border-gray-600 w-full"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
        <button type="submit" className="bg-green-600 hover:bg-green-700 px-4 py-2 rounded">
          Add
        </button>
        <button
    className="bg-red-600 hover:bg-red-700 px-4 py-1 rounded"
    onClick={() => {
      localStorage.removeItem("token")
      window.location.href = "/login"
    }}
  >
    Logout
  </button>
      </form>

      {Array.isArray(tasks) && tasks.length === 0 ? (
  <p>No tasks found</p>
) : (
  <ul className="space-y-2">
  {tasks.map((task) => (
    <li key={task.id} className="bg-gray-800 p-4 rounded-lg flex justify-between items-center">
      <div>
        <span>{task.title}</span>
        <span className="ml-2">{task.completed ? '✅' : '⏳'}</span>
      </div>
      <div className="space-x-2">
        <button
          className="bg-yellow-500 px-2 py-1 rounded text-sm"
          onClick={async () => {
            await API.put(`/api/tasks/toggle?id=${task.id}`)
            fetchTasks()
          }}
        >
          Toggle
        </button>
        <button
          className="bg-red-600 px-2 py-1 rounded text-sm"
          onClick={async () => {
            await API.delete(`/api/tasks/delete?id=${task.id}`)
            fetchTasks()
          }}
        >
          Delete
        </button>
      </div>
    </li>
  ))}
</ul>

)}

    </div>
  )
}

export default TaskList
