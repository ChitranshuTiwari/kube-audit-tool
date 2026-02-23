# Kube Audit Tool

> A hands-on DevOps portfolio project demonstrating Kubernetes security auditing with Go.

Go CLI to audit Kubernetes clusters for security best practices: resource limits, network policies, RBAC.

## Build

```bash
go build -o kube-audit .
```

## Usage

```bash
./kube-audit
# Uses default kubeconfig (~/.kube/config)
```

## License

MIT
