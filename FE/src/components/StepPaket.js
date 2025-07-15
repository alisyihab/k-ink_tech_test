"use client";

import { useEffect, useState } from "react";
import { Field, ErrorMessage } from "formik";
import http from "@/lib/http";

export default function StepPaket() {
  const [paketList, setPaketList] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchPaket = async () => {
      try {
        const res = await http.get("/paket");
        if (Array.isArray(res.data.data)) {
          setPaketList(res.data.data);
        } else {
          setError("Data paket tidak valid");
        }
      } catch (err) {
        console.error("Gagal ambil paket:", err);
        setError("Gagal memuat daftar paket");
      } finally {
        setLoading(false);
      }
    };

    fetchPaket();
  }, []);

  return (
    <div className="space-y-5">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Pilih Paket
        </label>

        {loading ? (
          <div className="text-gray-500 text-sm">Memuat paket...</div>
        ) : error ? (
          <div className="text-red-500 text-sm">{error}</div>
        ) : paketList.length === 0 ? (
          <div className="text-gray-500 text-sm">Belum ada paket tersedia</div>
        ) : (
          <Field
            as="select"
            name="paket_id"
            className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="">-- Pilih Paket --</option>
            {paketList.map((paket) => (
              <option key={paket.id} value={paket.id}>
                {paket.nama ?? "Tanpa Nama"} - {paket.level ?? "Unknown"} (
                {paket.wilayah ?? "-"}) - Rp{" "}
                {paket.harga?.toLocaleString() ?? "0"}
              </option>
            ))}
          </Field>
        )}

        <ErrorMessage
          name="paket_id"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>
    </div>
  );
}
