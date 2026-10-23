package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
func genFn0(p0 uint8) int16 {
    var v6 int64 = 9223372036854775807
    if false {
        var v7 uint16 = 587
        v6 = v6
        var v8 uint64 = 0
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 uint8 = 2
    v4 := []int{-2147483648}
    v5 := make(map[uint32]int)
    switch "hello" {
    case "":
        v3 = (^(v3))
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9)
        v3 = v3
        if (0.0 >= 1.5) {
            if true {
                v4 = append(v4, -1)
            }
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 float32 = ((-2.25 / (float32(0.0) - -2.25)) / 1.0)
                var v12 uint = 255
                v3 -= (v3 >> ((v12) & 63))
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 241, F1: 9223372036854775807, F2: 721}
            _ = v13
        }
        _ = fi9
    }
    var v14 float64 = 1.5
    v15 := GenS1{F0: 2, F1: 9223372036854775807, F2: 899}
    var v16 uint = 4294967295
    fmt.Println(v3, v3)
    var v17 int8 = -128
    var v18 int32 = 1
    for fi19 := 0; fi19 < 1; fi19++ {
        if ((-(fi19)) < fi19) {
            for fi20 := 0; fi20 < 13; fi20++ {
                fmt.Println(v18, v3)
                _ = fi20
            }
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v3 += (^(v3))
            _ = len(v2)
            if ((v16 & v16) >= 0) {
                var v22 int8 = int8(v16)
                var v23 float32 = 3.14159
                switch (fi21 * fi21) {
                case 65:
                    v23 = -2.25
                default:
                    v23 += 0.0
                }
                _ = v22
                _ = v23
            }
            _ = fi21
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
