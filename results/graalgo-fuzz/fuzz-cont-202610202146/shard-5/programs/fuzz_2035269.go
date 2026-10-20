package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint; F1 int64; F2 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS10 struct { F0 bool; F1 int }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int64; F1 uint8; F2 int16 }
type GenS13 struct { F0 string; F1 uint }
func genFn0() uint32 {
    var v6 uint8 = 149
    if (1 >= 1) {
        v7 := []int{}
        var v8 int32 = 1
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 4294967295}
    fmt.Println(v6)
    _ = v6
    _ = v9
    return 1
}

func main() {
    v1 := []int{2147483647, -1, 1}
    var v2 int = 682
    var v3 float64 = 0.1
    var v4 uint = 255
    v5 := make(map[uint32]string)
    v10 := GenS1{F0: 1, F1: 0}
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint8 = 128
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v4, v2)
        _ = fi13
    }
    var v14 int32 = -625
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint8 = 128
        if true {
            v14 = 653
            v17 := GenS1{F0: -1, F1: 18446744073709551615}
            _ = v17
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi15
        _ = v16
    }
    v3 -= (v3 * (1.5 / v3))
    var v18 uint16 = 32767
    switch "a" {
    case "foo bar":
        switch v2 {
        case 9223372036854775807:
            v1 = append(v1, (v2 / (((v2 % (((v2 / ((v2) | 1))) | 1))) | 1)))
        }
    default:
        v19 := GenS2{F0: (v14 & v14), F1: 2147483648}
        _ = v19
    }
    var v20 string = ("hello" + "\n\t")
    for fi21 := 0; fi21 < 7; fi21++ {
        v20 = v20
        _ = fi21
    }
    v20 = "\n\t"
    var v22 float64 = (1e10 / 0.1)
    switch v2 {
    case 4294967295:
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 float32 = (0.1 - 0.1)
            switch v20 {
            case "\n\t":
                v25 := GenS2{F0: 0, F1: 2147483648}
                _ = v25
            case "foo bar":
                var v26 float64 = v3
                _ = v26
            default:
                v27 := GenS1{F0: 116, F1: 4294967295}
                _ = v27
            }
            switch (fi23 << (((v4 & 400)) & 63)) {
            case -9223372036854775808:
                var v28 uint16 = 1
                _ = v28
            case -1:
                var v29 string = v20
                _ = v29
            case 222:
                v30 := GenS3{F0: 255, F1: (v14 >> ((v4) & 63)), F2: (v4 << ((uint(v18)) & 63))}
                _ = v30
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi23
            _ = v24
        }
    case 0:
        v4 -= (v4 + v4)
    case -9223372036854775808:
        v31 := GenS3{F0: 106, F1: (v14 & int32(v18)), F2: 2}
        _ = v31
    default:
        for fi32 := 0; fi32 < 20; fi32++ {
            if (true || true) {
                var v33 uint32 = 230
                _ = v33
            }
            v20 = v20
            _ = fi32
        }
    }
    if (((-587 >= -1) && false) && false) {
        switch (("foo bar" + v20) + v20) {
        case "foo bar":
            fmt.Println(v14, v14, v20)
        default:
            v22 += 1.5
        }
        for fi34 := 0; fi34 < 19; fi34++ {
            for fi35 := 0; fi35 < 9; fi35++ {
                if (-22 != 1) {
                    fmt.Println(v4, fi34, v4)
                    if (2 >= 287) {
                        v14 += (v14 * (v14 | -2147483648))
                        if (true && false) {
                            v14 = 2147483647
                            switch v2 {
                            case 42:
                                if len(v1) > 0 {
                                    v1[0] = (^(fi34))
                                }
                            case -407:
                                if false {
                                    v36 := GenS2{F0: 0, F1: uint32(v4)}
                                    _ = v36
                                }
                            case 4294967295:
                                if len(v1) > 0 {
                                    v1[0] = (v2 | fi34)
                                }
                            }
                        } else {
                            v37 := GenS4{F0: (v14 << (((^(v4))) & 63)), F1: (v4 | v4)}
                            v22 -= ((((float64(1.5) / (((v22 / v3) * (v3 - float64(0.0))) * ((float64(float32(v3)) + (v3 + v3)) * v22))) / (v22 + v3)) / ((v3 - v22) - 1.5)) / float64(((((float32(v22) * (1.5 + float32(0.0))) + 1.5) + 1e10) / 1e10)))
                            _ = v37
                        }
                    }
                }
                if (true || (true || false)) {
                    v5[255] = (v20 + (v20 + v20))
                } else {
                    fmt.Println(v20, fi35)
                }
                var v38 uint32 = 522
                v20 = v20
                _ = fi35
                _ = v38
            }
            for fi39 := 0; fi39 < 14; fi39++ {
                v40 := GenS3{F0: 3, F1: v14, F2: uint(v18)}
                v1 = append(v1, (fi39 << ((v4) & 63)))
                _ = fi39
                _ = v40
            }
            for fi41 := 0; fi41 < 9; fi41++ {
                if ((v4 >> (((v4 << (((^(v4))) & 63))) & 63)) >= v4) {
                    if (false && false) {
                        var v42 uint = 5
                        v43 := GenS5{F0: 1, F1: 779}
                        switch 983 {
                        case 2147483647:
                            v44 := GenS2{F0: (^(v14)), F1: genFn0()}
                            _ = v44
                        default:
                            fmt.Println(fi41)
                        }
                        if (true || false) {
                            if false {
                                switch ((v20 + v20) + "") {
                                case "hello":
                                    var v45 float64 = (1.5 / (v3 * v3))
                                    _ = v45
                                }
                            }
                            if (!false) {
                                if (((v4 < 18446744073709551615) || false) || false) {
                                    v22 = v22
                                }
                            }
                        }
                        v4 -= v42
                        v4 -= (v4 % (((v4 / ((2) | 1))) | 1))
                        _ = len(v1)
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(v18)
                    }
                }
                _ = fi41
            }
            _ = fi34
        }
        if true {
            _ = len(v1)
            v1 = append(v1, v2)
        }
    }
    fmt.Println(v3, v20)
    for fi46 := 0; fi46 < 2; fi46++ {
        for fi47 := 0; fi47 < 3; fi47++ {
            var v48 int8 = 106
            switch (-(v2)) {
            case 42:
                v48 = 83
            case 9223372036854775807:
                v49 := GenS6{F0: 9223372036854775807, F1: (true && true)}
                v50 := GenS7{F0: (v4 >> (((v4 >> ((320) & 63))) & 63)), F1: 0, F2: v20}
                var v51 uint = v4
                _ = v49
                _ = v50
                _ = v51
            case -770:
                if false {
                    v18 = 32767
                    switch (fi47 << ((0) & 63)) {
                    case -143:
                        if (77 > 0) {
                            if ((((v3 + ((((0.0 / v3) - (v22 * v3)) - (float64((1.5 - -2.25)) * v22)) + 3.14159)) / ((float64(3.14159) / (float64(3.14159) - -2.25)) - v3)) * v22) == 0.0) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v2 += int(v18)
                            v14 = v14
                        }
                    case 635:
                        fmt.Println(v18, v3)
                    }
                    var v52 uint8 = 58
                    if false {
                        switch (fi46 % ((0) | 1)) {
                        case -2147483648:
                            var v53 uint64 = 0
                            _ = v53
                        default:
                            var v54 uint16 = (v18 / (((v18 >> ((v4) & 63))) | 1))
                            _ = v54
                        }
                    }
                    _ = v52
                } else {
                    v55 := GenS8{F0: 2}
                    if false {
                        switch (fi46 % (((fi47 * v2)) | 1)) {
                        case -298:
                            if ((0 < 127) && true) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case -9223372036854775808:
                            if true {
                                v56 := GenS3{F0: 2, F1: (v14 >> ((v4) & 63)), F2: v4}
                                v48 = v48
                                switch 818 {
                                case 938:
                                    var v57 int16 = 0
                                    _ = v57
                                case 127:
                                    fmt.Println(v4, v4, v3)
                                }
                                fmt.Println(v4)
                                switch "\n\t" {
                                case "\n\t":
                                    switch (-(fi47)) {
                                    case 127:
                                        if (int32(v48) < 1) {
                                            v18 = (v18 * v18)
                                            var v58 uint8 = 2
                                            _ = v58
                                        } else {
                                            var v59 float64 = 3.14159
                                            var v60 float32 = 1.5
                                            v48 = (v48 + v48)
                                            _ = v59
                                            _ = v60
                                        }
                                    default:
                                        switch ((v20 + ((v20 + "a") + v20)) + "hello") {
                                        case "\n\t":
                                            _ = len(v1)
                                        case "a":
                                            v61 := GenS6{F0: -9223372036854775808, F1: false}
                                            _ = v61
                                        case "hello":
                                            var v62 uint32 = 2
                                            _ = v62
                                        }
                                    }
                                case "hello":
                                    var v63 string = (v20 + "hello")
                                    _ = v63
                                }
                                _ = v56
                            } else {
                                v64 := GenS9{F0: 2, F1: v18, F2: 1}
                                if len(v1) > 0 {
                                    v1[0] = (v2 ^ fi47)
                                }
                                var v65 string = ((v20 + (((("" + v20) + v20) + (v20 + "")) + (v20 + "a"))) + v20)
                                _ = v64
                                _ = v65
                            }
                        case 128:
                            v22 = (((v3 * ((0.0 / (v3 / v3)) / 1.5)) + float64(-2.25)) - 0.0)
                        }
                        var v66 int8 = -60
                        v67 := GenS10{F0: true, F1: 32767}
                        if (v3 <= v22) {
                            var v68 int32 = (v14 << ((v4) & 63))
                            _ = len(v1)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v68
                        }
                        _ = v66
                        _ = v67
                    } else {
                        v3 += v3
                        switch "" {
                        case "foo bar":
                            var v69 float32 = (1.5 / float32(float64(1.5)))
                            _ = v69
                        case "":
                            if (32767 < 375) {
                                if (((v18 == (v18 >> ((v4) & 63))) || (936 == 1)) && false) {
                                    v18 = v18
                                    v70 := GenS10{F0: (!true), F1: (fi46 ^ v2)}
                                    _ = v70
                                }
                                if true {
                                    switch v20 {
                                    case "a":
                                        switch (fi46 ^ (^(fi47))) {
                                        case 42:
                                            if (((v3 + float64(0.1)) * v3) < (((1.5 + (float64(0.1) * ((v22 / 1.5) + v22))) + v22) / (1.5 - ((float64(0.0) / -2.25) * v3)))) {
                                                fmt.Println(fi46, v18)
                                            } else {
                                                v71 := GenS4{F0: -350, F1: 497}
                                                fmt.Println(v3, fi46)
                                                if false {
                                                    v2 = -9223372036854775808
                                                    if true {
                                                        v3 -= (v22 - float64((1e10 * ((3.14159 / (float32(float64(1e10)) - 0.0)) * (3.14159 - (3.14159 * ((0.1 - 1.5) / 1.5)))))))
                                                        var v72 uint64 = 320
                                                        _ = v72
                                                    } else {
                                                        var v73 int8 = 30
                                                        _ = v73
                                                    }
                                                    var v74 uint32 = genFn0()
                                                    _ = len(v5)
                                                    _ = len(v1)
                                                    if ((fi47 ^ 1) != (fi47 - fi46)) {
                                                        var v75 int16 = 1000
                                                        _ = v75
                                                    }
                                                    fmt.Println(v74, fi47)
                                                    _ = v74
                                                } else {
                                                    v1 = append(v1, int(v4))
                                                    v14 += -336
                                                    var v76 uint = 0
                                                    _ = v76
                                                }
                                                _ = v71
                                            }
                                        case -9223372036854775808:
                                            delete(v5, 879)
                                        }
                                    case "hello":
                                        switch fi47 {
                                        case 42:
                                            fmt.Println(v3, fi46)
                                        case -9223372036854775808:
                                            if false {
                                                v77 := v5[1]
                                                switch (v2 / ((0) | 1)) {
                                                case 609:
                                                    var v78 string = (v77 + v77)
                                                    _ = v78
                                                case 32768:
                                                    var v79 int32 = 42
                                                    var v80 int32 = v79
                                                    _ = v79
                                                    _ = v80
                                                }
                                                switch "a" {
                                                case "":
                                                    var v81 float32 = 0.0
                                                    _ = v81
                                                case "a":
                                                    v82 := GenS3{F0: 1, F1: (^(v14)), F2: (v4 - v4)}
                                                    _ = v82
                                                case "foo bar":
                                                    if (false || (int(v4) < (v2 ^ fi46))) {
                                                        v83 := GenS11{F0: false}
                                                        _ = v83
                                                    }
                                                }
                                                _ = v77
                                            } else {
                                                var v84 uint16 = 255
                                                if (1 >= 1) {
                                                    if true {
                                                        switch ((v20 + ("hello" + (v20 + ""))) + v20) {
                                                        case "foo bar":
                                                            var v85 int32 = (v14 / ((v14) | 1))
                                                            _ = v85
                                                        case "a":
                                                            var v86 bool = (false && (!false))
                                                            _ = v86
                                                        case "\n\t":
                                                            _ = len(v5)
                                                        default:
                                                            v20 = ""
                                                        }
                                                        var v87 uint8 = 80
                                                        _ = v87
                                                    }
                                                    v5[42] = "foo bar"
                                                }
                                                v88 := GenS3{F0: 255, F1: 1, F2: v4}
                                                _ = v84
                                                _ = v88
                                            }
                                        case -931:
                                            delete(v5, 2)
                                        }
                                    case "":
                                        var v89 int32 = 0
                                        _ = v89
                                    }
                                } else {
                                    v90 := GenS12{F0: -9223372036854775808, F1: 1, F2: -519}
                                    fmt.Println(v4, v20, fi46)
                                    _ = v90
                                }
                            }
                        }
                        v20 = (v20 + (v20 + "\n\t"))
                    }
                    var v91 uint32 = 1
                    v2 -= fi47
                    _ = v55
                    _ = v91
                }
            default:
                if false {
                    v92 := GenS13{F0: (v20 + ((((v20 + ("hello" + "a")) + ("a" + ("hello" + "a"))) + v20) + "\n\t")), F1: (^(v4))}
                    _ = v92
                }
            }
            v93 := GenS9{F0: 2, F1: uint16(v4), F2: 0}
            _ = fi47
            _ = v48
            _ = v93
        }
        delete(v5, uint32(v4))
        _ = len(v1)
        _ = fi46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
