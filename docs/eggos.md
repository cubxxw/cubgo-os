# eggos

[toc]

## 环境

+ Windows11（环境无所谓）
+ docker（版本无所谓）

## 准备

💡 你可以选择直接`push`远程打包好的`eggos-ubuntu`镜像，这样的话，可以不用下面的步骤

```
docker pull 3293172751/cubos-ubuntu:1.0.2
```



**😍 强推的教程：**

+ [x] [docker教程](https://docker.nsddd.top)
+ [x] [Go语言教程](https://go.nsddd.top)
+ [x] [linux教程](https://github.com/3293172751/awesome-cs-course/blob/master/linux/README.md)

> 二种安装方式，一种是直接`docker pull`拉取镜像，第二种就是使用`dockerfile`，建议第二种。

### 第一种直接docker安装

使用`docker`安装Ubuntu（可以不用设置容器映射，直接push即可）

```bash
docker run -it exec --name ubuntu-eggos ubuntu /bin/bash
```

……. 安装`vim`、`go`、`git`等等等等基本工具，哎呀反正挺麻烦的，如果你不是选择直接`pull`我的镜像到本地，建议直接用`dockerfile`方式

```bash
# 虚拟机可以直接用sudo root 进入root用户，Linux教程中有设置root密码教程
apt-get install vim
apt-get install git
apt-get install golang
apt-get install net-tools  # 这个还需要安装一个net-tools 因为ipconfig 需要
```



### 第二种dockerfile

使用`dockerfile`直接安装镜像，⚡推荐使用



**📮 需求**

我们所需要的镜像中，不含有`vim`/`ifconfig`/`golang`等命令，所以我们需要重新按照那些命令

```text
docker commit
```

> 对于简单的可能会是一个很好的选择，但是对于随时变化的镜像来说，特别麻烦

可不可以一次性搞定？

> 某种镜像的增强，给我list做个清单，后续我加入任何功能，直接在list单子里面run，相当于多次提交
>
> **一次性添加，一次性成型。那么你就可以选择相信`dockerfile` 🚸**



## 下载eggos

直接使用ISO镜像来启动eggos，可以免去我们编译整个项目的麻烦，特别是我们刚开始的时候想快速体验eggos的功能的时候。

eggos在每个版本会生成一个ISO镜像文件，我们可以从github的release界面直接下载，从网址 https://github.com/icexin/eggos/releases 进入到eggos的release界面，点击`eggos.iso`下载。

![image-20221013162445155](http://sm.nsddd.top/smsmimage-20221013162445155.png?xxw@nsddd.top)



回到`Ubuntu`，我们需要用到`wget`远程下载，`tar`压缩工具下载。

```bash
apt-get install wget
apt install tar
```

 🧷 远程下载

> + [wget命令学习可以看这篇](https://github.com/3293172751/awesome-cs-course/blob/master/linux/linux-web/8.md)

```bash
wget -c -b  https://github.com/icexin/eggos/releases/download/v0.4.1/eggos_0.4.1_Linux_x86_64.tar.gz
```



## 安装qemu虚拟机

ubuntu的用户可以通过命令 `sudo apt install qemu-system-x86` 来安装qemu虚拟机。

```bash
apt install qemu-system-x86
```



## 启动qemu虚拟机

使用如下命令来启动虚拟机:

```sh
qemu-system-x86_64 -m 256M -nographic -no-reboot -serial mon:stdio -netdev user,id=eth0,hostfwd=tcp::8080-:80 -device e1000,netdev=eth0 -cdrom eggos.iso
```

> 没有图形界面终端，如windows的`wsl`需要加上`-nographic`，从而以非图形化方式启动qem

![image-20221013164722018](http://sm.nsddd.top/smsmimage-20221013164722018.png?xxw@nsddd.top)

 🔥上面的图片表示我们已经安装成功了



## 体验eggos

🚸 安装好了`eggos`后，我们就可以来体验下`eggos`了



### JavaScript解释器

> 我们可以像平常启动`node`、`ipython`、`jshell`那样来使用`js`

```
js
```

![image-20221013165350046](https://sm.nsddd.top/smsmimage-20221013165350046.png?xxw@nsddd.top)

⬇️ 接下来我们可以使用`js`一样来使用这个解释器。

```js
root@eggos# js
>>> console.log("this is cub-os github:https://github.com/3293172751")
this is cub-os github:https://github.com/3293172751
undefined

>>> xiongxinwei = new RegExp("xiongxinwei@mail.com")
/xiongxinwei@mail.com/
    
>>> var a = 12, b = 13
undefined

>>> console.log(a,b)
12 13
undefined
```

这个Javascript解释器内置了一个简单的`http.Get`方法，用于从一个url里面获取内容。

```Js
>>> var url = "http://baidu.com"
>>> resp = http.Get(url)
>>> console.log(resp)
<html>
<meta http-equiv="refresh" content="0;url=http://www.baidu.com/">
</html>
```

📜 对上面的解释：

它甚至有自动补全功能，输入`http.`之后敲击`<TAB>`键，就可以补全出`http.Get`。重复按`<TAB>`可以在多个候选结果间切换。

按`<Crtl+D>`来退出Javascript解释器。



## 退出、重新进入

我记录这个板块的原因是文档上并没有给出退出的方法，我尝试了下用`Ctrl + Z`退出，是成功的，但是后台并没有结束，就和`shell`下面的`Ctrl+C`和`Ctrl+D`的区别⚡ 

⬇️ 下面我分别开启了`8080`端口和`8081`端口，并且企图杀死`8080`进程并且重新进去

![image-20221013173405464](http://sm.nsddd.top/smimage-20221013173405464.png?xxw@nsddd.top)

再一次进入

> ⚠️ 注意，如果你并没有添加`qemu-system-x86_64`到环境变量，那么我想你必须要去`eggos`目录下执行。

```bash
qemu-system-x86_64 -m 256M -nographic -no-reboot -serial mon:stdio -netdev user,id=eth0,hostfwd=tcp::8080-:80 -device e1000,netdev=eth0 -cdrom eggos.iso
```



### 文件系统

`eggos`现在没有实现完整的文件系统，而是使用了 [afero](https://github.com/spf13/afero)作为文件系统的抽象接口。

> `afero`是一个Go 的文件系统抽象系统。
>
> + [afero翻译和笔记](./afero.md)
> + [samba协议介绍](./samba.md)

⬇️ 我们通过`mount`命令挂载一个`samba`文件系统来体验`eggos`的文件系统功能。

查看自己的`IP`地址

![image-20221013181120314](http://sm.nsddd.top/smimage-20221013181120314.png?xxw@nsddd.top)

💡简单的一个案例如下：

```bash
# 找到自己的ip
mount smb://icexin:eggos@172.17.0.3:445/sambashare /share
```



```sh
root@eggos# cd /share
root@eggos# ls
-rw-rw-rw- 111 fib.js
root@eggos# cat fib.js
function fib(n) {
        if (n == 1 || n == 2) {
                return 1;
        }
        return fib(n-1) + fib(n-2);
}

console.log(fib(10))

root@eggos# js fib.js
55
```

📜 对上面的解释：

首先我们通过mount命令将一个samba地址挂载到了本地的一个目录

+ `mount`的第一个参数是源地址URI，第二个参数是本地的挂载点，这里是`/share`
+ 源地址URI符合标准的URI格式，其中`smb`代表我们使用的是`samba`协议
+ `icexin`和`eggos`分别是samba的账号和密码，读者需要替换成自己的账号和密码
+ `172.28.90.3:445`为samba服务器的ip和端口，读者需要替换成自己的ip和端口

接着我们使用`cd`和`ls`命令切换到`/share`目录，并列出目录下的文件，目录下有一个`fib.js`。

然后我们使用`cat`命令打印出`fib.js`文件的内容，里面使用Javascript实现了斐波那契数列算法，并打印出第10个斐波那契数。

最后我们执行`js`命令，用`fib.js`作为参数，效果是执行了`fib.js`里面的Javascript代码，打印出结果`55`。



