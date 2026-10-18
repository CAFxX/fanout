package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint; F2 bool }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int64; F1 uint }
type GenS10 struct { F0 bool; F1 int16; F2 uint }
func genFn0(p0 string) int {
    v7 := []int{}
    p0 = p0
    _ = p0
    _ = v7
    return 933
}
func genFn1(p0 int8) uint8 {
    switch "hello" {
    case "a":
        fmt.Println(p0)
    case "foo bar":
        var v8 uint16 = 42
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            if false {
                v10 := GenS1{F0: -936}
                for fi11 := 0; fi11 < 4; fi11++ {
                    if true {
                        fmt.Println(fi9, p0, fi9)
                        switch (fi9 + 32767) {
                        case -237:
                            var v12 string = ("hello" + ("" + ""))
                            _ = v12
                        case 32767:
                            fmt.Println(p0, fi9, p0)
                        case 4294967295:
                            fmt.Println(fi11)
                        }
                        switch ("" + ("\n\t" + "foo bar")) {
                        case "a":
                            fmt.Println(p0)
                        default:
                            fmt.Println(fi11, p0)
                        }
                        fmt.Println(p0, fi9)
                    } else {
                        p0 = p0
                    }
                    _ = fi11
                }
                p0 -= (^(p0))
                _ = v10
            } else {
                fmt.Println(p0, fi9)
                if (561 == 1000) {
                    p0 += (p0 / ((p0) | 1))
                    p0 = p0
                    var v13 int32 = -2147483648
                    _ = v13
                }
                fmt.Println(p0, fi9)
            }
            var v14 string = ""
            _ = fi9
            _ = v14
        }
        if true {
            for fi15 := 0; fi15 < 2; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int = (fi15 & fi15)
                    _ = fi16
                    _ = v17
                }
                v18 := GenS2{F0: 953}
                _ = fi15
                _ = v18
            }
            if false {
                if true {
                    switch -596 {
                    case 0:
                        fmt.Println(p0)
                    case 9223372036854775807:
                        fmt.Println(p0)
                    case -1:
                        v19 := GenS1{F0: -32768}
                        _ = v19
                    default:
                        v20 := GenS1{F0: int16(p0)}
                        _ = v20
                    }
                    p0 -= (p0 * p0)
                    p0 = (^(p0))
                } else {
                    v21 := GenS2{F0: 0}
                    var v22 string = ("a" + "foo bar")
                    switch 32767 {
                    case -1:
                        fmt.Println(p0)
                    }
                    _ = v21
                    _ = v22
                }
            }
        } else {
            fmt.Println(p0)
            var v23 uint32 = 4294967295
            p0 -= (p0 * p0)
            _ = v23
        }
    }
    fmt.Println(p0)
    if false {
        v24 := GenS3{F0: -1, F1: 42}
        _ = v24
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint16 = 32767
    var v2 int = 0
    var v3 int64 = -9223372036854775808
    var v4 int16 = int16(v2)
    v5 := make(map[uint32]uint32)
    var v6 float64 = 1e10
    var v25 uint = 255
    v26 := v5[uint32(v25)]
    if (false && (genFn1(100) <= 183)) {
        v26 -= v26
        v27 := GenS1{F0: v4}
        _ = v27
    } else {
        v4 = (v4 % ((v4) | 1))
        var v28 int32 = -1
        _ = v28
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 uint = 145
        if (1 != -1) {
            for fi31 := 0; fi31 < 19; fi31++ {
                v4 = (v4 ^ v4)
                if (42 != 1) {
                    fmt.Println(v2, v6)
                }
                v3 = v3
                _ = fi31
            }
        } else {
            for fi32 := 0; fi32 < 7; fi32++ {
                v33 := GenS2{F0: v1}
                _ = fi32
                _ = v33
            }
            var v34 string = ("foo bar" + "\n\t")
            fmt.Println(v30)
            _ = v34
        }
        if true {
            var v35 int16 = 0
            var v36 int64 = 811
            _ = v35
            _ = v36
        }
        _ = fi29
        _ = v30
    }
    v6 += float64((float32(v6) - (1e10 * 3.14159)))
    if ((v1 * v1) >= (v1 - (v1 % ((42) | 1)))) {
        var v37 float32 = ((float32(v6) - (0.0 / ((float32(v6) * 1.5) - 3.14159))) - ((3.14159 * 1e10) + 1.5))
        _ = v37
    } else {
        if false {
            v38 := GenS4{F0: -3}
            _ = v38
        } else {
            var v39 uint32 = v26
            var v40 string = "foo bar"
            _ = v39
            _ = v40
        }
    }
    var v41 int8 = 1
    fmt.Println(v3)
    v2 += (^(v2))
    if (true || false) {
        if true {
            if (v41 != (v41 << ((v25) & 63))) {
                fmt.Println(v25, v3, v25)
            }
            if (!(!false)) {
                for fi42 := 0; fi42 < 9; fi42++ {
                    v43 := GenS5{F0: "\n\t", F1: (v25 + 1), F2: false}
                    var v44 float32 = (float32((v6 / (v6 / v6))) * -2.25)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
            if false {
                if (v6 == (((3.14159 - -2.25) - v6) / (((0.0 + (v6 * (1.5 - (v6 * ((v6 + ((float64(1.5) - v6) * v6)) + v6))))) / ((v6 + v6) - -2.25)) - 1.5))) {
                    if (!true) {
                        if ((2 != (v26 >> (((v25 >> (((v25 - v25)) & 63))) & 63))) && true) {
                            for fi45 := 0; fi45 < 14; fi45++ {
                                if (v6 == (v6 / 1.5)) {
                                    _ = len(v5)
                                    switch ("" + "a") {
                                    case "":
                                        fmt.Println(fi45)
                                    default:
                                        switch (-(v2)) {
                                        case -852:
                                            var v46 float64 = v6
                                            _ = v46
                                        }
                                    }
                                    var v47 string = "hello"
                                    _ = v47
                                } else {
                                    for fi48 := 0; fi48 < 8; fi48++ {
                                        if true {
                                            v49 := GenS6{F0: 255, F1: ("\n\t" + "hello")}
                                            _ = v49
                                        }
                                        switch "a" {
                                        case "hello":
                                            v50 := GenS5{F0: "hello", F1: (v25 - v25), F2: true}
                                            _ = v50
                                        case "":
                                            var v51 uint = v25
                                            _ = v51
                                        case "a":
                                            v52 := []int{32768, (fi48 & v2)}
                                            _ = v52
                                        default:
                                            switch -1 {
                                            case -2147483648:
                                                v53 := []int{fi45, fi48, (fi48 >> ((v25) & 63)), 32768}
                                                _ = v53
                                            case 128:
                                                v54 := GenS7{F0: (v41 << ((v25) & 63)), F1: -78}
                                                if true {
                                                    v55 := GenS8{F0: int32(v2), F1: fi48}
                                                    if ((v4 << (((^(v25))) & 63)) <= -946) {
                                                        v4 = (v4 ^ -1)
                                                        if true {
                                                            var v56 uint = v25
                                                            v4 += (v4 >> (((^(v56))) & 63))
                                                            _ = v56
                                                        }
                                                    } else {
                                                        fmt.Println(fi48, v41)
                                                        v57 := []int{(fi48 - genFn0("hello"))}
                                                        v58 := GenS7{F0: (v41 << ((uint(v1)) & 63)), F1: 0}
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                    _ = v55
                                                } else {
                                                    v6 = v6
                                                }
                                                v59 := GenS5{F0: (("foo bar" + "\n\t") + "hello"), F1: uint(v26), F2: true}
                                                _ = v54
                                                _ = v59
                                            case 841:
                                                var v60 int = (-(v2))
                                                _ = v60
                                            default:
                                                var v61 uint = v25
                                                _ = v61
                                            }
                                        }
                                        v41 = v41
                                        _ = fi48
                                    }
                                }
                                v62 := GenS9{F0: 9223372036854775807, F1: (v25 + (v25 | 157))}
                                var v63 uint = 0
                                _ = fi45
                                _ = v62
                                _ = v63
                            }
                        }
                        v41 += v41
                    } else {
                        v64 := GenS9{F0: v3, F1: v25}
                        _ = v64
                    }
                    switch "" {
                    case "":
                        for fi65 := 0; fi65 < 14; fi65++ {
                            switch 2147483647 {
                            case -238:
                                var v66 int32 = 42
                                _ = v66
                            case 9223372036854775807:
                                if true {
                                    v67 := []int{}
                                    fmt.Println(v2, v1, v2)
                                    _ = v67
                                } else {
                                    if ((!true) || true) {
                                        for fi68 := 0; fi68 < 11; fi68++ {
                                            var v69 int16 = (v4 + -636)
                                            if true {
                                                if ((((0.0 - (-2.25 * float32(v6))) - 1.5) > 1e10) && (-552 < 42)) {
                                                    fmt.Println(v69)
                                                    var v70 int8 = v41
                                                    _ = v70
                                                }
                                                var v71 string = "a"
                                                _ = v71
                                            }
                                            v3 = v3
                                            _ = fi68
                                            _ = v69
                                        }
                                    }
                                }
                            case 32768:
                                var v72 string = (("a" + (("foo bar" + "foo bar") + "hello")) + ("" + ("foo bar" + "")))
                                _ = v72
                            }
                            _ = fi65
                        }
                    case "foo bar":
                        if (4294967295 >= 195) {
                            v73 := []int{(^(v2)), v2}
                            if len(v73) >= 2 {
                                v73 = v73[1:len(v73)]
                            }
                            _ = v73
                        } else {
                            if (!true) {
                                v4 = v4
                                switch "foo bar" {
                                case "foo bar":
                                    for fi74 := 0; fi74 < 11; fi74++ {
                                        switch v2 {
                                        case -949:
                                            var v75 int64 = int64(v26)
                                            _ = v75
                                        case 2147483647:
                                            var v76 uint64 = 4294967295
                                            _ = v76
                                        case 128:
                                            fmt.Println(v26, fi74)
                                        default:
                                            v25 = (v25 / ((v25) | 1))
                                        }
                                        _ = fi74
                                    }
                                case "hello":
                                    for fi77 := 0; fi77 < 12; fi77++ {
                                        switch fi77 {
                                        case -2147483648:
                                            var v78 uint16 = v1
                                            _ = v78
                                        case -234:
                                            v26 = v26
                                        case 206:
                                            switch 2147483647 {
                                            case 4294967295:
                                                v3 += (v3 ^ -1)
                                            case 128:
                                                if ((-(v4)) >= v4) {
                                                    var v79 bool = ((38 >= (^(v41))) && (true || (false || true)))
                                                    switch genFn0("") {
                                                    case -27:
                                                        v6 = v6
                                                    case 827:
                                                        v80 := []int{}
                                                        _ = v80
                                                    }
                                                    _ = v79
                                                }
                                            default:
                                                v41 = (v41 + v41)
                                            }
                                        }
                                        _ = fi77
                                    }
                                case "\n\t":
                                    for fi81 := 0; fi81 < 8; fi81++ {
                                        var v82 int16 = (-(v4))
                                        var v83 int32 = 446
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                    }
                                default:
                                    v84 := []int{(v2 ^ (v2 / ((int(v26)) | 1)))}
                                    _ = v84
                                }
                            } else {
                                _ = len(v5)
                                if true {
                                    if ((v26 - (v26 >> ((876) & 63))) <= (v26 | 2147483648)) {
                                        var v85 uint16 = 42
                                        var v86 uint16 = v85
                                        _ = v85
                                        _ = v86
                                    }
                                    v87 := GenS6{F0: genFn1(v41), F1: "hello"}
                                    _ = v87
                                }
                                var v88 int8 = 12
                                _ = v88
                            }
                            var v89 bool = (true && (!false))
                            _ = v89
                        }
                    case "a":
                        delete(v5, v26)
                    default:
                        v90 := []int{32767, 2147483647, 884, genFn0("foo bar")}
                        _ = v90
                    }
                    v91 := GenS5{F0: "", F1: v25, F2: ((v3 / ((1) | 1)) != 42)}
                    _ = v91
                } else {
                    v92 := []int{}
                    if (genFn1(0) > 127) {
                        if (int64(v2) >= (v3 - v3)) {
                            switch (v2 - v2) {
                            case 32768:
                                for fi93 := 0; fi93 < 11; fi93++ {
                                    var v94 uint64 = 42
                                    switch "" {
                                    case "":
                                        _ = len(v92)
                                    case "hello":
                                        v2 = int(v3)
                                    }
                                    _ = fi93
                                    _ = v94
                                }
                            case 32767:
                                v2 = (v2 - (v2 % (((-(v2))) | 1)))
                            case -334:
                                for fi95 := 0; fi95 < 18; fi95++ {
                                    v3 -= 1
                                    v92 = append(v92, (^(v2)))
                                    for fi96 := 0; fi96 < 2; fi96++ {
                                        var v97 uint16 = (v1 * (v1 % ((42) | 1)))
                                        v2 -= fi96
                                        var v98 string = ("" + "")
                                        fmt.Println(v97, v1)
                                        var v99 int64 = (v3 & 733)
                                        _ = fi96
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                    _ = fi95
                                }
                            }
                        } else {
                            v92 = append(v92, -860)
                        }
                        v4 = int16(v3)
                        for fi100 := 0; fi100 < 12; fi100++ {
                            var v101 uint16 = v1
                            fmt.Println(v101)
                            _ = fi100
                            _ = v101
                        }
                    } else {
                        var v102 bool = (!((v26 - v26) < (^(v26))))
                        _ = v102
                    }
                    _ = v92
                }
            } else {
                v3 -= v3
                var v103 int32 = -1
                var v104 uint8 = 42
                _ = v103
                _ = v104
            }
        }
        var v105 uint64 = 0
        if false {
            fmt.Println(v26)
            if false {
                if (!true) {
                    switch "\n\t" {
                    case "foo bar":
                        v106 := GenS10{F0: false, F1: 42, F2: 65535}
                        _ = v106
                    default:
                        switch (v2 * genFn0((("hello" + ("a" + "")) + ""))) {
                        case 175:
                            delete(v5, 42)
                        case 32768:
                            if true {
                                v26 -= (v26 >> ((65535) & 63))
                            } else {
                                if false {
                                    for fi107 := 0; fi107 < 2; fi107++ {
                                        v26 = v26
                                        if (!(!false)) {
                                            v6 += (v6 / 1e10)
                                            var v108 bool = ((0.0 - 0.0) < 3.14159)
                                            _ = v108
                                        } else {
                                            if false {
                                                if (v25 == v25) {
                                                    var v109 int8 = (v41 >> (((v25 / ((v25) | 1))) & 63))
                                                    _ = v109
                                                }
                                                if (true || true) {
                                                    v26 -= (v26 & 1)
                                                    v105 = v105
                                                } else {
                                                    var v110 int = v2
                                                    _ = v110
                                                }
                                            } else {
                                                var v111 int16 = v4
                                                for fi112 := 0; fi112 < 4; fi112++ {
                                                    var v113 float32 = 1.5
                                                    if true {
                                                        var v114 uint16 = v1
                                                        fmt.Println(fi107, v2)
                                                        v113 += v113
                                                        switch ("\n\t" + (("\n\t" + "foo bar") + "a")) {
                                                        case "hello":
                                                            v115 := GenS9{F0: (v3 + -9223372036854775808), F1: (v25 + uint(v26))}
                                                            _ = v115
                                                        default:
                                                            v116 := []int{v2, 190, fi107}
                                                            _ = v116
                                                        }
                                                        _ = v114
                                                    }
                                                    _ = fi112
                                                    _ = v113
                                                }
                                                _ = v111
                                            }
                                            switch (fi107 >> ((v25) & 63)) {
                                            case 164:
                                                fmt.Println(v3, v41, fi107)
                                            case 2147483647:
                                                delete(v5, v26)
                                            case 127:
                                                for fi117 := 0; fi117 < 11; fi117++ {
                                                    v118 := []int{}
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v119 uint = (v25 * (v25 | (v25 % ((v25) | 1))))
                                                        _ = v119
                                                    case "a":
                                                        var v120 uint64 = (v105 - v105)
                                                        _ = v120
                                                    case "":
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v121 float32 = ((0.1 * 1e10) + 0.1)
                                                            _ = v121
                                                        case "foo bar":
                                                            if (!("\n\t" == (("" + ("a" + "\n\t")) + "hello"))) {
                                                                var v122 uint8 = 1
                                                                _ = v122
                                                            }
                                                        default:
                                                            var v123 bool = (false || true)
                                                            _ = v123
                                                        }
                                                    default:
                                                        fmt.Println(fi117, v4)
                                                    }
                                                    v41 = (v41 % ((int8(v4)) | 1))
                                                    _ = fi117
                                                    _ = v118
                                                }
                                            }
                                        }
                                        _ = fi107
                                    }
                                }
                                var v124 float64 = v6
                                _ = v124
                            }
                        }
                    }
                    var v125 int8 = -1
                    _ = v125
                }
                switch v2 {
                case 32767:
                    v4 = v4
                case 9223372036854775807:
                    var v126 float64 = v6
                    _ = v126
                case 1:
                    var v127 int = genFn0("a")
                    _ = v127
                }
            } else {
                v6 -= (float64((-2.25 + 0.1)) * v6)
                delete(v5, uint32(v1))
                v3 = v3
            }
        }
        _ = v105
    } else {
        if (!false) {
            fmt.Println(v1)
            var v128 bool = true
            _ = v128
        } else {
            var v129 uint32 = (^(v26))
            var v130 uint = 255
            _ = v129
            _ = v130
        }
    }
    v3 = v3
    v3 += v3
    v25 += (v25 - v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
