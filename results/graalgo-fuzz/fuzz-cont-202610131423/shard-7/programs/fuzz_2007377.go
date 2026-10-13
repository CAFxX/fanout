package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int32; F2 uint32 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 int64; F1 int64; F2 uint8 }
type GenS13 struct { F0 int32; F1 uint16; F2 int16 }
type GenS14 struct { F0 uint32; F1 uint }
type GenS15 struct { F0 int32; F1 int64; F2 bool }
type GenS16 struct { F0 uint; F1 int64 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 uint32; F1 string; F2 int }
type GenS20 struct { F0 uint32; F1 uint }
type GenS21 struct { F0 int; F1 bool; F2 uint16 }
type GenS22 struct { F0 int32; F1 uint8 }
type GenS23 struct { F0 bool }
func genFn0(p0 uint) int8 {
    p0 -= (p0 << (((^(p0))) & 63))
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 102}
        var v8 int64 = -9223372036854775808
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        v9 := []int{32768, 1}
        fmt.Println(p0)
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS1{F0: 1}
            for fi12 := 0; fi12 < 3; fi12++ {
                p0 -= p0
                _ = fi12
            }
            if false {
                v13 := GenS2{F0: -1, F1: p0, F2: 42}
                _ = v13
            } else {
                if true {
                    fmt.Println(fi10)
                }
                fmt.Println(fi10)
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        for fi14 := 0; fi14 < 15; fi14++ {
            if true {
                if (true && false) {
                    if false {
                        v15 := []int{int(p0), fi14, fi14}
                        fmt.Println(p0)
                        switch ("hello" + "hello") {
                        case "":
                            var v16 uint = 620
                            _ = v16
                        case "a":
                            v17 := GenS3{F0: 1, F1: (75 <= 618)}
                            _ = v17
                        case "foo bar":
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                        _ = v15
                    } else {
                        if true {
                            fmt.Println(p0, p0)
                            p0 -= (p0 >> ((663) & 63))
                        }
                        v18 := GenS1{F0: 127}
                        p0 = p0
                        _ = v18
                    }
                }
                v19 := []int{-2147483648, (fi14 >> ((648) & 63))}
                var v20 float64 = (0.1 - 1e10)
                _ = v19
                _ = v20
            }
            p0 = 1
            _ = fi14
        }
        if true {
            switch ("\n\t" + "a") {
            case "a":
                v21 := []int{}
                _ = v21
            default:
                v22 := GenS2{F0: 1, F1: (p0 * p0), F2: (p0 % (((p0 | 0)) | 1))}
                _ = v22
            }
            p0 += (^(p0))
            var v23 uint16 = 255
            _ = v23
        } else {
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int = (-(fi24))
                if true {
                    v26 := GenS2{F0: -2147483648, F1: p0, F2: p0}
                    p0 -= (p0 + p0)
                    if false {
                        var v27 uint64 = 42
                        switch "foo bar" {
                        case "\n\t":
                            var v28 string = "a"
                            _ = v28
                        case "a":
                            switch (fi24 * 32767) {
                            case 518:
                                if false {
                                    v27 -= (^(v27))
                                    v27 = v27
                                    if (32767 == -152) {
                                        p0 += (p0 * 255)
                                        var v29 int = (^(fi24))
                                        _ = v29
                                    }
                                }
                            }
                        case "hello":
                            var v30 string = ("a" + "hello")
                            _ = v30
                        default:
                            fmt.Println(p0, v27, fi24)
                        }
                        var v31 uint = p0
                        _ = v27
                        _ = v31
                    } else {
                        fmt.Println(fi24)
                    }
                    _ = v26
                }
                v32 := GenS4{F0: 1, F1: 1}
                _ = fi24
                _ = v25
                _ = v32
            }
        }
    }
    v33 := []int{}
    v34 := GenS5{F0: 1}
    _ = p0
    _ = v33
    _ = v34
    return -120
}
func genFn1() uint16 {
    v35 := GenS6{F0: (!(0.1 <= 1e10)), F1: -2147483648, F2: 0}
    v36 := []int{32768, 42}
    _ = v35
    _ = v36
    return 255
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 0.1
    var v3 uint = 65535
    v4 := make(map[int]int)
    var v5 float64 = ((float64(1.5) - v2) / v2)
    v3 -= v3
    var v37 int = 2147483647
    v38 := []int{-1}
    fmt.Println(v2)
    v3 = 2
    var v39 uint16 = 0
    switch v37 {
    case 32768:
        fmt.Println(v39, v5)
    case 4294967295:
        fmt.Println(v37)
    default:
        for fi40 := 0; fi40 < 12; fi40++ {
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 uint64 = 506
                v2 -= 0.0
                delete(v4, (fi41 << ((42) & 63)))
                _ = fi41
                _ = v42
            }
            if ("hello" > "hello") {
                switch ("\n\t" + "\n\t") {
                case "hello":
                    if (false || (!(true && (1 < 959)))) {
                        for fi43 := 0; fi43 < 5; fi43++ {
                            switch 127 {
                            case 1:
                                fmt.Println(fi40)
                            case -60:
                                v38 = append(v38, v37)
                            case 32767:
                                v2 = ((0.1 / ((v5 * v2) / float64(1e10))) / (float64((3.14159 * float32(float64(-2.25)))) + ((v2 * v2) * v5)))
                            default:
                                v2 = 0.1
                            }
                            switch "\n\t" {
                            case "hello":
                                v37 += v37
                            }
                            _ = fi43
                        }
                    }
                case "\n\t":
                    var v44 uint32 = 65535
                    _ = v44
                default:
                    v39 -= 1
                }
                switch (v37 ^ v37) {
                case 1:
                    if (true || ((-2147483648 >= 271) || (true && true))) {
                        v2 -= 1e10
                    }
                case 357:
                    v2 = 1.5
                default:
                    v45 := GenS7{F0: -850, F1: 42}
                    _ = v45
                }
                if (1 == 42) {
                    var v46 float32 = 0.1
                    var v47 float64 = v2
                    if true {
                        v48 := GenS8{F0: 42}
                        _ = v48
                    } else {
                        var v49 int8 = 0
                        _ = v49
                    }
                    _ = v46
                    _ = v47
                }
            } else {
                var v50 uint32 = 735
                _ = v50
            }
            delete(v1, "hello")
            _ = fi40
        }
    }
    v4[127] = 127
    var v51 bool = true
    var v52 int8 = 127
    if false {
        switch (v37 / ((v37) | 1)) {
        case -2147483648:
            for fi53 := 0; fi53 < 11; fi53++ {
                for fi54 := 0; fi54 < 13; fi54++ {
                    switch (("" + "foo bar") + ("foo bar" + "")) {
                    case "\n\t":
                        var v55 bool = false
                        _ = v55
                    case "foo bar":
                        v2 += (((float64(1.5) - float64(1.5)) - v5) - v5)
                    case "":
                        if v51 {
                            switch "" {
                            case "hello":
                                var v56 int32 = -2147483648
                                if v51 {
                                    if (3.14159 > v2) {
                                        var v57 int16 = 949
                                        _ = v57
                                    }
                                    fmt.Println(v56)
                                    v58 := GenS9{F0: v39, F1: -9223372036854775808}
                                    _ = v58
                                }
                                var v59 uint = (v3 % ((v3) | 1))
                                _ = v56
                                _ = v59
                            }
                        }
                    }
                    var v60 int32 = 159
                    var v61 int32 = v60
                    _ = fi54
                    _ = v60
                    _ = v61
                }
                var v62 int8 = v52
                _ = fi53
                _ = v62
            }
        case 4294967295:
            v51 = v51
        }
    } else {
        v63 := GenS7{F0: -320, F1: 446}
        for fi64 := 0; fi64 < 6; fi64++ {
            if false {
                v52 -= v52
                for fi65 := 0; fi65 < 12; fi65++ {
                    if v51 {
                        delete(v1, "a")
                        if v51 {
                            delete(v1, ("" + "\n\t"))
                            var v66 bool = true
                            _ = v66
                        } else {
                            if (!(-2.25 < 3.14159)) {
                                if v51 {
                                    v3 -= v3
                                    switch v37 {
                                    case 403:
                                        if (!v51) {
                                            switch "" {
                                            case "\n\t":
                                                if true {
                                                    v52 = v52
                                                    v2 = 0.1
                                                    if false {
                                                        if ((-(v37)) >= (fi65 << (((^(v3))) & 63))) {
                                                            v3 += (^(v3))
                                                            fmt.Println(v52)
                                                        } else {
                                                            v52 = -46
                                                            switch v37 {
                                                            case -869:
                                                                v52 += (v52 - (^(v52)))
                                                            case 0:
                                                                v2 += 1e10
                                                            default:
                                                                var v67 uint8 = 127
                                                                _ = v67
                                                            }
                                                            switch (v37 * -26) {
                                                            case 1:
                                                                if true {
                                                                    v39 -= v39
                                                                    v5 -= 3.14159
                                                                    if (-90 >= -20) {
                                                                        if v51 {
                                                                            fmt.Println(fi65, v5, v52)
                                                                        }
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v68 := GenS10{F0: 127, F1: 127}
                                                                        _ = v68
                                                                    case "\n\t":
                                                                        switch (v37 ^ (fi64 & (fi65 + (fi65 << ((230) & 63))))) {
                                                                        case 32768:
                                                                            var v69 bool = false
                                                                            _ = v69
                                                                        case 467:
                                                                            if v51 {
                                                                                if (!(!v51)) {
                                                                                    v39 = (v39 ^ (^(v39)))
                                                                                    var v70 uint64 = 458
                                                                                    if v51 {
                                                                                        var v71 uint32 = 860
                                                                                        var v72 uint16 = v39
                                                                                        var v73 int = (fi64 >> ((779) & 63))
                                                                                        _ = v71
                                                                                        _ = v72
                                                                                        _ = v73
                                                                                    }
                                                                                    _ = v70
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    if len(v38) >= 2 {
                                                                                        v38 = v38[1:len(v38)]
                                                                                    }
                                                                                default:
                                                                                    switch v37 {
                                                                                    case 4294967295:
                                                                                        v51 = (true || v51)
                                                                                    default:
                                                                                        v74 := v1[("a" + "")]
                                                                                        v51 = (9223372036854775807 > -9223372036854775808)
                                                                                        _ = v74
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v75 int8 = (v52 + (v52 ^ v52))
                                                                                _ = v75
                                                                            }
                                                                        default:
                                                                            var v76 uint16 = genFn1()
                                                                            _ = v76
                                                                        }
                                                                        if len(v38) > 0 {
                                                                            v38[0] = fi65
                                                                        }
                                                                        var v77 float64 = v2
                                                                        _ = v77
                                                                    }
                                                                    v78 := GenS11{F0: (4294967295 != 2147483647), F1: 255}
                                                                    _ = len(v38)
                                                                    _ = v78
                                                                }
                                                            case 32768:
                                                                var v79 float64 = v5
                                                                _ = v79
                                                            case 42:
                                                                v80 := GenS12{F0: 9223372036854775807, F1: 1, F2: 64}
                                                                _ = v80
                                                            }
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v81 float32 = ((3.14159 + float32(v5)) * (0.1 + 0.1))
                                                            _ = v81
                                                        }
                                                        var v82 float32 = (0.1 + (0.0 + (-2.25 * (3.14159 + 1e10))))
                                                        v83 := GenS2{F0: 1, F1: uint(v39), F2: v3}
                                                        v51 = v51
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v84 := GenS13{F0: 633, F1: v39, F2: 1000}
                                                        _ = v84
                                                    }
                                                }
                                            case "":
                                                v85 := GenS8{F0: -9223372036854775808}
                                                _ = v85
                                            case "a":
                                                v86 := GenS11{F0: v51, F1: 4294967295}
                                                _ = v86
                                            default:
                                                v2 += (v5 - (v5 - (1.5 + -2.25)))
                                            }
                                            v51 = false
                                            if len(v38) >= 2 {
                                                v38 = v38[1:len(v38)]
                                            }
                                        } else {
                                            var v87 uint64 = 1
                                            _ = v87
                                        }
                                    }
                                    switch (fi65 >> (((v3 & (^(v3)))) & 63)) {
                                    case 32768:
                                        var v88 uint64 = 1
                                        _ = v88
                                    case 269:
                                        if true {
                                            v89 := v1[("\n\t" + ("\n\t" + "hello"))]
                                            if ("a" > "") {
                                                if ((-2147483648 >= -1) || v51) {
                                                    var v90 int8 = genFn0((v3 - v3))
                                                    v2 = float64(float32((0.1 - v5)))
                                                    var v91 uint = (^(v3))
                                                    switch "\n\t" {
                                                    case "":
                                                        var v92 uint32 = 2147483647
                                                        _ = v92
                                                    case "hello":
                                                        v91 = (v91 % ((v3) | 1))
                                                    default:
                                                        v89 = (((v37 << ((255) & 63)) >= fi65) || ((v3 == (v3 << ((v3) & 63))) || v51))
                                                    }
                                                    _ = v90
                                                    _ = v91
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    switch (v37 - (fi65 + (v37 * (^(v37))))) {
                                                    case -61:
                                                        var v93 bool = (1 <= 9223372036854775807)
                                                        _ = v93
                                                    case -1:
                                                        if v51 {
                                                            switch (^(v37)) {
                                                            case -9223372036854775808:
                                                                fmt.Println(v3, fi65)
                                                            case 32768:
                                                                v37 = (fi64 * (fi64 * (fi64 + fi65)))
                                                            }
                                                            v3 = 18446744073709551615
                                                            fmt.Println(v39, v2, v5)
                                                        }
                                                        if true {
                                                            fmt.Println(v3, v5, fi64)
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                var v94 int8 = (v52 % (((v52 - (v52 * int8(v37)))) | 1))
                                                                _ = v94
                                                            }
                                                            if v51 {
                                                                fmt.Println(v89)
                                                            }
                                                        }
                                                        v95 := GenS14{F0: 2, F1: (v3 / ((v3) | 1))}
                                                        _ = v95
                                                    case 4294967295:
                                                        var v96 uint32 = 1
                                                        _ = v96
                                                    default:
                                                        delete(v1, ("foo bar" + "hello"))
                                                    }
                                                case "hello":
                                                    var v97 int8 = v52
                                                    _ = v97
                                                default:
                                                    if false {
                                                        var v98 uint16 = v39
                                                        v99 := GenS8{F0: -9223372036854775808}
                                                        _ = v98
                                                        _ = v99
                                                    } else {
                                                        var v100 uint16 = v39
                                                        _ = v100
                                                    }
                                                }
                                                v89 = (uint64(v3) != 0)
                                            } else {
                                                fmt.Println(v37)
                                                switch 2147483647 {
                                                case -2147483648:
                                                    switch fi64 {
                                                    case -132:
                                                        if v51 {
                                                            switch (fi64 + fi64) {
                                                            case -2147483648:
                                                                if v51 {
                                                                    v101 := GenS8{F0: -736}
                                                                    var v102 int16 = -379
                                                                    delete(v1, ("\n\t" + "\n\t"))
                                                                    _ = v101
                                                                    _ = v102
                                                                }
                                                            case 9223372036854775807:
                                                                fmt.Println(v52)
                                                            case 0:
                                                                v4[fi65] = (fi65 * (v37 & fi65))
                                                            }
                                                        } else {
                                                            v89 = (42 <= 0)
                                                            switch "a" {
                                                            case "a":
                                                                if len(v38) >= 2 {
                                                                    v38 = v38[1:len(v38)]
                                                                }
                                                            }
                                                            var v103 bool = v89
                                                            _ = v103
                                                        }
                                                    case -9223372036854775808:
                                                        fmt.Println(v5, v89, v5)
                                                    default:
                                                        var v104 uint16 = (v39 * v39)
                                                        _ = v104
                                                    }
                                                case 150:
                                                    v105 := GenS8{F0: 9223372036854775807}
                                                    _ = v105
                                                default:
                                                    var v106 uint64 = 2
                                                    _ = v106
                                                }
                                            }
                                            _ = v89
                                        } else {
                                            var v107 bool = v51
                                            if len(v38) > 0 {
                                                v38[0] = (v37 >> ((uint(v39)) & 63))
                                            }
                                            _ = v107
                                        }
                                    case -2147483648:
                                        delete(v1, "")
                                    }
                                    fmt.Println(v2)
                                } else {
                                    if true {
                                        if len(v38) >= 2 {
                                            v38 = v38[1:len(v38)]
                                        }
                                        v108 := GenS1{F0: uint8(v3)}
                                        if v51 {
                                            switch (fi64 & fi65) {
                                            case 4294967295:
                                                if v51 {
                                                    if (0.1 < (v2 * v5)) {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v109 int8 = (-(v52))
                                                            _ = v109
                                                        case "\n\t":
                                                            var v110 uint16 = (v39 * genFn1())
                                                            v4[v37] = (v37 - fi64)
                                                            switch "hello" {
                                                            case "":
                                                                var v111 uint8 = 198
                                                                _ = v111
                                                            default:
                                                                switch (v37 << ((v3) & 63)) {
                                                                case 400:
                                                                    fmt.Println(v5)
                                                                case 4294967295:
                                                                    switch int(v52) {
                                                                    case 0:
                                                                        var v112 int64 = 9223372036854775807
                                                                        _ = v112
                                                                    case 4294967295:
                                                                        if (!(("foo bar" + "foo bar") == "\n\t")) {
                                                                            var v113 string = ""
                                                                            v113 = v113
                                                                            if v51 {
                                                                                switch v113 {
                                                                                case "hello":
                                                                                    v38 = append(v38, fi64)
                                                                                case "\n\t":
                                                                                    v37 -= (fi64 + fi64)
                                                                                case "":
                                                                                    v52 += genFn0(1)
                                                                                }
                                                                                v110 -= v110
                                                                                v114 := GenS1{F0: 230}
                                                                                _ = v114
                                                                            }
                                                                            _ = v113
                                                                        }
                                                                    case -9223372036854775808:
                                                                        v115 := v1["\n\t"]
                                                                        _ = v115
                                                                    default:
                                                                        if len(v38) >= 2 {
                                                                            v38 = v38[1:len(v38)]
                                                                        }
                                                                    }
                                                                case 2147483647:
                                                                    var v116 int32 = 42
                                                                    _ = v116
                                                                }
                                                            }
                                                            switch (fi64 | (fi64 ^ (v37 / ((42) | 1)))) {
                                                            case 32768:
                                                                var v117 uint16 = (v39 >> ((uint(v110)) & 63))
                                                                v118 := GenS15{F0: 0, F1: 9223372036854775807, F2: v51}
                                                                _ = v117
                                                                _ = v118
                                                            case 1:
                                                                var v119 bool = ((255 == 3) && (false || false))
                                                                if (451 == 1) {
                                                                    var v120 bool = false
                                                                    v51 = false
                                                                    if ((1 >= 121) && (!v120)) {
                                                                        var v121 string = ("foo bar" + ("\n\t" + "\n\t"))
                                                                        fmt.Println(v121)
                                                                        v110 -= genFn1()
                                                                        _ = v121
                                                                    } else {
                                                                        var v122 uint64 = 42
                                                                        v123 := GenS16{F0: v3, F1: 881}
                                                                        v120 = ((0 != genFn1()) && v119)
                                                                        _ = v122
                                                                        _ = v123
                                                                    }
                                                                    var v124 float32 = (1.5 / 3.14159)
                                                                    if v120 {
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if v119 {
                                                                                if false {
                                                                                    v125 := GenS17{F0: -9223372036854775808}
                                                                                    _ = v125
                                                                                } else {
                                                                                    if (2 == 767) {
                                                                                        v126 := GenS4{F0: (^(v110)), F1: v110}
                                                                                        fmt.Println(v119, v39)
                                                                                        fmt.Println(v37)
                                                                                        _ = v126
                                                                                    }
                                                                                    var v127 uint16 = 42
                                                                                    v128 := GenS11{F0: v119, F1: 485}
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                }
                                                                            } else {
                                                                                if ((203 <= 42) && v51) {
                                                                                    v129 := GenS5{F0: -32768}
                                                                                    var v130 int8 = v52
                                                                                    var v131 uint = 255
                                                                                    if v51 {
                                                                                        if (v119 && v119) {
                                                                                            fmt.Println(v120, v130, fi65)
                                                                                            var v132 float32 = v124
                                                                                            _ = v132
                                                                                        } else {
                                                                                            _ = len(v4)
                                                                                            var v133 int32 = -1
                                                                                            v39 = (v39 + (v110 + v110))
                                                                                            _ = v133
                                                                                        }
                                                                                        v119 = false
                                                                                        var v134 int64 = 1
                                                                                        _ = v134
                                                                                    }
                                                                                    _ = v129
                                                                                    _ = v130
                                                                                    _ = v131
                                                                                }
                                                                                var v135 float64 = (v2 * 0.1)
                                                                                var v136 uint64 = 18446744073709551615
                                                                                _ = v135
                                                                                _ = v136
                                                                            }
                                                                        case "\n\t":
                                                                            if len(v38) >= 2 {
                                                                                v38 = v38[1:len(v38)]
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v3)
                                                                        }
                                                                    } else {
                                                                        v124 = 0.0
                                                                        fmt.Println(v3, v37)
                                                                        switch -1 {
                                                                        case 0:
                                                                            var v137 int = fi64
                                                                            _ = v137
                                                                        }
                                                                    }
                                                                    _ = v120
                                                                    _ = v124
                                                                }
                                                                var v138 float32 = 1.5
                                                                v139 := GenS18{F0: -386}
                                                                _ = v119
                                                                _ = v138
                                                                _ = v139
                                                            case 32767:
                                                                var v140 uint64 = 0
                                                                _ = v140
                                                            default:
                                                                var v141 int32 = int32(v37)
                                                                _ = v141
                                                            }
                                                            if (v110 >= (v110 * (v110 >> (((^(v3))) & 63)))) {
                                                                v2 = v5
                                                                fmt.Println(v52, v39)
                                                            } else {
                                                                if false {
                                                                    var v142 uint8 = 2
                                                                    fmt.Println(v142)
                                                                    _ = v142
                                                                }
                                                            }
                                                            _ = v110
                                                        default:
                                                            delete(v1, (("hello" + "a") + ("a" + "a")))
                                                        }
                                                        var v143 float32 = (-2.25 * 0.0)
                                                        _ = v143
                                                    }
                                                } else {
                                                    fmt.Println(v51)
                                                    if (2 != 42) {
                                                        v1["foo bar"] = (-1 >= 0)
                                                        var v144 bool = v51
                                                        v145 := GenS3{F0: 1, F1: false}
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    if true {
                                                        if true {
                                                            v146 := v4[(fi65 & int(v3))]
                                                            var v147 int8 = v52
                                                            _ = v146
                                                            _ = v147
                                                        } else {
                                                            v148 := GenS3{F0: 2147483647, F1: true}
                                                            var v149 int = fi65
                                                            _ = v148
                                                            _ = v149
                                                        }
                                                        var v150 uint8 = 0
                                                        v151 := GenS19{F0: 65535, F1: "\n\t", F2: int(v150)}
                                                        _ = v150
                                                        _ = v151
                                                    }
                                                }
                                            case 32767:
                                                switch (^(fi64)) {
                                                case -2147483648:
                                                    switch (fi65 - fi65) {
                                                    case -231:
                                                        v2 -= 3.14159
                                                    }
                                                case -541:
                                                    var v152 int64 = 42
                                                    v52 += (v52 + int8(v3))
                                                    _ = v152
                                                }
                                                v38 = append(v38, fi64)
                                                switch "foo bar" {
                                                case "hello":
                                                    v153 := GenS20{F0: 557, F1: (^(v3))}
                                                    _ = v153
                                                case "\n\t":
                                                    fmt.Println(v37, fi65, v5)
                                                default:
                                                    switch (-(fi65)) {
                                                    case -446:
                                                        v52 += (^(v52))
                                                    case -1:
                                                        switch 127 {
                                                        case -471:
                                                            var v154 uint = (v3 - v3)
                                                            _ = v154
                                                        case 128:
                                                            switch "foo bar" {
                                                            case "":
                                                                delete(v1, (("" + "") + ("\n\t" + (("hello" + "\n\t") + ("" + ("\n\t" + "foo bar"))))))
                                                            case "a":
                                                                v155 := GenS9{F0: v39, F1: 127}
                                                                _ = v155
                                                            case "foo bar":
                                                                v156 := GenS21{F0: 32767, F1: (true || v51), F2: (v39 - v39)}
                                                                var v157 float32 = 1e10
                                                                _ = v156
                                                                _ = v157
                                                            default:
                                                                v1["\n\t"] = v51
                                                                var v158 uint = 989
                                                                v52 = v52
                                                                _ = v158
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                var v159 bool = v51
                                                _ = v159
                                            }
                                            v5 -= 0.0
                                        }
                                        _ = v108
                                    } else {
                                        if (!v51) {
                                            fmt.Println(v3, v52, v5)
                                            var v160 int16 = -837
                                            _ = v160
                                        } else {
                                            var v161 uint8 = 255
                                            _ = v161
                                        }
                                        var v162 int8 = (v52 - (v52 * -51))
                                        if ((897 != 1) || (v51 || (v2 == (((1.5 + (v2 - -2.25)) - float64(-2.25)) / v5)))) {
                                            v39 = (v39 / ((v39) | 1))
                                            v162 += (v52 * 56)
                                        }
                                        fmt.Println(v2, v3)
                                        _ = v162
                                    }
                                }
                                v39 += genFn1()
                            } else {
                                _ = len(v38)
                                fmt.Println(v51, v3, v3)
                            }
                            var v163 int32 = 42
                            v2 = 0.1
                            _ = v163
                        }
                        switch ("foo bar" + (("a" + ("" + "\n\t")) + "hello")) {
                        case "hello":
                            v39 = (v39 ^ v39)
                        case "\n\t":
                            var v164 uint64 = 18446744073709551615
                            _ = v164
                        case "a":
                            var v165 float32 = 1.5
                            _ = v165
                        default:
                            if (9223372036854775807 < -9223372036854775808) {
                                switch fi65 {
                                case 127:
                                    var v166 int = 127
                                    _ = v166
                                case 853:
                                    v2 = (v5 + (v5 + v5))
                                case -1:
                                    v5 -= 0.1
                                default:
                                    fmt.Println(v52)
                                }
                            }
                        }
                    } else {
                        v167 := GenS22{F0: 0, F1: 128}
                        var v168 uint = 0
                        if ((v168 & v3) != (v3 & 474)) {
                            var v169 float64 = (0.0 + (v5 / (v2 - -2.25)))
                            v168 = v3
                            _ = v169
                        } else {
                            var v170 bool = (!false)
                            _ = v170
                        }
                        _ = v167
                        _ = v168
                    }
                    v39 += 42
                    _ = fi65
                }
                if true {
                    v3 -= 65535
                    if len(v38) >= 2 {
                        v38 = v38[1:len(v38)]
                    }
                } else {
                    v171 := GenS10{F0: 167, F1: 128}
                    var v172 uint8 = 187
                    v173 := GenS23{F0: (v52 > v52)}
                    _ = v171
                    _ = v172
                    _ = v173
                }
            }
            _ = fi64
        }
        _ = v63
    }
    var v174 int = (v37 << ((299) & 63))
    for fi175 := 0; fi175 < 1; fi175++ {
        var v176 int = (v37 % ((fi175) | 1))
        v174 -= (fi175 ^ -2147483648)
        var v177 int32 = 0
        _ = fi175
        _ = v176
        _ = v177
    }
    switch "" {
    case "hello":
        if ((v51 || v51) && true) {
            var v178 uint64 = 4294967295
            var v179 int32 = 1
            v1[(("a" + ("\n\t" + "")) + "hello")] = v51
            _ = v178
            _ = v179
        } else {
            fmt.Println(v3, v174)
        }
    }
    _ = len(v4)
    switch v174 {
    case 9223372036854775807:
        var v180 uint16 = (v39 & (^(v39)))
        _ = v180
    case 127:
        if (!true) {
            v181 := v1[("hello" + "\n\t")]
            _ = v181
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v174
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
