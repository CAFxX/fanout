package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint32; F1 bool; F2 int }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint8; F1 uint16; F2 string }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0() int32 {
    var v6 float32 = 3.14159
    v6 += 1e10
    _ = v6
    return 0
}
func genFn1(p0 uint64) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            v8 := []int{}
            _ = v8
        } else {
            v9 := GenS1{F0: 0, F1: 2147483648}
            _ = v9
        }
        fmt.Println(p0, p0)
        fmt.Println(p0, p0)
        _ = fi7
    }
    p0 = (p0 << ((255) & 63))
    switch int(p0) {
    case 32767:
        switch 736 {
        case 0:
            v10 := GenS1{F0: 9223372036854775807, F1: 721}
            _ = v10
        case 128:
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 2; fi12++ {
                    var v13 uint = 1
                    var v14 int32 = genFn0()
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = fi11
            }
        }
    case 460:
        switch "" {
        case "a":
            if true {
                p0 = p0
                if (-1 <= 2147483647) {
                    var v15 string = "foo bar"
                    v16 := []int{-9223372036854775808, 32768, 127}
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS2{F0: 990}
                    _ = v17
                }
                fmt.Println(p0)
            } else {
                var v18 int32 = genFn0()
                var v19 string = "foo bar"
                var v20 uint = 1
                _ = v18
                _ = v19
                _ = v20
            }
        case "foo bar":
            for fi21 := 0; fi21 < 12; fi21++ {
                if false {
                    v22 := []int{fi21}
                    _ = v22
                } else {
                    p0 = p0
                    for fi23 := 0; fi23 < 19; fi23++ {
                        p0 -= 42
                        var v24 int64 = -898
                        var v25 float64 = 0.1
                        var v26 string = ("foo bar" + (("hello" + "") + "\n\t"))
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                }
                p0 = (^(p0))
                _ = fi21
            }
        default:
            v27 := []int{}
            _ = v27
        }
    case -1:
        v28 := GenS2{F0: 217}
        _ = v28
    }
    if (false || true) {
        var v29 int64 = -794
        v30 := GenS3{F0: v29}
        _ = v29
        _ = v30
    } else {
        v31 := []int{-1, -9223372036854775808, 32767, 32768}
        var v32 uint16 = 0
        p0 += 2
        _ = v31
        _ = v32
    }
    _ = p0
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (-2.25 - 0.1)
    var v3 int8 = 0
    var v4 uint16 = 255
    v5 := []string{"\n\t", "a", "hello", "foo bar"}
    v33 := GenS1{F0: -823, F1: 465}
    delete(v1, 255)
    v2 -= ((1e10 * v2) * 1.5)
    for fi34 := 0; fi34 < 10; fi34++ {
        if false {
            switch ("" + "hello") {
            case "a":
                fmt.Println(v3)
            case "hello":
                v35 := v1[875]
                _ = v35
            default:
                v36 := GenS1{F0: 0, F1: 2147483647}
                _ = v36
            }
        } else {
            if (false || (!(!true))) {
                var v37 int = (fi34 % ((fi34) | 1))
                _ = v37
            } else {
                fmt.Println(v2)
                var v38 int = fi34
                _ = v38
            }
            var v39 string = ("hello" + "foo bar")
            for fi40 := 0; fi40 < 3; fi40++ {
                v39 = (v39 + "foo bar")
                _ = fi40
            }
            _ = v39
        }
        _ = fi34
    }
    v41 := GenS4{F0: 127, F1: 0}
    if (-731 >= -2147483648) {
        v2 -= (float64(float32((-2.25 * v2))) * 1.5)
    } else {
        var v42 int32 = 0
        v43 := v1[333]
        _ = v42
        _ = v43
    }
    v2 = ((float64(3.14159) / v2) + -2.25)
    if (int16(v4) == 42) {
        var v44 int8 = genFn1(0)
        switch 9223372036854775807 {
        case 42:
            v4 = v4
        case -305:
            if false {
                if (true && true) {
                    var v45 int64 = 9223372036854775807
                    var v46 uint16 = 32767
                    _ = v45
                    _ = v46
                } else {
                    for fi47 := 0; fi47 < 8; fi47++ {
                        _ = len(v5)
                        var v48 uint8 = 128
                        var v49 int32 = int32(v3)
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v5 = append(v5, ("" + "\n\t"))
                }
                v2 -= v2
                var v50 uint8 = 255
                _ = v50
            }
        case 678:
            var v51 float32 = 0.1
            _ = v51
        default:
            v52 := GenS5{F0: 227, F1: true, F2: -9223372036854775808}
            _ = v52
        }
        _ = v44
    }
    if false {
        switch -1 {
        case 32767:
            fmt.Println(v2)
        case 9223372036854775807:
            var v53 uint = 42
            _ = v53
        case 2147483647:
            v3 -= -50
        }
        v4 = 998
    }
    v54 := GenS6{F0: 0, F1: uint64(v4)}
    switch "\n\t" {
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            switch ("foo bar" + ((("foo bar" + "hello") + "a") + "\n\t")) {
            case "hello":
                v1[65535] = true
            case "foo bar":
                var v55 uint32 = 4294967295
                _ = v55
            case "":
                var v56 int16 = -832
                _ = v56
            default:
                v57 := GenS5{F0: 2147483647, F1: false, F2: 9223372036854775807}
                _ = v57
            }
        case "hello":
            _ = len(v5)
        case "a":
            if (!(v2 != v2)) {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        var v58 int64 = -9223372036854775808
                        v2 += (v2 * v2)
                        for fi59 := 0; fi59 < 18; fi59++ {
                            if true {
                                for fi60 := 0; fi60 < 15; fi60++ {
                                    var v61 int = fi60
                                    var v62 int16 = 0
                                    var v63 uint16 = 255
                                    var v64 float32 = 1.5
                                    switch (fi59 | 844) {
                                    case 32767:
                                        if len(v5) > 0 {
                                            v5[0] = "a"
                                        }
                                    case 4294967295:
                                        if ((v3 < (v3 >> ((336) & 63))) || (false && false)) {
                                            _ = len(v1)
                                        } else {
                                            v61 = (v61 - 42)
                                        }
                                    case 1:
                                        v65 := GenS5{F0: 1, F1: (!true), F2: (v61 % (((v61 << ((4294967295) & 63))) | 1))}
                                        _ = v65
                                    default:
                                        v66 := GenS6{F0: 121, F1: 768}
                                        if true {
                                            _ = len(v1)
                                            v4 = v63
                                            v58 += 815
                                            v61 += int(v4)
                                        }
                                        v1[53] = (0.1 < (v2 / -2.25))
                                        _ = v66
                                    }
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                var v67 int8 = -128
                                _ = v67
                            } else {
                                fmt.Println(fi59, v2, fi59)
                                if true {
                                    v68 := GenS3{F0: v58}
                                    v2 += ((1.5 * ((0.1 / 3.14159) + v2)) + v2)
                                    _ = v68
                                }
                            }
                            _ = fi59
                        }
                        _ = v58
                    } else {
                        if ((1.5 * v2) <= (v2 * (0.0 * 3.14159))) {
                            for fi69 := 0; fi69 < 6; fi69++ {
                                v70 := GenS1{F0: -513, F1: 2}
                                v2 -= 0.0
                                _ = fi69
                                _ = v70
                            }
                        } else {
                            if false {
                                var v71 float64 = v2
                                _ = v71
                            }
                            v72 := GenS2{F0: 2}
                            if (!(144 >= 223)) {
                                var v73 int8 = v3
                                _ = v73
                            }
                            _ = v72
                        }
                        v1[995] = true
                        var v74 int = 127
                        _ = v74
                    }
                default:
                    v3 = (^(v3))
                }
                for fi75 := 0; fi75 < 2; fi75++ {
                    for fi76 := 0; fi76 < 6; fi76++ {
                        fmt.Println(v3, v2)
                        var v77 int8 = (-(v3))
                        _ = fi76
                        _ = v77
                    }
                    v78 := GenS4{F0: 128, F1: 813}
                    _ = fi75
                    _ = v78
                }
                v79 := GenS3{F0: 1}
                _ = v79
            } else {
                var v80 float64 = 3.14159
                _ = v80
            }
        }
    }
    var v81 int = int(v3)
    var v82 uint32 = 446
    var v83 uint8 = 48
    var v84 int16 = 42
    for fi85 := 0; fi85 < 14; fi85++ {
        v82 -= (v82 + v82)
        for fi86 := 0; fi86 < 16; fi86++ {
            v87 := GenS4{F0: v83, F1: (v84 / ((v84) | 1))}
            var v88 uint32 = (v82 % ((v82) | 1))
            _ = fi86
            _ = v87
            _ = v88
        }
        _ = fi85
    }
    v82 = v82
    if false {
        if (((1e10 / 1.5) - (0.0 * 1e10)) < 1e10) {
            switch v81 {
            case 32:
                var v89 int8 = int8(v83)
                _ = v89
            }
            if ((!(!(v82 >= v82))) || true) {
                var v90 uint64 = 42
                for fi91 := 0; fi91 < 4; fi91++ {
                    var v92 int32 = -2147483648
                    for fi93 := 0; fi93 < 7; fi93++ {
                        if (v3 != v3) {
                            fmt.Println(v4)
                            switch "a" {
                            case "foo bar":
                                v92 -= 539
                            case "\n\t":
                                v94 := GenS7{F0: "a", F1: v90}
                                _ = v94
                            case "hello":
                                v4 += v4
                            }
                            delete(v1, 88)
                        } else {
                            switch -2147483648 {
                            case 592:
                                v95 := GenS5{F0: (v82 * 1), F1: (uint(v4) < 661), F2: v81}
                                _ = v95
                            default:
                                var v96 uint64 = (v90 >> ((4294967295) & 63))
                                _ = v96
                            }
                            v1[(v82 ^ uint32(v4))] = (v83 >= (v83 % ((v83) | 1)))
                            switch (fi91 << ((38) & 63)) {
                            case 2147483647:
                                v92 += int32(v4)
                            case 128:
                                if true {
                                    fmt.Println(v82)
                                    delete(v1, 4294967295)
                                    var v97 float64 = v2
                                    v90 = (v90 & 554)
                                    if (!true) {
                                        var v98 uint64 = v90
                                        _ = v98
                                    }
                                    _ = v97
                                }
                            }
                        }
                        _ = fi93
                    }
                    _ = fi91
                    _ = v92
                }
                v81 = v81
                _ = v90
            } else {
                switch v81 {
                case 32768:
                    switch "hello" {
                    case "foo bar":
                        v99 := GenS8{F0: ((v84 & (v84 / ((v84) | 1))) >= 1)}
                        _ = v99
                    case "":
                        for fi100 := 0; fi100 < 12; fi100++ {
                            var v101 bool = true
                            var v102 string = "hello"
                            switch v81 {
                            case 587:
                                for fi103 := 0; fi103 < 14; fi103++ {
                                    v104 := GenS9{F0: v83, F1: v4, F2: ""}
                                    v2 -= v2
                                    v2 -= (((float64(1e10) * (float64(0.0) * (v2 / (1.5 * (((v2 - -2.25) * v2) / (float64((((3.14159 + (0.1 * (1.5 / 3.14159))) * 1.5) * -2.25)) / (3.14159 - (v2 - ((1e10 + ((v2 - 0.1) + ((v2 * ((1.5 - float64((((0.0 + ((1.5 + (-2.25 - 1.5)) + (-2.25 / 1.5))) - 1e10) - (1.5 * ((0.1 / 1e10) / 1.5))))) - -2.25)) / (3.14159 - (3.14159 + v2))))) - (3.14159 * -2.25)))))))))) + v2) + (v2 + 0.0))
                                    _ = fi103
                                    _ = v104
                                }
                            case 934:
                                if ((v4 << ((65535) & 63)) >= (^(v4))) {
                                    fmt.Println(v101, v101, v84)
                                    v82 += 65535
                                    var v105 int64 = 0
                                    _ = v105
                                }
                            case 32767:
                                switch (^(v81)) {
                                case 32768:
                                    v82 = (^(v82))
                                case 32767:
                                    v106 := GenS10{F0: 85, F1: (v83 * 2)}
                                    _ = v106
                                default:
                                    v81 -= v81
                                }
                            }
                            _ = fi100
                            _ = v101
                            _ = v102
                        }
                    case "hello":
                        if true {
                            for fi107 := 0; fi107 < 11; fi107++ {
                                _ = len(v1)
                                _ = fi107
                            }
                            if true {
                                if (true && true) {
                                    if true {
                                        v108 := GenS11{F0: (-(v81)), F1: v83, F2: 42}
                                        _ = v108
                                    }
                                } else {
                                    var v109 uint8 = 0
                                    fmt.Println(v83, v83)
                                    for fi110 := 0; fi110 < 2; fi110++ {
                                        v5 = append(v5, (("hello" + ("hello" + "hello")) + "a"))
                                        _ = fi110
                                    }
                                    _ = v109
                                }
                                for fi111 := 0; fi111 < 18; fi111++ {
                                    v83 = 128
                                    _ = fi111
                                }
                                v112 := GenS2{F0: 1}
                                _ = v112
                            }
                        }
                    }
                case 586:
                    fmt.Println(v81, v4, v2)
                case 32767:
                    for fi113 := 0; fi113 < 11; fi113++ {
                        if ((false && false) || (!(v3 < (v3 % (((v3 >> ((2) & 63))) | 1))))) {
                            for fi114 := 0; fi114 < 10; fi114++ {
                                var v115 bool = false
                                if v115 {
                                    v2 -= 1e10
                                } else {
                                    var v116 int16 = (v84 ^ 472)
                                    _ = v116
                                }
                                var v117 uint = 18446744073709551615
                                _ = fi114
                                _ = v115
                                _ = v117
                            }
                        }
                        fmt.Println(v83, v4)
                        _ = fi113
                    }
                default:
                    for fi118 := 0; fi118 < 13; fi118++ {
                        v83 = v83
                        v3 -= 100
                        _ = fi118
                    }
                }
            }
            switch "\n\t" {
            case "a":
                v1[75] = ((0 < 423) && (65535 >= 703))
            case "foo bar":
                v119 := GenS1{F0: 9223372036854775807, F1: (v82 % (((v82 - v82)) | 1))}
                _ = v119
            default:
                v3 -= 31
            }
        }
        v84 -= -462
        _ = len(v1)
    } else {
        if true {
            var v120 uint64 = 0
            v82 = (^(v82))
            _ = v120
        } else {
            v1[(v82 / ((v82) | 1))] = false
        }
        var v121 float64 = float64(-2.25)
        _ = len(v1)
        _ = v121
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v41
    _ = v5
    _ = v54
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
