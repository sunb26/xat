import { Camera } from "@/components/camera";
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
import { Button, Input, Tooltip } from "@nextui-org/react";
import { CircleHelp } from "lucide-react";

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
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <label className="text-sm">Subtotal</label>
                  <Tooltip content="The pre-tax ammount">
                    <CircleHelp className="w-4 h-4 ml-2" />
                  </Tooltip>
                </div>
                <Input
                  required
                  type="number"
                  startContent="$"
                  placeholder="0.00"
                  className="max-w-52"
                />
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <label className="text-sm">GST/HST</label>
                  <Tooltip content="Goods and services tax/harmonized sales tax">
                    <CircleHelp className="w-4 h-4 ml-2" />
                  </Tooltip>
                </div>
                <Input
                  required
                  type="number"
                  startContent="$"
                  placeholder="0.00"
                  className="max-w-52"
                />
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <label className="text-sm">Gratuity</label>
                  <Tooltip content="The gratuity provided on this expense">
                    <CircleHelp className="w-4 h-4 ml-2" />
                  </Tooltip>
                </div>
                <Input
                  type="number"
                  startContent="$"
                  placeholder="0.00"
                  className="max-w-52"
                />
              </div>
              <div className="flex items-center justify-between">
                <label className="text-sm">Date</label>
                <Input required type="date" className="max-w-52" />
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <label className="text-sm">Total</label>
                  <Tooltip content="The sum of all items in this expense">
                    <CircleHelp className="w-4 h-4 ml-2" />
                  </Tooltip>
                </div>
                <Input
                  isReadOnly
                  type="number"
                  startContent="$"
                  placeholder="0.00"
                  className="max-w-52"
                />
              </div>
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
