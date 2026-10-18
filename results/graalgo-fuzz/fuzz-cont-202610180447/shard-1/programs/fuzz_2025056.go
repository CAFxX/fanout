package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int32; F2 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool; F1 string; F2 int16 }
type GenS12 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 uint8, p1 int16, p2 int8) int {
    var v7 int = -417
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v7 = fi9
            v7 += int(p0)
            _ = fi9
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            p0 = p0
            p2 = (p2 / (((p2 | (^(p2)))) | 1))
            switch ((("" + "a") + "hello") + "") {
            case "foo bar":
                v11 := []int{(fi10 - v7), (fi10 & 32768), (fi8 * -963)}
                _ = v11
            default:
                p0 -= p0
            }
            _ = fi10
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(p1, fi12)
        v13 := GenS1{F0: (p1 * p1), F1: 42}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        switch (("" + "hello") + "") {
        case "":
            var v15 int64 = -1
            _ = v15
        case "foo bar":
            var v16 bool = true
            _ = v16
        case "\n\t":
            if false {
                var v17 uint32 = 750
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 bool = true
                    switch fi14 {
                    case 42:
                        switch "" {
                        case "":
                            var v20 uint64 = 18446744073709551615
                            _ = v20
                        }
                    case 0:
                        var v21 uint32 = v17
                        _ = v21
                    default:
                        if (-848 == 0) {
                            var v22 uint32 = v17
                            fmt.Println(v22, v22)
                            v23 := GenS1{F0: p1, F1: 65535}
                            _ = v22
                            _ = v23
                        } else {
                            if (((p0 == 36) && (!v19)) || ((false || (!(!false))) || v19)) {
                                switch -1 {
                                case 2147483647:
                                    var v24 string = "foo bar"
                                    _ = v24
                                case 32767:
                                    v25 := GenS2{F0: 1, F1: p0}
                                    _ = v25
                                }
                                p1 -= (^(p1))
                            }
                        }
                    }
                    v19 = true
                    if (-171 != 924) {
                        if (42 > -9223372036854775808) {
                            p2 -= p2
                        }
                        fmt.Println(fi14, p1)
                    } else {
                        if v19 {
                            v17 = 2147483647
                        } else {
                            var v26 uint64 = 18446744073709551615
                            if ((^(p1)) == (p1 - (p1 ^ (p1 + (p1 >> ((0) & 63)))))) {
                                v17 = (v17 / (((v17 >> ((18446744073709551615) & 63))) | 1))
                                if (uint64(p0) == (v26 | uint64(p0))) {
                                    fmt.Println(v17)
                                    v27 := []int{}
                                    if v19 {
                                        fmt.Println(v7)
                                        var v28 int16 = (p1 ^ p1)
                                        _ = v28
                                    } else {
                                        p0 = (p0 & p0)
                                        fmt.Println(v7, v26, p2)
                                        var v29 uint16 = 65535
                                        if len(v27) >= 2 {
                                            v27 = v27[1:len(v27)]
                                        }
                                        _ = v29
                                    }
                                    _ = v27
                                }
                            } else {
                                fmt.Println(p2, v17)
                                if true {
                                    v17 = 65
                                    var v30 int = (^(fi18))
                                    _ = v30
                                }
                                switch (-(fi18)) {
                                case 2147483647:
                                    v31 := GenS2{F0: 0, F1: p0}
                                    _ = v31
                                }
                            }
                            p2 = int8(v17)
                            _ = v26
                        }
                        if ((727 == 488) && v19) {
                            if v19 {
                                if v19 {
                                    v32 := GenS1{F0: (p1 % ((0) | 1)), F1: 1}
                                    v17 -= 2147483647
                                    v33 := GenS2{F0: -278, F1: (p0 % (((p0 ^ (^(p0)))) | 1))}
                                    _ = v32
                                    _ = v33
                                }
                                v34 := []int{(fi14 + (-(v7))), -9223372036854775808}
                                _ = v34
                            } else {
                                var v35 float32 = 0.0
                                _ = v35
                            }
                            v36 := GenS2{F0: 1, F1: 238}
                            _ = v36
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v37 := GenS3{F0: 42}
                                _ = v37
                            default:
                                if ((!(v19 && (v19 || (!v19)))) || (213 > 812)) {
                                    var v38 int16 = p1
                                    v39 := GenS2{F0: -9223372036854775808, F1: (p0 >> ((0) & 63))}
                                    v19 = false
                                    _ = v38
                                    _ = v39
                                }
                            }
                        }
                        var v40 uint = uint(v17)
                        _ = v40
                    }
                    _ = fi18
                    _ = v19
                }
                v41 := GenS1{F0: (p1 - 1), F1: uint(v17)}
                _ = v17
                _ = v41
            } else {
                if true {
                    p1 = -1
                    fmt.Println(v7, v7)
                    p2 = p2
                } else {
                    v42 := GenS4{F0: 808, F1: (p1 - -691), F2: (p0 % ((2) | 1))}
                    switch "a" {
                    case "a":
                        for fi43 := 0; fi43 < 18; fi43++ {
                            var v44 uint8 = p0
                            _ = fi43
                            _ = v44
                        }
                    case "\n\t":
                        p1 -= p1
                    case "":
                        v45 := GenS3{F0: 100}
                        _ = v45
                    }
                    _ = v42
                }
                fmt.Println(fi14, p2)
                for fi46 := 0; fi46 < 1; fi46++ {
                    v47 := GenS5{F0: p2}
                    _ = fi46
                    _ = v47
                }
            }
        default:
            var v48 int8 = p2
            _ = v48
        }
        var v49 int8 = p2
        _ = fi14
        _ = v49
    }
    if true {
        var v50 uint8 = p0
        if true {
            if (!(2147483648 != 2)) {
                if true {
                    v51 := []int{}
                    _ = v51
                }
                p1 = p1
                fmt.Println(v7)
            }
        } else {
            var v52 int16 = p1
            if true {
                v50 = v50
            } else {
                v53 := []int{v7, (v7 - 42), 42, int(p2)}
                _ = v53
            }
            _ = v52
        }
        var v54 int32 = -845
        _ = v50
        _ = v54
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}
func genFn1(p0 uint, p1 int16) uint32 {
    switch "" {
    case "":
        v55 := GenS5{F0: -1}
        _ = v55
    case "a":
        v56 := []int{0, 32768}
        _ = v56
    case "hello":
        var v57 float64 = 3.14159
        _ = v57
    default:
        v58 := []int{32768, -943, 1}
        _ = v58
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        v60 := GenS6{F0: 525}
        _ = fi59
        _ = v60
    }
    var v61 uint64 = 1
    for fi62 := 0; fi62 < 6; fi62++ {
        v63 := GenS3{F0: 127}
        var v64 int = genFn0(218, -1, -128)
        if (100 == 65) {
            var v65 int16 = p1
            p0 -= 65535
            if true {
                var v66 int16 = v65
                if (2147483647 == fi62) {
                    v64 = (v64 % ((fi62) | 1))
                }
                _ = v66
            }
            _ = v65
        }
        _ = fi62
        _ = v63
        _ = v64
    }
    _ = p0
    _ = p1
    _ = v61
    return 528
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[uint32]bool)
    var v3 int16 = 0
    var v4 uint = 524
    v5 := make(map[int]uint32)
    var v6 float32 = 3.14159
    v6 = v6
    for fi67 := 0; fi67 < 8; fi67++ {
        fmt.Println(v6, v3, v4)
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "")
        }
        v2[1] = (false && false)
        _ = fi67
    }
    var v68 uint64 = 18446744073709551615
    for fi69 := 0; fi69 < 17; fi69++ {
        fmt.Println(fi69)
        v70 := GenS5{F0: 1}
        _ = fi69
        _ = v70
    }
    if ((false && (false || true)) && false) {
        switch ("hello" + "\n\t") {
        case "hello":
            var v71 float32 = 0.0
            _ = v71
        case "":
            if false {
                v6 = float32(1.5)
                fmt.Println(v4)
                var v72 uint8 = 2
                _ = v72
            } else {
                delete(v2, 253)
                if false {
                    fmt.Println(v3, v6, v6)
                    _ = len(v2)
                    if true {
                        switch genFn0(130, v3, 100) {
                        case 127:
                            switch -2147483648 {
                            case -9223372036854775808:
                                delete(v2, uint32(v4))
                            case 752:
                                v68 -= 18446744073709551615
                            case 9223372036854775807:
                                v1 = append(v1, "")
                            default:
                                for fi73 := 0; fi73 < 1; fi73++ {
                                    var v74 uint16 = 42
                                    _ = fi73
                                    _ = v74
                                }
                            }
                        case -2147483648:
                            var v75 int = -284
                            _ = v75
                        case 32767:
                            v6 -= 0.0
                        }
                    } else {
                        fmt.Println(v68)
                        v68 -= (v68 % ((v68) | 1))
                        fmt.Println(v4, v4, v4)
                    }
                } else {
                    v76 := GenS5{F0: 0}
                    _ = v76
                }
                if false {
                    var v77 int64 = -136
                    _ = v77
                } else {
                    switch ("" + "a") {
                    case "a":
                        fmt.Println(v4, v68)
                    case "foo bar":
                        switch 128 {
                        case 1:
                            if (!(-912 >= 42)) {
                                var v78 uint32 = 220
                                if len(v1) > 0 {
                                    v1[0] = ((((("foo bar" + "foo bar") + "foo bar") + "hello") + ("" + "")) + "foo bar")
                                }
                                v3 += -205
                                _ = v78
                            } else {
                                if true {
                                    fmt.Println(v68, v3)
                                }
                            }
                        case 0:
                            v3 += (v3 | v3)
                        case 4294967295:
                            var v79 int32 = -121
                            _ = v79
                        }
                    default:
                        var v80 uint8 = 0
                        _ = v80
                    }
                }
            }
        case "\n\t":
            fmt.Println(v3, v68, v68)
        default:
            if (!(false || (!(!false)))) {
                var v81 uint32 = 105
                if false {
                    switch ("hello" + "") {
                    case "\n\t":
                        var v82 uint8 = 0
                        _ = v82
                    case "a":
                        switch "hello" {
                        case "\n\t":
                            if ((v81 << ((v4) & 63)) >= v81) {
                                if true {
                                    v83 := GenS7{F0: 0, F1: 1, F2: 0}
                                    for fi84 := 0; fi84 < 11; fi84++ {
                                        if true {
                                            var v85 float32 = (v6 - v6)
                                            for fi86 := 0; fi86 < 2; fi86++ {
                                                v87 := GenS8{F0: false}
                                                _ = fi86
                                                _ = v87
                                            }
                                            if (0 != 255) {
                                                switch -205 {
                                                case 0:
                                                    switch fi84 {
                                                    case 834:
                                                        var v88 int8 = -52
                                                        _ = v88
                                                    default:
                                                        v1 = append(v1, "foo bar")
                                                    }
                                                }
                                                v1 = append(v1, (("foo bar" + "foo bar") + ("a" + "foo bar")))
                                            } else {
                                                v89 := GenS3{F0: 42}
                                                for fi90 := 0; fi90 < 8; fi90++ {
                                                    v91 := GenS9{F0: (v68 ^ 157), F1: 255, F2: 127}
                                                    v68 += (v68 & v68)
                                                    if false {
                                                        fmt.Println(v6, fi90)
                                                        if (!true) {
                                                            v92 := GenS10{F0: 100}
                                                            v4 = v4
                                                            _ = v92
                                                        }
                                                        var v93 int = (fi84 * (-(fi90)))
                                                        _ = v93
                                                    }
                                                    _ = fi90
                                                    _ = v91
                                                }
                                                _ = v89
                                            }
                                            _ = v85
                                        } else {
                                            v94 := GenS11{F0: (463 <= 0), F1: ("" + "foo bar"), F2: (v3 << ((v4) & 63))}
                                            _ = v94
                                        }
                                        v2[1] = true
                                        v3 = -1
                                        _ = fi84
                                    }
                                    var v95 int64 = 9223372036854775807
                                    _ = v83
                                    _ = v95
                                } else {
                                    v4 -= (v4 + v4)
                                    v96 := GenS7{F0: genFn1(57, 1), F1: 2147483647, F2: -202}
                                    _ = v96
                                }
                                var v97 uint64 = (v68 * (v68 << ((v4) & 63)))
                                fmt.Println(v81)
                                _ = v97
                            } else {
                                if true {
                                    v98 := GenS2{F0: int64(v68), F1: uint8(v68)}
                                    v81 = v81
                                    _ = v98
                                }
                            }
                        case "hello":
                            var v99 int8 = 1
                            _ = v99
                        default:
                            for fi100 := 0; fi100 < 5; fi100++ {
                                var v101 int32 = 611
                                fmt.Println(v3, v68)
                                _ = fi100
                                _ = v101
                            }
                        }
                    }
                    fmt.Println(v68)
                    var v102 int = 585
                    _ = v102
                } else {
                    v6 = (float32(1.5) + 1.5)
                    _ = len(v1)
                }
                v68 -= 2
                _ = v81
            } else {
                switch int(v4) {
                case 4294967295:
                    var v103 uint32 = 769
                    _ = v103
                case 32767:
                    for fi104 := 0; fi104 < 14; fi104++ {
                        var v105 int8 = -1
                        if (1 == 1) {
                            v106 := GenS12{F0: ("\n\t" + "a"), F1: 32767, F2: 42}
                            for fi107 := 0; fi107 < 9; fi107++ {
                                var v108 uint8 = 0
                                v3 = v3
                                _ = fi107
                                _ = v108
                            }
                            for fi109 := 0; fi109 < 15; fi109++ {
                                delete(v5, genFn0(0, (v3 + int16(fi109)), 21))
                                _ = fi109
                            }
                            _ = v106
                        }
                        _ = fi104
                        _ = v105
                    }
                case 0:
                    v1 = append(v1, "hello")
                }
            }
        }
    }
    _ = len(v5)
    v3 -= v3
    switch 127 {
    case 127:
        if true {
            for fi110 := 0; fi110 < 5; fi110++ {
                v6 -= float32((-2.25 - 1e10))
                _ = fi110
            }
        } else {
            if true {
                v3 -= int16(v68)
                var v111 uint16 = 32767
                fmt.Println(v3, v68, v68)
                _ = v111
            } else {
                var v112 uint16 = 947
                _ = v112
            }
            v3 += v3
            v113 := GenS12{F0: ("a" + ""), F1: 146, F2: 756}
            _ = v113
        }
    }
    v4 = v4
    v68 -= v68
    for fi114 := 0; fi114 < 7; fi114++ {
        delete(v5, (^(fi114)))
        _ = fi114
    }
    switch genFn0(uint8(v68), (v3 | v3), 16) {
    case 4294967295:
        if len(v1) > 0 {
            v1[0] = (("\n\t" + "foo bar") + "foo bar")
        }
    }
    v115 := GenS6{F0: v68}
    fmt.Println(v3, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
