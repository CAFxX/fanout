package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 string; F1 int; F2 uint8 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 int16; F1 int }
func genFn0(p0 uint16, p1 int32, p2 int16) uint16 {
    var v6 int8 = -1
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(v6, fi7, p2)
        p1 += (p1 << ((4294967295) & 63))
        p0 -= 32767
        _ = fi7
    }
    v8 := []int{0}
    for fi9 := 0; fi9 < 10; fi9++ {
        p1 = (p1 | p1)
        switch ((("foo bar" + "\n\t") + "hello") + "foo bar") {
        case "\n\t":
            switch ("hello" + "hello") {
            case "foo bar":
                var v10 string = ("foo bar" + "a")
                _ = v10
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    var v11 float32 = -2.25
                    _ = v11
                case "foo bar":
                    var v12 uint8 = 2
                    _ = v12
                case "a":
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 bool = false
                        _ = fi13
                        _ = v14
                    }
                }
            case "hello":
                if false {
                    switch ("" + "") {
                    case "a":
                        var v15 string = "foo bar"
                        _ = v15
                    case "":
                        for fi16 := 0; fi16 < 10; fi16++ {
                            v17 := GenS1{F0: 905, F1: 4294967295, F2: 59}
                            fmt.Println(fi16)
                            var v18 uint32 = 1
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    default:
                        if true {
                            for fi19 := 0; fi19 < 5; fi19++ {
                                _ = len(v8)
                                v20 := GenS2{F0: -1, F1: -982}
                                _ = len(v8)
                                if false {
                                    if len(v8) > 0 {
                                        v8[0] = 128
                                    }
                                } else {
                                    if false {
                                        if (1e10 >= (-2.25 / 1e10)) {
                                            var v21 int32 = 1
                                            _ = v21
                                        }
                                        if true {
                                            if false {
                                                var v22 int16 = int16(v6)
                                                v8 = append(v8, int(p0))
                                                fmt.Println(v6, v6, v22)
                                                _ = v22
                                            } else {
                                                if false {
                                                    var v23 int8 = -102
                                                    var v24 bool = (!true)
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                            fmt.Println(p1, p1, p1)
                                        }
                                        v6 = v6
                                    } else {
                                        if (fi19 == int(p1)) {
                                            var v25 uint32 = 2147483648
                                            if true {
                                                if false {
                                                    var v26 int64 = int64(v6)
                                                    _ = v26
                                                }
                                                if true {
                                                    if (39 == 255) {
                                                        var v27 int8 = (v6 * v6)
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        v28 := GenS3{F0: 2, F1: "\n\t"}
                                                        _ = v27
                                                        _ = v28
                                                    } else {
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        switch "foo bar" {
                                                        case "a":
                                                            switch (^(fi19)) {
                                                            case 32767:
                                                                var v29 uint64 = 4294967295
                                                                var v30 int32 = (p1 * (p1 - p1))
                                                                if ((fi9 + fi19) == (fi19 | fi19)) {
                                                                    if (("\n\t" + "a") == ("\n\t" + "a")) {
                                                                        var v31 uint16 = (p0 | p0)
                                                                        _ = v31
                                                                    } else {
                                                                        if true {
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "foo bar":
                                                                                if len(v8) >= 2 {
                                                                                    v8 = v8[1:len(v8)]
                                                                                }
                                                                            case "a":
                                                                                var v32 int8 = (v6 - int8(p0))
                                                                                _ = v32
                                                                            }
                                                                            v33 := GenS4{F0: uint8(v29), F1: 128}
                                                                            _ = v33
                                                                        }
                                                                        v8 = append(v8, (^(fi19)))
                                                                    }
                                                                    v34 := GenS2{F0: 578, F1: v30}
                                                                    _ = v34
                                                                } else {
                                                                    if false {
                                                                        if true {
                                                                            var v35 uint8 = 0
                                                                            _ = v35
                                                                        }
                                                                        fmt.Println(fi9, v25, v6)
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                    }
                                                                }
                                                                _ = v29
                                                                _ = v30
                                                            case -483:
                                                                p0 = (p0 << ((747) & 63))
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v6, v6, p1)
                                                        }
                                                        if (p2 >= p2) {
                                                            if (true || (!(!false))) {
                                                                if true {
                                                                    if true {
                                                                        v25 += v25
                                                                    }
                                                                    switch (fi19 - fi9) {
                                                                    case -826:
                                                                        _ = len(v8)
                                                                    case 32767:
                                                                        if false {
                                                                            p0 += p0
                                                                            var v36 uint = 260
                                                                            _ = v36
                                                                        } else {
                                                                            fmt.Println(fi19)
                                                                            if (int8(p0) >= v6) {
                                                                                var v37 bool = false
                                                                                if v37 {
                                                                                    var v38 uint16 = (p0 | p0)
                                                                                    switch (fi9 | fi19) {
                                                                                    case 32767:
                                                                                        switch fi9 {
                                                                                        case -2147483648:
                                                                                            var v39 uint16 = 65535
                                                                                            _ = v39
                                                                                        }
                                                                                    case -1:
                                                                                        if (!v37) {
                                                                                            var v40 uint64 = 1
                                                                                            _ = v40
                                                                                        }
                                                                                    default:
                                                                                        if len(v8) >= 2 {
                                                                                            v8 = v8[1:len(v8)]
                                                                                        }
                                                                                    }
                                                                                    _ = v38
                                                                                }
                                                                                _ = v37
                                                                            }
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                        }
                                                                    case 0:
                                                                        v41 := GenS1{F0: 1, F1: 42, F2: (v6 / ((v6) | 1))}
                                                                        _ = v41
                                                                    default:
                                                                        var v42 bool = (333 > -9223372036854775808)
                                                                        _ = v42
                                                                    }
                                                                }
                                                            } else {
                                                                p0 -= (p0 >> ((637) & 63))
                                                                if true {
                                                                    if false {
                                                                        var v43 int32 = p1
                                                                        _ = v43
                                                                    }
                                                                    p1 = (p1 % (((p1 & p1)) | 1))
                                                                }
                                                            }
                                                            switch 652 {
                                                            case 897:
                                                                var v44 float32 = ((((3.14159 - (1.5 / 3.14159)) - 1e10) + (0.1 / 1e10)) + 0.1)
                                                                _ = v44
                                                            }
                                                        }
                                                    }
                                                    var v45 uint32 = (v25 & v25)
                                                    if false {
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        p2 = p2
                                                        var v46 bool = (383 != p0)
                                                        _ = v46
                                                    }
                                                    switch 32768 {
                                                    case -2147483648:
                                                        v47 := GenS4{F0: 1, F1: 127}
                                                        _ = v47
                                                    case 42:
                                                        var v48 int64 = 1
                                                        if len(v8) > 0 {
                                                            v8[0] = int(v25)
                                                        }
                                                        _ = v48
                                                    case 32767:
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    }
                                                    _ = v45
                                                } else {
                                                    if false {
                                                        var v49 float32 = (0.0 / 1.5)
                                                        switch -472 {
                                                        case 42:
                                                            p2 -= p2
                                                        case 856:
                                                            var v50 string = "foo bar"
                                                            _ = v50
                                                        case 0:
                                                            if (100 == 45) {
                                                                var v51 float32 = (((v49 * v49) * 1.5) - v49)
                                                                var v52 float64 = -2.25
                                                                p2 -= (p2 + int16(fi19))
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                fmt.Println(p2)
                                                                p1 -= (^(p1))
                                                            }
                                                        default:
                                                            v6 = v6
                                                        }
                                                        if (-2.25 == 1.5) {
                                                            switch ("\n\t" + "\n\t") {
                                                            case "a":
                                                                v8 = append(v8, 128)
                                                            }
                                                        } else {
                                                            if (0 != 1) {
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                            } else {
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                            }
                                                            v8 = append(v8, fi19)
                                                            if (!((true || (!true)) || (-1 == -548))) {
                                                                v6 += (v6 / (((v6 << ((0) & 63))) | 1))
                                                                if len(v8) > 0 {
                                                                    v8[0] = fi9
                                                                }
                                                                var v53 string = ("" + "foo bar")
                                                                _ = v53
                                                            }
                                                        }
                                                        _ = v49
                                                    }
                                                }
                                                switch fi9 {
                                                case -2147483648:
                                                    v54 := GenS3{F0: 255, F1: "foo bar"}
                                                    _ = v54
                                                case 128:
                                                    p1 -= 1
                                                case 4294967295:
                                                    if (false || (true && true)) {
                                                        var v55 bool = ((-(v6)) <= v6)
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        _ = v55
                                                    }
                                                    fmt.Println(v25, fi19)
                                                    if false {
                                                        if false {
                                                            switch int(p0) {
                                                            case 9223372036854775807:
                                                                var v56 float64 = 0.1
                                                                _ = v56
                                                            case 127:
                                                                var v57 int = (fi9 & fi19)
                                                                _ = v57
                                                            default:
                                                                if false {
                                                                    p0 += p0
                                                                    v6 += v6
                                                                } else {
                                                                    v58 := GenS5{F0: "hello", F1: 0}
                                                                    _ = v58
                                                                }
                                                            }
                                                        } else {
                                                            v59 := GenS6{F0: 467}
                                                            p0 += p0
                                                            _ = v59
                                                        }
                                                    }
                                                    var v60 uint64 = 18446744073709551615
                                                    var v61 uint32 = (v25 ^ (v25 / (((v25 ^ 42)) | 1)))
                                                    if ((v6 ^ 18) < 42) {
                                                        if len(v8) > 0 {
                                                            v8[0] = fi9
                                                        }
                                                        v6 -= int8(v25)
                                                    } else {
                                                        var v62 int = 32768
                                                        if ((("" + "") + "\n\t") < "a") {
                                                            if true {
                                                                v62 -= fi19
                                                            }
                                                            if len(v8) > 0 {
                                                                v8[0] = (fi19 | (fi9 + fi9))
                                                            }
                                                            var v63 string = "a"
                                                            var v64 int8 = v6
                                                            fmt.Println(v62)
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            v8 = append(v8, (fi19 + (-(fi9))))
                                                            p0 = (p0 >> ((0) & 63))
                                                        }
                                                        v65 := GenS5{F0: ("\n\t" + "hello"), F1: 127}
                                                        if true {
                                                            p2 = -569
                                                            var v66 uint = 18446744073709551615
                                                            _ = v66
                                                        }
                                                        _ = v62
                                                        _ = v65
                                                    }
                                                    _ = v60
                                                    _ = v61
                                                default:
                                                    v67 := GenS7{F0: 450, F1: 4294967295}
                                                    _ = v67
                                                }
                                                if true {
                                                    var v68 int16 = (p2 | 42)
                                                    if true {
                                                        if true {
                                                            var v69 float64 = ((3.14159 - 1.5) / 3.14159)
                                                            _ = v69
                                                        }
                                                        fmt.Println(v68)
                                                        if len(v8) > 0 {
                                                            v8[0] = fi19
                                                        }
                                                        p2 = (v68 | (v68 | v68))
                                                        switch fi19 {
                                                        case -2147483648:
                                                            var v70 int32 = (p1 / ((p1) | 1))
                                                            _ = v70
                                                        }
                                                    }
                                                    var v71 string = ""
                                                    _ = v68
                                                    _ = v71
                                                } else {
                                                    var v72 int64 = 42
                                                    _ = v72
                                                }
                                            } else {
                                                var v73 string = ("hello" + (("hello" + ("foo bar" + "\n\t")) + "foo bar"))
                                                if false {
                                                    fmt.Println(v73)
                                                    fmt.Println(v6)
                                                    fmt.Println(v25, p0, p1)
                                                }
                                                var v74 uint64 = 369
                                                var v75 int16 = p2
                                                var v76 uint64 = (v74 / ((v74) | 1))
                                                _ = v73
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                            _ = v25
                                        } else {
                                            if false {
                                                _ = len(v8)
                                                if true {
                                                    switch (fi19 ^ (fi19 + fi9)) {
                                                    case 32767:
                                                        v77 := GenS5{F0: "foo bar", F1: 2}
                                                        _ = v77
                                                    case -2147483648:
                                                        v78 := GenS4{F0: 118, F1: 0}
                                                        _ = v78
                                                    default:
                                                        v8 = append(v8, (fi19 << ((uint(p0)) & 63)))
                                                    }
                                                    switch (("foo bar" + "a") + "") {
                                                    case "a":
                                                        v79 := GenS2{F0: -1, F1: p1}
                                                        var v80 int32 = p1
                                                        var v81 bool = false
                                                        _ = v79
                                                        _ = v80
                                                        _ = v81
                                                    case "hello":
                                                        var v82 int32 = 42
                                                        if ((!false) && true) {
                                                            v82 = v82
                                                        } else {
                                                            fmt.Println(p1, p2, fi9)
                                                            switch (fi9 >> ((445) & 63)) {
                                                            case 229:
                                                                var v83 int64 = 9223372036854775807
                                                                _ = v83
                                                            case 216:
                                                                var v84 uint = 42
                                                                _ = v84
                                                            case 0:
                                                                v6 = (v6 ^ (v6 - 42))
                                                            default:
                                                                v85 := GenS6{F0: 4294967295}
                                                                _ = v85
                                                            }
                                                        }
                                                        _ = v82
                                                    default:
                                                        p1 -= (p1 / ((p1) | 1))
                                                    }
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                }
                                            } else {
                                                v86 := GenS7{F0: p2, F1: 1}
                                                var v87 int32 = p1
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                        fmt.Println(p0, p1, v6)
                                    }
                                    p2 = (^(p2))
                                    v6 = 0
                                    p1 += p1
                                    v88 := GenS8{F0: (p0 ^ p0), F1: 32767}
                                    fmt.Println(fi9, fi9)
                                    _ = v88
                                }
                                _ = fi19
                                _ = v20
                            }
                            if len(v8) > 0 {
                                v8[0] = (fi9 / ((fi9) | 1))
                            }
                        }
                    }
                    p1 = p1
                } else {
                    switch (^(fi9)) {
                    case -226:
                        if ((!(!true)) && (-1 >= 260)) {
                            fmt.Println(v6, p2)
                            var v89 string = "\n\t"
                            _ = v89
                        }
                    case -73:
                        var v90 uint8 = 0
                        _ = v90
                    }
                }
            }
        case "foo bar":
            switch "hello" {
            case "a":
                if false {
                    v8 = append(v8, fi9)
                    for fi91 := 0; fi91 < 18; fi91++ {
                        v92 := GenS9{F0: "foo bar", F1: fi91, F2: 153}
                        if (!false) {
                            p0 -= 856
                            var v93 int64 = 0
                            var v94 uint16 = p0
                            var v95 int8 = v6
                            p2 += p2
                            if ((p1 != p1) || true) {
                                if len(v8) > 0 {
                                    v8[0] = (^(fi9))
                                }
                                if len(v8) > 0 {
                                    v8[0] = fi9
                                }
                            } else {
                                var v96 bool = false
                                _ = v96
                            }
                            _ = v93
                            _ = v94
                            _ = v95
                        }
                        _ = fi91
                        _ = v92
                    }
                } else {
                    v97 := GenS10{F0: 32768, F1: 1}
                    var v98 uint = 18446744073709551615
                    _ = v97
                    _ = v98
                }
            }
        }
        v6 = (v6 * int8(p2))
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return p0
}
func genFn1(p0 uint, p1 uint8, p2 int32) int16 {
    for fi99 := 0; fi99 < 13; fi99++ {
        p0 -= (p0 << (((p0 >> (((p0 % (((p0 + p0)) | 1))) & 63))) & 63))
        v100 := GenS11{F0: -432, F1: (fi99 & fi99)}
        _ = fi99
        _ = v100
    }
    var v101 bool = false
    fmt.Println(v101, v101, p2)
    p1 += (p1 / (((p1 / ((p1) | 1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v101
    return -1
}

func main() {
    v1 := []int{-1, -9223372036854775808, 0}
    var v2 int32 = 42
    var v3 float32 = ((0.1 * float32(0.1)) * (-2.25 - (-2.25 / 0.1)))
    var v4 float32 = -2.25
    var v5 uint32 = 65535
    var v102 uint16 = 1
    if ("" != "hello") {
        var v103 int = 941
        v3 -= -2.25
        _ = v103
    } else {
        fmt.Println(v3, v5, v3)
    }
    var v104 uint8 = 1
    v3 = 3.14159
    v5 = (^(v5))
    v3 += (0.1 / float32((1.5 - 1e10)))
    v104 -= (^(v104))
    var v105 string = "hello"
    v1 = append(v1, 32767)
    v102 -= v102
    if true {
        var v106 int = -187
        v106 -= 2147483647
        switch "\n\t" {
        case "foo bar":
            switch "\n\t" {
            case "hello":
                for fi107 := 0; fi107 < 10; fi107++ {
                    if ((fi107 << ((0) & 63)) > (v106 >> ((255) & 63))) {
                        if (0 <= 980) {
                            for fi108 := 0; fi108 < 14; fi108++ {
                                v106 = (v106 & int(v104))
                                if false {
                                    v1 = append(v1, v106)
                                    fmt.Println(v3, v102)
                                    var v109 uint32 = (v5 / (((v5 << ((uint(v5)) & 63))) | 1))
                                    _ = v109
                                }
                                v110 := GenS10{F0: fi107, F1: v5}
                                if len(v1) > 0 {
                                    v1[0] = v106
                                }
                                _ = fi108
                                _ = v110
                            }
                            fmt.Println(v102, v102, v106)
                        }
                        fmt.Println(v106, v104)
                        var v111 bool = ((v105 < (v105 + (v105 + v105))) || ((v5 | (v5 << ((167) & 63))) >= 2))
                        _ = v111
                    } else {
                        for fi112 := 0; fi112 < 17; fi112++ {
                            fmt.Println(v106, v105, v2)
                            v2 -= v2
                            _ = fi112
                        }
                        for fi113 := 0; fi113 < 6; fi113++ {
                            var v114 uint64 = 4294967295
                            var v115 string = v105
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                        var v116 bool = (((v5 ^ (v5 + (v5 % ((uint32(v104)) | 1)))) <= 731) || (4294967295 >= 42))
                        _ = v116
                    }
                    _ = fi107
                }
            case "a":
                var v117 bool = (((uint64(v102) > 2) || (((v102 - (v102 % ((genFn0((v102 + v102), 0, 0)) | 1))) != v102) || true)) && ((789 >= 1) && (361 == v102)))
                _ = v117
            default:
                v118 := GenS4{F0: (v104 / (((v104 << ((684) & 63))) | 1)), F1: v104}
                _ = v118
            }
        case "":
            v2 += (-(v2))
        case "a":
            v102 += genFn0(genFn0(v102, (v2 & v2), 32767), v2, genFn1(18446744073709551615, v104, v2))
        }
        _ = v106
    } else {
        v105 = (v105 + v105)
        v119 := GenS7{F0: 42, F1: 18446744073709551615}
        _ = v119
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v102:", v102)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v102
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
