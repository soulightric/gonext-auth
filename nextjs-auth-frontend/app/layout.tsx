import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Fiber Auth Demo",
  description: "Contoh integrasi Next.js dengan backend Golang Fiber",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="id">
      <body>{children}</body>
    </html>
  );
}
