package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int8; F1 int64; F2 uint }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint64; F1 int; F2 uint64 }
type GenS7 struct { F0 int; F1 uint64; F2 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 string; F1 uint64; F2 int }
type GenS12 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 int64) uint64 {
    if (1 <= -1000) {
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                p0 = p0
                var v8 int8 = -1
                var v9 uint8 = 9
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            var v10 int64 = (p0 >> ((2) & 63))
            v11 := []int{128, 9223372036854775807}
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: p0, F1: 0}
            var v14 float32 = -2.25
            p0 += (p0 & p0)
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    p0 += (p0 >> ((255) & 63))
    p0 -= (p0 - (p0 ^ p0))
    if (false && (false && false)) {
        var v15 float32 = 0.0
        _ = v15
    }
    v16 := GenS2{F0: 1, F1: 9223372036854775807}
    _ = p0
    _ = v16
    return 168
}
func genFn1(p0 string) uint32 {
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := []int{(fi17 ^ fi17), (fi17 << ((255) & 63)), (fi17 >> ((65535) & 63))}
        if (!false) {
            p0 = (p0 + p0)
            fmt.Println(p0, p0)
        } else {
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS3{F0: -106, F1: int64(fi19), F2: 255}
                var v21 int32 = 119
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        if ((!false) && (42 >= 1)) {
            fmt.Println(p0, p0)
        } else {
            fmt.Println(fi17)
            v22 := GenS1{F0: -423, F1: 255}
            for fi23 := 0; fi23 < 15; fi23++ {
                p0 = ""
                _ = fi23
            }
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    p0 = p0
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{}
    var v3 int8 = -128
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 65535
    var v6 uint64 = 1
    v3 += v3
    for fi24 := 0; fi24 < 8; fi24++ {
        v6 = v6
        fmt.Println(v3, v5)
        _ = fi24
    }
    fmt.Println(v5, v6)
    for fi25 := 0; fi25 < 10; fi25++ {
        if true {
            switch (fi25 % ((-368) | 1)) {
            case 164:
                var v26 uint8 = 230
                _ = v26
            default:
                var v27 uint16 = (v5 | 1)
                _ = v27
            }
        }
        _ = fi25
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 6; fi29++ {
            switch fi28 {
            case 1:
                v2 = append(v2, 128)
            case -554:
                v3 = (v3 - v3)
            case -309:
                var v30 uint64 = v6
                _ = v30
            default:
                fmt.Println(fi29)
            }
            delete(v4, 345)
            v6 += v6
            _ = fi29
        }
        switch "\n\t" {
        case "a":
            for fi31 := 0; fi31 < 10; fi31++ {
                var v32 string = (((("foo bar" + "") + "hello") + "") + ("foo bar" + "hello"))
                _ = fi31
                _ = v32
            }
        case "":
            v33 := v1[65535]
            _ = v33
        default:
            switch (("foo bar" + (("a" + "hello") + "\n\t")) + ("" + "\n\t")) {
            case "foo bar":
                if false {
                    switch 4294967295 {
                    case 32768:
                        v34 := GenS4{F0: ("a" + "hello"), F1: genFn0(-9223372036854775808)}
                        _ = v34
                    }
                    v35 := v4[4294967295]
                    if true {
                        _ = len(v1)
                        for fi36 := 0; fi36 < 19; fi36++ {
                            v35 = v35
                            var v37 int8 = (v3 << ((0) & 63))
                            _ = fi36
                            _ = v37
                        }
                    } else {
                        fmt.Println(v6, v35)
                        fmt.Println(v6)
                        for fi38 := 0; fi38 < 9; fi38++ {
                            var v39 uint = 2
                            _ = fi38
                            _ = v39
                        }
                    }
                    _ = v35
                }
            case "a":
                if (!((v3 - (v3 >> ((18446744073709551615) & 63))) == v3)) {
                    v40 := GenS1{F0: 9223372036854775807, F1: (v5 % ((234) | 1))}
                    fmt.Println(v3, v3)
                    _ = v40
                }
            case "hello":
                if true {
                    var v41 uint8 = 255
                    for fi42 := 0; fi42 < 17; fi42++ {
                        fmt.Println(v41, fi28, v41)
                        if (!(genFn1("") >= 65535)) {
                            v5 = (v5 | (^(v5)))
                            fmt.Println(v3, v5)
                        }
                        if (false || false) {
                            var v43 int64 = 348
                            var v44 uint64 = uint64(v41)
                            if (0.1 <= 1e10) {
                                v3 -= 14
                                v45 := GenS3{F0: int8(v43), F1: (v43 - int64(v44)), F2: 251}
                                switch fi42 {
                                case 9223372036854775807:
                                    v41 = (v41 & v41)
                                case 32767:
                                    fmt.Println(v44, v43, v41)
                                case -436:
                                    switch "foo bar" {
                                    case "":
                                        v2 = append(v2, -2147483648)
                                    case "hello":
                                        var v46 int8 = (v3 * 127)
                                        _ = v46
                                    default:
                                        if false {
                                            v47 := v1[genFn1("hello")]
                                            switch ("\n\t" + ("\n\t" + "a")) {
                                            case "\n\t":
                                                var v48 int32 = 2147483647
                                                _ = v48
                                            case "":
                                                if ((v47 || true) && ("\n\t" == ("hello" + "\n\t"))) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v3 = v3
                                                    }
                                                    switch (("\n\t" + "") + "") {
                                                    case "":
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi28 >> ((616) & 63))
                                                        }
                                                    }
                                                    v49 := GenS3{F0: 42, F1: 42, F2: 468}
                                                    _ = v49
                                                } else {
                                                    v50 := GenS5{F0: 4294967295, F1: fi28}
                                                    if ((v41 - 42) == (v41 & v41)) {
                                                        v51 := GenS5{F0: 663, F1: (fi42 ^ 128)}
                                                        _ = v51
                                                    }
                                                    delete(v4, 65535)
                                                    if false {
                                                        if (-554 == -1) {
                                                            v3 += v3
                                                            v44 = (v44 * 215)
                                                            v52 := GenS6{F0: (^(v6)), F1: fi42, F2: genFn0((v43 ^ (v43 >> ((228) & 63))))}
                                                            _ = v52
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                switch fi28 {
                                                                case -2147483648:
                                                                    if ((0.0 * 1e10) > float32(1.5)) {
                                                                        var v53 int = (^(fi28))
                                                                        v41 += (^(v41))
                                                                        _ = v53
                                                                    }
                                                                case 127:
                                                                    v1[genFn1("hello")] = v47
                                                                default:
                                                                    v43 = (v43 | v43)
                                                                }
                                                            case "a":
                                                                if true {
                                                                    fmt.Println(v6, v5, v6)
                                                                    switch -502 {
                                                                    case 32768:
                                                                        if true {
                                                                            fmt.Println(fi42, v43)
                                                                            var v54 float64 = -2.25
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                if false {
                                                                                    _ = len(v1)
                                                                                    var v55 uint64 = v44
                                                                                    if true {
                                                                                        var v56 uint32 = 2
                                                                                        _ = v56
                                                                                    }
                                                                                    var v57 int32 = 42
                                                                                    _ = v55
                                                                                    _ = v57
                                                                                }
                                                                            case "":
                                                                                v47 = true
                                                                            case "\n\t":
                                                                                v4[988] = 2147483647
                                                                            }
                                                                            _ = v54
                                                                        } else {
                                                                            var v58 int = fi28
                                                                            v47 = true
                                                                            v58 -= fi28
                                                                            _ = v58
                                                                        }
                                                                    default:
                                                                        switch 127 {
                                                                        case 329:
                                                                            var v59 int64 = v43
                                                                            _ = v59
                                                                        }
                                                                    }
                                                                    if v47 {
                                                                        switch (^(fi28)) {
                                                                        case 128:
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        case -2147483648:
                                                                            switch "" {
                                                                            case "hello":
                                                                                switch (fi42 ^ fi28) {
                                                                                case -9223372036854775808:
                                                                                    if true {
                                                                                        v5 = (v5 ^ 2)
                                                                                    }
                                                                                case -771:
                                                                                    fmt.Println(v5, v44, fi42)
                                                                                }
                                                                            case "foo bar":
                                                                                v44 = 18446744073709551615
                                                                            default:
                                                                                var v60 uint = 1
                                                                                _ = v60
                                                                            }
                                                                        default:
                                                                            v3 = (v3 / ((v3) | 1))
                                                                        }
                                                                    } else {
                                                                        var v61 int8 = 105
                                                                        var v62 uint8 = v41
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                }
                                                            }
                                                            var v63 int = int(v6)
                                                            v64 := v4[2147483647]
                                                            var v65 int8 = v3
                                                            var v66 float32 = 1.5
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        fmt.Println(v6)
                                                    }
                                                    fmt.Println(v44, v47, v41)
                                                    _ = v50
                                                }
                                            }
                                            _ = v47
                                        } else {
                                            fmt.Println(v43, fi28)
                                            var v67 int16 = int16(fi28)
                                            var v68 int16 = (v67 + 153)
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                }
                                _ = v45
                            }
                            _ = v43
                            _ = v44
                        } else {
                            v6 = (v6 / (((v6 << ((uint(v6)) & 63))) | 1))
                        }
                        _ = fi42
                    }
                    _ = v41
                }
            default:
                for fi69 := 0; fi69 < 13; fi69++ {
                    fmt.Println(v3, v3, fi28)
                    switch (("" + "\n\t") + "\n\t") {
                    case "\n\t":
                        var v70 uint32 = 1
                        _ = v70
                    case "hello":
                        var v71 float64 = -2.25
                        _ = v71
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v72 uint32 = 143
                        var v73 int = 9223372036854775807
                        fmt.Println(v3, fi69, v72)
                        v74 := GenS7{F0: v73, F1: (v6 + 2), F2: v6}
                        _ = v72
                        _ = v73
                        _ = v74
                    default:
                        var v75 float64 = 3.14159
                        _ = v75
                    }
                    _ = fi69
                }
            }
        }
        _ = fi28
    }
    v3 += (^(v3))
    var v76 int64 = 1
    var v77 bool = (!(42 >= v6))
    switch 16 {
    case 4294967295:
        delete(v4, 1)
    default:
        fmt.Println(v6, v5)
    }
    if false {
        var v78 uint32 = 245
        _ = v78
    } else {
        for fi79 := 0; fi79 < 9; fi79++ {
            if true {
                var v80 uint32 = 42
                _ = v80
            } else {
                _ = len(v2)
                for fi81 := 0; fi81 < 4; fi81++ {
                    v82 := GenS8{F0: (v3 | v3)}
                    var v83 bool = (!(false && (-9223372036854775808 < -9223372036854775808)))
                    if ((v76 * v76) >= -412) {
                        v84 := GenS4{F0: "a", F1: (^(v6))}
                        switch (fi81 - -9223372036854775808) {
                        case 128:
                            switch (("a" + "hello") + "foo bar") {
                            case "\n\t":
                                if false {
                                    switch ("foo bar" + "a") {
                                    case "\n\t":
                                        _ = len(v4)
                                    case "a":
                                        v6 -= 18446744073709551615
                                    default:
                                        v2 = append(v2, fi79)
                                    }
                                }
                            case "":
                                if v83 {
                                    fmt.Println(v77)
                                    switch ("hello" + "\n\t") {
                                    case "\n\t":
                                        if len(v2) > 0 {
                                            v2[0] = (fi81 - 42)
                                        }
                                    case "hello":
                                        v85 := GenS7{F0: (fi79 - fi79), F1: v6, F2: (v6 << ((42) & 63))}
                                        _ = v85
                                    case "a":
                                        v6 = v6
                                    default:
                                        var v86 string = "foo bar"
                                        _ = v86
                                    }
                                }
                            }
                        default:
                            if ((!true) && (959 < 249)) {
                                v3 -= (^(v3))
                            } else {
                                if ((v76 == (-(v76))) || v83) {
                                    if (65535 >= 255) {
                                        _ = len(v4)
                                        var v87 uint = 2
                                        _ = v87
                                    }
                                    var v88 string = "a"
                                    if true {
                                        var v89 int32 = 1
                                        _ = v89
                                    }
                                    v83 = false
                                    var v90 uint16 = v5
                                    switch fi79 {
                                    case 4294967295:
                                        v77 = ((((v83 || false) || v77) && true) || (v77 && (0 < -161)))
                                        var v91 int8 = v3
                                        v92 := GenS1{F0: (-(v76)), F1: (^(v5))}
                                        switch v88 {
                                        case "hello":
                                            v4[genFn1(("a" + v88))] = 2147483648
                                        default:
                                            v91 = 100
                                        }
                                        v2 = append(v2, -2147483648)
                                        _ = v91
                                        _ = v92
                                    case 42:
                                        var v93 uint16 = v90
                                        _ = v93
                                    case -730:
                                        var v94 uint32 = 692
                                        _ = v94
                                    default:
                                        fmt.Println(v83, v83, v90)
                                    }
                                    switch v88 {
                                    case "hello":
                                        v95 := v1[2147483647]
                                        _ = v95
                                    case "a":
                                        v96 := v4[2147483648]
                                        _ = v96
                                    default:
                                        var v97 uint8 = 1
                                        switch v88 {
                                        case "hello":
                                            v98 := GenS3{F0: v3, F1: v76, F2: 255}
                                            v99 := GenS1{F0: (v76 ^ (v76 ^ (^(v76)))), F1: v5}
                                            _ = v98
                                            _ = v99
                                        case "\n\t":
                                            if true {
                                                v3 = (v3 & 42)
                                                if true {
                                                    v97 = 208
                                                } else {
                                                    delete(v1, 2147483648)
                                                }
                                            } else {
                                                var v100 int8 = v3
                                                _ = v100
                                            }
                                        default:
                                            v101 := GenS8{F0: (^(v3))}
                                            _ = v101
                                        }
                                        if (255 != 65535) {
                                            switch (v88 + v88) {
                                            case "foo bar":
                                                var v102 string = v88
                                                _ = v102
                                            }
                                        } else {
                                            var v103 uint64 = v6
                                            var v104 uint16 = (v90 - (^(v5)))
                                            v4[4294967295] = 0
                                            v105 := GenS9{F0: v76}
                                            v106 := GenS8{F0: v3}
                                            _ = v103
                                            _ = v104
                                            _ = v105
                                            _ = v106
                                        }
                                        _ = v97
                                    }
                                    _ = v88
                                    _ = v90
                                }
                            }
                        }
                        var v107 uint8 = 127
                        _ = v107
                        _ = v84
                    } else {
                        if v83 {
                            v6 -= (v6 - (v6 / ((v6) | 1)))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v108 uint8 = 142
                            switch (fi81 | fi81) {
                            case 32767:
                                v83 = (18446744073709551615 > 0)
                            case 675:
                                var v109 uint16 = v5
                                _ = v109
                            default:
                                v110 := GenS10{F0: -32768, F1: 0}
                                _ = v110
                            }
                            _ = v108
                        } else {
                            v111 := GenS7{F0: (fi81 | fi79), F1: (v6 >> ((65535) & 63)), F2: v6}
                            switch 4294967295 {
                            case -1:
                                fmt.Println(fi81, v3, fi79)
                            case 1:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                if v83 {
                                    v76 += v76
                                    fmt.Println(v77, v5, fi79)
                                }
                            }
                            v83 = v77
                            _ = v111
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi81
                    _ = v82
                    _ = v83
                }
            }
            v112 := GenS11{F0: "", F1: (v6 ^ 42), F2: (^(fi79))}
            _ = fi79
            _ = v112
        }
    }
    fmt.Println(v6, v3)
    for fi113 := 0; fi113 < 1; fi113++ {
        fmt.Println(fi113, v5)
        if true {
            if len(v2) > 0 {
                v2[0] = fi113
            }
            v76 = v76
            for fi114 := 0; fi114 < 18; fi114++ {
                v3 += (v3 * -128)
                var v115 float64 = 3.14159
                v6 -= genFn0(v76)
                var v116 int8 = (v3 | v3)
                _ = fi114
                _ = v115
                _ = v116
            }
        } else {
            v5 = (v5 ^ (v5 & (v5 - (v5 + (v5 % (((^(v5))) | 1))))))
        }
        for fi117 := 0; fi117 < 5; fi117++ {
            v3 -= v3
            _ = fi117
        }
        _ = fi113
    }
    _ = len(v1)
    if true {
        for fi118 := 0; fi118 < 1; fi118++ {
            v5 -= v5
            switch "foo bar" {
            case "hello":
                v5 = 32767
            case "\n\t":
                var v119 int8 = v3
                _ = v119
            default:
                if v77 {
                    if (v77 || (true && (("\n\t" > "") && (!(-2147483648 >= 0))))) {
                        for fi120 := 0; fi120 < 17; fi120++ {
                            if v77 {
                                if v77 {
                                    _ = len(v2)
                                } else {
                                    v76 += v76
                                    switch "\n\t" {
                                    case "\n\t":
                                        if v77 {
                                            fmt.Println(v76)
                                            v77 = v77
                                            v1[2147483647] = true
                                        }
                                    case "":
                                        var v121 int16 = 1000
                                        _ = v121
                                    case "hello":
                                        v3 = int8(v76)
                                    default:
                                        var v122 uint64 = genFn0(-470)
                                        _ = v122
                                    }
                                    _ = len(v1)
                                }
                            }
                            v76 = v76
                            if v77 {
                                if (4294967295 == 18446744073709551615) {
                                    if len(v2) > 0 {
                                        v2[0] = -329
                                    }
                                } else {
                                    fmt.Println(v6)
                                    switch ("foo bar" + ("" + "")) {
                                    case "":
                                        var v123 bool = false
                                        _ = v123
                                    case "hello":
                                        v2 = append(v2, fi118)
                                    case "\n\t":
                                        v3 += (v3 & (v3 / ((v3) | 1)))
                                    default:
                                        if (!(genFn0((v76 * (v76 ^ v76))) < v6)) {
                                            fmt.Println(fi118, v3)
                                            switch ("" + "") {
                                            case "a":
                                                fmt.Println(v3)
                                            default:
                                                v4[uint32(v6)] = 42
                                            }
                                        }
                                        v76 -= 42
                                    }
                                }
                                v124 := GenS12{F0: 981, F1: v6, F2: 612}
                                _ = v124
                            }
                            _ = fi120
                        }
                        v5 -= v5
                    } else {
                        var v125 float32 = (3.14159 / ((1.5 * (-2.25 / 1e10)) - 0.1))
                        _ = v125
                    }
                    var v126 uint = 65535
                    _ = v126
                } else {
                    delete(v4, genFn1((("" + ("a" + "\n\t")) + "\n\t")))
                }
            }
            _ = fi118
        }
        delete(v4, genFn1(("hello" + "")))
    } else {
        v3 = (-(v3))
        if v77 {
            v127 := GenS2{F0: -42, F1: -1}
            _ = v127
        } else {
            v128 := GenS6{F0: (^(v6)), F1: 42, F2: genFn0(v76)}
            var v129 bool = false
            _ = v128
            _ = v129
        }
        switch "foo bar" {
        case "foo bar":
            if v77 {
                v130 := GenS8{F0: (v3 + -1)}
                _ = v130
            }
        case "a":
            v6 -= v6
        case "\n\t":
            v131 := GenS12{F0: 42, F1: v6, F2: 2147483647}
            _ = v131
        default:
            v5 += 95
        }
    }
    if v77 {
        if v77 {
            v3 = v3
        }
        var v132 string = "\n\t"
        _ = v132
    } else {
        var v133 uint = 2
        _ = v133
    }
    fmt.Println(v76, v5, v77)
    if false {
        var v134 float64 = (3.14159 + 3.14159)
        _ = v134
    } else {
        _ = len(v2)
        fmt.Println(v77, v3, v76)
    }
    v6 = (^(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
}
