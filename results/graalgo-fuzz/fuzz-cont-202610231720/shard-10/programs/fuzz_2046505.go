package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64; F1 int8; F2 bool }
type GenS10 struct { F0 string; F1 uint8; F2 int16 }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 uint16; F1 uint64; F2 string }
type GenS13 struct { F0 uint16; F1 uint64 }
type GenS14 struct { F0 uint64; F1 uint16 }
type GenS15 struct { F0 uint64; F1 uint; F2 int32 }
type GenS16 struct { F0 int64; F1 int16 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int32; F1 int32 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint32; F1 int64 }
type GenS21 struct { F0 int16; F1 string; F2 int8 }
func genFn0() bool {
    var v6 int = 222
    v6 += v6
    var v7 uint16 = 255
    var v8 uint64 = 497
    _ = v6
    _ = v7
    _ = v8
    return ((428 != 42) || (255 < 0))
}
func genFn1(p0 uint32) int16 {
    for fi9 := 0; fi9 < 15; fi9++ {
        p0 += p0
        fmt.Println(p0)
        _ = fi9
    }
    v10 := []int{}
    var v11 uint16 = 0
    v12 := GenS1{F0: 173}
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: 160}
        v11 += (v11 / ((2) | 1))
        _ = len(v10)
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 2
    var v3 bool = false
    v4 := make(map[string]int)
    v5 := []uint32{42, 2147483648}
    switch "\n\t" {
    case "hello":
        v15 := GenS2{F0: 63, F1: -406, F2: 0}
        _ = v15
    case "\n\t":
        fmt.Println(v1, v3)
    case "foo bar":
        if v3 {
            for fi16 := 0; fi16 < 1; fi16++ {
                v2 += 2
                v17 := GenS3{F0: (v1 * 383)}
                _ = fi16
                _ = v17
            }
            if true {
                v3 = true
            }
            if false {
                for fi18 := 0; fi18 < 5; fi18++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    delete(v4, ("a" + "hello"))
                    switch "" {
                    case "":
                        v19 := GenS2{F0: -128, F1: 678, F2: fi18}
                        _ = v19
                    }
                    _ = fi18
                }
                var v20 bool = (v3 || (!(false || false)))
                _ = v20
            } else {
                v1 = (v1 / (((v1 - v1)) | 1))
                fmt.Println(v1, v2)
                switch ("" + "a") {
                case "foo bar":
                    v21 := GenS1{F0: 2}
                    _ = v21
                case "":
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if (v3 && true) {
                            fmt.Println(fi22, v1, v3)
                            _ = len(v4)
                            if true {
                                var v23 int16 = 32767
                                var v24 float32 = -2.25
                                var v25 int8 = 100
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        }
                        _ = fi22
                    }
                case "\n\t":
                    var v26 int16 = 1
                    _ = v26
                default:
                    var v27 uint64 = v2
                    _ = v27
                }
            }
        }
    }
    v3 = ((32767 >= 127) || (42 != 32767))
    if false {
        v28 := GenS1{F0: 127}
        _ = v28
    } else {
        for fi29 := 0; fi29 < 20; fi29++ {
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := GenS4{F0: 42}
                var v32 int64 = -9223372036854775808
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = len(v4)
            _ = fi29
        }
        if (!false) {
            var v33 int32 = 519
            v33 -= 1
            if v3 {
                for fi34 := 0; fi34 < 1; fi34++ {
                    v35 := GenS4{F0: 1}
                    _ = fi34
                    _ = v35
                }
            } else {
                if v3 {
                    fmt.Println(v3, v1, v3)
                    for fi36 := 0; fi36 < 18; fi36++ {
                        var v37 uint16 = 155
                        if v3 {
                            v37 -= v37
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v2 -= (v2 % ((v2) | 1))
                        }
                        _ = fi36
                        _ = v37
                    }
                } else {
                    fmt.Println(v1)
                    v1 += v1
                    _ = len(v5)
                }
                if true {
                    for fi38 := 0; fi38 < 5; fi38++ {
                        for fi39 := 0; fi39 < 5; fi39++ {
                            switch "" {
                            case "":
                                var v40 uint8 = 58
                                _ = v40
                            case "hello":
                                if (!(127 >= 2)) {
                                    v33 += (v33 ^ (v33 >> (((v1 ^ (^(v1)))) & 63)))
                                    switch "\n\t" {
                                    case "a":
                                        v41 := GenS5{F0: -1}
                                        _ = v41
                                    case "\n\t":
                                        v33 -= v33
                                    case "foo bar":
                                        var v42 string = "hello"
                                        _ = v42
                                    default:
                                        v33 += v33
                                    }
                                    v5 = append(v5, 928)
                                }
                            default:
                                v43 := GenS3{F0: (^(v1))}
                                _ = v43
                            }
                            v44 := GenS6{F0: (v3 && ((-32768 != -179) && (v1 < 359)))}
                            fmt.Println(fi38, fi39, v2)
                            v33 -= (^(v33))
                            _ = fi39
                            _ = v44
                        }
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "\n\t":
                            if v3 {
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    v45 := GenS7{F0: 100}
                                    _ = v45
                                case "hello":
                                    fmt.Println(fi38, fi38, v3)
                                default:
                                    if (-27 == -128) {
                                        if (815 <= 255) {
                                            switch "" {
                                            case "\n\t":
                                                v2 += v2
                                            }
                                        }
                                    }
                                }
                                var v46 int8 = -128
                                _ = v46
                            }
                        default:
                            for fi47 := 0; fi47 < 3; fi47++ {
                                var v48 int8 = 127
                                var v49 uint32 = 921
                                var v50 uint16 = 42
                                _ = fi47
                                _ = v48
                                _ = v49
                                _ = v50
                            }
                        }
                        _ = fi38
                    }
                    if true {
                        for fi51 := 0; fi51 < 1; fi51++ {
                            switch (fi51 & fi51) {
                            case 1:
                                v33 = v33
                            }
                            v52 := GenS8{F0: (v1 ^ (^(v1)))}
                            fmt.Println(v3)
                            _ = fi51
                            _ = v52
                        }
                        v3 = (11 >= 95)
                        v53 := GenS5{F0: -1}
                        _ = v53
                    }
                } else {
                    if true {
                        v3 = (!(v2 >= (v2 >> ((v1) & 63))))
                        v54 := GenS2{F0: 100, F1: 0, F2: 0}
                        v55 := GenS8{F0: (v1 % (((v1 >> ((v1) & 63))) | 1))}
                        _ = v54
                        _ = v55
                    }
                    v56 := GenS9{F0: 18446744073709551615, F1: 127, F2: v3}
                    if (int64(v2) >= 1) {
                        for fi57 := 0; fi57 < 2; fi57++ {
                            v5 = append(v5, 160)
                            for fi58 := 0; fi58 < 6; fi58++ {
                                var v59 uint32 = 2
                                _ = fi58
                                _ = v59
                            }
                            _ = fi57
                        }
                        fmt.Println(v3)
                    } else {
                        switch (("hello" + "a") + "hello") {
                        case "a":
                            var v60 int16 = genFn1(217)
                            _ = v60
                        default:
                            switch int(v1) {
                            case -756:
                                v33 += v33
                            }
                        }
                        if false {
                            fmt.Println(v3)
                            v61 := v4["\n\t"]
                            fmt.Println(v3, v2, v33)
                            _ = v61
                        }
                        v1 -= v1
                    }
                    _ = v56
                }
                var v62 uint32 = 606
                _ = v62
            }
            _ = v33
        }
        v63 := GenS8{F0: (v1 | v1)}
        _ = v63
    }
    if v3 {
        switch 1 {
        case -2147483648:
            fmt.Println(v1, v2)
        }
        v64 := v4["a"]
        _ = v64
    }
    var v65 int8 = -1
    v2 = 1
    var v66 int32 = -521
    if false {
        switch "foo bar" {
        case "foo bar":
            v2 -= 42
        default:
            v2 -= (v2 << ((42) & 63))
        }
        switch 32767 {
        case 9223372036854775807:
            v67 := GenS2{F0: (v65 & (v65 / ((v65) | 1))), F1: 0, F2: -1}
            _ = v67
        }
    } else {
        v68 := GenS7{F0: v65}
        _ = v68
    }
    v69 := GenS10{F0: "hello", F1: 1, F2: genFn1(65535)}
    var v70 int32 = v66
    v1 += (v1 - (v1 + v1))
    switch ("foo bar" + "hello") {
    case "hello":
        if true {
            v71 := GenS8{F0: 1}
            if (((("" + "a") + "") + ("foo bar" + "")) <= "") {
                for fi72 := 0; fi72 < 15; fi72++ {
                    for fi73 := 0; fi73 < 16; fi73++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v74 uint = (v1 / ((v1) | 1))
                            _ = v74
                        case "hello":
                            if false {
                                v1 = (^(v1))
                                if (!v3) {
                                    v2 += (v2 + (v2 & v2))
                                    v3 = true
                                    if (65535 == 42) {
                                        switch (fi72 + fi72) {
                                        case 42:
                                            switch "hello" {
                                            case "foo bar":
                                                var v75 uint64 = (v2 / ((v2) | 1))
                                                switch "" {
                                                case "foo bar":
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v76 bool = true
                                                        _ = v76
                                                    case "":
                                                        if v3 {
                                                            if (65535 == 65535) {
                                                                if (42 <= -1) {
                                                                    if (true && false) {
                                                                        switch ("hello" + ("" + ("" + ("" + "")))) {
                                                                        case "foo bar":
                                                                            var v77 string = ("" + "")
                                                                            _ = v77
                                                                        case "":
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                var v78 bool = false
                                                                                _ = v78
                                                                            case "":
                                                                                v2 -= (v2 * v2)
                                                                                if v3 {
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    var v79 string = ("" + "foo bar")
                                                                                    _ = v79
                                                                                } else {
                                                                                    v3 = v3
                                                                                }
                                                                            }
                                                                            if (!(!(0.0 >= ((-2.25 - float64((-2.25 - 1e10))) - (1.5 / 1.5))))) {
                                                                                if true {
                                                                                    var v80 uint = (v1 % (((v1 + (v1 * 255))) | 1))
                                                                                    _ = v80
                                                                                }
                                                                                v81 := v4[((("\n\t" + "foo bar") + "\n\t") + "foo bar")]
                                                                                _ = v81
                                                                            } else {
                                                                                if (v3 || (!(42 > -762))) {
                                                                                    v82 := GenS10{F0: (("a" + "\n\t") + ""), F1: 1, F2: genFn1(2)}
                                                                                    if (v65 <= (^(v65))) {
                                                                                        switch (fi72 | fi73) {
                                                                                        case 4294967295:
                                                                                            var v83 int64 = -9223372036854775808
                                                                                            _ = v83
                                                                                        case 127:
                                                                                            if ((uint32(v1) >= 3) && v3) {
                                                                                                fmt.Println(fi73)
                                                                                                var v84 uint8 = 1
                                                                                                _ = v84
                                                                                            }
                                                                                        default:
                                                                                            v70 -= (v70 << ((uint(v2)) & 63))
                                                                                        }
                                                                                        var v85 float64 = 0.0
                                                                                        _ = v85
                                                                                    } else {
                                                                                        var v86 uint32 = 0
                                                                                        var v87 uint = (v1 >> (((v1 >> (((v1 * v1)) & 63))) & 63))
                                                                                        var v88 bool = v3
                                                                                        var v89 uint8 = 42
                                                                                        fmt.Println(v75, v86)
                                                                                        _ = v86
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                    }
                                                                                    var v90 uint32 = 42
                                                                                    _ = v82
                                                                                    _ = v90
                                                                                } else {
                                                                                    v2 = v2
                                                                                }
                                                                                v91 := GenS7{F0: (v65 - 127)}
                                                                                var v92 bool = ((!(!((v65 % ((-48) | 1)) > 15))) || (!(!(v3 || true))))
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                            v75 = (^(v75))
                                                                        case "\n\t":
                                                                            delete(v4, ((("\n\t" + "hello") + ("foo bar" + "hello")) + "hello"))
                                                                        default:
                                                                            if (v3 || ((((((-2.25 / 1.0) * 0.1) + 3.14159) - 0.1) + (1e10 / 1.5)) == 1.5)) {
                                                                                var v93 uint8 = 57
                                                                                _ = v93
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch ("a" + "hello") {
                                                                    case "hello":
                                                                        if (("a" >= "") && v3) {
                                                                            v94 := GenS11{F0: (^(v1)), F1: 128}
                                                                            _ = v94
                                                                        }
                                                                    case "\n\t":
                                                                        v4[""] = (fi73 % (((fi73 / ((fi73) | 1))) | 1))
                                                                    case "foo bar":
                                                                        if (!((v3 && true) && v3)) {
                                                                            v95 := GenS8{F0: 18446744073709551615}
                                                                            v96 := GenS12{F0: 42, F1: v75, F2: "\n\t"}
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                    var v97 int = fi73
                                                                    v1 = v1
                                                                    _ = v97
                                                                }
                                                                if (!(((("a" + ((("foo bar" + "") + "hello") + "")) > "\n\t") && (467 > 853)) && (-25 < -128))) {
                                                                    v65 = (v65 << (((v1 % ((18446744073709551615) | 1))) & 63))
                                                                    switch int(v65) {
                                                                    case -2147483648:
                                                                        v98 := GenS11{F0: 2, F1: -9223372036854775808}
                                                                        _ = v98
                                                                    case 1:
                                                                        v1 += (^(v1))
                                                                    default:
                                                                        if (!true) {
                                                                            if (v1 != (^(v1))) {
                                                                                v65 -= 42
                                                                                v99 := GenS1{F0: uint8(v2)}
                                                                                _ = v99
                                                                            } else {
                                                                                _ = len(v5)
                                                                                v4["hello"] = fi72
                                                                                v70 = (v70 << ((4294967295) & 63))
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi73, v3, v66)
                                                                    var v100 int = 4294967295
                                                                    _ = v100
                                                                }
                                                                v5 = append(v5, 888)
                                                                v1 = v1
                                                                v3 = (v1 <= v1)
                                                                v70 = (v70 + 2147483647)
                                                                var v101 float32 = (-2.25 - 0.1)
                                                                _ = v101
                                                            } else {
                                                                fmt.Println(v65, v66)
                                                                v102 := GenS11{F0: 199, F1: (fi72 + fi72)}
                                                                v4[("hello" + "")] = 1
                                                                _ = v102
                                                            }
                                                            fmt.Println(v2)
                                                        } else {
                                                            v66 -= v70
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v103 int = fi73
                                                                var v104 int8 = v65
                                                                var v105 float32 = (3.14159 + (1.5 * ((-2.25 * 0.0) * (-2.25 - 3.14159))))
                                                                _ = v103
                                                                _ = v104
                                                                _ = v105
                                                            case "\n\t":
                                                                var v106 float64 = 3.14159
                                                                _ = v106
                                                            }
                                                            v4["foo bar"] = (fi73 << ((4294967295) & 63))
                                                        }
                                                    case "hello":
                                                        var v107 int16 = int16(v2)
                                                        _ = v107
                                                    default:
                                                        v2 = (v2 / (((v2 % ((v75) | 1))) | 1))
                                                    }
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v75
                                            }
                                        default:
                                            fmt.Println(v2, v2, v66)
                                        }
                                        var v108 int64 = int64(v1)
                                        _ = v108
                                    } else {
                                        var v109 uint64 = v2
                                        if true {
                                            var v110 float64 = 3.14159
                                            fmt.Println(v3, v109)
                                            switch "hello" {
                                            case "a":
                                                v4[("" + (("foo bar" + ("a" + "hello")) + "\n\t"))] = (fi72 >> ((757) & 63))
                                            case "foo bar":
                                                if (-936 > fi73) {
                                                    if true {
                                                        v2 += (v109 % (((v2 >> ((v1) & 63))) | 1))
                                                        var v111 uint64 = 2
                                                        delete(v4, "hello")
                                                        _ = v111
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            case "hello":
                                                if (236 <= 128) {
                                                    fmt.Println(v110)
                                                    var v112 uint8 = 42
                                                    var v113 int8 = v65
                                                    v3 = (v1 == 4294967295)
                                                    _ = v112
                                                    _ = v113
                                                } else {
                                                    var v114 float64 = 0.0
                                                    var v115 uint16 = 32767
                                                    v109 -= v109
                                                    _ = v114
                                                    _ = v115
                                                }
                                            }
                                            _ = v110
                                        } else {
                                            v116 := GenS13{F0: 42, F1: v109}
                                            if true {
                                                if (!((^(v1)) >= 411)) {
                                                    v65 = (v65 / ((127) | 1))
                                                } else {
                                                    v117 := GenS13{F0: 65535, F1: (v109 / ((v2) | 1))}
                                                    if v3 {
                                                        fmt.Println(fi72, v2)
                                                    } else {
                                                        v118 := v4[(((("" + "a") + "") + ("a" + "a")) + "a")]
                                                        switch (fi72 / ((v118) | 1)) {
                                                        case 32767:
                                                            v119 := GenS2{F0: (v65 * v65), F1: (v66 << ((0) & 63)), F2: (v118 & v118)}
                                                            _ = v119
                                                        }
                                                        _ = len(v4)
                                                        _ = v118
                                                    }
                                                    _ = v117
                                                }
                                                if (true || (((v2 * 4294967295) < 2) && (((fi73 > (-(fi72))) || true) || v3))) {
                                                    if len(v5) > 0 {
                                                        v5[0] = 42
                                                    }
                                                    var v120 uint16 = 969
                                                    if (1 > 0) {
                                                        if v3 {
                                                            var v121 uint8 = 42
                                                            if (true && v3) {
                                                                fmt.Println(v1)
                                                                v120 = uint16(v121)
                                                                switch 127 {
                                                                case 42:
                                                                    v109 = 4294967295
                                                                }
                                                            } else {
                                                                if (2 > (v120 >> (((^(v1))) & 63))) {
                                                                    switch fi72 {
                                                                    case -2147483648:
                                                                        fmt.Println(v1)
                                                                    }
                                                                    var v122 float64 = 1e10
                                                                    switch (fi73 << ((v1) & 63)) {
                                                                    case 42:
                                                                        v123 := GenS6{F0: true}
                                                                        _ = v123
                                                                    }
                                                                    _ = len(v5)
                                                                    _ = v122
                                                                } else {
                                                                    var v124 uint32 = 65535
                                                                    if ((!(42 != v120)) && v3) {
                                                                        if (!(v66 <= 189)) {
                                                                            v125 := GenS13{F0: uint16(v109), F1: 1}
                                                                            if false {
                                                                                v65 = (v65 ^ v65)
                                                                            }
                                                                            fmt.Println(v3)
                                                                            _ = v125
                                                                        }
                                                                        if v3 {
                                                                            var v126 int8 = (v65 & v65)
                                                                            _ = v126
                                                                        } else {
                                                                            if false {
                                                                                var v127 float64 = -2.25
                                                                                v127 += v127
                                                                                switch (("\n\t" + "foo bar") + "a") {
                                                                                case "":
                                                                                    if v3 {
                                                                                        if ((v3 || v3) || v3) {
                                                                                            fmt.Println(v127, v70)
                                                                                            var v128 string = "a"
                                                                                            v129 := GenS2{F0: v65, F1: (v70 << ((v1) & 63)), F2: -934}
                                                                                            v3 = ((v3 && (v1 == (v1 >> ((65535) & 63)))) || v3)
                                                                                            fmt.Println(v109, v3, v109)
                                                                                            _ = len(v4)
                                                                                            _ = v128
                                                                                            _ = v129
                                                                                        }
                                                                                    } else {
                                                                                        var v130 uint32 = v124
                                                                                        v131 := GenS9{F0: v109, F1: (-(v65)), F2: (v120 > (v120 + v120))}
                                                                                        _ = v130
                                                                                        _ = v131
                                                                                    }
                                                                                case "hello":
                                                                                    fmt.Println(v109, v1, v120)
                                                                                case "\n\t":
                                                                                    var v132 string = "\n\t"
                                                                                    _ = v132
                                                                                default:
                                                                                    var v133 bool = ((v120 ^ v120) < (^(v120)))
                                                                                    if false {
                                                                                        v1 = v1
                                                                                        var v134 uint8 = (v121 << ((v1) & 63))
                                                                                        _ = v134
                                                                                    } else {
                                                                                        var v135 uint = 2
                                                                                        fmt.Println(v121, v121)
                                                                                        v65 = v65
                                                                                        _ = v135
                                                                                    }
                                                                                    _ = v133
                                                                                }
                                                                                _ = v127
                                                                            } else {
                                                                                v1 = (v1 | v1)
                                                                                var v136 string = "hello"
                                                                                v136 = "hello"
                                                                                var v137 uint16 = v120
                                                                                v138 := GenS2{F0: -128, F1: 2147483647, F2: 9223372036854775807}
                                                                                _ = v136
                                                                                _ = v137
                                                                                _ = v138
                                                                            }
                                                                            v124 = (v124 ^ v124)
                                                                            fmt.Println(v120, v3)
                                                                        }
                                                                    }
                                                                    _ = v124
                                                                }
                                                                if (1e10 >= 1.5) {
                                                                    if (0 < 0) {
                                                                        v109 = (v2 & v2)
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            var v139 string = "hello"
                                                                            switch 24 {
                                                                            case 2147483647:
                                                                                if (false || (!(!(v3 || true)))) {
                                                                                    v65 = v65
                                                                                } else {
                                                                                    fmt.Println(v120)
                                                                                    var v140 int = fi73
                                                                                    _ = v140
                                                                                }
                                                                            case 0:
                                                                                var v141 int64 = -132
                                                                                _ = v141
                                                                            case 42:
                                                                                fmt.Println(v1)
                                                                            }
                                                                            v142 := GenS14{F0: v109, F1: (^(v120))}
                                                                            _ = v139
                                                                            _ = v142
                                                                        case "":
                                                                            fmt.Println(v2)
                                                                        case "\n\t":
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        default:
                                                                            if v3 {
                                                                                if false {
                                                                                    if true {
                                                                                        fmt.Println(v1, fi73, v66)
                                                                                        _ = len(v4)
                                                                                    } else {
                                                                                        var v143 uint8 = v121
                                                                                        v144 := GenS15{F0: v109, F1: v1, F2: v70}
                                                                                        v145 := GenS1{F0: (v143 % (((v143 ^ (v143 % ((255) | 1)))) | 1))}
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                        _ = v145
                                                                                    }
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                }
                                                                                v70 = v66
                                                                            } else {
                                                                                v146 := v4[("hello" + "a")]
                                                                                _ = v146
                                                                            }
                                                                        }
                                                                        var v147 string = "foo bar"
                                                                        switch fi73 {
                                                                        case 4294967295:
                                                                            fmt.Println(v65, v147)
                                                                        case 127:
                                                                            if len(v5) > 0 {
                                                                                v5[0] = 724
                                                                            }
                                                                        }
                                                                        _ = v147
                                                                    }
                                                                } else {
                                                                    v148 := v4[("hello" + "")]
                                                                    if (int8(v120) == v65) {
                                                                        v66 = v66
                                                                        var v149 float32 = ((0.1 * float32((1.5 - ((((0.1 + 1.5) * 1.5) + 1e10) / 0.1)))) - (0.0 / (-2.25 + 3.14159)))
                                                                        _ = v149
                                                                    } else {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 2147483647
                                                                        }
                                                                        v120 -= v120
                                                                    }
                                                                    _ = v148
                                                                }
                                                                fmt.Println(fi72)
                                                            }
                                                            var v150 uint8 = 2
                                                            _ = v121
                                                            _ = v150
                                                        } else {
                                                            if (true || v3) {
                                                                _ = len(v5)
                                                                v66 += v70
                                                            }
                                                            if v3 {
                                                                fmt.Println(v109)
                                                                fmt.Println(fi72, fi73, v120)
                                                            } else {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    if (v65 > (v65 | (v65 | -128))) {
                                                                        v151 := v4[("hello" + "foo bar")]
                                                                        switch ((((("" + ("\n\t" + "a")) + "hello") + "") + (("foo bar" + "hello") + "foo bar")) + (("hello" + ("hello" + ("hello" + ("" + ("foo bar" + ("foo bar" + ("foo bar" + ("a" + "\n\t")))))))) + "a")) {
                                                                        case "":
                                                                            v2 += (^(v109))
                                                                        default:
                                                                            var v152 uint32 = uint32(v2)
                                                                            _ = v152
                                                                        }
                                                                        fmt.Println(v2, v151, v120)
                                                                        _ = v151
                                                                    } else {
                                                                        v66 += v66
                                                                        switch (("a" + "foo bar") + "foo bar") {
                                                                        case "foo bar":
                                                                            var v153 uint16 = (v120 ^ 255)
                                                                            _ = v153
                                                                        case "\n\t":
                                                                            v154 := GenS7{F0: (v65 + 0)}
                                                                            _ = v154
                                                                        default:
                                                                            var v155 float32 = 3.14159
                                                                            _ = v155
                                                                        }
                                                                        v3 = (!true)
                                                                    }
                                                                case "":
                                                                    _ = len(v4)
                                                                default:
                                                                    v2 += (v109 - (v2 + (^(v2))))
                                                                }
                                                            }
                                                            if (v3 || v3) {
                                                                var v156 bool = v3
                                                                v157 := GenS16{F0: 9223372036854775807, F1: 0}
                                                                _ = v156
                                                                _ = v157
                                                            }
                                                        }
                                                    }
                                                    switch (fi72 - (fi73 + fi72)) {
                                                    case 128:
                                                        v5 = append(v5, 65535)
                                                        v2 = v2
                                                    case -1:
                                                        delete(v4, "a")
                                                    default:
                                                        v120 = (v120 % ((v120) | 1))
                                                    }
                                                    v158 := GenS17{F0: v1}
                                                    _ = v120
                                                    _ = v158
                                                }
                                                var v159 uint8 = 2
                                                _ = v159
                                            } else {
                                                var v160 bool = (fi73 >= (fi72 + fi73))
                                                var v161 int64 = 0
                                                fmt.Println(v3, v109, fi72)
                                                _ = v160
                                                _ = v161
                                            }
                                            var v162 int16 = 32767
                                            _ = v116
                                            _ = v162
                                        }
                                        v163 := GenS18{F0: (v70 >> ((4294967295) & 63)), F1: (v66 & (v66 * (v66 / (((-(v70))) | 1))))}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v109
                                        _ = v163
                                    }
                                    var v164 bool = ((float32(0.0) + -2.25) == 0.1)
                                    var v165 uint = v1
                                    _ = v164
                                    _ = v165
                                } else {
                                    v2 = v2
                                }
                                fmt.Println(v2, fi73)
                            }
                        default:
                            v166 := GenS19{F0: v1}
                            _ = v166
                        }
                        var v167 int32 = v66
                        var v168 uint8 = 127
                        delete(v4, ("a" + ("a" + "a")))
                        _ = fi73
                        _ = v167
                        _ = v168
                    }
                    _ = fi72
                }
                v4["a"] = 2147483647
                v70 = (^(v70))
            } else {
                for fi169 := 0; fi169 < 19; fi169++ {
                    if (!true) {
                        if (v3 && v3) {
                            v170 := GenS15{F0: (v2 * v2), F1: v1, F2: (-(v66))}
                            for fi171 := 0; fi171 < 4; fi171++ {
                                fmt.Println(v2, v2)
                                _ = fi171
                            }
                            var v172 int64 = int64(v2)
                            _ = v170
                            _ = v172
                        }
                        v1 = (v1 ^ v1)
                    }
                    v1 += 1
                    _ = fi169
                }
            }
            for fi173 := 0; fi173 < 11; fi173++ {
                var v174 bool = v3
                for fi175 := 0; fi175 < 6; fi175++ {
                    v176 := GenS12{F0: 59, F1: 780, F2: "hello"}
                    if v3 {
                        if (-1 <= -1) {
                            var v177 uint8 = 128
                            _ = v177
                        }
                        fmt.Println(v2, v1, v66)
                        v178 := GenS20{F0: 0, F1: 9223372036854775807}
                        _ = v178
                    }
                    _ = fi175
                    _ = v176
                }
                _ = fi173
                _ = v174
            }
            _ = v71
        } else {
            switch 404 {
            case -273:
                fmt.Println(v2, v66)
            case 4294967295:
                v179 := GenS9{F0: 0, F1: v65, F2: (v65 < v65)}
                _ = v179
            case -573:
                fmt.Println(v3, v70, v65)
            default:
                for fi180 := 0; fi180 < 12; fi180++ {
                    var v181 string = ("hello" + "")
                    _ = fi180
                    _ = v181
                }
            }
            fmt.Println(v66, v2)
        }
    case "a":
        fmt.Println(v70, v70)
    default:
        var v182 int8 = v65
        _ = v182
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi183 := 0; fi183 < 4; fi183++ {
        var v184 uint8 = 128
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi183
        _ = v184
    }
    for fi185 := 0; fi185 < 19; fi185++ {
        var v186 bool = (704 > 61)
        for fi187 := 0; fi187 < 15; fi187++ {
            v188 := GenS21{F0: -50, F1: "foo bar", F2: (v65 % ((42) | 1))}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            switch fi187 {
            case 128:
                v189 := GenS17{F0: v1}
                _ = v189
            }
            _ = fi187
            _ = v188
        }
        _ = len(v4)
        _ = fi185
        _ = v186
    }
    if (v3 || true) {
        fmt.Println(v66)
    } else {
        fmt.Println(v1, v70)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v69
    _ = v70
}
