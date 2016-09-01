# gogo
简单易用的Go包管理工具

**本项目正在开发中**

## 安装

```bash
$ go install github.com/leizongmin/gogo
```

## 使用

首先创建包描述文件`package.yaml`:

```yaml
package: github.com/leizongmin/example
version: 0.0.0
import:
- package: gopkg.in/yaml.v2
- package: github.com/Masterminds/vcs
  version: ^1.2.0
- package: github.com/codegangsta/cli
- package: github.com/Masterminds/semver
  version: ^1.0.0
```

命令：

```bash
# 初始化，创建虚拟 _workspace 目录
$ gogo init

# 添加依赖并保存到 package.yaml
$ gogo vendor github.com/leizongmin/leisp
# 如果要更新可以添加 -u 参数
$ gogo vendor -u github.com/leizongmin/leisp

# 根据 pakcage.yaml 安装所有依赖
$ gogo install

# 构建
$ gogo build

# 执行 go 命令，以下命令相当于 go env
$ gogo - env

# 执行 go 文件
$ gogo run xxx.go

# 清理
$ gogo clean
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
