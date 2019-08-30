package drds

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

// RefreshJstMigrateDrdsDbAtomUrl invokes the drds.RefreshJstMigrateDrdsDbAtomUrl API synchronously
// api document: https://help.aliyun.com/api/drds/refreshjstmigratedrdsdbatomurl.html
func (client *Client) RefreshJstMigrateDrdsDbAtomUrl(request *RefreshJstMigrateDrdsDbAtomUrlRequest) (response *RefreshJstMigrateDrdsDbAtomUrlResponse, err error) {
	response = CreateRefreshJstMigrateDrdsDbAtomUrlResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshJstMigrateDrdsDbAtomUrlWithChan invokes the drds.RefreshJstMigrateDrdsDbAtomUrl API asynchronously
// api document: https://help.aliyun.com/api/drds/refreshjstmigratedrdsdbatomurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshJstMigrateDrdsDbAtomUrlWithChan(request *RefreshJstMigrateDrdsDbAtomUrlRequest) (<-chan *RefreshJstMigrateDrdsDbAtomUrlResponse, <-chan error) {
	responseChan := make(chan *RefreshJstMigrateDrdsDbAtomUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshJstMigrateDrdsDbAtomUrl(request)
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

// RefreshJstMigrateDrdsDbAtomUrlWithCallback invokes the drds.RefreshJstMigrateDrdsDbAtomUrl API asynchronously
// api document: https://help.aliyun.com/api/drds/refreshjstmigratedrdsdbatomurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefreshJstMigrateDrdsDbAtomUrlWithCallback(request *RefreshJstMigrateDrdsDbAtomUrlRequest, callback func(response *RefreshJstMigrateDrdsDbAtomUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshJstMigrateDrdsDbAtomUrlResponse
		var err error
		defer close(result)
		response, err = client.RefreshJstMigrateDrdsDbAtomUrl(request)
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

// RefreshJstMigrateDrdsDbAtomUrlRequest is the request struct for api RefreshJstMigrateDrdsDbAtomUrl
type RefreshJstMigrateDrdsDbAtomUrlRequest struct {
	*requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// RefreshJstMigrateDrdsDbAtomUrlResponse is the response struct for api RefreshJstMigrateDrdsDbAtomUrl
type RefreshJstMigrateDrdsDbAtomUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateRefreshJstMigrateDrdsDbAtomUrlRequest creates a request to invoke RefreshJstMigrateDrdsDbAtomUrl API
func CreateRefreshJstMigrateDrdsDbAtomUrlRequest() (request *RefreshJstMigrateDrdsDbAtomUrlRequest) {
	request = &RefreshJstMigrateDrdsDbAtomUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "RefreshJstMigrateDrdsDbAtomUrl", "drds", "openAPI")
	return
}

// CreateRefreshJstMigrateDrdsDbAtomUrlResponse creates a response to parse from RefreshJstMigrateDrdsDbAtomUrl response
func CreateRefreshJstMigrateDrdsDbAtomUrlResponse() (response *RefreshJstMigrateDrdsDbAtomUrlResponse) {
	response = &RefreshJstMigrateDrdsDbAtomUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
