docker-run:
	rm backend/.env || :
	docker compose --env-file release.env up --build
	cp .env backend/
