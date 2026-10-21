package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16 }
func genFn0() string {
    var v6 string = "\n\t"
    v7 := GenS1{F0: 18446744073709551615, F1: -470}
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 float32 = float32(float64((0.1 + 1e10)))
    var v2 int = 42
    v3 := []string{("a" + "foo bar"), (("foo bar" + (((("foo bar" + ("foo bar" + ("foo bar" + "a"))) + "foo bar") + ("a" + (("" + ("a" + "hello")) + "foo bar"))) + "a")) + "a")}
    v4 := []int{v2, (^(v2)), v2}
    var v5 bool = false
    var v8 uint32 = 1
    if v5 {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint32 = v8
            var v11 float64 = float64((v1 * (3.14159 + v1)))
            v12 := GenS2{F0: 388, F1: 1, F2: 42}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    } else {
        v5 = v5
        fmt.Println(v5)
    }
    if ((!v5) && v5) {
        v2 = v2
        _ = len(v3)
    }
    v13 := GenS2{F0: 150, F1: v8, F2: 0}
    switch (v2 % (((v2 << ((84) & 63))) | 1)) {
    case 613:
        v8 += (^(v8))
    case 127:
        v5 = (!true)
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        if (420 != 415) {
            if (false && (!v5)) {
                v15 := GenS3{F0: 2}
                _ = v15
            }
        } else {
            fmt.Println(fi14, fi14, fi14)
            var v16 float64 = float64(v1)
            _ = v16
        }
        fmt.Println(v5, fi14, v2)
        if ((v2 & fi14) > 32768) {
            for fi17 := 0; fi17 < 6; fi17++ {
                _ = len(v3)
                _ = fi17
            }
        }
        _ = fi14
    }
    if (365 > 0) {
        switch (v2 + v2) {
        case -2147483648:
            var v18 uint32 = 713
            _ = v18
        case 4294967295:
            v1 += v1
        default:
            var v19 uint64 = 42
            _ = v19
        }
    } else {
        v20 := GenS2{F0: 42, F1: 65535, F2: 51}
        fmt.Println(v2, v5)
        var v21 int64 = -882
        _ = v20
        _ = v21
    }
    v1 += v1
    v22 := GenS1{F0: 18446744073709551615, F1: 42}
    v23 := GenS2{F0: 2, F1: v8, F2: 42}
    v5 = (131 == -2147483648)
    v5 = ((!v5) && true)
    var v24 float32 = (((v1 + ((1.5 / v1) - v1)) - 1.5) - float32((1.5 * 1e10)))
    v5 = (((-339 <= 999) || false) && (false && v5))
    var v25 uint = 1
    fmt.Println(v8, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
