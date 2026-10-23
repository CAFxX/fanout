package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 uint32 }
type GenS5 struct { F0 string; F1 int32 }
func genFn0(p0 uint32, p1 uint8, p2 bool) string {
    p2 = ((p1 | p1) <= (p1 >> ((0) & 63)))
    fmt.Println(p1, p2)
    v7 := GenS1{F0: 908}
    var v8 uint32 = (p0 + 518)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 uint64, p1 int, p2 uint16) uint8 {
    v9 := GenS1{F0: 128}
    v10 := GenS1{F0: (p1 >> ((42) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 113
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]bool)
    var v3 float32 = 0.1
    var v4 uint16 = 0
    v5 := make(map[string]int)
    v6 := []uint32{2147483648, 999, 42}
    var v11 int = 373
    var v12 int8 = v1
    v1 = (v1 >> ((151) & 63))
    v13 := GenS1{F0: (v11 - (^(v11)))}
    v11 = 9223372036854775807
    v12 = v12
    if (!(!true)) {
        if (1000 < 1000) {
            delete(v5, "")
            v1 = (v12 - v1)
            v14 := GenS1{F0: (v11 + v11)}
            _ = v14
        }
    }
    if true {
        v11 -= (-(v11))
        v15 := GenS2{F0: v12, F1: 18446744073709551615, F2: 9223372036854775807}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 int32 = 2147483647
            var v18 uint64 = 195
            if (!false) {
                switch 1 {
                case 9223372036854775807:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        switch "hello" {
                        case "\n\t":
                            var v20 float32 = (v3 / (v3 / 0.1))
                            _ = v20
                        case "foo bar":
                            v21 := GenS1{F0: fi19}
                            _ = v21
                        case "a":
                            v3 = 3.14159
                        default:
                            var v22 string = (("foo bar" + "") + "foo bar")
                            _ = v22
                        }
                        if false {
                            _ = len(v6)
                            if (!true) {
                                if ((2 > 1) || false) {
                                    var v23 float64 = 0.1
                                    _ = v23
                                }
                            } else {
                                var v24 int16 = 32767
                                _ = v24
                            }
                        } else {
                            v25 := GenS3{F0: 906}
                            if (!false) {
                                var v26 uint32 = 65535
                                if (!true) {
                                    v27 := GenS4{F0: 249, F1: 42, F2: (^(v26))}
                                    v17 -= (v17 * int32(v12))
                                    v11 += (v11 / (((fi19 / (((fi16 ^ (^(fi16)))) | 1))) | 1))
                                    _ = v27
                                }
                                _ = v26
                            }
                            delete(v2, 4294967295)
                            _ = v25
                        }
                        _ = fi19
                    }
                case 1:
                    if len(v6) > 0 {
                        v6[0] = 469
                    }
                case -2147483648:
                    fmt.Println(v1)
                }
                _ = len(v5)
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        v4 -= (v4 * (v4 | (v4 % ((v4) | 1))))
        v28 := GenS5{F0: "", F1: -1}
        _ = v28
    }
    v29 := GenS2{F0: (v12 + v1), F1: 504, F2: 42}
    var v30 int32 = -375
    v3 = (float32(1e10) + float32((3.14159 + 0.0)))
    v31 := v2[384]
    v30 = (v30 * v30)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
