from absl import app, flags
from pip_requirements_parser import RequirementsFile

_INPUT = flags.DEFINE_string("input", None, "Path to input requirements.txt file.")
_OUTPUT = flags.DEFINE_string(
    "output", None, "Path to output patched requirements file."
)
_EXCLUDE = flags.DEFINE_multi_string(
    "exclude",
    [],
    "Exclude a package from requirements file. Repeat the flag to exclude multiple packages.",
)


def main(argv: list[str]) -> None:
    if len(argv) > 1:
        raise ValueError("too many arguments")
    if _INPUT.value is None or _OUTPUT.value is None:
        raise ValueError("expected some value for input and output")
    if _EXCLUDE.value is None:
        return
    requirements_file = RequirementsFile.from_file(_INPUT.value)
    requirements = (
        f"{requirement.dumps()}\n"
        for requirement in requirements_file.requirements
        if not requirement.req or requirement.req.name not in _EXCLUDE.value
    )
    with open(_OUTPUT.value, "w") as f:
        f.writelines(requirements)


if __name__ == "__main__":
    app.run(main)
