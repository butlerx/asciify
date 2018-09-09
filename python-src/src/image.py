from PIL import Image


def resize(image: Image, new_width: int = 100) -> Image:
    """
    resizes the image into the final width while maintaining aspect ratio
    """
    (old_width, old_height) = image.size
    return image.resize(
        (new_width, int(float(old_height) / float(old_width) * new_width))
    )


def grayscale(image: Image) -> Image:
    """
    returns the grayscale version of image
    """
    return image.convert("L")
