package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 string; F1 uint8; F2 uint }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 int32; F1 uint; F2 uint32 }
type GenS12 struct { F0 uint; F1 int }
type GenS13 struct { F0 int32; F1 int; F2 uint32 }
type GenS14 struct { F0 int32; F1 uint8 }
type GenS15 struct { F0 uint16; F1 string }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint8; F1 uint16 }
type GenS18 struct { F0 int32; F1 int32 }
type GenS19 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS20 struct { F0 uint; F1 int }
type GenS21 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0() int {
    v7 := []int{1, 127, 0}
    if true {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        var v8 string = (("" + "hello") + "foo bar")
        for fi9 := 0; fi9 < 6; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                if (!false) {
                    v8 = v8
                    v8 = "\n\t"
                    v8 = "foo bar"
                }
                v8 = "hello"
                switch (-(fi9)) {
                case -1:
                    if (42 < 2147483647) {
                        v8 = v8
                        v11 := GenS1{F0: 142}
                        v12 := GenS1{F0: 0}
                        _ = v11
                        _ = v12
                    } else {
                        v13 := GenS1{F0: 42}
                        switch -705 {
                        case 9223372036854775807:
                            if false {
                                var v14 string = (v8 + "")
                                _ = len(v7)
                                v14 = (v14 + "hello")
                                _ = v14
                            } else {
                                if len(v7) > 0 {
                                    v7[0] = fi10
                                }
                            }
                        case 992:
                            switch v8 {
                            case "hello":
                                switch fi10 {
                                case 909:
                                    v7 = append(v7, (fi10 / (((fi10 ^ fi10)) | 1)))
                                default:
                                    v15 := GenS1{F0: 255}
                                    _ = v15
                                }
                            default:
                                var v16 float64 = 1.5
                                _ = v16
                            }
                            v17 := GenS1{F0: 1}
                            _ = v17
                        case 4294967295:
                            v8 = v8
                        }
                        v8 = v8
                        _ = v13
                    }
                }
                var v18 float32 = -2.25
                _ = fi10
                _ = v18
            }
            fmt.Println(v8, v8)
            var v19 int8 = 127
            _ = fi9
            _ = v19
        }
        _ = v8
    }
    var v20 float32 = 1e10
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v20
    _ = v7
    return -207
}

func main() {
    var v1 int64 = -1
    var v2 float32 = 0.1
    var v3 int16 = -283
    v4 := make(map[int]int)
    var v5 string = ((("" + "hello") + "\n\t") + ("\n\t" + "a"))
    var v6 float64 = float64(1e10)
    v1 = (v1 >> ((18446744073709551615) & 63))
    if false {
        var v21 int32 = 2147483647
        _ = v21
    }
    var v22 bool = (955 != 42)
    v22 = v22
    v23 := GenS1{F0: 89}
    if false {
        v1 = -533
        if (v22 || (v22 && (v22 || (1e10 >= (v6 + v6))))) {
            for fi24 := 0; fi24 < 11; fi24++ {
                _ = len(v4)
                _ = fi24
            }
            var v25 string = ("a" + "a")
            var v26 int16 = v3
            _ = v25
            _ = v26
        }
        if (v2 > -2.25) {
            switch "hello" {
            case "foo bar":
                v3 = v3
            default:
                var v27 bool = (false || v22)
                _ = v27
            }
            v6 = (float64((v2 / v2)) * ((1.5 / ((v6 + (float64(3.14159) * 0.1)) + v6)) * v6))
        } else {
            delete(v4, 127)
            var v28 int64 = 217
            _ = v28
        }
    }
    var v29 int32 = -1
    var v30 int32 = 117
    var v31 int = 2147483647
    fmt.Println(v6)
    switch "a" {
    case "\n\t":
        var v32 float32 = ((((v2 / 3.14159) + v2) + float32((v6 * 1.5))) - -2.25)
        _ = v32
    case "a":
        if ((^(v29)) != v29) {
            var v33 bool = v22
            _ = v33
        } else {
            if (!false) {
                if (float64(float32(v6)) < 0.0) {
                    v3 = (v3 % ((868) | 1))
                    v34 := []int{(v31 ^ (v31 * (v31 & (v31 / ((v31) | 1))))), (v31 | (v31 - genFn0())), v31}
                    v22 = true
                    _ = v34
                } else {
                    v1 -= v1
                    fmt.Println(v30, v2, v30)
                }
            } else {
                if v22 {
                    v35 := v4[(v31 % (((v31 >> ((4294967295) & 63))) | 1))]
                    _ = v35
                }
                _ = len(v4)
                v36 := GenS2{F0: v22, F1: 255, F2: 42}
                _ = v36
            }
            var v37 int16 = (v3 << ((0) & 63))
            if (v1 < (v1 * 0)) {
                v38 := []int{}
                if v22 {
                    switch v31 {
                    case 127:
                        if ((42 >= -1) && v22) {
                            for fi39 := 0; fi39 < 3; fi39++ {
                                if v22 {
                                    v40 := GenS1{F0: 176}
                                    _ = v40
                                }
                                v41 := GenS2{F0: v22, F1: 65535, F2: 32767}
                                for fi42 := 0; fi42 < 19; fi42++ {
                                    v2 += (1.5 * ((v2 + ((v2 + v2) / v2)) - 1.5))
                                    v43 := GenS3{F0: -1, F1: (((v5 + v5) + v5) + v5)}
                                    _ = fi42
                                    _ = v43
                                }
                                _ = fi39
                                _ = v41
                            }
                            v30 = (v29 % ((42) | 1))
                            v44 := v4[v31]
                            _ = v44
                        }
                    default:
                        v45 := GenS4{F0: (^(v31))}
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 14; fi46++ {
                        if v22 {
                            switch (v5 + v5) {
                            case "foo bar":
                                for fi47 := 0; fi47 < 17; fi47++ {
                                    v29 -= v30
                                    v29 -= 1
                                    v48 := GenS5{F0: -98}
                                    fmt.Println(v30, fi47)
                                    switch "" {
                                    case "hello":
                                        if false {
                                            switch (v31 * 0) {
                                            case 229:
                                                v30 -= (v29 << ((150) & 63))
                                            default:
                                                v29 = (v30 - (v30 | -2147483648))
                                                v3 = (v3 ^ v3)
                                                if (65535 <= 2) {
                                                    switch (fi47 - fi47) {
                                                    case 1:
                                                        v49 := GenS6{F0: 253}
                                                        _ = v49
                                                    case 4294967295:
                                                        v50 := GenS7{F0: v37, F1: (^(v37))}
                                                        var v51 int16 = (v3 + (v37 & -32768))
                                                        _ = v50
                                                        _ = v51
                                                    case 32767:
                                                        v30 += 42
                                                        fmt.Println(v1, fi46)
                                                    }
                                                    var v52 int16 = 0
                                                    v29 -= (v29 * (v30 >> ((2) & 63)))
                                                    _ = v52
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        switch (fi46 / ((genFn0()) | 1)) {
                                        case 128:
                                            var v53 float64 = ((1e10 / 3.14159) + ((v6 * v6) / 1e10))
                                            _ = v53
                                        }
                                    case "":
                                        v3 = v3
                                    default:
                                        var v54 float64 = 0.1
                                        _ = v54
                                    }
                                    _ = fi47
                                    _ = v48
                                }
                            case "a":
                                var v55 uint64 = 0
                                _ = v55
                            case "\n\t":
                                switch (v5 + v5) {
                                case "\n\t":
                                    if true {
                                        v2 -= 1e10
                                    }
                                default:
                                    if (2 > 4294967295) {
                                        switch v5 {
                                        case "foo bar":
                                            if (v22 || false) {
                                                if false {
                                                    fmt.Println(v29, v31)
                                                }
                                            }
                                        case "a":
                                            var v56 float64 = 3.14159
                                            _ = v56
                                        case "\n\t":
                                            fmt.Println(v30, v5, v37)
                                        default:
                                            v57 := GenS5{F0: -128}
                                            _ = v57
                                        }
                                    } else {
                                        v1 = (^(v1))
                                    }
                                }
                            default:
                                if true {
                                    for fi58 := 0; fi58 < 19; fi58++ {
                                        switch "hello" {
                                        case "a":
                                            v59 := GenS7{F0: (v3 >> ((655) & 63)), F1: int16(fi58)}
                                            _ = v59
                                        case "":
                                            v6 -= v6
                                        case "hello":
                                            fmt.Println(v37, v6, fi46)
                                        default:
                                            v30 = 2147483647
                                        }
                                        v1 -= v1
                                        _ = fi58
                                    }
                                    v31 += v31
                                } else {
                                    _ = len(v38)
                                    v29 += 652
                                    for fi60 := 0; fi60 < 16; fi60++ {
                                        fmt.Println(v29)
                                        if (true && true) {
                                            if v22 {
                                                v38 = append(v38, fi60)
                                            } else {
                                                v6 = v6
                                            }
                                            var v61 float32 = v2
                                            v61 = (0.1 - v2)
                                            _ = v61
                                        } else {
                                            var v62 uint8 = 255
                                            v63 := v4[genFn0()]
                                            v64 := GenS1{F0: 2}
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                        v3 = (v37 * (v3 / (((v37 / ((v37) | 1))) | 1)))
                                        _ = fi60
                                    }
                                }
                            }
                            if (0 > (-(fi46))) {
                                if true {
                                    switch (fi46 ^ (fi46 % (((v31 & (fi46 >> ((1) & 63)))) | 1))) {
                                    case -201:
                                        var v65 int = 569
                                        _ = v65
                                    case 32768:
                                        _ = len(v4)
                                    case 4294967295:
                                        var v66 int32 = (v29 - v29)
                                        _ = v66
                                    }
                                    v6 = float64((1.5 - v2))
                                } else {
                                    for fi67 := 0; fi67 < 11; fi67++ {
                                        v38 = append(v38, (^(fi67)))
                                        _ = fi67
                                    }
                                    switch (v31 / ((genFn0()) | 1)) {
                                    case -2147483648:
                                        v29 -= v29
                                    case -761:
                                        v31 -= 128
                                    case 4294967295:
                                        var v68 int16 = (v37 % ((v37) | 1))
                                        _ = v68
                                    }
                                }
                                for fi69 := 0; fi69 < 17; fi69++ {
                                    var v70 uint8 = 0
                                    switch v5 {
                                    case "foo bar":
                                        if v22 {
                                            var v71 int64 = 1
                                            _ = v71
                                        }
                                    case "hello":
                                        if (((v22 || ((749 != 390) || v22)) || (v22 && false)) || true) {
                                            v6 += v6
                                            v2 += float32(1e10)
                                            if v22 {
                                                fmt.Println(v37, v6, v31)
                                                v72 := GenS8{F0: "hello", F1: (v70 | (v70 | 82)), F2: 0}
                                                _ = v72
                                            } else {
                                                v73 := GenS9{F0: 42, F1: 630}
                                                switch "hello" {
                                                case "a":
                                                    v37 = (^(v37))
                                                }
                                                if v22 {
                                                    v74 := GenS10{F0: 127, F1: 99}
                                                    switch (("hello" + (("\n\t" + (v5 + "foo bar")) + v5)) + (v5 + (v5 + ("hello" + v5)))) {
                                                    case "":
                                                        fmt.Println(v3, v70)
                                                    case "hello":
                                                        var v75 float32 = (((1.5 - float32((v6 - ((v6 * (v6 - -2.25)) * v6)))) - 1.5) - v2)
                                                        fmt.Println(fi46, v75, v70)
                                                        if false {
                                                            v76 := GenS11{F0: v30, F1: 18446744073709551615, F2: 2}
                                                            delete(v4, genFn0())
                                                            _ = v76
                                                        } else {
                                                            v3 += v3
                                                            v75 = v75
                                                            switch v5 {
                                                            case "foo bar":
                                                                v77 := GenS8{F0: v5, F1: (^(v70)), F2: 193}
                                                                _ = v77
                                                            case "hello":
                                                                _ = len(v4)
                                                                var v78 int16 = int16(v70)
                                                                if v22 {
                                                                    v38 = append(v38, fi46)
                                                                    var v79 uint32 = 2147483648
                                                                    v30 -= -503
                                                                    _ = v79
                                                                } else {
                                                                    var v80 float64 = v6
                                                                    switch (fi69 ^ (fi69 | genFn0())) {
                                                                    case -9223372036854775808:
                                                                        v6 += -2.25
                                                                    default:
                                                                        if true {
                                                                            _ = len(v38)
                                                                            if false {
                                                                                v6 += (((((v6 / float64(0.1)) - ((1.5 + v80) * v6)) + float64((v2 - v2))) / ((1.5 / (1.5 + (v80 - (float64(v2) - v6)))) * (1.5 * v80))) + v6)
                                                                                var v81 float32 = v2
                                                                                fmt.Println(v30, v30, v22)
                                                                                _ = v81
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v22, v5, v75)
                                                                        }
                                                                    }
                                                                    _ = v80
                                                                }
                                                                v82 := v4[genFn0()]
                                                                _ = v78
                                                                _ = v82
                                                            default:
                                                                var v83 uint32 = 2
                                                                _ = v83
                                                            }
                                                        }
                                                        _ = v75
                                                    default:
                                                        if len(v38) > 0 {
                                                            v38[0] = genFn0()
                                                        }
                                                    }
                                                    _ = v74
                                                } else {
                                                    var v84 uint32 = 65535
                                                    if true {
                                                        v5 = ((v5 + v5) + (v5 + "hello"))
                                                        v85 := GenS12{F0: 0, F1: 0}
                                                        fmt.Println(v70, v84, fi69)
                                                        if ((v3 >> ((18446744073709551615) & 63)) >= v3) {
                                                            v84 += v84
                                                            if v22 {
                                                                switch "" {
                                                                case "hello":
                                                                    v30 = 42
                                                                case "\n\t":
                                                                    if (!v22) {
                                                                        var v86 int32 = (v29 ^ (-(v30)))
                                                                        _ = v86
                                                                    }
                                                                default:
                                                                    switch (fi46 % (((v31 / ((v31) | 1))) | 1)) {
                                                                    case -9223372036854775808:
                                                                        if v22 {
                                                                            var v87 uint32 = (v84 | 499)
                                                                            var v88 uint16 = 233
                                                                            _ = v87
                                                                            _ = v88
                                                                        } else {
                                                                            if false {
                                                                                var v89 uint16 = 42
                                                                                fmt.Println(v1)
                                                                                _ = v89
                                                                            } else {
                                                                                var v90 string = (v5 + v5)
                                                                                v91 := GenS13{F0: (v30 >> ((556) & 63)), F1: (-(fi46)), F2: 317}
                                                                                _ = v90
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                    default:
                                                                        v92 := GenS12{F0: uint(v84), F1: (v31 << ((255) & 63))}
                                                                        _ = v92
                                                                    }
                                                                }
                                                                var v93 int32 = -653
                                                                v37 += (v3 % (((v3 ^ v37)) | 1))
                                                                _ = v93
                                                            } else {
                                                                var v94 int64 = int64(fi69)
                                                                fmt.Println(v84)
                                                                if (!(-1 > v1)) {
                                                                    switch 128 {
                                                                    case 9223372036854775807:
                                                                        switch (v31 % (((v31 & (^(v31)))) | 1)) {
                                                                        case 128:
                                                                            var v95 int8 = 1
                                                                            _ = v95
                                                                        default:
                                                                            v94 = v1
                                                                        }
                                                                    case 128:
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            fmt.Println(v29, v94, v6)
                                                                        case "a":
                                                                            v4[32767] = v31
                                                                            v96 := GenS14{F0: 42, F1: v70}
                                                                            var v97 string = v5
                                                                            v98 := v4[-534]
                                                                            _ = v96
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                    case -71:
                                                                        v94 -= (v94 << ((495) & 63))
                                                                    }
                                                                    var v99 uint = 791
                                                                    _ = v99
                                                                }
                                                                _ = v94
                                                            }
                                                        } else {
                                                            v1 = v1
                                                            v37 = -32768
                                                        }
                                                        if ((v2 / v2) < v2) {
                                                            var v100 int8 = 42
                                                            var v101 int8 = v100
                                                            _ = v100
                                                            _ = v101
                                                        } else {
                                                            v37 -= (^(v37))
                                                            fmt.Println(v3)
                                                        }
                                                        _ = v85
                                                    }
                                                    var v102 string = (v5 + v5)
                                                    v6 = v6
                                                    v103 := GenS2{F0: v22, F1: 2, F2: 255}
                                                    v30 = (v30 % ((v30) | 1))
                                                    if true {
                                                        switch (^(fi69)) {
                                                        case -1:
                                                            fmt.Println(v30, v22)
                                                        case -9223372036854775808:
                                                            v104 := GenS3{F0: 1, F1: v5}
                                                            _ = v104
                                                        case 835:
                                                            if v22 {
                                                                if v22 {
                                                                    if ((v70 + (v70 * (^(v70)))) < v70) {
                                                                        var v105 int = (fi46 % (((fi46 << ((1) & 63))) | 1))
                                                                        switch v5 {
                                                                        case "\n\t":
                                                                            if (!false) {
                                                                                if true {
                                                                                    v6 = float64(0.1)
                                                                                    v106 := GenS14{F0: (^(v30)), F1: (^(v70))}
                                                                                    v22 = true
                                                                                    _ = v106
                                                                                } else {
                                                                                    if (v22 || true) {
                                                                                        fmt.Println(v102, v5, v102)
                                                                                        v107 := GenS15{F0: 304, F1: (v5 + "hello")}
                                                                                        if len(v38) > 0 {
                                                                                            v38[0] = genFn0()
                                                                                        }
                                                                                        v29 += v30
                                                                                        v105 = (fi69 * v31)
                                                                                        _ = v107
                                                                                    } else {
                                                                                        v70 = (v70 | (v70 / ((uint8(v84)) | 1)))
                                                                                        var v108 uint8 = (v70 / ((v70) | 1))
                                                                                        if (0 <= 42) {
                                                                                            delete(v4, (v31 % (((-(v31))) | 1)))
                                                                                            if len(v38) >= 2 {
                                                                                                v38 = v38[1:len(v38)]
                                                                                            }
                                                                                            if ((v37 == 1000) && (!(!(1 <= 0)))) {
                                                                                                var v109 float32 = ((v2 / v2) * v2)
                                                                                                _ = v109
                                                                                            } else {
                                                                                                if true {
                                                                                                    if v22 {
                                                                                                        _ = len(v38)
                                                                                                        var v110 int = int(v3)
                                                                                                        v37 = v37
                                                                                                        _ = v110
                                                                                                    } else {
                                                                                                        v111 := GenS16{F0: 32767}
                                                                                                        _ = v111
                                                                                                    }
                                                                                                    v112 := GenS17{F0: v70, F1: 255}
                                                                                                    _ = v112
                                                                                                }
                                                                                                var v113 int = (v31 << ((2) & 63))
                                                                                                _ = len(v38)
                                                                                                if v22 {
                                                                                                    delete(v4, -2147483648)
                                                                                                    if len(v38) >= 2 {
                                                                                                        v38 = v38[1:len(v38)]
                                                                                                    }
                                                                                                    delete(v4, (fi69 - genFn0()))
                                                                                                }
                                                                                                _ = v113
                                                                                            }
                                                                                        }
                                                                                        _ = v108
                                                                                    }
                                                                                    var v114 uint = 2
                                                                                    if false {
                                                                                        v115 := GenS18{F0: v30, F1: v29}
                                                                                        if (v29 != 1) {
                                                                                            v30 -= int32(v37)
                                                                                        }
                                                                                        _ = v115
                                                                                    } else {
                                                                                        v30 = v29
                                                                                    }
                                                                                    _ = v114
                                                                                }
                                                                                v29 -= v30
                                                                            }
                                                                        }
                                                                        if len(v38) > 0 {
                                                                            v38[0] = 0
                                                                        }
                                                                        _ = v105
                                                                    }
                                                                    var v116 bool = true
                                                                    switch 4294967295 {
                                                                    case 127:
                                                                        v4[127] = (fi46 | (fi69 | v31))
                                                                    case 0:
                                                                        fmt.Println(v102, v2, v22)
                                                                    }
                                                                    _ = v116
                                                                } else {
                                                                    fmt.Println(v2)
                                                                    var v117 int32 = int32(v84)
                                                                    _ = v117
                                                                }
                                                                v3 += 1000
                                                            }
                                                        default:
                                                            v22 = (-96 >= -128)
                                                        }
                                                        var v118 int16 = 291
                                                        v3 = -32768
                                                        _ = v118
                                                    } else {
                                                        fmt.Println(v3, v70)
                                                        if true {
                                                            fmt.Println(v102)
                                                            var v119 string = v5
                                                            var v120 int8 = 0
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                        var v121 string = ("" + (v102 + v102))
                                                        _ = v121
                                                    }
                                                    _ = v102
                                                    _ = v103
                                                    _ = v84
                                                }
                                                _ = v73
                                            }
                                        } else {
                                            if len(v38) > 0 {
                                                v38[0] = genFn0()
                                            }
                                            v122 := GenS17{F0: (v70 * (v70 & v70)), F1: 65535}
                                            _ = v122
                                        }
                                    default:
                                        v123 := GenS3{F0: 100, F1: ("\n\t" + v5)}
                                        _ = v123
                                    }
                                    if v22 {
                                        switch v5 {
                                        case "":
                                            v124 := GenS5{F0: 0}
                                            _ = v124
                                        case "hello":
                                            v125 := GenS11{F0: (v29 * (v30 | v29)), F1: 18446744073709551615, F2: 2147483647}
                                            _ = v125
                                        case "foo bar":
                                            if v22 {
                                                switch ((v5 + v5) + "\n\t") {
                                                case "\n\t":
                                                    v126 := GenS13{F0: v29, F1: v31, F2: 0}
                                                    _ = v126
                                                default:
                                                    if len(v38) > 0 {
                                                        v38[0] = v31
                                                    }
                                                }
                                                if false {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if ((v1 != -1) || true) {
                                                            v127 := GenS19{F0: v1, F1: 4294967295, F2: 0}
                                                            v30 += v29
                                                            _ = v127
                                                        } else {
                                                            fmt.Println(v3, v30)
                                                            fmt.Println(v31)
                                                            var v128 uint = 42
                                                            var v129 int = 32768
                                                            _ = v128
                                                            _ = v129
                                                        }
                                                        var v130 string = v5
                                                        v3 -= 577
                                                        var v131 uint32 = 991
                                                        var v132 uint32 = 322
                                                        _ = v130
                                                        _ = v131
                                                        _ = v132
                                                    case "a":
                                                        _ = len(v4)
                                                    default:
                                                        v29 += (v30 & v30)
                                                        var v133 int32 = v29
                                                        _ = v133
                                                    }
                                                    fmt.Println(fi69, v70)
                                                    switch v5 {
                                                    case "hello":
                                                        switch int(v3) {
                                                        case -516:
                                                            switch v5 {
                                                            case "hello":
                                                                v38 = append(v38, (fi46 - (^(fi69))))
                                                            case "\n\t":
                                                                fmt.Println(v6, fi46)
                                                                var v134 uint8 = v70
                                                                switch v5 {
                                                                case "hello":
                                                                    if true {
                                                                        var v135 uint = 18446744073709551615
                                                                        var v136 float64 = float64((0.0 - v2))
                                                                        _ = v135
                                                                        _ = v136
                                                                    }
                                                                case "":
                                                                    v137 := GenS4{F0: -282}
                                                                    _ = v137
                                                                case "a":
                                                                    var v138 uint = 4294967295
                                                                    _ = v138
                                                                }
                                                                _ = v134
                                                            default:
                                                                v37 -= (v37 << ((171) & 63))
                                                            }
                                                            v30 = v29
                                                        }
                                                    case "":
                                                        fmt.Println(v2)
                                                    case "a":
                                                        v37 = 1
                                                    }
                                                } else {
                                                    switch genFn0() {
                                                    case -62:
                                                        var v139 uint = 42
                                                        _ = v139
                                                    case 9223372036854775807:
                                                        var v140 float64 = v6
                                                        _ = v140
                                                    }
                                                    switch ((((v5 + "foo bar") + ("foo bar" + "hello")) + "") + "foo bar") {
                                                    case "foo bar":
                                                        if false {
                                                            if len(v38) > 0 {
                                                                v38[0] = fi46
                                                            }
                                                            if false {
                                                                var v141 int = (-(fi46))
                                                                if len(v38) >= 2 {
                                                                    v38 = v38[1:len(v38)]
                                                                }
                                                                var v142 uint32 = 2147483647
                                                                var v143 float64 = v6
                                                                _ = v141
                                                                _ = v142
                                                                _ = v143
                                                            } else {
                                                                var v144 float32 = 3.14159
                                                                _ = v144
                                                            }
                                                        }
                                                    case "":
                                                        v22 = v22
                                                    case "hello":
                                                        v145 := GenS20{F0: 2, F1: int(v30)}
                                                        var v146 string = v5
                                                        v3 = (-(v37))
                                                        _ = v145
                                                        _ = v146
                                                    }
                                                    if len(v38) >= 2 {
                                                        v38 = v38[1:len(v38)]
                                                    }
                                                    var v147 int32 = (v30 ^ v29)
                                                    v148 := GenS3{F0: -17, F1: v5}
                                                    _ = v147
                                                    _ = v148
                                                }
                                            } else {
                                                var v149 uint = 601
                                                if len(v38) >= 2 {
                                                    v38 = v38[1:len(v38)]
                                                }
                                                v70 += (^(v70))
                                                var v150 string = "a"
                                                var v151 int = fi46
                                                v30 = v29
                                                _ = v149
                                                _ = v150
                                                _ = v151
                                            }
                                        default:
                                            if v22 {
                                                var v152 uint16 = 65535
                                                fmt.Println(v70)
                                                switch ((v5 + v5) + v5) {
                                                case "":
                                                    var v153 int32 = (v29 | -1)
                                                    _ = v153
                                                case "\n\t":
                                                    var v154 int64 = (^(v1))
                                                    _ = v154
                                                case "foo bar":
                                                    v155 := v4[genFn0()]
                                                    _ = v155
                                                default:
                                                    delete(v4, v31)
                                                    if v22 {
                                                        v156 := GenS5{F0: int8(v3)}
                                                        var v157 uint64 = 18446744073709551615
                                                        _ = v156
                                                        _ = v157
                                                    } else {
                                                        v158 := GenS19{F0: (v1 * (-(v1))), F1: 333, F2: uint32(v152)}
                                                        v159 := GenS17{F0: v70, F1: v152}
                                                        _ = v158
                                                        _ = v159
                                                    }
                                                    var v160 int8 = 127
                                                    fmt.Println(v30, v6)
                                                    fmt.Println(v1, v30)
                                                    _ = v160
                                                }
                                                switch (v5 + ((v5 + v5) + "\n\t")) {
                                                case "foo bar":
                                                    _ = len(v38)
                                                default:
                                                    if false {
                                                        switch ("" + (v5 + v5)) {
                                                        case "foo bar":
                                                            fmt.Println(v30, v2)
                                                        }
                                                        var v161 int16 = 42
                                                        v5 = (v5 + "foo bar")
                                                        v162 := GenS1{F0: (v70 | v70)}
                                                        _ = v161
                                                        _ = v162
                                                    } else {
                                                        v22 = true
                                                        fmt.Println(v37, v152, v70)
                                                    }
                                                }
                                                v31 -= fi69
                                                var v163 uint32 = 518
                                                _ = v152
                                                _ = v163
                                            } else {
                                                v3 += 1000
                                                var v164 uint16 = 0
                                                _ = v164
                                            }
                                        }
                                        v6 = v6
                                    } else {
                                        switch -1 {
                                        case -9223372036854775808:
                                            v1 += (v1 - int64(v31))
                                        }
                                        v165 := GenS21{F0: 18446744073709551615, F1: v3, F2: 255}
                                        fmt.Println(v37)
                                        _ = v165
                                    }
                                    var v166 bool = v22
                                    _ = fi69
                                    _ = v166
                                    _ = v70
                                }
                            }
                            _ = len(v38)
                        }
                        _ = fi46
                    }
                } else {
                    var v167 int64 = -116
                    v168 := v4[v31]
                    _ = v167
                    _ = v168
                }
                v169 := GenS16{F0: 1}
                _ = v169
                _ = v38
            } else {
                delete(v4, 776)
            }
            _ = v37
        }
    case "":
        var v170 string = "foo bar"
        _ = v170
    }
    var v171 int32 = (v30 - (v29 - 42))
    fmt.Println("v1:", v1)
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v171
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
