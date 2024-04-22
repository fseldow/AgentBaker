package nbcontractv1

import (
	"log"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
)

func TestNewNBContractBuilder(t *testing.T) {
	wantedResult := Configuration{
		Version:          contractVersion,
		KubeBinaryConfig: &KubeBinaryConfig{},
		ApiServerConfig:  &ApiServerConfig{},
		AuthConfig:       &AuthConfig{},
		ClusterConfig: &ClusterConfig{
			LoadBalancerConfig:   &LoadBalancerConfig{},
			ClusterNetworkConfig: &ClusterNetworkConfig{},
		},
		GpuConfig:              &GPUConfig{},
		TlsBootstrappingConfig: &TLSBootstrappingConfig{},
		KubeletConfig:          &KubeletConfig{},
		RuncConfig:             &RuncConfig{},
		ContainerdConfig:       &ContainerdConfig{},
		TeleportConfig:         &TeleportConfig{},
		CustomLinuxOsConfig: &CustomLinuxOSConfig{
			SysctlConfig: &SysctlConfig{},
			UlimitConfig: &UlimitConfig{},
		},
		HttpProxyConfig:          &HTTPProxyConfig{},
		CustomCloudConfig:        &CustomCloudConfig{},
		CustomSearchDomainConfig: &CustomSearchDomainConfig{},
	}
	tests := []struct {
		name string
		want *Configuration
	}{
		{
			name: "Test with nil configuration",
			want: &wantedResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNBContractBuilder().nodeBootstrapConfig; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNBContractConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNBContractBuilder_ApplyConfiguration(t *testing.T) {
	type fields struct {
		nBContractConfiguration *Configuration
	}
	wantedResult := &Configuration{
		Version:          contractVersion,
		KubeBinaryConfig: &KubeBinaryConfig{},
		ApiServerConfig:  &ApiServerConfig{},
		AuthConfig:       &AuthConfig{},
		ClusterConfig: &ClusterConfig{
			LoadBalancerConfig:   &LoadBalancerConfig{},
			ClusterNetworkConfig: &ClusterNetworkConfig{},
		},
		GpuConfig:              &GPUConfig{},
		TlsBootstrappingConfig: &TLSBootstrappingConfig{},
		KubeletConfig:          &KubeletConfig{},
		RuncConfig:             &RuncConfig{},
		ContainerdConfig:       &ContainerdConfig{},
		TeleportConfig:         &TeleportConfig{},
		CustomLinuxOsConfig: &CustomLinuxOSConfig{
			SysctlConfig: &SysctlConfig{},
			UlimitConfig: &UlimitConfig{},
		},
		HttpProxyConfig:          &HTTPProxyConfig{},
		CustomCloudConfig:        &CustomCloudConfig{},
		CustomSearchDomainConfig: &CustomSearchDomainConfig{},
	}
	tests := []struct {
		name   string
		fields fields
		want   *Configuration
	}{
		{
			name: "Test with nil configuration",
			fields: fields{
				nBContractConfiguration: &Configuration{},
			},
			want: wantedResult,
		},
		{
			name: "Apply nil AuthConfig configuration and expect AuthConfig in nBContractConfiguration to be non-nil",
			fields: fields{
				nBContractConfiguration: &Configuration{
					AuthConfig: nil,
				},
			},
			want: wantedResult,
		},
		{
			name: "Apply some configurations and expect them to be applied",
			fields: fields{
				nBContractConfiguration: &Configuration{
					AuthConfig: &AuthConfig{
						TargetCloud: "some-cloud",
					},
					LinuxAdminUsername: "testuser",
				},
			},
			want: func() *Configuration {
				tmpResult := NewNBContractBuilder().nodeBootstrapConfig
				tmpResult.AuthConfig.TargetCloud = "some-cloud"
				tmpResult.LinuxAdminUsername = "testuser"
				return tmpResult
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewNBContractBuilder()
			builder.ApplyConfiguration(tt.fields.nBContractConfiguration)
			if got := builder.nodeBootstrapConfig; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApplyConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNBContractBuilder_deepCopy(t *testing.T) {
	type Teststruct struct {
		A string
		B *int
		C bool
	}
	type args struct {
		src *Teststruct
		dst *Teststruct
	}
	tests := []struct {
		name        string
		args        args
		want        *Teststruct
		isChangeDst bool
	}{
		{
			name: "Test with empty src",
			args: args{
				src: &Teststruct{},
				dst: &Teststruct{},
			},
			want:        &Teststruct{},
			isChangeDst: false,
		},
		{
			name: "Test with non-empty src",
			args: args{
				src: &Teststruct{
					A: "some-string",
					B: to.IntPtr(123),
					C: true,
				},
				dst: &Teststruct{},
			},
			want: &Teststruct{
				A: "some-string",
				B: to.IntPtr(123),
				C: true,
			},
			isChangeDst: false,
		},
		{
			name: "Test with dst which has some existing values. Expect them to be overwritten",
			args: args{
				src: &Teststruct{
					A: "some-string",
					B: to.IntPtr(123),
				},
				dst: &Teststruct{
					A: "some-other-string",
					B: to.IntPtr(456),
					C: false,
				},
			},
			want: &Teststruct{
				A: "some-string",
				B: to.IntPtr(123),
				C: false,
			},
			isChangeDst: false,
		},
		{
			name: "After deepCopy, changes in dst should not affect src",
			args: args{
				src: &Teststruct{
					A: "some-string",
					B: to.IntPtr(123),
				},
				dst: &Teststruct{},
			},
			want: &Teststruct{
				A: "some-string",
				B: to.IntPtr(123),
			},
			isChangeDst: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nBCB := &NBContractBuilder{
				nodeBootstrapConfig: &Configuration{},
			}
			if err := nBCB.deepCopy(tt.args.src, tt.args.dst); err != nil {
				log.Printf("Failed to deep copy the configuration: %v", err)
			}
			log.Printf("dst = %v, src %v", tt.args.dst, tt.args.src)
			if got := tt.args.dst; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
			if tt.isChangeDst {
				tt.args.dst.A = "some-other-string"
				tt.args.dst.B = to.IntPtr(456)
				if got := tt.args.src; !reflect.DeepEqual(got, tt.want) {
					t.Errorf("src = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestNBContractBuilder_validateSemVer(t *testing.T) {
	type fields struct {
		nodeBootstrapConfig *Configuration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test with nil version and expect error",
			fields: fields{
				nodeBootstrapConfig: &Configuration{},
			},
			wantErr: true,
		},
		{
			name: "Test with invalid version and expect error",
			fields: fields{
				nodeBootstrapConfig: &Configuration{
					Version: "some-invalid-version",
				},
			},
			wantErr: true,
		},
		{
			name: "Test with valid version and expect no error",
			fields: fields{
				nodeBootstrapConfig: &Configuration{
					Version: contractVersion,
				},
			},
			wantErr: false,
		},
		{
			name: "Test with mismatch major version and expect error",
			fields: fields{
				nodeBootstrapConfig: &Configuration{
					Version: "v2.0.0",
				},
			},
			wantErr: true,
		},
		{
			name: "Test with mismatch minor version and expect no error",
			fields: fields{
				nodeBootstrapConfig: &Configuration{
					Version: "v1.1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nBCB := NewNBContractBuilder()
			nBCB.ApplyConfiguration(tt.fields.nodeBootstrapConfig)
			nBCB.nodeBootstrapConfig.Version = tt.fields.nodeBootstrapConfig.Version
			if err := nBCB.validateSemVer(); (err != nil) != tt.wantErr {
				t.Errorf("NBContractBuilder.validateSemVer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}