package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 uint8; F2 int }
type GenS3 struct { F0 int64; F1 uint16; F2 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint64; F1 uint64; F2 int }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 uint32; F1 string; F2 int }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 int; F1 int32; F2 bool }
type GenS12 struct { F0 string; F1 uint16; F2 int }
type GenS13 struct { F0 int32; F1 int8; F2 int }
type GenS14 struct { F0 int; F1 int; F2 uint }
type GenS15 struct { F0 int32; F1 int8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 bool; F1 uint }
type GenS19 struct { F0 bool; F1 int64; F2 int8 }
type GenS20 struct { F0 uint64; F1 uint }
type GenS21 struct { F0 uint; F1 bool; F2 uint32 }
type GenS22 struct { F0 int64; F1 int16 }
type GenS23 struct { F0 uint64; F1 int64; F2 string }
type GenS24 struct { F0 int; F1 int64; F2 uint }
type GenS25 struct { F0 int32; F1 int64 }
type GenS26 struct { F0 int; F1 uint32; F2 bool }
func genFn0(p0 int32, p1 bool, p2 uint64) bool {
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{}
        switch ("\n\t" + ("a" + "foo bar")) {
        case "a":
            var v9 uint64 = 859
            _ = v9
        case "":
            switch "foo bar" {
            case "a":
                for fi10 := 0; fi10 < 15; fi10++ {
                    v8 = append(v8, (-(fi10)))
                    var v11 int32 = 1
                    switch "" {
                    case "foo bar":
                        p2 -= p2
                    case "a":
                        var v12 uint8 = 129
                        _ = v12
                    }
                    _ = len(v8)
                    fmt.Println(fi7)
                    _ = len(v8)
                    _ = fi10
                    _ = v11
                }
            default:
                v13 := GenS1{F0: 173, F1: int(p2)}
                _ = v13
            }
        }
        _ = fi7
        _ = v8
    }
    p2 += p2
    v14 := []int{}
    for fi15 := 0; fi15 < 6; fi15++ {
        p1 = (p1 && p1)
        fmt.Println(p1, fi15, p1)
        _ = fi15
    }
    switch (("a" + ("foo bar" + ("" + (("" + "") + "\n\t")))) + "a") {
    case "\n\t":
        v14 = append(v14, -96)
    case "foo bar":
        if p1 {
            var v16 int8 = 1
            if (893 >= 4294967295) {
                switch "foo bar" {
                case "foo bar":
                    v17 := GenS2{F0: 258, F1: 42, F2: 2147483647}
                    _ = v17
                case "":
                    for fi18 := 0; fi18 < 2; fi18++ {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 int32 = (p0 ^ (p0 << ((4294967295) & 63)))
                            v21 := GenS3{F0: -9223372036854775808, F1: 32767, F2: (fi19 & 4294967295)}
                            _ = len(v14)
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS2{F0: (fi18 % (((fi18 - (fi18 & int(p0)))) | 1)), F1: 127, F2: 2147483647}
                        _ = fi18
                        _ = v22
                    }
                case "hello":
                    var v23 int32 = p0
                    _ = v23
                }
                switch "foo bar" {
                case "hello":
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
                switch -136 {
                case -9223372036854775808:
                    v14 = append(v14, -9223372036854775808)
                case -514:
                    var v24 uint64 = (p2 / (((p2 * 227)) | 1))
                    _ = v24
                case -2147483648:
                    p0 += (p0 % ((p0) | 1))
                }
            }
            v16 -= v16
            _ = v16
        }
    case "hello":
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS4{F0: (p2 | (^(p2)))}
            _ = fi25
            _ = v26
        }
    default:
        if len(v14) > 0 {
            v14[0] = 9223372036854775807
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return false
}
func genFn1() int8 {
    var v27 uint8 = 128
    if ((0.0 * 3.14159) < float64(0.1)) {
        v28 := []int{1, 67, 32768}
        for fi29 := 0; fi29 < 18; fi29++ {
            if len(v28) > 0 {
                v28[0] = (fi29 * (fi29 << ((930) & 63)))
            }
            v30 := GenS5{F0: false}
            _ = fi29
            _ = v30
        }
        _ = v28
    } else {
        for fi31 := 0; fi31 < 14; fi31++ {
            for fi32 := 0; fi32 < 11; fi32++ {
                v27 += v27
                v33 := GenS6{F0: "", F1: uint(v27), F2: 4294967295}
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
        for fi34 := 0; fi34 < 14; fi34++ {
            v35 := GenS1{F0: 242, F1: -2147483648}
            for fi36 := 0; fi36 < 8; fi36++ {
                var v37 int32 = 82
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                if false {
                    v27 -= v27
                    v27 = (v27 - (^(v27)))
                }
                _ = fi38
            }
            _ = fi34
            _ = v35
        }
    }
    var v39 uint8 = 158
    if true {
        var v40 int = 1
        v41 := GenS3{F0: 1, F1: 65535, F2: v40}
        _ = v40
        _ = v41
    } else {
        var v42 float32 = 1.5
        v27 += 55
        _ = v42
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        v44 := []int{fi43, (fi43 | fi43), fi43, (fi43 ^ (fi43 * 128))}
        var v45 int = fi43
        _ = fi43
        _ = v44
        _ = v45
    }
    _ = v27
    _ = v39
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 float64 = 1e10
    v4 := []string{}
    var v5 float64 = (((0.0 / (v3 - (0.1 + 3.14159))) / 1.0) / (1.5 + ((1e10 + (v3 / v3)) / (3.14159 * float64(0.1)))))
    v6 := make(map[uint32]bool)
    v5 = 1e10
    v4 = append(v4, (("hello" + ("\n\t" + "foo bar")) + "a"))
    for fi46 := 0; fi46 < 14; fi46++ {
        v5 -= v3
        _ = fi46
    }
    for fi47 := 0; fi47 < 9; fi47++ {
        if (3.14159 == 3.14159) {
            for fi48 := 0; fi48 < 7; fi48++ {
                if true {
                    v3 = v3
                    v5 += v5
                    v5 += v3
                    var v49 uint8 = 42
                    if true {
                        v3 -= float64(0.0)
                        if false {
                            v5 += v5
                        }
                        if (!(18446744073709551615 == 2)) {
                            var v50 float32 = 0.0
                            var v51 int32 = -1
                            v51 += v51
                            _ = v50
                            _ = v51
                        } else {
                            v5 = -2.25
                            v52 := GenS7{F0: 18446744073709551615, F1: 2, F2: (fi47 & (fi47 / ((2147483647) | 1)))}
                            v49 = (v49 >> ((2) & 63))
                            _ = v52
                        }
                    }
                    _ = v49
                }
                var v53 int16 = -626
                v53 -= v53
                _ = fi48
                _ = v53
            }
            switch fi47 {
            case 127:
                v3 += (v3 / (v5 + (v3 + 3.14159)))
            default:
                for fi54 := 0; fi54 < 13; fi54++ {
                    v3 = 0.1
                    _ = fi54
                }
            }
        }
        for fi55 := 0; fi55 < 13; fi55++ {
            v56 := GenS1{F0: 171, F1: fi47}
            v57 := GenS3{F0: 231, F1: 255, F2: -1}
            _ = fi55
            _ = v56
            _ = v57
        }
        switch "foo bar" {
        case "foo bar":
            if (false && (42 != 0)) {
                var v58 uint64 = 4294967295
                v3 -= -2.25
                _ = v58
            }
        default:
            _ = len(v2)
        }
        _ = fi47
    }
    if (0.0 < 3.14159) {
        switch "hello" {
        case "":
            switch 32768 {
            case 32768:
                delete(v1, 953)
            case 493:
                if true {
                    switch "" {
                    case "a":
                        if (((!(0.0 >= 1.5)) || (1 < 1)) || (1e10 < -2.25)) {
                            v59 := GenS8{F0: 65535, F1: false}
                            v60 := GenS4{F0: 1}
                            _ = v59
                            _ = v60
                        } else {
                            var v61 uint = 17
                            var v62 uint16 = 0
                            if true {
                                v63 := GenS2{F0: 42, F1: 42, F2: 32768}
                                for fi64 := 0; fi64 < 2; fi64++ {
                                    v65 := GenS9{F0: uint32(v62), F1: "", F2: fi64}
                                    for fi66 := 0; fi66 < 18; fi66++ {
                                        v3 += 1.5
                                        v4 = append(v4, "")
                                        if (!(false && (255 >= 0))) {
                                            var v67 int32 = -1
                                            v5 = (v3 + v5)
                                            _ = v67
                                        }
                                        _ = fi66
                                    }
                                    _ = fi64
                                    _ = v65
                                }
                                v4 = append(v4, "a")
                                _ = v63
                            }
                            _ = v61
                            _ = v62
                        }
                    case "hello":
                        var v68 uint32 = 65535
                        _ = v68
                    case "":
                        if true {
                            if true {
                                var v69 float32 = (-2.25 - 1.5)
                                if (1 <= -1) {
                                    v70 := GenS10{F0: false, F1: 42}
                                    if (215 < 208) {
                                        if (!true) {
                                            if true {
                                                v71 := GenS11{F0: 1, F1: 108, F2: false}
                                                var v72 int = -961
                                                _ = v71
                                                _ = v72
                                            }
                                            if false {
                                                var v73 uint64 = 64
                                                _ = v73
                                            }
                                            switch 0 {
                                            case -634:
                                                for fi74 := 0; fi74 < 7; fi74++ {
                                                    fmt.Println(v69, v3)
                                                    if (("hello" + "foo bar") >= "foo bar") {
                                                        var v75 uint16 = 384
                                                        v76 := GenS10{F0: ((!(v69 < (float32(((v5 / v3) * float64(3.14159))) + 1e10))) || true), F1: 2}
                                                        _ = v75
                                                        _ = v76
                                                    } else {
                                                        for fi77 := 0; fi77 < 15; fi77++ {
                                                            if (genFn1() < 0) {
                                                                v5 += 3.14159
                                                                v5 -= v5
                                                            }
                                                            _ = fi77
                                                        }
                                                        var v78 float64 = v5
                                                        _ = v78
                                                    }
                                                    if true {
                                                        fmt.Println(fi74, v5)
                                                        fmt.Println(fi74)
                                                        fmt.Println(v69)
                                                    } else {
                                                        v79 := v1[1]
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v79
                                                    }
                                                    _ = fi74
                                                }
                                            case 4294967295:
                                                switch -808 {
                                                case -2147483648:
                                                    fmt.Println(v3, v5, v3)
                                                case 1:
                                                    v69 += (((float32(v3) * v69) - v69) / v69)
                                                case -557:
                                                    switch "" {
                                                    case "foo bar":
                                                        v80 := GenS12{F0: "foo bar", F1: 1, F2: -2147483648}
                                                        _ = v80
                                                    case "a":
                                                        if false {
                                                            for fi81 := 0; fi81 < 11; fi81++ {
                                                                fmt.Println(v69, v69, v3)
                                                                for fi82 := 0; fi82 < 3; fi82++ {
                                                                    v69 = (0.1 / v69)
                                                                    _ = fi82
                                                                }
                                                                var v83 int64 = -1
                                                                _ = fi81
                                                                _ = v83
                                                            }
                                                            _ = len(v2)
                                                            var v84 int16 = 0
                                                            _ = v84
                                                        } else {
                                                            var v85 uint64 = 25
                                                            var v86 bool = (-2147483648 == -2147483648)
                                                            switch ("a" + "a") {
                                                            case "a":
                                                                var v87 int64 = 0
                                                                _ = v87
                                                            case "foo bar":
                                                                v88 := GenS13{F0: 42, F1: genFn1(), F2: 2147483647}
                                                                _ = v88
                                                            case "\n\t":
                                                                if (!(2147483647 >= 42)) {
                                                                    var v89 uint64 = (v85 | 1)
                                                                    for fi90 := 0; fi90 < 9; fi90++ {
                                                                        v91 := GenS5{F0: (!(v85 < v85))}
                                                                        _ = fi90
                                                                        _ = v91
                                                                    }
                                                                    _ = v89
                                                                } else {
                                                                    v92 := v6[1]
                                                                    _ = v92
                                                                }
                                                            }
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                    }
                                                default:
                                                    v3 += 3.14159
                                                }
                                            case 32768:
                                                v5 -= (v5 + v5)
                                            default:
                                                if (false && false) {
                                                    switch 159 {
                                                    case 32768:
                                                        v69 -= (v69 - v69)
                                                    case -54:
                                                        v93 := GenS5{F0: true}
                                                        _ = v93
                                                    }
                                                } else {
                                                    switch -1 {
                                                    case -9223372036854775808:
                                                        if (-299 < 42) {
                                                            var v94 int = -136
                                                            _ = len(v6)
                                                            _ = len(v2)
                                                            _ = v94
                                                        } else {
                                                            fmt.Println(v3, v5, v69)
                                                        }
                                                    }
                                                    for fi95 := 0; fi95 < 4; fi95++ {
                                                        if true {
                                                            v4 = append(v4, "hello")
                                                            switch (-(fi95)) {
                                                            case -758:
                                                                v3 = v5
                                                            case 4294967295:
                                                                v96 := GenS3{F0: 0, F1: 255, F2: fi95}
                                                                _ = v96
                                                            case -1:
                                                                v69 = v69
                                                            }
                                                        } else {
                                                            for fi97 := 0; fi97 < 14; fi97++ {
                                                                if true {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                } else {
                                                                    var v98 uint8 = 62
                                                                    v99 := GenS14{F0: (^(fi95)), F1: (fi97 ^ 262), F2: 145}
                                                                    if false {
                                                                        v69 += (((v69 + v69) - ((v69 / (0.1 - 0.0)) / (v69 * v69))) - 3.14159)
                                                                    }
                                                                    if true {
                                                                        var v100 float64 = (3.14159 / v5)
                                                                        var v101 int8 = 85
                                                                        _ = v100
                                                                        _ = v101
                                                                    } else {
                                                                        fmt.Println(v69)
                                                                        var v102 int64 = 985
                                                                        v102 -= (v102 >> ((0) & 63))
                                                                        _ = v102
                                                                    }
                                                                    var v103 int = (fi97 + fi95)
                                                                    v104 := GenS15{F0: 0, F1: 42}
                                                                    switch "" {
                                                                    case "":
                                                                        switch (^(v103)) {
                                                                        case 128:
                                                                            fmt.Println(v69, fi95, v3)
                                                                            v105 := GenS16{F0: 9223372036854775807}
                                                                            _ = v105
                                                                        case 32767:
                                                                            if false {
                                                                                var v106 int = fi97
                                                                                v3 -= v3
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = ("hello" + "a")
                                                                                }
                                                                                _ = v106
                                                                            } else {
                                                                                v69 = (0.1 - v69)
                                                                                var v107 float64 = v3
                                                                                v3 = (((v5 * 0.0) - v3) + v5)
                                                                                if (42 != -1) {
                                                                                    v2[(fi97 << ((2) & 63))] = (v103 | 165)
                                                                                }
                                                                                switch fi95 {
                                                                                case 387:
                                                                                    v3 = v5
                                                                                case 0:
                                                                                    if (829 == 1) {
                                                                                        v98 += v98
                                                                                        v108 := GenS12{F0: "", F1: 255, F2: (fi97 ^ (fi97 ^ v103))}
                                                                                        _ = v108
                                                                                    } else {
                                                                                        v109 := GenS7{F0: 2, F1: 2, F2: v103}
                                                                                        _ = v109
                                                                                    }
                                                                                default:
                                                                                    v110 := GenS17{F0: 4294967295}
                                                                                    _ = v110
                                                                                }
                                                                                _ = v107
                                                                            }
                                                                        }
                                                                    }
                                                                    switch (fi97 - fi97) {
                                                                    case 128:
                                                                        if (!(false && (!true))) {
                                                                            _ = len(v2)
                                                                            var v111 int64 = -189
                                                                            v3 -= ((v5 - float64(-2.25)) + (((((0.1 / float64(((float32((v5 / (v3 + float64(v69)))) / v69) / (((v69 - float32(v5)) - float32((1.5 + 0.1))) + ((v69 / (v69 * -2.25)) + (v69 - (-2.25 + (3.14159 * 1.5)))))))) - v3) / 1.0) + v3) + 0.0))
                                                                            if (v98 == v98) {
                                                                                v112 := GenS18{F0: false, F1: 0}
                                                                                _ = len(v4)
                                                                                v98 = 128
                                                                                _ = v112
                                                                            }
                                                                            var v113 uint = 18446744073709551615
                                                                            _ = v111
                                                                            _ = v113
                                                                        } else {
                                                                            v114 := GenS7{F0: 18446744073709551615, F1: 1, F2: fi95}
                                                                            v5 = (v3 / v5)
                                                                            switch (("" + "hello") + "\n\t") {
                                                                            case "a":
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = ""
                                                                                }
                                                                            case "foo bar":
                                                                                fmt.Println(fi97, v69, v69)
                                                                            case "":
                                                                                v98 -= v98
                                                                            }
                                                                            _ = len(v4)
                                                                            _ = v114
                                                                        }
                                                                    case -427:
                                                                        var v115 uint16 = 744
                                                                        _ = v115
                                                                    default:
                                                                        var v116 float32 = -2.25
                                                                        _ = v116
                                                                    }
                                                                    _ = v103
                                                                    _ = v104
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                                _ = len(v1)
                                                                v3 += 0.0
                                                                _ = fi97
                                                            }
                                                            v3 = (((float64(float32((v5 - (v3 / float64(1.5))))) / float64(float32(v3))) - 0.1) * 1.5)
                                                        }
                                                        _ = fi95
                                                    }
                                                    for fi117 := 0; fi117 < 1; fi117++ {
                                                        fmt.Println(v5, v5, v5)
                                                        v5 -= 1e10
                                                        _ = fi117
                                                    }
                                                }
                                            }
                                        } else {
                                            v118 := GenS5{F0: true}
                                            for fi119 := 0; fi119 < 7; fi119++ {
                                                for fi120 := 0; fi120 < 11; fi120++ {
                                                    if true {
                                                        var v121 int32 = 1
                                                        delete(v1, 2147483648)
                                                        if (!(!true)) {
                                                            if (!(32767 <= 1)) {
                                                                var v122 int8 = 42
                                                                switch (-(fi119)) {
                                                                case 9223372036854775807:
                                                                    fmt.Println(v69, fi119)
                                                                case 42:
                                                                    if true {
                                                                        if (644 < 18446744073709551615) {
                                                                            var v123 int = fi119
                                                                            _ = v123
                                                                        } else {
                                                                            v124 := v1[1]
                                                                            v3 = v5
                                                                            if v124 {
                                                                                var v125 int64 = 1
                                                                                _ = len(v4)
                                                                                var v126 string = "a"
                                                                                _ = v125
                                                                                _ = v126
                                                                            } else {
                                                                                v122 = v122
                                                                                delete(v6, 447)
                                                                                fmt.Println(v121, fi119)
                                                                            }
                                                                            _ = v124
                                                                        }
                                                                        if (((!false) && (false || false)) || (!false)) {
                                                                            var v127 uint8 = 67
                                                                            _ = v127
                                                                        }
                                                                        switch -817 {
                                                                        case -495:
                                                                            fmt.Println(fi120, v122)
                                                                        case 0:
                                                                            fmt.Println(v69)
                                                                        }
                                                                    } else {
                                                                        if true {
                                                                            var v128 int64 = 970
                                                                            fmt.Println(v122, v128)
                                                                            v129 := GenS1{F0: 128, F1: (-(fi120))}
                                                                            _ = v128
                                                                            _ = v129
                                                                        } else {
                                                                            var v130 int = 784
                                                                            v122 = v122
                                                                            _ = v130
                                                                        }
                                                                        if true {
                                                                            v131 := GenS19{F0: (1 < 148), F1: -9223372036854775808, F2: 0}
                                                                            v3 -= (-2.25 / ((v5 - 1.5) + 0.0))
                                                                            var v132 int32 = v121
                                                                            _ = v131
                                                                            _ = v132
                                                                        }
                                                                        var v133 int32 = (v121 + -2147483648)
                                                                        v3 += v5
                                                                        if false {
                                                                            fmt.Println(fi119, v5)
                                                                            if true {
                                                                                if (!true) {
                                                                                    var v134 uint = 2
                                                                                    if (((v69 + ((((v69 * float32(v3)) + v69) * 0.0) + 1e10)) == float32(3.14159)) || true) {
                                                                                        switch ("" + "hello") {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v122, v133, v5)
                                                                                        case "":
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                        if (fi119 >= (-(fi119))) {
                                                                                            var v135 uint64 = 578
                                                                                            _ = v135
                                                                                        } else {
                                                                                            switch "a" {
                                                                                            case "":
                                                                                                var v136 int16 = -448
                                                                                                _ = v136
                                                                                            case "hello":
                                                                                                switch fi120 {
                                                                                                case -680:
                                                                                                    switch int(v134) {
                                                                                                    case 32767:
                                                                                                        switch "" {
                                                                                                        case "a":
                                                                                                            v121 += v121
                                                                                                        case "hello":
                                                                                                            v5 += v5
                                                                                                        default:
                                                                                                            var v137 string = ""
                                                                                                            _ = v137
                                                                                                        }
                                                                                                    case -264:
                                                                                                        var v138 uint64 = 782
                                                                                                        _ = v138
                                                                                                    case 250:
                                                                                                        v122 -= 1
                                                                                                    default:
                                                                                                        var v139 uint = v134
                                                                                                        _ = v139
                                                                                                    }
                                                                                                case -272:
                                                                                                    if (!(0 >= 111)) {
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                v133 += v133
                                                                                                                if ("a" >= "hello") {
                                                                                                                    v1[2147483647] = false
                                                                                                                    delete(v6, 42)
                                                                                                                    _ = len(v4)
                                                                                                                    var v140 uint8 = 128
                                                                                                                    _ = v140
                                                                                                                } else {
                                                                                                                    var v141 int = fi120
                                                                                                                    _ = v141
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch "\n\t" {
                                                                                                            case "":
                                                                                                                v122 -= v122
                                                                                                            }
                                                                                                        }
                                                                                                        v142 := GenS20{F0: 1, F1: 255}
                                                                                                        _ = v142
                                                                                                    }
                                                                                                default:
                                                                                                    var v143 int64 = -718
                                                                                                    _ = v143
                                                                                                }
                                                                                            case "\n\t":
                                                                                                if (!(true && true)) {
                                                                                                    v2[int(v122)] = (fi120 >> ((v134) & 63))
                                                                                                }
                                                                                            default:
                                                                                                v144 := GenS15{F0: 42, F1: (-(v122))}
                                                                                                _ = v144
                                                                                            }
                                                                                            if (false && true) {
                                                                                                fmt.Println(v3)
                                                                                                fmt.Println(v122, v69, v3)
                                                                                            }
                                                                                        }
                                                                                        delete(v1, 0)
                                                                                        var v145 uint16 = 0
                                                                                        v146 := GenS19{F0: false, F1: 42, F2: (v122 & (v122 * genFn1()))}
                                                                                        _ = v145
                                                                                        _ = v146
                                                                                    }
                                                                                    _ = v134
                                                                                }
                                                                                v122 -= (v122 / ((1) | 1))
                                                                            }
                                                                            if len(v4) > 0 {
                                                                                v4[0] = "foo bar"
                                                                            }
                                                                        }
                                                                        _ = v133
                                                                    }
                                                                case -553:
                                                                    v147 := GenS21{F0: 42, F1: false, F2: 65535}
                                                                    _ = len(v4)
                                                                    switch ("foo bar" + "hello") {
                                                                    case "":
                                                                        var v148 uint8 = 42
                                                                        _ = v148
                                                                    }
                                                                    _ = v147
                                                                default:
                                                                    if true {
                                                                        delete(v6, 4294967295)
                                                                        if false {
                                                                            fmt.Println(v5)
                                                                            fmt.Println(v121)
                                                                        }
                                                                        var v149 int32 = v121
                                                                        _ = v149
                                                                    }
                                                                }
                                                                var v150 uint16 = 499
                                                                v151 := GenS13{F0: v121, F1: 42, F2: 563}
                                                                _ = v122
                                                                _ = v150
                                                                _ = v151
                                                            }
                                                        }
                                                        _ = v121
                                                    }
                                                    _ = fi120
                                                }
                                                v3 -= v3
                                                for fi152 := 0; fi152 < 5; fi152++ {
                                                    v153 := GenS22{F0: 42, F1: 32767}
                                                    if false {
                                                        v154 := GenS23{F0: 2, F1: 387, F2: "\n\t"}
                                                        if (474 != 255) {
                                                            var v155 uint = 18446744073709551615
                                                            v156 := GenS24{F0: 32768, F1: 1, F2: (v155 & v155)}
                                                            if (("hello" + "hello") > (("" + ("\n\t" + "")) + (("foo bar" + ("" + ("foo bar" + "\n\t"))) + "a"))) {
                                                                v157 := v6[1]
                                                                fmt.Println(v3, v155, fi119)
                                                                v158 := GenS25{F0: 677, F1: 1}
                                                                var v159 string = "hello"
                                                                _ = v157
                                                                _ = v158
                                                                _ = v159
                                                            }
                                                            _ = v155
                                                            _ = v156
                                                        }
                                                        _ = v154
                                                    }
                                                    _ = fi152
                                                    _ = v153
                                                }
                                                _ = fi119
                                            }
                                            _ = v118
                                        }
                                        v160 := GenS26{F0: -386, F1: 1, F2: (1 != 0)}
                                        _ = v160
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        if false {
                                            switch "hello" {
                                            case "foo bar":
                                                var v161 uint8 = 255
                                                _ = v161
                                            case "":
                                                v3 -= ((v3 + v5) - (v3 - 0.0))
                                            case "\n\t":
                                                var v162 string = "foo bar"
                                                _ = v162
                                            default:
                                                var v163 float32 = -2.25
                                                _ = v163
                                            }
                                        }
                                    case "":
                                        for fi164 := 0; fi164 < 10; fi164++ {
                                            fmt.Println(v5, v5, fi164)
                                            _ = fi164
                                        }
                                    default:
                                        _ = len(v4)
                                    }
                                    _ = v70
                                } else {
                                    v4 = append(v4, ("hello" + "a"))
                                }
                                _ = v69
                            } else {
                                var v165 uint = 879
                                if true {
                                    v5 += v3
                                    v5 += ((v5 / v5) / (v3 + 3.14159))
                                }
                                for fi166 := 0; fi166 < 5; fi166++ {
                                    var v167 string = "hello"
                                    var v168 uint32 = 123
                                    _ = fi166
                                    _ = v167
                                    _ = v168
                                }
                                _ = v165
                            }
                        } else {
                            if (false && ((v3 > float64((0.1 - 0.0))) || (!((42 > 120) && false)))) {
                                v169 := GenS19{F0: (1 < 128), F1: -9223372036854775808, F2: -128}
                                var v170 int32 = -2147483648
                                for fi171 := 0; fi171 < 6; fi171++ {
                                    v172 := GenS23{F0: 361, F1: -1, F2: ""}
                                    _ = fi171
                                    _ = v172
                                }
                                _ = v169
                                _ = v170
                            } else {
                                v5 = float64((1e10 + -2.25))
                                v2[2147483647] = 42
                                fmt.Println(v3)
                            }
                            var v173 uint = 4294967295
                            switch 32767 {
                            case 311:
                                for fi174 := 0; fi174 < 10; fi174++ {
                                    fmt.Println(v173, fi174, fi174)
                                    delete(v2, (fi174 ^ fi174))
                                    _ = fi174
                                }
                            case 222:
                                var v175 float64 = 0.1
                                _ = v175
                            default:
                                _ = len(v2)
                            }
                            _ = v173
                        }
                    default:
                        for fi176 := 0; fi176 < 6; fi176++ {
                            v4 = append(v4, "")
                            var v177 string = "\n\t"
                            _ = fi176
                            _ = v177
                        }
                    }
                } else {
                    var v178 float64 = 1.5
                    _ = v178
                }
            default:
                switch "" {
                case "foo bar":
                    v3 -= (0.1 * v5)
                case "hello":
                    if (genFn1() < 8) {
                        if false {
                            var v179 uint64 = 824
                            _ = v179
                        } else {
                            fmt.Println(v3, v3)
                        }
                        if (!((float64(-2.25) / (1e10 / (1.5 - 0.1))) > (((v3 - v5) * -2.25) - 1e10))) {
                            var v180 int = 148
                            v180 -= 4294967295
                            v5 -= v5
                            _ = v180
                        }
                    }
                default:
                    fmt.Println(v5)
                }
            }
        case "foo bar":
            var v181 string = "foo bar"
            _ = v181
        default:
            v182 := v1[42]
            _ = v182
        }
        switch 32768 {
        case 2147483647:
            fmt.Println(v3)
        case 4294967295:
            if (!(!false)) {
                v183 := v6[2147483647]
                _ = v183
            } else {
                var v184 float64 = 0.1
                _ = v184
            }
        default:
            switch 1 {
            case 32768:
                v3 = float64(1.5)
            case 1:
                for fi185 := 0; fi185 < 4; fi185++ {
                    var v186 int64 = -1
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    v4 = append(v4, "")
                    _ = fi185
                    _ = v186
                }
            default:
                v187 := GenS21{F0: 1, F1: true, F2: 1}
                _ = v187
            }
        }
        v5 = 0.0
    }
    var v188 float32 = 1e10
    fmt.Println(v188, v3, v3)
    var v189 uint = 2
    v188 -= (-2.25 + v188)
    fmt.Println(v3, v5)
    v5 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v188:", v188)
    fmt.Println("v189:", v189)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v188
    _ = v189
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
