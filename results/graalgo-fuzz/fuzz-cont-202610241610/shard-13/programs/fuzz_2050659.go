package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32; F1 int8; F2 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 int16; F1 int16; F2 bool }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint8; F1 string; F2 int64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64; F1 uint32 }
type GenS14 struct { F0 int; F1 uint64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int; F1 int64 }
type GenS17 struct { F0 bool; F1 uint64 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 uint }
func genFn0(p0 uint64, p1 uint32) uint64 {
    var v6 int8 = 42
    if ((!true) && (1 < 122)) {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 bool = ((255 <= 255) && true)
            p0 = (p0 - p0)
            if (!(1000 >= 42)) {
                switch (("" + "hello") + "hello") {
                case "foo bar":
                    var v9 int64 = 0
                    _ = v9
                case "a":
                    v10 := []int{(fi7 / ((fi7) | 1)), (fi7 ^ 639), -9223372036854775808, -2147483648}
                    _ = v10
                }
            } else {
                switch "" {
                case "":
                    v6 = (v6 & 101)
                case "a":
                    var v11 float32 = 0.0
                    _ = v11
                case "foo bar":
                    for fi12 := 0; fi12 < 12; fi12++ {
                        var v13 string = ("a" + "")
                        if true {
                            if ((v13 >= v13) && v8) {
                                v14 := GenS1{F0: 860, F1: p1, F2: 0}
                                v15 := GenS2{F0: 572, F1: (v6 << ((2) & 63)), F2: fi12}
                                _ = v14
                                _ = v15
                            }
                            switch fi7 {
                            case 128:
                                v6 = (v6 / ((v6) | 1))
                            case -713:
                                v16 := []int{(fi7 * 42), fi7, fi12}
                                _ = v16
                            case 9223372036854775807:
                                v17 := []int{fi7, (fi7 << ((18446744073709551615) & 63))}
                                _ = v17
                            default:
                                switch fi12 {
                                case 266:
                                    var v18 int32 = 561
                                    _ = v18
                                case 42:
                                    v19 := GenS2{F0: (p1 << ((4294967295) & 63)), F1: v6, F2: int(p0)}
                                    _ = v19
                                default:
                                    if (!true) {
                                        v20 := []int{47}
                                        var v21 float64 = float64((float32(((1e10 / (1.5 + (1.5 / 0.1))) + float64(1.5))) - -2.25))
                                        fmt.Println(p0)
                                        _ = v20
                                        _ = v21
                                    } else {
                                        v22 := []int{(fi7 ^ -786), fi7, (fi12 * (fi7 % ((fi7) | 1))), 32767}
                                        var v23 uint32 = (p1 - p1)
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            }
                        }
                        var v24 uint16 = 1
                        _ = fi12
                        _ = v13
                        _ = v24
                    }
                default:
                    switch "" {
                    case "a":
                        switch (fi7 * (fi7 % (((^(fi7))) | 1))) {
                        case -342:
                            var v25 uint32 = uint32(p0)
                            _ = v25
                        case -1:
                            for fi26 := 0; fi26 < 9; fi26++ {
                                var v27 uint16 = 65535
                                v28 := GenS3{F0: (p1 - p1), F1: -81}
                                _ = fi26
                                _ = v27
                                _ = v28
                            }
                        }
                    default:
                        if v8 {
                            fmt.Println(p1, v8, p0)
                            for fi29 := 0; fi29 < 1; fi29++ {
                                var v30 string = "foo bar"
                                v31 := GenS4{F0: 569, F1: 2, F2: (p1 / ((4294967295) | 1))}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            v32 := []int{}
                            _ = v32
                        } else {
                            v6 = v6
                            switch ("hello" + ("" + "a")) {
                            case "\n\t":
                                for fi33 := 0; fi33 < 12; fi33++ {
                                    var v34 int8 = (v6 | v6)
                                    var v35 float32 = 1e10
                                    _ = fi33
                                    _ = v34
                                    _ = v35
                                }
                            case "foo bar":
                                fmt.Println(p0, fi7, v8)
                            case "a":
                                if v8 {
                                    v8 = (544 >= 65535)
                                    p1 = (p1 & (^(p1)))
                                    v6 = (v6 | (^(v6)))
                                } else {
                                    fmt.Println(p0)
                                    var v36 int16 = -32768
                                    _ = v36
                                }
                            }
                            fmt.Println(p1)
                        }
                    }
                }
            }
            _ = fi7
            _ = v8
        }
    }
    p0 += (p0 % ((p0) | 1))
    v6 -= (v6 >> ((0) & 63))
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int16 = -1
    var v3 int = -2147483648
    var v4 int8 = -1
    v5 := []int{(v3 - 1), 9223372036854775807, (v3 % ((4294967295) | 1))}
    var v37 float32 = 0.1
    if false {
        v4 += v4
        for fi38 := 0; fi38 < 12; fi38++ {
            v39 := GenS5{F0: -555, F1: fi38}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi38
            _ = v39
        }
        v40 := GenS6{F0: "", F1: 43}
        _ = v40
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        fmt.Println(v3, v4, v4)
        _ = fi41
    }
    v37 += float32((0.0 - 0.1))
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            v43 := GenS6{F0: "foo bar", F1: 127}
            _ = fi42
            _ = v43
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi44 := 0; fi44 < 15; fi44++ {
        v4 = -93
        _ = fi44
    }
    var v45 uint8 = 0
    var v46 uint64 = v1
    switch ("" + "hello") {
    case "a":
        switch 128 {
        case 1:
            for fi47 := 0; fi47 < 17; fi47++ {
                _ = len(v5)
                _ = fi47
            }
        default:
            if true {
                if ((v45 << ((0) & 63)) > (v45 ^ v45)) {
                    var v48 uint = 42
                    _ = v48
                }
                v45 = v45
                v49 := GenS2{F0: 2147483647, F1: v4, F2: -1}
                _ = v49
            }
        }
    case "hello":
        v50 := GenS7{F0: 32767, F1: -32768, F2: (!((0.1 + (v37 * float32((((1.5 * 1.5) + 1.5) / float64(float32(1e10)))))) < v37))}
        _ = v50
    case "":
        v46 = v1
    default:
        v2 = -932
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ("hello" > "") {
        switch (v3 << ((449) & 63)) {
        case -668:
            switch "" {
            case "\n\t":
                for fi51 := 0; fi51 < 2; fi51++ {
                    if (true && (v37 == v37)) {
                        if (!(float64((v37 - v37)) > (1e10 / 3.14159))) {
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            default:
                                v45 -= 2
                            }
                            var v52 string = ("a" + "\n\t")
                            if ((715 != 0) || true) {
                                switch v52 {
                                case "foo bar":
                                    for fi53 := 0; fi53 < 16; fi53++ {
                                        v54 := GenS6{F0: v52, F1: (^(v45))}
                                        var v55 int8 = (-(v4))
                                        if false {
                                            switch (fi51 % (((fi53 * (v3 - 42))) | 1)) {
                                            case -815:
                                                var v56 int8 = -1
                                                _ = v56
                                            case 42:
                                                var v57 int8 = int8(v2)
                                                _ = v57
                                            default:
                                                var v58 string = "foo bar"
                                                _ = v58
                                            }
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v5 = append(v5, (-(fi53)))
                                            var v59 float32 = 0.0
                                            _ = v59
                                        }
                                        v60 := GenS2{F0: 590, F1: int8(v3), F2: (fi51 >> ((21) & 63))}
                                        switch v52 {
                                        case "\n\t":
                                            v46 += v46
                                        default:
                                            v1 = 735
                                        }
                                        _ = fi53
                                        _ = v54
                                        _ = v55
                                        _ = v60
                                    }
                                case "\n\t":
                                    if (!false) {
                                        for fi61 := 0; fi61 < 1; fi61++ {
                                            v62 := GenS7{F0: v2, F1: (v2 - (v2 - v2)), F2: true}
                                            _ = fi61
                                            _ = v62
                                        }
                                        var v63 uint = 4294967295
                                        _ = v63
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v64 uint = 2
                                        _ = v64
                                    }
                                case "hello":
                                    for fi65 := 0; fi65 < 1; fi65++ {
                                        var v66 string = v52
                                        if true {
                                            if ((245 < 0) || ((v2 * (v2 / ((v2) | 1))) == -32768)) {
                                                var v67 int = (fi65 >> ((898) & 63))
                                                _ = v67
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = 4294967295
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if (false && false) {
                                                    var v68 bool = (((v46 / ((v46) | 1)) != (v46 | (v46 | (^(v46))))) || (131 < 2))
                                                    var v69 int64 = 1
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    var v70 int16 = 579
                                                    v71 := GenS6{F0: v66, F1: v45}
                                                    v72 := GenS8{F0: (v2 ^ 1), F1: 276}
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi51 - (^(fi51)))
                                                    }
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                }
                                            }
                                            fmt.Println(fi51, v52)
                                            switch "" {
                                            case "foo bar":
                                                v73 := GenS9{F0: 255, F1: (v66 + "hello"), F2: 702}
                                                _ = v73
                                            default:
                                                switch ((v66 + v66) + ("\n\t" + v66)) {
                                                case "\n\t":
                                                    v66 = v52
                                                }
                                            }
                                            fmt.Println(v4)
                                            v37 = (-2.25 / v37)
                                        }
                                        v4 -= v4
                                        var v74 int64 = 0
                                        _ = fi65
                                        _ = v66
                                        _ = v74
                                    }
                                default:
                                    fmt.Println(v3, v52)
                                }
                                var v75 string = ("" + ("\n\t" + "a"))
                                var v76 int16 = v2
                                _ = v75
                                _ = v76
                            } else {
                                v52 = v52
                                v1 = v1
                            }
                            _ = v52
                        }
                    } else {
                        var v77 int8 = 127
                        if (false && false) {
                            var v78 bool = true
                            if false {
                                switch "foo bar" {
                                case "\n\t":
                                    v2 = (v2 / ((-855) | 1))
                                case "foo bar":
                                    for fi79 := 0; fi79 < 5; fi79++ {
                                        var v80 uint16 = 42
                                        v4 -= v4
                                        _ = fi79
                                        _ = v80
                                    }
                                default:
                                    switch (fi51 * fi51) {
                                    case -731:
                                        if (1 == 42) {
                                            v46 += (v46 >> ((18446744073709551615) & 63))
                                            var v81 uint64 = (v46 % (((v46 & v46)) | 1))
                                            v82 := GenS7{F0: (v2 - 701), F1: -32768, F2: v78}
                                            _ = v81
                                            _ = v82
                                        } else {
                                            fmt.Println(v46, v1, v77)
                                            v45 -= 165
                                        }
                                    case -660:
                                        if len(v5) > 0 {
                                            v5[0] = v3
                                        }
                                    }
                                }
                                v1 += (v46 << ((1) & 63))
                                var v83 string = ("a" + "foo bar")
                                _ = v83
                            }
                            v1 = (v1 - (^(v1)))
                            _ = v78
                        } else {
                            var v84 uint64 = 1
                            _ = v84
                        }
                        var v85 string = ""
                        _ = v77
                        _ = v85
                    }
                    if true {
                        var v86 string = "a"
                        switch (v86 + "\n\t") {
                        case "hello":
                            var v87 int32 = 1
                            _ = v87
                        case "a":
                            for fi88 := 0; fi88 < 12; fi88++ {
                                if false {
                                    v89 := GenS2{F0: 663, F1: v4, F2: 32767}
                                    v86 = v86
                                    switch v86 {
                                    case "foo bar":
                                        if (v86 == v86) {
                                            if true {
                                                v3 = fi51
                                                if true {
                                                    var v90 bool = (!(false || (!(v37 < 3.14159))))
                                                    var v91 float64 = 3.14159
                                                    _ = v90
                                                    _ = v91
                                                }
                                                v92 := GenS10{F0: (v4 - 42)}
                                                var v93 float32 = 0.1
                                                var v94 uint16 = 42
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            }
                                            var v95 uint64 = v1
                                            if true {
                                                switch v3 {
                                                case -182:
                                                    var v96 uint64 = v95
                                                    _ = v96
                                                case 2147483647:
                                                    v5 = append(v5, (-(fi51)))
                                                case -865:
                                                    fmt.Println(v95)
                                                default:
                                                    v37 -= (0.1 * v37)
                                                }
                                                if (v45 < 1) {
                                                    v97 := GenS11{F0: ("\n\t" + (v86 + v86))}
                                                    v98 := GenS11{F0: "\n\t"}
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    var v99 string = v86
                                                    _ = v99
                                                }
                                                v37 += v37
                                            } else {
                                                fmt.Println(v2)
                                                v46 = genFn0(v46, 42)
                                                v4 = int8(v2)
                                            }
                                            _ = v95
                                        } else {
                                            v46 = genFn0(genFn0(v46, 70), 65535)
                                        }
                                    }
                                    _ = v89
                                } else {
                                    v45 = 128
                                    if ((v3 | (fi88 << ((706) & 63))) > v3) {
                                        if false {
                                            var v100 bool = true
                                            _ = len(v5)
                                            _ = v100
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = (-(fi88))
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            switch (v86 + ("\n\t" + ((v86 + ("\n\t" + v86)) + (((v86 + "foo bar") + "hello") + ((v86 + "hello") + ""))))) {
                                            case "foo bar":
                                                fmt.Println(v46)
                                            default:
                                                if ((fi51 < fi51) && (!(!true))) {
                                                    _ = len(v5)
                                                } else {
                                                    if false {
                                                        v101 := GenS12{F0: 0}
                                                        v102 := GenS13{F0: -1, F1: 2147483647}
                                                        switch (fi88 + v3) {
                                                        case 32768:
                                                            v45 -= v45
                                                        default:
                                                            if (2147483647 <= 179) {
                                                                var v103 bool = true
                                                                _ = v103
                                                            }
                                                        }
                                                        fmt.Println(v4, v46)
                                                        _ = v101
                                                        _ = v102
                                                    } else {
                                                        v1 = (v46 / (((v46 | 2)) | 1))
                                                        var v104 int32 = 1
                                                        _ = v104
                                                    }
                                                    switch (v3 << ((65535) & 63)) {
                                                    case 128:
                                                        if (((v1 & v1) > (^(v46))) && false) {
                                                            switch (^(fi51)) {
                                                            case 128:
                                                                var v105 string = ""
                                                                _ = v105
                                                            default:
                                                                fmt.Println(v1, fi88)
                                                            }
                                                            if ((-2.25 / 1.0) == 1e10) {
                                                                var v106 int16 = v2
                                                                if (65535 == 658) {
                                                                    if false {
                                                                        var v107 uint32 = 2147483648
                                                                        v106 -= (v106 << ((0) & 63))
                                                                        v1 -= v1
                                                                        v106 -= (v106 % ((654) | 1))
                                                                        var v108 uint = 42
                                                                        var v109 int = (fi51 * (fi88 | v3))
                                                                        _ = len(v5)
                                                                        _ = v107
                                                                        _ = v108
                                                                        _ = v109
                                                                    } else {
                                                                        if false {
                                                                            var v110 uint32 = 65535
                                                                            _ = len(v5)
                                                                            _ = v110
                                                                        }
                                                                        v106 += (v106 / ((v106) | 1))
                                                                    }
                                                                }
                                                                _ = v106
                                                            } else {
                                                                var v111 int64 = -788
                                                                if len(v5) > 0 {
                                                                    v5[0] = (-(v3))
                                                                }
                                                                _ = v111
                                                            }
                                                            v112 := GenS12{F0: 1}
                                                            _ = v112
                                                        }
                                                    case 4294967295:
                                                        var v113 uint32 = 65535
                                                        _ = v113
                                                    default:
                                                        v37 = v37
                                                    }
                                                    v4 = (v4 | (v4 | v4))
                                                }
                                            }
                                        }
                                        v114 := GenS6{F0: ("a" + v86), F1: (^(v45))}
                                        v4 -= (v4 >> ((0) & 63))
                                        _ = v114
                                    }
                                    _ = len(v5)
                                    v3 = (fi51 * (fi51 + int(v45)))
                                }
                                _ = len(v5)
                                v115 := GenS13{F0: 1, F1: 65535}
                                _ = fi88
                                _ = v115
                            }
                        case "":
                            var v116 int64 = 0
                            _ = v116
                        default:
                            if (!(-2.25 < -2.25)) {
                                v46 -= (^(v1))
                            }
                        }
                        v1 -= (^(v1))
                        _ = v86
                    }
                    v5 = append(v5, (fi51 >> ((201) & 63)))
                    _ = fi51
                }
            case "":
                v4 = v4
            case "hello":
                fmt.Println(v3, v1)
            default:
                for fi117 := 0; fi117 < 20; fi117++ {
                    if (32768 < v3) {
                        fmt.Println(v37, v1)
                        switch ("a" + (("a" + "") + ((("\n\t" + "") + "a") + ("foo bar" + ("" + "a"))))) {
                        case "":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case "hello":
                            if false {
                                v118 := GenS5{F0: -1, F1: 331}
                                for fi119 := 0; fi119 < 11; fi119++ {
                                    v120 := GenS14{F0: 2147483647, F1: 1}
                                    v3 = fi117
                                    _ = fi119
                                    _ = v120
                                }
                                v121 := GenS11{F0: "hello"}
                                _ = v118
                                _ = v121
                            }
                        default:
                            v122 := GenS15{F0: 0}
                            _ = v122
                        }
                    }
                    var v123 bool = (!false)
                    _ = fi117
                    _ = v123
                }
            }
        default:
            v124 := GenS8{F0: (v2 << ((1) & 63)), F1: 65535}
            _ = v124
        }
        v125 := GenS16{F0: v3, F1: 42}
        _ = v125
    }
    _ = len(v5)
    v2 += int16(v4)
    switch (v3 % ((970) | 1)) {
    case 454:
        fmt.Println(v46)
    case 0:
        if (!false) {
            fmt.Println(v3, v46, v46)
            switch ("\n\t" + "foo bar") {
            case "a":
                v126 := GenS17{F0: (42 <= -785), F1: v46}
                _ = v126
            case "\n\t":
                v127 := GenS18{F0: int64(v4)}
                _ = v127
            case "hello":
                v45 -= 255
            }
        } else {
            switch (v3 * v3) {
            case 42:
                for fi128 := 0; fi128 < 10; fi128++ {
                    v45 -= v45
                    _ = fi128
                }
            }
        }
    default:
        v129 := GenS19{F0: 429}
        _ = v129
    }
    if false {
        for fi130 := 0; fi130 < 16; fi130++ {
            var v131 int32 = 42
            v132 := GenS18{F0: 438}
            var v133 uint = 65535
            _ = fi130
            _ = v131
            _ = v132
            _ = v133
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
