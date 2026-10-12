package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 int; F1 string; F2 uint }
type GenS3 struct { F0 string; F1 bool; F2 int32 }

func main() {
    var v1 uint16 = 2
    var v2 float64 = (((3.14159 + 0.0) / 1e10) + ((1.5 / (3.14159 - 0.0)) + (float64((1.5 / (((float32(float64(1.5)) / -2.25) * (-2.25 / float32((((1.5 + 0.0) / 1.0) - 0.0)))) + (0.1 / (3.14159 + 0.0))))) - 3.14159)))
    var v3 int8 = 1
    var v4 float64 = v2
    var v5 uint32 = 423
    fmt.Println(v2, v5, v2)
    v6 := []int{}
    switch ("foo bar" + (("\n\t" + "\n\t") + "\n\t")) {
    case "foo bar":
        v1 = (^(v1))
    case "\n\t":
        var v7 bool = (!(32767 == 32767))
        _ = v7
    }
    v3 += v3
    v8 := GenS1{F0: 255, F1: 35, F2: v3}
    fmt.Println(v5, v5, v2)
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 string = "hello"
        _ = fi9
        _ = v10
    }
    if true {
        v2 -= float64(1e10)
        var v11 uint8 = 128
        v2 -= (v2 / v4)
        _ = v11
    }
    v2 -= v2
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            if (v5 > (v5 | 4294967295)) {
                fmt.Println(v2)
                if true {
                    var v14 uint64 = 398
                    switch "foo bar" {
                    case "":
                        var v15 int16 = 1000
                        _ = v15
                    default:
                        var v16 uint16 = v1
                        v17 := GenS1{F0: 1, F1: 0, F2: v3}
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                } else {
                    if false {
                        if (true || true) {
                            if false {
                                switch fi13 {
                                case 9223372036854775807:
                                    switch 127 {
                                    case -9223372036854775808:
                                        v18 := GenS1{F0: 65535, F1: 42, F2: v3}
                                        _ = v18
                                    }
                                case 32767:
                                    var v19 int32 = int32(v5)
                                    _ = v19
                                default:
                                    fmt.Println(fi12, v3, v5)
                                }
                                switch "a" {
                                case "hello":
                                    if len(v6) > 0 {
                                        v6[0] = (fi13 ^ 32767)
                                    }
                                case "":
                                    fmt.Println(v5, v5)
                                default:
                                    if ((v5 % (((v5 + v5)) | 1)) <= (v5 + v5)) {
                                        v20 := GenS2{F0: 127, F1: "hello", F2: 0}
                                        fmt.Println(v4)
                                        _ = v20
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            if (0 < -103) {
                                                fmt.Println(fi12, v3)
                                                var v21 string = "foo bar"
                                                v22 := GenS2{F0: fi12, F1: "\n\t", F2: 42}
                                                _ = v21
                                                _ = v22
                                            }
                                        case "hello":
                                            var v23 uint32 = v5
                                            _ = v23
                                        case "":
                                            fmt.Println(v1, v4, fi13)
                                        default:
                                            v3 = v3
                                        }
                                    }
                                }
                                var v24 int32 = 912
                                _ = v24
                            }
                            fmt.Println(fi12, v4, v4)
                        }
                        fmt.Println(v3, fi12, v4)
                        if (false || (74 <= (^(v5)))) {
                            v1 = v1
                            if ((fi13 > (fi13 % (((^(fi13))) | 1))) || false) {
                                var v25 int8 = int8(v5)
                                _ = v25
                            } else {
                                v26 := GenS3{F0: ("\n\t" + "\n\t"), F1: (42 < 166), F2: -2147483648}
                                v1 = 2
                                var v27 bool = (1 != 0)
                                _ = v26
                                _ = v27
                            }
                        }
                    }
                    v6 = append(v6, -216)
                }
                if false {
                    v6 = append(v6, (fi13 & (fi12 >> ((42) & 63))))
                }
                var v28 uint = 42
                fmt.Println(v2)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                fmt.Println(v1, v1)
                _ = v28
            } else {
                var v29 float32 = 3.14159
                v3 += 0
                _ = v29
            }
            v4 -= (3.14159 * (0.1 * v2))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi13
        }
        _ = fi12
    }
    switch 385 {
    case -404:
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 string = "foo bar"
            if (!true) {
                var v32 uint = 591
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v32
            } else {
                v5 -= 692
                v5 -= 0
                v33 := GenS1{F0: 2, F1: 1, F2: int8(v1)}
                _ = v33
            }
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
