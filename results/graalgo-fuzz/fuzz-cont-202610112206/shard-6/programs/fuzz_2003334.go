package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint; F1 int16; F2 int32 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int; F1 string; F2 uint }
type GenS10 struct { F0 int16; F1 bool }
type GenS11 struct { F0 uint; F1 int32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint; F1 int; F2 bool }
func genFn0(p0 uint, p1 int) uint64 {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: 255}
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int64 = -9223372036854775808
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    switch "\n\t" {
    case "hello":
        var v10 int = (-(p1))
        _ = v10
    case "foo bar":
        if (true && ("" < "hello")) {
            if (!((42 <= 843) && (4294967295 < 256))) {
                v11 := []int{(p1 & p1), -234, p1, -2147483648}
                p1 -= -417
                if (!(false && (!true))) {
                    p0 = (p0 / ((p0) | 1))
                }
                _ = v11
            }
            v12 := []int{}
            p0 = 1
            _ = v12
        }
    }
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            switch "\n\t" {
            case "foo bar":
                switch "" {
                case "\n\t":
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        v14 := []int{}
                        _ = v14
                    default:
                        if true {
                            var v15 uint64 = 4294967295
                            fmt.Println(p1)
                            p1 = (^(p1))
                            _ = v15
                        }
                    }
                case "hello":
                    fmt.Println(p1, p0)
                }
            case "a":
                if false {
                    if true {
                        for fi16 := 0; fi16 < 13; fi16++ {
                            fmt.Println(fi16, fi16)
                            fmt.Println(fi16, p1)
                            p1 += (fi13 | p1)
                            var v17 bool = true
                            _ = fi16
                            _ = v17
                        }
                        if (int(p0) > p1) {
                            switch ("a" + "\n\t") {
                            case "a":
                                v18 := GenS2{F0: 32767, F1: 4294967295, F2: 113}
                                _ = v18
                            case "\n\t":
                                switch -2147483648 {
                                case 4294967295:
                                    fmt.Println(p0)
                                case -2147483648:
                                    var v19 int = fi13
                                    _ = v19
                                case 397:
                                    if (false || true) {
                                        var v20 float64 = 1e10
                                        v21 := []int{(fi13 / ((p1) | 1)), p1, -950, p1}
                                        p0 += (p0 / (((p0 - p0)) | 1))
                                        _ = v20
                                        _ = v21
                                    } else {
                                        fmt.Println(p0)
                                        for fi22 := 0; fi22 < 10; fi22++ {
                                            v23 := GenS2{F0: 1000, F1: 736, F2: 194}
                                            var v24 int = p1
                                            v25 := GenS2{F0: 0, F1: uint64(p0), F2: 128}
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        fmt.Println(p0, fi13)
                                    }
                                default:
                                    if (-79 < 100) {
                                        for fi26 := 0; fi26 < 15; fi26++ {
                                            if false {
                                                switch (fi13 | fi13) {
                                                case 1:
                                                    v27 := GenS3{F0: 0, F1: 42}
                                                    v28 := GenS3{F0: -1, F1: 2}
                                                    _ = v27
                                                    _ = v28
                                                case -2147483648:
                                                    p0 += (p0 << ((p0) & 63))
                                                default:
                                                    var v29 int = (-(fi13))
                                                    _ = v29
                                                }
                                                v30 := []int{fi26, fi13, fi13, 680}
                                                _ = v30
                                            }
                                            v31 := []int{fi13, (p1 << ((p0) & 63)), (-(p1))}
                                            v32 := GenS3{F0: -485, F1: 2}
                                            p1 += p1
                                            var v33 bool = true
                                            p1 -= fi26
                                            switch "foo bar" {
                                            case "hello":
                                                if ((p0 / ((p0) | 1)) > p0) {
                                                    switch "a" {
                                                    case "":
                                                        switch "hello" {
                                                        case "hello":
                                                            if len(v31) >= 2 {
                                                                v31 = v31[1:len(v31)]
                                                            }
                                                        default:
                                                            if len(v31) > 0 {
                                                                v31[0] = fi13
                                                            }
                                                        }
                                                    case "foo bar":
                                                        if len(v31) > 0 {
                                                            v31[0] = 4294967295
                                                        }
                                                    case "hello":
                                                        var v34 int16 = -1
                                                        _ = v34
                                                    default:
                                                        p0 += (p0 >> ((343) & 63))
                                                    }
                                                    fmt.Println(fi26, p1, p1)
                                                } else {
                                                    fmt.Println(fi13, fi26, p0)
                                                    switch "\n\t" {
                                                    case "hello":
                                                        p1 -= 983
                                                    case "foo bar":
                                                        v35 := GenS4{F0: 128, F1: -797, F2: 0}
                                                        _ = v35
                                                    case "\n\t":
                                                        p1 -= 4294967295
                                                    }
                                                    fmt.Println(p0, fi13, v33)
                                                }
                                            case "":
                                                if len(v31) > 0 {
                                                    v31[0] = 32768
                                                }
                                            default:
                                                p1 += p1
                                            }
                                            var v36 float64 = 3.14159
                                            _ = fi26
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                            _ = v36
                                        }
                                        var v37 int64 = -9223372036854775808
                                        var v38 int64 = v37
                                        _ = v37
                                        _ = v38
                                    } else {
                                        p0 -= 18446744073709551615
                                    }
                                }
                            case "hello":
                                if false {
                                    switch "foo bar" {
                                    case "a":
                                        p0 += p0
                                    case "foo bar":
                                        v39 := GenS5{F0: 2, F1: 0, F2: 1000}
                                        _ = v39
                                    }
                                    var v40 bool = false
                                    p1 -= 32768
                                    _ = v40
                                }
                            }
                            for fi41 := 0; fi41 < 10; fi41++ {
                                var v42 uint64 = 2
                                _ = fi41
                                _ = v42
                            }
                            fmt.Println(p0)
                        }
                        var v43 float64 = (1.5 - float64(0.1))
                        _ = v43
                    }
                    var v44 uint = (p0 - p0)
                    fmt.Println(fi13, fi13, v44)
                    _ = v44
                } else {
                    p1 = p1
                }
            }
            _ = fi13
        }
        var v45 bool = (((((((0.0 + (0.1 - 0.1)) / (3.14159 + 1.5)) + (1e10 / 3.14159)) + 0.0) + (3.14159 * -2.25)) < -2.25) || (-2147483648 >= 0))
        _ = v45
    }
    switch (^(p1)) {
    case -572:
        var v46 uint64 = 1
        _ = v46
    case 128:
        if false {
            switch "" {
            case "\n\t":
                switch ("" + "a") {
                case "a":
                    switch 4294967295 {
                    case 1:
                        switch 127 {
                        case 369:
                            v47 := []int{(p1 * p1), -2147483648}
                            _ = v47
                        case -401:
                            v48 := []int{}
                            _ = v48
                        default:
                            var v49 int64 = 9223372036854775807
                            _ = v49
                        }
                    case 128:
                        for fi50 := 0; fi50 < 1; fi50++ {
                            p1 -= p1
                            _ = fi50
                        }
                    default:
                        v51 := GenS6{F0: p0, F1: 42, F2: 42}
                        _ = v51
                    }
                case "\n\t":
                    switch (("hello" + "") + ((("hello" + "foo bar") + "\n\t") + "a")) {
                    case "a":
                        p1 += -415
                    }
                case "hello":
                    fmt.Println(p0)
                }
            case "foo bar":
                p1 = (^(p1))
            default:
                var v52 uint64 = 42
                _ = v52
            }
            for fi53 := 0; fi53 < 7; fi53++ {
                for fi54 := 0; fi54 < 19; fi54++ {
                    if false {
                        switch (^(fi53)) {
                        case 1:
                            v55 := GenS4{F0: 37, F1: 435, F2: 314}
                            _ = v55
                        case -395:
                            fmt.Println(p0)
                        }
                        switch fi54 {
                        case 32767:
                            fmt.Println(fi53, fi54, fi54)
                        case 42:
                            v56 := []int{}
                            _ = v56
                        }
                        p0 = p0
                    } else {
                        if (!true) {
                            var v57 int32 = -2147483648
                            _ = v57
                        } else {
                            var v58 uint32 = 2
                            _ = v58
                        }
                        switch "\n\t" {
                        case "foo bar":
                            p0 -= (p0 * (^(p0)))
                        default:
                            v59 := []int{0, int(p0), -9223372036854775808, p1}
                            _ = v59
                        }
                        switch "\n\t" {
                        case "":
                            p0 = p0
                        case "a":
                            v60 := GenS2{F0: 1, F1: 42, F2: uint8(p0)}
                            _ = v60
                        }
                    }
                    var v61 int32 = 333
                    p0 -= 42
                    _ = fi54
                    _ = v61
                }
                _ = fi53
            }
            if (true && (164 < 178)) {
                var v62 string = "foo bar"
                v63 := GenS1{F0: 0, F1: 224}
                _ = v62
                _ = v63
            }
        }
    case 9223372036854775807:
        v64 := GenS6{F0: (p0 / (((p0 << ((4294967295) & 63))) | 1)), F1: -942, F2: 251}
        _ = v64
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint16, p1 uint16) string {
    for fi65 := 0; fi65 < 7; fi65++ {
        p0 = (^(p1))
        p1 = (p1 & p1)
        _ = fi65
    }
    if false {
        var v66 float32 = 0.1
        v67 := GenS4{F0: 128, F1: -32768, F2: 1}
        fmt.Println(v66, p0)
        _ = v66
        _ = v67
    } else {
        if (0.1 <= float64(3.14159)) {
            var v68 string = ("a" + ("\n\t" + ("a" + "a")))
            for fi69 := 0; fi69 < 11; fi69++ {
                var v70 string = (v68 + "hello")
                p1 -= 1
                _ = fi69
                _ = v70
            }
            if false {
                fmt.Println(p0, p1)
                v71 := GenS7{F0: genFn0(975, 32768), F1: 0, F2: 4294967295}
                fmt.Println(p0, v68, v68)
                _ = v71
            } else {
                for fi72 := 0; fi72 < 20; fi72++ {
                    p1 = (^(p1))
                    v73 := GenS1{F0: genFn0(760, (-(fi72))), F1: 2}
                    if false {
                        for fi74 := 0; fi74 < 2; fi74++ {
                            var v75 uint8 = 0
                            switch (fi72 / ((-2147483648) | 1)) {
                            case -426:
                                var v76 uint16 = 42
                                _ = v76
                            case 42:
                                fmt.Println(v68)
                            }
                            _ = fi74
                            _ = v75
                        }
                        for fi77 := 0; fi77 < 6; fi77++ {
                            v68 = "\n\t"
                            fmt.Println(p1, v68)
                            _ = fi77
                        }
                        fmt.Println(p1, v68)
                    } else {
                        for fi78 := 0; fi78 < 8; fi78++ {
                            var v79 int8 = -127
                            v80 := []int{fi72, fi72, (fi72 / ((-9223372036854775808) | 1)), fi72}
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                    }
                    _ = fi72
                    _ = v73
                }
            }
            _ = v68
        }
        fmt.Println(p1)
        var v81 string = ("\n\t" + "a")
        _ = v81
    }
    var v82 int = 407
    switch v82 {
    case 0:
        switch (v82 / (((v82 % (((v82 / (((v82 << ((65535) & 63))) | 1))) | 1))) | 1)) {
        case -9223372036854775808:
            if (0 <= -128) {
                var v83 uint64 = 344
                switch v82 {
                case 1:
                    var v84 uint = 712
                    _ = v84
                }
                _ = v83
            } else {
                fmt.Println(v82)
                v85 := []int{}
                _ = v85
            }
        }
    }
    _ = p0
    _ = p1
    _ = v82
    return "hello"
}

func main() {
    v1 := []string{("" + ""), "hello"}
    var v2 string = ("hello" + "a")
    var v3 uint8 = 1
    v4 := []uint32{2147483647}
    var v5 uint64 = 2
    if len(v1) > 0 {
        v1[0] = v2
    }
    var v86 bool = ((-244 > 42) || (!((!true) || (true && true))))
    var v87 uint16 = 32767
    if (!v86) {
        switch -2147483648 {
        case 127:
            v88 := GenS8{F0: 113, F1: 32767}
            _ = v88
        case -441:
            fmt.Println(v5, v2)
        }
    }
    v5 -= 4294967295
    if true {
        if (1 > 100) {
            fmt.Println(v87, v3)
            var v89 float64 = (0.1 * 0.1)
            _ = v89
        }
        v90 := GenS3{F0: 2147483647, F1: (v3 - v3)}
        fmt.Println(v86)
        _ = v90
    }
    switch -839 {
    case 686:
        for fi91 := 0; fi91 < 2; fi91++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v92 int16 = 1
            _ = fi91
            _ = v92
        }
    case 127:
        for fi93 := 0; fi93 < 5; fi93++ {
            if (!((fi93 & 553) >= fi93)) {
                v3 = (v3 << ((255) & 63))
            }
            switch v2 {
            case "":
                var v94 int16 = 0
                _ = v94
            case "hello":
                switch ((v2 + ((v2 + ("a" + v2)) + (("a" + v2) + (v2 + (v2 + (v2 + v2)))))) + v2) {
                case "hello":
                    switch (fi93 + fi93) {
                    case 32768:
                        var v95 int16 = -596
                        _ = v95
                    case -2147483648:
                        if true {
                            var v96 uint = uint(v5)
                            _ = v96
                        } else {
                            v2 = v2
                            _ = len(v1)
                            v97 := GenS6{F0: 65535, F1: -218, F2: -486}
                            _ = v97
                        }
                    default:
                        v86 = ((0.1 * (((1e10 - 0.0) * (1e10 - float32((1e10 / (0.0 - -2.25))))) / float32(0.1))) > 1e10)
                    }
                case "a":
                    if len(v4) > 0 {
                        v4[0] = 606
                    }
                }
            }
            var v98 uint16 = (v87 & (^(v87)))
            _ = fi93
            _ = v98
        }
    default:
        v99 := GenS2{F0: -32768, F1: (^(v5)), F2: (v3 + v3)}
        _ = v99
    }
    if (2 >= 2147483647) {
        var v100 bool = true
        for fi101 := 0; fi101 < 6; fi101++ {
            var v102 int8 = 127
            _ = fi101
            _ = v102
        }
        if v86 {
            var v103 bool = false
            _ = v103
        } else {
            for fi104 := 0; fi104 < 19; fi104++ {
                switch v2 {
                case "hello":
                    switch 742 {
                    case 9223372036854775807:
                        v105 := GenS9{F0: fi104, F1: v2, F2: 2}
                        _ = v105
                    case 4294967295:
                        v86 = (-1 < -693)
                    case 2147483647:
                        if v86 {
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                            fmt.Println(v86, v2)
                            switch (fi104 | (fi104 / ((fi104) | 1))) {
                            case 2147483647:
                                if ((false || true) || v86) {
                                    switch fi104 {
                                    case 32768:
                                        for fi106 := 0; fi106 < 20; fi106++ {
                                            if v86 {
                                                if v86 {
                                                    fmt.Println(v3, fi106)
                                                    v107 := GenS7{F0: v5, F1: 32767, F2: (v5 / ((v5) | 1))}
                                                    _ = len(v1)
                                                    _ = v107
                                                } else {
                                                    v108 := GenS9{F0: 32767, F1: "a", F2: 65535}
                                                    _ = v108
                                                }
                                                if (47 >= 779) {
                                                    v3 = v3
                                                }
                                                v109 := GenS10{F0: -1, F1: (true && v100)}
                                                _ = v109
                                            } else {
                                                if v100 {
                                                    if v100 {
                                                        v5 = genFn0(669, (fi106 | (fi104 + (fi104 << ((255) & 63)))))
                                                        switch v2 {
                                                        case "foo bar":
                                                            v3 += v3
                                                        case "hello":
                                                            var v110 int32 = 1
                                                            _ = v110
                                                        }
                                                    } else {
                                                        if v86 {
                                                            v111 := GenS11{F0: 65535, F1: -1}
                                                            v100 = v100
                                                            _ = v111
                                                        } else {
                                                            v100 = (false || (2147483648 >= 2147483647))
                                                        }
                                                        var v112 int8 = 34
                                                        v87 -= v87
                                                        _ = v112
                                                    }
                                                    var v113 int32 = 42
                                                    _ = v113
                                                }
                                            }
                                            v3 = uint8(v87)
                                            if v100 {
                                                fmt.Println(v100, v3, v3)
                                            }
                                            _ = fi106
                                        }
                                    case -147:
                                        var v114 bool = v100
                                        _ = v114
                                    default:
                                        v115 := GenS11{F0: uint(v5), F1: -1}
                                        _ = v115
                                    }
                                    var v116 float32 = 0.1
                                    for fi117 := 0; fi117 < 8; fi117++ {
                                        var v118 int8 = 42
                                        _ = fi117
                                        _ = v118
                                    }
                                    _ = v116
                                } else {
                                    var v119 string = v2
                                    fmt.Println(v87, v5, v100)
                                    _ = v119
                                }
                            default:
                                _ = len(v1)
                            }
                        } else {
                            for fi120 := 0; fi120 < 9; fi120++ {
                                fmt.Println(v100, v5)
                                _ = fi120
                            }
                            switch (v2 + v2) {
                            case "hello":
                                if v86 {
                                    for fi121 := 0; fi121 < 16; fi121++ {
                                        var v122 uint8 = v3
                                        _ = fi121
                                        _ = v122
                                    }
                                    for fi123 := 0; fi123 < 9; fi123++ {
                                        switch (v2 + ("\n\t" + (((v2 + v2) + ("foo bar" + (v2 + ((v2 + v2) + (v2 + v2))))) + "\n\t"))) {
                                        case "hello":
                                            v124 := GenS12{F0: 42}
                                            _ = v124
                                        case "":
                                            fmt.Println(fi123, fi123)
                                        case "foo bar":
                                            var v125 uint64 = genFn0(65535, -324)
                                            _ = v125
                                        default:
                                            fmt.Println(fi104)
                                        }
                                        v87 = uint16(v3)
                                        v126 := GenS13{F0: v3}
                                        fmt.Println(v87, v3)
                                        v127 := GenS14{F0: v87}
                                        _ = fi123
                                        _ = v126
                                        _ = v127
                                    }
                                    if false {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                            }
                        }
                    default:
                        for fi128 := 0; fi128 < 6; fi128++ {
                            v2 = v2
                            _ = fi128
                        }
                    }
                default:
                    if (v86 && v100) {
                        if ((v2 + ((v2 + v2) + "")) > (v2 + v2)) {
                            v1 = append(v1, ("foo bar" + ((v2 + v2) + v2)))
                        } else {
                            _ = len(v4)
                        }
                        switch ((v2 + v2) + v2) {
                        case "foo bar":
                            var v129 bool = (-67 > (fi104 % ((fi104) | 1)))
                            _ = v129
                        case "hello":
                            var v130 int8 = -96
                            _ = v130
                        case "a":
                            var v131 float64 = float64((3.14159 + 0.0))
                            _ = v131
                        default:
                            for fi132 := 0; fi132 < 10; fi132++ {
                                fmt.Println(v3, fi104)
                                v2 = v2
                                _ = fi132
                            }
                        }
                        for fi133 := 0; fi133 < 16; fi133++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v100 = false
                            var v134 int8 = int8(v87)
                            _ = fi133
                            _ = v134
                        }
                    }
                }
                v86 = ((^(v3)) <= 127)
                _ = fi104
            }
        }
        _ = v100
    }
    var v135 int8 = 109
    fmt.Println(v87, v135, v2)
    if true {
        v86 = v86
        switch 32768 {
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 1:
            for fi136 := 0; fi136 < 11; fi136++ {
                if true {
                    v137 := GenS13{F0: (v3 ^ 80)}
                    _ = v137
                }
                v1 = append(v1, "hello")
                _ = fi136
            }
        case 2147483647:
            for fi138 := 0; fi138 < 10; fi138++ {
                _ = len(v1)
                _ = fi138
            }
        default:
            for fi139 := 0; fi139 < 2; fi139++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi139
            }
        }
    }
    var v140 int16 = 1000
    v141 := GenS15{F0: 255, F1: 32768, F2: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v135:", v135)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v135
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
}
