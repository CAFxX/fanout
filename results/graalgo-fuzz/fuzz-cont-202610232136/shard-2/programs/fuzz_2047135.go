package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 int64; F1 int; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS10 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 int8, p1 int, p2 uint16) int {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 14; fi4++ {
        p1 = fi4
        _ = fi4
    }
    var v5 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 uint, p2 int) int16 {
    v6 := []int{42, -2147483648, 763, p2}
    if (-126 < 100) {
        v7 := GenS1{F0: -652, F1: 42}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 0
    var v2 int16 = -1
    v3 := []int{-2147483648, 127, 32768, 21}
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            if true {
                var v9 uint64 = 0
                _ = v9
            } else {
                if len(v3) > 0 {
                    v3[0] = (fi8 * 271)
                }
                v2 = (v2 & (^(v2)))
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 int8 = 42
                _ = fi10
                _ = v11
            }
            _ = fi8
        }
        if (0.1 > (0.0 * 1e10)) {
            if true {
                var v12 uint32 = uint32(v1)
                v1 += v1
                fmt.Println(v1, v2)
                _ = v12
            } else {
                v1 = v1
            }
            v13 := GenS2{F0: -2147483648}
            var v14 uint8 = (v1 / (((v1 | (v1 * v1))) | 1))
            _ = v13
            _ = v14
        } else {
            switch "\n\t" {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = genFn0(127, 32767, 32767)
                }
            }
            v3 = append(v3, 198)
        }
        switch 331 {
        case 905:
            v1 += v1
        default:
            var v15 int8 = -1
            _ = v15
        }
    } else {
        var v16 uint = 4294967295
        switch "hello" {
        case "foo bar":
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 int = (fi17 - fi17)
                if false {
                    v19 := GenS1{F0: (v2 << ((v16) & 63)), F1: (v16 << ((v16) & 63))}
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
        default:
            if false {
                var v20 int64 = 1
                var v21 uint32 = 65535
                var v22 bool = false
                _ = v20
                _ = v21
                _ = v22
            } else {
                v2 -= (v2 + 42)
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := GenS1{F0: int16(v16), F1: 18446744073709551615}
                    var v25 uint16 = 32767
                    fmt.Println(fi23, v2)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
        _ = v16
    }
    var v26 int8 = -38
    v3 = append(v3, 127)
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 int8 = (v26 / ((v26) | 1))
        _ = fi27
        _ = v28
    }
    if len(v3) > 0 {
        v3[0] = genFn0(v26, 0, 255)
    }
    v1 -= v1
    var v29 int = -903
    if (42 == v1) {
        if true {
            var v30 int = (v29 % ((genFn0((v26 >> ((630) & 63)), v29, 65535)) | 1))
            _ = v30
        } else {
            if true {
                fmt.Println(v2, v1, v26)
            } else {
                for fi31 := 0; fi31 < 11; fi31++ {
                    switch (v29 % ((-361) | 1)) {
                    case 1:
                        if (-9223372036854775808 != 42) {
                            v3 = append(v3, v29)
                        }
                    case 10:
                        for fi32 := 0; fi32 < 18; fi32++ {
                            switch fi31 {
                            case 2147483647:
                                if true {
                                    v29 += fi31
                                    switch fi32 {
                                    case 32768:
                                        v33 := GenS1{F0: genFn1(true, 699, (fi31 - 127)), F1: 42}
                                        _ = v33
                                    case 4294967295:
                                        v26 += v26
                                    case 127:
                                        switch v29 {
                                        case 127:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case -687:
                                            if (!true) {
                                                v29 += 32768
                                            } else {
                                                var v34 uint32 = 895
                                                _ = v34
                                            }
                                        }
                                    }
                                    v26 = (v26 / ((v26) | 1))
                                } else {
                                    var v35 float32 = 0.1
                                    var v36 uint32 = 1
                                    _ = v35
                                    _ = v36
                                }
                            case 382:
                                if (!true) {
                                    v1 = 44
                                    fmt.Println(v1)
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = (-(fi31))
                                    }
                                }
                                v3 = append(v3, genFn0(100, v29, 1))
                                v26 = v26
                            default:
                                var v37 int32 = 575
                                _ = v37
                            }
                            if (-1 != 416) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v26, fi31)
                                v2 = v2
                            } else {
                                if (0.1 == ((((1.5 * 1e10) / float64(1e10)) / 1.0) + float64((float32(float64(3.14159)) / 0.1)))) {
                                    if true {
                                        v38 := GenS3{F0: 0}
                                        v2 = (v2 >> ((255) & 63))
                                        _ = len(v3)
                                        _ = v38
                                    }
                                    fmt.Println(v2)
                                }
                                _ = len(v3)
                                var v39 int32 = 0
                                _ = v39
                            }
                            v2 += v2
                            var v40 uint32 = 294
                            if false {
                                var v41 uint = 353
                                _ = len(v3)
                                _ = len(v3)
                                _ = v41
                            } else {
                                v42 := GenS4{F0: (fi31 >> ((42) & 63)), F1: v1, F2: "a"}
                                _ = v42
                            }
                            _ = fi32
                            _ = v40
                        }
                    case 32768:
                        v29 -= (v29 | -1)
                    default:
                        v2 += 107
                    }
                    for fi43 := 0; fi43 < 15; fi43++ {
                        v44 := GenS5{F0: 9223372036854775807, F1: fi43, F2: -1}
                        v29 += (fi31 | (fi43 << ((799) & 63)))
                        _ = fi43
                        _ = v44
                    }
                    var v45 uint = 0
                    _ = fi31
                    _ = v45
                }
            }
        }
        fmt.Println(v2, v26, v26)
        var v46 uint8 = (^(v1))
        _ = v46
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "a" {
        case "hello":
            v1 = v1
        case "a":
            if len(v3) > 0 {
                v3[0] = v29
            }
        default:
            fmt.Println(v1)
        }
        var v47 string = ("\n\t" + ("hello" + (("" + "\n\t") + ("\n\t" + (("foo bar" + ("" + ("a" + "a"))) + "hello")))))
        _ = v47
    }
    var v48 int32 = -395
    switch "\n\t" {
    case "\n\t":
        for fi49 := 0; fi49 < 12; fi49++ {
            v3 = append(v3, 42)
            _ = fi49
        }
    case "hello":
        v50 := GenS6{F0: ("hello" + ("foo bar" + ""))}
        _ = v50
    case "a":
        for fi51 := 0; fi51 < 9; fi51++ {
            if true {
                v52 := GenS6{F0: ("foo bar" + "hello")}
                var v53 int8 = (v26 % ((v26) | 1))
                v29 += -2147483648
                _ = v52
                _ = v53
            } else {
                if ((^(v1)) <= 1) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    fmt.Println(v48, v2, v1)
                    switch genFn0(95, (fi51 / ((int(v2)) | 1)), 42) {
                    case 127:
                        for fi54 := 0; fi54 < 20; fi54++ {
                            if (!(9223372036854775807 != -1)) {
                                var v55 float32 = (0.1 - float32(0.0))
                                fmt.Println(fi51, v55)
                                var v56 uint32 = 1
                                _ = v55
                                _ = v56
                            }
                            v3 = append(v3, 655)
                            fmt.Println(v26, v29)
                            if true {
                                var v57 int64 = 0
                                v58 := GenS7{F0: 42}
                                _ = v57
                                _ = v58
                            } else {
                                v59 := GenS5{F0: 9223372036854775807, F1: v29, F2: (v26 * v26)}
                                _ = v59
                            }
                            _ = fi54
                        }
                    case 0:
                        for fi60 := 0; fi60 < 9; fi60++ {
                            var v61 uint16 = 962
                            var v62 float32 = 1.5
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                    case 1:
                        for fi63 := 0; fi63 < 3; fi63++ {
                            v64 := GenS2{F0: 171}
                            v65 := GenS7{F0: 656}
                            if (true || true) {
                                var v66 int = (v29 + (fi51 / (((fi51 - v29)) | 1)))
                                _ = v66
                            }
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    default:
                        var v67 float32 = (3.14159 - 3.14159)
                        _ = v67
                    }
                    if (true || (v1 >= (v1 ^ (v1 ^ v1)))) {
                        for fi68 := 0; fi68 < 12; fi68++ {
                            if (!(957 > 2147483648)) {
                                if (("a" > "foo bar") || true) {
                                    var v69 int16 = (^(v2))
                                    if (v2 == genFn1(true, 1, (fi51 - 42))) {
                                        v29 -= (v29 * (v29 % (((-(fi51))) | 1)))
                                    } else {
                                        switch (fi68 / (((fi68 * genFn0((v26 << ((255) & 63)), v29, 845))) | 1)) {
                                        case -9223372036854775808:
                                            var v70 int16 = v2
                                            _ = v70
                                        case -2147483648:
                                            fmt.Println(v1, v69, v26)
                                        case -591:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            if (true && false) {
                                                v48 = v48
                                                fmt.Println(fi68, v1, fi51)
                                                if false {
                                                    var v71 int32 = -494
                                                    var v72 float32 = 0.0
                                                    var v73 float32 = ((3.14159 + (v72 + v72)) * float32((1e10 * (0.0 + -2.25))))
                                                    v74 := GenS1{F0: (v69 % ((v2) | 1)), F1: 135}
                                                    v3 = append(v3, -843)
                                                    _ = v71
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                }
                                            } else {
                                                var v75 uint32 = 1
                                                _ = v75
                                            }
                                        }
                                        fmt.Println(fi51)
                                        v76 := GenS6{F0: "hello"}
                                        if true {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v77 uint16 = 840
                                            v78 := GenS8{F0: (v1 * v1), F1: (v1 / ((1) | 1))}
                                            _ = v77
                                            _ = v78
                                        } else {
                                            if true {
                                                if ((2147483647 > 0) || (1e10 <= (1e10 + (float64(1e10) / 0.1)))) {
                                                    v69 = int16(v48)
                                                }
                                            } else {
                                                var v79 string = "foo bar"
                                                fmt.Println(v29, fi51)
                                                _ = v79
                                            }
                                            if (506 < -175) {
                                                var v80 float64 = (((float64((((3.14159 * ((((((1.5 - 0.1) + (0.1 / ((float32(0.1) + 1.5) + -2.25))) / 1.0) - 0.1) - 0.0) - 1e10)) / 1.0) / 1.0)) + 1e10) * -2.25) * 3.14159)
                                                _ = len(v3)
                                                _ = v80
                                            }
                                        }
                                        fmt.Println(v26, fi51, v26)
                                        _ = v76
                                    }
                                    var v81 uint8 = (v1 % (((v1 | v1)) | 1))
                                    _ = v69
                                    _ = v81
                                } else {
                                    var v82 uint = 352
                                    fmt.Println(v29, v48, v2)
                                    v82 = (^(v82))
                                    _ = v82
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                var v83 int = 2147483647
                                _ = v83
                            }
                            var v84 int32 = 312
                            var v85 bool = false
                            _ = fi68
                            _ = v84
                            _ = v85
                        }
                    }
                }
            }
            _ = fi51
        }
    default:
        v86 := GenS6{F0: "\n\t"}
        _ = v86
    }
    v87 := GenS8{F0: v1, F1: (v1 >> ((18446744073709551615) & 63))}
    v48 -= (v48 >> ((65535) & 63))
    v29 -= genFn0(v26, (v29 - v29), 42)
    var v88 float64 = (1e10 / 1.5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        for fi89 := 0; fi89 < 13; fi89++ {
            v90 := GenS4{F0: -2147483648, F1: v1, F2: "\n\t"}
            for fi91 := 0; fi91 < 17; fi91++ {
                switch (v29 + v29) {
                case 9223372036854775807:
                    v26 = v26
                case 979:
                    if (uint32(v1) <= 0) {
                        var v92 int8 = v26
                        v2 += v2
                        var v93 int32 = v48
                        _ = v92
                        _ = v93
                    }
                case 0:
                    switch (v29 * v29) {
                    case 494:
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    case 32768:
                        switch 32767 {
                        case -2147483648:
                            switch -921 {
                            case -1:
                                v29 += (fi91 ^ v29)
                                var v94 int16 = v2
                                _ = v94
                            case 879:
                                _ = len(v3)
                            case 32768:
                                v95 := GenS9{F0: 0, F1: true, F2: 1}
                                _ = v95
                            default:
                                var v96 uint = 65535
                                _ = v96
                            }
                        case 747:
                            v48 = (^(v48))
                        case 718:
                            v2 = (v2 / ((339) | 1))
                        default:
                            var v97 string = "foo bar"
                            v98 := GenS10{F0: false, F1: (-(v48)), F2: v26}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v97
                            _ = v98
                        }
                    case 4294967295:
                        var v99 int = (v29 - v29)
                        _ = v99
                    }
                default:
                    v100 := GenS5{F0: 0, F1: (v29 ^ (fi91 & fi91)), F2: (v26 / ((1) | 1))}
                    _ = v100
                }
                var v101 int32 = v48
                _ = fi91
                _ = v101
            }
            _ = fi89
            _ = v90
        }
        var v102 uint16 = 32767
        if (true || true) {
            v103 := GenS6{F0: ("" + "a")}
            _ = v103
        }
        _ = v102
    }
    var v104 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v104
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v48
    _ = v87
    _ = v88
}
