package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 string; F2 int16 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int16; F1 uint; F2 int32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 int16; F2 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 uint16; F2 int16 }
type GenS12 struct { F0 bool; F1 int16 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int32) int8 {
    v5 := GenS1{F0: -1}
    var v6 int = 1
    fmt.Println(p1, p1, v6)
    var v7 float32 = 0.0
    v8 := []int{-9223372036854775808, v6, (v6 & v6), 127}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1() uint {
    for fi9 := 0; fi9 < 20; fi9++ {
        switch (("" + "a") + "a") {
        case "a":
            v10 := []int{2147483647, -211}
            _ = v10
        case "hello":
            var v11 uint64 = 0
            _ = v11
        case "\n\t":
            v12 := GenS2{F0: 758, F1: "foo bar", F2: -32768}
            _ = v12
        default:
            var v13 uint = 211
            _ = v13
        }
        var v14 int16 = -1
        _ = fi9
        _ = v14
    }
    v15 := GenS3{F0: -6, F1: -351}
    var v16 string = ("foo bar" + "\n\t")
    v16 = "foo bar"
    _ = v15
    _ = v16
    return 4294967295
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]uint32)
    var v3 uint = 0
    var v4 bool = true
    if (("foo bar" + "\n\t") == ("a" + ("foo bar" + "foo bar"))) {
        for fi17 := 0; fi17 < 6; fi17++ {
            v3 += 18446744073709551615
            v18 := GenS2{F0: (v3 ^ (v3 - v3)), F1: "\n\t", F2: 0}
            if (127 != 31) {
                v19 := GenS1{F0: -9223372036854775808}
                switch (fi17 ^ (fi17 & (fi17 * (^(fi17))))) {
                case 4294967295:
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v21 := []int{(^(fi17)), (fi20 & 42)}
                        v4 = v4
                        var v22 string = (("foo bar" + "") + "hello")
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                case -1:
                    v4 = v4
                case 32768:
                    var v23 bool = (0 < 18446744073709551615)
                    _ = v23
                default:
                    var v24 bool = true
                    _ = v24
                }
                if true {
                    for fi25 := 0; fi25 < 16; fi25++ {
                        if (96 > -128) {
                            v4 = false
                            v26 := GenS4{F0: (v1 & (v1 + v1)), F1: v3, F2: 2147483647}
                            _ = v26
                        } else {
                            if true {
                                switch 2147483647 {
                                case 9223372036854775807:
                                    _ = len(v2)
                                case 0:
                                    v1 += (v1 & 0)
                                default:
                                    var v27 int8 = 1
                                    _ = v27
                                }
                                switch (fi17 / ((-684) | 1)) {
                                case 297:
                                    v28 := GenS2{F0: (v3 + 42), F1: "", F2: v1}
                                    _ = v28
                                }
                                switch (^(fi25)) {
                                case 42:
                                    delete(v2, "\n\t")
                                default:
                                    fmt.Println(v3, v1, v1)
                                    v29 := GenS5{F0: 2, F1: "hello"}
                                    v4 = (!v4)
                                    _ = v29
                                }
                            }
                            if (((((!((3.14159 < 1.5) || v4)) || ((-9223372036854775808 != 1) && v4)) || (-900 >= -2147483648)) && v4) || true) {
                                var v30 float64 = (((1e10 / (0.1 * 0.1)) / ((0.1 / -2.25) / ((((float64(float32(0.0)) - (3.14159 - (1e10 / (((1e10 + 3.14159) - -2.25) - (-2.25 / (1.5 * (0.0 + (-2.25 + 1.5)))))))) / 1.0) + ((-2.25 - 0.0) + 0.1)) + 0.1))) * -2.25)
                                fmt.Println(v1, v30)
                                _ = v30
                            } else {
                                switch (fi25 * fi25) {
                                case -9223372036854775808:
                                    v3 += 255
                                default:
                                    v1 -= -32768
                                }
                                v1 = (v1 & v1)
                            }
                            var v31 int64 = 293
                            _ = v31
                        }
                        if (750 > 42) {
                            v32 := v2["hello"]
                            fmt.Println(v3, v3)
                            var v33 uint8 = 2
                            _ = v32
                            _ = v33
                        }
                        _ = fi25
                    }
                    var v34 uint16 = 498
                    if true {
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v36 := v2["a"]
                            var v37 uint32 = v36
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        v3 -= (v3 & (v3 * 1))
                    }
                    _ = v34
                }
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 int32 = 1
            fmt.Println(v3, v4)
            fmt.Println(v1, v1)
            _ = fi38
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        if v4 {
            _ = len(v2)
            switch "" {
            case "a":
                for fi41 := 0; fi41 < 18; fi41++ {
                    if false {
                        v42 := GenS6{F0: -778}
                        fmt.Println(fi41)
                        _ = v42
                    }
                    _ = fi41
                }
            case "\n\t":
                v3 += (v3 * (v3 | v3))
            default:
                v4 = (!false)
            }
            if v4 {
                if v4 {
                    switch "\n\t" {
                    case "\n\t":
                        for fi43 := 0; fi43 < 9; fi43++ {
                            var v44 string = ""
                            _ = fi43
                            _ = v44
                        }
                    case "":
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v3 = genFn1()
                            fmt.Println(fi40, v3)
                            _ = fi45
                        }
                    }
                    v1 += v1
                }
                v4 = ("" <= "a")
                switch ("\n\t" + "hello") {
                case "\n\t":
                    for fi46 := 0; fi46 < 11; fi46++ {
                        v1 = v1
                        _ = fi46
                    }
                case "":
                    var v47 float32 = 0.0
                    _ = v47
                case "hello":
                    v3 -= v3
                default:
                    if (34 == 0) {
                        v3 -= v3
                    }
                }
            } else {
                _ = len(v2)
            }
        } else {
            v48 := []int{fi40, (fi40 - (-(fi40))), (fi40 + (fi40 - fi40)), fi40}
            _ = v48
        }
        if false {
            v4 = (float32(1e10) < (0.0 * (3.14159 - 3.14159)))
        }
        for fi49 := 0; fi49 < 2; fi49++ {
            if true {
                switch (("" + "\n\t") + "foo bar") {
                case "foo bar":
                    v50 := GenS7{F0: 2, F1: (^(v1)), F2: ((457 == 775) && v4)}
                    _ = v50
                case "":
                    if (!(!v4)) {
                        var v51 string = "foo bar"
                        v1 -= (v1 << ((genFn1()) & 63))
                        v1 += -1
                        var v52 float32 = 0.0
                        _ = v51
                        _ = v52
                    }
                default:
                    v3 = v3
                    v53 := GenS1{F0: 9223372036854775807}
                    _ = v53
                }
                v54 := []int{127, (fi40 << ((1) & 63)), fi40, fi40}
                v54 = append(v54, fi49)
                _ = v54
            }
            v3 = (v3 | v3)
            v55 := []int{fi40, -1, (fi49 | 9223372036854775807), (fi40 << (((v3 / ((v3) | 1))) & 63))}
            _ = fi49
            _ = v55
        }
        _ = fi40
    }
    var v56 uint = (^(v3))
    v56 = (v3 + v56)
    v57 := []int{}
    if (-1 == -1) {
        switch 0 {
        case 4294967295:
            v58 := v2[((("" + "foo bar") + "") + "a")]
            _ = v58
        default:
            for fi59 := 0; fi59 < 16; fi59++ {
                fmt.Println(v4, v56, v1)
                _ = fi59
            }
        }
        var v60 uint32 = 362
        _ = v60
    } else {
        if (true || (0 != uint16(v56))) {
            if (1e10 != ((float64(0.0) + 1.5) / float64(0.1))) {
                for fi61 := 0; fi61 < 4; fi61++ {
                    v3 -= v56
                    _ = fi61
                }
                v1 += v1
            } else {
                for fi62 := 0; fi62 < 16; fi62++ {
                    var v63 float32 = ((3.14159 + 3.14159) * ((1.5 - 3.14159) + ((0.1 - float32((0.1 - 1e10))) * -2.25)))
                    switch ("\n\t" + "") {
                    case "":
                        if v4 {
                            if (v56 != 18446744073709551615) {
                                var v64 uint8 = 2
                                switch "foo bar" {
                                case "\n\t":
                                    if ((!(((v1 != (v1 % ((1000) | 1))) || v4) && true)) && (42 >= -1)) {
                                        for fi65 := 0; fi65 < 9; fi65++ {
                                            var v66 int8 = 42
                                            switch ("\n\t" + "a") {
                                            case "foo bar":
                                                v66 += (v66 / ((genFn0(2, 42)) | 1))
                                            case "hello":
                                                switch 262 {
                                                case 42:
                                                    v67 := GenS1{F0: 9223372036854775807}
                                                    _ = v67
                                                case 385:
                                                    v1 -= (v1 - (v1 - v1))
                                                case 127:
                                                    var v68 int64 = 437
                                                    v69 := GenS8{F0: (v64 > v64)}
                                                    _ = v68
                                                    _ = v69
                                                }
                                            }
                                            _ = fi65
                                            _ = v66
                                        }
                                        v4 = (0 < -9223372036854775808)
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            v70 := GenS1{F0: 130}
                                            _ = v70
                                        case "":
                                            for fi71 := 0; fi71 < 8; fi71++ {
                                                var v72 uint8 = 127
                                                _ = len(v57)
                                                _ = fi71
                                                _ = v72
                                            }
                                        case "a":
                                            switch "a" {
                                            case "":
                                                var v73 int32 = 42
                                                _ = v73
                                            case "\n\t":
                                                v3 += 1
                                            default:
                                                v74 := GenS2{F0: v3, F1: "\n\t", F2: (^(v1))}
                                                _ = v74
                                            }
                                        }
                                    }
                                case "":
                                    v75 := GenS9{F0: "a"}
                                    _ = v75
                                }
                                _ = v64
                            } else {
                                if (42 <= 65535) {
                                    v63 += v63
                                }
                                if v4 {
                                    fmt.Println(v3, fi62)
                                    v76 := GenS6{F0: -656}
                                    _ = v76
                                }
                            }
                            v56 = (^(v3))
                            _ = len(v2)
                        } else {
                            v1 -= v1
                        }
                    case "a":
                        v63 += -2.25
                    }
                    switch "foo bar" {
                    case "":
                        if v4 {
                            if len(v57) > 0 {
                                v57[0] = (fi62 ^ (fi62 | fi62))
                            }
                        } else {
                            fmt.Println(fi62, fi62, v3)
                        }
                    case "foo bar":
                        var v77 uint32 = 2147483647
                        _ = v77
                    case "a":
                        if true {
                            var v78 int32 = 1
                            _ = v78
                        } else {
                            v57 = append(v57, fi62)
                            if (!(false && (((true || (!false)) && (255 != 255)) && (!v4)))) {
                                if v4 {
                                    var v79 int64 = -1
                                    if true {
                                        v80 := GenS5{F0: 0, F1: "hello"}
                                        switch (("" + "a") + "") {
                                        case "foo bar":
                                            fmt.Println(v3, v4, v63)
                                        }
                                        v63 += v63
                                        _ = v80
                                    } else {
                                        switch "a" {
                                        case "a":
                                            switch "\n\t" {
                                            case "\n\t":
                                                delete(v2, "\n\t")
                                            }
                                        case "foo bar":
                                            fmt.Println(v3, v3, v79)
                                        }
                                        v63 -= -2.25
                                    }
                                    _ = v79
                                }
                                if len(v57) >= 2 {
                                    v57 = v57[1:len(v57)]
                                }
                                var v81 float32 = (v63 - v63)
                                _ = v81
                            }
                        }
                    default:
                        v63 = 1e10
                    }
                    _ = fi62
                    _ = v63
                }
            }
            fmt.Println(v4, v3, v1)
        }
    }
    delete(v2, "a")
    v2["a"] = 2
    var v82 bool = (65535 != 65535)
    if v4 {
        var v83 string = "\n\t"
        for fi84 := 0; fi84 < 11; fi84++ {
            var v85 int32 = -2147483648
            var v86 int16 = 1
            v87 := GenS10{F0: 845}
            _ = fi84
            _ = v85
            _ = v86
            _ = v87
        }
        v56 = 2
        _ = v83
    }
    fmt.Println(v56, v56, v82)
    v88 := GenS11{F0: v1, F1: 65535, F2: v1}
    _ = len(v57)
    switch "hello" {
    case "a":
        switch "a" {
        case "":
            for fi89 := 0; fi89 < 9; fi89++ {
                if ("a" != "a") {
                    if false {
                        var v90 int8 = 42
                        if false {
                            v90 += v90
                        } else {
                            switch -2147483648 {
                            case 127:
                                v4 = v82
                            case -904:
                                v91 := GenS12{F0: true, F1: v1}
                                _ = v91
                            case 846:
                                fmt.Println(v1)
                            }
                        }
                        _ = v90
                    } else {
                        fmt.Println(v3)
                        if v82 {
                            if v82 {
                                var v92 uint32 = 2147483648
                                switch "\n\t" {
                                case "":
                                    v93 := GenS13{F0: fi89}
                                    _ = v93
                                case "foo bar":
                                    if (174 != 2) {
                                        v4 = ((v92 * 2147483647) < v92)
                                        for fi94 := 0; fi94 < 18; fi94++ {
                                            var v95 int16 = v1
                                            if ((0.0 + ((-2.25 * (-2.25 * (3.14159 + 1.5))) * 1e10)) < 1e10) {
                                                delete(v2, "")
                                            }
                                            v96 := GenS13{F0: (fi94 ^ 128)}
                                            _ = fi94
                                            _ = v95
                                            _ = v96
                                        }
                                    } else {
                                        v97 := GenS6{F0: 42}
                                        v4 = (!false)
                                        _ = v97
                                    }
                                case "hello":
                                    for fi98 := 0; fi98 < 17; fi98++ {
                                        var v99 string = (("a" + "hello") + "\n\t")
                                        var v100 int32 = -2147483648
                                        _ = fi98
                                        _ = v100
                                        _ = v99
                                    }
                                default:
                                    switch "a" {
                                    case "a":
                                        v101 := GenS5{F0: 18446744073709551615, F1: "a"}
                                        _ = v101
                                    case "":
                                        var v102 int = (fi89 << ((genFn1()) & 63))
                                        _ = v102
                                    }
                                }
                                for fi103 := 0; fi103 < 2; fi103++ {
                                    var v104 float32 = (3.14159 * 0.0)
                                    var v105 int8 = 0
                                    _ = fi103
                                    _ = v104
                                    _ = v105
                                }
                                _ = v92
                            } else {
                                fmt.Println(v56, v56)
                                v1 = v1
                                switch "a" {
                                case "\n\t":
                                    var v106 uint = genFn1()
                                    _ = v106
                                default:
                                    v2["a"] = 544
                                }
                            }
                            for fi107 := 0; fi107 < 10; fi107++ {
                                v56 -= 299
                                fmt.Println(v56, fi89, v4)
                                if (-2.25 == float32(3.14159)) {
                                    if v82 {
                                        v108 := v2["a"]
                                        v57 = append(v57, fi107)
                                        _ = v108
                                    } else {
                                        if v4 {
                                            v109 := GenS9{F0: ("hello" + ((("foo bar" + "foo bar") + "foo bar") + "hello"))}
                                            var v110 int16 = 42
                                            var v111 string = "a"
                                            _ = v109
                                            _ = v110
                                            _ = v111
                                        }
                                    }
                                }
                                fmt.Println(v3, v56, fi107)
                                _ = fi107
                            }
                            fmt.Println(v1)
                        }
                        var v112 uint32 = 65535
                        _ = v112
                    }
                    if (!v4) {
                        if ((v4 || (!false)) && (!(!v4))) {
                            var v113 uint = (v3 >> ((v3) & 63))
                            _ = v113
                        }
                    } else {
                        if (306 == 800) {
                            var v114 string = ""
                            v4 = v82
                            var v115 uint16 = 32767
                            _ = v114
                            _ = v115
                        } else {
                            if ("hello" > ("\n\t" + "")) {
                                for fi116 := 0; fi116 < 6; fi116++ {
                                    v117 := GenS9{F0: "a"}
                                    v56 += genFn1()
                                    fmt.Println(v4, v3, fi89)
                                    _ = fi116
                                    _ = v117
                                }
                                for fi118 := 0; fi118 < 4; fi118++ {
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(fi89, fi118, v3)
                                    case "a":
                                        var v119 string = ("a" + "a")
                                        _ = v119
                                    case "\n\t":
                                        v4 = (true || (!v82))
                                    default:
                                        fmt.Println(v82, fi118)
                                    }
                                    delete(v2, ("foo bar" + "foo bar"))
                                    v120 := GenS8{F0: (42 >= 0)}
                                    _ = fi118
                                    _ = v120
                                }
                                v121 := GenS9{F0: ("\n\t" + ("hello" + ("foo bar" + "")))}
                                _ = v121
                            } else {
                                switch ((("\n\t" + ("foo bar" + ("foo bar" + (("foo bar" + ("a" + ("hello" + ((("\n\t" + "\n\t") + "a") + "hello")))) + "a")))) + "foo bar") + ("\n\t" + "foo bar")) {
                                case "hello":
                                    if (true || (true && (v82 && (!false)))) {
                                        switch 0 {
                                        case -898:
                                            var v122 uint64 = 679
                                            _ = v122
                                        case 127:
                                            var v123 int32 = -2147483648
                                            _ = v123
                                        }
                                        for fi124 := 0; fi124 < 16; fi124++ {
                                            v82 = (228 < 218)
                                            _ = fi124
                                        }
                                    }
                                }
                                var v125 float64 = 0.0
                                _ = v125
                            }
                            for fi126 := 0; fi126 < 1; fi126++ {
                                if len(v57) > 0 {
                                    v57[0] = (^(fi126))
                                }
                                switch 2147483647 {
                                case -517:
                                    fmt.Println(v4)
                                case 4294967295:
                                    var v127 uint = v3
                                    _ = v127
                                case 1:
                                    v1 = -127
                                }
                                _ = fi126
                            }
                            v128 := GenS8{F0: v4}
                            _ = v128
                        }
                        var v129 float32 = float32((((3.14159 / 0.1) / 1.0) + 1.5))
                        fmt.Println(v129)
                        _ = v129
                    }
                }
                var v130 uint8 = 2
                for fi131 := 0; fi131 < 2; fi131++ {
                    if v82 {
                        var v132 bool = (!false)
                        _ = v132
                    } else {
                        v133 := GenS14{F0: 804}
                        _ = v133
                    }
                    v56 -= genFn1()
                    if false {
                        v4 = (2147483647 < 65535)
                    } else {
                        switch (("foo bar" + "hello") + "hello") {
                        case "foo bar":
                            _ = len(v2)
                        case "hello":
                            v134 := GenS14{F0: 2147483648}
                            _ = v134
                        case "a":
                            v130 = v130
                        default:
                            switch (fi89 & fi131) {
                            case 286:
                                var v135 uint32 = 0
                                _ = v135
                            }
                        }
                        v130 -= v130
                    }
                    _ = fi131
                }
                _ = fi89
                _ = v130
            }
        case "\n\t":
            var v136 int = -961
            _ = v136
        default:
            for fi137 := 0; fi137 < 13; fi137++ {
                var v138 int16 = 32767
                _ = fi137
                _ = v138
            }
        }
    default:
        if (!(!true)) {
            fmt.Println(v3, v1)
            switch "a" {
            case "a":
                fmt.Println(v56, v82, v1)
            case "hello":
                if len(v57) > 0 {
                    v57[0] = -667
                }
            default:
                fmt.Println(v1)
            }
        } else {
            v56 += v56
            v139 := v2[("hello" + ("a" + "foo bar"))]
            _ = v139
        }
    }
    if v82 {
        v140 := GenS10{F0: 4294967295}
        var v141 string = ("a" + "")
        _ = v140
        _ = v141
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v82
    _ = v88
}
