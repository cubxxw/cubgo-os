# samba协议

[toc]

## 基本使用

> 当前使用的是`docker`的`Ubuntu`镜像

```bash
sudo root
apt-get install samba
touch /etc/samba/smbpasswd
smbpasswd -a xionxinwei  # 设置用户名和密码
```

📮 修改配置文件

```
vim /etc/samba/smb.conf
```

📄 写入以下内容（`G`到最后一行，`o`下一行输入）

```json
[cub-os docker-ubuntu]
	comment = ubuntu
	path = /home/xiongxinwei
	writable = yes
	valid users = xiongxinwei
	available = yes
	create mask = 0777
	directory mask = 0777
	public = yes
```

> ⚠️ 注意，配置文件中最上方中括号里的[cub-os docker-ub]也是挂载路径的一部分

如果是上述配置，则需访问

>  smb://samba.username:samba.password@samba.url/cub-os docker-ubuntu/ /share

⚡ 重启服务

```
/etc/init.d/smbd restart
```

> Linux中套路一样，一般在`etc/init.d`下面都有对应的文件，使用`restart`刷新就好。

💡 验证

```
smb://IP地址
```

> 此时就是可以进去虚拟机里面的目录，这个就和`ftp`类似，我们也是把目录给共享出去l

## Samba的工作原理

Samba服务功能强大，这与其通信基于SMB/CIFS协议有关。SMB不仅提供目录和打印机共享，还支持认证、权限设置。在早期，SMB运行于NBT协议（NetBIOS over TCP/IP）上，使用UDP协议的137、138及TCP协议的139端口；后期SMB经过开发，可以直接运行于TCP/IP协议上，没有额外的NBT层，使用TCP协议的445端口。

## Samba的工作流程主要为四个阶段

### 协议协商

客户端在访问Samba服务器时，首先由客户端发送一个SMB negprot请求数据报，并列出它所支持的所有SMB协议版本。服务器在接收到请求信息后开始响应请求，并列出希望使用的协议版本，选择最优的SMB类型。如果没有可使用的协议版本则返回oXFFFFH信息，结束通信。



### 建立连接

当SMB协议版本确定后，客户端进程向服务器发起一个用户或共享的认证，这个过程是通过发送`SesssetupX`请求数据报实现的。客户端发送一对用户名和密码或一个简单密码到服务器，然后服务器通过发送一个`SesssetupX`请应答数据报来允许或拒绝本次连接。



### 访问共享资源

当客户端和服务器完成了协商和认证之后，它会发送一个Tcon或SMB TconX数据报并列出它想访问网络资源的名称，之后服务器会发送一个SMB TconX应答数据报以表示此次连接是否被接受或拒绝。



### 断开连接

连接到相应资源，SMB客户端就能够open SMB打开一个文件，通过read SMB读取文件，通过write SMB写入文件，通过close SMB关闭文件。



## Samba相关进程

Samba服务是由两个进程组成，分别是nmbd和smbd。

`nmbd`：其功能是进行NetBIOS名解析，并提供浏览服务显示网络上的共享资源列表。

`smbd`：其主要功能就是用来管理Samba服务器上的共享目录、打印机等，主要是针对网络上的共享资源进行管理的服务。当要访问服务器时，要查找共享文件，这时我们就要依靠smbd这个进程来管理数据传输。



## samba服务器的安全模式

samba服务器有share、user、server、domain和ads 五种安全模式，用来适应不同的企业服务器需求。

**（1）share安全级别模式**

客户端登录samba服务器，不需要输入用户名和密码就可以浏览samba服务器的资源，适用于公共的共享资源，安全性差，需要配合其他权限设置，保证samba服务器的安全性。

**（2）user安全级别模式**

客户端登录samba服务器，需要提交合法帐号和密码，经过服务器验证才可以访问共享资源，服务器默认为此级别模式。

**（3）server安全级别模式**

客户端需要将用户名和密码，提交到指定的一台samba服务器上进行验证，如果验证出现错误，客户端会用user级别访问。

**（4）domain安全级别模式**

如果samba服务器加入windows域环境中，验证工作服将由windows域控制器负责，domain级别的samba服务器只是成为域的成员客户端，并不具备服务器的特性，samba早期的版本就是使用此级别登录windows域滴。

**（5）ads安全级别模式**

当samba服务器使用ads安全级别加入到windows域环境中，其就具备了domain安全级别模式中所有的功能并可以具备域控制器的功能。