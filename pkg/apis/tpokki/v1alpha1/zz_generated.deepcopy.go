// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatlingTask) DeepCopyInto(out *GatlingTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatlingTask.
func (in *GatlingTask) DeepCopy() *GatlingTask {
	if in == nil {
		return nil
	}
	out := new(GatlingTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatlingTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatlingTaskList) DeepCopyInto(out *GatlingTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GatlingTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatlingTaskList.
func (in *GatlingTaskList) DeepCopy() *GatlingTaskList {
	if in == nil {
		return nil
	}
	out := new(GatlingTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatlingTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatlingTaskSpec) DeepCopyInto(out *GatlingTaskSpec) {
	*out = *in
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	in.ScenarioSpec.DeepCopyInto(&out.ScenarioSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatlingTaskSpec.
func (in *GatlingTaskSpec) DeepCopy() *GatlingTaskSpec {
	if in == nil {
		return nil
	}
	out := new(GatlingTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatlingTaskStatus) DeepCopyInto(out *GatlingTaskStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatlingTaskStatus.
func (in *GatlingTaskStatus) DeepCopy() *GatlingTaskStatus {
	if in == nil {
		return nil
	}
	out := new(GatlingTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSourceSpec) DeepCopyInto(out *GitSourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSourceSpec.
func (in *GitSourceSpec) DeepCopy() *GitSourceSpec {
	if in == nil {
		return nil
	}
	out := new(GitSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScenarioSpec) DeepCopyInto(out *ScenarioSpec) {
	*out = *in
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.GitSource = in.GitSource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScenarioSpec.
func (in *ScenarioSpec) DeepCopy() *ScenarioSpec {
	if in == nil {
		return nil
	}
	out := new(ScenarioSpec)
	in.DeepCopyInto(out)
	return out
}
