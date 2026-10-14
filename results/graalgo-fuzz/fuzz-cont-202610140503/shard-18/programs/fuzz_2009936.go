package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 uint16 = 65535
    var v4 int = 42
    var v5 uint8 = 145
    var v6 string = ""
    v7 := GenS1{F0: 344, F1: uint64(v3), F2: v5}
    v6 = v6
    delete(v2, 787)
    v4 += (^(v4))
    v8 := []int{}
    v9 := GenS2{F0: v4, F1: (v4 + v4)}
    if false {
        if len(v8) > 0 {
            v8[0] = 32768
        }
        var v10 uint8 = (^(v5))
        _ = v10
    }
    if false {
        var v11 uint32 = 162
        var v12 bool = false
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    v4 = (v4 << ((uint(v5)) & 63))
    v13 := v2[uint32(v3)]
    var v14 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
