import { ReceiptForm } from "@/components/receiptForm";
import { Card, CardBody, CardHeader, Input } from "@nextui-org/react";

export const ReceiptTable = () => {
  return (
    <Card className="p-6" id="receipt">
      <CardHeader className="text-lg font-semibold justify-between">
        Your Receipts <ReceiptForm />
      </CardHeader>
      <CardBody>
        <form className="flex flex-col gap-4">
          <Input required type="number" placeholder="Amount" startContent="$" />
          <Input
            required
            type="number"
            placeholder="GST/HST"
            startContent="$"
          />
          <Input required type="date" label="Date" />
        </form>
      </CardBody>
    </Card>
  );
};
