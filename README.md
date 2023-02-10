#BUILD INSTRUCTIONS (copied from microservices-demo development guide)

Prerequisites
Docker for Desktop.
kubectl (can be installed via gcloud components install kubectl)
skaffold 2.0+ (latest version recommended), a tool that builds and deploys Docker images in bulk.
A Google Cloud Project with Google Container Registry enabled.
Enable GCP APIs for Cloud Monitoring, Tracing, Profiler:
gcloud services enable monitoring.googleapis.com \
    cloudtrace.googleapis.com \
    cloudprofiler.googleapis.com
Minikube (optional - see Local Cluster)
Kind (optional - see Local Cluster)
Option 1: Google Kubernetes Engine (GKE)
💡 Recommended if you're using Google Cloud Platform and want to try it on a realistic cluster. Note: If your cluster has Workload Identity enabled, see these instructions

Create a Google Kubernetes Engine cluster and make sure kubectl is pointing to the cluster.

gcloud services enable container.googleapis.com
gcloud container clusters create demo --enable-autoupgrade \
    --enable-autoscaling --min-nodes=3 --max-nodes=10 --num-nodes=5 --zone=us-central1-a
kubectl get nodes
Enable Google Container Registry (GCR) on your GCP project and configure the docker CLI to authenticate to GCR:

gcloud services enable containerregistry.googleapis.com
gcloud auth configure-docker -q
In the root of this repository, run skaffold run --default-repo=gcr.io/[PROJECT_ID], where [PROJECT_ID] is your GCP project ID.

This command:

builds the container images
pushes them to GCR
applies the ./kubernetes-manifests deploying the application to Kubernetes.
Troubleshooting: If you get "No space left on device" error on Google Cloud Shell, you can build the images on Google Cloud Build: Enable the Cloud Build API, then run skaffold run -p gcb --default-repo=gcr.io/[PROJECT_ID] instead.

Find the IP address of your application, then visit the application on your browser to confirm installation.

kubectl get service frontend-external
Option 2 - Local Cluster
Launch a local Kubernetes cluster with one of the following tools:

To launch Minikube (tested with Ubuntu Linux). Please, ensure that the local Kubernetes cluster has at least:

4 CPUs
4.0 GiB memory
32 GB disk space
minikube start --cpus=4 --memory 4096 --disk-size 32g
To launch Docker for Desktop (tested with Mac/Windows). Go to Preferences:

choose “Enable Kubernetes”,
set CPUs to at least 3, and Memory to at least 6.0 GiB
on the "Disk" tab, set at least 32 GB disk space
To launch a Kind cluster:

kind create cluster
Run kubectl get nodes to verify you're connected to the respective control plane.

Run skaffold run (first time will be slow, it can take ~20 minutes). This will build and deploy the application. If you need to rebuild the images automatically as you refactor the code, run skaffold dev command.

Run kubectl get pods to verify the Pods are ready and running.

Run kubectl port-forward deployment/frontend 8080:8080 to forward a port to the frontend service.

Navigate to localhost:8080 to access the web frontend.

Cleanup
If you've deployed the application with skaffold run command, you can run skaffold delete to clean up the deployed resources.
