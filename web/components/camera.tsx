"use client";

import { Button, Card, CardBody, CardHeader, Image } from "@nextui-org/react";
import { useCallback, useRef, useState } from "react";
import Webcam from "react-webcam";

const videoConstraints = {
  width: 720,
  height: 360,
  facingMode: "user",
};

export const Camera = () => {
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

  return (
    <Card className="p-6" id="camera">
      <CardHeader className="text-lg font-semibold">
        Upload a Receipt
      </CardHeader>
      <CardBody className="flex flex-col items-center justify-center gap-4 min-h-96 border-dashed border-2 rounded-lg">
        {!isCaptureEnable && url == null ? (
          <Button onClick={() => setCaptureEnable(true)} color="primary">
            Enable Camera
          </Button>
        ) : (
          <></>
        )}
        {isCaptureEnable && (
          <>
            <div>
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
            </div>
            <Button onPress={capture} color="primary">
              Capture
            </Button>
          </>
        )}
        {url && (
          <>
            <div>
              <Image src={url} alt="receipt screenshot" />
            </div>
            <Button onPress={() => setUrl(null)} color="danger">
              Delete
            </Button>
          </>
        )}
      </CardBody>
    </Card>
  );
};
