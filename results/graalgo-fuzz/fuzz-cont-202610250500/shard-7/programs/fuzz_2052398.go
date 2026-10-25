package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int8; F2 string }
type GenS5 struct { F0 int64; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 string }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 int16; F1 uint16; F2 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int; F1 uint16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int64; F1 int32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 string; F1 int }
type GenS16 struct { F0 uint32 }

func main() {
    var v1 uint32 = 2
    var v2 uint32 = v1
    v3 := make(map[string]uint32)
    var v4 bool = (595 < 32768)
    switch ("foo bar" + (("" + "hello") + "a")) {
    case "\n\t":
        if ((3.14159 * 1e10) != float64(((0.0 - (float32(0.1) * -2.25)) - 1e10))) {
            for fi5 := 0; fi5 < 18; fi5++ {
                fmt.Println(fi5, v2)
                var v6 uint8 = 255
                _ = fi5
                _ = v6
            }
        }
    case "hello":
        for fi7 := 0; fi7 < 1; fi7++ {
            if false {
                if false {
                    v8 := []int{fi7, (-(fi7)), (fi7 << ((857) & 63))}
                    var v9 uint = 114
                    _ = v8
                    _ = v9
                } else {
                    v10 := []int{(fi7 & (fi7 ^ (fi7 + fi7))), (fi7 << ((255) & 63)), 4294967295}
                    _ = v10
                }
                delete(v3, "")
            }
            var v11 int = fi7
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint8 = 255
                _ = fi12
                _ = v13
            }
            _ = fi7
            _ = v11
        }
    case "":
        var v14 uint = 4294967295
        _ = v14
    }
    fmt.Println(v4, v4, v2)
    if (116 > 205) {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 float64 = (0.1 + float64(float32(1.5)))
            v17 := GenS1{F0: 323, F1: 339, F2: 4294967295}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int8 = 100
        if true {
            switch 32768 {
            case 32768:
                v19 := []int{4294967295}
                _ = v19
            case 937:
                var v20 int64 = -754
                _ = v20
            case 42:
                for fi21 := 0; fi21 < 12; fi21++ {
                    if (!(true && (128 != 18))) {
                        v22 := GenS1{F0: v1, F1: 0, F2: 794}
                        _ = v22
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v23 := []int{fi21, fi21}
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 18; fi24++ {
                            fmt.Println(fi24)
                            v25 := GenS2{F0: (v1 - v2), F1: (v18 << ((804) & 63))}
                            _ = fi24
                            _ = v25
                        }
                        if v4 {
                            var v26 int = 9223372036854775807
                            v18 = (v18 & v18)
                            _ = v26
                        } else {
                            if true {
                                if v4 {
                                    v18 = (v18 << ((89) & 63))
                                    var v27 uint = 578
                                    v28 := GenS2{F0: (v2 % (((v2 - (^(v2)))) | 1)), F1: 1}
                                    _ = v27
                                    _ = v28
                                }
                            } else {
                                v29 := []int{128, 0, (fi21 ^ fi21)}
                                delete(v3, "foo bar")
                                _ = v29
                            }
                            switch ("" + "") {
                            case "":
                                v30 := GenS2{F0: (^(v2)), F1: (^(v18))}
                                _ = v30
                            case "a":
                                v31 := GenS2{F0: 2147483647, F1: (v18 ^ (-(v18)))}
                                _ = v31
                            }
                            switch fi21 {
                            case 32768:
                                v4 = (189 < 1)
                            case 2147483647:
                                var v32 uint16 = 0
                                _ = v32
                            }
                        }
                    }
                    _ = fi21
                }
            default:
                v33 := GenS2{F0: (v1 | (^(v1))), F1: (v18 ^ v18)}
                _ = v33
            }
            v34 := GenS2{F0: 1, F1: 64}
            var v35 bool = v4
            _ = v34
            _ = v35
        } else {
            for fi36 := 0; fi36 < 10; fi36++ {
                if ((-525 == -728) && (!v4)) {
                    v37 := GenS3{F0: 753}
                    _ = v37
                } else {
                    var v38 int16 = -1
                    _ = v38
                }
                v3[("foo bar" + (("a" + (("" + "a") + ("hello" + "a"))) + "\n\t"))] = (v2 % ((4294967295) | 1))
                var v39 uint32 = v2
                _ = fi36
                _ = v39
            }
            var v40 float64 = 0.0
            _ = v40
        }
        _ = v18
    }
    v1 = 4294967295
    v41 := GenS3{F0: 17}
    var v42 uint8 = 0
    v4 = v4
    if v4 {
        switch (("\n\t" + ("foo bar" + (("\n\t" + "foo bar") + "foo bar"))) + "hello") {
        case "hello":
            for fi43 := 0; fi43 < 5; fi43++ {
                var v44 uint8 = v42
                if v4 {
                    v4 = v4
                    v45 := GenS1{F0: 0, F1: 0, F2: 2}
                    _ = v45
                } else {
                    switch "\n\t" {
                    case "":
                        if v4 {
                            if ("hello" > "foo bar") {
                                v4 = false
                            } else {
                                if v4 {
                                    var v46 uint8 = 46
                                    for fi47 := 0; fi47 < 8; fi47++ {
                                        v48 := GenS4{F0: 2147483647, F1: -94, F2: ""}
                                        _ = fi47
                                        _ = v48
                                    }
                                    var v49 uint16 = 1
                                    _ = v46
                                    _ = v49
                                }
                                var v50 int32 = -1
                                _ = v50
                            }
                            if v4 {
                                fmt.Println(v1, v1)
                                if ((v2 ^ (v1 + v1)) < (v1 % ((uint32(v42)) | 1))) {
                                    if ((("foo bar" + "hello") + "a") != "hello") {
                                        if (v4 && false) {
                                            _ = len(v3)
                                            var v51 int32 = -331
                                            _ = v51
                                        } else {
                                            if (("" + "hello") <= "") {
                                                v42 = (v42 & v44)
                                            } else {
                                                v1 -= (v1 * 65)
                                                for fi52 := 0; fi52 < 4; fi52++ {
                                                    v1 += (v2 + (v1 ^ (v1 % (((v2 | v1)) | 1))))
                                                    _ = fi52
                                                }
                                                v53 := GenS5{F0: 9223372036854775807, F1: -1, F2: 42}
                                                _ = v53
                                            }
                                            var v54 int32 = 42
                                            _ = v54
                                        }
                                    } else {
                                        fmt.Println(v44, v2)
                                        var v55 string = "foo bar"
                                        var v56 uint = 18446744073709551615
                                        _ = v55
                                        _ = v56
                                    }
                                }
                            } else {
                                for fi57 := 0; fi57 < 14; fi57++ {
                                    v4 = (v4 && (32767 != 32767))
                                    v58 := GenS3{F0: 18446744073709551615}
                                    var v59 float32 = ((1e10 / 1.0) - 0.0)
                                    var v60 float64 = 0.1
                                    switch 2147483647 {
                                    case -2147483648:
                                        v61 := GenS6{F0: v1, F1: 2147483647, F2: v42}
                                        _ = v61
                                    }
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                _ = len(v3)
                            }
                        } else {
                            if (false && (31 < 42)) {
                                var v62 uint64 = 18446744073709551615
                                var v63 int32 = 2147483647
                                _ = v62
                                _ = v63
                            }
                        }
                    case "foo bar":
                        if (fi43 > int(v2)) {
                            if v4 {
                                v42 = 127
                                v64 := v3[""]
                                if (v4 || (!v4)) {
                                    for fi65 := 0; fi65 < 15; fi65++ {
                                        switch (^(fi65)) {
                                        case 128:
                                            v66 := GenS1{F0: v1, F1: 42, F2: 4294967295}
                                            _ = v66
                                        case -80:
                                            fmt.Println(v44, v4)
                                        default:
                                            if (!v4) {
                                                fmt.Println(v1)
                                                if (!(v64 > (v2 + (v2 ^ 776)))) {
                                                    switch ("\n\t" + "foo bar") {
                                                    case "":
                                                        v67 := []int{}
                                                        _ = v67
                                                    default:
                                                        var v68 int32 = -1
                                                        _ = v68
                                                    }
                                                    var v69 float32 = ((3.14159 / 1.0) - float32(((3.14159 - float64(((-2.25 - float32(0.1)) * 1e10))) * 1.5)))
                                                    _ = v69
                                                }
                                            }
                                        }
                                        v2 += (v1 | v64)
                                        var v70 int32 = 699
                                        _ = fi65
                                        _ = v70
                                    }
                                }
                                _ = v64
                            }
                            v3[""] = (^(v2))
                            v71 := GenS7{F0: 2, F1: 222, F2: ("hello" + ("" + "\n\t"))}
                            _ = v71
                        } else {
                            for fi72 := 0; fi72 < 12; fi72++ {
                                if false {
                                    if v4 {
                                        if v4 {
                                            v73 := []int{}
                                            var v74 int8 = 127
                                            v42 = (v44 ^ 2)
                                            var v75 uint8 = v44
                                            switch "" {
                                            case "\n\t":
                                                if true {
                                                    if (false && (!((100 >= (v74 % (((v74 << ((255) & 63))) | 1))) && v4))) {
                                                        switch (fi43 >> ((351) & 63)) {
                                                        case 1:
                                                            fmt.Println(v74, v75)
                                                        }
                                                        switch "\n\t" {
                                                        case "a":
                                                            fmt.Println(fi72, v74)
                                                        case "hello":
                                                            v42 -= 42
                                                            if v4 {
                                                                var v76 int64 = -787
                                                                v77 := GenS1{F0: (^(v2)), F1: v76, F2: 65535}
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                            var v78 string = ("foo bar" + "\n\t")
                                                            v79 := GenS8{F0: 0, F1: (("a" + v78) == "")}
                                                            var v80 float32 = 1e10
                                                            _ = v78
                                                            _ = v79
                                                            _ = v80
                                                        case "foo bar":
                                                            v2 = v1
                                                        default:
                                                            var v81 uint8 = v44
                                                            _ = v81
                                                        }
                                                        switch (-(fi72)) {
                                                        case 647:
                                                            v1 = (v2 >> ((25) & 63))
                                                        default:
                                                            if (1e10 >= float64(0.0)) {
                                                                var v82 uint32 = 2147483648
                                                                var v83 uint = 42
                                                                v84 := GenS9{F0: -32768, F1: 28, F2: -458}
                                                                switch int(v74) {
                                                                case 0:
                                                                    var v85 uint32 = (v1 << (((v83 << ((0) & 63))) & 63))
                                                                    _ = v85
                                                                case 42:
                                                                    delete(v3, "\n\t")
                                                                default:
                                                                    v4 = (!(!((42 == 983) || false)))
                                                                }
                                                                _ = v82
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                            v86 := GenS9{F0: -1, F1: 108, F2: 1}
                                                            _ = v86
                                                        }
                                                    } else {
                                                        switch "foo bar" {
                                                        case "a":
                                                            if ((0.1 > 1e10) || v4) {
                                                                var v87 int = (fi72 & fi43)
                                                                fmt.Println(fi43, v4, fi43)
                                                                v74 -= v74
                                                                v4 = true
                                                                v2 -= v1
                                                                _ = v87
                                                            }
                                                        case "\n\t":
                                                            var v88 uint32 = (v1 << ((2) & 63))
                                                            if v4 {
                                                                switch 32767 {
                                                                case 4294967295:
                                                                    switch (fi72 << ((255) & 63)) {
                                                                    case 42:
                                                                        _ = len(v3)
                                                                    case 0:
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            switch (^(fi43)) {
                                                                            case -1:
                                                                                v89 := GenS7{F0: 42, F1: 831, F2: ("hello" + ("" + "a"))}
                                                                                _ = v89
                                                                            case -738:
                                                                                v90 := GenS4{F0: 735, F1: 127, F2: "\n\t"}
                                                                                _ = v90
                                                                            case -1000:
                                                                                if true {
                                                                                    v91 := GenS5{F0: 9223372036854775807, F1: 0, F2: (v44 * (v42 / (((v42 >> ((4294967295) & 63))) | 1)))}
                                                                                    if false {
                                                                                        if (fi72 >= fi72) {
                                                                                            var v92 uint64 = 4294967295
                                                                                            v88 = (v2 ^ (^(v1)))
                                                                                            if v4 {
                                                                                                v88 = (v88 | v2)
                                                                                            }
                                                                                            fmt.Println(v44, v74, v42)
                                                                                            var v93 bool = v4
                                                                                            _ = v92
                                                                                            _ = v93
                                                                                        } else {
                                                                                            v94 := GenS6{F0: v1, F1: -1, F2: v44}
                                                                                            v75 -= 90
                                                                                            v95 := GenS10{F0: (v74 | 100)}
                                                                                            var v96 int16 = 1000
                                                                                            _ = len(v3)
                                                                                            var v97 uint8 = (v42 / ((170) | 1))
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        }
                                                                                        v88 -= 0
                                                                                        delete(v3, ("foo bar" + "a"))
                                                                                    }
                                                                                    v4 = v4
                                                                                    _ = v91
                                                                                } else {
                                                                                    v98 := GenS11{F0: (fi72 % ((fi72) | 1)), F1: 42}
                                                                                    if v4 {
                                                                                        if v4 {
                                                                                            v73 = append(v73, (fi72 + fi43))
                                                                                        }
                                                                                        v99 := v3[("hello" + ("a" + "a"))]
                                                                                        _ = v99
                                                                                    } else {
                                                                                        if (!(!false)) {
                                                                                            switch 42 {
                                                                                            case -2147483648:
                                                                                                var v100 uint8 = 42
                                                                                                _ = v100
                                                                                            case 9223372036854775807:
                                                                                                v1 -= (v88 ^ 0)
                                                                                            case -783:
                                                                                                _ = len(v3)
                                                                                            default:
                                                                                                var v101 uint64 = 336
                                                                                                _ = v101
                                                                                            }
                                                                                            var v102 uint32 = v88
                                                                                            v1 -= 0
                                                                                            var v103 int64 = -602
                                                                                            v104 := GenS12{F0: (v88 % ((v1) | 1))}
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                        }
                                                                                        fmt.Println(v88, fi72, v1)
                                                                                    }
                                                                                    v4 = (!v4)
                                                                                    _ = v98
                                                                                }
                                                                            default:
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    v74 = (v74 >> ((125) & 63))
                                                                                case "a":
                                                                                    var v105 string = "foo bar"
                                                                                    _ = v105
                                                                                default:
                                                                                    if (!(-1 >= int16(fi43))) {
                                                                                        var v106 int16 = -635
                                                                                        _ = v106
                                                                                    } else {
                                                                                        var v107 int8 = -47
                                                                                        _ = v107
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            v1 += v88
                                                                        }
                                                                    case 128:
                                                                        v108 := GenS13{F0: 42, F1: 902}
                                                                        _ = v108
                                                                    default:
                                                                        delete(v3, ("\n\t" + "\n\t"))
                                                                    }
                                                                }
                                                                var v109 float32 = (1.5 * ((0.0 + ((1e10 / float32(3.14159)) / (1.5 / -2.25))) + -2.25))
                                                                _ = v109
                                                            } else {
                                                                v110 := GenS6{F0: v1, F1: 0, F2: (^(v44))}
                                                                var v111 uint32 = 65535
                                                                _ = v110
                                                                _ = v111
                                                            }
                                                            _ = v88
                                                        default:
                                                            if len(v73) >= 2 {
                                                                v73 = v73[1:len(v73)]
                                                            }
                                                        }
                                                        v4 = v4
                                                    }
                                                    if true {
                                                        var v112 int64 = -9223372036854775808
                                                        v113 := GenS1{F0: 1, F1: 9223372036854775807, F2: 18446744073709551615}
                                                        switch "" {
                                                        case "foo bar":
                                                            v42 = v44
                                                            var v114 int = (-(fi72))
                                                            _ = v114
                                                        }
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                } else {
                                                    switch "foo bar" {
                                                    case "a":
                                                        v73 = append(v73, (fi43 * (fi43 << ((4294967295) & 63))))
                                                    case "foo bar":
                                                        fmt.Println(v42)
                                                    case "":
                                                        fmt.Println(v2)
                                                    default:
                                                        var v115 int = fi72
                                                        _ = v115
                                                    }
                                                    if (0 < -501) {
                                                        if v4 {
                                                            var v116 string = "hello"
                                                            fmt.Println(fi72)
                                                            v117 := GenS14{F0: v2}
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        v44 += (v44 + uint8(v2))
                                                        var v118 uint16 = 0
                                                        fmt.Println(v1)
                                                        v74 += (v74 + -13)
                                                        _ = v118
                                                    } else {
                                                        v119 := v3[""]
                                                        v44 = v44
                                                        _ = v119
                                                    }
                                                    if len(v73) > 0 {
                                                        v73[0] = 0
                                                    }
                                                }
                                            case "hello":
                                                if (true && true) {
                                                    v120 := v3[(("" + "foo bar") + ("a" + "foo bar"))]
                                                    var v121 bool = (v4 || v4)
                                                    _ = v120
                                                    _ = v121
                                                }
                                            }
                                            v74 = (v74 ^ (v74 * -128))
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        }
                                        var v122 int32 = -107
                                        if v4 {
                                            var v123 int16 = 42
                                            v124 := []int{-942, fi43, (fi43 >> ((255) & 63))}
                                            var v125 int8 = -1
                                            _ = v123
                                            _ = v124
                                            _ = v125
                                        } else {
                                            v126 := GenS15{F0: "", F1: 42}
                                            switch (^(fi43)) {
                                            case 32768:
                                                if true {
                                                    if v4 {
                                                        v2 -= (^(v1))
                                                    } else {
                                                        v42 -= (v44 - v42)
                                                    }
                                                }
                                            default:
                                                v127 := GenS8{F0: 42, F1: v4}
                                                _ = v127
                                            }
                                            if v4 {
                                                v128 := []int{}
                                                var v129 float64 = ((((1.5 + float64(0.1)) - 0.1) - (float64((3.14159 / 1.5)) * 0.0)) * 0.1)
                                                v1 -= v1
                                                _ = v128
                                                _ = v129
                                            }
                                            _ = v126
                                        }
                                        _ = v122
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v42 = 36
                                    case "foo bar":
                                        v130 := v3[("a" + "\n\t")]
                                        _ = v130
                                    default:
                                        v2 = 65535
                                    }
                                    fmt.Println(v4, v42)
                                }
                                v131 := v3[("" + "foo bar")]
                                v42 -= (v42 / (((^(v44))) | 1))
                                _ = fi72
                                _ = v131
                            }
                            v132 := GenS10{F0: 0}
                            var v133 int8 = 42
                            _ = v132
                            _ = v133
                        }
                    case "a":
                        v2 = (v1 + v1)
                    }
                    if v4 {
                        fmt.Println(v2, v4, v4)
                        v44 += 42
                        v42 = (v42 | (v42 - (v44 & (v42 | 128))))
                    }
                    if ((float32((1e10 * 0.1)) - 3.14159) == 1.5) {
                        if true {
                            switch -589 {
                            case -74:
                                if ((!((fi43 / ((fi43) | 1)) >= fi43)) || (!(1 == 2))) {
                                    for fi134 := 0; fi134 < 10; fi134++ {
                                        var v135 bool = v4
                                        _ = fi134
                                        _ = v135
                                    }
                                    if (((24 >= 100) || v4) && false) {
                                        v4 = (!(true || v4))
                                        v136 := GenS3{F0: 1}
                                        _ = v136
                                    }
                                    v3["\n\t"] = v1
                                }
                            case 1:
                                v44 = v44
                            default:
                                v44 = v44
                            }
                        } else {
                            for fi137 := 0; fi137 < 13; fi137++ {
                                switch "" {
                                case "hello":
                                    if v4 {
                                        switch "a" {
                                        case "":
                                            delete(v3, "hello")
                                        }
                                        if v4 {
                                            switch "foo bar" {
                                            case "foo bar":
                                                if (true || (0 <= 655)) {
                                                    v3["foo bar"] = 894
                                                    if v4 {
                                                        v2 -= 4294967295
                                                        var v138 float32 = 1e10
                                                        fmt.Println(v2, v42, fi137)
                                                        _ = v138
                                                    }
                                                    v42 += 0
                                                }
                                            case "\n\t":
                                                var v139 uint = 134
                                                _ = v139
                                            case "hello":
                                                var v140 string = "hello"
                                                if v4 {
                                                    delete(v3, v140)
                                                } else {
                                                    var v141 float32 = 0.1
                                                    _ = v141
                                                }
                                                _ = v140
                                            }
                                            v142 := GenS5{F0: -531, F1: -1, F2: v42}
                                            switch (("\n\t" + "") + ("a" + "a")) {
                                            case "hello":
                                                var v143 int64 = 42
                                                _ = v143
                                            default:
                                                switch "" {
                                                case "":
                                                    var v144 uint16 = 32767
                                                    _ = v144
                                                case "hello":
                                                    var v145 uint = 364
                                                    _ = v145
                                                case "a":
                                                    if (127 != v44) {
                                                        if v4 {
                                                            var v146 uint16 = 255
                                                            v147 := []int{(fi43 >> ((1) & 63)), fi137, (fi43 & fi43)}
                                                            _ = v146
                                                            _ = v147
                                                        }
                                                        v42 += uint8(v1)
                                                    }
                                                default:
                                                    var v148 int = (fi43 / ((-9223372036854775808) | 1))
                                                    _ = v148
                                                }
                                            }
                                            _ = v142
                                        } else {
                                            var v149 float32 = 0.1
                                            _ = v149
                                        }
                                    }
                                default:
                                    if ((v42 & v42) <= 127) {
                                        if v4 {
                                            var v150 uint16 = 1
                                            _ = v150
                                        } else {
                                            var v151 uint8 = v42
                                            v2 -= (v2 ^ (v2 & 2147483647))
                                            _ = v151
                                        }
                                        var v152 string = "hello"
                                        fmt.Println(v152, v4)
                                        if v4 {
                                            v44 = (v42 | v44)
                                            var v153 uint = 42
                                            v154 := GenS6{F0: (v1 ^ (v2 * (^(v2)))), F1: -1, F2: 225}
                                            _ = v153
                                            _ = v154
                                        } else {
                                            v155 := GenS4{F0: -419, F1: 19, F2: "\n\t"}
                                            _ = v155
                                        }
                                        fmt.Println(v2)
                                        _ = v152
                                    }
                                }
                                var v156 uint32 = (v1 | (v1 ^ v1))
                                _ = fi137
                                _ = v156
                            }
                            var v157 int32 = -2147483648
                            _ = v157
                        }
                    } else {
                        v158 := GenS16{F0: v1}
                        v159 := GenS11{F0: fi43, F1: 1}
                        v160 := v3[(("a" + "") + "hello")]
                        _ = v158
                        _ = v159
                        _ = v160
                    }
                }
                v161 := []int{(fi43 | 32768), (fi43 - (fi43 + fi43))}
                _ = fi43
                _ = v161
                _ = v44
            }
        case "":
            v162 := v3["a"]
            _ = v162
        default:
            v163 := v3["\n\t"]
            _ = v163
        }
    }
    v2 += (v1 / (((v2 * (v1 >> ((65535) & 63)))) | 1))
    v2 += (v2 >> ((18446744073709551615) & 63))
    if (18446744073709551615 < 42) {
        v164 := []int{1}
        _ = len(v3)
        v42 = v42
        _ = v164
    }
    v1 = (v1 - v2)
    switch (("hello" + "a") + "") {
    case "":
        var v165 int32 = 406
        _ = v165
    case "foo bar":
        fmt.Println(v2, v4, v2)
    default:
        v166 := GenS5{F0: -9223372036854775808, F1: -9223372036854775808, F2: (v42 % ((v42) | 1))}
        _ = v166
    }
    for fi167 := 0; fi167 < 13; fi167++ {
        v168 := []int{}
        _ = fi167
        _ = v168
    }
    switch 376 {
    case 9223372036854775807:
        v2 += 777
    case 0:
        fmt.Println(v1, v2)
    default:
        for fi169 := 0; fi169 < 8; fi169++ {
            _ = len(v3)
            _ = fi169
        }
    }
    switch 2147483647 {
    case 9223372036854775807:
        v4 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
