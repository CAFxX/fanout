package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -1
    v3 := []uint32{582, 576}
    var v4 uint8 = 255
    v5 := make(map[string]uint32)
    var v6 bool = (3.14159 >= 3.14159)
    v7 := GenS1{F0: true, F1: 878}
    if (!(v2 <= (v2 % ((v2) | 1)))) {
        v8 := GenS1{F0: (((v2 * v2) < v2) && false), F1: 1}
        _ = v8
    }
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 2; fi10++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            v9 -= v9
            _ = fi11
        }
        _ = fi10
    }
    v9 = 42
    if v6 {
        v12 := GenS2{F0: (1 >= v4)}
        _ = v12
    }
    switch 2147483647 {
    case 1:
        switch "foo bar" {
        case "foo bar":
            for fi13 := 0; fi13 < 4; fi13++ {
                switch fi13 {
                case 2147483647:
                    switch fi13 {
                    case -9223372036854775808:
                        if v6 {
                            v14 := GenS3{F0: -9223372036854775808, F1: "\n\t"}
                            var v15 int = fi13
                            _ = v14
                            _ = v15
                        } else {
                            if v6 {
                                var v16 int = (-(fi13))
                                var v17 uint64 = 42
                                _ = v16
                                _ = v17
                            }
                        }
                    case -681:
                        if true {
                            v18 := GenS4{F0: v2, F1: 0}
                            _ = v18
                        } else {
                            v19 := GenS4{F0: v2, F1: 18446744073709551615}
                            _ = v19
                        }
                    default:
                        if v6 {
                            switch "a" {
                            case "foo bar":
                                v4 -= 70
                            }
                            switch 4294967295 {
                            case -9223372036854775808:
                                v2 = v2
                            case 905:
                                fmt.Println(fi13, v2)
                            case 32768:
                                if len(v3) > 0 {
                                    v3[0] = 2147483648
                                }
                            default:
                                fmt.Println(fi13, fi13, v9)
                            }
                        }
                    }
                }
                _ = fi13
            }
        }
    case -1:
        v4 += (v4 + (v4 / ((59) | 1)))
    default:
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := v5[("a" + ("\n\t" + "a"))]
            if (v4 <= (^(v4))) {
                fmt.Println(fi20)
                if false {
                    v9 -= 845
                } else {
                    v4 -= v4
                    var v22 float32 = (float32(-2.25) - 1.5)
                    if (18446744073709551615 < 4294967295) {
                        fmt.Println(v6, v22)
                        var v23 uint = 65535
                        _ = len(v3)
                        _ = v23
                    }
                    _ = v22
                }
                v2 = (v2 / ((v2) | 1))
            }
            _ = fi20
            _ = v21
        }
    }
    if v6 {
        var v24 bool = v6
        fmt.Println(v4, v24, v2)
        v4 += 109
        _ = v24
    } else {
        for fi25 := 0; fi25 < 19; fi25++ {
            v9 = 2
            v6 = ((v2 | (v2 ^ v2)) > v2)
            _ = fi25
        }
        fmt.Println(v2)
        v26 := GenS2{F0: (!(!(((0.1 * 0.0) != (3.14159 * -2.25)) || (!(0 != 0)))))}
        _ = v26
    }
    if v6 {
        var v27 uint16 = (v9 + v9)
        _ = v27
    }
    v9 = v9
    v28 := GenS2{F0: false}
    v4 -= 255
    v29 := GenS5{F0: 255}
    v2 = v2
    v9 += 952
    switch "" {
    case "":
        for fi30 := 0; fi30 < 16; fi30++ {
            v5["\n\t"] = 42
            v6 = v6
            _ = fi30
        }
    case "foo bar":
        if false {
            v3 = append(v3, 893)
            var v31 float32 = float32(((1e10 * 0.0) + 3.14159))
            _ = v31
        }
    case "hello":
        v9 -= v9
    }
    v4 += 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
