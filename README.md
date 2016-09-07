# gogo
简单易用的Go包管理工具

## 安装

```bash
$ go get -u github.com/leizongmin/gogo
```

## 使用

首先在包的根目录下创建包描述文件`package.yaml`:

```yaml
package: github.com/leizongmin/example
version: 0.0.0
```

说明：

+ `package`为包的完整路径，比如 GitHub 用户名是`leizongmin`，要创建的包是`example`，那么包的完整路径就是`github.com/leizongmin/example`。别人在使用时可以通过`import "github.com/leizongmin/example"`来引入
+ `version`为当前包的版本号，目前相关功能暂未实现

创建完`package.yaml`文件后，可以在包的根目录下执行以下命令：

```bash
# 第一步：初始化，创建虚拟 _workspace 目录和 vendor 目录
$ gogo init

# 第二步：添加依赖并保存到 package.yaml
$ gogo import github.com/leizongmin/leisp
# 如果要更新可以添加 -u 参数
$ gogo import -u github.com/leizongmin/leisp

# 第三步：构建
$ gogo build
```

使用过程中还可以执行以下命令：

```bash
# 根据 pakcage.yaml 安装所有依赖
$ gogo install

# 执行 go 命令，以下命令相当于 go env
$ gogo - env

# 删除 vendor 和 _workspace 目录
$ gogo clean
# 执行了 gogo clean 后，要继续使用时必须必须先执行 gogo init
```


## 工作目录`_workspace`

在执行`gogo init`之后，会创建一个`_workspace`目录，实际上这个目录为`GOPATH`环境变量的值，在配置编辑器的`GOPATH`时可设置为此`_workspace`目录。

可执行`gogo - env`查看 Go 打印出来的环境变量值。


## `gogo`开发环境配置

```bash
# 安装 gogo 工具
$ go install github.com/leizongmin/gogo

# 下载代码
$ git clone https://github.com/leizongmin/gogo.git && cd gogo

# 安装依赖模块
$ gogo install

# 修改文件之后，编译代码
$ gogo build

# 执行编译出来的 gogo 命令
$ ./bin/gogo
```


## License

```
MIT License

Copyright (c) 2016 Zongmin Lei <leizongmin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
