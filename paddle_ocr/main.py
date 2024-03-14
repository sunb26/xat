from absl import app
from absl import flags
from paddleocr import PaddleOCR
import paddle

_IMAGE_URI = flags.DEFINE_string("image_uri", None, "URI of the image to process.")


def main(argv: list[str]) -> None:
    if len(argv) > 1:
        raise ValueError("too many arguments")
    paddle.utils.run_check()
    ocr = PaddleOCR(use_angle_cls=True, lang="en")
    res = ocr.ocr(_IMAGE_URI.value)
    print(res)


if __name__ == "__main__":
    app.run(main)
