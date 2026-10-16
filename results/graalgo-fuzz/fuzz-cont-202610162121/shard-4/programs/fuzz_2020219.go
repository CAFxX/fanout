package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64 }
func genFn0(p0 uint32, p1 uint32, p2 uint64) int16 {
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int8 = -82
            switch (-(fi7)) {
            case 0:
                v10 := GenS1{F0: v6, F1: 255, F2: 928}
                _ = v10
            case 257:
                var v11 float64 = 1.5
                _ = v11
            case -2147483648:
                var v12 uint32 = 159
                _ = v12
            default:
                switch "hello" {
                case "":
                    fmt.Println(fi7)
                case "\n\t":
                    v13 := GenS2{F0: 696, F1: (v9 | (v9 | (^(v9))))}
                    _ = v13
                }
            }
            _ = fi8
            _ = v9
        }
        var v14 float64 = ((-2.25 + float64(-2.25)) - ((-2.25 / -2.25) * (float64(float32(float64(-2.25))) * ((float64((3.14159 - (float32(((0.0 + ((1.5 * 3.14159) - (1.5 * float64((1e10 * 0.1))))) / (0.0 - 1.5))) + 3.14159))) / 1.5) + float64((-2.25 * 3.14159))))))
        _ = fi7
        _ = v14
    }
    var v15 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return -711
}
func genFn1(p0 bool, p1 int64, p2 bool) uint {
    p1 = (p1 + p1)
    var v16 uint8 = 42
    switch 1 {
    case -31:
        switch (("a" + "") + "\n\t") {
        case "":
            if p2 {
                p1 += p1
                v17 := GenS1{F0: (v16 / ((v16) | 1)), F1: 42, F2: 310}
                _ = v17
            }
        case "foo bar":
            p1 = -9223372036854775808
        }
    }
    if ((941 != int(p1)) && (-667 >= 0)) {
        var v18 uint16 = 32767
        var v19 int8 = 0
        var v20 int = -830
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: (^(p1)), F1: (v16 * (v16 & v16)), F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v21
    return 748
}

func main() {
    var v1 uint8 = 217
    var v2 int = 102
    v3 := []string{}
    var v4 uint16 = 827
    var v5 uint32 = uint32(v1)
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 int32 = -1
        _ = fi22
        _ = v23
    }
    switch (v2 * v2) {
    case 202:
        v24 := GenS4{F0: 0, F1: -9223372036854775808}
        _ = v24
    case 636:
        var v25 int8 = 100
        _ = v25
    }
    v26 := GenS3{F0: 42, F1: (v1 * (v1 ^ (v1 ^ (v1 << ((65535) & 63))))), F2: 185}
    v27 := GenS4{F0: 597, F1: 734}
    switch "a" {
    case "\n\t":
        if (!false) {
            v28 := GenS5{F0: 42}
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 string = "\n\t"
                switch (-(v2)) {
                case 1:
                    v31 := GenS3{F0: -9223372036854775808, F1: uint8(v4), F2: v5}
                    _ = v31
                default:
                    v2 -= (^(v2))
                }
                v4 += (^(v4))
                _ = fi29
                _ = v30
            }
            var v32 int16 = -538
            _ = v28
            _ = v32
        } else {
            var v33 uint32 = v5
            for fi34 := 0; fi34 < 1; fi34++ {
                v33 = (v33 | 843)
                v5 -= (v5 << ((576) & 63))
                _ = fi34
            }
            _ = v33
        }
    case "a":
        fmt.Println(v4)
    default:
        for fi35 := 0; fi35 < 13; fi35++ {
            switch "a" {
            case "\n\t":
                switch v2 {
                case 128:
                    v36 := GenS6{F0: v1}
                    _ = v36
                case 4294967295:
                    switch "foo bar" {
                    case "":
                        for fi37 := 0; fi37 < 17; fi37++ {
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            _ = fi37
                        }
                    case "\n\t":
                        var v38 int64 = 1
                        _ = v38
                    }
                default:
                    v4 += (v4 | v4)
                }
            default:
                fmt.Println(v1, v5)
            }
            _ = fi35
        }
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS7{F0: v5, F1: v1}
        v2 -= (fi39 ^ fi39)
        _ = fi39
        _ = v40
    }
    if (("" + "foo bar") >= ("\n\t" + "hello")) {
        if true {
            for fi41 := 0; fi41 < 16; fi41++ {
                switch (v2 - (-(v2))) {
                case -814:
                    v5 -= v5
                case 0:
                    v2 = fi41
                default:
                    switch "foo bar" {
                    case "foo bar":
                        if (!(!((v4 % ((v4) | 1)) > (v4 ^ 32767)))) {
                            if true {
                                var v42 uint = 2
                                fmt.Println(v2)
                                _ = v42
                            }
                            var v43 uint8 = (v1 % ((v1) | 1))
                            v44 := GenS2{F0: uint16(v5), F1: -1}
                            _ = v43
                            _ = v44
                        } else {
                            v5 += 42
                            v45 := GenS4{F0: genFn0((v5 / ((v5) | 1)), 1, 349), F1: -1}
                            _ = v45
                        }
                    }
                }
                _ = fi41
            }
        }
        if (1000 <= 1000) {
            _ = len(v3)
            v1 = 128
            for fi46 := 0; fi46 < 5; fi46++ {
                v47 := GenS6{F0: v1}
                var v48 int64 = -1
                _ = fi46
                _ = v47
                _ = v48
            }
        }
    }
    v49 := GenS7{F0: 2147483648, F1: v1}
    var v50 uint64 = 2
    for fi51 := 0; fi51 < 6; fi51++ {
        v52 := GenS5{F0: 42}
        _ = fi51
        _ = v52
    }
    var v53 uint = 4294967295
    if false {
        if ((0.0 > (((((-2.25 * 0.1) / 1.5) + float64(0.1)) + ((0.0 * 0.0) - 0.0)) - 1.5)) && true) {
            for fi54 := 0; fi54 < 2; fi54++ {
                v55 := GenS8{F0: (^(v50)), F1: 1}
                _ = fi54
                _ = v55
            }
            switch -845 {
            case 127:
                v56 := GenS9{F0: -805}
                _ = v56
            case -2147483648:
                var v57 int16 = -36
                _ = v57
            default:
                var v58 float32 = float32((-2.25 / (float64(0.1) + (-2.25 + float64(1.5)))))
                _ = v58
            }
        } else {
            v1 = v1
            switch (^(v2)) {
            case 32768:
                switch v2 {
                case -2147483648:
                    if false {
                        if (v1 >= (v1 | (v1 >> (((v53 - v53)) & 63)))) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                fmt.Println(v2, v1, v2)
                                var v59 uint32 = uint32(v4)
                                if (1000 == -310) {
                                    _ = len(v3)
                                } else {
                                    fmt.Println(v2)
                                    v1 = (v1 + 87)
                                }
                                _ = v59
                            } else {
                                switch "\n\t" {
                                case "a":
                                    var v60 float32 = (float32(0.0) * -2.25)
                                    _ = v60
                                case "foo bar":
                                    for fi61 := 0; fi61 < 1; fi61++ {
                                        v50 -= (v50 + v50)
                                        v4 -= v4
                                        _ = fi61
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            switch "" {
                            case "a":
                                var v62 float32 = 0.1
                                _ = v62
                            }
                        } else {
                            v63 := GenS10{F0: (!false), F1: 0}
                            v3 = append(v3, ("\n\t" + "hello"))
                            _ = v63
                        }
                        fmt.Println(v1)
                    }
                case 127:
                    if true {
                        v53 = (^(v53))
                    }
                }
            case 2147483647:
                var v64 float64 = 0.1
                _ = v64
            case 1000:
                v1 -= (v1 % (((^(v1))) | 1))
            default:
                for fi65 := 0; fi65 < 11; fi65++ {
                    for fi66 := 0; fi66 < 10; fi66++ {
                        var v67 int32 = -1
                        v67 += 2147483647
                        if true {
                            switch (fi66 * fi65) {
                            case 1:
                                if false {
                                    var v68 float32 = 1.5
                                    _ = v68
                                } else {
                                    v69 := GenS11{F0: -1}
                                    _ = v69
                                }
                            case -815:
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            case 4294967295:
                                v70 := GenS10{F0: true, F1: -1}
                                _ = v70
                            }
                            _ = len(v3)
                            if true {
                                var v71 int32 = (^(v67))
                                _ = v71
                            }
                            fmt.Println(v50, fi66)
                            fmt.Println(v5, v5, v5)
                        } else {
                            switch "a" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            case "a":
                                if false {
                                    var v72 uint32 = (v5 * v5)
                                    _ = v72
                                }
                            }
                            v53 += (^(v53))
                            var v73 uint64 = 2
                            v2 += (-(v2))
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v2 = fi65
                            }
                            _ = len(v3)
                            if false {
                                var v74 int16 = -32768
                                _ = v74
                            }
                            _ = v73
                        }
                        _ = fi66
                        _ = v67
                    }
                    _ = fi65
                }
            }
            if (!true) {
                switch v2 {
                case 9223372036854775807:
                    if (v1 >= 85) {
                        fmt.Println(v4)
                        for fi75 := 0; fi75 < 2; fi75++ {
                            var v76 uint32 = 640
                            _ = fi75
                            _ = v76
                        }
                    }
                case -1:
                    _ = len(v3)
                case 1:
                    fmt.Println(v50, v53, v50)
                }
                fmt.Println(v53)
            } else {
                var v77 int64 = 382
                _ = v77
            }
        }
    }
    var v78 float64 = -2.25
    for fi79 := 0; fi79 < 11; fi79++ {
        v80 := GenS8{F0: (v50 & (v50 - (^(v50)))), F1: 1000}
        v81 := GenS12{F0: (v50 | v50)}
        fmt.Println(v2)
        _ = fi79
        _ = v80
        _ = v81
    }
    var v82 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v78:", v78)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
    _ = v78
    _ = v82
}
