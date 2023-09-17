GO = go
MAIN_FILE = src/vertex/main.go
NAME = vertex
EXAMPLE = 1

.PHONY: run build build_run

run:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running $(NAME)...                     |"
	@echo " ----------------------------------------------------"
	@echo
	
	@$(GO) run examples/example$(EXAMPLE).go

setup:
	@echo
	@echo
	@echo " ----------------------------------------------------"
	@echo "      Installing Packages for Dev version...          |"
	@echo " ----------------------------------------------------"
	@echo
	$(GO) mod tidy

build_run:
	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Building $(NAME)...                    |"
	@echo " ----------------------------------------------------"
	@echo

	@$(GO) build -o bin/$(NAME) $(EXAMPLE)

	@echo
	@echo " ----------------------------------------------------"
	@echo "|              Running $(NAME)...                     |"
	@echo " ----------------------------------------------------"
	@echo

	@./bin/$(NAME)