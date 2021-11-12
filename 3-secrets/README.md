# 3. Secrets

## TOC
- [1. How to create secrets](#1-how-to-create-secrets)
- [2. How to list secrets](#2-how-to-list-secrets)
- [3. How to describe a secret](#3-how-to-describe-a-secret)
- [4. How to inspect the secret value](#4-how-to-inspect-the-secret-value)
- [5. How to read a secret as an environment variable](#5-how-to-read-a-secret-as-an-environment-variable)

# 1. How to create secrets <div id='1-how-to-create-secrets'>

Not all data can be passed as environment variables because they are sensitive, like a database password. A more appropriate way to pass them is by using secrets.

Creating secrets in a k8s cluster is more or less like the bellow command:

```
kubectl create secret $TYPE $SECRET_NAME --from-literal KEY=VALUE
```

Where TYPE can be:
- generic;
- docker-registry;
- tls;


In our example, lets create a secret called `my-safe-secret` and lets add key-value pair where key is called `useful_video` which will be a link to an youtube video.

```
kubectl create secret generic my-safe-secret \
        --from-literal "useful_video=https://www.youtube.com/watch?v=9XcQNbHUwFE"
```

# 2. How to list secrets <div id='2-how-to-list-secrets'>

```
kubectl get secrets
```

# 3. How to describe a secret <div id='3-how-to-describe-a-secret'>

```
kubectl describe secret my-safe-secret
```

# 4. How to inspect the secret value <div id='4-how-to-inspect-the-secret-value'>

```
kubectl get secrets my-safe-secret -o json
```

# 5. How to read a secret as an environment variable <div id='5-how-to-read-a-secret-as-an-environment-variable'>

As we say in Brazil, "I don't want to fill sausage", so just take a look at [deployment.yml](deployment.yml) and you'll see how to use it.

NOTE: remember to add the host `app-4-of-tutorial.com` to your /etc/hosts and then run:

```
kubectl apply -f .
```

To consume the service:

```
curl https://app-4-of-tutorial.com -k
```
