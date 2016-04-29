package jsonprocessor

type BuildType struct{
	Language string
	BuildTool string
	Directory string
	Args []string
}

type jsonobject struct {
	Build []BuildType
}
