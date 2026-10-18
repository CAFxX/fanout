package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 string; F2 string }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int8; F2 uint16 }
type GenS9 struct { F0 uint16; F1 int64; F2 int }
type GenS10 struct { F0 uint64; F1 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int64; F1 int32; F2 int16 }
type GenS13 struct { F0 uint16; F1 uint }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int8; F1 uint }
type GenS16 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 int8; F1 int }
type GenS20 struct { F0 int16; F1 uint16 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 int32 }
type GenS23 struct { F0 string; F1 uint64; F2 uint32 }
type GenS24 struct { F0 uint32; F1 bool }
type GenS25 struct { F0 bool; F1 int32 }
type GenS26 struct { F0 string; F1 uint32 }
type GenS27 struct { F0 uint }
type GenS28 struct { F0 uint64 }
type GenS29 struct { F0 uint8; F1 string; F2 uint16 }
type GenS30 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS31 struct { F0 uint32; F1 uint64; F2 uint }
type GenS32 struct { F0 uint16; F1 int32 }
type GenS33 struct { F0 uint8; F1 uint }
type GenS34 struct { F0 string; F1 int32; F2 uint }
type GenS35 struct { F0 uint16; F1 int; F2 uint }
type GenS36 struct { F0 int64; F1 int; F2 int }
type GenS37 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS38 struct { F0 uint8 }
type GenS39 struct { F0 int16; F1 int16 }
type GenS40 struct { F0 uint32 }
type GenS41 struct { F0 uint32 }
type GenS42 struct { F0 int64; F1 uint64; F2 uint }
type GenS43 struct { F0 bool }
type GenS44 struct { F0 int16 }
type GenS45 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS46 struct { F0 uint8; F1 string; F2 int32 }
type GenS47 struct { F0 bool; F1 uint16; F2 int64 }
type GenS48 struct { F0 string; F1 int32; F2 int64 }
type GenS49 struct { F0 int16; F1 int32 }
type GenS50 struct { F0 int64; F1 uint; F2 int16 }
type GenS51 struct { F0 int; F1 int16 }
type GenS52 struct { F0 uint64; F1 int64 }
func genFn0() bool {
    v5 := []int{-2147483648, 9223372036854775807}
    if false {
        var v6 bool = false
        v5 = append(v5, 2147483647)
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 string = ("hello" + (("foo bar" + ("foo bar" + (("hello" + ("foo bar" + (("foo bar" + "") + "a"))) + ""))) + "foo bar"))
            var v9 uint16 = 255
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            _ = len(v5)
            fmt.Println(fi10)
            var v11 int8 = -1
            _ = fi10
            _ = v11
        }
        if (18446744073709551615 > 4294967295) {
            v12 := GenS1{F0: -2147483648}
            var v13 uint8 = 42
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 int16 = 1
                _ = fi14
                _ = v15
            }
            _ = v12
            _ = v13
        }
        switch "foo bar" {
        case "hello":
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 float32 = ((float32(1e10) * (((3.14159 / -2.25) / 1e10) / float32(3.14159))) - -2.25)
                fmt.Println(v17, fi16)
                var v18 bool = false
                _ = fi16
                _ = v17
                _ = v18
            }
        case "a":
            for fi19 := 0; fi19 < 12; fi19++ {
                fmt.Println(fi19)
                v20 := GenS1{F0: 2147483647}
                _ = fi19
                _ = v20
            }
        case "foo bar":
            v21 := GenS2{F0: (false || (42 <= 1)), F1: false, F2: ("a" + "foo bar")}
            _ = v21
        }
    }
    var v22 int64 = 0
    if true {
        if true {
            if (1e10 == (1e10 + 1.5)) {
                v23 := GenS1{F0: -2147483648}
                if len(v5) > 0 {
                    v5[0] = -9223372036854775808
                }
                _ = v23
            } else {
                var v24 float32 = 3.14159
                if (!true) {
                    v22 = (v22 % (((v22 + 0)) | 1))
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v25 := GenS1{F0: 0}
                        _ = v25
                    }
                    v24 += (3.14159 / (((0.1 - (v24 - (v24 * v24))) - v24) + v24))
                } else {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 float64 = (0.1 - float64((v24 + 0.0)))
                        _ = fi26
                        _ = v27
                    }
                    if false {
                        var v28 uint = 2
                        _ = v28
                    }
                }
                _ = v24
            }
        } else {
            if false {
                v29 := GenS2{F0: true, F1: (true && false), F2: "a"}
                for fi30 := 0; fi30 < 9; fi30++ {
                    switch (^(fi30)) {
                    case 635:
                        var v31 string = "foo bar"
                        _ = v31
                    default:
                        fmt.Println(fi30)
                    }
                    _ = fi30
                }
                _ = v29
            }
        }
        v22 += 827
        var v32 int16 = 42
        _ = v32
    }
    switch "hello" {
    case "hello":
        fmt.Println(v22)
    }
    _ = v22
    _ = v5
    return ((!false) && false)
}
func genFn1(p0 uint8, p1 uint8, p2 uint8) string {
    var v33 int16 = -20
    fmt.Println(p2)
    p2 += (p2 ^ (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return ("\n\t" + "foo bar")
}

func main() {
    var v1 float64 = 0.1
    var v2 string = "hello"
    var v3 uint64 = 37
    var v4 int8 = 42
    if (!("\n\t" > (v2 + ""))) {
        v34 := []int{}
        _ = v34
    }
    switch v2 {
    case "a":
        v2 = v2
    }
    v3 = (^(v3))
    var v35 uint8 = 1
    v3 += (v3 / ((1) | 1))
    v36 := []int{38, -1, -9223372036854775808}
    var v37 uint32 = 42
    var v38 uint16 = 42
    var v39 uint = 2
    fmt.Println(v37, v4, v37)
    for fi40 := 0; fi40 < 6; fi40++ {
        v39 = (v39 % ((65535) | 1))
        _ = fi40
    }
    fmt.Println(v2)
    if false {
        for fi41 := 0; fi41 < 5; fi41++ {
            var v42 uint = v39
            if true {
                fmt.Println(v2)
                var v43 int64 = int64(fi41)
                for fi44 := 0; fi44 < 15; fi44++ {
                    v42 = 255
                    var v45 int64 = int64(v38)
                    _ = fi44
                    _ = v45
                }
                _ = v43
            } else {
                if true {
                    if (v2 <= v2) {
                        v39 += (v42 >> (((^(v39))) & 63))
                        fmt.Println(fi41)
                        v46 := GenS1{F0: 0}
                        _ = v46
                    }
                    fmt.Println(v2)
                    v3 = 649
                } else {
                    switch "a" {
                    case "a":
                        switch (fi41 - (fi41 - fi41)) {
                        case 0:
                            fmt.Println(v39)
                        default:
                            v47 := GenS2{F0: true, F1: false, F2: (v2 + "")}
                            _ = v47
                        }
                    case "":
                        if (v1 < 1e10) {
                            if (114 == v4) {
                                v42 -= v42
                                if true {
                                    if (true && (447 < 32767)) {
                                        for fi48 := 0; fi48 < 18; fi48++ {
                                            var v49 int16 = -541
                                            _ = fi48
                                            _ = v49
                                        }
                                    }
                                    if len(v36) >= 2 {
                                        v36 = v36[1:len(v36)]
                                    }
                                    switch v2 {
                                    case "foo bar":
                                        if true {
                                            switch 618 {
                                            case 42:
                                                _ = len(v36)
                                            case 1:
                                                if len(v36) >= 2 {
                                                    v36 = v36[1:len(v36)]
                                                }
                                            case 0:
                                                for fi50 := 0; fi50 < 19; fi50++ {
                                                    if len(v36) > 0 {
                                                        v36[0] = 32768
                                                    }
                                                    var v51 int32 = -2147483648
                                                    fmt.Println(v38, v35)
                                                    switch v2 {
                                                    case "\n\t":
                                                        v52 := GenS1{F0: int32(fi50)}
                                                        _ = v52
                                                    case "a":
                                                        _ = len(v36)
                                                    case "foo bar":
                                                        v53 := GenS3{F0: (^(v39)), F1: v2, F2: (((v2 + v2) + v2) + (v2 + v2))}
                                                        _ = v53
                                                    }
                                                    v54 := GenS4{F0: (v38 & 65535), F1: v37, F2: (v37 - 4294967295)}
                                                    v55 := GenS4{F0: 65535, F1: 0, F2: v37}
                                                    _ = fi50
                                                    _ = v51
                                                    _ = v54
                                                    _ = v55
                                                }
                                            }
                                        }
                                    case "hello":
                                        if len(v36) >= 2 {
                                            v36 = v36[1:len(v36)]
                                        }
                                    case "\n\t":
                                        v56 := GenS4{F0: v38, F1: 1, F2: (v37 | uint32(v39))}
                                        _ = v56
                                    default:
                                        v36 = append(v36, (fi41 | (fi41 + (fi41 - fi41))))
                                    }
                                } else {
                                    var v57 bool = false
                                    _ = v57
                                }
                                if false {
                                    if true {
                                        if (!(1.5 > (float64(float32(float64(-2.25))) - 0.1))) {
                                            var v58 float64 = 1.5
                                            for fi59 := 0; fi59 < 13; fi59++ {
                                                fmt.Println(v4, v4)
                                                v60 := GenS3{F0: v42, F1: "hello", F2: v2}
                                                if (fi41 >= 127) {
                                                    var v61 int8 = v4
                                                    switch ((("foo bar" + (((v2 + "") + (("a" + v2) + "")) + (((("" + "") + "a") + v2) + ("hello" + "hello")))) + v2) + (v2 + "")) {
                                                    case "hello":
                                                        v62 := GenS3{F0: v42, F1: ((((v2 + v2) + v2) + "hello") + ""), F2: ((v2 + v2) + v2)}
                                                        _ = v62
                                                    case "\n\t":
                                                        var v63 bool = (944 > 158)
                                                        _ = v63
                                                    case "a":
                                                        fmt.Println(v39, v3)
                                                    default:
                                                        v35 += v35
                                                        switch "hello" {
                                                        case "a":
                                                            fmt.Println(v3)
                                                        case "hello":
                                                            v36 = append(v36, 0)
                                                        case "\n\t":
                                                            var v64 uint16 = (v38 - v38)
                                                            _ = v64
                                                        default:
                                                            var v65 int16 = 679
                                                            var v66 uint = (v39 / (((^(v42))) | 1))
                                                            var v67 float32 = ((1e10 / -2.25) / 1.5)
                                                            v61 += int8(v37)
                                                            if ((v38 | uint16(v42)) < (v38 / (((v38 >> (((v42 * 234)) & 63))) | 1))) {
                                                                v68 := GenS2{F0: false, F1: false, F2: (v2 + v2)}
                                                                if (!false) {
                                                                    v69 := GenS5{F0: v3}
                                                                    v35 += 163
                                                                    _ = v69
                                                                }
                                                                fmt.Println(v58)
                                                                if false {
                                                                    if len(v36) > 0 {
                                                                        v36[0] = -191
                                                                    }
                                                                    fmt.Println(v3)
                                                                    _ = len(v36)
                                                                    if false {
                                                                        switch (v2 + v2) {
                                                                        case "\n\t":
                                                                            if false {
                                                                                v1 += float64((v67 + (float32(0.0) / 3.14159)))
                                                                                fmt.Println(v65)
                                                                                var v70 float64 = -2.25
                                                                                _ = len(v36)
                                                                                _ = v70
                                                                            }
                                                                        case "hello":
                                                                            var v71 uint16 = (^(v38))
                                                                            _ = v71
                                                                        default:
                                                                            if (v35 <= (v35 & 2)) {
                                                                                fmt.Println(v38, v67, v1)
                                                                                if len(v36) > 0 {
                                                                                    v36[0] = (fi41 * (fi59 << ((703) & 63)))
                                                                                }
                                                                                var v72 int64 = 42
                                                                                _ = v72
                                                                            } else {
                                                                                v61 += -23
                                                                                v73 := GenS6{F0: "foo bar", F1: 1}
                                                                                switch v2 {
                                                                                case "hello":
                                                                                    _ = len(v36)
                                                                                case "foo bar":
                                                                                    fmt.Println(v66)
                                                                                case "a":
                                                                                    if (v65 > (v65 << ((4294967295) & 63))) {
                                                                                        switch 738 {
                                                                                        case -2147483648:
                                                                                            v1 -= v1
                                                                                        case 0:
                                                                                            v65 = v65
                                                                                        case 790:
                                                                                            v74 := GenS6{F0: "foo bar", F1: (v35 & (v35 ^ 144))}
                                                                                            _ = v74
                                                                                        }
                                                                                        _ = len(v36)
                                                                                    }
                                                                                }
                                                                                _ = v73
                                                                            }
                                                                            var v75 uint = 4294967295
                                                                            _ = v75
                                                                        }
                                                                        var v76 uint16 = 2
                                                                        switch ("foo bar" + v2) {
                                                                        case "hello":
                                                                            fmt.Println(v39, v61, v65)
                                                                        case "\n\t":
                                                                            v61 += v4
                                                                        }
                                                                        switch ((v2 + (v2 + "a")) + "foo bar") {
                                                                        case "a":
                                                                            v61 -= 127
                                                                        case "\n\t":
                                                                            v36 = append(v36, fi41)
                                                                        case "foo bar":
                                                                            fmt.Println(v58, v66, v61)
                                                                        default:
                                                                            if (42 == -1) {
                                                                                v35 = (v35 % (((v35 * v35)) | 1))
                                                                            } else {
                                                                                if len(v36) > 0 {
                                                                                    v36[0] = fi41
                                                                                }
                                                                                if (353 >= (v3 - v3)) {
                                                                                    if true {
                                                                                        _ = len(v36)
                                                                                    } else {
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            v1 += 1e10
                                                                                        case "a":
                                                                                            switch fi59 {
                                                                                            case 128:
                                                                                                v66 += 255
                                                                                            case 385:
                                                                                                v77 := GenS7{F0: 42}
                                                                                                _ = v77
                                                                                            case 1:
                                                                                                var v78 int32 = 1
                                                                                                _ = v78
                                                                                            default:
                                                                                                v36 = append(v36, (fi41 / ((-291) | 1)))
                                                                                            }
                                                                                        default:
                                                                                            v36 = append(v36, -9223372036854775808)
                                                                                        }
                                                                                        if false {
                                                                                            if (-128 >= int8(v66)) {
                                                                                                v1 = (-2.25 / v1)
                                                                                                var v79 float64 = float64(float32((1.5 - ((v1 / v1) * 1e10))))
                                                                                                _ = v79
                                                                                            }
                                                                                            v66 = 860
                                                                                            var v80 string = v2
                                                                                            _ = v80
                                                                                        } else {
                                                                                            var v81 uint16 = (v76 >> ((65535) & 63))
                                                                                            _ = v81
                                                                                        }
                                                                                    }
                                                                                    switch v2 {
                                                                                    case "foo bar":
                                                                                        v36 = append(v36, (fi41 << ((uint(v35)) & 63)))
                                                                                    case "":
                                                                                        v82 := GenS6{F0: (v2 + v2), F1: (v35 ^ 222)}
                                                                                        _ = v82
                                                                                    default:
                                                                                        v37 -= (v37 & v37)
                                                                                    }
                                                                                    v35 -= uint8(v42)
                                                                                } else {
                                                                                    var v83 int = (fi59 % (((fi59 | fi41)) | 1))
                                                                                    switch v83 {
                                                                                    case -9223372036854775808:
                                                                                        var v84 uint8 = (v35 / ((v35) | 1))
                                                                                        _ = v84
                                                                                    case 128:
                                                                                        if true {
                                                                                            v67 = v67
                                                                                        }
                                                                                    default:
                                                                                        _ = len(v36)
                                                                                    }
                                                                                    if len(v36) > 0 {
                                                                                        v36[0] = int(v38)
                                                                                    }
                                                                                    if false {
                                                                                        var v85 uint8 = 141
                                                                                        var v86 float64 = v1
                                                                                        var v87 uint32 = (^(v37))
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                        _ = v87
                                                                                    }
                                                                                    _ = v83
                                                                                }
                                                                                var v88 int16 = v65
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                        _ = v76
                                                                    } else {
                                                                        switch (v2 + ("" + v2)) {
                                                                        case "\n\t":
                                                                            var v89 int = 9223372036854775807
                                                                            _ = v89
                                                                        default:
                                                                            switch fi59 {
                                                                            case 127:
                                                                                if true {
                                                                                    var v90 float32 = 0.0
                                                                                    switch -782 {
                                                                                    case -349:
                                                                                        v91 := GenS1{F0: 2147483647}
                                                                                        _ = v91
                                                                                    case -1:
                                                                                        var v92 float32 = 3.14159
                                                                                        _ = v92
                                                                                    default:
                                                                                        v39 = 18446744073709551615
                                                                                    }
                                                                                    v93 := GenS7{F0: 42}
                                                                                    if false {
                                                                                        v65 += 1
                                                                                    }
                                                                                    _ = v90
                                                                                    _ = v93
                                                                                }
                                                                            case 107:
                                                                                var v94 uint8 = 30
                                                                                _ = v94
                                                                            case 367:
                                                                                var v95 int16 = v65
                                                                                _ = v95
                                                                            }
                                                                        }
                                                                        v96 := GenS8{F0: -2147483648, F1: (v4 << ((1) & 63)), F2: v38}
                                                                        _ = v96
                                                                    }
                                                                    var v97 int8 = v61
                                                                    v66 += (v42 + (^(v42)))
                                                                    var v98 int8 = v97
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                                _ = v68
                                                            }
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    }
                                                    var v99 uint64 = (v3 / ((4294967295) | 1))
                                                    _ = v61
                                                    _ = v99
                                                }
                                                _ = fi59
                                                _ = v60
                                            }
                                            _ = v58
                                        } else {
                                            var v100 uint16 = uint16(v37)
                                            switch ("" + "hello") {
                                            case "a":
                                                var v101 uint64 = v3
                                                _ = v101
                                            case "":
                                                var v102 int32 = 318
                                                _ = v102
                                            case "\n\t":
                                                v35 = (v35 & v35)
                                            default:
                                                var v103 int16 = 1
                                                _ = v103
                                            }
                                            var v104 uint = 1
                                            _ = v100
                                            _ = v104
                                        }
                                    }
                                }
                            }
                            switch -675 {
                            case 32768:
                                if (!(true && false)) {
                                    if (!false) {
                                        v105 := GenS3{F0: (v42 & (v39 & (v39 << (((v42 ^ (v39 % ((v42) | 1)))) & 63)))), F1: (v2 + "\n\t"), F2: ""}
                                        _ = v105
                                    }
                                    v106 := GenS9{F0: 706, F1: 571, F2: fi41}
                                    if true {
                                        var v107 int64 = -9223372036854775808
                                        var v108 uint64 = (v3 % (((v3 % ((499) | 1))) | 1))
                                        _ = v107
                                        _ = v108
                                    } else {
                                        if false {
                                            v109 := GenS4{F0: (v38 % (((^(v38))) | 1)), F1: (v37 >> (((v42 ^ v39)) & 63)), F2: (v37 >> (((v39 - v39)) & 63))}
                                            for fi110 := 0; fi110 < 1; fi110++ {
                                                v111 := GenS10{F0: v3, F1: 0}
                                                if (!false) {
                                                    v112 := GenS11{F0: 42}
                                                    v1 -= 1.5
                                                    if false {
                                                        if len(v36) > 0 {
                                                            v36[0] = 55
                                                        }
                                                        fmt.Println(v2)
                                                        if true {
                                                            v113 := GenS12{F0: 1, F1: 716, F2: -1}
                                                            v114 := GenS13{F0: v38, F1: v39}
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                        fmt.Println(fi110)
                                                    } else {
                                                        v2 = v2
                                                        v3 -= (v3 ^ v3)
                                                        var v115 uint64 = (v3 - (v3 / (((v3 - (^(v3)))) | 1)))
                                                        if ((v38 % ((1) | 1)) <= 2) {
                                                            if false {
                                                                var v116 float32 = 1e10
                                                                fmt.Println(v4)
                                                                _ = v116
                                                            }
                                                            v1 = v1
                                                        }
                                                        _ = v115
                                                    }
                                                    _ = v112
                                                } else {
                                                    v38 += uint16(v37)
                                                    if false {
                                                        if false {
                                                            var v117 uint = uint(v35)
                                                            _ = v117
                                                        } else {
                                                            var v118 int8 = (^(v4))
                                                            v36 = append(v36, (fi41 & (fi41 & (fi110 + fi41))))
                                                            _ = v118
                                                        }
                                                        if true {
                                                            if (-9223372036854775808 < 0) {
                                                                v119 := GenS14{F0: v38}
                                                                _ = v119
                                                            } else {
                                                                if (536 < 649) {
                                                                    v36 = append(v36, (fi41 | (-(fi41))))
                                                                    v36 = append(v36, 1)
                                                                    var v120 int = (^(fi41))
                                                                    var v121 int16 = 416
                                                                    _ = v120
                                                                    _ = v121
                                                                } else {
                                                                    var v122 bool = false
                                                                    if v122 {
                                                                        fmt.Println(v42, v3)
                                                                        var v123 uint = (^(v42))
                                                                        v2 = ""
                                                                        _ = v123
                                                                    } else {
                                                                        v36 = append(v36, fi41)
                                                                    }
                                                                    v37 = 1
                                                                    _ = v122
                                                                }
                                                                switch 2147483647 {
                                                                case 0:
                                                                    v3 = (v3 >> ((0) & 63))
                                                                case -2147483648:
                                                                    var v124 uint32 = v37
                                                                    _ = v124
                                                                }
                                                                var v125 int32 = -2147483648
                                                                _ = v125
                                                            }
                                                            var v126 uint32 = (v37 / (((v37 * v37)) | 1))
                                                            v38 += 42
                                                            _ = v126
                                                        }
                                                    } else {
                                                        fmt.Println(v37)
                                                    }
                                                }
                                                _ = fi110
                                                _ = v111
                                            }
                                            switch v2 {
                                            case "foo bar":
                                                for fi127 := 0; fi127 < 19; fi127++ {
                                                    v1 = float64(0.0)
                                                    v35 = (^(v35))
                                                    v128 := GenS9{F0: v38, F1: 9223372036854775807, F2: fi41}
                                                    _ = fi127
                                                    _ = v128
                                                }
                                            case "":
                                                fmt.Println(v3, v2, v2)
                                            }
                                            _ = v109
                                        }
                                        switch "\n\t" {
                                        case "a":
                                            switch 0 {
                                            case 4294967295:
                                                v129 := GenS5{F0: v3}
                                                _ = v129
                                            case -452:
                                                for fi130 := 0; fi130 < 3; fi130++ {
                                                    v39 = v39
                                                    if len(v36) >= 2 {
                                                        v36 = v36[1:len(v36)]
                                                    }
                                                    _ = fi130
                                                }
                                            }
                                        }
                                    }
                                    _ = v106
                                }
                            case 127:
                                _ = len(v36)
                            case -454:
                                _ = len(v36)
                            default:
                                switch 32768 {
                                case 0:
                                    if false {
                                        var v131 uint32 = 4294967295
                                        if len(v36) > 0 {
                                            v36[0] = (fi41 + fi41)
                                        }
                                        _ = len(v36)
                                        _ = v131
                                    } else {
                                        for fi132 := 0; fi132 < 8; fi132++ {
                                            if (42 == 2147483647) {
                                                switch (v2 + v2) {
                                                case "foo bar":
                                                    if len(v36) >= 2 {
                                                        v36 = v36[1:len(v36)]
                                                    }
                                                default:
                                                    if false {
                                                        switch ((v2 + v2) + "") {
                                                        case "hello":
                                                            if (2 <= v3) {
                                                                _ = len(v36)
                                                                v36 = append(v36, (fi132 | (^(fi41))))
                                                            }
                                                        default:
                                                            v2 = "a"
                                                        }
                                                        v133 := GenS15{F0: (v4 / ((-1) | 1)), F1: v42}
                                                        _ = v133
                                                    }
                                                }
                                                fmt.Println(fi41, v39, v38)
                                            } else {
                                                fmt.Println(fi132)
                                                _ = len(v36)
                                                v42 = uint(v35)
                                                switch v2 {
                                                case "\n\t":
                                                    var v134 int16 = -32768
                                                    _ = v134
                                                case "":
                                                    if ((!(v42 == (v39 | v42))) && false) {
                                                        if len(v36) >= 2 {
                                                            v36 = v36[1:len(v36)]
                                                        }
                                                        v135 := GenS16{F0: (v37 | (v37 ^ v37)), F1: v38, F2: 2}
                                                        if true {
                                                            var v136 uint16 = (v38 ^ v38)
                                                            switch int(v37) {
                                                            case 868:
                                                                if ((!false) && (true && ((^(v3)) >= 1))) {
                                                                    var v137 int = (fi41 ^ (fi132 - -1))
                                                                    v138 := GenS17{F0: (v35 / ((v35) | 1))}
                                                                    if (!false) {
                                                                        v137 -= int(v39)
                                                                    } else {
                                                                        v1 = 0.0
                                                                        v1 -= v1
                                                                        var v139 int64 = 307
                                                                        _ = v139
                                                                    }
                                                                    if (42 != 115) {
                                                                        var v140 float64 = 3.14159
                                                                        v137 = (v137 * 9223372036854775807)
                                                                        fmt.Println(v39)
                                                                        var v141 uint = v39
                                                                        switch 2147483647 {
                                                                        case -1:
                                                                            switch v2 {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v142 int32 = 669
                                                                                    v141 -= 255
                                                                                    _ = v142
                                                                                }
                                                                            default:
                                                                                v143 := GenS15{F0: (-(v4)), F1: v39}
                                                                                _ = v143
                                                                            }
                                                                        case 127:
                                                                            var v144 int = (^(v137))
                                                                            _ = v144
                                                                        default:
                                                                            if len(v36) > 0 {
                                                                                v36[0] = fi132
                                                                            }
                                                                            if ((false && (42 <= -1)) || true) {
                                                                                fmt.Println(v136, v141)
                                                                            }
                                                                            if (false && (0 >= -912)) {
                                                                                var v145 uint64 = (v3 ^ (v3 & (v3 & v3)))
                                                                                v3 = 42
                                                                                _ = v145
                                                                            } else {
                                                                                if true {
                                                                                    if len(v36) >= 2 {
                                                                                        v36 = v36[1:len(v36)]
                                                                                    }
                                                                                    fmt.Println(v42)
                                                                                    fmt.Println(fi132, v3)
                                                                                }
                                                                                v37 = (^(v37))
                                                                            }
                                                                        }
                                                                        v146 := GenS4{F0: (v136 / ((uint16(v3)) | 1)), F1: v37, F2: v37}
                                                                        if len(v36) > 0 {
                                                                            v36[0] = 32767
                                                                        }
                                                                        var v147 uint64 = (^(v3))
                                                                        _ = v140
                                                                        _ = v141
                                                                        _ = v146
                                                                        _ = v147
                                                                    }
                                                                    if (((^(v37)) == (v37 << ((65535) & 63))) || true) {
                                                                        var v148 bool = false
                                                                        v137 += fi132
                                                                        _ = v148
                                                                    }
                                                                    _ = v137
                                                                    _ = v138
                                                                }
                                                                fmt.Println(v3, v38, v3)
                                                                var v149 uint32 = 2147483647
                                                                v136 += (^(v136))
                                                                _ = v149
                                                            case 4294967295:
                                                                v150 := GenS18{F0: v35}
                                                                _ = v150
                                                            }
                                                            _ = v136
                                                        } else {
                                                            v151 := GenS19{F0: v4, F1: fi132}
                                                            _ = len(v36)
                                                            var v152 uint = (v42 / ((v39) | 1))
                                                            _ = v151
                                                            _ = v152
                                                        }
                                                        var v153 bool = false
                                                        if ((^(v38)) < 255) {
                                                            v1 -= v1
                                                        }
                                                        _ = v135
                                                        _ = v153
                                                    }
                                                case "a":
                                                    if len(v36) >= 2 {
                                                        v36 = v36[1:len(v36)]
                                                    }
                                                default:
                                                    if false {
                                                        v4 = v4
                                                        v38 = (v38 / ((895) | 1))
                                                    } else {
                                                        v42 -= 18446744073709551615
                                                        switch (v2 + (v2 + v2)) {
                                                        case "hello":
                                                            if (9223372036854775807 >= -702) {
                                                                var v154 int16 = -50
                                                                if (0.1 != v1) {
                                                                    switch (fi132 >> ((uint(v35)) & 63)) {
                                                                    case -373:
                                                                        var v155 int = fi41
                                                                        _ = v155
                                                                    case -9223372036854775808:
                                                                        _ = len(v36)
                                                                    }
                                                                    v156 := GenS20{F0: 232, F1: (v38 + 0)}
                                                                    _ = v156
                                                                }
                                                                if true {
                                                                    v35 = (v35 / ((42) | 1))
                                                                    var v157 uint32 = uint32(v35)
                                                                    _ = v157
                                                                } else {
                                                                    v35 = v35
                                                                }
                                                                _ = v154
                                                            } else {
                                                                fmt.Println(v35, fi41, v2)
                                                                if false {
                                                                    _ = len(v36)
                                                                    if len(v36) >= 2 {
                                                                        v36 = v36[1:len(v36)]
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            if false {
                                                                var v158 int16 = 1
                                                                if true {
                                                                    switch int(v38) {
                                                                    case 32767:
                                                                        switch (v2 + v2) {
                                                                        case "a":
                                                                            v4 += -86
                                                                        case "":
                                                                            if (true && (!false)) {
                                                                                if ((v1 != 0.1) && true) {
                                                                                    if ("" <= ("" + ((v2 + "a") + v2))) {
                                                                                        var v159 string = "a"
                                                                                        _ = v159
                                                                                    }
                                                                                    var v160 uint32 = (v37 / ((v37) | 1))
                                                                                    if true {
                                                                                        fmt.Println(v35)
                                                                                    }
                                                                                    _ = v160
                                                                                } else {
                                                                                    if false {
                                                                                        v4 = v4
                                                                                        v38 += v38
                                                                                    }
                                                                                }
                                                                                if ((v38 + 32767) < uint16(v3)) {
                                                                                    fmt.Println(v158, fi132, fi132)
                                                                                    if (18446744073709551615 <= v3) {
                                                                                        v161 := GenS21{F0: (v158 + (v158 ^ v158))}
                                                                                        if true {
                                                                                            v2 = v2
                                                                                            v1 = v1
                                                                                        } else {
                                                                                            if false {
                                                                                                fmt.Println(v1, v4)
                                                                                                v38 += (v38 + (v38 * (v38 ^ (v38 << ((0) & 63)))))
                                                                                            }
                                                                                            var v162 uint16 = v38
                                                                                            _ = v162
                                                                                        }
                                                                                        if (v37 == (v37 % (((v37 - v37)) | 1))) {
                                                                                            switch v2 {
                                                                                            case "a":
                                                                                                var v163 uint16 = (^(v38))
                                                                                                _ = v163
                                                                                            default:
                                                                                                fmt.Println(v35, v1)
                                                                                            }
                                                                                            var v164 string = v2
                                                                                            switch v2 {
                                                                                            case "\n\t":
                                                                                                v165 := GenS21{F0: (v158 * -32768)}
                                                                                                _ = v165
                                                                                            case "hello":
                                                                                                v35 = (^(v35))
                                                                                            case "a":
                                                                                                if false {
                                                                                                    if ((378 != 621) && false) {
                                                                                                        if len(v36) > 0 {
                                                                                                            v36[0] = (fi132 ^ fi41)
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch fi132 {
                                                                                                        case -215:
                                                                                                            var v166 float64 = 3.14159
                                                                                                            _ = v166
                                                                                                        case 42:
                                                                                                            v3 -= (v3 - (v3 + v3))
                                                                                                        case 32768:
                                                                                                            var v167 float32 = ((3.14159 + 1.5) - 3.14159)
                                                                                                            _ = v167
                                                                                                        default:
                                                                                                            v168 := GenS22{F0: 0}
                                                                                                            _ = v168
                                                                                                        }
                                                                                                        v36 = append(v36, -1)
                                                                                                        v1 = ((((v1 * ((v1 / 1.0) / float64((3.14159 - 1e10)))) + 1.5) - (1e10 + 1e10)) - v1)
                                                                                                        v36 = append(v36, fi132)
                                                                                                    }
                                                                                                    switch v2 {
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                v42 -= v39
                                                                                                                v169 := GenS23{F0: v2, F1: 42, F2: (v37 - (v37 - 1))}
                                                                                                                var v170 int = 32767
                                                                                                                _ = v169
                                                                                                                _ = v170
                                                                                                            } else {
                                                                                                                fmt.Println(v38)
                                                                                                                v38 += 464
                                                                                                                v164 = ""
                                                                                                            }
                                                                                                            if (!(!(255 != (v35 & (^(v35)))))) {
                                                                                                                v3 -= v3
                                                                                                                if len(v36) >= 2 {
                                                                                                                    v36 = v36[1:len(v36)]
                                                                                                                }
                                                                                                                switch fi41 {
                                                                                                                case -9223372036854775808:
                                                                                                                    var v171 float64 = v1
                                                                                                                    _ = v171
                                                                                                                }
                                                                                                            } else {
                                                                                                                if true {
                                                                                                                    if ((fi132 - (fi132 | int(v38))) >= int(v4)) {
                                                                                                                        var v172 int = (fi132 / ((2147483647) | 1))
                                                                                                                        _ = v172
                                                                                                                    } else {
                                                                                                                        v2 = v164
                                                                                                                    }
                                                                                                                    switch "a" {
                                                                                                                    case "a":
                                                                                                                        if true {
                                                                                                                            v3 -= v3
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if ((true && (true || true)) && true) {
                                                                                                                        if false {
                                                                                                                            switch 1 {
                                                                                                                            case 907:
                                                                                                                                v4 -= v4
                                                                                                                            default:
                                                                                                                                fmt.Println(v1)
                                                                                                                            }
                                                                                                                            switch (fi41 - (fi132 | fi132)) {
                                                                                                                            case 32767:
                                                                                                                                if ((0.1 - -2.25) != (v1 - ((float64(3.14159) + (((((v1 - -2.25) - 1e10) - (v1 / (1e10 / (float64(-2.25) - (0.0 + 1.5))))) + v1) * ((v1 - (v1 / 0.1)) - v1))) / v1))) {
                                                                                                                                    v173 := GenS24{F0: 989, F1: (uint8(v3) >= v35)}
                                                                                                                                    v174 := GenS22{F0: int32(v38)}
                                                                                                                                    if false {
                                                                                                                                        switch ("foo bar" + ((v2 + "\n\t") + "\n\t")) {
                                                                                                                                        case "\n\t":
                                                                                                                                            if false {
                                                                                                                                                v37 += (v37 >> (((v39 + 18446744073709551615)) & 63))
                                                                                                                                                var v175 uint8 = v35
                                                                                                                                                _ = v175
                                                                                                                                            } else {
                                                                                                                                                fmt.Println(v37, v37)
                                                                                                                                                var v176 int32 = -341
                                                                                                                                                _ = v176
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            if len(v36) >= 2 {
                                                                                                                                                v36 = v36[1:len(v36)]
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    fmt.Println(v42, v38)
                                                                                                                                    _ = v173
                                                                                                                                    _ = v174
                                                                                                                                }
                                                                                                                                var v177 int32 = 1
                                                                                                                                _ = v177
                                                                                                                            default:
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v4, v158, v2)
                                                                                                                                    if len(v36) > 0 {
                                                                                                                                        v36[0] = (fi132 & -9223372036854775808)
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    _ = len(v36)
                                                                                                                                    v178 := GenS1{F0: 0}
                                                                                                                                    var v179 int64 = 929
                                                                                                                                    _ = v178
                                                                                                                                    _ = v179
                                                                                                                                }
                                                                                                                                var v180 uint16 = v38
                                                                                                                                _ = v180
                                                                                                                            }
                                                                                                                            v42 = (^(v42))
                                                                                                                        } else {
                                                                                                                            if true {
                                                                                                                                v4 -= 0
                                                                                                                                if true {
                                                                                                                                    var v181 uint16 = v38
                                                                                                                                    if false {
                                                                                                                                        if false {
                                                                                                                                            switch (-(fi41)) {
                                                                                                                                            case 42:
                                                                                                                                                var v182 uint32 = 1
                                                                                                                                                _ = v182
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        var v183 int8 = v4
                                                                                                                                        var v184 uint32 = (v37 - v37)
                                                                                                                                        _ = v183
                                                                                                                                        _ = v184
                                                                                                                                    } else {
                                                                                                                                        switch fi132 {
                                                                                                                                        case 836:
                                                                                                                                            fmt.Println(v42, v39)
                                                                                                                                        }
                                                                                                                                        var v185 int = (fi41 | fi41)
                                                                                                                                        v186 := GenS25{F0: true, F1: 2147483647}
                                                                                                                                        _ = v185
                                                                                                                                        _ = v186
                                                                                                                                    }
                                                                                                                                    var v187 uint32 = (v37 - v37)
                                                                                                                                    _ = v181
                                                                                                                                    _ = v187
                                                                                                                                }
                                                                                                                                var v188 string = "foo bar"
                                                                                                                                _ = v188
                                                                                                                            }
                                                                                                                            var v189 uint16 = (v38 ^ v38)
                                                                                                                            _ = v189
                                                                                                                        }
                                                                                                                        var v190 uint = (^(v42))
                                                                                                                        _ = v190
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if (1e10 < 1e10) {
                                                                                                                        fmt.Println(fi41)
                                                                                                                    } else {
                                                                                                                        switch (fi41 * (-(fi41))) {
                                                                                                                        case -1:
                                                                                                                            var v191 int32 = 0
                                                                                                                            _ = v191
                                                                                                                        case 1:
                                                                                                                            fmt.Println(fi41, v38, fi132)
                                                                                                                        }
                                                                                                                        if len(v36) >= 2 {
                                                                                                                            v36 = v36[1:len(v36)]
                                                                                                                        }
                                                                                                                        v4 -= v4
                                                                                                                    }
                                                                                                                    _ = len(v36)
                                                                                                                    var v192 uint16 = v38
                                                                                                                    switch "hello" {
                                                                                                                    case "a":
                                                                                                                        v38 -= (v192 >> (((^(v39))) & 63))
                                                                                                                        if (true || true) {
                                                                                                                            switch "a" {
                                                                                                                            case "foo bar":
                                                                                                                                v158 += 0
                                                                                                                            case "a":
                                                                                                                                switch v2 {
                                                                                                                                case "":
                                                                                                                                    var v193 uint32 = v37
                                                                                                                                    _ = v193
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                var v194 uint64 = (v3 & (^(v3)))
                                                                                                                                _ = v194
                                                                                                                            }
                                                                                                                            v192 = (v192 ^ 255)
                                                                                                                            var v195 float32 = 3.14159
                                                                                                                            _ = v195
                                                                                                                        } else {
                                                                                                                            v39 -= 785
                                                                                                                        }
                                                                                                                    case "":
                                                                                                                        switch fi132 {
                                                                                                                        case -2147483648:
                                                                                                                            switch v2 {
                                                                                                                            case "hello":
                                                                                                                                fmt.Println(v39, v37, fi132)
                                                                                                                            default:
                                                                                                                                switch (v2 + (((("a" + v2) + v164) + ("foo bar" + v164)) + "foo bar")) {
                                                                                                                                case "foo bar":
                                                                                                                                    if ((^(v38)) > (v192 >> ((v39) & 63))) {
                                                                                                                                        var v196 float32 = (0.1 - float32(v1))
                                                                                                                                        v42 += 65535
                                                                                                                                        v1 = float64((1e10 / 1.5))
                                                                                                                                        _ = v196
                                                                                                                                    } else {
                                                                                                                                        v158 = (-(v158))
                                                                                                                                    }
                                                                                                                                case "a":
                                                                                                                                    v3 += (^(v3))
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case 127:
                                                                                                                            v197 := GenS26{F0: v2, F1: 2}
                                                                                                                            _ = v197
                                                                                                                        case 0:
                                                                                                                            v198 := GenS11{F0: (-(fi41))}
                                                                                                                            _ = v198
                                                                                                                        default:
                                                                                                                            if (-83 < (v4 << ((v39) & 63))) {
                                                                                                                                switch fi41 {
                                                                                                                                case 795:
                                                                                                                                    fmt.Println(v164, v4, v3)
                                                                                                                                case 32767:
                                                                                                                                    switch v164 {
                                                                                                                                    case "":
                                                                                                                                        switch v164 {
                                                                                                                                        case "\n\t":
                                                                                                                                            v36 = append(v36, (fi132 + (fi41 ^ fi41)))
                                                                                                                                        case "foo bar":
                                                                                                                                            var v199 string = "foo bar"
                                                                                                                                            fmt.Println(fi132)
                                                                                                                                            v37 += v37
                                                                                                                                            _ = v199
                                                                                                                                        case "":
                                                                                                                                            switch (fi41 & fi41) {
                                                                                                                                            case -583:
                                                                                                                                                v200 := GenS12{F0: 42, F1: 1, F2: v158}
                                                                                                                                                _ = v200
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            v201 := GenS17{F0: v35}
                                                                                                                                            _ = v201
                                                                                                                                        }
                                                                                                                                    case "\n\t":
                                                                                                                                        if len(v36) >= 2 {
                                                                                                                                            v36 = v36[1:len(v36)]
                                                                                                                                        }
                                                                                                                                    case "a":
                                                                                                                                        if (1.5 != v1) {
                                                                                                                                            switch v164 {
                                                                                                                                            case "":
                                                                                                                                                v202 := GenS27{F0: (^(v39))}
                                                                                                                                                _ = v202
                                                                                                                                            case "foo bar":
                                                                                                                                                if true {
                                                                                                                                                    v4 = -50
                                                                                                                                                    if false {
                                                                                                                                                        var v203 int16 = (v158 * v158)
                                                                                                                                                        var v204 float64 = (((v1 / float64(((0.1 - 0.1) * (3.14159 - (-2.25 * (3.14159 * float32(v1))))))) / ((0.1 * v1) * float64((0.0 * 0.0)))) - (((v1 - ((((1.5 + v1) - v1) - (((3.14159 / v1) * v1) * (v1 * v1))) * -2.25)) + float64(1e10)) + (v1 + ((0.1 * ((v1 - 1.5) + -2.25)) + float64(-2.25)))))
                                                                                                                                                        v205 := GenS27{F0: v39}
                                                                                                                                                        v36 = append(v36, (fi132 % ((128) | 1)))
                                                                                                                                                        _ = v203
                                                                                                                                                        _ = v204
                                                                                                                                                        _ = v205
                                                                                                                                                    } else {
                                                                                                                                                        if (32767 <= int(v35)) {
                                                                                                                                                            if len(v36) >= 2 {
                                                                                                                                                                v36 = v36[1:len(v36)]
                                                                                                                                                            }
                                                                                                                                                            switch fi41 {
                                                                                                                                                            case 0:
                                                                                                                                                                var v206 uint16 = 567
                                                                                                                                                                _ = v206
                                                                                                                                                            case 42:
                                                                                                                                                                v1 = 0.0
                                                                                                                                                            }
                                                                                                                                                            v36 = append(v36, 128)
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    var v207 uint8 = 127
                                                                                                                                                    _ = v207
                                                                                                                                                } else {
                                                                                                                                                    if false {
                                                                                                                                                        if (float64(-2.25) > 1.5) {
                                                                                                                                                            switch int(v37) {
                                                                                                                                                            case 2147483647:
                                                                                                                                                                if (v1 >= 3.14159) {
                                                                                                                                                                    fmt.Println(v2, v1)
                                                                                                                                                                    switch ("" + v164) {
                                                                                                                                                                    case "hello":
                                                                                                                                                                        _ = len(v36)
                                                                                                                                                                    case "":
                                                                                                                                                                        v164 = (((v2 + (("\n\t" + "hello") + "hello")) + "foo bar") + (v2 + v2))
                                                                                                                                                                    case "\n\t":
                                                                                                                                                                        v164 = ("hello" + "a")
                                                                                                                                                                        v208 := GenS28{F0: (v3 >> (((v42 & (v39 * (v42 ^ v39)))) & 63))}
                                                                                                                                                                        _ = v208
                                                                                                                                                                    }
                                                                                                                                                                    fmt.Println(v192, v3)
                                                                                                                                                                    var v209 uint = (^(v39))
                                                                                                                                                                    _ = v209
                                                                                                                                                                }
                                                                                                                                                            case -360:
                                                                                                                                                                v210 := GenS16{F0: v37, F1: (v192 % ((0) | 1)), F2: 494}
                                                                                                                                                                _ = v210
                                                                                                                                                            case 9223372036854775807:
                                                                                                                                                                v2 = ""
                                                                                                                                                            default:
                                                                                                                                                                var v211 uint = (v39 - v39)
                                                                                                                                                                _ = v211
                                                                                                                                                            }
                                                                                                                                                        } else {
                                                                                                                                                            if false {
                                                                                                                                                                v35 += 1
                                                                                                                                                                switch int(v3) {
                                                                                                                                                                case 128:
                                                                                                                                                                    _ = len(v36)
                                                                                                                                                                default:
                                                                                                                                                                    fmt.Println(fi132, v3)
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        v212 := GenS29{F0: v35, F1: v164, F2: (^(v38))}
                                                                                                                                                        var v213 int = (fi41 & fi132)
                                                                                                                                                        _ = v212
                                                                                                                                                        _ = v213
                                                                                                                                                    } else {
                                                                                                                                                        v192 = uint16(v35)
                                                                                                                                                        var v214 uint = (^(v39))
                                                                                                                                                        _ = v214
                                                                                                                                                    }
                                                                                                                                                    if (71 < (^(v35))) {
                                                                                                                                                        _ = len(v36)
                                                                                                                                                        var v215 uint32 = uint32(v35)
                                                                                                                                                        if len(v36) >= 2 {
                                                                                                                                                            v36 = v36[1:len(v36)]
                                                                                                                                                        }
                                                                                                                                                        if ((0.1 + -2.25) <= (0.1 / (((0.0 + 1e10) / -2.25) + (3.14159 / 1.0)))) {
                                                                                                                                                            var v216 uint64 = uint64(v37)
                                                                                                                                                            if ((v35 + (v35 + 42)) <= (v35 + v35)) {
                                                                                                                                                                switch (^(fi132)) {
                                                                                                                                                                case 32768:
                                                                                                                                                                    v217 := GenS30{F0: (v215 & (v215 + v215)), F1: -325, F2: 426}
                                                                                                                                                                    var v218 int16 = (v158 ^ (^(v158)))
                                                                                                                                                                    switch fi132 {
                                                                                                                                                                    case 127:
                                                                                                                                                                        switch 9223372036854775807 {
                                                                                                                                                                        case 1:
                                                                                                                                                                            v219 := GenS31{F0: 2147483647, F1: (v216 + 377), F2: 65535}
                                                                                                                                                                            _ = v219
                                                                                                                                                                        default:
                                                                                                                                                                            var v220 float64 = v1
                                                                                                                                                                            _ = v220
                                                                                                                                                                        }
                                                                                                                                                                    case 1:
                                                                                                                                                                        var v221 int = fi132
                                                                                                                                                                        _ = v221
                                                                                                                                                                    case 9223372036854775807:
                                                                                                                                                                        switch (fi132 & fi132) {
                                                                                                                                                                        case 810:
                                                                                                                                                                            switch "\n\t" {
                                                                                                                                                                            case "":
                                                                                                                                                                                v222 := GenS25{F0: true, F1: 0}
                                                                                                                                                                                _ = v222
                                                                                                                                                                            case "a":
                                                                                                                                                                                if (!((v4 - (v4 & (^(v4)))) != (v4 >> ((v42) & 63)))) {
                                                                                                                                                                                    if len(v36) > 0 {
                                                                                                                                                                                        v36[0] = fi41
                                                                                                                                                                                    }
                                                                                                                                                                                    switch int(v37) {
                                                                                                                                                                                    case 32768:
                                                                                                                                                                                        switch v2 {
                                                                                                                                                                                        case "hello":
                                                                                                                                                                                            switch "\n\t" {
                                                                                                                                                                                            case "\n\t":
                                                                                                                                                                                                if len(v36) > 0 {
                                                                                                                                                                                                    v36[0] = 78
                                                                                                                                                                                                }
                                                                                                                                                                                            default:
                                                                                                                                                                                                var v223 bool = ((v37 * (v37 + 2)) < v37)
                                                                                                                                                                                                _ = v223
                                                                                                                                                                                            }
                                                                                                                                                                                        case "a":
                                                                                                                                                                                            v2 = "\n\t"
                                                                                                                                                                                        case "foo bar":
                                                                                                                                                                                            var v224 uint32 = 1
                                                                                                                                                                                            _ = v224
                                                                                                                                                                                        }
                                                                                                                                                                                        v225 := GenS3{F0: (v39 << (((v42 | uint(v35))) & 63)), F1: "a", F2: ((((v164 + "") + v164) + v164) + "a")}
                                                                                                                                                                                        switch v164 {
                                                                                                                                                                                        case "hello":
                                                                                                                                                                                            switch (v164 + "hello") {
                                                                                                                                                                                            case "\n\t":
                                                                                                                                                                                                switch 9223372036854775807 {
                                                                                                                                                                                                case 128:
                                                                                                                                                                                                    var v226 int = (^(fi41))
                                                                                                                                                                                                    _ = v226
                                                                                                                                                                                                }
                                                                                                                                                                                            case "":
                                                                                                                                                                                                v42 = (v39 / ((v39) | 1))
                                                                                                                                                                                            default:
                                                                                                                                                                                                switch 128 {
                                                                                                                                                                                                case 34:
                                                                                                                                                                                                    var v227 int16 = -1
                                                                                                                                                                                                    _ = v227
                                                                                                                                                                                                default:
                                                                                                                                                                                                    v38 += v38
                                                                                                                                                                                                }
                                                                                                                                                                                            }
                                                                                                                                                                                        default:
                                                                                                                                                                                            var v228 int16 = -32768
                                                                                                                                                                                            _ = v228
                                                                                                                                                                                        }
                                                                                                                                                                                        _ = v225
                                                                                                                                                                                    case 1:
                                                                                                                                                                                        v37 = v37
                                                                                                                                                                                    }
                                                                                                                                                                                    v229 := GenS1{F0: 1}
                                                                                                                                                                                    _ = v229
                                                                                                                                                                                }
                                                                                                                                                                                v37 = 1
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    default:
                                                                                                                                                                        v39 = (v39 + (v39 ^ (v39 >> ((uint(v35)) & 63))))
                                                                                                                                                                        var v230 int = (fi41 << (((v39 * v39)) & 63))
                                                                                                                                                                        _ = v230
                                                                                                                                                                    }
                                                                                                                                                                    var v231 uint16 = (v38 % ((598) | 1))
                                                                                                                                                                    _ = v217
                                                                                                                                                                    _ = v218
                                                                                                                                                                    _ = v231
                                                                                                                                                                case 42:
                                                                                                                                                                    v232 := GenS28{F0: v216}
                                                                                                                                                                    _ = v232
                                                                                                                                                                }
                                                                                                                                                                if ("a" >= "") {
                                                                                                                                                                    if true {
                                                                                                                                                                        switch -2147483648 {
                                                                                                                                                                        case 1:
                                                                                                                                                                            _ = len(v36)
                                                                                                                                                                        case -622:
                                                                                                                                                                            v233 := GenS3{F0: (^(v39)), F1: ((v164 + (v2 + v164)) + v2), F2: v164}
                                                                                                                                                                            _ = v233
                                                                                                                                                                        case 32768:
                                                                                                                                                                            var v234 uint8 = 1
                                                                                                                                                                            _ = v234
                                                                                                                                                                        default:
                                                                                                                                                                            switch (((("\n\t" + v2) + v2) + v2) + ((((v2 + "") + v164) + "foo bar") + v2)) {
                                                                                                                                                                            case "a":
                                                                                                                                                                                var v235 bool = (!(false || true))
                                                                                                                                                                                _ = v235
                                                                                                                                                                            case "\n\t":
                                                                                                                                                                                var v236 bool = false
                                                                                                                                                                                _ = v236
                                                                                                                                                                            case "hello":
                                                                                                                                                                                switch (^(fi132)) {
                                                                                                                                                                                case 128:
                                                                                                                                                                                    var v237 bool = ((v158 & (v158 + 32767)) >= v158)
                                                                                                                                                                                    _ = v237
                                                                                                                                                                                case 696:
                                                                                                                                                                                    v158 += (v158 + v158)
                                                                                                                                                                                case -66:
                                                                                                                                                                                    v238 := GenS3{F0: 981, F1: v2, F2: "hello"}
                                                                                                                                                                                    fmt.Println(v1, v37, v216)
                                                                                                                                                                                    v239 := GenS20{F0: 1000, F1: (v38 >> (((v39 ^ v39)) & 63))}
                                                                                                                                                                                    _ = v238
                                                                                                                                                                                    _ = v239
                                                                                                                                                                                default:
                                                                                                                                                                                    v2 = "foo bar"
                                                                                                                                                                                }
                                                                                                                                                                            default:
                                                                                                                                                                                v4 = (v4 / ((126) | 1))
                                                                                                                                                                            }
                                                                                                                                                                            v240 := GenS21{F0: v158}
                                                                                                                                                                            fmt.Println(v35, fi132, v158)
                                                                                                                                                                            _ = v240
                                                                                                                                                                        }
                                                                                                                                                                        var v241 float32 = 1e10
                                                                                                                                                                        if false {
                                                                                                                                                                            var v242 int = (-(fi132))
                                                                                                                                                                            if false {
                                                                                                                                                                                v243 := GenS32{F0: 810, F1: 168}
                                                                                                                                                                                var v244 int16 = v158
                                                                                                                                                                                fmt.Println(v35, v1)
                                                                                                                                                                                v39 = uint(v35)
                                                                                                                                                                                _ = v243
                                                                                                                                                                                _ = v244
                                                                                                                                                                            } else {
                                                                                                                                                                                fmt.Println(v242)
                                                                                                                                                                                switch v2 {
                                                                                                                                                                                case "":
                                                                                                                                                                                    switch v2 {
                                                                                                                                                                                    case "a":
                                                                                                                                                                                        v37 = (v37 | (^(v37)))
                                                                                                                                                                                    }
                                                                                                                                                                                default:
                                                                                                                                                                                    v245 := GenS8{F0: 990, F1: v4, F2: uint16(v42)}
                                                                                                                                                                                    _ = v245
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                            _ = v242
                                                                                                                                                                        }
                                                                                                                                                                        _ = v241
                                                                                                                                                                    } else {
                                                                                                                                                                        switch -9223372036854775808 {
                                                                                                                                                                        case -367:
                                                                                                                                                                            v39 = v42
                                                                                                                                                                        case -463:
                                                                                                                                                                            if ((!true) || true) {
                                                                                                                                                                                fmt.Println(v216, fi132, fi41)
                                                                                                                                                                                v37 -= uint32(v192)
                                                                                                                                                                            } else {
                                                                                                                                                                                v1 -= (v1 + (0.1 * ((1.5 + (3.14159 - v1)) - 0.1)))
                                                                                                                                                                                v3 += 0
                                                                                                                                                                                switch (fi132 + fi41) {
                                                                                                                                                                                case 32767:
                                                                                                                                                                                    v36 = append(v36, int(v4))
                                                                                                                                                                                case -1:
                                                                                                                                                                                    v38 = 255
                                                                                                                                                                                case 32768:
                                                                                                                                                                                    var v246 float32 = 0.0
                                                                                                                                                                                    _ = v246
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                        default:
                                                                                                                                                                            fmt.Println(v2, v37)
                                                                                                                                                                        }
                                                                                                                                                                        v39 = (v39 % (((v39 ^ (v42 >> (((v39 >> (((v42 + 735)) & 63))) & 63)))) | 1))
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                v247 := GenS30{F0: (v215 << ((v42) & 63)), F1: 567, F2: v192}
                                                                                                                                                                _ = v247
                                                                                                                                                            }
                                                                                                                                                            _ = v216
                                                                                                                                                        } else {
                                                                                                                                                            var v248 uint16 = (v192 >> ((v42) & 63))
                                                                                                                                                            switch v2 {
                                                                                                                                                            case "\n\t":
                                                                                                                                                                fmt.Println(v37, v2)
                                                                                                                                                            case "hello":
                                                                                                                                                                v42 = 100
                                                                                                                                                            }
                                                                                                                                                            _ = v248
                                                                                                                                                        }
                                                                                                                                                        v4 += (v4 | 74)
                                                                                                                                                        _ = v215
                                                                                                                                                    } else {
                                                                                                                                                        var v249 int = (-(fi41))
                                                                                                                                                        _ = v249
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            case "a":
                                                                                                                                                fmt.Println(v158, v42)
                                                                                                                                            default:
                                                                                                                                                v158 = -1
                                                                                                                                            }
                                                                                                                                            fmt.Println(fi132, v38, v1)
                                                                                                                                            var v250 int32 = 2147483647
                                                                                                                                            v251 := GenS9{F0: uint16(v37), F1: int64(v35), F2: fi41}
                                                                                                                                            v252 := GenS13{F0: v192, F1: v39}
                                                                                                                                            if false {
                                                                                                                                                v36 = append(v36, 42)
                                                                                                                                            } else {
                                                                                                                                                v250 = int32(v192)
                                                                                                                                            }
                                                                                                                                            _ = v250
                                                                                                                                            _ = v251
                                                                                                                                            _ = v252
                                                                                                                                        } else {
                                                                                                                                            var v253 int8 = v4
                                                                                                                                            switch 4294967295 {
                                                                                                                                            case 128:
                                                                                                                                                if len(v36) >= 2 {
                                                                                                                                                    v36 = v36[1:len(v36)]
                                                                                                                                                }
                                                                                                                                            case 127:
                                                                                                                                                fmt.Println(fi41)
                                                                                                                                            case 900:
                                                                                                                                                v4 = 0
                                                                                                                                            default:
                                                                                                                                                switch 4294967295 {
                                                                                                                                                case 128:
                                                                                                                                                    v2 = v164
                                                                                                                                                default:
                                                                                                                                                    var v254 bool = true
                                                                                                                                                    _ = v254
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            fmt.Println(fi41, v158)
                                                                                                                                            _ = v253
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v255 uint = v39
                                                                                                                                        _ = v255
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    switch (v164 + v2) {
                                                                                                                                    case "\n\t":
                                                                                                                                        var v256 uint8 = 2
                                                                                                                                        _ = v256
                                                                                                                                    default:
                                                                                                                                        fmt.Println(v3)
                                                                                                                                    }
                                                                                                                                    var v257 bool = false
                                                                                                                                    _ = v257
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                v258 := GenS33{F0: (v35 << ((62) & 63)), F1: (^(v39))}
                                                                                                                                if true {
                                                                                                                                    var v259 int = -1
                                                                                                                                    _ = v259
                                                                                                                                }
                                                                                                                                _ = v258
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v192
                                                                                                                }
                                                                                                                if true {
                                                                                                                    if ((-(v4)) < (v4 ^ 85)) {
                                                                                                                        v36 = append(v36, fi41)
                                                                                                                        v260 := GenS32{F0: 255, F1: -2147483648}
                                                                                                                        v42 = 351
                                                                                                                        _ = v260
                                                                                                                    }
                                                                                                                    var v261 uint32 = (v37 * 742)
                                                                                                                    _ = v261
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(fi41, v4)
                                                                                                        }
                                                                                                    case "":
                                                                                                        if len(v36) >= 2 {
                                                                                                            v36 = v36[1:len(v36)]
                                                                                                        }
                                                                                                    case "a":
                                                                                                        switch (fi41 << ((v42) & 63)) {
                                                                                                        case 32767:
                                                                                                            fmt.Println(fi41, v158, v35)
                                                                                                        default:
                                                                                                            v2 = v2
                                                                                                        }
                                                                                                    default:
                                                                                                        var v262 string = "\n\t"
                                                                                                        _ = v262
                                                                                                    }
                                                                                                } else {
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        v263 := GenS34{F0: v2, F1: 0, F2: (v42 | (v39 << ((v39) & 63)))}
                                                                                                        _ = v263
                                                                                                    default:
                                                                                                        v2 = v2
                                                                                                    }
                                                                                                    if len(v36) >= 2 {
                                                                                                        v36 = v36[1:len(v36)]
                                                                                                    }
                                                                                                    v2 = "a"
                                                                                                }
                                                                                            }
                                                                                            _ = v164
                                                                                        } else {
                                                                                            if (((0.1 + 3.14159) != 1e10) || ((v2 + "foo bar") < "a")) {
                                                                                                var v264 uint64 = v3
                                                                                                _ = v264
                                                                                            } else {
                                                                                                var v265 float64 = 3.14159
                                                                                                _ = v265
                                                                                            }
                                                                                        }
                                                                                        _ = v161
                                                                                    } else {
                                                                                        v1 -= v1
                                                                                        v3 = (v3 + 356)
                                                                                        v38 = v38
                                                                                    }
                                                                                    if false {
                                                                                        v36 = append(v36, -2147483648)
                                                                                        if true {
                                                                                            if len(v36) >= 2 {
                                                                                                v36 = v36[1:len(v36)]
                                                                                            }
                                                                                            var v266 int8 = (v4 - (v4 | (^(v4))))
                                                                                            fmt.Println(v42, fi132)
                                                                                            _ = v266
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if ((v3 == 827) || false) {
                                                                                    if len(v36) >= 2 {
                                                                                        v36 = v36[1:len(v36)]
                                                                                    }
                                                                                    var v267 int64 = 1
                                                                                    if len(v36) > 0 {
                                                                                        v36[0] = (fi132 % ((fi41) | 1))
                                                                                    }
                                                                                    v268 := GenS35{F0: (v38 & (^(v38))), F1: int(v4), F2: 552}
                                                                                    _ = v267
                                                                                    _ = v268
                                                                                } else {
                                                                                    if len(v36) > 0 {
                                                                                        v36[0] = fi41
                                                                                    }
                                                                                    v37 = 65535
                                                                                    var v269 int32 = 0
                                                                                    _ = v269
                                                                                }
                                                                            }
                                                                        default:
                                                                            switch (v2 + (((v2 + (v2 + "\n\t")) + (v2 + "\n\t")) + (("a" + v2) + ((v2 + (("a" + v2) + v2)) + "")))) {
                                                                            case "a":
                                                                                if len(v36) >= 2 {
                                                                                    v36 = v36[1:len(v36)]
                                                                                }
                                                                            default:
                                                                                var v270 uint = (v39 % ((v42) | 1))
                                                                                _ = v270
                                                                            }
                                                                        }
                                                                        v271 := GenS25{F0: (!true), F1: 0}
                                                                        var v272 int32 = 590
                                                                        fmt.Println(v272)
                                                                        var v273 int16 = -1
                                                                        _ = v271
                                                                        _ = v272
                                                                        _ = v273
                                                                    case 128:
                                                                        if ((!false) || (0 < v3)) {
                                                                            if true {
                                                                                if (v37 < 0) {
                                                                                    var v274 uint16 = v38
                                                                                    if false {
                                                                                        var v275 uint16 = 255
                                                                                        v276 := GenS24{F0: v37, F1: true}
                                                                                        _ = v275
                                                                                        _ = v276
                                                                                    } else {
                                                                                        if len(v36) > 0 {
                                                                                            v36[0] = (fi132 << ((uint(v35)) & 63))
                                                                                        }
                                                                                        switch (fi132 | (fi132 % ((fi132) | 1))) {
                                                                                        case 2147483647:
                                                                                            switch "" {
                                                                                            case "hello":
                                                                                                if (!(false || (-83 != v158))) {
                                                                                                    var v277 float32 = 3.14159
                                                                                                    v3 += (v3 << ((v42) & 63))
                                                                                                    v2 = "hello"
                                                                                                    _ = v277
                                                                                                } else {
                                                                                                    var v278 uint16 = (v38 % ((v38) | 1))
                                                                                                    if (false && (0 < (^(v3)))) {
                                                                                                        v279 := GenS7{F0: 0}
                                                                                                        _ = v279
                                                                                                    } else {
                                                                                                        if (true || (!false)) {
                                                                                                            v42 -= 18446744073709551615
                                                                                                        }
                                                                                                        var v280 int16 = (v158 & (v158 * (-(v158))))
                                                                                                        v280 = v158
                                                                                                        _ = v280
                                                                                                    }
                                                                                                    if (v158 > 0) {
                                                                                                        switch "a" {
                                                                                                        case "hello":
                                                                                                            if len(v36) >= 2 {
                                                                                                                v36 = v36[1:len(v36)]
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            switch fi132 {
                                                                                                            case 1:
                                                                                                                v42 = v42
                                                                                                            case 32767:
                                                                                                                v278 -= v274
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            v36 = append(v36, 127)
                                                                                                        default:
                                                                                                            v278 += v274
                                                                                                        }
                                                                                                        v281 := GenS17{F0: v35}
                                                                                                        _ = v281
                                                                                                    } else {
                                                                                                        v158 -= v158
                                                                                                    }
                                                                                                    _ = v278
                                                                                                }
                                                                                            default:
                                                                                                v1 = 1e10
                                                                                            }
                                                                                        case 369:
                                                                                            var v282 uint32 = v37
                                                                                            _ = v282
                                                                                        case 0:
                                                                                            var v283 uint8 = 150
                                                                                            _ = v283
                                                                                        default:
                                                                                            fmt.Println(v38, v39, v38)
                                                                                            var v284 uint32 = 65535
                                                                                            fmt.Println(v35, v1, v42)
                                                                                            var v285 int64 = 0
                                                                                            if false {
                                                                                                var v286 bool = (((v285 + -9223372036854775808) != v285) || ((v37 * v284) >= v284))
                                                                                                var v287 float32 = ((1.5 + 3.14159) - (float32(v1) + (0.1 / ((1e10 + (3.14159 / float32(v1))) - 1e10))))
                                                                                                if len(v36) >= 2 {
                                                                                                    v36 = v36[1:len(v36)]
                                                                                                }
                                                                                                _ = v286
                                                                                                _ = v287
                                                                                            }
                                                                                            _ = v284
                                                                                            _ = v285
                                                                                        }
                                                                                    }
                                                                                    _ = v274
                                                                                }
                                                                                _ = len(v36)
                                                                            } else {
                                                                                if true {
                                                                                    var v288 uint = (v42 * v42)
                                                                                    var v289 int = (fi41 >> ((4294967295) & 63))
                                                                                    _ = v288
                                                                                    _ = v289
                                                                                }
                                                                            }
                                                                            var v290 int32 = 42
                                                                            v291 := GenS27{F0: (v39 % ((v39) | 1))}
                                                                            _ = v290
                                                                            _ = v291
                                                                        } else {
                                                                            v38 = (v38 % (((v38 * v38)) | 1))
                                                                            var v292 uint = (v39 * v39)
                                                                            var v293 string = ("" + "foo bar")
                                                                            _ = v292
                                                                            _ = v293
                                                                        }
                                                                    case -433:
                                                                        v4 = (^(v4))
                                                                    default:
                                                                        if true {
                                                                            switch fi132 {
                                                                            case 1:
                                                                                v38 = 0
                                                                            default:
                                                                                switch (fi132 >> ((v42) & 63)) {
                                                                                case -329:
                                                                                    v158 -= v158
                                                                                case 624:
                                                                                    switch 757 {
                                                                                    case 2147483647:
                                                                                        v2 = "hello"
                                                                                    case -349:
                                                                                        _ = len(v36)
                                                                                    case -9223372036854775808:
                                                                                        if false {
                                                                                            _ = len(v36)
                                                                                            v294 := GenS36{F0: 9223372036854775807, F1: (fi41 * 2147483647), F2: (fi41 | fi132)}
                                                                                            if false {
                                                                                                v38 += uint16(v37)
                                                                                                v37 = (v37 & uint32(v38))
                                                                                                var v295 float64 = v1
                                                                                                _ = v295
                                                                                            } else {
                                                                                                switch fi132 {
                                                                                                case -9223372036854775808:
                                                                                                    var v296 uint8 = v35
                                                                                                    switch v2 {
                                                                                                    case "hello":
                                                                                                        v42 = 0
                                                                                                    }
                                                                                                    _ = v296
                                                                                                case -2147483648:
                                                                                                    fmt.Println(v4, v42, v39)
                                                                                                default:
                                                                                                    if (-413 < -832) {
                                                                                                        var v297 float32 = 1.5
                                                                                                        v298 := GenS17{F0: 0}
                                                                                                        _ = v297
                                                                                                        _ = v298
                                                                                                    } else {
                                                                                                        var v299 uint16 = v38
                                                                                                        switch fi132 {
                                                                                                        case 4294967295:
                                                                                                            switch fi41 {
                                                                                                            case -674:
                                                                                                                switch fi132 {
                                                                                                                case 4294967295:
                                                                                                                    fmt.Println(v4)
                                                                                                                case 127:
                                                                                                                    var v300 uint8 = uint8(v3)
                                                                                                                    _ = v300
                                                                                                                default:
                                                                                                                    if (!(((!false) || false) && true)) {
                                                                                                                        fmt.Println(v4)
                                                                                                                        v301 := GenS37{F0: v37, F1: v37, F2: (v3 * (v3 >> (((v39 ^ v39)) & 63)))}
                                                                                                                        if len(v36) > 0 {
                                                                                                                            v36[0] = fi132
                                                                                                                        }
                                                                                                                        _ = v301
                                                                                                                    }
                                                                                                                }
                                                                                                            case 9223372036854775807:
                                                                                                                v37 += 0
                                                                                                            default:
                                                                                                                var v302 uint16 = v299
                                                                                                                _ = v302
                                                                                                            }
                                                                                                        case 42:
                                                                                                            var v303 string = (("\n\t" + v2) + v2)
                                                                                                            _ = v303
                                                                                                        default:
                                                                                                            if (v37 <= (v37 * uint32(v35))) {
                                                                                                                if ((true || (float64(0.1) != (1.5 / 1.0))) && false) {
                                                                                                                    fmt.Println(v39)
                                                                                                                    v304 := GenS38{F0: (v35 % (((v35 % ((v35) | 1))) | 1))}
                                                                                                                    if true {
                                                                                                                        var v305 uint8 = 139
                                                                                                                        v3 -= v3
                                                                                                                        v306 := GenS14{F0: v38}
                                                                                                                        var v307 uint16 = v299
                                                                                                                        _ = v305
                                                                                                                        _ = v306
                                                                                                                        _ = v307
                                                                                                                    }
                                                                                                                    var v308 uint32 = v37
                                                                                                                    _ = v304
                                                                                                                    _ = v308
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v299
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v294
                                                                                        } else {
                                                                                            var v309 string = v2
                                                                                            if (42 < 66) {
                                                                                                var v310 uint8 = (v35 >> ((v42) & 63))
                                                                                                v309 = ("\n\t" + "hello")
                                                                                                fmt.Println(v3, fi132, v158)
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    if (((v309 + v2) + (v309 + v2)) <= v309) {
                                                                                                        v42 -= v39
                                                                                                        var v311 float64 = v1
                                                                                                        _ = v311
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v1 -= v1
                                                                                                default:
                                                                                                    v42 = v39
                                                                                                }
                                                                                                var v312 int = fi41
                                                                                                v313 := GenS39{F0: -32768, F1: v158}
                                                                                                v36 = append(v36, (fi132 | v312))
                                                                                                _ = v310
                                                                                                _ = v312
                                                                                                _ = v313
                                                                                            }
                                                                                            _ = v309
                                                                                        }
                                                                                    default:
                                                                                        if true {
                                                                                            v314 := GenS30{F0: (v37 / ((v37) | 1)), F1: 0, F2: uint16(v3)}
                                                                                            _ = v314
                                                                                        } else {
                                                                                            v36 = append(v36, 393)
                                                                                            var v315 float32 = -2.25
                                                                                            _ = v315
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v35 -= v35
                                                                                }
                                                                            }
                                                                            var v316 uint = (v42 ^ (v42 + v42))
                                                                            v3 = 1
                                                                            _ = v316
                                                                        } else {
                                                                            fmt.Println(v158, v38)
                                                                            if true {
                                                                                v35 = v35
                                                                                fmt.Println(v37)
                                                                            } else {
                                                                                if (420 != -2147483648) {
                                                                                    v317 := GenS40{F0: v37}
                                                                                    _ = len(v36)
                                                                                    var v318 int16 = v158
                                                                                    _ = v317
                                                                                    _ = v318
                                                                                }
                                                                                v42 = (v42 << ((998) & 63))
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = len(v36)
                                                                } else {
                                                                    if (!(float64((float32(v1) * (0.0 - 1.5))) >= (v1 * v1))) {
                                                                        var v319 int32 = -2147483648
                                                                        fmt.Println(v38, fi132)
                                                                        if true {
                                                                            if ((fi41 | fi41) <= (fi41 + (fi41 | (fi132 << ((1) & 63))))) {
                                                                                v158 -= (v158 - v158)
                                                                            } else {
                                                                                var v320 bool = (571 > 98)
                                                                                if v320 {
                                                                                    switch v2 {
                                                                                    case "foo bar":
                                                                                        v36 = append(v36, 606)
                                                                                    }
                                                                                    fmt.Println(v158, v37)
                                                                                    var v321 bool = (true || v320)
                                                                                    _ = v321
                                                                                } else {
                                                                                    switch ((v2 + "\n\t") + (v2 + ("hello" + "\n\t"))) {
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            v2 = (v2 + v2)
                                                                                        } else {
                                                                                            if v320 {
                                                                                                var v322 uint8 = v35
                                                                                                v158 = v158
                                                                                                var v323 string = "\n\t"
                                                                                                _ = v322
                                                                                                _ = v323
                                                                                            } else {
                                                                                                v324 := GenS2{F0: (492 > v37), F1: (((1.5 / 1e10) - (0.1 / (float32(-2.25) / 0.1))) < 0.1), F2: ""}
                                                                                                _ = len(v36)
                                                                                                var v325 float64 = (-2.25 - ((((v1 / v1) - (v1 * ((v1 / ((float64(-2.25) * (((((v1 / v1) - (0.0 * v1)) * 0.1) - 0.1) * (-2.25 - v1))) / v1)) + (v1 + (((float64(-2.25) - (v1 + v1)) - (((v1 - v1) / (float64((3.14159 / 0.1)) - -2.25)) * 1.5)) - v1))))) / v1) * ((((v1 - v1) + (3.14159 - v1)) + float64(0.1)) * (((v1 + 3.14159) + v1) * -2.25))))
                                                                                                v3 = v3
                                                                                                if (true && v320) {
                                                                                                    switch v2 {
                                                                                                    case "":
                                                                                                        v326 := GenS41{F0: v37}
                                                                                                        _ = v326
                                                                                                    case "foo bar":
                                                                                                        if len(v36) > 0 {
                                                                                                            v36[0] = 51
                                                                                                        }
                                                                                                    default:
                                                                                                        switch v2 {
                                                                                                        case "":
                                                                                                            var v327 int64 = 9223372036854775807
                                                                                                            _ = v327
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v2)
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    v328 := GenS24{F0: v37, F1: (!(!((^(v3)) >= (v3 << ((uint(v38)) & 63)))))}
                                                                                                    _ = v328
                                                                                                }
                                                                                                _ = v324
                                                                                                _ = v325
                                                                                            }
                                                                                            if ((!false) && ((1e10 / 1.5) <= (-2.25 - (0.1 + 3.14159)))) {
                                                                                                v42 = (v39 >> ((v42) & 63))
                                                                                                v319 = 42
                                                                                                if len(v36) >= 2 {
                                                                                                    v36 = v36[1:len(v36)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v42 = 2
                                                                                    v39 = v39
                                                                                }
                                                                                if len(v36) >= 2 {
                                                                                    v36 = v36[1:len(v36)]
                                                                                }
                                                                                fmt.Println(v1, v320)
                                                                                v3 = v3
                                                                                _ = v320
                                                                            }
                                                                            if (!(fi132 <= -45)) {
                                                                                v42 += v42
                                                                                v329 := GenS42{F0: -9223372036854775808, F1: v3, F2: v39}
                                                                                v330 := GenS43{F0: ((1.5 >= v1) || false)}
                                                                                _ = v329
                                                                                _ = v330
                                                                            } else {
                                                                                if true {
                                                                                    fmt.Println(fi41)
                                                                                    if (v39 < 442) {
                                                                                        v36 = append(v36, 2147483647)
                                                                                    }
                                                                                    v331 := GenS44{F0: (v158 << (((^(v42))) & 63))}
                                                                                    _ = v331
                                                                                } else {
                                                                                    v42 = uint(v37)
                                                                                    v319 = (v319 << ((v42) & 63))
                                                                                }
                                                                                var v332 uint64 = 790
                                                                                _ = v332
                                                                            }
                                                                        }
                                                                        _ = v319
                                                                    }
                                                                }
                                                                _ = v158
                                                            } else {
                                                                _ = len(v36)
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        var v333 int32 = -2147483648
                                                        var v334 int64 = -9223372036854775808
                                                        _ = v333
                                                        _ = v334
                                                    }
                                                    if ((127 > v35) && false) {
                                                        if true {
                                                            if (((!true) || ((!((^(v35)) <= 42)) && false)) && (42 >= (v38 + (v38 & (v38 / ((v38) | 1)))))) {
                                                                var v335 uint8 = v35
                                                                v336 := GenS28{F0: v3}
                                                                _ = v335
                                                                _ = v336
                                                            } else {
                                                                v337 := GenS45{F0: (v35 + (v35 / ((v35) | 1))), F1: uint16(v35), F2: 1}
                                                                v35 -= (v35 % ((127) | 1))
                                                                _ = v337
                                                            }
                                                            var v338 uint32 = v37
                                                            if len(v36) > 0 {
                                                                v36[0] = -1
                                                            }
                                                            v3 = v3
                                                            if false {
                                                                if len(v36) >= 2 {
                                                                    v36 = v36[1:len(v36)]
                                                                }
                                                            }
                                                            _ = v338
                                                        } else {
                                                            switch fi132 {
                                                            case -1:
                                                                var v339 int8 = (^(v4))
                                                                if true {
                                                                    var v340 uint16 = 900
                                                                    _ = v340
                                                                } else {
                                                                    var v341 float32 = 1.5
                                                                    var v342 uint8 = 2
                                                                    v342 += v35
                                                                    if len(v36) > 0 {
                                                                        v36[0] = fi132
                                                                    }
                                                                    v2 = "foo bar"
                                                                    _ = v341
                                                                    _ = v342
                                                                }
                                                                if ("hello" <= v2) {
                                                                    if true {
                                                                        var v343 bool = (255 <= v35)
                                                                        _ = v343
                                                                    }
                                                                    fmt.Println(v38, v4, v1)
                                                                }
                                                                _ = v339
                                                            case 9223372036854775807:
                                                                v344 := GenS46{F0: (v35 % ((v35) | 1)), F1: v2, F2: 1}
                                                                _ = v344
                                                            }
                                                            fmt.Println(v38, v2)
                                                            v35 = (v35 + (^(v35)))
                                                        }
                                                    }
                                                    v36 = append(v36, (^(fi132)))
                                                }
                                            }
                                            _ = fi132
                                        }
                                        for fi345 := 0; fi345 < 19; fi345++ {
                                            fmt.Println(fi41, v1, v39)
                                            _ = fi345
                                        }
                                    }
                                case 619:
                                    for fi346 := 0; fi346 < 6; fi346++ {
                                        switch v2 {
                                        case "a":
                                            v347 := GenS47{F0: true, F1: v38, F2: -9223372036854775808}
                                            _ = v347
                                        case "hello":
                                            v35 = v35
                                        case "":
                                            v2 = v2
                                        default:
                                            switch ((v2 + "\n\t") + (v2 + "")) {
                                            case "a":
                                                v348 := GenS5{F0: uint64(v37)}
                                                _ = v348
                                            case "hello":
                                                var v349 int = (^(fi41))
                                                _ = v349
                                            case "":
                                                var v350 int16 = 1000
                                                _ = v350
                                            }
                                        }
                                        v42 = v42
                                        if false {
                                            v351 := GenS18{F0: uint8(v3)}
                                            _ = v351
                                        }
                                        _ = fi346
                                    }
                                default:
                                    if (!(v37 != (v37 | v37))) {
                                        if len(v36) > 0 {
                                            v36[0] = 4294967295
                                        }
                                    } else {
                                        v35 += 255
                                        var v352 int32 = 0
                                        for fi353 := 0; fi353 < 18; fi353++ {
                                            v36 = append(v36, (fi41 >> (((v42 + 2)) & 63)))
                                            if (((!false) && true) && (false && false)) {
                                                if (255 > v35) {
                                                    var v354 int16 = 32767
                                                    v355 := GenS48{F0: (v2 + v2), F1: 0, F2: 9223372036854775807}
                                                    _ = v354
                                                    _ = v355
                                                } else {
                                                    if (!(false && (true && true))) {
                                                        v3 += uint64(v38)
                                                    }
                                                }
                                                var v356 float32 = (3.14159 + 1e10)
                                                _ = v356
                                            }
                                            v35 += v35
                                            _ = fi353
                                        }
                                        _ = v352
                                    }
                                }
                            }
                        }
                    case "hello":
                        _ = len(v36)
                    }
                    fmt.Println(v3)
                    v1 = 1.5
                }
                switch fi41 {
                case -36:
                    fmt.Println(v42, v38)
                case 128:
                    v37 = 2147483647
                case 4294967295:
                    for fi357 := 0; fi357 < 16; fi357++ {
                        v2 = v2
                        v42 = v42
                        if false {
                            v358 := GenS49{F0: -1, F1: 2147483647}
                            v3 = (v3 % (((^(v3))) | 1))
                            var v359 uint = (v39 - (v39 - (^(v42))))
                            _ = v358
                            _ = v359
                        } else {
                            v36 = append(v36, fi41)
                            if ((^(v38)) < (v38 % ((v38) | 1))) {
                                v360 := GenS33{F0: (v35 >> ((255) & 63)), F1: 18446744073709551615}
                                _ = v360
                            } else {
                                var v361 int16 = 0
                                _ = v361
                            }
                            if true {
                                if len(v36) > 0 {
                                    v36[0] = -1
                                }
                            }
                        }
                        _ = fi357
                    }
                }
                if false {
                    if false {
                        switch fi41 {
                        case 655:
                            if false {
                                var v362 uint8 = v35
                                _ = v362
                            } else {
                                for fi363 := 0; fi363 < 5; fi363++ {
                                    var v364 string = "foo bar"
                                    var v365 int32 = -682
                                    _ = len(v36)
                                    _ = fi363
                                    _ = v364
                                    _ = v365
                                }
                                _ = len(v36)
                                if len(v36) > 0 {
                                    v36[0] = 32767
                                }
                            }
                        case 1:
                            v38 += uint16(v35)
                        case -281:
                            var v366 uint16 = (v38 & (v38 + (^(v38))))
                            _ = v366
                        }
                        if true {
                            v367 := GenS50{F0: -1, F1: uint(v37), F2: 1}
                            _ = v367
                        }
                        v36 = append(v36, 32768)
                    } else {
                        for fi368 := 0; fi368 < 3; fi368++ {
                            switch -1 {
                            case -1:
                                if ((-2.25 * 0.0) <= 0.1) {
                                    if true {
                                        v38 += (v38 / ((v38) | 1))
                                        var v369 uint64 = v3
                                        _ = v369
                                    } else {
                                        var v370 int32 = 309
                                        _ = v370
                                    }
                                } else {
                                    v36 = append(v36, (-(fi368)))
                                    v371 := GenS51{F0: 32768, F1: 430}
                                    _ = v371
                                }
                            case 32767:
                                var v372 float64 = v1
                                _ = v372
                            case -9223372036854775808:
                                var v373 int64 = 596
                                _ = v373
                            default:
                                v4 = -126
                            }
                            v374 := GenS40{F0: (v37 * (v37 % (((^(v37))) | 1)))}
                            _ = fi368
                            _ = v374
                        }
                    }
                } else {
                    v2 = v2
                    v375 := GenS38{F0: v35}
                    fmt.Println(v38)
                    _ = v375
                }
            }
            if true {
                v376 := GenS51{F0: fi41, F1: -473}
                for fi377 := 0; fi377 < 12; fi377++ {
                    fmt.Println(fi41, v39, v1)
                    if false {
                        v378 := GenS14{F0: v38}
                        v379 := GenS52{F0: v3, F1: 9223372036854775807}
                        _ = v378
                        _ = v379
                    } else {
                        var v380 uint = (v42 / ((v39) | 1))
                        v381 := GenS2{F0: (0 > 42), F1: (1.5 >= 0.1), F2: "a"}
                        v36 = append(v36, 657)
                        _ = v380
                        _ = v381
                    }
                    if (false || false) {
                        var v382 float32 = float32((1e10 / v1))
                        v3 = (v3 >> (((v42 * (^(v42)))) & 63))
                        _ = v382
                    }
                    _ = fi377
                }
                _ = v376
            } else {
                if ((v35 - (v35 / ((v35) | 1))) <= (v35 & uint8(v3))) {
                    v35 += uint8(v38)
                } else {
                    for fi383 := 0; fi383 < 17; fi383++ {
                        if (((true && (!(uint(v37) == (^(v39))))) || (!(!true))) && false) {
                            _ = len(v36)
                        } else {
                            v35 = uint8(v39)
                            var v384 string = "\n\t"
                            _ = v384
                        }
                        switch (fi41 ^ 9223372036854775807) {
                        case 1:
                            var v385 int8 = v4
                            _ = v385
                        case -550:
                            var v386 float64 = v1
                            _ = v386
                        default:
                            if len(v36) >= 2 {
                                v36 = v36[1:len(v36)]
                            }
                        }
                        v387 := GenS24{F0: v37, F1: (!false)}
                        _ = fi383
                        _ = v387
                    }
                }
            }
            _ = fi41
            _ = v42
        }
        v2 = v2
        _ = len(v36)
    } else {
        v388 := GenS21{F0: 32767}
        for fi389 := 0; fi389 < 14; fi389++ {
            _ = len(v36)
            _ = fi389
        }
        _ = v388
    }
    var v390 uint16 = v38
    for fi391 := 0; fi391 < 20; fi391++ {
        _ = len(v36)
        _ = fi391
    }
    v4 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v390:", v390)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v390
    _ = v4
}
