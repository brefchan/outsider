package env

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path"
	"strings"

	"github.com/bref/outsider/framework/contract"
)

type HadeEnv struct {
	folder string            // 代表.env所在目录
	maps   map[string]string //保存所有环境变量
}

func NewHadeEnv(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewHadeEnv param error")
	}

	// 读取folder文件
	folder := params[0].(string)

	// 实例化
	hadeEnv := &HadeEnv{
		folder: folder,
		maps:   map[string]string{"APP_ENV": contract.EnvDevelopment},
	}

	// 解析folder/.env文件
	filePath := path.Join(folder, ".env")

	// 打开文件 .env
	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()

		// 读取文件
		bRader := bufio.NewReader(file)
		for {
			// 按行进行读取
			line, _, c := bRader.ReadLine()
			if c == io.EOF {
				break
			}

			// 按照等号解析
			s := bytes.SplitN(line, []byte{'='}, 2)
			// 如果不符合规范
			if len(s) < 2 {
				continue
			}

			// 保存map
			key := string(s[0])
			val := string(s[1])
			hadeEnv.maps[key] = val
		}

	}

	// 获取当前程序的环境变量,并且覆盖.env文件下的变量
	for _, environ := range os.Environ() {
		pair := strings.SplitN(environ, "=", 2)
		if len(pair) < 2 {
			continue
		}
		hadeEnv.maps[pair[0]] = pair[1]

	}

	// 返回实例
	return hadeEnv, nil

}

// AppEnv 获取表示当前APP环境的变量APP_ENV
func (en *HadeEnv) AppEnv() string {
	return en.Get("APP_ENV")
}

// IsExist 判断一个环境变量是否有被设置
func (en *HadeEnv) IsExist(key string) bool {
	_, ok := en.maps[key]
	return ok
}

// Get 获取某个环境变量,如果没有设置,返回""
func (en *HadeEnv) Get(key string) string {
	if val, ok := en.maps[key]; ok {
		return val
	}
	return ""
}

// All 获取所有的环境变量 .env和运行环境变量融合后结果
func (en *HadeEnv) All() map[string]string {
	return en.maps
}
