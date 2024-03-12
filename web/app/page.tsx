import { Button, Input } from "@nextui-org/react";

export default function Home() {
  return (
    <main className="flex flex-col justify-between items-center p-24 min-h-screen">
      <div>
        <form className="flex flex-col gap-2">
          <Input type="text" label="Name" />
          <Input type="date" label="Date" />
          <Input type="number" label="Amount" />
          <Input type="number" label="GST/HST" />
          <Button color="primary">Submit</Button>
        </form>
      </div>
    </main>
  );
}
