# 4. PVC (Persistent Volume Claim)

## TOC
- [1. What is PVC](#1-what-is-pvc)
- [2. How to create PVCs](#2-how-to-create-pvcs)
- [3. How to see created file using application](#3-how-to-see-created-file-using-application)
- [4. How to list PersistentVolumes](#4-how-to-list-persistentvolumes)
- [5. How to describe PersistentVolumes](#5-how-to-describe-persistentvolumes)
- [6. How to list PersistentVolumesClaims](#6-how-to-list-persistentvolumesclaims)
- [7. How to describe PersistentVolumesClaims](#7-how-to-describe-persistentvolumesclaims)
- [8. How to delete those volume objects](#8-how-to-delete-thos-volume-objects)


## 1. What is PVC <div id='1-what-is-pvc'>

PVC stands for Persistent Volume Claim. It is an object that allows you to request volume when needed.
In k8s worlds there are three types of volumes:
- Volume:
    * it is related to volume at Pod level, once Pod dies volume is gone;
- PersistentVolume:
    * it is a persistent volume, which means that if a container uses that volume and Pod crashes, when new Pod gets up the new one will be able to use data defined at that volume;
- PersistentVolumeClaim:
    * it requests volume from a persistent volume;

A useful cenario where PVC is useful is related to databases. If the container running the database crashes data must be kept.

## 2. How to create PVCs <div id='2-how-to-create-pvcs'>

First, it is necessary create a path on minikube cluster. To do it:
- access the minikube shell by running bellow command:

```
minikube ssh
```

- create a path to we place the example file:

```
sudo mkdir /mnt/k8s
```

- write a file to that path to be read by the example application:

```
sudo sh -c "echo 'hi from a volume' > /mnt/k8s/dumbtext"
```

- check if file `/mnt/k8s/dumbtext` has content `hi from a volume`:

```
cat /mnt/k8s/dumbtext
```

Once you did it, it is necessary define a PersistentVolume before. Take a look at the persistent volume at [persistent-volume.yml](persistent-volume.yml) and also at file [pvc.yml](pvc.yml).


In our example, the PersistentVolume `simple-pv` is setting up 5Gi of memory, and PVC `simple-pvc` is requesting 2Gi.

## 3. How to see created file using application <div id='3-how-to-see-created-file-using-application'>

```
kubectl apply -f .
```

Note: do NOT forget to add host `app-5-of-tutorial.com` to /etc/hosts.

Then, run this curl:

```
curl https://app-5-of-tutorial.com -k
```

And you should see the content of file `/mnt/k8s/dumbtext`

## 4. How to list PersistentVolumes <div id='4-how-to-list-persistentvolumes'>

```
kubectl get pv
```

## 5. How to describe PersistentVolumes <div id='5-how-to-describe-persistentvolumes'>

```
kubectl describe pv simple-pv
```

The output will look like this one:

```
Name:            simple-pv
Labels:          type=local
Annotations:     pv.kubernetes.io/bound-by-controller: yes
Finalizers:      [kubernetes.io/pv-protection]
StorageClass:    manual
Status:          Bound
Claim:           default/simple-pvc
Reclaim Policy:  Retain
Access Modes:    RWO
VolumeMode:      Filesystem
Capacity:        5Gi
Node Affinity:   <none>
Message:         
Source:
    Type:          HostPath (bare host directory volume)
    Path:          /mnt/k8s
    HostPathType:  
Events:            <none>

```

## 6. How to list PersistentVolumeClaims <div id='6-how-to-list-persistentvolumesclaims'>

```
kubectl get pvc
```

## 7. How to describe PersistentVolumeClaims <div id='7-how-to-describe-persistentvolumesclaims'>

```
kubectl describe pvc simple-pvc
```

The output will look like this:

```
Name:          simple-pvc
Namespace:     default
StorageClass:  manual
Status:        Bound
Volume:        simple-pv
Labels:        <none>
Annotations:   pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      5Gi
Access Modes:  RWO
VolumeMode:    Filesystem
Used By:       app-5-5fd97bc89b-cp46g
Events:        <none>
```

## 8. How to delete those volume objects <div id='8-how-to-delete-thos-volume-objects'>

```
kubectl delete -f .
```
