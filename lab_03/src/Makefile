TESTFLAGS = -cover -v -run Test
BENCHFLAGS = -bench=. -benchmem -run Benchmark
BUILDFLAGS = -o ./out/app.exe -gcflags '-N -l'

OUTDIR = ./out/
PKG1 = ./array/

FILES = main.go 
FILES += $(PKG1)type.go $(PKG1)sort_comb.go $(PKG1)sort_insert.go $(PKG1)sort_pancake.go

.PHONY: app build run test bench clean

app: $(OUTDIR)app.exe
	$(OUTDIR)app.exe

build: $(OUTDIR)app.exe

run:
	go run main.go

test:
	@go test $(PKG1) $(TESTFLAGS)

bench:
	@go test $(PKG1) $(BENCHFLAGS)

clean:
	rm -rf $(OUTDIR)


./out/app.exe: $(FILES)
	go build $(BUILDFLAGS)