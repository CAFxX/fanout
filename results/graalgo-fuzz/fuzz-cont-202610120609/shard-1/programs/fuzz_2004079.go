package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 string; F1 uint; F2 int16 }
type GenS7 struct { F0 int64; F1 uint; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint }
type GenS9 struct { F0 int8; F1 uint32; F2 int64 }
type GenS10 struct { F0 uint8; F1 uint32; F2 uint32 }
func genFn0() uint16 {
    var v7 int16 = 1
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{fi8}
        fmt.Println(fi8, v7)
        if len(v9) > 0 {
            v9[0] = fi8
        }
        _ = fi8
        _ = v9
    }
    _ = v7
    return 42
}

func main() {
    var v1 int = -9223372036854775808
    v2 := []string{}
    var v3 float64 = 1e10
    var v4 int16 = 1000
    var v5 uint64 = 138
    v6 := make(map[int]int)
    v10 := GenS1{F0: 255, F1: v1}
    fmt.Println(v1)
    fmt.Println(v3, v1, v3)
    switch 32767 {
    case 341:
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 bool = (v4 >= v4)
            if v12 {
                delete(v6, (fi11 % ((v1) | 1)))
            }
            v1 = -13
            _ = fi11
            _ = v12
        }
    case -23:
        v13 := GenS1{F0: 2, F1: 0}
        _ = v13
    case 1:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint32 = 4294967295
            v2 = append(v2, "a")
            for fi16 := 0; fi16 < 20; fi16++ {
                switch fi16 {
                case -2147483648:
                    if (v4 < v4) {
                        v1 = fi16
                    }
                case 127:
                    v17 := GenS2{F0: -9223372036854775808}
                    _ = v17
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println(v3, v5)
    if (0 <= 42) {
        var v18 uint = 2
        if false {
            switch ("" + (("\n\t" + "hello") + ((("\n\t" + "\n\t") + "\n\t") + ((("foo bar" + ("\n\t" + "a")) + "a") + ("foo bar" + "\n\t"))))) {
            case "foo bar":
                if false {
                    fmt.Println(v1, v18, v3)
                    for fi19 := 0; fi19 < 1; fi19++ {
                        if true {
                            fmt.Println(v18, v4)
                        } else {
                            switch -9223372036854775808 {
                            case 42:
                                v6[(fi19 - -9223372036854775808)] = fi19
                            }
                            v3 = v3
                        }
                        _ = fi19
                    }
                    fmt.Println(v3, v18, v18)
                } else {
                    v20 := v6[v1]
                    _ = v20
                }
            case "hello":
                _ = len(v6)
            case "":
                _ = len(v6)
            }
            v3 -= v3
        }
        fmt.Println(v3, v1)
        _ = v18
    } else {
        for fi21 := 0; fi21 < 9; fi21++ {
            switch "hello" {
            case "foo bar":
                var v22 int64 = 1
                _ = v22
            case "":
                var v23 uint = 825
                _ = v23
            case "hello":
                v24 := v6[(^(fi21))]
                _ = v24
            }
            var v25 int32 = 0
            _ = fi21
            _ = v25
        }
        v5 = v5
    }
    switch 1 {
    case 4294967295:
        switch "hello" {
        case "foo bar":
            for fi26 := 0; fi26 < 17; fi26++ {
                switch (-(v1)) {
                case 4294967295:
                    v6[32768] = (v1 | (v1 >> ((255) & 63)))
                case 176:
                    if (0 >= -82) {
                        v5 -= v5
                        v27 := GenS3{F0: 2, F1: false}
                        _ = v27
                    }
                }
                for fi28 := 0; fi28 < 9; fi28++ {
                    var v29 float64 = ((v3 - v3) + float64(0.0))
                    v4 -= 1
                    if (1 < 32767) {
                        _ = len(v2)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v30 uint32 = 2147483648
                        v4 += (v4 & (v4 + (v4 ^ (v4 + v4))))
                        _ = v30
                    }
                    v31 := GenS4{F0: 286, F1: 766}
                    if false {
                        var v32 float64 = 1e10
                        v32 = (float64(((1e10 - -2.25) / ((0.0 * -2.25) / (-2.25 * ((0.1 - 3.14159) / float32((1.5 * ((((v3 * 0.0) - ((v29 / -2.25) / v29)) * ((v29 / (v32 / 1.5)) * -2.25)) / float64(-2.25))))))))) / float64(float32(v29)))
                        switch "" {
                        case "":
                            fmt.Println(fi26, v5, v4)
                            v32 -= 0.1
                        case "foo bar":
                            var v33 uint32 = 0
                            _ = v33
                        default:
                            var v34 uint16 = 1
                            _ = v34
                        }
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                    _ = v31
                }
                _ = fi26
            }
        case "a":
            for fi35 := 0; fi35 < 17; fi35++ {
                if ("a" == "foo bar") {
                    var v36 int8 = -128
                    v37 := GenS4{F0: 42, F1: 0}
                    var v38 string = ""
                    _ = v36
                    _ = v37
                    _ = v38
                } else {
                    v5 += v5
                    for fi39 := 0; fi39 < 7; fi39++ {
                        if false {
                            if false {
                                if (!false) {
                                    var v40 float32 = (1.5 - (((float32(v3) + 0.0) + 3.14159) - 1e10))
                                    _ = v40
                                } else {
                                    fmt.Println(fi35, v4)
                                    v41 := GenS5{F0: -2147483648, F1: 58, F2: 42}
                                    _ = v41
                                }
                                var v42 int64 = -9223372036854775808
                                if true {
                                    var v43 int8 = -128
                                    _ = v43
                                }
                                _ = v42
                            } else {
                                v5 += (^(v5))
                                var v44 uint16 = 42
                                _ = v44
                            }
                            if false {
                                if false {
                                    v1 += fi39
                                    v45 := GenS5{F0: -929, F1: 150, F2: 299}
                                    if (!(-991 < -434)) {
                                        var v46 uint16 = 65535
                                        fmt.Println(fi39)
                                        var v47 uint16 = (^(v46))
                                        var v48 string = ("hello" + "hello")
                                        if false {
                                            v1 += fi39
                                        } else {
                                            v49 := v6[(v1 * fi35)]
                                            _ = v49
                                        }
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    } else {
                                        fmt.Println(fi39, v5)
                                        v50 := v6[fi35]
                                        _ = v50
                                    }
                                    _ = v45
                                }
                                fmt.Println(v4, v3, v4)
                                fmt.Println(v5, v5, v4)
                                if ((v3 * (v3 * v3)) >= float64(3.14159)) {
                                    v1 = 42
                                    v51 := GenS6{F0: ("foo bar" + ("\n\t" + "\n\t")), F1: 16, F2: (v4 / ((v4) | 1))}
                                    v52 := v6[(fi39 ^ (-(fi35)))]
                                    v1 = (v1 - v52)
                                    switch (("a" + "\n\t") + (("\n\t" + ("foo bar" + "hello")) + "hello")) {
                                    case "foo bar":
                                        v5 = (v5 - (v5 ^ (v5 / (((v5 - (^(v5)))) | 1))))
                                    }
                                    switch 775 {
                                    case 443:
                                        if true {
                                            switch (fi35 & fi39) {
                                            case 0:
                                                v5 = (v5 | (v5 - v5))
                                            default:
                                                switch (fi39 & -1) {
                                                case 294:
                                                    if (!true) {
                                                        if (false || true) {
                                                            var v53 uint8 = 209
                                                            _ = v53
                                                        } else {
                                                            fmt.Println(v5, fi39, v1)
                                                            v3 = float64(-2.25)
                                                        }
                                                        fmt.Println(v3, v1, fi35)
                                                        v54 := GenS7{F0: 9, F1: 297, F2: -1}
                                                        var v55 uint32 = 218
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                case 42:
                                                    var v56 bool = true
                                                    _ = v56
                                                }
                                            }
                                            v57 := GenS6{F0: "\n\t", F1: 178, F2: v4}
                                            fmt.Println(v5, v1, v4)
                                            _ = v57
                                        } else {
                                            v1 = (v52 + (fi35 & (v52 & (-(v52)))))
                                            fmt.Println(v3, v4, fi35)
                                            if false {
                                                var v58 uint64 = 2
                                                _ = v58
                                            }
                                        }
                                    case 0:
                                        v5 -= (^(v5))
                                        var v59 int8 = -126
                                        switch "foo bar" {
                                        case "hello":
                                            v1 += int(v5)
                                        }
                                        _ = v59
                                    case -9223372036854775808:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        switch "hello" {
                                        case "a":
                                            v1 = (^(fi39))
                                        case "hello":
                                            v60 := GenS8{F0: 127, F1: 453}
                                            _ = v60
                                        default:
                                            var v61 int = (fi39 >> ((4294967295) & 63))
                                            _ = v61
                                        }
                                    }
                                    _ = v51
                                    _ = v52
                                } else {
                                    switch (("\n\t" + "a") + ("hello" + ("a" + "foo bar"))) {
                                    case "":
                                        var v62 uint16 = genFn0()
                                        _ = v62
                                    case "a":
                                        v1 -= -2147483648
                                    case "hello":
                                        delete(v6, 127)
                                    }
                                    var v63 int8 = 0
                                    switch "a" {
                                    case "hello":
                                        v3 -= v3
                                    case "":
                                        var v64 float64 = -2.25
                                        _ = v64
                                    case "foo bar":
                                        var v65 int64 = 689
                                        _ = v65
                                    default:
                                        if false {
                                            v2 = append(v2, "a")
                                            v63 += 110
                                            var v66 string = ""
                                            _ = v66
                                        }
                                    }
                                    var v67 uint32 = 2147483648
                                    var v68 float32 = 1e10
                                    _ = v63
                                    _ = v67
                                    _ = v68
                                }
                                v4 = -32768
                            } else {
                                var v69 float32 = -2.25
                                var v70 bool = ((2 == 2) || false)
                                _ = v69
                                _ = v70
                            }
                            v5 = 0
                            _ = len(v2)
                        }
                        _ = fi39
                    }
                    var v71 uint64 = v5
                    _ = v71
                }
                _ = fi35
            }
        case "\n\t":
            v72 := GenS9{F0: -128, F1: 2147483647, F2: 1}
            _ = v72
        default:
            v4 = v4
        }
    case 2147483647:
        if false {
            var v73 int8 = 29
            for fi74 := 0; fi74 < 20; fi74++ {
                if (447 < 42) {
                    _ = len(v2)
                } else {
                    if true {
                        if true {
                            var v75 int64 = -643
                            _ = v75
                        } else {
                            v1 -= fi74
                        }
                        v76 := GenS5{F0: 252, F1: 42, F2: 139}
                        _ = v76
                    }
                    v73 = -1
                    for fi77 := 0; fi77 < 12; fi77++ {
                        v73 -= 100
                        var v78 uint16 = 255
                        _ = fi77
                        _ = v78
                    }
                }
                _ = fi74
            }
            switch ("\n\t" + "\n\t") {
            case "a":
                if true {
                    for fi79 := 0; fi79 < 3; fi79++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v4 -= v4
                        v5 = 0
                        _ = fi79
                    }
                    if true {
                        var v80 uint = 4294967295
                        _ = v80
                    }
                    if (0 == -2147483648) {
                        var v81 int8 = 113
                        v81 += -128
                        _ = v81
                    } else {
                        fmt.Println(v5, v4)
                        if (true && false) {
                            for fi82 := 0; fi82 < 5; fi82++ {
                                var v83 string = (("" + "\n\t") + "a")
                                _ = fi82
                                _ = v83
                            }
                            var v84 string = ("" + "hello")
                            v85 := GenS4{F0: 4294967295, F1: 114}
                            _ = v84
                            _ = v85
                        } else {
                            v86 := GenS5{F0: 42, F1: 2, F2: 2147483648}
                            for fi87 := 0; fi87 < 3; fi87++ {
                                if ((1 == 0) && false) {
                                    var v88 bool = false
                                    _ = v88
                                }
                                switch v1 {
                                case -157:
                                    var v89 int = fi87
                                    _ = v89
                                }
                                switch "foo bar" {
                                case "hello":
                                    var v90 uint16 = 0
                                    _ = v90
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if false {
                                            v91 := GenS8{F0: 255, F1: 2}
                                            _ = v91
                                        } else {
                                            fmt.Println(v1, v3)
                                            if true {
                                                var v92 float64 = 3.14159
                                                for fi93 := 0; fi93 < 13; fi93++ {
                                                    var v94 int16 = (v4 ^ -32768)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = fi93
                                                    _ = v94
                                                }
                                                for fi95 := 0; fi95 < 17; fi95++ {
                                                    if true {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    } else {
                                                        var v96 uint = 255
                                                        var v97 float64 = ((v92 + ((0.0 - 3.14159) / v3)) * 0.1)
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    _ = fi95
                                                }
                                                _ = v92
                                            }
                                        }
                                    }
                                }
                                _ = fi87
                            }
                            v4 = 42
                            _ = v86
                        }
                    }
                }
            case "hello":
                v98 := GenS7{F0: 0, F1: 2, F2: 0}
                _ = v98
            case "":
                var v99 int64 = 9223372036854775807
                _ = v99
            default:
                v5 = v5
            }
            _ = v73
        }
    default:
        var v100 uint64 = (v5 << ((uint(v5)) & 63))
        _ = v100
    }
    v101 := GenS10{F0: 42, F1: 2, F2: 2147483648}
    v102 := GenS7{F0: 0, F1: 864, F2: 52}
    fmt.Println(v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v101
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
