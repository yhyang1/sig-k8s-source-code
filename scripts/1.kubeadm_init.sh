#!/bin/bash

images=(

k8s.gcr.io/kube-apiserver:v1.18.6
k8s.gcr.io/kube-controller-manager:v1.18.6
k8s.gcr.io/kube-scheduler:v1.18.6
k8s.gcr.io/kube-proxy:v1.18.6
k8s.gcr.io/pause:3.2
k8s.gcr.io/etcd:3.4.3-0
k8s.gcr.io/coredns:1.6.7

    kube-apiserver-amd64:v1.11.2
    kube-controller-manager-amd64:v1.11.2
    kube-scheduler-amd64:v1.11.2
    kube-proxy-amd64:v1.11.2
    pause:3.1
    etcd-amd64:3.2.18
    coredns:1.1.3

    pause-amd64:3.1

    kubernetes-dashboard-amd64:v1.10.0
    heapster-amd64:v1.5.4
    heapster-grafana-amd64:v5.0.4
    heapster-influxdb-amd64:v1.5.2
)

for imageName in ${images[@]} ; do
    docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
    docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
done

kubeadm init \
    --control-plane-endpoint=cluster-endpoint \
    --apiserver-advertise-address=172.25.58.161