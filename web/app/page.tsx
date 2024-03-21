import { Camera } from "@/components/camera";
import { MailingForm } from "@/components/mailing";
import { ReceiptForm } from "@/components/receiptForm";
import { UserInfo } from "@/components/userInfo";
import { Divider, Link } from "@nextui-org/react";

export default function Home() {
  return (
    <section className="flex flex-col items-center">
      <div className="flex justify-between p-6 mt-8 w-full max-w-4xl">
        <div className="hidden flex-col gap-4 pr-4 min-w-48 md:flex fixed">
          <h1>2024 Tax Return</h1>
          <Divider />
          <Link isBlock color="foreground" href="#userinfo">
            About you
          </Link>
          <Link isBlock color="foreground" href="#mailing">
            Mailing Address
          </Link>
          <Link isBlock color="foreground" href="#receipt">
            Add Receipts
          </Link>
          <Link isBlock color="foreground" href="#camera">
            Upload a Receipt
          </Link>
          <Link isBlock color="foreground">
            Summary
          </Link>
          <Link isBlock color="foreground">
            Submit
          </Link>
        </div>
        <div className="flex flex-col gap-8 w-full md:ml-48">
          <h1 className="text-2xl">Your 2024 Tax Return</h1>
          <UserInfo />
          <MailingForm />
          <ReceiptForm />
          <Camera />
        </div>
      </div>
    </section>
  );
}
