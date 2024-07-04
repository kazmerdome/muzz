

up:
	@docker-compose up -d

run-gateway:
	@go run cmd/gateway/main.go

build-gateway:
	@go build -o build/gateway cmd/gateway/main.go

sqlc:
	@go run github.com/sqlc-dev/sqlc/cmd/sqlc generate

# mockery
#
# generate mocks for a given interface |  usage: make name=HealthService mock
# optional filename                    |  usage: make name=HealthService filename=mock_health_service.go mock
# optional srcpkg                      |  usage: make name=HealthService srcpkg=github.com/yoyr/pkg mock
# optional structname                  |  usage: make name=HealthService structname=X mock
mock:
	@echo "Generating mock for $(name)..."
	@cmd="go run github.com/vektra/mockery/v2@latest --name=$(name) --recursive"; \
	if [ -n "$(filename)" ]; then \
		cmd="$$cmd --filename=$(filename)"; \
	fi; \
	if [ -n "$(srcpkg)" ]; then \
		cmd="$$cmd --srcpkg=$(srcpkg)"; \
	fi; \
	if [ -n "$(structname)" ]; then \
		cmd="$$cmd --structname=$(structname)"; \
	fi; \
	eval $$cmd

.PHONY: mock

.PHONY: gateway up sqlc mock

# postgres migration
#
postgres_uri="no pstgres_uri loaded from env"
env_file=.env
version=
ifneq ("$(wildcard $(env_file))","")
	include $(env_file)
endif

# Check if SSL is disabled and set the appropriate SSL mode
ifeq ($(POSTGRES_IS_SSL_DISABLED),true)
	ssl_mode=?sslmode=disable
else
	ssl_mode=
endif

# Construct the postgres URI
postgres_uri:=$(POSTGRES_URI)/$(POSTGRES_DATABASE)$(ssl_mode)

# jump to the next version
migrate-up:
	@migrate -path db/migration -database $(postgres_uri) --verbose up $(version)

# use a specific version | usage: make version=1 migrate-goto
migrate-goto:
	@migrate -path db/migration -database $(postgres_uri) --verbose goto $(version)

# force a specific version | usage: make version=1 migrate-force
migrate-force:
	@migrate -path db/migration -database $(postgres_uri) --verbose force $(version)

# jump back to the previous version
migrate-down:
	@migrate -path db/migration -database $(postgres_uri) --verbose down 1

# print current migration version
migrate-version:
	@migrate -path db/migration -database $(postgres_uri) --verbose version

.PHONY: migrate-up migrate-goto migrate-force migrate-down migrate-version
