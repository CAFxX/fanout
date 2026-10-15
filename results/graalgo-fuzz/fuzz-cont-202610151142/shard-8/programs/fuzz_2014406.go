package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
func genFn0(p0 uint32) string {
    p0 = (p0 >> ((4294967295) & 63))
    p0 += p0
    p0 += (^(p0))
    var v7 int64 = 42
    if true {
        for fi8 := 0; fi8 < 20; fi8++ {
            v7 -= (v7 << ((2) & 63))
            v7 -= v7
            var v9 string = "hello"
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: v7, F1: 0}
        _ = v10
    }
    _ = p0
    _ = v7
    return ("hello" + "a")
}

func main() {
    var v1 int = 1
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := []uint32{83}
    v5 := make(map[uint32]int)
    v6 := make(map[string]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    switch 42 {
    case 1:
        if (292 > 0) {
            switch 864 {
            case -2147483648:
                _ = len(v3)
            case 127:
                switch "hello" {
                case "hello":
                    v1 = 2147483647
                default:
                    for fi11 := 0; fi11 < 12; fi11++ {
                        fmt.Println(fi11, v1)
                        if (2147483647 >= -949) {
                            fmt.Println(fi11)
                            _ = len(v5)
                            for fi12 := 0; fi12 < 6; fi12++ {
                                v1 = (v1 % ((fi11) | 1))
                                if false {
                                    _ = len(v5)
                                } else {
                                    switch 4294967295 {
                                    case 950:
                                        if false {
                                            var v13 uint16 = 2
                                            _ = v13
                                        }
                                    case 4294967295:
                                        var v14 float64 = 0.0
                                        _ = v14
                                    case 1:
                                        v1 -= (v1 >> ((2) & 63))
                                    default:
                                        delete(v6, ("a" + "hello"))
                                    }
                                }
                                v2 = append(v2, 0)
                                _ = fi12
                            }
                        }
                        _ = fi11
                    }
                }
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
    default:
        var v15 float32 = 0.1
        _ = v15
    }
    switch v1 {
    case 0:
        var v16 int32 = 355
        _ = v16
    case 4294967295:
        var v17 int16 = -1
        _ = v17
    default:
        if false {
            var v18 int64 = -1
            _ = v18
        } else {
            fmt.Println(v1)
            v5[65535] = (v1 ^ v1)
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
    }
    var v19 uint8 = 255
    var v20 string = ("" + "hello")
    fmt.Println(v19, v1)
    if ((763 >= 597) || (!false)) {
        v19 = (v19 >> ((65535) & 63))
        if false {
            v19 = v19
            _ = len(v4)
            v19 = (v19 ^ 181)
        }
        var v21 bool = true
        _ = v21
    } else {
        var v22 float64 = (0.0 * -2.25)
        var v23 uint16 = 32767
        var v24 float64 = ((3.14159 - 0.1) - v22)
        _ = v22
        _ = v23
        _ = v24
    }
    v20 = v20
    var v25 float64 = (0.0 + 1.5)
    for fi26 := 0; fi26 < 13; fi26++ {
        switch "a" {
        case "foo bar":
            var v27 uint32 = 65535
            _ = v27
        case "hello":
            v1 -= -9223372036854775808
        }
        v25 -= 3.14159
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
