package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint16; F1 int; F2 int16 }

func main() {
    var v1 string = ""
    var v2 uint64 = 292
    var v3 bool = (!((false || (1 <= 32767)) || false))
    var v4 int16 = -749
    v5 := GenS1{F0: 928}
    v6 := []int{-1, 9223372036854775807, -2147483648, 128}
    v7 := GenS2{F0: -43, F1: 207, F2: (^(v2))}
    for fi8 := 0; fi8 < 19; fi8++ {
        if v3 {
            v2 = v2
            switch ((("" + v1) + v1) + (v1 + (v1 + "\n\t"))) {
            case "a":
                var v9 uint = 817
                _ = v9
            }
            v10 := GenS1{F0: 328}
            _ = v10
        }
        _ = fi8
    }
    var v11 int32 = -2147483648
    var v12 float64 = (0.1 - float64(((float32(1.5) + 1e10) / float32(3.14159))))
    v13 := GenS3{F0: 96, F1: (-(v11))}
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 bool = false
        _ = fi14
        _ = v15
    }
    v4 -= int16(v2)
    v2 -= v2
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint16 = 32767
        _ = fi16
        _ = v17
    }
    switch 32768 {
    case -749:
        fmt.Println(v11, v11, v12)
    case 1:
        v1 = ("" + "")
    default:
        var v18 int8 = -118
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        _ = len(v6)
        _ = fi19
    }
    v6 = append(v6, 1)
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int32 = (v11 - (v11 & v11))
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v23 := GenS4{F0: 65535, F1: 42, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
