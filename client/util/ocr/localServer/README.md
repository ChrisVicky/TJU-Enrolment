# Local Server running `ddddocr`

## Setup: python 3.10

1. Conda Env initialization with Python: 3.10: `conda create --name class-env python=3.10`
2. Activate Conda: `conda activate class-env`
3. Install necessary pkgs: `conda install pip`, then: `pip install ddddocr "fastapi[all]"`

## Execute

- `uvicorn main:app --reload`
