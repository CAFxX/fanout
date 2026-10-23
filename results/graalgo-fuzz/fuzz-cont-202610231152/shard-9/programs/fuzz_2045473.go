package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool; F1 int32; F2 uint8 }
type GenS11 struct { F0 uint32; F1 string }
type GenS12 struct { F0 string; F1 int64 }
type GenS13 struct { F0 bool; F1 uint8; F2 bool }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint8; F1 int8 }
type GenS16 struct { F0 int16; F1 int64 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 bool; F1 string }
type GenS20 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64, p1 uint64) uint {
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 1, F1: 128}
        _ = fi5
        _ = v6
    }
    p1 = p0
    fmt.Println(p0)
    p1 += (p1 | (p0 - (p0 * p1)))
    if (false || false) {
        if ((599 > 0) && ("" == "hello")) {
            if (0 > 68) {
                if false {
                    p1 -= (p1 << ((533) & 63))
                    for fi7 := 0; fi7 < 3; fi7++ {
                        v8 := GenS2{F0: 4294967295, F1: ("hello" + ("a" + "\n\t"))}
                        p0 -= p1
                        _ = fi7
                        _ = v8
                    }
                } else {
                    fmt.Println(p0)
                    fmt.Println(p1)
                }
                if false {
                    v9 := []int{0, 127}
                    _ = v9
                }
                v10 := GenS1{F0: 875, F1: 42}
                _ = v10
            } else {
                v11 := GenS3{F0: 1, F1: -1}
                if false {
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 int16 = 32767
                        v14 := []int{fi12, fi12}
                        for fi15 := 0; fi15 < 19; fi15++ {
                            fmt.Println(p0)
                            switch 0 {
                            case 42:
                                var v16 int64 = 42
                                _ = v16
                            case 688:
                                fmt.Println(v13, p1)
                            case -938:
                                if false {
                                    v13 += v13
                                    if len(v14) > 0 {
                                        v14[0] = fi15
                                    }
                                    p0 = (p1 >> ((2) & 63))
                                } else {
                                    var v17 float64 = (float64(((1.5 * -2.25) / ((1.5 * (1e10 + -2.25)) - (1.5 * 3.14159)))) - (1.5 + 0.0))
                                    v17 = (((1.5 / v17) * v17) + v17)
                                    _ = v17
                                }
                            }
                            v18 := GenS2{F0: uint32(p1), F1: "\n\t"}
                            _ = fi15
                            _ = v18
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                if ((!false) && (false && false)) {
                    var v19 uint64 = (p1 / ((p0) | 1))
                    v20 := []int{}
                    _ = v19
                    _ = v20
                }
                _ = v11
            }
            switch 128 {
            case -9223372036854775808:
                p1 -= p0
            case 127:
                v21 := GenS4{F0: 840}
                _ = v21
            default:
                p1 = (p0 / ((344) | 1))
            }
            v22 := GenS5{F0: 1, F1: 42}
            _ = v22
        } else {
            p1 = (p1 % ((p1) | 1))
        }
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 486
}

func main() {
    var v1 string = ((("\n\t" + "\n\t") + "a") + ("a" + (("a" + ("\n\t" + (("a" + "a") + "\n\t"))) + "")))
    var v2 int32 = -1
    v3 := []uint32{1, 289}
    var v4 int64 = -1
    for fi23 := 0; fi23 < 14; fi23++ {
        if false {
            for fi24 := 0; fi24 < 6; fi24++ {
                if false {
                    if (!true) {
                        var v25 uint8 = 42
                        v26 := GenS3{F0: 0, F1: 0}
                        var v27 bool = false
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        v28 := GenS6{F0: 9223372036854775807}
                        v2 -= (^(v2))
                        _ = v28
                    }
                    fmt.Println(v2, v4)
                    v1 = ("\n\t" + "a")
                }
                _ = fi24
            }
            var v29 int8 = 127
            fmt.Println(v29)
            _ = v29
        } else {
            v1 = v1
            if (1e10 > ((3.14159 / ((3.14159 * (3.14159 * 1.5)) * 0.1)) - (3.14159 - float64(1e10)))) {
                v2 = (^(v2))
                _ = len(v3)
            } else {
                for fi30 := 0; fi30 < 2; fi30++ {
                    var v31 int = (fi30 | fi30)
                    _ = fi30
                    _ = v31
                }
                fmt.Println(v1)
                v32 := GenS3{F0: 1000, F1: -1}
                _ = v32
            }
        }
        switch "hello" {
        case "hello":
            v4 += (v4 % (((v4 * (v4 - (v4 << ((0) & 63))))) | 1))
        case "":
            v2 += v2
        case "\n\t":
            v3 = append(v3, 42)
        }
        _ = fi23
    }
    var v33 uint32 = 111
    if (!true) {
        fmt.Println(v2, v1, v4)
        fmt.Println(v2, v1)
        v34 := GenS3{F0: 578, F1: 100}
        _ = v34
    } else {
        fmt.Println(v33, v33, v4)
    }
    switch 924 {
    case 9223372036854775807:
        v35 := GenS7{F0: 0, F1: 811, F2: (v33 << ((4294967295) & 63))}
        _ = v35
    case -9223372036854775808:
        v33 += (^(v33))
    }
    if false {
        if (("hello" == ("" + v1)) || true) {
            v36 := GenS5{F0: 0, F1: 0}
            v37 := GenS8{F0: 4294967295}
            if false {
                fmt.Println(v1, v1, v4)
            }
            _ = v36
            _ = v37
        }
        v3 = append(v3, (^(v33)))
        v2 += (v2 ^ (^(v2)))
    }
    if false {
        v1 = (v1 + ((v1 + "") + v1))
        if (!false) {
            switch 69 {
            case 779:
                v33 += (v33 & 451)
            case 2147483647:
                fmt.Println(v1)
            case 269:
                switch v1 {
                case "foo bar":
                    fmt.Println(v2, v4)
                case "hello":
                    v33 = (v33 * (v33 >> ((genFn0(42, 2)) & 63)))
                case "\n\t":
                    switch -23 {
                    case 1:
                        var v38 int = 128
                        _ = v38
                    case -398:
                        if ((v4 >> ((2) & 63)) < -100) {
                            if false {
                                v39 := GenS3{F0: 1, F1: -128}
                                _ = v39
                            } else {
                                var v40 int32 = v2
                                v2 -= v40
                                _ = v40
                            }
                            var v41 uint16 = 1
                            if (-9223372036854775808 >= v4) {
                                if true {
                                    for fi42 := 0; fi42 < 2; fi42++ {
                                        v3 = append(v3, v33)
                                        _ = len(v3)
                                        fmt.Println(v33)
                                        _ = fi42
                                    }
                                    var v43 uint8 = 42
                                    _ = v43
                                } else {
                                    v44 := GenS9{F0: 42}
                                    _ = v44
                                }
                                v41 -= 257
                                var v45 uint16 = (v41 % ((v41) | 1))
                                _ = v45
                            } else {
                                switch 4294967295 {
                                case 128:
                                    var v46 uint16 = v41
                                    _ = v46
                                case 9223372036854775807:
                                    switch -9223372036854775808 {
                                    case -1:
                                        v47 := GenS10{F0: false, F1: v2, F2: 124}
                                        _ = v47
                                    case -380:
                                        if len(v3) > 0 {
                                            v3[0] = (v33 - 0)
                                        }
                                    case -154:
                                        var v48 float32 = ((1.5 + float32((float64(0.1) - (0.0 * (-2.25 / ((0.0 * 3.14159) + ((-2.25 * 1.5) - float64(float32((0.1 / float64(((-2.25 + 1.5) - (-2.25 + (((-2.25 / 1e10) + 0.0) * 1.5)))))))))))))) - (-2.25 / float32(3.14159)))
                                        _ = v48
                                    }
                                case -9223372036854775808:
                                    v49 := GenS4{F0: v4}
                                    _ = v49
                                default:
                                    v4 = v4
                                }
                                fmt.Println(v33)
                            }
                            _ = v41
                        } else {
                            switch int(v4) {
                            case 279:
                                switch v1 {
                                case "":
                                    fmt.Println(v33, v4)
                                case "a":
                                    v33 = 42
                                }
                            default:
                                switch 42 {
                                case -1:
                                    if false {
                                        v2 = (v2 + (v2 * v2))
                                        v2 = v2
                                        if ((!false) || true) {
                                            var v50 string = (v1 + (v1 + "hello"))
                                            _ = v50
                                        }
                                    } else {
                                        var v51 float32 = (3.14159 - -2.25)
                                        _ = len(v3)
                                        switch "\n\t" {
                                        case "\n\t":
                                            if (!false) {
                                                var v52 uint16 = 544
                                                if (!(!true)) {
                                                    for fi53 := 0; fi53 < 16; fi53++ {
                                                        var v54 int16 = 32767
                                                        v1 = ("" + "foo bar")
                                                        switch 47 {
                                                        case 127:
                                                            if len(v3) > 0 {
                                                                v3[0] = (v33 % ((v33) | 1))
                                                            }
                                                        case 0:
                                                            v55 := GenS11{F0: 753, F1: v1}
                                                            _ = v55
                                                        case 2147483647:
                                                            for fi56 := 0; fi56 < 16; fi56++ {
                                                                var v57 uint = genFn0(968, 0)
                                                                v54 = 1
                                                                fmt.Println(v51, v2, fi53)
                                                                if (uint(v33) >= (v57 / ((v57) | 1))) {
                                                                    var v58 string = v1
                                                                    if false {
                                                                        var v59 uint64 = 587
                                                                        fmt.Println(fi56, v59, v52)
                                                                        v51 -= v51
                                                                        _ = v59
                                                                    }
                                                                    _ = v58
                                                                }
                                                                fmt.Println(v54)
                                                                _ = fi56
                                                                _ = v57
                                                            }
                                                        default:
                                                            var v60 int64 = 1
                                                            _ = v60
                                                        }
                                                        _ = fi53
                                                        _ = v54
                                                    }
                                                    switch -2147483648 {
                                                    case 0:
                                                        v2 = 42
                                                    case -265:
                                                        v61 := GenS12{F0: v1, F1: (^(v4))}
                                                        _ = v61
                                                    case 2147483647:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    var v62 uint = 0
                                                    _ = v62
                                                }
                                                if (2 > (^(v52))) {
                                                    if (v2 < (v2 ^ (^(v2)))) {
                                                        v63 := GenS4{F0: v4}
                                                        _ = v63
                                                    } else {
                                                        var v64 int8 = 100
                                                        switch 1 {
                                                        case -201:
                                                            v3 = append(v3, (v33 + uint32(v52)))
                                                        case 42:
                                                            var v65 int64 = v4
                                                            _ = v65
                                                        default:
                                                            v66 := GenS13{F0: true, F1: 0, F2: false}
                                                            _ = v66
                                                        }
                                                        switch 127 {
                                                        case -714:
                                                            if true {
                                                                v52 = v52
                                                                if (false && false) {
                                                                    v67 := GenS14{F0: true}
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (^(v33))
                                                                    }
                                                                    _ = v67
                                                                } else {
                                                                    fmt.Println(v52)
                                                                    for fi68 := 0; fi68 < 10; fi68++ {
                                                                        for fi69 := 0; fi69 < 1; fi69++ {
                                                                            var v70 uint64 = 1
                                                                            _ = fi69
                                                                            _ = v70
                                                                        }
                                                                        _ = len(v3)
                                                                        for fi71 := 0; fi71 < 8; fi71++ {
                                                                            var v72 uint8 = 102
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v33
                                                                            }
                                                                            _ = fi71
                                                                            _ = v72
                                                                        }
                                                                        _ = fi68
                                                                    }
                                                                    v64 = (v64 * (v64 % ((v64) | 1)))
                                                                }
                                                                switch ((v1 + v1) + ("a" + v1)) {
                                                                case "":
                                                                    fmt.Println(v1)
                                                                default:
                                                                    var v73 int8 = (v64 << ((65535) & 63))
                                                                    _ = v73
                                                                }
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v33 & v33)
                                                                }
                                                                for fi74 := 0; fi74 < 6; fi74++ {
                                                                    for fi75 := 0; fi75 < 2; fi75++ {
                                                                        fmt.Println(v33, v52, v1)
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (v33 % (((v33 ^ (^(v33)))) | 1))
                                                                        }
                                                                        v76 := GenS15{F0: 2, F1: (v64 ^ (^(v64)))}
                                                                        _ = fi75
                                                                        _ = v76
                                                                    }
                                                                    _ = fi74
                                                                }
                                                            }
                                                        case -105:
                                                            v77 := GenS15{F0: 72, F1: 100}
                                                            _ = v77
                                                        case 1:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            var v78 uint32 = v33
                                                            _ = v78
                                                        }
                                                        _ = v64
                                                    }
                                                    v79 := GenS16{F0: 1, F1: (-(v4))}
                                                    var v80 uint8 = 255
                                                    _ = v79
                                                    _ = v80
                                                }
                                                _ = v52
                                            }
                                        }
                                        _ = v51
                                    }
                                default:
                                    v2 -= 623
                                }
                            }
                            switch -1 {
                            case 4294967295:
                                v81 := GenS17{F0: -22}
                                _ = v81
                            case -1:
                                for fi82 := 0; fi82 < 7; fi82++ {
                                    for fi83 := 0; fi83 < 10; fi83++ {
                                        if false {
                                            if ((v2 * v2) > v2) {
                                                _ = len(v3)
                                            }
                                            v84 := GenS18{F0: 127}
                                            fmt.Println(fi82)
                                            _ = v84
                                        } else {
                                            v85 := GenS19{F0: true, F1: v1}
                                            v4 = 5
                                            if true {
                                                v4 -= 0
                                            }
                                            _ = v85
                                        }
                                        var v86 uint8 = 42
                                        _ = fi83
                                        _ = v86
                                    }
                                    v33 = (v33 * (^(v33)))
                                    _ = fi82
                                }
                            default:
                                v4 = v4
                            }
                        }
                    }
                }
            default:
                v87 := GenS18{F0: 0}
                _ = v87
            }
        }
        v2 -= (v2 + 388)
    }
    v4 = (v4 | v4)
    v1 = v1
    v88 := GenS2{F0: (v33 % ((v33) | 1)), F1: v1}
    if true {
        if ((v2 << ((4294967295) & 63)) <= (v2 / ((v2) | 1))) {
            for fi89 := 0; fi89 < 15; fi89++ {
                var v90 int = fi89
                for fi91 := 0; fi91 < 4; fi91++ {
                    v1 = ("" + (((((v1 + v1) + "foo bar") + "") + ((((v1 + "") + "") + v1) + v1)) + v1))
                    v2 += 1
                    _ = fi91
                }
                _ = fi89
                _ = v90
            }
            switch 4294967295 {
            case -492:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 2147483647:
                var v92 string = ((((v1 + v1) + ((v1 + ("\n\t" + v1)) + v1)) + "foo bar") + ("a" + v1))
                _ = v92
            }
        } else {
            if true {
                if ((-1 >= -1) || true) {
                    var v93 int = -9223372036854775808
                    var v94 uint32 = (v33 - v33)
                    var v95 bool = (!false)
                    _ = v93
                    _ = v94
                    _ = v95
                }
                v4 += v4
            }
            var v96 uint = 0
            _ = v96
        }
    }
    if (!(true || (true && (!(678 < 2))))) {
        var v97 int32 = v2
        _ = v97
    } else {
        var v98 int8 = 1
        for fi99 := 0; fi99 < 17; fi99++ {
            for fi100 := 0; fi100 < 1; fi100++ {
                if false {
                    fmt.Println(v2, v98, fi99)
                    switch (fi99 - (fi100 * (^(fi100)))) {
                    case 9223372036854775807:
                        v101 := GenS8{F0: 930}
                        _ = v101
                    }
                    fmt.Println(fi99)
                }
                v3 = append(v3, 0)
                if false {
                    fmt.Println(fi100, v1, v1)
                    v102 := GenS15{F0: 127, F1: (-(v98))}
                    switch v1 {
                    case "foo bar":
                        v103 := GenS20{F0: (!((1.5 - 0.1) < 1e10)), F1: 1}
                        _ = v103
                    }
                    var v104 int64 = int64(fi100)
                    var v105 bool = true
                    switch (^(fi100)) {
                    case 713:
                        v1 = ("\n\t" + v1)
                    }
                    var v106 int = (fi100 ^ 4294967295)
                    v98 = v98
                    _ = v102
                    _ = v104
                    _ = v105
                    _ = v106
                } else {
                    v4 = (v4 / (((-(v4))) | 1))
                    v4 = 1
                }
                _ = fi100
            }
            v4 = -648
            v1 = ("\n\t" + "hello")
            _ = fi99
        }
        v107 := GenS20{F0: true, F1: 1}
        _ = v107
        _ = v98
    }
    if (!true) {
        var v108 int8 = -1
        _ = v108
    } else {
        if len(v3) > 0 {
            v3[0] = (v33 * 275)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v88
}
