package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }

func main() {
    var v1 int32 = 849
    var v2 float64 = (0.1 / 3.14159)
    v3 := []int{-778, -9223372036854775808, 128}
    v4 := []string{"hello"}
    v5 := []uint32{874, 791, 360}
    v1 += v1
    var v6 int64 = 1
    var v7 uint = 255
    v6 = v6
    v7 = v7
    v8 := GenS1{F0: 2147483648}
    if true {
        if false {
            var v9 float64 = 3.14159
            _ = v9
        } else {
            if false {
                fmt.Println(v1, v7)
            }
        }
        switch "foo bar" {
        case "a":
            for fi10 := 0; fi10 < 12; fi10++ {
                _ = len(v5)
                v11 := GenS2{F0: 255, F1: 2, F2: 42}
                _ = fi10
                _ = v11
            }
        case "hello":
            var v12 uint = (v7 + v7)
            _ = v12
        case "":
            if len(v5) > 0 {
                v5[0] = 486
            }
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    } else {
        if false {
            if false {
                if ((1.5 * (0.1 / (3.14159 - 1e10))) == float32(1.5)) {
                    switch -313 {
                    case -132:
                        switch 1 {
                        case -9223372036854775808:
                            v1 -= 24
                        case 489:
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                        default:
                            switch -845 {
                            case 2147483647:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                    default:
                        var v13 uint64 = 552
                        _ = v13
                    }
                } else {
                    var v14 int = 42
                    _ = v14
                }
                v2 = (1.5 / (v2 * 3.14159))
            } else {
                v15 := GenS2{F0: 127, F1: 42, F2: 281}
                v1 = (v1 | (v1 << (((v7 / ((18446744073709551615) | 1))) & 63)))
                _ = v15
            }
        }
        switch ("foo bar" + (("a" + "") + "")) {
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "foo bar":
            _ = len(v5)
        case "a":
            v2 = float64((3.14159 * (-2.25 / 3.14159)))
        default:
            v16 := GenS2{F0: 127, F1: 9, F2: 65535}
            _ = v16
        }
    }
    v6 += v6
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v17 int32 = 1
    var v18 float64 = 1.5
    if (-1 == 927) {
        v6 += (^(v6))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
