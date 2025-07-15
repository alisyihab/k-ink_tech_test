'use client'

import { Field, ErrorMessage } from 'formik'

export default function StepBank() {
  return (
    <div className="space-y-5">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">Nama Bank</label>
        <Field
          name="bank_name"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Contoh: BCA, BRI, Mandiri"
        />
        <ErrorMessage name="bank_name" component="div" className="text-sm text-red-500 mt-1" />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">Nomor Rekening</label>
        <Field
          name="bank_account_no"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Contoh: 1234567890"
        />
        <ErrorMessage name="bank_account_no" component="div" className="text-sm text-red-500 mt-1" />
      </div>
    </div>
  )
}
