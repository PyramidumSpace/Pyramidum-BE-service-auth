SEP="========================================================"
GQLGEN ?= github.com/99designs/gqlgen

################################################################################################################

################################################################################################################
.PHONY: env

define ENV_SAMPLE
JWT_AUTH_SALT=bluh
POSTGRES_DSN=bluh
SERVER_PORT=8080
LOG_LEVEL=


endef
export ENV_SAMPLE
env:
	@if [ ! -f ".env" ];\
		then echo "$$ENV_SAMPLE" > .env;\
	 fi