# practical-k8s

Internal Tikal workshop on kubectl

# Install tooling

`kubectl` https://kubernetes.io/docs/tasks/tools/install-kubectl/
`kubectx` and `kubens` https://github.com/ahmetb/kubectx
`stern` (optional) for tailing logs https://github.com/wercker/stern

# Local kubernetes

## Docker desktop

## Kind OSX

https://kind.sigs.k8s.io/

```bash
brew install kind
./kind-cluster-with-local-registry.sh
kubectl create namespace practical-k8s
kubens practical-k8s
```

The original sh script can be found [here](https://kind.sigs.k8s.io/docs/user/local-registry/)

This will create a Kind k8s node in docker and a local registry running on port 5000.

## K3s Linux

https://k3s.io/

```bash
curl -L -O https://github.com/k3s-io/k3s/releases/download/v1.20.2%2Bk3s1/k3s
sudo k3s server --write-kubeconfig-mode 644
```


# Building the containers

## K3S & Kind

This one will build the images for the local repository.

```bash
make docker-other
```

## Docker Desktop

This one will build for the local docker with an arbitrary tag (to prevent registry pulling)

```bash
make docker
```

# Using the correct image in Pod definitions

Before you apply the yaml files with `kubectl apply`, you need to make sure that you're using the right image.

## K3S & Kind

```yaml
image: localhost:5000/ping:latest

image: localhost:5000/pong:latest
```

## Docker Desktop

```yaml
image: ping:1

image: pong:1
```
