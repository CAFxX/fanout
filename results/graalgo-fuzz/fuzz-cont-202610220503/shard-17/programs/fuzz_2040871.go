package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int16; F1 uint; F2 uint32 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 int64; F1 string; F2 int64 }
type GenS10 struct { F0 uint32 }
func genFn0() uint16 {
    if (-1 > 1) {
        var v7 uint8 = 122
        var v8 int8 = 1
        v9 := []int{168}
        _ = v7
        _ = v8
        _ = v9
    } else {
        switch -135 {
        case 32768:
            var v10 uint64 = 4294967295
            _ = v10
        case 1:
            var v11 float32 = 3.14159
            _ = v11
        default:
            var v12 int32 = 42
            _ = v12
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := []int{(fi13 | (fi13 >> ((807) & 63))), (fi13 ^ fi13)}
            fmt.Println(fi13)
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(fi13, fi13)
                fmt.Println(fi13, fi15)
                v16 := GenS1{F0: 255}
                var v17 uint32 = 900
                fmt.Println(v17, fi15, fi13)
                var v18 float64 = -2.25
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi13
            _ = v14
        }
        var v19 uint16 = 755
        _ = v19
    }
    v20 := GenS1{F0: 210}
    _ = v20
    return 212
}

func main() {
    var v1 float32 = 1e10
    var v2 int64 = 955
    v3 := []uint32{}
    var v4 string = ("foo bar" + "foo bar")
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    v2 = (^(v2))
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 uint64 = 0
            _ = fi22
            _ = v23
        }
        switch "\n\t" {
        case "\n\t":
            if true {
                v2 = (v2 & 1)
                var v24 int64 = (v2 * 42)
                _ = v24
            }
        case "foo bar":
            for fi25 := 0; fi25 < 1; fi25++ {
                fmt.Println(v1, v2, fi25)
                _ = fi25
            }
        }
        v4 = v4
        _ = fi21
    }
    v2 = v2
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 int32 = -1
        _ = fi26
        _ = v27
    }
    switch (("a" + v4) + (("" + ("" + ("" + v4))) + "hello")) {
    case "hello":
        v2 = -9223372036854775808
    case "":
        if false {
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 int8 = 1
                _ = fi28
                _ = v29
            }
        } else {
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 string = v4
                if (65535 == 0) {
                    if true {
                        v2 = 856
                        v2 -= -1
                    }
                    switch fi30 {
                    case 32767:
                        var v32 uint32 = 2
                        _ = v32
                    case 127:
                        if (genFn0() > 350) {
                            var v33 int8 = -1
                            var v34 uint32 = 2
                            switch fi30 {
                            case 9223372036854775807:
                                v35 := GenS2{F0: 1, F1: -1, F2: uint(v34)}
                                _ = v35
                            case -2147483648:
                                switch (^(fi30)) {
                                case 9223372036854775807:
                                    v1 -= v1
                                case -995:
                                    v36 := GenS2{F0: -532, F1: (v2 & int64(fi30)), F2: 39}
                                    _ = v36
                                case 127:
                                    if len(v3) > 0 {
                                        v3[0] = (v34 + (v34 % (((v34 - (v34 ^ 2147483647))) | 1)))
                                    }
                                }
                            case 42:
                                switch fi30 {
                                case 953:
                                    var v37 string = "\n\t"
                                    _ = v37
                                default:
                                    if false {
                                        v38 := GenS3{F0: genFn0(), F1: 0, F2: (v33 << ((4294967295) & 63))}
                                        if (65535 == 42) {
                                            v39 := GenS4{F0: -406, F1: ((v31 + "a") + "foo bar")}
                                            if true {
                                                v6[(^(fi30))] = v4
                                            } else {
                                                var v40 int8 = 0
                                                _ = v40
                                            }
                                            v1 = v1
                                            _ = v39
                                        } else {
                                            for fi41 := 0; fi41 < 15; fi41++ {
                                                _ = len(v6)
                                                _ = fi41
                                            }
                                            v4 = (v31 + (((v31 + v31) + v4) + ""))
                                        }
                                        if (874 > 1) {
                                            v3 = append(v3, (v34 - v34))
                                            if (v1 > v1) {
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483647
                                                }
                                                v31 = v31
                                                var v42 int64 = -1
                                                _ = v42
                                            }
                                            if (!true) {
                                                if true {
                                                    v2 += v2
                                                    v2 = v2
                                                } else {
                                                    _ = len(v3)
                                                    v1 = 1.5
                                                    var v43 int = int(v33)
                                                    _ = v43
                                                }
                                            } else {
                                                var v44 int8 = v33
                                                var v45 int64 = 266
                                                _ = v44
                                                _ = v45
                                            }
                                        }
                                        _ = v38
                                    }
                                }
                            default:
                                v1 = v1
                            }
                            _ = v33
                            _ = v34
                        }
                    case -9223372036854775808:
                        _ = len(v3)
                    }
                    var v46 uint = 4294967295
                    _ = v46
                } else {
                    v47 := GenS5{F0: true}
                    for fi48 := 0; fi48 < 10; fi48++ {
                        v49 := GenS5{F0: false}
                        _ = fi48
                        _ = v49
                    }
                    _ = v47
                }
                _ = fi30
                _ = v31
            }
            v1 -= v1
            v5[("\n\t" + v4)] = "foo bar"
        }
    default:
        switch ("foo bar" + v4) {
        case "a":
            for fi50 := 0; fi50 < 20; fi50++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v2, v2)
                _ = fi50
            }
        default:
            var v51 int8 = 0
            _ = v51
        }
    }
    _ = len(v3)
    v52 := GenS2{F0: 42, F1: (v2 ^ 9223372036854775807), F2: 660}
    fmt.Println(v2, v2)
    var v53 uint = 58
    v2 += -1
    if (true || false) {
        var v54 uint8 = 255
        switch "" {
        case "":
            if false {
                for fi55 := 0; fi55 < 1; fi55++ {
                    var v56 uint64 = 553
                    _ = fi55
                    _ = v56
                }
                v57 := GenS1{F0: 1}
                _ = v57
            } else {
                v4 = (v4 + "")
                fmt.Println(v1, v54)
                if (-59 >= 3) {
                    v54 += (v54 ^ v54)
                } else {
                    for fi58 := 0; fi58 < 18; fi58++ {
                        fmt.Println(fi58, v4)
                        _ = fi58
                    }
                    v4 = v4
                    v59 := v6[-2147483648]
                    _ = v59
                }
            }
        default:
            v54 += (v54 % ((100) | 1))
        }
        if (((!false) && false) && (false || false)) {
            if false {
                switch v4 {
                case "":
                    v60 := GenS2{F0: 1000, F1: v2, F2: (v53 * v53)}
                    _ = v60
                case "a":
                    var v61 string = ("" + (((v4 + v4) + v4) + ("a" + (("" + "\n\t") + ""))))
                    _ = v61
                default:
                    v62 := GenS6{F0: int(v53), F1: uint(v54)}
                    _ = v62
                }
            }
            switch v4 {
            case "foo bar":
                switch 1 {
                case -965:
                    switch 980 {
                    case 0:
                        v2 += v2
                    default:
                        var v63 string = (v4 + "hello")
                        _ = v63
                    }
                default:
                    var v64 uint64 = 18446744073709551615
                    _ = v64
                }
            case "":
                if (((2 > 42) || true) && (uint64(v54) == 2)) {
                    fmt.Println(v1, v4, v53)
                    var v65 float32 = 3.14159
                    _ = v65
                } else {
                    var v66 bool = (((1e10 / 1e10) >= 0.1) || true)
                    v67 := GenS7{F0: -42, F1: 304, F2: 1}
                    _ = v66
                    _ = v67
                }
            }
        }
        _ = v54
    }
    switch 40 {
    case 127:
        if false {
            switch (("foo bar" + "hello") + "foo bar") {
            case "foo bar":
                v68 := GenS8{F0: 18446744073709551615, F1: 4294967295}
                _ = v68
            case "a":
                switch ((v4 + v4) + ((((v4 + (v4 + "")) + "foo bar") + (v4 + "foo bar")) + "hello")) {
                case "foo bar":
                    var v69 uint64 = 4294967295
                    _ = v69
                case "hello":
                    if false {
                        for fi70 := 0; fi70 < 13; fi70++ {
                            v71 := GenS6{F0: (fi70 + fi70), F1: (v53 & (v53 & v53))}
                            if (!(!true)) {
                                var v72 uint64 = 0
                                for fi73 := 0; fi73 < 4; fi73++ {
                                    v72 = v72
                                    if (1 <= 2) {
                                        if true {
                                            if (fi70 < fi70) {
                                                v74 := GenS4{F0: 1000, F1: (((("hello" + (v4 + v4)) + "\n\t") + "") + (("hello" + v4) + (v4 + "a")))}
                                                var v75 float64 = 0.1
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                        switch ("foo bar" + (v4 + (v4 + (v4 + v4)))) {
                                        case "\n\t":
                                            if false {
                                                if false {
                                                    switch "" {
                                                    case "hello":
                                                        var v76 int = fi73
                                                        v76 = (^(v76))
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v6[v76] = v4
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v76
                                                    case "foo bar":
                                                        var v77 uint16 = uint16(v53)
                                                        _ = v77
                                                    default:
                                                        var v78 uint = (v53 - (^(v53)))
                                                        _ = v78
                                                    }
                                                }
                                            }
                                        case "foo bar":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            if false {
                                                var v79 uint16 = 0
                                                v80 := GenS9{F0: (v2 % ((9223372036854775807) | 1)), F1: ("" + "\n\t"), F2: v2}
                                                v72 -= 2
                                                v4 = v4
                                                _ = v79
                                                _ = v80
                                            } else {
                                                var v81 int64 = (v2 + v2)
                                                _ = v81
                                            }
                                        }
                                        if (v4 > (("hello" + (v4 + "\n\t")) + "\n\t")) {
                                            v82 := GenS2{F0: 32767, F1: (-(v2)), F2: v53}
                                            v53 = (v53 * v53)
                                            var v83 int64 = v2
                                            switch (fi70 ^ fi70) {
                                            case -2147483648:
                                                if (!(fi73 > (fi73 | (^(fi73))))) {
                                                    fmt.Println(v4)
                                                }
                                            case 715:
                                                v84 := GenS10{F0: 42}
                                                _ = v84
                                            case -685:
                                                var v85 int64 = (v2 | (v83 & (-(v2))))
                                                _ = v85
                                            default:
                                                v1 += ((0.1 * (float32(0.1) * ((((-2.25 / ((1.5 - float32(0.1)) + float32(0.0))) / 1e10) - (0.1 + v1)) - v1))) / -2.25)
                                            }
                                            _ = v82
                                            _ = v83
                                        } else {
                                            v72 -= 4294967295
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v2)
                                    v72 = 2
                                    _ = fi73
                                }
                                for fi86 := 0; fi86 < 1; fi86++ {
                                    if (false && (v2 <= 9223372036854775807)) {
                                        switch 127 {
                                        case 1:
                                            _ = len(v5)
                                        case 0:
                                            if true {
                                                var v87 int64 = v2
                                                _ = v87
                                            }
                                        case 367:
                                            v88 := GenS3{F0: 2, F1: 128, F2: 127}
                                            _ = v88
                                        }
                                        v2 = (-(v2))
                                    } else {
                                        var v89 bool = (0.0 < -2.25)
                                        switch fi70 {
                                        case 0:
                                            v90 := GenS10{F0: 622}
                                            _ = v90
                                        case 128:
                                            switch int(v2) {
                                            case 32767:
                                                var v91 uint16 = 42
                                                _ = v91
                                            case -9223372036854775808:
                                                v72 = uint64(v53)
                                            case 389:
                                                v92 := GenS6{F0: (^(fi70)), F1: 255}
                                                _ = v92
                                            }
                                        case -733:
                                            switch (fi70 % ((fi86) | 1)) {
                                            case 0:
                                                v2 += (v2 & (-(v2)))
                                            }
                                        default:
                                            var v93 int8 = 127
                                            _ = v93
                                        }
                                        _ = v89
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi86
                                }
                                _ = v72
                            }
                            _ = len(v6)
                            _ = fi70
                            _ = v71
                        }
                    }
                }
            default:
                v1 = 0.1
            }
        }
    case 32767:
        v94 := GenS5{F0: true}
        _ = v94
    case 284:
        fmt.Println(v1)
    default:
        var v95 int64 = 395
        _ = v95
    }
    var v96 uint64 = 2
    v97 := GenS9{F0: 1, F1: v4, F2: (-(v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v96
    _ = v97
}
