build:
	go build

test: build
	~/bin/maelstrom/maelstrom test -w echo --bin ./maelstrom-echo --node-count 1 --time-limit 10
