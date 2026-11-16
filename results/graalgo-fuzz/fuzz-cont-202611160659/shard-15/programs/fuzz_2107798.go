package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int; F2 string }
type GenS5 struct { F0 int; F1 int64; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string; F1 bool; F2 int16 }
type GenS12 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS13 struct { F0 bool; F1 int8; F2 string }
type GenS14 struct { F0 uint32; F1 uint32 }
type GenS15 struct { F0 int32; F1 int16 }
type GenS16 struct { F0 int32; F1 string; F2 int8 }
func genFn0() uint64 {
    v7 := []int{42}
    _ = len(v7)
    _ = v7
    return 2
}
func genFn1(p0 string, p1 uint16) uint {
    if (15 < 100) {
        if (-82 >= 110) {
            var v8 uint32 = 2147483647
            _ = v8
        } else {
            if true {
                switch 105 {
                case 9223372036854775807:
                    v9 := GenS1{F0: 65535, F1: 170}
                    _ = v9
                case -930:
                    if (true && (true && true)) {
                        var v10 string = p0
                        if ((!(!true)) || false) {
                            fmt.Println(v10, p0, v10)
                            p0 = v10
                            p0 = "hello"
                        }
                        _ = v10
                    }
                case 128:
                    var v11 int32 = -140
                    _ = v11
                }
                p0 = "\n\t"
            } else {
                if (true && ((-2147483648 < 1) || false)) {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        var v13 float64 = float64(-2.25)
                        for fi14 := 0; fi14 < 14; fi14++ {
                            var v15 uint = 18446744073709551615
                            _ = fi14
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                }
            }
        }
        var v16 uint = 818
        switch ((("foo bar" + p0) + p0) + ("hello" + p0)) {
        case "hello":
            if (!false) {
                var v17 bool = false
                v18 := GenS2{F0: 0}
                var v19 int8 = 1
                _ = v17
                _ = v18
                _ = v19
            }
        default:
            p0 = p0
        }
        _ = v16
    } else {
        p0 = (p0 + p0)
        fmt.Println(p1, p0)
    }
    if (-669 != -9223372036854775808) {
        p0 = ("foo bar" + p0)
        v20 := []int{}
        _ = v20
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 0
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    v5 := make(map[uint32]uint32)
    var v6 int64 = 86
    v2 = 1
    var v21 uint32 = 1
    var v22 int16 = -89
    v23 := GenS2{F0: 18446744073709551615}
    delete(v3, (v21 >> ((149) & 63)))
    if false {
        _ = len(v5)
        var v24 int8 = (v2 - 127)
        v25 := v3[(v21 ^ v21)]
        _ = v24
        _ = v25
    }
    if (!true) {
        var v26 uint16 = 32767
        if (false || false) {
            v22 = v22
            switch "" {
            case "foo bar":
                fmt.Println(v1)
            case "\n\t":
                switch 9223372036854775807 {
                case 42:
                    v22 = 32767
                default:
                    _ = len(v4)
                }
            case "a":
                for fi27 := 0; fi27 < 3; fi27++ {
                    if ((^(v21)) == v21) {
                        switch (fi27 / ((fi27) | 1)) {
                        case 9223372036854775807:
                            for fi28 := 0; fi28 < 3; fi28++ {
                                var v29 int8 = 115
                                _ = fi28
                                _ = v29
                            }
                        case -1:
                            if false {
                                for fi30 := 0; fi30 < 4; fi30++ {
                                    switch v1 {
                                    case "":
                                        v21 -= v21
                                    }
                                    var v31 uint16 = (^(v26))
                                    _ = fi30
                                    _ = v31
                                }
                                if (!(!true)) {
                                    if false {
                                        for fi32 := 0; fi32 < 1; fi32++ {
                                            switch v1 {
                                            case "foo bar":
                                                v33 := GenS1{F0: (v21 << ((uint(v26)) & 63)), F1: 42}
                                                _ = v33
                                            case "hello":
                                                v34 := []int{}
                                                _ = v34
                                            default:
                                                if false {
                                                    var v35 float32 = float32(1.5)
                                                    _ = v35
                                                } else {
                                                    switch 32768 {
                                                    case -751:
                                                        if true {
                                                            v26 -= (v26 & v26)
                                                        } else {
                                                            var v36 uint = uint(v21)
                                                            v21 += v21
                                                            switch (("hello" + v1) + "foo bar") {
                                                            case "hello":
                                                                v37 := []int{(^(fi32)), (fi27 % ((int(v2)) | 1)), -1}
                                                                v21 += 0
                                                                _ = v37
                                                            case "":
                                                                if (false && false) {
                                                                    var v38 int = -551
                                                                    switch (v38 >> (((v36 ^ genFn1(("" + (v1 + v1)), (v26 & (v26 << (((v36 ^ v36)) & 63)))))) & 63)) {
                                                                    case -63:
                                                                        fmt.Println(v36)
                                                                    case 42:
                                                                        v39 := v4[v1]
                                                                        _ = v39
                                                                    default:
                                                                        fmt.Println(v22)
                                                                    }
                                                                    var v40 uint32 = (v21 % ((v21) | 1))
                                                                    _ = v38
                                                                    _ = v40
                                                                } else {
                                                                    var v41 uint = uint(v21)
                                                                    var v42 float32 = -2.25
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                                delete(v5, (v21 & (v21 * (v21 & uint32(v26)))))
                                                            case "a":
                                                                if false {
                                                                    v2 += v2
                                                                    var v43 int64 = (v6 ^ (v6 | v6))
                                                                    v44 := []int{(fi27 | 32767)}
                                                                    _ = v43
                                                                    _ = v44
                                                                } else {
                                                                    v21 += v21
                                                                    v45 := GenS1{F0: uint32(v36), F1: 2}
                                                                    _ = v45
                                                                }
                                                            default:
                                                                v22 = (-(v22))
                                                            }
                                                            if (fi27 < 128) {
                                                                var v46 float64 = 3.14159
                                                                v6 += v6
                                                                v47 := GenS2{F0: uint64(v21)}
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                            _ = v36
                                                        }
                                                    case 125:
                                                        v48 := []int{149, (fi27 + fi27), fi32}
                                                        _ = v48
                                                    }
                                                    v49 := GenS1{F0: v21, F1: 42}
                                                    _ = v49
                                                }
                                            }
                                            v6 += 1
                                            delete(v3, v21)
                                            switch "\n\t" {
                                            case "":
                                                v1 = "\n\t"
                                            default:
                                                v6 = 0
                                            }
                                            v50 := GenS3{F0: genFn0()}
                                            _ = len(v5)
                                            _ = fi32
                                            _ = v50
                                        }
                                        var v51 uint = 65535
                                        _ = v51
                                    } else {
                                        switch ("foo bar" + "a") {
                                        case "\n\t":
                                            v6 -= (v6 + 42)
                                        default:
                                            fmt.Println(v22, v21, v2)
                                        }
                                        v52 := []int{(fi27 % ((fi27) | 1)), (-(fi27))}
                                        v53 := GenS3{F0: 564}
                                        _ = v52
                                        _ = v53
                                    }
                                    var v54 uint16 = (v26 + (v26 + 909))
                                    _ = v54
                                }
                                switch -102 {
                                case 2147483647:
                                    switch v1 {
                                    case "":
                                        v55 := v5[(v21 % (((^(v21))) | 1))]
                                        _ = v55
                                    case "hello":
                                        if true {
                                            switch "\n\t" {
                                            case "":
                                                var v56 int8 = 1
                                                _ = v56
                                            case "hello":
                                                var v57 int64 = (v6 >> ((2) & 63))
                                                _ = v57
                                            case "foo bar":
                                                for fi58 := 0; fi58 < 4; fi58++ {
                                                    fmt.Println(v6, v2)
                                                    v59 := GenS4{F0: (v22 / ((v22) | 1)), F1: 128, F2: (v1 + (("" + v1) + ""))}
                                                    if (!false) {
                                                        delete(v4, (("a" + (v1 + "foo bar")) + v1))
                                                        var v60 string = (v1 + "\n\t")
                                                        _ = v60
                                                    }
                                                    _ = fi58
                                                    _ = v59
                                                }
                                            }
                                            var v61 float32 = (0.1 * 0.1)
                                            switch (v1 + "") {
                                            case "\n\t":
                                                v62 := GenS5{F0: fi27, F1: (v6 << ((42) & 63)), F2: (v6 + (v6 / (((v6 << ((genFn1(v1, (v26 | 42))) & 63))) | 1)))}
                                                _ = v62
                                            case "a":
                                                var v63 int32 = -1
                                                _ = v63
                                            case "foo bar":
                                                if false {
                                                    if false {
                                                        v2 = (-(v2))
                                                        var v64 float32 = v61
                                                        var v65 uint32 = v21
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        for fi66 := 0; fi66 < 13; fi66++ {
                                                            var v67 string = (v1 + "foo bar")
                                                            v1 = (v67 + "")
                                                            if (true && (v6 <= (^(v6)))) {
                                                                var v68 int16 = (v22 & (-(v22)))
                                                                _ = v68
                                                            }
                                                            _ = fi66
                                                            _ = v67
                                                        }
                                                        switch v1 {
                                                        case "hello":
                                                            switch v1 {
                                                            case "a":
                                                                v69 := GenS6{F0: v21}
                                                                _ = v69
                                                            case "hello":
                                                                if (!(!true)) {
                                                                    if (294 < v21) {
                                                                        for fi70 := 0; fi70 < 4; fi70++ {
                                                                            v26 -= 582
                                                                            if true {
                                                                                var v71 uint64 = 0
                                                                                switch 126 {
                                                                                case 311:
                                                                                    fmt.Println(v71, v21, fi27)
                                                                                }
                                                                                if (!false) {
                                                                                    delete(v5, v21)
                                                                                }
                                                                                _ = v71
                                                                            } else {
                                                                                v72 := GenS7{F0: (v22 / ((int16(v2)) | 1)), F1: v2}
                                                                                v73 := GenS1{F0: 216, F1: 223}
                                                                                v22 = -3
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                            fmt.Println(v26)
                                                                            _ = fi70
                                                                        }
                                                                        if false {
                                                                            v22 = (v22 >> ((689) & 63))
                                                                        }
                                                                        v74 := GenS8{F0: 0}
                                                                        _ = v74
                                                                    } else {
                                                                        _ = len(v3)
                                                                        v75 := []int{}
                                                                        var v76 uint64 = uint64(v26)
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                } else {
                                                                    for fi77 := 0; fi77 < 8; fi77++ {
                                                                        switch "hello" {
                                                                        case "":
                                                                            v78 := GenS9{F0: (^(v26))}
                                                                            _ = v78
                                                                        case "a":
                                                                            switch fi27 {
                                                                            case -1:
                                                                                if true {
                                                                                    fmt.Println(v1)
                                                                                    var v79 string = v1
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                            fmt.Println(fi77)
                                                                        default:
                                                                            v80 := GenS1{F0: (^(v21)), F1: uint8(v26)}
                                                                            _ = v80
                                                                        }
                                                                        var v81 uint = 255
                                                                        _ = fi77
                                                                        _ = v81
                                                                    }
                                                                }
                                                            default:
                                                                v82 := []int{(fi27 & (fi27 + (fi27 & (fi27 % ((fi27) | 1))))), fi27, (fi27 * -901)}
                                                                _ = v82
                                                            }
                                                        case "":
                                                            v83 := GenS9{F0: 2}
                                                            _ = v83
                                                        }
                                                        switch 4294967295 {
                                                        case 127:
                                                            if true {
                                                                switch (fi27 + (-(fi27))) {
                                                                case -1:
                                                                    for fi84 := 0; fi84 < 20; fi84++ {
                                                                        v26 = (v26 & (v26 ^ v26))
                                                                        if false {
                                                                            v85 := []int{4294967295, (fi27 * fi84), (^(fi84)), (fi84 + fi84)}
                                                                            v2 = -128
                                                                            _ = v85
                                                                        } else {
                                                                            v3[v21] = (fi27 << ((18446744073709551615) & 63))
                                                                            v86 := v3[285]
                                                                            _ = v86
                                                                        }
                                                                        v61 = 0.0
                                                                        _ = fi84
                                                                    }
                                                                case -2147483648:
                                                                    var v87 int = 128
                                                                    _ = v87
                                                                case 32767:
                                                                    v3[2147483648] = (fi27 + fi27)
                                                                }
                                                                fmt.Println(v22, v26)
                                                            }
                                                        case 62:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                fmt.Println(fi27, v2)
                                                            case "\n\t":
                                                                v1 = v1
                                                            case "a":
                                                                switch v1 {
                                                                case "hello":
                                                                    for fi88 := 0; fi88 < 13; fi88++ {
                                                                        v6 = v6
                                                                        v26 = (v26 << ((42) & 63))
                                                                        _ = fi88
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                v61 += 1.5
                                            }
                                            _ = v61
                                        } else {
                                            v3[v21] = -2147483648
                                            fmt.Println(v22, v21, v1)
                                        }
                                    }
                                case 127:
                                    switch "hello" {
                                    case "foo bar":
                                        v2 = (v2 ^ (v2 + v2))
                                    }
                                case 32768:
                                    var v89 uint16 = 1
                                    _ = v89
                                }
                            } else {
                                var v90 uint32 = (v21 - 65535)
                                v91 := GenS4{F0: v22, F1: (fi27 ^ (fi27 << ((42) & 63))), F2: v1}
                                _ = v90
                                _ = v91
                            }
                        case 1:
                            var v92 float64 = ((0.1 + (3.14159 * (float64(((float32(0.1) / 1.5) / (0.0 - (0.1 * 1.5)))) - (float64(3.14159) + 1.5)))) - 1e10)
                            _ = v92
                        }
                        v93 := GenS10{F0: (-(fi27))}
                        if true {
                            for fi94 := 0; fi94 < 18; fi94++ {
                                v5[v21] = v21
                                var v95 uint = 935
                                _ = fi94
                                _ = v95
                            }
                            var v96 uint32 = 1
                            switch ((v1 + v1) + (v1 + v1)) {
                            case "a":
                                var v97 float32 = 0.1
                                _ = v97
                            case "foo bar":
                                if (((^(v26)) <= 815) && (1 >= -2147483648)) {
                                    v2 = v2
                                    for fi98 := 0; fi98 < 8; fi98++ {
                                        if (false || false) {
                                            var v99 bool = (uint32(v26) == 839)
                                            switch v1 {
                                            case "a":
                                                v6 -= v6
                                            case "":
                                                fmt.Println(v99, v2)
                                            default:
                                                v1 = v1
                                                v21 = 2147483647
                                            }
                                            v100 := GenS11{F0: v1, F1: v99, F2: v22}
                                            v101 := GenS3{F0: 571}
                                            var v102 int64 = (v6 / (((v6 ^ -1)) | 1))
                                            var v103 uint64 = 2
                                            _ = v100
                                            _ = v101
                                            _ = v102
                                            _ = v103
                                            _ = v99
                                        } else {
                                            v26 = 255
                                            v21 -= (^(v21))
                                        }
                                        fmt.Println(v21, v22, v96)
                                        if true {
                                            v104 := GenS7{F0: 32767, F1: 111}
                                            switch v1 {
                                            case "\n\t":
                                                v2 -= int8(fi98)
                                            default:
                                                v21 -= v21
                                            }
                                            v26 -= (v26 | (v26 | v26))
                                            _ = v104
                                        }
                                        _ = fi98
                                    }
                                    fmt.Println(v26, v1, fi27)
                                }
                            default:
                                v105 := GenS12{F0: v26, F1: (-(v2)), F2: 286}
                                _ = v105
                            }
                            _ = v96
                        } else {
                            for fi106 := 0; fi106 < 14; fi106++ {
                                v1 = ""
                                v107 := []int{int(v22), 32767, int(v6), 42}
                                var v108 int8 = (-(v2))
                                switch ("foo bar" + v1) {
                                case "hello":
                                    switch 1 {
                                    case 1:
                                        switch (fi27 - (fi27 / ((74) | 1))) {
                                        case 896:
                                            v107 = append(v107, fi106)
                                            if (810 == 42) {
                                                if true {
                                                    if len(v107) > 0 {
                                                        v107[0] = fi106
                                                    }
                                                    fmt.Println(v21)
                                                    v109 := GenS13{F0: false, F1: 42, F2: v1}
                                                    _ = v109
                                                } else {
                                                    var v110 uint64 = 2
                                                    var v111 float32 = (((1e10 - -2.25) - -2.25) * 1.5)
                                                    _ = v110
                                                    _ = v111
                                                }
                                                if true {
                                                    _ = len(v5)
                                                    v1 = v1
                                                    v6 += v6
                                                } else {
                                                    switch (-(fi106)) {
                                                    case 32767:
                                                        v112 := v3[(v21 ^ 2147483647)]
                                                        _ = v112
                                                    }
                                                    v113 := GenS8{F0: 1}
                                                    v114 := GenS14{F0: 368, F1: (v21 * v21)}
                                                    _ = v113
                                                    _ = v114
                                                }
                                                var v115 float32 = (1.5 / float32(0.1))
                                                v116 := GenS13{F0: (1 < (v26 ^ (v26 - 255))), F1: (v108 << ((genFn1(("hello" + (v1 + ("" + ("\n\t" + v1)))), 0)) & 63)), F2: "a"}
                                                var v117 float32 = 1.5
                                                _ = v115
                                                _ = v116
                                                _ = v117
                                            } else {
                                                switch v1 {
                                                case "\n\t":
                                                    v2 += 100
                                                case "foo bar":
                                                    v26 = v26
                                                }
                                                var v118 int = (fi106 & fi106)
                                                _ = v118
                                            }
                                            v6 += v6
                                            if len(v107) > 0 {
                                                v107[0] = (fi27 & (fi106 >> ((219) & 63)))
                                            }
                                            var v119 float64 = (1e10 - ((0.0 - 0.0) - 0.1))
                                            _ = v119
                                        case 32767:
                                            switch v1 {
                                            case "hello":
                                                var v120 int64 = (-(v6))
                                                _ = v120
                                            case "a":
                                                var v121 uint64 = 1
                                                fmt.Println(v108, v22, v26)
                                                v122 := GenS1{F0: v21, F1: 128}
                                                _ = v121
                                                _ = v122
                                            case "foo bar":
                                                v6 = (v6 << ((65535) & 63))
                                            default:
                                                switch v1 {
                                                case "foo bar":
                                                    var v123 uint8 = 139
                                                    _ = v123
                                                case "\n\t":
                                                    var v124 int16 = v22
                                                    _ = v124
                                                }
                                            }
                                        case 4294967295:
                                            var v125 float64 = 0.1
                                            _ = v125
                                        default:
                                            v2 = (^(v2))
                                        }
                                    }
                                case "foo bar":
                                    var v126 float32 = (1e10 - 1.5)
                                    _ = v126
                                case "":
                                    switch fi27 {
                                    case 4294967295:
                                        var v127 int8 = 1
                                        _ = v127
                                    case -332:
                                        var v128 bool = (genFn0() >= uint64(v26))
                                        _ = v128
                                    }
                                default:
                                    fmt.Println(v6, v21, v2)
                                }
                                _ = fi106
                                _ = v107
                                _ = v108
                            }
                            if true {
                                for fi129 := 0; fi129 < 1; fi129++ {
                                    delete(v4, "hello")
                                    var v130 int32 = 42
                                    v131 := []int{fi129, fi129}
                                    _ = fi129
                                    _ = v130
                                    _ = v131
                                }
                            } else {
                                v132 := GenS15{F0: -1, F1: 810}
                                _ = v132
                            }
                            var v133 uint32 = (v21 % ((4294967295) | 1))
                            _ = v133
                        }
                        _ = v93
                    } else {
                        for fi134 := 0; fi134 < 11; fi134++ {
                            v1 = v1
                            _ = fi134
                        }
                        v21 = v21
                    }
                    v135 := []int{fi27, fi27, fi27}
                    _ = fi27
                    _ = v135
                }
            default:
                var v136 bool = true
                _ = v136
            }
        }
        v137 := GenS16{F0: -914, F1: "foo bar", F2: v2}
        _ = v137
        _ = v26
    } else {
        var v138 int32 = 0
        var v139 int8 = (-(v2))
        _ = v138
        _ = v139
    }
    v2 = 33
    v21 = v21
    for fi140 := 0; fi140 < 15; fi140++ {
        for fi141 := 0; fi141 < 19; fi141++ {
            v142 := GenS14{F0: v21, F1: v21}
            v21 += v21
            _ = fi141
            _ = v142
        }
        v21 -= 408
        _ = fi140
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
