# gogo

简单易用的Go包管理工具

## 目标

* 使得 Go 项目可以脱离全局设置的 `GOPATH` 指定的路径
* 兼容 go 命令及其他第三方包管理工具，比如 dep
* 兼容主流编辑器

## 安装

```bash
go get -u github.com/leizongmin/gogo
```

## 使用

首先在包的根目录下创建包描述文件`package.yaml`:

```yaml
package: github.com/leizongmin/example
version: 0.0.0
```

说明：

* `package`为包的完整路径，比如 GitHub 用户名是`leizongmin`，要创建的包是`example`，那么包的完整路径就是`github.com/leizongmin/example`。别人在使用时可以通过`import "github.com/leizongmin/example"`来引入
* `version`为当前包的版本号，目前相关功能暂未实现

创建完`package.yaml`文件后，可以在包的根目录下执行以下命令：

```bash
# 第一步：初始化，创建虚拟 _workspace 目录和 vendor 目录
gogo init

# 第二步：添加依赖并保存到 package.yaml
gogo import github.com/leizongmin/leisp

# 第三步：构建
gogo build

# 开发时可以使用以下命令构建并直接执行
gogo dev
# 如果有参数可以直接在后面添加
gogo dev arg1 arg2
```

使用过程中还可以执行以下命令：

```bash
# 根据 pakcage.yaml 安装所有依赖
gogo install

# 执行 go 命令，以下命令相当于 go env
gogo - env

# 执行其他命令
gogo run env

# 删除 vendor 和 _workspace 目录
gogo clean
# 执行了 gogo clean 后，要继续使用时必须必须先执行 gogo init
```

## 工作目录`_workspace`

在执行`gogo init`之后，会创建一个`_workspace`目录，实际上这个目录为`GOPATH`环境变量的值，在配置编辑器的`GOPATH`时可设置为此`_workspace`目录。

可执行`gogo - env`查看 Go 打印出来的环境变量值。

### 在 Visual Studio Code 编辑器中使用

使用 **ms-vscode.go** 插件，添加以下编辑器配置（推荐修改文件 `.vscode/settings/json`）：

```json
{
  "go.gopath": "${workspaceRoot}/_workspace:/_YOUR_GLOBAL_GOPATH_"
}
```

说明：其中 `_YOUR_GLOBAL_GOPATH_` 原本的全局 GOPATH 路径，主要是用于 Go 插件安装必须的命令行工具。

## `gogo`开发环境配置

```bash
# 安装 gogo 工具
go get -u github.com/leizongmin/gogo

# 下载代码
git clone https://github.com/leizongmin/gogo.git && cd gogo

# 初始化
gogo init

# 安装依赖模块
gogo install

# 修改文件之后，编译代码
gogo build

# 执行编译出来的 gogo 命令
./bin/gogo
```

## 案例

* [gogo](https://github.com/leizongmin/gogo) - 简单易用的Go包管理工具
* [leisp](https://github.com/leizongmin/leisp) - The leisp programming language written in Go

## License

```text
MIT License

Copyright (c) 2016-2017 Zongmin Lei <leizongmin@gmail.com>

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
