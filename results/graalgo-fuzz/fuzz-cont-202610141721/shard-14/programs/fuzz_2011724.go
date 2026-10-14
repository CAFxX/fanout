package main

import (
    "fmt"
)


func main() {
    var v1 uint = 18446744073709551615
    v2 := []uint32{2, 84, 1}
    var v3 string = ((("a" + "\n\t") + "a") + "")
    var v4 int = 0
    var v5 uint32 = 2147483647
    v5 = 4294967295
    v2 = append(v2, 65535)
    var v6 float32 = 0.1
    var v7 string = v3
    var v8 uint64 = 42
    fmt.Println(v6, v8)
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
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
