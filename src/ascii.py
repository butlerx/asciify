"""generate  image to ascii"""
from .image import Image, grayscale, resize


def pixel_to_char(image: Image, buckets: int = 25) -> str:
    """
    replaces every pixel with a character whose intensity is similar
    """
    ascii_chars = [" ", ",", ":", ";", "+", "*", "?", "%", "S", "#", "@"][::-1]
    return "".join(
        [ascii_chars[pixel_value // buckets] for pixel_value in list(image.getdata())]
    )


def to_ascii(image: Image, new_width: int = 100) -> str:
    """
    Converts and Image too ascii image
    """
    pixels = pixel_to_char(grayscale(resize(image, new_width)))

    # Construct the image from the character list
    return "\n".join(
        [
            str(pixels[index : index + new_width]).rstrip()
            for index in range(0, len(pixels), new_width)
        ]
    )
