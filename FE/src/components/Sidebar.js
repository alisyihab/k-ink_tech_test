"use client";

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import toast from "react-hot-toast";

const menus = [
  { label: "Dashboard", href: "/dashboard" },
  { label: "Member", href: "/dashboard/member" },
  { label: "Paket", href: "/dashboard/paket" },
  { label: "Tree", href: "/dashboard/tree" },
];

export default function Sidebar() {
  const pathname = usePathname();
  const router = useRouter();

  const handleLogout = () => {
    localStorage.removeItem("token");
    toast.success("Logout berhasil");
    router.push("/login");
  };

  return (
    <aside className="w-64 bg-white p-8 border-r h-screen flex flex-col shadow-lg transition-all duration-300">
      <h1 className="text-2xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 mb-8 animate-fade-in-down">
        Smart Reward
      </h1>

      <nav className="space-y-3 flex-1">
        {menus.map((m) => (
          <Link
            key={m.href}
            href={m.href}
            className={`block px-4 py-3 rounded-lg text-lg font-medium transition-all duration-200 ${
              pathname === m.href
                ? "bg-gradient-to-r from-blue-600 to-indigo-600 text-white shadow-md"
                : "text-gray-600 hover:bg-gray-100 hover:text-blue-600"
            }`}
          >
            {m.label}
          </Link>
        ))}
      </nav>

      <button
        onClick={handleLogout}
        className="mt-6 px-4 py-3 bg-gradient-to-r from-red-600 to-red-700 text-white rounded-lg font-semibold hover:from-red-700 hover:to-red-800 hover:scale-105 transition-all duration-200"
      >
        Logout
      </button>
    </aside>
  );
}
