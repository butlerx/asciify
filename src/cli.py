"""cli interface"""
from argparse import ArgumentParser, Namespace, RawDescriptionHelpFormatter
from os import getcwd
from textwrap import dedent


def get_args() -> Namespace:
    """gets cli arguments"""
    parser = ArgumentParser(
        prog="ascii-logo",
        formatter_class=RawDescriptionHelpFormatter,
        usage="%(prog)s [options] image.png",
        description=dedent(
            """
            Every project needs a logo, so why not an ascii one.
            Convert image to ascii.
            Supports injecting image to readme
            """
        ),
    )
    parser.add_argument(
        "image",
        metavar="image.png",
        type=str,
        help="Image to convert, Supports remote url",
    )
    parser.add_argument(
        "--print", "-p", action="store_true", help="Print image to terminal"
    )
    parser.add_argument(
        "--path",
        metavar="./folder/to/output",
        default=getcwd(),
        type=str,
        help="Path to output txt file too (defaults to current dir)",
    )
    parser.add_argument(
        "--readme",
        "-r",
        type=str,
        help="Path to readme.md to inject ascii too",
        metavar="README.md",
    )
    parser.add_argument(
        "--width",
        "-w",
        default=100,
        type=int,
        help="Width of output (default:100)",
        metavar="100",
    )
    return parser.parse_args()
