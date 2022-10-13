## afero

[toc]

## 概述

Afero 是一个文件系统框架，提供与任何文件系统交互的简单、统一和通用的 API，作为提供接口、类型和方法的抽象层。Afero 具有非常干净的界面和简单的设计，没有不必要的构造函数或初始化方法。

Afero 还是一个库，它提供了一组可互操作的后端文件系统的基本集，使使用 afero 变得容易，同时保留了 os 和 ioutil 包的所有功能和优势。

与单独使用 os 包相比，Afero 提供了显着改进，最显着的是能够在不依赖磁盘的情况下创建模拟和测试文件系统。

它适用于您考虑使用 OS 包的任何情况，因为它提供了一个额外的抽象，可以在测试期间轻松使用内存支持的文件系统。它还增加了对 http 文件系统的支持，以实现完全的互操作性。

### Afero 功能

+ 用于访问各种文件系统的单一一致 API
+ 多种文件系统类型之间的互操作
+ 一组接口来鼓励和加强后端之间的互操作性
+ 原子跨平台内存支持的文件系统
+ 通过将多个文件系统组合为一个来支持组合（联合）文件系统
+ 修改现有文件系统的专用后端（只读，Regexp 过滤）
+ 一组从 io、ioutil 和 hugo 移植的实用程序函数，以实现 afero 感知
+ go 1.16 文件系统抽象的包装器`io/fs.FS`

## 使用 Afero

Afero 易于使用且易于采用。

您可以使用 Afero 的几种不同方式：

+ 单独使用接口来定义您自己的文件系统。
+ 操作系统包的包装器。
+ 为应用程序的不同部分定义不同的文件系统。
+ 在测试时使用 Afero 模拟文件系统

### 第 1 步：安装 Afero

首先使用 go get 安装最新版本的库。

```
$ go get github.com/spf13/afero
```

接下来在您的应用程序中包含 Afero。

```
import "github.com/spf13/afero"
```

### 第 2 步：声明后端

首先定义一个包变量并将其设置为指向文件系统的指针。

```
var AppFs = afero.NewMemMapFs()

or

var AppFs = afero.NewOsFs()
```

重要的是要注意，如果您重复复合文字，您将使用一个全新的、独立的文件系统。对于 OsF，它仍将使用相同的底层文件系统，但会降低根据需要放入其他文件系统的能力。

### 第 3 步：像使用 OS 包一样使用它

在整个应用程序中，您可以像往常一样使用任何功能和方法。

因此，如果我之前的申请有：

```
os.Open("/tmp/foo")
```

我们将其替换为：

```
AppFs.Open("/tmp/foo")
```

`AppFs`是我们上面定义的变量。

### 所有可用功能的列表

可用的文件系统方法：

```
Chmod(name string, mode os.FileMode) : error
Chown(name string, uid, gid int) : error
Chtimes(name string, atime time.Time, mtime time.Time) : error
Create(name string) : File, error
Mkdir(name string, perm os.FileMode) : error
MkdirAll(path string, perm os.FileMode) : error
Name() : string
Open(name string) : File, error
OpenFile(name string, flag int, perm os.FileMode) : File, error
Remove(name string) : error
RemoveAll(path string) : error
Rename(oldname, newname string) : error
Stat(name string) : os.FileInfo, error
```

可用的文件接口和方法：

```
io.Closer
io.Reader
io.ReaderAt
io.Seeker
io.Writer
io.WriterAt

Name() : string
Readdir(count int) : []os.FileInfo, error
Readdirnames(n int) : []string, error
Stat() : os.FileInfo, error
Sync() : error
Truncate(size int64) : error
WriteString(s string) : ret int, err error
```

在某些应用程序中，定义一个简单地导出文件系统变量以便从任何地方轻松访问的新包可能是有意义的。

### 使用 Afero 的实用功能

Afero 提供了一组函数，使底层文件系统的使用变得更加容易。这些功能主要是从 io 和 ioutil 移植过来的，其中一些是为 Hugo 开发的。

afero 实用程序支持所有与 afero 兼容的后端。

实用程序列表包括：

```
DirExists(path string) (bool, error)
Exists(path string) (bool, error)
FileContainsBytes(filename string, subslice []byte) (bool, error)
GetTempDir(subPath string) string
IsDir(path string) (bool, error)
IsEmpty(path string) (bool, error)
ReadDir(dirname string) ([]os.FileInfo, error)
ReadFile(filename string) ([]byte, error)
SafeWriteReader(path string, r io.Reader) (err error)
TempDir(dir, prefix string) (name string, err error)
TempFile(dir, prefix string) (f File, err error)
Walk(root string, walkFn filepath.WalkFunc) error
WriteFile(filename string, data []byte, perm os.FileMode) error
WriteReader(path string, r io.Reader) (err error)
```

有关完整列表，请参阅[Afero 的 GoDoc](https://godoc.org/github.com/spf13/afero)

它们有两种不同的使用方法。您可以直接调用它们，其中每个函数的第一个参数将是文件系统，或者您可以声明一个新`Afero`的自定义类型，用于将这些函数作为方法绑定到给定的文件系统。

#### 直接调用实用程序

```
fs := new(afero.MemMapFs)
f, err := afero.TempFile(fs,"", "ioutil-test")
```

#### 通过 Afero 拨打电话

```
fs := afero.NewMemMapFs()
afs := &afero.Afero{Fs: fs}
f, err := afs.TempFile("", "ioutil-test")
```

### 使用 Afero 进行测试

使用模拟文件系统进行测试有很大的好处。每次初始化时它都处于完全空白状态，并且无论操作系统如何都可以轻松重现。您可以根据自己的喜好创建文件，并且文件访问速度会很快，同时还可以让您免于删除临时文件、Windows 文件锁定等所有烦人的问题。MemMapFs 后端非常适合测试。

+ 比在磁盘上执行 I/O 操作快得多
+ 避免安全问题和权限
+ 更多的控制。'rm -rf /' 充满信心
+ 测试设置要容易得多
+ 无需测试清理

实现此目的的一种方法是定义一个如上所述的变量。在您的应用程序中，这将在测试期间设置为 afero.NewOsFs()，您可以将其设置为 afero.NewMemMapFs()。

让每个测试初始化一个空白的 slate 内存后端并不少见。为此，我将`appFS = afero.NewOsFs()`在我的应用程序代码中定义我的适当位置。这种方法确保测试是顺序独立的，没有测试依赖于早期测试留下的状态。

然后在我的测试中，我会为每个测试初始化一个新的 MemMapFs：

```
func TestExist(t *testing.T) {
	appFS := afero.NewMemMapFs()
	// create test files and directories
	appFS.MkdirAll("src/a", 0755)
	afero.WriteFile(appFS, "src/a/b", []byte("file b"), 0644)
	afero.WriteFile(appFS, "src/c", []byte("file c"), 0644)
	name := "src/c"
	_, err := appFS.Stat(name)
	if os.IsNotExist(err) {
		t.Errorf("file \"%s\" does not exist.\n", name)
	}
}
```

## 可用的后端

### 操作系统本机

#### OsFs

第一个只是对本机操作系统调用的包装。这使得它非常易于使用，因为所有调用都与现有的操作系统调用相同。它还使您的代码在操作期间使用操作系统和在测试期间或根据需要使用模拟文件系统变得微不足道。

```
appfs := afero.NewOsFs()
appfs.MkdirAll("src/a", 0755)
```

### 内存支持存储

#### MemMapFs

Afero 还提供了一个完全原子内存支持的文件系统，非常适合用于模拟并在不需要持久性时加速不必要的磁盘 io。它是完全并发的，可以安全地在 go 例程中工作。

```
mm := afero.NewMemMapFs()
mm.MkdirAll("src/a", 0755)
```

##### 内存文件

作为 MemMapFs 的一部分，Afero 还提供了一个原子的、完全并发的内存支持文件实现。这可以轻松用于其他内存支持的文件系统。计划是使用 InMemoryFile 添加一个基数树内存存储文件系统。

### 网络接口

#### SftpFs

Afero 具有对安全文件传输协议 (sftp) 的实验性支持。可用于通过加密通道执行文件操作。

#### GCSF

Afero 为 Google Cloud Storage (GCS) 提供实验性支持。您可以将 `GOOGLE_APPLICATION_CREDENTIALS_JSON`env 变量设置为 JSON 凭证，也可以使用`opts`in `NewGcsFS`配置对 GCS 存储桶的访问权限。

现有实现的一些已知限制：

+ 不支持 Chmod - GCS ACL 可能会映射到 *nix 样式权限，但这会增加另一个级别的复杂性，并且在此版本中被忽略。
+ 不支持 Chtimes - 可以使用属性进行模拟（隐式设置 gcs a/m-times），但这留给另一个版本。
+ 不是线程安全的 - 还假设所有文件操作都是通过 GcsF 的同一实例完成的。不保证不同 GcsFs 实例之间的文件操作是一致的。

### 过滤后端

#### BasePathFs

BasePathFs 将所有操作限制在 Fs 内的给定路径上。在调用源 Fs 之前，此 Fs 上的操作的给定文件名将附加基本路径。

```
bp := afero.NewBasePathFs(afero.NewOsFs(), "/base/path")
```

#### 只读Fs

源 Fs 周围的薄包装器，提供只读视图。

```
fs := afero.NewReadOnlyFs(afero.NewOsFs())
_, err := fs.Create("/file.txt")
// err = syscall.EPERM
```

## 正则表达式

文件名的过滤视图，任何与传递的正则表达式不匹配的文件都将被视为不存在。不会创建与提供的正则表达式不匹配的文件。目录不被过滤。

```
fs := afero.NewRegexpFs(afero.NewMemMapFs(), regexp.MustCompile(`\.txt$`))
_, err := fs.Create("/file.html")
// err = syscall.ENOENT
```

#### HttpFs

Afero 提供了一个与 http 兼容的后端，它可以包装任何现有的后端。

Http 包需要一个稍微特殊的 Open 版本，它返回一个 http.File 类型。

Afero 提供了一个满足此要求的 httpFs 文件系统。任何 Afero 文件系统都可以用作 httpFs。

```
httpFs := afero.NewHttpFs(<ExistingFS>)
fileserver := http.FileServer(httpFs.Dir(<PATH>))
http.Handle("/", fileserver)
```

### 复合后端

Afero 提供了让两个（或更多）文件系统充当单个文件系统的能力。

#### CacheOnReadFs

CacheOnReadFs 将懒惰地将任何访问的文件从基础层复制到覆盖层中。随后的读取将直接从覆盖中提取，允许请求在覆盖中创建时的缓存持续时间内。

如果基本文件系统是可写的，对文件的任何更改都将首先对基本文件进行，然后对覆盖层进行。首先将调用写入打开文件句柄，如覆盖`Write()`或`Truncate()`覆盖。

要仅将文件写入覆盖，您可以直接使用覆盖 Fs（而不是通过联合 Fs）。

在给定时间缓存层中的文件。持续时间，缓存持续时间为 0 表示“永远”，这意味着永远不会从基础重新请求文件。

只读基础将使覆盖也成为只读，但当缓存层中不存在（或过时）文件时，仍将文件从基础复制到覆盖。

```
base := afero.NewOsFs()
layer := afero.NewMemMapFs()
ufs := afero.NewCacheOnReadFs(base, layer, 100 * time.Second)
```

#### CopyOnWriteFs()

CopyOnWriteFs 是一个只读的基础文件系统，顶部有一个潜在的可写层。

读取操作将首先在覆盖层中查找，如果在那里找不到，将从基础提供文件。

文件系统的更改只会在覆盖中进行。

任何修改仅在基础中找到的文件的尝试都会在修改之前将该文件复制到覆盖层（包括打开具有可写句柄的文件）。

当前不允许删除和重命名仅存在于基础层中的文件。如果文件存在于基础层和覆盖层中，则仅覆盖层将被删除/重命名。

```
	base := afero.NewOsFs()
	roBase := afero.NewReadOnlyFs(base)
	ufs := afero.NewCopyOnWriteFs(roBase, afero.NewMemMapFs())

	fh, _ = ufs.Create("/home/test/file2.txt")
	fh.WriteString("This is a test")
	fh.Close()
```

在此示例中，所有写入操作将仅发生在内存 (MemMapFs) 中，而基本文件系统 (OsFs) 保持不变。

### 所需/可能的后端

以下是我们希望有人实施的可能后端的简短列表：

+ SSH
+ S3

## 关于该项目

### 名字里有什么

Afero 来自拉丁词根 Ad-Facere。

**“广告”**是一个前缀，意思是“到”。

**“Facere”**是词根“faciō”构成“make or do”的一种形式。

afero 的字面意思是“制作”或“做”，这似乎非常适合允许人们制作文件和目录并用它们做事的库。

与 Afero 同根的英文单词是“affair”。Affair 具有相同的概念，但作为名词，它的意思是“制造或完成的事情”或“特定类型的对象”。

与我的其他一些库（雨果、眼镜蛇、毒蛇）不同，它的谷歌搜索非常好，这也很好。