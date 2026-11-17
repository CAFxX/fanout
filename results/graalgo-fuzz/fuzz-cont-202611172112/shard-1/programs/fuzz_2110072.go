package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 int8; F2 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int, p1 uint64) int64 {
    v6 := GenS1{F0: 0, F1: -1, F2: 189}
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(fi8, fi7, fi7)
            _ = fi8
        }
        _ = fi7
    }
    var v9 uint16 = 32767
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := []int{32767, -532, 4294967295, 32768}
    v2 := make(map[string]uint32)
    var v3 int8 = -120
    var v4 float32 = (1.5 - -2.25)
    var v5 uint32 = 0
    if (((2 > 243) || true) || (true && true)) {
        if (235 == 4294967295) {
            var v10 int32 = -289
            for fi11 := 0; fi11 < 2; fi11++ {
                fmt.Println(v10)
                _ = fi11
            }
            _ = v10
        }
        v4 += v4
        var v12 int = -704
        _ = v12
    } else {
        switch "a" {
        case "foo bar":
            if ((!(-9223372036854775808 == 9223372036854775807)) || (!(int(v5) < -9223372036854775808))) {
                for fi13 := 0; fi13 < 4; fi13++ {
                    switch (("hello" + "a") + (("" + "hello") + "hello")) {
                    case "foo bar":
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v4 -= float32((-2.25 / -2.25))
                            _ = fi14
                        }
                    case "a":
                        var v15 uint16 = 65535
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := GenS2{F0: 32767, F1: int32(v3)}
                        v5 = 0
                        v18 := v2["a"]
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = fi13
                }
                for fi19 := 0; fi19 < 12; fi19++ {
                    if true {
                        v4 += (((-2.25 * float32(1e10)) / -2.25) - v4)
                        v4 = ((v4 - v4) - v4)
                        _ = len(v2)
                    } else {
                        v20 := GenS3{F0: 1, F1: v5}
                        _ = v20
                    }
                    var v21 uint8 = 0
                    _ = fi19
                    _ = v21
                }
                if (!false) {
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        v3 = -128
                    case "a":
                        v3 = (v3 / (((-(v3))) | 1))
                    case "foo bar":
                        fmt.Println(v4)
                    }
                } else {
                    var v22 float32 = 1.5
                    if (42 > 42) {
                        var v23 uint32 = (v5 * (v5 & (^(v5))))
                        _ = v23
                    } else {
                        v24 := v2["hello"]
                        if ((42 < (-(v3))) && true) {
                            v5 = v5
                            v4 = float32(0.0)
                            if (false || true) {
                                var v25 uint8 = 77
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    if false {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        fmt.Println(v4)
                                    } else {
                                        v5 -= v5
                                    }
                                    _ = fi26
                                }
                                _ = v25
                            } else {
                                for fi27 := 0; fi27 < 6; fi27++ {
                                    if ((!(!(461 < 963))) && true) {
                                        for fi28 := 0; fi28 < 18; fi28++ {
                                            switch (("foo bar" + ("\n\t" + ("a" + ("hello" + "\n\t")))) + "hello") {
                                            case "":
                                                var v29 uint64 = 795
                                                var v30 bool = false
                                                var v31 uint64 = (^(v29))
                                                switch "foo bar" {
                                                case "":
                                                    if (!(true && v30)) {
                                                        v32 := GenS4{F0: ("\n\t" + "")}
                                                        var v33 int64 = -1
                                                        fmt.Println(v4, v30, fi28)
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                case "hello":
                                                    v34 := v2["\n\t"]
                                                    _ = v34
                                                case "\n\t":
                                                    switch "a" {
                                                    case "hello":
                                                        var v35 uint16 = 2
                                                        _ = v35
                                                    case "\n\t":
                                                        v30 = v30
                                                    }
                                                }
                                                if ((v31 ^ 0) <= (^(v29))) {
                                                    var v36 uint16 = 2
                                                    v37 := GenS5{F0: (v36 | (v36 | 42)), F1: v29}
                                                    var v38 uint16 = (v36 >> ((1) & 63))
                                                    _ = v36
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    if (!(238 < 368)) {
                                                        var v39 uint16 = 570
                                                        _ = v39
                                                    } else {
                                                        var v40 string = ("\n\t" + (("\n\t" + ("a" + ("\n\t" + "\n\t"))) + "hello"))
                                                        if len(v1) > 0 {
                                                            v1[0] = (fi28 * fi27)
                                                        }
                                                        _ = len(v1)
                                                        _ = v40
                                                    }
                                                    switch "foo bar" {
                                                    case "":
                                                        if len(v1) > 0 {
                                                            v1[0] = -9223372036854775808
                                                        }
                                                    case "a":
                                                        if false {
                                                            v41 := GenS2{F0: 44, F1: 2147483647}
                                                            var v42 int8 = (-(v3))
                                                            v43 := GenS3{F0: 2, F1: 42}
                                                            _ = v41
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                    case "hello":
                                                        var v44 string = ("foo bar" + "hello")
                                                        _ = v44
                                                    default:
                                                        v29 = v31
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = fi28
                                                    }
                                                    if v30 {
                                                        if (857 > 65535) {
                                                            v31 -= (v29 * v29)
                                                        } else {
                                                            v3 -= v3
                                                            v45 := GenS2{F0: 260, F1: 42}
                                                            if v30 {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v5 += v24
                                                                    var v46 uint64 = v31
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v47 := GenS6{F0: 0}
                                                                        _ = v47
                                                                    case "hello":
                                                                        v22 = v4
                                                                    case "\n\t":
                                                                        v29 += (v46 / ((v31) | 1))
                                                                        v24 -= v5
                                                                        v24 -= v24
                                                                    default:
                                                                        v30 = v30
                                                                    }
                                                                    _ = v46
                                                                case "":
                                                                    v22 = (v4 / -2.25)
                                                                default:
                                                                    fmt.Println(fi27, v29)
                                                                }
                                                                v48 := GenS7{F0: v24}
                                                                switch (fi27 - -9223372036854775808) {
                                                                case -9223372036854775808:
                                                                    v2["a"] = 2147483647
                                                                case -542:
                                                                    fmt.Println(v29)
                                                                case 128:
                                                                    v31 = v29
                                                                default:
                                                                    if (false || ("\n\t" > "\n\t")) {
                                                                        if v30 {
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                if v30 {
                                                                                    v22 = 3.14159
                                                                                    v3 = (-(v3))
                                                                                    if (!(!v30)) {
                                                                                        var v49 int8 = v3
                                                                                        _ = len(v2)
                                                                                        _ = v49
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v50 int16 = 0
                                                                            v51 := GenS2{F0: 32767, F1: 270}
                                                                            switch "a" {
                                                                            case "":
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                            _ = v50
                                                                            _ = v51
                                                                        } else {
                                                                            v30 = v30
                                                                            _ = len(v2)
                                                                            if (!false) {
                                                                                var v52 uint64 = v31
                                                                                v4 = (v4 / (3.14159 / (0.0 * v22)))
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                        if v30 {
                                                                            v2["a"] = 124
                                                                            var v53 int16 = -1
                                                                            switch "" {
                                                                            case "hello":
                                                                                v54 := GenS3{F0: 65535, F1: v24}
                                                                                _ = v54
                                                                            case "foo bar":
                                                                                var v55 int64 = 469
                                                                                _ = v55
                                                                            default:
                                                                                v56 := GenS7{F0: v5}
                                                                                _ = v56
                                                                            }
                                                                            _ = v53
                                                                        }
                                                                    } else {
                                                                        var v57 float32 = (3.14159 * -2.25)
                                                                        _ = v57
                                                                    }
                                                                }
                                                                _ = v48
                                                            }
                                                            _ = v45
                                                        }
                                                        v30 = v30
                                                    }
                                                }
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            case "\n\t":
                                                if ((42 == 0) && true) {
                                                    var v58 int = (-(fi27))
                                                    _ = v58
                                                } else {
                                                    v3 += (^(v3))
                                                    v59 := GenS8{F0: 1}
                                                    _ = v59
                                                }
                                            }
                                            delete(v2, "hello")
                                            _ = fi28
                                        }
                                        var v60 float32 = float32(0.1)
                                        if (42 >= -9223372036854775808) {
                                            if true {
                                                var v61 bool = ("foo bar" <= (("a" + "foo bar") + ("" + ((("foo bar" + "") + "foo bar") + "foo bar"))))
                                                var v62 uint = 255
                                                _ = v61
                                                _ = v62
                                            } else {
                                                for fi63 := 0; fi63 < 20; fi63++ {
                                                    var v64 uint8 = 42
                                                    _ = fi63
                                                    _ = v64
                                                }
                                            }
                                            v22 += float32(float64(v4))
                                            for fi65 := 0; fi65 < 20; fi65++ {
                                                v66 := v2[("" + ("\n\t" + "\n\t"))]
                                                fmt.Println(v22, fi65)
                                                v3 -= 127
                                                if (v4 == -2.25) {
                                                    switch -368 {
                                                    case 4294967295:
                                                        if false {
                                                            var v67 bool = false
                                                            if (443 < 9223372036854775807) {
                                                                var v68 bool = v67
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v69 float64 = float64((((1e10 + v22) / v4) * v22))
                                                                    fmt.Println(v5, v3, v5)
                                                                    v70 := GenS3{F0: 42, F1: v24}
                                                                    _ = v69
                                                                    _ = v70
                                                                case "hello":
                                                                    if v68 {
                                                                        if (-167 != 1000) {
                                                                            fmt.Println(fi65, v22)
                                                                            v71 := GenS4{F0: "foo bar"}
                                                                            _ = v71
                                                                        }
                                                                        v60 = v60
                                                                        v5 = (v66 + 1)
                                                                        var v72 int = (fi27 / ((fi27) | 1))
                                                                        v73 := v2[(("foo bar" + "foo bar") + ("" + (("foo bar" + "\n\t") + "foo bar")))]
                                                                        fmt.Println(v4, v24)
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                case "":
                                                                    fmt.Println(v22, v66, v4)
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                                switch "foo bar" {
                                                                case "":
                                                                    if v68 {
                                                                        if (!false) {
                                                                            var v74 int8 = v3
                                                                            var v75 float32 = v22
                                                                            _ = v74
                                                                            _ = v75
                                                                        } else {
                                                                            var v76 int8 = (v3 >> ((4294967295) & 63))
                                                                            if v67 {
                                                                                v76 = (v76 % ((int8(v24)) | 1))
                                                                                v77 := GenS9{F0: 65535, F1: 1, F2: 728}
                                                                                _ = v77
                                                                            } else {
                                                                                fmt.Println(v3, v4, v3)
                                                                                v78 := GenS10{F0: 2}
                                                                                fmt.Println(v24, fi65)
                                                                                _ = v78
                                                                            }
                                                                            var v79 int8 = 0
                                                                            _ = v76
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                case "a":
                                                                    v4 -= (3.14159 * float32(float64(-2.25)))
                                                                default:
                                                                    var v80 float32 = float32(((1e10 + float64(v60)) + (-2.25 * 1.5)))
                                                                    _ = v80
                                                                }
                                                                _ = v68
                                                            } else {
                                                                var v81 int8 = 2
                                                                if (4294967295 != v66) {
                                                                    var v82 bool = (!(!((1 == 70) || v67)))
                                                                    _ = v82
                                                                } else {
                                                                    var v83 int32 = 0
                                                                    fmt.Println(v24)
                                                                    _ = v83
                                                                }
                                                                switch "a" {
                                                                case "hello":
                                                                    if (v67 || v67) {
                                                                        fmt.Println(v5, v67, v24)
                                                                    } else {
                                                                        _ = len(v1)
                                                                        v84 := GenS11{F0: (!((v3 == (v3 - v81)) || v67)), F1: -9223372036854775808}
                                                                        v67 = false
                                                                        _ = v84
                                                                    }
                                                                case "a":
                                                                    fmt.Println(v66, v5)
                                                                case "":
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v85 := v2["a"]
                                                                        _ = v85
                                                                    }
                                                                }
                                                                _ = v81
                                                            }
                                                            v66 += (v5 ^ 657)
                                                            if false {
                                                                var v86 int16 = -32768
                                                                v1 = append(v1, (fi27 ^ 128))
                                                                v22 -= 3.14159
                                                                _ = v86
                                                            }
                                                            _ = v67
                                                        } else {
                                                            switch -553 {
                                                            case -162:
                                                                fmt.Println(v4, v22, v22)
                                                            case 32767:
                                                                v87 := v2[("a" + ("foo bar" + "\n\t"))]
                                                                _ = v87
                                                            default:
                                                                v88 := GenS5{F0: 233, F1: 2}
                                                                _ = v88
                                                            }
                                                            switch (^(fi65)) {
                                                            case 468:
                                                                v2["\n\t"] = (v5 + v24)
                                                            case 897:
                                                                fmt.Println(v24)
                                                            }
                                                        }
                                                    }
                                                }
                                                v89 := GenS2{F0: 2, F1: 0}
                                                _ = fi65
                                                _ = v66
                                                _ = v89
                                            }
                                        }
                                        _ = v60
                                    }
                                    _ = fi27
                                }
                                if false {
                                    v90 := GenS11{F0: true, F1: 1}
                                    fmt.Println(v3, v3)
                                    _ = v90
                                }
                            }
                        }
                        _ = len(v2)
                        _ = v24
                    }
                    _ = v22
                }
            }
        case "a":
            switch 979 {
            case 42:
                fmt.Println(v5)
            case 77:
                fmt.Println(v3, v3, v5)
            case 843:
                v5 -= (v5 | (v5 / ((2) | 1)))
            default:
                var v91 float32 = v4
                _ = v91
            }
        default:
            v92 := GenS8{F0: 128}
            _ = v92
        }
        switch ((("" + ((("\n\t" + ("hello" + "")) + "foo bar") + ("foo bar" + (("foo bar" + "foo bar") + (("hello" + "") + ""))))) + "a") + ("hello" + (((("a" + "\n\t") + "hello") + (("" + ("" + "hello")) + ("foo bar" + ""))) + "foo bar"))) {
        case "hello":
            v93 := GenS10{F0: 2}
            _ = v93
        default:
            for fi94 := 0; fi94 < 19; fi94++ {
                fmt.Println(v4)
                var v95 int16 = -189
                var v96 uint32 = (^(v5))
                _ = fi94
                _ = v95
                _ = v96
            }
        }
        var v97 int32 = 648
        _ = v97
    }
    _ = len(v1)
    switch 32767 {
    case -110:
        switch ("hello" + "") {
        case "foo bar":
            for fi98 := 0; fi98 < 17; fi98++ {
                var v99 uint32 = v5
                switch (fi98 - (fi98 & fi98)) {
                case 32767:
                    var v100 uint64 = 2
                    _ = v100
                default:
                    for fi101 := 0; fi101 < 4; fi101++ {
                        if ("\n\t" <= "a") {
                            fmt.Println(v5, v5, v4)
                            var v102 uint64 = 4294967295
                            v3 += -1
                            _ = v102
                        }
                        var v103 int8 = (^(v3))
                        var v104 int = (fi101 & (fi101 & fi98))
                        _ = fi101
                        _ = v103
                        _ = v104
                    }
                }
                _ = fi98
                _ = v99
            }
        case "hello":
            for fi105 := 0; fi105 < 7; fi105++ {
                v4 += v4
                v5 += v5
                _ = fi105
            }
        }
    case 32768:
        v106 := GenS12{F0: 127, F1: 9223372036854775807}
        _ = v106
    }
    for fi107 := 0; fi107 < 1; fi107++ {
        v108 := v2["a"]
        for fi109 := 0; fi109 < 5; fi109++ {
            var v110 int64 = -956
            _ = fi109
            _ = v110
        }
        var v111 float32 = v4
        _ = fi107
        _ = v108
        _ = v111
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 += v5
    var v112 float64 = -2.25
    v113 := GenS9{F0: 228, F1: -128, F2: 42}
    var v114 bool = (0 > (^(v5)))
    if v114 {
        v5 = v5
    } else {
        var v115 uint32 = 866
        switch ("" + "foo bar") {
        case "\n\t":
            v116 := GenS7{F0: v115}
            _ = v116
        case "foo bar":
            if (1e10 <= v112) {
                fmt.Println(v114)
                v117 := v2["foo bar"]
                _ = v117
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        fmt.Println(v115, v115, v4)
        _ = v115
    }
    if ((-1 != -9223372036854775808) && (!((v114 && (!(v114 && (-128 < v3)))) && true))) {
        var v118 uint32 = v5
        for fi119 := 0; fi119 < 11; fi119++ {
            if (!(!v114)) {
                var v120 int16 = 42
                v121 := GenS13{F0: fi119}
                for fi122 := 0; fi122 < 7; fi122++ {
                    var v123 float64 = (v112 / v112)
                    _ = fi122
                    _ = v123
                }
                _ = v120
                _ = v121
            } else {
                v1 = append(v1, (fi119 % (((-(fi119))) | 1)))
            }
            _ = fi119
        }
        for fi124 := 0; fi124 < 1; fi124++ {
            if false {
                var v125 int64 = -21
                _ = v125
            }
            v1 = append(v1, (fi124 ^ (fi124 >> ((1) & 63))))
            _ = fi124
        }
        _ = v118
    } else {
        _ = len(v1)
        var v126 uint16 = uint16(v5)
        _ = len(v2)
        _ = v126
    }
    if (-1 == 1) {
        v3 += 127
        v127 := GenS14{F0: 1, F1: 0}
        for fi128 := 0; fi128 < 7; fi128++ {
            var v129 bool = (1 > int32(fi128))
            fmt.Println(v112, v5, v114)
            _ = fi128
            _ = v129
        }
        _ = v127
    }
    if false {
        v130 := GenS2{F0: 71, F1: 42}
        var v131 int16 = -32768
        var v132 int16 = v131
        _ = v130
        _ = v131
        _ = v132
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v112
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
