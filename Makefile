# Define variables for the resources and commands
KUBECTL = kubectl
GO = go
DEPLOYMENT_DIR = ./manifests
SNAPSHOT_CMD = go run cmd/main.go snapshot

# Define the default target
.PHONY: all
all: deploy snapshot

.PHONY: deploy
deploy:
	@echo "Deploying Kubernetes resources..."
	$(KUBECTL) apply -f $(DEPLOYMENT_DIR)


.PHONY: snapshot
snapshot:
	@echo "Running snapshot build..."
	$(SNAPSHOT_CMD)
