package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 int16; F1 int64 }
type GenS13 struct { F0 int32; F1 uint16; F2 int16 }
type GenS14 struct { F0 uint16; F1 string; F2 uint16 }
type GenS15 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int = -9223372036854775808
    var v2 float32 = 0.1
    v3 := make(map[int]uint32)
    var v4 uint16 = 0
    var v5 uint16 = (^(v4))
    v6 := make(map[uint32]bool)
    var v7 uint16 = v5
    var v8 int32 = 894
    fmt.Println(v4)
    if (true || false) {
        v4 = (v5 & v5)
        v7 = v4
        v4 -= (v7 << ((255) & 63))
    } else {
        var v9 int8 = -1
        for fi10 := 0; fi10 < 18; fi10++ {
            switch "hello" {
            case "\n\t":
                v11 := []int{fi10, (v1 - v1)}
                _ = v11
            case "a":
                fmt.Println(v1, v9, fi10)
            case "":
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 int32 = (v8 * int32(fi12))
                    _ = fi12
                    _ = v13
                }
            default:
                var v14 bool = (2 > 186)
                _ = v14
            }
            _ = fi10
        }
        _ = v9
    }
    switch "a" {
    case "\n\t":
        var v15 int32 = int32(v4)
        _ = v15
    }
    var v16 uint16 = (v4 - 240)
    fmt.Println(v2, v7, v5)
    if (!false) {
        for fi17 := 0; fi17 < 9; fi17++ {
            switch "\n\t" {
            case "a":
                if (!true) {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if false {
                            var v19 float64 = 0.0
                            _ = len(v3)
                            _ = v19
                        }
                        _ = fi18
                    }
                    fmt.Println(v16, v1)
                    var v20 uint64 = 42
                    _ = v20
                } else {
                    fmt.Println(v4, v16)
                    switch fi17 {
                    case 32768:
                        var v21 float64 = -2.25
                        _ = v21
                    case 128:
                        var v22 uint8 = 2
                        _ = v22
                    }
                }
            case "\n\t":
                switch v1 {
                case 1:
                    var v23 uint = 1
                    _ = v23
                case 32768:
                    v24 := GenS1{F0: 42, F1: 100, F2: 9223372036854775807}
                    _ = v24
                }
            case "hello":
                for fi25 := 0; fi25 < 5; fi25++ {
                    if true {
                        v16 = (v7 % ((796) | 1))
                        switch v1 {
                        case -984:
                            if false {
                                v2 = 3.14159
                            } else {
                                if (int32(v1) <= (v8 ^ (^(v8)))) {
                                    fmt.Println(v2, v1, v5)
                                    v8 -= v8
                                    switch "\n\t" {
                                    case "a":
                                        v26 := GenS2{F0: fi17, F1: 2147483647, F2: v5}
                                        _ = v26
                                    case "hello":
                                        v27 := GenS1{F0: 244, F1: -41, F2: 9223372036854775807}
                                        _ = v27
                                    default:
                                        var v28 uint32 = 2147483648
                                        _ = v28
                                    }
                                } else {
                                    var v29 uint16 = v4
                                    v30 := []int{}
                                    v29 = (v29 % ((v5) | 1))
                                    _ = v29
                                    _ = v30
                                }
                            }
                        case 1:
                            v7 = (v7 * 65535)
                        case 122:
                            switch ("foo bar" + ("\n\t" + ("" + ""))) {
                            case "":
                                if (-9223372036854775808 >= -9223372036854775808) {
                                    v31 := GenS2{F0: -447, F1: 160, F2: v4}
                                    _ = v31
                                } else {
                                    v32 := GenS1{F0: 42, F1: -59, F2: -624}
                                    _ = v32
                                }
                                var v33 float32 = ((v2 / (0.0 - -2.25)) / float32(0.1))
                                _ = v33
                            case "foo bar":
                                v7 = v7
                            default:
                                v34 := []int{(v1 ^ fi25), fi17, (v1 & (fi17 << ((18446744073709551615) & 63)))}
                                _ = v34
                            }
                        default:
                            var v35 uint64 = 42
                            _ = v35
                        }
                        var v36 int16 = -32768
                        _ = v36
                    } else {
                        var v37 uint64 = 2
                        v38 := GenS2{F0: fi17, F1: (v8 - (v8 - 2147483647)), F2: 255}
                        _ = v37
                        _ = v38
                    }
                    _ = fi25
                }
            }
            _ = fi17
        }
        if false {
            if true {
                var v39 uint64 = 18446744073709551615
                _ = v39
            }
            var v40 int64 = 9223372036854775807
            _ = v40
        } else {
            var v41 int64 = 42
            v8 -= v8
            _ = v41
        }
        v42 := GenS1{F0: 2, F1: 0, F2: 1}
        _ = v42
    }
    var v43 uint32 = 666
    if true {
        _ = len(v3)
        var v44 int32 = -298
        _ = v44
    }
    fmt.Println(v43, v7)
    if true {
        if false {
            v45 := GenS3{F0: (("foo bar" + "\n\t") + (("foo bar" + "foo bar") + ("foo bar" + ("hello" + "hello")))), F1: 0}
            _ = v45
        } else {
            var v46 uint16 = 679
            v47 := []int{int(v43), (^(v1)), 128}
            var v48 uint16 = (v7 * v5)
            _ = v46
            _ = v47
            _ = v48
        }
    }
    v49 := GenS3{F0: "", F1: 282}
    fmt.Println(v4, v5, v8)
    for fi50 := 0; fi50 < 13; fi50++ {
        var v51 int64 = -1
        if false {
            for fi52 := 0; fi52 < 8; fi52++ {
                v1 = 127
                if false {
                    var v53 uint16 = 68
                    var v54 string = ""
                    fmt.Println(v5, v51, v4)
                    _ = v53
                    _ = v54
                } else {
                    fmt.Println(v4, v51)
                    var v55 int = (v1 >> ((42) & 63))
                    v43 += v43
                    _ = v55
                }
                switch fi50 {
                case -2147483648:
                    var v56 int16 = -394
                    _ = v56
                case 128:
                    fmt.Println(v51, v51, v2)
                case -572:
                    var v57 int16 = 42
                    _ = v57
                }
                _ = fi52
            }
            if false {
                v58 := GenS3{F0: "", F1: 1}
                _ = v58
            }
            v8 = v8
        } else {
            switch "\n\t" {
            case "":
                if (1 == 1) {
                    v4 = (v7 & v7)
                    v59 := []int{(v1 | (v1 + v1)), -1, (-(v1)), v1}
                    _ = v59
                } else {
                    if ((1e10 == 1e10) || (((-2.25 - 1e10) + (0.0 + 1e10)) >= 1e10)) {
                        var v60 int16 = 0
                        v7 = 65535
                        _ = v60
                    }
                    v7 = v5
                    if (true && ((-(v51)) <= -9223372036854775808)) {
                        for fi61 := 0; fi61 < 18; fi61++ {
                            switch "hello" {
                            case "\n\t":
                                if ((v51 ^ v51) >= -9223372036854775808) {
                                    v62 := []int{fi61, (v1 ^ fi50), fi50, fi50}
                                    _ = v62
                                }
                            case "a":
                                v63 := GenS4{F0: true, F1: v51}
                                _ = v63
                            case "hello":
                                v64 := GenS1{F0: 59, F1: 100, F2: 9223372036854775807}
                                _ = v64
                            }
                            switch 1 {
                            case -520:
                                fmt.Println(fi50, fi50)
                            case 127:
                                delete(v3, fi50)
                            default:
                                v65 := GenS5{F0: v1}
                                _ = v65
                            }
                            v66 := []int{}
                            var v67 uint = 2
                            _ = fi61
                            _ = v66
                            _ = v67
                        }
                    }
                }
            default:
                if true {
                    switch -759 {
                    case 371:
                        v68 := []int{v1, (^(v1))}
                        _ = v68
                    case -9223372036854775808:
                        if false {
                            fmt.Println(v1, v7, v4)
                            for fi69 := 0; fi69 < 20; fi69++ {
                                if true {
                                    v7 = (v4 << ((4294967295) & 63))
                                    if true {
                                        var v70 uint32 = v43
                                        switch fi69 {
                                        case 2147483647:
                                            v5 -= 65535
                                            _ = len(v6)
                                        case 739:
                                            switch ("" + ("a" + ("" + ("foo bar" + "")))) {
                                            case "":
                                                v71 := GenS6{F0: -2147483648, F1: v8}
                                                _ = v71
                                            case "a":
                                                var v72 bool = (true || (v8 < -2147483648))
                                                v8 -= v8
                                                _ = v72
                                            case "foo bar":
                                                v1 = (fi69 << ((18446744073709551615) & 63))
                                            }
                                        case 127:
                                            v73 := []int{4294967295, v1}
                                            _ = v73
                                        }
                                        _ = v70
                                    }
                                    if true {
                                        v74 := GenS1{F0: 2, F1: 0, F2: 498}
                                        v1 = fi50
                                        var v75 string = "a"
                                        _ = v74
                                        _ = v75
                                    }
                                    var v76 uint16 = (v16 >> ((413) & 63))
                                    if true {
                                        v7 = (v7 / (((^(v5))) | 1))
                                        switch fi50 {
                                        case 32767:
                                            fmt.Println(v7, v43)
                                        case -9223372036854775808:
                                            if ((1 <= 18446744073709551615) || true) {
                                                if false {
                                                    var v77 int64 = v51
                                                    v78 := []int{-1, 4294967295, fi50, (fi50 / ((9223372036854775807) | 1))}
                                                    v4 -= 785
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    switch (^(v1)) {
                                                    case 0:
                                                        v79 := v6[(v43 | (v43 << ((18446744073709551615) & 63)))]
                                                        _ = v79
                                                    case 2147483647:
                                                        var v80 bool = (uint32(v16) >= 2)
                                                        _ = v80
                                                    default:
                                                        var v81 uint = 429
                                                        _ = v81
                                                    }
                                                    fmt.Println(v4, v8, v4)
                                                }
                                            } else {
                                                var v82 int8 = 100
                                                v83 := []int{}
                                                var v84 uint64 = 42
                                                if (375 >= 42) {
                                                    v1 = (fi50 / ((fi69) | 1))
                                                    var v85 bool = (-628 > 2147483647)
                                                    v82 = (v82 % (((v82 + v82)) | 1))
                                                    switch (fi69 ^ (-(fi50))) {
                                                    case 1:
                                                        fmt.Println(fi69)
                                                    case 0:
                                                        v86 := GenS6{F0: 42, F1: v8}
                                                        _ = v86
                                                    case 510:
                                                        fmt.Println(v5)
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            if true {
                                                                v87 := GenS7{F0: 42}
                                                                if ((!true) || v85) {
                                                                    v1 = v1
                                                                    fmt.Println(v51, v8)
                                                                    var v88 string = (("foo bar" + "a") + "foo bar")
                                                                    _ = v88
                                                                }
                                                                v16 = v76
                                                                _ = v87
                                                            } else {
                                                                v8 = 2147483647
                                                                if v85 {
                                                                    if ((v8 >> ((4294967295) & 63)) > (v8 ^ 257)) {
                                                                        var v89 int8 = 1
                                                                        var v90 uint = 65535
                                                                        _ = v89
                                                                        _ = v90
                                                                    } else {
                                                                        var v91 int16 = 42
                                                                        var v92 int = v1
                                                                        v76 += v76
                                                                        _ = v91
                                                                        _ = v92
                                                                    }
                                                                    switch 9223372036854775807 {
                                                                    case 128:
                                                                        v93 := GenS6{F0: 761, F1: (v8 >> ((4294967295) & 63))}
                                                                        _ = v93
                                                                    case 9223372036854775807:
                                                                        switch ("" + "foo bar") {
                                                                        case "\n\t":
                                                                            if true {
                                                                                var v94 uint = 4294967295
                                                                                var v95 uint64 = 42
                                                                                _ = v94
                                                                                _ = v95
                                                                            } else {
                                                                                var v96 int16 = -32768
                                                                                var v97 uint8 = 42
                                                                                v98 := GenS2{F0: fi69, F1: int32(v82), F2: (^(v76))}
                                                                                _ = v96
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                        }
                                                                    case 909:
                                                                        v43 = v43
                                                                    }
                                                                    var v99 uint64 = (v84 + uint64(v43))
                                                                    _ = v99
                                                                }
                                                                var v100 uint8 = 165
                                                                v85 = true
                                                                if (v85 && (v76 > v76)) {
                                                                    v43 = (v43 / ((v43) | 1))
                                                                    if (v84 >= (^(v84))) {
                                                                        fmt.Println(v2, v84)
                                                                        if v85 {
                                                                            v8 = (v8 & v8)
                                                                            fmt.Println(v4, v5)
                                                                            v101 := GenS3{F0: "hello", F1: 1}
                                                                            if v85 {
                                                                                var v102 uint64 = 943
                                                                                _ = v102
                                                                            }
                                                                            _ = v101
                                                                        } else {
                                                                            var v103 int64 = v51
                                                                            _ = v103
                                                                        }
                                                                        switch ("hello" + "a") {
                                                                        case "":
                                                                            if v85 {
                                                                                if len(v83) >= 2 {
                                                                                    v83 = v83[1:len(v83)]
                                                                                }
                                                                                var v104 uint64 = v84
                                                                                v1 += 2147483647
                                                                                _ = v104
                                                                            }
                                                                        case "hello":
                                                                            v105 := GenS7{F0: 255}
                                                                            v83 = append(v83, (fi69 % ((42) | 1)))
                                                                            _ = v105
                                                                        case "a":
                                                                            var v106 uint8 = 2
                                                                            _ = v106
                                                                        default:
                                                                            if (938 != (v84 * (v84 & (^(v84))))) {
                                                                                switch (fi50 | (fi69 % ((128) | 1))) {
                                                                                case 1:
                                                                                    fmt.Println(v82, v5, v16)
                                                                                case -579:
                                                                                    v107 := GenS7{F0: 927}
                                                                                    _ = v107
                                                                                case 285:
                                                                                    v85 = v85
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    if (65535 != 32767) {
                                                                                        v5 = 65535
                                                                                        if (v85 && v85) {
                                                                                            v76 += v16
                                                                                            switch (-(fi50)) {
                                                                                            case 9223372036854775807:
                                                                                                v76 = v4
                                                                                            case 1:
                                                                                                v43 = v43
                                                                                            default:
                                                                                                v108 := GenS8{F0: v51, F1: ""}
                                                                                                _ = v108
                                                                                            }
                                                                                            var v109 uint32 = v43
                                                                                            _ = v109
                                                                                        }
                                                                                    } else {
                                                                                        var v110 int32 = (v8 & v8)
                                                                                        _ = v110
                                                                                    }
                                                                                }
                                                                                var v111 uint8 = (v100 | 1)
                                                                                _ = v111
                                                                            } else {
                                                                                if v85 {
                                                                                    v84 = v84
                                                                                    var v112 int32 = -2147483648
                                                                                    v113 := v3[4294967295]
                                                                                    _ = v112
                                                                                    _ = v113
                                                                                } else {
                                                                                    var v114 float64 = 3.14159
                                                                                    if len(v83) > 0 {
                                                                                        v83[0] = (fi69 >> ((18446744073709551615) & 63))
                                                                                    }
                                                                                    delete(v3, (-(fi50)))
                                                                                    _ = v114
                                                                                }
                                                                                var v115 int64 = v51
                                                                                _ = len(v83)
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                    }
                                                                    if ((v51 % ((-9223372036854775808) | 1)) == (v51 / ((v51) | 1))) {
                                                                        v43 += v43
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            var v116 string = "hello"
                                                                            _ = v116
                                                                        case "\n\t":
                                                                            v117 := GenS9{F0: v43, F1: ((-(v8)) <= v8)}
                                                                            _ = v117
                                                                        default:
                                                                            if (((((3.14159 / 3.14159) - float64(v2)) * 1.5) - (1e10 - (((float64(v2) * 0.0) / 0.1) + ((1e10 * -2.25) / float64(v2))))) == float64(0.0)) {
                                                                                v4 = uint16(v84)
                                                                                if (fi50 != int(v8)) {
                                                                                    if v85 {
                                                                                        var v118 float64 = (0.1 * 0.1)
                                                                                        var v119 uint16 = (v7 & 717)
                                                                                        _ = v118
                                                                                        _ = v119
                                                                                    } else {
                                                                                        var v120 uint = 793
                                                                                        if false {
                                                                                            switch fi50 {
                                                                                            case -9223372036854775808:
                                                                                                if v85 {
                                                                                                    v121 := GenS10{F0: (v84 | v84)}
                                                                                                    v122 := GenS8{F0: int64(fi50), F1: ""}
                                                                                                    if (((float32(0.1) * (((v2 - v2) * float32(float64(((float32(((float64(v2) - 0.1) - 1e10)) - v2) * v2)))) * 0.0)) - (1e10 + v2)) == 1.5) {
                                                                                                        delete(v3, (^(v1)))
                                                                                                        switch 9223372036854775807 {
                                                                                                        case 1:
                                                                                                            switch "" {
                                                                                                            case "hello":
                                                                                                                if v85 {
                                                                                                                    v16 = 42
                                                                                                                    _ = len(v6)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        if ((v100 * v100) == v100) {
                                                                                                            fmt.Println(v8, v85)
                                                                                                            v82 -= (v82 / ((v82) | 1))
                                                                                                            v83 = append(v83, v1)
                                                                                                            var v123 uint32 = (v43 & (v43 ^ (v43 - v43)))
                                                                                                            _ = v123
                                                                                                        } else {
                                                                                                            var v124 uint32 = v43
                                                                                                            _ = v124
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v121
                                                                                                    _ = v122
                                                                                                } else {
                                                                                                    var v125 uint32 = v43
                                                                                                    _ = v125
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                if v85 {
                                                                                                    fmt.Println(v100)
                                                                                                    v85 = false
                                                                                                } else {
                                                                                                    v8 -= 42
                                                                                                }
                                                                                            }
                                                                                            _ = len(v83)
                                                                                        } else {
                                                                                            v126 := GenS6{F0: (v8 % ((v8) | 1)), F1: 0}
                                                                                            switch 4294967295 {
                                                                                            case 96:
                                                                                                if (!(!true)) {
                                                                                                    var v127 uint = v120
                                                                                                    switch (fi69 / ((v1) | 1)) {
                                                                                                    case -9223372036854775808:
                                                                                                        v128 := GenS1{F0: (v100 ^ (v100 | (v100 ^ 0))), F1: 17, F2: v51}
                                                                                                        _ = v128
                                                                                                    case 4294967295:
                                                                                                        switch ((("\n\t" + ((("\n\t" + "a") + "foo bar") + ("" + ("\n\t" + "")))) + ("" + "")) + "\n\t") {
                                                                                                        case "":
                                                                                                            fmt.Println(v16, fi50, v76)
                                                                                                        case "a":
                                                                                                            v120 -= 589
                                                                                                            var v129 int64 = (v51 ^ (v51 + (v51 % ((int64(v8)) | 1))))
                                                                                                            v130 := v3[(fi69 - (v1 >> ((255) & 63)))]
                                                                                                            if v85 {
                                                                                                                switch "foo bar" {
                                                                                                                case "":
                                                                                                                    if (4294967295 > 0) {
                                                                                                                        fmt.Println(v100, v82)
                                                                                                                        fmt.Println(v127, v1)
                                                                                                                    } else {
                                                                                                                        if len(v83) >= 2 {
                                                                                                                            v83 = v83[1:len(v83)]
                                                                                                                        }
                                                                                                                        v6[147] = (!v85)
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v5 += 32767
                                                                                                                default:
                                                                                                                    var v131 uint8 = v100
                                                                                                                    _ = v131
                                                                                                                }
                                                                                                            }
                                                                                                            v7 = (v76 ^ 32767)
                                                                                                            if (("hello" + "") == "") {
                                                                                                                v83 = append(v83, v1)
                                                                                                                if (!v85) {
                                                                                                                    fmt.Println(fi69, v1)
                                                                                                                    v132 := GenS11{F0: "\n\t", F1: "hello"}
                                                                                                                    v133 := v6[(^(v130))]
                                                                                                                    _ = v132
                                                                                                                    _ = v133
                                                                                                                }
                                                                                                                var v134 int16 = 42
                                                                                                                if (true || (v127 > (^(v127)))) {
                                                                                                                    if ((v127 >> ((v120) & 63)) < 2) {
                                                                                                                        v135 := GenS1{F0: uint8(v84), F1: -69, F2: v129}
                                                                                                                        _ = v135
                                                                                                                    }
                                                                                                                    v5 = (^(v16))
                                                                                                                }
                                                                                                                if false {
                                                                                                                    v76 -= (v7 & v4)
                                                                                                                }
                                                                                                                v129 = (^(v129))
                                                                                                                _ = v134
                                                                                                            }
                                                                                                            _ = v129
                                                                                                            _ = v130
                                                                                                        case "foo bar":
                                                                                                            var v136 float32 = v2
                                                                                                            _ = v136
                                                                                                        }
                                                                                                    case 2147483647:
                                                                                                        if true {
                                                                                                            v83 = append(v83, (fi69 % ((9223372036854775807) | 1)))
                                                                                                            var v137 string = ""
                                                                                                            var v138 bool = v85
                                                                                                            _ = v137
                                                                                                            _ = v138
                                                                                                        } else {
                                                                                                            var v139 int8 = (v82 / ((int8(v120)) | 1))
                                                                                                            _ = v139
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v127
                                                                                                } else {
                                                                                                    v51 -= -9223372036854775808
                                                                                                    v4 -= v5
                                                                                                    v51 = (v51 - (v51 * v51))
                                                                                                }
                                                                                            case 127:
                                                                                                var v140 uint32 = v43
                                                                                                fmt.Println(v1)
                                                                                                _ = v140
                                                                                            default:
                                                                                                v7 -= v4
                                                                                            }
                                                                                            var v141 float32 = -2.25
                                                                                            var v142 float64 = 0.0
                                                                                            v51 = -1
                                                                                            _ = v126
                                                                                            _ = v141
                                                                                            _ = v142
                                                                                        }
                                                                                        _ = v120
                                                                                    }
                                                                                    fmt.Println(v76, v1)
                                                                                    v51 += v51
                                                                                }
                                                                                v143 := GenS12{F0: 1000, F1: v51}
                                                                                if v85 {
                                                                                    v4 += v7
                                                                                    if v85 {
                                                                                        if len(v83) >= 2 {
                                                                                            v83 = v83[1:len(v83)]
                                                                                        }
                                                                                        if v85 {
                                                                                            var v144 uint16 = 1
                                                                                            _ = v144
                                                                                        } else {
                                                                                            if v85 {
                                                                                                switch (-(v1)) {
                                                                                                case -1:
                                                                                                    var v145 int8 = 100
                                                                                                    _ = v145
                                                                                                case 127:
                                                                                                    var v146 uint64 = v84
                                                                                                    _ = v146
                                                                                                default:
                                                                                                    if (v84 > (v84 ^ (^(v84)))) {
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            if v85 {
                                                                                                                v7 -= 597
                                                                                                                if v85 {
                                                                                                                    v4 += uint16(v100)
                                                                                                                    v147 := GenS5{F0: (fi69 / (((fi69 ^ int(v43))) | 1))}
                                                                                                                    fmt.Println(v4, fi50, v43)
                                                                                                                    _ = v147
                                                                                                                }
                                                                                                                switch "\n\t" {
                                                                                                                case "a":
                                                                                                                    var v148 bool = (v84 != 0)
                                                                                                                    var v149 string = ("foo bar" + "hello")
                                                                                                                    fmt.Println(fi50, v5)
                                                                                                                    _ = v148
                                                                                                                    _ = v149
                                                                                                                default:
                                                                                                                    fmt.Println(v2, v2, v51)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(fi69)
                                                                                                    } else {
                                                                                                        if (v43 == (v43 >> ((0) & 63))) {
                                                                                                            if (v5 >= 15) {
                                                                                                                _ = len(v6)
                                                                                                                fmt.Println(v5)
                                                                                                            }
                                                                                                            fmt.Println(v2, v1, v2)
                                                                                                            _ = len(v6)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if v85 {
                                                                                                var v150 uint16 = v76
                                                                                                var v151 float64 = ((1e10 - ((1e10 / (-2.25 * 3.14159)) / float64(-2.25))) / 1e10)
                                                                                                v152 := GenS9{F0: v43, F1: v85}
                                                                                                _ = v150
                                                                                                _ = v151
                                                                                                _ = v152
                                                                                            }
                                                                                            v8 += (-(v8))
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = len(v83)
                                                                                _ = v143
                                                                            }
                                                                        }
                                                                        if (v43 == uint32(v4)) {
                                                                            if v85 {
                                                                                if v85 {
                                                                                    var v153 int16 = 42
                                                                                    _ = v153
                                                                                }
                                                                                if v85 {
                                                                                    v154 := GenS3{F0: ("\n\t" + ("" + "a")), F1: 428}
                                                                                    var v155 float64 = 1.5
                                                                                    _ = v154
                                                                                    _ = v155
                                                                                } else {
                                                                                    v8 = int32(v82)
                                                                                    if true {
                                                                                        if len(v83) > 0 {
                                                                                            v83[0] = (fi69 % (((v1 / ((int(v43)) | 1))) | 1))
                                                                                        }
                                                                                        fmt.Println(v2, v16)
                                                                                    } else {
                                                                                        v84 -= 2
                                                                                        var v156 uint32 = uint32(v5)
                                                                                        _ = v156
                                                                                    }
                                                                                    v83 = append(v83, -416)
                                                                                }
                                                                            } else {
                                                                                var v157 float64 = ((0.0 - float64(float32(0.0))) - (float64(-2.25) - 3.14159))
                                                                                _ = v157
                                                                            }
                                                                        } else {
                                                                            var v158 bool = (true || v85)
                                                                            v159 := GenS13{F0: v8, F1: (v7 / ((v76) | 1)), F2: 760}
                                                                            v1 -= int(v5)
                                                                            _ = v158
                                                                            _ = v159
                                                                        }
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v160 uint = 65535
                                                                            _ = v160
                                                                        default:
                                                                            if ((^(v84)) != 235) {
                                                                                v5 = 926
                                                                                _ = len(v3)
                                                                                if false {
                                                                                    if true {
                                                                                        var v161 float32 = (3.14159 - v2)
                                                                                        _ = v161
                                                                                    } else {
                                                                                        if (!false) {
                                                                                            var v162 int16 = int16(fi69)
                                                                                            v83 = append(v83, int(v162))
                                                                                            _ = v162
                                                                                        }
                                                                                        if ((v5 & v76) >= v16) {
                                                                                            v8 = (v8 % (((v8 % ((1) | 1))) | 1))
                                                                                            v8 = v8
                                                                                        }
                                                                                    }
                                                                                    if v85 {
                                                                                        if (2 < 1) {
                                                                                            v84 = v84
                                                                                            var v163 int32 = (v8 * v8)
                                                                                            if v85 {
                                                                                                v43 += v43
                                                                                            }
                                                                                            v164 := GenS8{F0: (v51 * -9223372036854775808), F1: "hello"}
                                                                                            _ = v163
                                                                                            _ = v164
                                                                                        }
                                                                                    } else {
                                                                                        v7 -= v76
                                                                                        switch (fi50 & fi50) {
                                                                                        case 42:
                                                                                            var v165 uint64 = (^(v84))
                                                                                            _ = v165
                                                                                        default:
                                                                                            var v166 int64 = (v51 & -1)
                                                                                            _ = v166
                                                                                        }
                                                                                    }
                                                                                    var v167 uint8 = (v100 & (v100 >> ((673) & 63)))
                                                                                    _ = v167
                                                                                } else {
                                                                                    if v85 {
                                                                                        v85 = (((v43 / ((v43) | 1)) == (v43 ^ (v43 + (^(v43))))) && v85)
                                                                                        var v168 float32 = float32(0.0)
                                                                                        var v169 int64 = -139
                                                                                        _ = v168
                                                                                        _ = v169
                                                                                    }
                                                                                    _ = len(v3)
                                                                                    v170 := GenS6{F0: 1, F1: (v8 | (v8 ^ (v8 - (v8 | (v8 & -445)))))}
                                                                                    if v85 {
                                                                                        if ((-2.25 * (-2.25 / (3.14159 - -2.25))) <= float64(v2)) {
                                                                                            switch (("" + "a") + ((("foo bar" + "foo bar") + "foo bar") + "hello")) {
                                                                                            case "a":
                                                                                                if ((v1 >> ((4294967295) & 63)) > fi69) {
                                                                                                    var v171 int16 = 907
                                                                                                    _ = v171
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v16 += uint16(v84)
                                                                                            case "foo bar":
                                                                                                fmt.Println(fi50, v51)
                                                                                            default:
                                                                                                if true {
                                                                                                    var v172 int16 = 42
                                                                                                    _ = v172
                                                                                                } else {
                                                                                                    if v85 {
                                                                                                        switch (-(v1)) {
                                                                                                        case -752:
                                                                                                            fmt.Println(v43, v76, v8)
                                                                                                        default:
                                                                                                            var v173 uint16 = v4
                                                                                                            _ = v173
                                                                                                        }
                                                                                                        if len(v83) >= 2 {
                                                                                                            v83 = v83[1:len(v83)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        v76 += v16
                                                                                                        v174 := GenS7{F0: 65535}
                                                                                                        _ = v174
                                                                                                    }
                                                                                                    switch (fi69 * 42) {
                                                                                                    case 127:
                                                                                                        var v175 int = v1
                                                                                                        _ = v175
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v176 := GenS14{F0: 0, F1: "\n\t", F2: (v16 << ((uint(v4)) & 63))}
                                                                                            if true {
                                                                                                var v177 uint16 = (^(v76))
                                                                                                v83 = append(v83, v1)
                                                                                                _ = v177
                                                                                            }
                                                                                            v51 -= (-(v51))
                                                                                            v1 += fi69
                                                                                            _ = v176
                                                                                        }
                                                                                        v7 = (^(v76))
                                                                                    }
                                                                                    fmt.Println(v43, v8, v84)
                                                                                    _ = v170
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v100
                                                            }
                                                        case "hello":
                                                            v178 := GenS7{F0: 0}
                                                            _ = v178
                                                        case "":
                                                            var v179 int16 = 1
                                                            _ = v179
                                                        default:
                                                            var v180 uint = 854
                                                            _ = v180
                                                        }
                                                        v181 := v3[(fi50 | v1)]
                                                        v85 = ((^(v43)) != (v43 % ((2147483647) | 1)))
                                                        _ = v181
                                                    default:
                                                        if (162 >= v8) {
                                                            fmt.Println(fi69)
                                                            fmt.Println(v51)
                                                        }
                                                    }
                                                    _ = v85
                                                }
                                                fmt.Println(v1, v16)
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                            }
                                        case 347:
                                            var v182 string = (("a" + "foo bar") + "foo bar")
                                            var v183 float64 = (1e10 / (((3.14159 / ((-2.25 / (1.5 - 3.14159)) / -2.25)) - 1.5) - (1.5 / ((1e10 - 1.5) * 0.1))))
                                            _ = v182
                                            _ = v183
                                        default:
                                            var v184 int16 = 32767
                                            _ = v184
                                        }
                                    }
                                    _ = v76
                                }
                                v3[(v1 >> ((71) & 63))] = (v43 & 14)
                                _ = fi69
                            }
                            var v185 string = "hello"
                            _ = v185
                        } else {
                            v3[(v1 * (v1 >> ((2) & 63)))] = v43
                            if false {
                                var v186 int32 = (v8 + v8)
                                v187 := GenS4{F0: (1 > -62), F1: 9223372036854775807}
                                _ = v186
                                _ = v187
                            } else {
                                if ((250 >= 1) || (1e10 < v2)) {
                                    var v188 int16 = -541
                                    switch (v1 * v1) {
                                    case 32768:
                                        for fi189 := 0; fi189 < 9; fi189++ {
                                            switch 127 {
                                            case 32768:
                                                if ((true || (4294967295 <= 1)) && (false || (v8 <= 42))) {
                                                    v7 = v7
                                                    fmt.Println(v7, v51, v8)
                                                } else {
                                                    v190 := GenS15{F0: true, F1: 42}
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v51 += v51
                                                    case "":
                                                        switch ((((("" + "") + (("a" + "foo bar") + "foo bar")) + "a") + "") + "hello") {
                                                        case "":
                                                            v16 -= v4
                                                        case "a":
                                                            v191 := GenS13{F0: v8, F1: 2, F2: (v188 - -1)}
                                                            _ = v191
                                                        }
                                                    default:
                                                        fmt.Println(v8)
                                                    }
                                                    var v192 int = fi50
                                                    var v193 string = "foo bar"
                                                    if (1e10 > ((float64(v2) * 1e10) * 0.1)) {
                                                        v194 := GenS8{F0: (v51 - v51), F1: v193}
                                                        v8 = (v8 - -450)
                                                        _ = v194
                                                    } else {
                                                        v188 -= (^(v188))
                                                        v195 := v6[757]
                                                        if v195 {
                                                            var v196 uint8 = 42
                                                            v197 := GenS5{F0: v1}
                                                            v2 = v2
                                                            _ = v196
                                                            _ = v197
                                                        } else {
                                                            fmt.Println(v193, v5, v195)
                                                        }
                                                        _ = v195
                                                    }
                                                    _ = v190
                                                    _ = v192
                                                    _ = v193
                                                }
                                            }
                                            _ = fi189
                                        }
                                    }
                                    _ = v188
                                } else {
                                    fmt.Println(v43, v43, v5)
                                }
                                v4 = v7
                            }
                            switch (fi50 / (((fi50 | -9223372036854775808)) | 1)) {
                            case 128:
                                v1 = -675
                            case 9223372036854775807:
                                fmt.Println(v16, v7, v51)
                            case 32767:
                                fmt.Println(v5, v2)
                            }
                        }
                    }
                }
            }
            v51 = (v51 - int64(fi50))
        }
        _ = fi50
        _ = v51
    }
    for fi198 := 0; fi198 < 4; fi198++ {
        v4 = v16
        for fi199 := 0; fi199 < 10; fi199++ {
            var v200 float64 = 0.0
            if true {
                var v201 int32 = -134
                v202 := GenS15{F0: (-1 > int64(v8)), F1: 42}
                _ = v201
                _ = v202
            } else {
                v43 -= v43
            }
            _ = fi199
            _ = v200
        }
        var v203 bool = ((1 != -32768) && (false || true))
        _ = fi198
        _ = v203
    }
    for fi204 := 0; fi204 < 15; fi204++ {
        for fi205 := 0; fi205 < 13; fi205++ {
            if (42 == 127) {
                fmt.Println(fi204)
                v206 := v6[2147483647]
                _ = v206
            } else {
                switch "foo bar" {
                case "foo bar":
                    v7 = v5
                    v207 := v3[(v1 ^ (fi204 ^ fi204))]
                    _ = v207
                }
                v208 := []int{fi204, 923, (fi204 * fi204)}
                _ = v208
            }
            fmt.Println(v8, v43, v16)
            _ = fi205
        }
        fmt.Println(v5, v16)
        _ = fi204
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
