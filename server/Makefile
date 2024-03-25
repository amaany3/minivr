#############################
# middleware
#############################
.PHONY: middleware.up
middleware.up:
	@docker-compose up -d

.PHONY: middleware.down
middleware.down:
	@docker-compose down

.PHONY: middleware.reset
middleware.reset:
	@docker-compose down --remove-orphans -v

.PHONY: middleware.psql
middleware.psql:
	PGPASSWORD=pgpass psql -h postgres.minivr.local -p 5432 -U minivr -d minivr
