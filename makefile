
setup:
	@echo "Setting up the project dependencies ..."
	@make tidy 
	@make deps-upgrade
	@echo "Project dependencies are set up."

# ~~~ Modules support ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
tidy:
	go mod tidy

deps-reset:
	git checkout -- go.mod
	go mod tidy

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

.PHONY: setup