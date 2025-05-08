# Kubernetes Deployment Instructions

This document provides detailed instructions for deploying the WhoAre application on Kubernetes.

## Prerequisites

- Kubernetes cluster (version 1.28 or higher)
- kubectl configured with cluster access
- Docker images built and available in your container registry
- Basic understanding of Kubernetes concepts

## Deployment Steps

### 1. Configure Environment Variables

Create a ConfigMap and Secret for your environment variables:

```bash
# Create ConfigMap
kubectl create configmap whoare-config \
  --from-literal=ENV=prod \
  --from-literal=DB_DEBUG=false

# Create Secret (replace with your actual values)
kubectl create secret generic whoare-secrets \
  --from-literal=DB_DSN="your-database-connection-string" \
  --from-literal=JWT_SECRET="your-jwt-secret"
```

### 2. Deploy the Application

Apply the Kubernetes manifests in the following order:

```bash
# 1. Deploy ConfigMap and Secrets
kubectl apply -f configmap.yaml
kubectl apply -f secrets.yaml

# 2. Deploy the Database (if using the provided PostgreSQL deployment)
kubectl apply -f postgres.yaml

# 3. Deploy the Application
kubectl apply -f deployment.yaml

# 4. Deploy the Service
kubectl apply -f service.yaml

# 5. Deploy the Ingress (if using)
kubectl apply -f ingress.yaml
```

### 3. Verify Deployment

Check the status of your deployment:

```bash
# Check pods
kubectl get pods

# Check services
kubectl get services

# Check ingress
kubectl get ingress

# Check logs
kubectl logs -f deployment/whoare
```

## Configuration Details

### Resource Requirements

The application is configured with the following resource requirements:

- CPU: 500m (0.5 cores)
- Memory: 512Mi

Adjust these values based on your needs and cluster capacity.

### Health Checks

The application includes:
- Liveness probe: Checks if the application is running
- Readiness probe: Checks if the application is ready to receive traffic

### Scaling

To scale the application:

```bash
# Scale to 3 replicas
kubectl scale deployment whoare --replicas=3
```

## Monitoring and Maintenance

### Logs

View application logs:
```bash
# View logs for all pods
kubectl logs -l app=whoare

# View logs for a specific pod
kubectl logs <pod-name>
```

### Troubleshooting

Common issues and solutions:

1. **Pod not starting**
   ```bash
   kubectl describe pod <pod-name>
   kubectl logs <pod-name>
   ```

2. **Service not accessible**
   ```bash
   kubectl describe service whoare
   kubectl get endpoints whoare
   ```

3. **Database connection issues**
   ```bash
   kubectl logs -l app=whoare | grep "database"
   ```

## Security Considerations

1. **Secrets Management**
   - Never commit secrets to version control
   - Use Kubernetes Secrets or a secrets management solution
   - Rotate secrets regularly

2. **Network Policies**
   - Implement network policies to restrict pod-to-pod communication
   - Use TLS for all external communications

3. **RBAC**
   - Implement proper Role-Based Access Control
   - Use service accounts with minimal required permissions

## Backup and Recovery

### Database Backup

```bash
# Backup PostgreSQL database
kubectl exec -it <postgres-pod> -- pg_dump -U postgres whoare > backup.sql
```

### Application Backup

```bash
# Backup ConfigMaps and Secrets
kubectl get configmap whoare-config -o yaml > configmap-backup.yaml
kubectl get secret whoare-secrets -o yaml > secrets-backup.yaml
```

## Cleanup

To remove the deployment:

```bash
# Delete all resources
kubectl delete -f .

# Or delete specific resources
kubectl delete deployment whoare
kubectl delete service whoare
kubectl delete ingress whoare
kubectl delete configmap whoare-config
kubectl delete secret whoare-secrets
```

## Best Practices

1. **Resource Management**
   - Set appropriate resource requests and limits
   - Monitor resource usage
   - Implement horizontal pod autoscaling

2. **High Availability**
   - Deploy across multiple nodes
   - Use pod anti-affinity rules
   - Implement proper health checks

3. **Security**
   - Keep Kubernetes and all components updated
   - Implement network policies
   - Use RBAC
   - Regular security audits

4. **Monitoring**
   - Set up proper monitoring and alerting
   - Monitor application metrics
   - Monitor cluster health

## Additional Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Kubernetes Best Practices](https://kubernetes.io/docs/concepts/configuration/overview/)
- [Kubernetes Security](https://kubernetes.io/docs/concepts/security/)
