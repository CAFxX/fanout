package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int; F1 uint; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint8; F1 uint }
type GenS9 struct { F0 int16; F1 string; F2 uint64 }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 int8; F1 string }
type GenS12 struct { F0 int32; F1 uint64; F2 int32 }
type GenS13 struct { F0 int32; F1 int32; F2 bool }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 string; F1 int }
type GenS16 struct { F0 uint16; F1 int64 }
type GenS17 struct { F0 bool; F1 uint64 }
type GenS18 struct { F0 uint32; F1 uint64 }
type GenS19 struct { F0 string }
type GenS20 struct { F0 int; F1 int16 }
type GenS21 struct { F0 int64; F1 uint64; F2 int8 }
type GenS22 struct { F0 uint16; F1 uint32 }
type GenS23 struct { F0 int; F1 int16; F2 uint8 }
func genFn0() int8 {
    v4 := []int{128, 4294967295, -59, 698}
    v4 = append(v4, 42)
    _ = v4
    return 100
}
func genFn1(p0 uint16, p1 int8, p2 uint16) int16 {
    p1 = genFn0()
    if false {
        v5 := []int{32768, -142, 127, 127}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int16 = 0
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int = 1
            _ = fi8
            _ = v9
        }
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v7
    }
    var v10 int16 = -32768
    v11 := GenS1{F0: 71, F1: 750}
    v12 := v3[""]
    v13 := []int{v12, (v12 / ((v12) | 1)), (v12 & v12)}
    v14 := GenS1{F0: v12, F1: (v12 - v12)}
    v15 := GenS1{F0: -1, F1: v12}
    v16 := v2[v12]
    v16 = (v16 - 42)
    switch v12 {
    case 2147483647:
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS1{F0: (fi17 & v12), F1: 127}
            fmt.Println(fi17, v16, fi17)
            switch ("\n\t" + "\n\t") {
            case "hello":
                for fi19 := 0; fi19 < 12; fi19++ {
                    v10 = (v10 + v10)
                    _ = fi19
                }
            case "":
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 uint16 = 65535
                    _ = fi20
                    _ = v21
                }
            default:
                if true {
                    v12 += fi17
                    v12 -= (fi17 / (((fi17 * v12)) | 1))
                }
            }
            _ = fi17
            _ = v18
        }
    default:
        switch ("foo bar" + "foo bar") {
        case "a":
            v22 := GenS1{F0: (-(v12)), F1: v12}
            _ = v22
        case "hello":
            v23 := GenS1{F0: v12, F1: (^(v12))}
            _ = v23
        case "":
            v24 := GenS1{F0: v12, F1: v12}
            _ = v24
        }
    }
    _ = len(v13)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v25 := GenS1{F0: -2147483648, F1: 1}
    if false {
        if (!(18446744073709551615 != 42)) {
            fmt.Println(v12)
        } else {
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 int64 = -1
                v12 = 32767
                _ = fi26
                _ = v27
            }
            if (386 < 2147483647) {
                if (-45 >= 0) {
                    v10 -= 1000
                    var v28 float32 = -2.25
                    _ = v28
                } else {
                    var v29 float64 = ((float64(0.1) * 1e10) / ((0.1 * 0.1) + ((float64(0.0) / 0.1) + (-2.25 - (((3.14159 + ((1.5 - 1.5) * 1e10)) / 0.1) * -2.25)))))
                    var v30 int16 = -1
                    switch 4294967295 {
                    case 2147483647:
                        for fi31 := 0; fi31 < 15; fi31++ {
                            for fi32 := 0; fi32 < 1; fi32++ {
                                if false {
                                    var v33 int64 = int64(fi32)
                                    _ = v33
                                } else {
                                    fmt.Println(v10)
                                }
                                _ = len(v13)
                                var v34 int16 = (v30 % ((v10) | 1))
                                _ = fi32
                                _ = v34
                            }
                            for fi35 := 0; fi35 < 9; fi35++ {
                                v16 = v16
                                var v36 string = "a"
                                var v37 string = (v36 + "foo bar")
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            _ = fi31
                        }
                    case 435:
                        _ = len(v13)
                    }
                    _ = v29
                    _ = v30
                }
                v38 := GenS1{F0: (v12 + (v12 << ((753) & 63))), F1: -423}
                fmt.Println(v12, v12)
                _ = v38
            } else {
                if false {
                    switch "a" {
                    case "a":
                        var v39 uint = 123
                        _ = v39
                    case "\n\t":
                        v40 := GenS2{F0: "hello"}
                        _ = v40
                    case "":
                        v41 := GenS2{F0: ""}
                        _ = v41
                    default:
                        var v42 uint8 = 239
                        _ = v42
                    }
                } else {
                    for fi43 := 0; fi43 < 7; fi43++ {
                        var v44 float64 = ((1e10 + (0.1 - 3.14159)) * (0.0 * (-2.25 * -2.25)))
                        if (-888 >= 1) {
                            v2[(fi43 - fi43)] = (v16 << ((149) & 63))
                            var v45 int32 = 2147483647
                            _ = v45
                        }
                        _ = fi43
                        _ = v44
                    }
                    for fi46 := 0; fi46 < 15; fi46++ {
                        v47 := GenS2{F0: "hello"}
                        if (true && (!true)) {
                            if false {
                                if false {
                                    var v48 uint16 = 2
                                    fmt.Println(fi46, v12)
                                    for fi49 := 0; fi49 < 17; fi49++ {
                                        if true {
                                            v12 += (fi46 % ((fi46) | 1))
                                        }
                                        _ = fi49
                                    }
                                    _ = v48
                                } else {
                                    if (!true) {
                                        var v50 uint64 = 2
                                        var v51 uint32 = v16
                                        _ = v50
                                        _ = v51
                                    } else {
                                        var v52 int16 = (v10 * (v10 << ((271) & 63)))
                                        if (false || ((^(v16)) <= (v16 * (v16 >> ((507) & 63))))) {
                                            var v53 string = ""
                                            switch v53 {
                                            case "":
                                                v16 += (v16 * (^(v16)))
                                            case "hello":
                                                for fi54 := 0; fi54 < 8; fi54++ {
                                                    var v55 uint = 219
                                                    if false {
                                                        var v56 float32 = 0.1
                                                        if (-1 >= 9223372036854775807) {
                                                            switch "a" {
                                                            case "a":
                                                                v57 := GenS3{F0: 255, F1: (42 >= 125)}
                                                                _ = v57
                                                            case "\n\t":
                                                                var v58 uint32 = v16
                                                                _ = v58
                                                            case "hello":
                                                                v55 = v55
                                                            }
                                                            switch (v53 + v53) {
                                                            case "a":
                                                                delete(v3, (v53 + v53))
                                                            }
                                                        } else {
                                                            v53 = "\n\t"
                                                            v59 := GenS4{F0: fi54, F1: 65535, F2: 549}
                                                            switch fi46 {
                                                            case -9223372036854775808:
                                                                v60 := GenS5{F0: -1}
                                                                _ = v60
                                                            case -858:
                                                                if ((^(v10)) == (v52 >> ((v55) & 63))) {
                                                                    v2[(fi46 / ((fi46) | 1))] = (v16 >> ((4294967295) & 63))
                                                                    var v61 int32 = 42
                                                                    v62 := GenS6{F0: 42, F1: 42}
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    _ = len(v13)
                                                                }
                                                                _ = len(v2)
                                                                if (0.0 < 0.1) {
                                                                    var v63 int32 = 42
                                                                    switch 42 {
                                                                    case 520:
                                                                        v64 := GenS6{F0: 0, F1: 255}
                                                                        _ = v64
                                                                    case 128:
                                                                        if false {
                                                                            v13 = append(v13, 42)
                                                                            switch ("a" + ((v53 + v53) + "\n\t")) {
                                                                            case "foo bar":
                                                                                v53 = v53
                                                                            case "\n\t":
                                                                                switch v53 {
                                                                                case "hello":
                                                                                    switch (-(fi46)) {
                                                                                    case -972:
                                                                                        v65 := GenS7{F0: -9223372036854775808, F1: v53}
                                                                                        _ = v65
                                                                                    case 4294967295:
                                                                                        v2[-633] = v16
                                                                                    case -638:
                                                                                        switch (fi54 - fi54) {
                                                                                        case 0:
                                                                                            if true {
                                                                                                fmt.Println(v56, v16, v55)
                                                                                            } else {
                                                                                                if len(v13) > 0 {
                                                                                                    v13[0] = 766
                                                                                                }
                                                                                                var v66 uint16 = 385
                                                                                                v67 := v1[(v16 % (((v16 / (((v16 / (((v16 + v16)) | 1))) | 1))) | 1))]
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                            }
                                                                                        case 4294967295:
                                                                                            if (18446744073709551615 > 0) {
                                                                                                switch v12 {
                                                                                                case 0:
                                                                                                    v56 = v56
                                                                                                }
                                                                                                _ = len(v2)
                                                                                                switch v53 {
                                                                                                case "a":
                                                                                                    var v68 uint16 = 410
                                                                                                    _ = v68
                                                                                                case "foo bar":
                                                                                                    if (4294967295 == 42) {
                                                                                                        var v69 int8 = 100
                                                                                                        if false {
                                                                                                            v56 = v56
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v70 := GenS3{F0: 32767, F1: false}
                                                                                                                v71 := GenS7{F0: 9223372036854775807, F1: v53}
                                                                                                                _ = v70
                                                                                                                _ = v71
                                                                                                            }
                                                                                                            if ((v63 * (^(v63))) < v63) {
                                                                                                                _ = len(v3)
                                                                                                                var v72 int64 = 42
                                                                                                                if len(v13) > 0 {
                                                                                                                    v13[0] = v12
                                                                                                                }
                                                                                                                _ = v72
                                                                                                            }
                                                                                                            if (false && false) {
                                                                                                                var v73 uint = (v55 * 18446744073709551615)
                                                                                                                _ = v73
                                                                                                            } else {
                                                                                                                if (!(9223372036854775807 <= -1)) {
                                                                                                                    v74 := GenS5{F0: (v69 / (((-(v69))) | 1))}
                                                                                                                    _ = v74
                                                                                                                }
                                                                                                                if true {
                                                                                                                    v69 = (v69 ^ v69)
                                                                                                                } else {
                                                                                                                    v16 = v16
                                                                                                                    var v75 uint64 = 23
                                                                                                                    _ = v75
                                                                                                                }
                                                                                                                if (!true) {
                                                                                                                    if true {
                                                                                                                        fmt.Println(v10, v56, fi46)
                                                                                                                        v69 = 100
                                                                                                                        if false {
                                                                                                                            var v76 float32 = ((-2.25 * v56) + v56)
                                                                                                                            var v77 uint = v55
                                                                                                                            switch ((v53 + v53) + "") {
                                                                                                                            case "a":
                                                                                                                                _ = len(v13)
                                                                                                                            case "hello":
                                                                                                                                switch 32767 {
                                                                                                                                case 4294967295:
                                                                                                                                    v56 += v56
                                                                                                                                }
                                                                                                                            case "foo bar":
                                                                                                                                var v78 uint8 = 26
                                                                                                                                _ = v78
                                                                                                                            default:
                                                                                                                                var v79 uint = 0
                                                                                                                                _ = v79
                                                                                                                            }
                                                                                                                            var v80 int8 = (v69 >> (((v77 % ((v77) | 1))) & 63))
                                                                                                                            var v81 int32 = v63
                                                                                                                            v52 += v52
                                                                                                                            fmt.Println(v16, v80)
                                                                                                                            _ = v76
                                                                                                                            _ = v77
                                                                                                                            _ = v80
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if (false && true) {
                                                                                                                            if false {
                                                                                                                                v10 += (v10 ^ (v52 ^ genFn1(562, genFn0(), 65535)))
                                                                                                                            } else {
                                                                                                                                v10 -= int16(fi46)
                                                                                                                            }
                                                                                                                            _ = len(v3)
                                                                                                                            var v82 uint64 = 0
                                                                                                                            _ = v82
                                                                                                                        }
                                                                                                                        var v83 int16 = v10
                                                                                                                        var v84 int16 = v10
                                                                                                                        _ = v83
                                                                                                                        _ = v84
                                                                                                                    }
                                                                                                                    v13 = append(v13, (v12 % ((fi46) | 1)))
                                                                                                                } else {
                                                                                                                    v85 := GenS8{F0: 160, F1: v55}
                                                                                                                    var v86 uint32 = v16
                                                                                                                    _ = v85
                                                                                                                    _ = v86
                                                                                                                }
                                                                                                            }
                                                                                                            v10 -= genFn1(2, -128, 2)
                                                                                                            fmt.Println(fi46, v56)
                                                                                                        }
                                                                                                        var v87 int = (^(fi46))
                                                                                                        _ = v69
                                                                                                        _ = v87
                                                                                                    }
                                                                                                case "hello":
                                                                                                    var v88 float64 = -2.25
                                                                                                    _ = v88
                                                                                                default:
                                                                                                    var v89 uint = (v55 >> ((4294967295) & 63))
                                                                                                    var v90 uint8 = 214
                                                                                                    _ = v89
                                                                                                    _ = v90
                                                                                                }
                                                                                            } else {
                                                                                                v52 -= v10
                                                                                                if false {
                                                                                                    v53 = (((v53 + (v53 + (v53 + v53))) + (v53 + v53)) + "foo bar")
                                                                                                } else {
                                                                                                    var v91 string = (((v53 + ("hello" + v53)) + v53) + v53)
                                                                                                    if (float64(v56) > 1.5) {
                                                                                                        v1[v16] = 0
                                                                                                        v92 := GenS9{F0: -915, F1: "foo bar", F2: 0}
                                                                                                        delete(v3, (((v53 + "\n\t") + ("a" + ((v53 + v53) + v91))) + v91))
                                                                                                        _ = v92
                                                                                                    }
                                                                                                    _ = v91
                                                                                                }
                                                                                                v16 = 0
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            fmt.Println(v16, v56)
                                                                                        }
                                                                                    default:
                                                                                        _ = len(v13)
                                                                                    }
                                                                                case "a":
                                                                                    if len(v13) > 0 {
                                                                                        v13[0] = fi46
                                                                                    }
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        v56 += 1.5
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            if (-2.25 >= 1.5) {
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    v55 = (v55 ^ v55)
                                                                                                case "a":
                                                                                                    v93 := GenS10{F0: 0, F1: v52}
                                                                                                    _ = v93
                                                                                                case "hello":
                                                                                                    v94 := GenS11{F0: -128, F1: v53}
                                                                                                    _ = v94
                                                                                                }
                                                                                                if len(v13) >= 2 {
                                                                                                    v13 = v13[1:len(v13)]
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v56 = v56
                                                                                            if (!true) {
                                                                                                if (false && ((((v56 + 0.0) / (v56 * v56)) / v56) <= 3.14159)) {
                                                                                                    v12 += (fi54 | v12)
                                                                                                    var v95 uint32 = (^(v16))
                                                                                                    fmt.Println(v63, v16)
                                                                                                    v96 := GenS4{F0: 2147483647, F1: 42, F2: (v63 / (((^(v63))) | 1))}
                                                                                                    var v97 uint = v55
                                                                                                    _ = v95
                                                                                                    _ = v96
                                                                                                    _ = v97
                                                                                                } else {
                                                                                                    v52 = 0
                                                                                                    v98 := GenS12{F0: v63, F1: 42, F2: (v63 | (v63 % ((v63) | 1)))}
                                                                                                    _ = v98
                                                                                                }
                                                                                            } else {
                                                                                                if len(v13) > 0 {
                                                                                                    v13[0] = v12
                                                                                                }
                                                                                                var v99 uint64 = 250
                                                                                                if len(v13) > 0 {
                                                                                                    v13[0] = fi54
                                                                                                }
                                                                                                _ = v99
                                                                                            }
                                                                                            fmt.Println(v53, v10, v53)
                                                                                        }
                                                                                    }
                                                                                    var v100 string = v53
                                                                                    _ = v100
                                                                                }
                                                                            case "hello":
                                                                                if (255 != 2) {
                                                                                    v16 = 756
                                                                                    switch -965 {
                                                                                    case -2147483648:
                                                                                        v10 -= (v10 % ((v10) | 1))
                                                                                    case 2147483647:
                                                                                        fmt.Println(fi46, v16)
                                                                                    case 9223372036854775807:
                                                                                        switch (v12 % ((fi46) | 1)) {
                                                                                        case -1:
                                                                                            v63 -= v63
                                                                                        default:
                                                                                            var v101 float64 = ((-2.25 - (1e10 + -2.25)) / 1.0)
                                                                                            _ = v101
                                                                                        }
                                                                                        var v102 int8 = 100
                                                                                        _ = v102
                                                                                    }
                                                                                    if true {
                                                                                        if false {
                                                                                            _ = len(v13)
                                                                                        } else {
                                                                                            var v103 uint8 = 255
                                                                                            var v104 bool = ((0 > 0) || ((^(v103)) == v103))
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                        }
                                                                                        v105 := GenS6{F0: 314, F1: 32767}
                                                                                        v106 := GenS13{F0: v63, F1: 721, F2: false}
                                                                                        if false {
                                                                                            switch v12 {
                                                                                            case 353:
                                                                                                if true {
                                                                                                    switch "a" {
                                                                                                    case "hello":
                                                                                                        switch (^(fi46)) {
                                                                                                        case 32767:
                                                                                                            v107 := GenS14{F0: genFn1(567, 100, 633)}
                                                                                                            _ = v107
                                                                                                        case -874:
                                                                                                            if ((3.14159 < v56) || true) {
                                                                                                                v108 := GenS15{F0: (v53 + v53), F1: fi46}
                                                                                                                fmt.Println(v55, v16, fi46)
                                                                                                                if false {
                                                                                                                    var v109 float64 = 0.1
                                                                                                                    v110 := GenS13{F0: v63, F1: v63, F2: ((v10 >> ((0) & 63)) < genFn1(0, 55, 42))}
                                                                                                                    _ = v109
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                                _ = v108
                                                                                                            } else {
                                                                                                                v111 := GenS16{F0: 65535, F1: 9223372036854775807}
                                                                                                                if (!true) {
                                                                                                                    fmt.Println(fi54)
                                                                                                                    switch v53 {
                                                                                                                    case "":
                                                                                                                        fmt.Println(fi46, v16, v63)
                                                                                                                    case "foo bar":
                                                                                                                        delete(v2, -683)
                                                                                                                    default:
                                                                                                                        if ((v10 | v10) == -552) {
                                                                                                                            var v112 float32 = float32(1.5)
                                                                                                                            var v113 int8 = -128
                                                                                                                            _ = v112
                                                                                                                            _ = v113
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                var v114 int16 = (v52 - genFn1(255, 42, 65535))
                                                                                                                _ = v111
                                                                                                                _ = v114
                                                                                                            }
                                                                                                        case 32768:
                                                                                                            switch "hello" {
                                                                                                            case "a":
                                                                                                                v55 -= (^(v55))
                                                                                                            case "\n\t":
                                                                                                                switch ((v53 + (v53 + v53)) + (v53 + "foo bar")) {
                                                                                                                case "\n\t":
                                                                                                                    if (!false) {
                                                                                                                        var v115 uint64 = 31
                                                                                                                        fmt.Println(v115)
                                                                                                                        var v116 uint = (v55 / ((42) | 1))
                                                                                                                        fmt.Println(v55)
                                                                                                                        var v117 int32 = (v63 - (v63 << (((v116 % (((v116 ^ 42)) | 1))) & 63)))
                                                                                                                        var v118 uint16 = 594
                                                                                                                        v13 = append(v13, 32767)
                                                                                                                        _ = v115
                                                                                                                        _ = v116
                                                                                                                        _ = v117
                                                                                                                        _ = v118
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    if ((false || (-1 > (-(v63)))) && (false && true)) {
                                                                                                                        var v119 string = (v53 + "foo bar")
                                                                                                                        _ = v119
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v120 := GenS11{F0: -128, F1: "a"}
                                                                                                    v56 += v56
                                                                                                    var v121 float32 = 0.1
                                                                                                    _ = v120
                                                                                                    _ = v121
                                                                                                }
                                                                                            case 0:
                                                                                                v122 := GenS11{F0: -128, F1: "foo bar"}
                                                                                                _ = v122
                                                                                            default:
                                                                                                if true {
                                                                                                    v56 += float32(0.1)
                                                                                                    var v123 bool = false
                                                                                                    if len(v13) > 0 {
                                                                                                        v13[0] = v12
                                                                                                    }
                                                                                                    _ = v123
                                                                                                }
                                                                                            }
                                                                                            delete(v1, v16)
                                                                                            fmt.Println(v56, v52)
                                                                                        } else {
                                                                                            delete(v1, (v16 - (v16 & v16)))
                                                                                            var v124 int16 = 32767
                                                                                            switch v53 {
                                                                                            case "foo bar":
                                                                                                fmt.Println(fi54, fi54, v55)
                                                                                            default:
                                                                                                switch (fi54 >> ((18446744073709551615) & 63)) {
                                                                                                case 795:
                                                                                                    v13 = append(v13, fi46)
                                                                                                case 32767:
                                                                                                    v124 += int16(v16)
                                                                                                }
                                                                                            }
                                                                                            _ = v124
                                                                                        }
                                                                                        v125 := GenS3{F0: 42, F1: (((v16 ^ (v16 % ((v16) | 1))) <= (v16 + (v16 + 1))) || false)}
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                        _ = v125
                                                                                    } else {
                                                                                        v126 := GenS17{F0: true, F1: 327}
                                                                                        if (float32((-2.25 * -2.25)) < (v56 + v56)) {
                                                                                            v1[0] = v16
                                                                                            if (-1 > 85) {
                                                                                                if true {
                                                                                                    var v127 float64 = (((1e10 * -2.25) - float64(float32(((0.0 * (3.14159 + float64(v56))) + -2.25)))) + -2.25)
                                                                                                    _ = v127
                                                                                                } else {
                                                                                                    var v128 uint64 = 18446744073709551615
                                                                                                    _ = v128
                                                                                                }
                                                                                                var v129 uint = 65535
                                                                                                fmt.Println(v56, fi54)
                                                                                                _ = v129
                                                                                            } else {
                                                                                                var v130 int32 = int32(v55)
                                                                                                _ = v130
                                                                                            }
                                                                                            var v131 bool = (18446744073709551615 >= 255)
                                                                                            _ = v131
                                                                                        }
                                                                                        _ = v126
                                                                                    }
                                                                                } else {
                                                                                    var v132 uint = 58
                                                                                    _ = v132
                                                                                }
                                                                            }
                                                                            v52 = (v10 + (v10 ^ (v52 + v52)))
                                                                        }
                                                                    }
                                                                    v133 := GenS11{F0: -1, F1: ((v53 + "hello") + (("a" + v53) + ((v53 + v53) + (v53 + ("" + v53)))))}
                                                                    _ = v133
                                                                    _ = v63
                                                                } else {
                                                                    switch fi46 {
                                                                    case 29:
                                                                        delete(v1, (^(v16)))
                                                                    case 128:
                                                                        var v134 uint16 = 255
                                                                        _ = v134
                                                                    }
                                                                }
                                                            }
                                                            _ = v59
                                                        }
                                                        var v135 int64 = 0
                                                        v53 = v53
                                                        v12 = v12
                                                        _ = v135
                                                        _ = v56
                                                    } else {
                                                        if (1 > -762) {
                                                            v136 := GenS18{F0: (v16 - v16), F1: 2}
                                                            v137 := GenS19{F0: (v53 + (v53 + "foo bar"))}
                                                            _ = v136
                                                            _ = v137
                                                        }
                                                        var v138 float64 = 1e10
                                                        v10 -= v10
                                                        _ = v138
                                                    }
                                                    _ = fi54
                                                    _ = v55
                                                }
                                            default:
                                                v139 := GenS20{F0: (-(v12)), F1: -32768}
                                                _ = v139
                                            }
                                            _ = v53
                                        }
                                        var v140 uint = 791
                                        _ = v140
                                        _ = v52
                                    }
                                    if (!false) {
                                        switch ("a" + ((("\n\t" + "") + "") + ("hello" + ((("\n\t" + "\n\t") + (((("\n\t" + ("foo bar" + ("a" + "foo bar"))) + (("" + "a") + ("a" + ""))) + "") + (("\n\t" + "\n\t") + (("\n\t" + "a") + ("foo bar" + ("a" + (("hello" + "foo bar") + "\n\t"))))))) + "foo bar")))) {
                                        case "":
                                            switch v12 {
                                            case 0:
                                                if (!(!((2147483647 < -2147483648) || false))) {
                                                    v12 = (fi46 >> ((893) & 63))
                                                    var v141 float64 = 1.5
                                                    fmt.Println(fi46, v12)
                                                    _ = v141
                                                }
                                            default:
                                                switch "hello" {
                                                case "foo bar":
                                                    var v142 string = ("" + ("hello" + "foo bar"))
                                                    _ = v142
                                                case "hello":
                                                    var v143 uint32 = v16
                                                    _ = v143
                                                case "\n\t":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v144 int16 = (v10 >> ((42) & 63))
                                                        _ = v144
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                }
                                            }
                                        default:
                                            for fi145 := 0; fi145 < 1; fi145++ {
                                                if len(v13) >= 2 {
                                                    v13 = v13[1:len(v13)]
                                                }
                                                var v146 uint8 = 0
                                                _ = fi145
                                                _ = v146
                                            }
                                        }
                                        if (0.1 == 0.1) {
                                            v10 = v10
                                            v12 -= v12
                                            if false {
                                                for fi147 := 0; fi147 < 15; fi147++ {
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    _ = fi147
                                                }
                                                _ = len(v2)
                                            } else {
                                                for fi148 := 0; fi148 < 1; fi148++ {
                                                    switch fi46 {
                                                    case 32768:
                                                        var v149 uint8 = 255
                                                        v3[((("\n\t" + (((("foo bar" + "hello") + "") + (("\n\t" + "\n\t") + ("\n\t" + "hello"))) + "\n\t")) + "\n\t") + "foo bar")] = fi46
                                                        v150 := GenS15{F0: ("\n\t" + (("" + "") + "a")), F1: (fi46 / (((fi148 ^ v12)) | 1))}
                                                        v151 := GenS21{F0: 42, F1: 347, F2: -128}
                                                        v16 = 334
                                                        _ = v149
                                                        _ = v150
                                                        _ = v151
                                                    case 127:
                                                        var v152 bool = false
                                                        _ = v152
                                                    }
                                                    fmt.Println(v10, v12)
                                                    _ = fi148
                                                }
                                            }
                                        }
                                    } else {
                                        var v153 uint64 = 4294967295
                                        for fi154 := 0; fi154 < 4; fi154++ {
                                            fmt.Println(v153)
                                            fmt.Println(v153, v153)
                                            if (!false) {
                                                v153 -= uint64(v16)
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(v10, v16, v16)
                                                case "hello":
                                                    switch 830 {
                                                    case 32768:
                                                        fmt.Println(fi154, v10)
                                                    default:
                                                        v16 -= v16
                                                    }
                                                case "a":
                                                    v155 := GenS22{F0: 1, F1: v16}
                                                    _ = v155
                                                default:
                                                    fmt.Println(v16)
                                                }
                                                v153 -= v153
                                                v153 -= (v153 & 0)
                                                if ((-2.25 >= 3.14159) && (uint32(v153) < 193)) {
                                                    if (42 >= 23) {
                                                        var v156 uint = 741
                                                        _ = v156
                                                    }
                                                }
                                            }
                                            _ = fi154
                                        }
                                        var v157 uint8 = 0
                                        _ = v153
                                        _ = v157
                                    }
                                }
                                switch ("hello" + "\n\t") {
                                case "hello":
                                    v16 += 65535
                                default:
                                    v16 -= v16
                                }
                            }
                            for fi158 := 0; fi158 < 7; fi158++ {
                                switch (fi158 / (((fi158 & v12)) | 1)) {
                                case 1:
                                    var v159 int = fi158
                                    _ = v159
                                case -1:
                                    var v160 uint64 = 4294967295
                                    _ = v160
                                case 9223372036854775807:
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                default:
                                    _ = len(v13)
                                }
                                _ = fi158
                            }
                            v16 -= v16
                        } else {
                            fmt.Println(v10)
                            if (1 > 1) {
                                if (v10 != 42) {
                                    var v161 uint = 4294967295
                                    fmt.Println(v16, v10)
                                    _ = v161
                                } else {
                                    v162 := GenS23{F0: (fi46 - fi46), F1: 32767, F2: 127}
                                    for fi163 := 0; fi163 < 15; fi163++ {
                                        v16 = v16
                                        switch (("foo bar" + "") + ("\n\t" + "foo bar")) {
                                        case "foo bar":
                                            v12 = v12
                                        case "":
                                            var v164 int32 = 1
                                            switch ("hello" + "hello") {
                                            case "":
                                                if (1 >= 743) {
                                                    var v165 float64 = 1e10
                                                    var v166 uint32 = v16
                                                    _ = v165
                                                    _ = v166
                                                }
                                            }
                                            _ = v164
                                        }
                                        var v167 uint32 = 0
                                        _ = fi163
                                        _ = v167
                                    }
                                    _ = v162
                                }
                            }
                        }
                        v12 += -9223372036854775808
                        _ = fi46
                        _ = v47
                    }
                }
            }
        }
        fmt.Println(v10)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
