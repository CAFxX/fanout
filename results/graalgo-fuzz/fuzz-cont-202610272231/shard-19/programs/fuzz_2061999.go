package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8; F1 int32; F2 string }

func main() {
    v1 := []uint32{2147483647, 0, 98, 0}
    var v2 float64 = 0.1
    var v3 float64 = (v2 + -2.25)
    var v4 int16 = -32768
    var v5 uint64 = 2
    var v6 uint = 2
    for fi7 := 0; fi7 < 7; fi7++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            var v8 int64 = 0
            _ = v8
        } else {
            var v9 uint64 = (v5 << (((v6 << ((v6) & 63))) & 63))
            var v10 uint = 42
            var v11 int64 = 518
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 int64 = -1
        _ = fi7
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v5 += v5
        fmt.Println(v4, v2, v6)
        _ = fi13
    }
    v4 = v4
    v6 += v6
    var v14 int32 = 0
    fmt.Println(v5)
    var v15 uint = 1
    v16 := GenS1{F0: uint32(v15), F1: true}
    var v17 int64 = 1
    switch "" {
    case "hello":
        v18 := GenS1{F0: 984, F1: true}
        _ = v18
    default:
        v15 = (v6 - 2)
    }
    if true {
        fmt.Println(v6)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v6, v6, v5)
    v19 := GenS2{F0: 1, F1: v14, F2: ("foo bar" + "")}
    var v20 int16 = (v4 + (v4 ^ v4))
    v21 := GenS1{F0: 65535, F1: true}
    var v22 uint32 = 816
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
