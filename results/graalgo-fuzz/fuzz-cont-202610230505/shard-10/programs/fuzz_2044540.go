package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int32 }
func genFn0() string {
    v5 := []int{1, 32768}
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: fi6}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    var v8 uint16 = 65535
    _ = v5
    _ = v8
    return ("\n\t" + ("foo bar" + "\n\t"))
}

func main() {
    var v1 uint8 = 233
    v2 := []uint32{2147483647, 42, 0}
    var v3 int16 = -32768
    v4 := make(map[uint32]int)
    v3 = -1
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v3, fi9, fi9)
        v1 = (v1 & (v1 & v1))
        var v10 int64 = -1
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 42
    v12 := GenS1{F0: -9223372036854775808}
    if (v11 == (^(v1))) {
        v11 = v1
        _ = len(v2)
        v13 := GenS2{F0: 32767, F1: 232}
        _ = v13
    }
    v3 = v3
    for fi14 := 0; fi14 < 8; fi14++ {
        v1 -= v1
        var v15 uint = 812
        _ = fi14
        _ = v15
    }
    var v16 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
