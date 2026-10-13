package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    v1 := []string{"hello", ""}
    v2 := make(map[uint32]string)
    var v3 int8 = 100
    var v4 bool = (false && false)
    var v5 int64 = -513
    v3 -= v3
    v6 := GenS1{F0: (("hello" + "\n\t") + ("hello" + "")), F1: 42}
    v3 = -1
    v5 += v5
    var v7 int32 = 2147483647
    for fi8 := 0; fi8 < 15; fi8++ {
        if len(v1) > 0 {
            v1[0] = ""
        }
        fmt.Println(v7, v3, v3)
        var v9 int64 = 0
        _ = fi8
        _ = v9
    }
    v4 = v4
    for fi10 := 0; fi10 < 16; fi10++ {
        v3 += (v3 / ((v3) | 1))
        _ = fi10
    }
    fmt.Println(v5, v7)
    var v11 uint16 = 42
    v2[2] = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
