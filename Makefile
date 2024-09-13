# Variables for coverage and profiling files
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html
CPU_PROFILE := cpu.prof
MEM_PROFILE := mem.prof

# Default target runs tests with race conditions
.PHONY: all
all: test

# Test with race detection, coverage report generation
.PHONY: test
test:
	go test ./... -v -race -covermode=atomic -coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

# Test without race detection, coverage report generation
.PHONY: test-no-race
test-no-race:
	go test ./... -v -covermode=atomic -coverprofile=$(COVERAGE_FILE)
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

# Format the code
.PHONY: format
format:
	go fmt ./...

# Run benchmarks with memory allocation statistics
.PHONY: bench
bench:
	go test ./... -bench . -benchmem -cpu=1

# Generate CPU and memory profiles while running benchmarks
.PHONY: profile-bench
profile-bench:
	go test ./... -bench . -cpuprofile=$(CPU_PROFILE) -memprofile=$(MEM_PROFILE) -cpu=1

# Generate CPU profiling report
.PHONY: cpu-report
cpu-report:
	go tool pprof $(CPU_PROFILE)

# Generate memory profiling report
.PHONY: mem-report
mem-report:
	go tool pprof $(MEM_PROFILE)

# Clean up profiling and coverage files
.PHONY: clean
clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML) $(CPU_PROFILE) $(MEM_PROFILE)
