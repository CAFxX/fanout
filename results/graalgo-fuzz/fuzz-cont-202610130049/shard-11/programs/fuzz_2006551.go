package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint16; F1 string; F2 int8 }

func main() {
    v1 := []string{}
    var v2 int = 113
    v3 := []uint32{2147483647, 1, 698, 443}
    v4 := []int{v2, v2, v2, (v2 / ((v2) | 1))}
    v5 := []string{"", ("\n\t" + "")}
    _ = len(v4)
    v2 += v2
    if (0.0 > (1e10 + -2.25)) {
        if len(v4) > 0 {
            v4[0] = (v2 - v2)
        }
        v2 -= -249
    } else {
        if (0 >= 88) {
            fmt.Println(v2)
        }
        var v6 int8 = 1
        _ = v6
    }
    fmt.Println(v2)
    if (978 != 2147483647) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v7 int16 = 32767
        _ = v7
    }
    if (!(0 > 0)) {
        switch ((("hello" + "a") + (("hello" + "hello") + "a")) + "foo bar") {
        case "a":
            if ((-1 <= -772) && true) {
                v2 += (v2 + 0)
                _ = len(v4)
            } else {
                for fi8 := 0; fi8 < 7; fi8++ {
                    v2 = v2
                    v9 := GenS1{F0: 119}
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                    _ = fi8
                    _ = v9
                }
                var v10 float32 = 0.0
                _ = v10
            }
        default:
            var v11 uint8 = 42
            _ = v11
        }
        if false {
            var v12 int32 = -692
            v12 += (-(v12))
            if (float32(3.14159) != (1.5 / 1.0)) {
                if ((((1e10 - 1e10) * -2.25) - float64(3.14159)) > 1e10) {
                    var v13 int = -9223372036854775808
                    _ = v13
                }
                v14 := GenS1{F0: 319}
                v2 -= -2147483648
                _ = v14
            } else {
                v2 += (-(v2))
                v2 = (v2 ^ (v2 >> ((396) & 63)))
            }
            _ = v12
        } else {
            if false {
                if false {
                    v15 := GenS2{F0: 127, F1: ""}
                    _ = v15
                }
            } else {
                switch "\n\t" {
                case "":
                    var v16 uint8 = 121
                    _ = v16
                default:
                    v17 := GenS3{F0: 255, F1: 2}
                    _ = v17
                }
                v2 = v2
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 string = "foo bar"
                    v19 = (v19 + v19)
                    for fi20 := 0; fi20 < 18; fi20++ {
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                if (0 == 1) {
                                    var v21 int64 = -1
                                    var v22 int16 = 32767
                                    _ = v21
                                    _ = v22
                                } else {
                                    v2 -= fi20
                                    var v23 uint32 = 1
                                    _ = v23
                                }
                                var v24 string = "hello"
                                _ = v24
                            } else {
                                fmt.Println(fi20, v19)
                                if (!false) {
                                    var v25 int16 = 0
                                    v26 := GenS4{F0: 32767, F1: 0}
                                    v27 := GenS4{F0: 1, F1: 9223372036854775807}
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                }
                            }
                        case "\n\t":
                            fmt.Println(v2, v2, v19)
                        }
                        _ = fi20
                    }
                    _ = fi18
                    _ = v19
                }
            }
            v28 := GenS5{F0: 1, F1: "foo bar", F2: 127}
            _ = v28
        }
    }
    v2 += (v2 & v2)
    v3 = append(v3, 42)
    v2 += v2
    if len(v1) > 0 {
        v1[0] = (("\n\t" + "\n\t") + "hello")
    }
    v2 = (v2 << ((279) & 63))
    var v29 float64 = 1e10
    var v30 uint32 = 511
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
