package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 int64, p1 int32) int64 {
    v5 := []int{129, 128, 1, 672}
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            p1 = p1
            _ = fi7
        }
        _ = fi6
    }
    p1 = (p1 << ((2) & 63))
    if false {
        var v8 uint8 = 2
        _ = v8
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v9 uint8 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    var v4 uint32 = 185
    var v10 int = 32767
    var v11 float32 = float32(1.5)
    v12 := GenS1{F0: 304}
    v13 := GenS1{F0: 224}
    v11 += 0.1
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := []int{v10, 9223372036854775807, 9223372036854775807}
        if (4294967295 < 0) {
            v16 := GenS2{F0: "hello", F1: "foo bar"}
            v10 = (fi14 % ((v10) | 1))
            _ = len(v15)
            _ = v16
        } else {
            _ = len(v3)
            v11 -= (0.0 + v11)
        }
        var v17 uint32 = v4
        _ = fi14
        _ = v15
        _ = v17
    }
    fmt.Println(v10)
    fmt.Println(v11, v10)
    var v18 bool = false
    var v19 float64 = float64((v11 + 1e10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
