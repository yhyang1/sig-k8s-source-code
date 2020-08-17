#!/bin/bash

kubeadm reset

iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X

ipvsadm -C

yum remove -y kubelet kubeadm kubectl