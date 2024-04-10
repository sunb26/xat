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
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Button, Input, Tooltip } from "@nextui-org/react";
import { ChevronsUpDown, CircleHelp, CircleX, Plus } from "lucide-react";
import { useState } from "react";

const data = [
  { id: 1, item: "Mouse", price: 80.23 },
  { id: 2, item: "Headphones", price: 354.32 },
  { id: 3, item: "Monitor", price: 310.34 },
];

export const ReceiptForm = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [items, setItems] = useState(data);

  const addItem = () => {
    const index = items.length > 0 ? items[items.length - 1].id + 1 : 1;
    setItems([...items, { id: index, item: "", price: 0.0 }]);
  };

  const removeItem = (index: number) => {
    const newItems = [...items];
    newItems.splice(index, 1);
    setItems(newItems);
  };

  const handleItemChange = (
    id: number,
    key: string,
    value: string | number
  ) => {
    const updatedItems = [...items];
    updatedItems[id] = { ...updatedItems[id], [key]: value };
    setItems(updatedItems);
  };

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
              Please double-check the following information is correct.
            </DrawerDescription>
          </DrawerHeader>
          <div className="p-4">
            <form className="flex flex-col gap-4">
              <Collapsible open={isOpen} onOpenChange={setIsOpen}>
                <div className="flex items-center justify-between ">
                  <h1 className="text-sm font-semibold">Edit receipt items</h1>
                  <CollapsibleTrigger asChild>
                    <Button variant="ghost" size="sm" className="w-9 p-0">
                      <ChevronsUpDown className="h-4 w-4" />
                      <span className="sr-only">Toggle</span>
                    </Button>
                  </CollapsibleTrigger>
                </div>
                <CollapsibleContent>
                  <ScrollArea className="h-64 rounded-md border m-2 pt-1 px-2">
                    {items.map((item, index) => (
                      <div className="flex" key={item.id}>
                        <Input
                          variant="underlined"
                          type="text"
                          placeholder="Item"
                          size="sm"
                          value={item.item}
                          onChange={(e) =>
                            handleItemChange(index, "item", e.target.value)
                          }
                        />
                        <Input
                          variant="underlined"
                          type="number"
                          startContent="$"
                          size="sm"
                          value={item.price.toString()}
                          onChange={(e) =>
                            handleItemChange(
                              index,
                              "price",
                              Number.parseFloat(e.target.value)
                            )
                          }
                        />
                        <Button
                          isIconOnly
                          variant="light"
                          onPress={() => removeItem(index)}
                        >
                          <CircleX className="w-5 h-5" />
                        </Button>
                      </div>
                    ))}
                    <Button
                      isIconOnly
                      variant="bordered"
                      className="w-full my-2"
                      onPress={() => addItem()}
                    >
                      <Plus />
                    </Button>
                  </ScrollArea>
                </CollapsibleContent>
              </Collapsible>
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <label className="text-sm">Subtotal</label>
                  <Tooltip content="The pre-tax amount.">
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
                  <Tooltip content="Goods and Services Tax/Harmonized Sales Tax.">
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
                  <Tooltip content="The gratuity listed on this receipt.">
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
                <div className="flex items-center">
                  <label className="text-sm">Date</label>
                  <Tooltip content="The date this receipt was issued.">
                    <CircleHelp className="w-4 h-4 ml-2" />
                  </Tooltip>
                </div>
                <Input required type="date" className="max-w-52" />
              </div>
              <div className="flex items-center justify-between">
                <div className="flex items-center">
                  <label className="text-sm">Total</label>
                  <Tooltip content="The sum of all expenses on this receipt after tax.">
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
