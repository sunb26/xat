import lzma

from absl import app, flags

_INPUT = flags.DEFINE_string("input", None, "Path to input .xz file.")
_OUTPUT = flags.DEFINE_string("output", None, "Path to output decompressed file.")


def main(argv: list[str]) -> None:
    if len(argv) > 1:
        raise ValueError("too many arguments")
    if _INPUT.value is None or _OUTPUT.value is None:
        raise ValueError("expected some value for input and output")
    with lzma.open(_INPUT.value, "rb") as f:
        content = f.read()
    with open(_OUTPUT.value, "wb") as f:
        f.write(content)


if __name__ == "__main__":
    app.run(main)
