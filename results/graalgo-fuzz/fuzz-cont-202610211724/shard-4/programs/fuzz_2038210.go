package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint8; F1 int; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint; F1 bool; F2 int32 }
type GenS15 struct { F0 bool; F1 uint16 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 int; F1 int32 }
func genFn0() string {
    v4 := []int{}
    v5 := GenS1{F0: true}
    var v6 int64 = 0
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    _ = v4
    _ = v5
    _ = v6
    return ("hello" + "foo bar")
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{1, 42, 65535}
    var v3 int8 = (v1 ^ (v1 + (v1 ^ (-(v1)))))
    var v7 int64 = 42
    var v8 int64 = v7
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            var v10 uint16 = 32767
            if ((((false && (fi9 >= int(v3))) || (2147483647 != 0)) && (!true)) && true) {
                switch (fi9 % ((fi9) | 1)) {
                case 1:
                    for fi11 := 0; fi11 < 14; fi11++ {
                        var v12 uint32 = 61
                        v7 = 42
                        var v13 int32 = 0
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                default:
                    v14 := GenS1{F0: false}
                    _ = v14
                }
                fmt.Println(v7, v7)
                if ((v10 & (v10 << ((42) & 63))) != v10) {
                    if (-2147483648 > -1) {
                        fmt.Println(fi9, v10)
                        v1 = (v1 + v3)
                        for fi15 := 0; fi15 < 4; fi15++ {
                            var v16 uint16 = v10
                            v16 += (v10 + v16)
                            _ = fi15
                            _ = v16
                        }
                    }
                    v7 -= (v7 | 9223372036854775807)
                    var v17 uint64 = 42
                    _ = v17
                } else {
                    v10 = (v10 * 720)
                    var v18 uint = 65535
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    v2 = append(v2, 42)
                    if ((511 >= 42) && (int64(v10) != (v8 & v7))) {
                        var v20 float64 = 0.1
                        _ = v20
                    }
                    v21 := GenS1{F0: true}
                    _ = fi19
                    _ = v21
                }
                v22 := GenS1{F0: false}
                _ = v22
            }
            if len(v2) > 0 {
                v2[0] = 127
            }
            _ = v10
        } else {
            var v23 int64 = v7
            for fi24 := 0; fi24 < 8; fi24++ {
                switch (fi9 % ((fi9) | 1)) {
                case 948:
                    switch fi9 {
                    case 128:
                        var v25 uint64 = 4294967295
                        _ = v25
                    case 127:
                        fmt.Println(v1)
                    case 295:
                        v26 := GenS1{F0: (4294967295 != 1)}
                        _ = v26
                    default:
                        if false {
                            if false {
                                fmt.Println(fi24)
                                v23 = 0
                            }
                        }
                    }
                case 32768:
                    if false {
                        if (442 > 928) {
                            if ((v1 % (((v1 << ((257) & 63))) | 1)) <= v1) {
                                switch 128 {
                                case -9223372036854775808:
                                    v27 := GenS2{F0: 126, F1: (^(v1)), F2: 510}
                                    _ = v27
                                case 1:
                                    fmt.Println(v1, v23)
                                }
                                _ = len(v2)
                                fmt.Println(fi9, fi24, v7)
                            }
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    switch (("\n\t" + "\n\t") + (("foo bar" + "a") + "hello")) {
                    case "a":
                        v8 = v23
                    default:
                        v2 = append(v2, 65535)
                    }
                case 128:
                    var v28 uint8 = 127
                    _ = v28
                default:
                    if (((0 >= 2147483647) && true) || (!(v1 > (v1 >> ((458) & 63))))) {
                        v3 += -128
                        var v29 int8 = -128
                        _ = v29
                    }
                }
                v1 = (v1 + 1)
                v2 = append(v2, 722)
                _ = fi24
            }
            v7 -= 0
            _ = v23
        }
        _ = fi9
    }
    if ((-1 <= 127) || ((3.14159 != 1e10) && false)) {
        v3 -= v3
        switch -9223372036854775808 {
        case -2147483648:
            var v30 uint16 = 32767
            _ = v30
        case -1:
            fmt.Println(v7)
        case 1:
            if (!(v3 > v3)) {
                v31 := GenS1{F0: (!true)}
                var v32 uint = 155
                if (127 >= 1) {
                    if false {
                        var v33 int64 = (v7 | (v8 << ((v32) & 63)))
                        v1 = 1
                        if (-37 <= 320) {
                            switch 128 {
                            case 0:
                                var v34 int32 = -504
                                _ = v34
                            default:
                                fmt.Println(v7, v1)
                            }
                            var v35 int32 = 1
                            if ((v35 + (v35 - v35)) == (v35 - v35)) {
                                if false {
                                    switch "hello" {
                                    case "a":
                                        v8 += (v33 / ((v7) | 1))
                                    }
                                } else {
                                    switch int(v32) {
                                    case 0:
                                        if (!(true && true)) {
                                            for fi36 := 0; fi36 < 19; fi36++ {
                                                v35 = 42
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if true {
                                                    for fi37 := 0; fi37 < 7; fi37++ {
                                                        _ = len(v2)
                                                        switch ("foo bar" + "hello") {
                                                        case "a":
                                                            v3 -= -1
                                                        case "foo bar":
                                                            v2 = append(v2, 2)
                                                        case "\n\t":
                                                            v38 := GenS3{F0: (!true), F1: true}
                                                            _ = v38
                                                        default:
                                                            var v39 uint = (v32 / ((18446744073709551615) | 1))
                                                            _ = v39
                                                        }
                                                        var v40 uint16 = 65535
                                                        _ = fi37
                                                        _ = v40
                                                    }
                                                    if (2 <= 0) {
                                                        for fi41 := 0; fi41 < 18; fi41++ {
                                                            v8 -= v7
                                                            switch fi36 {
                                                            case 32767:
                                                                var v42 uint8 = 255
                                                                _ = v42
                                                            case 9223372036854775807:
                                                                v43 := GenS4{F0: 1, F1: (fi36 - (fi36 * fi36))}
                                                                _ = v43
                                                            default:
                                                                v44 := GenS4{F0: fi41, F1: fi36}
                                                                _ = v44
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = fi41
                                                        }
                                                        var v45 int64 = v33
                                                        var v46 uint32 = 2147483648
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    v3 += (v3 % ((v1) | 1))
                                                }
                                                _ = fi36
                                            }
                                        } else {
                                            var v47 uint64 = 18446744073709551615
                                            var v48 uint = (v32 & 255)
                                            _ = v47
                                            _ = v48
                                        }
                                    case -2147483648:
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                    default:
                                        v49 := GenS3{F0: (v7 != (v8 & (v8 << ((v32) & 63)))), F1: (2147483648 < 2147483647)}
                                        _ = v49
                                    }
                                    if (0 >= 2147483647) {
                                        var v50 int8 = (v3 >> ((1) & 63))
                                        _ = v50
                                    }
                                    for fi51 := 0; fi51 < 14; fi51++ {
                                        switch -554 {
                                        case 0:
                                            v2 = append(v2, 42)
                                        case 32768:
                                            var v52 float32 = (((((3.14159 - 0.1) + (-2.25 / 0.1)) * ((((0.0 * (float32((((float64(0.0) * (1.5 - 0.0)) + 0.0) * 3.14159)) - ((-2.25 + (1e10 * 1.5)) + 3.14159))) * 3.14159) * float32(1e10)) - 3.14159)) / (3.14159 + 0.0)) + 0.1)
                                            _ = v52
                                        case -988:
                                            switch "\n\t" {
                                            case "hello":
                                                for fi53 := 0; fi53 < 19; fi53++ {
                                                    if false {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v54 uint16 = 42
                                                            _ = v54
                                                        }
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    var v55 uint = v32
                                                    switch fi51 {
                                                    case 32767:
                                                        v35 = v35
                                                    case 1:
                                                        var v56 int = (fi51 >> (((^(v32))) & 63))
                                                        _ = v56
                                                    }
                                                    var v57 bool = (1 > 2)
                                                    _ = fi53
                                                    _ = v55
                                                    _ = v57
                                                }
                                            default:
                                                for fi58 := 0; fi58 < 4; fi58++ {
                                                    if false {
                                                        fmt.Println(fi51, v35, fi58)
                                                        var v59 bool = true
                                                        _ = v59
                                                    }
                                                    _ = fi58
                                                }
                                            }
                                        }
                                        _ = fi51
                                    }
                                }
                                v60 := GenS1{F0: false}
                                _ = v60
                            } else {
                                var v61 int = 128
                                _ = v61
                            }
                            _ = v35
                        } else {
                            var v62 string = "a"
                            fmt.Println(v1, v1)
                            _ = v62
                        }
                        _ = v33
                    } else {
                        v63 := GenS5{F0: 54, F1: 42, F2: 617}
                        for fi64 := 0; fi64 < 9; fi64++ {
                            fmt.Println(fi64, v1)
                            v7 += 303
                            _ = fi64
                        }
                        _ = v63
                    }
                }
                _ = v31
                _ = v32
            }
        }
    } else {
        if true {
            for fi65 := 0; fi65 < 18; fi65++ {
                v1 += v3
                var v66 uint = 0
                _ = fi65
                _ = v66
            }
            _ = len(v2)
            if (747 < 42) {
                _ = len(v2)
                v67 := GenS6{F0: 42}
                _ = v67
            }
        } else {
            fmt.Println(v7, v8, v7)
            v68 := GenS7{F0: 9223372036854775807}
            switch "hello" {
            case "a":
                var v69 uint32 = 1
                _ = v69
            default:
                v8 -= (v7 ^ v8)
            }
            _ = v68
        }
        if true {
            var v70 int32 = 2147483647
            if true {
                if (!false) {
                    var v71 int8 = (v3 % ((v1) | 1))
                    v72 := GenS8{F0: 127, F1: 0}
                    if (!true) {
                        v73 := GenS7{F0: 120}
                        _ = v73
                    } else {
                        var v74 int64 = v7
                        switch 128 {
                        case 2147483647:
                            switch "" {
                            case "foo bar":
                                for fi75 := 0; fi75 < 6; fi75++ {
                                    switch "hello" {
                                    case "foo bar":
                                        v76 := GenS7{F0: (fi75 >> ((0) & 63))}
                                        _ = v76
                                    case "\n\t":
                                        if (1e10 == 1e10) {
                                            if true {
                                                for fi77 := 0; fi77 < 9; fi77++ {
                                                    if (v70 <= 0) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if false {
                                                            fmt.Println(v1)
                                                        }
                                                    } else {
                                                        v71 = -1
                                                    }
                                                    _ = fi77
                                                }
                                            } else {
                                                if (501 > 0) {
                                                    if true {
                                                        fmt.Println(v70, v3, v3)
                                                        var v78 uint = 4294967295
                                                        var v79 uint = (v78 + v78)
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                } else {
                                                    var v80 int16 = 81
                                                    var v81 bool = false
                                                    _ = v80
                                                    _ = v81
                                                }
                                                v71 = (v3 * v3)
                                            }
                                            switch 32767 {
                                            case 0:
                                                for fi82 := 0; fi82 < 9; fi82++ {
                                                    if true {
                                                        if (215 == 153) {
                                                            v2 = append(v2, 2)
                                                        } else {
                                                            v83 := GenS9{F0: (^(v70)), F1: -1}
                                                            v70 -= (^(v70))
                                                            _ = len(v2)
                                                            _ = v83
                                                        }
                                                    } else {
                                                        if false {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v71 = -1
                                                        }
                                                        var v84 int = fi82
                                                        fmt.Println(fi82)
                                                        _ = v84
                                                    }
                                                    _ = fi82
                                                }
                                            case 32768:
                                                if true {
                                                    switch (fi75 - fi75) {
                                                    case 42:
                                                        v3 = (v1 * 0)
                                                    }
                                                    var v85 int64 = (v7 % (((v8 - (v74 * 1))) | 1))
                                                    _ = v85
                                                }
                                            default:
                                                v3 += v1
                                            }
                                            v86 := GenS9{F0: v70, F1: (fi75 & int(v70))}
                                            _ = v86
                                        } else {
                                            var v87 uint16 = 42
                                            _ = v87
                                        }
                                    case "":
                                        var v88 int32 = (v70 | int32(v7))
                                        _ = v88
                                    default:
                                        fmt.Println(v71)
                                    }
                                    _ = fi75
                                }
                            case "\n\t":
                                var v89 int8 = (v3 * 127)
                                _ = v89
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        default:
                            var v90 uint = 255
                            _ = v90
                        }
                        _ = v74
                    }
                    _ = v71
                    _ = v72
                } else {
                    v2 = append(v2, 4294967295)
                    for fi91 := 0; fi91 < 14; fi91++ {
                        v7 = v8
                        _ = fi91
                    }
                }
                v92 := GenS10{F0: false}
                _ = v92
            } else {
                var v93 uint8 = 2
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v93
            }
            _ = v70
        }
    }
    if true {
        v94 := GenS11{F0: 67, F1: false}
        var v95 uint16 = 255
        for fi96 := 0; fi96 < 2; fi96++ {
            if true {
                v97 := GenS12{F0: 1, F1: 478}
                var v98 int8 = 42
                if len(v2) > 0 {
                    v2[0] = 497
                }
                _ = v97
                _ = v98
            }
            for fi99 := 0; fi99 < 7; fi99++ {
                v7 += (v8 ^ 9223372036854775807)
                if false {
                    if true {
                        var v100 int = (-(fi99))
                        _ = v100
                    } else {
                        var v101 int = (fi99 ^ (fi96 << ((272) & 63)))
                        _ = v101
                    }
                    fmt.Println(fi96, v8, v95)
                } else {
                    if (4294967295 > 2147483647) {
                        if true {
                            switch "hello" {
                            case "\n\t":
                                if ((fi99 * (fi96 % ((fi96) | 1))) >= fi99) {
                                    _ = len(v2)
                                } else {
                                    v102 := GenS13{F0: -32768}
                                    v103 := GenS14{F0: 42, F1: false, F2: 1}
                                    v95 = 2
                                    _ = v102
                                    _ = v103
                                }
                            case "a":
                                v104 := GenS10{F0: false}
                                _ = v104
                            case "":
                                if ((1 >= 2) || ((fi96 | 42) != 32768)) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v1 = v3
                                    fmt.Println(fi99)
                                }
                            default:
                                var v105 float64 = (1.5 + 0.1)
                                _ = v105
                            }
                        }
                    } else {
                        switch 901 {
                        case -2147483648:
                            v2 = append(v2, 65535)
                        case 902:
                            var v106 bool = (180 > (v95 / (((v95 * (v95 ^ (^(v95))))) | 1)))
                            _ = v106
                        case 42:
                            switch (-(fi96)) {
                            case 42:
                                v1 -= v1
                            case 420:
                                fmt.Println(v95)
                            }
                        }
                    }
                    var v107 int8 = int8(v95)
                    var v108 uint64 = 1
                    _ = v107
                    _ = v108
                }
                _ = fi99
            }
            _ = fi96
        }
        _ = v94
        _ = v95
    } else {
        switch "" {
        case "a":
            v109 := GenS6{F0: v1}
            _ = v109
        case "":
            for fi110 := 0; fi110 < 2; fi110++ {
                if (((!false) && true) || true) {
                    _ = len(v2)
                    switch fi110 {
                    case 128:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -141:
                        if true {
                            v7 -= -9223372036854775808
                            v8 += -1
                            v111 := GenS15{F0: true, F1: 1}
                            _ = v111
                        } else {
                            var v112 int16 = 1
                            var v113 int8 = -1
                            _ = v112
                            _ = v113
                        }
                    case -1:
                        v8 = 1
                    }
                    var v114 int64 = -394
                    _ = v114
                } else {
                    v8 -= (-(v7))
                }
                _ = fi110
            }
        case "hello":
            var v115 uint = 42
            _ = v115
        }
        for fi116 := 0; fi116 < 20; fi116++ {
            for fi117 := 0; fi117 < 8; fi117++ {
                _ = len(v2)
                switch ("foo bar" + "\n\t") {
                case "a":
                    if (42 != 177) {
                        v8 = (v8 >> ((0) & 63))
                    }
                case "\n\t":
                    v118 := GenS16{F0: v3}
                    _ = v118
                case "foo bar":
                    var v119 float64 = (3.14159 + -2.25)
                    _ = v119
                default:
                    if (((("hello" + (("\n\t" + "a") + "\n\t")) + (("hello" + "") + ("\n\t" + (("foo bar" + ("\n\t" + "")) + "a")))) + "\n\t") == "foo bar") {
                        var v120 float32 = 1.5
                        v121 := GenS1{F0: (2 > 0)}
                        fmt.Println(v3, fi117, fi117)
                        _ = v120
                        _ = v121
                    }
                }
                _ = fi117
            }
            v3 = (-(v1))
            _ = fi116
        }
        v7 = 42
    }
    fmt.Println(v8)
    for fi122 := 0; fi122 < 1; fi122++ {
        var v123 int32 = 42
        _ = fi122
        _ = v123
    }
    for fi124 := 0; fi124 < 9; fi124++ {
        for fi125 := 0; fi125 < 4; fi125++ {
            fmt.Println(v8, fi124, v8)
            _ = fi125
        }
        _ = fi124
    }
    if (v3 <= 39) {
        for fi126 := 0; fi126 < 12; fi126++ {
            v127 := GenS17{F0: 42}
            _ = fi126
            _ = v127
        }
        if false {
            switch -511 {
            case 1:
                var v128 uint16 = 718
                _ = v128
            default:
                if ((^(v3)) == v3) {
                    for fi129 := 0; fi129 < 20; fi129++ {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            var v130 float32 = 1.5
                            _ = len(v2)
                            var v131 uint64 = 42
                            _ = v130
                            _ = v131
                        }
                        _ = fi129
                    }
                    _ = len(v2)
                    switch 703 {
                    case 127:
                        if (0.1 == 0.0) {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                        }
                    case -769:
                        var v132 int = 419
                        _ = v132
                    default:
                        v133 := GenS18{F0: "a"}
                        _ = v133
                    }
                }
            }
            fmt.Println(v1, v3, v7)
        }
    }
    if len(v2) > 0 {
        v2[0] = 585
    }
    v134 := GenS12{F0: 831, F1: 729}
    _ = len(v2)
    if (42 >= 42) {
        var v135 uint64 = 297
        _ = v135
    }
    if true {
        if false {
            switch "hello" {
            case "\n\t":
                switch 2147483647 {
                case 2147483647:
                    if false {
                        v3 = -128
                    }
                case 127:
                    v136 := GenS2{F0: -1, F1: v3, F2: 213}
                    _ = v136
                default:
                    v3 -= v1
                }
            case "foo bar":
                fmt.Println(v3)
            case "hello":
                var v137 uint32 = 2147483647
                _ = v137
            }
        }
        if len(v2) > 0 {
            v2[0] = 690
        }
        switch (("\n\t" + "a") + "hello") {
        case "":
            var v138 bool = (!false)
            _ = v138
        case "foo bar":
            if (316 > 234) {
                v7 = (v7 / (((v7 - v8)) | 1))
            }
        }
    } else {
        switch "" {
        case "a":
            fmt.Println(v7, v8, v1)
        default:
            fmt.Println(v1, v8, v7)
        }
        if false {
            v1 += v1
        } else {
            for fi139 := 0; fi139 < 14; fi139++ {
                for fi140 := 0; fi140 < 6; fi140++ {
                    var v141 bool = (1 <= 0)
                    _ = fi140
                    _ = v141
                }
                switch -981 {
                case 32767:
                    v142 := GenS6{F0: 5}
                    _ = v142
                case 909:
                    v8 = 661
                }
                if (!(false && true)) {
                    fmt.Println(v7, v8)
                }
                _ = fi139
            }
            var v143 uint = 65535
            _ = v143
        }
    }
    v1 -= int8(v8)
    if (1e10 <= 1.5) {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    } else {
        for fi144 := 0; fi144 < 1; fi144++ {
            _ = len(v2)
            v145 := GenS19{F0: fi144, F1: 42}
            _ = fi144
            _ = v145
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v146 uint16 = 0
    switch "" {
    case "a":
        var v147 string = "hello"
        _ = v147
    case "foo bar":
        fmt.Println(v146, v7, v1)
    case "\n\t":
        var v148 uint64 = 720
        _ = v148
    default:
        v146 -= v146
    }
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v134
    _ = v146
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
