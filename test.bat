set GOPATH=%~dp0\\..\\..\\..\\..
set GOBIN=%~dp0\\bin
echo %GOPATH%
echo %GOBIN%

go test -v tests/x_test.go tests/nil_test.go tests/array_test.go 