CC=go build
OBJ=ask
SRC=src/*.go

debug:
	${CC} -o build/debug/${OBJ} ${SRC}

prod:
	${CC} -o build/prod/${OBJ} ${SRC}

install:
	touch $(HOME)/.config/ask-jef/ask.env
	echo "MODEL=" > $(HOME)/.config/ask-jef/ask.env
	${CC} -o /usr/local/bin/${OBJ} ${SRC}
