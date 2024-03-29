package utiles

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
)

type configWrapper struct {
	*container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
}

func CreateContainer(ctx *svc.ServiceContext, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) error {
	jwtToken, endpointsId, err := GetNewJwt(ctx)
	if err != nil {
		logx.Errorf("GetNewJwt error: %v", err)
		return err
	}
	client := NewCustomClient(jwtToken)
	baseURL := domain + "/api/endpoints/" + endpointsId
	url := baseURL + "/docker/containers/create?name=" + containerName
	body := configWrapper{
		Config:           config,
		HostConfig:       hostConfig,
		NetworkingConfig: networkingConfig,
	}
	postData, err := json.Marshal(body)
	if err != nil {
		logx.Errorf("json.Marshal error: %v", err)
		return err
	}
	resp, err := client.SendRequest("POST", url, postData)
	if err != nil {
		logx.Errorf("SendRequest error: %v", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logx.Errorf("Body.Close error: %v", err)
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotModified {
		errorResponse := ErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			// 在此处处理JSON解码错误
			logx.Errorf("json.NewDecoder error: %v", err)
			return err
		}
		logx.Errorf("CreateContainer error: %v", errorResponse.Message)
		return fmt.Errorf("CreateContainer error: %v", errorResponse.Message)
	}
	return nil
}
