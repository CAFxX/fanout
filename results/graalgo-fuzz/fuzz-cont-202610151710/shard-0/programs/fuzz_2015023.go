package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint; F1 int; F2 string }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    var v3 float32 = 1.5
    var v4 int32 = 917
    var v5 string = ("hello" + "hello")
    v4 = (v4 >> ((427) & 63))
    v4 = v4
    for fi6 := 0; fi6 < 5; fi6++ {
        if false {
            var v7 int = -191
            switch "foo bar" {
            case "a":
                v8 := GenS1{F0: 42, F1: -182}
                _ = v8
            case "":
                var v9 int64 = -673
                _ = v9
            case "hello":
                v4 += (v4 << ((255) & 63))
            }
            v10 := GenS2{F0: 2, F1: -2147483648}
            _ = v10
            _ = v7
        }
        v1[""] = true
        if false {
            v11 := GenS3{F0: 42}
            var v12 int8 = 0
            var v13 float32 = 3.14159
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi6
    }
    var v14 uint64 = 59
    v15 := GenS1{F0: 42, F1: -171}
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint = 410
            _ = fi17
            _ = v18
        }
        v19 := GenS3{F0: 42}
        for fi20 := 0; fi20 < 7; fi20++ {
            v1["hello"] = (false && (!(v14 == (v14 * 1))))
            if ((v4 <= v4) && (18446744073709551615 == (^(v14)))) {
                var v21 float32 = (v3 - float32((3.14159 / -2.25)))
                if ((fi16 - (fi20 ^ fi20)) <= fi16) {
                    var v22 bool = (249 > 0)
                    var v23 uint32 = 42
                    _ = v22
                    _ = v23
                }
                _ = v21
            } else {
                var v24 float64 = (1e10 * 3.14159)
                _ = v24
            }
            _ = fi20
        }
        _ = fi16
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = ((v5 + v5) + ("a" + v5))
    }
    if true {
        v3 = (1.5 - 1.5)
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 float64 = -2.25
            switch fi25 {
            case 127:
                v26 = 0.1
            case 477:
                v14 = 18446744073709551615
            }
            v27 := v1[(((("a" + v5) + ((v5 + (v5 + v5)) + v5)) + ((v5 + v5) + v5)) + v5)]
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            v4 -= (^(v4))
            fmt.Println(v3, v3)
            _ = fi29
        }
        v30 := GenS4{F0: (v4 & v4), F1: v14}
        for fi31 := 0; fi31 < 12; fi31++ {
            v3 = v3
            _ = fi31
        }
        _ = fi28
        _ = v30
    }
    if false {
        if true {
            if true {
                for fi32 := 0; fi32 < 5; fi32++ {
                    v33 := GenS1{F0: 1, F1: 42}
                    _ = fi32
                    _ = v33
                }
            }
            if (3.14159 > -2.25) {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v14 = v14
                    fmt.Println(v14, v4, v4)
                }
                if len(v2) > 0 {
                    v2[0] = ((v5 + (v5 + (v5 + v5))) + ("\n\t" + "hello"))
                }
                _ = len(v2)
            } else {
                v4 += v4
                if (!(!false)) {
                    var v34 int8 = -1
                    v3 += v3
                    var v35 int32 = v4
                    _ = v34
                    _ = v35
                }
            }
        }
    } else {
        if len(v2) > 0 {
            v2[0] = v5
        }
        if false {
            var v36 int16 = 1
            for fi37 := 0; fi37 < 11; fi37++ {
                for fi38 := 0; fi38 < 12; fi38++ {
                    switch (-(fi38)) {
                    case 0:
                        v39 := GenS5{F0: (v14 ^ v14), F1: true, F2: -52}
                        _ = v39
                    case 4294967295:
                        if (240 != 65535) {
                            var v40 int16 = v36
                            if false {
                                v2 = append(v2, ("" + v5))
                                if false {
                                    switch ("\n\t" + (v5 + ("a" + v5))) {
                                    case "":
                                        if (("hello" + v5) <= "hello") {
                                            v41 := GenS6{F0: 2}
                                            _ = v41
                                        }
                                        v42 := GenS4{F0: (v4 & (v4 ^ v4)), F1: v14}
                                        switch fi37 {
                                        case -976:
                                            var v43 int16 = (v40 ^ (-(v36)))
                                            _ = v43
                                        case 127:
                                            if false {
                                                switch ("" + (("\n\t" + "\n\t") + "foo bar")) {
                                                case "":
                                                    if (!true) {
                                                        fmt.Println(fi38, v5)
                                                        v36 += 890
                                                        if false {
                                                            _ = len(v2)
                                                            v44 := GenS7{F0: v4, F1: v4}
                                                            v4 += v4
                                                            v45 := v1[(v5 + (v5 + "hello"))]
                                                            v5 = (v5 + "\n\t")
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            v5 = "\n\t"
                                                        }
                                                        v1[((v5 + (v5 + ("foo bar" + (v5 + v5)))) + "a")] = (2 >= 142)
                                                        if (!(!true)) {
                                                            fmt.Println(v36, fi38)
                                                        }
                                                        v46 := GenS1{F0: 100, F1: 359}
                                                        _ = v46
                                                    }
                                                case "foo bar":
                                                    var v47 bool = (true || false)
                                                    _ = v47
                                                case "a":
                                                    if false {
                                                        switch (fi37 & fi37) {
                                                        case 32767:
                                                            if ((fi38 % ((0) | 1)) > (fi38 * (-(fi38)))) {
                                                                if (((v5 + v5) + v5) != "\n\t") {
                                                                    var v48 int = 2147483647
                                                                    _ = v48
                                                                }
                                                                var v49 int16 = 271
                                                                var v50 uint64 = (v14 - (v14 / ((66) | 1)))
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                var v51 bool = false
                                                                _ = v51
                                                            }
                                                        }
                                                        v52 := GenS6{F0: 13}
                                                        var v53 uint32 = 65535
                                                        fmt.Println(fi37)
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        var v54 string = v5
                                                        if (true || (-237 == 9223372036854775807)) {
                                                            var v55 int16 = (v36 * 32767)
                                                            _ = v55
                                                        } else {
                                                            v14 = (v14 - v14)
                                                        }
                                                        _ = v54
                                                    }
                                                default:
                                                    _ = len(v2)
                                                }
                                            } else {
                                                v56 := GenS7{F0: v4, F1: (-(v4))}
                                                v57 := GenS3{F0: 128}
                                                _ = len(v1)
                                                switch (fi37 / ((42) | 1)) {
                                                case 1:
                                                    _ = len(v2)
                                                case 0:
                                                    v58 := GenS8{F0: 42}
                                                    _ = v58
                                                default:
                                                    var v59 uint32 = 748
                                                    _ = v59
                                                }
                                                v14 += (^(v14))
                                                _ = v56
                                                _ = v57
                                            }
                                        case 32768:
                                            v2 = append(v2, ((v5 + v5) + v5))
                                        }
                                        _ = v42
                                    }
                                    v5 = "hello"
                                }
                                v1[v5] = false
                            } else {
                                switch 709 {
                                case 42:
                                    var v60 int = 4294967295
                                    _ = v60
                                case 127:
                                    switch "a" {
                                    case "":
                                        v36 = v36
                                    }
                                default:
                                    var v61 uint = 241
                                    _ = v61
                                }
                                fmt.Println(v5, v5, v4)
                                switch ((v5 + "foo bar") + "hello") {
                                case "foo bar":
                                    switch (("\n\t" + v5) + ((v5 + v5) + ((v5 + v5) + v5))) {
                                    case "\n\t":
                                        v62 := GenS6{F0: 2}
                                        _ = v62
                                    case "hello":
                                        var v63 bool = (!(v14 >= (v14 / (((v14 - 24)) | 1))))
                                        var v64 float32 = (float32(((float64(v3) * 0.0) * ((0.1 * (1e10 + ((float64((v3 / 3.14159)) + (0.1 - (1e10 * 1e10))) - (0.1 + 0.1)))) + 0.1))) - v3)
                                        _ = v63
                                        _ = v64
                                    case "":
                                        switch fi37 {
                                        case 127:
                                            v14 += 18446744073709551615
                                        }
                                    default:
                                        if (2147483647 >= 1) {
                                            switch (("hello" + v5) + (v5 + "hello")) {
                                            case "foo bar":
                                                v65 := GenS9{F0: 42}
                                                _ = v65
                                            case "a":
                                                v36 += v36
                                            default:
                                                fmt.Println(v3, fi37)
                                            }
                                            v66 := GenS1{F0: 70, F1: 1}
                                            _ = len(v1)
                                            if len(v2) > 0 {
                                                v2[0] = v5
                                            }
                                            fmt.Println(v4, v3)
                                            _ = v66
                                        }
                                    }
                                }
                                v40 = v36
                                switch 32768 {
                                case -2147483648:
                                    switch fi37 {
                                    case 149:
                                        var v67 int = -943
                                        _ = v67
                                    default:
                                        v5 = ("hello" + "foo bar")
                                    }
                                default:
                                    if true {
                                        v36 = 42
                                        if false {
                                            switch v5 {
                                            case "":
                                                switch (fi37 + (^(fi38))) {
                                                case 9223372036854775807:
                                                    if ((48 <= 1) && true) {
                                                        var v68 int8 = -126
                                                        _ = v68
                                                    } else {
                                                        v1[v5] = true
                                                    }
                                                default:
                                                    var v69 uint32 = 2
                                                    _ = v69
                                                }
                                            case "foo bar":
                                                v3 += 1e10
                                            case "hello":
                                                switch ((("foo bar" + (((v5 + (("hello" + "a") + v5)) + v5) + (v5 + v5))) + v5) + "a") {
                                                case "hello":
                                                    var v70 int8 = 19
                                                    _ = v70
                                                case "foo bar":
                                                    v3 = (1e10 + v3)
                                                default:
                                                    fmt.Println(v40)
                                                }
                                            }
                                            v14 += (v14 & v14)
                                            v14 = (v14 * (v14 >> ((42) & 63)))
                                        }
                                    }
                                }
                            }
                            switch fi38 {
                            case 9223372036854775807:
                                v36 = (v40 & -32768)
                            case 127:
                                if (true && false) {
                                    v71 := GenS4{F0: v4, F1: 0}
                                    var v72 int8 = 42
                                    var v73 int = fi38
                                    switch fi37 {
                                    case 1:
                                        switch v5 {
                                        case "foo bar":
                                            v5 = v5
                                        case "hello":
                                            fmt.Println(v14, v4)
                                        default:
                                            _ = len(v2)
                                        }
                                        fmt.Println(v40, v3, v36)
                                        var v74 string = ""
                                        _ = v74
                                    default:
                                        _ = len(v1)
                                    }
                                    switch ((("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) + v5) + v5) {
                                    case "hello":
                                        switch (fi38 - 32768) {
                                        case 1:
                                            fmt.Println(v73, v73, v36)
                                        case 42:
                                            fmt.Println(v3)
                                        case -701:
                                            if true {
                                                v75 := v1[(v5 + (v5 + "foo bar"))]
                                                _ = v75
                                            }
                                        }
                                        fmt.Println(v36, v3, v5)
                                    case "\n\t":
                                        v72 = v72
                                    }
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                            }
                            if (127 <= 0) {
                                fmt.Println(v40, v14)
                                v76 := v1[""]
                                var v77 int16 = int16(v4)
                                _ = v76
                                _ = v77
                            }
                            _ = v40
                        } else {
                            if (127 != -28) {
                                var v78 int = (fi38 / ((fi38) | 1))
                                _ = v78
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch 42 {
                        case 1:
                            switch 2147483647 {
                            case -350:
                                switch ((v5 + v5) + v5) {
                                case "foo bar":
                                    v3 = v3
                                case "":
                                    if false {
                                        var v79 int16 = (v36 / ((-1) | 1))
                                        if true {
                                            var v80 uint8 = 128
                                            _ = v80
                                        } else {
                                            switch fi38 {
                                            case 9223372036854775807:
                                                if false {
                                                    fmt.Println(v4, v36, v3)
                                                }
                                            case 127:
                                                switch v5 {
                                                case "foo bar":
                                                    v2 = append(v2, v5)
                                                case "\n\t":
                                                    switch v5 {
                                                    case "":
                                                        v36 = int16(v4)
                                                    }
                                                case "hello":
                                                    v79 = (v79 * 943)
                                                default:
                                                    v1[v5] = false
                                                }
                                            case 417:
                                                if (true || true) {
                                                    fmt.Println(v3, v5)
                                                    v81 := GenS5{F0: v14, F1: ((v36 & -1) != (v36 % (((v36 & v36)) | 1))), F2: (fi37 ^ (fi38 % (((fi37 ^ 573)) | 1)))}
                                                    _ = v81
                                                }
                                            default:
                                                v82 := GenS5{F0: 177, F1: true, F2: -42}
                                                _ = v82
                                            }
                                            fmt.Println(v3, fi37, v14)
                                            delete(v1, "\n\t")
                                        }
                                        var v83 int64 = -903
                                        v83 = v83
                                        _ = v79
                                        _ = v83
                                    } else {
                                        v4 += -1
                                        v36 += (v36 / ((32767) | 1))
                                        fmt.Println(fi38)
                                    }
                                case "\n\t":
                                    switch (fi37 >> ((4294967295) & 63)) {
                                    case 1:
                                        var v84 int8 = 0
                                        _ = v84
                                    case 32768:
                                        var v85 int = (fi37 << ((768) & 63))
                                        _ = v85
                                    case 42:
                                        v5 = (("a" + v5) + v5)
                                    }
                                }
                                v3 = 1e10
                            case 127:
                                var v86 int32 = 90
                                var v87 int16 = 1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v86
                                _ = v87
                            case 2147483647:
                                fmt.Println(v36)
                            }
                        }
                        fmt.Println(fi37, v14)
                    }
                    _ = fi38
                }
                _ = fi37
            }
            _ = v36
        } else {
            var v88 uint64 = (^(v14))
            v89 := GenS5{F0: v14, F1: (-204 <= 42), F2: int(v4)}
            _ = v88
            _ = v89
        }
        for fi90 := 0; fi90 < 18; fi90++ {
            if (false && true) {
                if (42 == 508) {
                    fmt.Println(v4, v5)
                    v5 = "foo bar"
                    for fi91 := 0; fi91 < 15; fi91++ {
                        var v92 int8 = 100
                        v5 = "a"
                        v92 -= 98
                        _ = fi91
                        _ = v92
                    }
                } else {
                    for fi93 := 0; fi93 < 1; fi93++ {
                        v94 := GenS10{F0: uint8(v14)}
                        _ = fi93
                        _ = v94
                    }
                }
            }
            for fi95 := 0; fi95 < 11; fi95++ {
                if false {
                    if (!true) {
                        v96 := GenS11{F0: 65535, F1: false}
                        if true {
                            if false {
                                v5 = "foo bar"
                            }
                            v5 = (v5 + "a")
                        } else {
                            v1[v5] = (true || false)
                            var v97 string = ((v5 + "foo bar") + "hello")
                            _ = v97
                        }
                        var v98 float32 = (v3 / 3.14159)
                        _ = v96
                        _ = v98
                    } else {
                        var v99 int8 = int8(v4)
                        if true {
                            switch (v5 + "hello") {
                            case "a":
                                if (v5 <= "\n\t") {
                                    _ = len(v2)
                                    if true {
                                        v14 += 1
                                        v1[v5] = false
                                    } else {
                                        var v100 uint32 = 42
                                        _ = v100
                                    }
                                    delete(v1, v5)
                                } else {
                                    v101 := GenS1{F0: v99, F1: -1}
                                    v1[""] = (true || (!((v14 / (((v14 << ((928) & 63))) | 1)) < v14)))
                                    _ = v101
                                }
                            }
                            var v102 int = (fi90 % ((fi90) | 1))
                            var v103 int16 = 0
                            _ = v102
                            _ = v103
                        }
                        if (v3 > 1e10) {
                            var v104 float32 = v3
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v104
                        } else {
                            var v105 int = fi95
                            var v106 uint8 = 81
                            v107 := GenS7{F0: (v4 % (((v4 - 0)) | 1)), F1: 132}
                            var v108 int16 = -624
                            _ = len(v1)
                            if (!false) {
                                var v109 uint64 = (v14 << ((890) & 63))
                                switch ("a" + v5) {
                                case "a":
                                    switch ((v5 + "\n\t") + v5) {
                                    case "a":
                                        var v110 uint64 = 2
                                        _ = v110
                                    case "hello":
                                        var v111 int = (-(fi90))
                                        _ = v111
                                    case "foo bar":
                                        if false {
                                            v108 = 1000
                                            v1["hello"] = ((1 >= -1) || false)
                                        }
                                    default:
                                        if (true && true) {
                                            var v112 float32 = v3
                                            switch (fi90 / ((int(v14)) | 1)) {
                                            case -2147483648:
                                                fmt.Println(v108, v4, v5)
                                            }
                                            _ = v112
                                        }
                                    }
                                default:
                                    var v113 uint16 = 32767
                                    _ = v113
                                }
                                switch v5 {
                                case "a":
                                    v114 := GenS12{F0: 679, F1: (v105 / ((-470) | 1)), F2: v5}
                                    _ = v114
                                case "\n\t":
                                    var v115 int16 = v108
                                    _ = v115
                                default:
                                    var v116 float64 = (0.0 - (0.1 * 0.1))
                                    _ = v116
                                }
                                _ = v109
                            } else {
                                fmt.Println(v108, v14)
                                v117 := GenS6{F0: (v106 << ((214) & 63))}
                                v14 -= (^(v14))
                                v105 -= (v105 * 920)
                                if ("hello" != "\n\t") {
                                    fmt.Println(v108)
                                    fmt.Println(v99)
                                    v2 = append(v2, v5)
                                } else {
                                    v118 := GenS2{F0: 562, F1: (^(v105))}
                                    var v119 int32 = (v4 << ((42) & 63))
                                    _ = v118
                                    _ = v119
                                }
                                _ = v117
                            }
                            v120 := GenS12{F0: 42, F1: 9223372036854775807, F2: v5}
                            _ = v105
                            _ = v106
                            _ = v107
                            _ = v108
                            _ = v120
                        }
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v121 float64 = ((float64(-2.25) / -2.25) + 0.1)
                            var v122 uint16 = 744
                            var v123 float64 = -2.25
                            _ = v121
                            _ = v122
                            _ = v123
                        } else {
                            fmt.Println(fi95, v3)
                            var v124 bool = (-2.25 < 0.0)
                            _ = v124
                        }
                        _ = v99
                    }
                    v5 = v5
                    var v125 float32 = v3
                    _ = v125
                }
                _ = fi95
            }
            for fi126 := 0; fi126 < 6; fi126++ {
                if false {
                    var v127 uint = 4294967295
                    _ = v127
                } else {
                    fmt.Println(v5, fi90)
                    var v128 int32 = (^(v4))
                    _ = v128
                }
                delete(v1, "\n\t")
                if ((-(v4)) < 480) {
                    v14 = (^(v14))
                    var v129 string = (v5 + ("foo bar" + ((v5 + v5) + v5)))
                    _ = v129
                } else {
                    var v130 int32 = 322
                    _ = v130
                }
                _ = fi126
            }
            _ = fi90
        }
    }
    _ = len(v2)
    var v131 uint64 = (^(v14))
    if (1 < 42) {
        switch "hello" {
        case "foo bar":
            v3 += float32(((3.14159 / float64(1.5)) * ((float64(float32(1.5)) - 0.0) * (float64(v3) - (0.0 * 0.1)))))
        case "\n\t":
            var v132 uint64 = 4294967295
            _ = v132
        default:
            for fi133 := 0; fi133 < 11; fi133++ {
                v131 = (v14 - 274)
                _ = fi133
            }
        }
        var v134 bool = true
        var v135 int = 128
        _ = v134
        _ = v135
    }
    v14 += v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v131:", v131)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v131
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
