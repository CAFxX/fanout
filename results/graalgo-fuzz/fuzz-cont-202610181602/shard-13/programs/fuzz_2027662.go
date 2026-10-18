package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
func genFn0(p0 uint16, p1 string) uint8 {
    var v7 uint8 = 128
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    v3 := []uint32{65535, 0, 2147483648, 0}
    v4 := make(map[string]int)
    var v5 int8 = 42
    var v6 int8 = v5
    for fi8 := 0; fi8 < 14; fi8++ {
        if false {
            v5 = (v5 & (v6 & v5))
        } else {
            v6 -= (v5 - int8(fi8))
        }
        _ = fi8
    }
    if false {
        v5 -= (v6 | v5)
        if (false && (2147483648 < 2147483648)) {
            var v9 uint8 = 127
            var v10 int32 = 2147483647
            _ = len(v4)
            _ = v10
            _ = v9
        } else {
            v6 -= 70
        }
        _ = len(v4)
    }
    if false {
        v6 = v6
        v11 := GenS1{F0: 128, F1: (-1 < 610)}
        var v12 int = 1
        _ = v11
        _ = v12
    } else {
        fmt.Println(v5, v6)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v13 := GenS1{F0: 179, F1: true}
    switch -664 {
    case -9223372036854775808:
        for fi14 := 0; fi14 < 3; fi14++ {
            if false {
                v6 += (v6 ^ 77)
                fmt.Println(v5, fi14, v6)
            } else {
                v5 += 127
            }
            v4["hello"] = fi14
            if true {
                v6 -= (v5 ^ (v5 & v6))
            } else {
                var v15 uint8 = 128
                var v16 int64 = int64(fi14)
                _ = v15
                _ = v16
            }
            _ = fi14
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 bool = ((false && true) && false)
        var v19 uint32 = 42
        var v20 int16 = 1
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint64 = 2
        _ = fi21
        _ = v22
    }
    v6 = v6
    switch 577 {
    case 9223372036854775807:
        var v23 float32 = (0.1 + (0.1 * float32(1e10)))
        _ = v23
    case -503:
        var v24 int64 = 1
        _ = v24
    }
    v5 -= -106
    var v25 int8 = (v5 & 42)
    if ((false && true) && (1 > 4294967295)) {
        if true {
            if true {
                _ = len(v4)
            } else {
                fmt.Println(v6)
            }
        }
        for fi26 := 0; fi26 < 6; fi26++ {
            fmt.Println(v25, fi26, v6)
            v27 := GenS1{F0: 104, F1: (((3.14159 - (float64(0.1) - float64(float32(float64(-2.25))))) >= 1.5) || (-32768 >= 1))}
            _ = fi26
            _ = v27
        }
    }
    if true {
        var v28 uint32 = 515
        v6 = (v25 & 100)
        _ = v28
    }
    v5 += v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
