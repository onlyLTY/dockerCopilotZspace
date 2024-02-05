package utiles

import (
	"encoding/json"
	dockerBackend "github.com/docker/docker/api/types/backend"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	url2 "net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func RestoreContainer(ctx *svc.ServiceContext, filename string, taskID string) error {
	var backupList []string
	var url string
	basePath := os.Getenv("BACKUP_DIR") // 从环境变量中获取备份目录
	if basePath == "" {
		basePath = "/data/backups" // 如果环境变量未设置，使用默认值
	}
	fullPath := filepath.Join(basePath, filename+".json")
	oldProgress := svc.TaskProgress{
		TaskID:     taskID,
		Percentage: 0,
		Message:    "",
		Name:       "",
		DetailMsg:  "",
		IsDone:     false,
	}
	oldProgress.Name = "恢复容器"
	content, err := os.ReadFile(fullPath)
	if err != nil {
		logx.Error("Failed to read file: %s", err)
		oldProgress.Percentage = 0
		oldProgress.Message = "读取文件失败或者未找到文件。请确认文件名仅由大小写字母、数字和短横线组成"
		oldProgress.DetailMsg = err.Error()
		oldProgress.IsDone = true
		ctx.UpdateProgress(taskID, oldProgress)
	}
	var configList []dockerBackend.ContainerCreateConfig
	err = json.Unmarshal(content, &configList)
	if err != nil {
		logx.Error("Failed to parse json: %s", err)
		oldProgress.Percentage = 0
		oldProgress.Message = "解析文件失败"
		oldProgress.DetailMsg = err.Error()
		oldProgress.IsDone = true
		ctx.UpdateProgress(taskID, oldProgress)
	}
	jwtToken, endpointsId, err := GetNewJwt(ctx)
	if err != nil {
		oldProgress.Message = "连接Docker失败"
		oldProgress.DetailMsg = err.Error()
		oldProgress.IsDone = true
		ctx.UpdateProgress(taskID, oldProgress)
		return err
	}
	client := NewCustomClient(jwtToken)
	baseURL := domain + "/api/endpoints/" + endpointsId
	for i, containerInfo := range configList {
		info := "正在恢复第" + strconv.Itoa(i+1) + "个容器"
		oldProgress.Percentage = int(float64(i) / float64(len(configList)) * 100)
		oldProgress.Message = info
		oldProgress.DetailMsg = info
		ctx.UpdateProgress(taskID, oldProgress)
		params := url2.Values{}
		imageName := strings.Split(containerInfo.Config.Image, ":")
		params.Add("fromImage", imageName[0])
		if len(imageName) > 1 {
			params.Add("tag", imageName[1])
		} else {
			params.Add("tag", "latest")
		}
		url = baseURL + "/docker/images/create?" + params.Encode()
		reader, err := client.SendRequest("POST", url, nil)
		if err != nil {
			oldProgress.Message = "拉取新镜像失败"
			oldProgress.DetailMsg = err.Error()
			oldProgress.IsDone = true
			ctx.UpdateProgress(taskID, oldProgress)
			return err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logx.Errorf("Body.Close error: %v", err)
			}
		}(reader.Body)
		decodePullResp(reader.Body, ctx, taskID)
		err = CreateContainer(ctx, containerInfo.Config, containerInfo.HostConfig, containerInfo.NetworkingConfig, containerInfo.Name)
		if err != nil {
			logx.Error("Failed to create container: %s", err)
			info = "正在恢复第" + strconv.Itoa(i+1) + "个容器"
			backupList = append(backupList, containerInfo.Name+"恢复失败"+err.Error())
			continue
		} else {
			backupList = append(backupList, containerInfo.Name+"恢复成功")
		}
	}
	oldProgress.Percentage = 100
	oldProgress.DetailMsg = strings.Join(backupList, ",\n")
	oldProgress.Message = "恢复完成"
	oldProgress.IsDone = true
	ctx.UpdateProgress(taskID, oldProgress)
	return nil
}
