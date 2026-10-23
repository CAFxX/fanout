package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int, p1 int16, p2 uint64) int32 {
    var v4 float64 = (((float64(1.5) * (float64(-2.25) * 1.5)) - (1.5 - ((1e10 - (0.0 / (0.1 / 1.5))) * 0.0))) * (((1e10 / ((1e10 / 1.5) - (float64(3.14159) - ((0.0 + ((0.1 - ((1.5 / 3.14159) * (1e10 - 0.1))) / 1.5)) - (-2.25 + 0.0))))) * 0.1) + 0.0))
    if true {
        var v5 int = (p0 * -842)
        if false {
            fmt.Println(v4)
            v4 -= float64(-2.25)
        } else {
            switch "foo bar" {
            case "hello":
                var v6 uint = 83
                _ = v6
            default:
                fmt.Println(v5, p0, p1)
            }
            v7 := []int{-1, v5, v5, p0}
            _ = v7
        }
        fmt.Println(p1, v4)
        _ = v5
    }
    v8 := []int{(-(p0)), (p0 >> ((1) & 63)), (p0 & p0)}
    for fi9 := 0; fi9 < 1; fi9++ {
        if (65535 >= 795) {
            for fi10 := 0; fi10 < 19; fi10++ {
                fmt.Println(p2, p2, p1)
                p0 = (p0 | (p0 << ((2) & 63)))
                if len(v8) > 0 {
                    v8[0] = fi9
                }
                var v11 uint32 = 492
                p1 -= p1
                _ = fi10
                _ = v11
            }
        } else {
            v12 := GenS1{F0: 1, F1: 771, F2: 232}
            var v13 uint64 = p2
            v8 = append(v8, 9223372036854775807)
            _ = v12
            _ = v13
        }
        var v14 int64 = -1
        _ = fi9
        _ = v14
    }
    if (true || false) {
        if (!(1e10 >= (-2.25 - (0.1 * (float32((v4 * v4)) * (0.1 + 0.0)))))) {
            if (p0 == p0) {
                fmt.Println(p1, v4, p1)
            } else {
                p0 = (p0 % ((p0) | 1))
                fmt.Println(p0)
                var v15 int = (-(p0))
                _ = v15
            }
        } else {
            fmt.Println(p0)
            var v16 int16 = 1000
            _ = v16
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 42
}
func genFn1(p0 uint16, p1 int16, p2 uint64) uint {
    switch 127 {
    case 4294967295:
        p2 -= p2
    case 42:
        fmt.Println(p0)
    }
    var v17 int16 = -32768
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS2{F0: p0, F1: 42}
        v20 := GenS1{F0: 786, F1: 186, F2: 9223372036854775807}
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 565
}

func main() {
    var v1 int32 = 0
    var v2 bool = (312 != -9223372036854775808)
    var v3 int64 = 1
    if true {
        v21 := GenS2{F0: 407, F1: 32767}
        fmt.Println(v2)
        _ = v21
    } else {
        var v22 int16 = -1
        for fi23 := 0; fi23 < 16; fi23++ {
            fmt.Println(v3)
            switch fi23 {
            case 0:
                switch fi23 {
                case 32767:
                    v24 := GenS1{F0: 42, F1: 4294967295, F2: (^(v3))}
                    _ = v24
                }
            case 42:
                v25 := []int{(^(fi23)), (fi23 % ((fi23) | 1)), fi23}
                _ = v25
            case -2147483648:
                var v26 bool = (2 >= genFn1(65535, 0, 94))
                _ = v26
            default:
                var v27 uint32 = 4294967295
                _ = v27
            }
            _ = fi23
        }
        fmt.Println(v3)
        _ = v22
    }
    v1 = -422
    v28 := []int{32767, -1, 32767, 868}
    v2 = (4294967295 > 79)
    var v29 bool = true
    for fi30 := 0; fi30 < 9; fi30++ {
        v3 -= 9223372036854775807
        _ = fi30
    }
    v31 := GenS1{F0: -377, F1: 42, F2: (v3 >> ((445) & 63))}
    v1 += v1
    v29 = (65 > genFn1(303, -1, 194))
    var v32 float32 = (0.1 * -2.25)
    v28 = append(v28, -2147483648)
    v33 := GenS1{F0: 0, F1: 18446744073709551615, F2: v3}
    v34 := GenS1{F0: 42, F1: 42, F2: 124}
    var v35 uint32 = 354
    for fi36 := 0; fi36 < 17; fi36++ {
        v37 := GenS2{F0: 32767, F1: 1}
        v38 := GenS3{F0: 127, F1: (v35 & 2147483647)}
        _ = fi36
        _ = v37
        _ = v38
    }
    if ((v35 * (v35 & 204)) >= v35) {
        fmt.Println(v3, v2, v29)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
