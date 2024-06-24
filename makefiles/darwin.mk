# Makefile extensions for darwin.

# -----------------------------------------------------------------------------
# Variables
# -----------------------------------------------------------------------------


# -----------------------------------------------------------------------------
# OS specific targets
# -----------------------------------------------------------------------------

.PHONY: clean-osarch-specific
clean-osarch-specific:
	@rm -f  $(GOPATH)/bin/$(PROGRAM_NAME) || true
	@rm -f  $(MAKEFILE_DIRECTORY)/coverage.html || true
	@rm -f  $(MAKEFILE_DIRECTORY)/coverage.out || true
	@rm -fr /tmp/sqlite || true
	@rm -fr $(TARGET_DIRECTORY) || true


.PHONY: coverage-osarch-specific
coverage-osarch-specific:
	@go test -v -exec macos_exec_dyld.sh -coverprofile=coverage.out -p 1 ./...
	@go tool cover -html="coverage.out" -o coverage.html
	@open file://$(MAKEFILE_DIRECTORY)/coverage.html


.PHONY: hello-world-osarch-specific
hello-world-osarch-specific:
	@echo "Hello World, from darwin."


.PHONY: run-osarch-specific
run-osarch-specific:
	@go run main.go


.PHONY: setup-osarch-specific
setup-osarch-specific:
	@echo "No setup required."


.PHONY: test-osarch-specific
test-osarch-specific:
	@go test -v -p 1 ./...

# -----------------------------------------------------------------------------
# Makefile targets supported only by this platform.
# -----------------------------------------------------------------------------

.PHONY: only-darwin
only-darwin:
	@echo "Only darwin has this Makefile target."
