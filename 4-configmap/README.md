# 4. ConfigMaps

## TOC
- [1. What is ConfigMap](#1-what-is-configmap)
- [2. How to create ConfigMaps](#2-how-to-create-configmaps)
- [3. How to use ConfigMaps](#3-how-to-use-configmaps)
- [4. How to list ConfigMaps](#4-how-to-list-configmaps)
- [5. How to describe ConfigMaps](#5-how-to-describe-configmaps)

# 1. What is ConfigMap <div id='1-what-is-configmap'>

ConfigMap are objects that allows to group and share configurable data between k8s objects.

# 2. How to create ConfigMaps <div id='2-how-to-create-configmaps'>

Just take a look at [configmap.yml](configmap.yml) file.

# 3. How to use ConfigMaps <div id='3-how-to-use-configmaps'>

Just take a look at [deployment.yml](deployment.yml) file.

# 4. How to list ConfigMaps <div id='4-how-to-list-configmaps'>

```
kubectl get configmaps
```

# 5. How to describe ConfigMaps <div id='5-how-to-describe-configmaps'>

```
kubectl describe configmaps app-4-config-map-for-local-development 
```
