package utiles

import (
	"encoding/json"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	MyType "github.com/onlyLTY/dockerCopilotZspace/zspace/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

func GetContainerList(ctx *svc.ServiceContext) ([]MyType.Container, error) {
	jwtToken, endpointsId, err := GetNewJwt(ctx)
	if err != nil {
		logx.Errorf("GetNewJwt error: %v", err)
		return nil, err
	}
	client := NewCustomClient(jwtToken)
	baseURL := domain + "/api/endpoints/" + endpointsId
	url := baseURL + "/docker/containers/json?all=true"
	resp, err := client.SendRequest("GET", url, nil)
	if err != nil {
		logx.Errorf("SendRequest error: %v", err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logx.Error("io.ReadAll error: %v", err)
		return nil, err
	}
	var dockerContainerList []MyType.Container
	err = json.Unmarshal(body, &dockerContainerList)
	if err != nil {
		logx.Errorf("json.Unmarshal error: %v", err)
		return nil, err
	}
	return dockerContainerList, nil
}

func CheckImageUpdate(ctx *svc.ServiceContext, containerListData []MyType.Container) []MyType.Container {
	for i, v := range containerListData {
		if _, ok := ctx.HubImageInfo.Data[v.ImageID]; ok {
			if ctx.HubImageInfo.Data[v.ImageID].NeedUpdate {
				containerListData[i].Update = true
			}
		}
	}
	return containerListData
}
