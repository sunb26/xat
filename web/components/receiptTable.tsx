"use client";

import { ReceiptForm } from "@/components/receiptForm";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Popover,
  PopoverContent,
  PopoverTrigger,
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
        change: "Updated gratuity",
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
          <Button isIconOnly variant="light">
            <History className="w-5 h-5" />
          </Button>
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
