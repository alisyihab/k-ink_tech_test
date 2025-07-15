"use client";

import Sidebar from "@/components/Sidebar";
import requireAuth from "@/lib/requireAuth";

function DashboardLayout({ children }) {
  return (
    <div className="min-h-screen flex">
      <Sidebar />
      <main className="flex-1 p-8 bg-gray-50">{children}</main>
    </div>
  );
}

export default requireAuth(DashboardLayout);
