package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 int }

func main() {
    var v1 int = -1
    var v2 float64 = (3.14159 / (((float64(((((((1.5 + (1e10 + 1.5)) + 1e10) * float32(0.1)) / 0.1) * 3.14159) + 0.1)) - 3.14159) + (((0.0 * (((1.5 / 1e10) + 0.0) + (1e10 - 0.0))) - (((float64(0.1) / 3.14159) / 0.1) + 0.1)) + 0.0)) + (1e10 / 1.0)))
    v3 := []uint32{2, 65535, 811}
    var v4 float64 = 0.0
    v5 := make(map[int]uint32)
    v2 += 1e10
    if true {
        v1 = v1
        v6 := GenS1{F0: -2, F1: 4294967295}
        _ = v6
    }
    if (!(187 < 255)) {
        v2 -= (v2 * v2)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 1, F1: 1}
        switch "\n\t" {
        case "a":
            for fi9 := 0; fi9 < 1; fi9++ {
                if true {
                    v10 := v5[fi7]
                    if (1 > v10) {
                        v3 = append(v3, (^(v10)))
                        switch "foo bar" {
                        case "hello":
                            if ((!true) || (42 != 9)) {
                                fmt.Println(v10, v1)
                                var v11 float32 = 3.14159
                                _ = v11
                            }
                        default:
                            v5[266] = 482
                        }
                        v12 := GenS2{F0: 128, F1: 0, F2: 4294967295}
                        _ = v12
                    }
                    _ = v10
                } else {
                    switch ("a" + "") {
                    case "hello":
                        var v13 uint = 0
                        _ = v13
                    default:
                        var v14 uint32 = 268
                        _ = v14
                    }
                }
                _ = fi9
            }
        case "hello":
            var v15 int16 = 415
            _ = v15
        default:
            if len(v3) > 0 {
                v3[0] = 200
            }
        }
        fmt.Println(v1, v1)
        _ = fi7
        _ = v8
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        switch 2147483647 {
        case -966:
            v3 = append(v3, 2147483647)
        }
        delete(v5, 524)
        _ = fi16
    }
    if (42 < 480) {
        var v17 uint8 = 42
        _ = v17
    } else {
        if true {
            if true {
                switch "" {
                case "\n\t":
                    v1 = v1
                case "foo bar":
                    var v18 int16 = 1000
                    _ = v18
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = 751
                    }
                default:
                    for fi19 := 0; fi19 < 14; fi19++ {
                        var v20 uint8 = 255
                        fmt.Println(v1)
                        v21 := GenS2{F0: -520, F1: 1, F2: (v1 | v1)}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS2{F0: (v1 - fi22), F1: 65535, F2: 665}
                    var v24 string = "hello"
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 9; fi25++ {
                    var v26 int8 = -47
                    var v27 bool = true
                    fmt.Println(v26)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            } else {
                v28 := v5[9223372036854775807]
                _ = v28
            }
            v5[1] = 42
            var v29 int = (v1 | (^(v1)))
            _ = v29
        }
    }
    if (0 >= 255) {
        v30 := GenS2{F0: (v1 | -2147483648), F1: 2, F2: 4294967295}
        _ = v30
    } else {
        var v31 uint64 = 18446744073709551615
        _ = v31
    }
    v1 = (v1 * (v1 | (^(v1))))
    if false {
        if false {
            var v32 float32 = 3.14159
            var v33 uint8 = 203
            _ = v32
            _ = v33
        }
    } else {
        var v34 uint32 = 2147483647
        if false {
            v2 = 0.0
            var v35 float64 = (v2 / v2)
            var v36 int = (v1 % ((v1) | 1))
            _ = v35
            _ = v36
        } else {
            v37 := GenS1{F0: 42, F1: 18446744073709551615}
            v1 = (v1 | (v1 << ((42) & 63)))
            _ = v37
        }
        v2 += ((v2 / v4) + 1e10)
        _ = v34
    }
    v4 -= v4
    var v38 uint = 255
    for fi39 := 0; fi39 < 8; fi39++ {
        v1 += fi39
        _ = fi39
    }
    switch 4294967295 {
    case -2147483648:
        for fi40 := 0; fi40 < 12; fi40++ {
            v41 := GenS1{F0: -8, F1: 162}
            _ = fi40
            _ = v41
        }
    }
    fmt.Println(v38, v38)
    v42 := GenS1{F0: 127, F1: 40}
    v38 = (v38 % ((v38) | 1))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v5
}
