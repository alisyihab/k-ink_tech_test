"use client";

import { Field, ErrorMessage } from "formik";

export default function StepAlamat() {
  return (
    <div className="space-y-5">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Alamat Lengkap
        </label>
        <Field
          as="textarea"
          name="alamat_lengkap"
          rows="3"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Contoh: Jl. Mawar No. 10, Jakarta Timur"
        />
        <ErrorMessage
          name="alamat_lengkap"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Manager Area
        </label>
        <Field
          name="manager_area"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Opsional"
        />
        <ErrorMessage
          name="manager_area"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div className="flex items-center space-x-2">
        <Field type="checkbox" name="kirim_ke_stockist" className="h-4 w-4" />
        <label className="text-sm text-gray-700">Kirim ke Stockist</label>
      </div>
    </div>
  );
}
