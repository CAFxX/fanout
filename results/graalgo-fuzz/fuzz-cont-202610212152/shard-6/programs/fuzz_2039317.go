package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint; F1 int; F2 int16 }
type GenS12 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS13 struct { F0 uint32; F1 int32; F2 bool }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int16; F1 bool; F2 uint }
type GenS19 struct { F0 int64; F1 bool }
type GenS20 struct { F0 uint32; F1 uint64 }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 uint; F1 int; F2 uint64 }
func genFn0(p0 uint8, p1 string, p2 string) uint64 {
    var v6 uint64 = 0
    v7 := GenS1{F0: 65535}
    v8 := []int{704}
    v6 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (^(v6))
}
func genFn1(p0 int32, p1 int32) uint16 {
    if (578 >= 1) {
        switch "hello" {
        case "":
            fmt.Println(p0, p0)
        default:
            var v9 uint16 = 255
            _ = v9
        }
        v10 := GenS1{F0: 4294967295}
        _ = v10
    } else {
        v11 := []int{437}
        for fi12 := 0; fi12 < 12; fi12++ {
            if true {
                switch "foo bar" {
                case "foo bar":
                    var v13 uint64 = 42
                    _ = v13
                case "":
                    p0 += (p1 | p1)
                case "hello":
                    var v14 bool = (!(1 <= 0))
                    _ = v14
                default:
                    var v15 int32 = p0
                    _ = v15
                }
            }
            if false {
                var v16 uint32 = 2
                p1 += (p0 - p1)
                _ = v16
            } else {
                if false {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        var v18 bool = (!false)
                        v18 = (0 != 0)
                        if (42 != 65535) {
                            if len(v11) > 0 {
                                v11[0] = fi17
                            }
                        } else {
                            var v19 int8 = 42
                            v18 = true
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    var v20 float64 = 0.0
                    p1 += (^(p0))
                    _ = v20
                } else {
                    var v21 bool = true
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v21
                }
                if false {
                    fmt.Println(fi12, fi12, fi12)
                } else {
                    var v22 uint8 = 127
                    var v23 uint = 42
                    _ = v22
                    _ = v23
                }
            }
            v24 := GenS2{F0: 4294967295, F1: 2, F2: 1}
            _ = fi12
            _ = v24
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 uint16 = 1
            switch ("a" + (("" + (("foo bar" + "\n\t") + "foo bar")) + ("\n\t" + "foo bar"))) {
            case "foo bar":
                p0 += p0
            case "a":
                p1 -= (p1 << ((159) & 63))
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi25
            _ = v26
        }
        _ = v11
    }
    var v27 float32 = 3.14159
    var v28 int32 = (p1 | (p0 & (p0 ^ 320)))
    v29 := []int{9223372036854775807}
    var v30 uint16 = 65535
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v29
    _ = v30
    return v30
}

func main() {
    v1 := []uint32{1, 4294967295, 2147483648, 0}
    v2 := []int{1, -9223372036854775808}
    var v3 float32 = (float32(1e10) * 3.14159)
    v4 := []uint32{2, 834, 322}
    var v5 int = 128
    var v31 string = ""
    var v32 bool = false
    fmt.Println(v32)
    v33 := GenS1{F0: 0}
    v3 -= ((v3 + 0.1) - 0.1)
    fmt.Println(v32)
    fmt.Println(v32, v5, v5)
    var v34 int64 = -614
    v3 += v3
    v5 += (v5 - v5)
    fmt.Println(v34, v3)
    switch ("foo bar" + "a") {
    case "hello":
        v35 := GenS2{F0: 0, F1: 255, F2: 2}
        _ = v35
    case "a":
        v31 = "foo bar"
    case "\n\t":
        if false {
            if true {
                v36 := GenS3{F0: -144, F1: 2147483648}
                _ = v36
            }
            for fi37 := 0; fi37 < 15; fi37++ {
                if (1 == 212) {
                    if (207 <= 246) {
                        if true {
                            v32 = v32
                            v32 = v32
                        } else {
                            v3 = (v3 * v3)
                            if len(v2) > 0 {
                                v2[0] = 32767
                            }
                            switch 42 {
                            case 0:
                                v32 = (2147483647 == -82)
                            default:
                                if ((791 == 358) || v32) {
                                    if v32 {
                                        var v38 bool = v32
                                        switch v31 {
                                        case "a":
                                            fmt.Println(fi37, v5, v3)
                                        case "\n\t":
                                            var v39 int16 = 42
                                            _ = v39
                                        case "hello":
                                            switch -1 {
                                            case -1:
                                                var v40 float32 = (v3 - v3)
                                                _ = v40
                                            case 2147483647:
                                                for fi41 := 0; fi41 < 6; fi41++ {
                                                    var v42 int64 = (v34 >> ((4294967295) & 63))
                                                    if ((fi41 < fi37) && (((0.1 + v3) <= (0.0 * -2.25)) && ((((-(v5)) >= (^(v5))) && (1.5 < 0.1)) && v32))) {
                                                        fmt.Println(v34)
                                                    }
                                                    var v43 uint8 = 255
                                                    if (-9223372036854775808 > v42) {
                                                        if false {
                                                            if (0 == 42) {
                                                                var v44 uint64 = 1
                                                                switch v31 {
                                                                case "\n\t":
                                                                    v31 = ((("" + "") + ((v31 + ("hello" + v31)) + "a")) + ((("foo bar" + ((v31 + ("" + "a")) + v31)) + "") + (("hello" + (v31 + (v31 + ("\n\t" + (v31 + (v31 + (v31 + ((((v31 + v31) + v31) + (v31 + "\n\t")) + ("\n\t" + ("\n\t" + v31)))))))))) + (v31 + "a"))))
                                                                }
                                                                var v45 int32 = -1
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                v43 = 0
                                                            }
                                                        }
                                                        switch v5 {
                                                        case 32767:
                                                            switch (v31 + (v31 + "a")) {
                                                            case "hello":
                                                                v38 = (255 == 42)
                                                            case "foo bar":
                                                                v31 = "hello"
                                                            case "a":
                                                                var v46 uint32 = 0
                                                                var v47 uint32 = v46
                                                                v3 -= v3
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        default:
                                                            if (false && v32) {
                                                                var v48 uint64 = genFn0(242, v31, "hello")
                                                                switch (("foo bar" + "") + ("hello" + ("a" + (v31 + v31)))) {
                                                                case "foo bar":
                                                                    fmt.Println(v3, v31)
                                                                case "hello":
                                                                    var v49 uint64 = (v48 / ((genFn0(2, ((((((v31 + (v31 + ((("\n\t" + v31) + (v31 + v31)) + "hello"))) + ("a" + "hello")) + v31) + v31) + v31) + "hello"), v31)) | 1))
                                                                    _ = v49
                                                                default:
                                                                    if false {
                                                                        v50 := GenS3{F0: -32768, F1: 643}
                                                                        if v38 {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        }
                                                                        switch (-(fi41)) {
                                                                        case 1:
                                                                            if v38 {
                                                                                v51 := GenS4{F0: -259}
                                                                                switch 1 {
                                                                                case 1:
                                                                                    v42 = (v34 * v42)
                                                                                case 4294967295:
                                                                                    _ = len(v2)
                                                                                }
                                                                                if (true && (127 < 0)) {
                                                                                    v42 = v42
                                                                                } else {
                                                                                    if (v38 && (1 <= 208)) {
                                                                                        fmt.Println(v5)
                                                                                    }
                                                                                    v2 = append(v2, int(v48))
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 1
                                                                                    }
                                                                                }
                                                                                _ = v51
                                                                            } else {
                                                                                v5 += (fi41 % (((fi37 - 128)) | 1))
                                                                                v38 = (true && (!(2147483647 != 0)))
                                                                            }
                                                                        }
                                                                        v52 := GenS3{F0: 617, F1: 281}
                                                                        var v53 bool = (!v38)
                                                                        _ = v50
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                }
                                                                _ = v48
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = (fi37 | (fi37 << ((446) & 63)))
                                                                }
                                                            }
                                                        }
                                                        var v54 uint = 42
                                                        _ = v54
                                                    } else {
                                                        v55 := GenS3{F0: 1000, F1: 1}
                                                        _ = v55
                                                    }
                                                    _ = fi41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            default:
                                                var v56 float64 = 1e10
                                                _ = v56
                                            }
                                        default:
                                            v32 = (1 >= 2147483647)
                                        }
                                        v57 := GenS5{F0: 524}
                                        _ = v38
                                        _ = v57
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    if (((128 < 190) || (true && (1 >= 255))) && false) {
                                        switch v5 {
                                        case -541:
                                            _ = len(v1)
                                        case -9223372036854775808:
                                            var v58 int32 = -1
                                            _ = v58
                                        case 128:
                                            switch (-(fi37)) {
                                            case 127:
                                                var v59 uint16 = 0
                                                _ = v59
                                            case 42:
                                                if (false || v32) {
                                                    if ((v31 + v31) != ("a" + v31)) {
                                                        switch "\n\t" {
                                                        case "":
                                                            switch fi37 {
                                                            case -2147483648:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            case -1:
                                                                if (v34 >= v34) {
                                                                    for fi60 := 0; fi60 < 1; fi60++ {
                                                                        var v61 int32 = 0
                                                                        var v62 uint64 = 18446744073709551615
                                                                        _ = fi60
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    for fi63 := 0; fi63 < 15; fi63++ {
                                                                        v64 := GenS2{F0: 409, F1: 56, F2: 4294967295}
                                                                        v65 := GenS2{F0: 18446744073709551615, F1: 276, F2: 731}
                                                                        _ = fi63
                                                                        _ = v64
                                                                        _ = v65
                                                                    }
                                                                    var v66 int8 = -128
                                                                    _ = v66
                                                                } else {
                                                                    v34 -= -9223372036854775808
                                                                    switch v5 {
                                                                    case 0:
                                                                        if v32 {
                                                                            v31 = "a"
                                                                        } else {
                                                                            var v67 uint8 = 42
                                                                            _ = v67
                                                                        }
                                                                    case 20:
                                                                        var v68 string = ((v31 + ("foo bar" + v31)) + v31)
                                                                        _ = v68
                                                                    default:
                                                                        for fi69 := 0; fi69 < 15; fi69++ {
                                                                            if true {
                                                                                var v70 uint8 = 42
                                                                                _ = v70
                                                                            }
                                                                            _ = fi69
                                                                        }
                                                                    }
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                        case "a":
                                                            var v71 string = (((v31 + ("a" + ((v31 + v31) + v31))) + ("hello" + v31)) + "")
                                                            _ = v71
                                                        default:
                                                            v2 = append(v2, (-(fi37)))
                                                        }
                                                    } else {
                                                        for fi72 := 0; fi72 < 14; fi72++ {
                                                            switch (v31 + "\n\t") {
                                                            case "\n\t":
                                                                if (v32 && ((false || v32) && v32)) {
                                                                    switch ((("foo bar" + (("a" + ((v31 + (v31 + ((v31 + "foo bar") + "hello"))) + "")) + v31)) + ((v31 + v31) + v31)) + "\n\t") {
                                                                    case "a":
                                                                        v5 -= fi72
                                                                    case "":
                                                                        var v73 float32 = v3
                                                                        _ = v73
                                                                    case "hello":
                                                                        v34 -= 9223372036854775807
                                                                    }
                                                                }
                                                            case "hello":
                                                                fmt.Println(fi37)
                                                            default:
                                                                v1 = append(v1, 85)
                                                            }
                                                            v32 = v32
                                                            switch 791 {
                                                            case -92:
                                                                v3 += v3
                                                            case 127:
                                                                if true {
                                                                    var v74 uint32 = 1
                                                                    fmt.Println(fi37, v34)
                                                                    v5 += (v5 / (((v5 ^ int(v74))) | 1))
                                                                    _ = v74
                                                                } else {
                                                                    var v75 int8 = 0
                                                                    if v32 {
                                                                        if (!v32) {
                                                                            _ = len(v4)
                                                                            v76 := GenS6{F0: 2, F1: genFn0(22, "a", v31), F2: 0}
                                                                            v5 = fi72
                                                                            _ = v76
                                                                        } else {
                                                                            fmt.Println(v32, v32)
                                                                            if (!(32768 >= fi37)) {
                                                                                v77 := GenS5{F0: 695}
                                                                                if (0.1 > (-2.25 + (((1.5 * 3.14159) + v3) + v3))) {
                                                                                    if (!((v5 >> ((425) & 63)) != fi72)) {
                                                                                        var v78 uint8 = 228
                                                                                        _ = v78
                                                                                    }
                                                                                    var v79 uint32 = 1
                                                                                    if true {
                                                                                        switch ((v31 + "a") + "\n\t") {
                                                                                        case "":
                                                                                            var v80 float32 = v3
                                                                                            _ = v80
                                                                                        default:
                                                                                            var v81 float32 = v3
                                                                                            _ = v81
                                                                                        }
                                                                                        v82 := GenS7{F0: 0, F1: ("" + ("hello" + v31))}
                                                                                        v32 = (-293 != 0)
                                                                                        if false {
                                                                                            var v83 uint8 = 0
                                                                                            var v84 uint16 = genFn1(0, 0)
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                        } else {
                                                                                            fmt.Println(v3, v31)
                                                                                        }
                                                                                        if v32 {
                                                                                            var v85 string = v31
                                                                                            if v32 {
                                                                                                v86 := GenS8{F0: false}
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = v5
                                                                                                }
                                                                                                var v87 int8 = v75
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                            }
                                                                                            var v88 float64 = (1.5 - (float64(v3) / (-2.25 / 1.0)))
                                                                                            _ = v85
                                                                                            _ = v88
                                                                                        }
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v82
                                                                                    }
                                                                                    _ = v79
                                                                                }
                                                                                _ = v77
                                                                            }
                                                                        }
                                                                        if v32 {
                                                                            v89 := GenS3{F0: -32768, F1: 65535}
                                                                            var v90 int32 = 30
                                                                            var v91 int16 = 1
                                                                            _ = v89
                                                                            _ = v90
                                                                            _ = v91
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(v75, v5)
                                                                                v75 = (v75 - v75)
                                                                                switch (^(fi37)) {
                                                                                case 32768:
                                                                                    v92 := GenS9{F0: 474}
                                                                                    _ = len(v1)
                                                                                    _ = v92
                                                                                case -1:
                                                                                    var v93 int64 = (v34 & v34)
                                                                                    _ = v93
                                                                                case 0:
                                                                                    v34 -= v34
                                                                                }
                                                                            } else {
                                                                                v5 = (^(v5))
                                                                                if v32 {
                                                                                    if ((true && (255 == 42)) && (v34 != (v34 / (((v34 & v34)) | 1)))) {
                                                                                        v94 := GenS10{F0: 4294967295}
                                                                                        _ = v94
                                                                                    } else {
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = 0
                                                                                        }
                                                                                        fmt.Println(v32)
                                                                                    }
                                                                                }
                                                                                switch (fi37 | v5) {
                                                                                case -864:
                                                                                    var v95 uint = 18446744073709551615
                                                                                    _ = v95
                                                                                }
                                                                            }
                                                                            var v96 uint16 = 167
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                    _ = v75
                                                                }
                                                            case 476:
                                                                var v97 uint16 = 350
                                                                if len(v2) > 0 {
                                                                    v2[0] = 238
                                                                }
                                                                var v98 float32 = 0.0
                                                                v99 := GenS11{F0: 2, F1: fi37, F2: 42}
                                                                _ = v97
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                            _ = fi72
                                                        }
                                                    }
                                                    var v100 int64 = v34
                                                    if len(v2) > 0 {
                                                        v2[0] = fi37
                                                    }
                                                    _ = v100
                                                }
                                            case -2147483648:
                                                v2 = append(v2, (fi37 >> ((432) & 63)))
                                            }
                                        }
                                        for fi101 := 0; fi101 < 20; fi101++ {
                                            if v32 {
                                                var v102 bool = ((int32(fi37) <= 42) && (!true))
                                                _ = v102
                                            }
                                            var v103 float64 = -2.25
                                            _ = fi101
                                            _ = v103
                                        }
                                    }
                                    if v32 {
                                        var v104 uint32 = 4294967295
                                        _ = v104
                                    } else {
                                        switch (v5 / ((2147483647) | 1)) {
                                        case -819:
                                            var v105 uint64 = 394
                                            _ = v105
                                        default:
                                            var v106 uint64 = 998
                                            _ = v106
                                        }
                                    }
                                    if v32 {
                                        if v32 {
                                            fmt.Println(v3, fi37)
                                            fmt.Println(fi37, v5, fi37)
                                            v5 = -261
                                        }
                                    } else {
                                        fmt.Println(fi37, fi37, v31)
                                        v107 := GenS12{F0: (127 < 191), F1: 0, F2: genFn0(255, "a", "a")}
                                        _ = v107
                                    }
                                }
                            }
                        }
                    }
                    for fi108 := 0; fi108 < 15; fi108++ {
                        _ = len(v1)
                        v32 = true
                        _ = fi108
                    }
                    v3 = 0.1
                } else {
                    var v109 bool = (1 >= 0)
                    v4 = append(v4, 2147483647)
                    var v110 int16 = 1000
                    _ = v109
                    _ = v110
                }
                if (((!v32) && (!v32)) && v32) {
                    var v111 uint8 = 39
                    if v32 {
                        v3 = float32(((-2.25 + (1e10 + ((3.14159 + (float64((v3 + 0.1)) * ((((1.5 + float64((v3 - v3))) / -2.25) - 0.0) - -2.25))) * 3.14159))) + 0.0))
                        for fi112 := 0; fi112 < 15; fi112++ {
                            v113 := GenS5{F0: 309}
                            v1 = append(v1, 4294967295)
                            _ = fi112
                            _ = v113
                        }
                        var v114 float64 = ((0.1 + (0.0 * -2.25)) + 0.1)
                        _ = v114
                    }
                    _ = v111
                } else {
                    var v115 uint32 = 4294967295
                    v34 = 9223372036854775807
                    var v116 uint8 = 69
                    _ = v115
                    _ = v116
                }
                _ = fi37
            }
        } else {
            _ = len(v4)
            if true {
                v117 := GenS13{F0: 2, F1: -1, F2: (!(v32 || v32))}
                var v118 int16 = 1
                _ = v117
                _ = v118
            }
            var v119 bool = false
            _ = v119
        }
    default:
        for fi120 := 0; fi120 < 1; fi120++ {
            var v121 uint32 = 177
            v31 = (v31 + "hello")
            for fi122 := 0; fi122 < 13; fi122++ {
                if ((v34 | v34) != (-(v34))) {
                    var v123 float64 = -2.25
                    switch "foo bar" {
                    case "foo bar":
                        var v124 uint = 255
                        _ = v124
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = v121
                        }
                    case "":
                        var v125 string = v31
                        _ = v125
                    }
                    switch v31 {
                    case "hello":
                        var v126 int8 = 50
                        _ = v126
                    case "a":
                        var v127 int = v5
                        _ = v127
                    case "foo bar":
                        fmt.Println(fi122, v32)
                    default:
                        v128 := GenS14{F0: (fi122 | 42)}
                        _ = v128
                    }
                    _ = len(v1)
                    if v32 {
                        if (v31 == (((v31 + ("a" + v31)) + "") + "foo bar")) {
                            var v129 int16 = 1
                            var v130 int8 = 127
                            if (-1 > 42) {
                                switch "" {
                                case "foo bar":
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v121, v32, v3)
                                    fmt.Println(fi120, v130)
                                }
                                var v131 int16 = v129
                                _ = v131
                            } else {
                                v5 += (v5 / (((fi120 + fi122)) | 1))
                                var v132 float32 = v3
                                _ = v132
                            }
                            _ = v129
                            _ = v130
                        }
                        v2 = append(v2, (v5 / (((fi120 - (fi120 * fi122))) | 1)))
                        if v32 {
                            var v133 uint32 = 95
                            var v134 float32 = v3
                            fmt.Println(v3, fi122, v133)
                            _ = v133
                            _ = v134
                        } else {
                            fmt.Println(fi120, v31)
                            var v135 int = v5
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v136 := GenS9{F0: 42}
                            v123 += v123
                            v137 := GenS6{F0: 42, F1: 0, F2: 160}
                            if (((v121 / ((1) | 1)) == 2) || true) {
                                var v138 uint32 = v121
                                if (fi122 > (v135 - v5)) {
                                    if v32 {
                                        if true {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            if (!v32) {
                                                v138 += (^(v121))
                                            }
                                            switch (v31 + v31) {
                                            case "\n\t":
                                                fmt.Println(fi122, v138, v34)
                                            case "foo bar":
                                                if ((v32 && false) || true) {
                                                    if len(v4) > 0 {
                                                        v4[0] = 216
                                                    }
                                                    v5 = (v135 - -2147483648)
                                                } else {
                                                    switch v31 {
                                                    case "a":
                                                        var v139 uint = 65535
                                                        _ = v139
                                                    case "\n\t":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case "hello":
                                                        fmt.Println(v123)
                                                    default:
                                                        fmt.Println(v3, v5)
                                                    }
                                                    var v140 uint8 = 2
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v141 string = v31
                                                    _ = v140
                                                    _ = v141
                                                }
                                            case "":
                                                if v32 {
                                                    if (("" > (((v31 + "foo bar") + (v31 + (v31 + ""))) + "\n\t")) && ((v34 >= v34) || ((v3 + 0.1) > float32(float64((((3.14159 / (3.14159 + (-2.25 * 1.5))) - v3) + float32((v123 + v123)))))))) {
                                                        var v142 int32 = -617
                                                        switch (((("foo bar" + "a") + (("foo bar" + (v31 + v31)) + "a")) + v31) + (v31 + "foo bar")) {
                                                        case "foo bar":
                                                            v143 := GenS9{F0: 991}
                                                            _ = v143
                                                        case "hello":
                                                            var v144 uint = 471
                                                            _ = v144
                                                        case "":
                                                            var v145 float32 = (0.1 + v3)
                                                            v145 = v145
                                                            var v146 int32 = -464
                                                            fmt.Println(v135, v34)
                                                            _ = v145
                                                            _ = v146
                                                        }
                                                        _ = len(v4)
                                                        v147 := GenS15{F0: (-(v142))}
                                                        _ = v142
                                                        _ = v147
                                                    }
                                                    if true {
                                                        v148 := GenS10{F0: (v138 >> ((14) & 63))}
                                                        v5 -= v5
                                                        _ = len(v2)
                                                        _ = v148
                                                    } else {
                                                        if v32 {
                                                            if len(v1) > 0 {
                                                                v1[0] = (^(v121))
                                                            }
                                                            fmt.Println(v121)
                                                        } else {
                                                            if (241 >= 15) {
                                                                _ = len(v2)
                                                                if true {
                                                                    var v149 int16 = -611
                                                                    if (317 != 213) {
                                                                        v2 = append(v2, (fi120 / ((fi120) | 1)))
                                                                    }
                                                                    _ = v149
                                                                } else {
                                                                    v121 = v138
                                                                }
                                                            } else {
                                                                fmt.Println(v3)
                                                                switch v31 {
                                                                case "\n\t":
                                                                    if v32 {
                                                                        var v150 int32 = -189
                                                                        _ = v150
                                                                    }
                                                                case "":
                                                                    if v32 {
                                                                        var v151 uint = 104
                                                                        v121 -= v121
                                                                        _ = v151
                                                                    }
                                                                default:
                                                                    v152 := GenS3{F0: 0, F1: v138}
                                                                    _ = v152
                                                                }
                                                            }
                                                            v153 := GenS16{F0: 1, F1: 185, F2: 0}
                                                            _ = v153
                                                        }
                                                        var v154 uint = 65535
                                                        if (uint64(v138) < 2) {
                                                            if false {
                                                                var v155 uint8 = 128
                                                                _ = v155
                                                            } else {
                                                                var v156 uint32 = v121
                                                                v157 := GenS17{F0: (false && v32)}
                                                                switch 32768 {
                                                                case 0:
                                                                    v121 -= v121
                                                                default:
                                                                    fmt.Println(v135, v31, v34)
                                                                }
                                                                _ = v156
                                                                _ = v157
                                                            }
                                                            switch v31 {
                                                            case "foo bar":
                                                                v138 -= (v138 - 42)
                                                                var v158 int16 = 42
                                                                _ = v158
                                                            default:
                                                                v159 := GenS18{F0: -920, F1: false, F2: (v154 / ((v154) | 1))}
                                                                _ = v159
                                                            }
                                                        }
                                                        _ = v154
                                                    }
                                                    var v160 uint16 = 0
                                                    if (!((v138 & v138) < 917)) {
                                                        var v161 uint32 = (v138 % ((v138) | 1))
                                                        _ = v161
                                                    }
                                                    _ = v160
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                        v162 := GenS3{F0: 191, F1: v138}
                                        fmt.Println(fi120)
                                        _ = v162
                                    } else {
                                        var v163 int16 = -32768
                                        switch fi120 {
                                        case 128:
                                            switch (fi122 + fi120) {
                                            case -721:
                                                v163 = (v163 % ((v163) | 1))
                                            case 776:
                                                v163 = (v163 % ((v163) | 1))
                                            case 9223372036854775807:
                                                v3 += ((v3 - v3) + 1.5)
                                            }
                                            if (1 == -1) {
                                                fmt.Println(fi120, fi122, v3)
                                            }
                                        default:
                                            switch fi120 {
                                            case 128:
                                                if (100 <= -1) {
                                                    if len(v2) > 0 {
                                                        v2[0] = v135
                                                    }
                                                } else {
                                                    if ((26 != (-(v34))) && v32) {
                                                        var v164 uint16 = 514
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v165 int64 = (^(v34))
                                                        switch fi122 {
                                                        case 32767:
                                                            var v166 float32 = float32(v123)
                                                            _ = v166
                                                        default:
                                                            v31 = v31
                                                        }
                                                        var v167 uint = 2
                                                        _ = v164
                                                        _ = v165
                                                        _ = v167
                                                    } else {
                                                        v168 := GenS3{F0: (v163 >> ((186) & 63)), F1: v121}
                                                        v34 -= (v34 & v34)
                                                        _ = v168
                                                    }
                                                }
                                            case -45:
                                                fmt.Println(v32)
                                            case 2147483647:
                                                if v32 {
                                                    v121 = (^(v138))
                                                    if (v32 || (v163 >= v163)) {
                                                        switch (("" + v31) + "a") {
                                                        case "hello":
                                                            var v169 uint8 = 103
                                                            _ = v169
                                                        }
                                                    } else {
                                                        v170 := GenS12{F0: (-1 >= -257), F1: 42, F2: 548}
                                                        if (4294967295 <= 65535) {
                                                            if v32 {
                                                                switch ((v31 + v31) + "\n\t") {
                                                                case "foo bar":
                                                                    v34 = 9223372036854775807
                                                                case "":
                                                                    _ = len(v2)
                                                                }
                                                                v171 := GenS14{F0: fi120}
                                                                var v172 float64 = (1.5 / v123)
                                                                _ = v171
                                                                _ = v172
                                                            } else {
                                                                _ = len(v2)
                                                            }
                                                        } else {
                                                            v173 := GenS19{F0: (v34 + (-(v34))), F1: (0 < 829)}
                                                            _ = v173
                                                        }
                                                        _ = v170
                                                    }
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    if (23 >= 0) {
                                                        if false {
                                                            v3 = float32(v123)
                                                            fmt.Println(v163)
                                                            var v174 uint16 = 349
                                                            _ = v174
                                                        }
                                                    } else {
                                                        v163 = 1
                                                        var v175 int32 = 1
                                                        v3 = v3
                                                        _ = v175
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = v138
                                                    }
                                                }
                                            }
                                        }
                                        _ = v163
                                    }
                                    if (v138 == 1) {
                                        v3 -= 3.14159
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if (((v34 >> ((255) & 63)) == (v34 >> ((616) & 63))) || v32) {
                                            var v176 int32 = int32(v34)
                                            if false {
                                                v121 += v138
                                                var v177 float32 = 3.14159
                                                var v178 int16 = 32767
                                                _ = v177
                                                _ = v178
                                            } else {
                                                var v179 uint16 = 1
                                                v135 = (fi122 & fi122)
                                                if false {
                                                    fmt.Println(v138, v135, v135)
                                                    var v180 uint16 = v179
                                                    switch (fi120 * fi122) {
                                                    case 1:
                                                        if v32 {
                                                            var v181 int32 = (v176 ^ (-(v176)))
                                                            var v182 uint64 = 405
                                                            _ = v181
                                                            _ = v182
                                                        }
                                                    }
                                                    _ = v180
                                                }
                                                _ = v179
                                            }
                                            v123 = v123
                                            _ = v176
                                        } else {
                                            v183 := GenS20{F0: (^(v121)), F1: 955}
                                            v3 -= ((v3 - v3) / v3)
                                            _ = v183
                                        }
                                        var v184 int32 = -315
                                        v185 := GenS21{F0: -654}
                                        _ = v184
                                        _ = v185
                                    }
                                    var v186 float32 = v3
                                    v187 := GenS5{F0: v138}
                                    var v188 int = fi122
                                    _ = v186
                                    _ = v187
                                    _ = v188
                                } else {
                                    fmt.Println(fi120, v135, v123)
                                    if ((v32 && ((!(2147483648 > (v138 + v121))) || (-105 > int8(v121)))) && ((v34 & (v34 % ((519) | 1))) == (-(v34)))) {
                                        v3 = 1e10
                                        v32 = (!(274 >= 4294967295))
                                    }
                                }
                                v4 = append(v4, (v121 & (v121 % ((v138) | 1))))
                                _ = v138
                            }
                            _ = v135
                            _ = v136
                            _ = v137
                        }
                    }
                    v189 := GenS22{F0: 773, F1: (^(fi120)), F2: 18446744073709551615}
                    _ = v123
                    _ = v189
                } else {
                    fmt.Println(v31)
                }
                var v190 int64 = -9223372036854775808
                fmt.Println(v121, v3, v3)
                _ = fi122
                _ = v190
            }
            _ = fi120
            _ = v121
        }
    }
    v3 -= (((0.0 - (((0.1 / v3) + v3) * (v3 / 1.5))) / float32(1.5)) - 3.14159)
    for fi191 := 0; fi191 < 6; fi191++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        switch v31 {
        case "\n\t":
            var v192 int8 = -7
            _ = v192
        }
        v193 := GenS10{F0: 65535}
        _ = fi191
        _ = v193
    }
    var v194 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v194:", v194)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v194
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
