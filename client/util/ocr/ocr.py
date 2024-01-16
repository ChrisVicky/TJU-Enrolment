# ocr.py
import ddddocr
import sys


def main(image_path):
    ocr = ddddocr.DdddOcr(show_ad=False)
    with open(image_path, "rb") as f:
        img_bytes = f.read()
    text = ocr.classification(img_bytes)
    print(text)


if __name__ == "__main__":
    main(sys.argv[1])
