"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { fetchMe, clearToken, getToken, User } from "@/lib/api";

export default function DashboardPage() {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = getToken();

    if (!token) {
      router.replace("/login");
      return;
    }

    fetchMe(token)
      .then(setUser)
      .catch((err) => {
        // Token invalid/expired → paksa login ulang
        clearToken();
        setError(err instanceof Error ? err.message : "Sesi berakhir");
        router.replace("/login");
      })
      .finally(() => setLoading(false));
  }, [router]);

  function handleLogout() {
    clearToken();
    router.push("/login");
  }

  if (loading) {
    return (
      <div className="page">
        <p style={{ color: "var(--muted)" }}>Memuat...</p>
      </div>
    );
  }

  if (!user) {
    return (
      <div className="page">
        <div className="card">
          <div className="error-box">{error || "Tidak ada sesi aktif"}</div>
        </div>
      </div>
    );
  }

  return (
    <div className="page">
      <div className="card dashboard">
        <h1>Halo, {user.name} 👋</h1>
        <p className="subtitle">Kamu berhasil login</p>

        <div className="user-row">
          <span>ID</span>
          <span>{user.id}</span>
        </div>
        <div className="user-row">
          <span>Nama</span>
          <span>{user.name}</span>
        </div>
        <div className="user-row">
          <span>Email</span>
          <span>{user.email}</span>
        </div>
        <div className="user-row">
          <span>Terdaftar sejak</span>
          <span>{new Date(user.created_at).toLocaleDateString("id-ID")}</span>
        </div>

        <div style={{ marginTop: 24 }}>
          <button className="btn btn-secondary" onClick={handleLogout}>
            Keluar
          </button>
        </div>
      </div>
    </div>
  );
}
