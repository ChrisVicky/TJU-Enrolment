# Tianjin University Enrolment Helper

## Features

- 通过 `Python` `ocr` 脚本完成验证码识别，快速进行登陆
- 通过 `goroutine` 特性提高请求频率，提高成功率
- 支持多用户同时进行
- 自动完成登陆、选课任务

## Usage

### 1. Configuration

- 修改 `config.toml` 适配自己的需求

  - `mv ./conf/example.toml config.toml`
  - 修改 `config.toml`

- 选择本地 ocr server 需要额外配置 ocr 的 python （ `./client/util/ocr/localDdddocr/README.md`)

### 2. Start OCR (if local)

- 查看其中的 README.md 进行配置

### 3. Start

- 使用预先编译好的代码：

  - Mac (Intel): `enrolment.mac`
  - Linux (x86_64): `enrolment.linux`
  - Windows (x86_64): `enrolment.exe`

- 使用 Golang
  - go run main.go

## Coding

```
.
|-- client              - Enrolment Client
|   `-- util            - Js Encoder & OCR Helper
|       |-- jsencoder
|       `-- ocr
|-- conf                - Configuration
|-- logger              - Logger Issues
|-- routine             - GoRoutine Helper
`-- runtime

15 directories
```
