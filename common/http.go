package common

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func PostWithMultipart(url, name, value string, file *os.File) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile(name, value)
	if err != nil {
		return nil, fmt.Errorf("创建文件失败")
	}

	_, _ = io.Copy(fileWriter, file)
	contentType := bodyWriter.FormDataContentType()
	_ = bodyWriter.Close()

	response, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return nil, fmt.Errorf("请求失败，请重新输入")
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败")
	}

	return b, nil
}
