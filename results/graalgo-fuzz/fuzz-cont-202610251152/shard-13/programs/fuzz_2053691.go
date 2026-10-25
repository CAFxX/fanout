package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int; F1 int; F2 int8 }
type GenS5 struct { F0 int }
func genFn0() int32 {
    var v7 int16 = 0
    switch "hello" {
    case "foo bar":
        v7 -= (^(v7))
    case "a":
        switch -894 {
        case 42:
            v7 = (v7 / ((0) | 1))
        default:
            v7 -= v7
        }
    default:
        fmt.Println(v7)
    }
    v7 += 1
    var v8 uint = 255
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 bool = true
    v2 := []string{(("" + ("" + "a")) + "a"), ""}
    var v3 int = -9223372036854775808
    var v4 int8 = 10
    var v5 uint = 813
    v6 := make(map[int]bool)
    switch "\n\t" {
    case "hello":
        switch "\n\t" {
        case "":
            switch v3 {
            case 2147483647:
                fmt.Println(v4)
            case -606:
                if false {
                    if (((((0.0 * 0.0) - ((0.0 - (0.1 + -2.25)) - -2.25)) * 0.0) + (1.5 / -2.25)) > 3.14159) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi9 := 0; fi9 < 11; fi9++ {
                            switch "a" {
                            case "\n\t":
                                var v10 int16 = -1
                                _ = v10
                            }
                            _ = fi9
                        }
                    } else {
                        if ((65535 >= 2147483648) || v1) {
                            var v11 int = int(v5)
                            _ = v11
                        } else {
                            var v12 bool = v1
                            _ = v12
                        }
                    }
                    var v13 int64 = 9223372036854775807
                    v1 = v1
                    _ = v13
                } else {
                    v14 := GenS1{F0: v1}
                    _ = v14
                }
            default:
                var v15 uint8 = 2
                _ = v15
            }
        }
    default:
        switch v3 {
        case -464:
            for fi16 := 0; fi16 < 15; fi16++ {
                v2 = append(v2, "")
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := GenS2{F0: v5, F1: v5}
                    v4 = (v4 * 127)
                    v19 := GenS2{F0: 41, F1: 65535}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := GenS2{F0: v5, F1: v5}
                _ = fi16
                _ = v20
            }
        case -894:
            v1 = (((1.5 / ((-2.25 / 1.0) - 1e10)) - (0.0 * -2.25)) == -2.25)
        }
    }
    switch "" {
    case "\n\t":
        v21 := GenS2{F0: (v5 * 1), F1: v5}
        _ = v21
    }
    v22 := GenS2{F0: 65535, F1: (v5 << ((4294967295) & 63))}
    var v23 uint = 4294967295
    v23 = v5
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS3{F0: 1000, F1: (-(v4))}
        _ = fi24
        _ = v25
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v26 uint16 = 0
    var v27 int8 = 1
    v3 += (v3 << ((v5) & 63))
    v23 = v23
    if false {
        v26 = 36
        if (!(0 >= 653)) {
            _ = len(v2)
            v28 := GenS4{F0: (v3 - (-(v3))), F1: (v3 - (^(v3))), F2: (v4 & v4)}
            v2 = append(v2, "a")
            _ = v28
        }
        v23 = uint(v26)
    }
    var v29 int64 = -290
    fmt.Println(v23, v23, v5)
    fmt.Println(v26, v1)
    switch "\n\t" {
    case "a":
        var v30 int32 = -2147483648
        _ = v30
    case "":
        var v31 uint32 = uint32(v5)
        _ = v31
    default:
        _ = len(v2)
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        switch "a" {
        case "foo bar":
            v33 := GenS5{F0: (fi32 / ((fi32) | 1))}
            _ = v33
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
