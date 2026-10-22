package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 int16, p1 bool) uint {
    fmt.Println(p0)
    v5 := []int{2147483647, -2147483648, 127}
    fmt.Println(p1, p0)
    var v6 int64 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int16) uint16 {
    var v7 int64 = -1
    var v8 uint8 = 204
    _ = p0
    _ = v7
    _ = v8
    return 496
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 42
    v3 := make(map[int]int)
    v4 := make(map[string]int)
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(fi9)
            if ((-1 == (fi10 | (-(fi9)))) && ((!true) && false)) {
                var v11 string = "\n\t"
                v2 = v2
                _ = v11
            }
            switch fi10 {
            case 4294967295:
                if (!true) {
                    v2 -= v2
                    fmt.Println(v2, fi9)
                } else {
                    v3[127] = (fi9 * -9223372036854775808)
                    v12 := GenS1{F0: ("foo bar" + ""), F1: 42, F2: 2}
                    v13 := []int{(fi9 ^ (fi9 << ((255) & 63))), (fi10 >> ((511) & 63))}
                    _ = v12
                    _ = v13
                }
                fmt.Println(fi10, fi9, fi10)
            case 32767:
                v2 = v2
            case 32768:
                if true {
                    if false {
                        if (18446744073709551615 <= 1) {
                            v2 += v2
                        } else {
                            var v14 int8 = -117
                            _ = v14
                        }
                        switch (fi10 & (fi10 ^ fi9)) {
                        case 0:
                            v2 = 884
                        case -502:
                            v2 = 1
                        case 32768:
                            var v15 int64 = 0
                            _ = v15
                        }
                        fmt.Println(v2, fi10)
                    }
                } else {
                    v1[fi10] = fi9
                    if false {
                        v16 := GenS2{F0: 216}
                        _ = v16
                    } else {
                        var v17 int = -916
                        var v18 uint64 = 479
                        fmt.Println(v18)
                        _ = v17
                        _ = v18
                    }
                    switch "hello" {
                    case "foo bar":
                        v19 := GenS2{F0: 255}
                        _ = v19
                    case "\n\t":
                        var v20 string = ("foo bar" + "")
                        _ = v20
                    default:
                        v2 += v2
                    }
                }
            default:
                switch "\n\t" {
                case "foo bar":
                    var v21 int64 = 9223372036854775807
                    _ = v21
                case "a":
                    var v22 string = ""
                    _ = v22
                default:
                    v2 = 2147483647
                }
            }
            _ = fi10
        }
        v2 -= v2
        _ = fi9
    }
    var v23 uint8 = 158
    var v24 int8 = 127
    v24 -= v24
    var v25 uint64 = 4294967295
    v2 -= v2
    switch 352 {
    case 32768:
        fmt.Println(v25)
    case -343:
        v23 -= (v23 + (^(v23)))
    case 2147483647:
        v26 := GenS2{F0: v23}
        _ = v26
    default:
        var v27 uint64 = (v25 >> ((4294967295) & 63))
        _ = v27
    }
    fmt.Println(v23, v25)
    switch 1 {
    case -661:
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 uint64 = (v25 % ((322) | 1))
            _ = fi28
            _ = v29
        }
    case -1:
        v30 := GenS3{F0: 128, F1: 833}
        _ = v30
    }
    v31 := v1[4294967295]
    v24 = v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
