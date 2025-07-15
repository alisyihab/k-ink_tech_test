'use client'
import { useEffect, useState } from 'react'
import http from '@/lib/http'

export default function DashboardPage() {
  const [data, setData] = useState({ total_member: 0, total_paket: 0 })

  useEffect(() => {
    const fetchData = async () => {
      const res = await http.get('/api/dashboard')
      setData(res.data)
    }

    fetchData()
  }, [])

  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 gap-6 p-6">
      <div className="bg-white shadow p-6 rounded-xl">
        <h3 className="text-gray-700 text-lg font-medium">Total Member</h3>
        <p className="text-3xl font-bold text-blue-600">{data.total_member}</p>
      </div>
      <div className="bg-white shadow p-6 rounded-xl">
        <h3 className="text-gray-700 text-lg font-medium">Total Paket</h3>
        <p className="text-3xl font-bold text-green-600">{data.total_paket}</p>
      </div>
    </div>
  )
}
