package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int8; F2 uint64 }
func genFn0(p0 uint32, p1 uint, p2 uint16) int32 {
    var v7 uint = p1
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = 16
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(p1, p0, fi10)
            if (false || (!((!false) || true))) {
                v7 = v7
                v11 := GenS1{F0: (v9 - v9)}
                _ = v11
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 0
    v3 := []uint32{65535, 2147483648, 2147483647}
    var v4 int8 = v2
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    if (((-2.25 * 3.14159) * (((float32(1.5) + 1e10) / float32(-2.25)) + float32(((1.5 * (3.14159 - 0.0)) + 0.1)))) >= 1.5) {
        if (0 > 2) {
            var v12 float32 = (float32((-2.25 - 0.1)) / (0.1 / (float32(-2.25) / 1.0)))
            if (3.14159 != v12) {
                if true {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v4 -= v4
                        var v14 float64 = 3.14159
                        _ = fi13
                        _ = v14
                    }
                    var v15 int16 = 1000
                    _ = v15
                }
                v16 := GenS1{F0: 0}
                _ = v16
            } else {
                var v17 int64 = 42
                _ = v17
            }
            v12 = v12
            _ = v12
        } else {
            v2 = -28
        }
        switch -225 {
        case 32768:
            var v18 uint = 4294967295
            _ = v18
        case -2147483648:
            for fi19 := 0; fi19 < 17; fi19++ {
                delete(v5, "foo bar")
                var v20 int64 = 1
                var v21 bool = (int8(fi19) <= v2)
                _ = fi19
                _ = v20
                _ = v21
            }
        case 2147483647:
            delete(v5, ("\n\t" + "a"))
        default:
            fmt.Println(v4, v4)
        }
    }
    fmt.Println(v2, v4)
    fmt.Println(v4, v2)
    if false {
        var v22 uint8 = 148
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS1{F0: v22}
            var v25 int32 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: v22}
        _ = v22
        _ = v26
    }
    v2 += (v4 + -128)
    fmt.Println(v2)
    v2 = v2
    for fi27 := 0; fi27 < 6; fi27++ {
        if (42 < (-(fi27))) {
            v28 := GenS2{F0: int(v4)}
            var v29 float64 = 0.0
            v2 = (v2 * (v4 | v2))
            _ = v28
            _ = v29
        }
        _ = fi27
    }
    v4 -= v4
    delete(v1, "hello")
    v30 := GenS3{F0: "a", F1: v4, F2: 42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
