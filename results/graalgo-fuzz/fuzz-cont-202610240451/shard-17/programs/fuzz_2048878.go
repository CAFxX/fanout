package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 18446744073709551615
    var v3 string = ("" + ("a" + "a"))
    v3 = "\n\t"
    var v4 uint32 = 2147483647
    var v5 float32 = 3.14159
    fmt.Println(v3, v4, v4)
    v3 = v3
    _ = len(v1)
    if (-1 == -1) {
        for fi6 := 0; fi6 < 7; fi6++ {
            v4 += (v4 + v4)
            _ = fi6
        }
        fmt.Println(v3)
        v7 := []int{1, 0, 2147483647}
        _ = v7
    }
    var v8 uint64 = v2
    var v9 uint16 = 65535
    v4 = v4
    var v10 bool = (-803 < 42)
    v11 := GenS1{F0: 42, F1: 4294967295}
    v2 = v2
    v12 := GenS1{F0: 42, F1: uint32(v8)}
    if v10 {
        v13 := v1[v4]
        _ = v13
    }
    v14 := GenS1{F0: -1, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
