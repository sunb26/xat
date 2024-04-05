"use client";

import { useCallback } from "react";
import { ReceiptForm } from "@/components/receiptForm";
import {
  Card,
  CardBody,
  CardHeader,
  Table,
  TableBody,
  TableColumn,
  TableCell,
  TableHeader,
  TableRow,
  Button,
} from "@nextui-org/react";
import { History } from "lucide-react";

const data = [
  {
    id: 1,
    gst: 13.0,
    total: {
      total: 123.0,
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
        return <div>{data.total.total.toFixed(2)}</div>;
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
