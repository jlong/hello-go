CC      = go
CFLAGS  = build
RM      = rm -f


default: all

all: hello

hello: hello.go
	$(CC) $(CFLAGS) hello.go

clean:
	$(RM) hello

test: clean hello
	bats .

tests: test
