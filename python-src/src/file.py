"""functions for writeing to file"""


def inject_readme(readme_path: str, image: str) -> str:
    """inject ascii text into readme"""
    with open(readme_path, "r") as file:
        return inject(file.read(), image)


def inject(readme: str, image: str) -> str:
    """replace {{ascii}} with ascii image"""
    txt = readme.split("{{ascii}}", 1)
    return "{}\n```\n{}\n```\n{}".format(txt[0], image, txt[1])


def save(path: str, image: str):
    """write to file"""
    with open(path, "w") as file:
        file.write(image)
        file.close()
