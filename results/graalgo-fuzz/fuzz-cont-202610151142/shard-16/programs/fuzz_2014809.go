package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 int8; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint64; F1 int32; F2 bool }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint8; F1 int }
func genFn0(p0 string) bool {
    var v4 uint = 801
    var v5 int8 = 100
    fmt.Println(p0, v4)
    for fi6 := 0; fi6 < 16; fi6++ {
        p0 = ""
        for fi7 := 0; fi7 < 15; fi7++ {
            p0 = p0
            _ = fi7
        }
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 float32 = (float32(3.14159) + (1.5 - 3.14159))
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    if (!((0 == 42) || (2 >= 255))) {
        var v10 float32 = 1.5
        _ = v10
    } else {
        v11 := []int{32767, -702}
        fmt.Println(v5, p0, p0)
        for fi12 := 0; fi12 < 16; fi12++ {
            v11 = append(v11, 9223372036854775807)
            _ = fi12
        }
        _ = v11
    }
    _ = p0
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 int8, p1 int8) int32 {
    v13 := []int{-808, 9223372036854775807, 32768, 0}
    var v14 float32 = float32(0.0)
    switch -16 {
    case 127:
        v15 := GenS1{F0: true}
        _ = v15
    case 128:
        if false {
            var v16 float64 = -2.25
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            if len(v13) > 0 {
                v13[0] = 127
            }
            _ = v16
        }
    case 2147483647:
        p0 = (p0 >> ((440) & 63))
    }
    p1 += 42
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return -1
}

func main() {
    var v1 float64 = 3.14159
    var v2 string = "foo bar"
    var v3 float32 = 3.14159
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS2{F0: (fi17 ^ fi17)}
        _ = fi17
        _ = v18
    }
    var v19 string = v2
    var v20 int8 = 0
    var v21 uint16 = 582
    fmt.Println(v19, v1)
    var v22 int = -1
    v2 = v2
    if true {
        for fi23 := 0; fi23 < 19; fi23++ {
            v19 = (((v2 + (v2 + (v19 + (v2 + (v2 + ("foo bar" + "foo bar")))))) + (v2 + v2)) + ("foo bar" + "\n\t"))
            var v24 uint = 436
            var v25 uint8 = 0
            _ = fi23
            _ = v24
            _ = v25
        }
        switch v19 {
        case "":
            v26 := GenS1{F0: true}
            _ = v26
        case "foo bar":
            if true {
                if (true || true) {
                    v21 = v21
                    fmt.Println(v3)
                    var v27 uint32 = 42
                    _ = v27
                }
                var v28 uint = 572
                v29 := GenS3{F0: -32768}
                _ = v28
                _ = v29
            } else {
                var v30 uint = uint(v21)
                for fi31 := 0; fi31 < 4; fi31++ {
                    v32 := []int{fi31}
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
        default:
            v33 := GenS2{F0: v22}
            _ = v33
        }
        var v34 string = v2
        _ = v34
    }
    v21 += 65535
    if (272 != 0) {
        switch (v22 | (v22 * (v22 | (v22 | (v22 * v22))))) {
        case -741:
            for fi35 := 0; fi35 < 18; fi35++ {
                v21 += v21
                var v36 int = int(v21)
                fmt.Println(v19)
                _ = fi35
                _ = v36
            }
        default:
            v37 := GenS4{F0: 2147483647, F1: 0}
            _ = v37
        }
        v38 := GenS5{F0: 65535, F1: v21}
        _ = v38
    } else {
        v2 = v2
        if false {
            if true {
                for fi39 := 0; fi39 < 11; fi39++ {
                    if false {
                        v40 := []int{}
                        v20 += (^(v20))
                        var v41 uint16 = v21
                        _ = v40
                        _ = v41
                    } else {
                        var v42 int = 277
                        _ = v42
                    }
                    if true {
                        v22 = v22
                        var v43 uint8 = 44
                        fmt.Println(v20)
                        _ = v43
                    } else {
                        if false {
                            if true {
                                switch v2 {
                                case "hello":
                                    for fi44 := 0; fi44 < 8; fi44++ {
                                        v22 += v22
                                        v45 := GenS6{F0: 0}
                                        v46 := []int{}
                                        var v47 uint = 255
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                case "\n\t":
                                    v48 := GenS7{F0: -2147483648, F1: false}
                                    _ = v48
                                default:
                                    v20 += int8(v22)
                                }
                                v49 := []int{(v22 / ((-9223372036854775808) | 1)), 128}
                                var v50 uint64 = 524
                                _ = v49
                                _ = v50
                            }
                            v2 = v19
                        } else {
                            var v51 uint32 = 646
                            _ = v51
                        }
                    }
                    v2 = v19
                    _ = fi39
                }
            }
            v1 -= v1
            v2 = "foo bar"
        }
    }
    fmt.Println(v19, v3, v3)
    switch (v19 + (v2 + ((("a" + ("" + "hello")) + "") + (v19 + ((v19 + "\n\t") + ("\n\t" + ((((v2 + (("foo bar" + v2) + "a")) + v19) + "\n\t") + v19))))))) {
    case "foo bar":
        v22 -= 121
    case "hello":
        if (true || ((-1 > 42) || (!true))) {
            var v52 int8 = (-(v20))
            _ = v52
        } else {
            var v53 uint64 = 989
            if false {
                fmt.Println(v53)
            } else {
                fmt.Println(v53)
                v54 := GenS3{F0: 0}
                _ = v54
            }
            _ = v53
        }
    }
    if ((!(!true)) && (!(!true))) {
        if true {
            var v55 bool = (2147483647 < 65535)
            var v56 uint = 395
            switch -768 {
            case 1:
                var v57 bool = (42 > 4294967295)
                _ = v57
            case 32768:
                var v58 int32 = genFn1(v20, 100)
                _ = v58
            }
            _ = v55
            _ = v56
        }
        for fi59 := 0; fi59 < 6; fi59++ {
            switch v19 {
            case "a":
                v60 := []int{v22, v22, 1, (fi59 << ((943) & 63))}
                _ = v60
            case "\n\t":
                if true {
                    v61 := GenS8{F0: (v20 - (v20 >> ((442) & 63))), F1: 42, F2: 0}
                    fmt.Println(fi59, v20)
                    var v62 float64 = v1
                    _ = v61
                    _ = v62
                }
            case "":
                var v63 int32 = 229
                _ = v63
            default:
                fmt.Println(v20, fi59, v1)
            }
            var v64 float32 = float32(v1)
            _ = fi59
            _ = v64
        }
        for fi65 := 0; fi65 < 8; fi65++ {
            v20 += -1
            _ = fi65
        }
    } else {
        fmt.Println(v2, v22)
        for fi66 := 0; fi66 < 14; fi66++ {
            if (0 != 42) {
                v21 -= (v21 % ((v21) | 1))
                var v67 float64 = v1
                if (true || true) {
                    var v68 float32 = 0.0
                    v69 := GenS4{F0: (v22 % (((fi66 >> ((942) & 63))) | 1)), F1: 72}
                    _ = v68
                    _ = v69
                }
                _ = v67
            }
            switch fi66 {
            case 128:
                v70 := []int{(fi66 % ((2147483647) | 1)), -625, (fi66 + (-(v22))), v22}
                _ = v70
            default:
                if ((32767 != 1000) && (false || false)) {
                    for fi71 := 0; fi71 < 10; fi71++ {
                        var v72 uint = 2
                        if ("hello" != v19) {
                            v73 := GenS4{F0: -534, F1: v72}
                            switch (v22 ^ v22) {
                            case 42:
                                if false {
                                    fmt.Println(v22)
                                } else {
                                    var v74 uint32 = 801
                                    v2 = v19
                                    var v75 int32 = 0
                                    _ = v74
                                    _ = v75
                                }
                            default:
                                v1 = v1
                                var v76 bool = false
                                var v77 uint16 = v21
                                var v78 float32 = (v3 - v3)
                                v79 := GenS9{F0: 695, F1: 1, F2: v76}
                                _ = v76
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                            _ = v73
                        } else {
                            v20 = (v20 & (-(v20)))
                        }
                        _ = fi71
                        _ = v72
                    }
                } else {
                    switch v22 {
                    case -25:
                        v80 := GenS10{F0: -9223372036854775808, F1: v20}
                        _ = v80
                    case 9223372036854775807:
                        v81 := []int{}
                        _ = v81
                    default:
                        v20 += (v20 & 42)
                    }
                    switch (^(v22)) {
                    case 83:
                        for fi82 := 0; fi82 < 6; fi82++ {
                            v83 := GenS6{F0: 920}
                            switch (v2 + ((v2 + v2) + v2)) {
                            case "foo bar":
                                v20 = (-(v20))
                            default:
                                if true {
                                    var v84 string = v2
                                    _ = v84
                                }
                            }
                            _ = fi82
                            _ = v83
                        }
                    case 89:
                        var v85 float64 = 1.5
                        _ = v85
                    case 42:
                        switch ("\n\t" + v19) {
                        case "a":
                            var v86 float32 = 1.5
                            _ = v86
                        case "foo bar":
                            if true {
                                v87 := GenS3{F0: 32767}
                                var v88 uint16 = 42
                                v21 = v88
                                _ = v87
                                _ = v88
                            } else {
                                fmt.Println(v20, fi66)
                            }
                        case "":
                            switch v2 {
                            case "foo bar":
                                var v89 uint64 = 924
                                _ = v89
                            case "\n\t":
                                switch (v2 + (("hello" + "foo bar") + (v2 + ((v2 + ((v2 + (("a" + v2) + "")) + (v19 + (v2 + (("hello" + v19) + v2))))) + "")))) {
                                case "a":
                                    switch "" {
                                    case "":
                                        if (false && (((0 >= 903) && false) || ((^(v21)) != 32767))) {
                                            v90 := GenS9{F0: 582, F1: genFn1(12, v20), F2: ((1 < 59) && (false && false))}
                                            var v91 int32 = 1
                                            _ = v90
                                            _ = v91
                                        } else {
                                            v92 := []int{v22, fi66, fi66}
                                            v20 = v20
                                            _ = v92
                                        }
                                    case "foo bar":
                                        fmt.Println(v19, v20)
                                    default:
                                        var v93 uint16 = v21
                                        _ = v93
                                    }
                                case "\n\t":
                                    v94 := GenS9{F0: 428, F1: -707, F2: ((!true) || false)}
                                    _ = v94
                                case "foo bar":
                                    for fi95 := 0; fi95 < 20; fi95++ {
                                        if false {
                                            v2 = v2
                                            if (855 <= 189) {
                                                var v96 int32 = 0
                                                v21 += (v21 - (v21 & 42))
                                                v97 := []int{}
                                                _ = v96
                                                _ = v97
                                            }
                                            v98 := []int{}
                                            _ = v98
                                        }
                                        _ = fi95
                                    }
                                }
                            default:
                                v2 = ""
                            }
                        }
                    default:
                        v3 = 0.1
                    }
                    v99 := []int{(-(v22)), v22, (fi66 % ((v22) | 1))}
                    _ = v99
                }
            }
            _ = fi66
        }
    }
    v100 := GenS6{F0: 18446744073709551615}
    v101 := GenS11{F0: -980}
    var v102 int16 = -573
    var v103 uint16 = 65535
    v104 := GenS12{F0: 0, F1: (v22 * v22)}
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v103
    _ = v104
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
