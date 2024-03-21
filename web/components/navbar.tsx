"use client";

import { Navbar, NavbarContent, NavbarItem } from "@nextui-org/react";
import {
  RedirectToSignIn,
  SignedIn,
  SignedOut,
  UserButton,
} from "@clerk/clerk-react";

export const NavBar = () => {
  return (
    <>
      <SignedIn>
        <Navbar>
          <NavbarContent justify="start">
            <NavbarItem>Your return</NavbarItem>
            <NavbarItem>Saved</NavbarItem>
          </NavbarContent>
          <NavbarContent justify="end">
            <NavbarItem>Help</NavbarItem>
            <NavbarItem>
              <UserButton />
            </NavbarItem>
          </NavbarContent>
        </Navbar>
      </SignedIn>
      <SignedOut>
        <RedirectToSignIn />
      </SignedOut>
    </>
  );
};
