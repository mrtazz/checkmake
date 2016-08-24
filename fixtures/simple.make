# this is a comment

exanded = "$(simple)"
simple := "foo"

clean:
	rm bar
	rm foo

foo: bar
	touch foo

bar:
	touch bar

all: foo

.PHONY: all clean

.DEFAULT: all
