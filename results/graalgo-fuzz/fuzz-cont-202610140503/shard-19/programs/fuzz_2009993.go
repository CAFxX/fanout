package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 int32, p1 uint64) int8 {
    if false {
        for fi5 := 0; fi5 < 4; fi5++ {
            switch fi5 {
            case 539:
                var v6 uint = 4294967295
                _ = v6
            default:
                var v7 bool = false
                _ = v7
            }
            p0 = 610
            fmt.Println(p0, p0, p1)
            _ = fi5
        }
    } else {
        switch 128 {
        case 42:
            if (65535 < 42) {
                var v8 int32 = (p0 ^ -234)
                _ = v8
            } else {
                for fi9 := 0; fi9 < 8; fi9++ {
                    v10 := []int{(fi9 / (((fi9 % (((fi9 & 127)) | 1))) | 1)), 32767, 127, fi9}
                    v11 := GenS1{F0: fi9, F1: 42}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v12 := []int{127}
                v13 := GenS1{F0: -847, F1: 847}
                _ = v12
                _ = v13
            }
        case -9223372036854775808:
            fmt.Println(p1, p0)
        case -2147483648:
            fmt.Println(p0, p0)
        }
        p0 = (p0 | p0)
        if false {
            fmt.Println(p0, p0)
            p0 = p0
            v14 := []int{-2147483648}
            _ = v14
        }
    }
    p0 = (p0 & (p0 * p0))
    p1 = p1
    v15 := []int{32767}
    if (!false) {
        var v16 float64 = (1e10 * (1e10 * 1.5))
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v15
    return -118
}

func main() {
    v1 := []uint32{0, 178, 0, 2}
    var v2 float32 = ((3.14159 * 0.1) * 1e10)
    var v3 bool = true
    v4 := []string{""}
    if (2 >= 728) {
        var v17 uint32 = 446
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 uint32 = v17
            _ = fi18
            _ = v19
        }
        if (2 < 906) {
            v17 = v17
            if true {
                v2 += (v2 * float32(float64((float32((0.1 + 0.0)) * (v2 * 0.0)))))
            }
        }
        _ = v17
    }
    var v20 int = -9223372036854775808
    v3 = (128 >= (v20 | -968))
    v20 += (v20 + -730)
    var v21 uint = 2
    v3 = v3
    v3 = (0 <= 141)
    v4 = append(v4, "")
    v22 := GenS2{F0: v21, F1: (v21 >> (((v21 / (((v21 * (v21 * (^(v21))))) | 1))) & 63))}
    v3 = v3
    v21 = 2
    v20 += v20
    var v23 int16 = -356
    v4 = append(v4, "a")
    if true {
        for fi24 := 0; fi24 < 16; fi24++ {
            _ = len(v1)
            _ = fi24
        }
    } else {
        var v25 string = "foo bar"
        _ = v25
    }
    var v26 uint64 = 42
    var v27 bool = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
