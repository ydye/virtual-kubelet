#!/bin/bash

set -e

nodename_prefix=`hostname`

/usr/bin/virtual-kubelet --provider mock --provider-config /etc/virtual-kubelet/vk-config.yml --enable-node-lease=true --node-number 150 --nodename $nodename_prefix --kubeconfig /etc/virtual-kubelet/kubeconfig

