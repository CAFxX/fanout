package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint; F1 string; F2 int }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 int32, p1 uint32) uint {
    v5 := []int{-1, 128, 9223372036854775807}
    v6 := GenS1{F0: 236}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint = 65535
    var v2 uint32 = 628
    v3 := []uint32{v2, (^(v2)), v2}
    var v4 bool = true
    if (9223372036854775807 != 0) {
        v7 := GenS2{F0: "", F1: v4}
        _ = v7
    } else {
        v1 -= v1
        var v8 int8 = 100
        _ = v8
    }
    fmt.Println(v4)
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "" {
        case "":
            if v4 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v10 := GenS2{F0: "", F1: v4}
                v1 = (v1 + (v1 * 1))
                _ = v10
            }
        default:
            switch ("\n\t" + ("\n\t" + "hello")) {
            case "":
                for fi11 := 0; fi11 < 12; fi11++ {
                    if (128 >= 109) {
                        v12 := GenS3{F0: v1, F1: ("foo bar" + "foo bar"), F2: 32768}
                        v4 = (!(-2147483648 >= 2147483647))
                        var v13 uint8 = 128
                        _ = v12
                        _ = v13
                    } else {
                        switch "hello" {
                        case "a":
                            switch "hello" {
                            case "":
                                v2 = 2147483647
                            case "hello":
                                v2 = 4294967295
                            case "foo bar":
                                var v14 int8 = 1
                                _ = v14
                            }
                        case "\n\t":
                            var v15 int64 = -9223372036854775808
                            _ = v15
                        }
                    }
                    v4 = (((v2 & 458) >= v2) || v4)
                    v2 -= (v2 ^ (^(v2)))
                    if (!v4) {
                        v16 := GenS4{F0: 822, F1: v4}
                        _ = v16
                    }
                    fmt.Println(v1, fi9, fi11)
                    _ = fi11
                }
            case "hello":
                fmt.Println(v2, fi9)
            case "a":
                if len(v3) > 0 {
                    v3[0] = (^(v2))
                }
            default:
                v2 += (v2 << (((v1 / (((v1 << ((v1) & 63))) | 1))) & 63))
            }
        }
        _ = fi9
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint16 = 32767
        _ = fi17
        _ = v18
    }
    var v19 float64 = (0.1 / 1e10)
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint = 0
        _ = fi20
        _ = v21
    }
    if v4 {
        v1 = v1
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint32 = v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 255}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (v2 < v2) {
        var v25 uint8 = 127
        _ = v25
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS2{F0: "", F1: (1 <= 2147483647)}
            _ = fi26
            _ = v27
        }
        v3 = append(v3, v2)
    }
    v28 := GenS1{F0: 2}
    var v29 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
