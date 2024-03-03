# Tianjin University Enrollment Helper

## Features

- 自动完成登陆、选课任务
- 通过 `Python` `ocr` 脚本完成验证码识别，快速进行登陆
- 通过 `goroutine` 特性提高请求频率，提高成功率
- 支持多用户同时进行

## Usage

### 1. Configuration

```toml
# Pid 为学期编号，可在选课页面通过 F12 查看如下代码：
# <input type="hidden" name="electionProfile.id" value="2808">
Pid = "2808"

[Program]
# 每个 Account 选课的线程数量
Threads = 4


# # Script: Python Code
# [Ocr]
# api = "./scipt.py"
# type = 0

# # Remote Ocr server
# [Ocr]
# api = "https://learning.twt.edu.cn/ocr"
# type = 1

# Local ocr server
[Ocr]
api = "http://127.0.0.1:8000/uploadfile/"
type = 2

# [[Account]]
# no = "3020202184" # 学号
# password = "0407Christopher!" # 密码
# courses = { 02172 = "毕业设计（论文）", 01320 = "形势与政策" }
# comment = "刘锦帆"

[[Account]]
no = "3023207213"
password = "114514henghenga~"
courses = { 02688 = "体育B-羽毛球", 06487 = "体育B-游泳" }
comment = "陈祎唯"
```
