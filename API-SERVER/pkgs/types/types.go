package types


type Job struct {
	TypeMeta
	Spec JobSpec
}


type TypeMeta struct {

	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`

	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,2,opt,name=apiVersion"`
}



type JobSpec struct {
	Containers []Container
	RestartPolicy string


}

type Container struct {
	Name string
	Image string
	Command []string
	Args []string
	Env []EnvVar
}


type EnvVar struct {
	Name string
	Value string
	
}


type Namespace struct{
	TypeMeta
	Metadata Metadata `json:"metadata"`

}


type Metadata struct {
	Name string `json:"name"`

}
