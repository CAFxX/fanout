package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 uint; F2 uint8 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint16; F1 string }
type GenS13 struct { F0 uint; F1 int }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint32, p1 uint, p2 string) uint8 {
    v5 := []int{42, -11, -9223372036854775808}
    var v6 float32 = -2.25
    if false {
        v5 = append(v5, 42)
        switch p2 {
        case "foo bar":
            var v7 int64 = 1
            _ = v7
        default:
            fmt.Println(p0, v6, v6)
        }
        var v8 uint8 = 0
        _ = v8
    } else {
        v9 := GenS1{F0: p1, F1: p0, F2: 0}
        fmt.Println(p1, v6, p2)
        switch (("hello" + p2) + "hello") {
        case "\n\t":
            p2 = "\n\t"
        case "foo bar":
            fmt.Println(v6, p0)
        case "a":
            v10 := GenS1{F0: (p1 * uint(p0)), F1: (^(p0)), F2: 127}
            _ = v10
        default:
            for fi11 := 0; fi11 < 18; fi11++ {
                p1 -= (^(p1))
                var v12 int8 = 127
                _ = fi11
                _ = v12
            }
        }
        _ = v9
    }
    switch 498 {
    case -531:
        var v13 int64 = -717
        _ = v13
    default:
        switch "hello" {
        case "\n\t":
            switch "a" {
            case "a":
                var v14 string = (p2 + (((((p2 + "foo bar") + "") + ("" + "hello")) + p2) + "hello"))
                _ = v14
            case "":
                if true {
                    p0 = (p0 | 551)
                } else {
                    for fi15 := 0; fi15 < 5; fi15++ {
                        if false {
                            var v16 uint = 462
                            _ = v16
                        }
                        _ = fi15
                    }
                    if ((p0 / (((p0 * p0)) | 1)) > (^(p0))) {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            for fi18 := 0; fi18 < 4; fi18++ {
                                if (!false) {
                                    if false {
                                        var v19 uint = p1
                                        if (((float32(0.1) * ((v6 * v6) / -2.25)) == (v6 / v6)) || false) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v20 := GenS1{F0: (^(v19)), F1: p0, F2: 42}
                                            fmt.Println(fi18)
                                            _ = v20
                                        }
                                        var v21 bool = false
                                        _ = v19
                                        _ = v21
                                    }
                                    p2 = "hello"
                                    if false {
                                        p0 = p0
                                        fmt.Println(p2)
                                        var v22 int32 = 287
                                        _ = v22
                                    } else {
                                        v5 = append(v5, fi17)
                                        var v23 bool = (1 < 128)
                                        v23 = v23
                                        p2 = p2
                                        _ = v23
                                    }
                                } else {
                                    fmt.Println(p0, p1)
                                    if (42 != 477) {
                                        if true {
                                            v24 := GenS2{F0: 539}
                                            var v25 float64 = 1.5
                                            var v26 uint32 = 42
                                            var v27 int8 = 100
                                            fmt.Println(fi18, v25, v6)
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        }
                                        p0 += (p0 + p0)
                                        v28 := GenS3{F0: ((42 != 1) || (float32((1e10 + float64((-2.25 * (0.1 * v6))))) != v6)), F1: 255, F2: p0}
                                        _ = v28
                                    }
                                    p1 -= p1
                                }
                                var v29 uint64 = 18446744073709551615
                                if true {
                                    var v30 int8 = 1
                                    v5 = append(v5, -2147483648)
                                    v31 := GenS4{F0: v29, F1: p2}
                                    switch "a" {
                                    case "a":
                                        fmt.Println(fi17, v29, v6)
                                    case "":
                                        if true {
                                            v32 := GenS1{F0: p1, F1: (p0 / ((716) | 1)), F2: fi17}
                                            v5 = append(v5, (fi17 / (((fi17 & 747)) | 1)))
                                            var v33 uint = p1
                                            _ = v32
                                            _ = v33
                                        } else {
                                            if (2147483647 == 4294967295) {
                                                fmt.Println(p0, p1, fi18)
                                                v29 -= uint64(p1)
                                            } else {
                                                switch p2 {
                                                case "foo bar":
                                                    v5 = append(v5, fi18)
                                                default:
                                                    p0 = p0
                                                }
                                                var v34 uint8 = 1
                                                _ = v34
                                            }
                                        }
                                    default:
                                        v30 -= v30
                                    }
                                    var v35 uint32 = 2147483647
                                    _ = v30
                                    _ = v31
                                    _ = v35
                                }
                                fmt.Println(p0)
                                _ = fi18
                                _ = v29
                            }
                            if ((3.14159 + 0.0) <= -2.25) {
                                v36 := GenS3{F0: ("a" != p2), F1: 65535, F2: (p0 << (((p1 / (((p1 % ((p1) | 1))) | 1))) & 63))}
                                _ = v36
                            }
                            _ = fi17
                        }
                        for fi37 := 0; fi37 < 2; fi37++ {
                            for fi38 := 0; fi38 < 6; fi38++ {
                                switch (fi38 + fi37) {
                                case -984:
                                    if ((false && (!true)) || true) {
                                        v6 -= v6
                                        if false {
                                            switch int(p0) {
                                            case 42:
                                                fmt.Println(p0, fi37, p2)
                                            case -56:
                                                v5 = append(v5, -493)
                                            case 644:
                                                fmt.Println(p0)
                                            default:
                                                var v39 uint32 = p0
                                                _ = v39
                                            }
                                        }
                                        p2 = p2
                                    }
                                    var v40 int = 128
                                    v41 := GenS5{F0: (((p2 + p2) + (((p2 + "hello") + p2) + (p2 + (p2 + p2)))) + p2)}
                                    _ = v40
                                    _ = v41
                                case 4294967295:
                                    fmt.Println(fi37, v6)
                                case 1:
                                    if true {
                                        v6 += -2.25
                                    } else {
                                        var v42 int = (fi38 % ((fi38) | 1))
                                        switch (p2 + (p2 + "")) {
                                        case "a":
                                            var v43 int64 = 0
                                            _ = v43
                                        case "":
                                            var v44 bool = true
                                            _ = v44
                                        }
                                        var v45 int16 = 1
                                        _ = v42
                                        _ = v45
                                    }
                                }
                                _ = fi38
                            }
                            fmt.Println(fi37)
                            switch ("a" + (((("foo bar" + p2) + (p2 + p2)) + "hello") + p2)) {
                            case "\n\t":
                                p2 = p2
                            case "foo bar":
                                for fi46 := 0; fi46 < 16; fi46++ {
                                    fmt.Println(fi37, p0)
                                    if (134 < 0) {
                                        v6 = v6
                                        p1 += (p1 ^ p1)
                                        if (9223372036854775807 >= -9223372036854775808) {
                                            switch ("\n\t" + ((p2 + p2) + (p2 + p2))) {
                                            case "foo bar":
                                                v47 := GenS6{F0: 1, F1: true, F2: 65535}
                                                p0 += (p0 & p0)
                                                switch (fi37 / (((-(fi37))) | 1)) {
                                                case 0:
                                                    if false {
                                                        if true {
                                                            v5 = append(v5, 807)
                                                        } else {
                                                            var v48 float32 = -2.25
                                                            if true {
                                                                fmt.Println(p2)
                                                                if (!(2 < 4294967295)) {
                                                                    v49 := GenS3{F0: true, F1: 0, F2: p0}
                                                                    _ = v49
                                                                } else {
                                                                    v50 := GenS6{F0: (^(p0)), F1: false, F2: 32767}
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    if (!false) {
                                                                        var v51 int8 = -1
                                                                        switch -1 {
                                                                        case 0:
                                                                            v5 = append(v5, fi37)
                                                                        default:
                                                                            v52 := GenS5{F0: ((p2 + (p2 + "\n\t")) + p2)}
                                                                            _ = v52
                                                                        }
                                                                        p1 = (^(p1))
                                                                        v51 += (v51 * 0)
                                                                        p0 -= 1
                                                                        _ = v51
                                                                    } else {
                                                                        var v53 int32 = 1
                                                                        var v54 int16 = int16(p0)
                                                                        _ = v53
                                                                        _ = v54
                                                                    }
                                                                    _ = v50
                                                                }
                                                                switch (("\n\t" + p2) + p2) {
                                                                case "":
                                                                    v5 = append(v5, (fi37 - fi46))
                                                                case "hello":
                                                                    switch p2 {
                                                                    case "a":
                                                                        var v55 float64 = (-2.25 - -2.25)
                                                                        _ = v55
                                                                    case "\n\t":
                                                                        v56 := GenS7{F0: (fi46 * (^(fi46)))}
                                                                        _ = v56
                                                                    }
                                                                default:
                                                                    fmt.Println(p1)
                                                                }
                                                                fmt.Println(v6, fi46)
                                                                v57 := GenS8{F0: (p2 + (((p2 + "") + "") + (p2 + p2))), F1: 255, F2: 1}
                                                                _ = v57
                                                            }
                                                            v58 := GenS9{F0: 0, F1: 0}
                                                            _ = v48
                                                            _ = v58
                                                        }
                                                    } else {
                                                        p0 += (p0 ^ (^(p0)))
                                                        p0 -= (p0 * (^(p0)))
                                                        if (true || false) {
                                                            p1 -= (p1 ^ 952)
                                                            var v59 float64 = 3.14159
                                                            v60 := GenS10{F0: p2, F1: p1}
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        fmt.Println(fi37, v6, v6)
                                                        var v61 string = (p2 + "")
                                                        _ = v61
                                                    }
                                                }
                                                _ = v47
                                            default:
                                                v6 += 0.0
                                            }
                                        }
                                    } else {
                                        fmt.Println(v6, fi46)
                                        switch (p2 + ("foo bar" + ("hello" + "hello"))) {
                                        case "":
                                            p1 = (p1 * p1)
                                        case "foo bar":
                                            if len(v5) > 0 {
                                                v5[0] = 128
                                            }
                                        case "hello":
                                            if (int(p0) <= fi46) {
                                                p1 += 785
                                            } else {
                                                if (-6 == 0) {
                                                    var v62 float32 = v6
                                                    var v63 float32 = -2.25
                                                    _ = v62
                                                    _ = v63
                                                }
                                                p1 = 1
                                                p1 = p1
                                            }
                                        }
                                    }
                                    _ = fi46
                                }
                            case "hello":
                                fmt.Println(p0, fi37, p1)
                            }
                            _ = fi37
                        }
                        v64 := GenS3{F0: false, F1: uint16(p1), F2: p0}
                        _ = v64
                    }
                }
            case "foo bar":
                var v65 int = -2147483648
                _ = v65
            }
        }
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -32768
    var v3 uint64 = 2
    v4 := []uint32{263, 4294967295, 42}
    v3 = v3
    v3 += 4294967295
    for fi66 := 0; fi66 < 6; fi66++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v67 uint64 = (^(v3))
        _ = fi66
        _ = v67
    }
    if len(v4) > 0 {
        v4[0] = 2
    }
    fmt.Println(v2)
    var v68 int32 = 42
    v69 := GenS11{F0: 2}
    if true {
        for fi70 := 0; fi70 < 14; fi70++ {
            for fi71 := 0; fi71 < 6; fi71++ {
                switch (^(fi70)) {
                case -284:
                    switch (fi70 >> ((65535) & 63)) {
                    case -1:
                        if (255 < 18446744073709551615) {
                            var v72 int32 = v68
                            var v73 uint = 1
                            switch "foo bar" {
                            case "":
                                if (true || (830 > 687)) {
                                    if (false && ((v73 * (v73 % ((v73) | 1))) > 42)) {
                                        var v74 int32 = v68
                                        var v75 float32 = 3.14159
                                        _ = v74
                                        _ = v75
                                    } else {
                                        switch 32768 {
                                        case 737:
                                            fmt.Println(v3)
                                        }
                                        v68 -= (v72 & v72)
                                    }
                                    v76 := GenS12{F0: 65535, F1: "hello"}
                                    _ = v76
                                }
                            case "foo bar":
                                if false {
                                    if true {
                                        var v77 uint = v73
                                        _ = v77
                                    }
                                    v78 := v1[(^(fi71))]
                                    _ = v78
                                } else {
                                    var v79 float32 = 1e10
                                    _ = v79
                                }
                            case "hello":
                                v2 -= v2
                            default:
                                v68 = (v72 | -1)
                            }
                            _ = v72
                            _ = v73
                        } else {
                            v4 = append(v4, 65535)
                            if len(v4) > 0 {
                                v4[0] = 2
                            }
                            v1[42] = 2147483647
                        }
                    }
                case 4294967295:
                    fmt.Println(fi71)
                default:
                    v68 += v68
                }
                v2 = v2
                _ = fi71
            }
            _ = fi70
        }
        for fi80 := 0; fi80 < 15; fi80++ {
            fmt.Println(v2, v68, fi80)
            v2 = v2
            _ = fi80
        }
    }
    if ((124 == 65535) && true) {
        switch "" {
        case "":
            v81 := v1[42]
            _ = v81
        case "foo bar":
            var v82 int64 = 42
            _ = v82
        }
        v83 := GenS13{F0: 171, F1: 32768}
        _ = v83
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v84 := GenS9{F0: -21, F1: 1}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v84
    }
    v68 += (v68 % (((v68 - 158)) | 1))
    fmt.Println(v2)
    for fi85 := 0; fi85 < 6; fi85++ {
        var v86 uint64 = 0
        switch 42 {
        case 32768:
            _ = len(v1)
        case 128:
            v68 += 422
        case -1:
            v2 -= v2
        default:
            if (false && (-1 < -46)) {
                v87 := GenS3{F0: false, F1: 0, F2: 2147483647}
                v88 := GenS14{F0: ("a" + "hello")}
                _ = v87
                _ = v88
            }
        }
        _ = fi85
        _ = v86
    }
    _ = len(v4)
    var v89 float32 = 0.1
    v3 = v3
    switch 42 {
    case -2147483648:
        if (-2.25 > v89) {
            switch ("foo bar" + ("\n\t" + ("hello" + ""))) {
            case "":
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    v90 := GenS8{F0: "", F1: 4294967295, F2: 127}
                    v91 := GenS15{F0: v2}
                    _ = v90
                    _ = v91
                }
            case "\n\t":
                v89 = -2.25
            case "hello":
                v1[-1] = 466
            }
            switch -491 {
            case 128:
                var v92 int64 = 0
                _ = v92
            case 127:
                v3 = (v3 - v3)
            default:
                if (!(!(!(((v68 != (v68 + v68)) && (false || (!(uint(v3) == 1)))) || false)))) {
                    if (3.14159 <= float64(((v89 - 1.5) + v89))) {
                        switch 32767 {
                        case 128:
                            switch 239 {
                            case 0:
                                for fi93 := 0; fi93 < 7; fi93++ {
                                    fmt.Println(fi93)
                                    if (!(false || ((true && (!true)) || false))) {
                                        var v94 uint = 2
                                        if true {
                                            v95 := GenS14{F0: ("hello" + "a")}
                                            switch ("hello" + "hello") {
                                            case "a":
                                                _ = len(v1)
                                            case "foo bar":
                                                switch fi93 {
                                                case -1:
                                                    if true {
                                                        var v96 uint64 = v3
                                                        for fi97 := 0; fi97 < 6; fi97++ {
                                                            switch (fi93 & fi97) {
                                                            case 42:
                                                                v89 = v89
                                                                fmt.Println(v3)
                                                            default:
                                                                switch fi93 {
                                                                case 223:
                                                                    var v98 int16 = (v2 * -1)
                                                                    _ = v98
                                                                case -837:
                                                                    var v99 int = fi97
                                                                    _ = v99
                                                                default:
                                                                    v100 := v1[fi93]
                                                                    _ = v100
                                                                }
                                                            }
                                                            fmt.Println(v68, fi97, fi93)
                                                            _ = fi97
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v96
                                                    }
                                                }
                                            case "\n\t":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            default:
                                                v101 := GenS14{F0: "\n\t"}
                                                _ = v101
                                            }
                                            var v102 int64 = 9223372036854775807
                                            _ = v102
                                            _ = v95
                                        } else {
                                            var v103 uint8 = 127
                                            _ = v103
                                        }
                                        for fi104 := 0; fi104 < 15; fi104++ {
                                            if (!false) {
                                                v89 -= float32(3.14159)
                                                v2 = v2
                                                var v105 int = (fi93 | fi93)
                                                v2 = v2
                                                switch ("\n\t" + "") {
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "":
                                                        var v106 int16 = (-(v2))
                                                        _ = v106
                                                    default:
                                                        var v107 int32 = v68
                                                        _ = v107
                                                    }
                                                }
                                                var v108 int8 = 0
                                                _ = v105
                                                _ = v108
                                            }
                                            var v109 string = ("foo bar" + "a")
                                            if true {
                                                switch -1 {
                                                case 506:
                                                    switch ((v109 + "foo bar") + (v109 + v109)) {
                                                    case "":
                                                        fmt.Println(v2, fi93)
                                                        var v110 uint16 = 1
                                                        _ = v110
                                                    }
                                                case 127:
                                                    var v111 bool = (0 > 1)
                                                    v1[fi104] = 1
                                                    _ = v111
                                                default:
                                                    v112 := GenS3{F0: false, F1: 676, F2: uint32(v3)}
                                                    _ = v112
                                                }
                                            }
                                            _ = fi104
                                            _ = v109
                                        }
                                        _ = v94
                                    } else {
                                        v3 += v3
                                        var v113 float64 = (-2.25 + (((3.14159 * ((1e10 / float64((v89 * 0.1))) / 1.0)) - float64(3.14159)) / (-2.25 + -2.25)))
                                        switch fi93 {
                                        case -516:
                                            v3 += 1
                                        case 42:
                                            if (4294967295 < 227) {
                                                v114 := GenS13{F0: 2, F1: -1}
                                                _ = v114
                                            }
                                        default:
                                            for fi115 := 0; fi115 < 10; fi115++ {
                                                var v116 uint = 4294967295
                                                v89 = v89
                                                _ = fi115
                                                _ = v116
                                            }
                                        }
                                        _ = v113
                                    }
                                    _ = fi93
                                }
                            }
                        default:
                            v117 := GenS15{F0: v2}
                            _ = v117
                        }
                        if false {
                            v118 := GenS14{F0: "foo bar"}
                            _ = v118
                        } else {
                            for fi119 := 0; fi119 < 4; fi119++ {
                                var v120 uint16 = 32767
                                _ = fi119
                                _ = v120
                            }
                            var v121 string = "hello"
                            _ = v121
                        }
                    } else {
                        var v122 float64 = 3.14159
                        v123 := GenS16{F0: 84, F1: 17}
                        if (-3 > 0) {
                            fmt.Println(v122, v89, v2)
                            var v124 bool = (true || ((!false) && (-9223372036854775808 != 42)))
                            _ = v124
                        }
                        _ = v122
                        _ = v123
                    }
                    var v125 int = 1
                    _ = v125
                } else {
                    for fi126 := 0; fi126 < 6; fi126++ {
                        if false {
                            v89 = v89
                            for fi127 := 0; fi127 < 20; fi127++ {
                                var v128 uint = 42
                                _ = fi127
                                _ = v128
                            }
                            var v129 uint64 = v3
                            _ = v129
                        }
                        _ = fi126
                    }
                }
            }
        } else {
            var v130 int32 = 1
            var v131 uint16 = 1
            fmt.Println(v2, v3)
            _ = v130
            _ = v131
        }
    case -9223372036854775808:
        switch -253 {
        case -1:
            fmt.Println(v2, v2)
        case 766:
            _ = len(v4)
        }
    case 51:
        if true {
            var v132 int64 = -9223372036854775808
            _ = v132
        }
    default:
        for fi133 := 0; fi133 < 16; fi133++ {
            v3 = 0
            if (-9223372036854775808 != (fi133 / ((fi133) | 1))) {
                delete(v1, fi133)
            }
            _ = fi133
        }
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v3, v3, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v68:", v68)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v69
    _ = v89
}
