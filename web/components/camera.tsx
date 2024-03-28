"use client";

import {
  Drawer,
  DrawerClose,
  DrawerContent,
  DrawerTrigger,
} from "@/components/ui/drawer";
import { Button, Image } from "@nextui-org/react";
import { useCallback, useRef, useState } from "react";
import Webcam from "react-webcam";

const videoConstraints = {
  width: 720,
  height: 480,
  facingMode: "user",
};

export const Camera = () => {
  const [isCaptureEnable, setCaptureEnable] = useState<boolean>(true);
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
        <Button size="sm">Open camera</Button>
      </DrawerTrigger>
      <DrawerContent>
        <div className="flex flex-col p-4 gap-4 min-h-96 mx-auto w-full max-w-md">
          {isCaptureEnable && (
            <>
              {/*
            NOTE: https://github.com/microsoft/TypeScript/issues/31147
            // @ts-ignore*/}
              <Webcam
                audio={false}
                width={540}
                height={360}
                ref={webcamRef}
                screenshotFormat="image/jpeg"
                videoConstraints={videoConstraints}
                className="rounded-md"
              />
              <Button onPress={capture} color="primary">
                Capture
              </Button>
            </>
          )}
          {url && (
            <>
              <Image src={url} alt="receipt screenshot" />
              <div className="flex gap-2">
                <Button onPress={resetCamera} className="w-full">
                  Re-take
                </Button>
                <Button color="primary" className="w-full">
                  Continue
                </Button>
              </div>
            </>
          )}
          <DrawerClose asChild>
            <Button variant="bordered">Cancel</Button>
          </DrawerClose>
        </div>
      </DrawerContent>
    </Drawer>
  );
};
