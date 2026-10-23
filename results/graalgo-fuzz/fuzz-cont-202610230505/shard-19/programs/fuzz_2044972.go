package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint; F1 uint; F2 string }
type GenS5 struct { F0 int8; F1 string; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 uint; F1 int32; F2 int16 }
type GenS11 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -9223372036854775808
    var v3 int = 127
    if false {
        for fi4 := 0; fi4 < 19; fi4++ {
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 int64 = 1
                v7 := []int{128, fi4}
                v2 -= (v2 % ((v2) | 1))
                _ = fi5
                _ = v6
                _ = v7
            }
            for fi8 := 0; fi8 < 20; fi8++ {
                v3 = v3
                v3 = 127
                delete(v1, 2)
                _ = fi8
            }
            v3 += (v3 ^ 9223372036854775807)
            _ = fi4
        }
        var v9 uint = 744
        _ = len(v1)
        _ = v9
    }
    if (!false) {
        var v10 int64 = 42
        var v11 uint64 = 1
        var v12 uint16 = 42
        _ = v10
        _ = v11
        _ = v12
    } else {
        v13 := []int{v3, v3, 2147483647}
        var v14 float64 = (((3.14159 - -2.25) * 0.0) / (-2.25 + (0.1 / 0.1)))
        _ = v13
        _ = v14
    }
    v3 -= v3
    if (!((false && (float32(0.1) != 1.5)) || (!false))) {
        v3 = v3
        v15 := GenS1{F0: 2, F1: 0}
        _ = v15
    } else {
        switch ("hello" + "") {
        case "a":
            var v16 uint32 = 255
            _ = v16
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS1{F0: 2, F1: 2}
            if (128 > 71) {
                for fi19 := 0; fi19 < 3; fi19++ {
                    delete(v1, 64)
                    v20 := GenS2{F0: ""}
                    switch (fi17 >> ((614) & 63)) {
                    case -450:
                        v2 = v2
                    case 4294967295:
                        if false {
                            var v21 bool = (127 != 127)
                            _ = v21
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                v22 := []int{v3, 1, -9223372036854775808}
                if (0 < 148) {
                    switch -374 {
                    case -9223372036854775808:
                        v3 = (fi17 & (fi17 << ((604) & 63)))
                    case -317:
                        switch (v3 + -9223372036854775808) {
                        case 4294967295:
                            var v23 int64 = (v2 << ((255) & 63))
                            _ = v23
                        default:
                            var v24 string = "hello"
                            _ = v24
                        }
                    case 32768:
                        var v25 uint = 1
                        _ = v25
                    }
                    v2 += v2
                }
                _ = v22
            } else {
                switch ("" + ("\n\t" + "hello")) {
                case "foo bar":
                    var v26 int = (^(v3))
                    _ = v26
                default:
                    v1[2147483647] = (!(!false))
                }
            }
            _ = fi17
            _ = v18
        }
        var v27 uint8 = 155
        _ = v27
    }
    v28 := []int{v3, -147, v3, v3}
    var v29 uint = 65535
    var v30 bool = true
    if true {
        if true {
            switch "a" {
            case "foo bar":
                if (128 <= uint8(v29)) {
                    v29 -= v29
                    var v31 int8 = 95
                    _ = v31
                } else {
                    var v32 uint16 = 450
                    var v33 uint64 = 369
                    _ = v32
                    _ = v33
                }
            default:
                var v34 bool = true
                _ = v34
            }
            var v35 int16 = 0
            if true {
                v36 := GenS1{F0: 0, F1: 255}
                if (-2.25 > 0.0) {
                    var v37 float64 = -2.25
                    _ = v37
                }
                switch v3 {
                case -1:
                    v38 := GenS3{F0: -128, F1: 0}
                    _ = v38
                case 783:
                    v30 = v30
                default:
                    switch "\n\t" {
                    case "":
                        var v39 int = (v3 ^ (v3 << (((v29 | 4294967295)) & 63)))
                        _ = v39
                    default:
                        for fi40 := 0; fi40 < 1; fi40++ {
                            for fi41 := 0; fi41 < 10; fi41++ {
                                if true {
                                    var v42 uint16 = 255
                                    _ = len(v28)
                                    var v43 int32 = 42
                                    var v44 int64 = v2
                                    fmt.Println(v2, v44, v29)
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                                if true {
                                    var v45 int32 = 2147483647
                                    switch "" {
                                    case "a":
                                        var v46 int8 = 127
                                        _ = v46
                                    case "":
                                        var v47 int16 = v35
                                        _ = v47
                                    case "hello":
                                        v48 := GenS4{F0: v29, F1: (v29 / (((v29 % ((v29) | 1))) | 1)), F2: (("foo bar" + ("foo bar" + "foo bar")) + "foo bar")}
                                        _ = v48
                                    default:
                                        switch 127 {
                                        case -1:
                                            switch "\n\t" {
                                            case "hello":
                                                var v49 int = fi41
                                                _ = v49
                                            case "a":
                                                v45 -= v45
                                            }
                                        case 373:
                                            var v50 float32 = float32(3.14159)
                                            _ = v50
                                        default:
                                            var v51 uint = v29
                                            _ = v51
                                        }
                                    }
                                    fmt.Println(fi41)
                                    _ = v45
                                }
                                v35 -= (v35 % (((-(v35))) | 1))
                                _ = fi41
                            }
                            v52 := GenS3{F0: -124, F1: -1}
                            _ = fi40
                            _ = v52
                        }
                    }
                }
                _ = v36
            } else {
                var v53 int = (v3 << ((v29) & 63))
                switch v53 {
                case -9223372036854775808:
                    if false {
                        v53 -= (v53 | -2147483648)
                    }
                case 42:
                    switch (^(v3)) {
                    case 42:
                        v2 = int64(v35)
                    case 127:
                        for fi54 := 0; fi54 < 8; fi54++ {
                            var v55 bool = v30
                            if (!(v55 && true)) {
                                v3 = 2147483647
                            } else {
                                var v56 uint = v29
                                for fi57 := 0; fi57 < 1; fi57++ {
                                    fmt.Println(v55)
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                    _ = fi57
                                }
                                fmt.Println(fi54, v2)
                                _ = v56
                            }
                            var v58 uint32 = 42
                            _ = fi54
                            _ = v55
                            _ = v58
                        }
                    default:
                        var v59 string = ("" + "\n\t")
                        _ = v59
                    }
                case 1:
                    fmt.Println(v53, v3, v2)
                }
                var v60 string = "foo bar"
                _ = v53
                _ = v60
            }
            _ = v35
        } else {
            var v61 int16 = 976
            _ = v61
        }
        if (-128 > 42) {
            if (!true) {
                var v62 float64 = 0.0
                _ = v62
            } else {
                var v63 int8 = -128
                v29 += (v29 + (^(v29)))
                for fi64 := 0; fi64 < 5; fi64++ {
                    var v65 int16 = -920
                    var v66 uint32 = 2147483648
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                _ = v63
            }
            if ((v2 % ((v2) | 1)) < (v2 << ((255) & 63))) {
                for fi67 := 0; fi67 < 20; fi67++ {
                    if (true || ((v30 && true) || (!v30))) {
                        if false {
                            var v68 int32 = -184
                            var v69 int32 = 544
                            _ = v68
                            _ = v69
                        }
                    } else {
                        if len(v28) > 0 {
                            v28[0] = v3
                        }
                        if ((((true && (-2.25 >= (1.5 - 0.0))) || v30) || v30) && (42 >= 777)) {
                            if (93 <= 1) {
                                if v30 {
                                    var v70 string = "hello"
                                    _ = v70
                                }
                            }
                        } else {
                            v71 := GenS2{F0: "foo bar"}
                            for fi72 := 0; fi72 < 7; fi72++ {
                                _ = len(v1)
                                if v30 {
                                    v29 -= (v29 << ((42) & 63))
                                } else {
                                    switch ("\n\t" + "foo bar") {
                                    case "\n\t":
                                        _ = len(v1)
                                    default:
                                        var v73 string = "\n\t"
                                        _ = v73
                                    }
                                    fmt.Println(fi72, v3)
                                }
                                fmt.Println(v3, v29, fi72)
                                _ = fi72
                            }
                            v29 = 18446744073709551615
                            _ = v71
                        }
                        v2 += 482
                    }
                    _ = fi67
                }
                v74 := GenS5{F0: -65, F1: "", F2: 117}
                switch "" {
                case "foo bar":
                    var v75 uint8 = 128
                    _ = v75
                case "":
                    for fi76 := 0; fi76 < 17; fi76++ {
                        var v77 float64 = ((3.14159 - 3.14159) + (0.0 * (1e10 + 1e10)))
                        v78 := GenS4{F0: 65535, F1: (v29 << ((42) & 63)), F2: "a"}
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                case "a":
                    var v79 float64 = 3.14159
                    _ = v79
                default:
                    fmt.Println(v3)
                }
                _ = v74
            } else {
                if v30 {
                    v29 += (v29 % ((v29) | 1))
                    if (-128 >= -103) {
                        switch 9223372036854775807 {
                        case 42:
                            for fi80 := 0; fi80 < 5; fi80++ {
                                switch "a" {
                                case "a":
                                    if len(v28) > 0 {
                                        v28[0] = 0
                                    }
                                case "hello":
                                    var v81 uint16 = 255
                                    _ = v81
                                case "foo bar":
                                    v82 := GenS6{F0: 1}
                                    _ = v82
                                }
                                if (fi80 < fi80) {
                                    for fi83 := 0; fi83 < 14; fi83++ {
                                        v2 = v2
                                        switch ("foo bar" + "\n\t") {
                                        case "a":
                                            var v84 uint8 = 114
                                            _ = v84
                                        default:
                                            fmt.Println(v30, v3)
                                        }
                                        var v85 uint64 = 811
                                        v85 = (v85 & v85)
                                        _ = fi83
                                        _ = v85
                                    }
                                }
                                _ = fi80
                            }
                        }
                    }
                    switch v3 {
                    case 1:
                        v86 := GenS7{F0: 47}
                        _ = v86
                    }
                } else {
                    for fi87 := 0; fi87 < 17; fi87++ {
                        for fi88 := 0; fi88 < 3; fi88++ {
                            _ = len(v1)
                            var v89 int32 = -1
                            _ = fi88
                            _ = v89
                        }
                        _ = fi87
                    }
                    for fi90 := 0; fi90 < 18; fi90++ {
                        v91 := GenS5{F0: -128, F1: "\n\t", F2: 100}
                        v29 = v29
                        if (v29 > 18446744073709551615) {
                            var v92 uint16 = 0
                            switch "hello" {
                            case "":
                                var v93 bool = ((v30 || v30) && (!(97 == 18446744073709551615)))
                                _ = v93
                            default:
                                v29 += 4294967295
                            }
                            var v94 float64 = -2.25
                            _ = v92
                            _ = v94
                        } else {
                            fmt.Println(v2, v2, v3)
                            if (!true) {
                                v95 := v1[986]
                                switch -1 {
                                case 1:
                                    v30 = ((255 > 42) && (-32768 <= 1))
                                case -9223372036854775808:
                                    if false {
                                        if (2147483647 != -478) {
                                            var v96 int32 = 1
                                            for fi97 := 0; fi97 < 20; fi97++ {
                                                v3 += 127
                                                _ = fi97
                                            }
                                            _ = v96
                                        }
                                    } else {
                                        var v98 float32 = (-2.25 * 3.14159)
                                        var v99 int64 = (v2 | (v2 - -343))
                                        v100 := GenS8{F0: 127}
                                        _ = v100
                                        _ = v98
                                        _ = v99
                                    }
                                default:
                                    for fi101 := 0; fi101 < 11; fi101++ {
                                        var v102 uint8 = 127
                                        _ = len(v1)
                                        v103 := GenS4{F0: (v29 - v29), F1: 589, F2: ("\n\t" + ("\n\t" + ((((("hello" + "") + (("hello" + "a") + ((("hello" + (("\n\t" + (("\n\t" + "\n\t") + "foo bar")) + "foo bar")) + "foo bar") + "foo bar"))) + ((("\n\t" + "hello") + ("a" + (("a" + "\n\t") + "a"))) + "a")) + ("a" + "\n\t")) + ("a" + (((("foo bar" + "") + "\n\t") + ("" + "\n\t")) + ("foo bar" + "a"))))))}
                                        _ = fi101
                                        _ = v102
                                        _ = v103
                                    }
                                }
                                var v104 uint8 = 255
                                _ = v104
                                _ = v95
                            }
                        }
                        _ = fi90
                        _ = v91
                    }
                }
            }
            v30 = ("" != "foo bar")
        } else {
            switch ("hello" + "foo bar") {
            case "foo bar":
                fmt.Println(v29)
            }
            fmt.Println(v2)
            if v30 {
                fmt.Println(v30, v3, v30)
                var v105 int16 = -1
                v106 := GenS7{F0: -128}
                _ = v105
                _ = v106
            } else {
                for fi107 := 0; fi107 < 2; fi107++ {
                    var v108 uint = (v29 + v29)
                    _ = fi107
                    _ = v108
                }
            }
        }
        if (v30 || false) {
            for fi109 := 0; fi109 < 19; fi109++ {
                if true {
                    v2 += v2
                    var v110 bool = v30
                    _ = v110
                } else {
                    v1[42] = (!v30)
                }
                var v111 bool = (255 < 0)
                if len(v28) > 0 {
                    v28[0] = v3
                }
                _ = fi109
                _ = v111
            }
            for fi112 := 0; fi112 < 18; fi112++ {
                fmt.Println(v30, v3)
                for fi113 := 0; fi113 < 9; fi113++ {
                    v30 = v30
                    fmt.Println(v29, v30, v3)
                    _ = fi113
                }
                if true {
                    if (479 > 42) {
                        if (!(!false)) {
                            switch ("hello" + "a") {
                            case "hello":
                                if v30 {
                                    var v114 float64 = 0.0
                                    var v115 float64 = (3.14159 - (1.5 - (0.0 - v114)))
                                    _ = v114
                                    _ = v115
                                }
                            case "a":
                                v30 = v30
                            default:
                                if v30 {
                                    var v116 int32 = -2147483648
                                    var v117 uint = v29
                                    _ = v116
                                    _ = v117
                                }
                            }
                            v2 = v2
                            if (v30 || v30) {
                                if len(v28) > 0 {
                                    v28[0] = (v3 + v3)
                                }
                                var v118 int32 = 545
                                fmt.Println(v2, fi112, v30)
                                _ = v118
                            } else {
                                v29 -= (v29 % ((65535) | 1))
                                var v119 bool = v30
                                v119 = (0 == 1)
                                _ = v119
                            }
                        }
                        v29 = (v29 ^ (v29 << ((v29) & 63)))
                    }
                    if v30 {
                        var v120 int8 = 84
                        _ = v120
                    }
                } else {
                    var v121 int = (-(v3))
                    var v122 float32 = -2.25
                    for fi123 := 0; fi123 < 6; fi123++ {
                        var v124 float64 = 0.1
                        v125 := GenS9{F0: 27, F1: ""}
                        v30 = (72 == 2)
                        _ = fi123
                        _ = v124
                        _ = v125
                    }
                    _ = v121
                    _ = v122
                }
                _ = fi112
            }
            for fi126 := 0; fi126 < 6; fi126++ {
                if (!((0 <= 42) && false)) {
                    var v127 int16 = 1
                    fmt.Println(fi126, v127)
                    _ = v127
                }
                var v128 int32 = 2147483647
                switch "a" {
                case "\n\t":
                    fmt.Println(v2, fi126)
                }
                _ = fi126
                _ = v128
            }
        } else {
            var v129 string = "a"
            var v130 int16 = -32768
            _ = v129
            _ = v130
        }
    } else {
        var v131 uint = (v29 >> ((v29) & 63))
        if (true && v30) {
            switch "" {
            case "":
                if false {
                    if false {
                        delete(v1, 830)
                        var v132 uint32 = 2147483647
                        var v133 float64 = float64(float32((3.14159 + (0.0 - (float64(1e10) - (3.14159 * (1.5 / 1.0)))))))
                        _ = v132
                        _ = v133
                    } else {
                        switch (^(v3)) {
                        case 127:
                            var v134 uint16 = 32767
                            _ = v134
                        case -2147483648:
                            v131 = v29
                        }
                    }
                    if v30 {
                        v135 := GenS3{F0: -1, F1: -1}
                        v29 = (v131 / ((0) | 1))
                        _ = v135
                    }
                    var v136 int64 = v2
                    _ = v136
                } else {
                    fmt.Println(v3)
                }
            case "foo bar":
                switch (v3 | (v3 + (v3 << ((v29) & 63)))) {
                case -9223372036854775808:
                    if true {
                        v137 := GenS5{F0: -1, F1: "foo bar", F2: 100}
                        v138 := v1[42]
                        v30 = (1e10 == (0.1 / -2.25))
                        _ = v137
                        _ = v138
                    } else {
                        v139 := GenS4{F0: v29, F1: 42, F2: "\n\t"}
                        _ = v139
                    }
                case 42:
                    v140 := GenS10{F0: (v29 + (v131 * v131)), F1: -383, F2: 42}
                    _ = v140
                case -877:
                    v3 = (v3 % (((-(v3))) | 1))
                }
            case "\n\t":
                v141 := GenS11{F0: ("" + "")}
                _ = v141
            default:
                _ = len(v1)
            }
            switch "a" {
            case "\n\t":
                var v142 uint64 = 909
                _ = v142
            case "foo bar":
                v143 := v1[65535]
                _ = v143
            }
        } else {
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            var v144 uint8 = 128
            if (false || v30) {
                var v145 uint32 = 2147483648
                var v146 int = (v3 | (v3 % ((127) | 1)))
                for fi147 := 0; fi147 < 8; fi147++ {
                    fmt.Println(v144, v30, fi147)
                    _ = fi147
                }
                _ = v145
                _ = v146
            } else {
                var v148 string = "\n\t"
                var v149 float64 = 0.1
                _ = v148
                _ = v149
            }
            _ = v144
        }
        switch v3 {
        case 539:
            var v150 uint32 = 840
            _ = v150
        case 873:
            fmt.Println(v131)
        }
        _ = v131
    }
    var v151 int16 = 558
    var v152 uint8 = 2
    var v153 int = v3
    v29 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v151
    _ = v152
    _ = v153
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
