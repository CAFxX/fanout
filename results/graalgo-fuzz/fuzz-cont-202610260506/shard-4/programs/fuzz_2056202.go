package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 bool }
type GenS2 struct { F0 int64; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int8; F1 int64; F2 int8 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 int32) int16 {
    v6 := []int{32767, 128, 32767, -2147483648}
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 18446744073709551615, F1: fi7, F2: true}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 int32 = 927
    var v2 uint16 = 255
    var v3 int8 = -128
    v4 := []uint32{2147483647}
    var v5 uint64 = 42
    fmt.Println(v2)
    var v9 float32 = ((float32((float64((3.14159 / 0.1)) - float64(0.1))) - float32(3.14159)) + (0.0 + (float32((0.1 * 0.0)) * (((0.1 / 3.14159) * float32(0.1)) + 1.5))))
    v2 = 42
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if false {
                var v12 float64 = (((0.0 * 1e10) * (1.5 + ((0.0 * (-2.25 - -2.25)) - -2.25))) + 3.14159)
                fmt.Println(v9, v9)
                switch 4294967295 {
                case 21:
                    v1 = v1
                case 2147483647:
                    switch "a" {
                    case "hello":
                        var v13 uint16 = (v2 / (((v2 - (v2 - v2))) | 1))
                        switch 2147483647 {
                        case 32768:
                            fmt.Println(v9)
                        case -2147483648:
                            _ = len(v4)
                        case 2147483647:
                            switch fi10 {
                            case 1:
                                switch ("\n\t" + ("foo bar" + "")) {
                                case "hello":
                                    var v14 int = fi11
                                    _ = v14
                                case "a":
                                    var v15 uint16 = v13
                                    fmt.Println(v13)
                                    v2 -= (^(v2))
                                    _ = v15
                                case "\n\t":
                                    if (uint8(v2) == 1) {
                                        if (false || (3.14159 <= (v9 * v9))) {
                                            if (0 >= 2) {
                                                var v16 float32 = v9
                                                if (!true) {
                                                    switch fi11 {
                                                    case -21:
                                                        _ = len(v4)
                                                    case 9223372036854775807:
                                                        var v17 uint = 18446744073709551615
                                                        _ = v17
                                                    case 32768:
                                                        var v18 int64 = -9223372036854775808
                                                        _ = v18
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    switch (-(fi10)) {
                                                    case 1:
                                                        v19 := GenS1{F0: (v5 * 2), F1: (fi11 | (fi11 | fi10)), F2: (!(fi11 <= -1))}
                                                        _ = v19
                                                    default:
                                                        v13 += (v13 + (v13 | v13))
                                                        v4 = append(v4, 65535)
                                                        v20 := GenS1{F0: (v5 >> ((361) & 63)), F1: fi10, F2: (!(-32768 <= -1))}
                                                        _ = v20
                                                    }
                                                }
                                                var v21 int32 = v1
                                                _ = v16
                                                _ = v21
                                            }
                                            fmt.Println(v1)
                                        }
                                        if ((("hello" + ("foo bar" + ("foo bar" + ""))) + ((("a" + "a") + "a") + ("a" + ""))) < "") {
                                            if len(v4) > 0 {
                                                v4[0] = 4294967295
                                            }
                                            fmt.Println(v9)
                                        }
                                        v22 := GenS2{F0: -23, F1: uint8(v5), F2: v3}
                                        _ = v22
                                    } else {
                                        var v23 uint = 31
                                        fmt.Println(fi11, v5)
                                        var v24 float64 = float64((v9 * v9))
                                        if (-401 < 0) {
                                            fmt.Println(v5, fi11, v2)
                                            v25 := GenS2{F0: 1, F1: 0, F2: (-(v3))}
                                            var v26 uint32 = 517
                                            _ = v25
                                            _ = v26
                                        }
                                        var v27 string = ("hello" + "hello")
                                        _ = v23
                                        _ = v24
                                        _ = v27
                                    }
                                }
                            case 0:
                                if true {
                                    v12 = (float64(v9) - float64(float32((3.14159 + 1.5))))
                                }
                            case 126:
                                v28 := GenS1{F0: (v5 >> ((255) & 63)), F1: int(v2), F2: (!true)}
                                _ = v28
                            }
                        }
                        _ = v13
                    case "":
                        if false {
                            var v29 uint = 2
                            var v30 string = (("a" + "hello") + "\n\t")
                            _ = v29
                            _ = v30
                        }
                    }
                case 9223372036854775807:
                    var v31 uint64 = v5
                    _ = v31
                default:
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            v2 = (v2 - (v2 - v2))
                            _ = len(v4)
                            v2 -= (v2 - 773)
                        } else {
                            var v32 int8 = v3
                            v32 = 42
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v32
                        }
                        v5 = 4294967295
                        fmt.Println(v3, fi10)
                    case "a":
                        fmt.Println(v9)
                    case "hello":
                        var v33 uint16 = v2
                        _ = v33
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
                _ = v12
            }
            v4 = append(v4, 65535)
            if (-667 == 1) {
                if true {
                    var v34 float64 = 0.1
                    switch "hello" {
                    case "hello":
                        var v35 int16 = -213
                        _ = v35
                    case "":
                        var v36 int8 = int8(v2)
                        _ = v36
                    case "a":
                        if true {
                            v4 = append(v4, 4294967295)
                            switch fi10 {
                            case 4294967295:
                                var v37 uint8 = 42
                                _ = v37
                            }
                        } else {
                            if false {
                                var v38 int64 = 606
                                v2 -= 1
                                var v39 string = "a"
                                _ = v38
                                _ = v39
                            }
                        }
                        if false {
                            v40 := GenS3{F0: 988, F1: -9223372036854775808, F2: uint(v5)}
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v9, v9)
                            case "":
                                v41 := GenS4{F0: v2}
                                _ = v41
                            default:
                                var v42 uint = 42
                                _ = v42
                            }
                            v4 = append(v4, 2147483648)
                            _ = v40
                        }
                        var v43 uint8 = 109
                        fmt.Println(v5, v1, v9)
                        v43 += 56
                        if len(v4) > 0 {
                            v4[0] = 4294967295
                        }
                        if true {
                            fmt.Println(v5)
                            fmt.Println(fi11, v34, fi11)
                            var v44 uint32 = 4294967295
                            _ = v44
                        } else {
                            if true {
                                if (false || true) {
                                    fmt.Println(fi10, v9)
                                    var v45 float32 = -2.25
                                    _ = v45
                                }
                            } else {
                                if false {
                                    v43 = uint8(v2)
                                    var v46 int32 = 864
                                    v4 = append(v4, 525)
                                    _ = v46
                                } else {
                                    v47 := GenS4{F0: 42}
                                    switch (fi10 << ((256) & 63)) {
                                    case 124:
                                        var v48 string = "a"
                                        _ = v48
                                    case 2147483647:
                                        var v49 string = ("hello" + "hello")
                                        _ = v49
                                    case -2147483648:
                                        fmt.Println(v5, v2, v3)
                                    default:
                                        if (!false) {
                                            v5 = 1
                                        }
                                    }
                                    _ = v47
                                }
                            }
                            v9 = float32(0.1)
                        }
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        var v50 int16 = 630
                        v43 -= v43
                        _ = v43
                        _ = v50
                    default:
                        fmt.Println(v5, v3)
                    }
                    var v51 int = -987
                    var v52 int16 = 32767
                    _ = v34
                    _ = v51
                    _ = v52
                }
                var v53 int16 = 1
                fmt.Println(v3, v1)
                var v54 uint8 = 128
                v4 = append(v4, 532)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v9 += v9
                fmt.Println(v5, v1)
                v55 := GenS5{F0: "a", F1: (!true)}
                _ = v53
                _ = v54
                _ = v55
            }
            _ = fi11
        }
        v2 = 42
        _ = fi10
    }
    if false {
        if (!true) {
            fmt.Println(v3)
            if false {
                fmt.Println(v5, v5)
            } else {
                if true {
                    switch 42 {
                    case 674:
                        for fi56 := 0; fi56 < 2; fi56++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if false {
                                for fi57 := 0; fi57 < 20; fi57++ {
                                    var v58 uint32 = 18
                                    fmt.Println(v58, fi57, v1)
                                    _ = fi57
                                    _ = v58
                                }
                                if false {
                                    var v59 float32 = v9
                                    _ = v59
                                } else {
                                    v3 += (v3 << ((4294967295) & 63))
                                }
                                switch fi56 {
                                case 9223372036854775807:
                                    v9 += 1.5
                                case 147:
                                    switch -1 {
                                    case 32767:
                                        var v60 int32 = v1
                                        _ = v60
                                    case 42:
                                        for fi61 := 0; fi61 < 10; fi61++ {
                                            v4 = append(v4, 616)
                                            if ((3.14159 == 1.5) && true) {
                                                var v62 string = "a"
                                                _ = v62
                                            } else {
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    v4 = append(v4, 981)
                                                case "hello":
                                                    fmt.Println(fi56, fi56)
                                                case "":
                                                    var v63 int32 = v1
                                                    _ = v63
                                                }
                                                v9 += -2.25
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v64 uint8 = 255
                                                _ = v64
                                            }
                                            _ = fi61
                                        }
                                    default:
                                        var v65 int16 = genFn0(v1)
                                        _ = v65
                                    }
                                }
                            } else {
                                switch fi56 {
                                case 9223372036854775807:
                                    var v66 int16 = 0
                                    _ = v66
                                }
                            }
                            v67 := GenS4{F0: (v2 - 2)}
                            _ = fi56
                            _ = v67
                        }
                    case 590:
                        if (!true) {
                            if false {
                                v2 = v2
                            }
                            for fi68 := 0; fi68 < 8; fi68++ {
                                v69 := GenS6{F0: (v3 & int8(v2)), F1: 223, F2: (v3 ^ 83)}
                                fmt.Println(v9, v2, v2)
                                _ = fi68
                                _ = v69
                            }
                            if false {
                                fmt.Println(v2)
                                var v70 uint = 4294967295
                                var v71 int32 = (-(v1))
                                _ = v70
                                _ = v71
                            }
                        }
                    default:
                        v5 -= v5
                    }
                }
            }
            var v72 float64 = 0.0
            _ = v72
        } else {
            var v73 uint32 = 42
            switch 127 {
            case 128:
                if len(v4) > 0 {
                    v4[0] = (^(v73))
                }
            case -1:
                var v74 float32 = 0.1
                _ = v74
            case 42:
                if false {
                    v75 := GenS4{F0: (^(v2))}
                    var v76 float32 = v9
                    _ = v75
                    _ = v76
                } else {
                    if false {
                        v9 -= -2.25
                    }
                    var v77 uint16 = (v2 % ((v2) | 1))
                    _ = len(v4)
                    _ = v77
                }
            default:
                v78 := GenS6{F0: 0, F1: 1, F2: (v3 ^ 127)}
                _ = v78
            }
            _ = v73
        }
        v5 = v5
        for fi79 := 0; fi79 < 13; fi79++ {
            var v80 bool = false
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            _ = fi79
            _ = v80
        }
    } else {
        switch ("a" + "a") {
        case "":
            v3 = v3
        case "a":
            v2 += v2
        }
    }
    v2 += v2
    fmt.Println(v2)
    if (!false) {
        var v81 int = -561
        _ = v81
    }
    var v82 uint8 = 42
    switch int(v2) {
    case -41:
        v83 := GenS5{F0: (((((("hello" + (("" + "hello") + "")) + "foo bar") + ("a" + "hello")) + "") + "hello") + "foo bar"), F1: false}
        _ = v83
    case 9223372036854775807:
        v2 = 443
    default:
        v84 := GenS1{F0: v5, F1: 127, F2: true}
        _ = v84
    }
    for fi85 := 0; fi85 < 8; fi85++ {
        v4 = append(v4, 798)
        var v86 uint64 = uint64(v82)
        if (("foo bar" + ((("hello" + "foo bar") + ("" + "hello")) + (((("foo bar" + "foo bar") + "\n\t") + ("\n\t" + "hello")) + "hello"))) > "\n\t") {
            fmt.Println(v82, v86)
            for fi87 := 0; fi87 < 18; fi87++ {
                fmt.Println(v2, v1, fi87)
                fmt.Println(v82, v5, v1)
                if true {
                    v88 := GenS1{F0: (v86 >> ((255) & 63)), F1: (fi87 | (fi87 % ((-730) | 1))), F2: (genFn0(288) <= 42)}
                    _ = len(v4)
                    _ = v88
                } else {
                    v89 := GenS3{F0: 42, F1: fi87, F2: 688}
                    _ = v89
                }
                _ = fi87
            }
        } else {
            v90 := GenS7{F0: 2147483647}
            switch ("hello" + (("" + ("foo bar" + "hello")) + "a")) {
            case "foo bar":
                if false {
                    v82 += (v82 ^ (v82 - v82))
                    var v91 int32 = -711
                    _ = v91
                }
            case "hello":
                switch "\n\t" {
                case "foo bar":
                    for fi92 := 0; fi92 < 6; fi92++ {
                        v93 := GenS1{F0: v5, F1: int(v82), F2: true}
                        var v94 uint16 = (v2 << ((253) & 63))
                        v3 = (v3 + v3)
                        _ = fi92
                        _ = v93
                        _ = v94
                    }
                case "hello":
                    if true {
                        if ("\n\t" > "foo bar") {
                            var v95 int = (fi85 << ((255) & 63))
                            _ = v95
                        } else {
                            switch fi85 {
                            case -269:
                                var v96 int32 = -2147483648
                                _ = v96
                            case -1:
                                switch "\n\t" {
                                case "hello":
                                    v4 = append(v4, 2147483648)
                                default:
                                    switch (fi85 + -745) {
                                    case 128:
                                        for fi97 := 0; fi97 < 4; fi97++ {
                                            var v98 int8 = int8(v5)
                                            v82 = v82
                                            _ = fi97
                                            _ = v98
                                        }
                                    case 4294967295:
                                        v82 -= v82
                                    case -855:
                                        v99 := GenS6{F0: (v3 / (((^(v3))) | 1)), F1: -1, F2: (v3 & v3)}
                                        _ = v99
                                    }
                                }
                            case 1:
                                for fi100 := 0; fi100 < 16; fi100++ {
                                    var v101 uint = 65535
                                    var v102 int8 = int8(v82)
                                    if len(v4) > 0 {
                                        v4[0] = 2147483648
                                    }
                                    _ = fi100
                                    _ = v101
                                    _ = v102
                                }
                            default:
                                v9 = v9
                            }
                            var v103 int = 273
                            v4 = append(v4, 2147483648)
                            _ = v103
                        }
                        v1 -= (v1 / (((-(v1))) | 1))
                        v86 = v86
                    }
                case "a":
                    v9 -= (1.5 - float32(float64(v9)))
                }
            case "a":
                var v104 string = "\n\t"
                _ = v104
            default:
                v4 = append(v4, 2147483648)
            }
            v5 += (v86 % ((190) | 1))
            _ = v90
        }
        _ = fi85
        _ = v86
    }
    for fi105 := 0; fi105 < 17; fi105++ {
        var v106 uint32 = 65535
        _ = fi105
        _ = v106
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v82
    _ = v9
}
