package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int16; F2 int32 }
type GenS4 struct { F0 int; F1 int8; F2 int8 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int64; F2 bool }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32; F1 int64 }
type GenS12 struct { F0 int; F1 bool; F2 int32 }
type GenS13 struct { F0 int32; F1 string; F2 bool }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint64; F1 int32; F2 int64 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS20 struct { F0 uint8 }
func genFn0(p0 int8) uint32 {
    p0 = p0
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6)
        var v7 string = "a"
        _ = fi6
        _ = v7
    }
    p0 += (p0 >> ((42) & 63))
    _ = p0
    return 322
}
func genFn1() uint8 {
    for fi8 := 0; fi8 < 1; fi8++ {
        switch fi8 {
        case -1:
            var v9 int = (fi8 / (((fi8 / ((fi8) | 1))) | 1))
            _ = v9
        }
        _ = fi8
    }
    var v10 uint16 = 32767
    v10 -= v10
    switch "" {
    case "\n\t":
        for fi11 := 0; fi11 < 17; fi11++ {
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                var v12 uint16 = (v10 >> ((111) & 63))
                _ = v12
            case "a":
                if false {
                    v10 = (v10 >> ((65535) & 63))
                    if true {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            v14 := GenS1{F0: ("hello" + ("a" + "foo bar")), F1: 619}
                            v10 = v10
                            var v15 uint = 402
                            v10 += (^(v10))
                            fmt.Println(v15, fi11)
                            var v16 uint64 = 2
                            v17 := GenS1{F0: ("hello" + ("hello" + "hello")), F1: 42}
                            fmt.Println(v15, v16, fi11)
                            v15 = (v15 & v15)
                            v16 = v16
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        if false {
                            v18 := GenS1{F0: ("\n\t" + "\n\t"), F1: 42}
                            var v19 int64 = -765
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS1{F0: "", F1: -495}
                        _ = v20
                    } else {
                        v21 := GenS2{F0: "hello", F1: uint8(v10)}
                        switch "a" {
                        case "":
                            if true {
                                v22 := GenS3{F0: 72, F1: 42, F2: -2147483648}
                                _ = v22
                            }
                        default:
                            fmt.Println(fi11, v10)
                        }
                        switch "a" {
                        case "\n\t":
                            for fi23 := 0; fi23 < 12; fi23++ {
                                var v24 int = fi23
                                v25 := []int{(v24 >> ((42) & 63))}
                                if ((v10 * (v10 ^ v10)) != (v10 | v10)) {
                                    if len(v25) > 0 {
                                        v25[0] = (^(fi11))
                                    }
                                }
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        case "foo bar":
                            var v26 uint16 = v10
                            _ = v26
                        case "":
                            if (!false) {
                                fmt.Println(fi11, v10)
                                v27 := []int{(fi11 - fi11), fi11, (fi11 >> ((804) & 63)), (fi11 + (fi11 * fi11))}
                                var v28 float32 = float32(0.0)
                                _ = v27
                                _ = v28
                            }
                        }
                        _ = v21
                    }
                    v10 = (v10 | (v10 | v10))
                }
            case "foo bar":
                if (1e10 >= (3.14159 / 1.5)) {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        var v30 uint64 = 4294967295
                        _ = fi29
                        _ = v30
                    }
                }
            }
            if (0.1 > 0.1) {
                v10 += (v10 ^ v10)
                switch ("a" + "a") {
                case "a":
                    switch fi11 {
                    case 2147483647:
                        for fi31 := 0; fi31 < 2; fi31++ {
                            fmt.Println(v10)
                            switch fi31 {
                            case 1:
                                v10 -= (^(v10))
                            case -9223372036854775808:
                                if false {
                                    var v32 int16 = 32767
                                    _ = v32
                                } else {
                                    v10 -= v10
                                    if true {
                                        fmt.Println(fi11, v10)
                                        var v33 bool = (66 > 1)
                                        v34 := []int{(fi11 + 571), 2147483647, (fi31 + fi31), 32767}
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v10 += v10
                                        var v35 uint16 = v10
                                        _ = v35
                                    }
                                }
                                if (!("" != ((("" + "hello") + (("foo bar" + "") + "\n\t")) + (("\n\t" + "a") + ("hello" + "a"))))) {
                                    if true {
                                        var v36 int = (fi11 + (fi31 - fi11))
                                        var v37 uint8 = 1
                                        _ = v36
                                        _ = v37
                                    } else {
                                        var v38 int32 = -1
                                        _ = v38
                                    }
                                    v10 += 32767
                                    if true {
                                        v39 := []int{fi11, fi31, fi11, fi31}
                                        v10 += v10
                                        _ = v39
                                    } else {
                                        switch "a" {
                                        case "a":
                                            var v40 int = 9223372036854775807
                                            v41 := GenS1{F0: "", F1: -951}
                                            _ = v40
                                            _ = v41
                                        case "":
                                            var v42 uint8 = 128
                                            _ = v42
                                        }
                                        switch ("" + "hello") {
                                        case "a":
                                            if true {
                                                v43 := GenS1{F0: "hello", F1: 1}
                                                fmt.Println(fi11, fi11, fi31)
                                                v44 := []int{}
                                                _ = v43
                                                _ = v44
                                            } else {
                                                fmt.Println(fi31, fi11, fi11)
                                                if (false || true) {
                                                    var v45 string = "\n\t"
                                                    var v46 uint = 2
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v47 := GenS2{F0: ("\n\t" + ""), F1: 127}
                                                    _ = v47
                                                }
                                            }
                                        case "foo bar":
                                            fmt.Println(fi31, v10, v10)
                                        default:
                                            v48 := []int{127, (fi11 >> ((203) & 63)), -9223372036854775808, fi11}
                                            _ = v48
                                        }
                                    }
                                } else {
                                    var v49 int16 = 32767
                                    v50 := []int{}
                                    v51 := GenS3{F0: 255, F1: (-(v49)), F2: 2147483647}
                                    v10 -= (v10 - 1)
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                            case 2147483647:
                                fmt.Println(fi31)
                            }
                            if true {
                                var v52 bool = false
                                _ = v52
                            } else {
                                v53 := []int{42}
                                v10 = (^(v10))
                                if false {
                                    if len(v53) > 0 {
                                        v53[0] = 2147483647
                                    }
                                    v10 += v10
                                    switch "hello" {
                                    case "hello":
                                        v53 = append(v53, -419)
                                        v10 = (^(v10))
                                    case "foo bar":
                                        switch "a" {
                                        case "":
                                            fmt.Println(v10)
                                        default:
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(fi11, v10, v10)
                                            case "":
                                                v10 = v10
                                            case "hello":
                                                fmt.Println(fi31, v10, fi31)
                                            }
                                        }
                                    case "":
                                        v10 += v10
                                        _ = len(v53)
                                        fmt.Println(fi31, fi11)
                                    }
                                    v54 := GenS3{F0: 128, F1: -32768, F2: 1}
                                    _ = v54
                                }
                                _ = v53
                            }
                            if (!true) {
                                fmt.Println(fi11, v10)
                                if true {
                                    var v55 int64 = 9223372036854775807
                                    var v56 uint64 = 1
                                    fmt.Println(fi11)
                                    v57 := GenS1{F0: "foo bar", F1: -777}
                                    fmt.Println(v10, v55)
                                    var v58 int32 = 1
                                    var v59 float64 = 1.5
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                } else {
                                    v10 += v10
                                }
                                var v60 int32 = 1
                                _ = v60
                            } else {
                                if true {
                                    v61 := []int{(-(fi31))}
                                    if ("foo bar" != "foo bar") {
                                        v10 += v10
                                        var v62 float64 = float64(1.5)
                                        fmt.Println(fi11, v62, fi11)
                                        _ = v62
                                    } else {
                                        v10 = v10
                                        v63 := GenS3{F0: 42, F1: -603, F2: 42}
                                        switch 1 {
                                        case 127:
                                            var v64 uint8 = 254
                                            _ = v64
                                        case 2147483647:
                                            var v65 int16 = -929
                                            _ = v65
                                        }
                                        _ = v63
                                    }
                                    switch (fi11 / ((fi31) | 1)) {
                                    case 127:
                                        v66 := GenS4{F0: fi31, F1: -104, F2: -75}
                                        _ = v66
                                    case 584:
                                        v10 += 989
                                    default:
                                        if len(v61) > 0 {
                                            v61[0] = (^(fi31))
                                        }
                                    }
                                    _ = v61
                                } else {
                                    v67 := []int{25, fi11, (fi31 ^ 358)}
                                    v68 := GenS5{F0: (fi31 % ((fi11) | 1)), F1: 32767}
                                    _ = v67
                                    _ = v68
                                }
                                switch ((("foo bar" + "\n\t") + ("" + "")) + (("hello" + "a") + "foo bar")) {
                                case "foo bar":
                                    if (-601 <= 0) {
                                        v69 := GenS6{F0: 2}
                                        var v70 int16 = -631
                                        _ = v69
                                        _ = v70
                                    }
                                }
                            }
                            v10 -= v10
                            _ = fi31
                        }
                    case 707:
                        fmt.Println(fi11)
                    case 0:
                        if (((-2.25 - (1e10 / 3.14159)) - 3.14159) > float64(1.5)) {
                            var v71 int32 = 272
                            v72 := []int{(fi11 + (^(fi11))), (fi11 | 2147483647), -1}
                            fmt.Println(v10)
                            _ = v71
                            _ = v72
                        } else {
                            var v73 int64 = 42
                            _ = v73
                        }
                    default:
                        for fi74 := 0; fi74 < 11; fi74++ {
                            fmt.Println(fi11, fi74)
                            v75 := []int{(^(fi11)), 32768, fi11, -606}
                            if true {
                                v76 := GenS2{F0: ("foo bar" + "foo bar"), F1: 1}
                                _ = v76
                            }
                            _ = fi74
                            _ = v75
                        }
                    }
                default:
                    switch "a" {
                    case "hello":
                        if ((255 >= 18446744073709551615) || (!((2 <= 1) || (1 > 1)))) {
                            v10 = (v10 - 255)
                            if (true && false) {
                                v77 := []int{fi11}
                                var v78 float32 = (0.1 - 3.14159)
                                _ = v77
                                _ = v78
                            } else {
                                if (fi11 <= fi11) {
                                    var v79 string = "a"
                                    _ = v79
                                } else {
                                    if (float32(((-2.25 * 0.0) * float64(float32(1.5)))) <= (float32(0.1) + (3.14159 - 3.14159))) {
                                        if true {
                                            for fi80 := 0; fi80 < 1; fi80++ {
                                                fmt.Println(fi11, v10)
                                                v10 -= v10
                                                var v81 float32 = 1e10
                                                var v82 uint8 = 39
                                                var v83 uint8 = 128
                                                _ = fi80
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            }
                                            v84 := []int{}
                                            _ = v84
                                        }
                                    } else {
                                        var v85 int64 = -565
                                        v85 = (v85 | (v85 & v85))
                                        _ = v85
                                    }
                                    var v86 uint = 18446744073709551615
                                    _ = v86
                                }
                                v10 = v10
                                var v87 uint16 = 193
                                _ = v87
                            }
                        }
                    default:
                        fmt.Println(v10, fi11)
                    }
                }
                for fi88 := 0; fi88 < 17; fi88++ {
                    switch "\n\t" {
                    case "hello":
                        var v89 uint8 = 42
                        _ = v89
                    case "a":
                        var v90 uint64 = 42
                        _ = v90
                    case "foo bar":
                        if (-9223372036854775808 == 1) {
                            v91 := []int{fi88, (fi11 ^ (fi11 & fi88)), fi88, (fi88 % (((-(fi88))) | 1))}
                            switch (fi88 * 966) {
                            case -779:
                                v91 = append(v91, (-(fi11)))
                            case -2147483648:
                                if (18446744073709551615 >= 4294967295) {
                                    fmt.Println(fi11, v10, fi88)
                                    var v92 string = "hello"
                                    _ = v92
                                }
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    switch "hello" {
                                    case "":
                                        var v93 uint8 = 239
                                        _ = v93
                                    case "a":
                                        var v94 int16 = 42
                                        _ = v94
                                    default:
                                        var v95 int16 = 948
                                        _ = v95
                                    }
                                case "hello":
                                    if true {
                                        v10 += 2
                                        if false {
                                            v96 := GenS7{F0: 128, F1: 42, F2: (!true)}
                                            if (-880 < -6) {
                                                var v97 uint32 = 4294967295
                                                _ = v97
                                            } else {
                                                v10 = (v10 & 42)
                                                if false {
                                                    v98 := GenS7{F0: 127, F1: -1, F2: (!(65535 != 18446744073709551615))}
                                                    if false {
                                                        v10 -= (v10 - v10)
                                                    }
                                                    if len(v91) >= 2 {
                                                        v91 = v91[1:len(v91)]
                                                    }
                                                    _ = v98
                                                }
                                                switch (-(fi11)) {
                                                case 9223372036854775807:
                                                    v10 = 0
                                                case -539:
                                                    switch (fi88 ^ -52) {
                                                    case 117:
                                                        if (!(-1 > 1000)) {
                                                            v10 = 1
                                                            v99 := GenS3{F0: 74, F1: -1, F2: -1}
                                                            if len(v91) >= 2 {
                                                                v91 = v91[1:len(v91)]
                                                            }
                                                            _ = v99
                                                        }
                                                    default:
                                                        var v100 int64 = 418
                                                        _ = v100
                                                    }
                                                case -1:
                                                    var v101 uint8 = 184
                                                    _ = v101
                                                }
                                            }
                                            _ = v96
                                        } else {
                                            fmt.Println(fi11)
                                            _ = len(v91)
                                        }
                                        var v102 uint16 = (v10 / (((v10 << ((0) & 63))) | 1))
                                        _ = v102
                                    }
                                default:
                                    if (!false) {
                                        var v103 uint32 = 2147483648
                                        _ = v103
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            v104 := GenS8{F0: -9223372036854775808, F1: 563}
                                            _ = v104
                                        case "hello":
                                            v91 = append(v91, (^(fi88)))
                                        case "":
                                            _ = len(v91)
                                        }
                                    }
                                }
                            }
                            v105 := GenS8{F0: fi11, F1: 547}
                            _ = v105
                            _ = v91
                        } else {
                            var v106 float32 = ((-2.25 + (0.1 + float32(1e10))) / 0.1)
                            v107 := []int{}
                            switch -832 {
                            case 128:
                                if len(v107) > 0 {
                                    v107[0] = (-(fi88))
                                }
                            case 32768:
                                var v108 uint = 255
                                _ = v108
                            }
                            v10 -= (v10 / ((255) | 1))
                            if (2 > 65535) {
                                if len(v107) >= 2 {
                                    v107 = v107[1:len(v107)]
                                }
                                var v109 uint64 = 813
                                _ = v109
                            }
                            _ = v106
                            _ = v107
                        }
                    default:
                        v110 := GenS8{F0: fi11, F1: 0}
                        if false {
                            switch fi11 {
                            case 127:
                                fmt.Println(v10, fi11, fi88)
                            case 50:
                                if (fi88 <= 4294967295) {
                                    var v111 float32 = 1.5
                                    v111 = 0.1
                                    _ = v111
                                } else {
                                    if false {
                                        v112 := GenS4{F0: fi11, F1: 100, F2: -116}
                                        v113 := []int{}
                                        v10 -= (v10 + v10)
                                        _ = v112
                                        _ = v113
                                    }
                                    fmt.Println(fi11, fi88)
                                    v10 = 145
                                }
                            case 0:
                                switch "a" {
                                case "foo bar":
                                    var v114 string = "foo bar"
                                    _ = v114
                                case "a":
                                    if true {
                                        var v115 float32 = 1.5
                                        if false {
                                            var v116 int = fi11
                                            v116 += 4294967295
                                            _ = v116
                                        }
                                        _ = v115
                                    }
                                }
                            }
                        } else {
                            if (fi88 < 42) {
                                v10 += (v10 + v10)
                            } else {
                                if true {
                                    var v117 int8 = -21
                                    v118 := []int{(fi88 & fi88), (-(fi11)), 42, fi11}
                                    if (!(genFn0(v117) != genFn0(-1))) {
                                        v117 = v117
                                        if false {
                                            if len(v118) >= 2 {
                                                v118 = v118[1:len(v118)]
                                            }
                                            fmt.Println(fi88, fi88)
                                            if len(v118) > 0 {
                                                v118[0] = fi88
                                            }
                                        }
                                        v117 += (v117 / (((v117 + v117)) | 1))
                                        if (!("foo bar" >= ((("a" + "a") + ("hello" + "a")) + "a"))) {
                                            var v119 bool = (v10 <= v10)
                                            v117 = v117
                                            v10 = 255
                                            _ = v119
                                        } else {
                                            var v120 int32 = 42
                                            _ = len(v118)
                                            _ = v120
                                        }
                                    } else {
                                        v10 = v10
                                        if false {
                                            if false {
                                                var v121 uint16 = v10
                                                var v122 int = 4294967295
                                                _ = v121
                                                _ = v122
                                            }
                                        }
                                    }
                                    _ = v117
                                    _ = v118
                                }
                                v123 := []int{fi88, fi88}
                                switch "foo bar" {
                                case "foo bar":
                                    var v124 uint32 = 1
                                    _ = v124
                                case "":
                                    if ((-319 != -788) && (65535 <= 255)) {
                                        var v125 uint32 = genFn0(127)
                                        var v126 int8 = 1
                                        if ((float64(((1.5 - ((0.1 + ((1.5 - (0.0 - float32((-2.25 * 0.0)))) - ((-2.25 * 1.5) - (1.5 * -2.25)))) * 0.1)) + 0.1)) < 1e10) || false) {
                                            if true {
                                                if false {
                                                    if len(v123) >= 2 {
                                                        v123 = v123[1:len(v123)]
                                                    }
                                                    v10 = 0
                                                    if len(v123) > 0 {
                                                        v123[0] = 32768
                                                    }
                                                } else {
                                                    var v127 uint8 = 255
                                                    var v128 int = fi88
                                                    v129 := GenS7{F0: v127, F1: 0, F2: (814 == -2147483648)}
                                                    _ = v127
                                                    _ = v128
                                                    _ = v129
                                                }
                                                switch ("\n\t" + "\n\t") {
                                                case "\n\t":
                                                    fmt.Println(v125)
                                                    switch fi11 {
                                                    case -1:
                                                        v126 = v126
                                                    case 9223372036854775807:
                                                        if true {
                                                            switch fi11 {
                                                            case 32768:
                                                                v125 = genFn0(v126)
                                                            default:
                                                                var v130 uint64 = 443
                                                                _ = v130
                                                            }
                                                            switch 127 {
                                                            case -453:
                                                                v126 -= v126
                                                            case 4294967295:
                                                                v126 += v126
                                                            case 2147483647:
                                                                v10 = (v10 - 1)
                                                            default:
                                                                v131 := GenS9{F0: 127}
                                                                _ = v131
                                                            }
                                                            v126 += (v126 - (v126 / ((100) | 1)))
                                                        }
                                                    default:
                                                        var v132 int8 = -117
                                                        switch (fi11 & (fi11 + fi11)) {
                                                        case 32767:
                                                            if len(v123) > 0 {
                                                                v123[0] = fi88
                                                            }
                                                            v133 := GenS10{F0: 1}
                                                            if ((0.1 < -2.25) || (128 > 0)) {
                                                                v126 = (-(v126))
                                                            }
                                                            _ = v133
                                                        case 128:
                                                            v134 := GenS11{F0: v125, F1: 42}
                                                            _ = v134
                                                        case 995:
                                                            v125 += 4294967295
                                                        }
                                                        if (525 > 32767) {
                                                            fmt.Println(fi88, fi11, fi11)
                                                            if (true || ((((float64(1.5) - 0.1) * 3.14159) + (1.5 * (0.0 - 1e10))) <= 1e10)) {
                                                                if false {
                                                                    if len(v123) > 0 {
                                                                        v123[0] = fi11
                                                                    }
                                                                    if (228 != 875) {
                                                                        if ((!false) || false) {
                                                                            if false {
                                                                                var v135 uint16 = v10
                                                                                if false {
                                                                                    if true {
                                                                                        var v136 uint = 940
                                                                                        v123 = append(v123, fi11)
                                                                                        _ = v136
                                                                                    }
                                                                                    v135 = (v10 + (v135 >> ((626) & 63)))
                                                                                } else {
                                                                                    if true {
                                                                                        v123 = append(v123, (fi88 ^ fi11))
                                                                                        v137 := GenS6{F0: 95}
                                                                                        if (int16(v125) <= 1) {
                                                                                            v132 = 0
                                                                                            v132 += (^(v132))
                                                                                        }
                                                                                        _ = v137
                                                                                    }
                                                                                    var v138 float32 = float32(1e10)
                                                                                    fmt.Println(v125, v132)
                                                                                    _ = v138
                                                                                }
                                                                                v132 -= (v132 & v126)
                                                                                _ = v135
                                                                            }
                                                                            _ = len(v123)
                                                                            if len(v123) > 0 {
                                                                                v123[0] = fi11
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v139 float32 = (-2.25 - (3.14159 + -2.25))
                                                                        _ = v139
                                                                    }
                                                                } else {
                                                                    fmt.Println(v126)
                                                                    if len(v123) > 0 {
                                                                        v123[0] = fi11
                                                                    }
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    v132 -= v132
                                                                }
                                                            } else {
                                                                if (false || false) {
                                                                    if false {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v132 = v126
                                                                        }
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            var v140 int32 = 720
                                                                            _ = v140
                                                                        case "\n\t":
                                                                            if false {
                                                                                v126 += 100
                                                                            }
                                                                        case "hello":
                                                                            var v141 float32 = 3.14159
                                                                            if false {
                                                                                var v142 uint16 = 255
                                                                                v143 := GenS12{F0: (-(fi11)), F1: (true && ((-553 > -1) && (!false))), F2: 2147483647}
                                                                                if len(v123) >= 2 {
                                                                                    v123 = v123[1:len(v123)]
                                                                                }
                                                                                _ = v142
                                                                                _ = v143
                                                                            }
                                                                            switch 4294967295 {
                                                                            case 42:
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    v125 = 327
                                                                                case "foo bar":
                                                                                    v144 := GenS13{F0: 0, F1: ((("foo bar" + "a") + ("\n\t" + ("\n\t" + ""))) + ""), F2: false}
                                                                                    _ = v144
                                                                                }
                                                                            case -2147483648:
                                                                                fmt.Println(v141)
                                                                            case -363:
                                                                                v145 := GenS14{F0: (fi11 % (((fi11 << ((18446744073709551615) & 63))) | 1))}
                                                                                _ = v145
                                                                            default:
                                                                                fmt.Println(v10, v141)
                                                                            }
                                                                            _ = v141
                                                                        }
                                                                        if true {
                                                                            var v146 bool = (false && false)
                                                                            _ = v146
                                                                        } else {
                                                                            v147 := GenS7{F0: 180, F1: 42, F2: true}
                                                                            _ = v147
                                                                        }
                                                                    }
                                                                }
                                                                if (0 >= v10) {
                                                                    v148 := GenS14{F0: fi11}
                                                                    v132 = (-(v126))
                                                                    var v149 bool = (!true)
                                                                    v150 := GenS15{F0: v10, F1: 32767, F2: (v10 << ((0) & 63))}
                                                                    v151 := GenS8{F0: fi11, F1: 65535}
                                                                    _ = v148
                                                                    _ = v149
                                                                    _ = v150
                                                                    _ = v151
                                                                }
                                                                var v152 int32 = 486
                                                                _ = v152
                                                            }
                                                        }
                                                        _ = v132
                                                    }
                                                    if (!false) {
                                                        switch fi88 {
                                                        case 0:
                                                            if len(v123) > 0 {
                                                                v123[0] = fi11
                                                            }
                                                        case 1:
                                                            var v153 string = "foo bar"
                                                            _ = v153
                                                        case -1:
                                                            var v154 int64 = 945
                                                            _ = v154
                                                        }
                                                    } else {
                                                        fmt.Println(v125, v125, v125)
                                                        var v155 float32 = ((0.1 + 0.0) - 1.5)
                                                        if len(v123) > 0 {
                                                            v123[0] = -2147483648
                                                        }
                                                        _ = v155
                                                    }
                                                default:
                                                    _ = len(v123)
                                                }
                                                var v156 int = fi11
                                                _ = v156
                                            }
                                            if false {
                                                v157 := GenS2{F0: "hello", F1: 255}
                                                var v158 string = ("" + ("foo bar" + ("\n\t" + "hello")))
                                                _ = v157
                                                _ = v158
                                            }
                                            fmt.Println(v125, fi11, fi11)
                                        } else {
                                            var v159 uint16 = (v10 << ((906) & 63))
                                            v160 := GenS16{F0: 255}
                                            _ = v159
                                            _ = v160
                                        }
                                        var v161 uint8 = 42
                                        _ = v125
                                        _ = v126
                                        _ = v161
                                    } else {
                                        if (!(true && ("a" >= "a"))) {
                                            if false {
                                                fmt.Println(fi88)
                                            } else {
                                                if (float32(1.5) == float32((1.5 * ((-2.25 - (0.0 + 1e10)) / 1.0)))) {
                                                    if false {
                                                        fmt.Println(fi88, fi88, fi88)
                                                    } else {
                                                        var v162 uint64 = 18446744073709551615
                                                        var v163 int32 = 1
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                    var v164 int8 = -128
                                                    if false {
                                                        var v165 int32 = 42
                                                        _ = v165
                                                    }
                                                    _ = v164
                                                } else {
                                                    var v166 int16 = -1
                                                    _ = v166
                                                }
                                                switch "" {
                                                case "hello":
                                                    switch fi88 {
                                                    case -1:
                                                        if false {
                                                            switch "" {
                                                            case "\n\t":
                                                                var v167 uint32 = 4294967295
                                                                _ = v167
                                                            case "a":
                                                                v10 += v10
                                                            case "hello":
                                                                var v168 uint16 = v10
                                                                _ = v168
                                                            default:
                                                                if false {
                                                                    switch -2147483648 {
                                                                    case -622:
                                                                        var v169 int32 = 1
                                                                        var v170 uint64 = 4294967295
                                                                        _ = v169
                                                                        _ = v170
                                                                    case 127:
                                                                        fmt.Println(fi11, fi88, v10)
                                                                    }
                                                                } else {
                                                                    v10 = (^(v10))
                                                                }
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                v10 = (^(v10))
                                                            default:
                                                                v10 -= (v10 >> ((729) & 63))
                                                            }
                                                            var v171 uint16 = v10
                                                            _ = v171
                                                        } else {
                                                            v172 := GenS8{F0: fi11, F1: 65535}
                                                            fmt.Println(v10, fi88)
                                                            fmt.Println(fi11)
                                                            _ = v172
                                                        }
                                                    default:
                                                        var v173 int = 127
                                                        _ = v173
                                                    }
                                                }
                                                v10 += (v10 & v10)
                                            }
                                            var v174 uint32 = 2
                                            if true {
                                                if (v10 > 2) {
                                                    var v175 int32 = 0
                                                    v174 -= v174
                                                    var v176 string = ""
                                                    _ = v175
                                                    _ = v176
                                                }
                                                v177 := GenS17{F0: 0, F1: -954, F2: -9223372036854775808}
                                                var v178 uint16 = uint16(v174)
                                                switch fi88 {
                                                case 905:
                                                    v179 := GenS6{F0: 1}
                                                    if (2147483648 < genFn0(-128)) {
                                                        v123 = append(v123, -437)
                                                        var v180 uint64 = 259
                                                        v180 += (v180 - v180)
                                                        _ = v180
                                                    }
                                                    _ = v179
                                                default:
                                                    switch -887 {
                                                    case 127:
                                                        v123 = append(v123, fi11)
                                                    case -9223372036854775808:
                                                        switch -9223372036854775808 {
                                                        case 42:
                                                            v10 = 1
                                                        case 32768:
                                                            var v181 bool = (!(1 >= 128))
                                                            _ = v181
                                                        }
                                                    case 9223372036854775807:
                                                        v10 -= 871
                                                    default:
                                                        switch "" {
                                                        case "foo bar":
                                                            v10 = (v178 / (((v10 & 808)) | 1))
                                                        case "":
                                                            if true {
                                                                v174 += (v174 << ((820) & 63))
                                                            } else {
                                                                if true {
                                                                    var v182 int64 = 0
                                                                    var v183 uint = 0
                                                                    var v184 int64 = v182
                                                                    var v185 int16 = 0
                                                                    _ = v182
                                                                    _ = v183
                                                                    _ = v184
                                                                    _ = v185
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                var v186 uint64 = 0
                                                _ = v177
                                                _ = v178
                                                _ = v186
                                            }
                                            _ = v174
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(fi11, fi88)
                                }
                                _ = v123
                            }
                            v187 := []int{1}
                            _ = v187
                        }
                        _ = v110
                    }
                    var v188 uint32 = 2147483648
                    if (v10 > 997) {
                        var v189 uint = 2
                        _ = v189
                    } else {
                        v190 := GenS18{F0: fi11}
                        if false {
                            v188 += (^(v188))
                            if (65535 >= v188) {
                                if (v188 > v188) {
                                    v191 := []int{fi88, 32767}
                                    var v192 float64 = 0.1
                                    _ = v191
                                    _ = v192
                                }
                                v10 -= 406
                                v193 := GenS16{F0: 2}
                                _ = v193
                            } else {
                                v194 := []int{(fi88 | (^(fi88))), (fi11 + fi11)}
                                _ = v194
                            }
                            v195 := []int{2147483647, (fi88 / (((^(fi88))) | 1)), fi88, fi11}
                            v10 -= v10
                            var v196 float64 = (3.14159 * 0.0)
                            v195 = append(v195, fi11)
                            _ = v195
                            _ = v196
                        } else {
                            v197 := []int{(fi88 - 42), fi11, fi88}
                            v188 -= (v188 | v188)
                            var v198 uint = 2
                            _ = v197
                            _ = v198
                        }
                        var v199 int32 = -1
                        _ = v190
                        _ = v199
                    }
                    _ = fi88
                    _ = v188
                }
            }
            v200 := []int{}
            _ = fi11
            _ = v200
        }
    case "":
        var v201 uint64 = 249
        _ = v201
    case "hello":
        switch "a" {
        case "a":
            var v202 float32 = -2.25
            _ = v202
        }
    default:
        for fi203 := 0; fi203 < 12; fi203++ {
            fmt.Println(v10, fi203)
            var v204 uint = 1
            if (2 < uint8(v204)) {
                v204 = 42
                if true {
                    v205 := []int{}
                    v10 = (v10 + v10)
                    for fi206 := 0; fi206 < 14; fi206++ {
                        fmt.Println(v10, fi203)
                        _ = fi206
                    }
                    _ = v205
                } else {
                    var v207 float64 = (1e10 * 1e10)
                    if false {
                        fmt.Println(fi203, fi203)
                    } else {
                        v208 := []int{}
                        for fi209 := 0; fi209 < 5; fi209++ {
                            v207 += v207
                            var v210 int64 = 42
                            _ = fi209
                            _ = v210
                        }
                        _ = v208
                    }
                    _ = v207
                }
                v211 := GenS13{F0: -1, F1: ((("a" + "hello") + "foo bar") + ""), F2: true}
                _ = v211
            } else {
                var v212 uint64 = 1
                v213 := GenS8{F0: fi203, F1: 18446744073709551615}
                v10 = v10
                _ = v212
                _ = v213
            }
            _ = fi203
            _ = v204
        }
    }
    _ = v10
    return 25
}

func main() {
    var v1 int16 = -324
    var v2 float32 = ((-2.25 + 0.0) * 0.1)
    v3 := make(map[string]bool)
    var v4 string = "\n\t"
    v5 := []uint32{4294967295, 937}
    v2 = (((0.0 * (v2 + (v2 + (((float32((-2.25 / 1.0)) - v2) - float32(-2.25)) - v2)))) / (float32(-2.25) / v2)) + ((float32(float64((v2 / v2))) / v2) * v2))
    var v214 float32 = ((v2 + v2) - v2)
    if true {
        for fi215 := 0; fi215 < 10; fi215++ {
            if true {
                var v216 int8 = 43
                var v217 uint16 = 32767
                v1 += v1
                _ = v216
                _ = v217
            }
            v218 := GenS13{F0: 42, F1: ((v4 + (v4 + (v4 + "\n\t"))) + v4), F2: (618 <= 758)}
            if true {
                switch (((("foo bar" + ("\n\t" + (v4 + v4))) + v4) + (v4 + "\n\t")) + (v4 + "hello")) {
                case "hello":
                    fmt.Println(v214, v1, v1)
                case "\n\t":
                    var v219 int16 = 0
                    _ = v219
                }
                var v220 int32 = -2147483648
                _ = v220
            }
            _ = fi215
            _ = v218
        }
        v2 = float32(0.0)
        var v221 uint32 = 39
        _ = v221
    }
    v4 = "\n\t"
    v222 := GenS19{F0: 1, F1: 585, F2: 100}
    for fi223 := 0; fi223 < 19; fi223++ {
        switch 2147483647 {
        case 9223372036854775807:
            var v224 uint8 = 42
            _ = v224
        case 2147483647:
            var v225 float64 = (0.0 / 0.1)
            _ = v225
        }
        var v226 int32 = 574
        v227 := GenS20{F0: 127}
        _ = fi223
        _ = v226
        _ = v227
    }
    var v228 float32 = v214
    v1 -= (v1 * v1)
    if false {
        v229 := GenS6{F0: 198}
        var v230 uint32 = 65535
        _ = v229
        _ = v230
    } else {
        v3[v4] = ((18446744073709551615 <= 2) || (true || false))
        var v231 float32 = (((float32(0.0) - ((v214 * (v2 * v2)) / v214)) * float32(0.0)) - (-2.25 + v2))
        v4 = "hello"
        _ = v231
    }
    var v232 int = 583
    v214 += v214
    if (true && true) {
        var v233 int8 = 42
        _ = len(v5)
        v5 = append(v5, 4294967295)
        _ = v233
    }
    v214 = v214
    v232 += v232
    if true {
        v228 = (v214 * v228)
    }
    v3[v4] = false
    v234 := GenS4{F0: 32767, F1: -101, F2: int8(v1)}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v214:", v214)
    fmt.Println("v228:", v228)
    fmt.Println("v232:", v232)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v214
    _ = v222
    _ = v228
    _ = v232
    _ = v234
    _ = v3
    _ = v4
    _ = v5
}
