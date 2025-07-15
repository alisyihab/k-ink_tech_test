"use client";

import { useEffect, useState } from "react";
import http from "@/lib/http";
import toast from "react-hot-toast";

export default function PaketPage() {
  const [pakets, setPakets] = useState([]);
  const [open, setOpen] = useState(false);
  const [form, setForm] = useState({
    wilayah: "",
    tipe: "",
    level: "",
    nama: "",
    harga: "",
  });

  /* --------------------------- fetch paket list -------------------------- */
  const loadData = async () => {
    const res = await http.get("/paket");
    setPakets(res.data.data || []);
  };
  useEffect(() => {
    loadData();
  }, []);

  /* ------------------------------ handlers ------------------------------- */
  const handleChange = (e) =>
    setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await http.post("/api/paket", { ...form, harga: Number(form.harga) });
      toast.success("Paket berhasil ditambah");
      setOpen(false);
      setForm({ wilayah: "", tipe: "", level: "", nama: "", harga: "" });
      loadData();
    } catch (err) {
      toast.error(err.response?.data?.message || "Gagal menambah paket");
    }
  };

  /* ------------------------------- render ------------------------------- */
  return (
    <div className="max-w-5xl mx-auto p-8">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-3xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 animate-fade-in-down">
          Daftar Paket
        </h1>
        <button
          onClick={() => setOpen(true)}
          className="px-6 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 text-white rounded-full font-semibold hover:from-blue-700 hover:to-indigo-700 hover:scale-105 transition-all duration-200"
        >
          Tambah Paket
        </button>
      </div>

      {/* table list */}
      <div className="overflow-x-auto bg-white/90 backdrop-blur-md rounded-2xl shadow-xl">
        <table className="w-full text-sm">
          <thead className="bg-gray-50 text-gray-600">
            <tr>
              <th className="px-6 py-4 text-left font-semibold">Nama</th>
              <th className="px-6 py-4 text-left font-semibold">Level</th>
              <th className="px-6 py-4 text-left font-semibold">Wilayah</th>
              <th className="px-6 py-4 text-left font-semibold">Harga</th>
            </tr>
          </thead>
          <tbody>
            {pakets.map((p) => (
              <tr
                key={p.id}
                className="border-t border-gray-200 hover:bg-gray-50 transition-all duration-200"
              >
                <td className="px-6 py-4">{p.nama}</td>
                <td className="px-6 py-4">{p.level}</td>
                <td className="px-6 py-4">{p.wilayah}</td>
                <td className="px-6 py-4">
                  Rp {p.harga.toLocaleString("id-ID")}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {/* ---------- Modal ---------- */}
      {open && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
          <div className="bg-white w-full max-w-lg rounded-2xl p-8 shadow-2xl animate-fade-in-up">
            <h2 className="text-2xl font-bold text-gray-800 mb-6">
              Tambah Paket
            </h2>

            <form onSubmit={handleSubmit} className="space-y-4">
              {["wilayah", "tipe", "level", "nama", "harga"].map((f) => (
                <div key={f}>
                  <input
                    name={f}
                    value={form[f]}
                    onChange={handleChange}
                    placeholder={f.charAt(0).toUpperCase() + f.slice(1)}
                    type={f === "harga" ? "number" : "text"}
                    className="w-full px-4 py-3 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all duration-200"
                    required
                  />
                </div>
              ))}

              <div className="flex justify-end gap-3 pt-6">
                <button
                  type="button"
                  className="px-6 py-3 bg-gray-200 text-gray-700 rounded-full font-semibold hover:bg-gray-300 transition-all duration-200"
                  onClick={() => setOpen(false)}
                >
                  Batal
                </button>
                <button
                  type="submit"
                  className="px-6 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 text-white rounded-full font-semibold hover:from-blue-700 hover:to-indigo-700 hover:scale-105 transition-all duration-200"
                >
                  Simpan
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
