package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 string; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint8; F2 bool }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int16 }
func genFn0(p0 string) uint {
    var v4 int8 = -128
    switch (p0 + (p0 + ("\n\t" + "\n\t"))) {
    case "foo bar":
        var v5 uint = 255
        _ = v5
    case "a":
        fmt.Println(v4)
    case "\n\t":
        v6 := []int{-9223372036854775808, -2147483648}
        _ = v6
    }
    fmt.Println(p0)
    v4 = v4
    switch (p0 + p0) {
    case "\n\t":
        switch 127 {
        case 32768:
            if (94 != 255) {
                v7 := GenS1{F0: -877, F1: 255, F2: -2147483648}
                var v8 float32 = -2.25
                _ = v7
                _ = v8
            }
        case 128:
            v9 := []int{}
            _ = v9
        case 9223372036854775807:
            var v10 bool = ((false && false) && false)
            _ = v10
        default:
            v11 := []int{-1, 1, 1, 32768}
            _ = v11
        }
    case "foo bar":
        v12 := GenS1{F0: 9223372036854775807, F1: 255, F2: 42}
        _ = v12
    case "a":
        var v13 uint64 = 748
        _ = v13
    }
    _ = p0
    _ = v4
    return 4294967295
}
func genFn1(p0 int, p1 int32, p2 int16) int32 {
    for fi14 := 0; fi14 < 9; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            switch -9223372036854775808 {
            case 0:
                if false {
                    v16 := GenS1{F0: p0, F1: 4294967295, F2: fi15}
                    if (float32(1.5) <= float32(1e10)) {
                        if false {
                            v17 := []int{4294967295, p0, 1}
                            _ = v17
                        } else {
                            v18 := []int{(p0 + (fi14 ^ (^(fi15)))), fi14}
                            var v19 float64 = 0.0
                            _ = v18
                            _ = v19
                        }
                        var v20 string = "\n\t"
                        p2 = p2
                        _ = v20
                    } else {
                        switch "" {
                        case "\n\t":
                            v21 := GenS1{F0: 32767, F1: 270, F2: p0}
                            _ = v21
                        case "foo bar":
                            v22 := GenS2{F0: 65535, F1: 1, F2: (p2 ^ 32767)}
                            _ = v22
                        default:
                            var v23 uint16 = 0
                            _ = v23
                        }
                    }
                    _ = v16
                }
            case 32768:
                v24 := GenS2{F0: 4294967295, F1: 25, F2: -32768}
                _ = v24
            case -353:
                v25 := GenS2{F0: 42, F1: 156, F2: 42}
                _ = v25
            }
            _ = fi15
        }
        switch fi14 {
        case -1:
            fmt.Println(fi14, fi14)
        case 118:
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 uint64 = 0
                if false {
                    p2 += p2
                } else {
                    v28 := []int{p0}
                    _ = v28
                }
                var v29 uint32 = 47
                _ = fi26
                _ = v27
                _ = v29
            }
        }
        _ = fi14
    }
    v30 := GenS2{F0: 942, F1: 128, F2: (p2 % ((1000) | 1))}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return p1
}

func main() {
    var v1 int64 = 249
    var v2 uint64 = 975
    var v3 uint8 = 136
    if false {
        if false {
            v31 := GenS3{F0: v1, F1: 993, F2: 42}
            var v32 int32 = 42
            _ = v31
            _ = v32
        }
        v33 := GenS3{F0: v1, F1: 262, F2: 902}
        _ = v33
    } else {
        v3 -= (v3 >> ((18446744073709551615) & 63))
        if true {
            v3 = v3
            switch 42 {
            case 2147483647:
                v2 -= v2
            case 127:
                if (255 >= 1) {
                    if (-128 <= 100) {
                        if (330 >= 450) {
                            switch 4294967295 {
                            case 4294967295:
                                if (!(42 < 1)) {
                                    v3 -= (^(v3))
                                    v34 := []int{}
                                    _ = v34
                                }
                            case 127:
                                var v35 int64 = v1
                                _ = v35
                            default:
                                for fi36 := 0; fi36 < 16; fi36++ {
                                    for fi37 := 0; fi37 < 18; fi37++ {
                                        v38 := GenS4{F0: (v3 ^ v3), F1: -9223372036854775808}
                                        v2 += 42
                                        _ = fi37
                                        _ = v38
                                    }
                                    var v39 uint16 = 0
                                    if ("\n\t" < "hello") {
                                        var v40 int8 = 108
                                        _ = v40
                                    } else {
                                        if true {
                                            var v41 float32 = (1.5 - 1.5)
                                            var v42 uint = genFn0("")
                                            _ = v41
                                            _ = v42
                                        }
                                        for fi43 := 0; fi43 < 12; fi43++ {
                                            v44 := GenS4{F0: v3, F1: 187}
                                            switch ("" + "foo bar") {
                                            case "foo bar":
                                                v45 := GenS4{F0: (v3 % (((v3 >> ((694) & 63))) | 1)), F1: (fi43 * (fi36 * (fi36 % (((fi36 % (((fi43 * fi43)) | 1))) | 1))))}
                                                v46 := GenS2{F0: 676, F1: v3, F2: 42}
                                                var v47 int8 = 0
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            case "a":
                                                if ((true || false) && ((!false) && ((^(v1)) != (^(v1))))) {
                                                    v48 := []int{128, (fi36 / ((127) | 1))}
                                                    if ((v2 | v2) < v2) {
                                                        fmt.Println(fi43, v3)
                                                    }
                                                    v39 = v39
                                                    _ = v48
                                                } else {
                                                    var v49 int16 = int16(v3)
                                                    v50 := GenS2{F0: 411, F1: 86, F2: v49}
                                                    var v51 uint8 = v3
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case "\n\t":
                                                if true {
                                                    var v52 int64 = 908
                                                    v53 := GenS4{F0: v3, F1: (-(fi36))}
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    if false {
                                                        v39 = 65535
                                                        if ((((("hello" + ("" + "foo bar")) + "") + "hello") <= ("a" + "a")) || true) {
                                                            v3 -= (v3 / (((v3 / (((^(v3))) | 1))) | 1))
                                                            fmt.Println(v39)
                                                            fmt.Println(v1)
                                                            v54 := GenS5{F0: (v3 / ((v3) | 1)), F1: 2, F2: (v39 / ((1) | 1))}
                                                            _ = v54
                                                        }
                                                    }
                                                    fmt.Println(v1, v3)
                                                    v55 := GenS6{F0: 42, F1: 65535}
                                                    var v56 string = ("foo bar" + "hello")
                                                    var v57 bool = false
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                }
                                            default:
                                                v58 := GenS7{F0: 4294967295}
                                                _ = v58
                                            }
                                            _ = fi43
                                            _ = v44
                                        }
                                    }
                                    _ = fi36
                                    _ = v39
                                }
                            }
                            for fi59 := 0; fi59 < 20; fi59++ {
                                var v60 int32 = 229
                                if false {
                                    v61 := GenS2{F0: 65535, F1: v3, F2: 1}
                                    switch "foo bar" {
                                    case "a":
                                        var v62 int16 = 1000
                                        _ = v62
                                    case "hello":
                                        for fi63 := 0; fi63 < 17; fi63++ {
                                            v60 = (-(v60))
                                            _ = fi63
                                        }
                                    case "":
                                        switch "hello" {
                                        case "\n\t":
                                            switch -1 {
                                            case -1:
                                                v1 += (v1 << ((0) & 63))
                                            case 32768:
                                                fmt.Println(v2)
                                            default:
                                                if (!(float32(float64(-2.25)) == -2.25)) {
                                                    v64 := GenS8{F0: (v2 | (v2 | (v2 ^ 809))), F1: "foo bar", F2: genFn1(fi59, int32(v2), 32767)}
                                                    for fi65 := 0; fi65 < 5; fi65++ {
                                                        if ((true && true) && (!false)) {
                                                            v2 -= 1
                                                            v66 := GenS9{F0: 2, F1: 1, F2: true}
                                                            _ = v66
                                                        } else {
                                                            v67 := GenS10{F0: 65535, F1: -22}
                                                            v60 -= v60
                                                            if true {
                                                                if true {
                                                                    v68 := GenS1{F0: (fi65 + fi65), F1: genFn0("foo bar"), F2: int(v3)}
                                                                    fmt.Println(fi59, fi59)
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        v69 := []int{}
                                                                        v2 += (v2 | 805)
                                                                        v70 := GenS11{F0: 1}
                                                                        _ = v69
                                                                        _ = v70
                                                                    case "foo bar":
                                                                        v60 += v60
                                                                    case "":
                                                                        v3 = (v3 ^ (v3 >> ((846) & 63)))
                                                                    }
                                                                    _ = v68
                                                                }
                                                                var v71 uint64 = (v2 / (((v2 ^ (v2 % ((v2) | 1)))) | 1))
                                                                _ = v71
                                                            }
                                                            _ = v67
                                                        }
                                                        _ = fi65
                                                    }
                                                    var v72 uint = 0
                                                    _ = v64
                                                    _ = v72
                                                } else {
                                                    fmt.Println(v2)
                                                    fmt.Println(v2)
                                                }
                                            }
                                        case "":
                                            var v73 uint8 = (v3 << ((130) & 63))
                                            _ = v73
                                        }
                                    default:
                                        for fi74 := 0; fi74 < 10; fi74++ {
                                            v75 := []int{(fi74 + 247), (fi74 >> ((301) & 63)), (^(fi59))}
                                            var v76 uint16 = 255
                                            var v77 int64 = v1
                                            _ = fi74
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                    _ = v61
                                } else {
                                    var v78 uint64 = 4294967295
                                    v79 := GenS10{F0: 42, F1: 42}
                                    v80 := GenS12{F0: 1}
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                                fmt.Println(v3, v2)
                                _ = fi59
                                _ = v60
                            }
                            var v81 float64 = 0.1
                            _ = v81
                        } else {
                            var v82 float64 = (1e10 + 1.5)
                            if false {
                                var v83 bool = (true && true)
                                v84 := GenS8{F0: (v2 ^ (v2 - v2)), F1: ("foo bar" + (("a" + ("" + (((((((("a" + "a") + ("hello" + "foo bar")) + "a") + "") + "\n\t") + "\n\t") + ("" + ("\n\t" + ""))) + ("\n\t" + "a")))) + ("foo bar" + "a"))), F2: -1000}
                                _ = v83
                                _ = v84
                            }
                            _ = v82
                        }
                        fmt.Println(v1, v3, v1)
                    } else {
                        v85 := GenS5{F0: 42, F1: 42, F2: 65535}
                        _ = v85
                    }
                    if false {
                        switch 127 {
                        case 0:
                            v86 := []int{}
                            _ = v86
                        case 127:
                            if (!false) {
                                var v87 uint16 = 1
                                _ = v87
                            } else {
                                v88 := GenS12{F0: -30}
                                _ = v88
                            }
                        case 2147483647:
                            fmt.Println(v3)
                        }
                        v89 := GenS8{F0: (^(v2)), F1: ("a" + "foo bar"), F2: -1}
                        _ = v89
                    } else {
                        v90 := GenS12{F0: 1000}
                        v91 := GenS6{F0: -21, F1: 559}
                        v2 = v2
                        _ = v90
                        _ = v91
                    }
                } else {
                    switch -9223372036854775808 {
                    case 765:
                        var v92 int = int(v2)
                        _ = v92
                    case 2147483647:
                        var v93 float32 = ((0.1 * 1e10) + (float32((3.14159 / 0.1)) - -2.25))
                        _ = v93
                    case 127:
                        switch "foo bar" {
                        case "a":
                            v1 = (v1 ^ int64(v3))
                        case "hello":
                            fmt.Println(v1)
                        }
                    }
                    v94 := []int{128, int(v2), 1}
                    var v95 bool = ((685 >= 93) && ((v1 >> ((255) & 63)) <= v1))
                    _ = v94
                    _ = v95
                }
            }
        }
    }
    if true {
        for fi96 := 0; fi96 < 19; fi96++ {
            fmt.Println(fi96)
            _ = fi96
        }
        for fi97 := 0; fi97 < 2; fi97++ {
            v1 += v1
            _ = fi97
        }
        var v98 int64 = v1
        _ = v98
    }
    v2 = 2
    switch 926 {
    case 42:
        var v99 bool = false
        _ = v99
    case 658:
        var v100 bool = (1 >= (v2 + v2))
        _ = v100
    default:
        v1 = v1
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        var v101 int = -691
        _ = v101
    default:
        if false {
            v102 := []int{}
            v1 -= v1
            v102 = append(v102, 4294967295)
            _ = v102
        }
    }
    v103 := GenS7{F0: (v2 + v2)}
    fmt.Println(v3, v2, v3)
    v104 := GenS7{F0: v2}
    v105 := GenS1{F0: -9223372036854775808, F1: genFn0("\n\t"), F2: int(v3)}
    fmt.Println(v3, v2, v3)
    v1 = v1
    v106 := []int{}
    _ = len(v106)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v106:", len(v106), v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v103
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v3
}
