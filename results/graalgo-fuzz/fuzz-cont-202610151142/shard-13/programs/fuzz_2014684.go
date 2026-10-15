package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 bool; F1 bool; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int16) int32 {
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 9223372036854775807, F1: 113, F2: 0}
        v9 := GenS2{F0: true, F1: (-1 != -158), F2: (fi7 & fi7)}
        p0 -= 1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS3{F0: (p0 << ((1) & 63))}
    fmt.Println(p0)
    _ = p0
    _ = v10
    return -2147483648
}
func genFn1(p0 int, p1 string, p2 uint16) uint16 {
    p2 = p2
    var v11 int8 = -86
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 2147483647
    v3 := make(map[uint32]uint32)
    var v4 string = "foo bar"
    var v5 uint32 = 0
    var v6 uint = 65535
    v12 := GenS1{F0: -255, F1: 109, F2: 18446744073709551615}
    var v13 int8 = 1
    if true {
        fmt.Println(v5, v6, v6)
        fmt.Println(v6)
    }
    var v14 uint64 = 2
    if (!true) {
        for fi15 := 0; fi15 < 18; fi15++ {
            fmt.Println(v14, v6)
            _ = fi15
        }
        fmt.Println(v4, v13)
        v16 := []int{}
        _ = v16
    }
    v2 = (v2 & -417)
    var v17 int32 = v2
    v14 += v14
    switch (v4 + "hello") {
    case "a":
        var v18 bool = (float64(1e10) < 1.5)
        _ = v18
    default:
        fmt.Println(v13, v2, v2)
    }
    if (65535 == 42) {
        v19 := v3[v5]
        if true {
            var v20 int32 = 42
            var v21 float64 = (0.1 + (-2.25 * float64(3.14159)))
            if (true && false) {
                if (true || (0 <= (v14 + v14))) {
                    var v22 bool = (-711 < 128)
                    _ = v22
                }
            } else {
                v13 = (^(v13))
                if true {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        fmt.Println(v14)
                        for fi24 := 0; fi24 < 19; fi24++ {
                            v25 := GenS3{F0: 0}
                            _ = fi24
                            _ = v25
                        }
                        v26 := GenS4{F0: ""}
                        _ = fi23
                        _ = v26
                    }
                } else {
                    fmt.Println(v2)
                    var v27 int32 = (^(v17))
                    _ = v27
                }
            }
            _ = v20
            _ = v21
        } else {
            var v28 float32 = ((((0.0 - 0.0) / float32(float64(1e10))) / 0.1) - (3.14159 + -2.25))
            v4 = v4
            delete(v1, v4)
            _ = v28
        }
        v29 := GenS3{F0: -265}
        _ = v19
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
