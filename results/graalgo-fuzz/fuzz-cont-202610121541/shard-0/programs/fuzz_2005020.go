package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []uint32{1, 4294967295, 493}
    v2 := []uint32{}
    var v3 int = 2147483647
    var v4 int16 = 0
    fmt.Println(v3, v3)
    switch "hello" {
    case "\n\t":
        fmt.Println(v4, v4)
    case "foo bar":
        if false {
            v4 = (v4 >> ((1) & 63))
            switch (^(v3)) {
            case 0:
                var v5 int = 2147483647
                _ = v5
            case -421:
                var v6 uint64 = 977
                _ = v6
            case 9223372036854775807:
                if true {
                    for fi7 := 0; fi7 < 11; fi7++ {
                        fmt.Println(fi7)
                        _ = fi7
                    }
                } else {
                    var v8 int64 = 71
                    if false {
                        v4 += -32768
                        fmt.Println(v4, v4, v8)
                    }
                    for fi9 := 0; fi9 < 12; fi9++ {
                        if true {
                            switch (fi9 + v3) {
                            case 4294967295:
                                fmt.Println(v4)
                            case 42:
                                v4 -= v4
                            case -1:
                                v10 := GenS1{F0: "\n\t", F1: v8}
                                _ = v10
                            default:
                                v8 = v8
                            }
                            if (int16(fi9) >= v4) {
                                v8 -= 1
                                v11 := GenS2{F0: 157}
                                if true {
                                    v3 = -1
                                } else {
                                    fmt.Println(v4, v3)
                                }
                                _ = v11
                            } else {
                                v12 := GenS1{F0: "", F1: 1}
                                v13 := GenS1{F0: ("" + ("a" + "\n\t")), F1: 546}
                                v14 := GenS2{F0: 2}
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                            v1 = append(v1, 757)
                        } else {
                            if (920 != 1) {
                                v15 := GenS2{F0: 1}
                                var v16 float64 = 0.1
                                _ = v15
                                _ = v16
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v17 := GenS1{F0: "\n\t", F1: v8}
                            _ = v17
                        }
                        var v18 int8 = 100
                        _ = fi9
                        _ = v18
                    }
                    _ = v8
                }
            default:
                v3 = v3
            }
        } else {
            if true {
                if true {
                    var v19 uint = 65535
                    v3 -= (v3 / (((^(v3))) | 1))
                    _ = v19
                }
            }
            var v20 uint16 = 32767
            v3 = (v3 - v3)
            _ = v20
        }
    default:
        switch (-(v3)) {
        case -268:
            var v21 uint8 = 2
            _ = v21
        default:
            v3 = (v3 - (v3 / ((v3) | 1)))
        }
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    v3 = 494
    for fi22 := 0; fi22 < 19; fi22++ {
        v3 = (v3 * (-(v3)))
        _ = fi22
    }
    if false {
        v3 = v3
        if ((true || (false && ((100 >= 100) || ((703 < 1) || false)))) || false) {
            v4 -= 693
            var v23 int = 128
            _ = v23
        } else {
            v2 = append(v2, 65535)
            var v24 uint64 = 18446744073709551615
            _ = v24
        }
        var v25 uint64 = 0
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 bool = (false || true)
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: ("foo bar" + "a"), F1: 9223372036854775807}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
