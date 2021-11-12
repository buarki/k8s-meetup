# 1. Getting started

## Goal of this part

The goal here is to *JUST GET STARTED WITH K8S*. It is always good having something done with what you're learning.

To get started with k8s lets just spin up a small service inside of a local cluster using [minikube](https://v1-18.docs.kubernetes.io/docs/tasks/tools/install-minikube/).

## TOC
- [1. Deployment](#1-deployment)
- [2. Service](#2-service)
- [3. Ingress](#3-ingress)
- [4. Start local cluster](#4-start-local-cluster)
- [5. Deploying a small service](#5-deploying-a-small-service)
    * [5.1 Adjusting host](#51-adjusting-host)
    * [5.2 Asking for something that is true :)](#52-asking-for-something-that-is-true)
- [6. How to see data about Pods?](#6-how-to-see-data-about-pods)
    * [6.1 List Pods running](#61-list-pods-running)
    * [6.2 Get datails about Pods](#62-get-details-about-pods)
- [7. How to see data about Deployments?](#7-how-to-see-data-about-deployments)
    * [7.1 List Deployments running](#71-list-deployments-running)
    * [7.2 Get datails about Deployments](#72-get-details-about-deployments)
    * [7.3 How can I check if Deployment really maintains the configs I set?](#73-how-can-i-check-if-deployment-really-maintains-the-configs-i-set)
- [8. How to see data about Service?](#8-how-to-see-data-about-service)
    * [8.1 List Services running](#81-list-services-running)
    * [8.2 Get datails about Services](#82-get-details-about-services)
- [9. How to see data about Ingress?](#9-how-to-see-data-about-ingress)
    * [9.1 List Ingresses running](#91-list-ingresses-running)
    * [9.2 Get datails about Ingress](#92-get-details-about-ingress)

## 1. Deployment <div id='1-deployment'>
- it creates a blueprint of how we want to have our containers, and ensures their availability as we can configure number of replicas, ports..... 
- it takes care of desired state we've configures;
- it also useful because it gives one IP to interact with Services;

## 2. Service <div id='2-service'>
- when a Pod is deployed k8s provides to it an IP (you can check them by running `kubectl get pods -o wide`);
- if a Pod crashes and another application was calling it directy using Pod's IP that application will no longer work;
- if an application calls a Pod directly by its IP and the Pod crashes the application will no longer be able to use that IP;
- so, to not be couple to Pod's IP we should use Services, because it provides a reliable IP and it is able to send traffic to Pods in which labels match;
- recommended type to be used is `ClusterIP`, because it manages traffic inside of cluster, and do NOT allow something outside of cluster to call Pods;

## 3. Ingress <div id='3-ingress'>
- this object sends traffic from outside to Services inside cluster based on rules (hosts, paths...);
- the job we need to do is provide a configuration about how ingress should work and the Ingress controller will take care of keeping our setup;
- I recomend you to use [`ingress-nginx`](https://kubernetes.github.io/ingress-nginx/deploy/#provider-specific-steps);


## 4. Start local cluster <div id='4-start-local-cluster'>
Once you have minikube stalled with propper driver for your system you can start minikube with following command (on my setup I'm using virtualbox as driver):

```
minikube start --kubernetes-version=v1.21.5 --driver=virtualbox
``` 

To check if cluster is working fine you can ask minikube what is its current IP:

```
minikube ip
```

If everything is fine you'll see an IP.

We also need to enable ingress on minikube, it can be done with following command:

```
minikube addons enable ingress
```

## 5. Deploying a small service <div id='5-deploying-a-small-service'>

To spin up a small service, just run on this repository:

```
kubectl apply -f .
```

You should see the following log:

```
service/app-1-service created
deployment.apps/app-1 created
ingress.networking.k8s.io/app-1-ingress created
```

### 5.1 Adjusting host <div id='51-adjusting-host'>
On [the ingress definition](ingress.yml) is configured a host called `tellmesomethingtruthy.com`. To make you able to call this one using `curl`, postman or your browser we need to configure that host under `/etc/hosts`. To do it, just append to that file the current minikube ip followed by that host:

```
echo "$(minikube ip) tellmesomethingtruthy.com" >> /etc/hosts
```


### 5.2 Asking for something that is true :) <div id='52-asking-for-something-that-is-true'>

If everything is fine we can consume the service we deployed:

```
curl https://tellmesomethingtruthy.com -k 
```

## 6. How to see data about Pods? <div id='6-how-to-see-data-about-pods'>

### 6.1 List Pods running <div id='61-list-pods-running'>

```
kubectl get pods
```

### 6.2 Get datails about Pods<div id='62-get-details-about-pods'>

```
kubectl describe pods $POD_NAME
```

The output will be something like this:

```
Name:         app-1-69f58b6fd7-jnzwp
Namespace:    default
Priority:     0
Node:         minikube/192.168.99.106
Start Time:   Thu, 11 Nov 2021 21:07:05 -0300
Labels:       instance=app-1
              pod-template-hash=69f58b6fd7
Annotations:  <none>
Status:       Running
IP:           172.17.0.4
IPs:
  IP:           172.17.0.4
Controlled By:  ReplicaSet/app-1-69f58b6fd7
Containers:
  app-1:
    Container ID:   docker://54d3380c6d77ba02fde4a307f3aa9b715022a8559fe4ba6024976539364ca763
    Image:          oreasek/app-1
    Image ID:       docker-pullable://oreasek/app-1@sha256:49f4135169506bc9d8e1c9cfaa967ac18d24f12536615fc053a437199788090a
    Port:           8080/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Thu, 11 Nov 2021 21:07:11 -0300
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-k8kk8 (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             True 
  ContainersReady   True 
  PodScheduled      True 
Volumes:
  kube-api-access-k8kk8:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age    From               Message
  ----    ------     ----   ----               -------
  Normal  Scheduled  9m49s  default-scheduler  Successfully assigned default/app-1-69f58b6fd7-jnzwp to minikube
  Normal  Pulling    9m48s  kubelet            Pulling image "oreasek/app-1"
  Normal  Pulled     9m43s  kubelet            Successfully pulled image "oreasek/app-1" in 4.596011163s
  Normal  Created    9m43s  kubelet            Created container app-1
  Normal  Started    9m43s  kubelet            Started container app-1
```

## 7. How to see data about Deployments? <div id='7-how-to-see-data-about-deployments'>

### 7.1 List Deployments running <div id='71-list-deployments-running'>

```
kubectl get deployments
```

### 7.2 Get datails about Deployments<div id='72-get-details-about-deployments'>

```
kubectl describe deployment app-1
```

The output will be something like this:

```
Name:                   app-1
Namespace:              default
CreationTimestamp:      Thu, 11 Nov 2021 21:07:05 -0300
Labels:                 <none>
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               instance=app-1
Replicas:               2 desired | 2 updated | 2 total | 2 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  instance=app-1
  Containers:
   app-1:
    Image:        oreasek/app-1
    Port:         8080/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   app-1-69f58b6fd7 (2/2 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  15m   deployment-controller  Scaled up replica set app-1-69f58b6fd7 to 2
```

### 7.3 How can I check if Deployment really maintains the configs I set? <div id='73-how-can-i-check-if-deployment-really-maintains-the-configs-i-set'>

I suggest this for you:
- open two terminal on your machine and let then in such a way you can see their output simultaneously;
- in one of them run this command to be watching Pods running on cluster:

```
watch kubectl get pods
```

- on the other one kill one of the Pods using following command:

```
kubectl delete pod $POD_NAME
```

- on the watching terminal you'll see that one Pod is being terminated while another one is being deployed again :)


## 8. How to see data about Service?

## 8.1 List Services running <div id='81-list-services-running'>

```
kubectl get svc
```

## 8.2 Get datails about Services <div id='82-get-details-about-services'>

```
 kubectl describe svc app-1-service
```

The output will be something like this:

```
Name:              app-1-service
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          instance=app-1
Type:              ClusterIP
IP Family Policy:  SingleStack
IP Families:       IPv4
IP:                10.105.44.191
IPs:               10.105.44.191
Port:              <unset>  8080/TCP
TargetPort:        8080/TCP
Endpoints:         172.17.0.3:8080,172.17.0.4:8080
Session Affinity:  None
Events:            <none>
```

## 9. How to see data about Ingress? <div id='9-how-to-see-data-about-ingress'>

## 9.1 List Ingresses running <div id='91-list-ingresses-running'>

```
kubectl get ingresses
```

## 9.2 Get datails about Ingresse <div id='92-get-details-about-ingress'>

```
kubectl describe ingress app-1-ingress
```

The output will be something like this:

```
Name:             app-1-ingress
Namespace:        default
Address:          192.168.99.106
Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
Rules:
  Host                       Path  Backends
  ----                       ----  --------
  tellmesomethingtruthy.com  
                             /   app-1-service:8080 (172.17.0.3:8080,172.17.0.4:8080)
Annotations:                 kubernetes.io/ingress.class: nginx
                             nginx.ingress.kubernetes.io/rewrite-target: /
Events:
  Type    Reason  Age                From                      Message
  ----    ------  ----               ----                      -------
  Normal  Sync    33m (x2 over 34m)  nginx-ingress-controller  Scheduled for sync
```


