package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 int64; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint }
func genFn0(p0 uint8, p1 uint, p2 int64) int8 {
    v5 := []int{2147483647, 9223372036854775807, -1, 9223372036854775807}
    v6 := GenS1{F0: p1, F1: p2}
    var v7 int32 = 890
    var v8 float32 = 0.0
    var v9 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 17
}
func genFn1(p0 int) uint64 {
    fmt.Println(p0)
    v10 := GenS2{F0: 16, F1: 0, F2: ("a" + "hello")}
    _ = p0
    _ = v10
    return 18446744073709551615
}

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 93
    var v3 uint16 = 0
    var v4 bool = (!false)
    v11 := []int{9223372036854775807, 128}
    v3 += 32767
    v4 = (!true)
    if v4 {
        var v12 float64 = ((v1 / v1) * v1)
        var v13 float64 = 1e10
        _ = v12
        _ = v13
    }
    if ((((((v1 / float64(1e10)) + (0.0 / v1)) / v1) * float64(((1e10 + float32(1.5)) + 1e10))) > (v1 - ((0.0 / v1) * 3.14159))) && ((v4 && ((v4 || v4) || v4)) && v4)) {
        var v14 int32 = -1
        var v15 uint16 = 216
        _ = v14
        _ = v15
    } else {
        if (42 > 32767) {
            v16 := GenS2{F0: 0, F1: -87, F2: ("\n\t" + ("" + ("\n\t" + "")))}
            v17 := GenS3{F0: (^(v3))}
            if ((!v4) && ((v4 || (-2147483648 >= -372)) || v4)) {
                v1 = v1
                var v18 int32 = 42
                _ = v18
            } else {
                v19 := GenS4{F0: "\n\t"}
                fmt.Println(v4)
                _ = v19
            }
            _ = v16
            _ = v17
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 uint32 = 505
        _ = fi20
        _ = v21
    }
    var v22 float32 = -2.25
    _ = len(v11)
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            _ = len(v11)
            _ = fi24
        }
        _ = fi23
    }
    if (32767 > -1) {
        var v25 uint32 = 975
        if true {
            for fi26 := 0; fi26 < 3; fi26++ {
                v2 += v2
                _ = fi26
            }
            fmt.Println(v3, v2)
        } else {
            _ = len(v11)
            v27 := GenS4{F0: "foo bar"}
            _ = v27
        }
        _ = v25
    } else {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := GenS5{F0: -2147483648, F1: 9223372036854775807}
            _ = fi28
            _ = v29
        }
        var v30 float64 = 0.0
        _ = v30
    }
    v22 = (v22 - v22)
    v31 := GenS4{F0: ""}
    if false {
        v4 = (!(v4 && false))
    }
    var v32 int64 = 42
    if ((v2 + 1) == v2) {
        for fi33 := 0; fi33 < 19; fi33++ {
            for fi34 := 0; fi34 < 10; fi34++ {
                switch ("foo bar" + "") {
                case "foo bar":
                    v35 := GenS5{F0: 2147483647, F1: fi34}
                    _ = v35
                case "\n\t":
                    if ((v3 * v3) >= 0) {
                        v3 += v3
                        v22 -= (v22 + v22)
                    }
                case "hello":
                    v11 = append(v11, fi34)
                default:
                    var v36 uint16 = v3
                    _ = v36
                }
                switch ((("" + "a") + "") + "") {
                case "a":
                    v37 := GenS6{F0: 336, F1: 2}
                    _ = v37
                case "\n\t":
                    var v38 uint16 = v3
                    _ = v38
                case "":
                    var v39 uint = 989
                    _ = v39
                default:
                    if v4 {
                        v1 = 1e10
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        v1 = 3.14159
                    }
                }
                _ = fi34
            }
            v4 = v4
            _ = fi33
        }
        _ = len(v11)
    }
    switch 4294967295 {
    case 4294967295:
        v40 := GenS3{F0: v3}
        _ = v40
    case 9223372036854775807:
        if v4 {
            fmt.Println(v2, v4)
        }
    case 127:
        v41 := GenS7{F0: v2, F1: 0, F2: 42}
        _ = v41
    default:
        if false {
            v1 = 0.1
        } else {
            for fi42 := 0; fi42 < 10; fi42++ {
                var v43 uint32 = 6
                v3 -= (v3 * v3)
                for fi44 := 0; fi44 < 5; fi44++ {
                    var v45 uint16 = (v3 % ((255) | 1))
                    _ = len(v11)
                    _ = fi44
                    _ = v45
                }
                _ = fi42
                _ = v43
            }
        }
    }
    if v4 {
        v46 := GenS3{F0: (v3 * (v3 << ((70) & 63)))}
        var v47 int64 = v32
        _ = v46
        _ = v47
    }
    v48 := GenS4{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v48
}
