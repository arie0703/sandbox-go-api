# Variables
IMAGE_NAME = microservice
TAG = latest
ECR_REGISTRY = $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com
ECR_REPOSITORY = $(ECR_REGISTRY)/$(IMAGE_NAME)

# Default values
AWS_REGION ?= ap-northeast-1

# Build Docker image
build:
	docker build -t $(IMAGE_NAME):$(TAG) .

# Tag image for ECR
tag:
	docker tag $(IMAGE_NAME):$(TAG) $(ECR_REPOSITORY):$(TAG)

# Login to ECR
ecr-login:
	aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin $(ECR_REGISTRY)

# Push image to ECR
push:
	docker push $(ECR_REPOSITORY):$(TAG)

# Build and push (complete workflow)
build-push: build ecr-login tag push

# Clean local images
clean:
	docker rmi $(IMAGE_NAME):$(TAG) || true
	docker rmi $(ECR_REPOSITORY):$(TAG) || true

# Show help
help:
	@echo "Available commands:"
	@echo "  build       - Build Docker image locally"
	@echo "  tag         - Tag image for ECR"
	@echo "  ecr-login   - Login to Amazon ECR"
	@echo "  push        - Push image to ECR"
	@echo "  build-push  - Build, tag, and push image to ECR"
	@echo "  clean       - Remove local Docker images"
	@echo "  help        - Show this help message"
	@echo ""
	@echo "Required environment variables:"
	@echo "  AWS_ACCOUNT_ID - Your AWS Account ID"
	@echo "  AWS_REGION     - AWS Region (default: ap-northeast-1)"

.PHONY: build tag ecr-login push build-push clean help