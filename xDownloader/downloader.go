package xDownloader

import (
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/unix755/xtools/xExec"
)

// Download 下载
func Download(fileUrl string, outputFile string, outputFolder string) (err error) {
	// 使用get方法连接url
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 解析输出文件名
	if outputFile == "" {
		_, m, _ := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
		if m["filename"] != "" {
			outputFile = m["filename"]
		} else {
			outputFile = path.Base(resp.Request.URL.Path)
		}
	}

	// 仅有这一种情况需要进行文件与输出路径的拼接
	if filepath.Dir(outputFile) == "." && outputFolder != "" {
		outputFile = filepath.Join(outputFolder, outputFile)
	}

	// 创建输出文件夹
	err = os.MkdirAll(filepath.Dir(outputFile), 0755)
	if err != nil {
		return err
	}

	// 创建输出文件
	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	// 将数据写入到输出文件中
	_, err = io.Copy(f, resp.Body)
	return err
}

// DownloadWithCurl 使用 curl 下载
func DownloadWithCurl(fileUrl string, outputFile string, outputFolder string) (err error) {
	// 参数中未指定输出文件名时, 解析连接中的文件名
	if outputFile == "" {
		// 建立连接
		resp, err := http.Get(fileUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// 解析文件名
		_, m, _ := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
		if m["filename"] != "" {
			outputFile = m["filename"]
		} else {
			outputFile = path.Base(resp.Request.URL.Path)
		}
	}

	// 仅有这一种情况需要进行文件与输出路径的拼接
	if filepath.Dir(outputFile) == "." && outputFolder != "" {
		outputFile = filepath.Join(outputFolder, outputFile)
	}

	// 创建最终输出文件夹
	err = os.MkdirAll(filepath.Dir(outputFile), 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("curl", "-Lo", outputFile, fileUrl)
	return xExec.Run(cmd)
}
