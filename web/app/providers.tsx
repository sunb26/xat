"use client";

import { NextUIProvider } from "@nextui-org/react";
import { useRouter } from "next/navigation";
import { ClerkProvider } from "@clerk/clerk-react";

const PUBLISHABLE_KEY = process.env.NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY;

if (!PUBLISHABLE_KEY) {
  throw new Error("Missing Publishable Key");
}

export function Providers({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  return (
    <ClerkProvider publishableKey={PUBLISHABLE_KEY}>
      <NextUIProvider navigate={router.push}>{children}</NextUIProvider>
    </ClerkProvider>
  );
}
