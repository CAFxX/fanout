package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0() int8 {
    var v4 int32 = 0
    var v5 uint32 = 2
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint32 = (v1 | v1)
    var v3 uint64 = 144
    v6 := []int{4294967295}
    if false {
        fmt.Println(v3, v1)
        switch -445 {
        case 0:
            var v7 float32 = -2.25
            _ = v7
        case 32767:
            v8 := GenS1{F0: 512, F1: true, F2: 32767}
            _ = v8
        }
    }
    v2 -= v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!false) {
        v2 = (^(v2))
        var v9 uint = 65535
        v10 := GenS2{F0: -2147483648, F1: 24}
        _ = v10
        _ = v9
    }
    var v11 uint16 = 255
    if (1.5 != 1e10) {
        v12 := GenS1{F0: 42, F1: true, F2: v11}
        _ = v12
    } else {
        var v13 int32 = -225
        var v14 string = "hello"
        var v15 uint32 = 76
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(v11)
    var v16 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
