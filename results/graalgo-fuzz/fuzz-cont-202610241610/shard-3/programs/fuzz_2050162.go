package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
type GenS2 struct { F0 int; F1 int64; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint32; F1 uint32 }
func genFn0() uint64 {
    var v6 uint = 0
    v7 := GenS1{F0: 309, F1: 18446744073709551615, F2: 32767}
    v8 := []int{}
    var v9 int16 = -569
    if (793 < 1) {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 int16 = v9
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: 1, F1: v6, F2: (v9 / (((-(v11))) | 1))}
                _ = fi12
                _ = v13
            }
            var v14 string = ("\n\t" + "a")
            _ = fi10
            _ = v11
            _ = v14
        }
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (-2.25 + 0.1)
    var v3 int32 = -1
    var v4 int = -9223372036854775808
    var v5 uint32 = 466
    if true {
        v15 := GenS2{F0: 1, F1: 42, F2: 858}
        fmt.Println(v2, v5, v2)
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := []int{893, 0, v4}
        var v18 string = "\n\t"
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if false {
            v5 -= (v5 & 1)
            v20 := []int{}
            v21 := GenS3{F0: 0, F1: 2}
            _ = v20
            _ = v21
        }
        _ = fi19
    }
    var v22 int16 = int16(v4)
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS4{F0: 335, F1: ((("foo bar" + ("" + "hello")) + "\n\t") + ("foo bar" + "hello"))}
        if false {
            if (int16(fi23) <= int16(v3)) {
                v25 := GenS5{F0: false, F1: 65535}
                v26 := GenS1{F0: (v5 + (^(v5))), F1: 0, F2: (v22 % (((v22 % (((v22 - 1)) | 1))) | 1))}
                for fi27 := 0; fi27 < 17; fi27++ {
                    var v28 uint64 = 2
                    _ = fi27
                    _ = v28
                }
                _ = v25
                _ = v26
            } else {
                for fi29 := 0; fi29 < 11; fi29++ {
                    if (((v2 <= v2) && false) || (false && false)) {
                        var v30 int = (v4 - (fi29 & (fi23 | 1)))
                        var v31 uint16 = 2
                        _ = v30
                        _ = v31
                    }
                    _ = fi29
                }
            }
            switch v4 {
            case -1:
                v32 := []int{fi23}
                _ = v32
            case -2147483648:
                var v33 uint16 = 255
                _ = v33
            case 0:
                switch (v4 / ((127) | 1)) {
                case 355:
                    var v34 int64 = -1
                    _ = v34
                case 9223372036854775807:
                    v3 -= v3
                }
            }
        }
        if ((1 != -12) || true) {
            fmt.Println(v22, v22, v22)
        } else {
            v35 := GenS6{F0: v3, F1: v5, F2: 42}
            v2 -= (((((0.1 + ((v2 + (float64(1e10) * float64((0.1 - (0.0 / 1.0))))) + (v2 / (v2 + -2.25)))) / v2) * ((v2 * v2) + v2)) + v2) + ((3.14159 / -2.25) - -2.25))
            _ = v35
        }
        _ = fi23
        _ = v24
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        fmt.Println(fi36, v22)
        _ = fi36
    }
    v4 = 2147483647
    for fi37 := 0; fi37 < 16; fi37++ {
        v38 := GenS2{F0: (v4 & 953), F1: 42, F2: 736}
        _ = fi37
        _ = v38
    }
    v39 := []int{v4}
    if false {
        v40 := GenS2{F0: v4, F1: -1, F2: 9223372036854775807}
        var v41 float32 = 0.1
        v5 -= v5
        _ = v40
        _ = v41
    }
    fmt.Println(v3, v4, v3)
    var v42 int16 = -933
    if true {
        fmt.Println(v5, v42, v2)
        switch "\n\t" {
        case "\n\t":
            v3 = (^(v3))
        case "foo bar":
            for fi43 := 0; fi43 < 2; fi43++ {
                v39 = append(v39, (fi43 ^ 9223372036854775807))
                for fi44 := 0; fi44 < 4; fi44++ {
                    if len(v39) > 0 {
                        v39[0] = (fi43 >> ((255) & 63))
                    }
                    _ = fi44
                }
                _ = fi43
            }
        case "hello":
            if len(v39) > 0 {
                v39[0] = (v4 << ((4294967295) & 63))
            }
        default:
            switch (v4 % ((v4) | 1)) {
            case 42:
                var v45 int8 = 0
                _ = v45
            default:
                switch v4 {
                case -165:
                    var v46 uint32 = v5
                    _ = v46
                default:
                    v22 += v22
                }
            }
        }
    } else {
        v5 += (v5 | v5)
        var v47 uint64 = 4294967295
        _ = v47
    }
    fmt.Println(v3, v5)
    for fi48 := 0; fi48 < 14; fi48++ {
        if false {
            if (-9223372036854775808 != -673) {
                var v49 uint16 = 42
                _ = len(v39)
                fmt.Println(v5)
                _ = v49
            } else {
                var v50 uint8 = 127
                v5 = 259
                _ = v50
            }
            v51 := GenS3{F0: uint64(v5), F1: 0}
            var v52 bool = (-95 < 77)
            _ = v51
            _ = v52
        }
        for fi53 := 0; fi53 < 20; fi53++ {
            switch v4 {
            case 4294967295:
                v39 = append(v39, (fi48 + (fi48 ^ fi53)))
            }
            if (-2.25 == 0.1) {
                var v54 int = 9223372036854775807
                switch "a" {
                case "hello":
                    var v55 uint32 = v5
                    _ = v55
                default:
                    v22 += (^(v42))
                }
                _ = v54
            } else {
                fmt.Println(v22)
                if (4294967295 > 365) {
                    v5 -= (v5 | (^(v5)))
                } else {
                    if (uint16(v5) >= 0) {
                        var v56 int8 = 102
                        v5 += (v5 / ((v5) | 1))
                        _ = len(v1)
                        _ = v56
                    } else {
                        var v57 uint8 = 255
                        _ = v57
                    }
                }
            }
            _ = fi53
        }
        for fi58 := 0; fi58 < 16; fi58++ {
            v59 := GenS7{F0: 2147483648, F1: (v5 + 1)}
            v60 := GenS6{F0: (^(v3)), F1: 51, F2: 282}
            v42 = v22
            _ = fi58
            _ = v59
            _ = v60
        }
        _ = fi48
    }
    v61 := GenS2{F0: (v4 - v4), F1: 9223372036854775807, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v61
}
