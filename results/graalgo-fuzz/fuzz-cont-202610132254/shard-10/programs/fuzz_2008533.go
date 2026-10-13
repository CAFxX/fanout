package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() int16 {
    var v4 float64 = (0.1 + 0.0)
    var v5 string = "\n\t"
    var v6 float64 = v4
    var v7 uint64 = 42
    var v8 int64 = 9223372036854775807
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{65535, 0, 65535, 2147483647}
    var v2 uint8 = 12
    v3 := []uint32{4294967295, 0, 1, 72}
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int32 = 0
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 1, F1: ""}
    fmt.Println(v2)
    v2 = 42
    v12 := GenS1{F0: 65535, F1: (("" + "foo bar") + "a")}
    var v13 uint = 18446744073709551615
    var v14 int32 = -527
    var v15 uint = 18446744073709551615
    switch "hello" {
    case "hello":
        if (42 != -1) {
            v16 := GenS2{F0: v14}
            v1 = append(v1, 943)
            v14 = -2147483648
            _ = v16
        } else {
            var v17 float64 = (1.5 + (1e10 - 1e10))
            v18 := GenS2{F0: (v14 - v14)}
            if (32768 > -683) {
                for fi19 := 0; fi19 < 6; fi19++ {
                    v20 := GenS3{F0: 18446744073709551615, F1: (fi19 | fi19)}
                    _ = fi19
                    _ = v20
                }
            }
            _ = v17
            _ = v18
        }
    case "foo bar":
        var v21 float64 = float64(0.0)
        _ = v21
    }
    if false {
        var v22 float64 = 1.5
        _ = v22
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi23
    }
    if ((v14 | (v14 ^ v14)) != (v14 + v14)) {
        switch "hello" {
        case "a":
            v13 += (v13 % (((v13 * (^(v15)))) | 1))
        case "\n\t":
            _ = len(v3)
        }
        v24 := GenS4{F0: ("" + "foo bar"), F1: 42}
        v25 := GenS2{F0: v14}
        _ = v24
        _ = v25
    }
    if (229 != 1) {
        if len(v3) > 0 {
            v3[0] = 921
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v13 += v13
        switch -18 {
        case 32767:
            if len(v1) > 0 {
                v1[0] = 805
            }
        case 9223372036854775807:
            for fi26 := 0; fi26 < 19; fi26++ {
                fmt.Println(v13, fi26)
                for fi27 := 0; fi27 < 17; fi27++ {
                    if false {
                        fmt.Println(v2, fi26)
                        fmt.Println(v13, fi27)
                    } else {
                        fmt.Println(v13, v14, v2)
                        v28 := GenS1{F0: 0, F1: "foo bar"}
                        _ = v28
                    }
                    v15 += 0
                    if true {
                        v3 = append(v3, 0)
                        if len(v3) > 0 {
                            v3[0] = 577
                        }
                        v14 += -269
                    } else {
                        v29 := GenS2{F0: v14}
                        v13 = 18446744073709551615
                        _ = v29
                    }
                    var v30 int32 = v14
                    var v31 uint8 = 158
                    _ = fi27
                    _ = v30
                    _ = v31
                }
                _ = fi26
            }
        case 520:
            var v32 uint32 = 4294967295
            _ = v32
        default:
            v2 = (v2 * 112)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
