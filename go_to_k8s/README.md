![image](https://tvax2.sinaimg.cn/large/007F3CC8ly1h5fjq325qhj315i0p645p.jpg)

k8s部署mysql,redis

k8s部署go 
https://blog.csdn.net/qq_35349114/article/details/115488514



如没安装hyper-v，执行以下脚本。

hyper-v.cmd

```bash
pushd "%~dp0"
 
dir /b %SystemRoot%\servicing\Packages\*Hyper-V*.mum >hyper-v.txt
 
for /f %%i in ('findstr /i . hyper-v.txt 2^>nul') do dism /online /norestart /add-package:"%SystemRoot%\servicing\Packages\%%i"
 
del hyper-v.txt
 
Dism /online /enable-feature /featurename:Microsoft-Hyper-V-All /LimitAccess /ALL
```


如果不能直接安装，使用克隆方式试试：https://blog.csdn.net/weixin_41915314/article/details/113619475

