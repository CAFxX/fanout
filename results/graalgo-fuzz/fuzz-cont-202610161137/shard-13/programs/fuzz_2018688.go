package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0(p0 int32, p1 int8) uint16 {
    v5 := []int{-596, 32768}
    if (int64(p1) >= -1) {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint64 = 42
            if (false && true) {
                var v8 uint32 = 333
                var v9 int64 = int64(v8)
                _ = v8
                _ = v9
            } else {
                _ = len(v5)
                v7 -= 42
                var v10 int64 = 247
                _ = v10
            }
            _ = fi6
            _ = v7
        }
        if false {
            for fi11 := 0; fi11 < 2; fi11++ {
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 int = fi12
                    v13 = v13
                    _ = fi12
                    _ = v13
                }
                if false {
                    if ((p1 * (p1 * (p1 % ((-123) | 1)))) == p1) {
                        if true {
                            p0 -= (p0 / ((66) | 1))
                            p1 = 100
                        }
                    }
                    p0 -= -583
                } else {
                    p0 += 2147483647
                    p1 = -128
                }
                _ = fi11
            }
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 263
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]uint32)
    var v3 string = "foo bar"
    v4 := make(map[uint32]bool)
    var v14 uint64 = 18446744073709551615
    switch ("foo bar" + ("" + ((v3 + "\n\t") + v3))) {
    case "\n\t":
        var v15 float64 = float64((3.14159 + ((1e10 - 1.5) * ((3.14159 + (0.0 * -2.25)) - 3.14159))))
        _ = v15
    case "a":
        if (true && false) {
            v3 = v3
            v14 = 4294967295
        } else {
            v16 := GenS1{F0: (v14 ^ (^(v14)))}
            _ = v16
        }
    case "hello":
        _ = len(v4)
    }
    if (42 <= 42) {
        v1 = v1
    }
    v3 = "hello"
    v1 += v1
    v1 += (v1 / ((-361) | 1))
    if (843 == 494) {
        fmt.Println(v1, v14)
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 int64 = 42
            var v19 float64 = (0.1 - 3.14159)
            v18 = v18
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := GenS2{F0: -60, F1: 833}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS2{F0: -1, F1: 42}
            _ = fi21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        v3 = v3
        _ = fi23
    }
    var v24 float32 = 0.0
    v14 = 0
    v25 := GenS1{F0: (v14 | (^(v14)))}
    fmt.Println(v3, v14)
    v26 := []int{-968, -9223372036854775808, -9223372036854775808}
    for fi27 := 0; fi27 < 8; fi27++ {
        fmt.Println(v24, fi27)
        _ = fi27
    }
    switch int(v14) {
    case 13:
        v28 := GenS2{F0: 0, F1: 270}
        _ = v28
    default:
        fmt.Println(v3, v14, v14)
    }
    var v29 float32 = (v24 / ((v24 / 1e10) + ((((((((1e10 + (1.5 + (((v24 * (v24 / v24)) * float32((((-2.25 - 0.1) * 1.5) + float64(v24)))) / v24))) / 1.0) / v24) + 1.5) * v24) / v24) * 1.5) / -2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
