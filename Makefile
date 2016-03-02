ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOPATH := ${GOPATH}:${ROOT}
export GOPATH := ${GOPATH}



test: test-converter test-grid test-plast test-ext-stock test-ext-lines test-window test-svg

test-svg:
	@go test ./

test-grid:
	@cd ./grid && go test ./

test-plast:
	@cd ./plast && go test ./

test-converter:
	@cd ./converter && go test ./

test-window:
	@cd ./window && go test ./


# Extentions
test-ext-stock:
	@cd ./ext/stock && go test ./

test-ext-lines:
	@cd ./ext/lines && go test ./


# reinstall
reinstall:
	@git push
	@go get -u github.com/iostrovok/svg-charts/points
	@go get -u github.com/iostrovok/svg-charts/colors
	@go get -u github.com/iostrovok/svg-charts/grid
	@go get -u github.com/iostrovok/svg-charts/converter
	@go get -u github.com/iostrovok/svg-charts/plast
	@go get -u github.com/iostrovok/svg-charts/window
	@go get -u github.com/iostrovok/svg-charts/ext/stock/
	@go get -u github.com/iostrovok/svg-charts/ext/lines/
	@go get -u github.com/iostrovok/svg-charts

# Examples
example-one:
	@go run -v ./example/expamle_1.go
