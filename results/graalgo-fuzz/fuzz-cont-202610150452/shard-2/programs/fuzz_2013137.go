package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -1
    var v2 int32 = -182
    v3 := []string{"\n\t", "", "hello"}
    var v4 int8 = (v1 & (v1 & (^(v1))))
    var v5 int8 = (v4 / ((v4) | 1))
    v2 = v2
    v5 = 42
    var v6 float64 = 3.14159
    v6 += v6
    v3 = append(v3, ("foo bar" + "\n\t"))
    var v7 uint32 = 2147483648
    var v8 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
