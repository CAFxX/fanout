package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 1
    v2 := make(map[string]bool)
    var v3 bool = true
    v4 := []string{(("\n\t" + "hello") + ""), "a", ("a" + ("a" + "hello"))}
    var v5 int8 = 0
    var v6 string = "\n\t"
    _ = len(v4)
    v5 = v5
    var v7 uint32 = 912
    var v8 int16 = 1
    v1 = (v1 + v1)
    v8 = v8
    var v9 float32 = (float32(0.1) + (1e10 - 1e10))
    fmt.Println(v3, v5)
    fmt.Println(v8)
    var v10 int64 = 0
    delete(v2, ("" + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
