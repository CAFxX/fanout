package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int64; F1 int32; F2 int64 }
type GenS8 struct { F0 int; F1 string; F2 uint64 }
type GenS9 struct { F0 int; F1 uint64; F2 uint32 }
type GenS10 struct { F0 uint; F1 int64; F2 string }
type GenS11 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS12 struct { F0 int64; F1 bool }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 bool; F1 int32 }
func genFn0(p0 uint, p1 int, p2 uint8) uint16 {
    var v4 uint8 = 1
    var v5 uint8 = (v4 / ((84) | 1))
    v6 := GenS1{F0: (("foo bar" + "a") + "\n\t"), F1: 127}
    v7 := GenS2{F0: -1, F1: (^(p1)), F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 uint16) int32 {
    if (3.14159 < (((1e10 * ((1.5 / 0.1) - (float64(-2.25) * (1e10 * (0.0 + (3.14159 * -2.25)))))) * 0.0) * 0.1)) {
        fmt.Println(p0)
    }
    switch 893 {
    case 42:
        p0 = (^(p0))
    case 128:
        if true {
            fmt.Println(p0)
        }
    case -265:
        if (800 != 42) {
            if true {
                v8 := []int{}
                _ = v8
            } else {
                for fi9 := 0; fi9 < 6; fi9++ {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(fi10)
                        var v11 bool = false
                        _ = fi10
                        _ = v11
                    }
                    var v12 uint16 = (^(p0))
                    if false {
                        v12 -= genFn0(355, fi9, 128)
                        var v13 uint = 523
                        p0 -= (p0 & 32767)
                        _ = v13
                    } else {
                        var v14 int16 = 1
                        for fi15 := 0; fi15 < 10; fi15++ {
                            v16 := []int{fi9, fi15}
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = fi9
                    _ = v12
                }
                p0 = p0
                p0 += (p0 << ((4294967295) & 63))
            }
        } else {
            v17 := []int{}
            _ = v17
        }
    default:
        if false {
            fmt.Println(p0)
            switch 127 {
            case 29:
                switch "\n\t" {
                case "\n\t":
                    if (-28 == -128) {
                        v18 := GenS2{F0: -379, F1: 933, F2: 348}
                        for fi19 := 0; fi19 < 15; fi19++ {
                            var v20 string = ("" + "hello")
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 uint8 = 42
                            v23 := []int{(fi21 << ((386) & 63))}
                            var v24 uint32 = 65535
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        var v25 bool = true
                        v26 := GenS2{F0: -1, F1: -2147483648, F2: 1}
                        _ = v25
                        _ = v26
                    }
                case "hello":
                    switch "" {
                    case "a":
                        switch 128 {
                        case -146:
                            for fi27 := 0; fi27 < 12; fi27++ {
                                switch -1 {
                                case 128:
                                    switch fi27 {
                                    case -311:
                                        var v28 int8 = 61
                                        _ = v28
                                    case 127:
                                        fmt.Println(p0)
                                    case 4294967295:
                                        v29 := []int{}
                                        _ = v29
                                    }
                                default:
                                    p0 += p0
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    v30 := GenS3{F0: false}
                                    _ = v30
                                default:
                                    var v31 bool = true
                                    _ = v31
                                }
                                _ = fi27
                            }
                        case 127:
                            p0 -= p0
                        }
                    case "hello":
                        if false {
                            for fi32 := 0; fi32 < 14; fi32++ {
                                switch ("foo bar" + ("foo bar" + "hello")) {
                                case "\n\t":
                                    v33 := []int{(fi32 * (fi32 * (fi32 * fi32))), (fi32 & fi32), fi32, fi32}
                                    _ = v33
                                case "hello":
                                    p0 = (p0 | p0)
                                case "foo bar":
                                    v34 := GenS3{F0: (float64(-2.25) != 0.0)}
                                    _ = v34
                                }
                                p0 = 65535
                                _ = fi32
                            }
                        }
                    case "":
                        v35 := []int{-464, 9223372036854775807, -9223372036854775808}
                        _ = v35
                    }
                case "foo bar":
                    var v36 int = 2147483647
                    _ = v36
                }
            }
        } else {
            v37 := []int{2147483647}
            fmt.Println(p0)
            if false {
                p0 -= (p0 % ((487) | 1))
                if len(v37) > 0 {
                    v37[0] = -2147483648
                }
            }
            _ = v37
        }
    }
    switch "a" {
    case "foo bar":
        p0 -= (p0 / ((65535) | 1))
    case "a":
        v38 := GenS4{F0: 1}
        _ = v38
    case "hello":
        v39 := []int{-1, 32768, 9223372036854775807, -9223372036854775808}
        _ = v39
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            if ((462 >= 4294967295) || true) {
                var v40 string = "foo bar"
                _ = v40
            } else {
                fmt.Println(p0)
            }
        case "\n\t":
            switch 2147483647 {
            case 32767:
                switch "hello" {
                case "a":
                    p0 += 2
                }
            case -9223372036854775808:
                v41 := GenS1{F0: (("\n\t" + "a") + "\n\t"), F1: 100}
                _ = v41
            case 2147483647:
                p0 = p0
            default:
                for fi42 := 0; fi42 < 20; fi42++ {
                    p0 += 741
                    var v43 string = (("a" + "hello") + "\n\t")
                    var v44 float64 = (((0.0 - float64(float32(3.14159))) / ((0.1 * (0.1 - -2.25)) - 1.5)) / 1.0)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
            }
        case "hello":
            switch 0 {
            case -568:
                v45 := GenS4{F0: -668}
                _ = v45
            }
        }
        switch "a" {
        case "\n\t":
            var v46 float64 = float64(-2.25)
            _ = v46
        }
    } else {
        p0 = 763
        for fi47 := 0; fi47 < 1; fi47++ {
            var v48 uint16 = 462
            for fi49 := 0; fi49 < 4; fi49++ {
                if true {
                    var v50 int = fi49
                    p0 = (^(p0))
                    _ = v50
                }
                _ = fi49
            }
            v51 := []int{}
            _ = fi47
            _ = v48
            _ = v51
        }
    }
    if true {
        v52 := []int{-9223372036854775808}
        if true {
            if (42 >= 465) {
                var v53 int64 = -1
                for fi54 := 0; fi54 < 6; fi54++ {
                    var v55 int16 = -32768
                    _ = fi54
                    _ = v55
                }
                _ = v53
            } else {
                var v56 bool = ((true && true) && (!false))
                _ = v56
            }
            switch 144 {
            case 0:
                switch "hello" {
                case "a":
                    if true {
                        v57 := GenS5{F0: 32767}
                        p0 -= 293
                        fmt.Println(p0)
                        _ = v57
                    } else {
                        if len(v52) >= 2 {
                            v52 = v52[1:len(v52)]
                        }
                    }
                default:
                    if false {
                        v52 = append(v52, 190)
                        var v58 uint32 = uint32(p0)
                        _ = v58
                    }
                }
            default:
                if true {
                    v59 := GenS6{F0: ("foo bar" + ("\n\t" + ("a" + ("hello" + "")))), F1: 446}
                    if (true || (float32((-2.25 + ((1.5 - float64((-2.25 / 0.1))) * float64(0.1)))) != (((3.14159 - 1.5) / 0.1) + 3.14159))) {
                        var v60 int8 = 10
                        var v61 uint64 = 42
                        fmt.Println(v61, v60)
                        _ = v60
                        _ = v61
                    } else {
                        var v62 bool = true
                        _ = v62
                    }
                    _ = v59
                }
            }
            p0 = 65535
        } else {
            var v63 float32 = (3.14159 / (1e10 * (0.0 + float32(1.5))))
            for fi64 := 0; fi64 < 11; fi64++ {
                var v65 bool = ("a" <= "a")
                fmt.Println(fi64, fi64, v65)
                var v66 string = ""
                _ = fi64
                _ = v65
                _ = v66
            }
            _ = v63
        }
        _ = v52
    } else {
        var v67 int8 = -128
        v68 := []int{0, -9223372036854775808, -322, 831}
        for fi69 := 0; fi69 < 10; fi69++ {
            if (genFn0(178, fi69, 42) < genFn0(255, (fi69 - fi69), 128)) {
                for fi70 := 0; fi70 < 17; fi70++ {
                    p0 = 230
                    fmt.Println(fi69, fi69, fi70)
                    if true {
                        if true {
                            p0 = (p0 % (((p0 & genFn0(1, fi70, 255))) | 1))
                            v67 -= v67
                            if (0 == 0) {
                                if (0.1 <= 1.5) {
                                    var v71 int8 = -128
                                    v71 += v71
                                    _ = v71
                                } else {
                                    var v72 uint64 = 4294967295
                                    switch "foo bar" {
                                    case "":
                                        if false {
                                            fmt.Println(p0, v72)
                                        } else {
                                            v72 -= 1
                                            v73 := GenS5{F0: 0}
                                            _ = v73
                                        }
                                    case "a":
                                        if (-139 > 1) {
                                            if false {
                                                if true {
                                                    if len(v68) > 0 {
                                                        v68[0] = (fi70 - 32768)
                                                    }
                                                    var v74 int64 = -9223372036854775808
                                                    _ = v74
                                                } else {
                                                    var v75 string = "foo bar"
                                                    var v76 float64 = (3.14159 - 3.14159)
                                                    if len(v68) > 0 {
                                                        v68[0] = int(v67)
                                                    }
                                                    _ = v75
                                                    _ = v76
                                                }
                                            }
                                        } else {
                                            v67 -= -91
                                            v67 -= v67
                                        }
                                    case "foo bar":
                                        v77 := GenS2{F0: -833, F1: (fi69 | fi70), F2: 423}
                                        _ = v77
                                    }
                                    _ = v72
                                }
                            }
                        }
                        if (!false) {
                            var v78 string = ("a" + ("a" + ""))
                            if (9223372036854775807 > 1) {
                                v79 := GenS3{F0: (1.5 >= (0.1 + 1.5))}
                                if false {
                                    var v80 string = v78
                                    _ = v80
                                } else {
                                    fmt.Println(p0, fi69, p0)
                                    var v81 uint8 = 1
                                    switch (fi70 + fi69) {
                                    case 42:
                                        v78 = "\n\t"
                                    case 128:
                                        if false {
                                            v78 = v78
                                            switch fi69 {
                                            case -9223372036854775808:
                                                var v82 int32 = 0
                                                _ = v82
                                            case 836:
                                                var v83 uint = 65535
                                                _ = v83
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    _ = len(v68)
                                                case "foo bar":
                                                    var v84 int64 = -9223372036854775808
                                                    _ = v84
                                                case "":
                                                    var v85 float64 = 0.0
                                                    _ = v85
                                                default:
                                                    if (uint(p0) > 1) {
                                                        switch (-(fi69)) {
                                                        case -427:
                                                            if ((p0 % (((p0 ^ 65535)) | 1)) >= p0) {
                                                                switch fi69 {
                                                                case 2147483647:
                                                                    v78 = (v78 + v78)
                                                                case -2147483648:
                                                                    v67 = (v67 ^ (-(v67)))
                                                                }
                                                                var v86 int8 = 0
                                                                v67 -= 1
                                                                _ = v86
                                                            }
                                                        case 9223372036854775807:
                                                            p0 = (^(p0))
                                                        case -305:
                                                            var v87 uint32 = 737
                                                            _ = v87
                                                        default:
                                                            fmt.Println(v67, v67)
                                                        }
                                                        v88 := GenS7{F0: 0, F1: -2147483648, F2: 969}
                                                        _ = v88
                                                    }
                                                }
                                            }
                                        }
                                    case 164:
                                        fmt.Println(p0, fi70)
                                    }
                                    _ = v81
                                }
                                _ = v79
                            }
                            var v89 uint64 = 4294967295
                            var v90 float32 = 3.14159
                            _ = v78
                            _ = v89
                            _ = v90
                        }
                        if (false && false) {
                            switch fi69 {
                            case 0:
                                v91 := GenS8{F0: (fi70 + -9223372036854775808), F1: "\n\t", F2: 196}
                                _ = v91
                            case 2147483647:
                                var v92 int16 = -1
                                _ = v92
                            case 902:
                                v93 := GenS8{F0: fi70, F1: "a", F2: 0}
                                _ = v93
                            default:
                                if (42 >= 9223372036854775807) {
                                    fmt.Println(fi70, fi69)
                                }
                            }
                        } else {
                            fmt.Println(p0, fi69)
                            fmt.Println(fi69, v67, p0)
                        }
                    } else {
                        fmt.Println(fi69)
                        var v94 string = "a"
                        var v95 int64 = 42
                        _ = v94
                        _ = v95
                    }
                    if (((false || true) && true) || ((0.0 * (1e10 - -2.25)) >= (((1.5 - 3.14159) * 1e10) * (1e10 * (((0.0 / 0.1) * 3.14159) - -2.25))))) {
                        fmt.Println(v67)
                    }
                    _ = fi70
                }
                v67 = v67
                var v96 int16 = -977
                _ = v96
            }
            _ = fi69
        }
        _ = v67
        _ = v68
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]bool)
    var v3 int = -1
    if false {
        v97 := GenS5{F0: -32768}
        if ((!(!true)) || (!(2 != 931))) {
            fmt.Println(v1)
            var v98 float64 = (-2.25 + 0.1)
            if (!true) {
                for fi99 := 0; fi99 < 19; fi99++ {
                    delete(v2, (fi99 >> ((42) & 63)))
                    _ = fi99
                }
            } else {
                v1 = (v1 + genFn0(4294967295, (v3 | v3), 0))
                var v100 int32 = -188
                _ = v100
            }
            _ = v98
        } else {
            if (!(v1 == genFn0(4294967295, -927, 108))) {
                v101 := []int{(v3 & v3), 0, v3}
                v3 -= v3
                _ = v101
            } else {
                v3 -= (-(v3))
                switch (("hello" + "") + "") {
                case "\n\t":
                    var v102 int16 = 1
                    _ = v102
                }
            }
        }
        _ = v97
    }
    switch (v3 | v3) {
    case -728:
        if false {
            v103 := []int{v3, (v3 >> ((4294967295) & 63))}
            if (65535 >= 2147483647) {
                var v104 uint64 = 1
                var v105 uint8 = 0
                v106 := GenS9{F0: v3, F1: 2, F2: 2147483647}
                _ = v104
                _ = v105
                _ = v106
            }
            var v107 uint8 = 220
            _ = v103
            _ = v107
        } else {
            v108 := []int{(v3 % (((v3 - (v3 - 128))) | 1)), v3}
            _ = v108
        }
    case 42:
        v1 -= 1
    }
    v3 = v3
    v109 := GenS10{F0: 685, F1: 0, F2: "\n\t"}
    if true {
        v110 := GenS4{F0: 1000}
        _ = v110
    }
    for fi111 := 0; fi111 < 20; fi111++ {
        switch "\n\t" {
        case "":
            v1 = (v1 << ((65535) & 63))
        default:
            switch "foo bar" {
            case "hello":
                if false {
                    v3 = (v3 ^ fi111)
                }
            case "a":
                switch (v3 % (((v3 >> ((466) & 63))) | 1)) {
                case 128:
                    if false {
                        v3 += fi111
                        var v112 uint = 771
                        _ = v112
                    }
                case -106:
                    v3 = fi111
                }
            default:
                v3 += (-(fi111))
            }
        }
        _ = fi111
    }
    var v113 uint64 = 333
    if ((4294967295 >= 392) || true) {
        fmt.Println(v1)
    }
    switch (^(v3)) {
    case -489:
        switch (v3 | v3) {
        case -2147483648:
            v114 := GenS11{F0: 4294967295, F1: 1, F2: genFn0(4294967295, (v3 % ((v3) | 1)), 1)}
            _ = v114
        case 1:
            fmt.Println(v3)
        case 2147483647:
            var v115 float64 = 1.5
            _ = v115
        default:
            for fi116 := 0; fi116 < 8; fi116++ {
                var v117 uint32 = 42
                v113 = (v113 - v113)
                for fi118 := 0; fi118 < 11; fi118++ {
                    var v119 bool = (0 <= 1)
                    v1 -= (^(v1))
                    _ = fi118
                    _ = v119
                }
                _ = fi116
                _ = v117
            }
        }
    case -935:
        v120 := []int{1}
        _ = v120
    case 42:
        if true {
            if (!(-128 <= 42)) {
                if (!true) {
                    switch "" {
                    case "a":
                        if true {
                            if (42 <= -1) {
                                fmt.Println(v1)
                            }
                        } else {
                            var v121 float32 = -2.25
                            v122 := GenS5{F0: -588}
                            _ = v121
                            _ = v122
                        }
                    case "hello":
                        v1 = (^(v1))
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        for fi123 := 0; fi123 < 16; fi123++ {
                            v113 = uint64(v1)
                            _ = fi123
                        }
                    case "hello":
                        if true {
                            v124 := GenS1{F0: ("foo bar" + "\n\t"), F1: 42}
                            _ = v124
                        }
                    case "":
                        v125 := GenS12{F0: -1, F1: false}
                        _ = v125
                    }
                    for fi126 := 0; fi126 < 16; fi126++ {
                        var v127 int32 = int32(v3)
                        _ = fi126
                        _ = v127
                    }
                    switch ("a" + "hello") {
                    case "foo bar":
                        var v128 uint16 = 128
                        _ = v128
                    case "":
                        v3 = (v3 * (v3 + v3))
                    case "\n\t":
                        switch 293 {
                        case 4294967295:
                            v129 := GenS13{F0: 1000}
                            _ = v129
                        case 127:
                            var v130 float32 = (-2.25 / 3.14159)
                            _ = v130
                        case 0:
                            for fi131 := 0; fi131 < 5; fi131++ {
                                v132 := []int{fi131, fi131, (^(v3))}
                                v3 -= fi131
                                fmt.Println(v3, v3, v1)
                                _ = fi131
                                _ = v132
                            }
                        default:
                            fmt.Println(v1)
                        }
                    default:
                        v133 := v2[v3]
                        _ = v133
                    }
                }
                fmt.Println(v3)
                for fi134 := 0; fi134 < 5; fi134++ {
                    var v135 int = (v3 ^ fi134)
                    var v136 int8 = -69
                    if ("a" >= ("" + "\n\t")) {
                        v137 := []int{fi134, -61}
                        v138 := GenS14{F0: (42 > 2147483647), F1: int32(v136)}
                        switch (v135 << ((42) & 63)) {
                        case 4294967295:
                            v137 = append(v137, (^(fi134)))
                        case 678:
                            for fi139 := 0; fi139 < 6; fi139++ {
                                var v140 int = int(v136)
                                var v141 string = "hello"
                                switch v135 {
                                case -1:
                                    _ = len(v2)
                                case 32767:
                                    v135 -= 0
                                case 42:
                                    if (!(7 > genFn1(v1))) {
                                        var v142 float32 = float32(3.14159)
                                        v1 = v1
                                        v143 := GenS9{F0: fi139, F1: (v113 ^ v113), F2: uint32(v113)}
                                        _ = v142
                                        _ = v143
                                    }
                                }
                                if true {
                                    _ = len(v2)
                                } else {
                                    if (true || ((fi139 * (fi134 % ((9223372036854775807) | 1))) < v3)) {
                                        fmt.Println(fi134, v1)
                                    }
                                    switch (-(fi134)) {
                                    case 2147483647:
                                        if len(v137) > 0 {
                                            v137[0] = (v140 - fi134)
                                        }
                                    case 9223372036854775807:
                                        v137 = append(v137, v135)
                                    }
                                }
                                _ = len(v137)
                                _ = fi139
                                _ = v140
                                _ = v141
                            }
                        case -178:
                            _ = len(v2)
                        default:
                            var v144 uint32 = 2147483647
                            _ = v144
                        }
                        _ = v137
                        _ = v138
                    } else {
                        if (true || false) {
                            for fi145 := 0; fi145 < 12; fi145++ {
                                fmt.Println(v1, v1)
                                var v146 int8 = (v136 ^ v136)
                                _ = fi145
                                _ = v146
                            }
                        } else {
                            delete(v2, -87)
                        }
                    }
                    _ = fi134
                    _ = v135
                    _ = v136
                }
            } else {
                v1 = (v1 % ((v1) | 1))
            }
            var v147 uint64 = 0
            v113 += v113
            _ = v147
        } else {
            if ("foo bar" <= "a") {
                var v148 int16 = 42
                _ = v148
            } else {
                v149 := v2[v3]
                for fi150 := 0; fi150 < 6; fi150++ {
                    v3 += fi150
                    _ = fi150
                }
                _ = v149
            }
            if (1 > 590) {
                var v151 float64 = 3.14159
                _ = v151
            } else {
                var v152 int32 = 0
                fmt.Println(v1, v1, v152)
                _ = v152
            }
        }
    default:
        v1 = (v1 >> ((18446744073709551615) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v109
    _ = v113
    _ = v2
    _ = v3
}
