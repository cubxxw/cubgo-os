# 搭建开发工具

[toc]

## 基本工具

基本工具如下：

+ go1.13.x
+ mage构建工具
+ gcc和gdb，需要能生成elf32格式的。
+ qemu虚拟机

各平台的安装步骤如下所述。

**Go语言安装包：**

```bash
Go语言编译器 #
$ curl -LO https://golang.google.cn/dl/go1.13.15.linux-amd64.tar.gz
$ mkdir $HOME/local && tar xf -C $HOME/local go1.13.15.linux-amd64.tar.gz
# 将$HOME/local/go/bin加入PATH环境变量，这里以bash为例
$ echo 'export PATH=$HOME/local/go/bin:$PATH' >> $HOME/.bashrc
# 将$HOME/go/bin也加入环境变量，方便执行编译出来的go工具
$ echo 'export PATH=$HOME/go/bin:$PATH' >> $HOME/.bashrc
$ source $HOME/.bashrc
# 设置GOPROXY环境变量，加速package下载
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn
```



### mage构建工具

mage是一个类似make的构建工具，但它的makefile是用go语言编写的，可以跨平台，eggos最初是用的makefile，后面切换到mage上了。安装mage也非常简单。

```sh
$ go get github.com/magefile/mage
```



### C语言开发套件

```sh
$ sudo apt-get install build-essential gdb
```



### Qemu虚拟机

```sh
$ sudo apt-get install qemu-system-x86
```



## IDE开发环境

我们使用`vscode`来开发eggos。项目默认附带了vscode的配置文件，直接使用vscode打开eggos项目即可开始修改代码。

### 安装的插件

+ [Go语言插件](https://marketplace.visualstudio.com/items?itemName=golang.go)，用于代码补全，跳转与高亮
+ [GDB插件](https://marketplace.visualstudio.com/items?itemName=webfreak.debug)，用于对内核进行debug，如断点调试



### Debug配置 

需要根据实际情况稍微修改一下`.vscode/launch.json`里面的配置

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
         {
            "name": "dlv-gdb",
            "type": "go",
            "request": "attach",
            "mode": "local",
            "cwd": "${workspaceFolder}",
            "debugAdapter": "legacy",
            "processId": 1234,
            "backend": "gdbstub",
            "dlvFlags": [
                "${workspaceRoot}/kernel.elf",
            ],
        },
        {
            "type": "gdb",
            "request": "attach",
            "name": "Attach to qemu",
            "executable": "./kernel.elf",
            "target": ":1234",
            "remote": true,
            "cwd": "${workspaceRoot}",
            "valuesFormatting": "parseText",
            // 这里根据情况进行替换，linux就用gdb
            "gdbpath": "gdb", 
            // MacOS用下面配置
            // "gdbpath":"x86_64-elf-gdb",  
        },
    ]
}
```

