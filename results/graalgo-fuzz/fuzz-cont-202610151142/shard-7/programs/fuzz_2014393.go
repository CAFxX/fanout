package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint32) uint {
    v6 := []int{}
    if len(v6) > 0 {
        v6[0] = 69
    }
    p0 -= (p0 % ((28) | 1))
    _ = len(v6)
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 uint = 65535
    var v2 float32 = (((float32(1e10) / 1e10) - float32(1.5)) + -2.25)
    var v3 float64 = 1e10
    var v4 float32 = v2
    var v5 uint8 = 1
    switch -2147483648 {
    case 42:
        var v7 uint64 = 0
        _ = v7
    case 2147483647:
        var v8 string = ("" + ("" + "\n\t"))
        _ = v8
    case 32767:
        var v9 int64 = -1
        _ = v9
    }
    v4 = ((v2 - (float32((0.0 / 1.5)) + (v4 + v2))) + v2)
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: -7}
        v12 := GenS1{F0: 940}
        var v13 uint16 = 2
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v1 = (v1 / ((uint(v5)) | 1))
    fmt.Println(v2, v1, v3)
    if (203 > v1) {
        v2 += (((3.14159 / v2) * float32(v3)) * 1.5)
        v5 = (^(v5))
        if (float64(v2) <= 3.14159) {
            fmt.Println(v5)
            fmt.Println(v5)
        }
    }
    var v14 uint32 = 726
    switch "a" {
    case "\n\t":
        var v15 int = 371
        _ = v15
    default:
        v16 := GenS2{F0: (v5 * (v5 >> ((2) & 63))), F1: -32768}
        _ = v16
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
