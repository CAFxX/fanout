package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint16 }
func genFn0() uint32 {
    switch "hello" {
    case "\n\t":
        var v6 float32 = -2.25
        _ = v6
    case "hello":
        var v7 uint32 = 2
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 string = "hello"
        _ = fi8
        _ = v9
    }
    return 2147483647
}
func genFn1(p0 int8) uint64 {
    p0 = p0
    var v10 int8 = 1
    var v11 int16 = -32768
    var v12 uint = 255
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 170
}

func main() {
    var v1 uint16 = 65535
    v2 := []int{}
    var v3 uint32 = 267
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]bool)
    var v13 int16 = -1
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint32 = genFn0()
        if false {
            for fi16 := 0; fi16 < 14; fi16++ {
                if true {
                    v13 = -206
                    fmt.Println(v13, fi16)
                } else {
                    fmt.Println(fi14)
                    if false {
                        var v17 int = fi16
                        _ = v17
                    }
                    v15 = (^(v15))
                }
                v18 := GenS1{F0: -2147483648, F1: 128, F2: v1}
                _ = fi16
                _ = v18
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                _ = len(v2)
                var v20 int16 = v13
                v15 -= (^(v15))
                _ = fi19
                _ = v20
            }
        }
        _ = fi14
        _ = v15
    }
    if true {
        var v21 int64 = 9223372036854775807
        var v22 float32 = (0.0 + (3.14159 * (float32(3.14159) - 0.0)))
        _ = v21
        _ = v22
    } else {
        v2 = append(v2, 42)
    }
    var v23 uint64 = 42
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int8 = 100
        _ = fi24
        _ = v25
    }
    var v26 float64 = ((0.1 / 3.14159) + float64((0.0 * 0.1)))
    var v27 int32 = 0
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 uint16 = (v1 / (((v1 / (((v1 >> ((uint(v3)) & 63))) | 1))) | 1))
        _ = fi28
        _ = v29
    }
    v3 = (v3 | uint32(v23))
    var v30 uint16 = uint16(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
