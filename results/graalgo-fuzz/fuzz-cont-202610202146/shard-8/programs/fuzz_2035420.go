package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 string; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint; F1 uint64; F2 int }
type GenS9 struct { F0 string; F1 int16; F2 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 uint32; F2 bool }
type GenS12 struct { F0 uint; F1 uint16; F2 uint }
type GenS13 struct { F0 int32; F1 bool; F2 bool }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int8; F1 int8; F2 int16 }
type GenS16 struct { F0 int16; F1 int32; F2 int16 }
type GenS17 struct { F0 bool; F1 bool; F2 uint }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint32; F1 uint; F2 int8 }
func genFn0(p0 uint16, p1 int8) bool {
    switch "a" {
    case "":
        switch "" {
        case "foo bar":
            var v7 int32 = 71
            _ = v7
        case "":
            v8 := []int{}
            _ = v8
        default:
            var v9 string = "a"
            _ = v9
        }
    case "\n\t":
        var v10 float64 = 3.14159
        _ = v10
    }
    if (true || ((-903 == -2147483648) || (!false))) {
        var v11 uint = 18446744073709551615
        v12 := []int{207, 42, 128}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    return (false && false)
}

func main() {
    var v1 int = 42
    var v2 int32 = 1
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 113
    var v5 uint64 = 872
    v6 := []string{"foo bar", (("hello" + "") + "\n\t")}
    var v13 uint32 = 49
    if false {
        fmt.Println(v1)
        _ = len(v3)
        switch ("hello" + "hello") {
        case "\n\t":
            v14 := GenS1{F0: true, F1: -74}
            _ = v14
        case "foo bar":
            v4 = 0
        case "hello":
            fmt.Println(v2, v13)
        }
    }
    if false {
        if true {
            if (!true) {
                v2 = v2
            } else {
                for fi15 := 0; fi15 < 6; fi15++ {
                    switch 1 {
                    case -1:
                        v16 := GenS1{F0: false, F1: (v2 % ((-1) | 1))}
                        _ = v16
                    case 127:
                        switch ("" + "") {
                        case "a":
                            if (9223372036854775807 < -566) {
                                var v17 uint8 = (v4 << ((18446744073709551615) & 63))
                                _ = v17
                            } else {
                                if true {
                                    if true {
                                        var v18 float32 = 0.1
                                        if true {
                                            if true {
                                                v2 -= int32(v5)
                                                if ((v2 % ((42) | 1)) == v2) {
                                                    for fi19 := 0; fi19 < 14; fi19++ {
                                                        var v20 uint = 42
                                                        v21 := GenS1{F0: (!false), F1: (v2 | 1)}
                                                        _ = fi19
                                                        _ = v20
                                                        _ = v21
                                                    }
                                                } else {
                                                    fmt.Println(v5)
                                                    v22 := v3[(v13 % ((v13) | 1))]
                                                    _ = v22
                                                }
                                                var v23 int16 = 1
                                                _ = v23
                                            } else {
                                                switch "" {
                                                case "":
                                                    v24 := GenS1{F0: false, F1: (v2 % ((42) | 1))}
                                                    _ = v24
                                                case "foo bar":
                                                    for fi25 := 0; fi25 < 2; fi25++ {
                                                        v26 := GenS1{F0: true, F1: (v2 - v2)}
                                                        var v27 uint16 = 32767
                                                        _ = fi25
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                }
                                                for fi28 := 0; fi28 < 16; fi28++ {
                                                    if (!(58 >= -128)) {
                                                        var v29 string = ("" + ("foo bar" + (("hello" + ((("" + "foo bar") + "") + "\n\t")) + "hello")))
                                                        _ = v29
                                                    } else {
                                                        v3[1] = 2
                                                        v13 -= v13
                                                        fmt.Println(v18, v1)
                                                    }
                                                    _ = len(v6)
                                                    v30 := GenS2{F0: (v2 >> ((1) & 63)), F1: 0, F2: 947}
                                                    fmt.Println(fi15, v2)
                                                    _ = fi28
                                                    _ = v30
                                                }
                                                v18 += 3.14159
                                            }
                                            v31 := GenS3{F0: 1000}
                                            _ = v31
                                        }
                                        _ = len(v6)
                                        _ = v18
                                    }
                                } else {
                                    if false {
                                        for fi32 := 0; fi32 < 9; fi32++ {
                                            if false {
                                                v33 := GenS2{F0: -1, F1: (^(v4)), F2: -9223372036854775808}
                                                _ = v33
                                            } else {
                                                if true {
                                                    fmt.Println(v1, v4, v5)
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v34 uint64 = (v5 + (v5 >> ((65535) & 63)))
                                                        _ = v34
                                                    case "hello":
                                                        v1 = (fi32 << ((2) & 63))
                                                    case "a":
                                                        if false {
                                                            v35 := GenS4{F0: ((v13 == (v13 | v13)) && false), F1: ("\n\t" + "hello"), F2: "\n\t"}
                                                            _ = v35
                                                        } else {
                                                            var v36 uint16 = 255
                                                            switch "hello" {
                                                            case "a":
                                                                if true {
                                                                    switch (fi15 / ((fi32) | 1)) {
                                                                    case 543:
                                                                        v1 += (v1 ^ 1)
                                                                    case 4294967295:
                                                                        v36 -= (^(v36))
                                                                    }
                                                                    v37 := GenS5{F0: fi15}
                                                                    v38 := GenS6{F0: -987, F1: ""}
                                                                    v1 += fi32
                                                                    _ = v37
                                                                    _ = v38
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v2)
                                                                        v3[v13] = (v13 - (v13 % ((uint32(v5)) | 1)))
                                                                    }
                                                                    if (!(v36 >= (^(v36)))) {
                                                                        var v39 uint8 = 128
                                                                        if true {
                                                                            if false {
                                                                                var v40 int32 = (v2 % ((v2) | 1))
                                                                                _ = v40
                                                                            }
                                                                            v41 := GenS7{F0: "\n\t", F1: true}
                                                                            _ = v41
                                                                        }
                                                                        v2 += v2
                                                                        _ = v39
                                                                    }
                                                                    var v42 float32 = 1.5
                                                                    _ = v42
                                                                }
                                                            case "":
                                                                v43 := GenS4{F0: (!(true && (true || false))), F1: "foo bar", F2: "\n\t"}
                                                                _ = v43
                                                            case "\n\t":
                                                                if ((-2.25 / 1e10) < 1.5) {
                                                                    _ = len(v3)
                                                                    var v44 float64 = ((0.1 * 0.1) / 1.5)
                                                                    _ = v44
                                                                }
                                                            default:
                                                                _ = len(v6)
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            _ = v36
                                                        }
                                                    default:
                                                        v45 := GenS8{F0: 1, F1: (^(v5)), F2: (fi32 - 4294967295)}
                                                        var v46 uint8 = (v4 * (v4 / ((42) | 1)))
                                                        delete(v3, (v13 >> ((520) & 63)))
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    v4 = (v4 << ((255) & 63))
                                                } else {
                                                    v1 = -2147483648
                                                    if (!true) {
                                                        v5 = v5
                                                    }
                                                }
                                                var v47 int = (fi32 % ((fi32) | 1))
                                                _ = v47
                                            }
                                            var v48 string = "foo bar"
                                            if (-2.25 > 3.14159) {
                                                var v49 bool = false
                                                _ = v49
                                            } else {
                                                v3[0] = (v13 & v13)
                                            }
                                            if false {
                                                var v50 int8 = 79
                                                var v51 uint64 = v5
                                                _ = v50
                                                _ = v51
                                            } else {
                                                var v52 uint8 = v4
                                                if false {
                                                    switch "hello" {
                                                    case "":
                                                        v53 := GenS9{F0: "a", F1: 0, F2: 100}
                                                        _ = v53
                                                    default:
                                                        v54 := GenS10{F0: (v52 + 233)}
                                                        _ = v54
                                                    }
                                                }
                                                v1 += (v1 >> ((0) & 63))
                                                _ = v52
                                            }
                                            _ = fi32
                                            _ = v48
                                        }
                                        v3[v13] = v13
                                        if (-1 == 42) {
                                            if (true && (2 > (^(v13)))) {
                                                var v55 int64 = 1
                                                v6 = append(v6, "foo bar")
                                                _ = v55
                                            }
                                        } else {
                                            var v56 int = (-(v1))
                                            _ = v56
                                        }
                                    } else {
                                        var v57 int8 = 42
                                        _ = v57
                                    }
                                    for fi58 := 0; fi58 < 6; fi58++ {
                                        switch (fi58 >> ((uint(v5)) & 63)) {
                                        case -2147483648:
                                            var v59 string = "a"
                                            _ = v59
                                        case 4294967295:
                                            switch ("\n\t" + "a") {
                                            case "hello":
                                                v5 = v5
                                                var v60 uint32 = (v13 << ((746) & 63))
                                                var v61 uint16 = 1
                                                _ = v60
                                                _ = v61
                                            case "foo bar":
                                                v62 := GenS11{F0: 335, F1: (v13 / (((v13 << ((263) & 63))) | 1)), F2: false}
                                                _ = v62
                                            default:
                                                delete(v3, 271)
                                            }
                                        case 127:
                                            var v63 int64 = -896
                                            _ = v63
                                        default:
                                            if ((true || (!(v2 <= (v2 / ((v2) | 1))))) && ((3.14159 >= (3.14159 * float64(-2.25))) && true)) {
                                                switch 42 {
                                                case 9223372036854775807:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                case -9223372036854775808:
                                                    if ((-2.25 <= 1.5) || (((true && true) && false) && true)) {
                                                        var v64 int64 = -808
                                                        _ = v64
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v65 uint8 = v4
                                                            _ = v65
                                                        }
                                                    }
                                                case 32767:
                                                    fmt.Println(fi58, v4, fi58)
                                                }
                                                var v66 uint16 = 32767
                                                if true {
                                                    v67 := GenS12{F0: 1, F1: 255, F2: 42}
                                                    fmt.Println(v4, v5)
                                                    _ = v67
                                                } else {
                                                    switch "" {
                                                    case "\n\t":
                                                        v68 := GenS13{F0: v2, F1: (46 >= -128), F2: (-990 != 0)}
                                                        _ = v68
                                                    case "hello":
                                                        if (true && false) {
                                                            switch ("a" + ("foo bar" + (("hello" + ("a" + (("foo bar" + "a") + ("hello" + "")))) + "\n\t"))) {
                                                            case "hello":
                                                                switch fi15 {
                                                                case 664:
                                                                    v4 = (v4 & v4)
                                                                }
                                                            }
                                                            if false {
                                                                if (!(!(false || (9223372036854775807 >= 748)))) {
                                                                    var v69 uint64 = (v5 >> ((uint(v13)) & 63))
                                                                    var v70 float32 = (0.0 + (3.14159 / 1.5))
                                                                    _ = v69
                                                                    _ = v70
                                                                } else {
                                                                    fmt.Println(fi58, fi58, v5)
                                                                    fmt.Println(v5, fi15, v5)
                                                                }
                                                                v5 = (v5 * uint64(v66))
                                                                if (true || false) {
                                                                    fmt.Println(v4, v13)
                                                                    v71 := GenS9{F0: ("hello" + "foo bar"), F1: -32768, F2: 106}
                                                                    v5 -= (v5 - (^(v5)))
                                                                    _ = v71
                                                                }
                                                            }
                                                            var v72 uint = 2
                                                            v2 -= 2147483647
                                                            _ = v72
                                                        } else {
                                                            var v73 uint = 42
                                                            _ = v73
                                                        }
                                                    case "":
                                                        v3[(v13 ^ (^(v13)))] = v13
                                                    default:
                                                        var v74 int32 = v2
                                                        _ = v74
                                                    }
                                                    if (uint16(v4) == (v66 | v66)) {
                                                        fmt.Println(v1, fi58)
                                                        var v75 uint8 = v4
                                                        _ = v75
                                                    }
                                                }
                                                v2 = (v2 + (v2 % (((v2 - 0)) | 1)))
                                                switch (v1 - 9223372036854775807) {
                                                case 0:
                                                    var v76 int8 = 42
                                                    _ = v76
                                                }
                                                v77 := v3[v13]
                                                fmt.Println(v77)
                                                _ = v66
                                                _ = v77
                                            }
                                        }
                                        var v78 uint8 = v4
                                        _ = fi58
                                        _ = v78
                                    }
                                    v4 -= (v4 * (v4 >> ((2) & 63)))
                                }
                            }
                        case "hello":
                            if len(v6) > 0 {
                                v6[0] = "hello"
                            }
                        case "foo bar":
                            switch 1 {
                            case 32768:
                                var v79 uint8 = v4
                                _ = v79
                            case -179:
                                v4 = (^(v4))
                            }
                        default:
                            v4 -= v4
                        }
                    default:
                        var v80 uint32 = (v13 ^ 119)
                        _ = v80
                    }
                    _ = fi15
                }
                for fi81 := 0; fi81 < 20; fi81++ {
                    v2 = -1
                    _ = fi81
                }
            }
            if false {
                v82 := GenS2{F0: (^(v2)), F1: (v4 ^ v4), F2: int64(v2)}
                fmt.Println(v2)
                _ = v82
            } else {
                v83 := GenS6{F0: 0, F1: ""}
                switch (-(v1)) {
                case 974:
                    fmt.Println(v2, v1)
                }
                _ = v83
            }
            if len(v6) > 0 {
                v6[0] = ("" + "")
            }
        } else {
            var v84 int8 = 100
            var v85 int = (v1 + 128)
            for fi86 := 0; fi86 < 3; fi86++ {
                v84 += v84
                if ((!(1e10 > 0.0)) || false) {
                    fmt.Println(v4)
                } else {
                    if true {
                        fmt.Println(v85, fi86, v13)
                        var v87 int64 = 42
                        switch "foo bar" {
                        case "hello":
                            v88 := v3[(v13 / ((167) | 1))]
                            _ = v88
                        case "\n\t":
                            for fi89 := 0; fi89 < 7; fi89++ {
                                v90 := GenS8{F0: 42, F1: v5, F2: (fi89 + (fi86 | fi89))}
                                switch -2147483648 {
                                case -1:
                                    fmt.Println(v13, v84, v13)
                                case 9223372036854775807:
                                    var v91 int8 = (v84 << ((105) & 63))
                                    v92 := GenS14{F0: -9223372036854775808}
                                    v2 += (v2 / ((v2) | 1))
                                    _ = v91
                                    _ = v92
                                }
                                v13 -= (^(v13))
                                v93 := GenS9{F0: ("" + "\n\t"), F1: 32767, F2: (v84 - (^(v84)))}
                                if false {
                                    switch 42 {
                                    case 42:
                                        fmt.Println(fi89, v87)
                                    case 2147483647:
                                        _ = len(v6)
                                    }
                                    fmt.Println(v85, v13, v4)
                                } else {
                                    v3[v13] = v13
                                    if len(v6) > 0 {
                                        v6[0] = "\n\t"
                                    }
                                }
                                v1 = 345
                                switch (^(fi86)) {
                                case 2147483647:
                                    fmt.Println(v85, v4)
                                case 836:
                                    var v94 string = ("foo bar" + "hello")
                                    _ = v94
                                default:
                                    v4 = 111
                                }
                                var v95 float64 = 1e10
                                _ = fi89
                                _ = v90
                                _ = v93
                                _ = v95
                            }
                        case "":
                            switch "" {
                            case "a":
                                if true {
                                    switch "hello" {
                                    case "foo bar":
                                        var v96 float64 = 3.14159
                                        _ = v96
                                    case "hello":
                                        _ = len(v3)
                                    case "":
                                        v97 := GenS15{F0: (v84 ^ (v84 - v84)), F1: (v84 >> ((1) & 63)), F2: -1}
                                        _ = v97
                                    }
                                } else {
                                    var v98 int32 = (^(v2))
                                    v5 -= (v5 * v5)
                                    if false {
                                        var v99 int8 = -53
                                        v100 := v3[v13]
                                        _ = v100
                                        _ = v99
                                    }
                                    _ = v98
                                }
                            case "hello":
                                for fi101 := 0; fi101 < 14; fi101++ {
                                    v102 := GenS16{F0: -649, F1: (v2 % ((2147483647) | 1)), F2: -1}
                                    var v103 int16 = int16(v5)
                                    if (!false) {
                                        var v104 float32 = ((float32((float64((1.5 + (1e10 + 0.0))) - 3.14159)) / (0.1 - -2.25)) * 1.5)
                                        fmt.Println(v84)
                                        _ = v104
                                    }
                                    _ = fi101
                                    _ = v102
                                    _ = v103
                                }
                            case "":
                                switch 530 {
                                case 1:
                                    for fi105 := 0; fi105 < 16; fi105++ {
                                        var v106 uint32 = v13
                                        _ = fi105
                                        _ = v106
                                    }
                                case 127:
                                    if (false && true) {
                                        if (((0.0 + 1.5) >= float64(1e10)) || (!(true && (true || (!((v4 << ((872) & 63)) >= 1)))))) {
                                            switch "foo bar" {
                                            case "a":
                                                if (4294967295 == (v13 | v13)) {
                                                    fmt.Println(v4)
                                                    switch "a" {
                                                    case "":
                                                        for fi107 := 0; fi107 < 4; fi107++ {
                                                            var v108 uint16 = 255
                                                            var v109 int64 = v87
                                                            _ = fi107
                                                            _ = v108
                                                            _ = v109
                                                        }
                                                    default:
                                                        v110 := GenS17{F0: (!false), F1: (uint16(v13) <= 2), F2: 4294967295}
                                                        _ = v110
                                                    }
                                                    fmt.Println(v4, v4, v87)
                                                }
                                            case "foo bar":
                                                v111 := v3[v13]
                                                _ = v111
                                            }
                                            for fi112 := 0; fi112 < 13; fi112++ {
                                                v2 += (v2 + v2)
                                                var v113 uint16 = 42
                                                var v114 int32 = -610
                                                _ = fi112
                                                _ = v113
                                                _ = v114
                                            }
                                            v115 := GenS18{F0: "a"}
                                            _ = v115
                                        }
                                    } else {
                                        if false {
                                            v13 -= 65535
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = len(v6)
                                        } else {
                                            switch fi86 {
                                            case -9223372036854775808:
                                                switch "" {
                                                case "\n\t":
                                                    var v116 float32 = (1e10 * (1e10 / 0.1))
                                                    _ = v116
                                                case "foo bar":
                                                    v84 -= v84
                                                default:
                                                    var v117 int64 = -36
                                                    _ = v117
                                                }
                                            case 32768:
                                                v4 += (v4 << ((2) & 63))
                                            default:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            for fi118 := 0; fi118 < 7; fi118++ {
                                                _ = len(v6)
                                                _ = fi118
                                            }
                                            var v119 uint8 = (v4 << ((uint(v5)) & 63))
                                            _ = v119
                                        }
                                        var v120 int64 = v87
                                        _ = v120
                                    }
                                case -2147483648:
                                    for fi121 := 0; fi121 < 4; fi121++ {
                                        delete(v3, (v13 * 2147483648))
                                        var v122 uint32 = 65535
                                        _ = fi121
                                        _ = v122
                                    }
                                default:
                                    v123 := GenS14{F0: (^(v87))}
                                    _ = v123
                                }
                            default:
                                var v124 uint64 = 42
                                _ = v124
                            }
                        default:
                            for fi125 := 0; fi125 < 8; fi125++ {
                                var v126 uint32 = v13
                                v6 = append(v6, "\n\t")
                                _ = fi125
                                _ = v126
                            }
                        }
                        _ = v87
                    } else {
                        v4 += v4
                        var v127 int16 = -32768
                        _ = v127
                    }
                    v84 = (v84 + (v84 + v84))
                }
                _ = fi86
            }
            _ = v84
            _ = v85
        }
    } else {
        delete(v3, 902)
        var v128 int = 2147483647
        _ = v128
    }
    v13 -= v13
    if (!false) {
        if ("a" > ("a" + ((("a" + ("\n\t" + "")) + "foo bar") + "a"))) {
            v1 = (v1 % (((v1 - (v1 % (((v1 & 2147483647)) | 1)))) | 1))
        } else {
            for fi129 := 0; fi129 < 10; fi129++ {
                _ = len(v6)
                v130 := GenS7{F0: ("a" + ("" + "")), F1: ((((0.0 * -2.25) + (1e10 * 1e10)) - (0.0 * 1.5)) != 0.0)}
                var v131 float64 = 0.1
                _ = fi129
                _ = v130
                _ = v131
            }
            v3[2] = 1
            if (-2.25 == 1.5) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v132 float64 = float64(1e10)
                fmt.Println(v13, v5, v2)
                _ = v132
            } else {
                for fi133 := 0; fi133 < 19; fi133++ {
                    for fi134 := 0; fi134 < 15; fi134++ {
                        var v135 int64 = -1
                        v5 += v5
                        _ = fi134
                        _ = v135
                    }
                    _ = fi133
                }
            }
        }
        var v136 int = v1
        v137 := GenS7{F0: "a", F1: false}
        _ = v136
        _ = v137
    } else {
        var v138 uint16 = 957
        v139 := GenS4{F0: (false && false), F1: "a", F2: "hello"}
        _ = v138
        _ = v139
    }
    for fi140 := 0; fi140 < 10; fi140++ {
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v4 = v4
            switch 0 {
            case -742:
                fmt.Println(v2)
            default:
                if false {
                    var v141 int16 = 1
                    var v142 int32 = (v2 * v2)
                    _ = v141
                    _ = v142
                } else {
                    fmt.Println(v2)
                    var v143 float32 = (float32(0.1) * 0.0)
                    for fi144 := 0; fi144 < 9; fi144++ {
                        var v145 string = ""
                        _ = fi144
                        _ = v145
                    }
                    _ = v143
                }
            }
        }
        switch 1 {
        case 42:
            v5 += 1
        case -9223372036854775808:
            v146 := GenS19{F0: (v13 | (v13 >> ((65535) & 63))), F1: 255, F2: 127}
            _ = v146
        default:
            v5 = (v5 * v5)
        }
        v3[(^(v13))] = v13
        _ = fi140
    }
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
