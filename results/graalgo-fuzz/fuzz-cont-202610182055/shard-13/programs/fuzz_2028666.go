package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint; F1 uint16; F2 string }
type GenS6 struct { F0 int; F1 uint8; F2 int32 }
type GenS7 struct { F0 int16; F1 uint32; F2 uint }
type GenS8 struct { F0 int; F1 uint; F2 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64; F1 string; F2 int16 }
type GenS12 struct { F0 uint8; F1 int; F2 uint16 }
type GenS13 struct { F0 string; F1 int16 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint32; F1 int64; F2 int }
type GenS16 struct { F0 int32; F1 uint; F2 uint16 }
type GenS17 struct { F0 uint; F1 bool; F2 uint16 }
func genFn0(p0 uint16, p1 uint) uint16 {
    p1 = (p1 * (p1 + p1))
    var v7 string = "hello"
    _ = p0
    _ = p1
    _ = v7
    return p0
}
func genFn1() int32 {
    var v8 float32 = (0.0 * 3.14159)
    switch ("foo bar" + "\n\t") {
    case "hello":
        v9 := GenS1{F0: 42}
        _ = v9
    default:
        v8 += float32(0.1)
    }
    if (!(false || (!true))) {
        var v10 int = 9223372036854775807
        v11 := []int{v10}
        var v12 int = v10
        _ = v10
        _ = v11
        _ = v12
    } else {
        v8 += ((v8 * v8) + 0.1)
    }
    v8 = float32(((((3.14159 * (1e10 + float64((v8 - float32(1e10))))) * (0.0 * (1.5 * 0.1))) + 3.14159) + 0.1))
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: 2147483647}
        switch fi13 {
        case -9223372036854775808:
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS1{F0: 267}
                _ = fi15
                _ = v16
            }
        default:
            v8 = -2.25
        }
        if true {
            switch (fi13 + (fi13 / ((fi13) | 1))) {
            case -2147483648:
                var v17 uint32 = 4294967295
                _ = v17
            case 32767:
                v18 := []int{fi13, (fi13 + -44), 127}
                _ = v18
            case -601:
                v19 := GenS1{F0: -2147483648}
                _ = v19
            default:
                var v20 int64 = 42
                _ = v20
            }
            var v21 uint64 = 42
            for fi22 := 0; fi22 < 1; fi22++ {
                if true {
                    if false {
                        v23 := GenS2{F0: 373, F1: fi22, F2: -2147483648}
                        fmt.Println(fi22)
                        var v24 uint64 = v21
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(fi13)
                    v25 := []int{1, fi13}
                    _ = v25
                } else {
                    var v26 int8 = 0
                    _ = v26
                }
                var v27 uint16 = 475
                v21 = v21
                _ = fi22
                _ = v27
            }
            _ = v21
        }
        _ = fi13
        _ = v14
    }
    _ = v8
    return 0
}

func main() {
    var v1 uint = 65535
    var v2 string = "foo bar"
    var v3 uint64 = 0
    v4 := make(map[string]int)
    var v5 uint16 = 221
    var v6 uint64 = (v3 / ((97) | 1))
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 int16 = 0
            switch fi28 {
            case 9223372036854775807:
                v1 = 4294967295
            case -657:
                if ((981 == v30) && ((float64(((-2.25 + 0.1) + -2.25)) + 0.1) <= -2.25)) {
                    switch (fi28 | fi29) {
                    case 768:
                        v30 -= v30
                    default:
                        var v31 uint32 = 0
                        _ = v31
                    }
                }
            }
            _ = fi29
            _ = v30
        }
        var v32 uint8 = 42
        v33 := []int{}
        _ = fi28
        _ = v32
        _ = v33
    }
    v34 := []int{128, int(v6), 666, -9223372036854775808}
    v5 -= v5
    _ = len(v34)
    switch ("" + v2) {
    case "":
        v2 = "foo bar"
    case "a":
        v35 := GenS3{F0: -9223372036854775808, F1: (!false)}
        _ = v35
    case "foo bar":
        switch (v2 + "\n\t") {
        case "\n\t":
            for fi36 := 0; fi36 < 6; fi36++ {
                v3 -= 18446744073709551615
                v37 := GenS3{F0: -897, F1: (1 <= 2)}
                switch "\n\t" {
                case "a":
                    if ((v3 * v3) <= (v3 * 22)) {
                        v6 = uint64(v1)
                        var v38 int8 = 127
                        v3 += 562
                        _ = v38
                    } else {
                        v34 = append(v34, fi36)
                        delete(v4, v2)
                        v39 := GenS4{F0: int(v1), F1: false}
                        _ = v39
                    }
                case "hello":
                    if false {
                        fmt.Println(v2, v6, v3)
                        switch v2 {
                        case "":
                            switch (("\n\t" + v2) + ((v2 + "a") + ("\n\t" + (v2 + ((v2 + (v2 + v2)) + "a"))))) {
                            case "a":
                                fmt.Println(v2, v2)
                            case "":
                                for fi40 := 0; fi40 < 5; fi40++ {
                                    delete(v4, v2)
                                    var v41 float64 = 1e10
                                    var v42 int32 = 749
                                    _ = fi40
                                    _ = v41
                                    _ = v42
                                }
                            case "hello":
                                if (!(((!(-1 == 42)) || (v5 == v5)) && (!true))) {
                                    fmt.Println(v2, v1)
                                    var v43 bool = (!false)
                                    fmt.Println(v3, v1)
                                    _ = v43
                                } else {
                                    if true {
                                        var v44 int32 = 42
                                        fmt.Println(v44, v6)
                                        _ = v44
                                    }
                                }
                            }
                        case "hello":
                            switch v2 {
                            case "hello":
                                for fi45 := 0; fi45 < 3; fi45++ {
                                    if false {
                                        _ = len(v4)
                                        v6 = v3
                                    }
                                    _ = fi45
                                }
                            case "a":
                                if (0 <= 1) {
                                    v46 := GenS4{F0: fi36, F1: true}
                                    _ = v46
                                }
                            case "foo bar":
                                switch "\n\t" {
                                case "hello":
                                    for fi47 := 0; fi47 < 17; fi47++ {
                                        v3 += v3
                                        if (v6 < uint64(v5)) {
                                            var v48 uint8 = uint8(v3)
                                            _ = v48
                                        } else {
                                            v3 += (v6 | v3)
                                            v49 := GenS3{F0: -1, F1: (((v5 >> ((v1) & 63)) >= 321) && false)}
                                            _ = v49
                                        }
                                        var v50 uint = (v1 / ((v1) | 1))
                                        delete(v4, ("foo bar" + (v2 + (v2 + "a"))))
                                        fmt.Println(v1, fi47)
                                        if (0.0 == (-2.25 + 1.5)) {
                                            v51 := v4["\n\t"]
                                            v50 -= (v1 + (^(v1)))
                                            _ = v51
                                        } else {
                                            v6 -= v3
                                        }
                                        v52 := GenS5{F0: (v50 - v50), F1: v5, F2: "a"}
                                        _ = fi47
                                        _ = v50
                                        _ = v52
                                    }
                                case "a":
                                    switch (v2 + v2) {
                                    case "\n\t":
                                        if len(v34) >= 2 {
                                            v34 = v34[1:len(v34)]
                                        }
                                    case "a":
                                        var v53 uint = (v1 & uint(v6))
                                        _ = v53
                                    case "hello":
                                        v54 := GenS3{F0: 1, F1: ((!false) || (true || (true && (42 == 100))))}
                                        _ = v54
                                    default:
                                        var v55 bool = false
                                        _ = v55
                                    }
                                case "":
                                    if false {
                                        var v56 uint32 = 573
                                        v57 := v4[v2]
                                        var v58 uint16 = 1
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                default:
                                    var v59 uint32 = 1
                                    _ = v59
                                }
                            }
                        case "foo bar":
                            v60 := GenS1{F0: 1}
                            _ = v60
                        default:
                            if false {
                                v6 = v6
                            } else {
                                for fi61 := 0; fi61 < 20; fi61++ {
                                    if true {
                                        v3 -= (v6 - v3)
                                        v62 := GenS6{F0: fi36, F1: 128, F2: 2147483647}
                                        _ = v62
                                    } else {
                                        var v63 int = (fi61 | 9223372036854775807)
                                        v64 := GenS1{F0: -247}
                                        _ = v63
                                        _ = v64
                                    }
                                    _ = fi61
                                }
                                if false {
                                    if (((-2.25 + 3.14159) * 0.1) >= 1.5) {
                                        switch fi36 {
                                        case 128:
                                            fmt.Println(v5, v2)
                                        case 9223372036854775807:
                                            if (true || false) {
                                                switch (fi36 & -426) {
                                                case -234:
                                                    var v65 uint16 = v5
                                                    _ = v65
                                                case -956:
                                                    v1 = v1
                                                case 9223372036854775807:
                                                    v34 = append(v34, (fi36 << ((v1) & 63)))
                                                default:
                                                    v66 := v4[v2]
                                                    _ = v66
                                                }
                                            } else {
                                                if false {
                                                    v1 = (v1 - v1)
                                                    _ = len(v34)
                                                } else {
                                                    switch (v2 + (v2 + v2)) {
                                                    case "foo bar":
                                                        for fi67 := 0; fi67 < 1; fi67++ {
                                                            v1 = (v1 >> ((v1) & 63))
                                                            _ = fi67
                                                        }
                                                    }
                                                    if true {
                                                        switch (fi36 - (fi36 + 42)) {
                                                        case 2147483647:
                                                            for fi68 := 0; fi68 < 19; fi68++ {
                                                                var v69 int64 = 42
                                                                v5 -= (v5 >> ((v1) & 63))
                                                                _ = fi68
                                                                _ = v69
                                                            }
                                                        case 32767:
                                                            for fi70 := 0; fi70 < 15; fi70++ {
                                                                v3 = (v3 & v3)
                                                                if (!((v5 >> ((42) & 63)) > 65535)) {
                                                                    switch fi36 {
                                                                    case -9223372036854775808:
                                                                        var v71 float64 = 0.0
                                                                        _ = v71
                                                                    case 0:
                                                                        fmt.Println(fi70, v1)
                                                                    case 128:
                                                                        if false {
                                                                            if (545 == 42) {
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    if true {
                                                                                        v2 = (v2 + v2)
                                                                                    } else {
                                                                                        v6 = v6
                                                                                    }
                                                                                }
                                                                            }
                                                                            switch "" {
                                                                            case "hello":
                                                                                v34 = append(v34, fi36)
                                                                            }
                                                                            v4[v2] = 9223372036854775807
                                                                        } else {
                                                                            var v72 uint16 = uint16(v3)
                                                                            _ = v72
                                                                        }
                                                                    default:
                                                                        v1 -= v1
                                                                    }
                                                                    v73 := GenS2{F0: 2, F1: fi70, F2: 1}
                                                                    v2 = "a"
                                                                    _ = v73
                                                                } else {
                                                                    v74 := GenS7{F0: 1000, F1: uint32(v1), F2: (v1 ^ uint(v5))}
                                                                    if (v1 >= 0) {
                                                                        fmt.Println(v3, v3)
                                                                        v5 = (v5 >> ((18446744073709551615) & 63))
                                                                    }
                                                                    v5 = 255
                                                                    if false {
                                                                        if len(v34) > 0 {
                                                                            v34[0] = (fi36 * (fi70 << ((v1) & 63)))
                                                                        }
                                                                        var v75 float64 = 0.1
                                                                        switch fi70 {
                                                                        case 4294967295:
                                                                            var v76 uint = v1
                                                                            v77 := GenS4{F0: fi70, F1: true}
                                                                            fmt.Println(v3)
                                                                            _ = v76
                                                                            _ = v77
                                                                        default:
                                                                            v2 = (("hello" + "hello") + v2)
                                                                        }
                                                                        _ = v75
                                                                    } else {
                                                                        var v78 int64 = -807
                                                                        _ = v78
                                                                    }
                                                                    _ = v74
                                                                }
                                                                v5 += v5
                                                                _ = fi70
                                                            }
                                                        case 152:
                                                            v79 := GenS8{F0: (fi36 % ((fi36) | 1)), F1: v1, F2: (v2 + "\n\t")}
                                                            _ = v79
                                                        }
                                                        var v80 uint = (v1 >> (((^(v1))) & 63))
                                                        _ = v80
                                                    } else {
                                                        if (1 > 159) {
                                                            if (fi36 > (fi36 - (fi36 % ((int(v1)) | 1)))) {
                                                                for fi81 := 0; fi81 < 3; fi81++ {
                                                                    var v82 uint = v1
                                                                    _ = fi81
                                                                    _ = v82
                                                                }
                                                                v83 := GenS6{F0: fi36, F1: 157, F2: 2147483647}
                                                                fmt.Println(fi36, v1)
                                                                _ = v83
                                                            } else {
                                                                for fi84 := 0; fi84 < 15; fi84++ {
                                                                    v85 := GenS9{F0: (v1 % ((v1) | 1))}
                                                                    v86 := GenS9{F0: v1}
                                                                    var v87 uint32 = 175
                                                                    if len(v34) > 0 {
                                                                        v34[0] = fi84
                                                                    }
                                                                    _ = fi84
                                                                    _ = v85
                                                                    _ = v86
                                                                    _ = v87
                                                                }
                                                                if true {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        fmt.Println(v6, fi36, v1)
                                                                    default:
                                                                        var v88 uint64 = v3
                                                                        _ = v88
                                                                    }
                                                                    v6 += 18446744073709551615
                                                                }
                                                                v3 = 0
                                                            }
                                                        } else {
                                                            v34 = append(v34, fi36)
                                                        }
                                                        v89 := GenS10{F0: ""}
                                                        v90 := GenS8{F0: (fi36 | fi36), F1: v1, F2: "\n\t"}
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                    fmt.Println(fi36, v2)
                                                }
                                                for fi91 := 0; fi91 < 13; fi91++ {
                                                    v92 := GenS2{F0: 425, F1: fi36, F2: genFn1()}
                                                    var v93 int16 = 1
                                                    _ = fi91
                                                    _ = v92
                                                    _ = v93
                                                }
                                                for fi94 := 0; fi94 < 10; fi94++ {
                                                    _ = len(v34)
                                                    if (uint32(v1) == 0) {
                                                        switch v2 {
                                                        case "a":
                                                            v3 -= v3
                                                        }
                                                    } else {
                                                        var v95 bool = true
                                                        fmt.Println(v2, v3, v6)
                                                        v96 := GenS11{F0: 545, F1: v2, F2: 1}
                                                        var v97 bool = (((^(v1)) <= 65535) || v95)
                                                        _ = v95
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    _ = fi94
                                                }
                                            }
                                        case -631:
                                            _ = len(v34)
                                        default:
                                            var v98 int = fi36
                                            _ = v98
                                        }
                                        switch ("" + ("" + (v2 + v2))) {
                                        case "foo bar":
                                            switch int(v5) {
                                            case -1:
                                                var v99 uint = v1
                                                _ = v99
                                            }
                                        }
                                        var v100 uint = (^(v1))
                                        _ = v100
                                    }
                                    v1 = v1
                                    v3 += (v3 ^ (v6 << (((v1 % ((0) | 1))) & 63)))
                                } else {
                                    var v101 bool = false
                                    _ = v101
                                }
                                v5 += (^(v5))
                            }
                        }
                        var v102 int32 = -1
                        _ = v102
                    } else {
                        var v103 bool = true
                        _ = len(v34)
                        var v104 uint16 = genFn0((v5 & genFn0((^(v5)), v1)), (^(v1)))
                        _ = v103
                        _ = v104
                    }
                }
                _ = fi36
                _ = v37
            }
        case "foo bar":
            v3 += v3
        }
    }
    for fi105 := 0; fi105 < 7; fi105++ {
        var v106 uint8 = 176
        switch ("a" + "hello") {
        case "foo bar":
            switch v2 {
            case "a":
                if false {
                    for fi107 := 0; fi107 < 9; fi107++ {
                        v1 = (v1 << ((v1) & 63))
                        _ = fi107
                    }
                    var v108 uint = 255
                    _ = v108
                }
            case "\n\t":
                switch (fi105 << ((uint(v106)) & 63)) {
                case 127:
                    for fi109 := 0; fi109 < 16; fi109++ {
                        if (!true) {
                            var v110 uint = v1
                            _ = v110
                        }
                        fmt.Println(fi109, v5, v6)
                        v2 = "\n\t"
                        if (false && (0 > (v1 ^ uint(v106)))) {
                            var v111 uint32 = 42
                            v112 := GenS12{F0: (v106 ^ 2), F1: fi105, F2: 32767}
                            _ = v111
                            _ = v112
                        } else {
                            if false {
                                _ = len(v34)
                                if false {
                                    var v113 int16 = 1
                                    var v114 uint = (v1 | (^(v1)))
                                    _ = v113
                                    _ = v114
                                }
                                if (-9223372036854775808 != -418) {
                                    switch ((v2 + v2) + v2) {
                                    case "hello":
                                        v115 := GenS2{F0: 2, F1: fi109, F2: 2147483647}
                                        _ = v115
                                    case "foo bar":
                                        switch fi109 {
                                        case 0:
                                            var v116 float32 = -2.25
                                            _ = v116
                                        default:
                                            v117 := GenS6{F0: (fi105 - int(v106)), F1: (v106 - v106), F2: 287}
                                            _ = v117
                                        }
                                    case "":
                                        var v118 bool = (1 <= 1)
                                        _ = v118
                                    default:
                                        delete(v4, (v2 + ("" + (v2 + ("\n\t" + "hello")))))
                                    }
                                } else {
                                    var v119 int8 = 1
                                    v6 = v6
                                    v119 += (v119 << ((v1) & 63))
                                    switch (("hello" + v2) + v2) {
                                    case "":
                                        if false {
                                            v2 = v2
                                            delete(v4, v2)
                                            var v120 uint32 = uint32(v3)
                                            v121 := GenS13{F0: v2, F1: -1}
                                            _ = v120
                                            _ = v121
                                        }
                                    case "foo bar":
                                        if false {
                                            switch fi109 {
                                            case 757:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v122 int64 = -157
                                                    _ = v122
                                                default:
                                                    var v123 uint64 = 18446744073709551615
                                                    _ = v123
                                                }
                                            case -1:
                                                var v124 string = ("\n\t" + (("" + "a") + "a"))
                                                _ = v124
                                            case -2147483648:
                                                if ((v119 >> (((v1 / ((0) | 1))) & 63)) < (-(v119))) {
                                                    if false {
                                                        delete(v4, v2)
                                                        v125 := GenS14{F0: 964}
                                                        _ = v125
                                                    } else {
                                                        var v126 uint32 = 4294967295
                                                        delete(v4, "hello")
                                                        delete(v4, (("" + "") + v2))
                                                        _ = v126
                                                    }
                                                    if true {
                                                        v1 = v1
                                                    } else {
                                                        var v127 int32 = 2147483647
                                                        var v128 uint16 = (v5 | uint16(v106))
                                                        _ = v127
                                                        _ = v128
                                                    }
                                                }
                                            default:
                                                if false {
                                                    v4[v2] = fi109
                                                } else {
                                                    var v129 uint = 4294967295
                                                    _ = v129
                                                }
                                            }
                                            switch (fi105 * (^(fi109))) {
                                            case 4294967295:
                                                v130 := v4[v2]
                                                _ = v130
                                            case -650:
                                                var v131 string = ("hello" + v2)
                                                _ = v131
                                            default:
                                                var v132 int = (fi105 - fi109)
                                                _ = v132
                                            }
                                        } else {
                                            v5 = v5
                                        }
                                    }
                                    v133 := GenS5{F0: v1, F1: 32767, F2: "hello"}
                                    _ = v119
                                    _ = v133
                                }
                            }
                            if len(v34) >= 2 {
                                v34 = v34[1:len(v34)]
                            }
                            if true {
                                var v134 uint16 = (^(v5))
                                fmt.Println(v1, v6)
                                switch (fi105 >> ((v1) & 63)) {
                                case 207:
                                    var v135 int64 = 42
                                    _ = v135
                                case 683:
                                    var v136 uint16 = v5
                                    _ = v136
                                }
                                var v137 int = (fi105 << ((v1) & 63))
                                _ = v134
                                _ = v137
                            }
                        }
                        var v138 int32 = 2147483647
                        fmt.Println(v3)
                        _ = fi109
                        _ = v138
                    }
                case 42:
                    v6 -= v6
                case -569:
                    switch ("\n\t" + v2) {
                    case "hello":
                        v139 := GenS5{F0: (v1 >> ((42) & 63)), F1: v5, F2: "a"}
                        _ = v139
                    case "a":
                        _ = len(v34)
                    }
                }
            case "foo bar":
                v140 := v4[(v2 + v2)]
                _ = v140
            default:
                if len(v34) > 0 {
                    v34[0] = (fi105 * (fi105 & 32767))
                }
            }
        case "":
            fmt.Println(fi105, v6, v5)
        default:
            fmt.Println(v106)
        }
        if true {
            for fi141 := 0; fi141 < 15; fi141++ {
                var v142 int = fi141
                v6 = 4294967295
                fmt.Println(v6, fi105)
                _ = fi141
                _ = v142
            }
            for fi143 := 0; fi143 < 4; fi143++ {
                var v144 int32 = genFn1()
                v145 := GenS15{F0: 65535, F1: 9223372036854775807, F2: int(v6)}
                if false {
                    var v146 int = fi143
                    var v147 string = ("hello" + v2)
                    _ = v146
                    _ = v147
                }
                _ = fi143
                _ = v144
                _ = v145
            }
            switch "hello" {
            case "":
                for fi148 := 0; fi148 < 8; fi148++ {
                    var v149 int32 = 42
                    v150 := GenS16{F0: (-(v149)), F1: (v1 | (^(v1))), F2: v5}
                    _ = fi148
                    _ = v149
                    _ = v150
                }
            case "a":
                var v151 uint8 = v106
                _ = v151
            case "foo bar":
                fmt.Println(v2, v6, v5)
            }
        } else {
            v106 += v106
            fmt.Println(v1)
        }
        _ = fi105
        _ = v106
    }
    v152 := GenS17{F0: v1, F1: (true && (-777 > 332)), F2: genFn0(genFn0(32767, v1), 18446744073709551615)}
    v153 := GenS1{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
