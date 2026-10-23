package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 string; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int8; F2 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint16; F1 int32 }
type GenS13 struct { F0 int; F1 bool }
type GenS14 struct { F0 string; F1 uint16; F2 uint16 }
type GenS15 struct { F0 uint8; F1 uint }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 uint16; F1 uint8 }
type GenS18 struct { F0 int; F1 int8; F2 int32 }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS21 struct { F0 int32; F1 uint; F2 int16 }
type GenS22 struct { F0 int64; F1 int; F2 int16 }
type GenS23 struct { F0 uint16; F1 int32 }
type GenS24 struct { F0 bool; F1 string }
type GenS25 struct { F0 uint16; F1 int64 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 int16; F1 int64; F2 uint16 }
type GenS28 struct { F0 int8; F1 int }
func genFn0() bool {
    var v5 uint8 = 127
    fmt.Println(v5)
    var v6 float32 = ((0.1 + 1e10) / 3.14159)
    if (true && true) {
        v6 -= v6
        if (!(!false)) {
            var v7 int8 = 0
            _ = v7
        }
        v8 := GenS1{F0: 2147483647}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            v6 -= ((0.0 + v6) / v6)
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 int32 = 2147483647
                switch ("a" + "foo bar") {
                case "hello":
                    var v12 uint16 = 5
                    _ = v12
                case "foo bar":
                    var v13 bool = false
                    _ = v13
                default:
                    v5 = 127
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        if true {
            v6 = 1.5
        } else {
            if false {
                var v14 bool = false
                _ = v14
            } else {
                for fi15 := 0; fi15 < 5; fi15++ {
                    v5 = (v5 / ((v5) | 1))
                    v16 := []int{(fi15 | (-(fi15)))}
                    _ = fi15
                    _ = v16
                }
                switch "foo bar" {
                case "\n\t":
                    v17 := GenS2{F0: "\n\t", F1: ((("\n\t" + "\n\t") + ("foo bar" + "")) + (("\n\t" + "hello") + "\n\t")), F2: -419}
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        switch "a" {
                        case "\n\t":
                            v19 := GenS1{F0: 42}
                            _ = v19
                        case "hello":
                            if (-9223372036854775808 < (^(fi18))) {
                                for fi20 := 0; fi20 < 3; fi20++ {
                                    if true {
                                        v21 := []int{(fi20 % ((fi18) | 1)), fi18, 4294967295}
                                        v6 = v6
                                        _ = v21
                                    } else {
                                        var v22 int8 = 30
                                        _ = v22
                                    }
                                    _ = fi20
                                }
                                for fi23 := 0; fi23 < 5; fi23++ {
                                    v24 := GenS2{F0: ("\n\t" + "hello"), F1: "hello", F2: 1}
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        default:
                            fmt.Println(v5, v6, v6)
                        }
                        fmt.Println(v5, v6)
                        _ = fi18
                    }
                }
                var v25 int = 32767
                _ = v25
            }
        }
    }
    _ = v5
    _ = v6
    return (!true)
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = ""
    var v3 int8 = -102
    v4 := make(map[string]string)
    if false {
        var v26 uint64 = 230
        if true {
            v27 := []int{127}
            v3 = (-(v3))
            _ = v27
        }
        fmt.Println(v1)
        _ = v26
    }
    var v28 uint16 = 381
    switch ((v2 + "foo bar") + v2) {
    case "":
        var v29 int = 2147483647
        _ = v29
    case "a":
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int8 = (v3 >> ((65535) & 63))
            _ = fi30
            _ = v31
        }
    case "hello":
        if true {
            v32 := GenS3{F0: 255}
            v3 = (v3 ^ v3)
            var v33 string = v2
            _ = v32
            _ = v33
        }
    default:
        v34 := GenS3{F0: 128}
        _ = v34
    }
    v35 := GenS2{F0: v2, F1: ((v2 + v2) + ((((v2 + v2) + "hello") + v2) + "")), F2: 0}
    fmt.Println(v2, v2)
    v36 := []int{128, 42, 32767, 127}
    if (v3 == (v3 << ((4294967295) & 63))) {
        for fi37 := 0; fi37 < 2; fi37++ {
            v1 = (v1 * v1)
            v3 = v3
            _ = fi37
        }
        switch 127 {
        case 750:
            if (12 < 65535) {
                v38 := GenS1{F0: 0}
                for fi39 := 0; fi39 < 10; fi39++ {
                    switch "" {
                    case "hello":
                        v28 -= v28
                    case "":
                        delete(v4, v2)
                    case "foo bar":
                        var v40 uint16 = 717
                        _ = v40
                    }
                    _ = fi39
                }
                if false {
                    fmt.Println(v3)
                }
                _ = v38
            } else {
                if true {
                    v28 += v28
                    v4[(v2 + (v2 + v2))] = ("\n\t" + "")
                    switch (((v2 + v2) + ("a" + (v2 + ("a" + (v2 + (v2 + (v2 + "a"))))))) + "\n\t") {
                    case "foo bar":
                        if ((false && (true && (!(1 < 65535)))) && false) {
                            v3 -= (v3 * (v3 % ((v3) | 1)))
                        } else {
                            var v41 uint16 = v28
                            _ = v41
                        }
                    }
                } else {
                    var v42 int16 = 1
                    if false {
                        if ((v42 > -127) || true) {
                            fmt.Println(v28, v42, v42)
                            v2 = "a"
                            if false {
                                fmt.Println(v42, v3)
                                switch (("hello" + "a") + "hello") {
                                case "a":
                                    for fi43 := 0; fi43 < 4; fi43++ {
                                        v42 -= (v42 + v42)
                                        v28 = (v28 + 255)
                                        _ = fi43
                                    }
                                case "foo bar":
                                    fmt.Println(v3)
                                default:
                                    for fi44 := 0; fi44 < 15; fi44++ {
                                        fmt.Println(v3, v2, fi44)
                                        v42 = v42
                                        _ = fi44
                                    }
                                }
                                var v45 int = 0
                                _ = v45
                            }
                        } else {
                            fmt.Println(v42, v42)
                        }
                        var v46 uint32 = 885
                        _ = v46
                    }
                    _ = v42
                }
            }
        case 2147483647:
            if false {
                var v47 string = v2
                _ = v47
            } else {
                if (!false) {
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                }
                v28 = (v28 ^ (v28 << ((2) & 63)))
                if true {
                    v1 += (v1 * v1)
                } else {
                    for fi48 := 0; fi48 < 16; fi48++ {
                        if false {
                            if (0 >= 4294967295) {
                                for fi49 := 0; fi49 < 5; fi49++ {
                                    if true {
                                        var v50 int8 = (v3 >> ((2) & 63))
                                        _ = v50
                                    } else {
                                        if len(v36) >= 2 {
                                            v36 = v36[1:len(v36)]
                                        }
                                    }
                                    switch fi49 {
                                    case 9223372036854775807:
                                        if ((v3 % (((v3 / ((v3) | 1))) | 1)) > (v3 ^ (v3 | (-(v3))))) {
                                            switch (fi49 % ((fi49) | 1)) {
                                            case 171:
                                                if len(v36) >= 2 {
                                                    v36 = v36[1:len(v36)]
                                                }
                                            case 128:
                                                v1 -= ((((v1 * ((0.0 + 3.14159) - v1)) + (v1 / ((((0.1 + v1) / (v1 - (0.1 * (3.14159 + v1)))) + (float64(-2.25) - float64(0.1))) * float64((3.14159 * -2.25))))) + v1) + 0.1)
                                            case -875:
                                                if true {
                                                    _ = len(v36)
                                                    var v51 int8 = 1
                                                    _ = v51
                                                } else {
                                                    var v52 uint32 = 683
                                                    _ = v52
                                                }
                                            default:
                                                v2 = (v2 + "a")
                                            }
                                        } else {
                                            v3 = v3
                                            var v53 int = fi49
                                            var v54 int16 = 254
                                            v55 := GenS3{F0: 128}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                    case 2147483647:
                                        v28 -= 0
                                        var v56 int16 = -1
                                        if len(v36) > 0 {
                                            v36[0] = fi49
                                        }
                                        _ = v56
                                    case -1:
                                        var v57 float32 = float32(v1)
                                        _ = v57
                                    }
                                    _ = fi49
                                }
                            } else {
                                for fi58 := 0; fi58 < 13; fi58++ {
                                    v28 += (v28 << ((4294967295) & 63))
                                    v59 := GenS2{F0: "foo bar", F1: "a", F2: 468}
                                    _ = fi58
                                    _ = v59
                                }
                                delete(v4, ("a" + v2))
                            }
                            fmt.Println(v3, fi48, v3)
                        } else {
                            _ = len(v4)
                        }
                        switch -877 {
                        case -976:
                            var v60 int64 = 9223372036854775807
                            _ = v60
                        case -375:
                            for fi61 := 0; fi61 < 4; fi61++ {
                                v62 := GenS2{F0: v2, F1: (v2 + "a"), F2: 42}
                                var v63 uint16 = v28
                                v1 = 0.1
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                        case 128:
                            var v64 int32 = -1
                            _ = v64
                        }
                        _ = fi48
                    }
                    for fi65 := 0; fi65 < 15; fi65++ {
                        if true {
                            var v66 int8 = v3
                            for fi67 := 0; fi67 < 1; fi67++ {
                                if (("hello" + "hello") <= ("\n\t" + ("a" + ""))) {
                                    if (!(v1 <= v1)) {
                                        var v68 int8 = v66
                                        switch v2 {
                                        case "\n\t":
                                            var v69 int = (fi65 / ((-9223372036854775808) | 1))
                                            _ = v69
                                        case "hello":
                                            if ((v1 + 0.0) < (v1 / float64(float32((v1 + (1.5 - float64(1e10))))))) {
                                                v70 := GenS3{F0: 0}
                                                _ = v70
                                            }
                                        case "a":
                                            fmt.Println(fi65, v3)
                                        default:
                                            if (fi67 < (^(fi67))) {
                                                var v71 int64 = -1
                                                delete(v4, v2)
                                                v71 = (-(v71))
                                                if (65535 >= 0) {
                                                    v71 = (-(v71))
                                                    var v72 int8 = int8(v71)
                                                    v71 += -86
                                                    if true {
                                                        var v73 uint = 2
                                                        switch fi67 {
                                                        case 127:
                                                            var v74 int8 = -118
                                                            _ = v74
                                                        case 128:
                                                            v75 := GenS4{F0: (v73 | 4294967295), F1: (v28 * (^(v28)))}
                                                            _ = v75
                                                        }
                                                        var v76 uint64 = 2
                                                        _ = v73
                                                        _ = v76
                                                    } else {
                                                        v77 := GenS1{F0: 0}
                                                        var v78 bool = (558 != 42)
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    fmt.Println(fi65, v71, v66)
                                                    var v79 uint16 = v28
                                                    var v80 uint16 = v79
                                                    _ = v72
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    v81 := GenS1{F0: 0}
                                                    _ = v81
                                                }
                                                _ = v71
                                            } else {
                                                var v82 float64 = v1
                                                v4[(v2 + "\n\t")] = "a"
                                                v4[(v2 + "hello")] = ("hello" + v2)
                                                _ = v82
                                            }
                                        }
                                        var v83 uint32 = 65535
                                        switch fi67 {
                                        case -9223372036854775808:
                                            if true {
                                                var v84 float64 = v1
                                                _ = v84
                                            }
                                        default:
                                            switch (fi65 | (fi65 >> ((0) & 63))) {
                                            case -609:
                                                v85 := GenS5{F0: 150, F1: (v68 % (((v66 | v68)) | 1)), F2: 127}
                                                _ = v85
                                            case -309:
                                                v3 += v3
                                            }
                                        }
                                        fmt.Println(v68, v3, v66)
                                        if true {
                                            var v86 uint64 = 18446744073709551615
                                            if false {
                                                if false {
                                                    var v87 uint64 = 18446744073709551615
                                                    var v88 int16 = -228
                                                    var v89 uint64 = 212
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                }
                                                v90 := GenS6{F0: 0}
                                                v83 = v83
                                                _ = v90
                                            }
                                            v3 -= (v3 << ((580) & 63))
                                            v91 := GenS6{F0: 437}
                                            var v92 uint64 = v86
                                            _ = v86
                                            _ = v91
                                            _ = v92
                                        } else {
                                            if false {
                                                if false {
                                                    v93 := GenS4{F0: 65535, F1: 0}
                                                    _ = v93
                                                } else {
                                                    v1 = float64((0.0 / ((((0.0 + 1.5) / 3.14159) / 1.5) / -2.25)))
                                                    var v94 uint16 = (v28 + 65535)
                                                    if (127 != 20) {
                                                        var v95 int8 = (^(v3))
                                                        switch (-(fi65)) {
                                                        case 228:
                                                            var v96 float32 = (1e10 * (3.14159 * float32((v1 - v1))))
                                                            _ = v96
                                                        default:
                                                            if true {
                                                                v3 = (v66 % (((v3 + v66)) | 1))
                                                                v97 := GenS6{F0: 0}
                                                                _ = v97
                                                            } else {
                                                                if (!true) {
                                                                    var v98 uint = 65535
                                                                    var v99 float32 = 1.5
                                                                    v100 := GenS5{F0: 42, F1: -11, F2: -128}
                                                                    _ = v100
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                                v66 -= (v3 + v95)
                                                                var v101 uint64 = 18446744073709551615
                                                                var v102 bool = (!true)
                                                                _ = v101
                                                                _ = v102
                                                            }
                                                        }
                                                        var v103 string = ""
                                                        _ = v103
                                                        _ = v95
                                                    } else {
                                                        v104 := GenS4{F0: 2, F1: 32767}
                                                        _ = v104
                                                    }
                                                    _ = v94
                                                }
                                                v105 := GenS1{F0: 0}
                                                v66 -= (v3 | v66)
                                                _ = v105
                                            } else {
                                                v106 := GenS3{F0: 33}
                                                if false {
                                                    v83 = v83
                                                    var v107 float32 = (1.5 * (0.0 / float32(1.5)))
                                                    fmt.Println(v66, v3, fi67)
                                                    v68 = v68
                                                    _ = v107
                                                }
                                                _ = v106
                                            }
                                            if false {
                                                v68 -= v66
                                                v28 = 42
                                                if (!false) {
                                                    var v108 bool = (((v2 + ("hello" + v2)) == "") && false)
                                                    _ = v108
                                                } else {
                                                    v66 += v68
                                                }
                                            } else {
                                                var v109 uint16 = (v28 & v28)
                                                _ = v109
                                            }
                                        }
                                        v66 = -128
                                        _ = len(v36)
                                        var v110 uint = 578
                                        var v111 int16 = 42
                                        if true {
                                            v3 = (v68 / (((v68 % ((v68) | 1))) | 1))
                                            fmt.Println(v66)
                                            if (v110 >= 1) {
                                                if false {
                                                    if false {
                                                        _ = len(v4)
                                                    }
                                                }
                                            } else {
                                                v112 := GenS6{F0: -1}
                                                switch v2 {
                                                case "a":
                                                    if true {
                                                        if ("foo bar" < "\n\t") {
                                                            var v113 uint32 = v83
                                                            if true {
                                                                if ((^(v83)) > 0) {
                                                                    var v114 int64 = -915
                                                                    _ = v114
                                                                } else {
                                                                    if (true && (v2 != "")) {
                                                                        var v115 uint8 = 255
                                                                        _ = v115
                                                                    } else {
                                                                        var v116 float64 = (v1 + v1)
                                                                        v113 = v113
                                                                        _ = v116
                                                                    }
                                                                    fmt.Println(v83, v83, v3)
                                                                    var v117 int = (fi65 / ((fi65) | 1))
                                                                    _ = v117
                                                                }
                                                                switch (fi65 + fi67) {
                                                                case 0:
                                                                    switch fi67 {
                                                                    case 32768:
                                                                        if (9223372036854775807 != -37) {
                                                                            _ = len(v36)
                                                                            if (42 != (v28 ^ 65535)) {
                                                                                if false {
                                                                                    if false {
                                                                                        if (v113 < v113) {
                                                                                            v118 := GenS4{F0: (v110 / (((v110 & (v110 % (((v110 / (((^(v110))) | 1))) | 1)))) | 1)), F1: v28}
                                                                                            _ = len(v4)
                                                                                            v119 := GenS7{F0: 9223372036854775807, F1: v28}
                                                                                            _ = v118
                                                                                            _ = v119
                                                                                        } else {
                                                                                            if true {
                                                                                                if (true && true) {
                                                                                                    var v120 uint8 = 0
                                                                                                    delete(v4, "a")
                                                                                                    _ = v120
                                                                                                } else {
                                                                                                    if (!(2 <= 42)) {
                                                                                                        v4[""] = ("foo bar" + v2)
                                                                                                        if len(v36) >= 2 {
                                                                                                            v36 = v36[1:len(v36)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v121 uint8 = 1
                                                                                                        fmt.Println(v111, fi67, v1)
                                                                                                        _ = v121
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v122 uint64 = 2
                                                                                                _ = v122
                                                                                            }
                                                                                            switch -533 {
                                                                                            case 30:
                                                                                                switch ((((v2 + ("foo bar" + ((("a" + v2) + "\n\t") + v2))) + ("foo bar" + ("a" + "\n\t"))) + (("a" + ("\n\t" + v2)) + v2)) + "foo bar") {
                                                                                                case "\n\t":
                                                                                                    _ = len(v36)
                                                                                                case "":
                                                                                                    fmt.Println(v1, v113)
                                                                                                }
                                                                                            case -546:
                                                                                                if len(v36) > 0 {
                                                                                                    v36[0] = int(v68)
                                                                                                }
                                                                                            default:
                                                                                                v68 = 36
                                                                                            }
                                                                                            var v123 int8 = (v68 * (v66 ^ -68))
                                                                                            _ = v123
                                                                                        }
                                                                                        if ((v110 / ((v110) | 1)) < v110) {
                                                                                            var v124 int64 = 1
                                                                                            _ = v124
                                                                                        }
                                                                                    } else {
                                                                                        v125 := GenS5{F0: 159, F1: v3, F2: v66}
                                                                                        _ = v125
                                                                                    }
                                                                                    var v126 int8 = (v3 ^ v3)
                                                                                    var v127 float64 = 3.14159
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                } else {
                                                                                    var v128 uint = (v110 - v110)
                                                                                    _ = v128
                                                                                }
                                                                            }
                                                                            if (1e10 != 0.1) {
                                                                                fmt.Println(v28, v111, v111)
                                                                                fmt.Println(v66, fi67, v68)
                                                                                var v129 uint32 = v83
                                                                                _ = v129
                                                                            } else {
                                                                                var v130 float64 = (-2.25 - v1)
                                                                                v111 = v111
                                                                                _ = v130
                                                                            }
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi65, v113, v28)
                                                                    }
                                                                case 1:
                                                                    v131 := GenS4{F0: (^(v110)), F1: v28}
                                                                    v36 = append(v36, (fi65 + fi67))
                                                                    if true {
                                                                        v2 = v2
                                                                        var v132 uint8 = 1
                                                                        if false {
                                                                            if false {
                                                                                if len(v36) > 0 {
                                                                                    v36[0] = (^(fi65))
                                                                                }
                                                                            } else {
                                                                                switch v2 {
                                                                                case "hello":
                                                                                    switch v2 {
                                                                                    case "\n\t":
                                                                                        switch int(v111) {
                                                                                        case -1:
                                                                                            if false {
                                                                                                v133 := GenS8{F0: v111}
                                                                                                _ = v133
                                                                                            } else {
                                                                                                var v134 uint16 = v28
                                                                                                _ = v134
                                                                                            }
                                                                                        default:
                                                                                            v111 = 650
                                                                                        }
                                                                                    case "hello":
                                                                                        if false {
                                                                                            v83 -= (^(v83))
                                                                                        }
                                                                                    case "a":
                                                                                        var v135 int32 = -46
                                                                                        _ = v135
                                                                                    default:
                                                                                        if true {
                                                                                            v132 = (v132 >> (((v110 - 185)) & 63))
                                                                                            switch ("" + v2) {
                                                                                            case "a":
                                                                                                delete(v4, v2)
                                                                                            case "foo bar":
                                                                                                if (v3 <= (v66 % ((v66) | 1))) {
                                                                                                    if false {
                                                                                                        v83 = (v113 ^ v113)
                                                                                                        if ((v28 << (((^(v110))) & 63)) >= 519) {
                                                                                                            v136 := GenS2{F0: v2, F1: v2, F2: 737}
                                                                                                            _ = v136
                                                                                                        }
                                                                                                        v2 = (("\n\t" + v2) + v2)
                                                                                                    }
                                                                                                    v66 = v66
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v137 uint = v110
                                                                                            v138 := GenS2{F0: ("a" + (((v2 + v2) + v2) + "foo bar")), F1: ("" + ("" + v2)), F2: 1}
                                                                                            var v139 uint16 = 792
                                                                                            if false {
                                                                                                if true {
                                                                                                    fmt.Println(v2, v139, fi65)
                                                                                                }
                                                                                                fmt.Println(v132, v111, v66)
                                                                                            }
                                                                                            delete(v4, v2)
                                                                                            fmt.Println(v111)
                                                                                            _ = v137
                                                                                            _ = v138
                                                                                            _ = v139
                                                                                        }
                                                                                    }
                                                                                case "\n\t":
                                                                                    v140 := GenS9{F0: true, F1: v111}
                                                                                    v141 := GenS5{F0: v132, F1: -128, F2: -128}
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                case "foo bar":
                                                                                    var v142 int = fi65
                                                                                    _ = v142
                                                                                default:
                                                                                    fmt.Println(v111)
                                                                                }
                                                                                fmt.Println(v68, v113)
                                                                            }
                                                                            var v143 uint = 0
                                                                            v111 += (^(v111))
                                                                            switch v2 {
                                                                            case "":
                                                                                var v144 int8 = 0
                                                                                _ = v144
                                                                            default:
                                                                                fmt.Println(v28, v110)
                                                                            }
                                                                            if (!(v110 >= (v143 + v110))) {
                                                                                var v145 uint16 = v28
                                                                                fmt.Println(v2, v110)
                                                                                v68 += 42
                                                                                v111 += (^(v111))
                                                                                _ = v145
                                                                            } else {
                                                                                v146 := GenS5{F0: (v132 << ((2) & 63)), F1: (v68 - (v68 | int8(v132))), F2: (v66 ^ 81)}
                                                                                v147 := v4[v2]
                                                                                _ = v146
                                                                                _ = v147
                                                                            }
                                                                            var v148 int16 = (v111 << ((v110) & 63))
                                                                            fmt.Println(fi65)
                                                                            if len(v36) > 0 {
                                                                                v36[0] = fi67
                                                                            }
                                                                            _ = v143
                                                                            _ = v148
                                                                        }
                                                                        _ = v132
                                                                    }
                                                                    v68 = int8(v113)
                                                                    _ = v131
                                                                case -2147483648:
                                                                    if (!false) {
                                                                        if true {
                                                                            v113 = (v83 << ((v110) & 63))
                                                                            var v149 float64 = 1e10
                                                                            if true {
                                                                                if (v83 > (v83 * 1)) {
                                                                                    if ((((0.0 + (-2.25 - 1e10)) / 1.5) / -2.25) == 0.0) {
                                                                                        v150 := GenS10{F0: 0}
                                                                                        v151 := GenS11{F0: uint8(v110)}
                                                                                        _ = v150
                                                                                        _ = v151
                                                                                    }
                                                                                    var v152 int64 = 42
                                                                                    _ = len(v36)
                                                                                    _ = v152
                                                                                } else {
                                                                                    var v153 float32 = 0.0
                                                                                    v4[""] = (v2 + "")
                                                                                    _ = v153
                                                                                }
                                                                                switch fi65 {
                                                                                case -906:
                                                                                    fmt.Println(v3, v111, v111)
                                                                                }
                                                                                switch v2 {
                                                                                case "":
                                                                                    switch v2 {
                                                                                    case "hello":
                                                                                        if true {
                                                                                            var v154 uint16 = (v28 ^ (v28 - v28))
                                                                                            _ = v154
                                                                                        } else {
                                                                                            var v155 float32 = 1.5
                                                                                            if true {
                                                                                                var v156 float64 = (3.14159 + 0.0)
                                                                                                fmt.Println(v66)
                                                                                                if ((-58 >= 42) || true) {
                                                                                                    v157 := GenS4{F0: v110, F1: (v28 ^ uint16(v110))}
                                                                                                    _ = v157
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if (v155 >= 0.0) {
                                                                                                            switch 42 {
                                                                                                            case 4294967295:
                                                                                                                var v158 uint32 = v113
                                                                                                                _ = v158
                                                                                                            case 32767:
                                                                                                                v66 = (v66 + v68)
                                                                                                            case -9223372036854775808:
                                                                                                                v3 += v3
                                                                                                            }
                                                                                                            if true {
                                                                                                                _ = len(v36)
                                                                                                            }
                                                                                                            switch "a" {
                                                                                                            case "":
                                                                                                                if true {
                                                                                                                    switch (v2 + "foo bar") {
                                                                                                                    case "":
                                                                                                                        switch (fi67 % (((fi65 & 2147483647)) | 1)) {
                                                                                                                        case 2147483647:
                                                                                                                            fmt.Println(v28, v113)
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        _ = len(v36)
                                                                                                                    }
                                                                                                                    delete(v4, ((v2 + "\n\t") + "foo bar"))
                                                                                                                    v159 := GenS12{F0: v28, F1: 1}
                                                                                                                    _ = v159
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        fmt.Println(v28, v156)
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                v36 = append(v36, 0)
                                                                                                                v160 := GenS13{F0: int(v113), F1: false}
                                                                                                                _ = v160
                                                                                                            } else {
                                                                                                                if (4294967295 != 1) {
                                                                                                                    v36 = append(v36, -2147483648)
                                                                                                                } else {
                                                                                                                    fmt.Println(v68)
                                                                                                                    if ((-(fi65)) == (fi67 ^ fi67)) {
                                                                                                                        var v161 int32 = 406
                                                                                                                        if false {
                                                                                                                            v1 = v156
                                                                                                                        }
                                                                                                                        var v162 int64 = 9223372036854775807
                                                                                                                        _ = v161
                                                                                                                        _ = v162
                                                                                                                    }
                                                                                                                    fmt.Println(v2, v1)
                                                                                                                }
                                                                                                            }
                                                                                                            var v163 bool = true
                                                                                                            _ = v163
                                                                                                        }
                                                                                                        if len(v36) > 0 {
                                                                                                            v36[0] = fi65
                                                                                                        }
                                                                                                    } else {
                                                                                                        v66 = 1
                                                                                                        var v164 uint32 = 121
                                                                                                        if false {
                                                                                                            fmt.Println(v1)
                                                                                                            var v165 bool = (-1 == 0)
                                                                                                            if v165 {
                                                                                                                switch (v2 + "a") {
                                                                                                                case "foo bar":
                                                                                                                    v166 := GenS14{F0: ((v2 + v2) + v2), F1: 32767, F2: v28}
                                                                                                                    _ = v166
                                                                                                                case "":
                                                                                                                    _ = len(v36)
                                                                                                                default:
                                                                                                                    var v167 int32 = 0
                                                                                                                    _ = v167
                                                                                                                }
                                                                                                                var v168 uint = 408
                                                                                                                var v169 uint16 = v28
                                                                                                                var v170 uint64 = 18446744073709551615
                                                                                                                if (2 < 184) {
                                                                                                                    var v171 int8 = (v66 / (((v3 + int8(v111))) | 1))
                                                                                                                    var v172 int64 = 1
                                                                                                                    v173 := GenS15{F0: 22, F1: (v168 << ((0) & 63))}
                                                                                                                    _ = v171
                                                                                                                    _ = v172
                                                                                                                    _ = v173
                                                                                                                }
                                                                                                                _ = len(v36)
                                                                                                                _ = v168
                                                                                                                _ = v169
                                                                                                                _ = v170
                                                                                                            }
                                                                                                            _ = v165
                                                                                                        }
                                                                                                        _ = v164
                                                                                                    }
                                                                                                    v174 := GenS12{F0: (^(v28)), F1: 616}
                                                                                                    _ = v174
                                                                                                }
                                                                                                _ = v156
                                                                                            }
                                                                                            _ = v155
                                                                                        }
                                                                                    case "a":
                                                                                        var v175 uint16 = (v28 | (v28 + (^(v28))))
                                                                                        _ = v175
                                                                                    default:
                                                                                        var v176 int64 = 0
                                                                                        _ = v176
                                                                                    }
                                                                                case "a":
                                                                                    switch (fi67 / (((fi67 - 32768)) | 1)) {
                                                                                    case -2147483648:
                                                                                        var v177 uint32 = (v113 ^ v83)
                                                                                        _ = v177
                                                                                    case 2147483647:
                                                                                        if true {
                                                                                            switch ("a" + (v2 + "foo bar")) {
                                                                                            case "":
                                                                                                var v178 float64 = v1
                                                                                                _ = v178
                                                                                            case "\n\t":
                                                                                                v179 := v4[(v2 + ((v2 + v2) + v2))]
                                                                                                _ = v179
                                                                                            case "a":
                                                                                                v1 = v149
                                                                                            default:
                                                                                                v110 += (^(v110))
                                                                                            }
                                                                                        } else {
                                                                                            var v180 int = (fi65 >> ((v110) & 63))
                                                                                            switch v2 {
                                                                                            case "\n\t":
                                                                                                var v181 bool = true
                                                                                                _ = v181
                                                                                            case "a":
                                                                                                var v182 int = (fi65 / ((fi65) | 1))
                                                                                                _ = v182
                                                                                            case "hello":
                                                                                                var v183 uint = v110
                                                                                                _ = v183
                                                                                            default:
                                                                                                v3 = 42
                                                                                            }
                                                                                            v180 -= 32767
                                                                                            _ = v180
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        v110 = 65535
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v184 int64 = -9223372036854775808
                                                                                    _ = v184
                                                                                default:
                                                                                    var v185 uint16 = 42
                                                                                    _ = v185
                                                                                }
                                                                            } else {
                                                                                v149 = v1
                                                                                if true {
                                                                                    switch 2147483647 {
                                                                                    case 127:
                                                                                        _ = len(v4)
                                                                                    }
                                                                                    v68 -= int8(v83)
                                                                                }
                                                                                v3 += v68
                                                                            }
                                                                            _ = v149
                                                                        } else {
                                                                            v68 = (-(v3))
                                                                        }
                                                                        v110 = v110
                                                                        if true {
                                                                            switch "hello" {
                                                                            case "a":
                                                                                v110 += (v110 & v110)
                                                                            case "hello":
                                                                                var v186 bool = false
                                                                                _ = v186
                                                                            default:
                                                                                v187 := GenS11{F0: 42}
                                                                                _ = v187
                                                                            }
                                                                        } else {
                                                                            v2 = v2
                                                                            var v188 float32 = -2.25
                                                                            _ = v188
                                                                        }
                                                                    } else {
                                                                        v189 := GenS7{F0: 0, F1: (^(v28))}
                                                                        if false {
                                                                            v190 := GenS16{F0: 795}
                                                                            v113 += 42
                                                                            v68 -= (v66 * -128)
                                                                            _ = v190
                                                                        } else {
                                                                            v191 := GenS17{F0: (v28 << ((v110) & 63)), F1: 167}
                                                                            delete(v4, "\n\t")
                                                                            var v192 bool = false
                                                                            _ = v191
                                                                            _ = v192
                                                                        }
                                                                        v110 -= (v110 | v110)
                                                                        _ = v189
                                                                    }
                                                                default:
                                                                    v36 = append(v36, fi67)
                                                                }
                                                                var v193 bool = false
                                                                _ = v193
                                                            }
                                                            var v194 float64 = ((v1 - 1e10) / v1)
                                                            _ = v113
                                                            _ = v194
                                                        } else {
                                                            switch (v2 + (v2 + (v2 + "\n\t"))) {
                                                            case "a":
                                                                _ = len(v4)
                                                            case "\n\t":
                                                                _ = len(v36)
                                                            default:
                                                                v83 = (v83 >> ((225) & 63))
                                                            }
                                                            v195 := GenS17{F0: (v28 * 15), F1: 228}
                                                            fmt.Println(v28, v110, fi65)
                                                            v196 := GenS5{F0: 128, F1: (v68 * v68), F2: (v66 % ((v66) | 1))}
                                                            v1 = 1e10
                                                            if true {
                                                                v3 -= int8(fi65)
                                                            }
                                                            v197 := GenS10{F0: 2147483647}
                                                            _ = v195
                                                            _ = v196
                                                            _ = v197
                                                        }
                                                        if true {
                                                            var v198 int64 = 9223372036854775807
                                                            v198 += v198
                                                            _ = len(v4)
                                                            var v199 bool = (uint32(v110) < 2147483648)
                                                            switch 9223372036854775807 {
                                                            case 897:
                                                                if len(v36) >= 2 {
                                                                    v36 = v36[1:len(v36)]
                                                                }
                                                            }
                                                            _ = v198
                                                            _ = v199
                                                        }
                                                        delete(v4, v2)
                                                    } else {
                                                        v36 = append(v36, (fi65 & fi67))
                                                        v200 := GenS18{F0: fi67, F1: v68, F2: -367}
                                                        if ((!((fi65 % ((fi65) | 1)) <= (fi65 * (-(fi65))))) && false) {
                                                            v3 += 0
                                                            var v201 int32 = -1
                                                            var v202 bool = (v201 > v201)
                                                            var v203 uint32 = v83
                                                            switch int(v66) {
                                                            case 117:
                                                                var v204 int16 = v111
                                                                _ = v204
                                                            case 127:
                                                                v68 = (v68 >> (((v110 & 18446744073709551615)) & 63))
                                                            }
                                                            var v205 uint16 = v28
                                                            _ = v201
                                                            _ = v202
                                                            _ = v203
                                                            _ = v205
                                                        }
                                                        _ = v200
                                                    }
                                                case "hello":
                                                    _ = len(v36)
                                                }
                                                v206 := GenS8{F0: 167}
                                                _ = v112
                                                _ = v206
                                            }
                                        } else {
                                            if (v83 == v83) {
                                                var v207 float32 = (0.1 / 3.14159)
                                                v3 -= v68
                                                v3 -= (^(v3))
                                                var v208 int32 = 709
                                                _ = v207
                                                _ = v208
                                            } else {
                                                v68 -= v68
                                                if false {
                                                    var v209 int32 = 1
                                                    var v210 int32 = (v209 + v209)
                                                    if true {
                                                        fmt.Println(v28, v110)
                                                    } else {
                                                        var v211 bool = false
                                                        if len(v36) > 0 {
                                                            v36[0] = int(v68)
                                                        }
                                                        switch fi65 {
                                                        case 32767:
                                                            v111 -= -541
                                                        case -9223372036854775808:
                                                            var v212 int8 = 77
                                                            _ = v212
                                                        case 1:
                                                            var v213 uint16 = (v28 ^ uint16(v83))
                                                            _ = v213
                                                        default:
                                                            v214 := GenS1{F0: v210}
                                                            _ = v214
                                                        }
                                                        _ = v211
                                                    }
                                                    _ = v209
                                                    _ = v210
                                                }
                                                var v215 int8 = int8(v111)
                                                _ = v215
                                            }
                                            if (17 == 0) {
                                                v216 := GenS4{F0: (^(v110)), F1: v28}
                                                switch ((("" + v2) + v2) + v2) {
                                                case "hello":
                                                    v68 = 42
                                                case "\n\t":
                                                    if len(v36) >= 2 {
                                                        v36 = v36[1:len(v36)]
                                                    }
                                                }
                                                v217 := GenS19{F0: 1}
                                                fmt.Println(fi67)
                                                _ = v216
                                                _ = v217
                                            } else {
                                                var v218 string = v2
                                                var v219 int8 = v68
                                                _ = v218
                                                _ = v219
                                            }
                                        }
                                        v220 := GenS20{F0: (v83 | (v83 | v83)), F1: 42, F2: (^(v111))}
                                        _ = v110
                                        _ = v111
                                        _ = v220
                                        _ = v68
                                        _ = v83
                                    } else {
                                        var v221 uint64 = 0
                                        _ = v221
                                    }
                                    if (4294967295 != 65535) {
                                        var v222 float64 = ((v1 / (v1 + v1)) + float64((-2.25 + 0.1)))
                                        _ = v222
                                    }
                                    v223 := GenS21{F0: 42, F1: 65535, F2: 42}
                                    _ = v223
                                }
                                if false {
                                    var v224 uint16 = 255
                                    _ = v224
                                } else {
                                    if false {
                                        v2 = (("" + v2) + (v2 + v2))
                                    } else {
                                        if len(v36) > 0 {
                                            v36[0] = 32767
                                        }
                                    }
                                    if true {
                                        var v225 uint16 = (v28 + v28)
                                        _ = v225
                                    }
                                }
                                _ = len(v4)
                                switch v2 {
                                case "\n\t":
                                    if (true || (2 < 128)) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v226 string = "a"
                                            _ = v226
                                        case "\n\t":
                                            v227 := GenS22{F0: -1, F1: (fi65 / ((261) | 1)), F2: -1}
                                            _ = v227
                                        }
                                        v66 = v66
                                        v1 = float64(1.5)
                                    }
                                default:
                                    if (255 <= 2) {
                                        var v228 int16 = -423
                                        if true {
                                            var v229 uint16 = (v28 & (v28 * (^(v28))))
                                            if (((((0.1 - 1.5) + ((0.1 * 0.0) * 0.0)) + ((1.5 * 0.0) / -2.25)) - 1e10) >= (0.1 * -2.25)) {
                                                fmt.Println(v66, v3, v229)
                                                v229 = v28
                                            } else {
                                                v66 += (v3 + (^(v66)))
                                                var v230 int32 = -2147483648
                                                v228 = v228
                                                _ = v230
                                            }
                                            _ = v229
                                        } else {
                                            var v231 uint32 = 1
                                            _ = v231
                                        }
                                        if (2147483648 == 4294967295) {
                                            var v232 bool = (2 < 310)
                                            if v232 {
                                                delete(v4, v2)
                                            } else {
                                                var v233 uint32 = 42
                                                delete(v4, v2)
                                                v234 := GenS20{F0: (v233 / ((v233) | 1)), F1: 2, F2: v228}
                                                fmt.Println(v232)
                                                _ = v233
                                                _ = v234
                                            }
                                            var v235 int8 = 127
                                            _ = v232
                                            _ = v235
                                        } else {
                                            v228 -= 32767
                                            v228 = (-(v228))
                                            v236 := GenS19{F0: 42}
                                            fmt.Println(v228)
                                            _ = v236
                                        }
                                        var v237 uint8 = 191
                                        _ = v228
                                        _ = v237
                                    } else {
                                        if (!(42 >= 2)) {
                                            v238 := GenS21{F0: -2147483648, F1: 88, F2: -75}
                                            v239 := GenS23{F0: v28, F1: 242}
                                            switch v2 {
                                            case "foo bar":
                                                fmt.Println(fi65, fi65, fi67)
                                                var v240 float32 = -2.25
                                                var v241 string = ("foo bar" + v2)
                                                var v242 uint16 = (v28 | (v28 * 255))
                                                _ = v240
                                                _ = v241
                                                _ = v242
                                            case "\n\t":
                                                fmt.Println(v66, v1)
                                            case "a":
                                                v2 = v2
                                            }
                                            _ = v238
                                            _ = v239
                                        } else {
                                            fmt.Println(fi65)
                                            fmt.Println(v66, fi65)
                                        }
                                        if false {
                                            v28 = (v28 ^ v28)
                                        }
                                        _ = len(v4)
                                    }
                                }
                                _ = fi67
                            }
                            switch (fi65 << ((255) & 63)) {
                            case 1:
                                v1 = 0.1
                            case -2147483648:
                                var v243 uint32 = 2
                                _ = v243
                            default:
                                var v244 int64 = 266
                                _ = v244
                            }
                            _ = v66
                        }
                        _ = fi65
                    }
                }
            }
        case 35:
            var v245 uint32 = 626
            _ = v245
        default:
            switch 32767 {
            case -591:
                _ = len(v4)
            default:
                var v246 int = 550
                _ = v246
            }
        }
    } else {
        for fi247 := 0; fi247 < 14; fi247++ {
            for fi248 := 0; fi248 < 1; fi248++ {
                if (4294967295 == 2) {
                    var v249 int32 = -569
                    _ = len(v36)
                    _ = v249
                } else {
                    var v250 uint16 = v28
                    _ = v250
                }
                var v251 uint = 838
                var v252 int16 = 42
                v253 := GenS24{F0: (false && true), F1: v2}
                _ = fi248
                _ = v251
                _ = v252
                _ = v253
            }
            _ = fi247
        }
        var v254 uint64 = 1
        _ = v254
    }
    fmt.Println(v3, v1)
    v3 = -128
    v255 := GenS25{F0: (v28 << ((65535) & 63)), F1: 9223372036854775807}
    switch 4294967295 {
    case -395:
        for fi256 := 0; fi256 < 10; fi256++ {
            switch (((v2 + ((v2 + v2) + ("a" + "\n\t"))) + ((v2 + "") + v2)) + (("\n\t" + "\n\t") + "foo bar")) {
            case "":
                v257 := GenS26{F0: ((((v2 + ("a" + v2)) >= v2) && (!false)) || true)}
                _ = v257
            case "a":
                v258 := GenS8{F0: 1}
                _ = v258
            case "hello":
                var v259 int8 = v3
                _ = v259
            }
            if false {
                switch v2 {
                case "":
                    if true {
                        for fi260 := 0; fi260 < 20; fi260++ {
                            var v261 int16 = 935
                            v1 -= (0.1 / v1)
                            _ = fi260
                            _ = v261
                        }
                    } else {
                        for fi262 := 0; fi262 < 3; fi262++ {
                            if len(v36) >= 2 {
                                v36 = v36[1:len(v36)]
                            }
                            v2 = ((v2 + ((v2 + ("" + v2)) + v2)) + "hello")
                            _ = fi262
                        }
                        v4[v2] = "a"
                        if true {
                            switch 32768 {
                            case -884:
                                v1 -= (v1 + v1)
                            case 0:
                                for fi263 := 0; fi263 < 9; fi263++ {
                                    _ = len(v36)
                                    _ = fi263
                                }
                            case 1:
                                delete(v4, (("a" + ("\n\t" + v2)) + v2))
                            default:
                                v264 := GenS16{F0: 4294967295}
                                _ = v264
                            }
                            var v265 int16 = 0
                            v3 = (v3 % ((v3) | 1))
                            _ = v265
                        } else {
                            for fi266 := 0; fi266 < 18; fi266++ {
                                var v267 float32 = -2.25
                                v3 = v3
                                fmt.Println(fi256)
                                _ = fi266
                                _ = v267
                            }
                            fmt.Println(v3, v3, v28)
                        }
                    }
                case "foo bar":
                    v1 = 0.0
                default:
                    fmt.Println(v3, fi256, fi256)
                }
            }
            if true {
                if true {
                    if (-2147483648 != -2147483648) {
                        for fi268 := 0; fi268 < 4; fi268++ {
                            var v269 int16 = 42
                            if false {
                                var v270 int16 = v269
                                var v271 bool = ((^(v3)) > v3)
                                _ = v270
                                _ = v271
                            } else {
                                v1 = float64(-2.25)
                                var v272 uint16 = (v28 * 114)
                                fmt.Println(v272, v3)
                                _ = v272
                            }
                            delete(v4, v2)
                            _ = fi268
                            _ = v269
                        }
                        v36 = append(v36, fi256)
                        v1 -= (v1 - 1e10)
                    }
                }
                v3 += (v3 + v3)
                var v273 int16 = -32768
                _ = v273
            }
            _ = fi256
        }
    case 4294967295:
        if len(v36) >= 2 {
            v36 = v36[1:len(v36)]
        }
    case 127:
        if (-736 > 0) {
            var v274 float64 = v1
            v274 += 1e10
            _ = v274
        } else {
            v28 -= (v28 + (v28 % ((v28) | 1)))
        }
    }
    v2 = v2
    fmt.Println(v28, v28, v28)
    v275 := GenS27{F0: 0, F1: 0, F2: (^(v28))}
    v276 := GenS27{F0: -1, F1: 42, F2: (v28 | 482)}
    v1 = (v1 * v1)
    v3 = int8(v28)
    for fi277 := 0; fi277 < 1; fi277++ {
        v278 := GenS28{F0: (v3 ^ v3), F1: 9223372036854775807}
        _ = fi277
        _ = v278
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v255
    _ = v275
    _ = v276
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
