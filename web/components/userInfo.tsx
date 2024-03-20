import { Card, CardBody, CardHeader, Input } from "@nextui-org/react";

export const UserInfo = () => {
  return (
    <Card className="p-6" id="userinfo">
      <CardHeader className="text-lg font-semibold">About You</CardHeader>
      <CardBody>
        <form className="flex flex-col gap-4">
          <h1>Name</h1>
          <div className="flex gap-2">
            <Input required type="text" placeholder="First name" />
            <Input type="text" placeholder="Middle name" />
            <Input required type="text" placeholder="Last name" />
          </div>
          <h1>Social Insurance Number</h1>
          <Input required type="number" placeholder="SIN" />
          <h1>Date of Birth</h1>
          <Input required type="date" />
        </form>
      </CardBody>
    </Card>
  );
};
