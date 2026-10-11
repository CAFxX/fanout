package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int64; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint8; F1 bool; F2 int64 }
type GenS11 struct { F0 int32; F1 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string; F1 int16 }
type GenS14 struct { F0 uint16; F1 int32 }
type GenS15 struct { F0 uint8; F1 int; F2 uint64 }
type GenS16 struct { F0 string; F1 int; F2 int }
func genFn0(p0 int, p1 uint16, p2 string) uint32 {
    var v5 uint = 1
    switch 191 {
    case -2147483648:
        var v6 float32 = 1.5
        _ = v6
    case 973:
        var v7 bool = (105 != 48)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = float32(1e10)
    var v3 uint8 = 2
    var v4 uint8 = v3
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{fi8, (-(fi8))}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := []int{fi10}
        v12 := v1[(-(fi10))]
        var v13 float64 = 1e10
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        if (!((565 <= 485) && (-1 != 2147483647))) {
            switch "foo bar" {
            case "hello":
                v15 := GenS1{F0: 0, F1: (!(-9223372036854775808 == 42)), F2: 109}
                _ = v15
            }
            v16 := GenS1{F0: 916, F1: false, F2: 143}
            fmt.Println(v2)
            _ = v16
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := GenS1{F0: -32768, F1: (3.14159 == -2.25), F2: (v4 >> ((4294967295) & 63))}
            var v19 int32 = -542
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
    }
    var v20 int8 = 42
    var v21 uint32 = 2147483648
    switch 32768 {
    case -1:
        var v22 int32 = 2147483647
        _ = v22
    case 128:
        v23 := GenS2{F0: 886, F1: (v21 & v21)}
        _ = v23
    case 0:
        v3 += 42
    default:
        var v24 uint = 0
        _ = v24
    }
    if ((!true) && (int16(v21) > -365)) {
        var v25 uint32 = v21
        v26 := GenS3{F0: ("\n\t" + (("foo bar" + "hello") + "a")), F1: 455}
        _ = v25
        _ = v26
    } else {
        if ((false || false) && (false || true)) {
            for fi27 := 0; fi27 < 3; fi27++ {
                if false {
                    var v28 int32 = 2147483647
                    for fi29 := 0; fi29 < 18; fi29++ {
                        _ = len(v1)
                        _ = fi29
                    }
                    var v30 int8 = (v20 << ((761) & 63))
                    _ = v28
                    _ = v30
                } else {
                    v31 := GenS3{F0: "", F1: 4294967295}
                    var v32 uint = 65535
                    v33 := GenS4{F0: "foo bar"}
                    _ = v31
                    _ = v32
                    _ = v33
                }
                v3 += (v3 + 128)
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            var v34 uint8 = v4
                            var v35 float32 = (float32(0.0) / -2.25)
                            _ = v34
                            _ = v35
                        }
                    case "hello":
                        if (!(!false)) {
                            var v36 float32 = 0.1
                            for fi37 := 0; fi37 < 5; fi37++ {
                                switch ("hello" + ("a" + "a")) {
                                case "":
                                    if true {
                                        fmt.Println(fi37, v4, v3)
                                        v38 := GenS3{F0: "", F1: 18446744073709551615}
                                        switch fi37 {
                                        case 731:
                                            v39 := GenS5{F0: -57, F1: -1, F2: (v20 ^ (v20 + -93))}
                                            _ = v39
                                        default:
                                            v3 += (^(v3))
                                            v40 := []int{fi37, (fi27 - (-(fi37)))}
                                            switch (fi27 << ((42) & 63)) {
                                            case -560:
                                                v41 := GenS1{F0: int16(v4), F1: true, F2: 42}
                                                _ = v41
                                            case 4294967295:
                                                _ = len(v40)
                                            case -9223372036854775808:
                                                v20 = (v20 | v20)
                                            default:
                                                v42 := v1[(fi27 >> ((18446744073709551615) & 63))]
                                                _ = v42
                                            }
                                            _ = v40
                                        }
                                        var v43 int = (fi27 | fi27)
                                        v43 = (v43 + fi37)
                                        var v44 float32 = ((-2.25 - v36) + v2)
                                        v4 -= v3
                                        v45 := GenS6{F0: 4294967295}
                                        if (!(2147483647 <= 42)) {
                                            v46 := GenS3{F0: "foo bar", F1: 255}
                                            _ = v46
                                        } else {
                                            v47 := GenS5{F0: v20, F1: -933, F2: v20}
                                            var v48 uint32 = 983
                                            _ = v47
                                            _ = v48
                                        }
                                        _ = v38
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                case "foo bar":
                                    v2 -= v36
                                default:
                                    if (128 < fi27) {
                                        switch fi37 {
                                        case -1:
                                            if true {
                                                var v49 string = "a"
                                                switch int(v4) {
                                                case -143:
                                                    var v50 int64 = 9223372036854775807
                                                    _ = v50
                                                case 4294967295:
                                                    var v51 bool = (v49 > (v49 + v49))
                                                    _ = v51
                                                case -1:
                                                    v52 := []int{fi27, fi37}
                                                    _ = v52
                                                default:
                                                    if (1 != int64(fi27)) {
                                                        fmt.Println(v21)
                                                    }
                                                }
                                                _ = v49
                                            }
                                        default:
                                            var v53 uint16 = 2
                                            if (18446744073709551615 > 18446744073709551615) {
                                                v36 = v2
                                                fmt.Println(v53, v3, v4)
                                            }
                                            _ = v53
                                        }
                                        fmt.Println(v2, v2)
                                    } else {
                                        v3 += v3
                                        var v54 int32 = 710
                                        if false {
                                            v55 := GenS7{F0: (v21 * (v21 + uint32(v4)))}
                                            _ = v55
                                        } else {
                                            if (fi37 > fi37) {
                                                fmt.Println(v4, v54)
                                                v2 -= (((v36 + float32(0.0)) / 0.1) + v2)
                                                v54 = 2147483647
                                            } else {
                                                v56 := []int{fi37, -9223372036854775808, fi27, (fi27 << ((255) & 63))}
                                                switch "a" {
                                                case "":
                                                    v4 = v4
                                                default:
                                                    v57 := GenS6{F0: (fi37 >> ((4294967295) & 63))}
                                                    _ = v57
                                                }
                                                v20 = v20
                                                _ = v56
                                            }
                                            fmt.Println(v54, v54)
                                            v58 := v1[(fi37 | fi27)]
                                            _ = v58
                                        }
                                        _ = v54
                                    }
                                }
                                v3 -= uint8(v21)
                                var v59 float32 = (0.1 / 1.0)
                                _ = fi37
                                _ = v59
                            }
                            var v60 float32 = (v2 * v2)
                            _ = v36
                            _ = v60
                        }
                    default:
                        fmt.Println(v4, v4, v20)
                    }
                    v61 := GenS2{F0: uint16(v4), F1: 65535}
                    if false {
                        if true {
                            v62 := GenS5{F0: 100, F1: 1, F2: 42}
                            v63 := []int{fi27}
                            _ = v62
                            _ = v63
                        } else {
                            for fi64 := 0; fi64 < 2; fi64++ {
                                var v65 string = ("\n\t" + (("hello" + "") + "hello"))
                                _ = fi64
                                _ = v65
                            }
                            _ = len(v1)
                            if (v4 >= (v4 + (v4 / (((v4 % (((v4 % (((v3 * v3)) | 1))) | 1))) | 1)))) {
                                v66 := GenS4{F0: ((("hello" + ("hello" + "foo bar")) + "\n\t") + "a")}
                                fmt.Println(v4, v3, v4)
                                for fi67 := 0; fi67 < 5; fi67++ {
                                    fmt.Println(v4, v3)
                                    fmt.Println(v20, v3, fi27)
                                    var v68 uint = 65535
                                    _ = fi67
                                    _ = v68
                                }
                                _ = v66
                            }
                        }
                        if true {
                            v2 += v2
                            var v69 uint16 = 1
                            v70 := []int{}
                            _ = v69
                            _ = v70
                        }
                    } else {
                        var v71 int32 = 0
                        fmt.Println(v3)
                        _ = v71
                    }
                    _ = v61
                } else {
                    v72 := GenS4{F0: ("\n\t" + "hello")}
                    _ = v72
                }
                _ = fi27
            }
            fmt.Println(v2)
        }
        var v73 uint16 = 401
        v20 += (-(v20))
        _ = v73
    }
    v3 = 255
    switch 1 {
    case 720:
        if true {
            for fi74 := 0; fi74 < 4; fi74++ {
                v20 -= (^(v20))
                for fi75 := 0; fi75 < 9; fi75++ {
                    if (1e10 >= 0.0) {
                        if (!((v3 >> ((216) & 63)) == (v3 - 1))) {
                            switch (fi74 * (fi75 + fi75)) {
                            case -378:
                                if true {
                                    if true {
                                        if false {
                                            v20 = v20
                                            v4 -= (v4 + v3)
                                            fmt.Println(v20, fi75, v3)
                                        }
                                    } else {
                                        if (false && ((!((-1 != 1000) || true)) && false)) {
                                            v2 += 0.0
                                            if true {
                                                v4 -= v3
                                                if (((9223372036854775807 != -914) || false) && true) {
                                                    _ = len(v1)
                                                    if false {
                                                        switch (^(fi74)) {
                                                        case 42:
                                                            switch (fi75 << ((uint(v21)) & 63)) {
                                                            case 128:
                                                                _ = len(v1)
                                                                if (2147483647 == 467) {
                                                                    if true {
                                                                        var v76 bool = (!true)
                                                                        var v77 uint64 = 0
                                                                        _ = v76
                                                                        _ = v77
                                                                    } else {
                                                                        if (v20 != (-(v20))) {
                                                                            var v78 int8 = (v20 << ((35) & 63))
                                                                            _ = v78
                                                                        } else {
                                                                            if true {
                                                                                v3 += (v3 % ((42) | 1))
                                                                                v21 += (v21 >> ((uint(v4)) & 63))
                                                                                switch fi74 {
                                                                                case 0:
                                                                                    var v79 float32 = (((3.14159 * v2) + v2) + (float32(1.5) - (((v2 / (v2 * v2)) / v2) / 1.0)))
                                                                                    _ = v79
                                                                                }
                                                                                var v80 float64 = (((1.5 + 3.14159) * 1.5) - (1.5 / ((1e10 / -2.25) - (1e10 * ((1e10 * 0.0) * 1.5)))))
                                                                                v21 = v21
                                                                                var v81 string = ("" + "foo bar")
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                v82 := []int{}
                                                                                switch fi74 {
                                                                                case 42:
                                                                                    var v83 uint = 358
                                                                                    _ = v83
                                                                                case 603:
                                                                                    if (("\n\t" + (("" + "foo bar") + ((("hello" + "foo bar") + "a") + "hello"))) >= "a") {
                                                                                        v84 := GenS5{F0: 100, F1: 700, F2: v20}
                                                                                        var v85 int16 = 32767
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    } else {
                                                                                        var v86 string = "hello"
                                                                                        _ = v86
                                                                                    }
                                                                                case -2147483648:
                                                                                    var v87 int16 = 502
                                                                                    _ = v87
                                                                                }
                                                                                var v88 uint32 = (v21 << ((4294967295) & 63))
                                                                                _ = v82
                                                                                _ = v88
                                                                            }
                                                                            v89 := []int{fi74, fi75, (fi74 * -821)}
                                                                            if false {
                                                                                fmt.Println(v20, v4)
                                                                            } else {
                                                                                if true {
                                                                                    var v90 uint16 = uint16(v4)
                                                                                    _ = v90
                                                                                }
                                                                                if true {
                                                                                    if (32768 != 42) {
                                                                                        v21 = uint32(v4)
                                                                                        var v91 uint = 322
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            v2 -= v2
                                                                                        default:
                                                                                            fmt.Println(v20, v4, v3)
                                                                                        }
                                                                                        _ = v91
                                                                                    }
                                                                                }
                                                                                switch "a" {
                                                                                case "":
                                                                                    var v92 uint = 42
                                                                                    _ = v92
                                                                                }
                                                                            }
                                                                            _ = v89
                                                                        }
                                                                        var v93 int64 = 777
                                                                        var v94 string = ("hello" + (("\n\t" + "foo bar") + "hello"))
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                    v20 -= v20
                                                                } else {
                                                                    var v95 int16 = -32768
                                                                    switch "" {
                                                                    case "hello":
                                                                        var v96 uint32 = (^(v21))
                                                                        _ = v96
                                                                    }
                                                                    _ = v95
                                                                }
                                                                if true {
                                                                    switch "" {
                                                                    case "":
                                                                        switch fi74 {
                                                                        case 42:
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                var v97 uint = 407
                                                                                _ = v97
                                                                            default:
                                                                                _ = len(v1)
                                                                            }
                                                                        default:
                                                                            var v98 int64 = 42
                                                                            _ = v98
                                                                        }
                                                                    case "foo bar":
                                                                        v21 = (^(v21))
                                                                    case "a":
                                                                        v2 = float32((0.1 + 1e10))
                                                                    }
                                                                    v99 := GenS7{F0: genFn0((fi75 | -1), 65535, ("\n\t" + "\n\t"))}
                                                                    v20 = (v20 + (v20 << ((502) & 63)))
                                                                    _ = v99
                                                                } else {
                                                                    switch 695 {
                                                                    case 9223372036854775807:
                                                                        if true {
                                                                            switch "" {
                                                                            case "a":
                                                                                v20 -= (v20 | v20)
                                                                            case "hello":
                                                                                v3 += 85
                                                                            case "foo bar":
                                                                                v3 += v3
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                var v100 float64 = 1.5
                                                                                _ = v100
                                                                            }
                                                                            v1[(-(fi74))] = "hello"
                                                                        }
                                                                    }
                                                                }
                                                            case -5:
                                                                v101 := []int{fi74, 642, fi74}
                                                                _ = v101
                                                            }
                                                        }
                                                        if true {
                                                            fmt.Println(v20)
                                                            v102 := GenS4{F0: ("\n\t" + "foo bar")}
                                                            v20 -= (v20 * (v20 + (v20 / ((127) | 1))))
                                                            _ = v102
                                                        } else {
                                                            if (!true) {
                                                                if (!((-2.25 + 0.0) >= 0.0)) {
                                                                    v3 -= (v4 & (^(v3)))
                                                                    fmt.Println(v2, v4)
                                                                }
                                                                var v103 float32 = v2
                                                                switch (fi74 + (^(fi75))) {
                                                                case 529:
                                                                    if ((int16(fi74) == 1) || false) {
                                                                        switch fi75 {
                                                                        case 32768:
                                                                            v103 = v103
                                                                        default:
                                                                            var v104 int8 = (v20 - v20)
                                                                            _ = v104
                                                                        }
                                                                        fmt.Println(v21, v3)
                                                                    } else {
                                                                        switch fi75 {
                                                                        case 32767:
                                                                            if (int32(v21) < int32(v3)) {
                                                                                v3 = uint8(v21)
                                                                                v105 := GenS7{F0: v21}
                                                                                fmt.Println(v103, fi75)
                                                                                _ = v105
                                                                            }
                                                                        case 0:
                                                                            v4 -= (v4 - v3)
                                                                        case -9223372036854775808:
                                                                            if (true || true) {
                                                                                if (4294967295 < 18446744073709551615) {
                                                                                    var v106 int = fi75
                                                                                    v107 := GenS8{F0: v20}
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                }
                                                                                var v108 int = fi75
                                                                                var v109 int64 = 1
                                                                                switch -146 {
                                                                                case -2147483648:
                                                                                    var v110 float64 = 1e10
                                                                                    _ = v110
                                                                                case -9223372036854775808:
                                                                                    v4 -= (v3 / ((127) | 1))
                                                                                case 68:
                                                                                    var v111 uint16 = 255
                                                                                    _ = v111
                                                                                default:
                                                                                    var v112 uint32 = genFn0(v108, 1, ("a" + "\n\t"))
                                                                                    _ = v112
                                                                                }
                                                                                _ = v108
                                                                                _ = v109
                                                                            }
                                                                        }
                                                                        v21 = 89
                                                                        var v113 uint = 4294967295
                                                                        var v114 int64 = -9223372036854775808
                                                                        _ = v113
                                                                        _ = v114
                                                                    }
                                                                case -2147483648:
                                                                    v115 := GenS3{F0: "", F1: 1}
                                                                    _ = v115
                                                                case 966:
                                                                    _ = len(v1)
                                                                default:
                                                                    v116 := GenS5{F0: -128, F1: 9223372036854775807, F2: (v20 | 100)}
                                                                    _ = v116
                                                                }
                                                                _ = v103
                                                            } else {
                                                                var v117 int32 = 684
                                                                v118 := []int{}
                                                                v119 := GenS7{F0: v21}
                                                                _ = v117
                                                                _ = v118
                                                                _ = v119
                                                            }
                                                            fmt.Println(v2)
                                                        }
                                                        var v120 uint64 = 42
                                                        v121 := GenS9{F0: 2147483647}
                                                        _ = v120
                                                        _ = v121
                                                    }
                                                    v4 = (v3 % (((v4 | 2)) | 1))
                                                }
                                            }
                                            v122 := []int{}
                                            if len(v122) >= 2 {
                                                v122 = v122[1:len(v122)]
                                            }
                                            v3 = (v3 * 255)
                                            _ = v122
                                        } else {
                                            v21 = (^(v21))
                                            var v123 uint = 4294967295
                                            _ = v123
                                        }
                                    }
                                }
                            case -99:
                                v124 := []int{fi75, fi74, (^(fi75))}
                                _ = v124
                            case 1:
                                var v125 int16 = -32768
                                _ = v125
                            default:
                                var v126 float32 = v2
                                _ = v126
                            }
                            if (-2147483648 != fi74) {
                                var v127 uint = 4294967295
                                if true {
                                    var v128 uint16 = uint16(v127)
                                    _ = v128
                                } else {
                                    v4 += v4
                                    var v129 uint64 = 1
                                    v21 -= (v21 - genFn0(fi74, 32767, ""))
                                    _ = v129
                                }
                                delete(v1, (fi75 << ((0) & 63)))
                                _ = v127
                            }
                        } else {
                            var v130 float32 = -2.25
                            _ = v130
                        }
                        var v131 int = (fi75 * fi75)
                        _ = v131
                    } else {
                        v4 = v3
                    }
                    switch (fi75 - (fi75 % (((fi75 ^ fi74)) | 1))) {
                    case 127:
                        v20 = 127
                    }
                    if false {
                        v132 := GenS6{F0: (fi74 - fi75)}
                        if false {
                            if (!((v20 % (((v20 & v20)) | 1)) < (v20 >> ((4294967295) & 63)))) {
                                v133 := GenS10{F0: 160, F1: ((fi74 * 32767) < (fi74 - fi75)), F2: 0}
                                if (-32768 <= 1000) {
                                    var v134 int8 = (v20 ^ (v20 - 127))
                                    if ((fi74 ^ fi75) < 347) {
                                        if (643 != v21) {
                                            switch fi74 {
                                            case 9223372036854775807:
                                                if false {
                                                    var v135 string = "foo bar"
                                                    if (v21 <= (v21 << ((65535) & 63))) {
                                                        var v136 int16 = 1000
                                                        v137 := GenS8{F0: v20}
                                                        _ = v136
                                                        _ = v137
                                                    }
                                                    v20 += (v20 / (((v134 % ((v134) | 1))) | 1))
                                                    v4 -= (v4 << ((483) & 63))
                                                    v138 := GenS6{F0: (fi75 * (fi74 << ((227) & 63)))}
                                                    _ = v135
                                                    _ = v138
                                                } else {
                                                    var v139 float32 = -2.25
                                                    switch (fi74 ^ -1) {
                                                    case 877:
                                                        v140 := GenS2{F0: 42, F1: (v21 | (v21 - 2147483648))}
                                                        _ = v140
                                                    case 42:
                                                        v141 := GenS11{F0: 42, F1: (v4 % ((v4) | 1))}
                                                        _ = v141
                                                    }
                                                    _ = v139
                                                }
                                            case 32768:
                                                var v142 int64 = -9223372036854775808
                                                _ = v142
                                            case 2147483647:
                                                v1[(-(fi74))] = ""
                                            }
                                            fmt.Println(v3, v3, v2)
                                            if (!(fi74 < (^(fi75)))) {
                                                var v143 uint16 = 794
                                                v3 = (v4 / ((255) | 1))
                                                v144 := GenS9{F0: genFn0(fi75, 1, "\n\t")}
                                                var v145 float32 = v2
                                                _ = v143
                                                _ = v144
                                                _ = v145
                                            } else {
                                                v21 += v21
                                                v134 -= v20
                                            }
                                            v4 -= v4
                                            var v146 bool = false
                                            var v147 float32 = (1.5 - (1e10 + v2))
                                            v146 = true
                                            _ = v146
                                            _ = v147
                                        }
                                    } else {
                                        v148 := v1[(^(fi74))]
                                        _ = v148
                                    }
                                    _ = v134
                                } else {
                                    _ = len(v1)
                                    v149 := GenS2{F0: 1, F1: (v21 * genFn0((fi74 | fi74), 32767, ""))}
                                    _ = v149
                                }
                                if ((!(1.5 <= -2.25)) && (!(65535 < (^(v21))))) {
                                    v2 = 0.0
                                } else {
                                    if (!false) {
                                        var v150 float64 = 1.5
                                        v151 := GenS9{F0: (v21 | genFn0(fi75, 65535, ""))}
                                        _ = v150
                                        _ = v151
                                    }
                                    var v152 int64 = 842
                                    var v153 int16 = 177
                                    _ = v152
                                    _ = v153
                                }
                                _ = v133
                            } else {
                                v1[9223372036854775807] = ("foo bar" + (("a" + "a") + ("\n\t" + "\n\t")))
                                if false {
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (!false) {
                                            v154 := GenS12{F0: ("a" + (("\n\t" + "foo bar") + ("" + "\n\t")))}
                                            _ = v154
                                        }
                                        if ((65535 >= 255) && true) {
                                            v20 = 3
                                        }
                                        v20 += (v20 | 79)
                                    }
                                }
                                v155 := GenS13{F0: ("hello" + "foo bar"), F1: 0}
                                _ = v155
                            }
                            var v156 uint8 = (v3 | 127)
                            if ((fi74 + (fi74 | (fi74 % ((fi74) | 1)))) <= 32768) {
                                v4 = v4
                                var v157 float64 = ((float64(v2) + (float64(v2) / float64(1.5))) - ((float64(1.5) + 0.1) - 0.0))
                                v2 = v2
                                _ = v157
                            } else {
                                fmt.Println(fi75)
                            }
                            if (1.5 >= (v2 + v2)) {
                                fmt.Println(v20)
                            }
                            v2 += (((float32(0.0) + v2) - ((v2 - v2) + v2)) - v2)
                            _ = v156
                        }
                        _ = v132
                    } else {
                        v158 := []int{(fi74 ^ (fi75 - fi75)), (fi75 | (fi74 & -2147483648)), (^(fi75)), (fi74 / ((-97) | 1))}
                        var v159 bool = (0 == 18446744073709551615)
                        _ = v158
                        _ = v159
                    }
                    _ = fi75
                }
                _ = fi74
            }
        } else {
            v4 = (v4 - 2)
            for fi160 := 0; fi160 < 2; fi160++ {
                var v161 int8 = 0
                for fi162 := 0; fi162 < 4; fi162++ {
                    v163 := GenS14{F0: 65535, F1: 2147483647}
                    _ = fi162
                    _ = v163
                }
                _ = fi160
                _ = v161
            }
        }
    case -329:
        if false {
            v164 := GenS15{F0: (^(v4)), F1: -103, F2: 1}
            _ = v164
        } else {
            delete(v1, 2147483647)
            switch "" {
            case "a":
                if (9223372036854775807 >= 42) {
                    v20 -= 1
                }
            }
            v21 -= genFn0(2147483647, 255, "hello")
        }
    case 0:
        v165 := GenS14{F0: 854, F1: 825}
        _ = v165
    }
    var v166 uint64 = 0
    switch 9223372036854775807 {
    case 32768:
        _ = len(v1)
    case -2147483648:
        delete(v1, -30)
    }
    v167 := []int{-588, 4294967295, 4294967295}
    fmt.Println(v166)
    v168 := GenS16{F0: "foo bar", F1: -954, F2: -2147483648}
    var v169 string = "hello"
    switch 32767 {
    case 32767:
        v1[-2147483648] = (v169 + v169)
    default:
        for fi170 := 0; fi170 < 17; fi170++ {
            _ = len(v1)
            _ = fi170
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v166:", v166)
    fmt.Println("slicelen v167:", len(v167), v167)
    fmt.Println("v169:", v169)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v166
    _ = v167
    _ = v168
    _ = v169
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
