package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 bool }
func genFn0() string {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{4294967295, (fi4 >> ((65535) & 63)), (-(fi4))}
        _ = fi4
        _ = v5
    }
    v6 := []int{}
    switch "\n\t" {
    case "foo bar":
        for fi7 := 0; fi7 < 17; fi7++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi7
        }
    default:
        v8 := GenS1{F0: 143}
        _ = v8
    }
    _ = v6
    return "a"
}
func genFn1(p0 int32, p1 uint16) int16 {
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float64 = 1e10
        var v11 uint32 = 106
        switch ("" + ("hello" + "foo bar")) {
        case "":
            for fi12 := 0; fi12 < 11; fi12++ {
                p0 = -1
                var v13 int64 = -220
                fmt.Println(v11, v10)
                _ = fi12
                _ = v13
            }
        case "hello":
            var v14 int32 = p0
            _ = v14
        default:
            v15 := GenS2{F0: 18446744073709551615}
            _ = v15
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: 4294967295}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "a" {
        case "":
            p0 = -1
        case "a":
            var v19 int16 = -32768
            _ = v19
        case "\n\t":
            switch ((("a" + "\n\t") + ("foo bar" + ("hello" + ""))) + (("" + "a") + ("hello" + ("a" + "a")))) {
            case "foo bar":
                fmt.Println(p1, fi18)
            case "hello":
                v20 := []int{(-(fi18)), fi18, fi18, (-(fi18))}
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 15; fi21++ {
                    switch fi21 {
                    case 32767:
                        var v22 uint = 1
                        _ = v22
                    case -169:
                        var v23 int16 = 1000
                        _ = v23
                    case -9223372036854775808:
                        if (255 > (^(p1))) {
                            p0 -= (-(p0))
                            var v24 int32 = p0
                            _ = v24
                        }
                    }
                    switch "hello" {
                    case "":
                        if ((2147483647 == 0) || true) {
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(p0, p0, p0)
                                switch "foo bar" {
                                case "foo bar":
                                    v25 := GenS1{F0: -1}
                                    _ = v25
                                case "a":
                                    v26 := []int{-994, (fi21 % ((fi18) | 1))}
                                    _ = v26
                                }
                                if true {
                                    v27 := []int{}
                                    _ = v27
                                }
                            case "":
                                v28 := []int{(fi18 ^ fi21), int(p0)}
                                _ = v28
                            case "a":
                                switch "" {
                                case "\n\t":
                                    if false {
                                        fmt.Println(p1, fi18, p0)
                                        p0 -= (p0 / ((1) | 1))
                                        v29 := []int{(fi18 & 373), (fi18 * fi18), fi21, (fi21 * 42)}
                                        v30 := GenS3{F0: (fi21 % ((fi18) | 1))}
                                        p0 = -1
                                        _ = v29
                                        _ = v30
                                    } else {
                                        p1 = p1
                                        v31 := []int{fi21, fi21}
                                        fmt.Println(p1, fi18)
                                        _ = v31
                                    }
                                }
                            default:
                                fmt.Println(fi21, p1, p1)
                            }
                        } else {
                            fmt.Println(fi21)
                            switch ("" + "foo bar") {
                            case "hello":
                                var v32 uint16 = (p1 << ((2) & 63))
                                var v33 float32 = (1e10 / -2.25)
                                switch "foo bar" {
                                case "a":
                                    if ("hello" == ("" + "a")) {
                                        v34 := []int{}
                                        fmt.Println(v33, v32)
                                        var v35 int64 = 9223372036854775807
                                        _ = v34
                                        _ = v35
                                    }
                                }
                                _ = v32
                                _ = v33
                            case "":
                                p0 -= (^(p0))
                                var v36 int = fi18
                                _ = v36
                            case "foo bar":
                                v37 := GenS4{F0: 42, F1: -1}
                                _ = v37
                            default:
                                fmt.Println(fi21)
                            }
                        }
                    case "a":
                        p1 += p1
                    case "foo bar":
                        fmt.Println(p1, p0)
                    }
                    p0 += p0
                    _ = fi21
                }
            }
        default:
            v38 := GenS4{F0: 669, F1: -725}
            _ = v38
        }
        if (fi18 <= fi18) {
            var v39 string = ""
            v40 := []int{(fi18 | fi18), fi18}
            var v41 float64 = float64((1.5 + (((((-2.25 + (1e10 / ((-2.25 * (1e10 - float32(float64(3.14159)))) - -2.25))) / float32((1.5 - 0.0))) - ((0.1 * 1e10) * 1e10)) - (1e10 + 1e10)) - (1.5 / 0.1))))
            _ = v39
            _ = v40
            _ = v41
        }
        if (true || true) {
            if (false || (82 < 0)) {
                v42 := []int{}
                var v43 float32 = 0.0
                var v44 float64 = -2.25
                _ = v42
                _ = v43
                _ = v44
            } else {
                var v45 int8 = -1
                if (!(290 > 1)) {
                    v46 := []int{0, (fi18 * fi18), fi18}
                    _ = v46
                }
                _ = v45
            }
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    return 752
}

func main() {
    v1 := []int{4294967295, 2147483647, 904, 42}
    var v2 uint8 = 0
    v3 := []uint32{}
    for fi47 := 0; fi47 < 13; fi47++ {
        fmt.Println(fi47, fi47)
        if false {
            for fi48 := 0; fi48 < 1; fi48++ {
                v49 := GenS5{F0: 1, F1: -988}
                if true {
                    v2 = 30
                    fmt.Println(fi47)
                }
                if false {
                    if (!(-1 == 34)) {
                        var v50 uint64 = 0
                        var v51 int16 = 848
                        var v52 float32 = (1e10 + ((0.1 + 0.0) * -2.25))
                        _ = v50
                        _ = v51
                        _ = v52
                    } else {
                        if true {
                            fmt.Println(fi48, v2)
                            if (false && true) {
                                var v53 uint = 1
                                if len(v1) > 0 {
                                    v1[0] = -118
                                }
                                var v54 float64 = (float64(1e10) + (0.1 / 1.0))
                                _ = v53
                                _ = v54
                            }
                            if true {
                                if (4294967295 == 2) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (!true) {
                                        var v55 string = ("" + "")
                                        v2 = (v2 << ((2) & 63))
                                        _ = v55
                                    }
                                    var v56 int16 = genFn1(-28, 42)
                                    _ = v56
                                } else {
                                    v2 += v2
                                    v57 := GenS6{F0: -1}
                                    _ = v57
                                }
                            }
                            v3 = append(v3, 4294967295)
                            if true {
                                var v58 bool = (540 <= 0)
                                _ = v58
                            }
                        }
                        if (42 >= 524) {
                            var v59 int64 = 42
                            _ = v59
                        } else {
                            v60 := GenS7{F0: fi48, F1: false}
                            fmt.Println(fi47, fi47, v2)
                            var v61 uint16 = 0
                            v62 := GenS8{F0: 42}
                            v61 = (v61 - (v61 / ((v61) | 1)))
                            if (!("a" >= ("\n\t" + "a"))) {
                                v63 := GenS4{F0: 108, F1: 42}
                                if false {
                                    v61 += v61
                                }
                                switch (fi48 >> ((42) & 63)) {
                                case -9223372036854775808:
                                    v61 += v61
                                case 23:
                                    v2 += (v2 % ((v2) | 1))
                                case 4294967295:
                                    v61 -= 0
                                }
                                _ = v63
                            }
                            if len(v1) > 0 {
                                v1[0] = fi48
                            }
                            var v64 int32 = 950
                            v65 := GenS1{F0: (fi47 % ((fi48) | 1))}
                            _ = v60
                            _ = v61
                            _ = v62
                            _ = v64
                            _ = v65
                        }
                        _ = len(v1)
                    }
                    var v66 uint32 = 2147483647
                    _ = v66
                } else {
                    switch ((("hello" + "\n\t") + "") + "\n\t") {
                    case "foo bar":
                        var v67 int = (fi47 / ((fi47) | 1))
                        _ = v67
                    }
                    var v68 uint = 255
                    v2 = (v2 % (((^(v2))) | 1))
                    _ = v68
                }
                _ = fi48
                _ = v49
            }
            for fi69 := 0; fi69 < 6; fi69++ {
                if true {
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 = v2
                        v2 += (v2 - (v2 >> ((166) & 63)))
                    }
                    fmt.Println(v2, fi47)
                }
                if len(v1) > 0 {
                    v1[0] = (^(fi47))
                }
                _ = fi69
            }
            v70 := GenS6{F0: 2147483647}
            _ = v70
        } else {
            switch "\n\t" {
            case "":
                v2 -= 255
            default:
                v71 := GenS9{F0: (v2 << ((2) & 63)), F1: (fi47 % (((-(fi47))) | 1))}
                _ = v71
            }
            fmt.Println(v2, v2)
        }
        if (!true) {
            fmt.Println(fi47)
            var v72 int8 = 13
            _ = v72
        }
        _ = fi47
    }
    v2 += 127
    if (true && ((int(v2) <= 94) && false)) {
        fmt.Println(v2)
        _ = len(v3)
        for fi73 := 0; fi73 < 15; fi73++ {
            if (18446744073709551615 < 4294967295) {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                if (!((v2 * (^(v2))) == v2)) {
                    v2 -= 1
                    if len(v1) > 0 {
                        v1[0] = (fi73 >> ((18446744073709551615) & 63))
                    }
                    switch "a" {
                    case "a":
                        v1 = append(v1, (-(fi73)))
                    case "":
                        for fi74 := 0; fi74 < 2; fi74++ {
                            v2 -= (v2 / ((v2) | 1))
                            if false {
                                v75 := GenS6{F0: -2147483648}
                                _ = v75
                            } else {
                                var v76 uint8 = 42
                                if false {
                                    var v77 float32 = float32(-2.25)
                                    v76 += v2
                                    var v78 uint32 = 4294967295
                                    var v79 int8 = -2
                                    v78 += (^(v78))
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                } else {
                                    var v80 float64 = 0.0
                                    fmt.Println(v2, v2, fi73)
                                    if (1 != 127) {
                                        v3 = append(v3, 65535)
                                        v81 := GenS3{F0: -2147483648}
                                        _ = v81
                                    }
                                    _ = v80
                                }
                                if true {
                                    if false {
                                        fmt.Println(v2)
                                        var v82 uint = 18446744073709551615
                                        v2 -= (v2 + v76)
                                        _ = v82
                                    }
                                    if false {
                                        var v83 uint32 = 2147483648
                                        var v84 uint16 = 65535
                                        _ = v83
                                        _ = v84
                                    }
                                    var v85 float64 = 1.5
                                    _ = v85
                                } else {
                                    var v86 uint16 = 1
                                    _ = v86
                                }
                                _ = v76
                            }
                            _ = fi74
                        }
                    default:
                        if (-829 <= -1) {
                            var v87 int32 = 1
                            if ("\n\t" < "") {
                                if (("" >= ("" + "")) && (false || true)) {
                                    v2 = v2
                                    var v88 float32 = 0.1
                                    _ = v88
                                }
                                for fi89 := 0; fi89 < 15; fi89++ {
                                    if len(v1) > 0 {
                                        v1[0] = fi89
                                    }
                                    if false {
                                        var v90 int16 = 32767
                                        _ = v90
                                    }
                                    var v91 int8 = -1
                                    _ = fi89
                                    _ = v91
                                }
                                if false {
                                    v87 = v87
                                    v1 = append(v1, fi73)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            _ = len(v3)
                            _ = v87
                        } else {
                            switch fi73 {
                            case -383:
                                var v92 uint = 284
                                _ = v92
                            case 128:
                                var v93 int16 = 42
                                _ = v93
                            case -1:
                                switch fi73 {
                                case 4294967295:
                                    if len(v1) > 0 {
                                        v1[0] = 1
                                    }
                                }
                            }
                            fmt.Println(fi73, fi73)
                        }
                    }
                }
                fmt.Println(v2)
            } else {
                v3 = append(v3, 65535)
            }
            v2 = v2
            v2 += v2
            _ = fi73
        }
    } else {
        v94 := GenS5{F0: 9223372036854775807, F1: -909}
        var v95 int32 = 2147483647
        var v96 uint8 = v2
        _ = v94
        _ = v95
        _ = v96
    }
    switch 32768 {
    case 2147483647:
        if (("" + (("hello" + ((("\n\t" + ((("foo bar" + "") + "") + "")) + "a") + ("" + ""))) + "hello")) >= "foo bar") {
            var v97 uint8 = v2
            _ = v97
        }
    case 4294967295:
        switch 128 {
        case -1:
            var v98 uint = 804
            _ = v98
        }
    case 0:
        for fi99 := 0; fi99 < 18; fi99++ {
            if true {
                v2 = (v2 & (v2 ^ (v2 - 200)))
            }
            v100 := GenS1{F0: fi99}
            v2 = (v2 * 68)
            _ = fi99
            _ = v100
        }
    default:
        switch 128 {
        case 735:
            v101 := GenS2{F0: 0}
            _ = v101
        default:
            v2 = (v2 | v2)
        }
    }
    if false {
        fmt.Println(v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if ((0.1 <= 3.14159) && true) {
        v102 := GenS9{F0: (v2 & 146), F1: 128}
        var v103 uint32 = 0
        _ = v102
        _ = v103
    }
    var v104 uint16 = 741
    v105 := GenS10{F0: false}
    v104 = v104
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v104
    _ = v105
    _ = v2
    _ = v3
}
