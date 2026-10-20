package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint, p1 uint32) int16 {
    v4 := []int{438, 32767, 32767}
    if len(v4) > 0 {
        v4[0] = -963
    }
    _ = p0
    _ = p1
    _ = v4
    return 32767
}
func genFn1() uint64 {
    v5 := []int{32768, 527}
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 42, F1: 4294967295}
        if len(v5) > 0 {
            v5[0] = fi6
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 42
    var v9 uint8 = 128
    _ = v5
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int8 = -1
    var v2 int32 = -1
    var v3 int64 = 541
    v3 += (-(v3))
    v10 := []int{}
    var v11 uint = 1
    fmt.Println(v2, v3)
    for fi12 := 0; fi12 < 7; fi12++ {
        v3 -= -9223372036854775808
        _ = fi12
    }
    v1 = 1
    v3 += (-(v3))
    fmt.Println(v3, v3)
    var v13 uint64 = 1
    if (!(!false)) {
        for fi14 := 0; fi14 < 19; fi14++ {
            v2 += 65
            _ = fi14
        }
        var v15 float64 = (3.14159 / 1.0)
        v3 += (v3 ^ 1)
        _ = v15
    }
    var v16 uint32 = 847
    if len(v10) > 0 {
        v10[0] = -310
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        switch 9223372036854775807 {
        case -999:
            var v18 uint16 = 42
            _ = v18
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v16
    _ = v2
    _ = v3
}
