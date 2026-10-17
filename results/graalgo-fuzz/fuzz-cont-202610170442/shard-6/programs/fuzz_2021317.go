package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) int32 {
    var v7 uint16 = 65535
    p0 += (p0 / ((uint8(v7)) | 1))
    v8 := GenS1{F0: -2147483648, F1: (v7 - v7)}
    v7 = uint16(p0)
    _ = p0
    _ = v7
    _ = v8
    return 283
}
func genFn1() uint8 {
    var v9 uint32 = 515
    v10 := GenS2{F0: true}
    v11 := []int{}
    var v12 uint8 = 1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 128
}

func main() {
    var v1 int8 = -128
    var v2 int = -2147483648
    v3 := []int{}
    var v4 int8 = v1
    v5 := []uint32{}
    v6 := make(map[uint32]bool)
    v6[2] = false
    switch "hello" {
    case "":
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint64 = 42
            var v15 uint32 = 65535
            _ = fi13
            _ = v14
            _ = v15
        }
    case "foo bar":
        _ = len(v6)
    case "\n\t":
        v16 := GenS1{F0: 4294967295, F1: 590}
        _ = v16
    default:
        v17 := GenS2{F0: (0.0 != 1.5)}
        _ = v17
    }
    v18 := GenS1{F0: v2, F1: 2}
    v4 = v1
    fmt.Println(v2, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v19 := v6[4294967295]
    v20 := GenS1{F0: v2, F1: 423}
    if false {
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := GenS1{F0: fi21, F1: 4}
            switch "hello" {
            case "foo bar":
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 float64 = ((-2.25 + (-2.25 + 0.0)) + (((3.14159 - 3.14159) * 0.1) - 0.1))
                    var v25 int16 = -32768
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            var v26 float32 = 3.14159
            _ = fi21
            _ = v22
            _ = v26
        }
    }
    var v27 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
