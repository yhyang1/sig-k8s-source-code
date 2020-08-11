#!/bin/bash

kubeadm init \
    --control-plane-endpoint=cluster-endpoint \
    --apiserver-advertise-address=172.25.58.161

