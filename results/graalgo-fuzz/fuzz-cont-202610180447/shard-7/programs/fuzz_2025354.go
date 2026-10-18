package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8, p1 uint8, p2 bool) int16 {
    if (42 != 0) {
        v7 := GenS1{F0: 0}
        var v8 uint16 = 32767
        v9 := GenS2{F0: -935, F1: 1}
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(fi11, p1)
            _ = fi11
        }
        _ = fi10
    }
    switch 0 {
    case 1:
        fmt.Println(p2, p2, p0)
    case 127:
        var v12 uint16 = 42
        _ = v12
    default:
        if (!p2) {
            p1 = (p1 * (p0 ^ 127))
            switch 194 {
            case 0:
                p0 -= p1
            }
        } else {
            fmt.Println(p1, p2, p0)
            v13 := []int{9223372036854775807, 42, 60}
            fmt.Println(p1, p2, p0)
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 float32 = 1e10
            if p2 {
                if false {
                    var v17 int = (fi14 & fi15)
                    var v18 float32 = -2.25
                    _ = v17
                    _ = v18
                } else {
                    var v19 int8 = 76
                    var v20 int32 = -2147483648
                    p0 -= 2
                    _ = v19
                    _ = v20
                }
            } else {
                var v21 uint = 4294967295
                var v22 int16 = 42
                v16 = v16
                _ = v21
                _ = v22
            }
            var v23 uint = 1
            _ = fi15
            _ = v16
            _ = v23
        }
        fmt.Println(p1, fi14)
        var v24 uint8 = p0
        _ = fi14
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    var v1 float32 = (1.5 + (0.1 * 1e10))
    var v2 uint32 = 4294967295
    var v3 uint8 = 0
    var v4 uint16 = 42
    var v5 uint64 = 297
    var v6 int16 = 0
    v2 -= v2
    v25 := GenS3{F0: 1, F1: v2}
    v26 := []int{976}
    var v27 int32 = 1
    v28 := GenS2{F0: 717, F1: 65535}
    v6 += genFn0(42, v3, (!(1e10 < 3.14159)))
    v29 := GenS4{F0: -1}
    if (v27 <= 2147483647) {
        v27 = (v27 >> ((24) & 63))
        v5 = 18446744073709551615
    } else {
        var v30 int16 = v6
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
