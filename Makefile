run-server:
	go run server/run.go server/config.go server/server.go --host="127.0.0.1" --port=8085

run-client:
	go run client/run.go client/config.go client/client.go --host="127.0.0.1" --port=8085