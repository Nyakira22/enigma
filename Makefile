.PHONY: setup up d b ps node

setup:
	@make up
	@make ps
d:
	docker compose down
up:
	docker compose up -d --build
ps:
	docker ps
exec:
	docker exec -it backend_enigma bash