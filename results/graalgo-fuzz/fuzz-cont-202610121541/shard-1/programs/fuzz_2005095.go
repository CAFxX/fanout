package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint32; F2 bool }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 string; F1 int; F2 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 string }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 bool; F1 int32; F2 bool }
type GenS13 struct { F0 uint32; F1 uint32; F2 bool }
func genFn0(p0 int32, p1 bool) bool {
    p1 = (true && p1)
    p1 = true
    _ = p0
    _ = p1
    return ("\n\t" != "")
}

func main() {
    var v1 int32 = 0
    var v2 uint64 = 1
    v3 := make(map[int]string)
    if ((v1 + int32(v2)) != (v1 / ((v1) | 1))) {
        if ((true && (!true)) || (!((v1 + (-(v1))) == 0))) {
            for fi4 := 0; fi4 < 19; fi4++ {
                if false {
                    switch "a" {
                    case "hello":
                        var v5 float32 = ((0.0 - 0.1) + float32(1.5))
                        _ = v5
                    case "a":
                        v6 := []int{fi4, fi4}
                        _ = v6
                    }
                    v7 := []int{}
                    _ = v7
                }
                v2 += v2
                for fi8 := 0; fi8 < 13; fi8++ {
                    v9 := []int{-135, 32767, -9223372036854775808, (fi8 / (((fi8 >> ((826) & 63))) | 1))}
                    fmt.Println(fi8, v1, v1)
                    _ = len(v3)
                    _ = fi8
                    _ = v9
                }
                _ = fi4
            }
            var v10 int64 = int64(v2)
            for fi11 := 0; fi11 < 5; fi11++ {
                v2 += (v2 | v2)
                v12 := GenS1{F0: ""}
                if true {
                    v13 := GenS2{F0: v1, F1: false}
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        var v15 uint32 = 2
                        v16 := GenS1{F0: ""}
                        if ("hello" <= ("" + ("" + "\n\t"))) {
                            var v17 int8 = 127
                            _ = v17
                        } else {
                            var v18 uint = 645
                            var v19 float32 = ((1e10 * 3.14159) * (1.5 / 0.1))
                            _ = v18
                            _ = v19
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    }
    if (-9223372036854775808 <= 42) {
        var v20 uint16 = 886
        var v21 int16 = 828
        _ = v20
        _ = v21
    }
    switch -9223372036854775808 {
    case 0:
        if (-2147483648 != v1) {
            v22 := GenS2{F0: 2147483647, F1: true}
            var v23 uint = 99
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := []int{-808, fi24, 127}
                var v26 float64 = 0.0
                v23 -= (v23 * 508)
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v22
            _ = v23
        }
    case 32768:
        switch "hello" {
        case "\n\t":
            switch -2147483648 {
            case 42:
                if false {
                    switch "a" {
                    case "":
                        v27 := GenS2{F0: v1, F1: ((32767 >= 1) && (0.1 != 1.5))}
                        _ = v27
                    case "\n\t":
                        v28 := GenS1{F0: "\n\t"}
                        _ = v28
                    default:
                        v29 := GenS1{F0: ""}
                        _ = v29
                    }
                    var v30 float32 = (0.1 / -2.25)
                    for fi31 := 0; fi31 < 6; fi31++ {
                        v32 := GenS3{F0: (fi31 & 9223372036854775807), F1: v1}
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
            case 734:
                v1 = v1
            case 646:
                if (!(9223372036854775807 < 127)) {
                    var v33 int64 = -1
                    if ((3.14159 != 3.14159) || false) {
                        v1 = (-(v1))
                    }
                    _ = v33
                }
            }
        case "":
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 int16 = 1000
                var v36 float32 = 0.1
                _ = fi34
                _ = v35
                _ = v36
            }
        case "foo bar":
            if false {
                var v37 uint64 = v2
                if (2 >= 0) {
                    v37 -= v2
                    fmt.Println(v2)
                } else {
                    fmt.Println(v1, v2)
                    v1 -= v1
                    var v38 int64 = -1
                    _ = v38
                }
                v1 = (-(v1))
                _ = v37
            } else {
                delete(v3, 2147483647)
                v39 := GenS4{F0: true, F1: 2147483648, F2: (^(v2))}
                _ = v39
            }
        default:
            var v40 int16 = 1
            _ = v40
        }
    case 32767:
        var v41 bool = (2147483647 != 65535)
        _ = v41
    }
    v42 := GenS3{F0: -508, F1: v1}
    fmt.Println(v1, v1)
    if false {
        v2 += (v2 - v2)
        v43 := []int{-9223372036854775808, 4294967295, 0}
        for fi44 := 0; fi44 < 9; fi44++ {
            if false {
                if true {
                    if (true || true) {
                        fmt.Println(fi44, v2, v2)
                    }
                    v45 := GenS4{F0: (true || false), F1: 284, F2: 2}
                    _ = v45
                }
                v46 := v3[(fi44 ^ -2147483648)]
                v1 = (v1 + v1)
                _ = v46
            }
            if (-124 >= -36) {
                fmt.Println(v2)
                switch (fi44 + -9223372036854775808) {
                case 155:
                    if true {
                        v1 = (v1 - 89)
                        switch fi44 {
                        case -2147483648:
                            v2 -= v2
                        case -9223372036854775808:
                            v47 := v3[fi44]
                            _ = v47
                        }
                        v48 := GenS5{F0: 669, F1: 42, F2: true}
                        _ = v48
                    } else {
                        switch fi44 {
                        case 2147483647:
                            if (127 < 0) {
                                if false {
                                    if (((260 > 1) && (false && (1.5 >= 1.5))) && (1 <= -108)) {
                                        v49 := v3[(fi44 << ((691) & 63))]
                                        var v50 int64 = 1
                                        for fi51 := 0; fi51 < 12; fi51++ {
                                            v3[fi44] = "a"
                                            if false {
                                                var v52 bool = true
                                                var v53 uint32 = 98
                                                if (!v52) {
                                                    fmt.Println(v52, fi51)
                                                    v50 -= (v50 / ((v50) | 1))
                                                    if v52 {
                                                        v54 := GenS6{F0: 127, F1: 100}
                                                        if false {
                                                            switch ("\n\t" + v49) {
                                                            case "":
                                                                v55 := GenS4{F0: (65535 != 1), F1: v53, F2: (v2 << ((2) & 63))}
                                                                _ = v55
                                                            default:
                                                                if false {
                                                                    switch (fi51 >> ((65535) & 63)) {
                                                                    case -9223372036854775808:
                                                                        if len(v43) > 0 {
                                                                            v43[0] = fi51
                                                                        }
                                                                    case -927:
                                                                        switch v49 {
                                                                        case "foo bar":
                                                                            var v56 int = fi51
                                                                            _ = v56
                                                                        case "a":
                                                                            var v57 string = v49
                                                                            _ = v57
                                                                        default:
                                                                            if false {
                                                                                fmt.Println(fi44, v50, v2)
                                                                                if (255 < 0) {
                                                                                    if false {
                                                                                        if false {
                                                                                            var v58 uint = 77
                                                                                            v2 -= (v2 % ((666) | 1))
                                                                                            var v59 int = fi44
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                        } else {
                                                                                            v43 = append(v43, 127)
                                                                                            if (-32768 > -375) {
                                                                                                v53 += (^(v53))
                                                                                                var v60 uint64 = v2
                                                                                                _ = v60
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v43 = append(v43, (fi44 % ((fi51) | 1)))
                                                                                    v61 := GenS3{F0: (-(fi44)), F1: 377}
                                                                                    _ = v61
                                                                                } else {
                                                                                    v49 = "a"
                                                                                }
                                                                                if v52 {
                                                                                    var v62 uint16 = 32767
                                                                                    v50 = v50
                                                                                    _ = v62
                                                                                } else {
                                                                                    var v63 uint16 = 255
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        v2 = v2
                                                                    }
                                                                } else {
                                                                    v3[fi51] = v49
                                                                }
                                                            }
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                if len(v43) > 0 {
                                                                    v43[0] = (fi44 << ((uint(v53)) & 63))
                                                                }
                                                            }
                                                        }
                                                        _ = v54
                                                    }
                                                    v64 := GenS3{F0: fi51, F1: (^(v1))}
                                                    _ = v64
                                                }
                                                v65 := GenS1{F0: ("foo bar" + (v49 + "a"))}
                                                if true {
                                                    v52 = (!false)
                                                    var v66 uint32 = (v53 & v53)
                                                    _ = v66
                                                } else {
                                                    v52 = (1e10 != float64(1e10))
                                                    _ = len(v3)
                                                    v67 := v3[(fi51 << ((402) & 63))]
                                                    _ = v67
                                                }
                                                delete(v3, (fi51 * fi44))
                                                if (209 < 32767) {
                                                    v50 += (v50 % ((v50) | 1))
                                                    if (int64(v1) >= v50) {
                                                        if ((v49 + v49) != "\n\t") {
                                                            v1 = v1
                                                            if (((1.5 * 1e10) > 0.1) || v52) {
                                                                v1 = v1
                                                                var v68 int32 = 2147483647
                                                                _ = v68
                                                            } else {
                                                                switch 355 {
                                                                case -505:
                                                                    switch (fi44 - fi44) {
                                                                    case 127:
                                                                        if (128 == 42) {
                                                                            switch ("\n\t" + (v49 + "hello")) {
                                                                            case "a":
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    fmt.Println(fi51)
                                                                                case "hello":
                                                                                    fmt.Println(fi44, v49, fi51)
                                                                                }
                                                                            case "\n\t":
                                                                                v3[fi51] = v49
                                                                            }
                                                                            var v69 int16 = 433
                                                                            if ((!(v49 != (v49 + v49))) && v52) {
                                                                                if (567 <= 255) {
                                                                                    v53 = (v53 - (v53 - v53))
                                                                                    var v70 uint = 836
                                                                                    _ = v70
                                                                                }
                                                                                switch v49 {
                                                                                case "a":
                                                                                    v43 = append(v43, fi51)
                                                                                case "\n\t":
                                                                                    v1 += (v1 << ((679) & 63))
                                                                                }
                                                                                var v71 int32 = int32(fi51)
                                                                                _ = v71
                                                                            } else {
                                                                                switch fi51 {
                                                                                case 509:
                                                                                    var v72 int8 = int8(v69)
                                                                                    _ = v72
                                                                                case 127:
                                                                                    v73 := GenS3{F0: 4294967295, F1: v1}
                                                                                    _ = v73
                                                                                }
                                                                                v2 = (v2 + (^(v2)))
                                                                            }
                                                                            _ = v69
                                                                        } else {
                                                                            var v74 float64 = (3.14159 * 3.14159)
                                                                            switch v49 {
                                                                            case "hello":
                                                                                switch v49 {
                                                                                case "foo bar":
                                                                                    v75 := GenS1{F0: v49}
                                                                                    _ = v75
                                                                                case "":
                                                                                    v43 = append(v43, fi51)
                                                                                default:
                                                                                    v52 = (v1 >= (v1 * (v1 / (((v1 << ((4294967295) & 63))) | 1))))
                                                                                }
                                                                                v50 = -207
                                                                                var v76 float32 = ((1e10 + (3.14159 - 1.5)) + 1e10)
                                                                                _ = v76
                                                                            case "":
                                                                                v52 = v52
                                                                            }
                                                                            _ = v74
                                                                        }
                                                                    case 1:
                                                                        var v77 int8 = 100
                                                                        _ = v77
                                                                    case -56:
                                                                        v52 = v52
                                                                    }
                                                                case 4294967295:
                                                                    v78 := v3[fi51]
                                                                    _ = v78
                                                                default:
                                                                    v52 = (false && (v49 < ("" + v49)))
                                                                }
                                                                if ((float32(1.5) - -2.25) <= 0.1) {
                                                                    v79 := GenS3{F0: fi51, F1: -417}
                                                                    delete(v3, (fi44 & fi51))
                                                                    _ = v79
                                                                }
                                                                v49 = (((v49 + v49) + v49) + v49)
                                                            }
                                                        }
                                                        switch (^(fi44)) {
                                                        case 32768:
                                                            var v80 int64 = v50
                                                            _ = v80
                                                        case 32767:
                                                            var v81 float32 = float32((3.14159 - -2.25))
                                                            _ = v81
                                                        default:
                                                            _ = len(v43)
                                                        }
                                                        fmt.Println(v2)
                                                        var v82 int64 = (v50 << ((11) & 63))
                                                        v83 := GenS7{F0: v49, F1: (fi44 % ((fi44) | 1)), F2: 704}
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v1 = (v1 << ((65535) & 63))
                                                        v84 := GenS8{F0: (v2 % (((v2 & 18446744073709551615)) | 1))}
                                                        var v85 uint32 = (^(v53))
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    v49 = "hello"
                                                }
                                                _ = v52
                                                _ = v53
                                                _ = v65
                                            }
                                            var v86 int8 = 1
                                            _ = fi51
                                            _ = v86
                                        }
                                        _ = v49
                                        _ = v50
                                    } else {
                                        if true {
                                            switch (-(fi44)) {
                                            case 1:
                                                v87 := GenS5{F0: 42, F1: 4294967295, F2: (!(true || (!false)))}
                                                _ = v87
                                            case 0:
                                                delete(v3, -2147483648)
                                            case -728:
                                                for fi88 := 0; fi88 < 13; fi88++ {
                                                    fmt.Println(fi88)
                                                    _ = len(v43)
                                                    var v89 int32 = 2147483647
                                                    _ = fi88
                                                    _ = v89
                                                }
                                            }
                                        }
                                        var v90 uint16 = 1
                                        for fi91 := 0; fi91 < 20; fi91++ {
                                            if false {
                                                v1 -= (^(v1))
                                                switch (fi91 & 0) {
                                                case 1:
                                                    v92 := GenS9{F0: 2147483648}
                                                    v3[fi44] = "a"
                                                    _ = v92
                                                case 577:
                                                    var v93 bool = true
                                                    var v94 int16 = 42
                                                    var v95 uint = 255
                                                    v96 := GenS9{F0: 2147483648}
                                                    _ = v93
                                                    _ = v94
                                                    _ = v95
                                                    _ = v96
                                                default:
                                                    v90 -= (v90 + 65535)
                                                }
                                            }
                                            _ = fi91
                                        }
                                        _ = v90
                                    }
                                    switch (("a" + "a") + "\n\t") {
                                    case "a":
                                        for fi97 := 0; fi97 < 5; fi97++ {
                                            var v98 uint = 18446744073709551615
                                            v99 := GenS2{F0: v1, F1: false}
                                            var v100 uint16 = 65535
                                            _ = fi97
                                            _ = v100
                                            _ = v98
                                            _ = v99
                                        }
                                    }
                                } else {
                                    v2 += 18446744073709551615
                                    v1 = (v1 >> ((255) & 63))
                                    switch "a" {
                                    case "a":
                                        if false {
                                            for fi101 := 0; fi101 < 10; fi101++ {
                                                v1 -= -344
                                                switch 420 {
                                                case 4294967295:
                                                    var v102 bool = true
                                                    _ = v102
                                                case -32:
                                                    v1 = 0
                                                }
                                                _ = fi101
                                            }
                                        }
                                    case "foo bar":
                                        v1 += int32(v2)
                                    case "\n\t":
                                        v103 := GenS3{F0: -488, F1: 2147483647}
                                        _ = v103
                                    }
                                }
                                for fi104 := 0; fi104 < 13; fi104++ {
                                    if len(v43) > 0 {
                                        v43[0] = (fi104 % ((0) | 1))
                                    }
                                    v1 = (-(v1))
                                    v105 := GenS9{F0: 4294967295}
                                    var v106 int32 = 0
                                    _ = fi104
                                    _ = v105
                                    _ = v106
                                }
                            } else {
                                v107 := GenS3{F0: fi44, F1: v1}
                                if len(v43) >= 2 {
                                    v43 = v43[1:len(v43)]
                                }
                                _ = v107
                            }
                        default:
                            for fi108 := 0; fi108 < 16; fi108++ {
                                if (2 > 2) {
                                    switch ("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) {
                                    case "\n\t":
                                        _ = len(v43)
                                    default:
                                        v109 := GenS8{F0: (v2 | (v2 % ((v2) | 1)))}
                                        _ = v109
                                    }
                                    var v110 float64 = (-2.25 / 1.5)
                                    _ = v110
                                } else {
                                    if true {
                                        v2 -= (v2 + v2)
                                    } else {
                                        fmt.Println(fi44)
                                        var v111 float32 = float32(float64(3.14159))
                                        var v112 float32 = v111
                                        _ = v111
                                        _ = v112
                                    }
                                    var v113 int8 = 42
                                    if true {
                                        var v114 uint32 = 405
                                        _ = v114
                                    }
                                    _ = v113
                                }
                                _ = fi108
                            }
                        }
                    }
                case -635:
                    fmt.Println(fi44, v1)
                case 2147483647:
                    for fi115 := 0; fi115 < 12; fi115++ {
                        v116 := GenS2{F0: v1, F1: true}
                        switch "\n\t" {
                        case "hello":
                            v2 -= v2
                        case "foo bar":
                            var v117 int = (fi115 | (fi44 << ((1) & 63)))
                            _ = v117
                        case "\n\t":
                            if (!true) {
                                if true {
                                    if true {
                                        if (891 < 712) {
                                            var v118 float32 = (float32(-2.25) * (0.0 / -2.25))
                                            _ = v118
                                        }
                                        var v119 uint64 = 42
                                        _ = v119
                                    } else {
                                        _ = len(v3)
                                        v1 += (v1 % ((v1) | 1))
                                        v120 := GenS8{F0: v2}
                                        _ = v120
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v121 int8 = -120
                                        _ = v121
                                    case "hello":
                                        switch fi115 {
                                        case 127:
                                            v1 = (v1 >> ((965) & 63))
                                        case 457:
                                            v43 = append(v43, (-(fi44)))
                                            fmt.Println(v2, v2, fi115)
                                            v1 = v1
                                            var v122 uint = 42
                                            var v123 uint = (v122 << ((v122) & 63))
                                            _ = v122
                                            _ = v123
                                        case -692:
                                            v1 = v1
                                        }
                                        v1 = (v1 + 42)
                                        switch ((("foo bar" + (("hello" + ("" + "")) + "foo bar")) + ("foo bar" + "hello")) + "foo bar") {
                                        case "":
                                            var v124 bool = false
                                            _ = v124
                                        case "\n\t":
                                            switch 42 {
                                            case -852:
                                                v125 := GenS7{F0: "\n\t", F1: fi44, F2: 429}
                                                _ = v125
                                            case 32767:
                                                v2 = v2
                                            case 341:
                                                if (2147483648 == 2147483647) {
                                                    fmt.Println(fi115)
                                                } else {
                                                    _ = len(v3)
                                                    switch fi44 {
                                                    case 32768:
                                                        var v126 uint64 = v2
                                                        _ = v126
                                                    case -2147483648:
                                                        fmt.Println(fi115, v2, v2)
                                                    case 1:
                                                        v2 -= (v2 << ((436) & 63))
                                                    }
                                                }
                                            }
                                        case "hello":
                                            v127 := GenS1{F0: ""}
                                            _ = v127
                                        }
                                    case "\n\t":
                                        var v128 uint = 4294967295
                                        _ = v128
                                    }
                                } else {
                                    if (0 == (v1 - (^(v1)))) {
                                        v129 := v3[fi44]
                                        if (919 <= 255) {
                                            v130 := v3[(fi115 - 609)]
                                            v131 := GenS10{F0: 34, F1: ("foo bar" + ((v130 + v130) + v130))}
                                            _ = v130
                                            _ = v131
                                        } else {
                                            var v132 uint8 = 42
                                            var v133 int8 = 127
                                            fmt.Println(v1, v129)
                                            _ = v132
                                            _ = v133
                                        }
                                        _ = v129
                                    }
                                    var v134 uint = 0
                                    switch fi115 {
                                    case 128:
                                        delete(v3, fi44)
                                    case -9223372036854775808:
                                        v2 = v2
                                    case -2147483648:
                                        if false {
                                            v1 = 1
                                        } else {
                                            _ = len(v43)
                                            var v135 int64 = 9223372036854775807
                                            _ = v135
                                        }
                                    default:
                                        switch "hello" {
                                        case "":
                                            v3[fi44] = ""
                                        default:
                                            switch fi115 {
                                            case 2147483647:
                                                v136 := GenS6{F0: 255, F1: 87}
                                                v1 -= v1
                                                _ = v136
                                            case 525:
                                                if (1.5 == (-2.25 + 0.1)) {
                                                    if (-1 < -1) {
                                                        if false {
                                                            v1 = v1
                                                            if false {
                                                                var v137 bool = (true && (0.0 >= 0.0))
                                                                _ = v137
                                                            }
                                                            v138 := GenS9{F0: 2147483648}
                                                            _ = v138
                                                        } else {
                                                            delete(v3, int(v1))
                                                        }
                                                        v139 := GenS11{F0: (fi115 + -9223372036854775808), F1: (fi44 | (fi115 & 0))}
                                                        if ((((-128 != -128) && (!("hello" < "hello"))) && (19 == 127)) || (-2.25 > 0.1)) {
                                                            v134 += (v134 + 1)
                                                            switch (fi115 | (^(fi44))) {
                                                            case 32768:
                                                                var v140 uint64 = v2
                                                                _ = v140
                                                            default:
                                                                _ = len(v43)
                                                            }
                                                        } else {
                                                            var v141 int64 = 9223372036854775807
                                                            _ = v141
                                                        }
                                                        _ = v139
                                                    }
                                                    if (false || false) {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v142 int8 = -124
                                                            _ = v142
                                                        case "hello":
                                                            var v143 uint = (v134 - 800)
                                                            _ = v143
                                                        default:
                                                            v3[912] = "foo bar"
                                                        }
                                                        v1 -= 0
                                                        v144 := GenS10{F0: 1000, F1: ""}
                                                        _ = v144
                                                    }
                                                    switch "\n\t" {
                                                    case "hello":
                                                        switch 9223372036854775807 {
                                                        case -1:
                                                            v145 := v3[(fi44 % (((fi44 & fi44)) | 1))]
                                                            _ = v145
                                                        }
                                                    case "":
                                                        v43 = append(v43, -290)
                                                    case "\n\t":
                                                        var v146 float64 = 3.14159
                                                        _ = v146
                                                    default:
                                                        var v147 uint32 = 327
                                                        _ = v147
                                                    }
                                                } else {
                                                    v2 = (v2 * v2)
                                                    if true {
                                                        v1 = 42
                                                        v2 += uint64(v134)
                                                        switch "\n\t" {
                                                        case "a":
                                                            _ = len(v3)
                                                        default:
                                                            var v148 uint32 = 465
                                                            _ = v148
                                                        }
                                                    }
                                                }
                                            default:
                                                v1 = v1
                                                _ = len(v3)
                                                var v149 uint32 = 1
                                                _ = v149
                                            }
                                        }
                                    }
                                    _ = v134
                                }
                            } else {
                                if (4294967295 < 32767) {
                                    switch -478 {
                                    case -154:
                                        v150 := GenS12{F0: (65535 < 2147483648), F1: (v1 << ((42) & 63)), F2: false}
                                        _ = v150
                                    case 42:
                                        var v151 uint32 = 1
                                        _ = v151
                                    }
                                    delete(v3, fi115)
                                    fmt.Println(fi115, fi115, fi115)
                                } else {
                                    _ = len(v3)
                                }
                                v2 -= 0
                            }
                        }
                        _ = fi115
                        _ = v116
                    }
                default:
                    if len(v43) > 0 {
                        v43[0] = (^(fi44))
                    }
                }
                v1 += 1
            } else {
                var v152 int64 = 1
                v2 = (v2 / ((709) | 1))
                _ = v152
            }
            var v153 uint = 748
            _ = fi44
            _ = v153
        }
        _ = v43
    } else {
        delete(v3, int(v1))
    }
    v2 = 802
    v154 := GenS1{F0: (((("\n\t" + ("foo bar" + "foo bar")) + ("" + "hello")) + "hello") + "a")}
    switch 2147483647 {
    case 951:
        for fi155 := 0; fi155 < 6; fi155++ {
            var v156 float32 = (float32(((1e10 / 3.14159) * -2.25)) - -2.25)
            for fi157 := 0; fi157 < 19; fi157++ {
                v158 := []int{(fi157 / ((744) | 1)), -9223372036854775808, -1}
                v159 := GenS13{F0: 2147483647, F1: 875, F2: true}
                v1 = (v1 | v1)
                _ = fi157
                _ = v158
                _ = v159
            }
            var v160 uint = 0
            _ = fi155
            _ = v156
            _ = v160
        }
    case 1:
        v3[-2147483648] = "\n\t"
    default:
        v161 := v3[32768]
        _ = v161
    }
    v2 -= 798
    for fi162 := 0; fi162 < 1; fi162++ {
        v163 := GenS3{F0: 32768, F1: (v1 / (((v1 * (v1 ^ (^(v1))))) | 1))}
        var v164 uint64 = 4294967295
        v165 := []int{}
        _ = fi162
        _ = v163
        _ = v164
        _ = v165
    }
    v2 = (v2 << ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v154
    _ = v2
    _ = v3
    _ = v42
}
