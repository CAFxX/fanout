package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32 }
func genFn0() uint {
    v7 := GenS1{F0: -1}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint8 = 128
            _ = fi9
            _ = v10
        }
        var v11 float64 = -2.25
        v11 += v11
        _ = fi8
        _ = v11
    }
    var v12 uint64 = 0
    _ = v12
    _ = v7
    return 0
}
func genFn1() int64 {
    v13 := GenS1{F0: 42}
    var v14 uint32 = 42
    fmt.Println(v14)
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 int32 = -210
    v2 := make(map[int]int)
    v3 := []int{42, 32767, 1, 1}
    v4 := make(map[int]int)
    var v5 int8 = -1
    var v6 string = "\n\t"
    var v15 int32 = (v1 * 2147483647)
    var v16 float64 = ((1.5 * float64(float32((0.0 * 1e10)))) - 1.5)
    var v17 uint64 = 42
    if false {
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 uint8 = 42
            _ = fi18
            _ = v19
        }
        if (!true) {
            var v20 uint16 = 0
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := GenS2{F0: v17, F1: genFn1()}
                v23 := GenS3{F0: 0}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        }
        v5 = (v5 - v5)
    }
    if false {
        v1 = v1
        var v24 uint32 = 2
        _ = v24
    }
    _ = len(v4)
    v5 += v5
    var v25 uint16 = 1
    v15 += -613
    for fi26 := 0; fi26 < 2; fi26++ {
        v5 = v5
        _ = fi26
    }
    switch "hello" {
    case "hello":
        if (42 != 2) {
            var v27 uint32 = 2
            v28 := GenS2{F0: v17, F1: 0}
            _ = v27
            _ = v28
        }
    case "a":
        var v29 float32 = float32(1.5)
        _ = v29
    default:
        var v30 uint = 154
        _ = v30
    }
    v31 := GenS2{F0: v17, F1: 0}
    var v32 int = 9223372036854775807
    var v33 string = v6
    v34 := v4[v32]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
