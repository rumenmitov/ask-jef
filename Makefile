CC=go build
OBJ=ask
SRC=src/*.go

debug:
	${CC} -o build/debug/${OBJ} ${SRC}

prod:
	${CC} -o build/prod/${OBJ} ${SRC}
