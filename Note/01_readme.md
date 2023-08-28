# 数组 arr
## 声明定义
    pass
## 数组切片
- 切片的本质是视图
- 切片访问可以向后扩展，不可以向前扩展
- - 向后扩展不可以超过底层数据的长度
- - 索引访问不可以超过当前视图的长度

# Map
## 声明定义
    需要赋值, 若不赋值，则为 nil ：
    map [ keyType ] valueType
    map [ string ] int

    空map，不是 nil，而是 empty map:
    make(map [ keyType ] valueType)

```go
package main
// 赋值声明
var m1 = map[int]string {
    1: "a",
    2: "b",
    3: "c",
}

// 空map
m2 := make(map[int]string)

// nil
var m3 map[int]string
```

## key 底层原理
map 使用 hash 表，必须可以比较相等  
除了 slice，map，function 意外 的内建类型都可以作为 key，因为 slice，map，function 都不可以比较相等  

自定义 struct 类型不包含上述字段，也可以作为 key


## 遍历
遍历 map 时，返回的 key 是无序的，每次遍历的顺序都不一样  
使用 for range   
可是省略 key，只返回 value   
```go
package main

for key, value := range m1 {
    fmt.Println(key, value)
}

for key := range m1 {
fmt.Println(key)
}

for _, value := range m1 {
    fmt.Println(value)
}

```

## 访问 map

 使用 [key] 访问  
 若 key 不存在，不会报错，则返回 value 类型的零值
 
 其实是两个返回值，第一个是 value，第二个是 bool，表示是否存在

```go
package main

v1, isExist := m1[1]

```

删除 map 中的元素

    使用 delete 函数
    delete(map, key)
    若 key 不存在，不会报错


# 字符串
- 使用 utf8.RuneCountInString() 获取字符串长度    
- 使用 []byte() 转换为 byte 数组  
- 使用 len() 获取 byte字节 数组长度，并不是字符串长度    
- 使用 []rune() 转换为 rune 数组，rune 是 int32 的别名，表示一个 Unicode 码点  

strings.  包，提供了很多字符串操作函数
 
 
    Fields, Split, Join
    Contains, Index
    ToLower, ToUpper
    Trim, TrimRight, TrimLeft