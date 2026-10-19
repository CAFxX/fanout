package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int32, p1 uint, p2 uint64) int32 {
    p0 = p0
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: int8(fi7), F1: int64(fi7)}
        _ = fi7
        _ = v8
    }
    p1 = (p1 + (p1 % ((965) | 1)))
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return p0
}
func genFn1(p0 uint8, p1 int64, p2 int32) int {
    if true {
        v9 := GenS1{F0: 42, F1: p1}
        p0 -= p0
        if (false || false) {
            fmt.Println(p0, p2, p1)
            var v10 int8 = 1
            fmt.Println(p0)
            _ = v10
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                p0 -= (p0 ^ p0)
                _ = fi11
            }
        }
        _ = v9
    } else {
        p0 -= 222
        fmt.Println(p1)
        p1 = p1
    }
    p0 -= p0
    v12 := GenS2{F0: 1, F1: (p1 | 0)}
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int32 = (-(p2))
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            p2 = p2
            p0 += p0
            _ = fi16
        }
        fmt.Println(p0)
        _ = fi13
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    var v1 float64 = -2.25
    var v2 int8 = 1
    v3 := make(map[string]string)
    var v4 string = "\n\t"
    var v5 int16 = -342
    var v6 float64 = v1
    v17 := GenS2{F0: 1, F1: 42}
    v18 := GenS1{F0: (^(v2)), F1: 42}
    fmt.Println(v6)
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int8 = v2
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int16 = (v5 * v5)
            switch fi21 {
            case 2147483647:
                var v23 uint64 = 0
                _ = v23
            default:
                if true {
                    var v24 int8 = (^(v2))
                    _ = v24
                }
            }
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v4, v6)
    v6 = (v1 * 0.1)
    v1 = 3.14159
    var v25 int32 = 1
    v26 := GenS2{F0: 749, F1: 0}
    v25 += (v25 + (v25 / ((v25) | 1)))
    if true {
        v27 := GenS1{F0: -125, F1: -9223372036854775808}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 16; fi28++ {
            var v29 uint8 = 128
            v30 := GenS2{F0: 0, F1: 42}
            _ = fi28
            _ = v29
            _ = v30
        }
        v31 := GenS2{F0: 825, F1: -1}
        _ = v31
    }
    var v32 uint64 = 68
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
