# Tianjin University Enrollment Helper

## Features

- 自动完成登陆、选课任务
- 通过 `Python` `ocr` 脚本完成验证码识别，快速进行登陆
- 通过 `goroutine` 特性提高请求频率，提高成功率
- 支持多用户同时进行

## Usage

### 1. Configuration

```toml
[Program]
# 每个 course 选课的线程数量，越多越好，但由于CPU限制，太多了也会降低总体性能
Threads = 4


# # Script: Python Code
# [Ocr]
# api = "./scipt.py"
# type = 0

# # Remote Ocr server
# [Ocr]
# api = "https://remote.ocr.server"
# type = 1

# Local ocr server
[Ocr]
api = "http://127.0.0.1:8000/uploadfile/"
type = 2

# [[Account]]
# no = "ixxxv" # 学号
# password = "ixxxx" # 密码
# courses = { 02572 = "毕计", 01360 = "形与" }
# comment = "ffff"

[[Account]]
no = "303103049"
password = "11rrrga~"
courses = { 02688 = "B-毛球", 06587 = "育B-泳" }
comment = "rrrddd"
```

- 选择本地 ocr 则需要额外配置 ocr 的 python （在 ./client/util/ocr/localDdddocr/README.md)

### 2. Start OCR (if local)
