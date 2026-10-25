package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }

func main() {
    var v1 uint16 = 600
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483647
    var v4 uint64 = 4294967295
    if true {
        var v5 string = "hello"
        if ((true || (false || false)) && (!false)) {
            v1 += (v1 - (^(v1)))
        }
        var v6 bool = true
        _ = v5
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 42, F1: (v7 ^ v7), F2: 1}
    if (true && (0.0 >= (3.14159 * 3.14159))) {
        v9 := []int{-9223372036854775808, int(v7)}
        _ = v9
    }
    var v10 string = "foo bar"
    var v11 float32 = (1e10 / float32(0.1))
    v12 := GenS1{F0: 65535, F1: v7, F2: 363}
    fmt.Println(v10)
    v13 := GenS1{F0: 2, F1: v7, F2: v1}
    for fi14 := 0; fi14 < 16; fi14++ {
        v4 -= 2
        fmt.Println(v3)
        _ = fi14
    }
    var v15 bool = ((107 > 2) || true)
    var v16 bool = v15
    if v15 {
        var v17 uint32 = 0
        v18 := []int{int(v7), 11}
        _ = v17
        _ = v18
    }
    v19 := []int{}
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
