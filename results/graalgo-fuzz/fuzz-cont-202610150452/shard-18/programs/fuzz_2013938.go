package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int16; F2 int64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint8 }
func genFn0(p0 string, p1 uint16) int16 {
    v5 := []int{-1, 42, 692, -1}
    if (!true) {
        var v6 uint16 = (p1 ^ 255)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v7 float64 = 0.1
        _ = v6
        _ = v7
    }
    var v8 uint8 = 127
    v9 := GenS1{F0: -123}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    _ = v9
    return 582
}

func main() {
    v1 := []int{40, -9223372036854775808, 32767, -949}
    v2 := []uint32{4294967295, 4294967295, 66, 76}
    var v3 int = 127
    var v4 int16 = -1
    fmt.Println(v4, v3)
    var v10 string = "\n\t"
    var v11 uint = 0
    for fi12 := 0; fi12 < 14; fi12++ {
        _ = len(v2)
        _ = fi12
    }
    var v13 float32 = ((((1e10 / (1.5 - -2.25)) * 1.5) * (0.0 - 0.1)) - ((1.5 / 1e10) / (1.5 * (-2.25 - (-2.25 - 1.5)))))
    for fi14 := 0; fi14 < 20; fi14++ {
        if ((float64(float32(-2.25)) >= -2.25) || true) {
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint32 = 2147483647
                v17 := GenS2{F0: 0, F1: v4, F2: 1}
                switch (v3 | (^(v3))) {
                case 202:
                    v13 += (v13 / v13)
                default:
                    if (-2147483648 >= fi15) {
                        if (-731 != -1) {
                            if false {
                                fmt.Println(v11, v3)
                                _ = len(v2)
                                var v18 uint32 = (v16 / ((v16) | 1))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v18
                            } else {
                                v3 += v3
                                var v19 bool = (0 < -1)
                                _ = v19
                            }
                            v20 := GenS2{F0: -9223372036854775808, F1: v4, F2: 9223372036854775807}
                            v21 := GenS3{F0: 2147483648, F1: v4, F2: (v11 + (^(v11)))}
                            _ = v20
                            _ = v21
                        }
                        var v22 int16 = (v4 ^ int16(v11))
                        if (genFn0("foo bar", 1) == (v22 << ((v11) & 63))) {
                            fmt.Println(v22, v3)
                            v3 = v3
                            fmt.Println(v13)
                        } else {
                            var v23 uint16 = 570
                            _ = v23
                        }
                        _ = v22
                    } else {
                        v3 -= fi14
                    }
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi14
    }
    v11 += v11
    v13 += (v13 / v13)
    v13 += float32((1.5 * float64(v13)))
    var v24 uint32 = 42
    var v25 uint8 = 2
    v26 := GenS2{F0: 9223372036854775807, F1: v4, F2: 607}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v10 {
    case "hello":
        var v27 uint8 = v25
        _ = v27
    case "foo bar":
        var v28 int8 = -128
        _ = v28
    case "\n\t":
        v29 := GenS3{F0: v24, F1: (v4 ^ v4), F2: (v11 - v11)}
        _ = v29
    }
    v30 := GenS4{F0: 0, F1: 531}
    v31 := GenS5{F0: 1}
    if ("hello" > v10) {
        fmt.Println(v10, v25, v25)
    } else {
        if (((!(!false)) || true) && true) {
            v32 := GenS6{F0: 2, F1: (v25 << ((v11) & 63))}
            _ = v32
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
