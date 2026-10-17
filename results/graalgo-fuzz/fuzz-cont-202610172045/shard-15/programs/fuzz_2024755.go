package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0(p0 uint, p1 string) uint8 {
    var v7 uint64 = 42
    v8 := []int{-275, 9223372036854775807, 503}
    fmt.Println(p1)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    p1 = "hello"
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 141
}
func genFn1() uint8 {
    var v9 string = "a"
    var v10 uint32 = 613
    v11 := GenS1{F0: -2147483648, F1: false, F2: false}
    _ = v10
    _ = v11
    _ = v9
    return 128
}

func main() {
    v1 := []uint32{2147483647, 0, 2}
    var v2 int8 = -128
    var v3 string = ""
    v4 := make(map[uint32]int)
    var v5 int32 = 1
    var v6 float64 = 1.5
    v12 := GenS2{F0: -9223372036854775808, F1: 796}
    v2 += (v2 << ((659) & 63))
    v3 = ((v3 + v3) + v3)
    fmt.Println(v2, v2)
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v5, v2)
        v14 := v4[2]
        _ = fi13
        _ = v14
    }
    delete(v4, 0)
    var v15 uint64 = 0
    v16 := GenS3{F0: v15, F1: -698, F2: 42}
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int8 = (v2 + v2)
        var v19 int8 = (v2 / ((45) | 1))
        var v20 float64 = (0.0 / v6)
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    v3 = ("\n\t" + (v3 + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
