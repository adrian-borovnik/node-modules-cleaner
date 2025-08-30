# .PHONY: run build

# run:
# 	@case "$(filter-out $@,$(MAKECMDGOALS))" in \
# 		cli) echo "Running CLI..."; go run cmd/cli/main.go ;; \
# 		web) echo "Running Web..."; go run cmd/web/main.go ;; \
# 		*) echo "Unknown subcommand. Use: make run [cli|web]" ;; \
# 	esac

# build:
# 	@case "$(filter-out $@,$(MAKECMDGOALS))" in \
# 		cli) echo "Building CLI..."; go build -o bin/cli cmd/cli/main.go ;; \
# 		web) echo "Building Web..."; go build -o bin/web cmd/web/main.go ;; \
# 		*) echo "Unknown subcommand. Use: make build [cli|web]" ;; \
# 	esac

# # Prevent make from interpreting subcommands as targets
# %:
# 	@:
#
#

.PHONY: run build exec install

ARGS ?=

run:
	go run cmd/node-modules-cleanner/main.go $(ARGS)

build:
	go build -o bin/nmclean cmd/node-modules-cleanner/main.go

exec: build
	./bin/nmclean $(ARGS)

install: build
	sudo mv bin/nmclean /usr/local/bin/nmclean
