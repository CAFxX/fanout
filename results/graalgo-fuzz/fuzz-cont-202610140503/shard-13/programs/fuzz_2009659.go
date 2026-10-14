package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint64; F1 bool; F2 string }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int32; F1 int8; F2 string }
type GenS8 struct { F0 int16; F1 uint16; F2 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8; F1 uint8 }
type GenS12 struct { F0 uint8; F1 int; F2 uint }
type GenS13 struct { F0 int8; F1 bool; F2 int }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int16; F1 int; F2 uint64 }
type GenS17 struct { F0 uint64; F1 int8 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint8; F1 uint32; F2 uint64 }
func genFn0() uint64 {
    var v7 uint32 = 2
    if ((("a" + "\n\t") >= "foo bar") || (false && ((true || false) && (4294967295 == 727)))) {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint = 65535
            fmt.Println(v9, fi8, fi8)
            _ = fi8
            _ = v9
        }
        var v10 uint32 = v7
        var v11 uint32 = 851
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 string = ("a" + "hello")
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int64 = -432
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int = 1
                _ = fi16
                _ = v17
            }
            if (false && (((!true) && ((fi14 ^ fi14) >= 32768)) || ((1 <= 414) || false))) {
                var v18 uint32 = v7
                if true {
                    var v19 int8 = 100
                    if false {
                        v20 := []int{fi14, fi14, (fi14 & (fi14 & (fi14 / ((fi14) | 1)))), 32768}
                        _ = v20
                    } else {
                        if (529 <= 611) {
                            v21 := GenS1{F0: 406}
                            switch (^(fi14)) {
                            case 531:
                                if (!false) {
                                    if true {
                                        var v22 string = "\n\t"
                                        v18 += v7
                                        _ = v22
                                    }
                                    var v23 bool = true
                                    _ = v23
                                } else {
                                    v24 := GenS1{F0: 706}
                                    fmt.Println(v18, fi14)
                                    v7 = v7
                                    _ = v24
                                }
                            default:
                                v25 := []int{fi14}
                                _ = v25
                            }
                            switch (-(fi14)) {
                            case 447:
                                v15 -= v15
                            case 1:
                                v26 := []int{-9223372036854775808, fi14}
                                _ = v26
                            case -994:
                                var v27 uint16 = 255
                                _ = v27
                            }
                            _ = v21
                        }
                        if (v15 < v15) {
                            v28 := GenS1{F0: 0}
                            _ = v28
                        }
                        v29 := GenS1{F0: 961}
                        _ = v29
                    }
                    _ = v19
                }
                _ = v18
            }
            _ = fi14
            _ = v15
        }
    }
    _ = v7
    return uint64(v7)
}
func genFn1() int {
    var v30 string = "hello"
    v31 := GenS2{F0: 42, F1: 224}
    var v32 int64 = -9223372036854775808
    var v33 uint64 = genFn0()
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    return 123
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (true || false)
    var v3 uint64 = 290
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    var v6 bool = true
    v34 := GenS3{F0: ("\n\t" + "\n\t"), F1: 0, F2: 699}
    v3 = 18446744073709551615
    fmt.Println(v2, v6, v3)
    for fi35 := 0; fi35 < 16; fi35++ {
        _ = len(v4)
        for fi36 := 0; fi36 < 3; fi36++ {
            v3 += v3
            v37 := GenS4{F0: v3, F1: ((2147483647 > -1) || v6), F2: "foo bar"}
            v38 := GenS5{F0: (fi35 * (fi35 << ((1) & 63))), F1: 42}
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    v39 := v5[("" + "foo bar")]
    v39 = v39
    for fi40 := 0; fi40 < 13; fi40++ {
        v41 := GenS6{F0: 42, F1: 128}
        if ((42 >= 1) && v6) {
            switch fi40 {
            case 4294967295:
                for fi42 := 0; fi42 < 14; fi42++ {
                    switch "a" {
                    case "\n\t":
                        v6 = v6
                    case "hello":
                        var v43 uint32 = 2147483647
                        _ = v43
                    default:
                        var v44 int32 = 0
                        _ = v44
                    }
                    v39 = v39
                    v45 := v1[42]
                    _ = fi42
                    _ = v45
                }
            default:
                for fi46 := 0; fi46 < 12; fi46++ {
                    v47 := []int{genFn1(), (fi40 & 42), fi40, 4294967295}
                    _ = fi46
                    _ = v47
                }
            }
            v48 := v4[("hello" + "a")]
            if true {
                var v49 int64 = -10
                _ = v49
            } else {
                var v50 uint8 = 7
                var v51 uint8 = (^(v50))
                if (((fi40 & (fi40 + -2147483648)) >= (fi40 % ((fi40) | 1))) || (v48 > 2)) {
                    v52 := []int{-9223372036854775808, genFn1()}
                    if ((!(v2 || false)) || ((v48 + 42) != v48)) {
                        if (127 <= -1) {
                            switch "foo bar" {
                            case "a":
                                if true {
                                    if v39 {
                                        v50 = v51
                                        if (2147483647 > 1) {
                                            v53 := v4["\n\t"]
                                            var v54 int32 = 42
                                            _ = v53
                                            _ = v54
                                        }
                                        v39 = v39
                                    } else {
                                        fmt.Println(v2)
                                        v3 = genFn0()
                                        switch "" {
                                        case "a":
                                            if (!(v2 && (false && v6))) {
                                                _ = len(v52)
                                                v2 = (fi40 <= -9223372036854775808)
                                            }
                                        case "hello":
                                            for fi55 := 0; fi55 < 3; fi55++ {
                                                v39 = true
                                                delete(v5, "hello")
                                                var v56 uint32 = (v48 ^ v48)
                                                fmt.Println(fi40, v50, v50)
                                                var v57 uint64 = 18446744073709551615
                                                _ = fi55
                                                _ = v56
                                                _ = v57
                                            }
                                        case "\n\t":
                                            switch genFn1() {
                                            case -2147483648:
                                                for fi58 := 0; fi58 < 17; fi58++ {
                                                    v6 = true
                                                    v59 := GenS3{F0: "", F1: -1, F2: 1}
                                                    var v60 int64 = -1
                                                    _ = fi58
                                                    _ = v59
                                                    _ = v60
                                                }
                                            case -1:
                                                var v61 uint8 = (v51 * (v51 << ((65535) & 63)))
                                                _ = v61
                                            case 0:
                                                for fi62 := 0; fi62 < 1; fi62++ {
                                                    delete(v4, "a")
                                                    if (!false) {
                                                        switch "" {
                                                        case "foo bar":
                                                            var v63 int16 = 1
                                                            v50 = 105
                                                            v3 = (v3 / (((v3 % (((v3 % ((v3) | 1))) | 1))) | 1))
                                                            _ = v63
                                                        case "\n\t":
                                                            switch "" {
                                                            case "a":
                                                                v64 := GenS7{F0: 262, F1: 1, F2: ""}
                                                                _ = v64
                                                            case "foo bar":
                                                                _ = len(v52)
                                                            case "\n\t":
                                                                if (int16(fi62) >= 32767) {
                                                                    var v65 uint = 255
                                                                    var v66 int32 = -1
                                                                    _ = v65
                                                                    _ = v66
                                                                } else {
                                                                    v1[v48] = (fi62 | (fi62 >> ((0) & 63)))
                                                                    fmt.Println(v48, v3)
                                                                }
                                                            }
                                                        default:
                                                            v67 := GenS8{F0: 184, F1: 1, F2: 65535}
                                                            _ = v67
                                                        }
                                                        fmt.Println(fi62)
                                                        if v39 {
                                                            delete(v1, 42)
                                                        } else {
                                                            v50 += 0
                                                            v68 := GenS3{F0: "a", F1: 2147483647, F2: 255}
                                                            v39 = v6
                                                            _ = v68
                                                        }
                                                    } else {
                                                        v2 = (-2147483648 >= 1)
                                                    }
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v69 int32 = -2147483648
                                                        _ = v69
                                                    case "hello":
                                                        v70 := GenS9{F0: genFn0()}
                                                        _ = v70
                                                    default:
                                                        v50 -= 0
                                                    }
                                                    _ = fi62
                                                }
                                            }
                                        default:
                                            var v71 int = genFn1()
                                            _ = v71
                                        }
                                    }
                                    delete(v4, "\n\t")
                                }
                            case "hello":
                                for fi72 := 0; fi72 < 18; fi72++ {
                                    v2 = ((false && v39) || (-2.25 < (0.0 * (3.14159 / float32(float64(1.5))))))
                                    if (!(!v2)) {
                                        var v73 int8 = -56
                                        _ = v73
                                    } else {
                                        var v74 int8 = 1
                                        _ = v74
                                    }
                                    var v75 float64 = 1.5
                                    v3 = v3
                                    _ = fi72
                                    _ = v75
                                }
                            }
                            fmt.Println(v50, v50, v51)
                            for fi76 := 0; fi76 < 11; fi76++ {
                                v77 := v1[v48]
                                _ = fi76
                                _ = v77
                            }
                        } else {
                            for fi78 := 0; fi78 < 11; fi78++ {
                                v48 = (v48 * (^(v48)))
                                v79 := v4["foo bar"]
                                _ = fi78
                                _ = v79
                            }
                            var v80 int32 = 981
                            _ = v80
                        }
                        v39 = v39
                    }
                    var v81 bool = (!(!((v3 & 0) > uint64(v50))))
                    _ = v52
                    _ = v81
                } else {
                    var v82 uint = 547
                    v50 -= 1
                    _ = v82
                }
                _ = v50
                _ = v51
            }
            _ = v48
        } else {
            var v83 int8 = 42
            if v2 {
                switch (("a" + ("" + "")) + ("foo bar" + ((((("" + "a") + "a") + "\n\t") + "hello") + ("" + "\n\t")))) {
                case "a":
                    _ = len(v1)
                }
            } else {
                var v84 uint32 = 723
                switch (fi40 | genFn1()) {
                case -2147483648:
                    _ = len(v1)
                case -179:
                    if ((-(v83)) != (v83 ^ (v83 << ((18446744073709551615) & 63)))) {
                        v39 = (v6 && v39)
                    } else {
                        if v2 {
                            v85 := GenS10{F0: (v83 / ((v83) | 1))}
                            _ = v85
                        }
                    }
                }
                if v39 {
                    var v86 string = "foo bar"
                    _ = v86
                }
                _ = v84
            }
            var v87 int64 = 9223372036854775807
            _ = v83
            _ = v87
        }
        var v88 uint = 356
        _ = fi40
        _ = v41
        _ = v88
    }
    v3 -= v3
    var v89 float32 = 0.0
    var v90 int64 = 42
    var v91 int8 = 30
    if true {
        v6 = (18446744073709551615 < 208)
    } else {
        switch "hello" {
        case "foo bar":
            if v6 {
                for fi92 := 0; fi92 < 5; fi92++ {
                    v93 := GenS5{F0: -1, F1: 0}
                    _ = fi92
                    _ = v93
                }
            }
        case "hello":
            for fi94 := 0; fi94 < 13; fi94++ {
                var v95 int32 = 355
                for fi96 := 0; fi96 < 13; fi96++ {
                    fmt.Println(v6, fi94)
                    var v97 uint64 = 0
                    switch fi96 {
                    case 444:
                        v3 = (^(v97))
                    case -9223372036854775808:
                        switch genFn1() {
                        case 32768:
                            switch ("\n\t" + "a") {
                            case "foo bar":
                                switch 32767 {
                                case 9223372036854775807:
                                    if (!v2) {
                                        switch (("\n\t" + ("hello" + "hello")) + ("hello" + "hello")) {
                                        case "a":
                                            if (!v2) {
                                                v98 := GenS4{F0: v3, F1: ((((float64(v89) / 1e10) + 0.1) / 0.1) == (((0.0 / 3.14159) / 1.0) / 1.0)), F2: "a"}
                                                v5[("foo bar" + ("\n\t" + ("a" + "a")))] = (!(255 > 603))
                                                v3 -= v3
                                                v99 := GenS4{F0: 18446744073709551615, F1: v2, F2: "\n\t"}
                                                var v100 int8 = v91
                                                _ = v100
                                                _ = v98
                                                _ = v99
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    var v101 float64 = ((0.1 / -2.25) + (0.0 * 3.14159))
                                                    _ = v101
                                                case "a":
                                                    var v102 uint = 0
                                                    _ = v102
                                                case "\n\t":
                                                    switch "" {
                                                    case "hello":
                                                        if ((((v89 + float32((0.1 + (-2.25 / (1.5 * (0.0 + 0.1)))))) * 3.14159) > (1e10 * v89)) || (!(v90 > 1))) {
                                                            if false {
                                                                var v103 int16 = 1000
                                                                v104 := GenS11{F0: 0, F1: 1}
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v105 := GenS8{F0: -865, F1: 428, F2: 0}
                                                        _ = v105
                                                    case "foo bar":
                                                        switch "" {
                                                        case "a":
                                                            if v39 {
                                                                switch ("" + "foo bar") {
                                                                case "hello":
                                                                    v106 := GenS8{F0: -1, F1: 249, F2: 0}
                                                                    v97 = v3
                                                                    var v107 uint32 = 2147483647
                                                                    switch "" {
                                                                    case "":
                                                                        v89 = v89
                                                                        var v108 uint = 643
                                                                        _ = v108
                                                                    case "foo bar":
                                                                        v91 = v91
                                                                    case "\n\t":
                                                                        v97 -= genFn0()
                                                                    }
                                                                    _ = v106
                                                                    _ = v107
                                                                case "\n\t":
                                                                    delete(v1, 626)
                                                                default:
                                                                    v109 := []int{fi96, genFn1()}
                                                                    _ = v109
                                                                }
                                                                var v110 bool = (!(1 >= 2147483648))
                                                                v2 = (770 < 237)
                                                                _ = v110
                                                            }
                                                        case "foo bar":
                                                            v3 -= 1
                                                        }
                                                    default:
                                                        v39 = v2
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            if ((fi96 >= (fi96 - (fi96 & (fi96 + (^(fi96)))))) || true) {
                                                var v111 int16 = -1
                                                _ = v111
                                            }
                                        case "":
                                            var v112 uint8 = 4
                                            _ = v112
                                        default:
                                            if (v90 <= v90) {
                                                _ = len(v1)
                                            }
                                            var v113 uint = 42
                                            switch fi96 {
                                            case -902:
                                                v39 = (!true)
                                            case 1:
                                                fmt.Println(v39, v6)
                                            case 9223372036854775807:
                                                v2 = ("" == "\n\t")
                                            }
                                            _ = v113
                                        }
                                        switch genFn1() {
                                        case 127:
                                            var v114 uint64 = 0
                                            _ = v114
                                        case 4294967295:
                                            var v115 bool = v6
                                            _ = v115
                                        case 32768:
                                            if true {
                                                var v116 int = (fi94 << ((1) & 63))
                                                _ = v116
                                            } else {
                                                v95 = (^(v95))
                                            }
                                        }
                                        v117 := GenS12{F0: 253, F1: fi96, F2: 65535}
                                        var v118 uint16 = 820
                                        _ = v117
                                        _ = v118
                                    } else {
                                        var v119 float64 = (3.14159 - -2.25)
                                        var v120 int8 = (^(v91))
                                        if true {
                                            var v121 int8 = int8(v90)
                                            if (v39 || true) {
                                                if (((v3 * (v97 >> ((65535) & 63))) >= v97) && v39) {
                                                    v122 := GenS2{F0: genFn1(), F1: 255}
                                                    var v123 int16 = -102
                                                    v90 -= -568
                                                    _ = v122
                                                    _ = v123
                                                } else {
                                                    switch (fi94 / ((genFn1()) | 1)) {
                                                    case 627:
                                                        var v124 float64 = v119
                                                        _ = v124
                                                    default:
                                                        v125 := GenS13{F0: v121, F1: v39, F2: (fi94 * fi96)}
                                                        _ = v125
                                                    }
                                                    v126 := GenS14{F0: 1}
                                                    _ = v126
                                                }
                                                switch fi94 {
                                                case 760:
                                                    if (!(42 == 161)) {
                                                        v127 := GenS15{F0: 1}
                                                        var v128 uint32 = 691
                                                        var v129 uint64 = 42
                                                        v130 := GenS16{F0: -32768, F1: fi94, F2: 42}
                                                        v4["foo bar"] = (^(v128))
                                                        _ = v127
                                                        _ = v128
                                                        _ = v129
                                                        _ = v130
                                                    }
                                                case 42:
                                                    if (v6 || false) {
                                                        fmt.Println(v3, v91, v39)
                                                        if v39 {
                                                            v90 -= v90
                                                        }
                                                    } else {
                                                        v121 = (-(v120))
                                                    }
                                                }
                                                v131 := GenS17{F0: 1, F1: v120}
                                                v121 += (v91 | -1)
                                                switch "" {
                                                case "":
                                                    if (946 >= v95) {
                                                        var v132 int64 = (v90 >> ((465) & 63))
                                                        v132 += int64(v95)
                                                        v133 := GenS14{F0: 851}
                                                        _ = v132
                                                        _ = v133
                                                    }
                                                case "hello":
                                                    v120 = v121
                                                default:
                                                    var v134 int = (fi96 % ((-445) | 1))
                                                    _ = v134
                                                }
                                                fmt.Println(v91, fi94)
                                                v135 := GenS9{F0: (v3 * (v3 % ((v97) | 1)))}
                                                _ = len(v5)
                                                v136 := GenS18{F0: 4294967295}
                                                var v137 int = (fi94 & fi94)
                                                _ = v131
                                                _ = v135
                                                _ = v136
                                                _ = v137
                                            } else {
                                                var v138 float32 = float32(0.0)
                                                _ = v138
                                            }
                                            _ = v121
                                        } else {
                                            switch (fi96 / ((fi94) | 1)) {
                                            case 586:
                                                v89 = (v89 - ((v89 / (v89 + v89)) + (((3.14159 * (v89 / v89)) / -2.25) + (1.5 + v89))))
                                            case 4294967295:
                                                var v139 bool = (-980 <= -259)
                                                _ = v139
                                            case 32768:
                                                fmt.Println(v89, fi94, fi96)
                                            default:
                                                v120 = -31
                                            }
                                        }
                                        _ = v119
                                        _ = v120
                                    }
                                case 32767:
                                    v3 = v3
                                    v5[(("a" + "\n\t") + ("" + ("" + "")))] = true
                                    var v140 float64 = 1e10
                                    _ = v140
                                default:
                                    fmt.Println(v97, fi96)
                                }
                                var v141 int8 = (v91 >> ((2) & 63))
                                _ = v141
                            case "\n\t":
                                var v142 int = (fi96 | (-(fi94)))
                                _ = v142
                            }
                            if v6 {
                                v143 := GenS19{F0: 120, F1: 2, F2: v3}
                                v5["a"] = (v89 <= 0.0)
                                _ = v143
                            } else {
                                v144 := []int{-9223372036854775808, (-(fi94)), 485, fi96}
                                switch "a" {
                                case "hello":
                                    v3 = (^(v97))
                                case "a":
                                    switch -191 {
                                    case 476:
                                        delete(v4, (("\n\t" + "foo bar") + "hello"))
                                    }
                                case "\n\t":
                                    v89 = v89
                                }
                                _ = v144
                            }
                            if false {
                                if (v6 || true) {
                                    switch (("\n\t" + "\n\t") + "foo bar") {
                                    case "hello":
                                        v145 := []int{425, genFn1(), genFn1(), 2147483647}
                                        _ = v145
                                    }
                                }
                            }
                            v146 := GenS6{F0: 255, F1: 1}
                            v147 := GenS12{F0: 2, F1: fi96, F2: 255}
                            _ = v146
                            _ = v147
                        case 184:
                            fmt.Println(v91, v90)
                        default:
                            if v6 {
                                switch "" {
                                case "":
                                    if false {
                                        v148 := GenS11{F0: 41, F1: 255}
                                        fmt.Println(v91, v97)
                                        fmt.Println(v89, v2)
                                        _ = v148
                                    }
                                default:
                                    v89 = (v89 * (1e10 / v89))
                                }
                                v149 := GenS4{F0: v3, F1: ((v89 + v89) != v89), F2: "hello"}
                                if (976 <= 1) {
                                    var v150 int64 = (v90 << ((679) & 63))
                                    _ = v150
                                } else {
                                    v151 := []int{fi94}
                                    v95 = v95
                                    var v152 uint64 = 1
                                    v153 := GenS3{F0: "a", F1: -2147483648, F2: 558}
                                    v89 = v89
                                    _ = v151
                                    _ = v152
                                    _ = v153
                                }
                                _ = v149
                            } else {
                                var v154 int8 = 100
                                _ = v154
                            }
                        }
                    case -1:
                        v155 := GenS3{F0: (("hello" + ("foo bar" + "foo bar")) + ((("a" + "a") + (("a" + "\n\t") + "")) + "hello")), F1: (v95 - v95), F2: 255}
                        _ = v155
                    }
                    v156 := []int{fi96}
                    _ = fi96
                    _ = v156
                    _ = v97
                }
                _ = fi94
                _ = v95
            }
        }
        var v157 uint32 = 4294967295
        switch "a" {
        case "\n\t":
            v158 := GenS19{F0: 1, F1: (v157 - v157), F2: 0}
            _ = v158
        case "a":
            var v159 int8 = (v91 >> ((18446744073709551615) & 63))
            _ = v159
        default:
            var v160 int8 = (v91 + v91)
            _ = v160
        }
        _ = v157
    }
    var v161 float32 = v89
    for fi162 := 0; fi162 < 18; fi162++ {
        v163 := v1[2147483648]
        _ = fi162
        _ = v163
    }
    var v164 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v161:", v161)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v161
    _ = v164
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v90
    _ = v91
}
