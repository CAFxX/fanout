package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 int32) string {
    v4 := []int{2147483647, 1}
    var v5 bool = ((!true) && false)
    p0 = 65535
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ("hello" + (((("foo bar" + "foo bar") + ("\n\t" + "\n\t")) + "") + "a"))
}

func main() {
    var v1 int8 = -96
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    fmt.Println(v1)
    fmt.Println(v1)
    fmt.Println(v1)
    v3[9] = 127
    for fi6 := 0; fi6 < 19; fi6++ {
        switch "" {
        case "\n\t":
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint16 = 1
                _ = fi7
                _ = v8
            }
        }
        _ = fi6
    }
    v1 = (v1 % ((0) | 1))
    _ = len(v3)
    switch 171 {
    case 32767:
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(v1, v1)
            v1 -= (v1 & v1)
            _ = fi9
        }
    case 200:
        if (65535 == 0) {
            v10 := []int{-297}
            v11 := GenS1{F0: v1, F1: -9223372036854775808, F2: 2}
            if (32767 == 0) {
                v1 = v1
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v12 uint8 = 42
                _ = v12
            }
            _ = v10
            _ = v11
        } else {
            v13 := GenS1{F0: 3, F1: 0, F2: 1}
            _ = v13
        }
    default:
        v14 := GenS1{F0: v1, F1: 833, F2: 48}
        _ = v14
    }
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 7; fi16++ {
        v1 = (v1 | (v1 * v1))
        _ = fi16
    }
    switch "" {
    case "a":
        if true {
            for fi17 := 0; fi17 < 5; fi17++ {
                switch (fi17 << ((255) & 63)) {
                case 4294967295:
                    if (false && ((fi17 << ((42) & 63)) < (fi17 % ((fi17) | 1)))) {
                        switch "a" {
                        case "foo bar":
                            v18 := GenS1{F0: -86, F1: (v15 + (v15 % (((^(v15))) | 1))), F2: 42}
                            _ = v18
                        case "\n\t":
                            v15 = v15
                        case "":
                            if false {
                                v19 := v3[1]
                                v1 -= -6
                                _ = v19
                            }
                        default:
                            for fi20 := 0; fi20 < 6; fi20++ {
                                v21 := []int{fi17, -649, (fi20 % ((fi17) | 1))}
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                }
                fmt.Println(v1)
                switch "foo bar" {
                case "a":
                    _ = len(v2)
                }
                _ = fi17
            }
        }
    case "\n\t":
        fmt.Println(v1, v1)
    case "foo bar":
        var v22 int64 = 315
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v24 := GenS2{F0: -1}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            v15 = 42
            v27 := GenS2{F0: -439}
            _ = fi26
            _ = v27
        }
        fmt.Println(v1)
        v1 -= 45
        _ = fi25
    }
    v28 := []int{-342, 9223372036854775807, 889, -1}
    _ = len(v2)
    fmt.Println(v15, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
}
