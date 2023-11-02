# 小工具

# Install

```bash
go install github.com/xiaoxuan6/tools@install
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

