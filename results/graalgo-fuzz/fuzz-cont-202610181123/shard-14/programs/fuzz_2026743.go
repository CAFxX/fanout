package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    v1 := []uint32{2147483647, 42}
    var v2 int64 = -136
    var v3 float32 = (-2.25 + 3.14159)
    var v4 string = ""
    fmt.Println(v3)
    var v5 bool = (false || true)
    v1 = append(v1, 480)
    var v6 uint32 = 42
    v7 := GenS1{F0: 42, F1: 2}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 bool = (!(65535 == 1))
    var v9 uint16 = 930
    v1 = append(v1, (^(v6)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
