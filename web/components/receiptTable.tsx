"use client";

import { ReceiptForm } from "@/components/receiptForm";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Slider,
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@nextui-org/react";
import { History } from "lucide-react";
import { useCallback } from "react";

const data = [
  {
    id: 1,
    gst: 13.0,
    date: new Date("2024-04-04T18:25:43.511Z"),
    total: 123.0,
    items: {
      subtotal: 100.0,
      gratuity: 10.0,
    },
    history: [
      {
        id: 2,
        change: "Updated gratuity value",
        date: "2024-04-04",
      },
      {
        id: 1,
        change: "Created Receipt",
        date: "2024-04-01",
      },
    ],
  },
];
type Receipt = (typeof data)[0];

const columns = [
  { name: "TOTAL", id: "total" },
  { name: "GST/HST", id: "gst" },
  { name: "DATE", id: "date" },
  { name: "HISTORY", id: "history" },
];
type Column = (typeof columns)[0];

export const ReceiptTable = () => {
  const renderCell = useCallback((data: Receipt, columnKey: string) => {
    const cellValue = data[columnKey as keyof Receipt];

    switch (columnKey) {
      case "gst":
        return <div>${data.gst.toFixed(2)}</div>;
      case "date":
        return <div>{data.date.toLocaleDateString()}</div>;
      case "total":
        return (
          <Popover placement="right">
            <PopoverTrigger>
              <Button variant="light">${data.total.toFixed(2)}</Button>
            </PopoverTrigger>
            <PopoverContent className="p-4">
              <div className="flex flex-col gap-0.5">
                {Object.entries(data.items).map(([key, value]) => (
                  <div key={key} className="flex justify-between gap-3">
                    <span>{key.charAt(0).toUpperCase() + key.slice(1)}:</span>
                    <span>${value.toFixed(2)}</span>
                  </div>
                ))}
                <div className="flex justify-between gap-3">
                  <span>GST/HST</span>
                  <span>${data.gst.toFixed(2)}</span>
                </div>
              </div>
            </PopoverContent>
          </Popover>
        );
      case "history":
        return (
          <Sheet>
            <SheetTrigger asChild>
              <Button isIconOnly variant="light">
                <History className="w-5 h-5" />
              </Button>
            </SheetTrigger>
            <SheetContent>
              <SheetHeader>
                <SheetTitle>Receipt History</SheetTitle>
                <SheetDescription>
                  Here are all the changes made to this receipt
                </SheetDescription>
              </SheetHeader>
              <div className="h-full p-6">
                {data.history.map((item) => (
                  <div className="flex gap-4 mb-4" key={item.id}>
                    <Slider
                      isDisabled
                      orientation="vertical"
                      size="sm"
                      color="foreground"
                      maxValue={1}
                      minValue={0}
                      defaultValue={0.8}
                      className="h-16"
                    />
                    <div className="flex flex-col gap-2">
                      <div className="font-semibold">{item.date}</div>
                      <div className="text-sm">{item.change}</div>
                    </div>
                  </div>
                ))}
              </div>
            </SheetContent>
          </Sheet>
        );
      default:
        return cellValue;
    }
  }, []);

  return (
    <Card className="p-6" id="receipt">
      <CardHeader className="text-lg font-semibold justify-between">
        Your Receipts <ReceiptForm />
      </CardHeader>
      <CardBody>
        <Table removeWrapper aria-label="Receipts table">
          <TableHeader columns={columns}>
            {(column: Column) => (
              <TableColumn key={column.id}>{column.name}</TableColumn>
            )}
          </TableHeader>
          <TableBody emptyContent={"No rows to display."} items={data}>
            {(item: Receipt) => (
              <TableRow key={item.id}>
                {(columnKey: string) => (
                  // @ts-ignore
                  <TableCell>{renderCell(item, columnKey)}</TableCell>
                )}
              </TableRow>
            )}
          </TableBody>
        </Table>
      </CardBody>
    </Card>
  );
};
