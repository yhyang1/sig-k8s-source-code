#!/bin/bash

kubeadm --v=5 config images pull --kubernetes-version v1.18.6 --image-repository registry.cn-zhangjiakou.aliyuncs.com/k8sx

kubeadm init \
  --control-plane-endpoint=cluster-endpoint \
  --apiserver-advertise-address=172.25.58.161 \
  --kubernetes-version v1.18.6 \
  --image-repository registry.cn-zhangjiakou.aliyuncs.com/k8sx

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml