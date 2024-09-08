# Define the shell to use
SHELL := /bin/bash

# Define the scripts
CREATE_SCRIPT := create-openings-stub.sh
DELETE_SCRIPT := delete-openings-stub.sh

# Default target
.PHONY: all
all: help

# Help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make create USERS=<number_of_users>"
	@echo "  make delete START_ID=<start_id> END_ID=<end_id>"
	@echo ""
	@echo "Targets:"
	@echo "  create   - Create job openings"
	@echo "  delete   - Delete job openings by ID range"
	@echo "  help     - Show this help message"

# Create target
.PHONY: create
create:
ifndef USERS
	$(error USERS is not set. Usage: make create USERS=<number_of_users>)
endif
	@chmod +x $(CREATE_SCRIPT)
	@./$(CREATE_SCRIPT) $(USERS)

# Delete target
.PHONY: delete
delete:
ifndef START_ID
	$(error START_ID is not set. Usage: make delete START_ID=<start_id> END_ID=<end_id>)
endif
ifndef END_ID
	$(error END_ID is not set. Usage: make delete START_ID=<start_id> END_ID=<end_id>)
endif
	@chmod +x $(DELETE_SCRIPT)
	@./$(DELETE_SCRIPT) $(START_ID) $(END_ID)