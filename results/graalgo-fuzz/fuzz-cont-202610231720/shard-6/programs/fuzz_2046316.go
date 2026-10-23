package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0(p0 uint32, p1 uint8) uint16 {
    fmt.Println(p0, p0)
    v7 := []int{676, 538, 32767, 1}
    switch "" {
    case "":
        v8 := GenS1{F0: -32768, F1: (("foo bar" + ("foo bar" + "")) + (("\n\t" + "hello") + "foo bar"))}
        _ = v8
    case "hello":
        v9 := GenS2{F0: 42}
        _ = v9
    case "\n\t":
        switch "foo bar" {
        case "a":
            fmt.Println(p1)
        case "foo bar":
            var v10 float32 = (((0.0 - (((float32(0.0) + 1e10) / 3.14159) * (float32(0.1) + (0.0 - 0.1)))) * -2.25) / 1.5)
            _ = v10
        case "":
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 bool = (!((!(42 > -911)) || true))
                if len(v7) > 0 {
                    v7[0] = (fi11 % ((fi11) | 1))
                }
                _ = fi11
                _ = v12
            }
        }
    }
    v13 := GenS1{F0: -1, F1: "a"}
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return 1
}

func main() {
    var v1 float32 = -2.25
    var v2 uint8 = 42
    var v3 int64 = -9223372036854775808
    var v4 int64 = v3
    var v5 uint16 = 42
    var v6 int16 = 601
    switch -9223372036854775808 {
    case 4294967295:
        v14 := []int{-744, 32767, 492, 32767}
        _ = v14
    case 0:
        fmt.Println(v1, v4, v2)
    }
    fmt.Println(v2)
    if true {
        v15 := []int{-631, 0}
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 19; fi17++ {
                fmt.Println(v3, fi16, v2)
                fmt.Println(v6, fi16, v3)
                switch "\n\t" {
                case "\n\t":
                    var v18 uint8 = (^(v2))
                    _ = v18
                case "foo bar":
                    switch fi17 {
                    case 1:
                        v19 := GenS3{F0: 42, F1: (v4 - int64(v2)), F2: v2}
                        _ = v19
                    case 2147483647:
                        v5 = v5
                    case 0:
                        var v20 uint32 = 65535
                        _ = v20
                    default:
                        if false {
                            fmt.Println(v4)
                            switch fi17 {
                            case 128:
                                switch "hello" {
                                case "\n\t":
                                    v21 := GenS2{F0: (v3 * (v4 + (^(v3))))}
                                    _ = v21
                                }
                            case -9223372036854775808:
                                v1 = v1
                            default:
                                switch fi17 {
                                case 288:
                                    var v22 uint32 = 670
                                    _ = v22
                                case -175:
                                    v4 = (v4 >> ((865) & 63))
                                case 42:
                                    var v23 uint32 = 65535
                                    _ = v23
                                default:
                                    if (v5 > 1) {
                                        v3 += v3
                                    } else {
                                        var v24 float32 = ((1e10 / v1) + -2.25)
                                        var v25 int64 = (-(v4))
                                        _ = v24
                                        _ = v25
                                    }
                                }
                            }
                        }
                    }
                case "a":
                    var v26 int = (^(fi17))
                    _ = v26
                }
                _ = fi17
            }
            v27 := GenS4{F0: 886, F1: (v2 << ((88) & 63)), F2: v5}
            _ = fi16
            _ = v27
        }
        _ = len(v15)
        _ = v15
    } else {
        v5 = v5
    }
    var v28 uint16 = v5
    var v29 int16 = v6
    if false {
        v30 := []int{22}
        v28 -= 729
        var v31 int16 = (v6 % (((v6 << ((2) & 63))) | 1))
        _ = v30
        _ = v31
    }
    v32 := []int{32767, 4294967295}
    v5 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
