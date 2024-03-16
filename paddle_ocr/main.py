import os
from pathlib import Path
from typing import Annotated, Any

import pykka
import uvicorn
from absl import app, flags
from fastapi import FastAPI, File
from fastapi.responses import JSONResponse
from paddleocr import PaddleOCR

_MOCK = flags.DEFINE_boolean(
    "mock",
    False,
    "Skips inference and outputs constant mock example with consistent structure.",
)


class Model(pykka.ThreadingActor):
    _MOCK_RES = [
        [
            [[441.0, 174.0], [1166.0, 176.0], [1165.0, 222.0], [441.0, 221.0]],
            ("ACKNOWLEDGEMENTS", 0.9971134662628174),
        ],
        [
            [[403.0, 346.0], [1204.0, 348.0], [1204.0, 384.0], [402.0, 383.0]],
            ("We would like to thank all the designers and", 0.9761400818824768),
        ],
        [
            [[403.0, 396.0], [1204.0, 398.0], [1204.0, 434.0], [402.0, 433.0]],
            ("contributors who have been involved in the", 0.9791957139968872),
        ],
    ]

    def __init__(self, mock: bool) -> None:
        super().__init__()
        self._mock = mock
        paddle_model = Path(os.environ["PADDLE_MODEL"])
        self._ocr = PaddleOCR(
            use_angle_cls=True,
            lang="en",
            # NOTE: model dirs taken from paddleocr.py (containing PaddleOCR)
            det_model_dir=str(paddle_model / "en_PP-OCRv3_det_infer"),
            rec_model_dir=str(paddle_model / "en_PP-OCRv4_rec_infer"),
            cls_model_dir=str(paddle_model / "ch_ppocr_mobile_v2.0_cls_infer"),
        )

    def infer(self, uri_or_content: str | bytes) -> list[list[Any]]:
        if self._mock:
            return type(self)._MOCK_RES
        return self._ocr.ocr(uri_or_content)


def main(argv: list[str]) -> None:
    print(argv)
    print([f"{key}={flag.value}" for key, flag in flags.FLAGS.__flags.items()])
    model = Model.start(mock=_MOCK.value).proxy()
    api = FastAPI()

    @api.get("/api/v1/ocr/uri/{uri:path}")
    async def _api_v1_ocr_uri(uri: str) -> JSONResponse:
        return JSONResponse(content=model.infer(uri).get())

    @api.post("/api/v1/ocr/content")
    async def _api_v1_ocr_content(content: Annotated[bytes, File()]) -> JSONResponse:
        return JSONResponse(content=model.infer(content).get())

    uvicorn.run(api, host="127.0.0.1", port=3001)


if __name__ == "__main__":
    app.run(main)
