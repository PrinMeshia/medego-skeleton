BINARY_NAME=MedeGoApp.exe

build:
	@go mod tidy
	@go mod vendor
	@go build -o tmp/${BINARY_NAME} .
	@echo MedeGo built!

run: build
	@echo Starting MedeGo...
	@.\tmp\${BINARY_NAME} &
	@echo MedeGo started!

clean:
	@echo Cleaning...
	@go clean
	@del .\tmp\${BINARY_NAME}
	@echo Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
    
stop:
	@echo Stopping MedeGo...
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped MedeGo

restart: stop start
