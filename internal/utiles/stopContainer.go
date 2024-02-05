package utiles

import (
	"encoding/json"
	"fmt"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
)

func StopContainer(ctx *svc.ServiceContext, id string) error {
	jwtToken, endpointsId, err := GetNewJwt(ctx)
	if err != nil {
		logx.Errorf("GetNewJwt error: %v", err)
		return err
	}
	client := NewCustomClient(jwtToken)
	baseURL := domain + "/api/endpoints/" + endpointsId
	url := baseURL + "/docker/containers/" + id + "/stop"
	resp, err := client.SendRequest("POST", url, nil)
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
	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotModified {
		errorResponse := ErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			// 在此处处理JSON解码错误
			logx.Errorf("json.NewDecoder error: %v", err)
			return err
		}
		logx.Errorf("StopContainer error: %v", errorResponse.Message)
		return fmt.Errorf("StopContainer error: %v", errorResponse.Message)
	}
	return nil
}
