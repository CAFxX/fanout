package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int16; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint8; F2 int64 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint16; F1 int; F2 int16 }
type GenS10 struct { F0 int16; F1 int64; F2 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8; F1 int64 }
type GenS13 struct { F0 int32; F1 string }
type GenS14 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS15 struct { F0 uint; F1 int8 }
type GenS16 struct { F0 int64 }
func genFn0(p0 uint, p1 int16, p2 uint64) uint {
    p1 -= p1
    for fi5 := 0; fi5 < 14; fi5++ {
        switch "hello" {
        case "a":
            v6 := []int{(-(fi5))}
            _ = v6
        }
        var v7 int16 = p1
        var v8 uint64 = 841
        _ = fi5
        _ = v7
        _ = v8
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (p0 % ((18446744073709551615) | 1))
}
func genFn1() int8 {
    switch 643 {
    case -117:
        if (9223372036854775807 != 42) {
            v10 := []int{1, 0, 672, 1}
            var v11 float64 = 0.0
            v11 = (((v11 / (v11 + (v11 / float64(1.5)))) * v11) / v11)
            _ = v10
            _ = v11
        }
    case -2147483648:
        switch "hello" {
        case "hello":
            if false {
                for fi12 := 0; fi12 < 8; fi12++ {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        fmt.Println(fi13)
                        v14 := []int{}
                        _ = fi13
                        _ = v14
                    }
                    var v15 uint = 0
                    _ = fi12
                    _ = v15
                }
                var v16 uint64 = 7
                _ = v16
            }
        default:
            var v17 int = 127
            _ = v17
        }
    default:
        var v18 uint16 = 255
        _ = v18
    }
    v19 := []int{}
    for fi20 := 0; fi20 < 7; fi20++ {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        fmt.Println(fi20)
        if (127 != 35) {
            if true {
                fmt.Println(fi20)
                switch fi20 {
                case 32767:
                    fmt.Println(fi20)
                case -964:
                    v21 := GenS1{F0: -550, F1: 4294967295, F2: 0}
                    _ = v21
                case 127:
                    switch "" {
                    case "a":
                        var v22 uint64 = 562
                        _ = v22
                    case "":
                        var v23 uint = 1000
                        _ = v23
                    case "\n\t":
                        var v24 uint = 0
                        _ = v24
                    default:
                        var v25 uint16 = 145
                        _ = v25
                    }
                }
            }
        } else {
            v26 := GenS1{F0: 0, F1: 739, F2: 4294967295}
            if false {
                for fi27 := 0; fi27 < 9; fi27++ {
                    if false {
                        var v28 uint8 = 1
                        v28 = 255
                        var v29 string = "foo bar"
                        _ = v28
                        _ = v29
                    } else {
                        fmt.Println(fi20, fi20)
                        switch "a" {
                        case "\n\t":
                            var v30 float64 = -2.25
                            var v31 uint64 = 718
                            v30 -= v30
                            _ = v30
                            _ = v31
                        case "a":
                            var v32 uint8 = 116
                            _ = v32
                        case "hello":
                            fmt.Println(fi27)
                        }
                    }
                    var v33 bool = false
                    _ = fi27
                    _ = v33
                }
                v34 := GenS1{F0: -1, F1: 4294967295, F2: 2}
                _ = v34
            } else {
                var v35 uint16 = 32767
                _ = v35
            }
            _ = v26
        }
        _ = fi20
    }
    _ = v19
    return 65
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]int)
    v3 := []uint32{4294967295, 65535, 1}
    var v4 int8 = 0
    var v36 uint64 = 533
    v36 = (^(v36))
    if ((v4 ^ v4) != (v4 / ((v4) | 1))) {
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := GenS2{F0: (^(v4))}
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 3; fi39++ {
            v40 := GenS2{F0: (v4 & v4)}
            if (true || (false || false)) {
                fmt.Println(v36, v4)
            } else {
                fmt.Println(fi39, v1, fi39)
                v4 += v4
                var v41 float32 = 3.14159
                _ = v41
            }
            if (!false) {
                var v42 uint32 = 2147483647
                v1 = v1
                if ("\n\t" == "a") {
                    for fi43 := 0; fi43 < 13; fi43++ {
                        fmt.Println(fi39)
                        v44 := GenS3{F0: (v42 << ((18446744073709551615) & 63)), F1: "foo bar"}
                        _ = fi43
                        _ = v44
                    }
                    v1 = v1
                } else {
                    fmt.Println(v4, v4)
                    v45 := GenS2{F0: genFn1()}
                    _ = v45
                }
                _ = v42
            }
            _ = fi39
            _ = v40
        }
    } else {
        var v46 uint32 = 2
        var v47 float32 = float32(1.5)
        fmt.Println(v4, v46)
        _ = v46
        _ = v47
    }
    v36 = (v36 | (v36 & (v36 | (v36 & (v36 & (^(v36)))))))
    for fi48 := 0; fi48 < 18; fi48++ {
        v49 := GenS4{F0: v36}
        switch "a" {
        case "a":
            v36 += 42
        case "":
            v50 := GenS5{F0: -2147483648}
            _ = v50
        }
        v51 := v2[fi48]
        _ = fi48
        _ = v49
        _ = v51
    }
    v1 += 3.14159
    v52 := GenS6{F0: true, F1: int16(v4), F2: 219}
    v36 += (v36 >> ((genFn0(27, 618, 4294967295)) & 63))
    var v53 int64 = 981
    if ((false || ((32768 != 128) && false)) || true) {
        v53 = (^(v53))
        for fi54 := 0; fi54 < 17; fi54++ {
            var v55 uint = 65535
            for fi56 := 0; fi56 < 16; fi56++ {
                var v57 int16 = 42
                switch int(v53) {
                case 9223372036854775807:
                    fmt.Println(v1, fi56)
                case 424:
                    var v58 uint16 = 65535
                    _ = v58
                case 128:
                    if true {
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        if ((!(1 < 1)) || false) {
                            fmt.Println(v53, v1, fi56)
                            v59 := GenS7{F0: 804, F1: 42, F2: -9223372036854775808}
                            _ = v59
                        }
                        v60 := GenS8{F0: (^(v53)), F1: ((-(v4)) <= (-(v4)))}
                        switch fi54 {
                        case 42:
                            v36 += 791
                        case 1:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            _ = len(v3)
                        }
                        _ = v60
                    }
                default:
                    v57 = v57
                }
                v4 = v4
                _ = fi56
                _ = v57
            }
            v1 = float64(float32(v1))
            _ = fi54
            _ = v55
        }
    } else {
        fmt.Println(v53, v1, v4)
        switch "foo bar" {
        case "a":
            var v61 int8 = v4
            _ = v61
        default:
            v2[4294967295] = 924
        }
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        for fi63 := 0; fi63 < 11; fi63++ {
            v36 += (^(v36))
            _ = fi63
        }
        v64 := GenS9{F0: 1, F1: fi62, F2: 1000}
        _ = fi62
        _ = v64
    }
    if false {
        v53 -= -1
        v4 -= 127
        fmt.Println(v53, v4)
    }
    if false {
        var v65 uint = 255
        _ = v65
    } else {
        switch "a" {
        case "":
            v1 = float64(0.0)
        default:
            v1 -= v1
        }
        for fi66 := 0; fi66 < 6; fi66++ {
            fmt.Println(v1, v53, v1)
            v67 := v2[-405]
            _ = fi66
            _ = v67
        }
        if (!(true && true)) {
            switch "\n\t" {
            case "":
                fmt.Println(v53)
            case "hello":
                v68 := GenS4{F0: 751}
                _ = v68
            default:
                var v69 int32 = 367
                _ = v69
            }
        } else {
            if (761 >= 42) {
                for fi70 := 0; fi70 < 18; fi70++ {
                    var v71 string = ("hello" + (("\n\t" + ("a" + "hello")) + ("" + "\n\t")))
                    v1 += 3.14159
                    var v72 int16 = 1000
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                if (0 != uint(v36)) {
                    if (!false) {
                        switch ("" + "foo bar") {
                        case "":
                            var v73 uint8 = 1
                            _ = v73
                        }
                        fmt.Println(v36, v4, v53)
                        var v74 string = (((("\n\t" + "a") + "a") + ("hello" + "a")) + ((("a" + "foo bar") + "\n\t") + "hello"))
                        _ = v74
                    }
                    _ = len(v2)
                    var v75 int32 = 1
                    _ = v75
                } else {
                    switch -9223372036854775808 {
                    case 4294967295:
                        switch -2147483648 {
                        case 490:
                            v4 -= genFn1()
                        case 2147483647:
                            var v76 uint16 = 42
                            _ = v76
                        default:
                            for fi77 := 0; fi77 < 1; fi77++ {
                                fmt.Println(v36, v36)
                                _ = len(v2)
                                _ = fi77
                            }
                        }
                    case 32768:
                        for fi78 := 0; fi78 < 10; fi78++ {
                            switch (fi78 + fi78) {
                            case 127:
                                if (!true) {
                                    delete(v2, fi78)
                                }
                            case -601:
                                for fi79 := 0; fi79 < 9; fi79++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v80 int = fi78
                                        _ = v80
                                    }
                                    _ = fi79
                                }
                            case 42:
                                var v81 int = 302
                                _ = v81
                            }
                            for fi82 := 0; fi82 < 19; fi82++ {
                                if (((float32(v1) * 1e10) * 3.14159) >= 1.5) {
                                    _ = len(v2)
                                    switch (fi78 / ((32767) | 1)) {
                                    case 0:
                                        v4 = (v4 % (((v4 - (v4 | (v4 - (v4 * v4))))) | 1))
                                        var v83 int16 = 42
                                        fmt.Println(v4, fi82)
                                        _ = v83
                                    }
                                }
                                if (!(!(0 < 4294967295))) {
                                    var v84 int = -2147483648
                                    v53 = v53
                                    v84 -= -1
                                    _ = v84
                                } else {
                                    if false {
                                        v53 += (v53 % ((9223372036854775807) | 1))
                                        v85 := GenS8{F0: v53, F1: (v53 == 992)}
                                        fmt.Println(v36)
                                        _ = v85
                                    }
                                }
                                v4 += (v4 | v4)
                                _ = len(v2)
                                var v86 string = ("\n\t" + "")
                                _ = fi82
                                _ = v86
                            }
                            _ = fi78
                        }
                    case 1:
                        v53 += (v53 + v53)
                    default:
                        if (1 < 0) {
                            v36 = 0
                            if ("a" == "") {
                                var v87 string = ("a" + "hello")
                                v4 += (^(v4))
                                var v88 uint16 = 32767
                                _ = v87
                                _ = v88
                            } else {
                                v53 += v53
                                var v89 float32 = (((1.5 / 3.14159) * float32(-2.25)) - 1e10)
                                if (54 > 42) {
                                    for fi90 := 0; fi90 < 8; fi90++ {
                                        switch -1 {
                                        case -2147483648:
                                            var v91 int64 = v53
                                            _ = v91
                                        case 42:
                                            v36 = (v36 ^ v36)
                                        default:
                                            var v92 int8 = (v4 << ((1) & 63))
                                            _ = v92
                                        }
                                        _ = fi90
                                    }
                                } else {
                                    if true {
                                        switch ("\n\t" + ("a" + ("" + ("" + ("\n\t" + (("\n\t" + "") + "hello")))))) {
                                        case "\n\t":
                                            var v93 int16 = 1
                                            _ = v93
                                        case "foo bar":
                                            _ = len(v3)
                                        default:
                                            if ((("a" + (("\n\t" + "") + ("foo bar" + "hello"))) + "") > "foo bar") {
                                                if (false && (((3.14159 / v1) + (v1 - -2.25)) < -2.25)) {
                                                    v3 = append(v3, 65535)
                                                    for fi94 := 0; fi94 < 6; fi94++ {
                                                        if ((v4 == v4) || (false && (148 == 1))) {
                                                            if (42 != 1000) {
                                                                switch (fi94 / ((fi94) | 1)) {
                                                                case 684:
                                                                    switch -9223372036854775808 {
                                                                    case -9223372036854775808:
                                                                        for fi95 := 0; fi95 < 4; fi95++ {
                                                                            v96 := GenS10{F0: 42, F1: (v53 % (((v53 * (v53 | (v53 % (((v53 & v53)) | 1))))) | 1)), F2: 2147483647}
                                                                            v53 = (-(v53))
                                                                            fmt.Println(v53)
                                                                            _ = fi95
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                case 0:
                                                                    for fi97 := 0; fi97 < 1; fi97++ {
                                                                        v98 := GenS11{F0: v4}
                                                                        _ = fi97
                                                                        _ = v98
                                                                    }
                                                                case 4294967295:
                                                                    var v99 uint = 65535
                                                                    _ = v99
                                                                default:
                                                                    if (v53 <= (v53 * -9223372036854775808)) {
                                                                        fmt.Println(v36, v53)
                                                                        v1 = ((v1 * v1) * v1)
                                                                    }
                                                                }
                                                            }
                                                            v100 := GenS10{F0: -32768, F1: (v53 - 9223372036854775807), F2: -1}
                                                            _ = v100
                                                        }
                                                        for fi101 := 0; fi101 < 15; fi101++ {
                                                            var v102 string = "a"
                                                            _ = fi101
                                                            _ = v102
                                                        }
                                                        var v103 bool = true
                                                        _ = fi94
                                                        _ = v103
                                                    }
                                                } else {
                                                    var v104 bool = false
                                                    var v105 uint64 = 253
                                                    v106 := GenS5{F0: 52}
                                                    _ = v104
                                                    _ = v105
                                                    _ = v106
                                                }
                                                if ((42 < v36) && (1.5 >= -2.25)) {
                                                    if true {
                                                        for fi107 := 0; fi107 < 5; fi107++ {
                                                            if (472 > -883) {
                                                                switch fi107 {
                                                                case -871:
                                                                    var v108 uint8 = 1
                                                                    _ = v108
                                                                case 42:
                                                                    for fi109 := 0; fi109 < 9; fi109++ {
                                                                        switch ("" + "hello") {
                                                                        case "\n\t":
                                                                            v53 += 1
                                                                        case "foo bar":
                                                                            var v110 uint32 = 788
                                                                            _ = v110
                                                                        case "hello":
                                                                            v111 := v2[(fi107 >> ((743) & 63))]
                                                                            _ = v111
                                                                        default:
                                                                            if (true && ((true && false) || (("\n\t" > "foo bar") || false))) {
                                                                                var v112 uint64 = (^(v36))
                                                                                var v113 uint32 = 953
                                                                                if false {
                                                                                    v112 = v112
                                                                                    fmt.Println(v112, fi109)
                                                                                }
                                                                                _ = v112
                                                                                _ = v113
                                                                            }
                                                                        }
                                                                        _ = len(v2)
                                                                        _ = fi109
                                                                    }
                                                                case -670:
                                                                    var v114 int32 = 1
                                                                    _ = v114
                                                                }
                                                                if ("foo bar" < ((("foo bar" + "") + ("\n\t" + "")) + ((("foo bar" + "foo bar") + ("a" + (("a" + "foo bar") + "foo bar"))) + "foo bar"))) {
                                                                    v115 := GenS4{F0: (v36 % ((v36) | 1))}
                                                                    for fi116 := 0; fi116 < 12; fi116++ {
                                                                        if false {
                                                                            var v117 uint8 = 144
                                                                            var v118 float64 = v1
                                                                            v119 := GenS3{F0: 2147483648, F1: ((("foo bar" + "\n\t") + "") + "\n\t")}
                                                                            _ = v117
                                                                            _ = v118
                                                                            _ = v119
                                                                        }
                                                                        var v120 bool = false
                                                                        _ = fi116
                                                                        _ = v120
                                                                    }
                                                                    v89 = (v89 - (v89 + (0.0 - v89)))
                                                                    _ = v115
                                                                } else {
                                                                    v36 -= v36
                                                                    switch (fi107 & (fi107 >> ((635) & 63))) {
                                                                    case -2147483648:
                                                                        for fi121 := 0; fi121 < 11; fi121++ {
                                                                            if (42 >= 42) {
                                                                                v122 := GenS11{F0: (v4 / ((0) | 1))}
                                                                                _ = v122
                                                                            }
                                                                            var v123 uint16 = 2
                                                                            _ = fi121
                                                                            _ = v123
                                                                        }
                                                                    case -9223372036854775808:
                                                                        v3 = append(v3, 2)
                                                                    default:
                                                                        v89 = 3.14159
                                                                    }
                                                                    var v124 bool = ((^(v53)) < -1)
                                                                    _ = v124
                                                                }
                                                            } else {
                                                                if false {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        for fi125 := 0; fi125 < 14; fi125++ {
                                                                            v4 = v4
                                                                            v89 = 1.5
                                                                            v126 := GenS6{F0: true, F1: 0, F2: 255}
                                                                            _ = fi125
                                                                            _ = v126
                                                                        }
                                                                    case "hello":
                                                                        switch fi107 {
                                                                        case 42:
                                                                            if (!true) {
                                                                                if true {
                                                                                    if true {
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            delete(v2, fi107)
                                                                                        default:
                                                                                            var v127 int16 = 3
                                                                                            _ = v127
                                                                                        }
                                                                                        fmt.Println(v4)
                                                                                    } else {
                                                                                        var v128 uint = genFn0(255, 42, v36)
                                                                                        var v129 int32 = -2147483648
                                                                                        v53 -= v53
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                    }
                                                                                }
                                                                                fmt.Println(v53, v89)
                                                                            }
                                                                        }
                                                                    }
                                                                    var v130 int64 = (v53 / ((int64(v4)) | 1))
                                                                    var v131 bool = (2 > 307)
                                                                    _ = v130
                                                                    _ = v131
                                                                } else {
                                                                    switch fi107 {
                                                                    case 127:
                                                                        for fi132 := 0; fi132 < 10; fi132++ {
                                                                            if false {
                                                                                v4 = 19
                                                                                fmt.Println(v4)
                                                                            } else {
                                                                                var v133 uint8 = 207
                                                                                switch 128 {
                                                                                case -2147483648:
                                                                                    var v134 uint64 = (v36 >> ((0) & 63))
                                                                                    _ = v134
                                                                                case 830:
                                                                                    fmt.Println(v89, v53)
                                                                                default:
                                                                                    var v135 uint16 = uint16(v36)
                                                                                    _ = v135
                                                                                }
                                                                                _ = v133
                                                                            }
                                                                            var v136 int16 = int16(fi132)
                                                                            _ = fi132
                                                                            _ = v136
                                                                        }
                                                                    case -740:
                                                                        if true {
                                                                            var v137 int8 = (v4 >> ((0) & 63))
                                                                            switch fi107 {
                                                                            case -1:
                                                                                v138 := GenS6{F0: false, F1: 42, F2: 148}
                                                                                _ = v138
                                                                            }
                                                                            _ = v137
                                                                        }
                                                                    case 32768:
                                                                        v139 := v2[(fi107 % ((-843) | 1))]
                                                                        _ = v139
                                                                    default:
                                                                        if false {
                                                                            var v140 uint = 30
                                                                            _ = v140
                                                                        } else {
                                                                            for fi141 := 0; fi141 < 20; fi141++ {
                                                                                if ((^(v53)) < (v53 << ((473) & 63))) {
                                                                                    if false {
                                                                                        switch "\n\t" {
                                                                                        case "foo bar":
                                                                                            var v142 uint8 = 194
                                                                                            _ = v142
                                                                                        case "a":
                                                                                            v143 := GenS9{F0: 0, F1: fi107, F2: 382}
                                                                                            _ = v143
                                                                                        default:
                                                                                            delete(v2, fi107)
                                                                                        }
                                                                                        var v144 uint32 = 65535
                                                                                        fmt.Println(v1, v1)
                                                                                        _ = v144
                                                                                    }
                                                                                    var v145 int64 = v53
                                                                                    _ = v145
                                                                                }
                                                                                v146 := GenS12{F0: (v4 * 0), F1: (v53 * -1)}
                                                                                var v147 uint8 = 127
                                                                                fmt.Println(v89, v89, v4)
                                                                                _ = fi141
                                                                                _ = v146
                                                                                _ = v147
                                                                            }
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v36 = 18446744073709551615
                                                                            case "hello":
                                                                                v148 := GenS13{F0: 42, F1: ((("" + "hello") + "a") + "foo bar")}
                                                                                _ = v148
                                                                            default:
                                                                                var v149 string = "\n\t"
                                                                                _ = v149
                                                                            }
                                                                            for fi150 := 0; fi150 < 6; fi150++ {
                                                                                switch ("foo bar" + ("" + "")) {
                                                                                case "foo bar":
                                                                                    var v151 int16 = 42
                                                                                    _ = v151
                                                                                case "\n\t":
                                                                                    delete(v2, fi107)
                                                                                default:
                                                                                    if false {
                                                                                        switch fi107 {
                                                                                        case 4294967295:
                                                                                            fmt.Println(v89, v4)
                                                                                        default:
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                        delete(v2, (^(fi150)))
                                                                                        v152 := GenS10{F0: 32767, F1: v53, F2: -789}
                                                                                        v153 := GenS14{F0: v36, F1: 127, F2: (v36 | (v36 % ((v36) | 1)))}
                                                                                        _ = v152
                                                                                        _ = v153
                                                                                    }
                                                                                }
                                                                                _ = fi150
                                                                            }
                                                                        }
                                                                    }
                                                                    v154 := GenS9{F0: 65535, F1: fi107, F2: -955}
                                                                    for fi155 := 0; fi155 < 11; fi155++ {
                                                                        fmt.Println(v53, fi107)
                                                                        var v156 uint8 = 2
                                                                        _ = len(v2)
                                                                        _ = fi155
                                                                        _ = v156
                                                                    }
                                                                    _ = v154
                                                                }
                                                                switch int(v36) {
                                                                case 0:
                                                                    if false {
                                                                        for fi157 := 0; fi157 < 17; fi157++ {
                                                                            fmt.Println(v89, fi157, v53)
                                                                            _ = fi157
                                                                        }
                                                                        v3 = append(v3, 65535)
                                                                    }
                                                                case -9223372036854775808:
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        _ = len(v2)
                                                                    }
                                                                default:
                                                                    for fi158 := 0; fi158 < 2; fi158++ {
                                                                        v4 += (v4 << ((genFn0(503, 42, (v36 % ((v36) | 1)))) & 63))
                                                                        _ = fi158
                                                                    }
                                                                }
                                                                fmt.Println(v89)
                                                            }
                                                            _ = fi107
                                                        }
                                                    }
                                                    v4 = (v4 % ((100) | 1))
                                                }
                                                v89 = v89
                                            } else {
                                                switch ("hello" + ("foo bar" + "hello")) {
                                                case "":
                                                    if len(v3) > 0 {
                                                        v3[0] = 555
                                                    }
                                                case "a":
                                                    var v159 int = 4294967295
                                                    _ = v159
                                                default:
                                                    if true {
                                                        v160 := GenS3{F0: 303, F1: ((("foo bar" + "") + "\n\t") + "hello")}
                                                        _ = v160
                                                    } else {
                                                        fmt.Println(v89, v53, v36)
                                                        for fi161 := 0; fi161 < 11; fi161++ {
                                                            v1 = v1
                                                            _ = fi161
                                                        }
                                                        v4 -= -128
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        var v162 uint8 = 42
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        for fi163 := 0; fi163 < 5; fi163++ {
                                            v36 = 666
                                            _ = fi163
                                        }
                                        _ = v162
                                    }
                                    if false {
                                        switch ("hello" + "\n\t") {
                                        case "foo bar":
                                            var v164 uint8 = 1
                                            _ = v164
                                        default:
                                            switch 42 {
                                            case 2147483647:
                                                fmt.Println(v4, v53, v89)
                                            case -545:
                                                var v165 float64 = (((v1 + float64(((v89 - v89) / (v89 - v89)))) / ((1e10 - float64(float32(0.0))) + (v1 * v1))) * v1)
                                                _ = v165
                                            case 0:
                                                switch "a" {
                                                case "":
                                                    v89 += v89
                                                case "foo bar":
                                                    switch ("" + "") {
                                                    case "\n\t":
                                                        v166 := GenS15{F0: 2, F1: v4}
                                                        _ = v166
                                                    }
                                                case "hello":
                                                    v89 += (3.14159 - 0.1)
                                                }
                                            default:
                                                v167 := v2[127]
                                                _ = v167
                                            }
                                        }
                                        v168 := GenS7{F0: (^(v53)), F1: 42, F2: v53}
                                        _ = v168
                                    }
                                }
                                _ = v89
                            }
                        } else {
                            v169 := GenS3{F0: 42, F1: "foo bar"}
                            for fi170 := 0; fi170 < 8; fi170++ {
                                v1 = 1.5
                                v171 := GenS12{F0: genFn1(), F1: v53}
                                var v172 uint8 = uint8(v36)
                                _ = fi170
                                _ = v171
                                _ = v172
                            }
                            _ = v169
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    var v173 string = "foo bar"
                    _ = v173
                }
            } else {
                switch ("a" + "hello") {
                case "hello":
                    if (!true) {
                        fmt.Println(v36, v1)
                        v36 -= v36
                    } else {
                        var v174 uint64 = 42
                        v174 += v36
                        v175 := GenS4{F0: v174}
                        _ = v174
                        _ = v175
                    }
                case "foo bar":
                    var v176 uint64 = v36
                    _ = v176
                case "\n\t":
                    v177 := GenS16{F0: -9223372036854775808}
                    _ = v177
                default:
                    fmt.Println(v1, v1, v1)
                }
            }
            for fi178 := 0; fi178 < 2; fi178++ {
                if (-1 == -1) {
                    _ = len(v3)
                }
                switch 32767 {
                case 128:
                    var v179 float32 = (float32((((float64((((0.1 + (3.14159 + 0.1)) * (0.0 - 0.1)) * 1e10)) + v1) * 0.1) / (v1 - v1))) - (1.5 * (0.1 * 1.5)))
                    _ = v179
                case -246:
                    var v180 bool = (false || (45 == 255))
                    _ = v180
                case -1:
                    v53 = v53
                }
                v36 = v36
                _ = fi178
            }
            v1 = (1e10 / (v1 - 3.14159))
        }
    }
    v36 += v36
    if (!(571 < 1)) {
        v181 := GenS11{F0: genFn1()}
        v182 := GenS13{F0: -174, F1: "foo bar"}
        var v183 int = 42
        _ = v181
        _ = v182
        _ = v183
    }
    for fi184 := 0; fi184 < 1; fi184++ {
        switch fi184 {
        case 4294967295:
            _ = len(v3)
        }
        _ = fi184
    }
    var v185 uint32 = 4294967295
    for fi186 := 0; fi186 < 10; fi186++ {
        v36 -= v36
        _ = fi186
    }
    fmt.Println("v1:", v1)
    fmt.Println("v185:", v185)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v185
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v52
    _ = v53
}
