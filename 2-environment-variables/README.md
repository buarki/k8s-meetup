# 2.Environment variables

## TOC
- [1. How to define environment variables](#1-how-to-define-environment-variables)
- [2. Deploy the app](#2-deploy-the-app)

## 1. How to define environment variables <div id='1-how-to-define-environment-variables'>
To define env vars you should add `env` attribute at containers definition, as you can see at [deployment.yml](/deployment.yml).

## 2. Deploy the app <div id='deploy-the-app'>

As you can see at [ingress.yml](/ingress.yml) we have the host `app-2-of-tutorial.com` defined, so `add it to your /etc/hosts´:

```
echo "$(minikube ip) app-2-of-tutorial.com" >> /etc/hosts
```

Once you configure hosts, deploy it:

```
kubectl apply -f .
```

Enjoy my surprise for you :)


