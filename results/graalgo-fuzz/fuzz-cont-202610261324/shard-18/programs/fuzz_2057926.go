package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
func genFn0(p0 int) uint16 {
    switch p0 {
    case 42:
        if (p0 >= p0) {
            p0 -= (p0 | (p0 ^ (p0 + -440)))
            if (!((42 != 586) && false)) {
                var v6 uint = 228
                v7 := []int{-1, (-(p0)), p0, p0}
                p0 += -1
                _ = v6
                _ = v7
            } else {
                var v8 uint64 = 10
                var v9 string = (("foo bar" + "foo bar") + "foo bar")
                _ = v8
                _ = v9
            }
        } else {
            p0 += -2147483648
        }
    case 32767:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 float32 = 3.14159
            _ = fi10
            _ = v11
        }
    case 127:
        v12 := []int{p0, p0, p0}
        _ = v12
    }
    var v13 int64 = -9223372036854775808
    _ = p0
    _ = v13
    return 42
}
func genFn1(p0 uint) int64 {
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 float64 = float64(1.5)
        _ = fi14
        _ = v15
    }
    v16 := []int{}
    fmt.Println(p0)
    var v17 float64 = -2.25
    _ = p0
    _ = v16
    _ = v17
    return 391
}

func main() {
    v1 := []int{}
    var v2 string = "a"
    var v3 uint8 = 248
    v4 := []uint32{2}
    var v5 uint8 = v3
    v18 := GenS1{F0: "a", F1: 0, F2: 693}
    v3 = 128
    v19 := GenS2{F0: 2147483648, F1: 420}
    var v20 int = 32767
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: "a", F1: -839, F2: 2147483647}
        var v23 float32 = (3.14159 * 3.14159)
        _ = fi21
        _ = v22
        _ = v23
    }
    switch -1 {
    case 32767:
        v5 = v5
    case -962:
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 float32 = ((((1.5 + (-2.25 / ((((0.1 * 0.0) - 1.5) * (((-2.25 * (((-2.25 * 3.14159) * 3.14159) - ((1.5 * 0.0) - (float32(((3.14159 * -2.25) + (0.1 - 1.5))) + 1.5)))) - (3.14159 / (-2.25 / 1e10))) - 1.5)) - 0.1))) / 1.0) / 1.5) / 3.14159)
            v26 := GenS2{F0: 809, F1: 2}
            _ = fi24
            _ = v25
            _ = v26
        }
    case -400:
        if ((v20 ^ (v20 | (v20 ^ (v20 & (v20 & (v20 >> ((18446744073709551615) & 63))))))) <= (^(v20))) {
            v27 := GenS3{F0: 2, F1: 0, F2: (^(v5))}
            if false {
                v3 += v5
                var v28 uint32 = 4294967295
                _ = v28
            }
            fmt.Println(v2, v20)
            _ = v27
        } else {
            switch -349 {
            case -1:
                var v29 uint8 = v5
                _ = v29
            case 110:
                v30 := GenS2{F0: 563, F1: 2}
                _ = v30
            }
            v31 := GenS2{F0: 2147483648, F1: 2147483647}
            v1 = append(v1, v20)
            _ = v31
        }
    default:
        var v32 int16 = 1
        _ = v32
    }
    var v33 int32 = 42
    fmt.Println(v20)
    v3 = 36
    if (v33 < 281) {
        var v34 int16 = 32767
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v34
    }
    v3 = (v3 | (v3 / (((v5 >> ((4294967295) & 63))) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
