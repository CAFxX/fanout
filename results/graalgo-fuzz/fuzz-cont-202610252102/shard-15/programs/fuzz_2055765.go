package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
func genFn0(p0 uint64, p1 int32, p2 string) uint8 {
    fmt.Println(p1, p1, p2)
    var v7 int = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 uint8) uint64 {
    p0 = true
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 10; fi9++ {
        p1 = 148
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return v8
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 3.14159
    var v3 int16 = -197
    v4 := []string{}
    v5 := make(map[string]string)
    var v6 float32 = v2
    v10 := GenS1{F0: -48}
    var v11 int8 = 42
    switch "a" {
    case "\n\t":
        var v12 float32 = float32(0.1)
        _ = v12
    default:
        v11 = (-(v11))
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float32 = v6
        _ = fi13
        _ = v14
    }
    v3 -= (v3 % ((v3) | 1))
    if true {
        v15 := GenS2{F0: 42, F1: false, F2: 599}
        v2 -= (0.1 * v2)
        var v16 int16 = (v3 ^ v3)
        _ = v15
        _ = v16
    } else {
        v11 = (v11 / ((108) | 1))
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    }
    var v17 float64 = 3.14159
    switch "foo bar" {
    case "a":
        v18 := v1[753]
        _ = v18
    case "\n\t":
        var v19 uint = 4294967295
        _ = v19
    case "":
        var v20 int32 = 42
        _ = v20
    default:
        for fi21 := 0; fi21 < 8; fi21++ {
            if false {
                var v22 int64 = -9223372036854775808
                _ = v22
            }
            v4 = append(v4, "\n\t")
            _ = fi21
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v1[-1] = 744
        for fi24 := 0; fi24 < 4; fi24++ {
            switch "" {
            case "a":
                v25 := GenS1{F0: v11}
                _ = v25
            case "":
                if len(v4) > 0 {
                    v4[0] = (("\n\t" + "a") + "hello")
                }
            default:
                v2 = ((1.5 * v2) * (v2 / v6))
            }
            _ = fi24
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS2{F0: 9223372036854775807, F1: (958 == 576), F2: 1}
        v3 += -320
        for fi28 := 0; fi28 < 5; fi28++ {
            v3 = (v3 * (v3 % ((-32768) | 1)))
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    if true {
        v29 := GenS1{F0: 127}
        var v30 int8 = -128
        _ = v29
        _ = v30
    }
    fmt.Println(v11)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
