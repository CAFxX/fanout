package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int; F2 uint }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 int32; F1 uint32; F2 int16 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int32; F1 int8; F2 uint64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint8; F1 uint64 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int16; F1 uint; F2 string }
type GenS19 struct { F0 int16; F1 bool }
func genFn0(p0 bool) int16 {
    v5 := GenS1{F0: p0, F1: 0}
    var v6 int64 = 145
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{fi8, (fi8 ^ (fi8 - fi8))}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 string = "a"
    var v2 int64 = 0
    var v3 int8 = 1
    var v4 int32 = -237
    if (1.5 == (3.14159 / (-2.25 + float32(((0.0 * float64((0.1 * 1e10))) + 1e10))))) {
        if ((!((v2 < 0) || true)) || (590 >= 2)) {
            v10 := []int{371, 0, -191}
            v11 := GenS1{F0: false, F1: 42}
            if true {
                v3 -= -128
            } else {
                v10 = append(v10, 32768)
                v12 := GenS2{F0: 127, F1: -1, F2: 65535}
                _ = v12
            }
            _ = v10
            _ = v11
        } else {
            var v13 int = -378
            if true {
                var v14 bool = ((!((v13 | (v13 % ((376) | 1))) < v13)) || true)
                v15 := []int{-207, 128, v13}
                _ = v14
                _ = v15
            } else {
                var v16 bool = true
                v2 = -9223372036854775808
                _ = v16
            }
            _ = v13
        }
        v17 := GenS1{F0: true, F1: 127}
        v3 = v3
        _ = v17
    } else {
        v2 = (v2 + v2)
        if false {
            var v18 bool = ((!false) || false)
            v2 -= 0
            for fi19 := 0; fi19 < 7; fi19++ {
                v18 = v18
                for fi20 := 0; fi20 < 9; fi20++ {
                    if false {
                        var v21 uint32 = 0
                        if true {
                            if false {
                                switch -1 {
                                case 4294967295:
                                    var v22 int8 = 100
                                    _ = v22
                                case -948:
                                    v23 := []int{}
                                    _ = v23
                                case 128:
                                    switch "a" {
                                    case "foo bar":
                                        if (1 <= 572) {
                                            if false {
                                                v18 = v18
                                                fmt.Println(fi20)
                                            }
                                        }
                                    case "a":
                                        switch (v1 + v1) {
                                        case "hello":
                                            v24 := []int{32767}
                                            _ = v24
                                        case "\n\t":
                                            v1 = v1
                                        case "a":
                                            v21 = (^(v21))
                                        }
                                    default:
                                        switch v1 {
                                        case "a":
                                            fmt.Println(v3)
                                        case "hello":
                                            fmt.Println(v21, v1, v3)
                                        }
                                        fmt.Println(v4, v3, v18)
                                        var v25 bool = false
                                        _ = v25
                                    }
                                default:
                                    v26 := []int{fi19, fi20}
                                    _ = v26
                                }
                                v27 := []int{fi20, 0}
                                _ = v27
                            }
                            if ((v18 && (float32(3.14159) >= 3.14159)) || v18) {
                                if (v3 < 100) {
                                    fmt.Println(fi19, v18)
                                    fmt.Println(v18, v1, fi20)
                                    if (uint16(v21) < 570) {
                                        v28 := []int{fi20, fi20}
                                        v3 = (v3 - 42)
                                        _ = v28
                                    } else {
                                        v29 := []int{}
                                        _ = v29
                                    }
                                } else {
                                    v30 := GenS2{F0: 180, F1: (^(fi20)), F2: 65535}
                                    fmt.Println(v3, fi20, v3)
                                    v3 -= 1
                                    _ = v30
                                }
                                v31 := GenS3{F0: 18446744073709551615, F1: 2, F2: 493}
                                v32 := []int{(fi19 & -343)}
                                _ = v31
                                _ = v32
                            }
                            fmt.Println(v18)
                        }
                        _ = v21
                    }
                    v4 -= (v4 ^ (v4 << ((2) & 63)))
                    _ = fi20
                }
                fmt.Println(v4, v4)
                _ = fi19
            }
            _ = v18
        }
        v3 -= int8(v4)
    }
    var v33 uint32 = 0
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 int16 = 859
        if false {
            var v36 float64 = (((0.0 + 0.0) - 0.0) + float64(0.1))
            _ = v36
        } else {
            v3 -= v3
        }
        var v37 uint = 42
        _ = fi34
        _ = v35
        _ = v37
    }
    v4 = (v4 | v4)
    v38 := GenS2{F0: 0, F1: 747, F2: 42}
    if true {
        for fi39 := 0; fi39 < 6; fi39++ {
            for fi40 := 0; fi40 < 5; fi40++ {
                var v41 uint32 = (v33 >> ((191) & 63))
                v42 := GenS4{F0: (genFn0(((^(v41)) > v33)) < -795), F1: 374}
                if true {
                    fmt.Println(v3)
                    if false {
                        v4 = (v4 >> ((42) & 63))
                        if (v3 >= (v3 >> ((18446744073709551615) & 63))) {
                            fmt.Println(v33, v1, fi40)
                            if true {
                                v1 = v1
                                v43 := []int{fi40, fi39, -1}
                                _ = v43
                            }
                        }
                        var v44 int64 = (^(v2))
                        _ = v44
                    } else {
                        fmt.Println(v41, v3, fi39)
                        v45 := []int{(-(fi40))}
                        switch (v1 + "foo bar") {
                        case "\n\t":
                            if ((v3 ^ (v3 << ((280) & 63))) == (v3 & v3)) {
                                if (!(1 == 970)) {
                                    switch v1 {
                                    case "a":
                                        _ = len(v45)
                                    case "":
                                        v4 = v4
                                    default:
                                        var v46 int64 = (^(v2))
                                        _ = v46
                                    }
                                }
                                if (3.14159 >= 0.0) {
                                    switch (fi40 | -2147483648) {
                                    case 980:
                                        if (false || false) {
                                            if false {
                                                v2 -= (v2 | 1)
                                                if false {
                                                    var v47 uint8 = 2
                                                    _ = v47
                                                } else {
                                                    if (true && (!true)) {
                                                        if true {
                                                            if (!(true && false)) {
                                                                if (!true) {
                                                                    v48 := GenS4{F0: (42 != 0), F1: 1}
                                                                    v2 = (v2 - v2)
                                                                    _ = len(v45)
                                                                    _ = v48
                                                                }
                                                                if len(v45) > 0 {
                                                                    v45[0] = (fi39 | fi40)
                                                                }
                                                            } else {
                                                                v2 -= v2
                                                            }
                                                            if (57 == 1) {
                                                                var v49 int64 = v2
                                                                var v50 string = "foo bar"
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                switch 720 {
                                                                case -579:
                                                                    v33 -= v41
                                                                case 42:
                                                                    if false {
                                                                        switch v1 {
                                                                        case "foo bar":
                                                                            if len(v45) > 0 {
                                                                                v45[0] = fi40
                                                                            }
                                                                        case "\n\t":
                                                                            v1 = v1
                                                                        }
                                                                    } else {
                                                                        v51 := GenS5{F0: v4, F1: (v2 * (v2 & v2))}
                                                                        _ = v51
                                                                    }
                                                                }
                                                                if true {
                                                                    var v52 float64 = -2.25
                                                                    v53 := GenS6{F0: v4, F1: 42, F2: -745}
                                                                    _ = v52
                                                                    _ = v53
                                                                } else {
                                                                    fmt.Println(v2)
                                                                    var v54 bool = (v4 <= (v4 & -677))
                                                                    _ = v54
                                                                }
                                                                if (!(!true)) {
                                                                    fmt.Println(v41, v1)
                                                                }
                                                            }
                                                            if false {
                                                                if len(v45) > 0 {
                                                                    v45[0] = 32767
                                                                }
                                                                v3 = (v3 & 127)
                                                                switch 128 {
                                                                case 4294967295:
                                                                    var v55 string = v1
                                                                    _ = v55
                                                                case -1:
                                                                    fmt.Println(v4)
                                                                default:
                                                                    v56 := GenS1{F0: true, F1: 42}
                                                                    _ = v56
                                                                }
                                                                if (65535 < 193) {
                                                                    fmt.Println(fi40)
                                                                    if len(v45) >= 2 {
                                                                        v45 = v45[1:len(v45)]
                                                                    }
                                                                } else {
                                                                    var v57 float32 = (0.0 + 1.5)
                                                                    v58 := GenS5{F0: (v4 % (((v4 / ((v4) | 1))) | 1)), F1: (v2 % ((v2) | 1))}
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                v1 = "hello"
                                                                v1 = ("foo bar" + "\n\t")
                                                            }
                                                        }
                                                        switch (v1 + v1) {
                                                        case "\n\t":
                                                            switch (v1 + ("\n\t" + v1)) {
                                                            case "\n\t":
                                                                switch fi39 {
                                                                case 42:
                                                                    if false {
                                                                        var v59 int16 = 1000
                                                                        fmt.Println(v2)
                                                                        _ = v59
                                                                    } else {
                                                                        switch (v1 + ((v1 + ("foo bar" + "")) + v1)) {
                                                                        case "\n\t":
                                                                            v60 := GenS5{F0: (v4 << ((734) & 63)), F1: (v2 * (-(v2)))}
                                                                            _ = v60
                                                                        case "":
                                                                            switch ("hello" + (v1 + v1)) {
                                                                            case "hello":
                                                                                var v61 int32 = v4
                                                                                v62 := GenS7{F0: -49, F1: 307}
                                                                                switch ("\n\t" + "a") {
                                                                                case "hello":
                                                                                    if ((-2.25 * -2.25) > -2.25) {
                                                                                        var v63 int16 = 1000
                                                                                        if true {
                                                                                            v64 := GenS3{F0: 714, F1: 127, F2: 65535}
                                                                                            _ = len(v45)
                                                                                            v61 = v61
                                                                                            _ = v64
                                                                                        } else {
                                                                                            v65 := GenS8{F0: -32768}
                                                                                            v41 += (^(v41))
                                                                                            _ = v65
                                                                                        }
                                                                                        _ = v63
                                                                                    }
                                                                                }
                                                                                if len(v45) >= 2 {
                                                                                    v45 = v45[1:len(v45)]
                                                                                }
                                                                                switch v1 {
                                                                                case "hello":
                                                                                    if false {
                                                                                        switch fi39 {
                                                                                        case 9223372036854775807:
                                                                                            if len(v45) > 0 {
                                                                                                v45[0] = (fi39 % ((514) | 1))
                                                                                            }
                                                                                        case -1:
                                                                                            if false {
                                                                                                v33 += v33
                                                                                            } else {
                                                                                                v4 = v61
                                                                                                v66 := GenS5{F0: 966, F1: v2}
                                                                                                _ = v66
                                                                                            }
                                                                                        case 2147483647:
                                                                                            var v67 uint = 2
                                                                                            _ = v67
                                                                                        }
                                                                                        var v68 float32 = ((-2.25 / 1.0) - 0.0)
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            if (((v2 << ((0) & 63)) == (^(v2))) || false) {
                                                                                                if len(v45) > 0 {
                                                                                                    v45[0] = (-(fi40))
                                                                                                }
                                                                                                var v69 int8 = int8(v61)
                                                                                                _ = v69
                                                                                            } else {
                                                                                                var v70 int8 = v3
                                                                                                v2 = (^(v2))
                                                                                                _ = v70
                                                                                            }
                                                                                        }
                                                                                        _ = v68
                                                                                    } else {
                                                                                        v4 = 763
                                                                                        var v71 float32 = 0.0
                                                                                        _ = v71
                                                                                    }
                                                                                }
                                                                                _ = v61
                                                                                _ = v62
                                                                            }
                                                                            var v72 uint16 = 899
                                                                            _ = v72
                                                                        case "hello":
                                                                            v1 = v1
                                                                            v1 = v1
                                                                            if ((0 != 586) && true) {
                                                                                fmt.Println(fi40, v33)
                                                                                v73 := GenS1{F0: (!(-2.25 <= 0.1)), F1: 123}
                                                                                switch v1 {
                                                                                case "":
                                                                                    var v74 float32 = 3.14159
                                                                                    _ = v74
                                                                                case "\n\t":
                                                                                    v3 = (v3 | 1)
                                                                                    if true {
                                                                                        if ((!false) && true) {
                                                                                            v75 := GenS9{F0: 61, F1: (!(v3 <= -60))}
                                                                                            v76 := GenS7{F0: (v4 >> ((42) & 63)), F1: 2}
                                                                                            _ = v75
                                                                                            _ = v76
                                                                                        } else {
                                                                                            v41 += v41
                                                                                            fmt.Println(v33, fi39, v1)
                                                                                            var v77 uint32 = (^(v41))
                                                                                            _ = v77
                                                                                        }
                                                                                        if len(v45) > 0 {
                                                                                            v45[0] = (fi39 % (((fi39 ^ fi40)) | 1))
                                                                                        }
                                                                                        v78 := GenS7{F0: 42, F1: 2}
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                                _ = v73
                                                                            } else {
                                                                                v1 = ("" + (v1 + ((v1 + ("\n\t" + "")) + ("\n\t" + ("hello" + "foo bar")))))
                                                                                var v79 int32 = -826
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                    }
                                                                case 32767:
                                                                    var v80 uint64 = 228
                                                                    _ = v80
                                                                }
                                                            case "a":
                                                                v1 = v1
                                                            }
                                                        default:
                                                            switch (v1 + (((v1 + "a") + v1) + "a")) {
                                                            case "\n\t":
                                                                v41 = (v33 ^ (v41 + v33))
                                                            case "hello":
                                                                v45 = append(v45, fi39)
                                                            case "foo bar":
                                                                v45 = append(v45, -432)
                                                            default:
                                                                v3 = (v3 / ((-3) | 1))
                                                            }
                                                        }
                                                    } else {
                                                        if len(v45) >= 2 {
                                                            v45 = v45[1:len(v45)]
                                                        }
                                                        var v81 bool = (18446744073709551615 > 0)
                                                        _ = v81
                                                    }
                                                    fmt.Println(v33, v1)
                                                }
                                            }
                                        } else {
                                            var v82 int64 = 0
                                            v1 = "hello"
                                            switch (fi40 << ((4294967295) & 63)) {
                                            case 879:
                                                switch int(v4) {
                                                case 9223372036854775807:
                                                    v1 = ((v1 + (v1 + v1)) + v1)
                                                case 570:
                                                    v45 = append(v45, 4294967295)
                                                case -9223372036854775808:
                                                    v33 -= v33
                                                }
                                            }
                                            _ = v82
                                        }
                                    }
                                    var v83 bool = (!false)
                                    fmt.Println(v4, v33)
                                    switch fi40 {
                                    case 9223372036854775807:
                                        fmt.Println(v2)
                                    case -535:
                                        if v83 {
                                            if v83 {
                                                if len(v45) > 0 {
                                                    v45[0] = fi39
                                                }
                                                var v84 int16 = -32768
                                                v83 = v83
                                                _ = v84
                                            } else {
                                                v1 = (v1 + (v1 + v1))
                                            }
                                            switch "a" {
                                            case "hello":
                                                switch (fi39 << ((423) & 63)) {
                                                case -292:
                                                    if len(v45) > 0 {
                                                        v45[0] = (-(fi40))
                                                    }
                                                }
                                            case "":
                                                if len(v45) >= 2 {
                                                    v45 = v45[1:len(v45)]
                                                }
                                            default:
                                                if v83 {
                                                    var v85 int64 = 195
                                                    v41 = 108
                                                    _ = v85
                                                } else {
                                                    if false {
                                                        if false {
                                                            var v86 int16 = 32767
                                                            _ = v86
                                                        } else {
                                                            var v87 uint8 = 127
                                                            if v83 {
                                                                v41 = (v41 >> ((255) & 63))
                                                                if v83 {
                                                                    switch ("" + ((v1 + "hello") + v1)) {
                                                                    case "a":
                                                                        v88 := GenS6{F0: 2147483647, F1: (v41 & (v33 << ((4294967295) & 63))), F2: genFn0((!v83))}
                                                                        _ = v88
                                                                    case "\n\t":
                                                                        var v89 int16 = -1
                                                                        _ = v89
                                                                    }
                                                                    v90 := GenS10{F0: -1, F1: v3, F2: 945}
                                                                    switch fi40 {
                                                                    case 1:
                                                                        var v91 uint32 = v41
                                                                        _ = v91
                                                                    }
                                                                    var v92 string = v1
                                                                    _ = v90
                                                                    _ = v92
                                                                }
                                                            } else {
                                                                fmt.Println(v4)
                                                                var v93 float64 = 0.1
                                                                if v83 {
                                                                    fmt.Println(v2, fi40, v3)
                                                                    switch fi40 {
                                                                    case -207:
                                                                        var v94 float64 = v93
                                                                        _ = v94
                                                                    case 0:
                                                                        switch fi40 {
                                                                        case 640:
                                                                            v93 += v93
                                                                            fmt.Println(v2)
                                                                            v33 = 4294967295
                                                                        case 0:
                                                                            if ((v3 * (v3 + v3)) <= 100) {
                                                                                var v95 float64 = (0.0 - v93)
                                                                                var v96 int = 9223372036854775807
                                                                                _ = v95
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v97 uint16 = 32767
                                                                        _ = v97
                                                                    }
                                                                    var v98 uint32 = v41
                                                                    _ = v98
                                                                }
                                                                _ = v93
                                                            }
                                                            if len(v45) > 0 {
                                                                v45[0] = (fi39 >> ((0) & 63))
                                                            }
                                                            _ = v87
                                                        }
                                                    } else {
                                                        if (41 == 822) {
                                                            _ = len(v45)
                                                        } else {
                                                            if false {
                                                                switch v1 {
                                                                case "hello":
                                                                    if v83 {
                                                                        var v99 uint16 = 1
                                                                        v99 = (v99 << ((624) & 63))
                                                                        _ = v99
                                                                    }
                                                                case "a":
                                                                    var v100 int64 = v2
                                                                    _ = v100
                                                                default:
                                                                    v41 = 486
                                                                }
                                                                fmt.Println(v83, v2, v1)
                                                            }
                                                            fmt.Println(v33)
                                                            var v101 bool = true
                                                            _ = v101
                                                        }
                                                        fmt.Println(v41, v3)
                                                        switch (fi40 + -925) {
                                                        case -2147483648:
                                                            if ((^(v41)) >= 65535) {
                                                                v102 := GenS11{F0: 0}
                                                                v33 = (v41 % ((v41) | 1))
                                                                var v103 int32 = (v4 + int32(fi40))
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                        case -1:
                                                            switch ("hello" + (v1 + ("foo bar" + ("hello" + (((v1 + ("\n\t" + "\n\t")) + ("a" + ((v1 + ((v1 + ((v1 + ((v1 + v1) + v1)) + v1)) + "")) + v1))) + (v1 + "hello")))))) {
                                                            case "foo bar":
                                                                switch fi39 {
                                                                case 32768:
                                                                    v104 := GenS12{F0: (v33 / (((v41 - 351)) | 1)), F1: v4}
                                                                    _ = v104
                                                                case 32767:
                                                                    fmt.Println(v33, fi40)
                                                                case 42:
                                                                    v105 := GenS3{F0: 201, F1: 255, F2: 791}
                                                                    _ = v105
                                                                }
                                                                fmt.Println(v41)
                                                            default:
                                                                v106 := GenS13{F0: (-(v2))}
                                                                _ = v106
                                                            }
                                                        case 2147483647:
                                                            var v107 float64 = float64(3.14159)
                                                            _ = v107
                                                        }
                                                    }
                                                }
                                            }
                                            switch "" {
                                            case "foo bar":
                                                v4 -= int32(v33)
                                            case "\n\t":
                                                v108 := GenS2{F0: 128, F1: fi40, F2: 0}
                                                _ = v108
                                            default:
                                                var v109 uint32 = (v41 * v33)
                                                _ = v109
                                            }
                                        } else {
                                            v110 := GenS1{F0: v83, F1: 127}
                                            v1 = ""
                                            _ = v110
                                        }
                                    case 128:
                                        var v111 int = fi39
                                        _ = v111
                                    default:
                                        if ((0.0 + 0.1) > ((-2.25 + (3.14159 - ((0.0 * (float64(3.14159) + (((1.5 * 0.0) - 1e10) - 0.1))) * 3.14159))) / 1.0)) {
                                            v41 = (v41 >> ((255) & 63))
                                            v112 := GenS14{F0: 1, F1: 0}
                                            if v83 {
                                                v2 -= (v2 % ((-857) | 1))
                                                fmt.Println(v1, fi39, v4)
                                                switch v1 {
                                                case "\n\t":
                                                    v3 += v3
                                                    v41 += v41
                                                    fmt.Println(v3)
                                                case "":
                                                    v113 := GenS5{F0: 2147483647, F1: 0}
                                                    _ = v113
                                                case "hello":
                                                    if (!((255 != 0) && ((v2 / ((v2) | 1)) != v2))) {
                                                        if (2 > 473) {
                                                            fmt.Println(v33, v2)
                                                        }
                                                        switch (-(fi40)) {
                                                        case 128:
                                                            fmt.Println(v83, fi40, fi39)
                                                        default:
                                                            if ((2 != 65535) || v83) {
                                                                fmt.Println(fi40)
                                                                v41 += (v41 + v33)
                                                                v41 -= (^(v41))
                                                            }
                                                        }
                                                        var v114 float64 = float64((((0.1 / (3.14159 + ((((0.0 * ((1.5 - 1.5) - 1e10)) * 1.5) + (float32((0.1 / 1.0)) * 3.14159)) + (1e10 * 3.14159)))) - 1e10) * 0.1))
                                                        _ = v114
                                                    } else {
                                                        v115 := GenS15{F0: ("\n\t" + "hello")}
                                                        _ = v115
                                                    }
                                                }
                                            }
                                            _ = v112
                                        }
                                    }
                                    v116 := GenS16{F0: 42, F1: 42, F2: 42}
                                    _ = v116
                                    _ = v83
                                }
                            }
                        }
                        _ = v45
                    }
                    switch v1 {
                    case "a":
                        v117 := GenS17{F0: 255}
                        _ = v117
                    case "\n\t":
                        v41 += (^(v33))
                    }
                }
                v33 -= v41
                _ = fi40
                _ = v41
                _ = v42
            }
            v4 = (v4 << ((42) & 63))
            _ = fi39
        }
        if ((v4 <= (v4 + -1)) || false) {
            var v118 uint = 2
            if true {
                switch -37 {
                case 0:
                    v119 := GenS13{F0: v2}
                    _ = v119
                case -2147483648:
                    fmt.Println(v2, v3)
                }
                if false {
                    for fi120 := 0; fi120 < 16; fi120++ {
                        v33 += (^(v33))
                        if (!(v1 < v1)) {
                            fmt.Println(v1, v2)
                            if (("a" + "hello") == v1) {
                                for fi121 := 0; fi121 < 19; fi121++ {
                                    var v122 uint16 = 65535
                                    var v123 uint16 = (v122 | 327)
                                    _ = fi121
                                    _ = v122
                                    _ = v123
                                }
                                if true {
                                    for fi124 := 0; fi124 < 17; fi124++ {
                                        switch -926 {
                                        case -2147483648:
                                            v4 += v4
                                        case 42:
                                            if ((!true) && true) {
                                                v1 = (v1 + ("\n\t" + v1))
                                                if (!(!false)) {
                                                    v125 := []int{}
                                                    if true {
                                                        var v126 uint64 = 42
                                                        v118 -= (^(v118))
                                                        _ = v126
                                                    } else {
                                                        var v127 int64 = v2
                                                        _ = len(v125)
                                                        _ = v127
                                                    }
                                                    if ((!(!false)) || true) {
                                                        var v128 int16 = genFn0((int32(v3) == v4))
                                                        _ = v128
                                                    }
                                                    _ = v125
                                                } else {
                                                    v129 := []int{}
                                                    var v130 uint32 = 42
                                                    var v131 uint64 = 437
                                                    var v132 uint64 = v131
                                                    v133 := GenS18{F0: 1, F1: 65535, F2: v1}
                                                    _ = v129
                                                    _ = v130
                                                    _ = v131
                                                    _ = v132
                                                    _ = v133
                                                }
                                            } else {
                                                var v134 int16 = 1000
                                                v1 = v1
                                                _ = v134
                                            }
                                        case 32768:
                                            v4 = (v4 + (-(v4)))
                                        }
                                        _ = fi124
                                    }
                                    var v135 uint64 = 141
                                    _ = v135
                                } else {
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        var v136 int16 = 42
                                        _ = v136
                                    }
                                }
                                switch (("foo bar" + (v1 + ((v1 + v1) + v1))) + (v1 + (((v1 + "") + "") + v1))) {
                                case "foo bar":
                                    v137 := GenS6{F0: (v4 * 258), F1: (v33 - (v33 / (((v33 << ((v118) & 63))) | 1))), F2: 960}
                                    _ = v137
                                case "\n\t":
                                    v138 := []int{}
                                    _ = v138
                                }
                            }
                        } else {
                            fmt.Println(fi120, v2)
                            v139 := GenS5{F0: v4, F1: v2}
                            _ = v139
                        }
                        _ = fi120
                    }
                }
            }
            v4 = v4
            _ = v118
        }
    } else {
        var v140 uint16 = 65535
        fmt.Println(v1, v33, v33)
        _ = v140
    }
    var v141 uint32 = (v33 ^ v33)
    v3 = (v3 & (v3 & (v3 % (((v3 << ((18446744073709551615) & 63))) | 1))))
    var v142 uint = 255
    v143 := []int{2147483647}
    switch "foo bar" {
    case "foo bar":
        if (4294967295 <= 4294967295) {
            v3 = (v3 ^ (v3 + v3))
        } else {
            v144 := GenS19{F0: -32768, F1: (!true)}
            _ = v144
        }
    case "a":
        for fi145 := 0; fi145 < 7; fi145++ {
            var v146 string = v1
            if (!(false || true)) {
                var v147 uint16 = 42
                _ = v147
            }
            _ = fi145
            _ = v146
        }
    }
    fmt.Println(v142)
    for fi148 := 0; fi148 < 11; fi148++ {
        var v149 int16 = genFn0(true)
        v149 += int16(v141)
        _ = fi148
        _ = v149
    }
    var v150 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("slicelen v143:", len(v143), v143)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v141
    _ = v142
    _ = v143
    _ = v150
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
}
