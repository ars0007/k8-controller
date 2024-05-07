# k8-controller
K8-controller creates custom controller and webhook

## Description
Custom resource takes two numbers and performs division operation.
Webhook performs validation that numbers should not be Negative or Zero

### To Deploy on the cluster

```sh
make docker-build IMG=<some-registry>/k8-controller:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/k8-controller:tag
```

**Create Custom resource kind `DivideResource`:**

```sh
kubectl create -f api/v1/example.yaml
```




#

