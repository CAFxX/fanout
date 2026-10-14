package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint8 = 255
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    v1 = 2
    v3["foo bar"] = 0
    fmt.Println(v1)
    v5 := GenS1{F0: -103, F1: 42, F2: 2}
    var v6 bool = false
    var v7 int8 = -1
    v8 := v3[("foo bar" + "\n\t")]
    _ = len(v4)
    var v9 float32 = (float32(1e10) * (float32(0.0) + -2.25))
    v8 = v8
    var v10 uint = 4294967295
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint64 = 4294967295
        switch "hello" {
        case "\n\t":
            _ = len(v2)
        case "foo bar":
            if (((float32(((float64(0.1) / 1.0) / (((3.14159 - (0.0 + 3.14159)) * ((float64(v9) / (float64((float32(1e10) / float32((1e10 + 0.0)))) + 1.5)) / 1.5)) * float64(((v9 - 1.5) / v9))))) <= float32(3.14159)) || v6) || ("" != "foo bar")) {
                for fi13 := 0; fi13 < 12; fi13++ {
                    if len(v2) > 0 {
                        v2[0] = v8
                    }
                    switch "foo bar" {
                    case "a":
                        switch ("hello" + "\n\t") {
                        case "":
                            fmt.Println(v1, fi11)
                        case "\n\t":
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v8, v8, v9)
                            case "":
                                var v14 uint16 = 42
                                _ = v14
                            }
                        default:
                            switch -1 {
                            case 32767:
                                delete(v4, fi13)
                            case -919:
                                if (v6 || (-1 >= -58)) {
                                    switch (fi11 / (((fi13 / (((^(fi11))) | 1))) | 1)) {
                                    case 2147483647:
                                        switch (fi11 % ((-9223372036854775808) | 1)) {
                                        case 42:
                                            switch ((("foo bar" + "hello") + "hello") + "foo bar") {
                                            case "\n\t":
                                                v10 = (^(v10))
                                            case "":
                                                v10 = v10
                                            case "a":
                                                if len(v2) > 0 {
                                                    v2[0] = 301
                                                }
                                                var v15 uint16 = 538
                                                _ = v15
                                            default:
                                                v9 -= v9
                                            }
                                        case 4294967295:
                                            var v16 uint64 = v12
                                            _ = v16
                                        }
                                        _ = len(v2)
                                    case -1:
                                        var v17 uint8 = (v1 - v1)
                                        _ = v17
                                    case -95:
                                        if len(v2) > 0 {
                                            v2[0] = 529
                                        }
                                    default:
                                        v18 := GenS2{F0: 42, F1: (v7 - v7), F2: v6}
                                        _ = v18
                                    }
                                    v19 := GenS3{F0: (v10 << ((241) & 63)), F1: fi11}
                                    _ = v19
                                } else {
                                    v10 -= (v10 ^ 18446744073709551615)
                                    if ((!(-1 < -59)) || false) {
                                        switch "" {
                                        case "\n\t":
                                            var v20 int = int(v12)
                                            _ = v20
                                        default:
                                            v21 := v3["hello"]
                                            _ = v21
                                        }
                                    } else {
                                        var v22 uint8 = (v1 | 42)
                                        var v23 uint = v10
                                        _ = v22
                                        _ = v23
                                    }
                                    var v24 int16 = 1
                                    if v6 {
                                        v24 = (^(v24))
                                    } else {
                                        var v25 uint16 = 2
                                        var v26 float32 = float32(0.1)
                                        v27 := GenS2{F0: -2147483648, F1: v7, F2: v6}
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v24
                                }
                            case -9223372036854775808:
                                v9 += (v9 - 1e10)
                            default:
                                v7 = -8
                            }
                        }
                    case "":
                        var v28 uint64 = (v12 & 1)
                        _ = v28
                    default:
                        var v29 int32 = 2147483647
                        _ = v29
                    }
                    _ = fi13
                }
                v30 := GenS2{F0: -1, F1: 127, F2: true}
                v3[("foo bar" + "foo bar")] = 2
                _ = v30
            }
        }
        if len(v2) > 0 {
            v2[0] = 797
        }
        _ = fi11
        _ = v12
    }
    if (float64(v9) != -2.25) {
        switch 32767 {
        case 32767:
            v2 = append(v2, (v8 ^ v8))
        case 9223372036854775807:
            v31 := v3["\n\t"]
            _ = v31
        case 127:
            var v32 uint16 = uint16(v10)
            _ = v32
        }
        if false {
            for fi33 := 0; fi33 < 13; fi33++ {
                v34 := GenS4{F0: 1, F1: 1}
                var v35 uint8 = (v1 % ((v1) | 1))
                _ = fi33
                _ = v34
                _ = v35
            }
        } else {
            v36 := GenS4{F0: 1, F1: 0}
            _ = v36
        }
        _ = len(v2)
    }
    if v6 {
        var v37 uint8 = (v1 + v1)
        v10 = v10
        _ = v37
    } else {
        fmt.Println(v6, v8)
    }
    var v38 uint16 = uint16(v1)
    v39 := v4[-2147483648]
    var v40 uint64 = 18446744073709551615
    var v41 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
