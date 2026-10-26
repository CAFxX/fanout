package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
func genFn0(p0 string, p1 int32) uint8 {
    v6 := []int{317, -322}
    fmt.Println(p0, p1)
    p0 = ""
    _ = p0
    _ = p1
    _ = v6
    return 243
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]uint32)
    var v3 int16 = -32768
    var v4 uint32 = 398
    var v5 uint64 = 107
    var v7 int32 = 0
    switch 2147483647 {
    case 1:
        v1 = 65535
    default:
        var v8 int = -9223372036854775808
        _ = v8
    }
    var v9 uint32 = 4294967295
    if (-2.25 == 0.0) {
        var v10 int = -1
        fmt.Println(v10)
        var v11 int = v10
        _ = v10
        _ = v11
    }
    if false {
        v12 := []int{}
        var v13 int16 = -1
        v14 := GenS1{F0: true, F1: -9223372036854775808, F2: "\n\t"}
        _ = v12
        _ = v13
        _ = v14
    }
    v3 += 1
    var v15 string = ""
    fmt.Println(v7)
    for fi16 := 0; fi16 < 14; fi16++ {
        fmt.Println(v7, v5)
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
