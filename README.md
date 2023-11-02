# 小工具

# Install

```bash
go install github.com/xiaoxuan6/tools@latest
```     

## 翻译

```bash
tools t -c "test"
or 
echo "test" | tools t -s
cat "test.txt" | tools t -s
or
tools t -f "a.txt"
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
