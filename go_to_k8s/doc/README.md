[toc]

## 部署kubernetes-dashboard


recommended.yaml


kubectl create -f recommended.yaml
或者执行
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta5/aio/deploy/recommended.yaml
```

演示图：

![image](https://tva3.sinaimg.cn/large/007F3CC8ly1h5grffd0xpj31hc0sn7wh.jpg)



---



## 安装kubernetes-dashboard.yaml

地址：https://github.com/AliyunContainerService/k8s-for-docker-desktop/blob/master/kube-system-default.yaml

![image](https://tvax4.sinaimg.cn/large/007F3CC8ly1h5grjy98jsj31hc0scngb.jpg)



---



## 创建token

```bash
eyJhbGciOiJSUzI1NiIsImtpZCI6IklLa092MlRnWGpLVzBMemhpVWp2QzJhb016YUVqckl3T0QwS2Q2QVhIVTgifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI1YWE3NGEyNy1jMmEwLTRlOGYtOTAzNy04NmUyYjIwNmZiNmIiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06ZGVmYXVsdCJ9.OmULOF5pHG3a8ttdEpvFDJ3LMbhBPu-QEzltZvCXHw6GbiULjuWJO1YZh6qhEHMp8G1JnctW4RE310NOUkaZL1VLnT1IlnNqwv3gi8eQsjCAN1DbzgSBpBN_zSKSm4b-6DbtKavAJg76xrq5W32gQIvMyx4xX3I57QuHvaOe54HaX7twK4FWHbbsa6uv2_kUVuSw75U0wXUMdxoz5BsrVFH32bVLT56BwcKSjZ6WSEkGI2oeYaYuwINI6eJbW0ghYRuoyNJ7OiMYA3OSKm7ZQ1Hcj7Nmhubnd0DCRrxNugOw1N5YcFJDeaG59CcuqGpVwG9suvA4zq-myV_zDb8R_Q
```

![image](https://tva4.sinaimg.cn/large/007F3CC8ly1h5grl6idtgj314f0scave.jpg)

建议保存起来，不保存也可以找回的。

## 访问

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/


## 创建dashboard账户

