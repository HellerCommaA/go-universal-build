# Go Universal Build System
This is a universal `build` system powered by Go Lang.  
All things considered, it's an expirement by me to further my knowledge of Go and building a semi-large-scale-system in the language.  

The intent is that you can pass a `build.json` file along to the compiled executable to build your multiple projects simultaneiously and have them output to a common directory.  

This started as an idea I had while trying to come up with something for work to compile a C#, C++, and multiple Java projects to a common directory.  I figured, I may as well try Go over a hacky-Bash script.  


# Changelog
v 0.0.1 -- Initial commits -- Will not run or do much of anything. Working on figuring out Go.


# goenv
`
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/adam/golang"
GORACE=""
GOROOT="/usr/local/opt/go/libexec"
GOTOOLDIR="/usr/local/opt/go/libexec/pkg/tool/darwin_amd64"
GO15VENDOREXPERIMENT="1"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fno-common"
CXX="clang++"
CGO_ENABLED="1"
`