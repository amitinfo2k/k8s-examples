# Example to run the tcpdump side container and serve the pcap via ngix

## build nginx image


## build tcpd image


## run sample k8s pod service


## run sample k8s pod with persistent volume service



## API to start PCAP


POST http://<node-ip>:30808/tcpdump

Data:

{
   "action" : "start",
   "pcap_file": "app-name" 

} 

Note : pcap_file -- is ignored as it will automatically append the pod name.

## API to stop PCAP


POST http://<node-ip>:30808/tcpdump

Data:

{
   "action" : "stop",
   "pcap_file": "app-name" 

} 


