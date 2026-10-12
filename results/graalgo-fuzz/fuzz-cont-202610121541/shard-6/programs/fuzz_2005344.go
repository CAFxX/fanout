package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }
type GenS2 struct { F0 int32; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 float64 = ((-2.25 - (0.1 + 3.14159)) + -2.25)
    var v5 int = 32767
    v6 := []string{(("" + "") + "hello"), "foo bar"}
    v1 -= v1
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            if len(v6) > 0 {
                v6[0] = ("\n\t" + "\n\t")
            }
            _ = fi8
        }
        _ = fi7
    }
    if true {
        fmt.Println(v5)
        if false {
            fmt.Println(v1)
            for fi9 := 0; fi9 < 13; fi9++ {
                if true {
                    if ("\n\t" < ((("hello" + (("a" + "a") + ((("a" + "foo bar") + "foo bar") + ""))) + "foo bar") + "")) {
                        v4 = v4
                        var v10 uint = 1
                        _ = v10
                    }
                } else {
                    switch fi9 {
                    case 9223372036854775807:
                        var v11 int64 = v1
                        _ = v11
                    case -879:
                        v12 := GenS1{F0: 0, F1: false, F2: 941}
                        _ = v12
                    case 942:
                        v4 += 1e10
                    }
                    fmt.Println(fi9, fi9, v4)
                }
                _ = fi9
            }
            switch 42 {
            case 32767:
                var v13 uint16 = 42
                _ = v13
            default:
                v14 := GenS1{F0: 797, F1: (0 > 13), F2: -1}
                _ = v14
            }
        } else {
            fmt.Println(v4, v5)
        }
        fmt.Println(v4, v4, v1)
    } else {
        v15 := GenS2{F0: -1, F1: v1, F2: 0}
        _ = v15
    }
    if false {
        v1 -= (v1 & 0)
        var v16 int32 = 0
        v16 -= v16
        _ = v16
    }
    var v17 int = v5
    if true {
        v18 := GenS3{F0: -1, F1: -128, F2: 2147483648}
        switch 124 {
        case 42:
            v17 = 378
        case 1:
            fmt.Println(v4, v17)
        case 0:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            switch 42 {
            case 1:
                v19 := GenS1{F0: 2147483648, F1: true, F2: -253}
                _ = v19
            case 719:
                v20 := GenS4{F0: 489}
                _ = v20
            default:
                v17 -= 32767
            }
        }
        _ = v18
    } else {
        v21 := GenS1{F0: 42, F1: (18446744073709551615 >= 1), F2: 486}
        v6 = append(v6, "foo bar")
        v17 = v5
        _ = v21
    }
    if true {
        var v22 int64 = (v1 * (v1 ^ 42))
        _ = v22
    }
    var v23 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
