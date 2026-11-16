package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 string }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint32; F2 bool }
type GenS7 struct { F0 uint; F1 bool; F2 uint16 }
type GenS8 struct { F0 int8; F1 bool; F2 int16 }
type GenS9 struct { F0 uint32 }

func main() {
    var v1 bool = true
    var v2 uint8 = 42
    v3 := []int{553}
    var v4 int64 = -1
    var v5 float32 = ((1.5 / 1e10) + 1.5)
    var v6 uint64 = uint64(v2)
    if len(v3) > 0 {
        v3[0] = 127
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            switch fi7 {
            case 9223372036854775807:
                fmt.Println(v5, v6, v2)
            case 42:
                if (1 > int8(v4)) {
                    var v9 uint64 = v6
                    _ = v9
                }
                var v10 uint16 = 65535
                if (-1 >= -1) {
                    if true {
                        v6 += 18446744073709551615
                        fmt.Println(v5, v5, v6)
                    }
                    v11 := GenS1{F0: -1}
                    var v12 int = (fi8 - fi7)
                    var v13 int = (fi7 - v12)
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = v10
            case -177:
                switch ("hello" + "a") {
                case "hello":
                    var v14 int16 = 735
                    _ = v14
                }
            default:
                fmt.Println(fi8)
            }
            v1 = ((float64(1.5) == (-2.25 + 0.1)) || (!v1))
            _ = fi8
        }
        if v1 {
            v15 := GenS2{F0: 126}
            var v16 float32 = 1.5
            if (145 == 928) {
                if ("a" > "hello") {
                    var v17 uint = 4294967295
                    var v18 float32 = ((0.0 / v5) / float32(1e10))
                    _ = v17
                    _ = v18
                }
                switch "hello" {
                case "\n\t":
                    var v19 int = (fi7 & 127)
                    _ = v19
                case "a":
                    if v1 {
                        v3 = append(v3, (fi7 & fi7))
                        v16 -= (v5 * -2.25)
                    } else {
                        v20 := GenS2{F0: 1}
                        var v21 int32 = 802
                        _ = v20
                        _ = v21
                    }
                case "":
                    for fi22 := 0; fi22 < 2; fi22++ {
                        if false {
                            _ = len(v3)
                            v6 = (v6 - 4294967295)
                            v23 := GenS1{F0: 1000}
                            var v24 float32 = (((v16 + 0.0) + (3.14159 + 0.0)) / v5)
                            v6 = (v6 % ((v6) | 1))
                            var v25 float64 = (3.14159 * float64(v16))
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        if ((-(v4)) == 936) {
                            v26 := GenS1{F0: -32768}
                            switch 0 {
                            case 2147483647:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                if len(v3) > 0 {
                                    v3[0] = fi22
                                }
                            }
                            _ = v26
                        }
                        _ = fi22
                    }
                default:
                    if false {
                        var v27 uint16 = 255
                        v2 = 2
                        _ = v27
                    } else {
                        v28 := GenS2{F0: -1}
                        _ = v28
                    }
                }
                var v29 uint32 = 212
                _ = v29
            }
            _ = v15
            _ = v16
        }
        var v30 uint8 = 128
        _ = fi7
        _ = v30
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        if (2147483648 <= 2147483647) {
            var v32 uint32 = 93
            _ = v32
        }
        _ = fi31
    }
    v4 = (^(v4))
    if true {
        v33 := GenS3{F0: -32768, F1: 973, F2: ("foo bar" + "")}
        _ = v33
    }
    switch -2147483648 {
    case 153:
        var v34 bool = v1
        _ = v34
    case -441:
        for fi35 := 0; fi35 < 2; fi35++ {
            v36 := GenS2{F0: 127}
            if (!v1) {
                v4 += v4
            }
            _ = fi35
            _ = v36
        }
    case -378:
        var v37 int8 = -1
        _ = v37
    default:
        v1 = ("\n\t" <= (("" + "") + ""))
    }
    if (v1 && v1) {
        v4 = v4
        v5 += 0.1
        v38 := GenS3{F0: 32767, F1: -1, F2: "a"}
        _ = v38
    } else {
        v4 -= (v4 / ((-9223372036854775808) | 1))
        switch 42 {
        case -1:
            for fi39 := 0; fi39 < 2; fi39++ {
                if v1 {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi40 := 0; fi40 < 9; fi40++ {
                        var v41 uint32 = 65535
                        if len(v3) > 0 {
                            v3[0] = fi40
                        }
                        _ = fi40
                        _ = v41
                    }
                    if (v1 || v1) {
                        for fi42 := 0; fi42 < 5; fi42++ {
                            v5 = v5
                            v6 += 2
                            v5 += -2.25
                            _ = fi42
                        }
                        v6 = (v6 >> ((4294967295) & 63))
                        var v43 int32 = 1
                        _ = v43
                    } else {
                        v44 := GenS1{F0: -214}
                        _ = v44
                    }
                } else {
                    if v1 {
                        for fi45 := 0; fi45 < 10; fi45++ {
                            v6 = (v6 - (v6 % (((v6 << ((460) & 63))) | 1)))
                            if ((v6 & uint64(v2)) != 2) {
                                v1 = (252 != -364)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v4 += int64(v2)
                                fmt.Println(v1, v2)
                            } else {
                                var v46 float32 = ((v5 + v5) - v5)
                                v3 = append(v3, int(v6))
                                _ = v46
                            }
                            if (!v1) {
                                var v47 int = fi39
                                fmt.Println(v6)
                                _ = v47
                            }
                            switch ("a" + "a") {
                            case "":
                                fmt.Println(v6)
                            case "foo bar":
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v6 -= v6
                                    case "hello":
                                        v48 := GenS3{F0: 0, F1: -32768, F2: (("" + "a") + "a")}
                                        _ = v48
                                    }
                                    v5 += v5
                                    _ = len(v3)
                                }
                                var v49 int16 = 1
                                v6 = v6
                                v4 -= v4
                                _ = v49
                            default:
                                v5 -= v5
                            }
                            v4 = v4
                            switch "" {
                            case "hello":
                                var v50 uint8 = (^(v2))
                                _ = v50
                            default:
                                var v51 int = 655
                                _ = v51
                            }
                            _ = fi45
                        }
                        fmt.Println(v2, v4)
                        switch -195 {
                        case 32767:
                            for fi52 := 0; fi52 < 5; fi52++ {
                                if v1 {
                                    fmt.Println(fi39, fi52)
                                    fmt.Println(fi39)
                                    var v53 int32 = 42
                                    _ = v53
                                } else {
                                    fmt.Println(v5, fi52)
                                }
                                fmt.Println(fi39, v2, v2)
                                if (v1 && v1) {
                                    v5 += v5
                                }
                                _ = fi52
                            }
                        case 4294967295:
                            if v1 {
                                v2 += v2
                            } else {
                                if true {
                                    fmt.Println(fi39, fi39, v4)
                                    v4 -= int64(fi39)
                                    v2 -= (v2 / ((v2) | 1))
                                } else {
                                    if true {
                                        switch fi39 {
                                        case 1:
                                            for fi54 := 0; fi54 < 5; fi54++ {
                                                _ = len(v3)
                                                if (1 < 714) {
                                                    v55 := GenS2{F0: 1}
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if true {
                                                            v1 = ((v2 / (((v2 % ((v2) | 1))) | 1)) > v2)
                                                            v56 := GenS2{F0: -1}
                                                            var v57 bool = (455 <= 1)
                                                            _ = v56
                                                            _ = v57
                                                        } else {
                                                            switch "a" {
                                                            case "":
                                                                var v58 int32 = -2147483648
                                                                _ = v58
                                                            case "\n\t":
                                                                if len(v3) > 0 {
                                                                    v3[0] = 994
                                                                }
                                                            default:
                                                                var v59 int64 = 42
                                                                _ = v59
                                                            }
                                                        }
                                                    case "a":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        fmt.Println(fi54, v5, v5)
                                                        switch ("foo bar" + ("" + "hello")) {
                                                        case "\n\t":
                                                            v6 -= 354
                                                        case "foo bar":
                                                            v4 = 42
                                                        default:
                                                            switch fi54 {
                                                            case -364:
                                                                var v60 int32 = 332
                                                                _ = v60
                                                            case 0:
                                                                v4 -= -9223372036854775808
                                                            case 42:
                                                                v2 = (v2 * (v2 & 2))
                                                            }
                                                            if v1 {
                                                                if (!v1) {
                                                                    var v61 int16 = 1
                                                                    v3 = append(v3, (^(fi54)))
                                                                    _ = v61
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = 0
                                                                }
                                                                v62 := GenS4{F0: 1, F1: 127}
                                                                _ = v62
                                                            } else {
                                                                v2 += 128
                                                                if (false && (468 > 743)) {
                                                                    v6 += 4294967295
                                                                    if len(v3) > 0 {
                                                                        v3[0] = fi39
                                                                    }
                                                                }
                                                                var v63 float32 = float32((1e10 * 3.14159))
                                                                var v64 int = fi39
                                                                v65 := GenS5{F0: 65535}
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            if ((v5 == (v5 / v5)) || v1) {
                                                                var v66 int16 = -744
                                                                v6 += v6
                                                                fmt.Println(v5, v2)
                                                                _ = v66
                                                            } else {
                                                                switch (fi39 | 9223372036854775807) {
                                                                case 42:
                                                                    v3 = append(v3, fi39)
                                                                case 1:
                                                                    v67 := GenS5{F0: 2}
                                                                    _ = v67
                                                                default:
                                                                    var v68 uint8 = 0
                                                                    _ = v68
                                                                }
                                                                var v69 uint64 = v6
                                                                _ = v69
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v4 = 42
                                                            case "hello":
                                                                fmt.Println(v1, fi39)
                                                            }
                                                        }
                                                    default:
                                                        v70 := GenS6{F0: 0, F1: 2147483648, F2: v1}
                                                        _ = v70
                                                    }
                                                    _ = v55
                                                }
                                                v71 := GenS6{F0: 0, F1: 824, F2: ((-2147483648 < 287) || false)}
                                                v3 = append(v3, (fi54 ^ 799))
                                                if (-32768 >= int16(v6)) {
                                                    var v72 uint = 99
                                                    switch (("hello" + ("foo bar" + "a")) + (("foo bar" + ("hello" + ("\n\t" + "hello"))) + "hello")) {
                                                    case "hello":
                                                        v1 = false
                                                    case "\n\t":
                                                        v3 = append(v3, fi54)
                                                    default:
                                                        var v73 string = "hello"
                                                        _ = v73
                                                    }
                                                    _ = v72
                                                } else {
                                                    v1 = (v1 && (-1 != fi54))
                                                    if false {
                                                        v1 = v1
                                                    } else {
                                                        switch (("a" + "\n\t") + "a") {
                                                        case "\n\t":
                                                            v5 += v5
                                                        case "foo bar":
                                                            v6 -= 241
                                                        case "hello":
                                                            v74 := GenS5{F0: 691}
                                                            _ = v74
                                                        }
                                                        v6 -= v6
                                                        switch fi54 {
                                                        case 0:
                                                            v75 := GenS5{F0: 255}
                                                            _ = v75
                                                        default:
                                                            var v76 bool = (v2 <= (v2 % ((v2) | 1)))
                                                            _ = v76
                                                        }
                                                        v2 = v2
                                                        var v77 uint = 830
                                                        _ = v77
                                                    }
                                                }
                                                _ = fi54
                                                _ = v71
                                            }
                                        case -1:
                                            var v78 int8 = 100
                                            _ = v78
                                        default:
                                            v5 -= v5
                                        }
                                        v79 := GenS7{F0: 76, F1: (!false), F2: 103}
                                        _ = v79
                                    }
                                }
                            }
                        default:
                            if v1 {
                                for fi80 := 0; fi80 < 18; fi80++ {
                                    if (-2.25 < 1e10) {
                                        switch 747 {
                                        case -843:
                                            fmt.Println(v6, v1, v2)
                                        case -587:
                                            var v81 int64 = 544
                                            _ = v81
                                        case 32768:
                                            if (float32((1e10 + 1.5)) >= 1e10) {
                                                if true {
                                                    if ((v6 | v6) <= v6) {
                                                        var v82 int16 = 42
                                                        _ = v82
                                                    } else {
                                                        v83 := GenS7{F0: 4294967295, F1: v1, F2: 0}
                                                        var v84 float32 = 3.14159
                                                        var v85 float32 = (((3.14159 - v5) * ((v84 / 3.14159) * float32(1.5))) + (-2.25 - ((v84 + v5) + v84)))
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    v86 := GenS2{F0: 42}
                                                    if false {
                                                        var v87 int = (-(fi39))
                                                        v88 := GenS7{F0: 4294967295, F1: (1.5 < (3.14159 * 0.1)), F2: 42}
                                                        var v89 uint = 4294967295
                                                        _ = v87
                                                        _ = v88
                                                        _ = v89
                                                    } else {
                                                        v90 := GenS8{F0: int8(v4), F1: (v1 && (-32768 <= 1000)), F2: 32767}
                                                        _ = v90
                                                    }
                                                    _ = v86
                                                }
                                                v6 += 275
                                            }
                                        }
                                        v4 += v4
                                        if true {
                                            if v1 {
                                                if (v2 >= 110) {
                                                    fmt.Println(v1, fi39, v6)
                                                }
                                                var v91 int8 = 100
                                                v91 += -68
                                                _ = v91
                                            }
                                            v1 = (65535 < 0)
                                        } else {
                                            var v92 int = (fi39 + -9223372036854775808)
                                            _ = v92
                                        }
                                    }
                                    _ = fi80
                                }
                            } else {
                                v3 = append(v3, 42)
                                fmt.Println(v6, fi39, v4)
                            }
                        }
                    }
                }
                _ = fi39
            }
        }
    }
    if false {
        if len(v3) > 0 {
            v3[0] = 9223372036854775807
        }
        for fi93 := 0; fi93 < 6; fi93++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v94 := GenS1{F0: -486}
            v2 = (v2 & v2)
            _ = fi93
            _ = v94
        }
        for fi95 := 0; fi95 < 7; fi95++ {
            v96 := GenS9{F0: 471}
            _ = fi95
            _ = v96
        }
    } else {
        for fi97 := 0; fi97 < 7; fi97++ {
            var v98 uint32 = 1
            _ = fi97
            _ = v98
        }
        var v99 int8 = -36
        _ = v99
    }
    for fi100 := 0; fi100 < 16; fi100++ {
        var v101 uint32 = 2147483648
        _ = fi100
        _ = v101
    }
    if (42 < 2147483648) {
        switch 32767 {
        case -2147483648:
            var v102 float64 = 0.1
            _ = v102
        default:
            if (-128 <= -128) {
                v4 = 0
            }
        }
    }
    v3 = append(v3, 0)
    var v103 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
