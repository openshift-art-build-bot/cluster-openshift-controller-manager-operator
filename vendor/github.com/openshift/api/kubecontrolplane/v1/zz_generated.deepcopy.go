// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	osin_v1 "github.com/openshift/api/osin/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatorConfig) DeepCopyInto(out *AggregatorConfig) {
	*out = *in
	out.ProxyClientInfo = in.ProxyClientInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatorConfig.
func (in *AggregatorConfig) DeepCopy() *AggregatorConfig {
	if in == nil {
		return nil
	}
	out := new(AggregatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Arguments) DeepCopyInto(out *Arguments) {
	{
		in := &in
		*out = make(Arguments, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Arguments.
func (in Arguments) DeepCopy() Arguments {
	if in == nil {
		return nil
	}
	out := new(Arguments)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAPIServerConfig) DeepCopyInto(out *KubeAPIServerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.GenericAPIServerConfig.DeepCopyInto(&out.GenericAPIServerConfig)
	in.AuthConfig.DeepCopyInto(&out.AuthConfig)
	out.AggregatorConfig = in.AggregatorConfig
	out.KubeletClientInfo = in.KubeletClientInfo
	in.UserAgentMatchingConfig.DeepCopyInto(&out.UserAgentMatchingConfig)
	out.ImagePolicyConfig = in.ImagePolicyConfig
	out.ProjectConfig = in.ProjectConfig
	if in.ServiceAccountPublicKeyFiles != nil {
		in, out := &in.ServiceAccountPublicKeyFiles, &out.ServiceAccountPublicKeyFiles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OAuthConfig != nil {
		in, out := &in.OAuthConfig, &out.OAuthConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(osin_v1.OAuthConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.APIServerArguments != nil {
		in, out := &in.APIServerArguments, &out.APIServerArguments
		*out = make(map[string]Arguments, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAPIServerConfig.
func (in *KubeAPIServerConfig) DeepCopy() *KubeAPIServerConfig {
	if in == nil {
		return nil
	}
	out := new(KubeAPIServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeAPIServerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAPIServerImagePolicyConfig) DeepCopyInto(out *KubeAPIServerImagePolicyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAPIServerImagePolicyConfig.
func (in *KubeAPIServerImagePolicyConfig) DeepCopy() *KubeAPIServerImagePolicyConfig {
	if in == nil {
		return nil
	}
	out := new(KubeAPIServerImagePolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAPIServerProjectConfig) DeepCopyInto(out *KubeAPIServerProjectConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAPIServerProjectConfig.
func (in *KubeAPIServerProjectConfig) DeepCopy() *KubeAPIServerProjectConfig {
	if in == nil {
		return nil
	}
	out := new(KubeAPIServerProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerConfig) DeepCopyInto(out *KubeControllerManagerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ServiceServingCert = in.ServiceServingCert
	out.ProjectConfig = in.ProjectConfig
	if in.ExtendedArguments != nil {
		in, out := &in.ExtendedArguments, &out.ExtendedArguments
		*out = make(map[string]Arguments, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerConfig.
func (in *KubeControllerManagerConfig) DeepCopy() *KubeControllerManagerConfig {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllerManagerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerProjectConfig) DeepCopyInto(out *KubeControllerManagerProjectConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerProjectConfig.
func (in *KubeControllerManagerProjectConfig) DeepCopy() *KubeControllerManagerProjectConfig {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerProjectConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConnectionInfo) DeepCopyInto(out *KubeletConnectionInfo) {
	*out = *in
	out.CertInfo = in.CertInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConnectionInfo.
func (in *KubeletConnectionInfo) DeepCopy() *KubeletConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(KubeletConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterAuthConfig) DeepCopyInto(out *MasterAuthConfig) {
	*out = *in
	if in.RequestHeader != nil {
		in, out := &in.RequestHeader, &out.RequestHeader
		if *in == nil {
			*out = nil
		} else {
			*out = new(RequestHeaderAuthenticationOptions)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.WebhookTokenAuthenticators != nil {
		in, out := &in.WebhookTokenAuthenticators, &out.WebhookTokenAuthenticators
		*out = make([]WebhookTokenAuthenticator, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterAuthConfig.
func (in *MasterAuthConfig) DeepCopy() *MasterAuthConfig {
	if in == nil {
		return nil
	}
	out := new(MasterAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestHeaderAuthenticationOptions) DeepCopyInto(out *RequestHeaderAuthenticationOptions) {
	*out = *in
	if in.ClientCommonNames != nil {
		in, out := &in.ClientCommonNames, &out.ClientCommonNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UsernameHeaders != nil {
		in, out := &in.UsernameHeaders, &out.UsernameHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GroupHeaders != nil {
		in, out := &in.GroupHeaders, &out.GroupHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraHeaderPrefixes != nil {
		in, out := &in.ExtraHeaderPrefixes, &out.ExtraHeaderPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestHeaderAuthenticationOptions.
func (in *RequestHeaderAuthenticationOptions) DeepCopy() *RequestHeaderAuthenticationOptions {
	if in == nil {
		return nil
	}
	out := new(RequestHeaderAuthenticationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceServingCert) DeepCopyInto(out *ServiceServingCert) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceServingCert.
func (in *ServiceServingCert) DeepCopy() *ServiceServingCert {
	if in == nil {
		return nil
	}
	out := new(ServiceServingCert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAgentDenyRule) DeepCopyInto(out *UserAgentDenyRule) {
	*out = *in
	in.UserAgentMatchRule.DeepCopyInto(&out.UserAgentMatchRule)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAgentDenyRule.
func (in *UserAgentDenyRule) DeepCopy() *UserAgentDenyRule {
	if in == nil {
		return nil
	}
	out := new(UserAgentDenyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAgentMatchRule) DeepCopyInto(out *UserAgentMatchRule) {
	*out = *in
	if in.HTTPVerbs != nil {
		in, out := &in.HTTPVerbs, &out.HTTPVerbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAgentMatchRule.
func (in *UserAgentMatchRule) DeepCopy() *UserAgentMatchRule {
	if in == nil {
		return nil
	}
	out := new(UserAgentMatchRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAgentMatchingConfig) DeepCopyInto(out *UserAgentMatchingConfig) {
	*out = *in
	if in.RequiredClients != nil {
		in, out := &in.RequiredClients, &out.RequiredClients
		*out = make([]UserAgentMatchRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DeniedClients != nil {
		in, out := &in.DeniedClients, &out.DeniedClients
		*out = make([]UserAgentDenyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAgentMatchingConfig.
func (in *UserAgentMatchingConfig) DeepCopy() *UserAgentMatchingConfig {
	if in == nil {
		return nil
	}
	out := new(UserAgentMatchingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookTokenAuthenticator) DeepCopyInto(out *WebhookTokenAuthenticator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookTokenAuthenticator.
func (in *WebhookTokenAuthenticator) DeepCopy() *WebhookTokenAuthenticator {
	if in == nil {
		return nil
	}
	out := new(WebhookTokenAuthenticator)
	in.DeepCopyInto(out)
	return out
}
