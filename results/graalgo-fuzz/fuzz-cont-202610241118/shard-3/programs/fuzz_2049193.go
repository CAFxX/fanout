package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int16 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 bool; F1 int64; F2 int32 }
type GenS9 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS10 struct { F0 uint64 }

func main() {
    var v1 uint32 = 4294967295
    v2 := []uint32{v1, (v1 >> ((42) & 63)), (v1 & 42)}
    v3 := make(map[int]bool)
    if false {
        v4 := GenS1{F0: 42}
        _ = v4
    }
    v1 += v1
    fmt.Println(v1)
    if true {
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := GenS1{F0: 225}
            var v7 string = ("" + "")
            for fi8 := 0; fi8 < 5; fi8++ {
                v7 = "\n\t"
                var v9 int8 = -5
                _ = fi8
                _ = v9
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        fmt.Println(v1)
    }
    v1 += (v1 & (^(v1)))
    fmt.Println(v1)
    if false {
        for fi10 := 0; fi10 < 12; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                switch "hello" {
                case "a":
                    if true {
                        if false {
                            fmt.Println(v1, fi10, v1)
                            var v12 int64 = 952
                            fmt.Println(fi11)
                            var v13 int16 = 120
                            var v14 bool = true
                            _ = v12
                            _ = v13
                            _ = v14
                        } else {
                            v15 := GenS2{F0: 2147483647, F1: fi10}
                            fmt.Println(fi11, fi10)
                            v2 = append(v2, v1)
                            if true {
                                v3[fi11] = (-1 <= 745)
                            } else {
                                v16 := GenS1{F0: 42}
                                if true {
                                    v1 = 2147483648
                                    var v17 int8 = 37
                                    _ = v17
                                }
                                delete(v3, fi11)
                                _ = v16
                            }
                            v18 := GenS1{F0: 127}
                            v1 = (v1 & v1)
                            fmt.Println(v1, fi10)
                            _ = v15
                            _ = v18
                        }
                    } else {
                        var v19 uint = 58
                        v3[fi10] = false
                        _ = v19
                    }
                    v20 := GenS1{F0: 1}
                    fmt.Println(fi10, fi10)
                    _ = v20
                case "hello":
                    var v21 uint64 = 4294967295
                    switch fi11 {
                    case 441:
                        var v22 int8 = 42
                        _ = v22
                    case 32767:
                        switch (fi10 | 935) {
                        case 32767:
                            fmt.Println(v21, fi11)
                        default:
                            if ((!false) || (v21 == v21)) {
                                delete(v3, (^(fi11)))
                                var v23 float64 = -2.25
                                v24 := GenS2{F0: (v1 - (v1 % ((v1) | 1))), F1: fi10}
                                var v25 float64 = (1e10 + ((3.14159 * 1.5) + v23))
                                var v26 uint8 = 20
                                var v27 float64 = v25
                                if (!(42 <= 114)) {
                                    var v28 uint = 233
                                    v21 = v21
                                    if (!false) {
                                        v27 = float64(1.5)
                                        if false {
                                            fmt.Println(v27)
                                            var v29 int8 = -128
                                            _ = v29
                                        }
                                    } else {
                                        var v30 int = (fi11 * 0)
                                        var v31 int16 = 32767
                                        _ = v30
                                        _ = v31
                                    }
                                    if true {
                                        v32 := GenS2{F0: (v1 - 2147483648), F1: fi11}
                                        switch 352 {
                                        case 42:
                                            v33 := GenS3{F0: -1, F1: -1}
                                            _ = v33
                                        }
                                        if (true || (fi11 == (fi11 - fi10))) {
                                            if (v28 <= v28) {
                                                if ((((v1 % ((v1) | 1)) == (v1 << (((^(v28))) & 63))) || ((fi11 == fi10) && (v21 <= v21))) && (-128 <= -39)) {
                                                    v25 = v27
                                                }
                                                var v34 int8 = 100
                                                _ = v34
                                            }
                                            var v35 uint8 = v26
                                            fmt.Println(v23, fi11)
                                            switch (^(fi11)) {
                                            case 2147483647:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    fmt.Println(v21, v35)
                                                case "hello":
                                                    fmt.Println(fi10)
                                                case "a":
                                                    var v36 float64 = v25
                                                    _ = v36
                                                default:
                                                    v3[fi11] = (!false)
                                                    var v37 uint64 = (v21 % ((v21) | 1))
                                                    _ = v37
                                                }
                                            case 32768:
                                                _ = len(v2)
                                            case 42:
                                                fmt.Println(fi10)
                                            default:
                                                var v38 string = ("hello" + ("" + ""))
                                                _ = v38
                                            }
                                            v39 := GenS4{F0: (v1 + (v1 >> (((v28 ^ 932)) & 63))), F1: 1, F2: 32767}
                                            v25 = 3.14159
                                            if ("hello" == "a") {
                                                fmt.Println(fi10, v21, fi11)
                                                switch fi10 {
                                                case 32768:
                                                    var v40 uint16 = 32767
                                                    _ = v40
                                                case -872:
                                                    if true {
                                                        v21 = 4294967295
                                                        var v41 uint16 = 255
                                                        _ = v41
                                                    }
                                                case 331:
                                                    if false {
                                                        fmt.Println(fi10)
                                                        v42 := GenS1{F0: (v35 + (v26 / (((v35 | (v35 - v26))) | 1)))}
                                                        _ = v42
                                                    }
                                                default:
                                                    _ = len(v3)
                                                }
                                            } else {
                                                v28 = 1
                                                if len(v2) > 0 {
                                                    v2[0] = v1
                                                }
                                            }
                                            if false {
                                                v43 := GenS1{F0: v26}
                                                switch "hello" {
                                                case "\n\t":
                                                    if false {
                                                        v3[(fi10 * (fi11 * fi10))] = false
                                                    }
                                                case "":
                                                    switch "hello" {
                                                    case "a":
                                                        v44 := GenS5{F0: 0, F1: 2147483648}
                                                        _ = v44
                                                    }
                                                }
                                                v2 = append(v2, (v1 / (((v1 % (((v1 << ((v28) & 63))) | 1))) | 1)))
                                                if (v27 > (v27 * v23)) {
                                                    var v45 int64 = 9223372036854775807
                                                    if len(v2) > 0 {
                                                        v2[0] = (v1 << ((v28) & 63))
                                                    }
                                                    _ = v45
                                                } else {
                                                    v25 = (float64((-2.25 / -2.25)) / v23)
                                                    if (1 != 0) {
                                                        switch "hello" {
                                                        case "a":
                                                            var v46 int64 = -955
                                                            _ = v46
                                                        case "foo bar":
                                                            v35 += v35
                                                        }
                                                        var v47 uint8 = (v26 % ((122) | 1))
                                                        v27 = v25
                                                        _ = v47
                                                    }
                                                }
                                                _ = v43
                                            }
                                            _ = v35
                                            _ = v39
                                        } else {
                                            var v48 bool = false
                                            if false {
                                                var v49 uint64 = (v21 & (v21 | v21))
                                                switch ("" + ("foo bar" + ("" + (("foo bar" + ("\n\t" + "foo bar")) + (((("foo bar" + "") + ("" + "")) + "\n\t") + "\n\t"))))) {
                                                case "\n\t":
                                                    v50 := GenS1{F0: uint8(v28)}
                                                    _ = v50
                                                case "a":
                                                    v2 = append(v2, (v1 << (((v28 - v28)) & 63)))
                                                }
                                                v51 := GenS4{F0: (v1 % (((^(v1))) | 1)), F1: -790, F2: 559}
                                                switch ("" + "hello") {
                                                case "foo bar":
                                                    var v52 uint32 = (v1 | v1)
                                                    _ = v52
                                                default:
                                                    v53 := GenS4{F0: v1, F1: 876, F2: -1}
                                                    _ = v53
                                                }
                                                _ = v49
                                                _ = v51
                                            } else {
                                                v54 := GenS1{F0: 128}
                                                fmt.Println(fi11, v27)
                                                var v55 int8 = -107
                                                _ = v54
                                                _ = v55
                                            }
                                            _ = v48
                                        }
                                        _ = v32
                                    }
                                    _ = v28
                                } else {
                                    var v56 int32 = -1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if true {
                                        v26 = 128
                                        v1 += v1
                                    }
                                    v56 = v56
                                    _ = v56
                                }
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        }
                    default:
                        switch 42 {
                        case 42:
                            if false {
                                v1 -= (^(v1))
                            } else {
                                v1 = 65535
                                if true {
                                    v21 = (v21 | (v21 << ((608) & 63)))
                                    delete(v3, -1)
                                    switch "hello" {
                                    case "hello":
                                        v21 += 1
                                    case "foo bar":
                                        v21 = v21
                                    case "\n\t":
                                        fmt.Println(v21, fi11, v21)
                                    }
                                }
                                if (1 != 169) {
                                    var v57 uint64 = v21
                                    _ = v57
                                } else {
                                    var v58 uint64 = (v21 >> ((255) & 63))
                                    if (1000 <= -32768) {
                                        v1 -= (v1 << ((2) & 63))
                                        if true {
                                            v59 := GenS2{F0: 1, F1: fi11}
                                            _ = v59
                                        }
                                        delete(v3, (-(fi11)))
                                    }
                                    v60 := GenS1{F0: 1}
                                    _ = v58
                                    _ = v60
                                }
                            }
                        case 128:
                            var v61 float64 = -2.25
                            _ = v61
                        case -354:
                            if (!true) {
                                v2 = append(v2, v1)
                                v62 := GenS5{F0: -1, F1: 2147483647}
                                _ = v62
                            }
                        }
                    }
                    _ = v21
                }
                _ = fi11
            }
            _ = len(v3)
            for fi63 := 0; fi63 < 6; fi63++ {
                var v64 uint8 = 2
                _ = fi63
                _ = v64
            }
            _ = fi10
        }
    } else {
        fmt.Println(v1)
        if false {
            var v65 float64 = 1.5
            v2 = append(v2, (v1 + (v1 | 65535)))
            if false {
                var v66 float64 = v65
                _ = v66
            } else {
                v67 := v3[-695]
                _ = v67
            }
            _ = v65
        } else {
            for fi68 := 0; fi68 < 14; fi68++ {
                v1 = v1
                for fi69 := 0; fi69 < 9; fi69++ {
                    v70 := GenS4{F0: (v1 + (v1 % ((v1) | 1))), F1: 885, F2: 0}
                    _ = fi69
                    _ = v70
                }
                var v71 int32 = -922
                _ = fi68
                _ = v71
            }
            for fi72 := 0; fi72 < 4; fi72++ {
                var v73 uint = 990
                v73 = (v73 >> ((v73) & 63))
                for fi74 := 0; fi74 < 19; fi74++ {
                    switch fi72 {
                    case 130:
                        var v75 uint8 = 128
                        _ = v75
                    case 128:
                        v1 += 0
                    case 0:
                        _ = len(v3)
                    }
                    var v76 int16 = -32768
                    _ = fi74
                    _ = v76
                }
                _ = fi72
                _ = v73
            }
            _ = len(v2)
        }
        switch "" {
        case "foo bar":
            var v77 uint16 = uint16(v1)
            _ = v77
        case "\n\t":
            fmt.Println(v1)
        case "hello":
            if false {
                if (float32(0.1) != 1e10) {
                    for fi78 := 0; fi78 < 12; fi78++ {
                        switch 32768 {
                        case -9223372036854775808:
                            v79 := GenS3{F0: 0, F1: 0}
                            _ = v79
                        default:
                            _ = len(v3)
                        }
                        _ = fi78
                    }
                    v1 = (v1 - 191)
                    fmt.Println(v1)
                } else {
                    if true {
                        v80 := GenS2{F0: (v1 << ((18446744073709551615) & 63)), F1: 153}
                        var v81 uint8 = 127
                        _ = v80
                        _ = v81
                    }
                    for fi82 := 0; fi82 < 17; fi82++ {
                        switch "a" {
                        case "a":
                            v3[(fi82 * 32767)] = (128 == 1)
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                switch (fi82 + (fi82 + fi82)) {
                                case 9223372036854775807:
                                    for fi83 := 0; fi83 < 19; fi83++ {
                                        v1 = v1
                                        _ = fi83
                                    }
                                default:
                                    var v84 uint = 654
                                    _ = v84
                                }
                            }
                        case "a":
                            v85 := GenS6{F0: ("a" + "foo bar")}
                            _ = v85
                        case "hello":
                            v1 = v1
                        default:
                            switch ("foo bar" + "a") {
                            case "a":
                                v1 += v1
                            case "":
                                for fi86 := 0; fi86 < 5; fi86++ {
                                    v2 = append(v2, (v1 ^ (v1 - v1)))
                                    _ = fi86
                                }
                            case "\n\t":
                                fmt.Println(v1, v1)
                            }
                        }
                        _ = fi82
                    }
                    v1 += 2
                }
                fmt.Println(v1)
                var v87 int64 = -9223372036854775808
                _ = v87
            } else {
                v88 := GenS4{F0: (v1 & v1), F1: -343, F2: 32767}
                if (808 <= 65535) {
                    var v89 float64 = -2.25
                    _ = v89
                }
                _ = v88
            }
        default:
            fmt.Println(v1)
        }
    }
    v90 := GenS6{F0: "hello"}
    var v91 uint64 = 1
    var v92 uint16 = 42
    fmt.Println(v92, v92, v1)
    v92 -= (^(v92))
    for fi93 := 0; fi93 < 20; fi93++ {
        if (true && ((v92 | 1) <= (^(v92)))) {
            for fi94 := 0; fi94 < 2; fi94++ {
                v95 := GenS7{F0: 1, F1: 42}
                var v96 int32 = -1
                if true {
                    v96 -= 2147483647
                    switch int(v91) {
                    case 32767:
                        v97 := GenS8{F0: (fi93 <= (fi94 << ((198) & 63))), F1: 186, F2: (v96 % ((878) | 1))}
                        switch "\n\t" {
                        case "hello":
                            var v98 string = ("\n\t" + "")
                            _ = v98
                        case "foo bar":
                            var v99 int16 = 1000
                            _ = v99
                        case "":
                            v91 = (v91 | v91)
                        default:
                            if (!(1 <= -49)) {
                                switch "foo bar" {
                                case "":
                                    v91 -= v91
                                }
                                if (!(-118 <= 0)) {
                                    var v100 uint64 = 424
                                    if (!("foo bar" >= "a")) {
                                        if false {
                                            v100 = v91
                                        } else {
                                            var v101 uint64 = v100
                                            fmt.Println(fi93)
                                            switch int(v92) {
                                            case -2147483648:
                                                switch -2 {
                                                case 42:
                                                    switch (fi94 ^ int(v96)) {
                                                    case 9223372036854775807:
                                                        v91 = v100
                                                    case 46:
                                                        v96 = (v96 % ((510) | 1))
                                                    case 755:
                                                        v102 := GenS3{F0: int8(v101), F1: 450}
                                                        _ = v102
                                                    default:
                                                        var v103 bool = (!(255 < uint(v91)))
                                                        _ = v103
                                                    }
                                                default:
                                                    fmt.Println(v91)
                                                }
                                            case 0:
                                                fmt.Println(v96, fi94, v92)
                                            }
                                            _ = v101
                                        }
                                    } else {
                                        v100 -= v91
                                        v100 = (v100 << ((25) & 63))
                                        var v104 uint32 = v1
                                        var v105 uint16 = v92
                                        var v106 float32 = 0.0
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                    }
                                    v3[int(v96)] = true
                                    _ = v100
                                } else {
                                    switch ("hello" + "") {
                                    case "hello":
                                        v1 = v1
                                    case "\n\t":
                                        fmt.Println(fi93, v1)
                                    default:
                                        switch (("a" + "a") + ("\n\t" + (("a" + ("foo bar" + "")) + "hello"))) {
                                        case "a":
                                            if ("hello" <= (("\n\t" + ((("hello" + "a") + "") + ("foo bar" + "\n\t"))) + ("hello" + "hello"))) {
                                                v2 = append(v2, 2)
                                                v107 := GenS1{F0: 185}
                                                _ = v107
                                            } else {
                                                v96 -= (v96 - (v96 / ((42) | 1)))
                                                var v108 uint = 2
                                                v96 -= v96
                                                _ = v108
                                            }
                                        }
                                    }
                                    v1 -= v1
                                    switch 949 {
                                    case -918:
                                        var v109 uint8 = 232
                                        _ = v109
                                    }
                                    if true {
                                        var v110 float64 = 0.1
                                        _ = v110
                                    } else {
                                        var v111 float32 = 1e10
                                        _ = len(v2)
                                        _ = v111
                                    }
                                    fmt.Println(v1, v1, fi93)
                                }
                                var v112 float32 = (3.14159 * -2.25)
                                _ = v112
                            } else {
                                v96 = v96
                                v113 := GenS4{F0: v1, F1: 0, F2: -1}
                                var v114 float64 = 1e10
                                _ = v113
                                _ = v114
                            }
                        }
                        _ = v97
                    case 0:
                        v1 -= 4294967295
                    default:
                        var v115 int32 = -323
                        _ = v115
                    }
                    v116 := GenS9{F0: 244, F1: 42, F2: 170}
                    v92 += (^(v92))
                    var v117 float32 = 1.5
                    v96 = -2147483648
                    _ = v116
                    _ = v117
                } else {
                    if ((v96 / ((v96) | 1)) < (v96 << ((65535) & 63))) {
                        v92 = 32767
                        var v118 uint64 = v91
                        v119 := GenS10{F0: 1}
                        _ = v118
                        _ = v119
                    } else {
                        v92 = (v92 & 0)
                        v2 = append(v2, v1)
                    }
                }
                _ = fi94
                _ = v95
                _ = v96
            }
            for fi120 := 0; fi120 < 6; fi120++ {
                fmt.Println(v1)
                if true {
                    switch (fi120 << ((4294967295) & 63)) {
                    case -126:
                        fmt.Println(fi120, v1)
                    case -1:
                        fmt.Println(v92, v1)
                    case -383:
                        var v121 int16 = 321
                        _ = v121
                    }
                } else {
                    v122 := v3[fi93]
                    var v123 uint32 = 2147483647
                    _ = v122
                    _ = v123
                }
                var v124 int = 238
                _ = fi120
                _ = v124
            }
        } else {
            for fi125 := 0; fi125 < 8; fi125++ {
                fmt.Println(v91, fi125)
                if (631 >= 1) {
                    if false {
                        var v126 uint16 = v92
                        fmt.Println(v1)
                        _ = v126
                    }
                    v91 = (v91 % (((v91 + 42)) | 1))
                    var v127 float64 = (1e10 + (float64(float32((float64(((-2.25 * 0.0) - (1e10 - -2.25))) - 1.5))) + (((1e10 * 0.0) - 0.0) * 1.5)))
                    v92 = (v92 % ((255) | 1))
                    _ = v127
                }
                switch (fi125 / ((fi125) | 1)) {
                case -977:
                    var v128 float32 = ((float32(0.1) / (1.5 + ((-2.25 * float32(((-2.25 / 3.14159) - 0.1))) - 3.14159))) + (1e10 - 0.1))
                    _ = v128
                case -1:
                    switch (fi125 + fi125) {
                    case 168:
                        var v129 int16 = 209
                        _ = v129
                    }
                default:
                    v91 -= (v91 + 617)
                }
                var v130 float64 = 0.0
                if len(v2) > 0 {
                    v2[0] = (^(v1))
                }
                _ = fi125
                _ = v130
            }
            var v131 string = (("a" + (("hello" + ("foo bar" + "hello")) + ("" + "hello"))) + "a")
            _ = v131
        }
        if (float64(0.0) == 1.5) {
            fmt.Println(fi93, v1, v1)
            v92 -= v92
            if false {
                v92 += (v92 % ((v92) | 1))
                var v132 uint64 = v91
                fmt.Println(fi93, v1, v91)
                _ = v132
            } else {
                var v133 uint = 42
                v133 -= 395
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v133
            }
        }
        _ = fi93
    }
    if (true && false) {
        for fi134 := 0; fi134 < 15; fi134++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi134
        }
    }
    v135 := GenS6{F0: ("" + "foo bar")}
    var v136 int = 544
    fmt.Println("v1:", v1)
    fmt.Println("v136:", v136)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v90
    _ = v91
    _ = v92
}
