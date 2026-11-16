package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint; F1 int64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint64; F1 string }
type GenS18 struct { F0 int8; F1 uint16 }
type GenS19 struct { F0 bool; F1 int32; F2 int64 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 bool; F1 int64; F2 uint32 }
func genFn0(p0 int8, p1 uint) uint16 {
    p0 += (p0 >> ((p1) & 63))
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 255, F1: (fi4 | int(p1)), F2: p0}
        _ = fi4
        _ = v5
    }
    switch -454 {
    case 4294967295:
        p1 = 65535
    case -9223372036854775808:
        v6 := []int{1, 42, -726, 20}
        _ = v6
    case 9223372036854775807:
        if false {
            p0 += (p0 % ((10) | 1))
            switch ("a" + ("\n\t" + "hello")) {
            case "":
                for fi7 := 0; fi7 < 6; fi7++ {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 int32 = 42
                        if true {
                            v10 := GenS1{F0: 2147483648, F1: 4294967295, F2: p0}
                            _ = v10
                        } else {
                            v11 := GenS2{F0: 9223372036854775807}
                            _ = v11
                        }
                        if ("hello" == ("hello" + "")) {
                            v12 := GenS3{F0: 344, F1: 255}
                            p0 += (-(p0))
                            _ = v12
                        } else {
                            if true {
                                v13 := []int{fi7}
                                switch "foo bar" {
                                case "a":
                                    switch "\n\t" {
                                    case "a":
                                        if (0.0 == 0.1) {
                                            p0 = (p0 * -28)
                                        } else {
                                            switch (("a" + "\n\t") + "\n\t") {
                                            case "":
                                                p1 += (^(p1))
                                            case "\n\t":
                                                var v14 bool = (p0 == (^(p0)))
                                                _ = v14
                                            case "hello":
                                                var v15 int16 = 0
                                                _ = v15
                                            }
                                            v13 = append(v13, 42)
                                        }
                                    }
                                    var v16 uint = 827
                                    var v17 bool = ((("\n\t" + (((((("\n\t" + "hello") + "foo bar") + ("hello" + "\n\t")) + "foo bar") + "foo bar") + "hello")) + "\n\t") > "\n\t")
                                    _ = v16
                                    _ = v17
                                case "":
                                    var v18 float32 = 0.0
                                    _ = v18
                                }
                                var v19 string = "a"
                                _ = v13
                                _ = v19
                            } else {
                                var v20 int = 127
                                _ = v20
                            }
                        }
                        _ = fi8
                        _ = v9
                    }
                    _ = fi7
                }
            case "\n\t":
                if true {
                    fmt.Println(p1, p1)
                } else {
                    v21 := GenS2{F0: 441}
                    v22 := GenS2{F0: 42}
                    _ = v21
                    _ = v22
                }
            case "hello":
                v23 := []int{2147483647, -350, 442}
                _ = v23
            }
            var v24 uint16 = 0
            _ = v24
        } else {
            p1 -= 667
        }
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 int16 = 1
    v2 := []uint32{42}
    var v3 uint64 = 2
    if false {
        if (1.5 > -2.25) {
            for fi25 := 0; fi25 < 18; fi25++ {
                for fi26 := 0; fi26 < 3; fi26++ {
                    if false {
                        var v27 uint = 255
                        var v28 bool = ((!false) || (true || true))
                        if true {
                            if v28 {
                                v3 += 756
                                var v29 uint64 = (v3 << (((v27 + (v27 & (v27 + (v27 ^ v27))))) & 63))
                                v28 = ((-96 != 42) && true)
                                _ = v29
                            }
                        } else {
                            if false {
                                if (147 != 42) {
                                    fmt.Println(fi26)
                                    _ = len(v2)
                                }
                                switch "hello" {
                                case "foo bar":
                                    v30 := GenS3{F0: 2147483647, F1: 255}
                                    _ = v30
                                }
                                if (v28 || true) {
                                    switch "\n\t" {
                                    case "":
                                        switch "foo bar" {
                                        case "hello":
                                            v27 += v27
                                        case "foo bar":
                                            if (65535 >= v27) {
                                                switch "hello" {
                                                case "":
                                                    var v31 int = 0
                                                    _ = v31
                                                case "a":
                                                    switch 9223372036854775807 {
                                                    case 347:
                                                        switch (^(fi25)) {
                                                        case 127:
                                                            _ = len(v2)
                                                        }
                                                    case 624:
                                                        if (fi26 != (fi26 >> ((v27) & 63))) {
                                                            v3 = v3
                                                        }
                                                    case 42:
                                                        var v32 int16 = 32767
                                                        _ = v32
                                                    default:
                                                        var v33 uint64 = (v3 << ((0) & 63))
                                                        v3 = (^(v33))
                                                        var v34 int8 = 75
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                case "foo bar":
                                                    v1 -= (v1 / ((v1) | 1))
                                                }
                                                fmt.Println(v3, v27)
                                                v28 = v28
                                                var v35 int64 = 0
                                                _ = v35
                                            }
                                        }
                                    case "foo bar":
                                        if (!v28) {
                                            v28 = ((fi26 >= 4294967295) && (!v28))
                                        } else {
                                            v1 = (v1 >> ((166) & 63))
                                        }
                                    case "\n\t":
                                        v36 := GenS4{F0: 255}
                                        _ = v36
                                    default:
                                        fmt.Println(v3, v1, v27)
                                    }
                                    fmt.Println(v3, fi25)
                                }
                            }
                            v37 := GenS3{F0: 42, F1: 49}
                            var v38 uint16 = 118
                            _ = v37
                            _ = v38
                        }
                        _ = v27
                        _ = v28
                    } else {
                        if (609 > 1) {
                            var v39 uint = 65535
                            _ = v39
                        }
                    }
                    var v40 int16 = (v1 / ((19) | 1))
                    _ = fi26
                    _ = v40
                }
                v1 += v1
                _ = fi25
            }
            v41 := GenS5{F0: (v1 % ((-38) | 1)), F1: 2}
            _ = v41
        }
        v1 = -492
        for fi42 := 0; fi42 < 8; fi42++ {
            if (-828 <= 0) {
                v43 := GenS1{F0: 2, F1: fi42, F2: 100}
                _ = v43
            }
            var v44 int64 = 42
            _ = fi42
            _ = v44
        }
    }
    var v45 int32 = 2147483647
    switch 348 {
    case -989:
        _ = len(v2)
    case 1:
        _ = len(v2)
    default:
        var v46 string = "a"
        _ = v46
    }
    if (v45 <= (v45 | 0)) {
        var v47 int16 = v1
        _ = v47
    }
    if (("\n\t" + "") > "\n\t") {
        var v48 int8 = 42
        _ = v48
    } else {
        _ = len(v2)
    }
    v49 := GenS6{F0: (genFn0(100, 4294967295) < 105)}
    switch ("\n\t" + "hello") {
    case "hello":
        for fi50 := 0; fi50 < 11; fi50++ {
            v51 := GenS7{F0: v3, F1: (fi50 - fi50)}
            v52 := GenS8{F0: (v3 % ((v3) | 1)), F1: 1000}
            if len(v2) > 0 {
                v2[0] = 896
            }
            _ = fi50
            _ = v51
            _ = v52
        }
    default:
        var v53 int64 = -1
        _ = v53
    }
    for fi54 := 0; fi54 < 16; fi54++ {
        for fi55 := 0; fi55 < 6; fi55++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 = int16(v45)
            v56 := GenS7{F0: v3, F1: 127}
            _ = fi55
            _ = v56
        }
        _ = fi54
    }
    v1 = v1
    if false {
        if false {
            var v57 uint16 = 65535
            _ = v57
        } else {
            v58 := GenS8{F0: 615, F1: v1}
            _ = v58
        }
    }
    var v59 uint8 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v60 bool = false
    switch 42 {
    case 4294967295:
        switch -2147483648 {
        case -1:
            v60 = v60
        case -9223372036854775808:
            if false {
                v61 := GenS9{F0: uint16(v59), F1: 42}
                v60 = (v45 >= v45)
                for fi62 := 0; fi62 < 10; fi62++ {
                    var v63 int64 = 991
                    _ = fi62
                    _ = v63
                }
                _ = v61
            } else {
                v45 = (^(v45))
            }
        case 1:
            var v64 int64 = 42
            _ = v64
        default:
            var v65 int32 = v45
            _ = v65
        }
    case -1:
        v60 = false
    case -932:
        if (-1 <= -53) {
            for fi66 := 0; fi66 < 8; fi66++ {
                if (v1 <= -966) {
                    var v67 string = (("\n\t" + "\n\t") + "foo bar")
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v67
                } else {
                    v2 = append(v2, 65535)
                    _ = len(v2)
                    if true {
                        for fi68 := 0; fi68 < 14; fi68++ {
                            v69 := GenS3{F0: 2147483648, F1: 127}
                            _ = fi68
                            _ = v69
                        }
                        switch "hello" {
                        case "":
                            v1 -= v1
                        case "a":
                            if (!false) {
                                v1 = (v1 ^ v1)
                                var v70 bool = ((v1 % ((v1) | 1)) != v1)
                                _ = v70
                            } else {
                                v71 := GenS8{F0: (^(v3)), F1: (v1 - -177)}
                                for fi72 := 0; fi72 < 11; fi72++ {
                                    if v60 {
                                        switch fi72 {
                                        case -1:
                                            var v73 int8 = -1
                                            _ = v73
                                        default:
                                            v74 := GenS10{F0: 2}
                                            _ = v74
                                        }
                                    } else {
                                        switch (("hello" + ("\n\t" + "\n\t")) + "foo bar") {
                                        case "a":
                                            var v75 int8 = 100
                                            _ = v75
                                        case "foo bar":
                                            var v76 int64 = -1
                                            _ = v76
                                        }
                                    }
                                    if v60 {
                                        var v77 int = fi72
                                        _ = v77
                                    }
                                    _ = fi72
                                }
                                _ = v71
                            }
                        }
                    } else {
                        v45 += -2147483648
                        var v78 int8 = 1
                        v78 = (^(v78))
                        _ = v78
                    }
                }
                fmt.Println(v60, v3, fi66)
                v79 := GenS11{F0: 42, F1: 4294967295}
                _ = fi66
                _ = v79
            }
            if v60 {
                for fi80 := 0; fi80 < 16; fi80++ {
                    if v60 {
                        if (v1 >= v1) {
                            v1 = (v1 | (^(v1)))
                            for fi81 := 0; fi81 < 2; fi81++ {
                                v82 := GenS3{F0: 65, F1: (v59 & v59)}
                                _ = fi81
                                _ = v82
                            }
                            v1 -= 42
                        } else {
                            v1 += -32768
                            if (-2.25 >= (((0.0 - -2.25) * -2.25) * 0.0)) {
                                v45 = v45
                            } else {
                                v2 = append(v2, 0)
                                _ = len(v2)
                            }
                            if v60 {
                                var v83 int = fi80
                                v59 += 16
                                _ = v83
                            } else {
                                var v84 float64 = 3.14159
                                _ = v84
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = 60
                        }
                    } else {
                        if (v60 && (v59 == 1)) {
                            if (0 >= (fi80 / (((fi80 << ((18446744073709551615) & 63))) | 1))) {
                                v85 := GenS12{F0: "foo bar"}
                                v86 := GenS5{F0: -1, F1: 2}
                                _ = v85
                                _ = v86
                            }
                            v3 = 2
                        } else {
                            fmt.Println(v1)
                            for fi87 := 0; fi87 < 3; fi87++ {
                                v2 = append(v2, 65535)
                                if false {
                                    if true {
                                        v60 = v60
                                    } else {
                                        var v88 uint = 1
                                        v59 -= (v59 >> (((v88 % (((v88 & (v88 ^ 2))) | 1))) & 63))
                                        _ = v88
                                    }
                                    v1 += -644
                                    switch fi80 {
                                    case 4294967295:
                                        _ = len(v2)
                                    default:
                                        if ((v3 - (^(v3))) >= (v3 + (^(v3)))) {
                                            _ = len(v2)
                                        } else {
                                            switch "a" {
                                            case "foo bar":
                                                if ((((-2.25 + 0.0) * 0.1) == 0.1) && v60) {
                                                    if (false && true) {
                                                        if v60 {
                                                            var v89 int8 = 1
                                                            v3 += (^(v3))
                                                            _ = v89
                                                        }
                                                        switch (fi87 / ((fi80) | 1)) {
                                                        case 42:
                                                            switch "a" {
                                                            case "foo bar":
                                                                if (true || v60) {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                switch ("a" + "a") {
                                                                case "hello":
                                                                    if (uint8(v3) != (v59 % ((v59) | 1))) {
                                                                        fmt.Println(fi87, fi87)
                                                                        var v90 int64 = -1
                                                                        v60 = (((false && ((!((((float64((0.0 / 3.14159)) + ((1e10 * -2.25) / -2.25)) >= 0.0) || false) || (int16(v45) < (^(v1))))) || ("" >= "foo bar"))) || (255 >= uint8(v3))) && (32767 > v1))
                                                                        _ = v90
                                                                    } else {
                                                                        fmt.Println(v1, v45)
                                                                    }
                                                                case "a":
                                                                    switch ("" + "foo bar") {
                                                                    case "hello":
                                                                        var v91 int8 = 103
                                                                        _ = v91
                                                                    }
                                                                    if v60 {
                                                                        var v92 int8 = -128
                                                                        _ = v92
                                                                    } else {
                                                                        switch fi87 {
                                                                        case -413:
                                                                            v2 = append(v2, 0)
                                                                        }
                                                                        fmt.Println(fi80)
                                                                        switch fi80 {
                                                                        case -1:
                                                                            switch (-(fi80)) {
                                                                            case 42:
                                                                                fmt.Println(v3)
                                                                            case -237:
                                                                                var v93 float64 = (1.5 - 0.0)
                                                                                _ = v93
                                                                            }
                                                                        case 841:
                                                                            v45 += v45
                                                                        case -627:
                                                                            v45 -= (v45 / (((v45 | (v45 ^ (v45 ^ v45)))) | 1))
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 0
                                                                    }
                                                                }
                                                            case "hello":
                                                                var v94 int64 = 0
                                                                _ = v94
                                                            }
                                                            if v60 {
                                                                fmt.Println(v1, fi87)
                                                                var v95 uint32 = 2147483648
                                                                _ = v95
                                                            } else {
                                                                if (915 <= 42) {
                                                                    _ = len(v2)
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 1
                                                                    }
                                                                    v1 = v1
                                                                } else {
                                                                    var v96 bool = v60
                                                                    _ = v96
                                                                }
                                                                v2 = append(v2, 93)
                                                                if (v60 || (1 != (fi80 + fi87))) {
                                                                    v97 := GenS6{F0: (true || (!v60))}
                                                                    _ = v97
                                                                }
                                                            }
                                                        default:
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case "a":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            default:
                                                                v59 = v59
                                                            }
                                                        }
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v98 uint8 = v59
                                                    _ = v98
                                                }
                                            case "":
                                                _ = len(v2)
                                                v45 = v45
                                                v99 := GenS11{F0: 42, F1: (v3 << ((56) & 63))}
                                                fmt.Println(v1)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v99
                                            default:
                                                if v60 {
                                                    if len(v2) > 0 {
                                                        v2[0] = 2
                                                    }
                                                    var v100 uint = 724
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch "" {
                                                        case "a":
                                                            if true {
                                                                if (!false) {
                                                                    var v101 int32 = (-(v45))
                                                                    fmt.Println(v1)
                                                                    _ = v101
                                                                }
                                                                fmt.Println(v1)
                                                                if (v60 || true) {
                                                                    if v60 {
                                                                        v2 = append(v2, 2)
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 574
                                                                        }
                                                                        v102 := GenS8{F0: (^(v3)), F1: v1}
                                                                        if v60 {
                                                                            v103 := GenS13{F0: 255}
                                                                            var v104 uint32 = 65535
                                                                            _ = len(v2)
                                                                            _ = v103
                                                                            _ = v104
                                                                        }
                                                                        _ = v102
                                                                    }
                                                                    fmt.Println(fi80, fi80, fi87)
                                                                    if (-201 <= -1) {
                                                                        if ((v59 << ((255) & 63)) > v59) {
                                                                            v59 -= (v59 + (^(v59)))
                                                                        } else {
                                                                            v105 := GenS9{F0: 0, F1: 88}
                                                                            var v106 int64 = int64(v3)
                                                                            var v107 int = (fi87 | fi87)
                                                                            v45 = (-(v45))
                                                                            fmt.Println(v59, v1, fi87)
                                                                            _ = v105
                                                                            _ = v106
                                                                            _ = v107
                                                                        }
                                                                        fmt.Println(v59, v100)
                                                                        v59 -= v59
                                                                    }
                                                                }
                                                            }
                                                        case "\n\t":
                                                            if len(v2) > 0 {
                                                                v2[0] = 4294967295
                                                            }
                                                        case "hello":
                                                            v60 = false
                                                        default:
                                                            v100 -= v100
                                                        }
                                                    case "a":
                                                        var v108 uint16 = 32767
                                                        _ = v108
                                                    case "hello":
                                                        if (!false) {
                                                            switch "a" {
                                                            case "":
                                                                v60 = v60
                                                            default:
                                                                if v60 {
                                                                    if (0.1 >= 1.5) {
                                                                        var v109 float64 = 3.14159
                                                                        v45 -= (v45 + -600)
                                                                        _ = v109
                                                                    } else {
                                                                        v110 := GenS5{F0: (^(v1)), F1: v100}
                                                                        _ = v110
                                                                    }
                                                                    switch "hello" {
                                                                    case "":
                                                                        v111 := GenS7{F0: 1, F1: (fi80 + 0)}
                                                                        _ = v111
                                                                    case "\n\t":
                                                                        var v112 uint8 = v59
                                                                        _ = v112
                                                                    case "foo bar":
                                                                        v45 = -2147483648
                                                                    }
                                                                } else {
                                                                    v113 := GenS14{F0: v100, F1: 19}
                                                                    switch (("\n\t" + "hello") + ("a" + "a")) {
                                                                    case "hello":
                                                                        var v114 float64 = (-2.25 * float64(0.1))
                                                                        _ = v114
                                                                    default:
                                                                        var v115 float32 = 1.5
                                                                        _ = v115
                                                                    }
                                                                    _ = v113
                                                                }
                                                            }
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = 4294967295
                                                            }
                                                        }
                                                    }
                                                    v116 := GenS15{F0: -32768}
                                                    var v117 uint64 = v3
                                                    if false {
                                                        fmt.Println(v45)
                                                        switch "foo bar" {
                                                        case "":
                                                            var v118 uint8 = (^(v59))
                                                            _ = v118
                                                        case "foo bar":
                                                            _ = len(v2)
                                                        }
                                                        v117 = 42
                                                    } else {
                                                        v119 := GenS12{F0: "a"}
                                                        v2 = append(v2, 2147483648)
                                                        v117 += v3
                                                        _ = v119
                                                    }
                                                    _ = v100
                                                    _ = v116
                                                    _ = v117
                                                }
                                            }
                                        }
                                    }
                                    var v120 int8 = 100
                                    _ = v120
                                } else {
                                    v121 := GenS4{F0: (^(v59))}
                                    v45 = -1
                                    v1 -= (v1 + (^(v1)))
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    _ = v121
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi87
                            }
                            var v122 int32 = v45
                            _ = v122
                        }
                        if v60 {
                            for fi123 := 0; fi123 < 14; fi123++ {
                                v59 += v59
                                _ = fi123
                            }
                        } else {
                            for fi124 := 0; fi124 < 7; fi124++ {
                                v45 += (-(v45))
                                v45 += (v45 * v45)
                                _ = fi124
                            }
                            var v125 int16 = (^(v1))
                            _ = v125
                        }
                    }
                    if (!((v60 && (42 >= 0)) && v60)) {
                        v126 := GenS3{F0: 42, F1: v59}
                        if (((42 < 104) && (!(v60 || v60))) || v60) {
                            v60 = (("foo bar" == "") && (!v60))
                            var v127 uint16 = genFn0(42, 250)
                            _ = v127
                        } else {
                            _ = len(v2)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if (v1 >= (v1 % ((-1) | 1))) {
                                if v60 {
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (!(!false)) {
                                            fmt.Println(v45)
                                        } else {
                                            v59 -= (v59 % ((v59) | 1))
                                        }
                                    default:
                                        for fi128 := 0; fi128 < 6; fi128++ {
                                            var v129 float64 = float64(((1e10 / -2.25) + 1e10))
                                            if (v60 || (!(!(0 == -261)))) {
                                                var v130 uint64 = (v3 - v3)
                                                if v60 {
                                                    if true {
                                                        var v131 uint = 836
                                                        _ = len(v2)
                                                        _ = v131
                                                    } else {
                                                        v60 = (v45 > 0)
                                                        fmt.Println(fi80, fi80, fi128)
                                                        var v132 uint16 = 300
                                                        v133 := GenS14{F0: 874, F1: 1}
                                                        var v134 uint8 = 84
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    }
                                                } else {
                                                    if (fi80 < fi80) {
                                                        v2 = append(v2, 2)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                                _ = v130
                                            } else {
                                                var v135 uint64 = 415
                                                var v136 float32 = 1e10
                                                switch "hello" {
                                                case "":
                                                    var v137 bool = (fi80 <= fi128)
                                                    _ = v137
                                                }
                                                _ = v135
                                                _ = v136
                                            }
                                            _ = fi128
                                            _ = v129
                                        }
                                    }
                                    var v138 string = ((("\n\t" + ((("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))) + "a") + ("a" + "foo bar"))) + "") + "")
                                    for fi139 := 0; fi139 < 7; fi139++ {
                                        if (65535 < 2) {
                                            fmt.Println(fi139)
                                            if (129 > 849) {
                                                v2 = append(v2, 1)
                                                switch 128 {
                                                case 1:
                                                    if true {
                                                        v140 := GenS16{F0: fi80}
                                                        v138 = ("hello" + v138)
                                                        var v141 int8 = -1
                                                        _ = v140
                                                        _ = v141
                                                    } else {
                                                        v1 = v1
                                                        v1 += v1
                                                        var v142 int32 = v45
                                                        _ = v142
                                                    }
                                                case 390:
                                                    v2 = append(v2, 2147483647)
                                                    switch v138 {
                                                    case "\n\t":
                                                        var v143 int = (-(fi80))
                                                        _ = v143
                                                    default:
                                                        var v144 float32 = (((1e10 * 3.14159) * 0.0) * ((3.14159 / 3.14159) * 0.0))
                                                        _ = v144
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = 2
                                                    }
                                                    switch v138 {
                                                    case "":
                                                        v3 = 42
                                                        if (v60 || (!true)) {
                                                            if (!v60) {
                                                                var v145 int64 = 0
                                                                _ = v145
                                                            }
                                                        } else {
                                                            var v146 int64 = -9223372036854775808
                                                            switch "a" {
                                                            case "\n\t":
                                                                v60 = true
                                                            case "a":
                                                                if (3.14159 >= 1e10) {
                                                                    fmt.Println(v1)
                                                                    v147 := GenS17{F0: v3, F1: v138}
                                                                    v148 := GenS18{F0: 100, F1: 544}
                                                                    _ = v147
                                                                    _ = v148
                                                                } else {
                                                                    fmt.Println(v45, fi80, fi80)
                                                                    if v60 {
                                                                        switch v138 {
                                                                        case "":
                                                                            v149 := GenS19{F0: true, F1: (v45 ^ (v45 % (((-(v45))) | 1))), F2: 42}
                                                                            _ = v149
                                                                        case "foo bar":
                                                                            var v150 int16 = v1
                                                                            _ = v150
                                                                        }
                                                                        v2 = append(v2, 571)
                                                                    }
                                                                }
                                                                var v151 int16 = -1
                                                                v138 = v138
                                                                var v152 uint64 = (v3 / ((0) | 1))
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v59 = 128
                                                                case "hello":
                                                                    v153 := GenS1{F0: 42, F1: int(v146), F2: -17}
                                                                    _ = v153
                                                                case "":
                                                                    var v154 float64 = 1e10
                                                                    _ = v154
                                                                }
                                                                _ = v151
                                                                _ = v152
                                                            }
                                                            v146 -= (v146 * (v146 % ((-1) | 1)))
                                                            v3 -= (v3 | (^(v3)))
                                                            fmt.Println(v146, v146, v45)
                                                            v59 = v59
                                                            _ = v146
                                                        }
                                                    default:
                                                        v2 = append(v2, 121)
                                                        v60 = (65535 == uint32(v3))
                                                        var v155 int16 = (v1 >> ((468) & 63))
                                                        v2 = append(v2, 2)
                                                        _ = v155
                                                    }
                                                case 9223372036854775807:
                                                    v59 -= 42
                                                }
                                                var v156 int16 = (v1 * v1)
                                                var v157 uint64 = v3
                                                _ = v156
                                                _ = v157
                                            } else {
                                                if (v60 && (false || false)) {
                                                    switch v138 {
                                                    case "":
                                                        if ((-2.25 + ((3.14159 * 0.1) + ((1e10 / float64((1e10 - -2.25))) / 1.5))) <= (0.0 - 1e10)) {
                                                            var v158 uint = 18446744073709551615
                                                            var v159 uint64 = (v3 | v3)
                                                            v160 := GenS16{F0: fi139}
                                                            _ = v158
                                                            _ = v159
                                                            _ = v160
                                                        }
                                                    case "hello":
                                                        v2 = append(v2, 386)
                                                    case "foo bar":
                                                        v161 := GenS18{F0: 88, F1: 42}
                                                        _ = v161
                                                    default:
                                                        v162 := GenS20{F0: 42}
                                                        _ = v162
                                                    }
                                                    var v163 string = (v138 + v138)
                                                    v164 := GenS21{F0: 2}
                                                    _ = v163
                                                    _ = v164
                                                }
                                                v165 := GenS7{F0: 305, F1: 0}
                                                _ = v165
                                            }
                                        } else {
                                            _ = len(v2)
                                            v1 -= (v1 - (v1 << ((407) & 63)))
                                        }
                                        _ = fi139
                                    }
                                    _ = v138
                                } else {
                                    var v166 uint32 = 1
                                    switch "" {
                                    case "hello":
                                        for fi167 := 0; fi167 < 15; fi167++ {
                                            v1 -= 42
                                            fmt.Println(fi80)
                                            var v168 uint32 = 1
                                            _ = fi167
                                            _ = v168
                                        }
                                    case "\n\t":
                                        v169 := GenS22{F0: 42}
                                        _ = v169
                                    default:
                                        fmt.Println(v59, v166, v166)
                                    }
                                    var v170 uint = 978
                                    _ = v166
                                    _ = v170
                                }
                                v171 := GenS4{F0: v59}
                                _ = v171
                            } else {
                                if (((int32(v59) > v45) && (2 != 819)) && ((v1 * 42) >= v1)) {
                                    for fi172 := 0; fi172 < 7; fi172++ {
                                        if v60 {
                                            v60 = (!v60)
                                            if true {
                                                v3 = 821
                                                v1 = v1
                                                var v173 int64 = int64(fi80)
                                                _ = v173
                                            } else {
                                                if (false || (false && false)) {
                                                    _ = len(v2)
                                                    if v60 {
                                                        switch (fi172 << ((4294967295) & 63)) {
                                                        case -1:
                                                            fmt.Println(v3)
                                                        case 9223372036854775807:
                                                            v2 = append(v2, 65535)
                                                        default:
                                                            var v174 uint = 42
                                                            _ = v174
                                                        }
                                                        switch ("a" + "foo bar") {
                                                        case "hello":
                                                            var v175 uint32 = 853
                                                            _ = v175
                                                        default:
                                                            if v60 {
                                                                switch "hello" {
                                                                case "a":
                                                                    v59 += 42
                                                                default:
                                                                    v176 := GenS17{F0: 1, F1: ""}
                                                                    _ = v176
                                                                }
                                                            } else {
                                                                if v60 {
                                                                    switch fi80 {
                                                                    case -406:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case 537:
                                                                        var v177 int16 = 749
                                                                        _ = v177
                                                                    default:
                                                                        v59 -= 2
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                } else {
                                                                    var v178 int = (fi80 & fi80)
                                                                    fmt.Println(v1, v1)
                                                                    var v179 int32 = -1
                                                                    _ = v178
                                                                    _ = v179
                                                                }
                                                                switch (fi80 + (-(fi80))) {
                                                                case 2147483647:
                                                                    if (1.5 >= ((3.14159 * 0.1) + 0.1)) {
                                                                        var v180 int = fi80
                                                                        _ = v180
                                                                    } else {
                                                                        if v60 {
                                                                            var v181 uint16 = 0
                                                                            v45 -= (v45 / ((-2147483648) | 1))
                                                                            _ = v181
                                                                        } else {
                                                                            v1 = (v1 | (v1 & v1))
                                                                            var v182 string = ("\n\t" + "a")
                                                                            _ = v182
                                                                        }
                                                                        v183 := GenS8{F0: v3, F1: 1000}
                                                                        if (!v60) {
                                                                            v3 = 18446744073709551615
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                        v59 = v59
                                                                        var v184 uint32 = 743
                                                                        _ = v183
                                                                        _ = v184
                                                                    }
                                                                case 894:
                                                                    v185 := GenS18{F0: -104, F1: uint16(v59)}
                                                                    _ = v185
                                                                }
                                                                var v186 float32 = 3.14159
                                                                _ = v186
                                                            }
                                                        }
                                                    }
                                                }
                                                var v187 uint = 65535
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v187
                                            }
                                            fmt.Println(fi172)
                                        }
                                        var v188 uint32 = 737
                                        _ = fi172
                                        _ = v188
                                    }
                                    var v189 float32 = (1e10 * float32(0.1))
                                    _ = v189
                                } else {
                                    v60 = v60
                                    var v190 bool = ((true || (!(true || false))) && (0 <= 1))
                                    _ = v190
                                }
                                var v191 float32 = 0.1
                                _ = v191
                            }
                        }
                        _ = v126
                    }
                    fmt.Println(v1)
                    _ = fi80
                }
            } else {
                var v192 uint64 = v3
                for fi193 := 0; fi193 < 18; fi193++ {
                    fmt.Println(v3, v1)
                    _ = fi193
                }
                var v194 bool = false
                _ = v192
                _ = v194
            }
        } else {
            v195 := GenS23{F0: true, F1: 1, F2: 994}
            _ = v195
        }
    default:
        v1 -= 32767
    }
    fmt.Println(v3)
    switch 127 {
    case 32767:
        var v196 int64 = 42
        _ = v196
    case 1:
        _ = len(v2)
    case -9223372036854775808:
        v59 -= (v59 >> ((0) & 63))
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v45 = (v45 + (v45 / (((v45 % (((v45 * -2147483648)) | 1))) | 1)))
    var v197 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v197:", v197)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v197
    _ = v2
    _ = v3
    _ = v45
    _ = v49
    _ = v59
    _ = v60
}
