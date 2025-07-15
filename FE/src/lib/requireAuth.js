"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";

export default function requireAuth(Component) {
  return function Protected(props) {
    const router = useRouter();

    useEffect(() => {
      const token = localStorage.getItem("token");
      if (!token) {
        toast.error("Silakan login dulu");
        router.replace("/login");
      }
    }, []);

    return <Component {...props} />;
  };
}
