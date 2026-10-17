package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 int64, p1 int64, p2 int64) uint64 {
    var v7 bool = (!("a" < ("a" + "a")))
    switch 565 {
    case 127:
        if (!v7) {
            if (!false) {
                fmt.Println(p1)
                v7 = (!(false || ("\n\t" < "\n\t")))
            } else {
                v8 := []int{4294967295}
                var v9 float64 = (-2.25 + float64(1e10))
                _ = v8
                _ = v9
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 4294967295
}

func main() {
    var v1 int8 = 127
    v2 := []uint32{2147483647}
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    var v5 uint8 = 0
    var v6 float64 = float64(((0.1 * float32(0.0)) * ((float32((((1e10 - -2.25) * -2.25) - (1e10 - ((float64(-2.25) - 0.0) - 1.5)))) + (0.1 * 0.0)) * 3.14159)))
    delete(v3, ("a" + ("" + ("foo bar" + "a"))))
    v10 := GenS1{F0: false}
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int16 = 1000
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(fi13, v6, v5)
        var v14 uint32 = 1
        _ = fi13
        _ = v14
    }
    fmt.Println(v6, v1, v1)
    v6 += (((-2.25 + (0.0 + v6)) * v6) + v6)
    var v15 uint = 725
    v15 += 65535
    v5 = 82
    v5 += v5
    fmt.Println(v1, v15)
    if (float32(1.5) != float32(float64(float32(v6)))) {
        v16 := GenS1{F0: true}
        for fi17 := 0; fi17 < 19; fi17++ {
            v5 -= (v5 + v5)
            v18 := GenS2{F0: v1}
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        var v19 int8 = v1
        _ = v19
    }
    v20 := GenS3{F0: 47, F1: 0}
    v15 = (v15 * 195)
    fmt.Println(v6)
    var v21 int8 = -128
    v22 := GenS4{F0: (^(v5)), F1: 429, F2: -9223372036854775808}
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v5)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
