"use client";

import { NextUIProvider } from "@nextui-org/react";
import { useRouter } from "next/navigation";
import { ClerkProvider } from "@clerk/clerk-react";

const FAKE_PUBLISHABLE_KEY =
  "pk_test_d2VsY29tZS1zdGFyZmlzaC02NS5jbGVyay5hY2NvdW50cy5kZXYk";

const PUBLISHABLE_KEY =
  process.env.NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY ??
  // HACK: use fake publishable key when there is no publishable key
  FAKE_PUBLISHABLE_KEY;

if (PUBLISHABLE_KEY === FAKE_PUBLISHABLE_KEY) {
  console.warn("using fake clerk publishable key");
}

export function Providers({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  return (
    <ClerkProvider publishableKey={PUBLISHABLE_KEY}>
      <NextUIProvider navigate={router.push}>{children}</NextUIProvider>
    </ClerkProvider>
  );
}
