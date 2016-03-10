ROOT := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOPATH := ${GOPATH}:${ROOT}
export GOPATH := ${GOPATH}



test: test-converter test-grid test-plast test-ext-stock test-ext-lines test-ext-hist test-window test-svg

test-svg:
	@echo "test"
	@go test ./

test-grid:
	@echo "test-grid"
	@cd ./grid && go test ./

test-plast:
	@echo "test-plast"
	@cd ./plast && go test ./

test-converter:
	@echo "test-converter"
	@cd ./converter && go test ./

test-window:
	@echo "test-window"
	@cd ./window && go test ./


# Extentions
test-ext-stock:
	@echo "test-ext-stock"
	@cd ./ext/stock && go test ./

test-ext-lines:
	@echo "test-ext-lines"
	@cd ./ext/lines && go test ./

test-ext-hist:
	@echo "test-ext-hist"
	@cd ./ext/hist && go test ./


# reinstall
reinstall:
	@echo "push"
	@git push
	@echo "go get"
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
