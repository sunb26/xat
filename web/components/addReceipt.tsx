"use client";

import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerTrigger,
} from "@/components/ui/drawer";
import { Input } from "@/components/ui/input";
import { useUser } from "@clerk/clerk-react";
import { Button, Image } from "@nextui-org/react";
import { useCallback, useRef, useState } from "react";
import Webcam from "react-webcam";

const videoConstraints = {
  width: 240,
  height: 480,
  facingMode: "user",
};

export const AddReciept = () => {
  const { isLoaded, user } = useUser();
  const [isCaptureEnable, setCaptureEnable] = useState<boolean>(false);
  const webcamRef = useRef<Webcam>(null);
  const [url, setUrl] = useState<string | null>(null);
  const [uploadedFile, setUploadedFile] = useState<File>();

  const capture = useCallback(() => {
    const imageSrc = webcamRef.current?.getScreenshot();
    if (imageSrc) {
      setUrl(imageSrc);
    }
    setCaptureEnable(false);
  }, []);

  function handleCancel() {
    setUrl(null);
    setCaptureEnable(false);
    setUploadedFile(undefined);
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    if (e.target.files?.[0]) {
      setUploadedFile(e.target.files[0]);
    }
  };

  const handleFileUpload = async () => {
    if (!uploadedFile) return;
    const userId = user?.id;

    const body = new FormData();
    body.append("file", uploadedFile);

    try {
      await fetch(
        "https://www.googleapis.com/upload/drive/v3/files?uploadType=media",
        {
          method: "POST",
          body: body,
          headers: {
            Authorization: `Bearer ${process.env.GOOGLE_API_KEY}`,
            "Content-Type": uploadedFile.type,
          },
        },
      );
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <Drawer>
      <DrawerTrigger asChild>
        <Button>Add receipt</Button>
      </DrawerTrigger>
      <DrawerContent>
        <div className="flex flex-col p-4 gap-4 min-h-96 mx-auto w-full max-w-md justify-between">
          <div className="flex flex-col w-full gap-4">
            {!url && !isCaptureEnable && (
              <Input type="file" onChange={handleChange} />
            )}
            {!url && !uploadedFile && !isCaptureEnable && (
              <div className="text-lg font-semibold mx-auto">or</div>
            )}
            {!uploadedFile && !isCaptureEnable && !url && (
              <Button onPress={() => setCaptureEnable(true)}>
                Take a picture
              </Button>
            )}
            {isCaptureEnable && (
              <>
                {/*
            NOTE: https://github.com/microsoft/TypeScript/issues/31147
            // @ts-ignore*/}
                <Webcam
                  audio={false}
                  width={240}
                  height={360}
                  ref={webcamRef}
                  screenshotFormat="image/jpeg"
                  videoConstraints={videoConstraints}
                  className="rounded-md mx-auto"
                />
                <Button onPress={capture} color="primary">
                  Capture
                </Button>
              </>
            )}
            {url && (
              <>
                <div className="mx-auto">
                  <Image src={url} alt="receipt screenshot" />
                </div>
                <div className="flex">
                  <Button onPress={handleCancel} className="w-full">
                    Re-take
                  </Button>
                </div>
              </>
            )}
          </div>
          <div className="flex gap-2">
            <DrawerClose asChild>
              <Button
                variant="bordered"
                onPress={handleCancel}
                className="w-full"
              >
                Cancel
              </Button>
            </DrawerClose>
            <Button
              isDisabled={!url && !uploadedFile}
              color="primary"
              className="w-full"
              onPress={handleFileUpload}
            >
              Continue
            </Button>
          </div>
        </div>
      </DrawerContent>
    </Drawer>
  );
};
