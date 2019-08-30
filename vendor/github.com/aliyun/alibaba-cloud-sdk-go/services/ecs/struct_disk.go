package ecs

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

// Disk is a nested struct in ecs response
type Disk struct {
	Category                      string                        `json:"Category" xml:"Category"`
	BdfId                         string                        `json:"BdfId" xml:"BdfId"`
	ImageId                       string                        `json:"ImageId" xml:"ImageId"`
	AutoSnapshotPolicyId          string                        `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
	DeleteAutoSnapshot            bool                          `json:"DeleteAutoSnapshot" xml:"DeleteAutoSnapshot"`
	EnableAutomatedSnapshotPolicy bool                          `json:"EnableAutomatedSnapshotPolicy" xml:"EnableAutomatedSnapshotPolicy"`
	DiskId                        string                        `json:"DiskId" xml:"DiskId"`
	Size                          int                           `json:"Size" xml:"Size"`
	IOPS                          int                           `json:"IOPS" xml:"IOPS"`
	RegionId                      string                        `json:"RegionId" xml:"RegionId"`
	MountInstanceNum              int                           `json:"MountInstanceNum" xml:"MountInstanceNum"`
	StorageSetId                  string                        `json:"StorageSetId" xml:"StorageSetId"`
	ResourceGroupId               string                        `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InstanceId                    string                        `json:"InstanceId" xml:"InstanceId"`
	Description                   string                        `json:"Description" xml:"Description"`
	Type                          string                        `json:"Type" xml:"Type"`
	ExpiredTime                   string                        `json:"ExpiredTime" xml:"ExpiredTime"`
	Device                        string                        `json:"Device" xml:"Device"`
	CreationTime                  string                        `json:"CreationTime" xml:"CreationTime"`
	IOPSRead                      int                           `json:"IOPSRead" xml:"IOPSRead"`
	SourceSnapshotId              string                        `json:"SourceSnapshotId" xml:"SourceSnapshotId"`
	StorageSetPartitionNumber     int                           `json:"StorageSetPartitionNumber" xml:"StorageSetPartitionNumber"`
	ProductCode                   string                        `json:"ProductCode" xml:"ProductCode"`
	Portable                      bool                          `json:"Portable" xml:"Portable"`
	KMSKeyId                      string                        `json:"KMSKeyId" xml:"KMSKeyId"`
	Encrypted                     bool                          `json:"Encrypted" xml:"Encrypted"`
	EnableAutoSnapshot            bool                          `json:"EnableAutoSnapshot" xml:"EnableAutoSnapshot"`
	DetachedTime                  string                        `json:"DetachedTime" xml:"DetachedTime"`
	DeleteWithInstance            bool                          `json:"DeleteWithInstance" xml:"DeleteWithInstance"`
	ZoneId                        string                        `json:"ZoneId" xml:"ZoneId"`
	DiskChargeType                string                        `json:"DiskChargeType" xml:"DiskChargeType"`
	PerformanceLevel              string                        `json:"PerformanceLevel" xml:"PerformanceLevel"`
	DiskName                      string                        `json:"DiskName" xml:"DiskName"`
	Status                        string                        `json:"Status" xml:"Status"`
	AttachedTime                  string                        `json:"AttachedTime" xml:"AttachedTime"`
	IOPSWrite                     int                           `json:"IOPSWrite" xml:"IOPSWrite"`
	Tags                          TagsInDescribeDisks           `json:"Tags" xml:"Tags"`
	MountInstances                MountInstances                `json:"MountInstances" xml:"MountInstances"`
	OperationLocks                OperationLocksInDescribeDisks `json:"OperationLocks" xml:"OperationLocks"`
}
