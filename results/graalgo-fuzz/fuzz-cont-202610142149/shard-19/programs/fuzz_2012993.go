package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 int32; F1 string; F2 int32 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint8; F1 int32; F2 int8 }
type GenS13 struct { F0 int64; F1 uint8; F2 uint }
type GenS14 struct { F0 int8; F1 string }
type GenS15 struct { F0 int8; F1 string }
type GenS16 struct { F0 string; F1 uint8 }
type GenS17 struct { F0 bool; F1 bool; F2 uint64 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 int; F1 uint32 }
type GenS20 struct { F0 int8; F1 int64; F2 uint }
type GenS21 struct { F0 uint32; F1 uint }
type GenS22 struct { F0 int16; F1 int32; F2 uint8 }
type GenS23 struct { F0 string; F1 int8 }
type GenS24 struct { F0 int16; F1 int32; F2 int }
type GenS25 struct { F0 int }
type GenS26 struct { F0 string }
type GenS27 struct { F0 int }
type GenS28 struct { F0 int8; F1 int64 }
type GenS29 struct { F0 int8 }
type GenS30 struct { F0 int64; F1 int32; F2 uint8 }
type GenS31 struct { F0 uint32; F1 uint32; F2 uint }
type GenS32 struct { F0 uint64 }
type GenS33 struct { F0 uint64 }
type GenS34 struct { F0 int8; F1 uint; F2 int8 }
type GenS35 struct { F0 uint }
type GenS36 struct { F0 uint32; F1 bool; F2 int }
type GenS37 struct { F0 int8 }
type GenS38 struct { F0 int }
type GenS39 struct { F0 int16; F1 uint8 }
type GenS40 struct { F0 int; F1 int64; F2 uint64 }
type GenS41 struct { F0 bool }
type GenS42 struct { F0 int32; F1 int8; F2 int16 }
type GenS43 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS44 struct { F0 uint64; F1 uint; F2 uint }
type GenS45 struct { F0 string; F1 string; F2 int16 }
type GenS46 struct { F0 int8 }
type GenS47 struct { F0 int16 }
type GenS48 struct { F0 string; F1 bool }
type GenS49 struct { F0 int; F1 int16; F2 int64 }
type GenS50 struct { F0 int; F1 int64 }
type GenS51 struct { F0 uint; F1 int8 }
type GenS52 struct { F0 int8; F1 uint32 }
type GenS53 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS54 struct { F0 uint16; F1 uint8 }
type GenS55 struct { F0 uint64 }
type GenS56 struct { F0 uint; F1 uint }
type GenS57 struct { F0 string; F1 bool; F2 int }
func genFn0(p0 int32, p1 int8) uint8 {
    var v7 float32 = float32(float64(float32(float64(((float32(1e10) + 1.5) + (float32((0.1 * 0.0)) / (float32((float64((-2.25 / 0.1)) - 0.1)) - (0.0 - 1.5))))))))
    p0 -= (p0 & (p0 & (^(p0))))
    p0 = p0
    for fi8 := 0; fi8 < 20; fi8++ {
        if (4294967295 < 2147483648) {
            for fi9 := 0; fi9 < 14; fi9++ {
                if (1.5 >= (float64((float32((1.5 - (-2.25 - ((3.14159 * 3.14159) - 1.5)))) - 0.0)) - -2.25)) {
                    v10 := GenS1{F0: 32768, F1: 125}
                    v11 := []int{(fi8 + -2147483648), fi9, (fi8 / (((fi8 * fi8)) | 1)), 42}
                    _ = v10
                    _ = v11
                } else {
                    if (0 > 0) {
                        if true {
                            if (4294967295 >= 403) {
                                p1 += p1
                                v12 := GenS2{F0: (p0 >> ((65535) & 63))}
                                var v13 uint64 = 466
                                if false {
                                    var v14 uint16 = 864
                                    var v15 int32 = (p0 / ((p0) | 1))
                                    v16 := GenS3{F0: -1}
                                    fmt.Println(fi8, v13)
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                _ = v12
                                _ = v13
                            }
                            v17 := GenS3{F0: 0}
                            _ = v17
                        } else {
                            switch "\n\t" {
                            case "hello":
                                v18 := []int{fi8, (fi8 / (((fi9 % (((fi9 ^ (-(fi9)))) | 1))) | 1)), 0, (fi9 >> ((4294967295) & 63))}
                                _ = v18
                            }
                        }
                    } else {
                        var v19 int64 = 42
                        _ = v19
                    }
                    switch (fi9 << ((990) & 63)) {
                    case 417:
                        fmt.Println(fi9, p0)
                    default:
                        v20 := []int{(fi8 * (fi8 << ((1) & 63))), fi8}
                        _ = v20
                    }
                }
                _ = fi9
            }
        } else {
            fmt.Println(v7)
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 183
}

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]int)
    var v3 uint32 = 2
    var v4 float32 = -2.25
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 42
    switch 32767 {
    case 690:
        var v21 uint64 = 1
        _ = v21
    case -2147483648:
        v22 := []int{-350, -2147483648}
        _ = v22
    case 32767:
        v23 := []int{1, -1, -69, -357}
        _ = v23
    default:
        delete(v2, (v3 + (v3 | 65535)))
    }
    var v24 uint8 = genFn0(0, (-(v1)))
    var v25 string = "a"
    delete(v2, (v3 << ((18446744073709551615) & 63)))
    fmt.Println(v24)
    for fi26 := 0; fi26 < 14; fi26++ {
        switch fi26 {
        case 9223372036854775807:
            if ("\n\t" >= (v25 + v25)) {
                switch (fi26 + fi26) {
                case 32767:
                    if (v1 > v1) {
                        v27 := GenS4{F0: 752, F1: (^(fi26))}
                        for fi28 := 0; fi28 < 6; fi28++ {
                            fmt.Println(fi26, v6)
                            v4 = 1.5
                            _ = fi28
                        }
                        _ = v27
                    }
                case 1:
                    var v29 int64 = 0
                    _ = v29
                default:
                    v3 += (v3 & v3)
                }
                for fi30 := 0; fi30 < 4; fi30++ {
                    fmt.Println(v3)
                    _ = fi30
                }
                switch fi26 {
                case 218:
                    delete(v5, 30)
                case -1:
                    var v31 int16 = -1
                    _ = v31
                case -9223372036854775808:
                    switch fi26 {
                    case 32768:
                        delete(v2, v3)
                    case 420:
                        v1 += v1
                    default:
                        fmt.Println(fi26, v1, fi26)
                    }
                }
            } else {
                delete(v5, v3)
                if (!(false && (42 > 32767))) {
                    v32 := GenS1{F0: fi26, F1: 42}
                    _ = v32
                } else {
                    v1 = v1
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v1 -= v1
                        var v34 int = fi26
                        var v35 int16 = -522
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            }
        case 0:
            if true {
                for fi36 := 0; fi36 < 3; fi36++ {
                    v37 := GenS1{F0: 4294967295, F1: 0}
                    if (255 != 1) {
                        switch (fi26 & (-(fi26))) {
                        case 113:
                            v4 = v4
                        case -889:
                            v38 := GenS2{F0: -1}
                            _ = v38
                        }
                    }
                    _ = fi36
                    _ = v37
                }
                for fi39 := 0; fi39 < 20; fi39++ {
                    v5[0] = v3
                    v40 := GenS5{F0: v3, F1: (v1 / (((-(v1))) | 1))}
                    var v41 int16 = 1000
                    v5[(v3 >> ((65535) & 63))] = v3
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                var v42 int = (-(fi26))
                _ = v42
            }
        case -327:
            switch v25 {
            case "":
                for fi43 := 0; fi43 < 16; fi43++ {
                    v6 += v24
                    var v44 float64 = -2.25
                    if (42 != fi26) {
                        var v45 float64 = 0.1
                        _ = v45
                    } else {
                        v6 -= v24
                        if true {
                            if true {
                                if true {
                                    v4 += (float32(((-2.25 - (float64((float32((v44 / (-2.25 + ((v44 + 3.14159) - (3.14159 / ((v44 - float64((float32(v44) + (v4 / float32((v44 / v44)))))) * (v44 + v44))))))) * v4)) + (v44 + v44))) * v44)) / float32(((1e10 * 1e10) / v44)))
                                    v44 = (v44 * 0.0)
                                    var v46 int64 = -339
                                    _ = v46
                                }
                                var v47 float32 = (v4 - v4)
                                _ = v47
                            } else {
                                switch v25 {
                                case "hello":
                                    v1 -= v1
                                case "\n\t":
                                    v44 -= v44
                                case "foo bar":
                                    switch fi43 {
                                    case 4294967295:
                                        if true {
                                            var v48 string = v25
                                            _ = v48
                                        } else {
                                            v49 := GenS5{F0: (^(v3)), F1: (v1 + (v1 - (^(v1))))}
                                            _ = v49
                                        }
                                    case -9223372036854775808:
                                        v6 += (v24 % ((uint8(v3)) | 1))
                                    default:
                                        if (190 > 0) {
                                            if true {
                                                if false {
                                                    if true {
                                                        v1 = (v1 % ((v1) | 1))
                                                        switch (fi26 * fi26) {
                                                        case 127:
                                                            fmt.Println(v3)
                                                            if (((v25 + (("hello" + v25) + (v25 + "hello"))) + v25) == "") {
                                                                v50 := GenS2{F0: 42}
                                                                switch "foo bar" {
                                                                case "":
                                                                    v51 := []int{9223372036854775807, fi43, (fi26 << ((18446744073709551615) & 63))}
                                                                    _ = v51
                                                                default:
                                                                    v44 += v44
                                                                }
                                                                if ((-2.25 + float64(v4)) >= v44) {
                                                                    v25 = (v25 + "\n\t")
                                                                    v1 = v1
                                                                }
                                                                v52 := GenS5{F0: v3, F1: v1}
                                                                _ = v50
                                                                _ = v52
                                                            }
                                                            v44 = -2.25
                                                        case -561:
                                                            switch v25 {
                                                            case "hello":
                                                                v53 := GenS6{F0: v1, F1: -729}
                                                                _ = v53
                                                            default:
                                                                var v54 int16 = 1
                                                                _ = v54
                                                            }
                                                        case 1:
                                                            v55 := v2[0]
                                                            _ = v55
                                                        default:
                                                            v56 := []int{fi43}
                                                            _ = v56
                                                        }
                                                    } else {
                                                        v1 -= (v1 ^ (v1 & (v1 | v1)))
                                                    }
                                                } else {
                                                    v3 += (^(v3))
                                                }
                                                var v57 int8 = 42
                                                _ = v57
                                            } else {
                                                if (true && (!(800 != 0))) {
                                                    fmt.Println(fi43)
                                                    var v58 int64 = 1
                                                    _ = v58
                                                } else {
                                                    v59 := GenS7{F0: (408 == 0), F1: -2147483648}
                                                    _ = v59
                                                }
                                                v60 := GenS2{F0: int32(v1)}
                                                _ = v60
                                            }
                                            if (!false) {
                                                var v61 int16 = 0
                                                switch v25 {
                                                case "a":
                                                    if false {
                                                        var v62 int64 = -9223372036854775808
                                                        var v63 int32 = 42
                                                        var v64 string = ("" + v25)
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v25 = ((v25 + "hello") + v25)
                                                        if false {
                                                            v65 := []int{42, (fi26 >> ((4294967295) & 63))}
                                                            v66 := GenS8{F0: (v61 + (v61 * (v61 * (v61 & (v61 & v61))))), F1: 42}
                                                            if (((v3 % ((v3) | 1)) >= v3) || (true || ((3.14159 + v4) >= (0.0 / v4)))) {
                                                                if false {
                                                                    fmt.Println(v1, v6, v1)
                                                                    if (849 <= 458) {
                                                                        v67 := v2[2147483648]
                                                                        _ = v67
                                                                    }
                                                                    if false {
                                                                        v24 = (v24 & (^(v6)))
                                                                    }
                                                                    var v68 int = (fi43 >> ((255) & 63))
                                                                    v69 := v2[(v3 >> ((906) & 63))]
                                                                    var v70 bool = false
                                                                    _ = v68
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                            } else {
                                                                v3 = (v3 & 4294967295)
                                                                v4 = (1e10 / v4)
                                                                if len(v65) > 0 {
                                                                    v65[0] = fi43
                                                                }
                                                                v71 := GenS9{F0: -1, F1: 1}
                                                                _ = v71
                                                            }
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                        if (v3 != (v3 / ((v3) | 1))) {
                                                            v72 := v5[v3]
                                                            v73 := GenS10{F0: 0, F1: v25, F2: 1}
                                                            _ = v72
                                                            _ = v73
                                                        } else {
                                                            fmt.Println(v44)
                                                            var v74 int16 = v61
                                                            if (((fi43 * fi43) <= fi43) && false) {
                                                                fmt.Println(fi26, fi26)
                                                            } else {
                                                                fmt.Println(v3)
                                                                switch fi43 {
                                                                case 2147483647:
                                                                    if false {
                                                                        v75 := []int{(fi43 - 368), (fi43 | -9223372036854775808)}
                                                                        _ = v75
                                                                    } else {
                                                                        v76 := GenS3{F0: (v74 << ((uint(v6)) & 63))}
                                                                        v24 = genFn0(42, 100)
                                                                        _ = v76
                                                                    }
                                                                case -272:
                                                                    v77 := GenS5{F0: v3, F1: (v1 >> ((891) & 63))}
                                                                    _ = v77
                                                                default:
                                                                    if true {
                                                                        if true {
                                                                            fmt.Println(v4, v3)
                                                                            v78 := GenS6{F0: 79, F1: -1}
                                                                            fmt.Println(v61, v74)
                                                                            _ = v78
                                                                        }
                                                                        delete(v2, (v3 + v3))
                                                                    } else {
                                                                        delete(v5, (v3 ^ 638))
                                                                        v79 := GenS11{F0: v61}
                                                                        switch 128 {
                                                                        case 9223372036854775807:
                                                                            switch (fi26 * (fi43 / ((-2147483648) | 1))) {
                                                                            case 128:
                                                                                v3 = v3
                                                                            case -2147483648:
                                                                                v80 := GenS3{F0: -819}
                                                                                switch fi43 {
                                                                                case 32768:
                                                                                    if false {
                                                                                        var v81 int8 = 100
                                                                                        _ = v81
                                                                                    } else {
                                                                                        if true {
                                                                                            fmt.Println(v24)
                                                                                            v82 := v2[(v3 * (v3 + v3))]
                                                                                            fmt.Println(fi43, v6, v74)
                                                                                            _ = v82
                                                                                        } else {
                                                                                            fmt.Println(v3)
                                                                                            switch (^(fi26)) {
                                                                                            case 0:
                                                                                                var v83 uint = 887
                                                                                                _ = v83
                                                                                            default:
                                                                                                v24 = v6
                                                                                            }
                                                                                            v6 += v24
                                                                                        }
                                                                                        fmt.Println(v6, v4)
                                                                                    }
                                                                                default:
                                                                                    v25 = v25
                                                                                    fmt.Println(v74, v25)
                                                                                }
                                                                                var v84 uint32 = (v3 / ((v3) | 1))
                                                                                _ = v80
                                                                                _ = v84
                                                                            case 1:
                                                                                var v85 uint8 = (v6 | (v24 * 2))
                                                                                _ = v85
                                                                            default:
                                                                                switch (v25 + v25) {
                                                                                case "foo bar":
                                                                                    v25 = ""
                                                                                }
                                                                            }
                                                                        case -1:
                                                                            v1 = 1
                                                                        case 780:
                                                                            v86 := GenS5{F0: 865, F1: 0}
                                                                            _ = v86
                                                                        default:
                                                                            fmt.Println(v61, v6, v6)
                                                                        }
                                                                        v44 = v44
                                                                        var v87 int8 = 117
                                                                        _ = v79
                                                                        _ = v87
                                                                    }
                                                                }
                                                            }
                                                            _ = v74
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v88 bool = false
                                                    _ = v88
                                                case "foo bar":
                                                    v4 -= 1.5
                                                }
                                                _ = v61
                                            }
                                            var v89 uint32 = v3
                                            var v90 uint32 = v3
                                            v2[(v90 + (v90 >> ((1) & 63)))] = int(v1)
                                            v91 := GenS1{F0: fi26, F1: 65535}
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                        } else {
                                            switch v25 {
                                            case "a":
                                                if (698 > 2) {
                                                    var v92 int = fi43
                                                    fmt.Println(v6)
                                                    _ = len(v2)
                                                    _ = v92
                                                }
                                            case "hello":
                                                if (!(1 == 1)) {
                                                    v93 := GenS4{F0: 65535, F1: (fi43 >> ((255) & 63))}
                                                    if true {
                                                        var v94 uint64 = 42
                                                        if (-773 == (^(fi26))) {
                                                            v44 += v44
                                                            var v95 int16 = 0
                                                            switch (^(fi26)) {
                                                            case 1:
                                                                switch v25 {
                                                                case "":
                                                                    if false {
                                                                        v96 := GenS12{F0: (v6 * (v24 * genFn0(0, v1))), F1: 2147483647, F2: v1}
                                                                        if false {
                                                                            var v97 bool = false
                                                                            _ = v97
                                                                        }
                                                                        delete(v2, (v3 * v3))
                                                                        if (false && (false || true)) {
                                                                            v98 := GenS13{F0: -9223372036854775808, F1: (v24 % ((v24) | 1)), F2: 974}
                                                                            _ = v98
                                                                        } else {
                                                                            v99 := GenS13{F0: 0, F1: uint8(v3), F2: 42}
                                                                            var v100 int8 = v1
                                                                            _ = v100
                                                                            _ = v99
                                                                        }
                                                                        _ = v96
                                                                    } else {
                                                                        if (int8(v95) < int8(v94)) {
                                                                            switch (fi26 >> ((255) & 63)) {
                                                                            case -313:
                                                                                if true {
                                                                                    v101 := GenS14{F0: 84, F1: v25}
                                                                                    v3 -= (^(v3))
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        var v102 uint16 = 0
                                                                                        _ = v102
                                                                                    }
                                                                                    switch fi43 {
                                                                                    case -1:
                                                                                        if true {
                                                                                            v103 := v5[v3]
                                                                                            _ = v103
                                                                                        } else {
                                                                                            var v104 int64 = -827
                                                                                            v105 := GenS2{F0: 117}
                                                                                            if (!true) {
                                                                                                switch fi26 {
                                                                                                case 42:
                                                                                                    fmt.Println(v94)
                                                                                                case 32767:
                                                                                                    v106 := GenS15{F0: v1, F1: v25}
                                                                                                    _ = v106
                                                                                                case 152:
                                                                                                    if true {
                                                                                                        fmt.Println(v94, v1, v3)
                                                                                                        var v107 uint = 0
                                                                                                        _ = v107
                                                                                                    } else {
                                                                                                        if (!true) {
                                                                                                            v104 = int64(v94)
                                                                                                            v3 -= v3
                                                                                                        } else {
                                                                                                            v44 += float64(v4)
                                                                                                            v108 := []int{(fi26 & fi26), -263, -841, (fi26 | (^(fi26)))}
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        v104 += (-(v104))
                                                                                                        v109 := []int{(fi26 << ((42) & 63))}
                                                                                                        _ = v109
                                                                                                    }
                                                                                                    if false {
                                                                                                        fmt.Println(v1, v4, v104)
                                                                                                        v110 := []int{fi43}
                                                                                                        var v111 uint16 = 65535
                                                                                                        _ = v110
                                                                                                        _ = v111
                                                                                                    }
                                                                                                default:
                                                                                                    var v112 uint8 = v24
                                                                                                    _ = v112
                                                                                                }
                                                                                            }
                                                                                            v113 := []int{}
                                                                                            v104 = (-(v104))
                                                                                            if true {
                                                                                                delete(v5, (v3 | uint32(v94)))
                                                                                                var v114 uint8 = v24
                                                                                                var v115 float64 = v44
                                                                                                _ = v114
                                                                                                _ = v115
                                                                                            } else {
                                                                                                if (1 > 127) {
                                                                                                    if ((v6 | v24) < v24) {
                                                                                                        v116 := GenS12{F0: (v6 * v24), F1: 861, F2: (v1 + (v1 / ((v1) | 1)))}
                                                                                                        _ = v116
                                                                                                    }
                                                                                                } else {
                                                                                                    if (!((v3 / (((v3 << ((42) & 63))) | 1)) != v3)) {
                                                                                                        switch -9223372036854775808 {
                                                                                                        case 32767:
                                                                                                            var v117 uint = 202
                                                                                                            _ = v117
                                                                                                        case -2147483648:
                                                                                                            fmt.Println(v6, v3, v95)
                                                                                                        default:
                                                                                                            v3 = (v3 * (v3 * v3))
                                                                                                        }
                                                                                                        fmt.Println(v3, fi26)
                                                                                                    }
                                                                                                    var v118 int8 = v1
                                                                                                    switch fi43 {
                                                                                                    case 1:
                                                                                                        v95 = (^(v95))
                                                                                                    case 127:
                                                                                                        if false {
                                                                                                            if ((!(42 >= 964)) || (int(v104) < fi26)) {
                                                                                                                var v119 float64 = (v44 / 0.1)
                                                                                                                switch (v25 + ("hello" + v25)) {
                                                                                                                case "":
                                                                                                                    v120 := GenS13{F0: 1, F1: v24, F2: 156}
                                                                                                                    _ = v120
                                                                                                                case "hello":
                                                                                                                    var v121 int8 = v1
                                                                                                                    _ = v121
                                                                                                                case "a":
                                                                                                                    if (!true) {
                                                                                                                        var v122 uint8 = 42
                                                                                                                        fmt.Println(v6, fi26)
                                                                                                                        v123 := v5[(v3 + (^(v3)))]
                                                                                                                        if false {
                                                                                                                            if false {
                                                                                                                                v123 += v3
                                                                                                                                var v124 int16 = v95
                                                                                                                                var v125 string = v25
                                                                                                                                _ = v124
                                                                                                                                _ = v125
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                v118 -= 1
                                                                                                                            } else {
                                                                                                                                v119 -= float64(0.0)
                                                                                                                            }
                                                                                                                            v126 := GenS16{F0: (v25 + "a"), F1: (v6 ^ 0)}
                                                                                                                            switch "foo bar" {
                                                                                                                            case "foo bar":
                                                                                                                                if (!true) {
                                                                                                                                    v4 += float32(v44)
                                                                                                                                    _ = len(v2)
                                                                                                                                    switch ("\n\t" + "") {
                                                                                                                                    case "":
                                                                                                                                        v127 := v2[v123]
                                                                                                                                        _ = v127
                                                                                                                                    default:
                                                                                                                                        switch "" {
                                                                                                                                        case "":
                                                                                                                                            var v128 uint32 = (^(v123))
                                                                                                                                            _ = v128
                                                                                                                                        case "a":
                                                                                                                                            v122 += (v24 % (((v6 * v122)) | 1))
                                                                                                                                        default:
                                                                                                                                            v5[v3] = 688
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    var v129 int16 = (v95 / ((v95) | 1))
                                                                                                                                    _ = v129
                                                                                                                                }
                                                                                                                            case "\n\t":
                                                                                                                                if false {
                                                                                                                                    v130 := GenS17{F0: false, F1: (1e10 != 3.14159), F2: 376}
                                                                                                                                    v2[(v3 - v123)] = fi43
                                                                                                                                    v131 := GenS18{F0: 1}
                                                                                                                                    _ = v130
                                                                                                                                    _ = v131
                                                                                                                                } else {
                                                                                                                                    switch v25 {
                                                                                                                                    case "foo bar":
                                                                                                                                        if ((18446744073709551615 > 0) || false) {
                                                                                                                                            fmt.Println(v3, v122, v123)
                                                                                                                                            if ((true && false) && ("" < "foo bar")) {
                                                                                                                                                var v132 uint32 = (v123 & 2147483647)
                                                                                                                                                v133 := GenS19{F0: fi26, F1: v3}
                                                                                                                                                _ = v132
                                                                                                                                                _ = v133
                                                                                                                                            }
                                                                                                                                            v123 = uint32(v94)
                                                                                                                                            if true {
                                                                                                                                                v119 = v44
                                                                                                                                            } else {
                                                                                                                                                if true {
                                                                                                                                                    v3 -= (v3 * v3)
                                                                                                                                                } else {
                                                                                                                                                    if len(v113) >= 2 {
                                                                                                                                                        v113 = v113[1:len(v113)]
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                v122 = 169
                                                                                                                                                if true {
                                                                                                                                                    delete(v2, (v123 / (((v123 | (v3 % ((v3) | 1)))) | 1)))
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            switch "" {
                                                                                                                                            case "hello":
                                                                                                                                                switch (fi26 & 32767) {
                                                                                                                                                case 0:
                                                                                                                                                    if ((v3 - v3) != 65535) {
                                                                                                                                                        switch (((v25 + v25) + ((v25 + ((v25 + (("foo bar" + v25) + ("a" + v25))) + v25)) + v25)) + v25) {
                                                                                                                                                        case "\n\t":
                                                                                                                                                            delete(v2, 463)
                                                                                                                                                        case "foo bar":
                                                                                                                                                            v119 = v44
                                                                                                                                                        }
                                                                                                                                                        v134 := GenS20{F0: -1, F1: 42, F2: 707}
                                                                                                                                                        _ = v134
                                                                                                                                                    } else {
                                                                                                                                                        v104 = v104
                                                                                                                                                        v135 := GenS10{F0: -351, F1: v25, F2: 2147483647}
                                                                                                                                                        if true {
                                                                                                                                                            switch (fi26 / ((fi43) | 1)) {
                                                                                                                                                            case -614:
                                                                                                                                                                var v136 uint64 = v94
                                                                                                                                                                _ = v136
                                                                                                                                                            }
                                                                                                                                                            fmt.Println(v94)
                                                                                                                                                            v137 := GenS19{F0: (fi26 & (^(fi43))), F1: (v123 >> ((1) & 63))}
                                                                                                                                                            var v138 uint32 = (^(v3))
                                                                                                                                                            _ = v137
                                                                                                                                                            _ = v138
                                                                                                                                                        }
                                                                                                                                                        fmt.Println(v6)
                                                                                                                                                        _ = v135
                                                                                                                                                    }
                                                                                                                                                case 42:
                                                                                                                                                    fmt.Println(fi26, v24, fi26)
                                                                                                                                                }
                                                                                                                                            case "":
                                                                                                                                                if false {
                                                                                                                                                    var v139 uint64 = (v94 % (((v94 & v94)) | 1))
                                                                                                                                                    v118 = v1
                                                                                                                                                    if (v118 > v1) {
                                                                                                                                                        var v140 int = (fi43 & (^(fi26)))
                                                                                                                                                        var v141 uint = 736
                                                                                                                                                        var v142 int16 = v95
                                                                                                                                                        _ = v140
                                                                                                                                                        _ = v141
                                                                                                                                                        _ = v142
                                                                                                                                                    } else {
                                                                                                                                                        var v143 uint8 = (v122 % (((^(v24))) | 1))
                                                                                                                                                        var v144 float32 = ((v4 * v4) + (v4 * (3.14159 * v4)))
                                                                                                                                                        _ = v143
                                                                                                                                                        _ = v144
                                                                                                                                                    }
                                                                                                                                                    _ = v139
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            var v145 float32 = float32((1.5 * v44))
                                                                                                                                            _ = v145
                                                                                                                                        }
                                                                                                                                    case "hello":
                                                                                                                                        v146 := GenS17{F0: true, F1: true, F2: v94}
                                                                                                                                        _ = v146
                                                                                                                                    case "":
                                                                                                                                        var v147 int64 = 1
                                                                                                                                        _ = v147
                                                                                                                                    }
                                                                                                                                    var v148 int16 = (v95 % (((-(v95))) | 1))
                                                                                                                                    _ = v148
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                _ = len(v2)
                                                                                                                            }
                                                                                                                            _ = v126
                                                                                                                        }
                                                                                                                        v113 = append(v113, 458)
                                                                                                                        fmt.Println(v1, fi43, v104)
                                                                                                                        if true {
                                                                                                                            var v149 uint64 = v94
                                                                                                                            v150 := GenS21{F0: v3, F1: 0}
                                                                                                                            if false {
                                                                                                                                var v151 int64 = (v104 << ((0) & 63))
                                                                                                                                fmt.Println(v1)
                                                                                                                                if (((-2.25 + (v119 / (v119 - v44))) * (v44 + 3.14159)) < v44) {
                                                                                                                                    v1 = (-(v1))
                                                                                                                                    if ((!false) && ((v151 | 0) != int64(v95))) {
                                                                                                                                        v152 := v2[(v3 + v3)]
                                                                                                                                        _ = v152
                                                                                                                                    }
                                                                                                                                    var v153 uint64 = 18446744073709551615
                                                                                                                                    _ = v153
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v94, v151, v24)
                                                                                                                                    var v154 string = (v25 + v25)
                                                                                                                                    switch (fi43 >> ((0) & 63)) {
                                                                                                                                    case 32768:
                                                                                                                                        delete(v5, (v3 / ((v3) | 1)))
                                                                                                                                    case 42:
                                                                                                                                        if true {
                                                                                                                                            v155 := GenS22{F0: (v95 | v95), F1: 2147483647, F2: genFn0(329, v1)}
                                                                                                                                            _ = v155
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v119, v123)
                                                                                                                                        }
                                                                                                                                    case 4294967295:
                                                                                                                                        v154 = "hello"
                                                                                                                                    }
                                                                                                                                    fmt.Println(v104, v149, v94)
                                                                                                                                    var v156 uint8 = genFn0(1, (v1 / ((-128) | 1)))
                                                                                                                                    _ = v154
                                                                                                                                    _ = v156
                                                                                                                                }
                                                                                                                                switch (fi26 % ((32768) | 1)) {
                                                                                                                                case 42:
                                                                                                                                    v3 += 2147483648
                                                                                                                                case -2147483648:
                                                                                                                                    v157 := GenS23{F0: ((v25 + (v25 + "foo bar")) + ("a" + (v25 + v25))), F1: (v1 + v1)}
                                                                                                                                    fmt.Println(v6, v25)
                                                                                                                                    if ((fi43 > 359) || (v94 >= v94)) {
                                                                                                                                        fmt.Println(v151, v122, fi43)
                                                                                                                                        fmt.Println(v3, v104, v6)
                                                                                                                                        var v158 int = (fi26 % ((127) | 1))
                                                                                                                                        _ = v158
                                                                                                                                    }
                                                                                                                                    var v159 int64 = (v151 | (v104 & 614))
                                                                                                                                    _ = v157
                                                                                                                                    _ = v159
                                                                                                                                case -9223372036854775808:
                                                                                                                                    if true {
                                                                                                                                        var v160 uint32 = v3
                                                                                                                                        v25 = (v25 + v25)
                                                                                                                                        _ = v160
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    if false {
                                                                                                                                        v161 := GenS24{F0: -1, F1: 2147483647, F2: fi43}
                                                                                                                                        v162 := GenS3{F0: 32767}
                                                                                                                                        v163 := GenS25{F0: fi43}
                                                                                                                                        _ = v161
                                                                                                                                        _ = v162
                                                                                                                                        _ = v163
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                if true {
                                                                                                                                    v164 := GenS19{F0: fi43, F1: (v3 * v123)}
                                                                                                                                    var v165 uint16 = 42
                                                                                                                                    _ = v164
                                                                                                                                    _ = v165
                                                                                                                                }
                                                                                                                                var v166 int64 = (v151 >> ((116) & 63))
                                                                                                                                _ = v151
                                                                                                                                _ = v166
                                                                                                                            } else {
                                                                                                                                if (v25 < "hello") {
                                                                                                                                    fmt.Println(v1)
                                                                                                                                    v167 := GenS26{F0: ""}
                                                                                                                                    _ = v167
                                                                                                                                } else {
                                                                                                                                    if false {
                                                                                                                                        var v168 uint64 = 18446744073709551615
                                                                                                                                        var v169 uint64 = (v94 + v94)
                                                                                                                                        _ = v168
                                                                                                                                        _ = v169
                                                                                                                                    } else {
                                                                                                                                        var v170 bool = true
                                                                                                                                        switch ((v25 + "foo bar") + v25) {
                                                                                                                                        case "foo bar":
                                                                                                                                            switch v25 {
                                                                                                                                            case "":
                                                                                                                                                if false {
                                                                                                                                                    v6 = (v122 % ((v6) | 1))
                                                                                                                                                } else {
                                                                                                                                                    v3 = (^(v3))
                                                                                                                                                }
                                                                                                                                            case "\n\t":
                                                                                                                                                v24 -= (v6 * genFn0(-1, (v1 + v118)))
                                                                                                                                            }
                                                                                                                                        case "a":
                                                                                                                                            v44 += v119
                                                                                                                                        case "hello":
                                                                                                                                            if (v4 == 3.14159) {
                                                                                                                                                v2[(v3 * (v3 - (v3 * (v3 / ((769) | 1)))))] = 2147483647
                                                                                                                                                if (v170 && ("\n\t" >= (v25 + v25))) {
                                                                                                                                                    var v171 int64 = (v104 << ((0) & 63))
                                                                                                                                                    _ = v171
                                                                                                                                                } else {
                                                                                                                                                    var v172 bool = v170
                                                                                                                                                    if len(v113) >= 2 {
                                                                                                                                                        v113 = v113[1:len(v113)]
                                                                                                                                                    }
                                                                                                                                                    v44 = v119
                                                                                                                                                    _ = v172
                                                                                                                                                }
                                                                                                                                                _ = len(v113)
                                                                                                                                            } else {
                                                                                                                                                fmt.Println(fi26)
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            if v170 {
                                                                                                                                                if len(v113) > 0 {
                                                                                                                                                    v113[0] = 9223372036854775807
                                                                                                                                                }
                                                                                                                                                var v173 uint16 = 1
                                                                                                                                                _ = v173
                                                                                                                                            } else {
                                                                                                                                                switch v25 {
                                                                                                                                                case "":
                                                                                                                                                    if false {
                                                                                                                                                        if true {
                                                                                                                                                            v123 = 42
                                                                                                                                                            var v174 int64 = v104
                                                                                                                                                            _ = v174
                                                                                                                                                        }
                                                                                                                                                        v94 -= (v94 - (^(v149)))
                                                                                                                                                    } else {
                                                                                                                                                        var v175 int64 = v104
                                                                                                                                                        var v176 int32 = 1
                                                                                                                                                        if v170 {
                                                                                                                                                            v149 -= 4294967295
                                                                                                                                                        } else {
                                                                                                                                                            switch (("\n\t" + (v25 + v25)) + ("a" + "a")) {
                                                                                                                                                            case "foo bar":
                                                                                                                                                                switch (((v25 + v25) + (v25 + v25)) + (v25 + v25)) {
                                                                                                                                                                case "hello":
                                                                                                                                                                    var v177 float32 = (((0.0 - v4) * v4) + v4)
                                                                                                                                                                    _ = v177
                                                                                                                                                                case "":
                                                                                                                                                                    if (((v1 << ((0) & 63)) <= (v118 ^ (v1 / (((^(v1))) | 1)))) && v170) {
                                                                                                                                                                        var v178 uint64 = (^(v94))
                                                                                                                                                                        _ = v178
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                if false {
                                                                                                                                                                    v123 = v3
                                                                                                                                                                    if (v122 >= (v122 % (((v24 | 2)) | 1))) {
                                                                                                                                                                        var v179 uint32 = 1
                                                                                                                                                                        _ = v179
                                                                                                                                                                    } else {
                                                                                                                                                                        var v180 int8 = -1
                                                                                                                                                                        _ = v180
                                                                                                                                                                    }
                                                                                                                                                                    var v181 uint16 = 0
                                                                                                                                                                    _ = v181
                                                                                                                                                                }
                                                                                                                                                                _ = len(v113)
                                                                                                                                                            case "\n\t":
                                                                                                                                                                var v182 float64 = float64(v4)
                                                                                                                                                                _ = v182
                                                                                                                                                            default:
                                                                                                                                                                var v183 int16 = (v95 - 32767)
                                                                                                                                                                _ = v183
                                                                                                                                                            }
                                                                                                                                                            v184 := GenS11{F0: v95}
                                                                                                                                                            switch fi26 {
                                                                                                                                                            case -2147483648:
                                                                                                                                                                var v185 string = (("\n\t" + ("a" + "a")) + v25)
                                                                                                                                                                _ = v185
                                                                                                                                                            case 0:
                                                                                                                                                                v186 := GenS1{F0: int(v95), F1: 994}
                                                                                                                                                                _ = v186
                                                                                                                                                            }
                                                                                                                                                            _ = v184
                                                                                                                                                        }
                                                                                                                                                        _ = v175
                                                                                                                                                        _ = v176
                                                                                                                                                    }
                                                                                                                                                    v104 += v104
                                                                                                                                                    v187 := v5[v3]
                                                                                                                                                    _ = v187
                                                                                                                                                case "\n\t":
                                                                                                                                                    fmt.Println(fi26, v122)
                                                                                                                                                case "a":
                                                                                                                                                    switch 694 {
                                                                                                                                                    case -2147483648:
                                                                                                                                                        if v170 {
                                                                                                                                                            if v170 {
                                                                                                                                                                var v188 uint32 = v123
                                                                                                                                                                if (!v170) {
                                                                                                                                                                    var v189 float32 = ((0.0 / float32(v44)) + 3.14159)
                                                                                                                                                                    switch v25 {
                                                                                                                                                                    case "a":
                                                                                                                                                                        v122 += 47
                                                                                                                                                                    case "":
                                                                                                                                                                        if (65535 == 32767) {
                                                                                                                                                                            v2[v3] = int(v95)
                                                                                                                                                                        } else {
                                                                                                                                                                            v6 -= (v6 % ((uint8(v3)) | 1))
                                                                                                                                                                            v190 := GenS1{F0: (-(fi43)), F1: 1}
                                                                                                                                                                            _ = v190
                                                                                                                                                                        }
                                                                                                                                                                    default:
                                                                                                                                                                        v149 = v149
                                                                                                                                                                    }
                                                                                                                                                                    _ = v189
                                                                                                                                                                } else {
                                                                                                                                                                    v1 += (v1 ^ int8(v95))
                                                                                                                                                                }
                                                                                                                                                                _ = v188
                                                                                                                                                            } else {
                                                                                                                                                                switch v25 {
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    if (255 < 255) {
                                                                                                                                                                        var v191 uint = 255
                                                                                                                                                                        v113 = append(v113, (-(fi43)))
                                                                                                                                                                        if v170 {
                                                                                                                                                                            v192 := GenS9{F0: -862, F1: v191}
                                                                                                                                                                            var v193 uint8 = uint8(v94)
                                                                                                                                                                            _ = v192
                                                                                                                                                                            _ = v193
                                                                                                                                                                        }
                                                                                                                                                                        _ = v191
                                                                                                                                                                    }
                                                                                                                                                                case "hello":
                                                                                                                                                                    if (!false) {
                                                                                                                                                                        switch "" {
                                                                                                                                                                        case "hello":
                                                                                                                                                                            v194 := GenS10{F0: -501, F1: v25, F2: 1}
                                                                                                                                                                            _ = v194
                                                                                                                                                                        case "a":
                                                                                                                                                                            fmt.Println(v25, v123, v119)
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            var v195 int32 = 42
                                                                                                                                                                            _ = v195
                                                                                                                                                                        default:
                                                                                                                                                                            var v196 uint8 = (v24 & 127)
                                                                                                                                                                            _ = v196
                                                                                                                                                                        }
                                                                                                                                                                        v197 := GenS23{F0: "\n\t", F1: int8(v104)}
                                                                                                                                                                        v119 -= 0.1
                                                                                                                                                                        _ = v197
                                                                                                                                                                    } else {
                                                                                                                                                                        v5[0] = v3
                                                                                                                                                                    }
                                                                                                                                                                default:
                                                                                                                                                                    if ((^(fi43)) <= (fi43 % ((fi26) | 1))) {
                                                                                                                                                                        v2[(^(v3))] = (fi43 % ((32768) | 1))
                                                                                                                                                                    } else {
                                                                                                                                                                        v198 := GenS10{F0: 0, F1: v25, F2: 0}
                                                                                                                                                                        v199 := GenS18{F0: 0}
                                                                                                                                                                        if v170 {
                                                                                                                                                                            var v200 uint8 = (^(v6))
                                                                                                                                                                            v201 := GenS27{F0: fi43}
                                                                                                                                                                            _ = v200
                                                                                                                                                                            _ = v201
                                                                                                                                                                        }
                                                                                                                                                                        _ = v198
                                                                                                                                                                        _ = v199
                                                                                                                                                                    }
                                                                                                                                                                    switch v25 {
                                                                                                                                                                    case "\n\t":
                                                                                                                                                                        var v202 bool = (!(v104 > (v104 - (v104 ^ v104))))
                                                                                                                                                                        _ = v202
                                                                                                                                                                    case "hello":
                                                                                                                                                                        var v203 uint32 = (v3 & (v3 % (((v123 >> ((uint(v24)) & 63))) | 1)))
                                                                                                                                                                        _ = v203
                                                                                                                                                                    }
                                                                                                                                                                    switch "foo bar" {
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        v95 += v95
                                                                                                                                                                    case "hello":
                                                                                                                                                                        var v204 string = ("foo bar" + (v25 + ""))
                                                                                                                                                                        _ = v204
                                                                                                                                                                    case "":
                                                                                                                                                                        if (((v95 ^ v95) != v95) || true) {
                                                                                                                                                                            switch -884 {
                                                                                                                                                                            case -2147483648:
                                                                                                                                                                                var v205 uint = 2
                                                                                                                                                                                _ = v205
                                                                                                                                                                            case 188:
                                                                                                                                                                                switch "a" {
                                                                                                                                                                                case "foo bar":
                                                                                                                                                                                    v5[v123] = v3
                                                                                                                                                                                case "a":
                                                                                                                                                                                    fmt.Println(v118, v95)
                                                                                                                                                                                }
                                                                                                                                                                            case 963:
                                                                                                                                                                                v122 = v6
                                                                                                                                                                                fmt.Println(v118)
                                                                                                                                                                            }
                                                                                                                                                                            delete(v5, uint32(v94))
                                                                                                                                                                        } else {
                                                                                                                                                                            v95 = 32767
                                                                                                                                                                            if v170 {
                                                                                                                                                                                switch ("hello" + "\n\t") {
                                                                                                                                                                                case "\n\t":
                                                                                                                                                                                    if (v25 >= v25) {
                                                                                                                                                                                        fmt.Println(fi26)
                                                                                                                                                                                        fmt.Println(v94)
                                                                                                                                                                                        v206 := GenS14{F0: (v1 % (((^(v118))) | 1)), F1: ("a" + v25)}
                                                                                                                                                                                        _ = v206
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v207 := GenS28{F0: -46, F1: v104}
                                                                                                                                                                                        var v208 uint64 = 1
                                                                                                                                                                                        _ = v207
                                                                                                                                                                                        _ = v208
                                                                                                                                                                                    }
                                                                                                                                                                                case "hello":
                                                                                                                                                                                    v113 = append(v113, 384)
                                                                                                                                                                                case "a":
                                                                                                                                                                                    var v209 uint = 18446744073709551615
                                                                                                                                                                                    _ = v209
                                                                                                                                                                                default:
                                                                                                                                                                                    v113 = append(v113, fi43)
                                                                                                                                                                                }
                                                                                                                                                                                var v210 string = v25
                                                                                                                                                                                v211 := GenS29{F0: v118}
                                                                                                                                                                                _ = v210
                                                                                                                                                                                _ = v211
                                                                                                                                                                            } else {
                                                                                                                                                                                switch "foo bar" {
                                                                                                                                                                                case "hello":
                                                                                                                                                                                    switch v25 {
                                                                                                                                                                                    case "":
                                                                                                                                                                                        v95 = v95
                                                                                                                                                                                    case "hello":
                                                                                                                                                                                        v95 -= v95
                                                                                                                                                                                    case "\n\t":
                                                                                                                                                                                        if (v24 == (v6 - (^(v122)))) {
                                                                                                                                                                                            if (fi43 < 319) {
                                                                                                                                                                                                switch fi43 {
                                                                                                                                                                                                case 924:
                                                                                                                                                                                                    delete(v5, (^(v3)))
                                                                                                                                                                                                default:
                                                                                                                                                                                                    var v212 uint16 = 442
                                                                                                                                                                                                    _ = v212
                                                                                                                                                                                                }
                                                                                                                                                                                            } else {
                                                                                                                                                                                                if (!true) {
                                                                                                                                                                                                    var v213 float32 = 0.1
                                                                                                                                                                                                    _ = v213
                                                                                                                                                                                                }
                                                                                                                                                                                            }
                                                                                                                                                                                            if false {
                                                                                                                                                                                                var v214 int = -294
                                                                                                                                                                                                _ = v214
                                                                                                                                                                                            }
                                                                                                                                                                                            var v215 int8 = int8(fi43)
                                                                                                                                                                                            _ = v215
                                                                                                                                                                                        } else {
                                                                                                                                                                                            var v216 float32 = float32((v44 + 3.14159))
                                                                                                                                                                                            v217 := GenS30{F0: v104, F1: -94, F2: 42}
                                                                                                                                                                                            _ = v216
                                                                                                                                                                                            _ = v217
                                                                                                                                                                                        }
                                                                                                                                                                                    }
                                                                                                                                                                                default:
                                                                                                                                                                                    var v218 bool = (true && v170)
                                                                                                                                                                                    _ = v218
                                                                                                                                                                                }
                                                                                                                                                                                v219 := v5[179]
                                                                                                                                                                                _ = v219
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                v3 -= uint32(v6)
                                                                                                                                                            }
                                                                                                                                                            v113 = append(v113, 32767)
                                                                                                                                                            _ = len(v113)
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                if v170 {
                                                                                                                                                    fmt.Println(v95, v170, v95)
                                                                                                                                                } else {
                                                                                                                                                    var v220 float32 = v4
                                                                                                                                                    var v221 int16 = (^(v95))
                                                                                                                                                    if ((fi26 % ((32768) | 1)) <= fi43) {
                                                                                                                                                        var v222 bool = (!(!true))
                                                                                                                                                        var v223 int = (fi43 % ((4294967295) | 1))
                                                                                                                                                        v94 -= 42
                                                                                                                                                        v221 = v221
                                                                                                                                                        var v224 string = ("a" + (("\n\t" + v25) + ("hello" + "hello")))
                                                                                                                                                        _ = v222
                                                                                                                                                        _ = v223
                                                                                                                                                        _ = v224
                                                                                                                                                    } else {
                                                                                                                                                        if len(v113) >= 2 {
                                                                                                                                                            v113 = v113[1:len(v113)]
                                                                                                                                                        }
                                                                                                                                                        var v225 int = fi26
                                                                                                                                                        _ = v225
                                                                                                                                                    }
                                                                                                                                                    _ = v220
                                                                                                                                                    _ = v221
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        _ = v170
                                                                                                                                    }
                                                                                                                                    v226 := GenS3{F0: (v95 & 0)}
                                                                                                                                    v3 = uint32(v6)
                                                                                                                                    _ = v226
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v149
                                                                                                                            _ = v150
                                                                                                                        }
                                                                                                                        _ = v122
                                                                                                                        _ = v123
                                                                                                                    }
                                                                                                                default:
                                                                                                                    _ = len(v113)
                                                                                                                }
                                                                                                                var v227 float32 = v4
                                                                                                                _ = v119
                                                                                                                _ = v227
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v118
                                                                                                }
                                                                                                v3 = v3
                                                                                                var v228 bool = (fi43 > (fi26 << ((236) & 63)))
                                                                                                _ = v228
                                                                                            }
                                                                                            _ = v104
                                                                                            _ = v105
                                                                                            _ = v113
                                                                                        }
                                                                                        var v229 uint64 = v94
                                                                                        v95 += -994
                                                                                        _ = v229
                                                                                    case -9223372036854775808:
                                                                                        var v230 int64 = int64(v3)
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v44, fi43)
                                                                                        case "":
                                                                                            v231 := GenS14{F0: v1, F1: (("hello" + (v25 + v25)) + "\n\t")}
                                                                                            v25 = v25
                                                                                            v44 += v44
                                                                                            _ = v231
                                                                                        case "foo bar":
                                                                                            v232 := []int{}
                                                                                            _ = v232
                                                                                        }
                                                                                        v233 := []int{int(v6)}
                                                                                        _ = v230
                                                                                        _ = v233
                                                                                    }
                                                                                    v6 += (v6 & (v6 - 2))
                                                                                    _ = v101
                                                                                }
                                                                                if true {
                                                                                    v2[77] = (fi43 % ((0) | 1))
                                                                                    v234 := v2[v3]
                                                                                    v235 := GenS28{F0: 0, F1: -974}
                                                                                    _ = v234
                                                                                    _ = v235
                                                                                } else {
                                                                                    v4 -= float32(v44)
                                                                                }
                                                                            case 2147483647:
                                                                                v6 = 128
                                                                            default:
                                                                                v44 = 0.0
                                                                            }
                                                                            switch v25 {
                                                                            case "":
                                                                                v5[(v3 << ((65535) & 63))] = 2147483648
                                                                            case "a":
                                                                                v236 := GenS10{F0: -2147483648, F1: "hello", F2: -432}
                                                                                _ = v236
                                                                            case "\n\t":
                                                                                v2[v3] = (fi26 ^ fi26)
                                                                            default:
                                                                                if true {
                                                                                    var v237 int = -232
                                                                                    _ = v237
                                                                                } else {
                                                                                    var v238 int32 = 0
                                                                                    _ = v238
                                                                                }
                                                                            }
                                                                            _ = len(v5)
                                                                        } else {
                                                                            var v239 bool = (true || ((1 >= uint16(v24)) && false))
                                                                            var v240 string = v25
                                                                            _ = v239
                                                                            _ = v240
                                                                        }
                                                                        v44 = ((((v44 + 1.5) / ((v44 - 0.0) - float64(float32(v44)))) + 1.5) / 0.1)
                                                                        v241 := GenS31{F0: (^(v3)), F1: (v3 + 852), F2: 255}
                                                                        _ = v241
                                                                    }
                                                                case "hello":
                                                                    v242 := GenS32{F0: 2}
                                                                    _ = v242
                                                                case "a":
                                                                    v243 := GenS31{F0: v3, F1: 1, F2: 1}
                                                                    _ = v243
                                                                }
                                                            case 32768:
                                                                delete(v2, v3)
                                                            default:
                                                                v24 -= (v24 & v24)
                                                            }
                                                            _ = v95
                                                        }
                                                        _ = v94
                                                    } else {
                                                        switch (fi26 ^ fi26) {
                                                        case -1:
                                                            v244 := GenS33{F0: 94}
                                                            _ = v244
                                                        case 32767:
                                                            var v245 int = fi43
                                                            _ = v245
                                                        default:
                                                            var v246 float32 = 1.5
                                                            _ = v246
                                                        }
                                                    }
                                                    var v247 int = fi26
                                                    var v248 uint8 = genFn0(0, (-(v1)))
                                                    _ = v247
                                                    _ = v248
                                                    _ = v93
                                                }
                                            default:
                                                switch fi43 {
                                                case 4294967295:
                                                    switch ("\n\t" + (("\n\t" + "a") + "")) {
                                                    case "\n\t":
                                                        var v249 uint = 65535
                                                        _ = v249
                                                    default:
                                                        fmt.Println(fi43, fi26, v4)
                                                        if true {
                                                            var v250 int64 = 0
                                                            var v251 int16 = 221
                                                            _ = v250
                                                            _ = v251
                                                        } else {
                                                            v252 := []int{-9223372036854775808, (fi26 & (fi43 | fi43)), (fi26 - fi43), fi26}
                                                            var v253 int16 = -1
                                                            v254 := GenS24{F0: -1, F1: 1, F2: (fi26 + (fi43 >> ((18446744073709551615) & 63)))}
                                                            _ = v252
                                                            _ = v253
                                                            _ = v254
                                                        }
                                                        switch 32767 {
                                                        case 0:
                                                            if true {
                                                                if (-404 <= 1) {
                                                                    var v255 float64 = 0.1
                                                                    _ = v255
                                                                }
                                                                var v256 int = fi43
                                                                switch v25 {
                                                                case "hello":
                                                                    if false {
                                                                        if ((^(v24)) > v6) {
                                                                            var v257 bool = (true && ((-2147483648 != -1) && true))
                                                                            var v258 int32 = 42
                                                                            v1 = v1
                                                                            _ = v257
                                                                            _ = v258
                                                                        } else {
                                                                            v259 := GenS34{F0: v1, F1: 42, F2: v1}
                                                                            v260 := GenS23{F0: (v25 + (v25 + ("foo bar" + "foo bar"))), F1: v1}
                                                                            var v261 int16 = int16(v3)
                                                                            _ = v259
                                                                            _ = v260
                                                                            _ = v261
                                                                        }
                                                                        if true {
                                                                            v262 := GenS35{F0: 1}
                                                                            var v263 float32 = v4
                                                                            _ = v262
                                                                            _ = v263
                                                                        } else {
                                                                            fmt.Println(v25)
                                                                            var v264 string = ((v25 + "a") + ((((v25 + "") + "foo bar") + v25) + "foo bar"))
                                                                            var v265 int8 = v1
                                                                            _ = v264
                                                                            _ = v265
                                                                        }
                                                                        v24 -= v6
                                                                    } else {
                                                                        if (v3 <= (v3 ^ (v3 - 4294967295))) {
                                                                            var v266 uint32 = v3
                                                                            _ = v266
                                                                        }
                                                                        v44 += ((v44 / v44) / float64(float32(-2.25)))
                                                                        fmt.Println(fi43)
                                                                    }
                                                                default:
                                                                    var v267 uint64 = 18446744073709551615
                                                                    _ = v267
                                                                }
                                                                _ = v256
                                                            } else {
                                                                var v268 uint8 = genFn0(1, 100)
                                                                _ = v268
                                                            }
                                                        }
                                                    }
                                                case -169:
                                                    v269 := v2[v3]
                                                    _ = v269
                                                default:
                                                    var v270 int64 = -1
                                                    _ = v270
                                                }
                                            }
                                        }
                                    }
                                default:
                                    fmt.Println(v3, v25, v6)
                                }
                                v271 := GenS4{F0: 914, F1: (fi43 * fi26)}
                                v6 = (^(v24))
                                v272 := GenS36{F0: (v3 ^ (v3 | (v3 / ((v3) | 1)))), F1: (!(v1 < v1)), F2: (fi26 & fi26)}
                                _ = v271
                                _ = v272
                            }
                            if (!(0 > 169)) {
                                v273 := GenS29{F0: (v1 / ((v1) | 1))}
                                switch 2147483647 {
                                case 128:
                                    switch "foo bar" {
                                    case "\n\t":
                                        v274 := []int{(fi26 >> ((4294967295) & 63)), 127}
                                        _ = v274
                                    case "foo bar":
                                        if (false || true) {
                                            if true {
                                                switch ((v25 + "hello") + v25) {
                                                case "":
                                                    fmt.Println(v1, v3)
                                                case "hello":
                                                    var v275 int16 = 510
                                                    _ = v275
                                                case "\n\t":
                                                    v44 -= float64(v4)
                                                default:
                                                    var v276 uint = 382
                                                    _ = v276
                                                }
                                            } else {
                                                v277 := GenS11{F0: -176}
                                                v25 = "\n\t"
                                                if true {
                                                    switch -800 {
                                                    case 127:
                                                        var v278 float64 = (1.5 * (v44 / v44))
                                                        _ = v278
                                                    case 1:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if true {
                                                                switch v25 {
                                                                case "hello":
                                                                    v279 := GenS12{F0: genFn0(-1, 0), F1: 1, F2: v1}
                                                                    _ = v279
                                                                case "\n\t":
                                                                    if (!(false || false)) {
                                                                        switch (fi43 & (-(fi26))) {
                                                                        case 128:
                                                                            v280 := GenS10{F0: 611, F1: (v25 + v25), F2: 680}
                                                                            _ = v280
                                                                        default:
                                                                            fmt.Println(v44, fi26, v44)
                                                                        }
                                                                        fmt.Println(fi43, v25, v6)
                                                                        v24 = 1
                                                                        v1 += (v1 >> ((2) & 63))
                                                                        v6 -= 127
                                                                    } else {
                                                                        var v281 uint = 0
                                                                        _ = v281
                                                                    }
                                                                case "":
                                                                    v282 := []int{(fi26 % ((128) | 1)), -209, int(v24), fi26}
                                                                    _ = v282
                                                                default:
                                                                    var v283 int64 = -1
                                                                    _ = v283
                                                                }
                                                            } else {
                                                                v4 = (float32(-2.25) * v4)
                                                                switch ((v25 + (("a" + "foo bar") + v25)) + "hello") {
                                                                case "a":
                                                                    var v284 uint16 = 54
                                                                    _ = v284
                                                                case "\n\t":
                                                                    if true {
                                                                        v285 := []int{(fi43 - fi26), fi26, fi26}
                                                                        v286 := GenS33{F0: 18446744073709551615}
                                                                        var v287 uint8 = (v6 / ((genFn0(2147483647, v1)) | 1))
                                                                        v288 := v2[(v3 << ((4294967295) & 63))]
                                                                        if false {
                                                                            var v289 uint32 = v3
                                                                            v290 := GenS3{F0: -1}
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                if (v4 >= v4) {
                                                                                    var v291 string = ((("\n\t" + (v25 + v25)) + v25) + (v25 + "hello"))
                                                                                    _ = v291
                                                                                } else {
                                                                                    var v292 int8 = (v1 + v1)
                                                                                    var v293 uint = 1
                                                                                    fmt.Println(v3, v292)
                                                                                    _ = v292
                                                                                    _ = v293
                                                                                }
                                                                            case "a":
                                                                                switch (v288 << ((722) & 63)) {
                                                                                case 42:
                                                                                    if (true && (337 == -801)) {
                                                                                        v24 += (v24 * v24)
                                                                                        v294 := GenS19{F0: fi43, F1: 2147483647}
                                                                                        if ((1.5 * float32((v44 + v44))) != -2.25) {
                                                                                            var v295 int8 = (v1 | (v1 / (((^(v1))) | 1)))
                                                                                            v3 = v3
                                                                                            _ = v295
                                                                                        }
                                                                                        _ = v294
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    v287 += v287
                                                                                default:
                                                                                    var v296 int = (fi26 - (-(fi26)))
                                                                                    _ = v296
                                                                                }
                                                                            case "":
                                                                                switch ("hello" + "") {
                                                                                case "foo bar":
                                                                                    v44 = v44
                                                                                case "hello":
                                                                                    if true {
                                                                                        switch (fi43 << ((65535) & 63)) {
                                                                                        case 127:
                                                                                            delete(v2, (v3 ^ (v3 | uint32(v6))))
                                                                                        }
                                                                                    } else {
                                                                                        v287 = v287
                                                                                        var v297 uint16 = 147
                                                                                        if ((v24 - v6) != (v287 % (((v24 * genFn0(1, (v1 >> ((42) & 63))))) | 1))) {
                                                                                            if false {
                                                                                                v298 := GenS37{F0: -1}
                                                                                                _ = v298
                                                                                            } else {
                                                                                                v289 = (v3 % (((v289 & v289)) | 1))
                                                                                                var v299 int64 = 0
                                                                                                _ = v299
                                                                                            }
                                                                                            fmt.Println(v24, v287, fi43)
                                                                                            v289 += 746
                                                                                        } else {
                                                                                            if (true && (uint32(v297) >= v3)) {
                                                                                                if false {
                                                                                                    v4 += 0.1
                                                                                                    v288 += -1
                                                                                                    v6 -= uint8(v3)
                                                                                                } else {
                                                                                                    var v300 uint16 = (v297 * 1)
                                                                                                    fmt.Println(v25)
                                                                                                    _ = len(v5)
                                                                                                    _ = v300
                                                                                                }
                                                                                                var v301 uint = 600
                                                                                                v3 = (^(v3))
                                                                                                _ = v301
                                                                                            }
                                                                                            var v302 int = 32767
                                                                                            _ = v302
                                                                                        }
                                                                                        _ = v297
                                                                                    }
                                                                                default:
                                                                                    v303 := GenS38{F0: (v288 + -2147483648)}
                                                                                    _ = v303
                                                                                }
                                                                            default:
                                                                                v289 = v3
                                                                            }
                                                                            _ = v289
                                                                            _ = v290
                                                                        }
                                                                        _ = v285
                                                                        _ = v286
                                                                        _ = v287
                                                                        _ = v288
                                                                    }
                                                                }
                                                                var v304 uint = 4294967295
                                                                _ = v304
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v277
                                            }
                                            if true {
                                                v305 := GenS39{F0: 42, F1: (v6 / ((v6) | 1))}
                                                _ = v305
                                            } else {
                                                var v306 int8 = int8(fi43)
                                                if true {
                                                    var v307 float64 = v44
                                                    if false {
                                                        switch (fi43 << ((901) & 63)) {
                                                        case -9223372036854775808:
                                                            v308 := GenS37{F0: int8(fi26)}
                                                            _ = v308
                                                        }
                                                        if ((0 >= 2) || (255 != 0)) {
                                                            v306 -= v306
                                                        }
                                                        v309 := GenS3{F0: -320}
                                                        _ = v309
                                                    }
                                                    var v310 uint16 = 2
                                                    var v311 float32 = v4
                                                    _ = v307
                                                    _ = v310
                                                    _ = v311
                                                } else {
                                                    var v312 int8 = (v1 / (((v306 >> ((1) & 63))) | 1))
                                                    fmt.Println(v306, v25)
                                                    _ = v312
                                                }
                                                var v313 uint32 = (v3 % ((2147483647) | 1))
                                                _ = v306
                                                _ = v313
                                            }
                                        } else {
                                            v314 := []int{}
                                            var v315 int = (fi43 * (fi26 % ((fi43) | 1)))
                                            _ = v314
                                            _ = v315
                                        }
                                    }
                                    if (v3 >= (v3 % ((v3) | 1))) {
                                        if true {
                                            if true {
                                                v316 := []int{fi43, fi26}
                                                v3 += v3
                                                _ = v316
                                            }
                                            switch ("\n\t" + "a") {
                                            case "\n\t":
                                                v4 += (1e10 / v4)
                                            default:
                                                var v317 float64 = v44
                                                _ = v317
                                            }
                                            var v318 float64 = v44
                                            v319 := []int{(fi26 / ((fi26) | 1)), (fi26 + fi43)}
                                            v320 := GenS40{F0: (fi26 >> ((255) & 63)), F1: int64(v24), F2: 0}
                                            v3 -= 2147483647
                                            var v321 int16 = -32
                                            _ = v318
                                            _ = v319
                                            _ = v320
                                            _ = v321
                                        } else {
                                            fmt.Println(v24, v25, v3)
                                            var v322 float32 = 3.14159
                                            _ = v322
                                        }
                                    } else {
                                        switch fi43 {
                                        case 811:
                                            v6 += genFn0(-2147483648, (v1 ^ (v1 & v1)))
                                        case 640:
                                            v323 := GenS41{F0: true}
                                            _ = v323
                                        }
                                        if false {
                                            if ((false || ((false && (!(0 >= 153))) && ((fi43 + (-(fi26))) >= fi26))) || false) {
                                                var v324 uint16 = 31
                                                v24 += v24
                                                _ = v324
                                            } else {
                                                var v325 uint64 = 0
                                                v24 += (^(v24))
                                                _ = v325
                                            }
                                            if true {
                                                if true {
                                                    if ((v1 >> ((665) & 63)) < 63) {
                                                        if (!false) {
                                                            v326 := GenS13{F0: 9223372036854775807, F1: v6, F2: 4294967295}
                                                            switch (v25 + (v25 + v25)) {
                                                            case "hello":
                                                                v327 := GenS18{F0: 244}
                                                                _ = v327
                                                            case "":
                                                                if false {
                                                                    v6 = v6
                                                                    if ((v25 > "a") && (-221 < 255)) {
                                                                        if false {
                                                                            _ = len(v2)
                                                                            v328 := GenS42{F0: 0, F1: -78, F2: 104}
                                                                            v4 += (v4 * float32((v44 * float64(0.1))))
                                                                            _ = v328
                                                                        } else {
                                                                            fmt.Println(fi43, v25, fi26)
                                                                            if true {
                                                                                if true {
                                                                                    v329 := GenS11{F0: -1}
                                                                                    var v330 int8 = (v1 << ((65535) & 63))
                                                                                    _ = v329
                                                                                    _ = v330
                                                                                } else {
                                                                                    v331 := GenS43{F0: (v3 % ((1) | 1)), F1: v24, F2: 106}
                                                                                    v44 -= (v44 * 0.1)
                                                                                    _ = v331
                                                                                }
                                                                                v24 = genFn0(-2147483648, (v1 * v1))
                                                                                switch v25 {
                                                                                case "\n\t":
                                                                                    v25 = (v25 + v25)
                                                                                case "a":
                                                                                    var v332 float64 = 3.14159
                                                                                    _ = v332
                                                                                case "foo bar":
                                                                                    var v333 int = (fi43 | int(v3))
                                                                                    _ = v333
                                                                                default:
                                                                                    v334 := GenS44{F0: 0, F1: uint(v3), F2: 1}
                                                                                    _ = v334
                                                                                }
                                                                                if (genFn0(-566, v1) <= (v24 - (v6 << ((1) & 63)))) {
                                                                                    v335 := GenS30{F0: -9223372036854775808, F1: 2147483647, F2: v24}
                                                                                    _ = v335
                                                                                }
                                                                            } else {
                                                                                v4 -= ((((v4 - v4) + ((0.0 - 0.1) - 1.5)) + 0.1) * 0.1)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v1)
                                                            }
                                                            switch "" {
                                                            case "":
                                                                v336 := []int{fi43, -742, 32767}
                                                                _ = v336
                                                            case "a":
                                                                var v337 bool = false
                                                                _ = v337
                                                            case "foo bar":
                                                                var v338 int64 = 78
                                                                _ = v338
                                                            default:
                                                                if false {
                                                                    v24 = (v6 * (^(v6)))
                                                                    fmt.Println(v1, v24, v6)
                                                                }
                                                            }
                                                            v25 = v25
                                                            v339 := GenS45{F0: v25, F1: v25, F2: 32767}
                                                            _ = v326
                                                            _ = v339
                                                        }
                                                        fmt.Println(v6, v25, v3)
                                                    }
                                                    v1 += 1
                                                    v24 -= 128
                                                    var v340 uint32 = 2147483648
                                                    switch fi26 {
                                                    case 695:
                                                        if (v4 != (v4 - 0.1)) {
                                                            delete(v2, (v3 & v340))
                                                            switch (fi43 % (((fi26 ^ fi43)) | 1)) {
                                                            case 32768:
                                                                v25 = v25
                                                                v6 -= (v6 % (((v24 ^ v6)) | 1))
                                                            case 4294967295:
                                                                v341 := GenS34{F0: 100, F1: 846, F2: v1}
                                                                _ = v341
                                                            case 1:
                                                                var v342 string = v25
                                                                _ = v342
                                                            default:
                                                                var v343 int8 = 92
                                                                _ = v343
                                                            }
                                                        }
                                                    case 201:
                                                        fmt.Println(fi43)
                                                    case -9223372036854775808:
                                                        v344 := []int{fi43, (fi43 * (fi26 % ((42) | 1)))}
                                                        _ = v344
                                                    }
                                                    if (int8(v24) < v1) {
                                                        var v345 bool = (619 != -891)
                                                        v346 := []int{fi26, 404}
                                                        fmt.Println(v1)
                                                        _ = v345
                                                        _ = v346
                                                    } else {
                                                        v25 = (v25 + v25)
                                                        v347 := GenS46{F0: 0}
                                                        v24 += 1
                                                        v348 := []int{9223372036854775807, fi26}
                                                        if len(v348) > 0 {
                                                            v348[0] = (fi26 + (-(fi43)))
                                                        }
                                                        v349 := GenS47{F0: 558}
                                                        _ = v347
                                                        _ = v348
                                                        _ = v349
                                                    }
                                                    _ = v340
                                                } else {
                                                    v350 := GenS6{F0: v1, F1: -32768}
                                                    v351 := GenS48{F0: v25, F1: (fi26 != (fi43 % (((fi43 + (fi43 >> ((23) & 63)))) | 1)))}
                                                    if (2147483647 == 2147483647) {
                                                        if (1 >= 42) {
                                                            if false {
                                                                v352 := GenS30{F0: 514, F1: 1, F2: (v24 + 105)}
                                                                _ = v352
                                                            } else {
                                                                fmt.Println(v3, v24, v4)
                                                            }
                                                            if (v4 <= v4) {
                                                                delete(v5, v3)
                                                                if (!(genFn0(42, -1) > (v24 >> ((2) & 63)))) {
                                                                    var v353 uint = uint(v24)
                                                                    _ = v353
                                                                }
                                                            }
                                                        }
                                                        _ = len(v5)
                                                        v354 := GenS22{F0: -293, F1: 42, F2: genFn0(0, v1)}
                                                        _ = v354
                                                    }
                                                    _ = v350
                                                    _ = v351
                                                }
                                                _ = len(v5)
                                                fmt.Println(v3, fi43)
                                            } else {
                                                var v355 int = (fi43 / (((fi26 + fi43)) | 1))
                                                _ = v355
                                            }
                                            v356 := []int{(-(fi43)), 2147483647, fi43}
                                            v357 := v5[244]
                                            if false {
                                                v358 := GenS12{F0: v24, F1: 1, F2: (-(v1))}
                                                _ = v358
                                            } else {
                                                switch v25 {
                                                case "foo bar":
                                                    switch (fi26 % ((fi26) | 1)) {
                                                    case 4294967295:
                                                        v44 = ((v44 * 1.5) + v44)
                                                    case -1:
                                                        if len(v356) >= 2 {
                                                            v356 = v356[1:len(v356)]
                                                        }
                                                    case 1:
                                                        v25 = "\n\t"
                                                    }
                                                default:
                                                    if ((v25 + ((("hello" + v25) + v25) + (((((((v25 + v25) + v25) + v25) + v25) + (((v25 + v25) + v25) + v25)) + v25) + v25))) > (v25 + ("foo bar" + ""))) {
                                                        v359 := GenS49{F0: 9223372036854775807, F1: 1000, F2: 9223372036854775807}
                                                        _ = v359
                                                    }
                                                }
                                            }
                                            v360 := GenS8{F0: 42, F1: 0}
                                            v361 := GenS50{F0: 439, F1: -9223372036854775808}
                                            _ = v356
                                            _ = v357
                                            _ = v360
                                            _ = v361
                                        }
                                    }
                                    var v362 uint16 = 2
                                    _ = v362
                                }
                                v1 -= (v1 + v1)
                                var v363 uint32 = (v3 & (^(v3)))
                                _ = v273
                                _ = v363
                            } else {
                                if (!(65535 < uint(v3))) {
                                    v6 = (^(v24))
                                    switch ((("a" + v25) + (v25 + v25)) + (v25 + v25)) {
                                    case "":
                                        var v364 uint64 = 995
                                        _ = v364
                                    case "hello":
                                        v6 += v6
                                    case "\n\t":
                                        var v365 bool = true
                                        _ = v365
                                    default:
                                        if (false && (!((65535 != 753) && true))) {
                                            var v366 uint16 = 896
                                            v44 = 1.5
                                            v4 = v4
                                            _ = v366
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            v367 := GenS19{F0: fi43, F1: (v3 + v3)}
                                            _ = v367
                                        default:
                                            if true {
                                                var v368 bool = (!true)
                                                fmt.Println(v1, v1)
                                                v369 := GenS37{F0: v1}
                                                _ = v368
                                                _ = v369
                                            }
                                        }
                                        var v370 uint32 = (v3 | v3)
                                        _ = v370
                                    }
                                } else {
                                    v371 := []int{-904, (fi26 / ((fi43) | 1))}
                                    var v372 uint8 = (v24 / ((127) | 1))
                                    _ = v371
                                    _ = v372
                                }
                            }
                        } else {
                            var v373 int64 = 1
                            switch (v25 + (("foo bar" + v25) + v25)) {
                            case "\n\t":
                                var v374 int8 = (v1 * (v1 + (v1 - -1)))
                                _ = v374
                            case "":
                                v1 = v1
                            case "hello":
                                if false {
                                    switch 2147483647 {
                                    case 127:
                                        fmt.Println(v6, v373)
                                    case 42:
                                        v375 := GenS51{F0: 18446744073709551615, F1: v1}
                                        v6 -= 2
                                        delete(v5, v3)
                                        if (false && true) {
                                            switch "hello" {
                                            case "a":
                                                if false {
                                                    var v376 int32 = int32(fi43)
                                                    _ = v376
                                                } else {
                                                    if (false && true) {
                                                        delete(v5, (v3 - (v3 << ((761) & 63))))
                                                    } else {
                                                        if true {
                                                            v24 = v6
                                                            v4 = ((v4 - (float32(v44) + v4)) - float32(((1e10 + ((0.0 - -2.25) + v44)) - (1e10 / float64(v4)))))
                                                            var v377 uint8 = 0
                                                            _ = v377
                                                        }
                                                        v373 = v373
                                                    }
                                                }
                                            case "\n\t":
                                                v1 = int8(fi43)
                                            default:
                                                v373 += (v373 & v373)
                                            }
                                            fmt.Println(v373)
                                            fmt.Println(v4, v25, v1)
                                        } else {
                                            v24 = v24
                                            var v378 float64 = 1.5
                                            _ = v378
                                        }
                                        _ = v375
                                    case 2147483647:
                                        var v379 int32 = -1
                                        _ = v379
                                    default:
                                        var v380 float64 = v44
                                        var v381 int32 = 192
                                        v382 := GenS52{F0: v1, F1: 2147483647}
                                        v373 += -9223372036854775808
                                        _ = v380
                                        _ = v381
                                        _ = v382
                                    }
                                    var v383 uint16 = 2
                                    v384 := v2[uint32(v383)]
                                    _ = v383
                                    _ = v384
                                } else {
                                    v44 = ((float64(v4) * float64(v4)) + v44)
                                    v385 := GenS43{F0: (v3 % ((v3) | 1)), F1: (^(v24)), F2: 255}
                                    fmt.Println(v3, v3)
                                    _ = v385
                                }
                                v3 -= (^(v3))
                                delete(v2, (v3 >> ((18446744073709551615) & 63)))
                            }
                            v386 := []int{(fi43 - (fi26 << ((996) & 63))), (fi26 & (fi43 % ((fi26) | 1))), -1}
                            _ = v373
                            _ = v386
                        }
                        v3 = v3
                    }
                    if (true && true) {
                        var v387 bool = false
                        _ = v387
                    }
                    v25 = (("\n\t" + v25) + v25)
                    v25 = v25
                    _ = fi43
                    _ = v44
                }
            case "hello":
                v1 = int8(v3)
            }
        }
        v388 := GenS52{F0: v1, F1: 2147483648}
        _ = fi26
        _ = v388
    }
    if (643 != 139) {
        var v389 int8 = v1
        v4 = 0.1
        switch -358 {
        case -2147483648:
            v390 := GenS38{F0: 1}
            _ = v390
        case 127:
            fmt.Println(v389, v25, v24)
        case 605:
            if false {
                v391 := v2[(v3 + 0)]
                switch v25 {
                case "a":
                    v1 = v389
                case "hello":
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v25)
                    }
                default:
                    if false {
                        v392 := v5[v3]
                        _ = v392
                    }
                }
                _ = v391
            } else {
                for fi393 := 0; fi393 < 12; fi393++ {
                    v3 = (^(v3))
                    var v394 uint8 = (v24 | v24)
                    v3 = v3
                    _ = fi393
                    _ = v394
                }
                if (true || (42 < 1)) {
                    v395 := GenS53{F0: 1, F1: 9223372036854775807, F2: 156}
                    for fi396 := 0; fi396 < 15; fi396++ {
                        var v397 int16 = -32768
                        var v398 uint8 = v24
                        if false {
                            var v399 uint32 = uint32(v24)
                            _ = v399
                        } else {
                            for fi400 := 0; fi400 < 7; fi400++ {
                                v401 := v5[v3]
                                fmt.Println(v401)
                                var v402 string = "\n\t"
                                _ = fi400
                                _ = v401
                                _ = v402
                            }
                            fmt.Println(v389)
                            v397 = -261
                        }
                        _ = fi396
                        _ = v397
                        _ = v398
                    }
                    v24 -= (^(v24))
                    _ = v395
                } else {
                    var v403 uint8 = 61
                    _ = v403
                }
                v6 -= 51
            }
        }
        _ = v389
    }
    if (true && ((true || true) || (18446744073709551615 != uint64(v3)))) {
        v24 -= (^(v24))
        switch "foo bar" {
        case "a":
            v404 := GenS46{F0: (^(v1))}
            _ = v404
        default:
            v405 := GenS42{F0: -2147483648, F1: v1, F2: 32767}
            _ = v405
        }
        switch v25 {
        case "":
            fmt.Println(v24, v6)
        case "a":
            v1 = (v1 / (((v1 - v1)) | 1))
        default:
            v406 := v5[4294967295]
            _ = v406
        }
    }
    var v407 int16 = 32767
    if ((!true) && (false || (1 > -9223372036854775808))) {
        for fi408 := 0; fi408 < 8; fi408++ {
            v407 = v407
            if true {
                var v409 uint64 = 1
                v5[4294967295] = 210
                _ = v409
            } else {
                var v410 uint = 42
                var v411 float64 = 3.14159
                v412 := GenS54{F0: 255, F1: (v6 / ((v6) | 1))}
                _ = v410
                _ = v411
                _ = v412
            }
            var v413 float64 = 1e10
            _ = fi408
            _ = v413
        }
        var v414 uint = 40
        if (-2147483648 >= -1) {
            if ((v24 << ((v414) & 63)) != (v24 >> (((v414 | (^(v414)))) & 63))) {
                switch ("hello" + v25) {
                case "\n\t":
                    v415 := GenS29{F0: v1}
                    _ = v415
                case "a":
                    v6 = (^(v24))
                case "foo bar":
                    var v416 uint16 = uint16(v6)
                    _ = v416
                }
            } else {
                fmt.Println(v6, v407)
                if true {
                    fmt.Println(v25, v25)
                    var v417 uint16 = 42
                    _ = v417
                } else {
                    for fi418 := 0; fi418 < 19; fi418++ {
                        v419 := GenS11{F0: v407}
                        v3 -= v3
                        _ = fi418
                        _ = v419
                    }
                    fmt.Println(v3, v3, v24)
                    var v420 bool = true
                    _ = v420
                }
                if (v1 <= (v1 + (v1 ^ v1))) {
                    if true {
                        delete(v5, v3)
                    }
                    v25 = ("" + v25)
                } else {
                    v421 := []int{}
                    _ = v421
                }
            }
        } else {
            v422 := GenS55{F0: uint64(v414)}
            if false {
                switch v25 {
                case "\n\t":
                    v3 = (v3 % ((v3) | 1))
                case "":
                    var v423 int8 = (v1 >> (((v414 >> ((v414) & 63))) & 63))
                    _ = v423
                default:
                    v424 := GenS56{F0: (v414 | v414), F1: 0}
                    _ = v424
                }
            }
            var v425 uint16 = 1
            _ = v422
            _ = v425
        }
        _ = v414
    }
    fmt.Println(v407, v407)
    fmt.Println(v1, v4)
    v426 := GenS57{F0: ("a" + "a"), F1: (uint64(v6) != 14), F2: 177}
    v3 -= (v3 ^ (v3 ^ 529))
    var v427 int = 817
    v25 = v25
    if false {
        switch 42 {
        case 725:
            var v428 uint64 = 0
            _ = v428
        case 9223372036854775807:
            var v429 uint8 = genFn0(-210, v1)
            _ = v429
        default:
            fmt.Println(v24, v407)
        }
        fmt.Println(v25, v24, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v407:", v407)
    fmt.Println("v427:", v427)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v407
    _ = v426
    _ = v427
    _ = v5
    _ = v6
}
