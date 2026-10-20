package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 bool }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 string; F1 string; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int32) uint64 {
    v7 := []int{128}
    p0 += p0
    for fi8 := 0; fi8 < 2; fi8++ {
        p0 += p0
        p0 = (p0 / (((p0 * (p0 * -599))) | 1))
        var v9 float32 = 1.5
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: -479, F1: -128, F2: false}
    _ = p0
    _ = v10
    _ = v7
    return 42
}

func main() {
    v1 := []int{-2147483648, -1}
    var v2 float32 = (3.14159 + float32(3.14159))
    var v3 int32 = -158
    var v4 bool = (!(false || ((((float32(float64(v2)) / v2) + v2) < v2) || false)))
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    switch 517 {
    case 355:
        v6[1] = 455
    case -2147483648:
        switch ("" + ("" + "foo bar")) {
        case "hello":
            _ = len(v5)
        case "foo bar":
            v3 += (v3 + (v3 << ((2) & 63)))
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        v6[2147483648] = 65535
    }
    if (1 > 18446744073709551615) {
        switch "\n\t" {
        case "foo bar":
            v3 -= 0
        case "":
            var v11 int8 = -110
            _ = v11
        case "hello":
            v4 = v4
        }
    }
    fmt.Println(v2, v2)
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS1{F0: -9223372036854775808, F1: -128, F2: (v3 > v3)}
        v1 = append(v1, (fi12 * fi12))
        _ = fi12
        _ = v13
    }
    switch 371 {
    case 4294967295:
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint8 = 155
            _ = fi14
            _ = v15
        }
    case -384:
        v4 = v4
    case 32768:
        v16 := GenS2{F0: 1000, F1: 1, F2: 42}
        _ = v16
    default:
        v17 := GenS1{F0: 32768, F1: -1, F2: (-2147483648 < 2147483647)}
        _ = v17
    }
    delete(v5, "")
    var v18 string = "\n\t"
    v19 := GenS1{F0: 573, F1: 42, F2: v4}
    v18 = (v18 + "")
    if false {
        for fi20 := 0; fi20 < 2; fi20++ {
            v2 += (v2 - v2)
            for fi21 := 0; fi21 < 4; fi21++ {
                switch "" {
                case "hello":
                    v22 := GenS1{F0: (fi20 - (fi20 ^ -595)), F1: 100, F2: v4}
                    _ = v22
                case "foo bar":
                    v1 = append(v1, -999)
                default:
                    v1 = append(v1, (fi20 ^ fi20))
                }
                _ = fi21
            }
            fmt.Println(v18, v4)
            _ = fi20
        }
        _ = len(v1)
    } else {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 uint = 814
            for fi25 := 0; fi25 < 10; fi25++ {
                if (-1 <= -742) {
                    fmt.Println(v18)
                } else {
                    if v4 {
                        v26 := GenS3{F0: "hello", F1: ("foo bar" + (v18 + v18)), F2: (!false)}
                        if (1 < 0) {
                            var v27 uint = 18446744073709551615
                            var v28 uint32 = 42
                            fmt.Println(fi25)
                            _ = v27
                            _ = v28
                        }
                        _ = v26
                    } else {
                        v2 = (float32(float64(v2)) - v2)
                        var v29 float32 = v2
                        if len(v1) > 0 {
                            v1[0] = (fi23 + 127)
                        }
                        _ = v29
                    }
                    v30 := GenS2{F0: -1, F1: 32767, F2: int(v24)}
                    v3 = -1
                    _ = v30
                }
                if len(v1) > 0 {
                    v1[0] = fi23
                }
                v4 = (594 < uint16(v24))
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
        v31 := GenS3{F0: "foo bar", F1: v18, F2: v4}
        _ = v31
    }
    delete(v6, 4294967295)
    v32 := GenS4{F0: 0, F1: 622}
    v33 := GenS2{F0: 0, F1: -32768, F2: 0}
    v3 -= v3
    switch -851 {
    case 32768:
        for fi34 := 0; fi34 < 3; fi34++ {
            for fi35 := 0; fi35 < 5; fi35++ {
                if len(v1) > 0 {
                    v1[0] = (fi35 << ((187) & 63))
                }
                v4 = (0 <= 42)
                _ = fi35
            }
            if false {
                _ = len(v6)
                v2 = ((v2 / (((0.1 * (0.0 / v2)) + 0.1) - v2)) * float32(0.1))
            }
            v2 -= v2
            _ = fi34
        }
    case 167:
        var v36 uint8 = 128
        _ = v36
    default:
        v37 := v6[4294967295]
        _ = v37
    }
    var v38 int64 = 402
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
