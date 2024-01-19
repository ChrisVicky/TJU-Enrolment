# `ddddocr` Setup

## Using `conda` to setup environment

```shell
conda create -n ddddocr python=3.8
conda activate ddddocr
conda install pip
pip install ddddocr
[[ $(python ocr.py test.png) == "3g5t" ]] && echo "OK" || echo "Error"
```
