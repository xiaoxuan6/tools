# 小工具

# Install

通过 `cmd` 安装

```bash
go install github.com/xiaoxuan6/tools@latest
```     

通过安装包安装，需要添加到环境变量中，[下载地址](https://github.com/xiaoxuan6/tools/releases)

## 翻译

```bash
tools t -c "test"
or 
echo "test" | tools t -s
cat "test.txt" | tools t -s
```

## qrcode

### generate

```bash
tools q -c "test"
or 
echo "test" | tools q -s
cat "test.txt" | tools q -s
```

### scan

```bash
tools q -f "qrcode.png"
 ```

## OCR

```bash
tools o -f ./23d3d34c-72a3-40ed-9281-5cf06566941b.jpg 
```

## clipboard2img

首先需要截图/复制到剪切板（主要解决 win10 使用 Shift+win+s/Shift+Ctrl+PrtSc 无法保存），然后执行如下：

```bash
tools c2i 
```