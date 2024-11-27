# Titanic

<div align="center">
  <img src="titanic.png" alt="Alt text">
</div>

## Overview
The **Pipeline Rollback Button** is a disaster recovery tool designed to simplify and accelerate the rollback process during application failures. It integrates with Kubernetes and Helm to provide instant recovery by restoring the application to a previously stable state. This project is ideal for environments requiring quick disaster mitigation, especially in CI/CD pipelines.

---


## Prerequisites
- **Tools & Libraries**:
  - [Helm](https://helm.sh/)
  - Go

- **Knowledge Requirements**:
  - Basic understanding of Kubernetes, Helm, and Python.
  - Familiarity with CI/CD concepts.

---

## Setup Instructions

### 1. Create a Local Kubernetes Cluster
- Use `kind` to create a local Kubernetes cluster:
  ```bash
  kind create cluster --name rollback-demo
  kubectl cluster-info --context kind-rollback-demo
