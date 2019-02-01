package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyScalingConfiguration invokes the ess.ModifyScalingConfiguration API synchronously
// api document: https://help.aliyun.com/api/ess/modifyscalingconfiguration.html
func (client *Client) ModifyScalingConfiguration(request *ModifyScalingConfigurationRequest) (response *ModifyScalingConfigurationResponse, err error) {
	response = CreateModifyScalingConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingConfigurationWithChan invokes the ess.ModifyScalingConfiguration API asynchronously
// api document: https://help.aliyun.com/api/ess/modifyscalingconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingConfigurationWithChan(request *ModifyScalingConfigurationRequest) (<-chan *ModifyScalingConfigurationResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingConfiguration(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyScalingConfigurationWithCallback invokes the ess.ModifyScalingConfiguration API asynchronously
// api document: https://help.aliyun.com/api/ess/modifyscalingconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScalingConfigurationWithCallback(request *ModifyScalingConfigurationRequest, callback func(response *ModifyScalingConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingConfigurationResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingConfiguration(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyScalingConfigurationRequest is the request struct for api ModifyScalingConfiguration
type ModifyScalingConfigurationRequest struct {
	*requests.RpcRequest
	ImageId                  string                                      `position:"Query" name:"ImageId"`
	Memory                   requests.Integer                            `position:"Query" name:"Memory"`
	IoOptimized              string                                      `position:"Query" name:"IoOptimized"`
	InstanceTypes            *[]string                                   `position:"Query" name:"InstanceTypes"  type:"Repeated"`
	InternetMaxBandwidthOut  requests.Integer                            `position:"Query" name:"InternetMaxBandwidthOut"`
	SecurityGroupId          string                                      `position:"Query" name:"SecurityGroupId"`
	KeyPairName              string                                      `position:"Query" name:"KeyPairName"`
	SpotPriceLimit           *[]ModifyScalingConfigurationSpotPriceLimit `position:"Query" name:"SpotPriceLimit"  type:"Repeated"`
	SystemDiskCategory       string                                      `position:"Query" name:"SystemDisk.Category"`
	UserData                 string                                      `position:"Query" name:"UserData"`
	HostName                 string                                      `position:"Query" name:"HostName"`
	PasswordInherit          requests.Boolean                            `position:"Query" name:"PasswordInherit"`
	ImageName                string                                      `position:"Query" name:"ImageName"`
	Override                 requests.Boolean                            `position:"Query" name:"Override"`
	DeploymentSetId          string                                      `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerAccount     string                                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string                                      `position:"Query" name:"OwnerAccount"`
	Cpu                      requests.Integer                            `position:"Query" name:"Cpu"`
	RamRoleName              string                                      `position:"Query" name:"RamRoleName"`
	OwnerId                  requests.Integer                            `position:"Query" name:"OwnerId"`
	DataDisk                 *[]ModifyScalingConfigurationDataDisk       `position:"Query" name:"DataDisk"  type:"Repeated"`
	ScalingConfigurationName string                                      `position:"Query" name:"ScalingConfigurationName"`
	Tags                     string                                      `position:"Query" name:"Tags"`
	ScalingConfigurationId   string                                      `position:"Query" name:"ScalingConfigurationId"`
	SpotStrategy             string                                      `position:"Query" name:"SpotStrategy"`
	InstanceName             string                                      `position:"Query" name:"InstanceName"`
	LoadBalancerWeight       requests.Integer                            `position:"Query" name:"LoadBalancerWeight"`
	SystemDiskSize           requests.Integer                            `position:"Query" name:"SystemDisk.Size"`
	InternetChargeType       string                                      `position:"Query" name:"InternetChargeType"`
}

// ModifyScalingConfigurationSpotPriceLimit is a repeated param struct in ModifyScalingConfigurationRequest
type ModifyScalingConfigurationSpotPriceLimit struct {
	InstanceType string `name:"InstanceType"`
	PriceLimit   string `name:"PriceLimit"`
}

// ModifyScalingConfigurationDataDisk is a repeated param struct in ModifyScalingConfigurationRequest
type ModifyScalingConfigurationDataDisk struct {
	SnapshotId         string `name:"SnapshotId"`
	Size               string `name:"Size"`
	Category           string `name:"Category"`
	Device             string `name:"Device"`
	DeleteWithInstance string `name:"DeleteWithInstance"`
}

// ModifyScalingConfigurationResponse is the response struct for api ModifyScalingConfiguration
type ModifyScalingConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyScalingConfigurationRequest creates a request to invoke ModifyScalingConfiguration API
func CreateModifyScalingConfigurationRequest() (request *ModifyScalingConfigurationRequest) {
	request = &ModifyScalingConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingConfiguration", "ess", "openAPI")
	return
}

// CreateModifyScalingConfigurationResponse creates a response to parse from ModifyScalingConfiguration response
func CreateModifyScalingConfigurationResponse() (response *ModifyScalingConfigurationResponse) {
	response = &ModifyScalingConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
