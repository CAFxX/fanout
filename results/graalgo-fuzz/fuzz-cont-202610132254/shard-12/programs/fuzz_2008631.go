package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 uint16) uint64 {
    for fi4 := 0; fi4 < 17; fi4++ {
        p0 -= p0
        var v5 int32 = int32(p0)
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: 710}
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1() bool {
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 uint32 = 42
    _ = v8
    return false
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = 32767
    var v3 int16 = (v2 + (v2 + v2))
    v9 := GenS1{F0: 9223372036854775807, F1: v2}
    v10 := GenS1{F0: int64(v1), F1: 331}
    if ((v2 >> ((0) & 63)) >= v2) {
        var v11 int = int(v2)
        if false {
            var v12 uint64 = 18446744073709551615
            v1 = (v1 - v1)
            v12 = genFn0(255)
            _ = v12
        } else {
            fmt.Println(v3)
        }
        v1 -= (v1 % (((v1 + (v1 - (^(v1))))) | 1))
        _ = v11
    }
    v13 := GenS1{F0: -9223372036854775808, F1: (v3 % (((v3 * (-(v2)))) | 1))}
    if true {
        for fi14 := 0; fi14 < 8; fi14++ {
            v3 += int16(fi14)
            _ = fi14
        }
        v15 := []int{669, 2147483647, 127, -386}
        _ = v15
    } else {
        var v16 float32 = -2.25
        var v17 float32 = float32(0.1)
        _ = v16
        _ = v17
    }
    var v18 int = 157
    fmt.Println(v2, v2)
    v18 -= 127
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
