#! /usr/bin/env python3
"""convert image to ascii"""
from io import BytesIO
from os.path import splitext
from re import IGNORECASE
from re import compile as _compile
from re import match
from sys import exit as _exit
from urllib.request import urlopen

from src import Image, get_args, inject_readme, save, to_ascii


def _main():
    try:
        args = get_args()
        image = to_ascii(
            Image.open(
                BytesIO(urlopen(args.image).read())
                if match(
                    _compile(
                        r"^(?:http|ftp)s?://"
                        r"(?:(?:[A-Z0-9](?:[A-Z0-9-]{0,61}[A-Z0-9])?\.)+(?:[A-Z]{2,6}\.?|[A-Z0-9-]{2,}\.?)|"
                        r"localhost|"
                        r"\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})"
                        r"(?::\d+)?"
                        r"(?:/?|[/?]\S+)$",
                        IGNORECASE,
                    ),
                    args.image,
                )
                else args.image
            ),
            args.width,
        )
        if args.print:
            print(image)
        elif args.readme:
            save(args.readme, inject_readme(args.readme, image))
        else:
            filename, _ = splitext(args.image)
            save("{}/{}.txt".format(args.path, filename), image)
        return 0
    except Exception as err:
        print("Unable to find image in", args.image)
        print(err)
        return 1


if __name__ == "__main__":
    _exit(_main())
