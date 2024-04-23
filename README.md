# 小工具

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/xiaoxuan6/tools/releaser.yml)
![GitHub Downloads (all assets, all releases)](https://img.shields.io/github/downloads/xiaoxuan6/tools/total)
![GitHub Downloads (all assets, latest release)](https://img.shields.io/github/downloads-pre/xiaoxuan6/tools/latest/total?sort=semver&style=plastic)
![GitHub License](https://img.shields.io/github/license/xiaoxuan6/tools)
![GitHub Repo stars](https://img.shields.io/github/stars/xiaoxuan6/tools?style=flat)
![GitHub forks](https://img.shields.io/github/forks/xiaoxuan6/tools?style=flat)
![GitHub Tag](https://img.shields.io/github/v/tag/xiaoxuan6/tools)

## Install

### CMD

```bash
go install github.com/xiaoxuan6/tools@latest
```     

### 安装包

[tools 下载地址](https://github.com/xiaoxuan6/tools/releases)

## Usage

### 翻译

```bash
tools t -c "test"
echo "test" | tools t
cat "test.txt" | tools t
tools t < .gitignore
```

### 二维码

#### generate

```bash
tools q -c "test"
echo "test" | tools q 
cat "test.txt" | tools q
```

#### scan

```bash
tools q -f "qrcode.png"
 ```

### OCR

```bash
tools o -f ./23d3d34c-72a3-40ed-9281-5cf06566941b.jpg 
```

### clipboard2img

首先需要截图/复制到剪切板（主要解决 win10 使用 Shift+win+s/Shift+Ctrl+PrtSc 无法保存），然后执行如下：

```bash
tools c2i 
```
