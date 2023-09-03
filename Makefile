dev: 
	set -a && source .env && set +a && air

docs-gen:
	swag init --parseDependency -q