// Copyright 2016 Marapongo, Inc. All rights reserved.

package clouds

// Arch selects a cloud infrastructure to target when compiling.
type Arch int

const (
	NoArch     Arch = iota // no target specified.
	AWSArch                // Amazon Web Services.
	GCPArch                // Google Cloud Platform.
	AzureArch              // Microsoft Azure.
	VMWareArch             // VMWare vSphere, etc.
)

const (
	noArch     = ""
	awsArch    = "aws"
	gcpArch    = "gcp"
	azureArch  = "azure"
	vmwareArch = "vmware"
)

// ArchMap maps human-friendly names to the Archs for those names.
var ArchMap = map[string]Arch{
	noArch:     NoArch,
	awsArch:    AWSArch,
	gcpArch:    GCPArch,
	azureArch:  AzureArch,
	vmwareArch: VMWareArch,
}

// ArchNames maps Archs to human-friendly names.
var ArchNames = map[Arch]string{
	NoArch:     noArch,
	AWSArch:    awsArch,
	GCPArch:    gcpArch,
	AzureArch:  azureArch,
	VMWareArch: vmwareArch,
}
