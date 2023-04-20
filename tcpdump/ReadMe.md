# Example to run the tcpdump side container and serve the pcap via ngix

## build nginx image

```
cd nginx
docker build -t amitinfo2k/nginx:1.0.0
```
## build tcpd image

```
cd nginx
docker build -t amitinfo2k/tcpdump-go:1.0.0
```

## run sample k8s pod service

kubectl create -f k8s-pod.yaml

## run sample k8s pod with persistent volume service

kubectl create -f k8s-pod-pv.yaml

## run sample k8s pod with persistent volume service in distributed manner separate ngix and app pod with sidecare tcpdump

kubectl create -f k8s-pod-dst.yaml

## API to start PCAP


POST http://<node-ip>:30808/tcpdump

Data:

```
{
   "action" : "start",
   "pcap_file": "app-name" 

} 
```
Note : pcap_file -- is ignored as it will automatically append the pod name.

## API to stop PCAP


POST http://<node-ip>:30808/tcpdump

Data:

```
{
   "action" : "stop",
   "pcap_file": "app-name" 

} 
```

