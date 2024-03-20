import { Card, CardBody, CardHeader, Input } from "@nextui-org/react";

export const MailingForm = () => {
  return (
    <Card className="p-6" id="mailing">
      <CardHeader className="text-lg font-semibold">Mailing Address</CardHeader>
      <CardBody>
        <form className="flex flex-col gap-4">
          <Input required type="text" placeholder="Address Line 1" />
          <div className="flex gap-2">
            <Input required type="text" placeholder="City" />
            <Input required type="text" placeholder="Province/Territory" />
            <Input required type="text" placeholder="Postal Code" />
          </div>
        </form>
      </CardBody>
    </Card>
  );
};
