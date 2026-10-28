package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint; F1 int8 }

func main() {
    var v1 int16 = 32767
    v2 := []uint32{42, 2, 0}
    var v3 int = -1
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 1}
    v1 = v1
    if (!false) {
        switch (v3 % ((0) | 1)) {
        case 1:
            if false {
                var v6 float64 = 1e10
                _ = v6
            } else {
                var v7 float64 = ((((0.0 + 1.5) / -2.25) / 1.0) / 1.0)
                if true {
                    fmt.Println(v3, v3, v3)
                    var v8 bool = (3.14159 != float32(((float64((3.14159 * 3.14159)) * 1.5) - v7)))
                    for fi9 := 0; fi9 < 16; fi9++ {
                        if (!false) {
                            switch 9223372036854775807 {
                            case 42:
                                v10 := GenS1{F0: 0}
                                _ = v10
                            }
                            var v11 uint32 = 4294967295
                            _ = v11
                        }
                        if (!((!true) && (1 != 2))) {
                            switch v3 {
                            case 127:
                                var v12 int = (v3 * 9223372036854775807)
                                _ = v12
                            default:
                                v1 = int16(fi9)
                            }
                        } else {
                            fmt.Println(v8, v7, v1)
                            switch (fi9 * v3) {
                            case 252:
                                var v13 float32 = (1.5 * 3.14159)
                                _ = v13
                            default:
                                if ((310 > 4294967295) && (957 > -469)) {
                                    var v14 uint8 = 1
                                    if (false || v8) {
                                        if (!((v8 || ((2 != 2) && (828 > 0))) && (v8 || v8))) {
                                            if (-2147483648 == -907) {
                                                v2 = append(v2, 65535)
                                            } else {
                                                v3 += (fi9 - (fi9 + -9223372036854775808))
                                                fmt.Println(v8)
                                            }
                                        }
                                        v15 := GenS2{F0: "foo bar", F1: true, F2: 1}
                                        _ = v15
                                    } else {
                                        if (v8 && v8) {
                                            v7 = v7
                                            v3 -= v3
                                            v16 := GenS1{F0: (v1 << ((65535) & 63))}
                                            _ = v16
                                        }
                                        v3 += fi9
                                    }
                                    _ = v14
                                }
                            }
                            v1 = 1
                        }
                        _ = fi9
                    }
                    _ = v8
                } else {
                    var v17 bool = (2147483648 <= 471)
                    v18 := GenS1{F0: (v1 & 292)}
                    v7 += v7
                    _ = v17
                    _ = v18
                }
                _ = v7
            }
        }
        v19 := GenS3{F0: 1, F1: -128}
        _ = v19
    } else {
        switch "foo bar" {
        case "\n\t":
            var v20 int = v3
            _ = v20
        case "a":
            if false {
                v21 := v4[1]
                fmt.Println(v21)
                _ = v21
            }
        default:
            if (true || false) {
                var v22 uint16 = 42
                _ = v22
            }
        }
        var v23 int = v3
        _ = v23
    }
    var v24 float64 = (float64((1.5 * (float32(0.0) / 0.1))) - (1e10 + float64(3.14159)))
    v24 -= v24
    _ = len(v2)
    var v25 float64 = v24
    if true {
        if (-2147483648 != 337) {
            var v26 uint8 = 182
            _ = v26
        }
        switch "\n\t" {
        case "hello":
            v1 += -1
        default:
            _ = len(v4)
        }
    }
    var v27 uint64 = 109
    if len(v2) > 0 {
        v2[0] = 696
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
