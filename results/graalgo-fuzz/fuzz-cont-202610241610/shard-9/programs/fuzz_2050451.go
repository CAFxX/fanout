package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint8; F1 int8; F2 uint }
func genFn0(p0 int8, p1 int64) int16 {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(p1, p1, p1)
        v7 := GenS1{F0: (-(p0)), F1: 127, F2: 1}
        if true {
            var v8 float32 = -2.25
            p1 += 0
            v9 := GenS1{F0: (p0 % (((^(p0))) | 1)), F1: 26, F2: (p0 << ((1) & 63))}
            _ = v8
            _ = v9
        } else {
            var v10 uint64 = 18446744073709551615
            fmt.Println(v10, v10)
            for fi11 := 0; fi11 < 8; fi11++ {
                switch "hello" {
                case "foo bar":
                    if (!((32767 != 2) && false)) {
                        var v12 uint8 = 42
                        if (!(true || ((260 > 520) && (!(!(v12 <= (^(v12)))))))) {
                            if false {
                                switch "" {
                                case "\n\t":
                                    var v13 uint32 = 890
                                    _ = v13
                                case "foo bar":
                                    switch -9223372036854775808 {
                                    case 425:
                                        v12 += (v12 & v12)
                                    case 128:
                                        p1 = 0
                                    default:
                                        v14 := GenS1{F0: 127, F1: (^(v12)), F2: (p0 & p0)}
                                        _ = v14
                                    }
                                case "hello":
                                    p1 = 42
                                    if false {
                                        v15 := GenS1{F0: (p0 >> ((2) & 63)), F1: 155, F2: (^(p0))}
                                        var v16 int = fi6
                                        v17 := GenS2{F0: 42, F1: 42}
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    } else {
                                        fmt.Println(fi11, v10, v10)
                                        var v18 uint16 = 0
                                        var v19 uint16 = v18
                                        _ = v18
                                        _ = v19
                                    }
                                    if true {
                                        p0 -= p0
                                    }
                                }
                                p0 -= 100
                            }
                            switch 864 {
                            case 32768:
                                if (2 != 18446744073709551615) {
                                    v20 := GenS1{F0: p0, F1: (v12 / ((255) | 1)), F2: int8(v10)}
                                    p1 += (p1 & (p1 % ((p1) | 1)))
                                    _ = v20
                                }
                            }
                            if true {
                                var v21 int8 = -27
                                _ = v21
                            }
                        }
                        _ = v12
                    } else {
                        var v22 int16 = -321
                        switch "a" {
                        case "hello":
                            var v23 bool = false
                            _ = v23
                        }
                        _ = v22
                    }
                case "":
                    v24 := GenS2{F0: (p0 % ((p0) | 1)), F1: 2147483647}
                    _ = v24
                case "a":
                    p0 -= (p0 + (p0 >> ((65535) & 63)))
                default:
                    v25 := GenS2{F0: p0, F1: 2147483647}
                    _ = v25
                }
                p0 = (p0 >> ((1) & 63))
                _ = fi11
            }
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    v26 := GenS2{F0: (p0 / ((p0) | 1)), F1: -2147483648}
    _ = p0
    _ = p1
    _ = v26
    return 1000
}

func main() {
    var v1 uint16 = 710
    v2 := make(map[uint32]int)
    var v3 float64 = 1.5
    var v4 string = ("a" + (("a" + "\n\t") + "foo bar"))
    var v5 int16 = 1
    var v27 int32 = -1
    if (-469 == 32767) {
        for fi28 := 0; fi28 < 7; fi28++ {
            switch 312 {
            case 32767:
                v29 := GenS2{F0: 1, F1: v27}
                _ = v29
            case 994:
                v30 := v2[562]
                _ = v30
            case -2147483648:
                v31 := GenS2{F0: 42, F1: (v27 ^ v27)}
                _ = v31
            }
            switch v4 {
            case "hello":
                var v32 bool = (2147483647 <= 2147483647)
                _ = v32
            }
            _ = fi28
        }
        fmt.Println(v5)
        fmt.Println(v4)
    }
    var v33 float64 = v3
    fmt.Println(v27, v4)
    fmt.Println(v1)
    v34 := []int{2147483647, -2147483648, 32768, 127}
    switch v4 {
    case "foo bar":
        v4 = ""
    case "a":
        if (true || true) {
            switch v4 {
            case "a":
                v5 -= v5
            case "\n\t":
                var v35 int16 = v5
                _ = v35
            case "foo bar":
                for fi36 := 0; fi36 < 6; fi36++ {
                    if false {
                        var v37 int16 = (v5 % ((v5) | 1))
                        _ = v37
                    }
                    var v38 int16 = -255
                    if false {
                        delete(v2, 797)
                        v39 := GenS3{F0: 32767, F1: v5}
                        var v40 string = v4
                        _ = v39
                        _ = v40
                    }
                    _ = fi36
                    _ = v38
                }
            default:
                v41 := GenS2{F0: 127, F1: -2147483648}
                _ = v41
            }
        } else {
            for fi42 := 0; fi42 < 17; fi42++ {
                for fi43 := 0; fi43 < 19; fi43++ {
                    v44 := GenS4{F0: (-(fi43)), F1: 18}
                    if false {
                        switch (v4 + ((v4 + v4) + "foo bar")) {
                        case "hello":
                            fmt.Println(fi43, v1, fi43)
                        default:
                            v45 := GenS5{F0: 1}
                            v4 = ""
                            _ = v45
                        }
                        _ = len(v34)
                        v46 := v2[496]
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
                v47 := GenS3{F0: v1, F1: (v5 - v5)}
                v1 = v1
                _ = fi42
                _ = v47
            }
            fmt.Println(v4, v3, v4)
        }
    }
    _ = len(v2)
    if (v4 != v4) {
        _ = len(v34)
        var v48 float64 = (float64(1.5) - v33)
        fmt.Println(v27, v4, v27)
        _ = v48
    } else {
        _ = len(v34)
        var v49 int64 = -1
        _ = v49
    }
    delete(v2, 2147483647)
    if (!((("hello" + v4) <= v4) && false)) {
        for fi50 := 0; fi50 < 3; fi50++ {
            v4 = v4
            var v51 string = (("" + "foo bar") + v4)
            _ = fi50
            _ = v51
        }
        if len(v34) > 0 {
            v34[0] = 690
        }
    } else {
        v33 += 0.1
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        v33 = v3
        var v53 bool = (918 < 4294967295)
        _ = fi52
        _ = v53
    }
    v4 = (v4 + v4)
    v54 := GenS4{F0: 128, F1: 0}
    switch (v4 + v4) {
    case "":
        switch 127 {
        case 264:
            if false {
                var v55 bool = (true && true)
                for fi56 := 0; fi56 < 1; fi56++ {
                    if v55 {
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                    }
                    if (((v1 == v1) || (0.1 != v3)) || ((!true) && true)) {
                        var v57 int64 = 205
                        _ = len(v2)
                        _ = v57
                    }
                    var v58 int32 = v27
                    _ = fi56
                    _ = v58
                }
                v59 := v2[993]
                _ = v55
                _ = v59
            } else {
                for fi60 := 0; fi60 < 16; fi60++ {
                    if (true && false) {
                        v61 := GenS2{F0: 0, F1: v27}
                        _ = v61
                    }
                    _ = fi60
                }
            }
        default:
            if ((v27 / ((v27) | 1)) > -383) {
                v5 -= (v5 + -1)
                v27 = v27
            }
        }
    default:
        for fi62 := 0; fi62 < 14; fi62++ {
            switch ("a" + "a") {
            case "foo bar":
                v27 += v27
            case "\n\t":
                if (!false) {
                    for fi63 := 0; fi63 < 17; fi63++ {
                        switch v4 {
                        case "\n\t":
                            v34 = append(v34, fi63)
                        }
                        v64 := GenS2{F0: -128, F1: (v27 - 0)}
                        _ = fi63
                        _ = v64
                    }
                    v65 := GenS6{F0: (((("foo bar" + v4) + v4) + "hello") + v4)}
                    for fi66 := 0; fi66 < 4; fi66++ {
                        var v67 bool = true
                        _ = fi66
                        _ = v67
                    }
                    _ = v65
                } else {
                    if true {
                        v3 -= 3.14159
                        var v68 float64 = (0.1 - v3)
                        var v69 uint8 = 1
                        _ = v68
                        _ = v69
                    } else {
                        var v70 int64 = -9223372036854775808
                        _ = v70
                    }
                }
            }
            if false {
                switch fi62 {
                case 398:
                    var v71 uint = 42
                    _ = v71
                case 32768:
                    v3 = 1e10
                default:
                    if (!((1 != 121) || ((true && true) && (1.5 == (0.1 - 0.1))))) {
                        v2[586] = int(v5)
                    } else {
                        fmt.Println(v27, v4)
                    }
                }
                v4 = ((v4 + v4) + ((v4 + "hello") + ("hello" + v4)))
                var v72 uint8 = 0
                _ = v72
            } else {
                fmt.Println(v27)
                v73 := GenS7{F0: 1}
                _ = v73
            }
            var v74 float32 = 1.5
            _ = fi62
            _ = v74
        }
    }
    for fi75 := 0; fi75 < 15; fi75++ {
        for fi76 := 0; fi76 < 19; fi76++ {
            if true {
                var v77 uint64 = 182
                if (!false) {
                    v77 += (v77 - v77)
                    if (!(true && ((v77 ^ v77) >= (v77 ^ v77)))) {
                        fmt.Println(fi76)
                        var v78 bool = ((genFn0(-39, 0) <= (v5 << ((615) & 63))) && true)
                        _ = v78
                    } else {
                        var v79 uint = 408
                        _ = v79
                    }
                    fmt.Println(fi75, fi76)
                } else {
                    if (v3 == (v33 + v33)) {
                        _ = len(v34)
                        _ = len(v34)
                        if false {
                            v1 = (v1 & v1)
                            fmt.Println(v77, v3, v33)
                            if false {
                                v80 := GenS7{F0: -1}
                                _ = v80
                            } else {
                                _ = len(v2)
                                v33 = -2.25
                            }
                        } else {
                            switch v4 {
                            case "":
                                if (v1 < v1) {
                                    var v81 float64 = (v33 / v33)
                                    fmt.Println(v81, v33, v5)
                                    v81 += 3.14159
                                    var v82 uint32 = 4294967295
                                    _ = v81
                                    _ = v82
                                }
                            case "\n\t":
                                v4 = v4
                            case "foo bar":
                                switch (fi76 - fi75) {
                                case 42:
                                    if (!true) {
                                        var v83 int8 = 127
                                        _ = v83
                                    }
                                case 524:
                                    var v84 int = fi75
                                    _ = v84
                                case -942:
                                    var v85 float32 = (1.5 * 1.5)
                                    _ = len(v2)
                                    delete(v2, 2147483647)
                                    delete(v2, 4294967295)
                                    _ = v85
                                default:
                                    var v86 int8 = -128
                                    _ = v86
                                }
                            }
                            _ = len(v2)
                            fmt.Println(v27, v4)
                            if true {
                                if true {
                                    switch (fi76 % (((fi75 * (^(fi76)))) | 1)) {
                                    case -9223372036854775808:
                                        var v87 int16 = 1
                                        _ = v87
                                    case 1:
                                        switch (v4 + "\n\t") {
                                        case "":
                                            if ((v77 % ((v77) | 1)) < v77) {
                                                var v88 int32 = -562
                                                var v89 uint16 = 255
                                                v90 := GenS5{F0: 255}
                                                _ = v88
                                                _ = v89
                                                _ = v90
                                            }
                                        default:
                                            var v91 float32 = 3.14159
                                            _ = v91
                                        }
                                    case 32768:
                                        v92 := GenS8{F0: 42, F1: 391}
                                        _ = v92
                                    }
                                } else {
                                    v27 += (v27 + (v27 ^ int32(v5)))
                                    var v93 uint32 = 0
                                    var v94 int16 = 338
                                    _ = v93
                                    _ = v94
                                }
                                v95 := GenS7{F0: -1}
                                var v96 int8 = 0
                                _ = v95
                                _ = v96
                            }
                            var v97 float64 = 1e10
                            _ = v97
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        if ((255 != 19) && ((fi76 - fi75) <= fi75)) {
                            switch (v4 + "") {
                            case "\n\t":
                                v3 += 1.5
                            case "hello":
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    v34 = append(v34, fi75)
                                case "foo bar":
                                    switch (fi76 - (^(fi76))) {
                                    case 4294967295:
                                        var v98 int16 = genFn0(1, 0)
                                        _ = v98
                                    case -1:
                                        if len(v34) >= 2 {
                                            v34 = v34[1:len(v34)]
                                        }
                                    }
                                default:
                                    if (v1 == (v1 / ((v1) | 1))) {
                                        if len(v34) >= 2 {
                                            v34 = v34[1:len(v34)]
                                        }
                                        v1 -= 1
                                    } else {
                                        fmt.Println(fi76)
                                        var v99 uint8 = 2
                                        _ = v99
                                    }
                                }
                            default:
                                v27 -= v27
                            }
                            v77 += uint64(v1)
                        } else {
                            v100 := GenS9{F0: 2, F1: 16, F2: 18446744073709551615}
                            v27 = (v27 % (((v27 << ((42) & 63))) | 1))
                            _ = v100
                        }
                    }
                    var v101 int16 = v5
                    _ = v101
                }
                _ = v77
            }
            var v102 string = ""
            var v103 float32 = 1.5
            _ = fi76
            _ = v102
            _ = v103
        }
        v104 := v2[218]
        _ = fi75
        _ = v104
    }
    for fi105 := 0; fi105 < 14; fi105++ {
        v106 := GenS5{F0: 32767}
        var v107 uint16 = (^(v1))
        v108 := GenS9{F0: 128, F1: 0, F2: 442}
        _ = fi105
        _ = v106
        _ = v107
        _ = v108
    }
    v1 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v54
}
