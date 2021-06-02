generate:
	docker build -t twirp-app .
	docker run -it --entrypoint protoc --rm -v $(CURDIR):/app twirp-app --go_out=. --twirp_out=. rpc/service.proto
