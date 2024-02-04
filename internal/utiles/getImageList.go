package utiles

import (
	"encoding/json"
	"fmt"
	DockerTypes "github.com/docker/docker/api/types"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	MyType "github.com/onlyLTY/dockerCopilotZspace/zspace/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"strings"
)

func GetImagesList(ctx *svc.ServiceContext) ([]MyType.Image, error) {
	jwtToken, endpointsId, err := GetNewJwt(ctx)
	if err != nil {
		logx.Errorf("GetNewJwt error: %v", err)
		return nil, err
	}
	client := NewCustomClient(jwtToken)
	baseURL := domain + "/api/endpoints/" + endpointsId
	url := baseURL + "/docker/images/json"
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
	var imageSummaryList []DockerTypes.ImageSummary
	err = json.Unmarshal(body, &imageSummaryList)
	if err != nil {
		logx.Errorf("json.Unmarshal error: %v", err)
		return nil, err
	}
	var imagesList []MyType.Image
	for _, img := range imageSummaryList {
		image := MyType.Image{
			ImageSummary: img,
			ImageName:    "",
			ImageTag:     "",
			InUsed:       false,
			SizeFormat:   "",
		}
		imagesList = append(imagesList, image)
	}
	//看不明白就不要看了，这内存反复地申请，如果你看明白了 给这改成指针吧，啥？我为啥不直接写指针，我懒癌犯了就这样，欢迎pr
	imagesList, err = checkImageInUsed(ctx, splitImageNameAndTag(calculateImageSize(imagesList)))
	if err != nil {
		return imagesList, err
	}
	return imagesList, nil
}

func splitImageNameAndTag(imagesList []MyType.Image) []MyType.Image {
	for i, image := range imagesList {
		if len(image.RepoTags) != 0 {
			imagesList[i].ImageName = strings.Split(image.RepoTags[0], ":")[0]
			imagesList[i].ImageTag = strings.Split(image.RepoTags[0], ":")[1]
		} else if len(image.RepoDigests) != 0 {
			imagesList[i].ImageName = strings.Split(image.RepoDigests[0], "@")[0]
			imagesList[i].ImageTag = "None"
		} else {
			imagesList[i].ImageName = "None"
			imagesList[i].ImageTag = "None"
		}
	}
	return imagesList
}
func checkImageInUsed(svc *svc.ServiceContext, imageList []MyType.Image) ([]MyType.Image, error) {
	list, err := GetContainerList(svc)
	if err != nil {
		return imageList, err
	}
	// 这里可以用mapreduce 我懒等pr
	for _, v := range list {
		for i, image := range imageList {
			if v.ImageID == image.ID {
				imageList[i].InUsed = true
				break
			}
		}
	}
	return imageList, nil
}
func calculateImageSize(imagesList []MyType.Image) []MyType.Image {
	for i := range imagesList {
		if imagesList[i].Size >= 1024*1024*1024 {
			imagesList[i].SizeFormat = // Convert size to gigabytes
				fmt.Sprintf("%d Gb", imagesList[i].Size/1024/1024/1024)
		} else {
			imagesList[i].SizeFormat = // Convert size to megabytes
				fmt.Sprintf("%d Mb", imagesList[i].Size/1024/1024)
		}
	}
	return imagesList
}
