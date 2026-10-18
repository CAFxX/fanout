package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
func genFn0(p0 string) int8 {
    fmt.Println(p0)
    p0 = "\n\t"
    var v6 float64 = 0.1
    var v7 string = p0
    var v8 uint8 = 1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -88
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 int = -9223372036854775808
    var v4 uint32 = 2
    v5 := make(map[string]string)
    switch v3 {
    case 0:
        var v9 int8 = 100
        _ = v9
    case 128:
        v10 := []int{-9223372036854775808, 9223372036854775807}
        _ = v10
    case 149:
        v3 -= (^(v3))
    }
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            switch "hello" {
            case "\n\t":
                for fi12 := 0; fi12 < 14; fi12++ {
                    if (-2147483648 > 0) {
                        v4 = 42
                    }
                    var v13 float32 = (0.1 / 0.1)
                    var v14 bool = ((v4 & (v4 << ((2) & 63))) < v4)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            default:
                if false {
                    fmt.Println(v4, v3)
                    v4 = (v4 ^ 2147483647)
                }
            }
            _ = fi11
        }
        fmt.Println(v3, v3)
    }
    v15 := []int{v3, 0, -783}
    var v16 uint16 = 44
    v15 = append(v15, (v3 | v3))
    for fi17 := 0; fi17 < 19; fi17++ {
        delete(v5, "")
        var v18 float32 = ((1e10 / 1.5) - (float32(0.1) + 1e10))
        _ = fi17
        _ = v18
    }
    switch (v3 * -160) {
    case -809:
        var v19 uint32 = v4
        _ = v19
    default:
        v20 := GenS1{F0: 9223372036854775807, F1: -713, F2: 4294967295}
        _ = v20
    }
    v4 -= v4
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int32 = 454
        if (-661 == -757) {
            var v23 int8 = int8(v22)
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        if (false || (false && true)) {
            var v25 bool = false
            _ = v25
        }
        _ = fi24
    }
    v4 += v4
    var v26 int = (v3 * v3)
    fmt.Println(v26)
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 uint32 = (^(v4))
        fmt.Println(v3)
        _ = fi27
        _ = v28
    }
    var v29 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
