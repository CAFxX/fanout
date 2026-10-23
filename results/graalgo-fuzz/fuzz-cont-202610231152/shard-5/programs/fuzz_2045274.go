package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 uint16; F1 int; F2 int }

func main() {
    var v1 float32 = 1.5
    v2 := []int{}
    var v3 float32 = float32(0.0)
    var v4 uint64 = 645
    v5 := make(map[uint32]bool)
    var v6 int16 = -1
    switch ("" + "a") {
    case "":
        var v7 string = "foo bar"
        _ = v7
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    default:
        fmt.Println(v6)
    }
    var v8 uint = 416
    var v9 uint = v8
    v10 := v5[42]
    v5[42] = v10
    switch "a" {
    case "a":
        v11 := GenS1{F0: 1, F1: 65535, F2: -1}
        _ = v11
    case "hello":
        v1 = v3
    case "\n\t":
        v12 := GenS1{F0: int64(v9), F1: 715, F2: 1}
        _ = v12
    default:
        if true {
            var v13 uint32 = 65535
            var v14 bool = (!(v6 != -381))
            _ = v13
            _ = v14
        }
    }
    v1 = v1
    for fi15 := 0; fi15 < 10; fi15++ {
        if len(v2) > 0 {
            v2[0] = (fi15 ^ fi15)
        }
        var v16 string = "a"
        _ = fi15
        _ = v16
    }
    v3 = v1
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS2{F0: 2, F1: fi17, F2: fi17}
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
