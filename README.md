# Titanic

<div align="center">
  <img src="titanic.png" alt="Alt text">
</div>

## Overview
The **Pipeline Rollback Button** is a disaster recovery tool designed to simplify and accelerate the rollback process during application failures. It integrates with Kubernetes and Helm to provide instant recovery by restoring the application to a previously stable state. This project is ideal for environments requiring quick disaster mitigation, especially in CI/CD pipelines.

---

## Features
- **Instant Rollback**: Revert Kubernetes deployments to their last stable state with a single click or command.
- **Pipeline Recovery**: Rollback CI/CD pipelines to prior artifacts or container image versions.
- **User-Friendly Interface**: Trigger rollbacks via CLI or a simple web-based UI (Streamlit).
- **Local Testing Support**: Fully testable on a local Kubernetes cluster using `kind`.
- **Optional Observability**: Integrate Prometheus and Grafana for real-time monitoring of deployments and rollbacks.

---

## Prerequisites
- **Tools & Libraries**:
  - [Kubernetes](https://kubernetes.io/) (with a local `kind` cluster)
  - [Helm](https://helm.sh/)
  - [Python 3.x](https://www.python.org/)
  - Python libraries:
    - `kubernetes`
    - `subprocess`
    - Optional: `streamlit` (for UI)

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
