TESTFLAGS = -cover -v -run Test
BENCHFLAGS = -bench=. -benchmem -run Benchmark
BUILDFLAGS = -o app.exe -gcflags '-N -l'

PKG1 = ./matrix/

FILES = main.go 
FILES += $(PKG1)type.go $(PKG1)mul_classic.go $(PKG1)mul_Winograd.go $(PKG1)mul_Winograd_optimized.go

.PHONY: app build run test clean

app: app.exe
	app.exe

build: app.exe

run:
	go run main.go

test:
	@go test $(PKG1) $(TESTFLAGS)

clean:
	rm -f app.exe


app.exe: $(FILES)
	go build $(BUILDFLAGS)