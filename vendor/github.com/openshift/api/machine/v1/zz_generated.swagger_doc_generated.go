package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_AlibabaCloudMachineProviderConfig = map[string]string{
	"":                  "AlibabaCloudMachineProviderConfig is the Schema for the alibabacloudmachineproviderconfig API Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"instanceType":      "The instance type of the instance.",
	"vpcId":             "The ID of the vpc",
	"regionId":          "The ID of the region in which to create the instance. You can call the DescribeRegions operation to query the most recent region list.",
	"zoneId":            "The ID of the zone in which to create the instance. You can call the DescribeZones operation to query the most recent region list.",
	"imageId":           "The ID of the image used to create the instance.",
	"dataDisk":          "DataDisks holds information regarding the extra disks attached to the instance",
	"securityGroups":    "SecurityGroups is a list of security group references to assign to the instance. A reference holds either the security group ID, the resource name, or the required tags to search. When more than one security group is returned for a tag search, all the groups are associated with the instance up to the maximum number of security groups to which an instance can belong. For more information, see the \"Security group limits\" section in Limits. https://www.alibabacloud.com/help/en/doc-detail/25412.htm",
	"bandwidth":         "Bandwidth describes the internet bandwidth strategy for the instance",
	"systemDisk":        "SystemDisk holds the properties regarding the system disk for the instance",
	"vSwitch":           "VSwitch is a reference to the vswitch to use for this instance. A reference holds either the vSwitch ID, the resource name, or the required tags to search. When more than one vSwitch is returned for a tag search, only the first vSwitch returned will be used. This parameter is required when you create an instance of the VPC type. You can call the DescribeVSwitches operation to query the created vSwitches.",
	"ramRoleName":       "RAMRoleName is the name of the instance Resource Access Management (RAM) role. This allows the instance to perform API calls as this specified RAM role.",
	"resourceGroup":     "ResourceGroup references the resource group to which to assign the instance. A reference holds either the resource group ID, the resource name, or the required tags to search. When more than one resource group are returned for a search, an error will be produced and the Machine will not be created. Resource Groups do not support searching by tags.",
	"tenancy":           "Tenancy specifies whether to create the instance on a dedicated host. Valid values:\n\ndefault: creates the instance on a non-dedicated host. host: creates the instance on a dedicated host. If you do not specify the DedicatedHostID parameter, Alibaba Cloud automatically selects a dedicated host for the instance. Empty value means no opinion and the platform chooses the a default, which is subject to change over time. Currently the default is `default`.",
	"userDataSecret":    "UserDataSecret contains a local reference to a secret that contains the UserData to apply to the instance",
	"credentialsSecret": "CredentialsSecret is a reference to the secret with alibabacloud credentials. Otherwise, defaults to permissions provided by attached RAM role where the actuator is running.",
	"tag":               "Tags are the set of metadata to add to an instance.",
}

func (AlibabaCloudMachineProviderConfig) SwaggerDoc() map[string]string {
	return map_AlibabaCloudMachineProviderConfig
}

var map_AlibabaCloudMachineProviderConfigList = map[string]string{
	"": "AlibabaCloudMachineProviderConfigList contains a list of AlibabaCloudMachineProviderConfig Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (AlibabaCloudMachineProviderConfigList) SwaggerDoc() map[string]string {
	return map_AlibabaCloudMachineProviderConfigList
}

var map_AlibabaCloudMachineProviderStatus = map[string]string{
	"":              "AlibabaCloudMachineProviderStatus is the Schema for the alibabacloudmachineproviderconfig API Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"instanceId":    "InstanceID is the instance ID of the machine created in alibabacloud",
	"instanceState": "InstanceState is the state of the alibabacloud instance for this machine",
	"conditions":    "Conditions is a set of conditions associated with the Machine to indicate errors or other status",
}

func (AlibabaCloudMachineProviderStatus) SwaggerDoc() map[string]string {
	return map_AlibabaCloudMachineProviderStatus
}

var map_AlibabaResourceReference = map[string]string{
	"":     "ResourceTagReference is a reference to a specific AlibabaCloud resource by ID, or tags. Only one of ID or Tags may be specified. Specifying more than one will result in a validation error.",
	"type": "type identifies the resource reference type for this entry.",
	"id":   "ID of resource",
	"name": "Name of the resource",
	"tags": "Tags is a set of metadata based upon ECS object tags used to identify a resource. For details about usage when multiple resources are found, please see the owning parent field documentation.",
}

func (AlibabaResourceReference) SwaggerDoc() map[string]string {
	return map_AlibabaResourceReference
}

var map_BandwidthProperties = map[string]string{
	"":                        "Bandwidth describes the bandwidth strategy for the network of the instance",
	"internetMaxBandwidthIn":  "InternetMaxBandwidthIn is the maximum inbound public bandwidth. Unit: Mbit/s. Valid values: When the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of this parameter are 1 to 10. Currently the default is `10` when outbound bandwidth is less than or equal to 10 Mbit/s. When the purchased outbound public bandwidth is greater than 10, the valid values are 1 to the InternetMaxBandwidthOut value. Currently the default is the value used for `InternetMaxBandwidthOut` when outbound public bandwidth is greater than 10.",
	"internetMaxBandwidthOut": "InternetMaxBandwidthOut is the maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100. When a value greater than 0 is used then a public IP address is assigned to the instance. Empty value means no opinion and the platform chooses the a default, which is subject to change over time. Currently the default is `0`",
}

func (BandwidthProperties) SwaggerDoc() map[string]string {
	return map_BandwidthProperties
}

var map_DataDiskProperties = map[string]string{
	"":                 "DataDisk contains the information regarding the datadisk attached to an instance",
	"Name":             "Name is the name of data disk N. If the name is specified the name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).\n\nEmpty value means the platform chooses a default, which is subject to change over time. Currently the default is `\"\"`.",
	"SnapshotID":       "SnapshotID is the ID of the snapshot used to create data disk N. Valid values of N: 1 to 16.\n\nWhen the DataDisk.N.SnapshotID parameter is specified, the DataDisk.N.Size parameter is ignored. The data disk is created based on the size of the specified snapshot. Use snapshots created after July 15, 2013. Otherwise, an error is returned and your request is rejected.",
	"Size":             "Size of the data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:\n\nValid values when DataDisk.N.Category is set to cloud_efficiency: 20 to 32768 Valid values when DataDisk.N.Category is set to cloud_ssd: 20 to 32768 Valid values when DataDisk.N.Category is set to cloud_essd: 20 to 32768 Valid values when DataDisk.N.Category is set to cloud: 5 to 2000 The value of this parameter must be greater than or equal to the size of the snapshot specified by the SnapshotID parameter.",
	"DiskEncryption":   "DiskEncryption specifies whether to encrypt data disk N.\n\nEmpty value means the platform chooses a default, which is subject to change over time. Currently the default is `disabled`.",
	"PerformanceLevel": "PerformanceLevel is the performance level of the ESSD used as as data disk N.  The N value must be the same as that in DataDisk.N.Category when DataDisk.N.Category is set to cloud_essd. Empty value means no opinion and the platform chooses a default, which is subject to change over time. Currently the default is `PL1`. Valid values:\n\nPL0: A single ESSD can deliver up to 10,000 random read/write IOPS. PL1: A single ESSD can deliver up to 50,000 random read/write IOPS. PL2: A single ESSD can deliver up to 100,000 random read/write IOPS. PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS. For more information about ESSD performance levels, see ESSDs.",
	"Category":         "Category describes the type of data disk N. Valid values: cloud_efficiency: ultra disk cloud_ssd: standard SSD cloud_essd: ESSD cloud: basic disk Empty value means no opinion and the platform chooses the a default, which is subject to change over time. Currently for non-I/O optimized instances of retired instance types, the default is `cloud`. Currently for other instances, the default is `cloud_efficiency`.",
	"KMSKeyID":         "KMSKeyID is the ID of the Key Management Service (KMS) key to be used by data disk N. Empty value means no opinion and the platform chooses the a default, which is subject to change over time. Currently the default is `\"\"` which is interpreted as do not use KMSKey encryption.",
	"DiskPreservation": "DiskPreservation specifies whether to release data disk N along with the instance. Empty value means no opinion and the platform chooses the a default, which is subject to change over time. Currently the default is `DeleteWithInstance`",
}

func (DataDiskProperties) SwaggerDoc() map[string]string {
	return map_DataDiskProperties
}

var map_SystemDiskProperties = map[string]string{
	"":                 "SystemDiskProperties contains the information regarding the system disk including performance, size, name, and category",
	"category":         "Category is the category of the system disk. Valid values: cloud_essd: ESSD. When the parameter is set to this value, you can use the SystemDisk.PerformanceLevel parameter to specify the performance level of the disk. cloud_efficiency: ultra disk. cloud_ssd: standard SSD. cloud: basic disk. Empty value means no opinion and the platform chooses the a default, which is subject to change over time. Currently for non-I/O optimized instances of retired instance types, the default is `cloud`. Currently for other instances, the default is `cloud_efficiency`.",
	"performanceLevel": "PerformanceLevel is the performance level of the ESSD used as the system disk. Valid values:\n\nPL0: A single ESSD can deliver up to 10,000 random read/write IOPS. PL1: A single ESSD can deliver up to 50,000 random read/write IOPS. PL2: A single ESSD can deliver up to 100,000 random read/write IOPS. PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS. Empty value means no opinion and the platform chooses a default, which is subject to change over time. Currently the default is `PL1`. For more information about ESSD performance levels, see ESSDs.",
	"name":             "Name is the name of the system disk. If the name is specified the name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-). Empty value means the platform chooses a default, which is subject to change over time. Currently the default is `\"\"`.",
	"size":             "Size is the size of the system disk. Unit: GiB. Valid values: 20 to 500. The value must be at least 20 and greater than or equal to the size of the image. Empty value means the platform chooses a default, which is subject to change over time. Currently the default is `40` or the size of the image depending on whichever is greater.",
}

func (SystemDiskProperties) SwaggerDoc() map[string]string {
	return map_SystemDiskProperties
}

var map_Tag = map[string]string{
	"":      "Tag  The tags of ECS Instance",
	"Key":   "Key is the name of the key pair",
	"Value": "Value is the value or data of the key pair",
}

func (Tag) SwaggerDoc() map[string]string {
	return map_Tag
}

var map_AWSResourceFilter = map[string]string{
	"":       "AWSResourceFilter is a filter used to identify an AWS resource",
	"name":   "Name of the filter. Filter names are case-sensitive.",
	"values": "Values includes one or more filter values. Filter values are case-sensitive.",
}

func (AWSResourceFilter) SwaggerDoc() map[string]string {
	return map_AWSResourceFilter
}

var map_AWSResourceReference = map[string]string{
	"":        "AWSResourceReference is a reference to a specific AWS resource by ID, ARN, or filters. Only one of ID, ARN or Filters may be specified. Specifying more than one will result in a validation error.",
	"type":    "Type determines how the reference will fetch the AWS resource.",
	"id":      "ID of resource.",
	"arn":     "ARN of resource.",
	"filters": "Filters is a set of filters used to identify a resource.",
}

func (AWSResourceReference) SwaggerDoc() map[string]string {
	return map_AWSResourceReference
}

var map_AWSFailureDomain = map[string]string{
	"":          "AWSFailureDomain configures failure domain information for the AWS platform.",
	"subnet":    "Subnet is a reference to the subnet to use for this instance.",
	"placement": "Placement configures the placement information for this instance.",
}

func (AWSFailureDomain) SwaggerDoc() map[string]string {
	return map_AWSFailureDomain
}

var map_AWSFailureDomainPlacement = map[string]string{
	"":                 "AWSFailureDomainPlacement configures the placement information for the AWSFailureDomain.",
	"availabilityZone": "AvailabilityZone is the availability zone of the instance.",
}

func (AWSFailureDomainPlacement) SwaggerDoc() map[string]string {
	return map_AWSFailureDomainPlacement
}

var map_AzureFailureDomain = map[string]string{
	"":     "AzureFailureDomain configures failure domain information for the Azure platform.",
	"zone": "Availability Zone for the virtual machine. If nil, the virtual machine should be deployed to no zone.",
}

func (AzureFailureDomain) SwaggerDoc() map[string]string {
	return map_AzureFailureDomain
}

var map_ControlPlaneMachineSet = map[string]string{
	"": "ControlPlaneMachineSet ensures that a specified number of control plane machine replicas are running at any given time. Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ControlPlaneMachineSet) SwaggerDoc() map[string]string {
	return map_ControlPlaneMachineSet
}

var map_ControlPlaneMachineSetList = map[string]string{
	"": "ControlPlaneMachineSetList contains a list of ControlPlaneMachineSet Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (ControlPlaneMachineSetList) SwaggerDoc() map[string]string {
	return map_ControlPlaneMachineSetList
}

var map_ControlPlaneMachineSetSpec = map[string]string{
	"":         "ControlPlaneMachineSet represents the configuration of the ControlPlaneMachineSet.",
	"state":    "State defines whether the ControlPlaneMachineSet is Active or Inactive. When Inactive, the ControlPlaneMachineSet will not take any action on the state of the Machines within the cluster. When Active, the ControlPlaneMachineSet will reconcile the Machines and will update the Machines as necessary. Once Active, a ControlPlaneMachineSet cannot be made Inactive. To prevent further action please remove the ControlPlaneMachineSet.",
	"replicas": "Replicas defines how many Control Plane Machines should be created by this ControlPlaneMachineSet. This field is immutable and cannot be changed after cluster installation. The ControlPlaneMachineSet only operates with 3 or 5 node control planes, 3 and 5 are the only valid values for this field.",
	"strategy": "Strategy defines how the ControlPlaneMachineSet will update Machines when it detects a change to the ProviderSpec.",
	"selector": "Label selector for Machines. Existing Machines selected by this selector will be the ones affected by this ControlPlaneMachineSet. It must match the template's labels. This field is considered immutable after creation of the resource.",
	"template": "Template describes the Control Plane Machines that will be created by this ControlPlaneMachineSet.",
}

func (ControlPlaneMachineSetSpec) SwaggerDoc() map[string]string {
	return map_ControlPlaneMachineSetSpec
}

var map_ControlPlaneMachineSetStatus = map[string]string{
	"":                    "ControlPlaneMachineSetStatus represents the status of the ControlPlaneMachineSet CRD.",
	"conditions":          "Conditions represents the observations of the ControlPlaneMachineSet's current state. Known .status.conditions.type are: Available, Degraded and Progressing.",
	"observedGeneration":  "ObservedGeneration is the most recent generation observed for this ControlPlaneMachineSet. It corresponds to the ControlPlaneMachineSets's generation, which is updated on mutation by the API Server.",
	"replicas":            "Replicas is the number of Control Plane Machines created by the ControlPlaneMachineSet controller. Note that during update operations this value may differ from the desired replica count.",
	"readyReplicas":       "ReadyReplicas is the number of Control Plane Machines created by the ControlPlaneMachineSet controller which are ready. Note that this value may be higher than the desired number of replicas while rolling updates are in-progress.",
	"updatedReplicas":     "UpdatedReplicas is the number of non-terminated Control Plane Machines created by the ControlPlaneMachineSet controller that have the desired provider spec and are ready. This value is set to 0 when a change is detected to the desired spec. When the update strategy is RollingUpdate, this will also coincide with starting the process of updating the Machines. When the update strategy is OnDelete, this value will remain at 0 until a user deletes an existing replica and its replacement has become ready.",
	"unavailableReplicas": "UnavailableReplicas is the number of Control Plane Machines that are still required before the ControlPlaneMachineSet reaches the desired available capacity. When this value is non-zero, the number of ReadyReplicas is less than the desired Replicas.",
}

func (ControlPlaneMachineSetStatus) SwaggerDoc() map[string]string {
	return map_ControlPlaneMachineSetStatus
}

var map_ControlPlaneMachineSetStrategy = map[string]string{
	"":     "ControlPlaneMachineSetStrategy defines the strategy for applying updates to the Control Plane Machines managed by the ControlPlaneMachineSet.",
	"type": "Type defines the type of update strategy that should be used when updating Machines owned by the ControlPlaneMachineSet. Valid values are \"RollingUpdate\" and \"OnDelete\". The current default value is \"RollingUpdate\".",
}

func (ControlPlaneMachineSetStrategy) SwaggerDoc() map[string]string {
	return map_ControlPlaneMachineSetStrategy
}

var map_ControlPlaneMachineSetTemplate = map[string]string{
	"":                                      "ControlPlaneMachineSetTemplate is a template used by the ControlPlaneMachineSet to create the Machines that it will manage in the future. ",
	"machineType":                           "MachineType determines the type of Machines that should be managed by the ControlPlaneMachineSet. Currently, the only valid value is machines_v1beta1_machine_openshift_io.",
	"machines_v1beta1_machine_openshift_io": "OpenShiftMachineV1Beta1Machine defines the template for creating Machines from the v1beta1.machine.openshift.io API group.",
}

func (ControlPlaneMachineSetTemplate) SwaggerDoc() map[string]string {
	return map_ControlPlaneMachineSetTemplate
}

var map_ControlPlaneMachineSetTemplateObjectMeta = map[string]string{
	"":            "ControlPlaneMachineSetTemplateObjectMeta is a subset of the metav1.ObjectMeta struct. It allows users to specify labels and annotations that will be copied onto Machines created from this template.",
	"labels":      "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels. This field must contain both the 'machine.openshift.io/cluster-api-machine-role' and 'machine.openshift.io/cluster-api-machine-type' labels, both with a value of 'master'. It must also contain a label with the key 'machine.openshift.io/cluster-api-cluster'.",
	"annotations": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
}

func (ControlPlaneMachineSetTemplateObjectMeta) SwaggerDoc() map[string]string {
	return map_ControlPlaneMachineSetTemplateObjectMeta
}

var map_FailureDomains = map[string]string{
	"":         "FailureDomain represents the different configurations required to spread Machines across failure domains on different platforms.",
	"platform": "Platform identifies the platform for which the FailureDomain represents. Currently supported values are AWS, Azure, and GCP.",
	"aws":      "AWS configures failure domain information for the AWS platform.",
	"azure":    "Azure configures failure domain information for the Azure platform.",
	"gcp":      "GCP configures failure domain information for the GCP platform.",
}

func (FailureDomains) SwaggerDoc() map[string]string {
	return map_FailureDomains
}

var map_GCPFailureDomain = map[string]string{
	"":     "GCPFailureDomain configures failure domain information for the GCP platform",
	"zone": "Zone is the zone in which the GCP machine provider will create the VM.",
}

func (GCPFailureDomain) SwaggerDoc() map[string]string {
	return map_GCPFailureDomain
}

var map_OpenShiftMachineV1Beta1MachineTemplate = map[string]string{
	"":               "OpenShiftMachineV1Beta1MachineTemplate is a template for the ControlPlaneMachineSet to create Machines from the v1beta1.machine.openshift.io API group.",
	"failureDomains": "FailureDomains is the list of failure domains (sometimes called availability zones) in which the ControlPlaneMachineSet should balance the Control Plane Machines. This will be merged into the ProviderSpec given in the template. This field is optional on platforms that do not require placement information.",
	"metadata":       "ObjectMeta is the standard object metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata Labels are required to match the ControlPlaneMachineSet selector.",
	"spec":           "Spec contains the desired configuration of the Control Plane Machines. The ProviderSpec within contains platform specific details for creating the Control Plane Machines. The ProviderSe should be complete apart from the platform specific failure domain field. This will be overriden when the Machines are created based on the FailureDomains field.",
}

func (OpenShiftMachineV1Beta1MachineTemplate) SwaggerDoc() map[string]string {
	return map_OpenShiftMachineV1Beta1MachineTemplate
}

var map_NutanixMachineProviderConfig = map[string]string{
	"":                  "NutanixMachineProviderConfig is the Schema for the nutanixmachineproviderconfigs API Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"cluster":           "cluster is to identify the cluster (the Prism Element under management of the Prism Central), in which the Machine's VM will be created. The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
	"image":             "image is to identify the rhcos image uploaded to the Prism Central (PC) The image identifier (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
	"subnets":           "subnets holds a list of identifiers (one or more) of the cluster's network subnets for the Machine's VM to connect to. The subnet identifiers (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
	"vcpusPerSocket":    "vcpusPerSocket is the number of vCPUs per socket of the VM",
	"vcpuSockets":       "vcpuSockets is the number of vCPU sockets of the VM",
	"memorySize":        "memorySize is the memory size (in Quantity format) of the VM The minimum memorySize is 2Gi bytes",
	"systemDiskSize":    "systemDiskSize is size (in Quantity format) of the system disk of the VM The minimum systemDiskSize is 20Gi bytes",
	"userDataSecret":    "userDataSecret is a local reference to a secret that contains the UserData to apply to the VM",
	"credentialsSecret": "credentialsSecret is a local reference to a secret that contains the credentials data to access Nutanix PC client",
}

func (NutanixMachineProviderConfig) SwaggerDoc() map[string]string {
	return map_NutanixMachineProviderConfig
}

var map_NutanixMachineProviderStatus = map[string]string{
	"":           "NutanixMachineProviderStatus is the type that will be embedded in a Machine.Status.ProviderStatus field. It contains nutanix-specific status information. Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"conditions": "conditions is a set of conditions associated with the Machine to indicate errors or other status",
	"vmUUID":     "vmUUID is the Machine associated VM's UUID The field is missing before the VM is created. Once the VM is created, the field is filled with the VM's UUID and it will not change. The vmUUID is used to find the VM when updating the Machine status, and to delete the VM when the Machine is deleted.",
}

func (NutanixMachineProviderStatus) SwaggerDoc() map[string]string {
	return map_NutanixMachineProviderStatus
}

var map_NutanixResourceIdentifier = map[string]string{
	"":     "NutanixResourceIdentifier holds the identity of a Nutanix PC resource (cluster, image, subnet, etc.)",
	"type": "Type is the identifier type to use for this resource.",
	"uuid": "uuid is the UUID of the resource in the PC.",
	"name": "name is the resource name in the PC",
}

func (NutanixResourceIdentifier) SwaggerDoc() map[string]string {
	return map_NutanixResourceIdentifier
}

var map_PowerVSMachineProviderConfig = map[string]string{
	"":                  "PowerVSMachineProviderConfig is the type that will be embedded in a Machine.Spec.ProviderSpec field for a PowerVS virtual machine. It is used by the PowerVS machine actuator to create a single Machine.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"userDataSecret":    "userDataSecret contains a local reference to a secret that contains the UserData to apply to the instance.",
	"credentialsSecret": "credentialsSecret is a reference to the secret with IBM Cloud credentials.",
	"serviceInstance":   "serviceInstance is the reference to the Power VS service on which the server instance(VM) will be created. Power VS service is a container for all Power VS instances at a specific geographic region. serviceInstance can be created via IBM Cloud catalog or CLI. supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli. More detail about Power VS service instance. https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server",
	"image":             "image is to identify the rhcos image uploaded to IBM COS bucket which is used to create the instance. supported image identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli.",
	"network":           "network is the reference to the Network to use for this instance. supported network identifier in PowerVSResource are Name, ID and RegEx and that can be obtained from IBM Cloud UI or IBM Cloud cli.",
	"keyPairName":       "keyPairName is the name of the KeyPair to use for SSH. The key pair will be exposed to the instance via the instance metadata service. On boot, the OS will copy the public keypair into the authorized keys for the core user.",
	"systemType":        "systemType is the System type used to host the instance. systemType determines the number of cores and memory that is available. Few of the supported SystemTypes are s922,e880,e980. e880 systemType available only in Dallas Datacenters. e980 systemType available in Datacenters except Dallas and Washington. When omitted, this means that the user has no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is s922 which is generally available.",
	"processorType":     "processorType is the VM instance processor type. It must be set to one of the following values: Dedicated, Capped or Shared. Dedicated: resources are allocated for a specific client, The hypervisor makes a 1:1 binding of a partition’s processor to a physical processor core. Shared: Shared among other clients. Capped: Shared, but resources do not expand beyond those that are requested, the amount of CPU time is Capped to the value specified for the entitlement. if the processorType is selected as Dedicated, then processors value cannot be fractional. When omitted, this means that the user has no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is Shared.",
	"processors":        "processors is the number of virtual processors in a virtual machine. when the processorType is selected as Dedicated the processors value cannot be fractional. maximum value for the Processors depends on the selected SystemType. when SystemType is set to e880 or e980 maximum Processors value is 143. when SystemType is set to s922 maximum Processors value is 15. minimum value for Processors depends on the selected ProcessorType. when ProcessorType is set as Shared or Capped, The minimum processors is 0.5. when ProcessorType is set as Dedicated, The minimum processors is 1. When omitted, this means that the user has no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The default is set based on the selected ProcessorType. when ProcessorType selected as Dedicated, the default is set to 1. when ProcessorType selected as Shared or Capped, the default is set to 0.5.",
	"memoryGiB":         "memoryGiB is the size of a virtual machine's memory, in GiB. maximum value for the MemoryGiB depends on the selected SystemType. when SystemType is set to e880 maximum MemoryGiB value is 7463 GiB. when SystemType is set to e980 maximum MemoryGiB value is 15307 GiB. when SystemType is set to s922 maximum MemoryGiB value is 942 GiB. The minimum memory is 32 GiB. When omitted, this means the user has no opinion and the platform is left to choose a reasonable default, which is subject to change over time. The current default is 32.",
}

func (PowerVSMachineProviderConfig) SwaggerDoc() map[string]string {
	return map_PowerVSMachineProviderConfig
}

var map_PowerVSMachineProviderStatus = map[string]string{
	"":                  "PowerVSMachineProviderStatus is the type that will be embedded in a Machine.Status.ProviderStatus field. It contains PowerVS-specific status information.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"conditions":        "conditions is a set of conditions associated with the Machine to indicate errors or other status",
	"instanceId":        "instanceId is the instance ID of the machine created in PowerVS instanceId uniquely identifies a Power VS server instance(VM) under a Power VS service. This will help in updating or deleting a VM in Power VS Cloud",
	"serviceInstanceID": "serviceInstanceID is the reference to the Power VS ServiceInstance on which the machine instance will be created. serviceInstanceID uniquely identifies the Power VS service By setting serviceInstanceID it will become easy and efficient to fetch a server instance(VM) within Power VS Cloud.",
	"instanceState":     "instanceState is the state of the PowerVS instance for this machine Possible instance states are Active, Build, ShutOff, Reboot This is used to display additional information to user regarding instance current state",
}

func (PowerVSMachineProviderStatus) SwaggerDoc() map[string]string {
	return map_PowerVSMachineProviderStatus
}

var map_PowerVSResource = map[string]string{
	"":      "PowerVSResource is a reference to a specific PowerVS resource by ID, Name or RegEx Only one of ID, Name or RegEx may be specified. Specifying more than one will result in a validation error.",
	"type":  "Type identifies the resource type for this entry. Valid values are ID, Name and RegEx",
	"id":    "ID of resource",
	"name":  "Name of resource",
	"regex": "Regex to find resource Regex contains the pattern to match to find a resource",
}

func (PowerVSResource) SwaggerDoc() map[string]string {
	return map_PowerVSResource
}

var map_PowerVSSecretReference = map[string]string{
	"":     "PowerVSSecretReference contains enough information to locate the referenced secret inside the same namespace.",
	"name": "Name of the secret.",
}

func (PowerVSSecretReference) SwaggerDoc() map[string]string {
	return map_PowerVSSecretReference
}

// AUTO-GENERATED FUNCTIONS END HERE
