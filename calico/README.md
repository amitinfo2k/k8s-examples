

# Example to split Calico network into two where one supports destination natting and other one doesn't

## Pre-requisite

- k8s cluster with defaul Calico network as primary CNI

## Steps of configuration 

1] Update the existing default ip pool (if u want to update existing default ip pool)
```
calicoctl replace -f default-ip-pool.yaml
```
2] Create other IP pool
```
calicoctl create -f my-ip-pool.yaml 
```
3] Check the ip pool config:

```
$ calicoctl get ipPool -o wide
NAME                CIDR              NAT     IPIPMODE   VXLANMODE   DISABLED   SELECTOR
default-pool        192.168.84.0/24   true    Always     Never       false      all()
no-nat-10.0.0.0-8   10.0.0.0/8        false   Never      Never       false      all()
```
4] Create a pod with following annotation:
```
annotations:
    cni.projectcalico.org/ipv4pools: "[\"no-nat-10.0.0.0-8\"]"
```
5] Create a Pod with ip from my ippool 
```
kubectl apply -f example-pod.yaml 
```
6] Verify  Pod Ips
```
$ kubectl get pods -o wide -w
 NAME             READY   STATUS    RESTARTS   AGE    IP            NODE    NOMINATED NODE   READINESS GATES
 pingtest-pool2   1/1     Running   0          164m   10.69.149.0   node1   <none>           <none>>
```
