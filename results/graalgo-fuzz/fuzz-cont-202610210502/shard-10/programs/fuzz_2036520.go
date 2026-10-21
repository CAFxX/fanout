package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint; F2 int }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 int64; F2 int64 }
type GenS9 struct { F0 uint; F1 int64; F2 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool; F1 uint8; F2 int16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int; F1 int64; F2 int32 }
type GenS17 struct { F0 bool; F1 uint32 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 uint; F1 uint }
type GenS20 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint64) int64 {
    fmt.Println(p1, p0)
    var v4 int = -34
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 188
}
func genFn1(p0 uint16) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    fmt.Println(p0)
    switch 377 {
    case 646:
        fmt.Println(p0)
    case -821:
        p0 += p0
    }
    var v5 uint64 = 847
    _ = p0
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]int)
    fmt.Println(v1)
    fmt.Println(v2)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            v1 -= (^(v1))
            _ = fi7
        }
        if false {
            v8 := GenS1{F0: -1, F1: fi6}
            if (fi6 == (fi6 >> ((42) & 63))) {
                switch (fi6 ^ (fi6 | fi6)) {
                case 128:
                    var v9 uint = 871
                    _ = v9
                case 645:
                    var v10 uint32 = 65535
                    _ = v10
                }
                var v11 int8 = 100
                v12 := v3["foo bar"]
                _ = v11
                _ = v12
            }
            var v13 int8 = 42
            _ = v13
            _ = v8
        } else {
            v3["a"] = (fi6 * (fi6 / ((fi6) | 1)))
            switch fi6 {
            case 309:
                for fi14 := 0; fi14 < 2; fi14++ {
                    var v15 int = 724
                    var v16 string = "\n\t"
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case 32767:
                fmt.Println(v2)
            default:
                v17 := []int{fi6, fi6, (fi6 % ((fi6) | 1)), (-(fi6))}
                _ = v17
            }
        }
        _ = fi6
    }
    if (true || (true || (false || false))) {
        for fi18 := 0; fi18 < 7; fi18++ {
            if ((genFn1(42) < 18446744073709551615) && true) {
                fmt.Println(v1, v2, v2)
            }
            _ = fi18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 float64 = (-2.25 - float64(1e10))
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS1{F0: 100, F1: -921}
                v24 := []int{fi19}
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi19
        }
        var v25 int32 = -460
        _ = v25
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 int32 = 432
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            if false {
                if (27 == 1) {
                    if (((1e10 / (1.5 / (3.14159 - 0.1))) >= 1.5) || false) {
                        v1 = (v1 << ((255) & 63))
                        if true {
                            if false {
                                v1 -= 2
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    v2 += v2
                                    _ = fi30
                                }
                                v1 += 42
                            }
                            var v31 int16 = 0
                            v32 := []int{(^(fi29)), (fi29 | fi29), 42, (fi29 % ((fi29) | 1))}
                            _ = v31
                            _ = v32
                        }
                        delete(v3, (("a" + "a") + "\n\t"))
                    } else {
                        var v33 uint32 = 65535
                        _ = v33
                    }
                    if true {
                        for fi34 := 0; fi34 < 7; fi34++ {
                            if false {
                                switch (fi34 / ((fi29) | 1)) {
                                case 4294967295:
                                    var v35 int = fi29
                                    _ = v35
                                case 277:
                                    var v36 float64 = 0.1
                                    _ = v36
                                case 9223372036854775807:
                                    switch fi34 {
                                    case 128:
                                        switch ((("\n\t" + "") + ("foo bar" + ("" + "\n\t"))) + (("foo bar" + (("\n\t" + (("hello" + ("" + "a")) + ("a" + ("foo bar" + "\n\t")))) + "a")) + "\n\t")) {
                                        case "hello":
                                            v37 := GenS2{F0: 185}
                                            _ = v37
                                        case "":
                                            var v38 string = "a"
                                            _ = v38
                                        }
                                    }
                                default:
                                    var v39 float32 = 1e10
                                    _ = v39
                                }
                                var v40 int32 = -777
                                if (18446744073709551615 != v1) {
                                    switch "" {
                                    case "a":
                                        v40 += (v40 * (v40 << ((65535) & 63)))
                                    case "\n\t":
                                        switch "" {
                                        case "foo bar":
                                            if true {
                                                v2 = -1
                                            } else {
                                                var v41 float32 = 1e10
                                                fmt.Println(fi34)
                                                _ = v41
                                            }
                                        case "a":
                                            fmt.Println(v1)
                                        }
                                    case "foo bar":
                                        if (true || (236 >= 0)) {
                                            v42 := []int{fi29}
                                            _ = v42
                                        } else {
                                            v43 := []int{(^(fi29)), fi29, (fi34 + fi34)}
                                            _ = v43
                                        }
                                    }
                                    switch "" {
                                    case "":
                                        v1 -= genFn1(255)
                                    case "foo bar":
                                        switch fi34 {
                                        case -120:
                                            var v44 uint16 = 0
                                            _ = v44
                                        default:
                                            v45 := GenS3{F0: "foo bar"}
                                            _ = v45
                                        }
                                    case "hello":
                                        switch ("foo bar" + ("hello" + "a")) {
                                        case "hello":
                                            switch "\n\t" {
                                            case "":
                                                v46 := GenS2{F0: 2}
                                                _ = v46
                                            case "hello":
                                                switch fi29 {
                                                case 960:
                                                    if (int(v40) >= fi34) {
                                                        var v47 int16 = -1
                                                        _ = v47
                                                    } else {
                                                        var v48 uint16 = 255
                                                        _ = v48
                                                    }
                                                case 4294967295:
                                                    v1 += 595
                                                case 344:
                                                    if (true && true) {
                                                        switch ("a" + "") {
                                                        case "a":
                                                            if (1.5 < ((3.14159 - (1e10 - 0.0)) - ((0.1 / float32(((0.0 - float64(1.5)) * (0.1 + 0.0)))) / (0.0 - (3.14159 - 1e10))))) {
                                                                var v49 int16 = -1
                                                                _ = v49
                                                            } else {
                                                                v1 = (v1 << ((566) & 63))
                                                                v50 := GenS4{F0: false, F1: 2, F2: fi34}
                                                                _ = v50
                                                            }
                                                        default:
                                                            v51 := GenS3{F0: ("\n\t" + ("hello" + ("a" + "")))}
                                                            _ = v51
                                                        }
                                                        delete(v3, "")
                                                        var v52 int16 = 1
                                                        _ = v52
                                                    } else {
                                                        var v53 uint32 = 65535
                                                        _ = v53
                                                    }
                                                default:
                                                    var v54 int = fi29
                                                    _ = v54
                                                }
                                            default:
                                                switch -9223372036854775808 {
                                                case 32767:
                                                    switch (-(fi34)) {
                                                    case 9:
                                                        fmt.Println(v40, fi29, v2)
                                                    case 32768:
                                                        fmt.Println(fi29)
                                                    case -1:
                                                        var v55 bool = true
                                                        _ = v55
                                                    }
                                                case 2147483647:
                                                    v56 := GenS5{F0: 1000, F1: v1}
                                                    _ = v56
                                                default:
                                                    v57 := GenS6{F0: 2147483648, F1: 65535}
                                                    _ = v57
                                                }
                                            }
                                        case "foo bar":
                                            v1 = v1
                                        case "":
                                            v1 = v1
                                        }
                                    default:
                                        delete(v3, ("\n\t" + "foo bar"))
                                    }
                                } else {
                                    if true {
                                        switch (-(fi29)) {
                                        case 32767:
                                            if (true && (128 < 127)) {
                                                var v58 int = (fi34 / ((fi34) | 1))
                                                fmt.Println(v58, v2, fi29)
                                                _ = v58
                                            }
                                            if true {
                                                var v59 int32 = v40
                                                var v60 int16 = int16(fi34)
                                                _ = v59
                                                _ = v60
                                            }
                                        case 9223372036854775807:
                                            v61 := []int{fi29, fi34, (-(fi34)), (fi34 + fi34)}
                                            var v62 float32 = (1.5 * 1e10)
                                            _ = v61
                                            _ = v62
                                        case 0:
                                            v63 := GenS5{F0: 1, F1: (^(v1))}
                                            var v64 string = "a"
                                            if (false || ((v1 + v1) <= (v1 & v1))) {
                                                v65 := v3[v64]
                                                if true {
                                                    var v66 string = "foo bar"
                                                    switch v65 {
                                                    case 128:
                                                        v40 -= 504
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                    _ = v66
                                                }
                                                var v67 uint16 = 46
                                                v68 := GenS7{F0: 996}
                                                _ = v65
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = v63
                                            _ = v64
                                        default:
                                            fmt.Println(v2, v2, v1)
                                            v2 = (v2 / ((42) | 1))
                                            fmt.Println(fi34)
                                        }
                                        if ((461 == 42) || (v2 >= 9223372036854775807)) {
                                            fmt.Println(v40)
                                            v2 = (v2 & (v2 ^ 389))
                                        }
                                        var v69 uint8 = 42
                                        _ = v69
                                    } else {
                                        if true {
                                            if (!(!true)) {
                                                fmt.Println(v1)
                                                if (false && (("" + "foo bar") > "hello")) {
                                                    v70 := GenS8{F0: true, F1: v2, F2: 0}
                                                    fmt.Println(v2)
                                                    var v71 uint8 = 128
                                                    fmt.Println(v40, v40, fi29)
                                                    _ = v70
                                                    _ = v71
                                                }
                                                if (!true) {
                                                    var v72 uint8 = 127
                                                    v2 = (v2 ^ (v2 | v2))
                                                    var v73 string = ("\n\t" + ("" + "\n\t"))
                                                    v1 = v1
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    v1 = v1
                                                    var v74 string = ""
                                                    _ = v74
                                                }
                                                if true {
                                                    switch (^(fi29)) {
                                                    case -449:
                                                        switch int(v2) {
                                                        case 32768:
                                                            var v75 int64 = -9223372036854775808
                                                            _ = v75
                                                        default:
                                                            fmt.Println(fi34, v1)
                                                        }
                                                    default:
                                                        var v76 bool = (!true)
                                                        _ = v76
                                                    }
                                                    fmt.Println(v1)
                                                    v40 += (v40 << ((255) & 63))
                                                    var v77 uint = 4294967295
                                                    _ = v77
                                                } else {
                                                    var v78 int32 = (v40 * v40)
                                                    v79 := []int{(fi29 << ((4294967295) & 63)), -9223372036854775808, fi34, int(v78)}
                                                    v80 := v3["a"]
                                                    v78 += (v40 % (((v78 / ((v40) | 1))) | 1))
                                                    var v81 string = "\n\t"
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                }
                                            }
                                            fmt.Println(v2)
                                        } else {
                                            if true {
                                                if true {
                                                    fmt.Println(fi29, v1)
                                                    var v82 int32 = (v40 >> ((4294967295) & 63))
                                                    _ = v82
                                                } else {
                                                    switch ("" + "hello") {
                                                    case "\n\t":
                                                        var v83 int64 = 9223372036854775807
                                                        if (!true) {
                                                            var v84 uint8 = 0
                                                            var v85 uint32 = 862
                                                            fmt.Println(v40, v84, v85)
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        _ = v83
                                                    case "":
                                                        v86 := []int{(fi34 | fi34), fi29, 2147483647, fi29}
                                                        _ = v86
                                                    case "hello":
                                                        var v87 uint64 = 0
                                                        _ = v87
                                                    }
                                                    v2 = (v2 << ((504) & 63))
                                                    v1 = (v1 ^ v1)
                                                    v88 := v3[""]
                                                    _ = v88
                                                }
                                                switch "a" {
                                                case "":
                                                    v2 = 0
                                                default:
                                                    fmt.Println(v40)
                                                }
                                            } else {
                                                if true {
                                                    switch -9223372036854775808 {
                                                    case 829:
                                                        fmt.Println(fi34, v1)
                                                    case 323:
                                                        var v89 uint32 = 4294967295
                                                        _ = v89
                                                    default:
                                                        var v90 int16 = 42
                                                        _ = v90
                                                    }
                                                    switch (fi34 ^ fi34) {
                                                    case 4294967295:
                                                        v1 = v1
                                                    case 268:
                                                        switch (fi34 >> ((415) & 63)) {
                                                        case 2147483647:
                                                            v1 -= genFn1(529)
                                                        }
                                                    case 32767:
                                                        v91 := GenS7{F0: 65535}
                                                        if true {
                                                            if true {
                                                                var v92 int8 = -128
                                                                var v93 int16 = -547
                                                                fmt.Println(v93, fi29, v92)
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                            fmt.Println(v2, v40, v40)
                                                            v40 += (v40 & -1)
                                                        }
                                                        v40 += (v40 ^ (-(v40)))
                                                        _ = v91
                                                    default:
                                                        switch fi34 {
                                                        case -9223372036854775808:
                                                            var v94 float64 = 1.5
                                                            _ = v94
                                                        case 0:
                                                            var v95 uint32 = 2147483647
                                                            _ = v95
                                                        case 32768:
                                                            v40 -= v40
                                                        }
                                                    }
                                                    fmt.Println(v2, v40, fi34)
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        if true {
                                                            v2 -= (v2 >> ((42) & 63))
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v3["\n\t"] = (^(fi34))
                                                            case "a":
                                                                if false {
                                                                    fmt.Println(v40, fi34, fi34)
                                                                    v96 := GenS7{F0: 0}
                                                                    _ = v96
                                                                }
                                                            case "":
                                                                fmt.Println(v1)
                                                            }
                                                            if ((v40 | v40) >= (v40 | 29)) {
                                                                v97 := GenS9{F0: 733, F1: v2, F2: true}
                                                                var v98 bool = true
                                                                _ = len(v3)
                                                                v99 := GenS1{F0: 1, F1: fi29}
                                                                v100 := GenS6{F0: 860, F1: 0}
                                                                _ = v100
                                                                _ = v97
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                        } else {
                                                            if (0 >= 42) {
                                                                delete(v3, ("" + "hello"))
                                                                v101 := []int{(fi34 | fi29)}
                                                                v40 = -2147483648
                                                                v40 += -22
                                                                _ = v101
                                                            } else {
                                                                if true {
                                                                    if false {
                                                                        v40 = (-(v40))
                                                                    }
                                                                    if (0 == 4294967295) {
                                                                        switch "" {
                                                                        case "a":
                                                                            fmt.Println(v40, fi34)
                                                                        case "":
                                                                            v102 := []int{793, (fi29 ^ (fi29 + int(v1))), (fi29 ^ (fi29 >> ((543) & 63)))}
                                                                            _ = v102
                                                                        default:
                                                                            if true {
                                                                                v2 += (v2 / (((^(v2))) | 1))
                                                                                v103 := GenS10{F0: fi29}
                                                                                var v104 uint8 = 0
                                                                                _ = v103
                                                                                _ = v104
                                                                            } else {
                                                                                if true {
                                                                                    if (float64(((0.1 * 1e10) / -2.25)) >= (((0.1 + -2.25) + ((0.0 / (0.0 + float64(((0.0 - (1e10 / 1e10)) / 0.1)))) + ((float64(0.0) - 1e10) / 1.0))) / 1.0)) {
                                                                                        if (!(2 >= 1)) {
                                                                                            if false {
                                                                                                v2 = -9223372036854775808
                                                                                                v3["hello"] = (fi34 & fi34)
                                                                                                v105 := []int{}
                                                                                                _ = v105
                                                                                            }
                                                                                            v40 += (v40 << ((42) & 63))
                                                                                            var v106 int8 = 0
                                                                                            v107 := GenS7{F0: 2}
                                                                                            if false {
                                                                                                var v108 int64 = (v2 >> ((4294967295) & 63))
                                                                                                if true {
                                                                                                    if (false && true) {
                                                                                                        fmt.Println(v2, v40)
                                                                                                    }
                                                                                                    var v109 uint32 = 42
                                                                                                    v106 += v106
                                                                                                    _ = v109
                                                                                                }
                                                                                                v110 := []int{(fi34 | -2147483648)}
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    var v111 uint = 8
                                                                                                    _ = v111
                                                                                                default:
                                                                                                    v1 = 18446744073709551615
                                                                                                }
                                                                                                _ = v108
                                                                                                _ = v110
                                                                                            }
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        }
                                                                                        if ("a" >= "\n\t") {
                                                                                            fmt.Println(fi34, v40)
                                                                                            v2 += (v2 & v2)
                                                                                            v40 += 1
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        if false {
                                                                                            var v112 int16 = 1
                                                                                            fmt.Println(v1, v112)
                                                                                            _ = v112
                                                                                        } else {
                                                                                            switch 1 {
                                                                                            case 127:
                                                                                                v113 := GenS11{F0: (v2 / ((-982) | 1))}
                                                                                                _ = v113
                                                                                            case 256:
                                                                                                if false {
                                                                                                    var v114 uint16 = 397
                                                                                                    if false {
                                                                                                        var v115 string = "foo bar"
                                                                                                        var v116 int = fi29
                                                                                                        v117 := GenS7{F0: 233}
                                                                                                        switch -692 {
                                                                                                        case 337:
                                                                                                            v118 := []int{}
                                                                                                            _ = v118
                                                                                                        case 32767:
                                                                                                            v119 := GenS12{F0: genFn0(true, (v1 * 18446744073709551615))}
                                                                                                            _ = v119
                                                                                                        case -831:
                                                                                                            v120 := []int{}
                                                                                                            _ = v120
                                                                                                        }
                                                                                                        _ = v115
                                                                                                        _ = v116
                                                                                                        _ = v117
                                                                                                    } else {
                                                                                                        v121 := []int{fi34, (fi34 / ((int(v2)) | 1)), fi29, (fi34 | fi29)}
                                                                                                        switch -413 {
                                                                                                        case 9223372036854775807:
                                                                                                            fmt.Println(v1)
                                                                                                        default:
                                                                                                            var v122 int32 = -542
                                                                                                            _ = v122
                                                                                                        }
                                                                                                        _ = v121
                                                                                                    }
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        if (42 <= 787) {
                                                                                                            var v123 uint64 = genFn1(v114)
                                                                                                            _ = v123
                                                                                                        } else {
                                                                                                            v2 = v2
                                                                                                            switch (^(fi29)) {
                                                                                                            case 42:
                                                                                                                switch (fi34 | 127) {
                                                                                                                case 1:
                                                                                                                    var v124 float32 = (float32((0.1 + (1e10 * ((float64(float32(0.0)) + 0.1) / (0.1 * 1.5))))) - float32(1.5))
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v125 := v3["\n\t"]
                                                                                                        var v126 int = (-(fi29))
                                                                                                        _ = v125
                                                                                                        _ = v126
                                                                                                    case "\n\t":
                                                                                                        if false {
                                                                                                            var v127 bool = (v114 == (^(v114)))
                                                                                                            var v128 string = "hello"
                                                                                                            _ = v127
                                                                                                            _ = v128
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        fmt.Println(fi34, fi29, v1)
                                                                                                    }
                                                                                                    _ = v114
                                                                                                }
                                                                                            default:
                                                                                                var v129 float64 = 0.1
                                                                                                _ = v129
                                                                                            }
                                                                                            v130 := GenS2{F0: 0}
                                                                                            _ = v130
                                                                                        }
                                                                                    }
                                                                                }
                                                                                delete(v3, (("\n\t" + ("" + (("a" + "\n\t") + (("" + ("a" + "foo bar")) + "\n\t")))) + ("\n\t" + "\n\t")))
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v2 -= v2
                                                            }
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v2 -= genFn0((false && false), v1)
                                                            default:
                                                                fmt.Println(v2)
                                                            }
                                                            v131 := GenS2{F0: uint8(v1)}
                                                            _ = v131
                                                        }
                                                    case "\n\t":
                                                        v132 := GenS13{F0: false, F1: 0, F2: 32767}
                                                        _ = v132
                                                    case "":
                                                        var v133 uint = 0
                                                        _ = v133
                                                    default:
                                                        var v134 int32 = (v40 - v40)
                                                        if (true || false) {
                                                            if (fi34 <= (fi34 >> ((18446744073709551615) & 63))) {
                                                                v135 := []int{-9223372036854775808, fi29, fi34}
                                                                var v136 bool = true
                                                                _ = v135
                                                                _ = v136
                                                            }
                                                            v1 = v1
                                                        }
                                                        v137 := GenS11{F0: (^(v2))}
                                                        _ = v134
                                                        _ = v137
                                                    }
                                                }
                                            }
                                            switch "hello" {
                                            case "a":
                                                v1 -= v1
                                            default:
                                                v138 := GenS13{F0: true, F1: 11, F2: 86}
                                                _ = v138
                                            }
                                            v1 += v1
                                        }
                                        v139 := v3["\n\t"]
                                        var v140 uint8 = 17
                                        if (true && (32767 > 630)) {
                                            v140 = 217
                                            fmt.Println(v2, v139)
                                        }
                                        _ = v139
                                        _ = v140
                                    }
                                    var v141 float32 = (0.0 / (3.14159 + (3.14159 * 0.1)))
                                    _ = v141
                                }
                                _ = v40
                            } else {
                                fmt.Println(v1)
                            }
                            v1 -= (v1 * (^(v1)))
                            v142 := GenS11{F0: v2}
                            _ = fi34
                            _ = v142
                        }
                        v2 -= genFn0(false, v1)
                        var v143 uint = 65535
                        _ = v143
                    }
                    var v144 string = ("foo bar" + "foo bar")
                    _ = v144
                } else {
                    v1 += v1
                    for fi145 := 0; fi145 < 19; fi145++ {
                        v146 := GenS14{F0: 42}
                        var v147 uint8 = 1
                        _ = fi145
                        _ = v146
                        _ = v147
                    }
                    var v148 float64 = float64(3.14159)
                    _ = v148
                }
                fmt.Println(fi29, v1, fi29)
            } else {
                for fi149 := 0; fi149 < 7; fi149++ {
                    var v150 bool = (42 == 2147483648)
                    v151 := []int{(-(fi29)), (fi149 ^ 24)}
                    _ = fi149
                    _ = v150
                    _ = v151
                }
                v2 = (v2 | v2)
            }
            v152 := []int{127, fi29, fi29}
            _ = fi29
            _ = v152
        }
        var v153 int8 = -4
        _ = v153
    }
    v3["foo bar"] = 127
    var v154 float64 = 0.1
    v2 += (v2 ^ -675)
    if true {
        v155 := GenS15{F0: 373}
        _ = v155
    } else {
        for fi156 := 0; fi156 < 17; fi156++ {
            v2 += v2
            for fi157 := 0; fi157 < 4; fi157++ {
                var v158 int = 32767
                v159 := []int{-756, (v158 | fi157), 2147483647, fi156}
                if len(v159) > 0 {
                    v159[0] = fi157
                }
                if true {
                    v2 = v2
                    if false {
                        var v160 string = "\n\t"
                        v161 := GenS16{F0: (fi156 + v158), F1: genFn0(((v1 % ((v1) | 1)) == v1), 944), F2: -721}
                        var v162 int8 = 1
                        _ = v160
                        _ = v161
                        _ = v162
                    }
                    if (false || (3.14159 > 1.5)) {
                        var v163 uint = 572
                        if len(v159) > 0 {
                            v159[0] = fi156
                        }
                        _ = v163
                    } else {
                        if len(v159) >= 2 {
                            v159 = v159[1:len(v159)]
                        }
                    }
                }
                _ = fi157
                _ = v158
                _ = v159
            }
            _ = fi156
        }
    }
    switch -139 {
    case -2147483648:
        var v164 uint = 4294967295
        _ = v164
    case 2147483647:
        var v165 int64 = -73
        _ = v165
    case 42:
        v166 := v3["a"]
        _ = v166
    default:
        v2 = -9223372036854775808
    }
    fmt.Println(v2, v2)
    v1 -= genFn1(56)
    for fi167 := 0; fi167 < 16; fi167++ {
        for fi168 := 0; fi168 < 2; fi168++ {
            v169 := []int{(^(fi167)), (fi168 >> ((4294967295) & 63)), fi168}
            if (!true) {
                fmt.Println(v154)
                if false {
                    switch "foo bar" {
                    case "hello":
                        var v170 int16 = 0
                        _ = v170
                    case "a":
                        var v171 string = "a"
                        _ = v171
                    case "":
                        if len(v169) >= 2 {
                            v169 = v169[1:len(v169)]
                        }
                    }
                    if true {
                        if (("foo bar" != "\n\t") || false) {
                            var v172 uint = 4294967295
                            var v173 int = 612
                            _ = v172
                            _ = v173
                        }
                        v1 -= (v1 * (^(v1)))
                        var v174 uint32 = 4294967295
                        var v175 int8 = 1
                        if (2 == 736) {
                            v176 := GenS10{F0: fi168}
                            if (-12 == v175) {
                                if (true && (false || false)) {
                                    var v177 uint32 = (v174 / ((900) | 1))
                                    if (v1 != v1) {
                                        if len(v169) > 0 {
                                            v169[0] = fi167
                                        }
                                        if false {
                                            var v178 int16 = int16(v175)
                                            _ = v178
                                        } else {
                                            switch ("foo bar" + "hello") {
                                            case "a":
                                                v2 = (v2 & 9223372036854775807)
                                                _ = len(v3)
                                                var v179 int64 = v2
                                                _ = v179
                                            case "\n\t":
                                                var v180 uint8 = 205
                                                _ = v180
                                            }
                                            switch "" {
                                            case "\n\t":
                                                var v181 uint8 = 253
                                                _ = v181
                                            }
                                            v182 := GenS11{F0: (^(v2))}
                                            _ = v182
                                        }
                                    } else {
                                        if (4294967295 <= (v1 + (v1 * (^(v1))))) {
                                            var v183 uint = 65535
                                            if false {
                                                v184 := GenS17{F0: false, F1: v177}
                                                fmt.Println(v1, v154, v1)
                                                _ = v184
                                            } else {
                                                v185 := GenS13{F0: true, F1: 128, F2: 432}
                                                v177 = v177
                                                v175 = 100
                                                if (!false) {
                                                    if (v154 >= (v154 / 1.0)) {
                                                        v154 = v154
                                                        v2 = (-(v2))
                                                        switch fi168 {
                                                        case 32768:
                                                            v186 := GenS10{F0: int(v183)}
                                                            _ = v186
                                                        default:
                                                            _ = len(v169)
                                                        }
                                                        v177 = v177
                                                        if false {
                                                            var v187 uint64 = v1
                                                            var v188 uint = v183
                                                            fmt.Println(v174, fi168)
                                                            var v189 int64 = (v2 << ((730) & 63))
                                                            _ = v187
                                                            _ = v188
                                                            _ = v189
                                                        } else {
                                                            fmt.Println(v175)
                                                            v3["foo bar"] = fi168
                                                        }
                                                    }
                                                }
                                                _ = v185
                                            }
                                            _ = v183
                                        }
                                        var v190 uint8 = 255
                                        _ = v190
                                    }
                                    fmt.Println(fi168, v1)
                                    _ = v177
                                }
                                v1 += (v1 % ((v1) | 1))
                            } else {
                                if (3.14159 >= float64(1.5)) {
                                    v2 -= 1
                                    v2 = (v2 << ((1) & 63))
                                    if true {
                                        switch (("hello" + "hello") + "foo bar") {
                                        case "hello":
                                            if false {
                                                var v191 uint8 = 2
                                                _ = v191
                                            }
                                        }
                                        var v192 uint8 = 36
                                        v3[(("\n\t" + "foo bar") + "hello")] = (fi167 - fi167)
                                        _ = v192
                                    } else {
                                        v169 = append(v169, fi167)
                                        v193 := GenS8{F0: true, F1: (v2 * (v2 & (v2 - int64(v175)))), F2: v2}
                                        _ = v193
                                    }
                                } else {
                                    v194 := GenS9{F0: 255, F1: genFn0((false && true), 4294967295), F2: false}
                                    var v195 uint = 42
                                    v196 := v3["foo bar"]
                                    _ = v194
                                    _ = v195
                                    _ = v196
                                }
                                switch fi168 {
                                case 424:
                                    v197 := GenS7{F0: 0}
                                    _ = v197
                                case 32767:
                                    v1 = (v1 * 225)
                                case 2147483647:
                                    v198 := GenS18{F0: 2}
                                    _ = v198
                                default:
                                    var v199 int8 = v175
                                    _ = v199
                                }
                            }
                            var v200 uint64 = v1
                            var v201 int32 = 924
                            v2 -= (v2 << ((559) & 63))
                            v202 := GenS10{F0: 127}
                            if ((149 < v1) || false) {
                                fmt.Println(v1, v201, v201)
                                if false {
                                    v200 -= v1
                                } else {
                                    v203 := GenS19{F0: 140, F1: 255}
                                    var v204 float32 = 0.0
                                    _ = v203
                                    _ = v204
                                }
                            } else {
                                v174 += v174
                            }
                            _ = v176
                            _ = v200
                            _ = v201
                            _ = v202
                        }
                        _ = v174
                        _ = v175
                    } else {
                        v2 += v2
                        delete(v3, "\n\t")
                    }
                    if (fi168 == fi168) {
                        _ = len(v3)
                        if (7 < -128) {
                            v2 -= (^(v2))
                            switch 2147483647 {
                            case -579:
                                v1 = (v1 ^ (v1 << ((uint(v1)) & 63)))
                            case -814:
                                v154 -= -2.25
                            }
                        }
                    } else {
                        v154 = -2.25
                        if false {
                            var v205 string = "foo bar"
                            var v206 uint64 = 4294967295
                            v207 := GenS19{F0: 4294967295, F1: 791}
                            v208 := GenS11{F0: (-(v2))}
                            _ = v205
                            _ = v206
                            _ = v207
                            _ = v208
                        }
                        delete(v3, "a")
                        var v209 bool = (true || true)
                        _ = v209
                    }
                }
                switch "\n\t" {
                case "":
                    var v210 uint32 = 659
                    _ = v210
                }
            }
            if len(v169) >= 2 {
                v169 = v169[1:len(v169)]
            }
            if false {
                switch fi168 {
                case 127:
                    var v211 int64 = v2
                    _ = v211
                case -232:
                    switch ("foo bar" + ("\n\t" + (("hello" + "a") + "\n\t"))) {
                    case "foo bar":
                        v3["a"] = fi168
                    default:
                        fmt.Println(fi168, fi167)
                    }
                }
            }
            _ = fi168
            _ = v169
        }
        v212 := GenS20{F0: 521}
        _ = fi167
        _ = v212
    }
    v2 -= v2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v154:", v154)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v154
    _ = v2
    _ = v3
}
