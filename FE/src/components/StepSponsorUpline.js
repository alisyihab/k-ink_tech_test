'use client'

import { Field, ErrorMessage } from 'formik'

export default function StepSponsorUpline() {
  return (
    <div className="space-y-5">
      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">ID Member Sponsor</label>
        <Field
          name="sponsor_id_member"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Opsional"
        />
        <ErrorMessage name="sponsor_id_member" component="div" className="text-sm text-red-500 mt-1" />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">Nama Sponsor</label>
        <Field
          name="sponsor_name"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Opsional"
        />
        <ErrorMessage name="sponsor_name" component="div" className="text-sm text-red-500 mt-1" />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">ID Member Upline</label>
        <Field
          name="upline_id_member"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Opsional"
        />
        <ErrorMessage name="upline_id_member" component="div" className="text-sm text-red-500 mt-1" />
      </div>

      <div>
        <label className="block text-sm font-medium text-gray-700 mb-1">Nama Upline</label>
        <Field
          name="upline_name"
          className="w-full border border-gray-300 rounded px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Opsional"
        />
        <ErrorMessage name="upline_name" component="div" className="text-sm text-red-500 mt-1" />
      </div>
    </div>
  )
}
