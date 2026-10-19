package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint16) uint8 {
    v7 := []int{-540, 4294967295}
    var v8 int64 = -1
    var v9 float64 = float64((0.1 / 1.0))
    var v10 uint64 = uint64(p1)
    v9 -= v9
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 46
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 129
    var v3 uint = (v2 + 65535)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 0
    if ((1 >= -1) || false) {
        var v11 uint32 = v6
        var v12 float32 = 1.5
        _ = v11
        _ = v12
    }
    if true {
        v13 := GenS1{F0: 42, F1: true, F2: false}
        delete(v5, v6)
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v2 += 255
        _ = fi14
    }
    v15 := []int{}
    v4["\n\t"] = (519 <= -699)
    var v16 uint64 = 4294967295
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = len(v15)
    var v17 uint64 = 156
    _ = len(v4)
    switch -400 {
    case 128:
        v3 = (^(v2))
    case 862:
        v1 = v1
    }
    v18 := GenS2{F0: "hello", F1: 32768}
    var v19 float32 = (((1e10 / 0.1) - 0.0) - float32(float64(3.14159)))
    v20 := GenS3{F0: 4}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
