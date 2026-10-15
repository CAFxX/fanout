package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 string; F1 string; F2 bool }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int64; F1 int16; F2 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 int8; F2 string }
type GenS9 struct { F0 uint; F1 uint8; F2 string }
type GenS10 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int16; F1 uint64; F2 string }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint32; F1 int16; F2 int8 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 uint; F1 uint; F2 string }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int8; F1 int64; F2 int32 }
type GenS22 struct { F0 uint }
func genFn0() int {
    switch 32768 {
    case -495:
        var v7 int16 = 806
        _ = v7
    case 2147483647:
        if (true && true) {
            var v8 uint32 = 65535
            fmt.Println(v8)
            _ = v8
        }
    case 427:
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int16 = 42
            _ = fi9
            _ = v10
        }
    }
    switch ("hello" + ("" + "")) {
    case "\n\t":
        var v11 uint32 = 0
        _ = v11
    case "a":
        var v12 uint32 = 65535
        _ = v12
    case "foo bar":
        v13 := []int{-111, 32768}
        _ = v13
    }
    v14 := []int{-536}
    if (0 <= 161) {
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint16 = 0
            v16 += 42
            switch (fi15 >> ((561) & 63)) {
            case -381:
                switch fi15 {
                case -859:
                    var v17 float32 = 3.14159
                    _ = v17
                default:
                    var v18 uint = 2
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 15; fi19++ {
                    if len(v14) > 0 {
                        v14[0] = fi19
                    }
                    _ = fi19
                }
            }
            _ = fi15
            _ = v16
        }
    } else {
        var v20 uint8 = 42
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 int32 = -1
        _ = fi21
        _ = v22
    }
    _ = v14
    return 4294967295
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]string)
    var v3 int64 = 9223372036854775807
    var v4 int8 = -70
    var v5 float32 = (v1 - (0.0 * 0.1))
    var v6 uint64 = 649
    v23 := GenS1{F0: (v4 & v4)}
    delete(v2, 830)
    v2[uint32(v6)] = "foo bar"
    if false {
        if false {
            v3 -= 9223372036854775807
            var v24 float64 = ((((1e10 + 1e10) / 1e10) / 1.5) - 1e10)
            v1 = float32((((v24 * ((v24 - 0.1) / ((v24 * (v24 * v24)) * -2.25))) / ((1.5 * v24) * v24)) * (-2.25 - (0.0 * v24))))
            _ = v24
        } else {
            switch "\n\t" {
            case "a":
                var v25 uint32 = 65535
                _ = v25
            case "foo bar":
                v26 := GenS1{F0: (v4 >> ((2) & 63))}
                _ = v26
            }
        }
    } else {
        for fi27 := 0; fi27 < 6; fi27++ {
            if (!true) {
                if false {
                    var v28 int = (-(fi27))
                    for fi29 := 0; fi29 < 13; fi29++ {
                        if ((552 > 459) && (("hello" + "hello") != "foo bar")) {
                            var v30 int = int(v6)
                            v28 -= 0
                            _ = v30
                        } else {
                            if false {
                                var v31 bool = (!(false || true))
                                v28 -= genFn0()
                                _ = v31
                            } else {
                                var v32 int64 = (-(v3))
                                if (("" + "") == (("hello" + (("foo bar" + "") + "foo bar")) + "hello")) {
                                    var v33 int16 = 1000
                                    _ = len(v2)
                                    _ = v33
                                } else {
                                    if ((false && ((!false) || false)) || (953 <= 0)) {
                                        fmt.Println(v4, v5)
                                        var v34 int16 = -168
                                        if true {
                                            v5 += float32((1.5 * (3.14159 / 0.1)))
                                            var v35 int64 = 9223372036854775807
                                            switch fi29 {
                                            case -2147483648:
                                                v36 := []int{(fi27 ^ fi27), (fi27 + (fi29 & (-(v28)))), (fi27 & 42)}
                                                _ = v36
                                            }
                                            _ = v35
                                        }
                                        switch ("\n\t" + ("a" + "hello")) {
                                        case "hello":
                                            if (true || (!true)) {
                                                v2[504] = ""
                                            }
                                        case "a":
                                            var v37 float32 = 3.14159
                                            _ = v37
                                        case "foo bar":
                                            v32 = -9223372036854775808
                                        }
                                        _ = v34
                                    }
                                    v3 += 42
                                    fmt.Println(fi29)
                                    if (false || ((-(v28)) >= (fi29 | int(v4)))) {
                                        if (!false) {
                                            switch "hello" {
                                            case "\n\t":
                                                var v38 float32 = ((float32(((((3.14159 * (-2.25 / -2.25)) / 1.5) * 0.0) / 1.0)) - -2.25) / v1)
                                                _ = v38
                                            case "foo bar":
                                                switch 4294967295 {
                                                case 2147483647:
                                                    var v39 uint = 1
                                                    _ = v39
                                                case 986:
                                                    v40 := GenS2{F0: 171, F1: int(v32)}
                                                    _ = v40
                                                case -9223372036854775808:
                                                    var v41 uint = 255
                                                    _ = v41
                                                }
                                            case "":
                                                fmt.Println(v3, v1)
                                            }
                                            v42 := []int{(fi27 - fi29), (fi27 + genFn0())}
                                            if (-1 == 2147483647) {
                                                var v43 uint8 = 44
                                                v3 = v32
                                                var v44 int64 = v3
                                                fmt.Println(fi27, v44)
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 float32 = (1.5 * v1)
                                            if true {
                                                switch "foo bar" {
                                                case "a":
                                                    v46 := GenS3{F0: 4294967295, F1: (v32 | (v32 | (^(v3))))}
                                                    _ = v46
                                                case "foo bar":
                                                    v47 := v2[2147483647]
                                                    _ = v47
                                                case "":
                                                    _ = len(v2)
                                                default:
                                                    var v48 float32 = (((v5 / (((0.0 - float32(((1e10 / 1e10) - (((float64(v5) + 0.1) * (((0.0 * (((3.14159 / (-2.25 + 1.5)) / float64(1e10)) / (1e10 / 0.1))) * 1e10) * (3.14159 - (0.1 / ((1.5 / 1.0) + float64(v5)))))) * ((-2.25 - 1.5) / (3.14159 * ((3.14159 / 1.0) - 0.0))))))) + 1.5) / v5)) / ((0.0 / v45) * (v5 / ((v1 / 1.5) - (v1 - (v5 * v5)))))) - v1)
                                                    _ = v48
                                                }
                                                var v49 int32 = 341
                                                _ = v49
                                            }
                                            _ = v42
                                            _ = v45
                                        }
                                        v28 -= fi27
                                    } else {
                                        if (-2147483648 == int32(fi29)) {
                                            v6 += v6
                                            var v50 int64 = v3
                                            var v51 uint = 2
                                            _ = v50
                                            _ = v51
                                        }
                                        if (-1 != -32768) {
                                            if (true && (0 == 65535)) {
                                                v1 -= v1
                                                if ((false || (577 <= -407)) || ((v4 % (((v4 & 127)) | 1)) >= 42)) {
                                                    fmt.Println(fi27, fi27)
                                                    fmt.Println(v4)
                                                    switch fi27 {
                                                    case 32768:
                                                        v5 -= (v5 / ((v5 / float32(0.1)) * float32(0.1)))
                                                    }
                                                    if true {
                                                        switch (("a" + ((((("" + (("foo bar" + "") + ("" + ("foo bar" + "a")))) + ("foo bar" + "")) + ("a" + "a")) + "\n\t") + "hello")) + "foo bar") {
                                                        case "":
                                                            v4 = (-(v4))
                                                        }
                                                        v52 := GenS4{F0: ("" + ""), F1: (("\n\t" + "foo bar") + "foo bar"), F2: (false || true)}
                                                        _ = v52
                                                    }
                                                    v32 -= 42
                                                }
                                                _ = len(v2)
                                            } else {
                                                v53 := []int{v28, 4294967295}
                                                if false {
                                                    switch (^(v28)) {
                                                    case 1:
                                                        v1 -= (v1 + (v5 + (v5 + 3.14159)))
                                                    default:
                                                        if false {
                                                            if (true && false) {
                                                                v53 = append(v53, (fi29 + (-(fi27))))
                                                                switch 2147483647 {
                                                                case -144:
                                                                    v28 = fi29
                                                                case 658:
                                                                    v54 := GenS5{F0: -1, F1: 512}
                                                                    _ = v54
                                                                }
                                                            } else {
                                                                var v55 int64 = v3
                                                                v3 -= v32
                                                                _ = v55
                                                            }
                                                            var v56 float32 = v5
                                                            var v57 int64 = 0
                                                            _ = v56
                                                            _ = v57
                                                        } else {
                                                            switch ("" + ("hello" + "hello")) {
                                                            case "\n\t":
                                                                var v58 float32 = (0.1 + (-2.25 / v1))
                                                                v2[uint32(v6)] = ("hello" + ("foo bar" + ("foo bar" + ("hello" + "\n\t"))))
                                                                fmt.Println(v5, v4)
                                                                _ = v58
                                                            case "hello":
                                                                var v59 uint8 = 123
                                                                _ = v59
                                                            default:
                                                                var v60 uint16 = 0
                                                                _ = v60
                                                            }
                                                            fmt.Println(fi29, v4)
                                                            v1 = v1
                                                        }
                                                    }
                                                    var v61 string = "foo bar"
                                                    v53 = append(v53, v28)
                                                    var v62 bool = (-174 >= 32767)
                                                    fmt.Println(v6)
                                                    _ = v61
                                                    _ = v62
                                                }
                                                switch "" {
                                                case "a":
                                                    v3 = (v32 / ((v3) | 1))
                                                case "foo bar":
                                                    v63 := GenS6{F0: (v3 ^ -1), F1: -243, F2: 2}
                                                    _ = v63
                                                case "\n\t":
                                                    if false {
                                                        switch v28 {
                                                        case 264:
                                                            if true {
                                                                var v64 string = "hello"
                                                                _ = v64
                                                            }
                                                        }
                                                        v53 = append(v53, (v28 + (fi29 ^ -2147483648)))
                                                        var v65 bool = ((((1.5 + float64(v5)) < 1e10) || (42 > 1)) || (-464 <= 9223372036854775807))
                                                        fmt.Println(v1, fi29, v3)
                                                        fmt.Println(v6)
                                                        _ = v65
                                                    }
                                                }
                                                _ = v53
                                            }
                                            var v66 uint32 = 487
                                            _ = v66
                                        } else {
                                            delete(v2, 42)
                                            var v67 uint8 = 0
                                            _ = v67
                                        }
                                        v68 := GenS2{F0: v28, F1: (fi27 | fi29)}
                                        _ = v68
                                    }
                                    v69 := []int{626, genFn0(), v28}
                                    var v70 int32 = 249
                                    _ = v69
                                    _ = v70
                                }
                                var v71 uint8 = 42
                                _ = v32
                                _ = v71
                            }
                            if ((!false) || false) {
                                switch 1 {
                                case 32768:
                                    fmt.Println(v28, v3)
                                default:
                                    if true {
                                        var v72 int = (fi29 / ((fi29) | 1))
                                        var v73 uint8 = 128
                                        v1 = ((3.14159 - -2.25) / (v1 + 3.14159))
                                        v73 += v73
                                        if true {
                                            if true {
                                                var v74 uint32 = 42
                                                if true {
                                                    v75 := GenS5{F0: 1000, F1: (^(fi27))}
                                                    _ = v75
                                                } else {
                                                    var v76 int16 = -32768
                                                    v77 := GenS7{F0: false}
                                                    _ = v76
                                                    _ = v77
                                                }
                                                _ = v74
                                            } else {
                                                var v78 int = (v72 % (((fi29 ^ genFn0())) | 1))
                                                var v79 string = ("" + "\n\t")
                                                _ = v78
                                                _ = v79
                                            }
                                            var v80 uint = 0
                                            var v81 int8 = v4
                                            _ = v80
                                            _ = v81
                                        } else {
                                            var v82 uint32 = 1
                                            _ = v82
                                        }
                                        switch (-(v28)) {
                                        case -9223372036854775808:
                                            if ((v4 & (-(v4))) == v4) {
                                                v83 := []int{2147483647, (v72 ^ (fi27 & fi27)), (v72 / ((v72) | 1)), fi27}
                                                _ = v83
                                            } else {
                                                var v84 uint = 0
                                                v85 := []int{-169, fi29, int(v4), int(v84)}
                                                var v86 uint8 = 127
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                            fmt.Println(v6, v5)
                                        }
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                v87 := v2[691]
                                if ((v28 << ((778) & 63)) < fi29) {
                                    switch 128 {
                                    case 4294967295:
                                        switch 1 {
                                        case 1:
                                            v6 += 474
                                        case 32768:
                                            var v88 int8 = v4
                                            v3 += v3
                                            _ = v88
                                        case 4294967295:
                                            fmt.Println(v4)
                                        }
                                    case 88:
                                        v89 := GenS8{F0: 48, F1: (v4 / ((v4) | 1)), F2: v87}
                                        fmt.Println(v5, v28)
                                        _ = v89
                                    case 19:
                                        delete(v2, 458)
                                    default:
                                        switch (fi29 - fi27) {
                                        case -832:
                                            if (76 > 107) {
                                                var v90 int32 = 2147483647
                                                v91 := GenS9{F0: 752, F1: 2, F2: "hello"}
                                                if (true || (("hello" + (v87 + v87)) > (v87 + v87))) {
                                                    fmt.Println(fi29, v6)
                                                }
                                                v92 := GenS9{F0: 18446744073709551615, F1: 42, F2: "a"}
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                            }
                                        case -2147483648:
                                            var v93 uint32 = 65535
                                            _ = v93
                                        case 9223372036854775807:
                                            switch (v87 + v87) {
                                            case "foo bar":
                                                v3 = 9223372036854775807
                                            case "\n\t":
                                                if (127 >= (v4 << ((326) & 63))) {
                                                    v94 := []int{genFn0(), fi27, (v28 ^ fi29), (v28 + fi27)}
                                                    _ = v94
                                                }
                                            }
                                        default:
                                            switch (((((((v87 + "hello") + "hello") + "a") + "foo bar") + (v87 + v87)) + v87) + v87) {
                                            case "a":
                                                if (true || false) {
                                                    var v95 uint32 = 42
                                                    _ = v95
                                                } else {
                                                    fmt.Println(v6)
                                                    v4 -= v4
                                                    switch v87 {
                                                    case "":
                                                        var v96 uint32 = 0
                                                        _ = v96
                                                    }
                                                    v97 := GenS10{F0: 42, F1: 1, F2: 738}
                                                    _ = v97
                                                }
                                                var v98 int64 = (-(v3))
                                                _ = v98
                                            case "\n\t":
                                                if true {
                                                    if false {
                                                        v87 = ("a" + v87)
                                                        var v99 int32 = -1
                                                        _ = v99
                                                    } else {
                                                        v100 := GenS1{F0: (^(v4))}
                                                        v3 = (^(v3))
                                                        v101 := GenS1{F0: v4}
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                    v102 := GenS7{F0: (39 != (fi27 - (fi29 >> ((0) & 63))))}
                                                    _ = v102
                                                }
                                            }
                                        }
                                    }
                                    switch v87 {
                                    case "":
                                        if true {
                                            v87 = "hello"
                                        } else {
                                            v28 -= int(v4)
                                            switch "foo bar" {
                                            case "foo bar":
                                                if true {
                                                    switch ("a" + ("hello" + ((v87 + "foo bar") + (v87 + "\n\t")))) {
                                                    case "foo bar":
                                                        if (v3 <= (v3 << ((2) & 63))) {
                                                            fmt.Println(v28, fi29)
                                                            v103 := v2[65535]
                                                            var v104 uint32 = 2
                                                            v5 = float32(((0.0 - 0.0) - 3.14159))
                                                            _ = v103
                                                            _ = v104
                                                        } else {
                                                            var v105 int8 = 66
                                                            _ = v105
                                                        }
                                                    default:
                                                        v106 := []int{(^(fi27)), v28, v28}
                                                        v4 -= (v4 >> ((65535) & 63))
                                                        _ = v106
                                                    }
                                                    switch 4294967295 {
                                                    case -163:
                                                        v107 := []int{fi27}
                                                        _ = v107
                                                    case 1:
                                                        if true {
                                                            var v108 int8 = 114
                                                            fmt.Println(fi29, v87, v4)
                                                            _ = v108
                                                        }
                                                    case -790:
                                                        fmt.Println(v6, v1)
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                }
                                            case "a":
                                                v6 += (v6 * v6)
                                            case "":
                                                v109 := GenS9{F0: 209, F1: 74, F2: v87}
                                                v87 = (v87 + v87)
                                                if false {
                                                    if (255 < 191) {
                                                        if true {
                                                            fmt.Println(v1)
                                                            v87 = (v87 + "\n\t")
                                                        } else {
                                                            v110 := []int{(v28 | (v28 | (-(fi29)))), (v28 << ((255) & 63)), v28, (fi27 & 972)}
                                                            v87 = (v87 + v87)
                                                            v111 := GenS11{F0: 2}
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v28)
                                                    v112 := GenS12{F0: 1, F1: (v6 + 888), F2: ("foo bar" + ((v87 + v87) + ((v87 + v87) + (v87 + "foo bar"))))}
                                                    if (!(!(1000 != 0))) {
                                                        fmt.Println(v6)
                                                        fmt.Println(v4, v87, v3)
                                                        v113 := GenS13{F0: (((v1 - (v5 + 0.1)) + 0.0) > 0.0), F1: 42}
                                                        _ = v113
                                                    } else {
                                                        if true {
                                                            if (v6 >= v6) {
                                                                var v114 uint16 = 748
                                                                _ = v114
                                                            }
                                                            switch v87 {
                                                            case "\n\t":
                                                                var v115 int16 = -32768
                                                                _ = v115
                                                            }
                                                            var v116 string = ("foo bar" + ("" + v87))
                                                            _ = v116
                                                        }
                                                        _ = len(v2)
                                                        var v117 uint64 = (v6 % ((82) | 1))
                                                        _ = v117
                                                    }
                                                    _ = v112
                                                }
                                                var v118 int64 = 1
                                                _ = v109
                                                _ = v118
                                            }
                                            if true {
                                                v119 := GenS8{F0: 606, F1: v4, F2: "foo bar"}
                                                var v120 float64 = 0.1
                                                fmt.Println(v120)
                                                _ = v119
                                                _ = v120
                                            } else {
                                                if true {
                                                    v121 := []int{-9223372036854775808}
                                                    v87 = v87
                                                    if ((v3 * v3) != (v3 << ((42) & 63))) {
                                                        v122 := GenS14{F0: 255}
                                                        fmt.Println(v28)
                                                        var v123 bool = (false && true)
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                    v124 := GenS15{F0: 91}
                                                    v121 = append(v121, fi29)
                                                    _ = v121
                                                    _ = v124
                                                } else {
                                                    var v125 float32 = v5
                                                    v126 := []int{(v28 - fi29), fi29}
                                                    _ = v125
                                                    _ = v126
                                                }
                                                v127 := []int{fi29, int(v6)}
                                                _ = v127
                                            }
                                        }
                                    }
                                    v87 = ("\n\t" + v87)
                                }
                                v128 := GenS16{F0: 65535, F1: -509, F2: (^(v4))}
                                switch (v87 + "hello") {
                                case "a":
                                    if (true || false) {
                                        if false {
                                            var v129 uint16 = 615
                                            _ = v129
                                        } else {
                                            if true {
                                                var v130 uint16 = 572
                                                var v131 int64 = v3
                                                var v132 int8 = v4
                                                _ = v130
                                                _ = v131
                                                _ = v132
                                            }
                                            v87 = (((((v87 + "\n\t") + v87) + "") + ((v87 + v87) + v87)) + v87)
                                            var v133 bool = (0 == -1)
                                            _ = v133
                                        }
                                        switch v87 {
                                        case "hello":
                                            v134 := []int{}
                                            _ = v134
                                        case "foo bar":
                                            v4 += v4
                                        }
                                    }
                                case "":
                                    v135 := GenS12{F0: 168, F1: 60, F2: "a"}
                                    _ = v135
                                case "\n\t":
                                    v136 := []int{(^(fi29)), (fi29 & (-(fi27))), fi29}
                                    _ = v136
                                }
                                _ = v128
                                _ = v87
                            }
                        }
                        _ = fi29
                    }
                    var v137 int = (fi27 | 32767)
                    _ = v137
                    _ = v28
                }
                fmt.Println(v1, v5)
            }
            if false {
                v4 += (v4 * 42)
            } else {
                if (255 < 255) {
                    var v138 int64 = 804
                    _ = v138
                } else {
                    if false {
                        for fi139 := 0; fi139 < 3; fi139++ {
                            if ((!(true && ((uint32(v6) > 1) || (-2147483648 == 2147483647)))) && true) {
                                switch fi139 {
                                case 2147483647:
                                    var v140 float32 = ((v1 + v5) - float32(float64(v5)))
                                    _ = v140
                                case 32767:
                                    var v141 float32 = v5
                                    _ = v141
                                case 472:
                                    switch 0 {
                                    case -70:
                                        fmt.Println(v6)
                                    case 1:
                                        v4 = 100
                                    default:
                                        v4 = (v4 | v4)
                                    }
                                }
                                v142 := GenS17{F0: (^(v6))}
                                _ = v142
                            }
                            if (v5 > (-2.25 + (0.1 - ((v5 - 0.0) - (0.0 - (v1 + v5)))))) {
                                v143 := GenS18{F0: 18446744073709551615, F1: 1, F2: ("" + ("hello" + "a"))}
                                v6 -= 74
                                _ = v143
                            } else {
                                switch "hello" {
                                case "a":
                                    if (612 == 2147483647) {
                                        v144 := GenS11{F0: 2}
                                        switch (fi27 ^ fi139) {
                                        case 468:
                                            if (genFn0() != 490) {
                                                if true {
                                                    switch int(v4) {
                                                    case 171:
                                                        if ((((0.0 * (1e10 / float64(((v5 * v1) + (1e10 - ((v1 - v5) / v5)))))) - 0.0) / (3.14159 / 3.14159)) != 1.5) {
                                                            var v145 bool = true
                                                            v146 := []int{fi139, (fi139 ^ fi27), (fi139 >> ((uint(v6)) & 63)), 28}
                                                            switch fi27 {
                                                            case 42:
                                                                if (42 > 0) {
                                                                    v145 = false
                                                                    delete(v2, 590)
                                                                    v146 = append(v146, (fi139 % (((^(fi139))) | 1)))
                                                                } else {
                                                                    v5 += (0.1 / v1)
                                                                    v6 += v6
                                                                    v147 := GenS18{F0: 65535, F1: 4294967295, F2: (("foo bar" + "foo bar") + "")}
                                                                    switch (fi27 + 319) {
                                                                    case -900:
                                                                        v148 := GenS15{F0: 337}
                                                                        _ = v148
                                                                    case -455:
                                                                        v5 = 3.14159
                                                                    case -1:
                                                                        var v149 int64 = v3
                                                                        _ = v149
                                                                    }
                                                                    _ = v147
                                                                }
                                                            case 2147483647:
                                                                var v150 uint = 805
                                                                _ = v150
                                                            case 32768:
                                                                if (389 == 42) {
                                                                    if v145 {
                                                                        switch (fi27 >> ((986) & 63)) {
                                                                        case -190:
                                                                            var v151 uint = 65535
                                                                            var v152 uint32 = 65535
                                                                            v151 = v151
                                                                            _ = v151
                                                                            _ = v152
                                                                        case -781:
                                                                            v145 = v145
                                                                            if (1 <= (v6 << ((42) & 63))) {
                                                                                var v153 bool = v145
                                                                                var v154 uint64 = (v6 & v6)
                                                                                if (-2.25 > (0.0 * (0.1 / 3.14159))) {
                                                                                    if len(v146) >= 2 {
                                                                                        v146 = v146[1:len(v146)]
                                                                                    }
                                                                                    var v155 int64 = (v3 + (v3 - int64(v4)))
                                                                                    v156 := GenS14{F0: 1}
                                                                                    if len(v146) > 0 {
                                                                                        v146[0] = fi139
                                                                                    }
                                                                                    _ = v155
                                                                                    _ = v156
                                                                                }
                                                                                _ = v153
                                                                                _ = v154
                                                                            }
                                                                        case 0:
                                                                            switch 307 {
                                                                            case 32767:
                                                                                var v157 int64 = (v3 ^ (v3 << ((2) & 63)))
                                                                                if len(v146) >= 2 {
                                                                                    v146 = v146[1:len(v146)]
                                                                                }
                                                                                _ = v157
                                                                            default:
                                                                                v5 = v1
                                                                            }
                                                                        default:
                                                                            var v158 uint16 = 0
                                                                            _ = len(v2)
                                                                            v146 = append(v146, 421)
                                                                            _ = v158
                                                                        }
                                                                        v145 = (1 >= 4294967295)
                                                                        v159 := GenS19{F0: -861}
                                                                        if (65535 == 18446744073709551615) {
                                                                            v1 -= (((float32((3.14159 + ((0.1 + 3.14159) + 1.5))) / (1e10 * ((v1 * 3.14159) + ((v5 - (float32(3.14159) * (1e10 - (0.0 + (((3.14159 / 3.14159) * float32(float64(0.1))) / 1e10))))) - (v1 - (v1 - ((v1 / 3.14159) + (v5 / (1.5 / 1e10))))))))) * (3.14159 / 1.0)) + v1)
                                                                            v3 += v3
                                                                        } else {
                                                                            var v160 uint32 = 4294967295
                                                                            var v161 int8 = v4
                                                                            _ = v160
                                                                            _ = v161
                                                                        }
                                                                        _ = v159
                                                                    } else {
                                                                        if (67 > 22) {
                                                                            switch fi27 {
                                                                            case -510:
                                                                                fmt.Println(fi27, v1)
                                                                            case -1:
                                                                                v162 := GenS20{F0: 682}
                                                                                _ = v162
                                                                            default:
                                                                                v4 += v4
                                                                            }
                                                                            v3 = 0
                                                                        }
                                                                        var v163 float64 = (((0.0 * -2.25) + 1e10) * 1e10)
                                                                        _ = v163
                                                                    }
                                                                    var v164 uint = 2
                                                                    _ = len(v2)
                                                                    _ = v164
                                                                } else {
                                                                    fmt.Println(fi27)
                                                                }
                                                            default:
                                                                var v165 string = "a"
                                                                _ = v165
                                                            }
                                                            var v166 int32 = -2147483648
                                                            _ = v145
                                                            _ = v146
                                                            _ = v166
                                                        }
                                                    case 9223372036854775807:
                                                        var v167 int16 = -1
                                                        _ = v167
                                                    }
                                                    v168 := []int{}
                                                    _ = v168
                                                }
                                                v4 -= v4
                                                v169 := v2[343]
                                                _ = v169
                                            }
                                        case 984:
                                            v170 := []int{127, fi27, 128}
                                            _ = v170
                                        case 42:
                                            var v171 int8 = v4
                                            _ = v171
                                        }
                                        fmt.Println(fi27, v1)
                                        _ = v144
                                    }
                                case "":
                                    if ((v3 % ((int64(fi139)) | 1)) != (v3 / ((8) | 1))) {
                                        v3 += 9223372036854775807
                                        _ = len(v2)
                                    } else {
                                        var v172 uint16 = 255
                                        _ = v172
                                    }
                                }
                                switch "hello" {
                                case "hello":
                                    var v173 uint8 = 127
                                    _ = v173
                                case "a":
                                    v174 := GenS21{F0: v4, F1: (v3 * (v3 + (^(v3)))), F2: -252}
                                    _ = v174
                                case "foo bar":
                                    var v175 uint64 = v6
                                    _ = v175
                                }
                            }
                            var v176 float64 = 0.0
                            _ = fi139
                            _ = v176
                        }
                        var v177 uint32 = 2
                        fmt.Println(v1, v5)
                        _ = v177
                    }
                }
                var v178 uint64 = 0
                _ = v178
            }
            for fi179 := 0; fi179 < 10; fi179++ {
                v180 := v2[142]
                fmt.Println(v180, fi179)
                v3 = v3
                _ = fi179
                _ = v180
            }
            _ = fi27
        }
    }
    v181 := []int{127, 128, 1, -2147483648}
    for fi182 := 0; fi182 < 1; fi182++ {
        v5 += ((v5 - v1) - 0.1)
        var v183 int32 = 1
        _ = fi182
        _ = v183
    }
    if len(v181) >= 2 {
        v181 = v181[1:len(v181)]
    }
    switch 127 {
    case 32767:
        v6 = 0
    case -9223372036854775808:
        var v184 uint8 = 128
        _ = v184
    default:
        v185 := GenS22{F0: 4294967295}
        _ = v185
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v181:", len(v181), v181)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v181
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
