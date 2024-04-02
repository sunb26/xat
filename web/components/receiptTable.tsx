"use client";

import { ReceiptForm } from "@/components/receiptForm";
import {
  Card,
  CardBody,
  CardHeader,
  Table,
  TableBody,
  TableColumn,
  TableHeader,
} from "@nextui-org/react";

export const ReceiptTable = () => {
  return (
    <Card className="p-6" id="receipt">
      <CardHeader className="text-lg font-semibold justify-between">
        Your Receipts <ReceiptForm />
      </CardHeader>
      <CardBody>
        <Table removeWrapper aria-label="Receipts table">
          <TableHeader>
            <TableColumn>GST/HST</TableColumn>
            <TableColumn>TOTAL</TableColumn>
            <TableColumn>HISTORY</TableColumn>
          </TableHeader>
          <TableBody emptyContent={"No rows to display."}>{[]}</TableBody>
        </Table>
      </CardBody>
    </Card>
  );
};
