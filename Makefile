BUILD=build

build:
	mkdir -p $(BUILD)

init:
	touch .env
	mkdir -p cmd
	mkdir -p internal
	mkdir -p config
	mkdir -p infrastructure
	mkdir -p log
	mkdir -p service