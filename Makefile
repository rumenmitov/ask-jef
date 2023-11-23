CC=go build
OBJ=ask
SRC=src/*.go

debug:
	${CC} -o build/debug/${OBJ} ${SRC}

prod:
	${CC} -o build/prod/${OBJ} ${SRC}

install:
	mkdir -p $(HOME)/.config/ask-jef
	touch $(HOME)/.config/ask-jef/ask.env
	echo "MODEL=" > $(HOME)/.config/ask-jef/ask.env
	mkdir -p $(HOME)/.cache/ask-jef
	${CC} -o /usr/local/bin/${OBJ} ${SRC}
