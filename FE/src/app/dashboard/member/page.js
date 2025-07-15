"use client";

import { useEffect, useState } from "react";
import http from "@/lib/http";
import useDebounce from "@/hooks/debounce";

export default function MemberPage() {
  const [members, setMembers] = useState([]);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [search, setSearch] = useState("");
  const debouncedSearch = useDebounce(search, 500);

  const fetchData = async () => {
    try {
      const res = await http.get(
        `/api/members?page=${page}&search=${debouncedSearch}`
      );
      setMembers(res.data.data);
      setTotalPages(res.data.total_pages || 1);
    } catch (err) {
      console.error("Gagal ambil data member:", err);
    }
  };

  useEffect(() => {
    fetchData();
  }, [page, debouncedSearch]);

  return (
    <div className="max-w-6xl mx-auto px-6 py-8">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 animate-fade-in-down">
          Daftar Member
        </h1>
        <input
          type="text"
          placeholder="Cari ID/Nama/Email/HP"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-full max-w-xs px-4 py-3 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all duration-200"
        />
      </div>

      <div className="overflow-auto bg-white/90 backdrop-blur-md rounded-2xl shadow-xl">
        <table className="w-full text-sm text-left">
          <thead className="bg-gray-50 text-gray-600">
            <tr>
              <th className="px-6 py-4 font-semibold">ID Member</th>
              <th className="px-6 py-4 font-semibold">Nama</th>
              <th className="px-6 py-4 font-semibold">Email</th>
              <th className="px-6 py-4 font-semibold">No HP</th>
              <th className="px-6 py-4 font-semibold">Paket</th>
            </tr>
          </thead>
          <tbody>
            {members.length === 0 ? (
              <tr>
                <td colSpan={5} className="px-6 py-4 text-center text-gray-500">
                  Belum ada data member.
                </td>
              </tr>
            ) : (
              members.map((m) => (
                <tr
                  key={m.id_member}
                  className="border-t border-gray-200 hover:bg-gray-50 transition-all duration-200"
                >
                  <td className="px-6 py-4">{m.id_member}</td>
                  <td className="px-6 py-4">{m.name}</td>
                  <td className="px-6 py-4">{m.email}</td>
                  <td className="px-6 py-4">{m.phone_number}</td>
                  <td className="px-6 py-4">{m.paket_snapshot?.nama || "-"}</td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <div className="flex justify-center items-center gap-4 mt-8">
        <button
          disabled={page === 1}
          onClick={() => setPage((p) => p - 1)}
          className="px-4 py-2 bg-gray-200 text-gray-700 rounded-full font-semibold hover:bg-gray-300 transition-all duration-200 disabled:opacity-40 disabled:cursor-not-allowed"
        >
          Prev
        </button>
        <span className="text-sm font-medium text-gray-600">
          Page {page} of {totalPages}
        </span>
        <button
          disabled={page === totalPages}
          onClick={() => setPage((p) => p + 1)}
          className="px-4 py-2 bg-gray-200 text-gray-700 rounded-full font-semibold hover:bg-gray-300 transition-all duration-200 disabled:opacity-40 disabled:cursor-not-allowed"
        >
          Next
        </button>
      </div>
    </div>
  );
}
