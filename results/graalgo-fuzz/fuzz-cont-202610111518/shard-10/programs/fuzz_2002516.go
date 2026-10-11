package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint32; F2 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS11 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 4294967295
    var v4 float32 = float32(-2.25)
    var v5 uint32 = (v3 + 65535)
    var v6 int = 9223372036854775807
    if true {
        fmt.Println(v4)
        if false {
            var v7 uint8 = 2
            if (!(v7 == v7)) {
                if false {
                    v7 += 211
                    if false {
                        if true {
                            if true {
                                if ((v7 + uint8(v3)) != v7) {
                                    var v8 int32 = 1
                                    v9 := GenS1{F0: "hello"}
                                    _ = v8
                                    _ = v9
                                } else {
                                    switch -9223372036854775808 {
                                    case -703:
                                        v7 += v7
                                    case -207:
                                        fmt.Println(v6, v7, v5)
                                    }
                                    var v10 int64 = 0
                                    _ = v10
                                }
                                v11 := []int{v6, 1}
                                _ = v11
                            } else {
                                v12 := GenS1{F0: "a"}
                                v13 := GenS2{F0: 1}
                                _ = v12
                                _ = v13
                            }
                        }
                        v14 := GenS3{F0: 42, F1: -688}
                        var v15 bool = true
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int64 = 42
                        _ = v16
                    }
                    switch v6 {
                    case -194:
                        fmt.Println(v3, v6, v3)
                    case -1:
                        for fi17 := 0; fi17 < 15; fi17++ {
                            var v18 uint32 = (^(v5))
                            switch (fi17 * (fi17 >> ((1) & 63))) {
                            case 32768:
                                if false {
                                    if (("" + "a") == "foo bar") {
                                        v1[v3] = (-(fi17))
                                    } else {
                                        if false {
                                            var v19 bool = true
                                            v3 = (v18 / (((v5 % ((v5) | 1))) | 1))
                                            _ = v19
                                        }
                                        fmt.Println(v5, v6, v4)
                                    }
                                    var v20 float32 = v4
                                    var v21 uint64 = 18446744073709551615
                                    _ = v20
                                    _ = v21
                                } else {
                                    fmt.Println(fi17, v5)
                                    var v22 uint64 = 946
                                    switch -700 {
                                    case 4294967295:
                                        fmt.Println(v6, v4, v22)
                                    case 32767:
                                        fmt.Println(fi17)
                                    case -1:
                                        if (3.14159 != (((1e10 * 3.14159) + 0.1) / (-2.25 + (((3.14159 / -2.25) + (3.14159 - -2.25)) - (((1.5 / 3.14159) * 3.14159) / (float64((v4 + float32(-2.25))) + 0.0)))))) {
                                            if ((-(v6)) > (^(v6))) {
                                                v22 = 729
                                                if true {
                                                    var v23 int32 = 0
                                                    v23 += (v23 << ((1) & 63))
                                                    switch fi17 {
                                                    case 128:
                                                        var v24 int16 = 614
                                                        _ = v24
                                                    case 2147483647:
                                                        var v25 uint = 274
                                                        _ = v25
                                                    case -493:
                                                        v26 := GenS3{F0: (v22 + v22), F1: -356}
                                                        _ = v26
                                                    }
                                                    _ = v23
                                                }
                                            } else {
                                                if (-104 != 244) {
                                                    for fi27 := 0; fi27 < 8; fi27++ {
                                                        v28 := GenS4{F0: v6, F1: 42}
                                                        v5 += 4294967295
                                                        if (1e10 != float64(0.0)) {
                                                            if ((128 == (v7 >> ((431) & 63))) || true) {
                                                                v6 = (^(fi17))
                                                            } else {
                                                                var v29 int16 = 32767
                                                                _ = v29
                                                            }
                                                            v4 += v4
                                                        }
                                                        if true {
                                                            if ((!(uint(v7) > 1)) || true) {
                                                                var v30 int64 = -1
                                                                v4 = (((1e10 / ((0.1 + (v4 + v4)) * v4)) - v4) - ((((v4 / ((float32(3.14159) - 3.14159) * (v4 + 3.14159))) - (-2.25 - v4)) * v4) * v4))
                                                                _ = v30
                                                            } else {
                                                                var v31 uint8 = (v7 << ((2) & 63))
                                                                _ = v31
                                                            }
                                                            var v32 float32 = v4
                                                            _ = v32
                                                        }
                                                        _ = fi27
                                                        _ = v28
                                                    }
                                                    for fi33 := 0; fi33 < 14; fi33++ {
                                                        var v34 int = fi33
                                                        v5 = 4294967295
                                                        _ = fi33
                                                        _ = v34
                                                    }
                                                } else {
                                                    if (v22 <= v22) {
                                                        v35 := GenS4{F0: v6, F1: 481}
                                                        if false {
                                                            for fi36 := 0; fi36 < 4; fi36++ {
                                                                var v37 int32 = 2147483647
                                                                v38 := GenS3{F0: (v22 << ((842) & 63)), F1: -9223372036854775808}
                                                                var v39 uint32 = v3
                                                                _ = fi36
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            fmt.Println(v5)
                                                        } else {
                                                            switch (fi17 ^ -9223372036854775808) {
                                                            case 42:
                                                                fmt.Println(v7)
                                                            case -349:
                                                                for fi40 := 0; fi40 < 4; fi40++ {
                                                                    if false {
                                                                        v41 := []int{928, -689, fi40}
                                                                        if (!(false && ((-843 <= 9223372036854775807) || false))) {
                                                                            fmt.Println(v6)
                                                                            fmt.Println(v3, v5)
                                                                        }
                                                                        _ = v41
                                                                    }
                                                                    var v42 uint32 = 0
                                                                    var v43 bool = false
                                                                    _ = fi40
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                            case 699:
                                                                var v44 bool = true
                                                                _ = v44
                                                            }
                                                        }
                                                        _ = v35
                                                    } else {
                                                        var v45 uint32 = (v3 << ((495) & 63))
                                                        _ = v45
                                                    }
                                                    if false {
                                                        if (0 >= 0) {
                                                            v46 := []int{}
                                                            var v47 uint32 = 1
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            switch "a" {
                                                            case "hello":
                                                                var v48 int64 = -456
                                                                _ = v48
                                                            case "":
                                                                if (true && true) {
                                                                    v4 -= float32(0.1)
                                                                    v49 := []int{(fi17 >> ((2) & 63))}
                                                                    for fi50 := 0; fi50 < 14; fi50++ {
                                                                        v51 := GenS1{F0: "a"}
                                                                        v5 = (v3 ^ uint32(v22))
                                                                        _ = fi50
                                                                        _ = v51
                                                                    }
                                                                    _ = v49
                                                                } else {
                                                                    fmt.Println(v5)
                                                                }
                                                            }
                                                            fmt.Println(v4, v6, v5)
                                                        }
                                                        if true {
                                                            v52 := []int{v6, v6}
                                                            v53 := GenS5{F0: (v7 ^ 80)}
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        v54 := v1[uint32(v7)]
                                                        _ = v54
                                                    }
                                                    fmt.Println(v22, v3, v6)
                                                }
                                            }
                                        }
                                    }
                                    _ = v22
                                }
                            }
                            _ = fi17
                            _ = v18
                        }
                    case 42:
                        switch ((("\n\t" + "foo bar") + "") + ((((("foo bar" + ("\n\t" + ("a" + "\n\t"))) + "foo bar") + "a") + "a") + "")) {
                        case "":
                            v55 := v2[78]
                            _ = v55
                        case "foo bar":
                            var v56 int32 = 0
                            _ = v56
                        default:
                            v4 = 1e10
                        }
                    default:
                        switch ("" + (("a" + "hello") + "foo bar")) {
                        case "":
                            v57 := []int{(v6 * v6)}
                            _ = v57
                        case "foo bar":
                            if (false || false) {
                                var v58 int = (-(v6))
                                _ = v58
                            } else {
                                var v59 float32 = float32(3.14159)
                                _ = v59
                            }
                        }
                    }
                } else {
                    for fi60 := 0; fi60 < 20; fi60++ {
                        for fi61 := 0; fi61 < 2; fi61++ {
                            v4 = 1.5
                            v62 := GenS6{F0: 8, F1: (v3 << ((316) & 63)), F2: false}
                            var v63 uint16 = 32767
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                        fmt.Println(v6, v7, v3)
                        _ = fi60
                    }
                }
            }
            _ = v7
        }
    } else {
        v64 := GenS5{F0: 42}
        var v65 uint32 = v3
        _ = v64
        _ = v65
    }
    fmt.Println(v5, v4)
    v66 := GenS6{F0: -128, F1: v5, F2: false}
    switch -2147483648 {
    case 802:
        v6 += 618
    case 2147483647:
        for fi67 := 0; fi67 < 20; fi67++ {
            v68 := GenS2{F0: 1}
            var v69 int32 = 738
            _ = fi67
            _ = v68
            _ = v69
        }
    case -426:
        switch v6 {
        case -1:
            fmt.Println(v6, v5, v3)
        case 1:
            for fi70 := 0; fi70 < 11; fi70++ {
                var v71 bool = false
                _ = fi70
                _ = v71
            }
        default:
            switch "\n\t" {
            case "foo bar":
                for fi72 := 0; fi72 < 3; fi72++ {
                    var v73 int32 = -746
                    v6 = (v6 % (((fi72 >> ((1) & 63))) | 1))
                    switch v6 {
                    case 2147483647:
                        switch 9223372036854775807 {
                        case -2147483648:
                            v74 := []int{v6, v6}
                            _ = v74
                        case 4294967295:
                            v75 := []int{(v6 | (v6 & 0)), 9223372036854775807}
                            _ = v75
                        case 128:
                            var v76 bool = true
                            _ = v76
                        default:
                            fmt.Println(v3, v5)
                        }
                    default:
                        if false {
                            v6 += (fi72 + fi72)
                        }
                    }
                    _ = fi72
                    _ = v73
                }
            default:
                switch (v6 & 0) {
                case 2147483647:
                    if false {
                        var v77 uint = 1
                        for fi78 := 0; fi78 < 14; fi78++ {
                            v79 := []int{fi78}
                            fmt.Println(v6, fi78, fi78)
                            fmt.Println(v77, v5, v6)
                            _ = fi78
                            _ = v79
                        }
                        _ = v77
                    }
                default:
                    for fi80 := 0; fi80 < 19; fi80++ {
                        v81 := GenS7{F0: (657 > 0)}
                        _ = fi80
                        _ = v81
                    }
                }
            }
        }
    default:
        switch "foo bar" {
        case "hello":
            if false {
                if (true || ((!true) || false)) {
                    v82 := v2[v5]
                    _ = v82
                } else {
                    v4 -= float32(-2.25)
                    v83 := []int{v6}
                    _ = v83
                }
            }
        case "":
            fmt.Println(v6, v4)
        case "foo bar":
            switch v6 {
            case 422:
                if ((98 == 127) && false) {
                    switch v6 {
                    case 1:
                        for fi84 := 0; fi84 < 17; fi84++ {
                            _ = len(v2)
                            v1[(v3 >> ((377) & 63))] = v6
                            _ = fi84
                        }
                    case -326:
                        v3 += 42
                    case -2147483648:
                        var v85 uint = uint(v3)
                        _ = v85
                    default:
                        switch ("a" + "\n\t") {
                        case "a":
                            var v86 string = ("foo bar" + ("a" + ("a" + "hello")))
                            _ = v86
                        case "":
                            v6 = (-(v6))
                        }
                    }
                    v6 = (v6 ^ v6)
                    if (false && (42 <= 65535)) {
                        if (42 != 100) {
                            v4 = v4
                        } else {
                            if (v4 >= float32((1e10 + (-2.25 * 1e10)))) {
                                switch (("" + "\n\t") + "a") {
                                case "foo bar":
                                    _ = len(v2)
                                default:
                                    v87 := GenS8{F0: (true || (509 < 255)), F1: 0}
                                    _ = v87
                                }
                                var v88 bool = false
                                var v89 bool = v88
                                _ = v88
                                _ = v89
                            }
                        }
                        var v90 int64 = 1
                        _ = v90
                    }
                }
            }
        }
    }
    if (v5 <= (v3 - 2)) {
        v91 := []int{(v6 ^ v6), v6, v6}
        if true {
            var v92 int = (v6 << ((18446744073709551615) & 63))
            v5 = 42
            _ = v92
        }
        v5 -= 65535
        _ = v91
    } else {
        for fi93 := 0; fi93 < 2; fi93++ {
            var v94 uint = 255
            _ = fi93
            _ = v94
        }
        v95 := []int{v6}
        _ = v95
    }
    v96 := GenS9{F0: 65535}
    if false {
        switch ("\n\t" + ("a" + "foo bar")) {
        case "a":
            v6 = (^(v6))
        case "hello":
            switch (("hello" + (("a" + ("a" + "")) + ((("foo bar" + "") + ("\n\t" + "")) + "a"))) + ("\n\t" + ("" + "a"))) {
            case "\n\t":
                for fi97 := 0; fi97 < 15; fi97++ {
                    switch 4294967295 {
                    case -185:
                        v98 := []int{fi97, fi97, fi97, v6}
                        _ = v98
                    default:
                        var v99 uint8 = 171
                        _ = v99
                    }
                    if false {
                        _ = len(v1)
                        if false {
                            v5 += (v3 + (v3 / ((v5) | 1)))
                            v4 = (0.0 / (1.5 * ((float32(-2.25) * 1.5) + 3.14159)))
                            if false {
                                if true {
                                    var v100 uint = 255
                                    _ = v100
                                } else {
                                    v6 = v6
                                    v4 = v4
                                    if true {
                                        if true {
                                            var v101 bool = true
                                            fmt.Println(v101, v6)
                                            _ = v101
                                        }
                                        var v102 int = v6
                                        var v103 int8 = 100
                                        _ = v102
                                        _ = v103
                                    }
                                }
                                v6 += v6
                            }
                        }
                        if true {
                            v104 := []int{9223372036854775807, v6, v6, (v6 % (((fi97 * fi97)) | 1))}
                            _ = v104
                        } else {
                            switch fi97 {
                            case -440:
                                v105 := []int{(fi97 & fi97), (v6 >> ((18446744073709551615) & 63))}
                                _ = v105
                            case -151:
                                v3 -= 42
                            }
                            fmt.Println(fi97, v5, v4)
                            if (!(42 > 128)) {
                                fmt.Println(v5, v3)
                                for fi106 := 0; fi106 < 7; fi106++ {
                                    var v107 uint8 = 0
                                    v5 -= (^(v5))
                                    v6 += fi97
                                    switch 2147483647 {
                                    case 1:
                                        var v108 float32 = (float32(1.5) - 1.5)
                                        _ = v108
                                    case 0:
                                        fmt.Println(v5, v5)
                                    case 32768:
                                        var v109 int8 = 0
                                        _ = v109
                                    }
                                    v3 = v5
                                    _ = fi106
                                    _ = v107
                                }
                                var v110 int8 = -1
                                _ = v110
                            } else {
                                var v111 int = (-(v6))
                                if (((((0 == 0) && (!true)) || ((1.5 / 1.0) <= 1e10)) && (true || (!(false || false)))) || true) {
                                    var v112 int64 = 455
                                    _ = v112
                                }
                                _ = v111
                            }
                        }
                    } else {
                        if true {
                            var v113 uint8 = 0
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                fmt.Println(v6, v5)
                            }
                            _ = v113
                        } else {
                            if (!true) {
                                fmt.Println(fi97, v6)
                                v5 = (v3 & (v5 ^ (v5 & (v3 & (v3 >> ((115) & 63))))))
                            } else {
                                v3 = (^(v5))
                                if true {
                                    fmt.Println(v6, v5)
                                    v114 := GenS10{F0: 1, F1: 127, F2: 1}
                                    _ = v114
                                } else {
                                    var v115 float64 = float64((v4 * 3.14159))
                                    var v116 int32 = 233
                                    _ = v115
                                    _ = v116
                                }
                                if (!(-2.25 == -2.25)) {
                                    v117 := GenS5{F0: 42}
                                    v4 -= ((float32((1.5 - 0.1)) * v4) - 0.0)
                                    var v118 int64 = 0
                                    _ = v117
                                    _ = v118
                                } else {
                                    var v119 uint = 947
                                    v120 := GenS2{F0: 2}
                                    if true {
                                        fmt.Println(v4, v119, fi97)
                                    }
                                    _ = v119
                                    _ = v120
                                }
                            }
                        }
                        v121 := v2[2147483647]
                        _ = v121
                    }
                    _ = fi97
                }
            }
        case "\n\t":
            var v122 bool = true
            _ = v122
        default:
            v4 += 1.5
        }
        switch v6 {
        case 361:
            var v123 int32 = 2147483647
            _ = v123
        case -9223372036854775808:
            fmt.Println(v4)
        case 585:
            if ("foo bar" == ((("foo bar" + "a") + ("hello" + (("" + (("\n\t" + "hello") + "a")) + "hello"))) + (("" + "hello") + "\n\t"))) {
                v124 := []int{0, v6, (v6 - v6)}
                switch 1 {
                case -9223372036854775808:
                    v3 = 2147483647
                }
                _ = v124
            } else {
                fmt.Println(v4)
                for fi125 := 0; fi125 < 15; fi125++ {
                    v126 := GenS11{F0: 0, F1: 1}
                    _ = fi125
                    _ = v126
                }
            }
        }
        for fi127 := 0; fi127 < 18; fi127++ {
            v3 = (v5 ^ (v3 - v3))
            _ = fi127
        }
    } else {
        for fi128 := 0; fi128 < 13; fi128++ {
            if false {
                if true {
                    v6 = v6
                    var v129 uint16 = 0
                    _ = v129
                }
            } else {
                var v130 uint = 523
                switch ("foo bar" + ("" + ((("\n\t" + ("\n\t" + "foo bar")) + "") + "a"))) {
                case "\n\t":
                    var v131 int8 = 1
                    _ = v131
                case "a":
                    if true {
                        fmt.Println(v5)
                        if (!false) {
                            if false {
                                fmt.Println(fi128, v6)
                                var v132 float32 = 0.0
                                v133 := GenS8{F0: (true || false), F1: 42}
                                _ = v132
                                _ = v133
                            } else {
                                v5 += 640
                            }
                            _ = len(v2)
                        }
                        fmt.Println(v3)
                    }
                }
                v2[v5] = v3
                _ = v130
            }
            _ = fi128
        }
        switch v6 {
        case 127:
            v3 = 2147483648
        case 85:
            switch "foo bar" {
            case "foo bar":
                v134 := GenS4{F0: (v6 ^ (v6 >> ((1) & 63))), F1: 1}
                _ = v134
            case "\n\t":
                for fi135 := 0; fi135 < 10; fi135++ {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi135)
                    case "hello":
                        var v136 uint64 = 4294967295
                        _ = v136
                    default:
                        var v137 uint = 4294967295
                        _ = v137
                    }
                    var v138 float64 = (-2.25 / ((0.1 / 1.0) * 1e10))
                    _ = fi135
                    _ = v138
                }
            case "":
                var v139 int8 = 1
                _ = v139
            }
        }
    }
    var v140 int16 = 1000
    var v141 bool = (0.1 <= float32(0.0))
    fmt.Println(v140)
    var v142 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v96
}
