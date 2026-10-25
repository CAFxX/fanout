package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int8 }
func genFn0() uint {
    var v5 uint8 = 42
    if (!(!false)) {
        var v6 float32 = 0.1
        v7 := GenS1{F0: 100, F1: 42, F2: 348}
        var v8 int8 = -1
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 uint64 = 2
        var v10 uint = 4294967295
        v11 := GenS2{F0: "a", F1: 30}
        _ = v10
        _ = v11
        _ = v9
    }
    _ = v5
    return 0
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]uint32)
    var v3 uint16 = 255
    v4 := make(map[uint32]string)
    if false {
        if true {
            v3 = v3
            v12 := GenS2{F0: "", F1: -1}
            _ = v12
        } else {
            v13 := []int{4294967295, -2147483648}
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v1, v3)
            _ = fi14
        }
    }
    var v15 uint32 = 65535
    v16 := []int{2147483647, 127}
    var v17 int8 = 100
    if (644 != 824) {
        v3 = 583
    }
    var v18 uint32 = v15
    switch "a" {
    case "hello":
        v17 -= 93
    case "a":
        for fi19 := 0; fi19 < 6; fi19++ {
            v15 = uint32(v1)
            _ = fi19
        }
    }
    v18 = 410
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 uint32 = (v18 - v18)
            _ = fi21
            _ = v22
        }
        fmt.Println(v3, v18)
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
