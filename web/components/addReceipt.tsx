"use client";

import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerTrigger,
} from "@/components/ui/drawer";
import { Button, Image } from "@nextui-org/react";
import React, { useCallback, useRef, useState } from "react";
import Webcam from "react-webcam";
import { Input } from "@/components/ui/input";

const videoConstraints = {
  width: 240,
  height: 480,
  facingMode: "user",
};

export const AddReciept = () => {
  const [isCaptureEnable, setCaptureEnable] = useState<boolean>(false);
  const webcamRef = useRef<Webcam>(null);
  const [url, setUrl] = useState<string | null>(null);
  const capture = useCallback(() => {
    const imageSrc = webcamRef.current?.getScreenshot();
    if (imageSrc) {
      setUrl(imageSrc);
    }
    setCaptureEnable(false);
  }, []);
  function resetCamera() {
    setUrl(null);
    setCaptureEnable(true);
  }

  return (
    <Drawer>
      <DrawerTrigger asChild>
        <Button>Add receipt</Button>
      </DrawerTrigger>
      <DrawerContent>
        <div className="flex flex-col p-4 gap-4 min-h-96 mx-auto w-full max-w-md">
          {!url && !isCaptureEnable && (
            <div className="flex flex-col gap-6">
              <Input type="file" />
              <div className="text-lg font-semibold">or</div>
              <Button onPress={() => setCaptureEnable(true)}>Take a picture</Button>
            </div>
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
              <div className="flex mx-auto">
                <Image src={url} alt="receipt screenshot" />
              </div>
              <div className="flex">
                <Button onPress={resetCamera} className="w-full">
                  Re-take
                </Button>
              </div>
            </>
          )}
          <Button isDisabled={!url}>Continue</Button>
          <DrawerClose asChild>
            <Button variant="bordered">Cancel</Button>
          </DrawerClose>
        </div>
      </DrawerContent>
    </Drawer>
  );
};
