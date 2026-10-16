package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int32; F2 int16 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 bool; F2 string }
type GenS10 struct { F0 string; F1 int32; F2 string }
type GenS11 struct { F0 uint8; F1 int8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint16; F1 uint32; F2 string }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint8; F1 int64 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint8; F1 int32 }
type GenS21 struct { F0 int8; F1 uint16; F2 int32 }
func genFn0(p0 uint8) int64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = 0
        _ = fi5
        _ = v6
    }
    if false {
        fmt.Println(p0)
        p0 += 128
    } else {
        switch 0 {
        case 127:
            p0 += 0
        case -601:
            fmt.Println(p0)
        case 1:
            var v7 int8 = 0
            _ = v7
        default:
            fmt.Println(p0)
        }
    }
    switch "a" {
    case "foo bar":
        v8 := []int{-9223372036854775808}
        _ = v8
    default:
        v9 := []int{4294967295, -1, -9223372036854775808, 0}
        _ = v9
    }
    fmt.Println(p0)
    p0 -= p0
    _ = p0
    return 786
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{32767, 67, 42}
    var v3 float64 = float64(3.14159)
    var v4 float64 = (v3 / float64(3.14159))
    switch 32767 {
    case 31:
        v10 := GenS1{F0: (1 == 182), F1: 774}
        _ = v10
    case -1:
        for fi11 := 0; fi11 < 16; fi11++ {
            switch "" {
            case "a":
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        v12 := GenS2{F0: -806, F1: (2 <= 1)}
                        _ = v12
                    case "":
                        var v13 int16 = -1
                        _ = v13
                    case "a":
                        var v14 float64 = v3
                        _ = v14
                    default:
                        var v15 bool = true
                        _ = v15
                    }
                    var v16 int32 = 117
                    _ = v16
                } else {
                    v3 = v3
                    _ = len(v2)
                }
            case "foo bar":
                if false {
                    v17 := GenS3{F0: -32768}
                    _ = v17
                }
            case "":
                if (!((42 < 120) && (100 == 1))) {
                    var v18 uint = 42
                    var v19 int64 = 9223372036854775807
                    for fi20 := 0; fi20 < 6; fi20++ {
                        v19 -= (-(v19))
                        v4 -= float64((float32(1e10) * (float32(v4) / 0.1)))
                        v21 := GenS4{F0: 1, F1: 535}
                        fmt.Println(v4)
                        _ = fi20
                        _ = v21
                    }
                    _ = v18
                    _ = v19
                } else {
                    var v22 string = "a"
                    _ = v22
                }
            default:
                for fi23 := 0; fi23 < 5; fi23++ {
                    fmt.Println(fi11, v3, v4)
                    v4 += float64(3.14159)
                    v24 := GenS2{F0: -587, F1: ((true || false) && true)}
                    _ = fi23
                    _ = v24
                }
            }
            switch "hello" {
            case "a":
                v25 := GenS5{F0: 1000}
                _ = v25
            case "hello":
                v26 := GenS6{F0: 42, F1: -2147483648, F2: 0}
                _ = v26
            case "foo bar":
                switch "a" {
                case "foo bar":
                    v4 = -2.25
                default:
                    for fi27 := 0; fi27 < 1; fi27++ {
                        var v28 uint64 = 1
                        _ = fi27
                        _ = v28
                    }
                }
            }
            _ = fi11
        }
    default:
        var v29 string = ("hello" + "a")
        _ = v29
    }
    switch ("a" + "\n\t") {
    case "a":
        if false {
            v30 := v1[""]
            if true {
                var v31 uint = 791
                _ = v31
            }
            if false {
                switch v30 {
                case -249:
                    var v32 int8 = 15
                    _ = v32
                case 42:
                    v33 := GenS7{F0: 0, F1: (18446744073709551615 != 0)}
                    _ = v33
                }
                switch 1 {
                case 4294967295:
                    v34 := GenS7{F0: -97, F1: false}
                    _ = v34
                default:
                    var v35 float32 = (((-2.25 * 0.1) / (0.1 + 0.1)) / (float32(float64(1.5)) / float32(float64((1.5 - ((float32(0.0) / ((1e10 - (3.14159 + 0.1)) + float32(v3))) / ((-2.25 * 1.5) - 0.1)))))))
                    _ = v35
                }
            }
            _ = v30
        }
    case "hello":
        v4 = (float64((float32(v3) * ((0.1 * (1e10 + (1e10 + -2.25))) + 0.1))) + v4)
    }
    var v36 float32 = float32(v3)
    for fi37 := 0; fi37 < 20; fi37++ {
        if true {
            switch fi37 {
            case -9223372036854775808:
                v36 = ((float32(-2.25) / v36) / (3.14159 / float32(v4)))
            case 1:
                switch fi37 {
                case 492:
                    var v38 bool = (!true)
                    _ = v38
                case 2147483647:
                    v3 += v3
                default:
                    switch ("\n\t" + ("" + "")) {
                    case "a":
                        fmt.Println(v3)
                    case "\n\t":
                        for fi39 := 0; fi39 < 6; fi39++ {
                            v40 := v1["a"]
                            if true {
                                v41 := GenS4{F0: 182, F1: 1}
                                var v42 uint32 = 98
                                v42 = v42
                                _ = v41
                                _ = v42
                            } else {
                                v4 = 3.14159
                                switch "" {
                                case "foo bar":
                                    if (98 == 670) {
                                        var v43 uint16 = 10
                                        var v44 float32 = 0.0
                                        v45 := GenS8{F0: 476}
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                case "a":
                                    var v46 int8 = -128
                                    _ = v46
                                case "":
                                    v4 = (v3 - 0.1)
                                }
                                v47 := GenS9{F0: 255, F1: true, F2: "hello"}
                                _ = v47
                            }
                            _ = fi39
                            _ = v40
                        }
                    }
                }
            case -21:
                if (42 != 1) {
                    v3 -= (3.14159 - (0.1 + ((-2.25 - (v3 * v4)) - 0.0)))
                    for fi48 := 0; fi48 < 10; fi48++ {
                        fmt.Println(v36)
                        var v49 int16 = 42
                        _ = fi48
                        _ = v49
                    }
                } else {
                    switch "a" {
                    case "hello":
                        v50 := GenS8{F0: 1}
                        _ = v50
                    case "":
                        for fi51 := 0; fi51 < 2; fi51++ {
                            v36 = float32(v4)
                            v4 = ((((1.5 / (v3 + v3)) - (((1.5 * 1e10) + ((0.1 * ((1.5 + (1.5 * 3.14159)) * (v4 / float64((((v36 - v36) * float32((v4 + v3))) + 1.5))))) - v4)) - 1.5)) * 3.14159) * 3.14159)
                            var v52 int64 = 140
                            v53 := GenS9{F0: 42, F1: true, F2: "foo bar"}
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    default:
                        if true {
                            if (1 < 2147483647) {
                                var v54 int8 = 127
                                switch (fi37 & -568) {
                                case -1:
                                    for fi55 := 0; fi55 < 11; fi55++ {
                                        v1["\n\t"] = fi37
                                        var v56 int = fi37
                                        if (false && ((fi55 >= (fi37 << ((1) & 63))) || false)) {
                                            var v57 float32 = 0.0
                                            fmt.Println(v36, v54)
                                            _ = v57
                                        } else {
                                            var v58 int64 = -1
                                            v56 += -584
                                            v54 += v54
                                            switch ("\n\t" + "") {
                                            case "":
                                                fmt.Println(v58)
                                            case "\n\t":
                                                v2 = append(v2, (fi37 << ((65535) & 63)))
                                            }
                                            v59 := GenS6{F0: v54, F1: -2147483648, F2: -1}
                                            v58 = v58
                                            _ = v58
                                            _ = v59
                                        }
                                        _ = fi55
                                        _ = v56
                                    }
                                case -182:
                                    for fi60 := 0; fi60 < 6; fi60++ {
                                        var v61 float32 = float32(v3)
                                        if (42 < 18446744073709551615) {
                                            var v62 uint8 = 0
                                            if (-1 < 0) {
                                                var v63 uint8 = 2
                                                if false {
                                                    if (v63 != (v63 | v62)) {
                                                        var v64 uint64 = 758
                                                        if (18446744073709551615 == 953) {
                                                            if len(v2) > 0 {
                                                                v2[0] = fi60
                                                            }
                                                            v3 += v4
                                                        } else {
                                                            v54 = v54
                                                        }
                                                        v36 = ((((v61 / -2.25) / (float32(v4) * v36)) / (v61 / (((float32(float64(0.1)) / (v36 * (v36 - (v36 - v61)))) - v61) / v36))) / 0.1)
                                                        _ = v64
                                                    }
                                                    if false {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        switch fi37 {
                                                        case -189:
                                                            if false {
                                                                v65 := GenS2{F0: -1, F1: false}
                                                                _ = v65
                                                            }
                                                        case 0:
                                                            if false {
                                                                fmt.Println(v62, v4, v4)
                                                                v66 := GenS10{F0: (("\n\t" + "\n\t") + ""), F1: -2147483648, F2: "foo bar"}
                                                                _ = v66
                                                            } else {
                                                                v62 = 127
                                                                if true {
                                                                    v61 = v61
                                                                }
                                                            }
                                                        default:
                                                            switch ("foo bar" + (("foo bar" + "foo bar") + ("a" + ("" + ("hello" + ((("foo bar" + "") + ("hello" + "foo bar")) + "\n\t")))))) {
                                                            case "\n\t":
                                                                v54 = (-(v54))
                                                            case "foo bar":
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v67 := GenS11{F0: v62, F1: v54}
                                                                    var v68 int64 = 0
                                                                    if (!true) {
                                                                        var v69 uint8 = 126
                                                                        var v70 string = "hello"
                                                                        v71 := GenS11{F0: (v63 / ((v62) | 1)), F1: (v54 & -2)}
                                                                        _ = len(v2)
                                                                        fmt.Println(fi37, fi60, v69)
                                                                        v72 := GenS12{F0: (!((711 > 65535) || false))}
                                                                        if len(v2) > 0 {
                                                                            v2[0] = fi37
                                                                        }
                                                                        _ = v69
                                                                        _ = v70
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                    fmt.Println(v61, v3, v4)
                                                                    _ = v67
                                                                    _ = v68
                                                                case "foo bar":
                                                                    if ((((fi37 % ((42) | 1)) < fi37) || (9223372036854775807 == 0)) && true) {
                                                                        var v73 int8 = (v54 - (v54 & (v54 / (((v54 ^ v54)) | 1))))
                                                                        v61 += (0.0 / v36)
                                                                        _ = v73
                                                                    } else {
                                                                        if true {
                                                                            v74 := GenS1{F0: (0 >= -32768), F1: 1}
                                                                            _ = v74
                                                                        } else {
                                                                            v75 := GenS4{F0: 255, F1: -1}
                                                                            switch fi37 {
                                                                            case 0:
                                                                                var v76 string = "foo bar"
                                                                                _ = v76
                                                                            case -9223372036854775808:
                                                                                var v77 uint16 = 32767
                                                                                _ = v77
                                                                            default:
                                                                                fmt.Println(v61)
                                                                            }
                                                                            v63 = 33
                                                                            v1["\n\t"] = fi60
                                                                            if (42 <= 2) {
                                                                                v54 = (v54 | v54)
                                                                            }
                                                                            _ = v75
                                                                        }
                                                                        var v78 uint = 0
                                                                        _ = v78
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v79 := GenS1{F0: true, F1: 32767}
                                                                        v80 := GenS2{F0: -867, F1: true}
                                                                        v36 = v36
                                                                        _ = v79
                                                                        _ = v80
                                                                    }
                                                                }
                                                            case "a":
                                                                v4 += (v3 - v4)
                                                            }
                                                        }
                                                    }
                                                    if false {
                                                        var v81 bool = true
                                                        switch "hello" {
                                                        case "a":
                                                            var v82 int = fi37
                                                            _ = v82
                                                        case "hello":
                                                            switch "" {
                                                            case "a":
                                                                fmt.Println(v3, v62)
                                                            case "":
                                                                var v83 float64 = v3
                                                                _ = v83
                                                            case "hello":
                                                                fmt.Println(v63, v4)
                                                            default:
                                                                v61 = v61
                                                            }
                                                            var v84 uint = 795
                                                            var v85 uint16 = 489
                                                            switch ("" + ("" + "a")) {
                                                            case "":
                                                                var v86 uint16 = v85
                                                                var v87 uint = v84
                                                                _ = v86
                                                                _ = v87
                                                            case "foo bar":
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi37
                                                                }
                                                            case "\n\t":
                                                                v3 = ((0.1 * ((float64((float32(1.5) - (float32(v4) + -2.25))) * (v4 * v3)) * 3.14159)) - 1.5)
                                                            default:
                                                                var v88 int8 = v54
                                                                _ = v88
                                                            }
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        _ = v81
                                                    }
                                                }
                                                v89 := GenS1{F0: true, F1: 0}
                                                _ = v63
                                                _ = v89
                                            }
                                            v3 = v3
                                            v4 += (1.5 * float64(float32(0.0)))
                                            switch "hello" {
                                            case "foo bar":
                                                v36 = float32(1e10)
                                            case "hello":
                                                fmt.Println(v3, fi60)
                                            }
                                            _ = v62
                                        }
                                        var v90 int16 = 222
                                        _ = fi60
                                        _ = v61
                                        _ = v90
                                    }
                                }
                                for fi91 := 0; fi91 < 17; fi91++ {
                                    v92 := GenS13{F0: true}
                                    v93 := GenS14{F0: 1, F1: 42, F2: 128}
                                    var v94 uint64 = 2
                                    _ = fi91
                                    _ = v92
                                    _ = v93
                                    _ = v94
                                }
                                _ = v54
                            } else {
                                var v95 uint8 = 2
                                for fi96 := 0; fi96 < 8; fi96++ {
                                    switch (-(fi96)) {
                                    case 380:
                                        delete(v1, "hello")
                                    }
                                    v2 = append(v2, (^(fi37)))
                                    _ = fi96
                                }
                                v3 = v3
                                _ = v95
                            }
                        }
                    }
                    switch (fi37 % (((fi37 + (fi37 >> ((714) & 63)))) | 1)) {
                    case -314:
                        if false {
                            for fi97 := 0; fi97 < 15; fi97++ {
                                fmt.Println(fi37)
                                v3 -= v3
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v98 := GenS6{F0: 100, F1: 0, F2: 385}
                                _ = fi97
                                _ = v98
                            }
                            v99 := GenS15{F0: 65535}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v99
                        } else {
                            switch (fi37 ^ (fi37 | 4294967295)) {
                            case 4294967295:
                                fmt.Println(v36, fi37)
                            default:
                                v4 = ((v4 / v3) - ((1.5 / v4) + (1e10 / ((float64(3.14159) * float64(v36)) / 0.1))))
                            }
                            v100 := GenS6{F0: 0, F1: 42, F2: 32767}
                            _ = v100
                        }
                    case 127:
                        v4 -= 0.1
                    case 4294967295:
                        if (800 < 1) {
                            for fi101 := 0; fi101 < 1; fi101++ {
                                fmt.Println(v3)
                                _ = fi101
                            }
                        } else {
                            var v102 uint64 = 2
                            switch 476 {
                            case -9223372036854775808:
                                fmt.Println(fi37, v3, v4)
                            case 32768:
                                _ = len(v2)
                            case 678:
                                v103 := GenS5{F0: -32768}
                                _ = v103
                            default:
                                var v104 float32 = v36
                                _ = v104
                            }
                            if (!(42 != 0)) {
                                v2 = append(v2, (fi37 - fi37))
                                v102 = (v102 % ((v102) | 1))
                            } else {
                                var v105 uint64 = (v102 + (v102 >> ((65535) & 63)))
                                v106 := GenS7{F0: -84, F1: (255 == 1)}
                                _ = v105
                                _ = v106
                            }
                            _ = v102
                        }
                    }
                }
            }
            if (!(0 <= 4294967295)) {
                if (!false) {
                    v3 = v4
                    v107 := GenS10{F0: (("hello" + "") + "foo bar"), F1: -986, F2: ""}
                    var v108 uint = 486
                    _ = v107
                    _ = v108
                } else {
                    var v109 float64 = v3
                    var v110 float64 = 3.14159
                    v111 := GenS15{F0: 425}
                    _ = v109
                    _ = v110
                    _ = v111
                }
            }
            fmt.Println(v4, fi37)
        } else {
            _ = len(v1)
            switch (-(fi37)) {
            case 0:
                for fi112 := 0; fi112 < 2; fi112++ {
                    fmt.Println(fi112)
                    _ = fi112
                }
            }
        }
        _ = fi37
    }
    v113 := GenS16{F0: 414, F1: 914, F2: (("a" + ("\n\t" + "")) + "")}
    if false {
        for fi114 := 0; fi114 < 17; fi114++ {
            v115 := GenS17{F0: 255}
            fmt.Println(fi114, v36, v3)
            _ = fi114
            _ = v115
        }
        v116 := GenS13{F0: (true || false)}
        _ = v116
    } else {
        v3 = float64(-2.25)
        delete(v1, "a")
        for fi117 := 0; fi117 < 20; fi117++ {
            v118 := GenS5{F0: -773}
            switch ((("hello" + "a") + ("\n\t" + ("hello" + "\n\t"))) + "a") {
            case "hello":
                v119 := GenS18{F0: 128, F1: 0}
                _ = v119
            case "\n\t":
                if (1 < 2) {
                    if (-710 != -9223372036854775808) {
                        if true {
                            for fi120 := 0; fi120 < 15; fi120++ {
                                if len(v2) > 0 {
                                    v2[0] = fi120
                                }
                                _ = fi120
                            }
                            v2 = append(v2, fi117)
                        } else {
                            _ = len(v2)
                            v3 -= (v4 + (0.1 + v4))
                            var v121 float32 = (-2.25 / -2.25)
                            _ = v121
                        }
                        if true {
                            v122 := v1[("foo bar" + "a")]
                            _ = v122
                        }
                        v36 = (3.14159 - v36)
                    }
                    if len(v2) > 0 {
                        v2[0] = (^(fi117))
                    }
                    if (false && true) {
                        if len(v2) > 0 {
                            v2[0] = fi117
                        }
                        switch ("hello" + "hello") {
                        case "a":
                            if ((fi117 & 127) == 4294967295) {
                                delete(v1, ("" + "\n\t"))
                            } else {
                                var v123 float32 = 1e10
                                _ = v123
                            }
                        default:
                            v4 = ((v4 + (float64(0.0) * -2.25)) * v3)
                        }
                        for fi124 := 0; fi124 < 16; fi124++ {
                            v4 += float64(float32(0.0))
                            v125 := v1[("hello" + "")]
                            _ = fi124
                            _ = v125
                        }
                    }
                } else {
                    fmt.Println(v36)
                }
            case "a":
                v3 = v3
            default:
                switch (fi117 >> ((42) & 63)) {
                case 32767:
                    v3 += v3
                case 42:
                    switch "" {
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        v126 := GenS19{F0: 32767}
                        _ = v126
                    }
                default:
                    var v127 string = "hello"
                    _ = v127
                }
            }
            _ = fi117
            _ = v118
        }
    }
    for fi128 := 0; fi128 < 16; fi128++ {
        var v129 float64 = v4
        for fi130 := 0; fi130 < 20; fi130++ {
            var v131 uint16 = 493
            _ = fi130
            _ = v131
        }
        for fi132 := 0; fi132 < 5; fi132++ {
            v3 = 0.0
            var v133 int64 = 1
            v129 -= (0.1 * (float64((float32(float64((3.14159 / float32(1e10)))) * v36)) * 0.1))
            v36 -= v36
            var v134 uint8 = 0
            _ = fi132
            _ = v133
            _ = v134
        }
        _ = fi128
        _ = v129
    }
    for fi135 := 0; fi135 < 3; fi135++ {
        v3 -= v3
        _ = len(v2)
        if (("" + "") < "hello") {
            var v136 uint = 0
            _ = v136
        } else {
            switch "a" {
            case "a":
                switch fi135 {
                case 974:
                    var v137 int32 = 2147483647
                    _ = v137
                case 32767:
                    v36 = 1e10
                }
            default:
                var v138 int64 = 42
                _ = v138
            }
        }
        _ = fi135
    }
    for fi139 := 0; fi139 < 16; fi139++ {
        var v140 bool = (170 == 825)
        switch ("foo bar" + "") {
        case "\n\t":
            var v141 int32 = 1
            _ = v141
        case "":
            var v142 int32 = 42
            _ = v142
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v4 = 0.0
        }
        for fi143 := 0; fi143 < 17; fi143++ {
            v3 = 3.14159
            fmt.Println(v140)
            v3 -= (1e10 / float64((1e10 / 1e10)))
            _ = fi143
        }
        _ = fi139
        _ = v140
    }
    v3 = v3
    for fi144 := 0; fi144 < 17; fi144++ {
        var v145 uint8 = 1
        _ = fi144
        _ = v145
    }
    v146 := GenS20{F0: 127, F1: 280}
    v147 := GenS21{F0: -112, F1: 2, F2: -1}
    switch "\n\t" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
    case "hello":
        var v148 string = ("a" + ("" + ""))
        _ = v148
    default:
        v36 += ((v36 + v36) * (((v36 / 1e10) / float32(v4)) + 0.0))
    }
    v4 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v113
    _ = v146
    _ = v147
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
