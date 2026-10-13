package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
func genFn0(p0 int32, p1 uint) bool {
    p1 = p1
    p1 = 4294967295
    v7 := []int{-2147483648}
    var v8 string = "foo bar"
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (-1 <= -9223372036854775808)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 66
    v3 := []int{110, 4294967295}
    v4 := make(map[uint32]uint32)
    var v5 bool = true
    var v6 uint = 737
    fmt.Println(v6, v2, v5)
    var v9 uint8 = (v2 % ((v2) | 1))
    fmt.Println(v6, v5, v2)
    v1[-791] = (935 < 4294967295)
    for fi10 := 0; fi10 < 4; fi10++ {
        v2 = v2
        switch fi10 {
        case -1:
            var v11 float64 = 0.0
            _ = v11
        case 42:
            switch "" {
            case "foo bar":
                if ((v2 % ((v9) | 1)) <= 214) {
                    fmt.Println(fi10, v6)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (1 > uint16(v6)) {
                        if (0 > 255) {
                            v9 = 127
                            if v5 {
                                if (0 < -2147483648) {
                                    if v5 {
                                        var v12 string = "\n\t"
                                        _ = v12
                                    } else {
                                        var v13 int64 = -1
                                        v13 = (v13 << ((v6) & 63))
                                        _ = v13
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                        } else {
                            v1[fi10] = (1 >= 859)
                        }
                        var v14 int32 = 2147483647
                        _ = v14
                    }
                }
            case "a":
                switch (^(fi10)) {
                case 2147483647:
                    _ = len(v4)
                default:
                    if (1000 > -32768) {
                        _ = len(v3)
                        if v5 {
                            var v15 bool = true
                            _ = v15
                        }
                        switch "foo bar" {
                        case "hello":
                            if v5 {
                                switch -403 {
                                case 952:
                                    v9 += v2
                                case 32768:
                                    var v16 float64 = 1.5
                                    _ = v16
                                default:
                                    v17 := GenS1{F0: -288, F1: 32767, F2: (9223372036854775807 > 9223372036854775807)}
                                    _ = v17
                                }
                            } else {
                                if false {
                                    switch 409 {
                                    case 314:
                                        v18 := GenS1{F0: -417, F1: (fi10 * 9223372036854775807), F2: v5}
                                        _ = v18
                                    }
                                } else {
                                    v19 := GenS2{F0: 9223372036854775807, F1: 4294967295}
                                    var v20 uint32 = 1
                                    _ = v19
                                    _ = v20
                                }
                                if v5 {
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v21 int32 = 42
                                        _ = v21
                                    case "":
                                        for fi22 := 0; fi22 < 19; fi22++ {
                                            _ = len(v3)
                                            v6 -= (^(v6))
                                            _ = fi22
                                        }
                                    default:
                                        var v23 uint32 = 2
                                        _ = v23
                                    }
                                }
                                v24 := GenS3{F0: (!(0 > 119)), F1: -128, F2: 1}
                                _ = v24
                            }
                        case "\n\t":
                            v5 = v5
                        case "foo bar":
                            fmt.Println(v6, v6, v5)
                        }
                    } else {
                        v2 = v2
                    }
                }
            case "hello":
                for fi25 := 0; fi25 < 3; fi25++ {
                    v5 = true
                    fmt.Println(fi25)
                    _ = fi25
                }
            default:
                for fi26 := 0; fi26 < 15; fi26++ {
                    var v27 int32 = -757
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if v5 {
                        _ = len(v4)
                    }
                    _ = fi26
                    _ = v27
                }
            }
        default:
            fmt.Println(v5)
        }
        var v28 uint16 = 65535
        _ = fi10
        _ = v28
    }
    v29 := GenS3{F0: true, F1: 100, F2: v6}
    fmt.Println(v6, v2, v9)
    v6 += (v6 % ((666) | 1))
    if v5 {
        v9 -= (v9 * v9)
        var v30 float64 = (((1e10 * 3.14159) - 1.5) + (0.1 * 3.14159))
        v30 -= v30
        _ = v30
    } else {
        v6 -= v6
        v31 := v1[112]
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
