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

#############################
# migration
#############################
.PHONY: migrate.diff.local
migrate.diff.local:
	@atlas schema diff \
		--from "postgres://minivr:pgpass@postgres.minivr.local:5432/minivr?sslmode=disable" \
		--to "file://schema.sql" \
		--dev-url docker://postgres/15/dev

.PHONY: migrate.apply.local
migrate.apply.local:
	@atlas schema apply \
		--url "postgres://minivr:pgpass@postgres.minivr.local:5432/minivr?sslmode=disable" \
		--to "file://schema.sql" \
		--dev-url docker://postgres/15/dev
