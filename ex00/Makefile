BINARY_NAME	:= main.out
RM				:= rm -f

all:	build

build:
	go build -o ${BINARY_NAME} main.go

test:
	go test -v ./... 
 
run:	build
	./${BINARY_NAME}
 
clean:
	go clean
	${RM} ${BINARY_NAME}

fclean:	clean

re:	fclean	all

format:
	gofmt -s -w .

.PHONY: all clean fclean re test format
