package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := []uint32{}
    v2 := []int{466, -2147483648, 32768}
    var v3 int16 = 0
    for fi4 := 0; fi4 < 16; fi4++ {
        v3 = (v3 << ((18446744073709551615) & 63))
        for fi5 := 0; fi5 < 20; fi5++ {
            fmt.Println(fi5, v3, v3)
            var v6 uint = 65535
            v7 := GenS1{F0: uint32(v6), F1: 0}
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    var v8 float32 = (float32(0.1) + -2.25)
    fmt.Println(v3, v8)
    var v9 float64 = 1.5
    var v10 uint16 = 42
    var v11 float32 = (-2.25 - float32(float64(v8)))
    switch "a" {
    case "a":
        if false {
            var v12 bool = true
            var v13 bool = v12
            v9 += v9
            _ = v12
            _ = v13
        }
    case "":
        v2 = append(v2, 849)
    }
    var v14 uint = 403
    v15 := GenS2{F0: 0}
    for fi16 := 0; fi16 < 20; fi16++ {
        if ((!true) && (35 != 2)) {
            var v17 int64 = -32
            v18 := GenS1{F0: 2, F1: (v10 & (v10 % ((v10) | 1)))}
            var v19 bool = (!(-24 > 100))
            _ = v17
            _ = v18
            _ = v19
        } else {
            v1 = append(v1, 0)
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(fi16)
            v21 := GenS3{F0: 1, F1: 4294967295}
            _ = fi20
            _ = v21
        }
        _ = fi16
    }
    if true {
        fmt.Println(v8, v3, v8)
        v1 = append(v1, 2147483647)
    }
    switch 4294967295 {
    case 4294967295:
        v22 := GenS1{F0: 931, F1: v10}
        _ = v22
    case 9223372036854775807:
        fmt.Println(v8, v14, v11)
    case 128:
        var v23 int16 = v3
        _ = v23
    default:
        if true {
            v14 += 212
        }
    }
    fmt.Println(v3)
    switch 32767 {
    case -9223372036854775808:
        for fi24 := 0; fi24 < 1; fi24++ {
            v3 += (v3 ^ 388)
            var v25 uint16 = 32767
            _ = fi24
            _ = v25
        }
    default:
        fmt.Println(v3, v14)
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := GenS3{F0: -2147483648, F1: (v14 | (v14 / ((v14) | 1)))}
        _ = fi26
        _ = v27
    }
    v11 -= ((v8 - ((v11 - (1e10 * ((0.1 - (v8 - v8)) + 1e10))) - v8)) + -2.25)
    v8 = ((float32((0.1 / 1.5)) * (0.1 - v8)) - (float32(((v9 + (v9 / (0.0 - 0.1))) * v9)) - (1e10 + 1.5)))
    v28 := GenS1{F0: 2147483648, F1: v10}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v8
    _ = v9
}
