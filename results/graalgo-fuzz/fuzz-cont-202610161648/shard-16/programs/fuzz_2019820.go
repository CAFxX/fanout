package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 string; F1 int64; F2 uint16 }
type GenS10 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS11 struct { F0 int64; F1 int32; F2 uint16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int32; F1 int64 }
type GenS14 struct { F0 int64; F1 int8; F2 uint32 }
type GenS15 struct { F0 uint; F1 int; F2 uint64 }
type GenS16 struct { F0 int64; F1 uint; F2 uint32 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 string; F1 int; F2 uint16 }
type GenS19 struct { F0 uint; F1 uint }
type GenS20 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS21 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS22 struct { F0 string; F1 uint64 }
type GenS23 struct { F0 int8; F1 uint32 }
type GenS24 struct { F0 int64; F1 int }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 int }
func genFn0(p0 uint32) int32 {
    var v4 int = 127
    p0 = (p0 << ((470) & 63))
    v5 := []int{v4, v4, 42, v4}
    switch (v4 % (((v4 - (^(v4)))) | 1)) {
    case 128:
        if len(v5) > 0 {
            v5[0] = (v4 >> ((42) & 63))
        }
    case -2147483648:
        if ((!true) || true) {
            p0 += (p0 & p0)
        }
    default:
        v6 := GenS1{F0: (p0 << ((1) & 63))}
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []int{(v1 / (((v1 >> ((247) & 63))) | 1)), 128, 42, (^(v1))}
    var v3 string = "a"
    fmt.Println(v3)
    switch 32768 {
    case 4294967295:
        switch -684 {
        case 1:
            v7 := GenS2{F0: 737, F1: 61, F2: 42}
            _ = v7
        case 32768:
            var v8 uint = 42
            _ = v8
        default:
            v1 += (v1 | (v1 - 128))
        }
    case 681:
        switch (("\n\t" + (v3 + (((v3 + (v3 + v3)) + v3) + (v3 + (v3 + (v3 + (v3 + v3))))))) + v3) {
        case "\n\t":
            var v9 int8 = 127
            _ = v9
        case "hello":
            v3 = "hello"
        case "foo bar":
            v10 := GenS2{F0: 1, F1: 2147483648, F2: 1}
            _ = v10
        }
    }
    switch (^(v1)) {
    case 32767:
        v11 := GenS2{F0: 0, F1: 2147483647, F2: -246}
        _ = v11
    case 0:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch -120 {
        case 32767:
            var v12 uint32 = 1
            _ = v12
        case -989:
            for fi13 := 0; fi13 < 10; fi13++ {
                v2 = append(v2, v1)
                var v14 bool = true
                v2 = append(v2, (v1 << ((42) & 63)))
                _ = fi13
                _ = v14
            }
        case -2147483648:
            if true {
                fmt.Println(v3, v3)
            } else {
                v1 = (v1 << ((308) & 63))
                v1 -= 2147483647
            }
        }
    }
    var v15 int16 = 42
    if len(v2) > 0 {
        v2[0] = v1
    }
    if false {
        fmt.Println(v3)
        for fi16 := 0; fi16 < 1; fi16++ {
            switch 2147483647 {
            case 4294967295:
                v17 := GenS3{F0: int16(v1)}
                _ = v17
            case -97:
                var v18 int8 = 127
                _ = v18
            case 92:
                if false {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        switch "hello" {
                        case "\n\t":
                            if ((v1 / (((fi16 & fi19)) | 1)) < 537) {
                                var v20 int16 = 1000
                                var v21 bool = (((!true) && (-299 < v1)) || (true || (((v3 + v3) <= (v3 + "hello")) && (1000 <= 1))))
                                var v22 int16 = v20
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                if false {
                                    switch (("hello" + "") + "\n\t") {
                                    case "foo bar":
                                        v2 = append(v2, 127)
                                    case "":
                                        v15 = (-(v15))
                                        v1 += int(v15)
                                    }
                                    switch v3 {
                                    case "":
                                        switch v1 {
                                        case -302:
                                            var v23 int32 = 42
                                            _ = v23
                                        case 2147483647:
                                            switch fi19 {
                                            case 1:
                                                if false {
                                                    v3 = v3
                                                    if true {
                                                        switch (v1 + v1) {
                                                        case 42:
                                                            v1 = (fi16 & v1)
                                                        case 245:
                                                            if false {
                                                                v2 = append(v2, (v1 % (((fi16 >> ((264) & 63))) | 1)))
                                                                var v24 int = 839
                                                                _ = v24
                                                            } else {
                                                                fmt.Println(fi19)
                                                            }
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                        v1 -= fi19
                                                        if (!(false && (true && true))) {
                                                            v3 = ""
                                                            if false {
                                                                v15 = v15
                                                            }
                                                            var v25 uint64 = 1
                                                            _ = v25
                                                        }
                                                        fmt.Println(v15, v3)
                                                    } else {
                                                        v3 = "foo bar"
                                                        if (-2.25 < 0.1) {
                                                            fmt.Println(v1, fi16, v3)
                                                            _ = len(v2)
                                                            v26 := GenS3{F0: -1}
                                                            fmt.Println(fi16, v3, v3)
                                                            _ = v26
                                                        }
                                                    }
                                                    var v27 uint64 = 18446744073709551615
                                                    _ = v27
                                                } else {
                                                    var v28 float32 = -2.25
                                                    v28 += 1e10
                                                    _ = v28
                                                }
                                                v29 := GenS4{F0: -221, F1: v3}
                                                _ = v29
                                            case 32768:
                                                fmt.Println(fi16)
                                            default:
                                                if (v3 <= v3) {
                                                    if false {
                                                        var v30 uint = 497
                                                        switch fi19 {
                                                        case 90:
                                                            if ((v3 + v3) == v3) {
                                                                var v31 float32 = (1e10 + float32(float64(0.1)))
                                                                v2 = append(v2, (fi19 / (((fi19 << ((v30) & 63))) | 1)))
                                                                var v32 uint8 = 255
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                        case 0:
                                                            var v33 int32 = genFn0(2147483647)
                                                            _ = v33
                                                        }
                                                        v2 = append(v2, 127)
                                                        _ = v30
                                                    } else {
                                                        var v34 string = (v3 + v3)
                                                        v34 = ((v34 + v34) + ("hello" + "foo bar"))
                                                        _ = v34
                                                    }
                                                    if true {
                                                        v35 := GenS5{F0: 128, F1: 4294967295}
                                                        v36 := GenS2{F0: 310, F1: 2, F2: -129}
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        v37 := GenS4{F0: -1, F1: v3}
                                                        _ = v37
                                                    }
                                                } else {
                                                    if (false || false) {
                                                        v3 = ("hello" + ("a" + v3))
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v38 int = fi16
                                                        _ = v38
                                                    } else {
                                                        if (2147483647 == -865) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        } else {
                                                            var v39 int64 = 803
                                                            var v40 string = v3
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi16 % ((9223372036854775807) | 1))
                                                    }
                                                }
                                            }
                                        case 32767:
                                            var v41 uint8 = 21
                                            _ = v41
                                        }
                                    default:
                                        switch v1 {
                                        case -1:
                                            var v42 int64 = 203
                                            _ = v42
                                        case -803:
                                            v1 -= -1
                                            switch "a" {
                                            case "a":
                                                var v43 int8 = 42
                                                _ = v43
                                            }
                                        case 4294967295:
                                            var v44 int16 = 1
                                            _ = v44
                                        }
                                    }
                                    if (!false) {
                                        if (32767 <= 42) {
                                            var v45 uint16 = 1
                                            v3 = ""
                                            if (1e10 > 1e10) {
                                                if (0.0 == -2.25) {
                                                    var v46 string = "hello"
                                                    fmt.Println(fi19, v15, v1)
                                                    v47 := GenS6{F0: v45, F1: 127, F2: 120}
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    if true {
                                                        v48 := GenS6{F0: (v45 * v45), F1: 1, F2: 1}
                                                        v15 = (v15 + 156)
                                                        v49 := GenS1{F0: 2147483647}
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = v1
                                                }
                                                v45 += 480
                                            }
                                            _ = v45
                                        } else {
                                            var v50 bool = (-1 <= -1)
                                            _ = v50
                                        }
                                        if (32768 < fi16) {
                                            var v51 string = ((v3 + ("hello" + ("foo bar" + "\n\t"))) + (("foo bar" + "a") + v3))
                                            _ = v51
                                        } else {
                                            var v52 uint16 = 42
                                            if true {
                                                switch -1 {
                                                case 20:
                                                    fmt.Println(fi19, v3)
                                                case 32768:
                                                    var v53 int8 = 0
                                                    _ = v53
                                                default:
                                                    v15 = (v15 | v15)
                                                    v2 = append(v2, (fi16 + (v1 | -978)))
                                                    switch "a" {
                                                    case "a":
                                                        if false {
                                                            var v54 int64 = -1
                                                            fmt.Println(v52, v15, v3)
                                                            var v55 int32 = -2147483648
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    default:
                                                        var v56 uint32 = 2147483648
                                                        if true {
                                                            var v57 bool = false
                                                            v56 -= (v56 ^ (v56 & 2147483648))
                                                            _ = v57
                                                        }
                                                        if (v56 > v56) {
                                                            v3 = v3
                                                        }
                                                        _ = v56
                                                    }
                                                }
                                            } else {
                                                v1 = 128
                                            }
                                            v58 := GenS4{F0: 1, F1: v3}
                                            _ = v52
                                            _ = v58
                                        }
                                    }
                                }
                                v1 += (fi19 + fi16)
                                v59 := GenS2{F0: 65535, F1: 1, F2: 42}
                                _ = v59
                            }
                        case "a":
                            var v60 int = -2147483648
                            _ = v60
                        case "hello":
                            fmt.Println(v15, v15, fi19)
                        }
                        if (209 > 0) {
                            fmt.Println(v3, fi16)
                            v15 = v15
                            v15 -= v15
                        } else {
                            if false {
                                v1 = (fi19 % ((42) | 1))
                            }
                            var v61 uint8 = 128
                            switch v1 {
                            case -9223372036854775808:
                                v62 := GenS4{F0: 431, F1: ""}
                                _ = v62
                            case 2147483647:
                                fmt.Println(v1)
                            case -599:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = v61
                        }
                        _ = fi19
                    }
                    var v63 int8 = -63
                    if false {
                        if (!false) {
                            v3 = "\n\t"
                        } else {
                            if true {
                                v15 -= (v15 + -1)
                                v64 := GenS4{F0: -1, F1: ((("a" + v3) + "hello") + ((v3 + v3) + v3))}
                                _ = v64
                            }
                            v65 := GenS3{F0: (v15 + (v15 << ((581) & 63)))}
                            _ = v65
                        }
                        switch "hello" {
                        case "\n\t":
                            for fi66 := 0; fi66 < 2; fi66++ {
                                v2 = append(v2, (fi16 % ((-55) | 1)))
                                switch -9223372036854775808 {
                                case 42:
                                    if false {
                                        if (206 < 563) {
                                            switch (fi66 + (fi66 >> ((42) & 63))) {
                                            case 9223372036854775807:
                                                _ = len(v2)
                                            }
                                            v3 = (v3 + v3)
                                            var v67 float64 = 1.5
                                            _ = v67
                                        }
                                        var v68 string = ((v3 + "hello") + v3)
                                        var v69 string = v68
                                        _ = v68
                                        _ = v69
                                    }
                                case 766:
                                    v63 = 42
                                case -107:
                                    fmt.Println(fi66)
                                }
                                _ = fi66
                            }
                        case "foo bar":
                            if false {
                                for fi70 := 0; fi70 < 8; fi70++ {
                                    _ = len(v2)
                                    v1 = (-(fi70))
                                    var v71 uint8 = 0
                                    _ = fi70
                                    _ = v71
                                }
                                v15 -= (v15 << ((65535) & 63))
                            }
                        }
                        v1 -= fi16
                    } else {
                        var v72 int64 = 469
                        for fi73 := 0; fi73 < 13; fi73++ {
                            _ = len(v2)
                            switch "hello" {
                            case "":
                                switch "hello" {
                                case "hello":
                                    fmt.Println(fi73)
                                default:
                                    v1 += (v1 * 1)
                                }
                            case "a":
                                switch v3 {
                                case "a":
                                    if ((v63 / ((v63) | 1)) != (v63 / ((0) | 1))) {
                                        fmt.Println(v3, v63, v63)
                                        var v74 int8 = v63
                                        var v75 int = (fi16 + (fi16 | fi73))
                                        v74 -= (v63 / ((-26) | 1))
                                        _ = v74
                                        _ = v75
                                    } else {
                                        if (true && (!((true || false) || (!(!(!(872 >= 18446744073709551615))))))) {
                                            v1 = -437
                                            var v76 int64 = (v72 + (v72 << ((190) & 63)))
                                            if len(v2) > 0 {
                                                v2[0] = fi16
                                            }
                                            if false {
                                                v77 := GenS5{F0: 255, F1: 1}
                                                switch (v1 * -405) {
                                                case 128:
                                                    fmt.Println(fi73, v72, v15)
                                                default:
                                                    v1 = 127
                                                    if (0 <= 65535) {
                                                        fmt.Println(v72)
                                                    } else {
                                                        v78 := GenS7{F0: -1, F1: 2}
                                                        _ = v78
                                                    }
                                                }
                                                v79 := GenS8{F0: -490, F1: 2147483648, F2: 2}
                                                _ = v77
                                                _ = v79
                                            }
                                            fmt.Println(v3)
                                            switch v1 {
                                            case 9223372036854775807:
                                                switch "hello" {
                                                case "hello":
                                                    switch 4294967295 {
                                                    case 127:
                                                        switch (fi73 >> ((65535) & 63)) {
                                                        case 357:
                                                            if ((v15 > v15) || ((false || ((1e10 * (((((0.1 + float64(1e10)) - 1.5) * 0.0) - 0.1) + (-2.25 * 3.14159))) > (((-2.25 * ((((0.0 * ((0.0 / 0.1) * (1.5 * 1e10))) / ((1.5 * -2.25) / 1.0)) * -2.25) * float64(3.14159))) * (-2.25 / 3.14159)) - 1.5))) && false)) {
                                                                v80 := GenS2{F0: 32767, F1: 2147483647, F2: v15}
                                                                _ = v80
                                                            } else {
                                                                switch (v3 + "") {
                                                                case "foo bar":
                                                                    var v81 int = 32768
                                                                    _ = v81
                                                                case "":
                                                                    fmt.Println(v15)
                                                                case "\n\t":
                                                                    if false {
                                                                        var v82 int = 486
                                                                        v1 += v82
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            _ = len(v2)
                                                                        case "hello":
                                                                            v82 = -1
                                                                        case "":
                                                                            if (!false) {
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    v2 = append(v2, (-(v82)))
                                                                                }
                                                                                v2 = append(v2, 32768)
                                                                                if (!true) {
                                                                                    v83 := GenS9{F0: ((v3 + (v3 + ("hello" + v3))) + v3), F1: int64(v63), F2: 0}
                                                                                    if (386 >= 240) {
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        v84 := GenS1{F0: 2}
                                                                                        _ = v84
                                                                                    }
                                                                                    if false {
                                                                                        v15 = -871
                                                                                    }
                                                                                    _ = v83
                                                                                } else {
                                                                                    v2 = append(v2, 4294967295)
                                                                                    switch (^(fi73)) {
                                                                                    case -950:
                                                                                        v72 = 197
                                                                                    case -726:
                                                                                        var v85 uint16 = 210
                                                                                        _ = v85
                                                                                    default:
                                                                                        v86 := GenS10{F0: 42, F1: 65535, F2: 1}
                                                                                        _ = v86
                                                                                    }
                                                                                    if false {
                                                                                        if false {
                                                                                            if (v3 < v3) {
                                                                                                switch 32767 {
                                                                                                case 127:
                                                                                                    var v87 bool = (1 < 153)
                                                                                                    _ = v87
                                                                                                case 115:
                                                                                                    var v88 int16 = 42
                                                                                                    _ = v88
                                                                                                default:
                                                                                                    var v89 bool = (!(false || false))
                                                                                                    _ = v89
                                                                                                }
                                                                                                switch v3 {
                                                                                                case "foo bar":
                                                                                                    v90 := GenS5{F0: 2, F1: 2}
                                                                                                    _ = v90
                                                                                                case "":
                                                                                                    v82 += -1
                                                                                                    var v91 int8 = int8(v15)
                                                                                                    var v92 int32 = 0
                                                                                                    _ = v91
                                                                                                    _ = v92
                                                                                                default:
                                                                                                    if false {
                                                                                                        var v93 int64 = (v76 & (-(v72)))
                                                                                                        v15 -= v15
                                                                                                        fmt.Println(v63, v3, v63)
                                                                                                        v15 += v15
                                                                                                        _ = v93
                                                                                                    } else {
                                                                                                        if (!true) {
                                                                                                            v94 := GenS1{F0: 872}
                                                                                                            fmt.Println(v1, v15)
                                                                                                            v95 := GenS11{F0: v76, F1: 1, F2: 65535}
                                                                                                            _ = v94
                                                                                                            _ = v95
                                                                                                        } else {
                                                                                                            v96 := GenS2{F0: 65535, F1: 52, F2: 633}
                                                                                                            var v97 int64 = v72
                                                                                                            var v98 bool = (((1e10 + 3.14159) - float32(0.1)) > 0.1)
                                                                                                            _ = v96
                                                                                                            _ = v97
                                                                                                            _ = v98
                                                                                                        }
                                                                                                        var v99 uint32 = 0
                                                                                                        _ = v99
                                                                                                    }
                                                                                                }
                                                                                                switch (fi73 % ((-218) | 1)) {
                                                                                                case -2147483648:
                                                                                                    v76 = -434
                                                                                                case 132:
                                                                                                    v76 -= v76
                                                                                                default:
                                                                                                    var v100 int64 = (v76 % ((v72) | 1))
                                                                                                    _ = v100
                                                                                                }
                                                                                                v101 := GenS12{F0: v3}
                                                                                                var v102 uint8 = 255
                                                                                                _ = v101
                                                                                                _ = v102
                                                                                            }
                                                                                            v103 := GenS10{F0: 208, F1: 937, F2: 127}
                                                                                            var v104 uint32 = 2147483648
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                        }
                                                                                        var v105 uint32 = 2147483647
                                                                                        _ = v105
                                                                                    } else {
                                                                                        switch fi73 {
                                                                                        case 128:
                                                                                            v106 := GenS12{F0: v3}
                                                                                            _ = v106
                                                                                        case 4294967295:
                                                                                            var v107 uint64 = 1
                                                                                            _ = v107
                                                                                        }
                                                                                        v72 = (v76 - int64(fi73))
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v108 bool = (42 == 746)
                                                                        v2 = append(v2, v82)
                                                                        switch (("hello" + (v3 + ("a" + v3))) + v3) {
                                                                        case "a":
                                                                            var v109 uint32 = 2147483647
                                                                            _ = v109
                                                                        case "foo bar":
                                                                            v2 = append(v2, v82)
                                                                        }
                                                                        fmt.Println(v72)
                                                                        _ = v108
                                                                        _ = v82
                                                                    } else {
                                                                        switch ((("" + (v3 + v3)) + "") + v3) {
                                                                        case "hello":
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        default:
                                                                            if (1 >= 0) {
                                                                                var v110 uint32 = 4294967295
                                                                                fmt.Println(v76, v72, v1)
                                                                                _ = v110
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = len(v2)
                                                                if false {
                                                                    var v111 uint16 = 65535
                                                                    v2 = append(v2, fi73)
                                                                    var v112 int32 = genFn0(2)
                                                                    v112 = v112
                                                                    switch (("hello" + ((v3 + v3) + "hello")) + v3) {
                                                                    case "\n\t":
                                                                        switch fi73 {
                                                                        case 42:
                                                                            v63 = v63
                                                                        }
                                                                    case "a":
                                                                        switch (v3 + "\n\t") {
                                                                        case "hello":
                                                                            fmt.Println(v72)
                                                                        case "a":
                                                                            v113 := GenS11{F0: v76, F1: -916, F2: v111}
                                                                            _ = v113
                                                                        case "":
                                                                            switch v3 {
                                                                            case "hello":
                                                                                fmt.Println(v112, v63)
                                                                            case "":
                                                                                switch ("hello" + ("\n\t" + ("a" + v3))) {
                                                                                case "\n\t":
                                                                                    fmt.Println(v72, v1, v63)
                                                                                case "hello":
                                                                                    v2 = append(v2, 42)
                                                                                case "":
                                                                                    if (906 == 901) {
                                                                                        v114 := GenS1{F0: 259}
                                                                                        _ = v114
                                                                                    } else {
                                                                                        if false {
                                                                                            var v115 int16 = v15
                                                                                            switch v3 {
                                                                                            case "\n\t":
                                                                                                var v116 int64 = v72
                                                                                                _ = v116
                                                                                            default:
                                                                                                v117 := GenS10{F0: v111, F1: 255, F2: (v63 >> ((0) & 63))}
                                                                                                _ = v117
                                                                                            }
                                                                                            v118 := GenS13{F0: -1, F1: 675}
                                                                                            if ((v112 % (((v112 * (v112 * (-(v112))))) | 1)) > genFn0(2)) {
                                                                                                switch (v1 << ((255) & 63)) {
                                                                                                case 1:
                                                                                                    fmt.Println(v76)
                                                                                                default:
                                                                                                    if true {
                                                                                                        v72 = 1
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v115
                                                                                            _ = v118
                                                                                        } else {
                                                                                            v72 -= v76
                                                                                        }
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = int(v15)
                                                                                        }
                                                                                        v72 = (-(v76))
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                switch fi16 {
                                                                                case 127:
                                                                                    var v119 bool = true
                                                                                    _ = v119
                                                                                case -1:
                                                                                    var v120 int32 = v112
                                                                                    _ = v120
                                                                                case 228:
                                                                                    var v121 int64 = v72
                                                                                    _ = v121
                                                                                }
                                                                                v112 -= 42
                                                                                v76 -= (v76 | (v72 | v72))
                                                                                v111 -= (^(v111))
                                                                            default:
                                                                                switch 32767 {
                                                                                case 32768:
                                                                                    _ = len(v2)
                                                                                case 32767:
                                                                                    v111 += (v111 + 42)
                                                                                case 171:
                                                                                    v15 -= v15
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v122 int64 = v72
                                                                            _ = v122
                                                                        }
                                                                    case "hello":
                                                                        v123 := GenS14{F0: (v72 & 0), F1: v63, F2: 310}
                                                                        _ = v123
                                                                    default:
                                                                        fmt.Println(v111, fi16, fi73)
                                                                    }
                                                                    _ = v111
                                                                    _ = v112
                                                                } else {
                                                                    if (18446744073709551615 <= 4294967295) {
                                                                        switch "" {
                                                                        case "":
                                                                            if true {
                                                                                fmt.Println(v3, fi16)
                                                                                v63 = v63
                                                                                v3 = ""
                                                                            }
                                                                        case "a":
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (^(fi16))
                                                                            }
                                                                            v63 += (v63 / ((v63) | 1))
                                                                            if ((fi16 | v1) < int(v15)) {
                                                                                v1 += fi16
                                                                            }
                                                                        case "\n\t":
                                                                            if (-628 == -2147483648) {
                                                                                switch (v3 + v3) {
                                                                                case "hello":
                                                                                    if (1 > v15) {
                                                                                        var v124 int8 = 127
                                                                                        var v125 float32 = ((1e10 * 3.14159) - float32(float64(0.0)))
                                                                                        switch v1 {
                                                                                        case 511:
                                                                                            var v126 uint16 = 932
                                                                                            _ = v126
                                                                                        case -75:
                                                                                            var v127 string = ((v3 + (v3 + v3)) + "")
                                                                                            _ = v127
                                                                                        case 945:
                                                                                            switch fi16 {
                                                                                            case 42:
                                                                                                if (42 < 42) {
                                                                                                    if true {
                                                                                                        var v128 bool = (false || false)
                                                                                                        v129 := GenS15{F0: 121, F1: int(v63), F2: 462}
                                                                                                        _ = v128
                                                                                                        _ = v129
                                                                                                    }
                                                                                                    switch v3 {
                                                                                                    case "hello":
                                                                                                        fmt.Println(v72)
                                                                                                    case "\n\t":
                                                                                                        var v130 float64 = -2.25
                                                                                                        var v131 int64 = (v76 & 197)
                                                                                                        switch ((v3 + "hello") + "foo bar") {
                                                                                                        case "\n\t":
                                                                                                            v76 = (-(v76))
                                                                                                        default:
                                                                                                            var v132 int = v1
                                                                                                            _ = v132
                                                                                                        }
                                                                                                        _ = v130
                                                                                                        _ = v131
                                                                                                    default:
                                                                                                        var v133 float64 = 1.5
                                                                                                        _ = v133
                                                                                                    }
                                                                                                    if (((true || false) && ((264 <= 139) || (int(v63) >= fi16))) || ((fi73 - (v1 & fi73)) > 127)) {
                                                                                                        v3 = v3
                                                                                                        switch v3 {
                                                                                                        case "a":
                                                                                                            fmt.Println(v3)
                                                                                                        case "foo bar":
                                                                                                            v134 := GenS16{F0: 11, F1: 4294967295, F2: 628}
                                                                                                            _ = v134
                                                                                                        default:
                                                                                                            v15 = v15
                                                                                                        }
                                                                                                        _ = len(v2)
                                                                                                        v1 -= fi16
                                                                                                        var v135 int = fi16
                                                                                                        var v136 float64 = 1e10
                                                                                                        var v137 int = -672
                                                                                                        _ = v135
                                                                                                        _ = v136
                                                                                                        _ = v137
                                                                                                    } else {
                                                                                                        v15 += v15
                                                                                                        var v138 uint = 255
                                                                                                        _ = v138
                                                                                                    }
                                                                                                } else {
                                                                                                    v139 := GenS4{F0: 9223372036854775807, F1: v3}
                                                                                                    v140 := GenS4{F0: v72, F1: "a"}
                                                                                                    _ = v139
                                                                                                    _ = v140
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                    }
                                                                                    v141 := GenS9{F0: v3, F1: v72, F2: 2}
                                                                                    if (!true) {
                                                                                        fmt.Println(v63, v1)
                                                                                        v72 += -393
                                                                                    }
                                                                                    _ = v141
                                                                                case "foo bar":
                                                                                    fmt.Println(v63, v3)
                                                                                case "a":
                                                                                    v3 = v3
                                                                                }
                                                                            } else {
                                                                                fmt.Println(fi16, fi16)
                                                                                v15 = (v15 * -32768)
                                                                                v2 = append(v2, -1)
                                                                            }
                                                                        default:
                                                                            switch fi16 {
                                                                            case 0:
                                                                                var v142 uint16 = 0
                                                                                _ = v142
                                                                            case 42:
                                                                                v63 -= (v63 << ((18446744073709551615) & 63))
                                                                            case 220:
                                                                                switch fi73 {
                                                                                case 29:
                                                                                    v15 += (v15 * v15)
                                                                                case 32768:
                                                                                    v3 = v3
                                                                                case 42:
                                                                                    v143 := GenS9{F0: (v3 + (v3 + v3)), F1: v72, F2: 65535}
                                                                                    _ = v143
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    v76 = 824
                                                                                    if (744 < 65535) {
                                                                                        v3 = "a"
                                                                                        fmt.Println(v1)
                                                                                        v63 = v63
                                                                                    } else {
                                                                                        _ = len(v2)
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v72 = 0
                                                                }
                                                                v72 = v72
                                                            }
                                                        case 247:
                                                            switch (v3 + "\n\t") {
                                                            case "foo bar":
                                                                var v144 uint16 = 0
                                                                if false {
                                                                    _ = len(v2)
                                                                }
                                                                v3 = v3
                                                                _ = v144
                                                            case "":
                                                                var v145 bool = (912 != 1)
                                                                _ = v145
                                                            default:
                                                                if ((true || false) || (2147483647 != 2147483647)) {
                                                                    var v146 int8 = (v63 / ((int8(v72)) | 1))
                                                                    _ = v146
                                                                }
                                                            }
                                                        case 2147483647:
                                                            v63 += v63
                                                        }
                                                    case -263:
                                                        v72 = v76
                                                    case 128:
                                                        v72 -= (-(v76))
                                                    }
                                                case "foo bar":
                                                    if (!true) {
                                                        var v147 int8 = -1
                                                        _ = v147
                                                    }
                                                case "\n\t":
                                                    switch "a" {
                                                    case "foo bar":
                                                        v148 := GenS4{F0: (v72 - v76), F1: "hello"}
                                                        _ = v148
                                                    case "hello":
                                                        v3 = "hello"
                                                    case "\n\t":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        if (!false) {
                                                            v149 := GenS16{F0: int64(v63), F1: 1, F2: 2}
                                                            var v150 int8 = 1
                                                            _ = v149
                                                            _ = v150
                                                        }
                                                    }
                                                default:
                                                    v1 -= 0
                                                }
                                            }
                                            _ = v76
                                        }
                                        switch (^(fi16)) {
                                        case -897:
                                            v151 := GenS17{F0: genFn0(961)}
                                            _ = v151
                                        case 42:
                                            v1 += 1
                                        case -9223372036854775808:
                                            fmt.Println(v15, v1, v3)
                                        default:
                                            var v152 uint64 = 4294967295
                                            _ = v152
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            v3 = ("a" + ("hello" + "foo bar"))
                                        case "a":
                                            var v153 int = fi73
                                            _ = v153
                                        case "hello":
                                            var v154 float64 = float64((0.1 * (1.5 / ((1e10 / 1e10) - (-2.25 + (0.0 * 0.0))))))
                                            _ = v154
                                        }
                                    }
                                case "\n\t":
                                    if len(v2) > 0 {
                                        v2[0] = 128
                                    }
                                }
                                switch fi73 {
                                case -9223372036854775808:
                                    if false {
                                        v155 := GenS18{F0: "hello", F1: (fi73 * fi73), F2: 0}
                                        if len(v2) > 0 {
                                            v2[0] = (fi16 << ((18446744073709551615) & 63))
                                        }
                                        _ = v155
                                    } else {
                                        v156 := GenS19{F0: 981, F1: 18446744073709551615}
                                        _ = v156
                                    }
                                default:
                                    var v157 uint8 = 127
                                    var v158 int32 = 2147483647
                                    _ = v157
                                    _ = v158
                                }
                                if true {
                                    var v159 int16 = -1
                                    v160 := GenS2{F0: 500, F1: 1, F2: -1}
                                    var v161 int32 = int32(v15)
                                    _ = v159
                                    _ = v160
                                    _ = v161
                                } else {
                                    v72 += -1
                                    v162 := GenS13{F0: 0, F1: v72}
                                    _ = v162
                                }
                                v3 = "foo bar"
                                v2 = append(v2, (fi16 | (^(v1))))
                            case "foo bar":
                                var v163 string = ((v3 + ("hello" + (v3 + "\n\t"))) + (("hello" + v3) + "a"))
                                _ = v163
                            default:
                                if ((v15 << ((255) & 63)) >= v15) {
                                    switch fi73 {
                                    case 1:
                                        var v164 int16 = (v15 & (-(v15)))
                                        _ = v164
                                    case -554:
                                        fmt.Println(fi73)
                                    case 42:
                                        switch "hello" {
                                        case "a":
                                            v165 := GenS20{F0: 491, F1: 2, F2: 0}
                                            _ = v165
                                        case "":
                                            var v166 int32 = 1
                                            _ = v166
                                        }
                                    default:
                                        var v167 float32 = 1e10
                                        _ = v167
                                    }
                                    if true {
                                        v1 -= v1
                                        if len(v2) > 0 {
                                            v2[0] = fi73
                                        }
                                    } else {
                                        var v168 int32 = -447
                                        _ = v168
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = fi16
                                    }
                                } else {
                                    fmt.Println(fi16)
                                }
                            }
                            _ = fi73
                        }
                        _ = v72
                    }
                    _ = v63
                }
            default:
                v169 := GenS15{F0: 42, F1: v1, F2: 42}
                _ = v169
            }
            if (false || true) {
                var v170 uint64 = 955
                v2 = append(v2, v1)
                switch (fi16 >> ((65535) & 63)) {
                case 32768:
                    if (true && ((!(-1 < (v15 & v15))) && false)) {
                        fmt.Println(fi16, v15, fi16)
                        switch v1 {
                        case 1:
                            for fi171 := 0; fi171 < 19; fi171++ {
                                v172 := GenS21{F0: 1, F1: (v15 | (v15 - -1)), F2: 875}
                                if len(v2) > 0 {
                                    v2[0] = -210
                                }
                                v173 := GenS20{F0: 2, F1: (v170 % (((v170 << ((485) & 63))) | 1)), F2: v170}
                                _ = fi171
                                _ = v172
                                _ = v173
                            }
                        case 32767:
                            fmt.Println(v3, v1)
                        }
                        v174 := GenS22{F0: (v3 + v3), F1: (v170 << ((2) & 63))}
                        _ = v174
                    } else {
                        for fi175 := 0; fi175 < 6; fi175++ {
                            var v176 int64 = 0
                            v176 -= (v176 / ((v176) | 1))
                            if len(v2) > 0 {
                                v2[0] = (fi175 + 127)
                            }
                            _ = fi175
                            _ = v176
                        }
                        v1 -= (v1 + -442)
                    }
                case 2147483647:
                    v177 := GenS23{F0: 127, F1: 2147483647}
                    _ = v177
                case -89:
                    switch v3 {
                    case "a":
                        switch (fi16 / (((v1 >> ((1) & 63))) | 1)) {
                        case 163:
                            v178 := GenS3{F0: int16(v1)}
                            _ = v178
                        case 9223372036854775807:
                            v1 = (v1 * (^(fi16)))
                        }
                    case "hello":
                        if true {
                            _ = len(v2)
                            for fi179 := 0; fi179 < 13; fi179++ {
                                v3 = v3
                                _ = fi179
                            }
                            for fi180 := 0; fi180 < 12; fi180++ {
                                fmt.Println(fi180, v3, fi180)
                                _ = fi180
                            }
                        }
                    case "foo bar":
                        fmt.Println(fi16)
                    default:
                        if false {
                            fmt.Println(v1, v170)
                        }
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = (fi16 % (((-(fi16))) | 1))
                    }
                }
                _ = v170
            } else {
                for fi181 := 0; fi181 < 2; fi181++ {
                    v182 := GenS24{F0: 9223372036854775807, F1: (v1 + 48)}
                    v3 = v3
                    var v183 uint32 = 0
                    if true {
                        v3 = v3
                    } else {
                        v184 := GenS25{F0: 42}
                        _ = v184
                    }
                    _ = fi181
                    _ = v182
                    _ = v183
                }
                v1 = 0
                for fi185 := 0; fi185 < 19; fi185++ {
                    var v186 uint16 = 65535
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                    _ = len(v2)
                    _ = fi185
                    _ = v186
                }
            }
            _ = fi16
        }
    }
    if true {
        fmt.Println(v1)
        fmt.Println(v1)
    }
    var v187 int = v1
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 997
    }
    var v188 float64 = 1e10
    if (563 == 0) {
        for fi189 := 0; fi189 < 2; fi189++ {
            v190 := GenS6{F0: 659, F1: 1, F2: -4}
            var v191 uint16 = 280
            var v192 int32 = 41
            _ = fi189
            _ = v190
            _ = v191
            _ = v192
        }
    } else {
        fmt.Println(v187)
        if (128 < 35) {
            v193 := GenS20{F0: 65535, F1: 4294967295, F2: 18446744073709551615}
            _ = v193
        } else {
            v15 -= (v15 / (((v15 + (-(v15)))) | 1))
            fmt.Println(v187)
        }
    }
    v15 -= (-(v15))
    v188 -= float64(1.5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (v188 <= 3.14159) {
        var v194 uint32 = 2
        for fi195 := 0; fi195 < 14; fi195++ {
            switch v187 {
            case 128:
                if (false || (v188 > (((v188 * (v188 / v188)) * v188) + v188))) {
                    v2 = append(v2, v187)
                    v196 := GenS24{F0: 42, F1: (v187 & (-(v1)))}
                    if true {
                        v15 += (v15 >> ((402) & 63))
                        v15 += v15
                    }
                    _ = v196
                }
            case 42:
                v197 := GenS26{F0: 4294967295}
                _ = v197
            case -1:
                fmt.Println(v1)
            }
            _ = fi195
        }
        _ = v194
    } else {
        v187 = 2147483647
        var v198 int16 = (v15 | (^(v15)))
        _ = v198
    }
    var v199 int32 = 2147483647
    for fi200 := 0; fi200 < 2; fi200++ {
        switch -980 {
        case -939:
            var v201 string = v3
            _ = v201
        }
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        var v202 bool = ((!(!(((v199 ^ v199) != -1) || false))) || true)
        _ = fi200
        _ = v202
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v187:", v187)
    fmt.Println("v188:", v188)
    fmt.Println("v199:", v199)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v187
    _ = v188
    _ = v199
    _ = v2
    _ = v3
}
