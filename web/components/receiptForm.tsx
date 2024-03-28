import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerDescription,
  DrawerFooter,
  DrawerHeader,
  DrawerTitle,
  DrawerTrigger,
} from "@/components/ui/drawer";
import { Button, Input } from "@nextui-org/react";
import { Camera } from "@/components/camera";

export const ReceiptForm = () => {
  return (
    <Drawer>
      <DrawerTrigger asChild>
        <Button>Add receipt</Button>
      </DrawerTrigger>
      <DrawerContent>
        <div className="mx-auto w-full max-w-md">
          <DrawerHeader>
            <DrawerTitle>
              <div className="flex items-center justify-between">
                Add a Receipt <Camera />
              </div>
            </DrawerTitle>
            <DrawerDescription>
              Fill in the form or take a picture
            </DrawerDescription>
          </DrawerHeader>
          <div className="p-4">
            <form className="flex flex-col gap-4">
              <Input required type="number" label="Subtotal" startContent="$" />
              <Input required type="number" label="GST/HST" startContent="$" />
              <Input type="number" label="Gratuity" startContent="$" />
              <Input required type="date" label="Date" />
              <Input disabled type="number" label="Total" />
            </form>
          </div>
          <DrawerFooter>
            <Button color="primary">Submit</Button>
            <DrawerClose asChild>
              <Button variant="bordered">Cancel</Button>
            </DrawerClose>
          </DrawerFooter>
        </div>
      </DrawerContent>
    </Drawer>
  );
};
