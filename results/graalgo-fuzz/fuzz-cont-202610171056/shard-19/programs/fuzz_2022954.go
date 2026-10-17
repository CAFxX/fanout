package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 uint; F2 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64; F1 int64; F2 int16 }
type GenS12 struct { F0 uint8; F1 int32 }
type GenS13 struct { F0 int16; F1 uint16 }
type GenS14 struct { F0 int64; F1 int }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 int64; F1 uint; F2 uint }
type GenS18 struct { F0 string; F1 uint16; F2 string }
type GenS19 struct { F0 int; F1 int32; F2 string }
type GenS20 struct { F0 uint64; F1 uint; F2 bool }
type GenS21 struct { F0 uint64; F1 int8 }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS24 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS25 struct { F0 int; F1 int; F2 uint16 }
func genFn0(p0 bool) string {
    fmt.Println(p0)
    fmt.Println(p0)
    fmt.Println(p0)
    var v7 int8 = 127
    v7 += v7
    _ = p0
    _ = v7
    return ((((("\n\t" + "a") + "hello") + "hello") + "foo bar") + "hello")
}
func genFn1(p0 uint32, p1 int64) int32 {
    for fi8 := 0; fi8 < 1; fi8++ {
        switch "\n\t" {
        case "a":
            switch fi8 {
            case 755:
                p1 -= (p1 | (p1 % ((p1) | 1)))
            case 9223372036854775807:
                var v9 uint16 = 187
                _ = v9
            case -9223372036854775808:
                p0 = p0
            default:
                if (!(255 != 255)) {
                    if false {
                        p0 = p0
                        var v10 float64 = 0.0
                        _ = v10
                    } else {
                        var v11 string = ""
                        var v12 uint = 65535
                        var v13 float64 = (1.5 + (1.5 + float64(1.5)))
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
            }
        case "foo bar":
            p0 -= 2147483648
        default:
            var v14 int32 = -809
            _ = v14
        }
        var v15 uint64 = 1
        _ = fi8
        _ = v15
    }
    fmt.Println(p0, p0)
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint32 = 643
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint64 = 786
            var v20 int16 = -1
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    v3 := []uint32{}
    var v4 uint32 = 42
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    var v21 uint = 255
    v21 = 895
    for fi22 := 0; fi22 < 2; fi22++ {
        fmt.Println(v21)
        v23 := GenS1{F0: -104}
        if len(v3) > 0 {
            v3[0] = v4
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v4, v21)
    fmt.Println(v4)
    for fi24 := 0; fi24 < 1; fi24++ {
        v21 = v21
        fmt.Println(fi24, v21)
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 uint = 65535
            fmt.Println(v4, fi24)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS1{F0: 0}
    var v28 int8 = -1
    fmt.Println(v21, v4, v21)
    for fi29 := 0; fi29 < 19; fi29++ {
        _ = len(v3)
        var v30 string = "a"
        v31 := GenS2{F0: 1}
        _ = fi29
        _ = v30
        _ = v31
    }
    if len(v3) > 0 {
        v3[0] = 561
    }
    fmt.Println(v4)
    v32 := v5["foo bar"]
    if v32 {
        switch -207 {
        case 42:
            v33 := GenS2{F0: 1}
            _ = v33
        default:
            v34 := GenS2{F0: 65535}
            _ = v34
        }
    } else {
        if (3.14159 >= -2.25) {
            var v35 int = 0
            _ = v35
        } else {
            _ = len(v2)
            if v32 {
                var v36 uint8 = 0
                v37 := GenS2{F0: 65535}
                v36 = v36
                _ = v36
                _ = v37
            } else {
                v38 := GenS3{F0: v28, F1: 42}
                if ((-2.25 / 0.1) >= ((1e10 / (1.5 / float32(((0.0 + (1e10 * float64(0.1))) + -2.25)))) * float32(0.1))) {
                    for fi39 := 0; fi39 < 20; fi39++ {
                        if true {
                            var v40 bool = v32
                            _ = v40
                        } else {
                            for fi41 := 0; fi41 < 15; fi41++ {
                                var v42 string = "\n\t"
                                if (!v32) {
                                    v43 := GenS1{F0: -342}
                                    v42 = v42
                                    v3 = append(v3, v4)
                                    _ = v43
                                }
                                var v44 uint32 = (v4 / ((v4) | 1))
                                _ = fi41
                                _ = v42
                                _ = v44
                            }
                        }
                        v45 := GenS4{F0: 2}
                        v32 = v32
                        _ = fi39
                        _ = v45
                    }
                    if v32 {
                        for fi46 := 0; fi46 < 13; fi46++ {
                            if v32 {
                                if v32 {
                                    switch 32767 {
                                    case 32768:
                                        v28 = (v28 << ((v21) & 63))
                                    case 73:
                                        var v47 string = (("" + "foo bar") + ("\n\t" + "\n\t"))
                                        _ = v47
                                    case 4294967295:
                                        var v48 float64 = (1e10 + 1e10)
                                        _ = v48
                                    default:
                                        for fi49 := 0; fi49 < 7; fi49++ {
                                            var v50 uint8 = 128
                                            v51 := v2[(^(v4))]
                                            v52 := GenS1{F0: 9223372036854775807}
                                            if (v21 != 187) {
                                                v50 -= 241
                                            } else {
                                                fmt.Println(v21)
                                                if true {
                                                    var v53 int32 = int32(v4)
                                                    switch fi49 {
                                                    case -411:
                                                        if false {
                                                            var v54 uint32 = v4
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v3 = append(v3, 2147483648)
                                                                v55 := GenS2{F0: 950}
                                                                _ = v55
                                                            }
                                                            v4 -= (v54 * 42)
                                                            _ = v54
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = 2147483648
                                                            }
                                                            var v56 uint64 = 0
                                                            _ = v56
                                                        }
                                                    case 4294967295:
                                                        v57 := v2[(v4 * 2147483647)]
                                                        _ = v57
                                                    }
                                                    if (71 == 2) {
                                                        if (v51 > ((("hello" + "hello") + ((((v51 + "\n\t") + ("" + "a")) + ("" + ("a" + "\n\t"))) + "")) + v51)) {
                                                            fmt.Println(v53)
                                                            var v58 uint16 = 32767
                                                            fmt.Println(v50)
                                                            v59 := GenS2{F0: (v58 * v58)}
                                                            switch v51 {
                                                            case "\n\t":
                                                                var v60 uint8 = v50
                                                                _ = v60
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        if true {
                                                            v61 := GenS1{F0: 1}
                                                            _ = v61
                                                        }
                                                        v62 := GenS5{F0: fi49, F1: (v50 << (((^(v21))) & 63))}
                                                        _ = v62
                                                    } else {
                                                        v63 := GenS1{F0: -9223372036854775808}
                                                        _ = v63
                                                    }
                                                    var v64 uint = uint(v4)
                                                    _ = v53
                                                    _ = v64
                                                }
                                                var v65 int8 = -111
                                                _ = v65
                                            }
                                            _ = fi49
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                }
                                switch "\n\t" {
                                case "":
                                    if v32 {
                                        v66 := GenS6{F0: 0}
                                        v67 := GenS4{F0: 2}
                                        _ = len(v5)
                                        _ = v66
                                        _ = v67
                                    }
                                case "hello":
                                    switch (fi46 + (fi46 | (fi46 / (((fi46 / ((fi46) | 1))) | 1)))) {
                                    case 127:
                                        v4 -= (v4 & (v4 % ((65535) | 1)))
                                    case 2147483647:
                                        switch "\n\t" {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "hello":
                                                v28 = v28
                                            case "foo bar":
                                                v68 := GenS5{F0: fi46, F1: 127}
                                                _ = v68
                                            case "\n\t":
                                                for fi69 := 0; fi69 < 6; fi69++ {
                                                    var v70 int64 = -1
                                                    _ = fi69
                                                    _ = v70
                                                }
                                            default:
                                                v71 := GenS7{F0: -1}
                                                _ = v71
                                            }
                                        case "":
                                            v4 = 304
                                        case "\n\t":
                                            for fi72 := 0; fi72 < 15; fi72++ {
                                                v73 := GenS8{F0: -9223372036854775808}
                                                v2[317] = "\n\t"
                                                _ = fi72
                                                _ = v73
                                            }
                                        }
                                    }
                                case "foo bar":
                                    var v74 uint64 = 18446744073709551615
                                    _ = v74
                                default:
                                    fmt.Println(v28, v21)
                                }
                            } else {
                                fmt.Println(fi46)
                                for fi75 := 0; fi75 < 5; fi75++ {
                                    var v76 int16 = 42
                                    _ = fi75
                                    _ = v76
                                }
                                v4 -= v4
                            }
                            _ = fi46
                        }
                        v77 := GenS9{F0: v4, F1: 255, F2: 0}
                        if false {
                            var v78 uint64 = 105
                            switch 32768 {
                            case -1:
                                v79 := GenS1{F0: -9223372036854775808}
                                _ = v79
                            case -2147483648:
                                switch -686 {
                                case 32768:
                                    v78 -= 2
                                case 32767:
                                    v32 = false
                                case 2147483647:
                                    delete(v5, "a")
                                }
                            }
                            fmt.Println(v4, v78)
                            _ = v78
                        }
                        _ = v77
                    } else {
                        if (!(-2147483648 < 0)) {
                            fmt.Println(v21, v4, v28)
                            v6[(v4 + 832)] = (v4 & v4)
                            switch -718 {
                            case -1:
                                for fi80 := 0; fi80 < 13; fi80++ {
                                    fmt.Println(v28)
                                    if ((v4 / ((1) | 1)) <= 843) {
                                        fmt.Println(v28, v21)
                                        if false {
                                            if v32 {
                                                var v81 uint8 = 128
                                                _ = v81
                                            }
                                            switch (fi80 % ((831) | 1)) {
                                            case 636:
                                                if (((v21 - 4294967295) == (v21 / ((v21) | 1))) && v32) {
                                                    v28 = v28
                                                }
                                            }
                                        } else {
                                            for fi82 := 0; fi82 < 4; fi82++ {
                                                v83 := GenS10{F0: 2}
                                                v21 = (v21 / (((v21 >> (((v21 / (((v21 | v21)) | 1))) & 63))) | 1))
                                                v84 := v5[""]
                                                _ = fi82
                                                _ = v83
                                                _ = v84
                                            }
                                        }
                                        _ = len(v1)
                                    } else {
                                        var v85 float32 = ((1.5 + float32((0.1 * ((-2.25 - 1e10) * ((-2.25 + (((float64(3.14159) * 0.0) * (3.14159 * 3.14159)) * (-2.25 + 1e10))) - 1e10))))) * (1.5 / -2.25))
                                        _ = v85
                                    }
                                    _ = fi80
                                }
                            }
                        } else {
                            v86 := GenS11{F0: -9223372036854775808, F1: 243, F2: 0}
                            _ = v86
                        }
                    }
                    v87 := GenS10{F0: 4294967295}
                    _ = v87
                } else {
                    var v88 uint32 = (v4 + (v4 | v4))
                    _ = v88
                }
                switch 32767 {
                case 1:
                    v5["foo bar"] = v32
                case 4294967295:
                    if (42 == 1) {
                        v89 := GenS12{F0: 2, F1: -1}
                        _ = v89
                    } else {
                        for fi90 := 0; fi90 < 5; fi90++ {
                            fmt.Println(v21)
                            if (1 > 338) {
                                if len(v3) > 0 {
                                    v3[0] = v4
                                }
                                if (int8(v21) > v28) {
                                    v4 = (v4 / (((^(v4))) | 1))
                                } else {
                                    fmt.Println(v21, v4)
                                    switch fi90 {
                                    case 825:
                                        if (v32 && (248 < 2)) {
                                            if (!v32) {
                                                v91 := GenS13{F0: 1, F1: uint16(v21)}
                                                var v92 float64 = 3.14159
                                                for fi93 := 0; fi93 < 4; fi93++ {
                                                    switch (fi93 & (fi93 + (fi93 | 223))) {
                                                    case 127:
                                                        fmt.Println(v92, fi93)
                                                    case 32767:
                                                        if (32767 <= 99) {
                                                            if v32 {
                                                                var v94 uint32 = (v4 & (v4 ^ v4))
                                                                fmt.Println(v21, v21, v94)
                                                                var v95 uint16 = 2
                                                                if (-36 == v28) {
                                                                    v96 := GenS14{F0: -1, F1: (fi90 | (fi93 | fi93))}
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v4
                                                                        }
                                                                    case "":
                                                                        delete(v2, 1)
                                                                    default:
                                                                        v97 := GenS10{F0: 1}
                                                                        v94 += (v94 & 494)
                                                                        _ = v97
                                                                    }
                                                                    v92 = v92
                                                                    if ((!(1 != 18446744073709551615)) || (!v32)) {
                                                                        var v98 int64 = 42
                                                                        v99 := GenS15{F0: 42}
                                                                        _ = v98
                                                                        _ = v99
                                                                    } else {
                                                                        v32 = (42 == 1)
                                                                    }
                                                                    v95 += v95
                                                                    delete(v1, 2147483647)
                                                                    _ = v96
                                                                } else {
                                                                    if (v92 != v92) {
                                                                        var v100 bool = v32
                                                                        var v101 bool = v32
                                                                        var v102 uint8 = 128
                                                                        _ = v100
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                    if v32 {
                                                                        _ = len(v3)
                                                                        var v103 int32 = 0
                                                                        _ = len(v3)
                                                                        _ = v103
                                                                    }
                                                                    delete(v5, "\n\t")
                                                                }
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                        }
                                                        v92 = v92
                                                        if v32 {
                                                            v32 = v32
                                                        }
                                                    default:
                                                        var v104 uint8 = 212
                                                        _ = v104
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = v4
                                                    }
                                                    v6[uint32(v21)] = v4
                                                    _ = fi93
                                                }
                                                _ = v91
                                                _ = v92
                                            }
                                            v105 := GenS16{F0: ("a" + (("" + (("a" + "a") + "foo bar")) + "foo bar")), F1: ("hello" + "")}
                                            _ = v105
                                        }
                                    case 128:
                                        for fi106 := 0; fi106 < 12; fi106++ {
                                            var v107 uint16 = 744
                                            switch (fi90 & fi90) {
                                            case -849:
                                                fmt.Println(v28, v107, v28)
                                            case -141:
                                                fmt.Println(v4, v4)
                                            case -2147483648:
                                                v3 = append(v3, (v4 << ((34) & 63)))
                                                v21 = v21
                                            }
                                            v107 = (v107 & (^(v107)))
                                            fmt.Println(v107, v107)
                                            v21 = (^(v21))
                                            v28 = v28
                                            if false {
                                                v108 := v6[v4]
                                                if (("foo bar" == ("a" + "hello")) || false) {
                                                    v109 := GenS17{F0: -9223372036854775808, F1: v21, F2: v21}
                                                    if ((-(v28)) < v28) {
                                                        v5[("" + ("foo bar" + "hello"))] = (v32 || true)
                                                        var v110 string = ("a" + "hello")
                                                        _ = v110
                                                    }
                                                    var v111 bool = ("hello" != "foo bar")
                                                    v112 := GenS8{F0: (fi106 + (fi90 % (((fi106 / ((fi106) | 1))) | 1)))}
                                                    _ = v109
                                                    _ = v111
                                                    _ = v112
                                                } else {
                                                    v32 = v32
                                                    v113 := GenS7{F0: v28}
                                                    _ = v113
                                                }
                                                if (!v32) {
                                                    v4 = (v4 & 2147483647)
                                                    v32 = (0 == -402)
                                                    v114 := GenS5{F0: (-(fi106)), F1: 46}
                                                    _ = v114
                                                } else {
                                                    var v115 int16 = 0
                                                    _ = v115
                                                }
                                                _ = v108
                                            } else {
                                                var v116 uint32 = v4
                                                var v117 float64 = 0.1
                                                _ = v116
                                                _ = v117
                                            }
                                            _ = fi106
                                            _ = v107
                                        }
                                    case -757:
                                        switch fi90 {
                                        case -2147483648:
                                            v118 := GenS11{F0: 629, F1: 587, F2: 1}
                                            _ = v118
                                        case 4294967295:
                                            v119 := GenS18{F0: "foo bar", F1: 102, F2: ""}
                                            _ = v119
                                        case 530:
                                            var v120 uint64 = 4294967295
                                            _ = v120
                                        }
                                    }
                                    switch fi90 {
                                    case 32767:
                                        v28 = 37
                                    case -1:
                                        switch ("foo bar" + "\n\t") {
                                        case "a":
                                            for fi121 := 0; fi121 < 9; fi121++ {
                                                var v122 uint16 = 255
                                                if (!(((!v32) || (-649 == 1000)) || (!(42 < 0)))) {
                                                    v4 -= 441
                                                    if ((v122 & (v122 % ((65535) | 1))) >= v122) {
                                                        var v123 int8 = v28
                                                        var v124 int32 = 899
                                                        v125 := GenS15{F0: v4}
                                                        if false {
                                                            v122 -= (v122 | (^(v122)))
                                                            var v126 uint16 = 65535
                                                            _ = v126
                                                        } else {
                                                            if (128 > 148) {
                                                                var v127 int16 = 696
                                                                v4 += 2147483648
                                                                var v128 int32 = (v124 % ((v124) | 1))
                                                                v124 = (v128 << (((v21 << ((v21) & 63))) & 63))
                                                                switch fi90 {
                                                                case 571:
                                                                    if true {
                                                                        fmt.Println(v127, v28, fi90)
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            switch "" {
                                                                            case "hello":
                                                                                fmt.Println(v4, v122, v128)
                                                                            }
                                                                        case "":
                                                                            if (true || (-1 == 42)) {
                                                                                var v129 uint8 = 197
                                                                                fmt.Println(v127, v21, v124)
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v130 float32 = 1e10
                                                                                v124 -= (v128 | v128)
                                                                                if v32 {
                                                                                    if v32 {
                                                                                        if v32 {
                                                                                            var v131 uint16 = v122
                                                                                            switch ("foo bar" + (("foo bar" + "hello") + "")) {
                                                                                            case "foo bar":
                                                                                                var v132 string = ""
                                                                                                _ = v132
                                                                                            case "\n\t":
                                                                                                v133 := GenS14{F0: 162, F1: -879}
                                                                                                _ = v133
                                                                                            case "":
                                                                                                fmt.Println(v21)
                                                                                            default:
                                                                                                var v134 int16 = 0
                                                                                                _ = v134
                                                                                            }
                                                                                            if v32 {
                                                                                                if v32 {
                                                                                                    if (v129 > (v129 - v129)) {
                                                                                                        var v135 bool = (!("\n\t" >= "a"))
                                                                                                        v136 := GenS19{F0: -462, F1: v128, F2: "a"}
                                                                                                        if v135 {
                                                                                                            var v137 bool = v32
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = (v4 << ((v21) & 63))
                                                                                                            }
                                                                                                            _ = v137
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                v128 += (v124 >> (((v21 - (^(v21)))) & 63))
                                                                                                            }
                                                                                                            switch (fi121 + (fi90 * (^(fi90)))) {
                                                                                                            case -574:
                                                                                                                var v138 uint8 = v129
                                                                                                                _ = v138
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v135
                                                                                                        _ = v136
                                                                                                    }
                                                                                                    fmt.Println(v122, v127, v21)
                                                                                                    var v139 int8 = (v123 | (v123 * 0))
                                                                                                    _ = v139
                                                                                                }
                                                                                                v3 = append(v3, 2147483647)
                                                                                                v128 = genFn1(v4, int64(fi121))
                                                                                                v140 := v1[(v4 / ((973) | 1))]
                                                                                                var v141 int32 = -2147483648
                                                                                                _ = v140
                                                                                                _ = v141
                                                                                            } else {
                                                                                                if v32 {
                                                                                                    v2[(v4 / ((v4) | 1))] = ""
                                                                                                    var v142 int32 = v124
                                                                                                    fmt.Println(v127, fi121, v122)
                                                                                                    v129 += v129
                                                                                                    var v143 bool = (true || (!(v124 == v142)))
                                                                                                    _ = v142
                                                                                                    _ = v143
                                                                                                }
                                                                                                v124 -= v124
                                                                                                var v144 int = fi121
                                                                                                _ = v144
                                                                                            }
                                                                                            _ = v131
                                                                                        } else {
                                                                                            v145 := GenS20{F0: 86, F1: (v21 ^ (^(v21))), F2: v32}
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            _ = v145
                                                                                        }
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        v128 += (v128 ^ (^(v124)))
                                                                                    }
                                                                                } else {
                                                                                    var v146 int8 = (-(v28))
                                                                                    _ = v146
                                                                                }
                                                                                _ = v129
                                                                                _ = v130
                                                                            }
                                                                        case "\n\t":
                                                                            var v147 uint32 = v4
                                                                            _ = v147
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v4, fi121, fi90)
                                                                    }
                                                                    v124 = (v128 << ((v21) & 63))
                                                                    v124 = int32(v122)
                                                                    var v148 uint64 = 42
                                                                    v149 := GenS5{F0: (^(fi90)), F1: 42}
                                                                    v150 := GenS13{F0: (v127 >> (((v21 / ((v21) | 1))) & 63)), F1: uint16(v148)}
                                                                    v21 -= (^(v21))
                                                                    _ = v148
                                                                    _ = v149
                                                                    _ = v150
                                                                }
                                                                if (true || (v127 == (v127 & 32767))) {
                                                                    v151 := GenS17{F0: 1, F1: (v21 ^ (v21 * (v21 % ((2) | 1)))), F2: 2}
                                                                    var v152 float32 = 3.14159
                                                                    _ = v151
                                                                    _ = v152
                                                                }
                                                                delete(v5, "\n\t")
                                                                switch fi121 {
                                                                case 32767:
                                                                    delete(v6, 20)
                                                                case -971:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 379
                                                                    }
                                                                case 42:
                                                                    v28 = (^(v28))
                                                                }
                                                                _ = v127
                                                                _ = v128
                                                            } else {
                                                                v153 := v2[v4]
                                                                if (fi121 > -1) {
                                                                    var v154 int32 = v124
                                                                    v3 = append(v3, v4)
                                                                    _ = v154
                                                                } else {
                                                                    v155 := GenS21{F0: 4294967295, F1: (-(v28))}
                                                                    v156 := GenS22{F0: (v21 | (v21 ^ v21))}
                                                                    _ = v155
                                                                    _ = v156
                                                                }
                                                                fmt.Println(fi121, fi90)
                                                                _ = v153
                                                            }
                                                        }
                                                        _ = v123
                                                        _ = v124
                                                        _ = v125
                                                    } else {
                                                        delete(v2, (v4 >> ((v21) & 63)))
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if (!v32) {
                                                                switch "a" {
                                                                case "":
                                                                    fmt.Println(fi90, v28)
                                                                case "a":
                                                                    var v157 int16 = 0
                                                                    var v158 string = "a"
                                                                    v159 := GenS19{F0: 509, F1: -2147483648, F2: v158}
                                                                    _ = v157
                                                                    _ = v158
                                                                    _ = v159
                                                                }
                                                            } else {
                                                                var v160 uint64 = 4294967295
                                                                fmt.Println(v21, v160)
                                                                _ = v160
                                                            }
                                                        case "\n\t":
                                                            v3 = append(v3, v4)
                                                        default:
                                                            var v161 int32 = 1
                                                            _ = v161
                                                        }
                                                        v162 := GenS15{F0: (v4 | v4)}
                                                        var v163 uint16 = (v122 >> ((v21) & 63))
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                    v28 = v28
                                                }
                                                _ = fi121
                                                _ = v122
                                            }
                                        }
                                    }
                                }
                            }
                            _ = fi90
                        }
                        v3 = append(v3, 4294967295)
                    }
                case 42:
                    if v32 {
                        for fi164 := 0; fi164 < 20; fi164++ {
                            v4 = (v4 ^ 549)
                            _ = fi164
                        }
                    } else {
                        var v165 uint32 = (v4 * v4)
                        _ = v165
                    }
                default:
                    switch (("foo bar" + ("" + "hello")) + "hello") {
                    case "a":
                        v28 = v28
                    default:
                        for fi166 := 0; fi166 < 16; fi166++ {
                            delete(v6, v4)
                            v28 = (v28 * v28)
                            v4 = (^(v4))
                            _ = fi166
                        }
                    }
                }
                _ = v38
            }
        }
        fmt.Println(v21)
        if (!v32) {
            for fi167 := 0; fi167 < 11; fi167++ {
                for fi168 := 0; fi168 < 9; fi168++ {
                    v169 := GenS23{F0: 412, F1: 55, F2: 0}
                    v21 -= (v21 << (((v21 + v21)) & 63))
                    var v170 float64 = (1.5 * float64(1e10))
                    _ = fi168
                    _ = v169
                    _ = v170
                }
                fmt.Println(v4, v4, v4)
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                _ = fi167
            }
        } else {
            for fi171 := 0; fi171 < 20; fi171++ {
                var v172 int8 = (v28 << (((v21 << (((v21 ^ v21)) & 63))) & 63))
                if false {
                    if v32 {
                        v21 = v21
                        v3 = append(v3, v4)
                        switch (fi171 / ((fi171) | 1)) {
                        case 127:
                            v6[(v4 * 477)] = 0
                        case 32767:
                            v21 = (v21 & v21)
                        }
                    } else {
                        if v32 {
                            if v32 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v21 += v21
                            } else {
                                v28 -= (v172 >> (((v21 * (^(v21)))) & 63))
                            }
                            var v173 bool = (v32 && (((0.0 - -2.25) / 1.0) == 0.1))
                            var v174 uint32 = (v4 + (^(v4)))
                            _ = v173
                            _ = v174
                        } else {
                            v175 := GenS24{F0: (v4 * (v4 / ((v4) | 1))), F1: 1, F2: v4}
                            _ = v175
                        }
                    }
                }
                _ = fi171
                _ = v172
            }
        }
    }
    switch "foo bar" {
    case "hello":
        v176 := GenS25{F0: 9223372036854775807, F1: 0, F2: 255}
        _ = v176
    case "\n\t":
        v4 = 65535
    default:
        v21 = v21
    }
    for fi177 := 0; fi177 < 10; fi177++ {
        for fi178 := 0; fi178 < 11; fi178++ {
            v28 += -128
            _ = fi178
        }
        _ = fi177
    }
    var v179 float32 = (0.1 * (3.14159 - 0.1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v179:", v179)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v179
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
