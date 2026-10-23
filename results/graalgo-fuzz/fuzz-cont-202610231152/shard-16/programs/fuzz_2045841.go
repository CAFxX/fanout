package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 bool; F2 int16 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 int64; F1 bool; F2 string }

func main() {
    var v1 int = 42
    v2 := make(map[int]int)
    v3 := []uint32{0, 1000, 975}
    v4 := make(map[string]string)
    var v5 uint = 18446744073709551615
    var v6 bool = false
    var v7 int8 = 1
    if v6 {
        for fi8 := 0; fi8 < 15; fi8++ {
            delete(v2, 128)
            if v6 {
                for fi9 := 0; fi9 < 17; fi9++ {
                    fmt.Println(fi8)
                    _ = fi9
                }
                v7 -= (v7 % ((96) | 1))
            }
            switch (v1 - (fi8 / ((v1) | 1))) {
            case 1:
                var v10 int = (-(v1))
                _ = v10
            case 42:
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch (("hello" + "") + ("a" + (("" + "hello") + ("a" + ("foo bar" + "foo bar"))))) {
                    case "":
                        if v6 {
                            v4[("\n\t" + "foo bar")] = "a"
                            switch ("foo bar" + "foo bar") {
                            case "\n\t":
                                v12 := GenS1{F0: 0, F1: false, F2: 342}
                                _ = v12
                            case "foo bar":
                                var v13 bool = v6
                                _ = v13
                            }
                        }
                    case "a":
                        _ = len(v3)
                    }
                    _ = fi11
                }
            }
            _ = fi8
        }
    } else {
        var v14 int = (v1 ^ (v1 / (((v1 >> (((v5 - v5)) & 63))) | 1)))
        v6 = (true || (v1 == (v14 - v1)))
        _ = v14
    }
    v7 = (v7 ^ v7)
    var v15 int64 = 9223372036854775807
    v16 := GenS1{F0: 128, F1: (-2.25 > 0.1), F2: 0}
    if (!(0.1 > -2.25)) {
        fmt.Println(v15, v1, v5)
        if (2 >= 1) {
            var v17 int8 = 127
            _ = v17
        }
    } else {
        _ = len(v3)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 -= (v1 + v1)
    for fi18 := 0; fi18 < 14; fi18++ {
        switch "a" {
        case "hello":
            switch fi18 {
            case 2147483647:
                var v19 uint32 = 2
                _ = v19
            case 32767:
                if v6 {
                    v1 = (v1 % (((fi18 / ((fi18) | 1))) | 1))
                    var v20 uint16 = 92
                    switch "a" {
                    case "foo bar":
                        switch "hello" {
                        case "\n\t":
                            if ((true || (42 > 2)) || v6) {
                                v21 := GenS1{F0: 128, F1: true, F2: v20}
                                _ = v21
                            } else {
                                _ = len(v3)
                            }
                        case "foo bar":
                            var v22 uint16 = (v20 + (^(v20)))
                            _ = v22
                        }
                    case "hello":
                        switch int(v15) {
                        case 669:
                            for fi23 := 0; fi23 < 15; fi23++ {
                                var v24 uint8 = 127
                                _ = fi23
                                _ = v24
                            }
                        case 42:
                            v7 = v7
                        default:
                            var v25 float64 = 0.1
                            _ = v25
                        }
                    }
                    _ = v20
                } else {
                    switch ("hello" + "a") {
                    case "foo bar":
                        if (((((-1 > -661) || (42 == -775)) && v6) && v6) || (-9223372036854775808 < v15)) {
                            fmt.Println(fi18, v1)
                        } else {
                            fmt.Println(v1, v7, v7)
                            fmt.Println(v15, v7, v5)
                            for fi26 := 0; fi26 < 18; fi26++ {
                                var v27 int16 = 0
                                v27 = 32767
                                _ = fi26
                                _ = v27
                            }
                        }
                    case "a":
                        for fi28 := 0; fi28 < 7; fi28++ {
                            if (42 == 888) {
                                _ = len(v4)
                                if v6 {
                                    v1 = (v1 * 9223372036854775807)
                                    var v29 int16 = -1
                                    var v30 int16 = 61
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                var v31 float32 = (-2.25 + 0.1)
                                fmt.Println(v5, v1)
                                _ = v31
                            }
                            _ = fi28
                        }
                    case "":
                        v32 := GenS1{F0: 85, F1: true, F2: 240}
                        _ = v32
                    default:
                        if ((v15 >> ((602) & 63)) <= int64(v5)) {
                            v33 := GenS2{F0: "a"}
                            for fi34 := 0; fi34 < 13; fi34++ {
                                if (int32(v7) < 0) {
                                    v35 := GenS3{F0: v7}
                                    fmt.Println(v15, v1)
                                    var v36 uint = (v5 - v5)
                                    _ = v35
                                    _ = v36
                                }
                                var v37 int = (v1 << ((255) & 63))
                                if false {
                                    if true {
                                        if (704 < 2) {
                                            fmt.Println(v15, fi34, v5)
                                            v7 -= v7
                                        }
                                    } else {
                                        switch 32767 {
                                        case 20:
                                            v38 := v4[(("foo bar" + ("hello" + "a")) + (("foo bar" + "hello") + ""))]
                                            v39 := GenS2{F0: v38}
                                            v6 = v6
                                            _ = v38
                                            _ = v39
                                        case -2147483648:
                                            v37 += int(v5)
                                        case 76:
                                            v40 := GenS3{F0: (v7 << ((1) & 63))}
                                            _ = v40
                                        }
                                        if ("" < "") {
                                            v41 := GenS2{F0: "\n\t"}
                                            var v42 int32 = 0
                                            if (0 != (v1 * 297)) {
                                                fmt.Println(v42, v1)
                                                var v43 uint = 385
                                                var v44 float32 = (float32(float64(0.0)) + -2.25)
                                                _ = v43
                                                _ = v44
                                            } else {
                                                var v45 int64 = v15
                                                if (42 >= (-(v7))) {
                                                    v7 = (v7 + -128)
                                                } else {
                                                    v4[""] = "hello"
                                                    delete(v4, "")
                                                }
                                                if v6 {
                                                    if (true || true) {
                                                        if v6 {
                                                            fmt.Println(v45, v7)
                                                            switch v1 {
                                                            case -250:
                                                                v7 = -29
                                                            case 1:
                                                                v46 := GenS1{F0: 2, F1: (!v6), F2: 32767}
                                                                _ = v46
                                                            case -2147483648:
                                                                var v47 uint16 = 613
                                                                _ = v47
                                                            }
                                                        }
                                                    } else {
                                                        v48 := GenS1{F0: 0, F1: ((v45 ^ -1) == (v15 + v15)), F2: 656}
                                                        v15 += (v15 - 9223372036854775807)
                                                        _ = v48
                                                    }
                                                }
                                                _ = v45
                                            }
                                            v49 := GenS1{F0: 1, F1: v6, F2: 2}
                                            _ = v41
                                            _ = v42
                                            _ = v49
                                        }
                                    }
                                    fmt.Println(v15, v5)
                                    if v6 {
                                        var v50 int = v37
                                        v51 := GenS3{F0: 127}
                                        _ = v50
                                        _ = v51
                                    } else {
                                        switch ("a" + "") {
                                        case "\n\t":
                                            if v6 {
                                                var v52 int64 = v15
                                                v4[""] = (("a" + (("a" + "a") + (("\n\t" + "\n\t") + ("\n\t" + "hello")))) + "foo bar")
                                                switch (fi18 >> ((v5) & 63)) {
                                                case 881:
                                                    var v53 int32 = -109
                                                    _ = v53
                                                case 32768:
                                                    var v54 string = ("hello" + "hello")
                                                    v55 := GenS2{F0: v54}
                                                    _ = v54
                                                    _ = v55
                                                default:
                                                    if v6 {
                                                        v4[("hello" + "foo bar")] = "a"
                                                        v52 = 42
                                                        if (255 >= 65535) {
                                                            switch "hello" {
                                                            case "a":
                                                                var v56 uint = 18446744073709551615
                                                                _ = v56
                                                            case "\n\t":
                                                                var v57 string = "foo bar"
                                                                _ = v57
                                                            }
                                                            fmt.Println(v5, fi34)
                                                        }
                                                        var v58 uint8 = 1
                                                        _ = v58
                                                    }
                                                }
                                                _ = v52
                                            } else {
                                                v1 -= (-(v1))
                                            }
                                        default:
                                            v1 += 32768
                                        }
                                        v15 += v15
                                        v15 += (v15 / ((v15) | 1))
                                        if false {
                                            if v6 {
                                                var v59 uint64 = 0
                                                if ((950 > -68) || v6) {
                                                    v6 = true
                                                }
                                                _ = v59
                                            }
                                        }
                                    }
                                }
                                _ = fi34
                                _ = v37
                            }
                            switch "\n\t" {
                            case "hello":
                                v60 := GenS3{F0: -1}
                                _ = v60
                            case "\n\t":
                                var v61 bool = v6
                                _ = v61
                            default:
                                var v62 float64 = (1.5 + 0.1)
                                _ = v62
                            }
                            _ = v33
                        }
                    }
                }
            }
        case "\n\t":
            v6 = false
        default:
            var v63 bool = true
            _ = v63
        }
        v15 = 9223372036854775807
        for fi64 := 0; fi64 < 4; fi64++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v65 string = ("\n\t" + "foo bar")
            _ = fi64
            _ = v65
        }
        _ = fi18
    }
    v7 += v7
    if (v6 || (!v6)) {
        var v66 uint32 = 791
        if (v6 || (296 < -752)) {
            var v67 bool = v6
            _ = v67
        }
        _ = v66
    }
    for fi68 := 0; fi68 < 11; fi68++ {
        if (567 >= 474) {
            if false {
                var v69 bool = (!(9223372036854775807 >= (v15 * (v15 + v15))))
                _ = v69
            }
        } else {
            switch ("hello" + "hello") {
            case "foo bar":
                for fi70 := 0; fi70 < 3; fi70++ {
                    v5 += 0
                    _ = fi70
                }
            case "a":
                var v71 int8 = (v7 / (((v7 / ((v7) | 1))) | 1))
                _ = v71
            }
            fmt.Println(v5)
        }
        v15 -= int64(v7)
        v7 = v7
        _ = fi68
    }
    for fi72 := 0; fi72 < 19; fi72++ {
        var v73 int64 = (^(v15))
        if ((0.1 * -2.25) < 1e10) {
            fmt.Println(v1, v7)
        }
        _ = fi72
        _ = v73
    }
    for fi74 := 0; fi74 < 13; fi74++ {
        var v75 uint32 = 488
        _ = fi74
        _ = v75
    }
    var v76 int = (v1 / (((v1 - v1)) | 1))
    if true {
        if v6 {
            var v77 uint = (^(v5))
            fmt.Println(v6, v77)
            _ = v77
        } else {
            for fi78 := 0; fi78 < 3; fi78++ {
                var v79 uint16 = 299
                var v80 bool = false
                v3 = append(v3, 65535)
                _ = fi78
                _ = v79
                _ = v80
            }
            for fi81 := 0; fi81 < 4; fi81++ {
                if (v6 || (v6 || v6)) {
                    v1 += -2147483648
                    for fi82 := 0; fi82 < 11; fi82++ {
                        var v83 string = "\n\t"
                        switch (("" + ("\n\t" + (v83 + "a"))) + (v83 + v83)) {
                        case "foo bar":
                            var v84 uint = v5
                            _ = v84
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                fmt.Println(v15, fi81)
                            case "foo bar":
                                v3 = append(v3, 1)
                            case "":
                                if false {
                                    switch (v83 + ((v83 + (("a" + v83) + "\n\t")) + "\n\t")) {
                                    case "foo bar":
                                        if true {
                                            switch (-(v76)) {
                                            case 103:
                                                var v85 uint64 = 18446744073709551615
                                                _ = v85
                                            case 1:
                                                if true {
                                                    v3 = append(v3, 42)
                                                }
                                            case -485:
                                                v5 += (v5 << ((381) & 63))
                                            default:
                                                v6 = (((v6 && v6) && ((v5 & v5) >= (v5 | v5))) && (v15 > 299))
                                            }
                                            switch 9223372036854775807 {
                                            case 32768:
                                                if (!v6) {
                                                    v86 := GenS2{F0: ""}
                                                    fmt.Println(v5, fi81, v1)
                                                    v83 = v83
                                                    v6 = (float32(0.1) >= ((((-2.25 + 0.1) / (1e10 / (float32((1.5 + 3.14159)) - -2.25))) - (1.5 + ((((3.14159 / float32((-2.25 + 1.5))) * float32(3.14159)) + (-2.25 / 1.0)) / 3.14159))) + (float32(float64(0.1)) - (float32(0.1) / (0.1 * 3.14159)))))
                                                    switch v83 {
                                                    case "foo bar":
                                                        switch v83 {
                                                        case "hello":
                                                            if false {
                                                                v7 = -73
                                                                var v87 uint64 = 1
                                                                _ = v87
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v88 uint16 = 255
                                                        if (v88 > 1) {
                                                            switch fi81 {
                                                            case -1:
                                                                v89 := GenS1{F0: 0, F1: ((1.5 - (((0.1 + (-2.25 + 1.5)) - -2.25) - 0.1)) == 0.0), F2: 0}
                                                                _ = v89
                                                            case -9223372036854775808:
                                                                v2[v76] = (v76 + v1)
                                                            case -2147483648:
                                                                _ = len(v2)
                                                            }
                                                            if ((!false) && v6) {
                                                                if (!true) {
                                                                    v3 = append(v3, 4294967295)
                                                                } else {
                                                                    v1 = (v76 << ((v5) & 63))
                                                                    var v90 int = v1
                                                                    v91 := GenS2{F0: v83}
                                                                    var v92 int32 = -181
                                                                    _ = v90
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                v83 = ("\n\t" + "")
                                                                if false {
                                                                    v7 += v7
                                                                } else {
                                                                    v83 = v83
                                                                    if v6 {
                                                                        var v93 float32 = -2.25
                                                                        v88 = 65535
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 2
                                                                        }
                                                                        _ = v93
                                                                    }
                                                                    var v94 bool = v6
                                                                    _ = v94
                                                                }
                                                                if v6 {
                                                                    var v95 int8 = v7
                                                                    var v96 uint8 = 147
                                                                    if (v6 || false) {
                                                                        v15 += 960
                                                                        if ((v7 % (((v95 - v7)) | 1)) <= v7) {
                                                                            _ = len(v2)
                                                                        } else {
                                                                            v1 += -75
                                                                            if v6 {
                                                                                _ = len(v3)
                                                                                var v97 uint8 = 0
                                                                                _ = v97
                                                                            }
                                                                            var v98 int = (fi82 ^ -936)
                                                                            _ = v98
                                                                        }
                                                                        fmt.Println(v5, v76, v15)
                                                                    } else {
                                                                        switch fi82 {
                                                                        case 32767:
                                                                            v7 += (v7 * 42)
                                                                        case 1:
                                                                            if (!(!v6)) {
                                                                                v99 := GenS2{F0: "\n\t"}
                                                                                if (v6 || (true || (v6 && (v6 || false)))) {
                                                                                    v5 = (v5 ^ (v5 * (v5 / ((279) | 1))))
                                                                                    v83 = "a"
                                                                                    v83 = ("\n\t" + (v83 + ("" + "\n\t")))
                                                                                }
                                                                                _ = v99
                                                                            } else {
                                                                                var v100 uint16 = v88
                                                                                var v101 uint64 = 1
                                                                                v96 = (v96 | (v96 >> ((v5) & 63)))
                                                                                _ = v100
                                                                                _ = v101
                                                                            }
                                                                        default:
                                                                            switch v83 {
                                                                            case "\n\t":
                                                                                v102 := v4[v83]
                                                                                _ = v102
                                                                            case "foo bar":
                                                                                if (int(v15) <= (v1 | 1)) {
                                                                                    if (!v6) {
                                                                                        if ((1 >= 0) && true) {
                                                                                            fmt.Println(v96)
                                                                                            var v103 int32 = 42
                                                                                            switch (v1 - 127) {
                                                                                            case 9223372036854775807:
                                                                                                if true {
                                                                                                    switch v83 {
                                                                                                    case "foo bar":
                                                                                                        v104 := GenS1{F0: (v96 + (^(v96))), F1: (v6 || (!v6)), F2: (v88 | (v88 & (v88 / ((32767) | 1))))}
                                                                                                        _ = v104
                                                                                                    case "\n\t":
                                                                                                        var v105 bool = v6
                                                                                                        _ = v105
                                                                                                    }
                                                                                                }
                                                                                            case 0:
                                                                                                v4["\n\t"] = v83
                                                                                            case -1:
                                                                                                var v106 string = "hello"
                                                                                                _ = v106
                                                                                            }
                                                                                            _ = v103
                                                                                        } else {
                                                                                            switch fi82 {
                                                                                            case 32767:
                                                                                                var v107 string = "hello"
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v108 := GenS4{F0: 2, F1: (v95 | (v95 / (((v7 ^ (v95 | 56))) | 1)))}
                                                                                        _ = v108
                                                                                    }
                                                                                } else {
                                                                                    var v109 int16 = 0
                                                                                    v110 := GenS5{F0: (v5 + (v5 >> (((v5 >> ((1) & 63))) & 63)))}
                                                                                    var v111 float64 = 0.0
                                                                                    switch ("foo bar" + v83) {
                                                                                    case "":
                                                                                        if (1 >= 4294967295) {
                                                                                            var v112 uint64 = 0
                                                                                            v7 += (^(v95))
                                                                                            _ = v112
                                                                                        }
                                                                                    }
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                            case "":
                                                                                var v113 float32 = 1e10
                                                                                _ = v113
                                                                            default:
                                                                                _ = len(v3)
                                                                            }
                                                                        }
                                                                        var v114 int32 = -23
                                                                        v115 := GenS3{F0: -36}
                                                                        if v6 {
                                                                            var v116 int16 = int16(fi81)
                                                                            v117 := v4[v83]
                                                                            v118 := GenS2{F0: (v83 + "a")}
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                        }
                                                                        fmt.Println(fi81, v7, v114)
                                                                        var v119 string = v83
                                                                        _ = v114
                                                                        _ = v115
                                                                        _ = v119
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 65535
                                                                    }
                                                                    v76 = fi82
                                                                    _ = v95
                                                                    _ = v96
                                                                } else {
                                                                    v120 := GenS1{F0: 32, F1: false, F2: v88}
                                                                    _ = v120
                                                                }
                                                            } else {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                fmt.Println(v7, v76, v7)
                                                                if (false && v6) {
                                                                    fmt.Println(v83, v5, v7)
                                                                    if v6 {
                                                                        var v121 int = (-(fi82))
                                                                        v3 = append(v3, 2147483647)
                                                                        if ((v15 >> ((18446744073709551615) & 63)) <= 1) {
                                                                            delete(v4, v83)
                                                                            switch v83 {
                                                                            case "":
                                                                                _ = len(v4)
                                                                            default:
                                                                                v122 := GenS6{F0: (v7 + (v7 | v7)), F1: (-418 < 0), F2: int16(v88)}
                                                                                _ = v122
                                                                            }
                                                                        }
                                                                        _ = len(v3)
                                                                        _ = v121
                                                                    }
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v88
                                                    case "hello":
                                                        fmt.Println(v7, v15, v7)
                                                        if len(v3) > 0 {
                                                            v3[0] = 91
                                                        }
                                                    }
                                                    v7 = v7
                                                    _ = v86
                                                } else {
                                                    fmt.Println(v7, v76, v83)
                                                    var v123 uint16 = 255
                                                    _ = v123
                                                }
                                            case 128:
                                                var v124 uint8 = 1
                                                _ = v124
                                            default:
                                                switch (fi81 * (-(v1))) {
                                                case 342:
                                                    var v125 uint16 = 79
                                                    _ = v125
                                                case -769:
                                                    var v126 float64 = (0.1 * (3.14159 / ((0.1 / (-2.25 * 1.5)) - (0.1 * (1e10 * 1.5)))))
                                                    _ = v126
                                                case -749:
                                                    v83 = (v83 + ((v83 + "\n\t") + (v83 + ((v83 + "") + "\n\t"))))
                                                default:
                                                    var v127 int64 = v15
                                                    _ = v127
                                                }
                                            }
                                            if (!v6) {
                                                if (0 > 255) {
                                                    if len(v3) > 0 {
                                                        v3[0] = 715
                                                    }
                                                    switch v83 {
                                                    case "a":
                                                        var v128 bool = (v6 || ((-1 != -32768) || (v7 != (v7 >> (((v5 << ((v5) & 63))) & 63)))))
                                                        _ = v128
                                                    case "hello":
                                                        v129 := GenS7{F0: (0.1 < 3.14159), F1: (255 != 42)}
                                                        _ = v129
                                                    default:
                                                        delete(v2, (fi81 - fi82))
                                                    }
                                                } else {
                                                    v130 := GenS1{F0: 127, F1: (((fi81 / (((v1 + (fi81 - 2147483647))) | 1)) <= -299) && (!v6)), F2: 2}
                                                    _ = v130
                                                }
                                                var v131 bool = (42 >= v5)
                                                fmt.Println(v83)
                                                _ = v131
                                            } else {
                                                v132 := GenS6{F0: (^(v7)), F1: (!((v15 ^ 637) <= (v15 | (v15 - (v15 + (v15 % ((v15) | 1))))))), F2: 852}
                                                fmt.Println(fi81, v7)
                                                v1 = (^(v76))
                                                fmt.Println(v76, v6, v83)
                                                v76 = 32768
                                                if v6 {
                                                    v133 := GenS4{F0: 42, F1: v7}
                                                    v3 = append(v3, 4294967295)
                                                    switch (v83 + v83) {
                                                    case "a":
                                                        v83 = (v83 + "hello")
                                                    default:
                                                        if v6 {
                                                            var v134 float32 = float32((-2.25 - ((3.14159 / 3.14159) + (0.0 * 0.0))))
                                                            fmt.Println(v134)
                                                            v135 := GenS7{F0: v6, F1: (-320 > 32767)}
                                                            _ = v134
                                                            _ = v135
                                                        }
                                                    }
                                                    _ = v133
                                                }
                                                var v136 uint64 = 1
                                                _ = v132
                                                _ = v136
                                            }
                                        } else {
                                            var v137 int32 = 0
                                            switch "" {
                                            case "a":
                                                fmt.Println(v6, v6, v76)
                                            case "":
                                                if len(v3) > 0 {
                                                    v3[0] = 145
                                                }
                                            }
                                            var v138 uint8 = 255
                                            _ = v137
                                            _ = v138
                                        }
                                    case "hello":
                                        v76 = v1
                                    default:
                                        v139 := GenS7{F0: v6, F1: (((1e10 - float64(0.1)) - (1.5 + float64(1.5))) < 1e10)}
                                        if (0.0 != (-2.25 / 0.1)) {
                                            switch "foo bar" {
                                            case "a":
                                                if (-1 > (v15 / (((v15 >> ((65535) & 63))) | 1))) {
                                                    v3 = append(v3, 1)
                                                } else {
                                                    if true {
                                                        var v140 bool = true
                                                        v3 = append(v3, 1)
                                                        _ = v140
                                                    }
                                                    switch v76 {
                                                    case 2147483647:
                                                        if v6 {
                                                            var v141 float32 = (-2.25 / (0.1 / 3.14159))
                                                            _ = v141
                                                        }
                                                    case 128:
                                                        var v142 int8 = v7
                                                        _ = v142
                                                    case -2147483648:
                                                        var v143 uint32 = 65535
                                                        _ = v143
                                                    default:
                                                        fmt.Println(v7, v76)
                                                    }
                                                    if true {
                                                        v144 := GenS5{F0: (^(v5))}
                                                        v145 := GenS2{F0: "a"}
                                                        v1 = fi81
                                                        fmt.Println(fi81, v7, v6)
                                                        if (true && (255 > 16)) {
                                                            var v146 uint64 = 18446744073709551615
                                                            _ = v146
                                                        }
                                                        var v147 int = 0
                                                        var v148 bool = v6
                                                        v149 := GenS8{F0: 127}
                                                        v150 := GenS3{F0: 127}
                                                        v7 = (-(v7))
                                                        var v151 int32 = -562
                                                        _ = v144
                                                        _ = v145
                                                        _ = v147
                                                        _ = v148
                                                        _ = v149
                                                        _ = v150
                                                        _ = v151
                                                    } else {
                                                        var v152 int64 = (^(v15))
                                                        v152 += -570
                                                        fmt.Println(fi82, v152)
                                                        _ = v152
                                                    }
                                                    var v153 float64 = 0.1
                                                    _ = v153
                                                }
                                                v15 = v15
                                            case "hello":
                                                switch 32768 {
                                                case 32768:
                                                    if v6 {
                                                        var v154 float32 = float32(0.0)
                                                        v155 := GenS6{F0: 2, F1: ((v7 | (v7 & (v7 & v7))) == (-(v7))), F2: 0}
                                                        _ = v154
                                                        _ = v155
                                                    }
                                                case 903:
                                                    fmt.Println(fi82, v83)
                                                default:
                                                    var v156 uint = (v5 & (v5 % ((255) | 1)))
                                                    _ = v156
                                                }
                                            case "":
                                                v157 := GenS2{F0: v83}
                                                _ = v157
                                            }
                                            if (v5 <= 0) {
                                                switch v83 {
                                                case "foo bar":
                                                    var v158 string = (("hello" + "") + ("" + (v83 + v83)))
                                                    _ = v158
                                                case "a":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "\n\t":
                                                    v15 += (v15 + 1)
                                                }
                                            }
                                            var v159 int = (^(v76))
                                            _ = v159
                                        } else {
                                            v83 = ""
                                            var v160 uint = 18446744073709551615
                                            switch v76 {
                                            case -2147483648:
                                                v161 := GenS3{F0: (v7 % (((v7 << ((v160) & 63))) | 1))}
                                                _ = v161
                                            case 81:
                                                v5 -= v160
                                            default:
                                                v162 := v2[v76]
                                                _ = v162
                                            }
                                            _ = v160
                                        }
                                        _ = v139
                                    }
                                } else {
                                    v163 := GenS9{F0: -882, F1: 296}
                                    _ = len(v3)
                                    var v164 float32 = 3.14159
                                    _ = v163
                                    _ = v164
                                }
                            }
                        }
                        if true {
                            var v165 float64 = 3.14159
                            fmt.Println(v5, v83, v15)
                            _ = v165
                        }
                        _ = fi82
                        _ = v83
                    }
                } else {
                    if v6 {
                        var v166 uint16 = 255
                        v3 = append(v3, 42)
                        var v167 uint32 = 266
                        _ = v166
                        _ = v167
                    }
                    if v6 {
                        v168 := GenS8{F0: 1}
                        for fi169 := 0; fi169 < 5; fi169++ {
                            v7 = -1
                            var v170 float64 = (3.14159 / 0.1)
                            v76 = fi81
                            _ = fi169
                            _ = v170
                        }
                        _ = v168
                    } else {
                        _ = len(v3)
                        if (940 == uint16(v5)) {
                            v3 = append(v3, 272)
                            v6 = (true || v6)
                        } else {
                            var v171 uint8 = 42
                            _ = v171
                        }
                    }
                }
                v172 := GenS10{F0: v15, F1: v6, F2: ("a" + "a")}
                _ = fi81
                _ = v172
            }
        }
        v173 := v4[("hello" + "")]
        var v174 int8 = (v7 ^ 51)
        _ = v173
        _ = v174
    } else {
        var v175 int8 = 100
        for fi176 := 0; fi176 < 10; fi176++ {
            v76 = (v76 / (((-(v76))) | 1))
            v3 = append(v3, 193)
            v5 = (v5 | v5)
            _ = fi176
        }
        _ = v175
    }
    switch "" {
    case "":
        v177 := GenS8{F0: 70}
        _ = v177
    case "hello":
        fmt.Println(v76, v6, v15)
    }
    var v178 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v178:", v178)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v15
    _ = v16
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
}
