package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS3 struct { F0 int64; F1 int8; F2 string }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 int; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint64; F2 uint }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 uint; F2 string }
type GenS12 struct { F0 int32; F1 int }
type GenS13 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS14 struct { F0 uint8; F1 uint16; F2 int32 }
func genFn0(p0 bool, p1 int8, p2 int16) string {
    if p0 {
        p1 = int8(p2)
        var v7 uint32 = 2147483648
        _ = v7
    }
    if p0 {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: (-(p2))}
            _ = fi8
            _ = v9
        }
        if p0 {
            p2 -= p2
            p1 = -128
        }
        fmt.Println(p1)
    }
    v10 := []int{32767, 1, 9223372036854775807, 113}
    var v11 float64 = float64(3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return "hello"
}
func genFn1(p0 uint16, p1 int) bool {
    if (!(!true)) {
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: 42}
            v14 := []int{fi12, fi12, p1}
            _ = fi12
            _ = v13
            _ = v14
        }
        if false {
            p0 += 798
            v15 := GenS1{F0: -32768}
            var v16 string = "foo bar"
            _ = v15
            _ = v16
        } else {
            switch (p1 - 127) {
            case 128:
                p1 -= 984
            default:
                v17 := []int{-2147483648, -199}
                _ = v17
            }
            if true {
                var v18 uint16 = p0
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 uint16 = (p0 ^ 0)
                    var v21 int = p1
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = v18
            }
            v22 := GenS2{F0: 2, F1: true, F2: 4294967295}
            _ = v22
        }
        if true {
            switch 127 {
            case 0:
                if (42 == 255) {
                    v23 := []int{p1}
                    _ = len(v23)
                    _ = v23
                }
            default:
                var v24 int8 = 100
                _ = v24
            }
            p1 = p1
        }
    } else {
        if true {
            v25 := []int{(p1 >> ((65535) & 63)), 628, (p1 * 4294967295), -128}
            if false {
                if len(v25) > 0 {
                    v25[0] = p1
                }
                var v26 float32 = 3.14159
                for fi27 := 0; fi27 < 17; fi27++ {
                    for fi28 := 0; fi28 < 13; fi28++ {
                        v26 = 0.0
                        p0 = 255
                        if len(v25) > 0 {
                            v25[0] = 127
                        }
                        if (2 < 42) {
                            var v29 int = 128
                            if false {
                                p1 = p1
                                var v30 uint16 = 765
                                _ = v30
                            } else {
                                if ((float64((((0.0 + v26) / v26) - (1.5 + (0.1 / ((v26 - float32((-2.25 - (0.0 + ((1.5 - 1e10) / (0.1 + 1.5)))))) * v26))))) / float64(v26)) != 1.5) {
                                    if false {
                                        if len(v25) > 0 {
                                            v25[0] = (fi27 | 2147483647)
                                        }
                                    } else {
                                        v26 += ((v26 - v26) * (((v26 - 1.5) * float32(3.14159)) - (1.5 - (-2.25 / v26))))
                                    }
                                    p0 += (p0 >> ((605) & 63))
                                    _ = len(v25)
                                } else {
                                    var v31 uint = 458
                                    fmt.Println(v29)
                                    _ = v31
                                }
                            }
                            _ = v29
                        } else {
                            v26 -= -2.25
                        }
                        p0 = 1
                        _ = fi28
                    }
                    var v32 int32 = -2147483648
                    switch "\n\t" {
                    case "foo bar":
                        var v33 float64 = ((1.5 * 0.0) * -2.25)
                        _ = v33
                    case "\n\t":
                        if (138 < 42) {
                            for fi34 := 0; fi34 < 11; fi34++ {
                                if len(v25) > 0 {
                                    v25[0] = 128
                                }
                                if (65535 == 0) {
                                    var v35 float32 = v26
                                    fmt.Println(v35)
                                    _ = v35
                                }
                                _ = fi34
                            }
                        } else {
                            switch p1 {
                            case 2147483647:
                                switch (^(p1)) {
                                case 2147483647:
                                    if false {
                                        var v36 uint16 = p0
                                        if true {
                                            switch (fi27 & -2147483648) {
                                            case -98:
                                                var v37 uint16 = (v36 | v36)
                                                _ = v37
                                            case 4294967295:
                                                if (1e10 >= (1e10 / 1.0)) {
                                                    var v38 int16 = 42
                                                    v38 += v38
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(v36, v26)
                                                    }
                                                    _ = v38
                                                }
                                            case 201:
                                                var v39 float32 = 0.0
                                                _ = v39
                                            }
                                            for fi40 := 0; fi40 < 16; fi40++ {
                                                _ = len(v25)
                                                switch ("\n\t" + "\n\t") {
                                                case "a":
                                                    fmt.Println(p0, p1, p1)
                                                default:
                                                    switch (fi40 & (fi27 / ((fi27) | 1))) {
                                                    case 0:
                                                        var v41 int32 = -73
                                                        _ = v41
                                                    default:
                                                        v42 := GenS2{F0: 453, F1: false, F2: 406}
                                                        _ = v42
                                                    }
                                                }
                                                p1 = 32768
                                                _ = fi40
                                            }
                                            for fi43 := 0; fi43 < 5; fi43++ {
                                                p1 += (fi27 / ((p1) | 1))
                                                fmt.Println(v32, v32)
                                                if true {
                                                    var v44 int32 = v32
                                                    _ = v44
                                                }
                                                v45 := GenS1{F0: -32768}
                                                if false {
                                                    p1 -= (fi43 % ((fi43) | 1))
                                                    switch fi27 {
                                                    case 4294967295:
                                                        var v46 int32 = v32
                                                        _ = v46
                                                    case 32767:
                                                        var v47 uint8 = 2
                                                        _ = v47
                                                    }
                                                    switch -416 {
                                                    case 127:
                                                        if len(v25) > 0 {
                                                            v25[0] = fi27
                                                        }
                                                    case 128:
                                                        switch ("a" + (("hello" + "hello") + ((("" + ("foo bar" + "hello")) + ("" + ("\n\t" + ""))) + ("\n\t" + "")))) {
                                                        case "":
                                                            p0 = v36
                                                        default:
                                                            v48 := GenS2{F0: 42, F1: (!(true || (((true && ((!(p1 < 1)) || true)) && false) && true))), F2: 2}
                                                            _ = v48
                                                        }
                                                    case -9223372036854775808:
                                                        if (1e10 >= ((3.14159 * 1e10) + -2.25)) {
                                                            v26 += v26
                                                        }
                                                    }
                                                } else {
                                                    v25 = append(v25, (fi27 << ((337) & 63)))
                                                    if len(v25) >= 2 {
                                                        v25 = v25[1:len(v25)]
                                                    }
                                                    fmt.Println(p0, fi43, p0)
                                                }
                                                _ = fi43
                                                _ = v45
                                            }
                                        }
                                        if false {
                                            switch (^(fi27)) {
                                            case 32768:
                                                p0 = (p0 + p0)
                                            case 32767:
                                                for fi49 := 0; fi49 < 1; fi49++ {
                                                    var v50 int8 = -127
                                                    _ = fi49
                                                    _ = v50
                                                }
                                            default:
                                                v26 -= v26
                                            }
                                            var v51 uint32 = 619
                                            fmt.Println(v26)
                                            _ = v51
                                        } else {
                                            v26 += v26
                                            v32 += v32
                                        }
                                        _ = v36
                                    }
                                default:
                                    var v52 uint16 = p0
                                    _ = v52
                                }
                            case -590:
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                            case 4294967295:
                                fmt.Println(p0, v32, v26)
                            }
                            v25 = append(v25, 4294967295)
                            if true {
                                var v53 string = "hello"
                                var v54 string = v53
                                _ = v53
                                _ = v54
                            } else {
                                v55 := GenS2{F0: 4294967295, F1: (true && false), F2: 435}
                                for fi56 := 0; fi56 < 4; fi56++ {
                                    v26 -= (v26 - ((v26 / 3.14159) + -2.25))
                                    fmt.Println(p0, v26, fi27)
                                    p1 = p1
                                    p0 = 42
                                    _ = fi56
                                }
                                for fi57 := 0; fi57 < 3; fi57++ {
                                    v58 := GenS3{F0: 9223372036854775807, F1: -128, F2: "a"}
                                    _ = fi57
                                    _ = v58
                                }
                                _ = v55
                            }
                        }
                    }
                    _ = fi27
                    _ = v32
                }
                _ = v26
            }
            if (true && true) {
                var v59 uint8 = 2
                p0 = p0
                _ = v59
            }
            _ = v25
        } else {
            for fi60 := 0; fi60 < 5; fi60++ {
                for fi61 := 0; fi61 < 20; fi61++ {
                    switch (fi61 | p1) {
                    case 32768:
                        if (true && true) {
                            if false {
                                var v62 string = "hello"
                                p0 -= p0
                                var v63 uint8 = 2
                                _ = v62
                                _ = v63
                            } else {
                                v64 := GenS4{F0: 42, F1: 14}
                                _ = v64
                            }
                            var v65 int = (p1 + (-(fi61)))
                            var v66 int16 = 1
                            var v67 float32 = 0.0
                            switch (v65 << ((4294967295) & 63)) {
                            case 32768:
                                var v68 uint64 = 18446744073709551615
                                _ = v68
                            case -9223372036854775808:
                                v67 -= float32(-2.25)
                            case -2147483648:
                                var v69 uint32 = 2147483647
                                _ = v69
                            default:
                                v70 := GenS3{F0: -739, F1: 127, F2: ""}
                                _ = v70
                            }
                            _ = v65
                            _ = v66
                            _ = v67
                        }
                    case 4294967295:
                        v71 := []int{(^(fi60)), (fi61 | fi60)}
                        _ = v71
                    case -468:
                        v72 := []int{fi60, (-(p1)), (p1 >> ((633) & 63))}
                        _ = v72
                    default:
                        var v73 int64 = 912
                        _ = v73
                    }
                    v74 := GenS2{F0: 65535, F1: true, F2: 128}
                    _ = fi61
                    _ = v74
                }
                switch (^(fi60)) {
                case 4294967295:
                    switch "" {
                    case "foo bar":
                        v75 := GenS2{F0: 2147483648, F1: (true && (fi60 >= fi60)), F2: 18446744073709551615}
                        _ = v75
                    case "hello":
                        p1 = (fi60 & (fi60 + p1))
                    case "a":
                        p1 += 2147483647
                    default:
                        var v76 uint8 = 180
                        _ = v76
                    }
                case 2147483647:
                    if (4294967295 == 2) {
                        p0 = p0
                        var v77 float64 = ((((-2.25 + ((3.14159 * float64(-2.25)) / ((3.14159 + (1.5 - -2.25)) + 0.1))) - float64(-2.25)) / ((((1.5 / 1e10) + -2.25) / (((float64(1e10) * (0.0 + ((1e10 / float64(1.5)) - (1e10 / (float64((float32(((-2.25 / -2.25) - -2.25)) / 1.0)) * ((0.1 - 3.14159) / float64(0.1))))))) - (((0.1 / 0.1) / (0.0 + 3.14159)) * 0.1)) / (-2.25 - ((float64(1e10) * ((0.1 - ((((1.5 / 1.0) + (-2.25 - -2.25)) / -2.25) - ((3.14159 - ((float64(1e10) + 0.0) - 1e10)) / (1e10 - 0.1)))) * (0.1 + 0.0))) / (-2.25 / 1.0))))) * -2.25)) * 3.14159)
                        if true {
                            var v78 string = (("a" + ("foo bar" + "\n\t")) + (("\n\t" + "a") + "foo bar"))
                            v77 = -2.25
                            _ = v78
                        }
                        _ = v77
                    } else {
                        var v79 int32 = 42
                        if true {
                            var v80 int32 = 1
                            var v81 float32 = (-2.25 - 1e10)
                            _ = v80
                            _ = v81
                        }
                        _ = v79
                    }
                default:
                    if (!true) {
                        p1 -= 42
                        if false {
                            if (!(-98 == -125)) {
                                if false {
                                    fmt.Println(p1, p0)
                                    if true {
                                        v82 := GenS5{F0: (p0 ^ p0), F1: 0}
                                        p1 += fi60
                                        fmt.Println(p0, p0)
                                        _ = v82
                                    } else {
                                        var v83 uint8 = 128
                                        var v84 uint8 = v83
                                        _ = v83
                                        _ = v84
                                    }
                                    v85 := []int{p1, (fi60 & 0), (p1 % ((-2147483648) | 1)), fi60}
                                    _ = v85
                                }
                            }
                            for fi86 := 0; fi86 < 3; fi86++ {
                                if ((((2147483647 == 4294967295) && true) && (!(false && (true && (((!true) && (!false)) || true))))) || true) {
                                    switch fi60 {
                                    case 0:
                                        p1 -= fi60
                                    default:
                                        p1 = 32768
                                    }
                                }
                                fmt.Println(fi60, p1)
                                fmt.Println(fi60, fi60)
                                _ = fi86
                            }
                        }
                    } else {
                        for fi87 := 0; fi87 < 16; fi87++ {
                            var v88 uint64 = 114
                            _ = fi87
                            _ = v88
                        }
                        v89 := []int{-2147483648, fi60}
                        p1 -= (p1 >> ((255) & 63))
                        _ = v89
                    }
                }
                _ = fi60
            }
        }
        var v90 bool = true
        p1 -= p1
        _ = v90
    }
    switch "" {
    case "\n\t":
        if false {
            switch ("\n\t" + "hello") {
            case "hello":
                p0 += 68
            case "foo bar":
                p1 = (p1 | (p1 * (-(p1))))
            case "":
                if true {
                    v91 := GenS1{F0: 32767}
                    _ = v91
                }
            }
        }
    }
    switch 127 {
    case 127:
        v92 := []int{p1, (p1 * (p1 ^ (p1 << ((255) & 63)))), 32767}
        _ = v92
    case 2147483647:
        switch "\n\t" {
        case "hello":
            var v93 uint16 = 1
            _ = v93
        case "a":
            v94 := []int{32768, p1, p1, p1}
            _ = v94
        case "foo bar":
            var v95 uint32 = 42
            _ = v95
        default:
            fmt.Println(p0, p1)
        }
    case -9223372036854775808:
        if (1.5 < float64(3.14159)) {
            v96 := []int{}
            if len(v96) > 0 {
                v96[0] = 1
            }
            var v97 string = "a"
            _ = v96
            _ = v97
        } else {
            v98 := GenS1{F0: 1}
            _ = v98
        }
    default:
        p0 = p0
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    var v5 float64 = 1.5
    v6 := make(map[int]uint32)
    v99 := GenS6{F0: 0}
    v5 -= 3.14159
    var v100 uint8 = 0
    if (!(!(18446744073709551615 > 1))) {
        v101 := GenS1{F0: 53}
        _ = v101
    } else {
        v100 = (v100 ^ v100)
        for fi102 := 0; fi102 < 9; fi102++ {
            switch "a" {
            case "foo bar":
                v103 := []int{}
                _ = v103
            }
            v2 += 0
            if (false || true) {
                switch fi102 {
                case 42:
                    v5 -= 0.1
                }
            } else {
                switch "hello" {
                case "\n\t":
                    var v104 int32 = v2
                    _ = v104
                case "foo bar":
                    if (4294967295 == 2) {
                        for fi105 := 0; fi105 < 5; fi105++ {
                            var v106 uint16 = 212
                            v100 += 229
                            v107 := v3[2147483647]
                            v100 = v100
                            fmt.Println(v2, v5)
                            v108 := GenS3{F0: 0, F1: 100, F2: v107}
                            _ = fi105
                            _ = v106
                            _ = v107
                            _ = v108
                        }
                        for fi109 := 0; fi109 < 15; fi109++ {
                            v110 := GenS4{F0: (^(v100)), F1: 297}
                            switch fi109 {
                            case 2147483647:
                                switch "hello" {
                                case "hello":
                                    var v111 uint = 38
                                    _ = v111
                                case "\n\t":
                                    _ = len(v6)
                                default:
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        fmt.Println(v100)
                                    case "hello":
                                        switch "hello" {
                                        case "":
                                            v5 += (0.0 / (((-2.25 + v5) + float64((float32(0.1) * -2.25))) + v5))
                                        case "foo bar":
                                            v112 := []int{}
                                            _ = v112
                                        case "a":
                                            var v113 int64 = 92
                                            _ = v113
                                        default:
                                            if (!(-1 < -2147483648)) {
                                                switch "hello" {
                                                case "":
                                                    v5 = 0.1
                                                }
                                                v100 += (v100 * (v100 + v100))
                                                var v114 int64 = -9223372036854775808
                                                _ = v114
                                            }
                                        }
                                    default:
                                        fmt.Println(v100, v100)
                                    }
                                }
                            case -1:
                                var v115 int16 = 42
                                _ = v115
                            }
                            _ = fi109
                            _ = v110
                        }
                        v2 = v2
                    } else {
                        v2 = v2
                        v116 := []int{fi102}
                        _ = v116
                    }
                case "hello":
                    if false {
                        v5 = (((0.0 / 1.5) / 1.0) * 0.0)
                        v117 := []int{}
                        _ = len(v3)
                        _ = v117
                    } else {
                        v100 = v100
                        var v118 float32 = -2.25
                        for fi119 := 0; fi119 < 2; fi119++ {
                            v118 = 1e10
                            v120 := GenS3{F0: 42, F1: -47, F2: ""}
                            _ = fi119
                            _ = v120
                        }
                        _ = v118
                    }
                default:
                    var v121 int8 = 42
                    _ = v121
                }
                if false {
                    v122 := []int{(-(fi102)), (-(fi102)), fi102}
                    if false {
                        v123 := GenS2{F0: 887, F1: true, F2: 0}
                        v2 += -1
                        _ = v123
                    } else {
                        v100 = (^(v100))
                        v2 -= v2
                        if ((v100 + (v100 + (^(v100)))) > 197) {
                            if false {
                                var v124 int32 = v2
                                switch "" {
                                case "":
                                    var v125 int = 42
                                    _ = v125
                                case "foo bar":
                                    v2 += 42
                                case "hello":
                                    var v126 float32 = (-2.25 + float32(v5))
                                    _ = v126
                                default:
                                    switch "foo bar" {
                                    case "":
                                        for fi127 := 0; fi127 < 12; fi127++ {
                                            var v128 uint64 = 1
                                            v100 -= 1
                                            _ = fi127
                                            _ = v128
                                        }
                                    case "a":
                                        for fi129 := 0; fi129 < 5; fi129++ {
                                            switch (fi129 ^ fi129) {
                                            case 687:
                                                switch fi129 {
                                                case -2147483648:
                                                    var v130 bool = (!(!((1 < 831) && true)))
                                                    _ = v130
                                                case 1:
                                                    var v131 bool = (177 == 852)
                                                    _ = v131
                                                default:
                                                    var v132 uint32 = 2147483647
                                                    _ = v132
                                                }
                                            case 127:
                                                if false {
                                                    if (799 >= 2) {
                                                        v5 = 3.14159
                                                        var v133 uint8 = v100
                                                        var v134 float64 = (float64((float32(1e10) - 0.0)) / v5)
                                                        _ = v133
                                                        _ = v134
                                                    } else {
                                                        v135 := GenS4{F0: 1, F1: 331}
                                                        v124 = (-(v124))
                                                        _ = v135
                                                    }
                                                }
                                            case -80:
                                                v100 += v100
                                            }
                                            _ = fi129
                                        }
                                    case "\n\t":
                                        if false {
                                            v2 += (v124 & -2147483648)
                                            switch fi102 {
                                            case -551:
                                                var v136 int64 = -1
                                                _ = v136
                                            case -186:
                                                if false {
                                                    v137 := GenS3{F0: 0, F1: -1, F2: "foo bar"}
                                                    if (float64((1.5 * ((3.14159 / 3.14159) + float32(float64(-2.25))))) == (v5 * float64((1.5 * -2.25)))) {
                                                        if true {
                                                            var v138 uint16 = 2
                                                            switch (fi102 | fi102) {
                                                            case 42:
                                                                switch (fi102 + fi102) {
                                                                case -9223372036854775808:
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if (42 <= 415) {
                                                                            for fi139 := 0; fi139 < 8; fi139++ {
                                                                                var v140 int16 = 0
                                                                                _ = fi139
                                                                                _ = v140
                                                                            }
                                                                            var v141 int64 = -1
                                                                            _ = v141
                                                                        } else {
                                                                            v142 := GenS3{F0: 9223372036854775807, F1: -84, F2: ""}
                                                                            var v143 uint64 = 1
                                                                            _ = v142
                                                                            _ = v143
                                                                        }
                                                                    case "a":
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v138 += v138
                                                                        case "hello":
                                                                            var v144 bool = false
                                                                            _ = v144
                                                                        default:
                                                                            var v145 string = "foo bar"
                                                                            _ = v145
                                                                        }
                                                                    }
                                                                default:
                                                                    v146 := GenS7{F0: -32768, F1: (fi102 << ((18446744073709551615) & 63)), F2: 2}
                                                                    _ = v146
                                                                }
                                                            case 32768:
                                                                delete(v3, 387)
                                                            }
                                                            for fi147 := 0; fi147 < 6; fi147++ {
                                                                switch fi147 {
                                                                case 255:
                                                                    var v148 uint = 4294967295
                                                                    _ = v148
                                                                case 128:
                                                                    var v149 string = ("a" + (("hello" + "\n\t") + "foo bar"))
                                                                    var v150 float64 = v5
                                                                    _ = v149
                                                                    _ = v150
                                                                case 491:
                                                                    switch -1 {
                                                                    case -382:
                                                                        v151 := GenS8{F0: -414, F1: 18446744073709551615, F2: 6}
                                                                        _ = v151
                                                                    case 0:
                                                                        var v152 uint64 = 4294967295
                                                                        _ = v152
                                                                    case -408:
                                                                        var v153 int64 = -1
                                                                        _ = v153
                                                                    default:
                                                                        v154 := GenS3{F0: 964, F1: 100, F2: ("" + "hello")}
                                                                        _ = v154
                                                                    }
                                                                default:
                                                                    if true {
                                                                        switch fi147 {
                                                                        case 4294967295:
                                                                            if false {
                                                                                v5 -= -2.25
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    if false {
                                                                                        v122 = append(v122, fi147)
                                                                                        v5 = float64(float32(((1.5 + 0.0) / (v5 + v5))))
                                                                                    }
                                                                                default:
                                                                                    var v155 float32 = 1.5
                                                                                    _ = v155
                                                                                }
                                                                                var v156 string = ("\n\t" + "\n\t")
                                                                                var v157 int = (fi147 - (fi102 << ((1) & 63)))
                                                                                _ = v156
                                                                                _ = v157
                                                                            } else {
                                                                                v158 := GenS9{F0: 0, F1: (true && true)}
                                                                                var v159 uint32 = 4294967295
                                                                                if len(v122) > 0 {
                                                                                    v122[0] = (^(fi102))
                                                                                }
                                                                                switch fi147 {
                                                                                case 9223372036854775807:
                                                                                    if len(v122) >= 2 {
                                                                                        v122 = v122[1:len(v122)]
                                                                                    }
                                                                                    v160 := v3[(v159 ^ (v159 >> ((255) & 63)))]
                                                                                    var v161 float64 = 1.5
                                                                                    _ = v160
                                                                                    _ = v161
                                                                                case -9223372036854775808:
                                                                                    switch (fi102 / ((int(v100)) | 1)) {
                                                                                    case -1:
                                                                                        fmt.Println(fi102)
                                                                                    case -2147483648:
                                                                                        v162 := GenS9{F0: 32767, F1: (fi147 > fi147)}
                                                                                        var v163 uint64 = 0
                                                                                        var v164 uint64 = v163
                                                                                        var v165 int = (-(fi102))
                                                                                        switch (fi147 - fi102) {
                                                                                        case 2147483647:
                                                                                            if (1.5 <= (float32((((float64((float32(v5) * (float32(float64(-2.25)) - 3.14159))) - (0.1 * (3.14159 + v5))) * (1.5 * v5)) / v5)) * (((0.0 - 0.0) * 0.1) / float32(((0.0 - v5) * v5))))) {
                                                                                                var v166 int32 = v124
                                                                                                fmt.Println(v5, v165)
                                                                                                _ = v166
                                                                                            }
                                                                                        case 32768:
                                                                                            if false {
                                                                                                switch (("hello" + "a") + "hello") {
                                                                                                case "":
                                                                                                    v164 = 4294967295
                                                                                                case "\n\t":
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v138, v100)
                                                                                                    case "hello":
                                                                                                        if (false && (v159 < v159)) {
                                                                                                            v167 := GenS5{F0: v138, F1: v2}
                                                                                                            v5 -= 1e10
                                                                                                            v100 += v100
                                                                                                            _ = v167
                                                                                                        } else {
                                                                                                            v168 := GenS3{F0: 42, F1: 100, F2: ""}
                                                                                                            v169 := GenS2{F0: (^(v159)), F1: (float32(v5) == float32(-2.25)), F2: (v164 / ((v164) | 1))}
                                                                                                            v5 += v5
                                                                                                            _ = v168
                                                                                                            _ = v169
                                                                                                        }
                                                                                                    default:
                                                                                                        v163 = uint64(v100)
                                                                                                    }
                                                                                                }
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        var v170 bool = ((!false) || (!false))
                                                                                                        _ = v170
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    if (!false) {
                                                                                                        _ = len(v1)
                                                                                                        if len(v122) > 0 {
                                                                                                            v122[0] = (v165 - v165)
                                                                                                        }
                                                                                                        fmt.Println(v100, v164, fi102)
                                                                                                        _ = len(v122)
                                                                                                        v138 -= (v138 / ((v138) | 1))
                                                                                                    }
                                                                                                case "a":
                                                                                                    fmt.Println(fi147, fi147)
                                                                                                default:
                                                                                                    v124 = v2
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v162
                                                                                        _ = v163
                                                                                        _ = v164
                                                                                        _ = v165
                                                                                    default:
                                                                                        var v171 int64 = 42
                                                                                        _ = v171
                                                                                    }
                                                                                }
                                                                                switch ("foo bar" + "foo bar") {
                                                                                case "\n\t":
                                                                                    v172 := GenS10{F0: 822}
                                                                                    _ = v172
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        if (9223372036854775807 != int64(v2)) {
                                                                                            switch "hello" {
                                                                                            case "hello":
                                                                                                if ((9223372036854775807 == 1) || (2 == 142)) {
                                                                                                    v124 = (v2 % (((-(v124))) | 1))
                                                                                                    v124 -= v2
                                                                                                    if (1 < 65535) {
                                                                                                        v100 += v100
                                                                                                        v173 := v1[("a" + "\n\t")]
                                                                                                        var v174 float64 = v5
                                                                                                        _ = v173
                                                                                                        _ = v174
                                                                                                    } else {
                                                                                                        v175 := GenS11{F0: -32768, F1: uint(v159), F2: "hello"}
                                                                                                        v138 += 907
                                                                                                        v138 -= (v138 << ((4294967295) & 63))
                                                                                                        _ = v175
                                                                                                    }
                                                                                                    v5 += v5
                                                                                                    var v176 string = (("a" + "hello") + ("\n\t" + (("a" + "\n\t") + "")))
                                                                                                    v122 = append(v122, fi102)
                                                                                                    _ = v176
                                                                                                }
                                                                                            case "a":
                                                                                                v124 = -2147483648
                                                                                            default:
                                                                                                var v177 uint8 = 128
                                                                                                _ = v177
                                                                                            }
                                                                                            if false {
                                                                                                if len(v122) > 0 {
                                                                                                    v122[0] = 42
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v100, v5, v100)
                                                                                        } else {
                                                                                            v122 = append(v122, fi147)
                                                                                            v100 = 9
                                                                                            var v178 float64 = (v5 * 1e10)
                                                                                            _ = v178
                                                                                        }
                                                                                        if true {
                                                                                            var v179 int = (-(fi102))
                                                                                            var v180 float64 = (v5 / 1e10)
                                                                                            var v181 float32 = 1e10
                                                                                            _ = v179
                                                                                            _ = v180
                                                                                            _ = v181
                                                                                        } else {
                                                                                            v182 := GenS12{F0: v124, F1: (-(fi147))}
                                                                                            v138 -= v138
                                                                                            _ = len(v122)
                                                                                            _ = v182
                                                                                        }
                                                                                    } else {
                                                                                        if true {
                                                                                            var v183 int8 = 42
                                                                                            _ = v183
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        if (!(!(false && (!false)))) {
                                                                                            var v184 uint32 = 1
                                                                                            if (float32(-2.25) == 0.1) {
                                                                                                var v185 uint32 = v184
                                                                                                _ = v185
                                                                                            }
                                                                                            var v186 uint32 = 275
                                                                                            _ = v184
                                                                                            _ = v186
                                                                                        } else {
                                                                                            if (("a" <= ("a" + "")) && true) {
                                                                                                var v187 int64 = 42
                                                                                                if (false && (v187 <= 1)) {
                                                                                                    v188 := GenS4{F0: (v100 & v100), F1: 1}
                                                                                                    v189 := GenS1{F0: -440}
                                                                                                    var v190 uint32 = v159
                                                                                                    _ = len(v122)
                                                                                                    v2 -= (v124 + v124)
                                                                                                    _ = v188
                                                                                                    _ = v189
                                                                                                    _ = v190
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v159 -= (^(v159))
                                                                                                        var v191 uint64 = 0
                                                                                                        _ = v191
                                                                                                    } else {
                                                                                                        fmt.Println(fi147, fi102, v187)
                                                                                                        fmt.Println(fi147)
                                                                                                        var v192 float32 = -2.25
                                                                                                        _ = v192
                                                                                                    }
                                                                                                    v159 = (^(v159))
                                                                                                    if false {
                                                                                                        if (false && (v2 == (v2 * v2))) {
                                                                                                            _ = len(v122)
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v193 float32 = (1.5 * ((-2.25 + 0.1) + 3.14159))
                                                                                                        _ = v193
                                                                                                    }
                                                                                                }
                                                                                                v5 = ((0.1 / 1.0) + (v5 / float64(1.5)))
                                                                                                _ = v187
                                                                                            } else {
                                                                                                _ = len(v4)
                                                                                                var v194 string = ""
                                                                                                _ = v194
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        switch (fi102 + 127) {
                                                                                        case 860:
                                                                                            v2 = v2
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(fi102, v138)
                                                                                        case 9223372036854775807:
                                                                                            v100 = 121
                                                                                        }
                                                                                        var v195 int64 = 835
                                                                                        _ = v195
                                                                                    }
                                                                                }
                                                                                _ = v158
                                                                                _ = v159
                                                                            }
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v3[1] = ("hello" + "hello")
                                                                        case "":
                                                                            var v196 uint8 = v100
                                                                            _ = v196
                                                                        default:
                                                                            v5 = v5
                                                                        }
                                                                        delete(v6, fi102)
                                                                    }
                                                                }
                                                                _ = fi147
                                                            }
                                                            _ = v138
                                                        }
                                                        for fi197 := 0; fi197 < 16; fi197++ {
                                                            v198 := GenS10{F0: 4294967295}
                                                            v5 -= (v5 - v5)
                                                            v5 += -2.25
                                                            _ = fi197
                                                            _ = v198
                                                        }
                                                        for fi199 := 0; fi199 < 1; fi199++ {
                                                            fmt.Println(v2, fi199, v5)
                                                            _ = fi199
                                                        }
                                                    }
                                                    _ = v137
                                                }
                                            case 32768:
                                                switch fi102 {
                                                case -357:
                                                    v2 = v2
                                                }
                                            default:
                                                v200 := GenS12{F0: (v2 ^ -1), F1: fi102}
                                                _ = v200
                                            }
                                        } else {
                                            var v201 int = fi102
                                            var v202 int16 = 163
                                            if false {
                                                v203 := GenS13{F0: 0, F1: v2, F2: 18446744073709551615}
                                                fmt.Println(v5)
                                                var v204 float64 = -2.25
                                                _ = v203
                                                _ = v204
                                            }
                                            _ = v201
                                            _ = v202
                                        }
                                    }
                                }
                                fmt.Println(v2, v5, v100)
                                _ = v124
                            }
                            switch (fi102 % ((401) | 1)) {
                            case -362:
                                var v205 int8 = 100
                                _ = v205
                            case 32768:
                                v5 -= 0.0
                            }
                            switch 4294967295 {
                            case 0:
                                var v206 float64 = (((v5 * -2.25) * 0.1) + 0.0)
                                _ = v206
                            case -538:
                                v100 = v100
                            }
                        } else {
                            var v207 uint32 = 986
                            if true {
                                v122 = append(v122, (fi102 << ((341) & 63)))
                                if false {
                                    v208 := GenS13{F0: 1, F1: (v2 | v2), F2: 42}
                                    var v209 int16 = 32767
                                    _ = v208
                                    _ = v209
                                } else {
                                    var v210 uint16 = 65535
                                    if true {
                                        for fi211 := 0; fi211 < 7; fi211++ {
                                            switch fi102 {
                                            case 2147483647:
                                                switch (("\n\t" + "foo bar") + ("foo bar" + "\n\t")) {
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "a":
                                                        if (!(true && false)) {
                                                            if len(v122) >= 2 {
                                                                v122 = v122[1:len(v122)]
                                                            }
                                                            var v212 uint = 2
                                                            v122 = append(v122, fi211)
                                                            _ = v212
                                                        }
                                                    case "":
                                                        v5 += v5
                                                    case "hello":
                                                        _ = len(v122)
                                                    default:
                                                        var v213 float64 = v5
                                                        _ = v213
                                                    }
                                                case "a":
                                                    v214 := v3[v207]
                                                    _ = v214
                                                }
                                            }
                                            if len(v122) > 0 {
                                                v122[0] = (fi211 ^ 42)
                                            }
                                            _ = len(v4)
                                            _ = fi211
                                        }
                                        fmt.Println(v2, v207)
                                        if false {
                                            delete(v4, (fi102 & (fi102 - fi102)))
                                            switch fi102 {
                                            case 4294967295:
                                                v2 -= (-(v2))
                                            case -9223372036854775808:
                                                for fi215 := 0; fi215 < 19; fi215++ {
                                                    fmt.Println(v5, fi102)
                                                    if true {
                                                        var v216 int64 = int64(fi215)
                                                        var v217 uint = 42
                                                        _ = v216
                                                        _ = v217
                                                    } else {
                                                        v218 := GenS14{F0: (v100 & v100), F1: (v210 - v210), F2: v2}
                                                        if (1 >= 42) {
                                                            var v219 uint32 = (v207 | (v207 - (v207 << ((2) & 63))))
                                                            v207 += 110
                                                            v220 := v6[(-(fi102))]
                                                            fmt.Println(fi102, fi215)
                                                            _ = v219
                                                            _ = v220
                                                        } else {
                                                            if false {
                                                                v221 := v6[4294967295]
                                                                _ = v221
                                                            }
                                                            delete(v1, "a")
                                                            var v222 int8 = 42
                                                            _ = v222
                                                        }
                                                        var v223 uint8 = (^(v100))
                                                        var v224 float64 = v5
                                                        v225 := GenS11{F0: 885, F1: 255, F2: "\n\t"}
                                                        _ = v218
                                                        _ = v223
                                                        _ = v224
                                                        _ = v225
                                                    }
                                                    _ = fi215
                                                }
                                            default:
                                                var v226 uint = 42
                                                _ = v226
                                            }
                                        }
                                    } else {
                                        v100 = (v100 * v100)
                                    }
                                    _ = v210
                                }
                            } else {
                                var v227 uint64 = 42
                                v207 = (v207 << ((1) & 63))
                                _ = v227
                            }
                            if true {
                                v2 -= (^(v2))
                            }
                            _ = v207
                        }
                    }
                    var v228 uint32 = 1
                    _ = v122
                    _ = v228
                }
                v229 := GenS2{F0: 65535, F1: (((v100 <= v100) && (!true)) && (0 == -128)), F2: 0}
                _ = v229
            }
            _ = fi102
        }
    }
    var v230 uint64 = 2
    v231 := GenS9{F0: -32768, F1: ((v100 | 1) != 128)}
    fmt.Println(v2, v5)
    var v232 uint = 65535
    if true {
        fmt.Println(v2)
        for fi233 := 0; fi233 < 3; fi233++ {
            if (!true) {
                var v234 int = (fi233 * (fi233 - fi233))
                _ = v234
            } else {
                v2 -= v2
            }
            v2 = (v2 >> ((255) & 63))
            switch fi233 {
            case 882:
                v2 = v2
            default:
                v5 = v5
            }
            _ = fi233
        }
    } else {
        var v235 int16 = -32768
        _ = v235
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v230:", v230)
    fmt.Println("v232:", v232)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v100
    _ = v2
    _ = v230
    _ = v231
    _ = v232
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
