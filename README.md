# k8swebhook
minimal k8s admission webhook template

caBundle in deploy.yaml is base64 encoded ca.pem

server/hook certificate dn must be hook.default.svc (if you keep service name and default namespace)
