package utiles

import (
	"encoding/json"
	dockerBackend "github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/api/types/network"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"path/filepath"
	"time"
)

func BackupContainer(ctx *svc.ServiceContext) error {
	containerList, err := GetContainerList(ctx)
	if err != nil {
		return err
	}
	var backupList []dockerBackend.ContainerCreateConfig
	for _, v := range containerList {
		inspect, err := GetContainerInspect(ctx, v.ID)
		if err != nil {
			logx.Error("获取容器信息失败" + err.Error())
			return err
		}
		var containerName string
		if len(v.Names) > 0 {
			containerName = v.Names[0][1:]
		} else {
			containerName = "get container name error"
			logx.Error("get container name error" + v.ID)
		}
		inspect.Config.Hostname = ""
		inspect.Image = inspect.Config.Image
		config := inspect.Config
		hostConfig := inspect.HostConfig
		networkingConfig := &network.NetworkingConfig{
			EndpointsConfig: inspect.NetworkSettings.Networks,
		}
		createConfig := dockerBackend.ContainerCreateConfig{Config: config, HostConfig: hostConfig, NetworkingConfig: networkingConfig, Name: containerName}
		backupList = append(backupList, createConfig)
	}
	jsonData, err := json.MarshalIndent(backupList, "", "  ")
	if err != nil {
		logx.Error("Error marshalling data:", err)
		return err
	}
	backupDir := os.Getenv("BACKUP_DIR") // 从环境变量中获取备份目录
	if backupDir == "" {
		backupDir = "/data/backups" // 如果环境变量未设置，使用默认值
	}
	_, err = os.Stat(backupDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(backupDir, 0755)
		if err != nil {
			logx.Error("Error creating backup directory:", err)
			return err
		}
	}
	currentDate := time.Now().Format("2006-01-02")
	fileName := "backup-" + currentDate + ".json"
	fullPath := filepath.Join(backupDir, fileName)
	err = os.WriteFile(fullPath, jsonData, 0644)
	if err != nil {
		logx.Error("Error writing to file:", err)
		return err
	}
	return nil
}
