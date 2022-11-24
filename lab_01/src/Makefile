TESTFLAGS = -cover -v -run Test
BENCHFLAGS = -bench=. -benchmem -run Benchmark
BUILDFLAGS = -o app.exe -gcflags '-N -l'

PKG1 = ./levenshtein/
PKG2 = ./utils/

FILES = main.go 
FILES += $(PKG1)dl_iterative.go $(PKG1)dl_recursive_cash.go $(PKG1)dl_recursive.go $(PKG1)l_iterative.go $(PKG1)type.go
FILES += $(PKG2)utils.go

.PHONY: app build run test bench clean

app: app.exe
	app.exe

build: app.exe

run:
	go run main.go

test:
	@-go test $(PKG1) $(TESTFLAGS)
	@echo 
	@go test $(PKG2) $(TESTFLAGS)

bench:
	@go test $(PKG1) $(BENCHFLAGS)

clean:
	rm -f app.exe


app.exe: $(FILES)
	go build $(BUILDFLAGS)