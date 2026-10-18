package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool; F1 int8; F2 uint64 }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 string; F1 uint8; F2 int8 }
type GenS15 struct { F0 uint8; F1 string }
type GenS16 struct { F0 string; F1 uint16 }
type GenS17 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS20 struct { F0 int16; F1 int64 }
type GenS21 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint) int {
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: (fi6 + fi6), F1: 0}
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: fi6, F1: 0}
            var v10 uint32 = 566
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    v11 := []int{-9223372036854775808}
    p0 = 593
    _ = p0
    _ = v11
    return 4294967295
}
func genFn1() int64 {
    v12 := []int{128}
    var v13 uint16 = 2
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            v13 += (v13 / ((0) | 1))
            v16 := GenS1{F0: (fi15 >> ((4294967295) & 63)), F1: 186}
            if len(v12) > 0 {
                v12[0] = fi15
            }
            v17 := GenS1{F0: (fi14 | (fi15 - (fi15 | fi15))), F1: 103}
            switch "foo bar" {
            case "hello":
                switch fi14 {
                case 1:
                    switch -316 {
                    case 32767:
                        switch genFn0(42) {
                        case 0:
                            v13 = 32767
                        }
                        if (true || false) {
                            var v18 uint16 = v13
                            var v19 int8 = 0
                            v20 := GenS1{F0: (fi14 | fi14), F1: 1}
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            var v21 string = (("hello" + "foo bar") + "\n\t")
                            _ = v21
                        }
                        if true {
                            v22 := GenS1{F0: (fi14 & fi15), F1: 42}
                            v13 -= v13
                            _ = v22
                        } else {
                            if true {
                                fmt.Println(fi14)
                            }
                            if (0 > 0) {
                                if true {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    var v23 int8 = 30
                                    var v24 int16 = 686
                                    _ = v23
                                    _ = v24
                                }
                                var v25 bool = (false || true)
                                _ = v25
                            } else {
                                fmt.Println(fi15, fi15, fi14)
                                var v26 int32 = -789
                                switch "" {
                                case "hello":
                                    v12 = append(v12, (fi14 - fi15))
                                case "foo bar":
                                    v27 := GenS1{F0: (fi14 + fi15), F1: 159}
                                    if (42 > 4294967295) {
                                        switch fi15 {
                                        case 0:
                                            if false {
                                                v26 -= (^(v26))
                                                var v28 int64 = -9223372036854775808
                                                v29 := GenS1{F0: -1, F1: 255}
                                                _ = v28
                                                _ = v29
                                            }
                                        case 1:
                                            switch genFn0(4294967295) {
                                            case 676:
                                                fmt.Println(v13)
                                            case 128:
                                                fmt.Println(v13, v13)
                                            }
                                        case -1:
                                            if true {
                                                _ = len(v12)
                                                if (1.5 > float64(float32(0.1))) {
                                                    fmt.Println(v13)
                                                    v13 -= 2
                                                    v26 = v26
                                                } else {
                                                    v12 = append(v12, fi15)
                                                }
                                            }
                                        default:
                                            var v30 int = 128
                                            _ = v30
                                        }
                                        if (!(-32768 >= -1)) {
                                            v31 := GenS2{F0: 42, F1: 401}
                                            _ = v31
                                        } else {
                                            var v32 float32 = 0.0
                                            v13 += v13
                                            _ = v32
                                        }
                                        v33 := GenS2{F0: -872, F1: 2147483647}
                                        _ = v33
                                    } else {
                                        if (true || true) {
                                            var v34 uint8 = 42
                                            _ = v34
                                        } else {
                                            var v35 bool = (!true)
                                            var v36 uint8 = 128
                                            v26 = -2147483648
                                            var v37 bool = v35
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                    }
                                    _ = v27
                                case "\n\t":
                                    if true {
                                        if len(v12) > 0 {
                                            v12[0] = genFn0(170)
                                        }
                                    }
                                default:
                                    v38 := GenS3{F0: 2147483647}
                                    _ = v38
                                }
                                _ = v26
                            }
                            v39 := GenS4{F0: 0, F1: true}
                            _ = v39
                        }
                        switch (fi14 % (((-(fi14))) | 1)) {
                        case 1:
                            var v40 int32 = -2147483648
                            _ = v40
                        case -666:
                            var v41 uint8 = 2
                            _ = v41
                        default:
                            if ((("a" + ("" + ("foo bar" + ""))) + "hello") >= "hello") {
                                if (!(!(1 < 213))) {
                                    v13 += (v13 >> ((255) & 63))
                                    v42 := GenS1{F0: (fi14 % (((fi15 << ((387) & 63))) | 1)), F1: 128}
                                    _ = len(v12)
                                    v43 := GenS5{F0: 1}
                                    _ = v42
                                    _ = v43
                                } else {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    if ((-(fi14)) != genFn0(749)) {
                                        var v44 int = fi14
                                        if true {
                                            var v45 float64 = 3.14159
                                            _ = v45
                                        }
                                        _ = v44
                                    } else {
                                        switch fi15 {
                                        case -802:
                                            v46 := GenS6{F0: 127, F1: 0}
                                            _ = v46
                                        case 934:
                                            v13 -= (^(v13))
                                        }
                                        var v47 float32 = 3.14159
                                        switch fi14 {
                                        case -376:
                                            var v48 int8 = int8(v13)
                                            switch (fi15 % ((fi15) | 1)) {
                                            case 128:
                                                if (138 >= -1) {
                                                    v12 = append(v12, 128)
                                                    v49 := GenS7{F0: 2, F1: 18, F2: 2}
                                                    _ = v49
                                                }
                                                var v50 uint16 = (v13 << ((2) & 63))
                                                switch (fi15 / ((386) | 1)) {
                                                case -9223372036854775808:
                                                    switch 42 {
                                                    case 127:
                                                        v48 -= v48
                                                    case 4294967295:
                                                        var v51 uint32 = 2147483647
                                                        var v52 float64 = float64(v47)
                                                        _ = v51
                                                        _ = v52
                                                    case 42:
                                                        v47 += (float32(-2.25) * 3.14159)
                                                    }
                                                case 2147483647:
                                                    switch (fi15 / (((fi14 / (((fi14 * (fi14 / ((775) | 1)))) | 1))) | 1)) {
                                                    case -9223372036854775808:
                                                        if len(v12) > 0 {
                                                            v12[0] = (fi15 | fi14)
                                                        }
                                                    default:
                                                        v13 = (v13 - 947)
                                                    }
                                                case -349:
                                                    v53 := GenS8{F0: 367}
                                                    fmt.Println(v50)
                                                    _ = v53
                                                }
                                                fmt.Println(v50, v13)
                                                var v54 uint64 = 671
                                                var v55 float64 = 1e10
                                                _ = v50
                                                _ = v54
                                                _ = v55
                                            case -9223372036854775808:
                                                fmt.Println(fi15, fi14)
                                            default:
                                                v47 -= v47
                                            }
                                            _ = v48
                                        case 9223372036854775807:
                                            var v56 uint64 = 826
                                            _ = v56
                                        case 42:
                                            v57 := GenS9{F0: false}
                                            _ = v57
                                        }
                                        v47 -= ((0.0 * v47) / v47)
                                        _ = v47
                                    }
                                    var v58 int16 = -32768
                                    _ = v58
                                }
                                var v59 float64 = (1e10 * (((3.14159 / ((float64(3.14159) - ((-2.25 - 1e10) + float64(float32(((((((0.1 * 3.14159) + (1e10 + 0.0)) / (1.5 * 1e10)) / 0.1) - 3.14159) / 1.0))))) + (float64(0.1) - ((0.0 * ((3.14159 - 1e10) + ((((0.0 + 0.1) / 0.1) + 3.14159) - 3.14159))) * -2.25)))) + 1e10) + ((0.1 + ((3.14159 - -2.25) / 3.14159)) / ((3.14159 * 1e10) / 3.14159))))
                                v13 = (v13 + (v13 % ((v13) | 1)))
                                _ = v59
                            } else {
                                v60 := GenS10{F0: 42}
                                v61 := GenS7{F0: 4294967295, F1: 995, F2: 2}
                                if (((^(v13)) > v13) && (190 != 255)) {
                                    var v62 uint64 = 412
                                    v62 -= (v62 - 2)
                                    v13 = v13
                                    v62 -= (v62 >> ((65535) & 63))
                                    v13 = 1
                                    _ = v62
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        fmt.Println(fi14, v13)
                                    case "":
                                        if false {
                                            if ((^(v13)) < (v13 ^ 826)) {
                                                if false {
                                                    if len(v12) > 0 {
                                                        v12[0] = fi14
                                                    }
                                                } else {
                                                    v13 = (v13 - (v13 * v13))
                                                    v63 := GenS3{F0: fi14}
                                                    _ = v63
                                                }
                                                var v64 uint = 0
                                                if true {
                                                    if false {
                                                        var v65 int = (fi15 << (((v64 >> (((v64 << (((^(v64))) & 63))) & 63))) & 63))
                                                        switch "" {
                                                        case "hello":
                                                            fmt.Println(v13)
                                                        case "\n\t":
                                                            v13 -= v13
                                                            v13 = v13
                                                            if true {
                                                                if ((1 < 1) && false) {
                                                                    v66 := GenS5{F0: 2147483648}
                                                                    _ = v66
                                                                } else {
                                                                    if (false || false) {
                                                                        fmt.Println(fi14, v65, fi15)
                                                                        if true {
                                                                            if len(v12) >= 2 {
                                                                                v12 = v12[1:len(v12)]
                                                                            }
                                                                            if false {
                                                                                if true {
                                                                                    if false {
                                                                                        if false {
                                                                                            var v67 int16 = -889
                                                                                            var v68 int32 = -948
                                                                                            var v69 string = ((("" + ("hello" + (("foo bar" + (("hello" + "") + "")) + "\n\t"))) + "\n\t") + ("hello" + "\n\t"))
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                            _ = v69
                                                                                        } else {
                                                                                            v70 := GenS1{F0: genFn0((^(v64))), F1: 0}
                                                                                            _ = v70
                                                                                        }
                                                                                        v13 = 71
                                                                                    }
                                                                                    v64 = v64
                                                                                    v65 += (fi14 << ((0) & 63))
                                                                                }
                                                                                if (!((0 > 4294967295) || (1e10 > 1.5))) {
                                                                                    var v71 int = int(v13)
                                                                                    if false {
                                                                                        v65 -= (v65 + genFn0(1))
                                                                                    }
                                                                                    _ = v71
                                                                                } else {
                                                                                    if len(v12) > 0 {
                                                                                        v12[0] = (fi15 >> (((v64 % ((v64) | 1))) & 63))
                                                                                    }
                                                                                    var v72 float32 = 1.5
                                                                                    _ = v72
                                                                                }
                                                                                v73 := GenS11{F0: 617}
                                                                                _ = v73
                                                                            }
                                                                            v65 += fi14
                                                                        } else {
                                                                            switch (("foo bar" + "a") + "a") {
                                                                            case "hello":
                                                                                var v74 int8 = -123
                                                                                v75 := GenS2{F0: -32768, F1: 0}
                                                                                _ = v74
                                                                                _ = v75
                                                                            default:
                                                                                var v76 int64 = 0
                                                                                _ = v76
                                                                            }
                                                                            var v77 bool = true
                                                                            _ = v77
                                                                        }
                                                                    }
                                                                    switch fi15 {
                                                                    case -9223372036854775808:
                                                                        v78 := GenS12{F0: true, F1: 100, F2: 42}
                                                                        _ = v78
                                                                    }
                                                                    fmt.Println(v64)
                                                                }
                                                                v64 = v64
                                                                switch v65 {
                                                                case -569:
                                                                    fmt.Println(fi15)
                                                                case 128:
                                                                    v79 := GenS7{F0: 4294967295, F1: 737, F2: 415}
                                                                    _ = v79
                                                                }
                                                            } else {
                                                                switch "hello" {
                                                                case "a":
                                                                    if (-2.25 <= (float32(float64(0.1)) + 0.0)) {
                                                                        if (!(128 == 1)) {
                                                                            var v80 float64 = 0.1
                                                                            _ = v80
                                                                        } else {
                                                                            v64 += v64
                                                                            var v81 int = int(v13)
                                                                            _ = v81
                                                                        }
                                                                    } else {
                                                                        var v82 uint16 = (^(v13))
                                                                        _ = v82
                                                                    }
                                                                }
                                                                switch "" {
                                                                case "":
                                                                    v83 := GenS13{F0: ("\n\t" + ""), F1: 65535}
                                                                    _ = v83
                                                                }
                                                                switch (fi14 - fi15) {
                                                                case 0:
                                                                    var v84 uint32 = 4294967295
                                                                    _ = v84
                                                                case 9223372036854775807:
                                                                    v64 += (v64 / ((v64) | 1))
                                                                    v65 -= fi15
                                                                    if true {
                                                                        v64 = v64
                                                                    }
                                                                default:
                                                                    switch (fi14 << (((v64 | v64)) & 63)) {
                                                                    case 32768:
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            var v85 int8 = 75
                                                                            _ = v85
                                                                        default:
                                                                            v86 := GenS14{F0: (("foo bar" + "a") + "\n\t"), F1: 127, F2: 0}
                                                                            _ = v86
                                                                        }
                                                                    case 435:
                                                                        v12 = append(v12, (^(fi15)))
                                                                    case -603:
                                                                        var v87 uint64 = 1
                                                                        _ = v87
                                                                    default:
                                                                        var v88 uint = v64
                                                                        _ = v88
                                                                    }
                                                                }
                                                                v64 = 1
                                                            }
                                                        case "a":
                                                            var v89 int32 = 2147483647
                                                            _ = v89
                                                        }
                                                        _ = v65
                                                    }
                                                    var v90 int64 = 9223372036854775807
                                                    var v91 uint = (v64 | (v64 * 18446744073709551615))
                                                    _ = v90
                                                    _ = v91
                                                }
                                                _ = v64
                                            } else {
                                                if false {
                                                    if true {
                                                        switch ("a" + ("a" + "")) {
                                                        case "foo bar":
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                        }
                                                        var v92 uint8 = 0
                                                        fmt.Println(v13, v13, fi15)
                                                        _ = v92
                                                    }
                                                    v13 = (v13 - 32767)
                                                }
                                            }
                                            fmt.Println(fi14, v13, fi14)
                                        } else {
                                            var v93 int64 = 42
                                            _ = v93
                                        }
                                    case "hello":
                                        switch genFn0(255) {
                                        case 127:
                                            fmt.Println(fi15, v13, fi15)
                                        case 235:
                                            fmt.Println(fi14, fi15)
                                        }
                                    default:
                                        if false {
                                            fmt.Println(fi14, fi15, fi14)
                                            _ = len(v12)
                                            var v94 uint64 = 18446744073709551615
                                            v95 := GenS7{F0: (v94 + 418), F1: 65535, F2: (^(v94))}
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                    v96 := GenS5{F0: 746}
                                    var v97 float32 = 3.14159
                                    _ = v96
                                    _ = v97
                                }
                                v13 = 32767
                                _ = v60
                                _ = v61
                            }
                            switch (("hello" + ("hello" + "foo bar")) + "") {
                            case "\n\t":
                                var v98 bool = true
                                _ = v98
                            case "":
                                var v99 uint64 = 2
                                _ = v99
                            default:
                                fmt.Println(fi14, fi15)
                            }
                            v100 := GenS15{F0: 127, F1: "foo bar"}
                            _ = v100
                        }
                        if (42 > 579) {
                            fmt.Println(fi14)
                            var v101 float64 = 0.1
                            _ = v101
                        }
                    case -378:
                        v102 := GenS14{F0: "hello", F1: 255, F2: 127}
                        _ = v102
                    case 560:
                        switch -9223372036854775808 {
                        case 32768:
                            switch ("foo bar" + ((("hello" + "foo bar") + ("foo bar" + ("hello" + ("" + ("hello" + "\n\t"))))) + "hello")) {
                            case "hello":
                                if (!((!(!false)) || false)) {
                                    switch -9223372036854775808 {
                                    case 956:
                                        if ((0.1 <= (0.0 * 1e10)) && (181 < (v13 << ((2) & 63)))) {
                                            v103 := GenS16{F0: "hello", F1: v13}
                                            _ = v103
                                        }
                                    }
                                    var v104 uint = 462
                                    _ = v104
                                } else {
                                    if true {
                                        var v105 int8 = -1
                                        _ = v105
                                    } else {
                                        switch ((("" + ("\n\t" + ("\n\t" + ("hello" + "hello")))) + "\n\t") + "a") {
                                        case "a":
                                            v12 = append(v12, (fi15 >> ((1) & 63)))
                                        case "hello":
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        }
                                        switch ("a" + ("\n\t" + "foo bar")) {
                                        case "hello":
                                            fmt.Println(v13)
                                        case "a":
                                            v106 := GenS2{F0: 32767, F1: 1}
                                            _ = v106
                                        case "\n\t":
                                            var v107 int = (^(fi14))
                                            _ = v107
                                        default:
                                            v108 := GenS16{F0: "hello", F1: (v13 | (^(v13)))}
                                            _ = v108
                                        }
                                        v109 := GenS14{F0: "a", F1: 1, F2: -70}
                                        v13 = v13
                                        _ = v109
                                    }
                                }
                            case "":
                                var v110 uint16 = v13
                                _ = v110
                            }
                        case 740:
                            var v111 uint = 42
                            _ = v111
                        }
                    default:
                        fmt.Println(v13, v13)
                        v13 = v13
                    }
                case 2147483647:
                    var v112 uint64 = 122
                    _ = v112
                case 128:
                    v13 -= (v13 * (v13 ^ 0))
                }
                var v113 int64 = 487
                _ = v113
            case "\n\t":
                v13 -= v13
            default:
                switch fi14 {
                case -1:
                    if ((v13 << ((2) & 63)) != (v13 + (v13 * (v13 | 295)))) {
                        if false {
                            if (!(!true)) {
                                v12 = append(v12, (fi14 % (((fi15 / ((4294967295) | 1))) | 1)))
                                if len(v12) > 0 {
                                    v12[0] = fi14
                                }
                            } else {
                                var v114 uint32 = 0
                                if (!true) {
                                    v13 = v13
                                    if true {
                                        var v115 uint = 255
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        var v116 uint = (^(v115))
                                        _ = v115
                                        _ = v116
                                    }
                                    fmt.Println(v114, fi15)
                                } else {
                                    v12 = append(v12, fi15)
                                }
                                if false {
                                    fmt.Println(fi14, v13, fi15)
                                }
                                var v117 bool = (1000 <= 42)
                                if ((!false) || ((fi14 > (fi15 & (fi14 >> ((65535) & 63)))) || false)) {
                                    v114 += (v114 + v114)
                                }
                                v118 := GenS17{F0: (^(v114)), F1: 2, F2: 0}
                                if len(v12) > 0 {
                                    v12[0] = (^(fi14))
                                }
                                _ = v114
                                _ = v117
                                _ = v118
                            }
                            var v119 int32 = -2147483648
                            _ = v119
                        } else {
                            v120 := GenS3{F0: genFn0(424)}
                            var v121 int16 = -1
                            switch (fi14 & (fi14 >> ((1) & 63))) {
                            case 9223372036854775807:
                                v12 = append(v12, fi15)
                            case 32768:
                                var v122 float32 = ((1e10 * 0.0) - 1e10)
                                _ = v122
                            }
                            _ = v120
                            _ = v121
                        }
                    } else {
                        v123 := GenS18{F0: 600}
                        _ = v123
                    }
                }
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi124 := 0; fi124 < 14; fi124++ {
            var v125 uint64 = 867
            v13 -= (v13 % ((2) | 1))
            v13 -= 65535
            if (!(genFn0(4294967295) > -822)) {
                if true {
                    if ((^(v13)) != v13) {
                        switch genFn0(1) {
                        case 450:
                            if (!true) {
                                if ((true && true) && (42 != 0)) {
                                    switch fi124 {
                                    case -2147483648:
                                        var v126 bool = (4294967295 >= fi14)
                                        _ = v126
                                    case 42:
                                        var v127 float32 = 3.14159
                                        _ = v127
                                    case -562:
                                        switch (fi14 * fi14) {
                                        case -1:
                                            v125 += (v125 >> ((4294967295) & 63))
                                        case 998:
                                            switch ("" + "\n\t") {
                                            case "":
                                                v13 += v13
                                            case "\n\t":
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                            case "a":
                                                var v128 int32 = 2147483647
                                                _ = v128
                                            }
                                        case 959:
                                            fmt.Println(v125, v125)
                                        }
                                    default:
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                    v125 -= v125
                                    fmt.Println(fi124)
                                } else {
                                    v129 := GenS9{F0: false}
                                    var v130 uint8 = 255
                                    _ = v129
                                    _ = v130
                                }
                                v131 := GenS4{F0: -569, F1: (3.14159 > (float64(float32(1.5)) + 1e10))}
                                v12 = append(v12, fi124)
                                v125 -= v125
                                var v132 bool = false
                                _ = v131
                                _ = v132
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    var v133 int8 = -1
                                    _ = v133
                                case "hello":
                                    if false {
                                        fmt.Println(v125)
                                        v13 = (v13 / ((v13) | 1))
                                        switch 32767 {
                                        case 128:
                                            var v134 bool = (false && (3.14159 <= 0.1))
                                            _ = v134
                                        case 9223372036854775807:
                                            v135 := GenS16{F0: "hello", F1: (v13 * (^(v13)))}
                                            _ = v135
                                        }
                                    }
                                }
                            }
                        case 32767:
                            fmt.Println(fi124, v125)
                        default:
                            var v136 uint16 = 1
                            _ = v136
                        }
                        v12 = append(v12, (fi14 >> ((4294967295) & 63)))
                        v13 = (v13 + 65535)
                    } else {
                        fmt.Println(v125, fi124, v13)
                        v13 += v13
                        var v137 string = ("hello" + ((((((((("" + ("\n\t" + (((((("foo bar" + "foo bar") + "a") + "hello") + "foo bar") + "\n\t") + ("a" + "")))) + ("hello" + (("" + "") + "\n\t"))) + ("\n\t" + ("\n\t" + "foo bar"))) + "foo bar") + "a") + ("" + ("a" + "\n\t"))) + "a") + ("" + ("a" + "\n\t"))) + ""))
                        _ = v137
                    }
                    v125 -= 4294967295
                }
            }
            fmt.Println(fi124)
            _ = fi124
            _ = v125
        }
        _ = fi14
    }
    v13 += (v13 - (^(v13)))
    for fi138 := 0; fi138 < 20; fi138++ {
        var v139 uint = 2
        v140 := GenS7{F0: 42, F1: 190, F2: 4294967295}
        _ = fi138
        _ = v139
        _ = v140
    }
    _ = v12
    _ = v13
    return 889
}

func main() {
    var v1 int8 = 1
    var v2 int64 = 0
    v3 := []uint32{339, 67}
    var v4 int8 = v1
    var v5 uint16 = 0
    v4 = -30
    v4 = 76
    for fi141 := 0; fi141 < 18; fi141++ {
        switch "" {
        case "":
            if false {
                if true {
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    v1 = -105
                } else {
                    v5 -= v5
                }
                var v142 bool = false
                v143 := GenS15{F0: 32, F1: "a"}
                _ = v142
                _ = v143
            } else {
                v3 = append(v3, 65535)
            }
        case "a":
            var v144 int64 = v2
            _ = v144
        case "hello":
            v145 := GenS17{F0: 1, F1: 2, F2: (v4 / (((v4 | v4)) | 1))}
            _ = v145
        default:
            _ = len(v3)
        }
        for fi146 := 0; fi146 < 2; fi146++ {
            v147 := GenS19{F0: 0, F1: v5, F2: -999}
            if (1000 <= 42) {
                v5 -= 255
                if (0.0 == 1.5) {
                    v1 -= (v1 & (v4 & 100))
                    v5 = (v5 & v5)
                    switch -9223372036854775808 {
                    case -704:
                        v148 := GenS5{F0: 2147483648}
                        _ = v148
                    default:
                        v1 += (-(v4))
                    }
                    var v149 bool = true
                    _ = v149
                }
                v150 := GenS20{F0: 1, F1: 1}
                v151 := GenS21{F0: 42, F1: 722}
                if true {
                    v152 := GenS16{F0: "hello", F1: (v5 * 0)}
                    v3 = append(v3, 228)
                    var v153 uint16 = 188
                    _ = v152
                    _ = v153
                }
                _ = v150
                _ = v151
            } else {
                v4 -= v4
                v154 := GenS11{F0: 528}
                _ = v154
            }
            var v155 uint16 = (^(v5))
            _ = fi146
            _ = v147
            _ = v155
        }
        if (genFn0(65535) <= fi141) {
            for fi156 := 0; fi156 < 9; fi156++ {
                var v157 bool = (!(42 < 18446744073709551615))
                var v158 bool = (!v157)
                _ = fi156
                _ = v157
                _ = v158
            }
            var v159 uint16 = (v5 % (((v5 & (v5 * 255))) | 1))
            v160 := GenS9{F0: true}
            _ = v159
            _ = v160
        } else {
            switch "\n\t" {
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = uint32(v5)
                }
            }
            v161 := GenS18{F0: 0}
            fmt.Println(v4, v4, v5)
            _ = v161
        }
        _ = fi141
    }
    fmt.Println(v1, v2)
    v5 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v162 int32 = 42
    fmt.Println(v162)
    v162 -= v162
    var v163 string = "hello"
    fmt.Println(v2, v162, v162)
    var v164 float32 = -2.25
    var v165 float64 = ((0.0 - 1.5) / 0.1)
    for fi166 := 0; fi166 < 7; fi166++ {
        for fi167 := 0; fi167 < 6; fi167++ {
            v4 = (v4 + v4)
            v4 = (v4 & v1)
            _ = fi167
        }
        _ = fi166
    }
    var v168 uint16 = (v5 / (((^(v5))) | 1))
    for fi169 := 0; fi169 < 11; fi169++ {
        if (!(v164 == ((v164 / v164) + v164))) {
            for fi170 := 0; fi170 < 20; fi170++ {
                v171 := GenS20{F0: -32768, F1: (v2 * (v2 ^ v2))}
                if false {
                    v1 = (v1 & v4)
                }
                _ = fi170
                _ = v171
            }
            for fi172 := 0; fi172 < 5; fi172++ {
                fmt.Println(v165, v164, v5)
                _ = fi172
            }
            v162 += (v162 & -403)
        } else {
            var v173 uint64 = 1
            for fi174 := 0; fi174 < 17; fi174++ {
                v173 -= v173
                _ = fi174
            }
            _ = v173
        }
        var v175 string = v163
        _ = fi169
        _ = v175
    }
    fmt.Println("v1:", v1)
    fmt.Println("v162:", v162)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("v168:", v168)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v162
    _ = v163
    _ = v164
    _ = v165
    _ = v168
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
