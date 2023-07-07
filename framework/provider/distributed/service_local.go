package distributed

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/bref/outsider/framework"
	"github.com/bref/outsider/framework/contract"
)

type LocalDistributedService struct {
	contract.Distributed
	container framework.Container
}

// NewLocalDistributedService 初始化本地分布式服务
func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	return &LocalDistributedService{container: container}, nil
}

// Select 为分布式选择器
func (s LocalDistributedService) Select(serviceName string, appID string, holdTime time.Duration) (selectAppID string, err error) {
	appService := s.container.MustMake(contract.AppKey).(contract.App)

	runtimeFolder := appService.RuntimeFolder()

	localFile := filepath.Join(runtimeFolder, "distribute_"+serviceName)

	// 打开文件锁
	file, err := os.OpenFile(localFile, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}

	// 尝试独占文件锁
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	// 因为使用LOCK_NB模式,获取不到锁的情况下会报错
	if err != nil {
		// 读取被选择的app id
		selectAppIDByt, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}

		return string(selectAppIDByt), err
	}

	// 此时已经获取了锁
	go func() {
		defer func() {
			// 释放文件锁
			syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
			// 释放文件
			file.Close()
			// 删除文件锁对应的文件
		}()

		// 创建选举结果有效的计时器
		timer := time.NewTimer(holdTime)

		// 等待计时器结束
		<-timer.C
	}()

	// 这里已经抢占到了,将抢占到的appID写入文件
	if _, err := file.WriteString(appID); err != nil {
		return "", err
	}
	return appID, nil
}
