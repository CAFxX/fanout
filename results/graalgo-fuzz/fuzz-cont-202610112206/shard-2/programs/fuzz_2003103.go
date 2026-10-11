package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint32 }
func genFn0() bool {
    var v6 int64 = -9223372036854775808
    switch -528 {
    case 320:
        var v7 uint16 = 2
        _ = v7
    }
    _ = v6
    return (((300 == -1) || (42 <= 42)) || (-1 >= -960))
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483647, 65535, 475}
    v4 := []uint32{42, 397, 147}
    v5 := []uint32{2}
    v3 = append(v3, 343)
    v1 = v1
    for fi8 := 0; fi8 < 1; fi8++ {
        v1 = fi8
        v9 := GenS1{F0: 0, F1: (!false), F2: (v1 ^ v1)}
        var v10 uint8 = 27
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        v1 -= v1
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: 2147483647, F1: (1e10 < 1e10), F2: (^(fi11))}
        var v13 uint64 = 600
        var v14 bool = true
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    if true {
        switch (v1 ^ (v1 ^ 164)) {
        case 128:
            if (("a" + "") == "foo bar") {
                v1 -= v1
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 float64 = 1e10
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 float32 = (((float32(((1e10 / -2.25) + 0.0)) / (0.1 / 3.14159)) + 0.1) * (0.0 / (((0.1 - 1e10) * 3.14159) / (1e10 / (((0.1 * (1e10 * float32(((v16 / (v16 + v16)) * v16)))) - -2.25) + 1e10)))))
                        _ = fi17
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
                var v19 string = "a"
                _ = v19
            } else {
                var v20 uint32 = 2147483647
                _ = v20
            }
        }
    }
    v21 := GenS1{F0: 42, F1: true, F2: -267}
    if (false && ((v1 / ((v1) | 1)) > 57)) {
        var v22 int64 = 1
        _ = v22
    }
    _ = len(v2)
    var v23 int16 = 1000
    v23 = v23
    switch "a" {
    case "foo bar":
        for fi24 := 0; fi24 < 17; fi24++ {
            v23 = -1
            switch (^(v1)) {
            case 42:
                v1 -= 128
            case 2147483647:
                var v25 uint16 = 1
                _ = v25
            case -2147483648:
                for fi26 := 0; fi26 < 5; fi26++ {
                    _ = len(v3)
                    _ = fi26
                }
            }
            _ = fi24
        }
    case "a":
        if true {
            v27 := GenS1{F0: -363, F1: false, F2: 803}
            _ = v27
        } else {
            var v28 float32 = float32(3.14159)
            _ = v28
        }
    default:
        var v29 int16 = (v23 | v23)
        _ = v29
    }
    v30 := GenS1{F0: 189, F1: false, F2: v1}
    v1 = 389
    for fi31 := 0; fi31 < 20; fi31++ {
        if ((-(v1)) < v1) {
            _ = len(v2)
            fmt.Println(v1, v23)
            switch (fi31 << ((319) & 63)) {
            case 750:
                for fi32 := 0; fi32 < 11; fi32++ {
                    v1 = (v1 / (((v1 >> ((18446744073709551615) & 63))) | 1))
                    _ = fi32
                }
            case 9223372036854775807:
                var v33 uint32 = 383
                _ = v33
            case -2147483648:
                if true {
                    if (!((2 < 2) && (!(true || (1 >= 288))))) {
                        v34 := GenS1{F0: -2147483648, F1: false, F2: fi31}
                        _ = v34
                    } else {
                        if ((625 != -9223372036854775808) || (18446744073709551615 < 399)) {
                            for fi35 := 0; fi35 < 3; fi35++ {
                                fmt.Println(v23, v1, fi35)
                                _ = fi35
                            }
                            v36 := GenS2{F0: 4294967295, F1: 255, F2: 2147483648}
                            _ = v36
                        }
                        switch v1 {
                        case -2147483648:
                            fmt.Println(fi31)
                        default:
                            fmt.Println(fi31)
                        }
                    }
                }
            default:
                v23 -= (v23 >> ((1) & 63))
            }
        } else {
            var v37 int16 = v23
            _ = v37
        }
        _ = fi31
    }
    if false {
        v1 -= 128
        v38 := GenS1{F0: 42, F1: true, F2: (v1 * 4294967295)}
        _ = v38
    }
    var v39 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v5
}
