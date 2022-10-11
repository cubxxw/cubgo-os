# cubgo-os

 ⚡cubgo-os（幼兽操作系统）由Go语言开发，可以运行基础应用~



## 背景

受清华的`rcore`项目影响



+ [在 x86 裸机上运行的 Go unikernel](https://github.com/icexin/eggos)  ==》 [文档](https://eggos.icexin.com/)

> 使用Go语言编写的运行在x86裸机上的unikernel
>
> 将Go程序运行在x86裸机上，全部使用Go语言编写(bootloader里面有少量汇编和c)，支持大部分Go的核心功能(GC，goroutine等)，包括大部分Go的标准库，同时也附带一个网络协议栈，从而能运行大部分基于`net`库的第三方库。
>
> 不同于传统的操作系统内核，eggos没有对用户态和内核态的代码进行隔离，整个unikernel运行在一个地址空间上。没有进程的概念以及进程间通信，但是完整支持了Go的goroutine和channel。另外也没有传统的ELF加载器，但是附带了一个Javascript的解释器。



+ [饼干研究操作系统](https://github.com/mit-pdos/biscuit)

> Biscuit 是用于 x86-64 CPU 的 Go 中的单片 POSIX 子集操作系统内核。编写它是为了研究使用带有垃圾收集的高级语言来实现具有通用架构风格的内核的性能权衡。你可以在这里找到关于 Biscuit 的研究论文： [https ://pdos.csail.mit.edu/projects/biscuit.html](https://pdos.csail.mit.edu/projects/biscuit.html)
>
> Biscuit 具有一些重要特性，可用于获得良好的应用程序性能：
>
> + 多核
> + 内核支持的线程
> + 具有并发、延迟和组提交的日志式 FS
> + 用于写入时复制和延迟映射的匿名/文件页面的虚拟内存
> + TCP/IP 堆栈
> + AHCI SATA 磁盘驱动程序
> + 英特尔 10Gb 网卡驱动程序
>
> Biscuit 还包括一个引导加载程序、一个部分 libc（“litc”）和一些用户空间程序，尽管我们可以使用 GRUB 或现有的 libc 实现，如 musl。
>
> 这个 repo 是 Go repo ( https://github.com/golang/go ) 的一个分支。几乎所有 Biscuit 的代码都在 biscuit/ 中。
>
> **功能列表**
>
> + Go的内置功能，如GC，goroutine，channel等
> + 一个支持行编辑的终端
> + 支持TCP/IP的协议栈
> + Go风格的VFS抽象，使用[afero](https://github.com/spf13/afero)
> + NES模拟器，使用[nes](https://github.com/fogleman/nes)
> + Javascript解释器，使用[otto](https://github.com/robertkrimen/otto)
> + GUI支持，使用[nucular](https://github.com/aarzilli/nucular)
> + 一些简单的应用，如(httpd, sshd)



+ [mit 地鼠操作系统](https://github.com/gopher-os/gopher-os)  ==> [官网](https://pdos.csail.mit.edu/projects/biscuit.html)

> 该项目的目标是使用[Go](https://golang.org/)构建一个 64 位 POSIX 兼容的无滴答内核，并使用与 Linux 兼容的系统调用实现。
>
> 这个项目不是要构建另一个操作系统，而是为了证明 Go 确实是编写在 ring-0 运行的低级代码的合适工具。
>
> **注意**：该项目仍处于早期开发阶段，尚未处于可用状态。事实上，如果你构建 ISO 并启动它，内核最终会因为`Kmain returned`错误而恐慌。
>
> 要了解有关当前项目状态和功能路线图的更多信息，请查看[状态](https://github.com/gopher-os/gopher-os/blob/master/STATUS.md)页面。



## 分析

Go语言编写操作系统内核一直是一个冷门的领域，难以想一个带runtime的语言如何在没有操作系统的裸机上启动运行，让语言自己的功能能正常运行都难，更别说做操作系统常见的内存管理，进程管理等。但也有人不断地尝试，比如以下几个项目：

**1.tinny**

[https://github.com/golang/go/commit/434f1e32a075d546f47943598aa5974d4a2492cegithub.com/golang/go/commit/434f1e32a075d546f47943598aa5974d4a2492ce](https://link.zhihu.com/?target=https%3A//github.com/golang/go/commit/434f1e32a075d546f47943598aa5974d4a2492ce)

这个是Go官方的早期版本的一个实现，能跑concurrent prime sieve，现在已经删掉了，可以看到在Go早期版本里面runtime还没有那么复杂的时候，移植到裸机其实挺简单的，想考古的同学可以试一下，里面用到的bootloader是著名的[xv6](https://link.zhihu.com/?target=https%3A//github.com/mit-pdos/xv6-public)的实现，因为Russ Cox自己就是xv6的作者之一。

**2.gopher-os**

[https://github.com/gopher-os/gopher-osgithub.com/gopher-os/gopher-os](https://link.zhihu.com/?target=https%3A//github.com/gopher-os/gopher-os)

这个通过修改go的编译过程以及重定向runtime符号实现了一个概念内核，项目本证明了go能脱离操作系统运行在裸机上，最终的效果是在图形化终端上打印了一下硬件信息，没有过多的其他应用。

**3.biscuit**



这是一个博士论文项目，完成度很高，多核，多线程，网络协议栈等等。这个项目是奔着写一个操作系统内核去的，其兼容了一些POSIX接口，根据论文里面的内容，可以跑redis和nginx。实现方式是hack了linux的amd64 arch相关代码，注入了自己的实现，代码整体分为两部分，一部分存在于go的标准库里面的runtime目录下，一部分是在项目的biscuit目录下。

我最早的想法也是借鉴这个项目，不过不是hack runtime的某个arch和os，而是新造了一个GOOS=eggos，编译的时候指定这个GOOS环境变量就能生成相应的内核，不过为了能让工具链和标准库能编译通过，修改了很多标准库代码，大部分只是加了一个编译tag，比如这个文件的这行代码，就需要在开头加上我自己的操作系统名称，类似这样的文件很多，这样做其实是不利于跟进go的版本升级的。



**4.unik**

[~eliasnaur/unik - sourcehut gitgit.sr.ht/~eliasnaur/unik](https://link.zhihu.com/?target=https%3A//git.sr.ht/~eliasnaur/unik)

这个项目是我在找Go实现的平台无关的GUI项目时发现的，作者写了一个Go的GUI项目，同时为了证明其跨平台能力用Go写了一个运行于裸机的应用程序，显卡驱动使用了virtio。作者还是挺牛的。我后来又深挖了一下，发现他之前就在Go的issue提过想让官方提供一个GOOS=none的实现，issue里的很多资料还是很有用的，具体讨论见这里

[https://github.com/golang/go/issues/35956github.com/golang/go/issues/35956](https://link.zhihu.com/?target=https%3A//github.com/golang/go/issues/35956)

这个项目对Go runtime的启动方式对我启发很大，在看到这个项目后，我把我之前的实现改成类似的方式，之后就摆脱了对Go runtime代码库的依赖。

总结一下几种实现裸机上运行Go的方式：

1. 直接加一个新的GOOS，GOARCH看自己要实现的平台，这种方式需要修改runtime，编译器以及标准库的代码，优点是实现比较直接，并且可以做一些定制，想更改runtime行为的时候代码不绕弯，缺点是与特定Go版本的代码绑定，不利于后续升级。
2. 魔改Go的编译链接过程，将runtime里面的os相关实现链接到自己的实现上，本质上跟第一种方案类似，但没有修改runtime代码，缺点是Makefile太复杂，并且很依赖特定版本的编译器。
3. 更改最终生成的ELF程序的入口，引导到自己实现的初始化代码，最后再跳转到Go的runtime入口。这种方案是目前我觉得最简洁的方式，即不用修改runtime代码，又用的是标准工具的能力，之后兼容更高的Go版本也不会有大改。

当然让Go程序在裸机上跑起来远不止上面说的那些，上面说的只是解决了runtime的bootstrap的问题，但让Go程序使用基本的new, interface，goroutine，channel等，我们至少还要解决以下几个问题：

1. **内存**，Go是一个带内存管理的语言，内存问题是我们首要解决的，不然就回到手工管理内存的时代。好消息是Go的runtime对OS的内存接口依赖很挺简单，主要有sysAlloc，sysReserve，sysMap等，如果采用第一种方式实现内核，可以参考wasm(runtime/mem_js.go)的实现，由于wasm整个是一个线性内存，跟裸机的内存模型一致，因此很有参考性。如果是后面几种实现，就要模拟linux的mmap系统调用，这个后面再说。
2. **线程，**线程对应Go GMP模型里面的M，用于执行某个Goroutine的代码。如果是第一种实现，可以参考wasm在runtime里面的代码，wasm是单线程实现，因此不用考虑多个M的情况，结果是wasm里面没有启动sysmon协程，纯粹靠编译器打桩的代码来实现协程调度。如果是后面几种实现，就要模拟linux的clone系统调用，这个后面再说。
3. **锁，**这里的锁特指runtime里面使用的**mutex和note**两种很重要的同步设施，主要用于同步多个M的。但在wasm里面因为只有一个线程，因此wasm的线程相关实现几乎是空操作（除了notesleep)，但Linux里面就用的是实实在在的futex实现的。还是看我们的实现，第一种方式就可以参考wasm的实现(runtime/lock_js.go)，如果是如果是后面几种实现，就要模拟linux的futex系统调用。

实现了上面几个基本功能之后，就能写一些纯内存操作的Go程序了，goroutine，channel啥的随便用，但要想使用标准库里面的文件，网络等功能还是需要后续的两个重要工作，**中断和系统调用**。

平时我们写Go用户应用程序的时候是不关注中断和系统调用这些概念的，因为Go的标准库帮我们封装了这些事情。比如我们写一个网络程序，在Go里面直接调用net.Conn.Read阻塞等待数据返回就可以了，在没有数据的时候，go的runtime挂起当前goroutine，同时在runtime内部使用了epoll_wait系统调用来等待fd的可读事件。在操作系统层面，收到epoll wait调用后，操作系统挂起当前等待的线程，调度其他进程。等网卡收到了数据包的时候，触发了操作系统的网络中断例程，经过操作系统的网络栈处理原始数据包之后，操作系统又唤醒了epoll wait线程，告诉它哪些fd可读，Go runtime这个时候再唤醒调用Read的goroutine来读取数据。

上面的整个过程对普通用户是透明的，但我们要跑在裸机上，上面的应用程序到硬件之间的Gap就需要我们亲自去填补了。关于如何处理中断和系统调用本身就可以另外写一篇文章了，这里简单说一下遇到的难点。

前面说过，Go除了sysmon协程辅助调度之外，在函数入口以及系统调用的地方加了一些指令来检查栈是否溢出，以及是否应该调度当前协程等。这些在用户程序里面没有问题，但在内核代码里面，这些就很致命了，这些检查代码一方面会假定当前有可用的G或者P，另一方面会发生潜在的协程调度，而在中断或者系统调用代码里面，是不能假定有可用的G或者P的，也不能发生调度，如何处理好这些问题真的很难。好消息是，Go runtime也会遇到这些问题，因此我们有现成的代码可借鉴，但坏消息是，写这些代码真的是如履薄冰，任何高级的Go 特性都不能使用，一点都不欢乐，因此大家看到我的这个项目里面晦涩代码的时候，其实当时写这些代码的人也很无奈。然而还有另外一个好消息，我把这些问题通过类似微内核的思想给解决了，具体可以看一下我的这个Go项目的代码。

说到这里就给大家介绍一下我写的这个项目。最早我萌生编写将Go跑在裸机上的项目是觉得Go的runtime做了很多类似操作系统的事情，比如goroutine对应进程，channel对应进程间通信，Go还有完善的内存管理，关键是Go编译后是一个二进制，不需要虚拟机。因为我上学的时候用C写过一个miniOS，觉得Go可能做一个简单的适配层就可以跑在裸机上了，说干就干，疫情期间有很多时间让我去思考和实验这个想法，最后还真成功了。

[https://github.com/icexin/eggosgithub.com/icexin/eggos](https://link.zhihu.com/?target=https%3A//github.com/icexin/eggos)



## 类似项目

+ 现代操作系统

+ [清华大学rcore项目](http://rcore-os.cn/rCore-Tutorial-Book-v3/)

> 用rust实现操作系统，这个项目目前第三版，很成熟了，我很喜欢
