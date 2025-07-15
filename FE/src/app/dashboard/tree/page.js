"use client";

import dynamic from "next/dynamic";
import { useEffect, useRef, useState } from "react";

const Tree = dynamic(() => import("react-d3-tree").then((mod) => mod.Tree), {
  ssr: false,
});

const sponsorTree = [
  {
    name: "Andi",
    children: [
      {
        name: "Budi",
        children: [{ name: "Citra" }],
      },
      {
        name: "Dedi",
        children: [{ name: "Eka" }],
      },
    ],
  },
];

const uplineTree = {
  name: "Sinta",
  children: [
    {
      name: "Agus",
      children: [{ name: "Andi" }],
    },
    {
      name: "Rina",
      children: [{ name: "Eka" }],
    },
  ],
};

export default function TreeSponsorUpline() {
  const sponsorRef = useRef(null);
  const uplineRef = useRef(null);

  const [sponsorDim, setSponsorDim] = useState({ width: 600, height: 400 });
  const [uplineDim, setUplineDim] = useState({ width: 600, height: 400 });

  useEffect(() => {
    if (sponsorRef.current) {
      setSponsorDim({
        width: sponsorRef.current.offsetWidth,
        height: sponsorRef.current.offsetHeight,
      });
    }
    if (uplineRef.current) {
      setUplineDim({
        width: uplineRef.current.offsetWidth,
        height: uplineRef.current.offsetHeight,
      });
    }
  }, []);

  return (
    <div className="max-w-7xl mx-auto px-6 py-8">
      <h1 className="text-3xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-blue-600 to-indigo-600 mb-8 animate-fade-in-down">
        Struktur Jaringan
      </h1>
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Sponsor Tree */}
        <div
          className="bg-white/90 backdrop-blur-md rounded-2xl shadow-xl p-6 transition-all duration-300 hover:shadow-2xl"
          ref={sponsorRef}
          style={{ height: 450 }}
        >
          <h2 className="text-xl font-bold text-gray-800 mb-4">Tree Sponsor</h2>
          <Tree
            data={sponsorTree}
            orientation="vertical"
            translate={{ x: sponsorDim.width / 2, y: 50 }}
          />
        </div>

        {/* Upline Tree */}
        <div
          className="bg-white/90 backdrop-blur-md rounded-2xl shadow-xl p-6 transition-all duration-300 hover:shadow-2xl"
          ref={uplineRef}
          style={{ height: 450 }}
        >
          <h2 className="text-xl font-bold text-gray-800 mb-4">Tree Upline</h2>
          <Tree
            data={uplineTree}
            orientation="vertical"
            translate={{ x: uplineDim.width / 2, y: 50 }}
          />
        </div>
      </div>
    </div>
  );
}
