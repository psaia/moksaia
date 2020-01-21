moksaia:
	go build -o moksaia cmd/moksaia/main.go

all: moksaia gen

gen:
	./moksaia

clean:
	rm moksaia && rm docs/*.html

.PHONY: gen clean
