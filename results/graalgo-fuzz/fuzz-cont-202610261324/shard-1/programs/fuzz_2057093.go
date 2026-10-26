package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 string; F1 uint16 }

func main() {
    v1 := []uint32{}
    v2 := []string{"\n\t", "\n\t", "hello", "hello"}
    var v3 int32 = 42
    var v4 int32 = (-(v3))
    var v5 int8 = 0
    fmt.Println(v5, v3, v5)
    v6 := GenS1{F0: 42, F1: 255, F2: 756}
    fmt.Println(v5, v4)
    var v7 uint = 0
    fmt.Println(v7, v4)
    _ = len(v2)
    fmt.Println(v3, v3)
    v1 = append(v1, 2)
    if (-2.25 == float32(1e10)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v8 int64 = 42
        for fi9 := 0; fi9 < 9; fi9++ {
            v5 = (v5 * (-(v5)))
            switch fi9 {
            case 32768:
                var v10 int64 = -9223372036854775808
                _ = v10
            case 128:
                var v11 int16 = 42
                _ = v11
            }
            _ = fi9
        }
        _ = v8
    } else {
        fmt.Println(v5)
        var v12 uint8 = 127
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int16 = -32768
            switch ("foo bar" + (("hello" + ("hello" + "foo bar")) + (("" + "foo bar") + ("hello" + "hello")))) {
            case "":
                if ((!true) || ((v14 <= 42) && (1 == 0))) {
                    v14 += (v14 * int16(fi13))
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v14)
                    case "\n\t":
                        v15 := GenS1{F0: int64(v14), F1: 127, F2: 65535}
                        _ = v15
                    case "a":
                        fmt.Println(v4, v3, v12)
                    default:
                        var v16 string = "hello"
                        _ = v16
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        if ((!true) || ("\n\t" >= "hello")) {
                            if false {
                                var v17 float32 = -2.25
                                var v18 bool = true
                                _ = v17
                                _ = v18
                            }
                        }
                        _ = len(v1)
                        var v19 int8 = v5
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 9; fi20++ {
                            var v21 uint = 255
                            _ = len(v1)
                            var v22 uint8 = v12
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    v7 = v7
                }
            default:
                for fi23 := 0; fi23 < 11; fi23++ {
                    if (!(18446744073709551615 <= (v7 ^ v7))) {
                        if len(v1) > 0 {
                            v1[0] = uint32(v7)
                        }
                    } else {
                        var v24 uint8 = 2
                        _ = v24
                    }
                    if false {
                        _ = len(v2)
                        v5 += (v5 >> (((v7 & 969)) & 63))
                    } else {
                        switch ("" + "a") {
                        case "foo bar":
                            var v25 float64 = (1.5 + ((float64((0.0 * 1e10)) + 1e10) + 0.0))
                            if (-1 == (fi23 + (fi13 - fi13))) {
                                switch fi23 {
                                case 2147483647:
                                    v26 := GenS2{F0: ("" + "hello"), F1: 255}
                                    _ = v26
                                case -727:
                                    fmt.Println(v4)
                                case 32768:
                                    fmt.Println(v14)
                                    if true {
                                        if (false && (18 >= 2)) {
                                            fmt.Println(v7, v12, v12)
                                        } else {
                                            var v27 int8 = v5
                                            _ = v27
                                        }
                                        if ((((^(v7)) != v7) && true) || true) {
                                            var v28 uint32 = 2147483648
                                            fmt.Println(v3, v14)
                                            _ = v28
                                        }
                                        v5 = (v5 / ((v5) | 1))
                                    }
                                default:
                                    var v29 float32 = (((1.5 / (3.14159 + 1e10)) / -2.25) * float32(v25))
                                    _ = v29
                                }
                            }
                            _ = v25
                        default:
                            var v30 float32 = 0.1
                            _ = v30
                        }
                        v12 = (v12 ^ v12)
                    }
                    _ = fi23
                }
            }
            for fi31 := 0; fi31 < 3; fi31++ {
                v7 = 0
                var v32 uint16 = 2
                _ = fi31
                _ = v32
            }
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        v34 := GenS2{F0: "foo bar", F1: 2}
        v3 = v3
        _ = fi33
        _ = v34
    }
    v35 := GenS1{F0: 0, F1: 42, F2: 262}
    v3 += 2147483647
    v4 -= (v3 << ((v7) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
