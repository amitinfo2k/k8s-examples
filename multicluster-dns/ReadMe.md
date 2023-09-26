
# Example for setting up Custom DNS server forwarder

 Example for setting up Custom DNS server to forward the dns query from particualr pod to another pod
 running in different cluster. 
 Assumption : 
 1] network reachability between both the cluster 
 2] pod and service cidr are different for both the clusters

# Steps for setup

1] Setup two k8s cluster say cluster1 and cluster2

2] Setup dns server in cluster1. before deploying the custom dns server you will need

DNS server service ip from the both the clusters also the dns suffix based on that the dns queries will
be forwarded

```
data:
  Corefile: |-
    <dns suffix to forward queries>:5353 {
      forward . <target dns service ip for cluster2>
    }
    .:5353 {
      forward .  <dns service ip for cluster1>
    }


e.g. 

data:
  Corefile: |-
    omec.svc.cluster.local:5353 {
      forward . 192.168.85.10
    }
    .:5353 {
      forward . 192.168.87.10
    }
```

on cluster1

3] install custom dns server

kubectl create -f my-dns-setup.yaml -n <namespace>


4] install test pod for testing dns queries

kubectl create -f test-pod.yaml -n <namespace>






