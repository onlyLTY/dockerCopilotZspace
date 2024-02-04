package utiles

import (
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

func GetContainerInspect(ctx *svc.ServiceContext, id string) (types.ContainerJSON, error) {
	jwtToken, endpointsId, err := GetNewJwt(ctx)
	if err != nil {
		logx.Errorf("GetNewJwt error: %v", err)
		return types.ContainerJSON{}, err
	}
	client := NewCustomClient(jwtToken)
	baseURL := domain + "/api/endpoints/" + endpointsId
	url := baseURL + "/docker/containers/" + id + "/json"
	resp, err := client.SendRequest("GET", url, nil)
	if err != nil {
		logx.Errorf("SendRequest error: %v", err)
		return types.ContainerJSON{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logx.Error("ioutil.ReadAll error: %v", err)
		return types.ContainerJSON{}, err
	}
	var containerInspect types.ContainerJSON
	err = json.Unmarshal(body, &containerInspect)
	if err != nil {
		return types.ContainerJSON{}, err
	}
	return containerInspect, nil
}
