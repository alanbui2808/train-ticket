# make is a build automation tool, part of MacOS
#	it reads file named "Makefile" and runs commands defined in it (targets)

APP_NAME = server

#	"Hey make, when someone says "run", execute:
run:
		go run ./cmd/${APP_NAME}/main.go

