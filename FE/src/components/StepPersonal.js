"use client";

import { Field, ErrorMessage } from "formik";

export default function StepPersonal() {
  return (
    <div className="space-y-5">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          ID Member
        </label>
        <Field
          name="id_member"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Contoh: MBR123"
        />
        <ErrorMessage
          name="id_member"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Nama Lengkap
        </label>
        <Field
          name="name"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage
          name="name"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Jenis Kelamin
        </label>
        <Field
          as="select"
          name="gender"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="">Pilih</option>
          <option value="L">Laki-laki</option>
          <option value="P">Perempuan</option>
        </Field>
        <ErrorMessage
          name="gender"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Tempat Lahir
        </label>
        <Field
          name="birth_place"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage
          name="birth_place"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Tanggal Lahir
        </label>
        <Field
          type="date"
          name="birth_date"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage
          name="birth_date"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Nomor HP
        </label>
        <Field
          name="phone_number"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage
          name="phone_number"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          Email
        </label>
        <Field
          type="email"
          name="email"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage
          name="email"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">
          No KTP
        </label>
        <Field
          name="no_ktp"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <ErrorMessage
          name="no_ktp"
          component="div"
          className="text-sm text-red-500 mt-1"
        />
      </div>
    </div>
  );
}
