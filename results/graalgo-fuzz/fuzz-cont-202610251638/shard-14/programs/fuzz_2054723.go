package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS6 struct { F0 string; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS11 struct { F0 uint64; F1 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 string; F1 int64 }
type GenS15 struct { F0 uint64 }
func genFn0() bool {
    v5 := GenS1{F0: "hello", F1: 42, F2: 127}
    if (2147483647 > 42) {
        v6 := GenS1{F0: "", F1: 128, F2: 89}
        var v7 int64 = -9223372036854775808
        var v8 float32 = -2.25
        _ = v6
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: "", F1: 127, F2: 100}
        var v10 int8 = 100
        _ = v10
        _ = v9
    }
    _ = v5
    return true
}
func genFn1(p0 int, p1 int32, p2 int16) uint16 {
    v11 := GenS1{F0: "foo bar", F1: 133, F2: -92}
    v12 := GenS1{F0: (("" + "a") + ("foo bar" + "hello")), F1: 2, F2: 100}
    p2 -= (p2 % ((1000) | 1))
    for fi13 := 0; fi13 < 11; fi13++ {
        if true {
            for fi14 := 0; fi14 < 11; fi14++ {
                if (42 >= 195) {
                    var v15 int64 = int64(fi14)
                    _ = v15
                } else {
                    var v16 float32 = -2.25
                    p0 += (fi14 + (p0 - (fi13 >> ((4294967295) & 63))))
                    v16 = (v16 - 0.0)
                    _ = v16
                }
                p0 = (-(fi14))
                _ = fi14
            }
        }
        switch (("\n\t" + "") + ("\n\t" + "a")) {
        case "foo bar":
            v17 := GenS1{F0: "foo bar", F1: 127, F2: -128}
            _ = v17
        case "a":
            v18 := GenS1{F0: ("foo bar" + "foo bar"), F1: 2, F2: 127}
            _ = v18
        }
        _ = fi13
    }
    switch 388 {
    case 128:
        v19 := GenS1{F0: "\n\t", F1: 0, F2: 127}
        _ = v19
    case 42:
        v20 := GenS2{F0: 314, F1: p0}
        _ = v20
    case 127:
        p1 = (p1 ^ (p1 | p1))
    default:
        if true {
            p1 -= int32(p0)
            var v21 uint = 42
            var v22 uint32 = 2
            _ = v21
            _ = v22
        } else {
            fmt.Println(p0)
            var v23 int32 = (p1 & p1)
            fmt.Println(v23)
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 int16 = 0
    v2 := []uint32{}
    var v3 bool = (-455 != 4294967295)
    var v4 uint8 = 0
    fmt.Println(v3, v1, v4)
    if false {
        switch ((("foo bar" + "a") + ("hello" + "\n\t")) + "a") {
        case "a":
            v3 = v3
        case "hello":
            if len(v2) > 0 {
                v2[0] = 431
            }
        case "foo bar":
            var v24 float64 = (((0.1 * float64(((0.0 - 1e10) * 0.0))) - (1e10 + -2.25)) / (float64(3.14159) + 1.5))
            _ = v24
        default:
            v4 = 127
        }
        v25 := GenS3{F0: 105, F1: "hello"}
        var v26 bool = (v4 >= v4)
        _ = v25
        _ = v26
    } else {
        fmt.Println(v1, v1, v4)
        fmt.Println(v3)
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := GenS2{F0: 0, F1: fi27}
        var v29 int64 = 1
        var v30 uint16 = 585
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v4)
    for fi31 := 0; fi31 < 16; fi31++ {
        v3 = (0 != -22)
        _ = fi31
    }
    switch 32767 {
    case -174:
        var v32 uint64 = 353
        _ = v32
    case -9223372036854775808:
        fmt.Println(v4, v4)
    case 391:
        v33 := GenS4{F0: -1, F1: 8}
        _ = v33
    }
    v34 := GenS4{F0: 500, F1: 65535}
    for fi35 := 0; fi35 < 6; fi35++ {
        switch ("\n\t" + (((("\n\t" + "foo bar") + "") + ((("a" + "\n\t") + ("" + "")) + ("\n\t" + "hello"))) + (("\n\t" + "foo bar") + ("\n\t" + "a")))) {
        case "foo bar":
            v1 = v1
        }
        switch "foo bar" {
        case "\n\t":
            v1 += v1
        case "":
            if (!(!(0.0 < (-2.25 * (3.14159 / 3.14159))))) {
                fmt.Println(v1)
            } else {
                v36 := GenS2{F0: 2, F1: (-(fi35))}
                _ = v36
            }
        default:
            switch "a" {
            case "foo bar":
                switch "a" {
                case "\n\t":
                    v37 := GenS2{F0: 2147483647, F1: -386}
                    _ = v37
                case "foo bar":
                    var v38 int = (-(fi35))
                    _ = v38
                default:
                    v39 := GenS4{F0: -958, F1: 1}
                    _ = v39
                }
            case "a":
                var v40 uint64 = 2
                _ = v40
            case "":
                var v41 string = (("foo bar" + "\n\t") + "\n\t")
                _ = v41
            default:
                var v42 int8 = 0
                _ = v42
            }
        }
        _ = fi35
    }
    v43 := GenS3{F0: -118, F1: ((("hello" + "a") + ((("a" + "") + "") + ("\n\t" + "hello"))) + "foo bar")}
    if (v4 <= (v4 - 42)) {
        for fi44 := 0; fi44 < 4; fi44++ {
            if (v3 && (!v3)) {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            } else {
                for fi45 := 0; fi45 < 17; fi45++ {
                    v1 -= (v1 * 42)
                    v46 := GenS2{F0: 199, F1: (fi44 & fi45)}
                    _ = fi45
                    _ = v46
                }
                v1 -= v1
            }
            _ = fi44
        }
    } else {
        if (v3 || (v3 && ((!(420 == 2147483648)) || ((65535 <= 42) || (0.0 == -2.25))))) {
            for fi47 := 0; fi47 < 1; fi47++ {
                _ = len(v2)
                switch fi47 {
                case 128:
                    fmt.Println(fi47)
                case -681:
                    v4 -= (v4 + 255)
                case 2147483647:
                    v1 = (v1 ^ v1)
                }
                _ = fi47
            }
            v48 := GenS1{F0: "hello", F1: (v4 - v4), F2: 33}
            _ = v48
        } else {
            switch (("" + "hello") + "foo bar") {
            case "hello":
                for fi49 := 0; fi49 < 18; fi49++ {
                    if ((-2.25 + float64((-2.25 + (0.0 / 1e10)))) > -2.25) {
                        switch ("" + ("foo bar" + "\n\t")) {
                        case "hello":
                            var v50 uint16 = 1
                            _ = v50
                        case "\n\t":
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v1)
                            case "a":
                                if (1 > -1) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v4 -= (^(v4))
                                }
                            default:
                                v51 := GenS5{F0: v4, F1: ((fi49 >> ((18446744073709551615) & 63)) > (-(fi49))), F2: 752}
                                _ = v51
                            }
                        default:
                            for fi52 := 0; fi52 < 3; fi52++ {
                                v1 -= (v1 & v1)
                                v1 -= v1
                                v1 += 1
                                _ = len(v2)
                                v1 += v1
                                _ = fi52
                            }
                        }
                        if (-128 <= int8(fi49)) {
                            v4 = (^(v4))
                            var v53 float64 = float64(1.5)
                            _ = v53
                        }
                    }
                    if true {
                        if false {
                            fmt.Println(fi49)
                            fmt.Println(v1, v1, fi49)
                            v1 += (v1 % ((1000) | 1))
                        }
                    }
                    _ = fi49
                }
            case "a":
                fmt.Println(v3)
            }
            _ = len(v2)
            for fi54 := 0; fi54 < 10; fi54++ {
                switch ("" + ("foo bar" + ("foo bar" + "hello"))) {
                case "a":
                    switch (fi54 >> ((4294967295) & 63)) {
                    case -1:
                        if len(v2) > 0 {
                            v2[0] = 978
                        }
                    default:
                        _ = len(v2)
                    }
                case "hello":
                    var v55 float32 = (1.5 - (1.5 / 1.5))
                    _ = v55
                case "":
                    switch "hello" {
                    case "":
                        var v56 bool = v3
                        _ = v56
                    case "\n\t":
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v57 := GenS6{F0: (("" + "hello") + "a"), F1: -1, F2: 42}
                            for fi58 := 0; fi58 < 13; fi58++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi58
                            }
                            _ = v57
                        } else {
                            v59 := GenS7{F0: 127, F1: 854}
                            v1 += v1
                            if len(v2) > 0 {
                                v2[0] = 780
                            }
                            _ = v59
                        }
                    case "hello":
                        v60 := GenS2{F0: 2147483647, F1: 156}
                        _ = v60
                    }
                }
                _ = fi54
            }
        }
        if v3 {
            if false {
                v4 = (^(v4))
                v61 := GenS1{F0: "", F1: (^(v4)), F2: -128}
                _ = v61
            }
            if v3 {
                var v62 uint32 = 2
                if v3 {
                    fmt.Println(v1)
                }
                for fi63 := 0; fi63 < 1; fi63++ {
                    v64 := GenS2{F0: 430, F1: fi63}
                    if v3 {
                        switch "\n\t" {
                        case "\n\t":
                            v1 += v1
                        case "hello":
                            switch (fi63 * fi63) {
                            case 975:
                                var v65 float32 = (0.1 * 1e10)
                                _ = v65
                            case -282:
                                if v3 {
                                    var v66 uint8 = v4
                                    var v67 string = ""
                                    v68 := GenS8{F0: 1}
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = (^(v62))
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v69 := GenS9{F0: "hello", F1: 65535}
                                    _ = v69
                                }
                            }
                        case "":
                            var v70 uint16 = 110
                            _ = v70
                        default:
                            v3 = v3
                        }
                        var v71 float64 = (1e10 / 1e10)
                        v72 := GenS9{F0: ("\n\t" + ("a" + "")), F1: 65535}
                        _ = v71
                        _ = v72
                    }
                    _ = fi63
                    _ = v64
                }
                _ = v62
            } else {
                if true {
                    fmt.Println(v3, v1, v3)
                    if v3 {
                        var v73 int8 = 0
                        switch (("" + "hello") + ("a" + "\n\t")) {
                        case "foo bar":
                            v4 += v4
                        }
                        _ = v73
                    }
                    var v74 string = "foo bar"
                    _ = v74
                } else {
                    if len(v2) > 0 {
                        v2[0] = 637
                    }
                    v3 = true
                }
                v2 = append(v2, 2147483648)
                var v75 uint32 = 1
                _ = v75
            }
            v76 := GenS8{F0: 255}
            _ = v76
        } else {
            v77 := GenS10{F0: 42, F1: 42, F2: 4294967295}
            _ = v77
        }
    }
    var v78 int = 1
    v1 += (v1 >> ((42) & 63))
    var v79 int32 = 42
    v80 := GenS11{F0: 4294967295, F1: 0}
    if (((v1 & (v1 / ((v1) | 1))) != v1) && (786 != 1)) {
        if v3 {
            for fi81 := 0; fi81 < 11; fi81++ {
                v78 += (v78 | -2147483648)
                var v82 uint64 = 236
                _ = fi81
                _ = v82
            }
            v83 := GenS12{F0: v78}
            _ = v83
        }
        var v84 float32 = -2.25
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v84
    } else {
        v2 = append(v2, 2147483648)
        for fi85 := 0; fi85 < 11; fi85++ {
            v79 -= (v79 % (((-(v79))) | 1))
            if (820 >= 0) {
                for fi86 := 0; fi86 < 15; fi86++ {
                    switch (fi86 + (fi86 & (-(fi85)))) {
                    case -1:
                        v2 = append(v2, 2)
                    case -68:
                        if (v3 && (586 == 284)) {
                            switch "\n\t" {
                            case "":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        } else {
                            if true {
                                var v87 float64 = (3.14159 / 1.5)
                                v79 = (v79 - v79)
                                v79 -= (v79 ^ v79)
                                v88 := GenS13{F0: (v1 & v1)}
                                _ = v87
                                _ = v88
                            } else {
                                v4 = (^(v4))
                                if v3 {
                                    v89 := GenS7{F0: 0, F1: 400}
                                    _ = v89
                                } else {
                                    var v90 uint16 = 42
                                    switch fi86 {
                                    case 32768:
                                        if (genFn1((v78 ^ fi86), v79, v1) > 447) {
                                            v1 += (v1 - (^(v1)))
                                            if (0 < 1) {
                                                fmt.Println(v4, fi85, v3)
                                            }
                                        }
                                    }
                                    _ = v90
                                }
                                switch "\n\t" {
                                case "":
                                    fmt.Println(fi86)
                                case "\n\t":
                                    var v91 uint8 = v4
                                    _ = v91
                                case "hello":
                                    v92 := GenS4{F0: 9223372036854775807, F1: 2}
                                    _ = v92
                                }
                                if (100 < 127) {
                                    var v93 int64 = 42
                                    if len(v2) > 0 {
                                        v2[0] = 624
                                    }
                                    _ = v93
                                }
                            }
                        }
                    }
                    if (!v3) {
                        v79 -= (v79 - v79)
                        v3 = ("foo bar" > ("a" + "hello"))
                        switch "\n\t" {
                        case "":
                            fmt.Println(v4, v3)
                        case "hello":
                            v2 = append(v2, 141)
                            var v94 float64 = 1.5
                            _ = v94
                        case "a":
                            var v95 uint32 = 0
                            _ = v95
                        default:
                            var v96 int32 = 42
                            _ = v96
                        }
                    }
                    _ = fi86
                }
                if true {
                    var v97 int64 = -501
                    var v98 bool = ((((!v3) || ((^(v79)) <= 1)) || v3) && (!(42 >= -9223372036854775808)))
                    _ = v97
                    _ = v98
                } else {
                    for fi99 := 0; fi99 < 8; fi99++ {
                        v100 := GenS14{F0: ((("foo bar" + "a") + "\n\t") + ((("foo bar" + "hello") + "\n\t") + ((("a" + "a") + "\n\t") + "foo bar"))), F1: -44}
                        if (!v3) {
                            v101 := GenS5{F0: (^(v4)), F1: true, F2: genFn1(fi85, 1, (v1 % (((v1 / ((v1) | 1))) | 1)))}
                            v1 = v1
                            fmt.Println(v78, fi99)
                            _ = v101
                        } else {
                            switch v78 {
                            case 120:
                                v4 -= (v4 * (v4 ^ 42))
                            case 1:
                                fmt.Println(v1, v4)
                            }
                        }
                        _ = fi99
                        _ = v100
                    }
                    var v102 int64 = 0
                    v1 = (-(v1))
                    _ = v102
                }
            }
            for fi103 := 0; fi103 < 12; fi103++ {
                v104 := GenS15{F0: 539}
                _ = fi103
                _ = v104
            }
            _ = fi85
        }
        var v105 int32 = -1
        _ = v105
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v78
    _ = v79
    _ = v80
}
