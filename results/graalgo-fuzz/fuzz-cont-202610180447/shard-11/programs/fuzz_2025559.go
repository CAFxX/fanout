package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint32; F1 int16; F2 int64 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 string; F1 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint16; F1 int }
type GenS15 struct { F0 int64; F1 int32 }
type GenS16 struct { F0 int; F1 uint32; F2 bool }
func genFn0(p0 uint8, p1 uint16) int32 {
    p0 += (p0 + p0)
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 int32 = -2147483648
            p0 = (p0 * (p0 >> ((uint(p0)) & 63)))
            _ = fi5
            _ = v6
        }
        switch ("hello" + ((("a" + "foo bar") + "foo bar") + "foo bar")) {
        case "a":
            for fi7 := 0; fi7 < 15; fi7++ {
                p1 = (p1 >> ((483) & 63))
                if true {
                    var v8 int64 = 42
                    v9 := GenS1{F0: uint32(p0)}
                    var v10 uint8 = p0
                    _ = v10
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
        case "\n\t":
            var v11 uint64 = 704
            _ = v11
        case "hello":
            p1 = p1
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := []int{(-(fi12)), fi12, fi4}
            var v14 int = fi12
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi4
    }
    v15 := GenS1{F0: 0}
    var v16 bool = false
    if ((float64(1e10) / (float64((-2.25 * (0.1 * ((0.0 * float32(1.5)) + (((-2.25 * 0.1) - float32((3.14159 * 1.5))) * 3.14159))))) + (-2.25 / 1.0))) < 1.5) {
        v17 := []int{-770, 0, 1}
        if ((1.5 > (3.14159 - 1.5)) && true) {
            var v18 uint16 = 88
            _ = v18
        }
        _ = v17
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            if v16 {
                var v20 uint8 = p0
                switch (-(fi19)) {
                case 32767:
                    v16 = true
                case 979:
                    v21 := GenS1{F0: 436}
                    _ = v21
                case 9223372036854775807:
                    if (1.5 > (0.1 - (3.14159 * 3.14159))) {
                        v22 := GenS2{F0: 127, F1: (360 < 1), F2: -1}
                        v23 := []int{fi19}
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 int64 = -1
                            fmt.Println(fi19, p1, v20)
                            _ = fi24
                            _ = v25
                        }
                        if (0.0 != ((0.1 + float64(float32(1.5))) + -2.25)) {
                            for fi26 := 0; fi26 < 2; fi26++ {
                                if v16 {
                                    if v16 {
                                        var v27 float64 = 1.5
                                        fmt.Println(fi19, v27)
                                        _ = v27
                                    }
                                    v28 := []int{}
                                    _ = v28
                                }
                                _ = fi26
                            }
                        } else {
                            fmt.Println(v16, p1)
                            v29 := []int{int(v20)}
                            _ = len(v29)
                            _ = v29
                        }
                        p0 -= p0
                    }
                default:
                    var v30 int8 = -1
                    _ = v30
                }
                _ = v20
            }
            _ = fi19
        }
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 int = 128
    var v2 bool = false
    var v3 uint8 = 170
    v31 := []int{(v1 ^ 32767), (-(v1)), v1, v1}
    for fi32 := 0; fi32 < 12; fi32++ {
        if v2 {
            v3 += v3
        } else {
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
        }
        v33 := GenS1{F0: 2147483648}
        _ = fi32
        _ = v33
    }
    switch ("\n\t" + "hello") {
    case "a":
        v3 = v3
    case "\n\t":
        v34 := GenS2{F0: int8(v3), F1: true, F2: 1}
        _ = v34
    default:
        v1 = 923
    }
    v1 = (v1 % ((-158) | 1))
    var v35 uint32 = uint32(v3)
    fmt.Println(v35)
    if ((!v2) || false) {
        if ((255 > 255) || (v3 < v3)) {
            v35 += (^(v35))
            v36 := GenS3{F0: 1, F1: -9223372036854775808, F2: 1}
            fmt.Println(v35, v1, v1)
            _ = v36
        }
        if len(v31) > 0 {
            v31[0] = 32768
        }
        for fi37 := 0; fi37 < 5; fi37++ {
            v38 := GenS1{F0: v35}
            if v2 {
                v39 := GenS4{F0: (1 == 854)}
                v31 = append(v31, 32768)
                if len(v31) > 0 {
                    v31[0] = (fi37 & (v1 >> ((914) & 63)))
                }
                _ = v39
            }
            var v40 float32 = 0.1
            _ = fi37
            _ = v38
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 5; fi41++ {
            v2 = (1 <= 91)
            _ = fi41
        }
        for fi42 := 0; fi42 < 2; fi42++ {
            var v43 float32 = (0.0 - -2.25)
            var v44 uint16 = 1
            _ = fi42
            _ = v43
            _ = v44
        }
        v45 := GenS2{F0: 27, F1: v2, F2: 0}
        _ = v45
    }
    fmt.Println(v1)
    for fi46 := 0; fi46 < 10; fi46++ {
        switch "\n\t" {
        case "foo bar":
            var v47 int8 = -39
            _ = v47
        default:
            for fi48 := 0; fi48 < 9; fi48++ {
                fmt.Println(fi48)
                _ = fi48
            }
        }
        switch fi46 {
        case -9223372036854775808:
            v3 -= 127
        case 629:
            _ = len(v31)
        default:
            v49 := GenS1{F0: 1}
            _ = v49
        }
        for fi50 := 0; fi50 < 19; fi50++ {
            v3 = (v3 % (((v3 - (v3 ^ (v3 * v3)))) | 1))
            var v51 uint = 2
            var v52 string = "hello"
            if (v35 != v35) {
                v2 = ((("a" + v52) + (("foo bar" + v52) + "")) > v52)
            } else {
                var v53 float64 = 0.1
                v54 := GenS5{F0: (v51 % ((v51) | 1)), F1: fi46}
                if (v2 || false) {
                    if true {
                        var v55 bool = false
                        _ = v55
                    } else {
                        v53 = float64(0.1)
                    }
                    fmt.Println(v52)
                    _ = len(v31)
                } else {
                    if true {
                        if v2 {
                            var v56 float32 = 0.1
                            v52 = ""
                            v53 = v53
                            if v2 {
                                if len(v31) > 0 {
                                    v31[0] = (-(fi50))
                                }
                                if v2 {
                                    fmt.Println(fi50, v2)
                                    v1 -= -45
                                    v3 += v3
                                    var v57 float32 = (v56 / (-2.25 * 1.5))
                                    if ((!(!(1 == 203))) && (int32(v35) < 42)) {
                                        switch (v1 & (fi50 + (fi46 ^ (v1 << ((943) & 63))))) {
                                        case 128:
                                            fmt.Println(v51, fi46, v52)
                                            if ((!(!v2)) || ((18446744073709551615 >= 1) && (4294967295 != 1))) {
                                                v58 := GenS3{F0: 42, F1: 9223372036854775807, F2: 1}
                                                _ = v58
                                            }
                                        case 4294967295:
                                            var v59 int16 = 1000
                                            v2 = v2
                                            _ = v59
                                        case 32767:
                                            switch "" {
                                            case "a":
                                                if v2 {
                                                    v2 = true
                                                    if (v51 <= v51) {
                                                        v57 = 3.14159
                                                        if len(v31) >= 2 {
                                                            v31 = v31[1:len(v31)]
                                                        }
                                                    } else {
                                                        _ = len(v31)
                                                    }
                                                    fmt.Println(v57)
                                                } else {
                                                    fmt.Println(v52, v1, v2)
                                                }
                                            }
                                        }
                                    } else {
                                        v60 := GenS6{F0: 369, F1: 1000, F2: 168}
                                        var v61 bool = ((-(v1)) != fi50)
                                        _ = v60
                                        _ = v61
                                    }
                                    if false {
                                        switch ("" + v52) {
                                        case "":
                                            if len(v31) > 0 {
                                                v31[0] = (^(v1))
                                            }
                                        case "a":
                                            var v62 float32 = (-2.25 / v56)
                                            _ = v62
                                        case "hello":
                                            if len(v31) > 0 {
                                                v31[0] = (fi50 - -461)
                                            }
                                        default:
                                            if ((false && (0 >= 1)) || v2) {
                                                _ = len(v31)
                                            } else {
                                                v63 := GenS7{F0: -9223372036854775808, F1: 111}
                                                switch fi46 {
                                                case 127:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v64 := GenS6{F0: v35, F1: 1000, F2: -1}
                                                        _ = v64
                                                    default:
                                                        switch (fi46 << ((v51) & 63)) {
                                                        case -2147483648:
                                                            v65 := GenS8{F0: (v35 / (((v35 | v35)) | 1))}
                                                            _ = v65
                                                        case -1:
                                                            v56 -= v57
                                                        default:
                                                            v66 := GenS9{F0: int(v35)}
                                                            _ = v66
                                                        }
                                                    }
                                                default:
                                                    if len(v31) >= 2 {
                                                        v31 = v31[1:len(v31)]
                                                    }
                                                    switch fi50 {
                                                    case 42:
                                                        v31 = append(v31, 32768)
                                                    case 103:
                                                        v2 = false
                                                    case -1:
                                                        var v67 float64 = v53
                                                        _ = v67
                                                    }
                                                }
                                                v51 = (v51 | uint(v35))
                                                var v68 bool = true
                                                v69 := GenS7{F0: 1, F1: 769}
                                                _ = v63
                                                _ = v68
                                                _ = v69
                                            }
                                        }
                                        var v70 float32 = (v57 - (v56 - v57))
                                        _ = v70
                                    }
                                    var v71 uint = 2
                                    _ = v57
                                    _ = v71
                                } else {
                                    if true {
                                        v53 = 3.14159
                                        if len(v31) >= 2 {
                                            v31 = v31[1:len(v31)]
                                        }
                                        var v72 int32 = 515
                                        _ = v72
                                    }
                                }
                            }
                            _ = v56
                        }
                    } else {
                        switch fi46 {
                        case -2147483648:
                            if (42 > -1) {
                                v31 = append(v31, 32767)
                                v31 = append(v31, fi46)
                                if v2 {
                                    fmt.Println(v1, v3)
                                    var v73 float32 = 0.0
                                    if ((v53 / 1.0) > v53) {
                                        _ = len(v31)
                                        if (63 != -128) {
                                            fmt.Println(v1, v51)
                                            fmt.Println(v3, v35, v2)
                                        }
                                        v74 := GenS10{F0: 42, F1: v51}
                                        _ = v74
                                    }
                                    _ = v73
                                }
                                if v2 {
                                    var v75 int32 = 42
                                    switch v52 {
                                    case "\n\t":
                                        fmt.Println(fi50, v3, v3)
                                    case "a":
                                        v76 := GenS2{F0: 100, F1: v2, F2: 1}
                                        _ = v76
                                    }
                                    v31 = append(v31, (-(v1)))
                                    var v77 bool = (v2 || ((v52 + (v52 + ("foo bar" + "\n\t"))) >= (v52 + (v52 + "foo bar"))))
                                    switch fi46 {
                                    case 238:
                                        v35 = (v35 + 2)
                                    case 42:
                                        var v78 string = v52
                                        _ = v78
                                    default:
                                        if v77 {
                                            switch (fi46 << (((v51 - 42)) & 63)) {
                                            case -389:
                                                v79 := GenS11{F0: 79}
                                                _ = v79
                                            case -606:
                                                if ((true && v2) || (128 >= 2)) {
                                                    var v80 uint64 = 2
                                                    _ = v80
                                                }
                                            }
                                        } else {
                                            v1 = 1
                                        }
                                    }
                                    if len(v31) > 0 {
                                        v31[0] = v1
                                    }
                                    _ = v75
                                    _ = v77
                                } else {
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    v81 := GenS12{F0: (v52 + "hello"), F1: (fi50 & v1)}
                                    v82 := GenS13{F0: -33}
                                    var v83 bool = (1e10 <= (float32(3.14159) - ((0.0 * float32((v53 + v53))) / -2.25)))
                                    _ = v81
                                    _ = v82
                                    _ = v83
                                }
                                var v84 int32 = 272
                                _ = v84
                            } else {
                                var v85 bool = v2
                                switch v1 {
                                case 9223372036854775807:
                                    if v2 {
                                        v31 = append(v31, fi50)
                                        var v86 uint64 = 2
                                        switch (v52 + "") {
                                        case "\n\t":
                                            v31 = append(v31, v1)
                                        case "a":
                                            var v87 int64 = 1
                                            _ = v87
                                        case "":
                                            v88 := GenS12{F0: "foo bar", F1: (fi46 % ((v1) | 1))}
                                            var v89 float64 = (((1e10 - ((v53 * 1e10) / (float64(float32(float64((0.0 + 0.0)))) / v53))) / 1.0) / v53)
                                            v1 = fi46
                                            _ = v88
                                            _ = v89
                                        }
                                        _ = v86
                                    }
                                case 1:
                                    fmt.Println(v35, v51, v2)
                                case 4294967295:
                                    v3 -= (^(v3))
                                default:
                                    var v90 float32 = 0.1
                                    _ = v90
                                }
                                fmt.Println(v51, v51)
                                _ = v85
                            }
                        case -464:
                            var v91 uint64 = 4294967295
                            _ = v91
                        case 2147483647:
                            v35 += (v35 & (v35 << (((^(v51))) & 63)))
                        default:
                            switch v52 {
                            case "":
                                v92 := GenS14{F0: 32767, F1: (fi46 ^ fi46)}
                                _ = v92
                            default:
                                var v93 uint = 496
                                v94 := GenS1{F0: v35}
                                switch ("a" + v52) {
                                case "foo bar":
                                    _ = len(v31)
                                }
                                _ = v93
                                _ = v94
                            }
                        }
                    }
                }
                var v95 int = (-(v1))
                _ = v53
                _ = v54
                _ = v95
            }
            if (("a" + "") >= v52) {
                v96 := GenS15{F0: -9223372036854775808, F1: -1}
                if ((v2 && (v35 <= (v35 / (((v35 ^ 65535)) | 1)))) && v2) {
                    v97 := GenS5{F0: v51, F1: (fi50 & 1)}
                    v98 := GenS15{F0: -1, F1: genFn0(v3, uint16(v35))}
                    _ = v97
                    _ = v98
                } else {
                    if v2 {
                        var v99 uint64 = 200
                        v2 = ((genFn0(v3, 1) >= 2147483647) || (0 >= 100))
                        _ = v99
                    } else {
                        v3 = (v3 << ((4294967295) & 63))
                        fmt.Println(v51, fi50, fi50)
                    }
                    v35 -= uint32(v3)
                }
                v1 -= 9223372036854775807
                _ = v96
            } else {
                _ = len(v31)
                v2 = (0 == 32767)
                fmt.Println(v2)
            }
            var v100 uint32 = (v35 & 0)
            if true {
                v101 := GenS16{F0: -24, F1: v35, F2: (!true)}
                _ = v101
            }
            _ = fi50
            _ = v100
            _ = v51
            _ = v52
        }
        _ = fi46
    }
    for fi102 := 0; fi102 < 14; fi102++ {
        for fi103 := 0; fi103 < 18; fi103++ {
            if (42 > -9223372036854775808) {
                if len(v31) > 0 {
                    v31[0] = v1
                }
            }
            _ = fi103
        }
        v3 += (v3 << ((0) & 63))
        _ = fi102
    }
    v2 = v2
    v3 = (v3 ^ 121)
    for fi104 := 0; fi104 < 8; fi104++ {
        v2 = (true || v2)
        v2 = (v2 || ((!(-32768 >= -32768)) && (1.5 < float32(1.5))))
        var v105 int32 = genFn0(v3, 2)
        _ = fi104
        _ = v105
    }
    if (true || (!v2)) {
        v1 = (v1 * (v1 + (^(v1))))
        v1 = v1
    }
    v2 = (v2 && v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
}
