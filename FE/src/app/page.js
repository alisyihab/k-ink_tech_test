// src/app/page.tsx
import Link from "next/link";

export default function LandingPage() {
  return (
    <main className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-blue-100 via-white to-blue-50 px-4">
      <div className="max-w-2xl text-center">
        <h1 className="text-4xl font-bold text-blue-700 mb-4">
          Selamat Datang di Program Smart Reward
        </h1>
        <p className="text-lg text-gray-700 mb-6">
          Dapatkan kesempatan luar biasa untuk meraih reward dengan bergabung dalam sistem bisnis terpercaya kami.
        </p>
        <Link
          href="/register"
          className="inline-block bg-blue-600 text-white px-6 py-3 rounded-lg shadow hover:bg-blue-700 transition"
        >
          Daftar Sekarang
        </Link>
      </div>

      <div className="mt-12 text-center text-sm text-gray-400">
        &copy; {new Date().getFullYear()} Multiverse Marketing
      </div>
    </main>
  );
}
