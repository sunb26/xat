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
    total: {
      subtotal: 100.0,
      gratuity: 10.0,
      total: 123.0,
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
  { name: "GST/HST ($)", id: "gst" },
  { name: "TOTAL ($)", id: "total" },
  { name: "HISTORY", id: "history" },
];
type Column = (typeof columns)[0];

export const ReceiptTable = () => {
  const renderCell = useCallback((data: Receipt, columnKey: string) => {
    const cellValue = data[columnKey as keyof Receipt];

    switch (columnKey) {
      case "gst":
        return <div>{data.gst.toFixed(2)}</div>;
      case "total":
        return (
          <Popover placement="right">
            <PopoverTrigger>
              <Button variant="light">{data.total.total.toFixed(2)}</Button>
            </PopoverTrigger>
            <PopoverContent>
              {Object.entries(data.total).map(([key, value]) => (
                <span key={key}>
                  {key.charAt(0).toUpperCase() + key.slice(1)}: $
                  {value.toFixed(2)}
                </span>
              ))}
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
