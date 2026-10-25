package main

import (
    "fmt"
)


func main() {
    var v1 string = ("a" + "hello")
    v2 := []uint32{416, 4294967295, 1}
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 56
    fmt.Println(v1, v5)
    v5 = v5
    fmt.Println(v5, v5)
    if false {
        v6 := v4[795]
        _ = v6
    }
    v5 = (^(v5))
    var v7 uint16 = 1
    fmt.Println(v5)
    delete(v3, 4294967295)
    var v8 int8 = -97
    fmt.Println(v7, v7, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
