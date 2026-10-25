package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int; F1 bool; F2 int8 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS8 struct { F0 int16; F1 string; F2 int32 }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 uint16; F1 int8; F2 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64 }
func genFn0(p0 int8, p1 uint64, p2 int) uint32 {
    p1 = (p1 - p1)
    v4 := []int{(^(p2)), p2, 32768, p2}
    for fi5 := 0; fi5 < 3; fi5++ {
        p0 += 127
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}
func genFn1() uint64 {
    for fi6 := 0; fi6 < 18; fi6++ {
        if true {
            if (!(true || (88 != 100))) {
                var v7 uint = 65535
                switch 1 {
                case 9223372036854775807:
                    v7 = v7
                case -2147483648:
                    for fi8 := 0; fi8 < 12; fi8++ {
                        v9 := GenS1{F0: "hello", F1: 1}
                        if ("foo bar" < (("" + ("\n\t" + "foo bar")) + ("hello" + (("hello" + "foo bar") + (("\n\t" + "foo bar") + "\n\t"))))) {
                            v7 = v7
                        }
                        _ = fi8
                        _ = v9
                    }
                default:
                    v7 = 0
                }
                v10 := []int{(fi6 & 32767), (fi6 | fi6), -2147483648, fi6}
                _ = v10
                _ = v7
            }
            var v11 uint32 = 775
            v12 := GenS1{F0: "", F1: 0}
            _ = v11
            _ = v12
        }
        v13 := []int{(fi6 << ((0) & 63)), 9223372036854775807, fi6, (fi6 * fi6)}
        for fi14 := 0; fi14 < 4; fi14++ {
            switch (fi14 | 32768) {
            case 393:
                var v15 int32 = 1
                _ = v15
            default:
                var v16 float64 = 3.14159
                _ = v16
            }
            var v17 string = ("\n\t" + (("hello" + ("\n\t" + "")) + ""))
            _ = fi14
            _ = v17
        }
        _ = fi6
        _ = v13
    }
    var v18 float32 = 0.0
    switch ("foo bar" + (((("foo bar" + ("a" + "")) + ("a" + "hello")) + "") + "a")) {
    case "hello":
        fmt.Println(v18)
    case "foo bar":
        if true {
            var v19 string = ""
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        switch fi20 {
        case 403:
            fmt.Println(v18, fi20)
        case 238:
            if (!(385 != 18446744073709551615)) {
                if ((((v18 - (v18 * 1.5)) - v18) + v18) < (((-2.25 / ((3.14159 + float32((float64(v18) + (1e10 * 0.0)))) + ((v18 - 1e10) * (1e10 * v18)))) - v18) / 1e10)) {
                    fmt.Println(v18, fi20)
                }
                switch (fi20 * fi20) {
                case 2147483647:
                    switch fi20 {
                    case -2147483648:
                        v18 = (v18 + (3.14159 / v18))
                    }
                case 128:
                    v21 := []int{(fi20 >> ((847) & 63)), fi20, fi20, (fi20 | fi20)}
                    _ = v21
                case 127:
                    var v22 uint64 = 0
                    _ = v22
                default:
                    v18 += (0.0 / (v18 / (float32(((((0.1 + float64((v18 + v18))) - 0.0) - -2.25) - (1.5 * -2.25))) + v18)))
                }
            } else {
                var v23 int16 = 42
                _ = v23
            }
        default:
            var v24 float32 = (1e10 / 3.14159)
            _ = v24
        }
        var v25 bool = (false || (false || ((false || true) || false)))
        _ = fi20
        _ = v25
    }
    switch 2147483647 {
    case 128:
        v18 -= (((3.14159 - v18) + 1e10) * (v18 / (v18 / -2.25)))
    }
    _ = v18
    return 904
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{}
    v3 := []string{"foo bar", "a", ("foo bar" + ("foo bar" + "foo bar")), (("" + "a") + "hello")}
    var v26 uint64 = genFn1()
    var v27 int16 = 0
    var v28 int = -720
    var v29 string = "a"
    v30 := GenS2{F0: 65535, F1: 100}
    if true {
        fmt.Println(v28)
        if false {
            for fi31 := 0; fi31 < 18; fi31++ {
                var v32 int8 = int8(v26)
                var v33 int64 = 9223372036854775807
                _ = fi31
                _ = v32
                _ = v33
            }
            v34 := GenS2{F0: 65535, F1: 6}
            _ = v34
        } else {
            var v35 int64 = 1
            var v36 float32 = (0.1 - (1e10 * (((0.0 - float32(1e10)) / 1.5) - (-2.25 / 1e10))))
            v27 += v27
            _ = v35
            _ = v36
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v37 string = ((v29 + v29) + ((("hello" + v29) + v29) + (v29 + ((v29 + v29) + v29))))
        if true {
            switch int(v1) {
            case 0:
                switch 0 {
                case 32768:
                    v27 = v27
                case -162:
                    v38 := GenS3{F0: 377, F1: 1}
                    _ = v38
                case -9223372036854775808:
                    switch (v28 * v28) {
                    case 42:
                        v39 := GenS3{F0: (v28 >> ((v1) & 63)), F1: 118}
                        _ = v39
                    case 127:
                        v29 = v37
                    default:
                        v40 := GenS4{F0: 4294967295, F1: 42}
                        _ = v40
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = v37
                    }
                }
            case -2147483648:
                v41 := GenS5{F0: (v28 ^ (v28 + v28)), F1: true, F2: 127}
                _ = v41
            default:
                var v42 int16 = int16(v26)
                _ = v42
            }
            var v43 int32 = 2147483647
            v2 = append(v2, v37)
            _ = v43
        } else {
            fmt.Println(v37, v26)
        }
        _ = v37
    }
    var v44 uint8 = 234
    v45 := GenS1{F0: "hello", F1: (v27 | v27)}
    var v46 float32 = 0.0
    switch v29 {
    case "hello":
        fmt.Println(v1, v26)
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        v47 := GenS6{F0: -9223372036854775808, F1: v27}
        _ = v47
    }
    if (false && (0 > 255)) {
        v28 = 32768
        v44 = v44
        switch 826 {
        case -9223372036854775808:
            v48 := GenS2{F0: (^(v1)), F1: 0}
            _ = v48
        case -307:
            if (genFn0(42, genFn1(), v28) > 780) {
                v27 += -1
                v1 -= v1
            }
        case 4294967295:
            v44 += (^(v44))
        }
    } else {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        for fi49 := 0; fi49 < 9; fi49++ {
            v3 = append(v3, v29)
            switch v29 {
            case "":
                for fi50 := 0; fi50 < 9; fi50++ {
                    v27 -= -219
                    _ = fi50
                }
            case "a":
                v28 = (fi49 % ((218) | 1))
            }
            _ = fi49
        }
    }
    if true {
        if ((v28 + v28) <= (v28 / ((v28) | 1))) {
            var v51 int8 = 59
            if len(v2) > 0 {
                v2[0] = v29
            }
            _ = v51
        }
        if false {
            fmt.Println(v29, v28)
            var v52 uint64 = v26
            v52 = (v52 & genFn1())
            _ = v52
        }
    }
    v53 := GenS7{F0: 2147483648, F1: 0, F2: 434}
    fmt.Println(v27)
    v26 = v26
    if (-2.25 != float64(v46)) {
        _ = len(v3)
    } else {
        v27 += (v27 * (v27 >> ((65535) & 63)))
        var v54 int64 = 0
        v55 := GenS4{F0: 405, F1: uint32(v26)}
        _ = v54
        _ = v55
    }
    if ((v44 << (((v1 - 18446744073709551615)) & 63)) < v44) {
        if len(v2) > 0 {
            v2[0] = v29
        }
        for fi56 := 0; fi56 < 12; fi56++ {
            switch "a" {
            case "\n\t":
                v1 += v1
            case "":
                var v57 float32 = (float32(float64(3.14159)) / v46)
                _ = v57
            case "hello":
                v28 = (fi56 >> (((v1 + (v1 + v1))) & 63))
            }
            switch v29 {
            case "a":
                var v58 uint16 = 899
                _ = v58
            case "":
                v59 := GenS3{F0: (fi56 & v28), F1: 0}
                _ = v59
            case "foo bar":
                for fi60 := 0; fi60 < 9; fi60++ {
                    if true {
                        v44 += (v44 + (v44 ^ v44))
                    }
                    var v61 float64 = -2.25
                    var v62 uint16 = 55
                    v63 := GenS5{F0: fi56, F1: false, F2: 127}
                    switch (^(fi56)) {
                    case 32767:
                        switch "" {
                        case "":
                            switch (^(fi60)) {
                            case -9223372036854775808:
                                v1 = (v1 * uint(v44))
                            case 32768:
                                switch v29 {
                                case "hello":
                                    switch ("hello" + (v29 + "foo bar")) {
                                    case "\n\t":
                                        fmt.Println(v1, v28)
                                    case "":
                                        v61 = v61
                                    case "hello":
                                        if (v27 > v27) {
                                            if (float32(3.14159) > ((((0.0 / (v46 / 3.14159)) + ((3.14159 - 1.5) * 1e10)) + v46) + ((v46 - float32(0.0)) * (0.1 + ((0.1 * 1.5) * 0.0))))) {
                                                var v64 int16 = int16(v26)
                                                if (1 == uint32(v26)) {
                                                    fmt.Println(v64, v61, v27)
                                                    if (!(695 < genFn0(-1, (v26 * genFn1()), fi60))) {
                                                        if (!(genFn0(-11, v26, fi56) > 0)) {
                                                            v65 := GenS1{F0: "hello", F1: -247}
                                                            v26 = 2
                                                            fmt.Println(v61, v29)
                                                            _ = v65
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = (v29 + ((("a" + v29) + "\n\t") + (v29 + v29)))
                                                            }
                                                            if false {
                                                                fmt.Println(v46)
                                                                v66 := GenS8{F0: 42, F1: v29, F2: 210}
                                                                if true {
                                                                    if true {
                                                                        var v67 int64 = -1
                                                                        var v68 uint64 = v26
                                                                        var v69 float64 = ((3.14159 - 1e10) / -2.25)
                                                                        _ = v67
                                                                        _ = v68
                                                                        _ = v69
                                                                    } else {
                                                                        switch fi60 {
                                                                        case 9223372036854775807:
                                                                            if true {
                                                                                v27 = (-(v27))
                                                                                var v70 uint = v1
                                                                                v62 -= v62
                                                                                _ = v70
                                                                            } else {
                                                                                if (!true) {
                                                                                    switch fi56 {
                                                                                    case 400:
                                                                                        if (true && (true && (!(v27 >= (v64 + v27))))) {
                                                                                            v2 = append(v2, ("a" + v29))
                                                                                        } else {
                                                                                            v26 -= v26
                                                                                            var v71 int16 = v27
                                                                                            var v72 float32 = (1e10 + v46)
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                        }
                                                                                    case 1:
                                                                                        switch (fi56 ^ (v28 & v28)) {
                                                                                        case 9223372036854775807:
                                                                                            v73 := GenS5{F0: (v28 ^ (fi60 % (((v28 >> ((uint(v62)) & 63))) | 1))), F1: false, F2: 0}
                                                                                            _ = v73
                                                                                        case 42:
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = v29
                                                                                                }
                                                                                            case "":
                                                                                                v3 = append(v3, "hello")
                                                                                            case "\n\t":
                                                                                                if true {
                                                                                                    v74 := GenS1{F0: v29, F1: 42}
                                                                                                    _ = v74
                                                                                                } else {
                                                                                                    fmt.Println(fi56, v28, v1)
                                                                                                    fmt.Println(fi56, v26, v61)
                                                                                                    _ = len(v3)
                                                                                                }
                                                                                            default:
                                                                                                v75 := GenS4{F0: 893, F1: 358}
                                                                                                if (!(2 >= 2)) {
                                                                                                    if true {
                                                                                                        switch 677 {
                                                                                                        case -1:
                                                                                                            var v76 int8 = 0
                                                                                                            _ = v76
                                                                                                        case 2147483647:
                                                                                                            v46 += v46
                                                                                                        case -552:
                                                                                                            if false {
                                                                                                                v44 -= (v44 - v44)
                                                                                                                fmt.Println(fi56, v64, v29)
                                                                                                            } else {
                                                                                                                switch v29 {
                                                                                                                case "hello":
                                                                                                                    v1 = v1
                                                                                                                case "a":
                                                                                                                    v46 = v46
                                                                                                                case "":
                                                                                                                    switch "a" {
                                                                                                                    case "hello":
                                                                                                                        v62 += uint16(v1)
                                                                                                                    default:
                                                                                                                        fmt.Println(fi56)
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        switch (fi60 ^ (-(fi56))) {
                                                                                                        case 2147483647:
                                                                                                            fmt.Println(v61, v62)
                                                                                                        case -1:
                                                                                                            switch v28 {
                                                                                                            case 42:
                                                                                                                v77 := GenS6{F0: 825, F1: -1}
                                                                                                                _ = v77
                                                                                                            }
                                                                                                        case 1:
                                                                                                            if ((((false && (v64 > (v64 / ((v27) | 1)))) && (!true)) || false) && true) {
                                                                                                                var v78 int16 = 32767
                                                                                                                v46 += (1.5 / v46)
                                                                                                                fmt.Println(v64, v78, v29)
                                                                                                                v61 = v61
                                                                                                                var v79 int8 = int8(v1)
                                                                                                                _ = v78
                                                                                                                _ = v79
                                                                                                            }
                                                                                                        default:
                                                                                                            v62 = v62
                                                                                                        }
                                                                                                        v1 = (v1 ^ v1)
                                                                                                        var v80 uint64 = (v26 ^ 74)
                                                                                                        fmt.Println(v27, v61, v61)
                                                                                                        switch (((((((v29 + (("" + v29) + ((v29 + v29) + "hello"))) + (((v29 + v29) + v29) + v29)) + v29) + (v29 + "hello")) + v29) + "") + v29) {
                                                                                                        case "":
                                                                                                            if false {
                                                                                                                v81 := GenS9{F0: 42, F1: 0}
                                                                                                                var v82 uint32 = 713
                                                                                                                v83 := GenS10{F0: v62, F1: -1, F2: int(v62)}
                                                                                                                _ = v81
                                                                                                                _ = v82
                                                                                                                _ = v83
                                                                                                            } else {
                                                                                                                switch ((v29 + v29) + "\n\t") {
                                                                                                                case "hello":
                                                                                                                    _ = len(v3)
                                                                                                                default:
                                                                                                                    fmt.Println(fi56, v61)
                                                                                                                }
                                                                                                                if (v62 >= (v62 + uint16(v26))) {
                                                                                                                    var v84 float64 = (v61 * (0.1 / ((v61 - 1e10) - (v61 * v61))))
                                                                                                                    fmt.Println(fi60)
                                                                                                                    if true {
                                                                                                                        fmt.Println(v44)
                                                                                                                    } else {
                                                                                                                        v3 = append(v3, v29)
                                                                                                                        switch "\n\t" {
                                                                                                                        case "foo bar":
                                                                                                                            if len(v2) >= 2 {
                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v84
                                                                                                                } else {
                                                                                                                    fmt.Println(fi60, v26)
                                                                                                                    fmt.Println(v28, v26, v27)
                                                                                                                    v64 -= (^(v64))
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(v46, v29)
                                                                                                        }
                                                                                                        fmt.Println(v64, v1, v26)
                                                                                                        _ = len(v2)
                                                                                                        _ = v80
                                                                                                    } else {
                                                                                                        fmt.Println(v26)
                                                                                                    }
                                                                                                    fmt.Println(v64, fi60)
                                                                                                } else {
                                                                                                    fmt.Println(v62, v27, v62)
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    v85 := GenS11{F0: (fi60 * fi56)}
                                                                                                    _ = v85
                                                                                                }
                                                                                                _ = v75
                                                                                            }
                                                                                        default:
                                                                                            var v86 uint16 = (v62 * v62)
                                                                                            var v87 uint = uint(v44)
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        }
                                                                                    default:
                                                                                        v61 = 1.5
                                                                                    }
                                                                                }
                                                                                v61 = 0.1
                                                                                v64 -= (v27 - v27)
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        fmt.Println(v46, fi56, fi56)
                                                                                    } else {
                                                                                        v62 = v62
                                                                                    }
                                                                                }
                                                                                var v88 uint32 = 2
                                                                                v1 = v1
                                                                                _ = v88
                                                                            }
                                                                        case -2147483648:
                                                                            fmt.Println(v64, fi60)
                                                                        }
                                                                        v89 := GenS12{F0: -9223372036854775808}
                                                                        _ = v89
                                                                    }
                                                                    v90 := GenS12{F0: -9223372036854775808}
                                                                    _ = v90
                                                                } else {
                                                                    if false {
                                                                        _ = len(v3)
                                                                        v64 = (v64 / (((v27 + (v27 ^ v27))) | 1))
                                                                    }
                                                                }
                                                                v91 := GenS8{F0: (^(v64)), F1: "hello", F2: 0}
                                                                _ = len(v3)
                                                                _ = v66
                                                                _ = v91
                                                            }
                                                            v2 = append(v2, (v29 + (v29 + ("foo bar" + v29))))
                                                        }
                                                        switch (fi60 | fi60) {
                                                        case 128:
                                                            var v92 int = (fi60 * (-(v28)))
                                                            _ = v92
                                                        case 9223372036854775807:
                                                            var v93 uint = v1
                                                            _ = v93
                                                        case 2147483647:
                                                            if true {
                                                                v1 += 944
                                                                v28 += 32767
                                                                v27 -= v27
                                                            }
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v94 int32 = 1
                                                        _ = v94
                                                    }
                                                    var v95 int32 = -903
                                                    v96 := GenS1{F0: v29, F1: (v27 * (^(v64)))}
                                                    _ = v95
                                                    _ = v96
                                                }
                                                _ = v64
                                            } else {
                                                switch (fi60 & fi56) {
                                                case 32768:
                                                    v97 := GenS6{F0: 9223372036854775807, F1: (v27 - 1)}
                                                    _ = v97
                                                case -135:
                                                    var v98 float64 = -2.25
                                                    _ = v98
                                                case -9223372036854775808:
                                                    v2 = append(v2, (((v29 + v29) + v29) + "hello"))
                                                }
                                            }
                                        } else {
                                            var v99 int8 = -35
                                            _ = v99
                                        }
                                    default:
                                        var v100 int = 86
                                        _ = v100
                                    }
                                case "foo bar":
                                    var v101 uint16 = v62
                                    _ = v101
                                case "a":
                                    v102 := GenS2{F0: (v1 >> ((v1) & 63)), F1: -1}
                                    _ = v102
                                }
                            default:
                                if len(v3) > 0 {
                                    v3[0] = v29
                                }
                                v2 = append(v2, v29)
                                v62 += 32767
                            }
                        case "a":
                            switch fi56 {
                            case 9223372036854775807:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    default:
                        _ = len(v3)
                    }
                    var v103 uint8 = v44
                    _ = fi60
                    _ = v103
                    _ = v61
                    _ = v62
                    _ = v63
                }
            }
            _ = fi56
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v44
    _ = v45
    _ = v46
    _ = v53
}
