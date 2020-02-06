# go struct拷贝工具
字段大概相似的struct可以通过该工具进行copy，可以简化代码，加快开发效率。

### 发生拷贝的条件：

1. 字段名字一样
2. 字段数据类型一样
3. 字段数据类型不同包，同类型名

### 使用方式举例：

```go
type DestStruct struct {
    Id int64
    Name string
}

var srcStruct = struct {
    Id int64
    Name string
} {
    Id: 10000,
    Name: "test"
}

structUtils := NewStructUtils()

var destStruct DestStruct
structUtils.CopyProperties(&destStruct, &srcStruct)
```

### 基准测试结果如下：

```text
goos: darwin
goarch: amd64
pkg: structutils
BenchmarkStructUtils_CopyProperties-12    	  500000	      2525 ns/op
PASS

Process finished with exit code 0
```

如果在使用过程遇到问题，或者发现bug，或者有更好的建议可以发邮件给我！ 欢迎沟通交流！