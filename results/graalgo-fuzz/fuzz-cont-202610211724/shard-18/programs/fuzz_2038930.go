package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
type GenS4 struct { F0 bool; F1 int32; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{2147483647, 47, 42}
    var v3 uint16 = 441
    v4 := []int{42, -973, -2147483648, -9223372036854775808}
    var v5 int8 = -115
    var v6 uint64 = 566
    v7 := GenS1{F0: 1, F1: v5}
    var v8 int32 = 394
    v1[848] = 32767
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(v8)
        _ = fi9
    }
    v6 -= v6
    if false {
        var v10 float32 = float32(((1e10 - 0.0) - float64(-2.25)))
        _ = v10
    } else {
        _ = len(v1)
        _ = len(v1)
        v6 -= v6
    }
    var v11 bool = (true && true)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v12 := GenS1{F0: 32768, F1: (v5 << ((255) & 63))}
    var v13 float64 = (-2.25 / 1.0)
    if v11 {
        v3 = (v3 ^ uint16(v6))
        switch "foo bar" {
        case "foo bar":
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                delete(v1, 4294967295)
            }
        case "a":
            switch 42 {
            case -724:
                var v14 uint = 18446744073709551615
                _ = v14
            case 128:
                var v15 uint = 4294967295
                _ = v15
            default:
                if false {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        if (v13 >= -2.25) {
                            if v11 {
                                v6 = v6
                            } else {
                                v5 = v5
                                v3 += 13
                                var v17 uint8 = 219
                                _ = v17
                            }
                            v18 := GenS2{F0: (v8 % (((v8 & v8)) | 1))}
                            switch (("hello" + "a") + (("\n\t" + "") + ((("foo bar" + "") + "foo bar") + "hello"))) {
                            case "a":
                                var v19 uint32 = 2
                                _ = v19
                            default:
                                v13 = v13
                            }
                            _ = v18
                        } else {
                            v20 := GenS2{F0: v8}
                            v8 = (v8 ^ -2147483648)
                            v21 := v1[2147483648]
                            _ = v20
                            _ = v21
                        }
                        if ((^(v6)) > (v6 << ((745) & 63))) {
                            v11 = v11
                            fmt.Println(v6, v11)
                            var v22 uint64 = (v6 / ((652) | 1))
                            _ = v22
                        } else {
                            v23 := GenS3{F0: 0, F1: (v8 | 2147483647), F2: 2}
                            _ = v23
                        }
                        _ = fi16
                    }
                    if (uint8(v3) != 2) {
                        v13 = v13
                        fmt.Println(v13, v6)
                    }
                    var v24 uint8 = 221
                    _ = v24
                }
            }
        default:
            v25 := GenS1{F0: 127, F1: v5}
            _ = v25
        }
    }
    v26 := GenS4{F0: ((v3 + (v3 << ((42) & 63))) >= (^(v3))), F1: v8, F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
