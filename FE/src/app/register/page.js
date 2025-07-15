'use client'

import RegisterFormWizard from "@/components/RegisterWizard"



export default function Register() {
  return (
    <main className="min-h-screen bg-gray-100 flex items-center justify-center px-4 py-10">
      <div className="w-full max-w-3xl">
        <RegisterFormWizard />
      </div>
    </main>
  )
}
