package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint8; F1 uint; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint; F1 int16; F2 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS13 struct { F0 uint32; F1 uint; F2 bool }
type GenS14 struct { F0 int32 }
func genFn0() string {
    var v7 float32 = -2.25
    v7 -= 1e10
    v8 := GenS1{F0: 4294967295, F1: -1, F2: 32767}
    v9 := GenS1{F0: 2, F1: 1, F2: 991}
    v10 := []int{-2147483648, 1, 1}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ("" + "a")
}
func genFn1(p0 uint16) int32 {
    switch "" {
    case "a":
        v11 := GenS1{F0: uint32(p0), F1: 42, F2: 65535}
        _ = v11
    case "":
        v12 := []int{}
        _ = v12
    case "hello":
        p0 += 255
    default:
        var v13 uint = 667
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint64 = 4294967295
        v16 := GenS2{F0: 150, F1: 898}
        var v17 float32 = (-2.25 - 1.5)
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        if ((446 < 301) || (!false)) {
            fmt.Println(fi18)
            if (-2.25 >= (-2.25 * float32(0.0))) {
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 int32 = 42
                    v21 := []int{fi19}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                p0 = p0
                switch "a" {
                case "foo bar":
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 bool = true
                        fmt.Println(fi18, fi22)
                        p0 += (^(p0))
                        v24 := []int{(fi18 >> ((396) & 63))}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
            } else {
                p0 = (p0 / ((p0) | 1))
            }
        }
        _ = fi18
    }
    _ = p0
    return 42
}

func main() {
    var v1 int = -602
    v2 := []int{v1, -1, 127}
    var v3 uint = 820
    var v4 float32 = (3.14159 * -2.25)
    var v5 uint32 = 594
    var v6 uint8 = 127
    v25 := GenS3{F0: "foo bar", F1: 35}
    v4 = 1.5
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS3{F0: ("foo bar" + "hello"), F1: fi26}
        _ = len(v2)
        _ = fi26
        _ = v27
    }
    fmt.Println(v3, v4, v3)
    var v28 uint32 = (v5 | uint32(v6))
    v3 -= (v3 & v3)
    if true {
        var v29 string = (("" + "\n\t") + "")
        var v30 uint64 = 652
        for fi31 := 0; fi31 < 20; fi31++ {
            v32 := GenS1{F0: (v5 / ((v28) | 1)), F1: 42, F2: 0}
            for fi33 := 0; fi33 < 14; fi33++ {
                v4 = (1e10 / v4)
                _ = fi33
            }
            if (-843 >= 9223372036854775807) {
                var v34 uint = uint(v30)
                if false {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        v36 := GenS3{F0: v29, F1: (^(v1))}
                        v1 += v1
                        _ = fi35
                        _ = v36
                    }
                    var v37 uint32 = (v5 + (v5 * (v28 << (((v3 / ((v3) | 1))) & 63))))
                    var v38 int16 = -32768
                    _ = v37
                    _ = v38
                } else {
                    v39 := GenS4{F0: (^(v6)), F1: 18446744073709551615, F2: 1}
                    var v40 uint64 = v30
                    _ = v39
                    _ = v40
                }
                _ = v34
            } else {
                for fi41 := 0; fi41 < 2; fi41++ {
                    _ = len(v2)
                    v1 = v1
                    v28 -= 42
                    fmt.Println(v3, fi41)
                    _ = fi41
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v4 += 3.14159
            }
            _ = fi31
            _ = v32
        }
        _ = v29
        _ = v30
    } else {
        var v42 float32 = (v4 * 3.14159)
        switch "a" {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = -950
            }
        case "a":
            v43 := GenS5{F0: 255, F1: uint32(v3)}
            _ = v43
        default:
            for fi44 := 0; fi44 < 14; fi44++ {
                switch ("foo bar" + "hello") {
                case "hello":
                    v45 := GenS6{F0: (v3 - (^(v3)))}
                    _ = v45
                case "":
                    fmt.Println(v3)
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    if (3.14159 > (1e10 - float64(1.5))) {
                        _ = len(v2)
                    } else {
                        fmt.Println(v6, v4)
                    }
                }
                _ = fi44
            }
        }
        _ = v42
    }
    v4 = 3.14159
    if true {
        v6 = (v6 / ((v6) | 1))
        var v46 uint = 42
        _ = v46
    }
    v3 = v3
    var v47 int16 = 0
    _ = len(v2)
    var v48 uint16 = 32767
    var v49 float32 = v4
    if len(v2) > 0 {
        v2[0] = (v1 % ((v1) | 1))
    }
    var v50 int = (v1 & (v1 >> (((v3 / ((701) | 1))) & 63)))
    if (!true) {
        var v51 uint = v3
        v48 += (^(v48))
        _ = v51
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        for fi53 := 0; fi53 < 8; fi53++ {
            var v54 uint = (^(v3))
            if (0 > 1) {
                if (v49 != 0.0) {
                    if len(v2) > 0 {
                        v2[0] = (v1 + v50)
                    }
                    var v55 int16 = v47
                    switch (^(fi52)) {
                    case -1:
                        v56 := GenS7{F0: 1, F1: v54}
                        _ = v56
                    case -2147483648:
                        v3 -= (v3 ^ (v3 % ((v3) | 1)))
                        v1 = -675
                        v1 = v50
                    case 127:
                        var v57 int = fi53
                        _ = v57
                    }
                    var v58 string = ("foo bar" + (("foo bar" + ("" + "a")) + ("hello" + ("foo bar" + ("" + "foo bar")))))
                    var v59 float32 = (v4 * ((v49 * (1.5 + (1e10 - (1.5 * -2.25)))) / 3.14159))
                    _ = v55
                    _ = v58
                    _ = v59
                } else {
                    v49 -= v49
                    if (841 < v54) {
                        switch ("" + ("a" + "hello")) {
                        case "a":
                            fmt.Println(v3, v1, v48)
                        case "":
                            switch (v1 >> (((v3 & (v3 & (v54 >> (((^(v3))) & 63))))) & 63)) {
                            case 9223372036854775807:
                                v5 += (v5 << (((v3 ^ v3)) & 63))
                            case -747:
                                fmt.Println(fi53, v1)
                            case -798:
                                var v60 float32 = v49
                                _ = v60
                            default:
                                if (false || false) {
                                    fmt.Println(v54, fi52)
                                    _ = len(v2)
                                } else {
                                    var v61 int16 = (-(v47))
                                    v3 += v54
                                    var v62 uint32 = (v5 / ((v5) | 1))
                                    _ = v61
                                    _ = v62
                                }
                                v63 := GenS3{F0: "hello", F1: -268}
                                _ = v63
                            }
                        default:
                            if true {
                                v50 = (fi53 % ((-983) | 1))
                            }
                        }
                        fmt.Println(fi53)
                    } else {
                        fmt.Println(v3)
                        v64 := GenS8{F0: 45}
                        switch (fi53 - (-(v1))) {
                        case -818:
                            fmt.Println(v47, v5)
                        case -658:
                            fmt.Println(fi52, v4)
                            v2 = append(v2, (fi52 * v1))
                        case 32768:
                            if (3.14159 == (0.1 - 0.1)) {
                                if (42 <= 1) {
                                    v3 = v54
                                    switch ("" + "\n\t") {
                                    case "":
                                        if ((380 != 35) && true) {
                                            v6 += v6
                                        } else {
                                            if (!true) {
                                                if (255 != (v54 >> (((v3 * 18446744073709551615)) & 63))) {
                                                    v5 = v28
                                                    if (!((v6 << (((v54 - 487)) & 63)) > (v6 ^ 255))) {
                                                        fmt.Println(v28, v50, v50)
                                                        v1 -= (fi52 << ((v3) & 63))
                                                    }
                                                    var v65 int16 = (-(v47))
                                                    _ = v65
                                                }
                                                if false {
                                                    v3 -= 433
                                                    switch v50 {
                                                    case -1:
                                                        switch fi52 {
                                                        case 2147483647:
                                                            if ((3.14159 + ((float64(0.0) - (float64(((v4 * 3.14159) + v49)) / (((float64(v4) + 1.5) * float64(v4)) - 0.0))) * 0.1)) != 0.1) {
                                                                if ((("\n\t" + "\n\t") + ("a" + "foo bar")) != "hello") {
                                                                    _ = len(v2)
                                                                    switch ("" + "") {
                                                                    case "foo bar":
                                                                        if (true && (v4 == v49)) {
                                                                            v2 = append(v2, 32767)
                                                                            var v66 int = (fi53 + 4294967295)
                                                                            _ = v66
                                                                        } else {
                                                                            if (!((!(1 != 100)) && false)) {
                                                                                fmt.Println(v3)
                                                                                switch 0 {
                                                                                case 9223372036854775807:
                                                                                    v67 := GenS5{F0: (v48 & (v48 / (((v48 | (v48 - v48))) | 1))), F1: (v5 & v5)}
                                                                                    _ = v67
                                                                                case -410:
                                                                                    var v68 float64 = (float64(((v49 + v4) + v49)) * 0.0)
                                                                                    _ = v68
                                                                                }
                                                                                v1 += -2147483648
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        if (false && (v6 > v6)) {
                                                                            var v69 uint16 = v48
                                                                            _ = v69
                                                                        } else {
                                                                            if ((v48 <= v48) && true) {
                                                                                if (true || ((3.14159 == 1.5) && (1 == -894))) {
                                                                                    if true {
                                                                                        v47 = (v47 - (v47 / ((v47) | 1)))
                                                                                        v2 = append(v2, (^(fi53)))
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            var v70 uint16 = 213
                                                                                            _ = v70
                                                                                        case "\n\t":
                                                                                            v1 += fi52
                                                                                            if (true || (true && false)) {
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 32767
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v49, v3)
                                                                                        }
                                                                                    } else {
                                                                                        var v71 uint64 = 42
                                                                                        if ((v71 / ((uint64(v54)) | 1)) != (v71 | 1)) {
                                                                                            if (uint(v48) > (v54 + (^(v54)))) {
                                                                                                v72 := GenS9{F0: (^(v3)), F1: (v47 & v47), F2: 0}
                                                                                                _ = v72
                                                                                            } else {
                                                                                                switch (v1 / ((fi52) | 1)) {
                                                                                                case 192:
                                                                                                    var v73 int64 = -1
                                                                                                    _ = len(v2)
                                                                                                    _ = v73
                                                                                                }
                                                                                                v6 = v6
                                                                                                fmt.Println(v54, v49, v54)
                                                                                            }
                                                                                            v1 = (fi53 * int(v5))
                                                                                        }
                                                                                        _ = v71
                                                                                    }
                                                                                }
                                                                                v74 := GenS5{F0: 108, F1: v5}
                                                                                v75 := GenS1{F0: 2, F1: (v47 * (v47 ^ int16(v5))), F2: v48}
                                                                                switch (fi52 << ((65535) & 63)) {
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(fi52, v48)
                                                                                case 4294967295:
                                                                                    v76 := GenS2{F0: 2147483648, F1: v3}
                                                                                    _ = v76
                                                                                default:
                                                                                    v77 := GenS8{F0: 38}
                                                                                    _ = v77
                                                                                }
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        var v78 int32 = 693
                                                                        _ = v78
                                                                    }
                                                                }
                                                                v49 -= (v49 * ((v49 / 1.5) - 3.14159))
                                                                _ = len(v2)
                                                                v79 := GenS10{F0: -1}
                                                                _ = v79
                                                            } else {
                                                                v80 := GenS3{F0: "foo bar", F1: (-(v1))}
                                                                _ = v80
                                                            }
                                                        case 32767:
                                                            v28 += v5
                                                        case 32768:
                                                            v6 -= v6
                                                        default:
                                                            if (true && true) {
                                                                if true {
                                                                    if (v48 < uint16(v5)) {
                                                                        var v81 float64 = float64((((v4 - (3.14159 / v4)) * (v49 + (0.1 - v49))) / 0.1))
                                                                        _ = v81
                                                                    } else {
                                                                        var v82 uint16 = 255
                                                                        _ = v82
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    if (4294967295 < uint64(v5)) {
                                                                        v3 += (v3 | (v54 & 196))
                                                                    } else {
                                                                        var v83 string = (((("a" + "\n\t") + "a") + "foo bar") + "\n\t")
                                                                        v49 = 0.0
                                                                        v48 -= uint16(v6)
                                                                        if false {
                                                                            var v84 uint32 = (v5 * v28)
                                                                            v4 += (v49 * 1e10)
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi53
                                                                            }
                                                                            _ = v84
                                                                        } else {
                                                                            var v85 uint64 = 42
                                                                            v86 := GenS11{F0: v28, F1: -128}
                                                                            v87 := GenS7{F0: (fi53 << ((v54) & 63)), F1: 65535}
                                                                            _ = v85
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                        v1 += (v1 + v50)
                                                                        _ = v83
                                                                    }
                                                                }
                                                                v47 += int16(v48)
                                                                if (!true) {
                                                                    v2 = append(v2, (-(v1)))
                                                                    var v88 float32 = (-2.25 + v49)
                                                                    _ = v88
                                                                } else {
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        fmt.Println(v49)
                                                                    default:
                                                                        if false {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case 1:
                                                        v89 := GenS6{F0: (v3 & v54)}
                                                        _ = v89
                                                    default:
                                                        if (!(genFn1(153) > 42)) {
                                                            v2 = append(v2, (fi53 & (-(v50))))
                                                            var v90 int64 = 1
                                                            fmt.Println(v5, v54)
                                                            _ = v90
                                                        } else {
                                                            switch (v50 | v50) {
                                                            case 2147483647:
                                                                var v91 uint64 = 42
                                                                _ = v91
                                                            case 42:
                                                                var v92 bool = (2 >= uint8(v48))
                                                                _ = v92
                                                            }
                                                        }
                                                    }
                                                }
                                                switch 42 {
                                                case 128:
                                                    v50 += 42
                                                case -9223372036854775808:
                                                    fmt.Println(v3, v28, v49)
                                                }
                                            } else {
                                                fmt.Println(fi53, v4)
                                            }
                                            var v93 int8 = 42
                                            v54 = v54
                                            v94 := GenS4{F0: (v6 & v6), F1: v54, F2: v47}
                                            _ = v93
                                            _ = v94
                                        }
                                    default:
                                        v49 -= (1.5 + (v4 - (float32(1.5) / (v49 + 3.14159))))
                                    }
                                }
                                v28 = v5
                            } else {
                                v95 := GenS10{F0: 0}
                                _ = v95
                            }
                        default:
                            v96 := GenS12{F0: 1, F1: 0, F2: v48}
                            _ = v96
                        }
                        _ = v64
                    }
                }
            } else {
                var v97 uint8 = (v6 + v6)
                v3 -= (v3 * 255)
                _ = v97
            }
            v50 = (fi53 / ((int(v6)) | 1))
            if false {
                fmt.Println(v48, fi53, v1)
            }
            if ((true || true) || ((true && true) && (genFn1(v48) != 0))) {
                v98 := GenS13{F0: v28, F1: (v54 >> (((v54 | (^(v54)))) & 63)), F2: false}
                if true {
                    fmt.Println(v50)
                    if (127 >= -4) {
                        fmt.Println(v47, fi52)
                        if true {
                            v99 := GenS5{F0: uint16(v28), F1: v5}
                            _ = v99
                        }
                        v100 := GenS14{F0: 42}
                        _ = v100
                    }
                } else {
                    var v101 string = ((("a" + "") + "foo bar") + (("hello" + "a") + ("hello" + (("a" + "a") + (("" + "a") + (("" + "\n\t") + "hello"))))))
                    _ = v101
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v98
            }
            _ = fi53
            _ = v54
        }
        switch (v50 - 128) {
        case -424:
            if true {
                v102 := GenS13{F0: 2147483648, F1: (^(v3)), F2: (-2.25 <= ((3.14159 - (0.0 - ((1e10 - ((float64((-2.25 * ((v49 - v49) * v4))) - 3.14159) - (float64(v49) * ((0.0 / (1e10 * ((-2.25 * 1e10) / 1.0))) * 1e10)))) + (0.0 * (0.1 / 1.5))))) - ((1e10 + ((1.5 * float64(float32((float64(v4) - (1.5 * 3.14159))))) + 3.14159)) + 1.5)))}
                _ = v102
            }
        default:
            var v103 int8 = -128
            _ = v103
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
