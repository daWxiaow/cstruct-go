set GOPATH=%~dp0\\..\\..\\..\\..
set GOBIN=%~dp0\\bin
echo %GOPATH%
echo %GOBIN%

cd benchmarks
go test -test.bench=".*" -count=1