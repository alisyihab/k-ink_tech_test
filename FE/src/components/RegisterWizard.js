"use client";

import { Formik, Form } from "formik";
import * as Yup from "yup";
import toast from "react-hot-toast";
import { useState } from "react";

import StepPersonal from "@/components/StepPersonal";
import StepBank from "@/components/StepBank";
import StepPaket from "@/components/StepPaket";
import StepSponsorUpline from "@/components/StepSponsorUpline";
import StepAlamat from "@/components/StepAlamat";
import http from "@/lib/http";

const steps = [
  { label: "Data Pribadi", component: StepPersonal },
  { label: "Info Bank", component: StepBank },
  { label: "Pilih Paket", component: StepPaket },
  { label: "Sponsor & Upline", component: StepSponsorUpline },
  { label: "Alamat", component: StepAlamat },
];

const initialValues = {
  id_member: "",
  name: "",
  gender: "",
  birth_place: "",
  birth_date: "",
  phone_number: "",
  email: "",
  no_ktp: "",
  bank_name: "",
  bank_account_no: "",
  sponsor_id_member: "",
  sponsor_name: "",
  upline_id_member: "",
  upline_name: "",
  paket_id: "",
  alamat_lengkap: "",
  manager_area: "",
  kirim_ke_stockist: false,
  tempat_daftar: "WEB",
};

const validationSchemas = [
  // Step 0: Data Pribadi
  Yup.object({
    id_member: Yup.string().required("ID Member wajib diisi"),
    name: Yup.string().required("Nama wajib diisi"),
    gender: Yup.string().required("Jenis kelamin wajib diisi"),
    birth_place: Yup.string().required("Tempat lahir wajib diisi"),
    birth_date: Yup.string().required("Tanggal lahir wajib diisi"),
    phone_number: Yup.string().required("No HP wajib diisi"),
    email: Yup.string()
      .email("Email tidak valid")
      .required("Email wajib diisi"),
    no_ktp: Yup.string().required("No KTP wajib diisi"),
  }),

  // Step 1: Info Bank
  Yup.object({
    bank_name: Yup.string().required("Nama bank wajib diisi"),
    bank_account_no: Yup.string().required("No rekening wajib diisi"),
  }),

  // Step 2: Pilih Paket
  Yup.object({
    paket_id: Yup.string().required("Paket wajib dipilih"),
  }),

  // Step 3: Sponsor & Upline (opsional)
  Yup.object({}),

  // Step 4: Alamat
  Yup.object({
    alamat_lengkap: Yup.string().required("Alamat wajib diisi"),
  }),
];

export default function RegisterFormWizard() {
  const [step, setStep] = useState(0);
  const StepComponent = steps[step].component;

  const isLastStep = step === steps.length - 1;

  const handleSubmit = async (values, actions) => {
    try {
      const res = await http.post("/member", values);
      toast.success("Registrasi berhasil!");
      actions.resetForm();
      setStep(0);
    } catch (err) {
      const errors = err.response?.data?.errors;
      if (errors) {
        actions.setErrors(errors);
      }
      toast.error(err.response?.data?.message || "Gagal registrasi");
    } finally {
      actions.setSubmitting(false);
    }
  };

  return (
    <div className="max-w-2xl mx-auto bg-white/90 backdrop-blur-md p-8 md:p-12 rounded-2xl shadow-xl mt-12 transition-all duration-300 hover:shadow-2xl">
      <h2 className="text-2xl md:text-3xl font-bold text-gray-800 mb-6 animate-fade-in-down">
        {steps[step].label}
      </h2>

      <Formik
        initialValues={initialValues}
        validationSchema={validationSchemas[step]}
        onSubmit={async (values, actions) => {
          if (isLastStep) {
            await handleSubmit(values, actions);
          } else {
            const errors = await actions.validateForm();

            const currentFields = Object.keys(validationSchemas[step].fields);
            const hasError = currentFields.some((field) => errors[field]);

            if (hasError) {
              actions.setTouched(
                currentFields.reduce((acc, field) => {
                  acc[field] = true;
                  return acc;
                }, {})
              );
              return;
            }

            setStep((s) => s + 1);
            actions.setTouched({});
          }
        }}
      >
        {({ isSubmitting }) => (
          <Form className="space-y-6">
            <StepComponent />

            <div className="flex justify-between mt-8 gap-4">
              {step > 0 && (
                <button
                  type="button"
                  onClick={() => setStep((s) => s - 1)}
                  className="px-6 py-3 bg-gray-200 text-gray-700 rounded-full font-semibold hover:bg-gray-300 transition-all duration-200"
                >
                  Kembali
                </button>
              )}
              <button
                type="submit"
                className="px-6 py-3 bg-gradient-to-r from-blue-600 to-indigo-600 text-white rounded-full font-semibold hover:from-blue-700 hover:to-indigo-700 hover:scale-105 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
                disabled={isSubmitting}
              >
                {isLastStep ? "Daftar Sekarang" : "Lanjut"}
              </button>
            </div>
          </Form>
        )}
      </Formik>
    </div>
  );
}
