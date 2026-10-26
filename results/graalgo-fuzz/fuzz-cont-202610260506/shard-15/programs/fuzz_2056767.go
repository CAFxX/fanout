package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 int64; F1 int32; F2 uint8 }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 string; F1 int }
func genFn0(p0 int, p1 int32, p2 uint32) int8 {
    if false {
        p1 = p1
    }
    switch p0 {
    case -847:
        p2 = (^(p2))
    case 32768:
        v5 := GenS1{F0: 372, F1: 42, F2: -211}
        _ = v5
    }
    fmt.Println(p0)
    var v6 uint8 = 208
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 65535
    v3 := []string{("a" + "a"), "\n\t", "", "hello"}
    v4 := make(map[int]string)
    v4[-1] = "\n\t"
    if (-1 > 9223372036854775807) {
        _ = len(v4)
        v7 := GenS1{F0: 488, F1: -1, F2: 1}
        _ = v7
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v8 := GenS1{F0: 560, F1: genFn0(32767, -1, 2147483647), F2: 561}
        _ = v8
    }
    delete(v4, -2147483648)
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 32767, F1: 100, F2: -137}
        v1 += (v1 % ((v1) | 1))
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int64 = -1
            switch fi9 {
            case -2147483648:
                var v13 string = (("hello" + "\n\t") + "hello")
                _ = v13
            case 1:
                v4[1] = "foo bar"
                if true {
                    v3 = append(v3, ((("" + "") + "") + "hello"))
                    var v14 uint = v1
                    _ = v14
                } else {
                    _ = len(v3)
                    v1 += 255
                }
            case 499:
                switch "\n\t" {
                case "foo bar":
                    if (!false) {
                        if true {
                            v1 = (v1 % ((v1) | 1))
                        } else {
                            v1 = 726
                        }
                    }
                    if ((-96 >= genFn0(fi9, -708, (v2 + (v2 * (v2 - v2))))) && true) {
                        var v15 string = ("\n\t" + "a")
                        switch 42 {
                        case 32767:
                            if ((!false) || ((-9223372036854775808 != int64(fi9)) || true)) {
                                v4[fi11] = v15
                                v16 := GenS1{F0: 32767, F1: 1, F2: 42}
                                v2 -= v2
                                v17 := GenS2{F0: 65535, F1: -1}
                                var v18 bool = false
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        case 32768:
                            _ = len(v4)
                        default:
                            if (0 != 2) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v2 += v2
                                v15 = ((v15 + ((v15 + "hello") + (v15 + ("\n\t" + (v15 + (v15 + "")))))) + ("" + (v15 + "\n\t")))
                            }
                            var v19 uint64 = 0
                            v1 -= 4294967295
                            var v20 int32 = 42
                            if (((v2 << ((uint(v2)) & 63)) > (v2 >> ((v1) & 63))) && (fi9 != 1)) {
                                var v21 uint32 = v2
                                _ = v21
                            }
                            _ = v19
                            _ = v20
                        }
                        delete(v4, (^(fi11)))
                        v22 := GenS3{F0: 127}
                        var v23 float32 = float32(0.0)
                        _ = v15
                        _ = v22
                        _ = v23
                    } else {
                        _ = len(v3)
                        if (-1 <= -1) {
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v4)
                            case "hello":
                                if false {
                                    v1 -= 65535
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v24 uint64 = 2
                                    v24 = v24
                                    _ = v24
                                }
                            case "a":
                                v4[fi9] = "foo bar"
                            }
                            switch -9223372036854775808 {
                            case 32767:
                                var v25 bool = (("\n\t" <= "foo bar") && ((0 > uint64(v1)) || false))
                                _ = v25
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        } else {
                            var v26 int64 = (v12 + int64(v2))
                            if (false && true) {
                                v26 = (^(v12))
                                v27 := GenS3{F0: 127}
                                var v28 int32 = 931
                                _ = v27
                                _ = v28
                            } else {
                                fmt.Println(v1)
                                delete(v4, -2147483648)
                                var v29 int16 = -32768
                                _ = v29
                            }
                            _ = v26
                        }
                        v30 := GenS1{F0: 2, F1: genFn0(fi9, int32(v2), (^(v2))), F2: 0}
                        switch ("foo bar" + ("a" + (("a" + "") + ("\n\t" + "foo bar")))) {
                        case "\n\t":
                            fmt.Println(v1, fi11)
                        case "hello":
                            v2 -= (^(v2))
                        case "":
                            fmt.Println(v1)
                        default:
                            var v31 int32 = 920
                            _ = v31
                        }
                        if false {
                            var v32 bool = true
                            _ = v32
                        } else {
                            switch 42 {
                            case -137:
                                v2 += 539
                            }
                            var v33 int16 = -32768
                            switch "hello" {
                            case "foo bar":
                                var v34 float32 = 1.5
                                _ = v34
                            case "hello":
                                v3 = append(v3, (("\n\t" + "foo bar") + (("" + ("" + "foo bar")) + "hello")))
                            default:
                                switch (("hello" + "a") + ("a" + (("foo bar" + ((("\n\t" + "") + "\n\t") + ((("a" + "a") + "foo bar") + "a"))) + "a"))) {
                                case "foo bar":
                                    switch -849 {
                                    case 42:
                                        v2 -= v2
                                    }
                                case "\n\t":
                                    if (42 != 65535) {
                                        if true {
                                            switch (fi11 << ((730) & 63)) {
                                            case -1:
                                                v35 := v4[(fi11 | (fi9 ^ (fi9 / ((fi11) | 1))))]
                                                _ = v35
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v33
                        }
                        switch (^(fi9)) {
                        case 32768:
                            var v36 float64 = 1.5
                            _ = v36
                        case 530:
                            v2 -= v2
                        case 4294967295:
                            if true {
                                switch fi9 {
                                case -1:
                                    v37 := GenS3{F0: 218}
                                    v12 = (v12 << (((v1 + (^(v1)))) & 63))
                                    _ = v37
                                case 735:
                                    var v38 float32 = float32((3.14159 / float64(-2.25)))
                                    _ = v38
                                case 0:
                                    var v39 uint = (v1 << ((291) & 63))
                                    _ = v39
                                default:
                                    if (941 >= 4294967295) {
                                        var v40 int8 = 113
                                        var v41 string = ("hello" + "")
                                        _ = v40
                                        _ = v41
                                    }
                                    v3 = append(v3, (("\n\t" + "\n\t") + ""))
                                }
                            } else {
                                var v42 uint32 = v2
                                _ = v42
                            }
                        default:
                            var v43 uint64 = 18446744073709551615
                            _ = v43
                        }
                        _ = v30
                    }
                }
            default:
                switch (fi11 & -928) {
                case -727:
                    v44 := GenS1{F0: 349, F1: 10, F2: 42}
                    _ = v44
                case 763:
                    switch fi11 {
                    case 307:
                        v45 := GenS3{F0: 2}
                        v2 += 1
                        _ = v45
                    case -1:
                        v46 := GenS1{F0: 255, F1: 1, F2: -381}
                        _ = v46
                    default:
                        switch -1 {
                        case 52:
                            switch "\n\t" {
                            case "hello":
                                _ = len(v4)
                            case "\n\t":
                                v1 += 2
                            case "foo bar":
                                v47 := GenS4{F0: 127, F1: ("hello" + "")}
                                _ = v47
                            }
                        case -1:
                            v48 := GenS5{F0: -1, F1: ((!(!((0.1 < 1.5) && true))) || true)}
                            fmt.Println(fi11, v1)
                            if false {
                                fmt.Println(fi9, v12)
                            } else {
                                switch "a" {
                                case "a":
                                    if (true && ((v1 / (((^(v1))) | 1)) > v1)) {
                                        if (!true) {
                                            v49 := GenS6{F0: (v12 / ((v12) | 1))}
                                            _ = v49
                                        }
                                        fmt.Println(fi9)
                                        var v50 int64 = v12
                                        var v51 uint64 = 18446744073709551615
                                        _ = v50
                                        _ = v51
                                    } else {
                                        var v52 int32 = -1
                                        v53 := GenS1{F0: 32767, F1: 65, F2: 32767}
                                        _ = v52
                                        _ = v53
                                    }
                                case "":
                                    if (true || true) {
                                        switch (fi11 ^ fi9) {
                                        case 1:
                                            var v54 int = -877
                                            _ = v54
                                        default:
                                            if false {
                                                var v55 uint32 = v2
                                                var v56 uint32 = 2147483648
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                    } else {
                                        if true {
                                            v57 := GenS4{F0: 100, F1: (("\n\t" + "a") + "foo bar")}
                                            fmt.Println(v12, v2)
                                            _ = v57
                                        } else {
                                            v58 := GenS7{F0: (v12 - v12), F1: 0}
                                            switch (fi11 % ((703) | 1)) {
                                            case 325:
                                                if false {
                                                    if true {
                                                        v59 := GenS3{F0: 0}
                                                        var v60 uint64 = 196
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        var v61 uint64 = 545
                                                        if ((v1 & (v1 | 501)) > (v1 & 538)) {
                                                            fmt.Println(v1, v2, fi11)
                                                            v12 += v12
                                                        }
                                                        var v62 int64 = 471
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                } else {
                                                    var v63 uint8 = 127
                                                    switch (fi11 / ((-290) | 1)) {
                                                    case -141:
                                                        var v64 uint8 = v63
                                                        _ = v64
                                                    }
                                                    _ = v63
                                                }
                                            case -2147483648:
                                                if (42 == -60) {
                                                    v1 -= v1
                                                }
                                            case 0:
                                                switch ("foo bar" + "a") {
                                                case "":
                                                    var v65 int = fi9
                                                    _ = v65
                                                }
                                            }
                                            if (!(!(false && false))) {
                                                if true {
                                                    if true {
                                                        if false {
                                                            switch (fi9 / (((fi11 % ((fi11) | 1))) | 1)) {
                                                            case -439:
                                                                v66 := GenS2{F0: 4294967295, F1: 42}
                                                                _ = v66
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v58
                                        }
                                        v67 := GenS8{F0: (v1 % ((v1) | 1)), F1: 485}
                                        v1 = v1
                                        _ = v67
                                    }
                                case "\n\t":
                                    v68 := GenS8{F0: 18446744073709551615, F1: 255}
                                    _ = v68
                                }
                            }
                            _ = v48
                        case -9223372036854775808:
                            v69 := GenS6{F0: (v12 + -590)}
                            _ = v69
                        default:
                            if ((-2.25 - 0.0) != (-2.25 + 3.14159)) {
                                v12 = (v12 & int64(v1))
                                var v70 float32 = 0.0
                                _ = v70
                            }
                        }
                    }
                case 127:
                    switch (^(fi11)) {
                    case 432:
                        fmt.Println(v12, fi11, fi11)
                    case -9223372036854775808:
                        if true {
                            if true {
                                v2 += 65535
                            } else {
                                var v71 bool = (true || false)
                                v72 := GenS8{F0: (v1 / ((v1) | 1)), F1: 255}
                                _ = v71
                                _ = v72
                            }
                        } else {
                            switch fi11 {
                            case 516:
                                fmt.Println(fi11)
                            default:
                                v1 = (v1 + 42)
                                var v73 int = 614
                                _ = v73
                            }
                        }
                    case 32767:
                        v12 -= v12
                    default:
                        v12 -= (-(v12))
                    }
                default:
                    fmt.Println(fi9)
                }
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    switch 32768 {
    case 42:
        v2 = v2
    case 0:
        switch 1 {
        case 32768:
            if (1 > 1) {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                var v74 uint = 65535
                _ = v74
            } else {
                if (uint8(v2) <= 255) {
                    v1 = v1
                }
            }
        default:
            v75 := GenS5{F0: 32767, F1: true}
            _ = v75
        }
    default:
        if (0 >= 1) {
            var v76 uint8 = uint8(v1)
            v2 = (v2 ^ (v2 - v2))
            if (32768 <= 283) {
                var v77 float32 = (-2.25 + 1e10)
                var v78 uint32 = v2
                _ = v77
                _ = v78
            }
            _ = v76
        }
    }
    v1 = v1
    v79 := GenS3{F0: 42}
    v3 = append(v3, "")
    switch "a" {
    case "hello":
        if true {
            var v80 int = 942
            _ = v80
        }
    default:
        fmt.Println(v2, v1)
    }
    if (true || (0.1 < (1e10 * float64(3.14159)))) {
        var v81 string = "foo bar"
        var v82 uint64 = 4294967295
        v82 += 4294967295
        _ = v81
        _ = v82
    } else {
        for fi83 := 0; fi83 < 20; fi83++ {
            switch fi83 {
            case 4294967295:
                v1 = v1
            }
            var v84 uint8 = uint8(v1)
            _ = fi83
            _ = v84
        }
        v2 += (v2 + v2)
        for fi85 := 0; fi85 < 11; fi85++ {
            v2 = (v2 ^ v2)
            if false {
                if false {
                    v86 := GenS9{F0: 9223372036854775807, F1: -1, F2: 42}
                    _ = v86
                } else {
                    var v87 float64 = 1.5
                    for fi88 := 0; fi88 < 8; fi88++ {
                        if false {
                            v87 = v87
                            var v89 uint64 = uint64(v2)
                            v87 -= 0.1
                            _ = v89
                        }
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + "")
                        }
                        _ = fi88
                    }
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    _ = v87
                }
                var v90 int16 = 0
                if (62 >= 127) {
                    var v91 int8 = 124
                    _ = v91
                }
                _ = v90
            } else {
                var v92 float32 = 1.5
                var v93 uint16 = 32767
                if true {
                    v2 += 2147483648
                    if (float64((1.5 + v92)) != 1.5) {
                        v94 := GenS10{F0: (fi85 - (fi85 + (fi85 % ((-9223372036854775808) | 1)))), F1: 0}
                        _ = v94
                    }
                }
                _ = v92
                _ = v93
            }
            _ = fi85
        }
    }
    v2 = 42
    switch 32767 {
    case 1:
        v2 = v2
    case 531:
        v1 = v1
    case 163:
        v1 = 42
    default:
        var v95 bool = false
        _ = v95
    }
    switch 128 {
    case 9223372036854775807:
        var v96 string = "foo bar"
        _ = v96
    case 4294967295:
        var v97 float32 = 3.14159
        _ = v97
    case 123:
        delete(v4, 894)
    }
    if false {
        if (-36 < 0) {
            var v98 float64 = 0.1
            v99 := GenS4{F0: 42, F1: (("a" + "\n\t") + "a")}
            v98 = -2.25
            _ = v98
            _ = v99
        } else {
            var v100 bool = ((!true) && (127 == 42))
            _ = v100
        }
        if true {
            if (!false) {
                var v101 uint32 = (v2 - v2)
                _ = v101
            } else {
                v2 = v2
                fmt.Println(v2, v2)
                var v102 float64 = (1.5 - (0.0 * ((3.14159 + 1e10) - ((((float64((float32(-2.25) / 1.0)) - 0.1) * 0.1) + (-2.25 * 1.5)) - (1e10 + 1.5)))))
                _ = v102
            }
            v3 = append(v3, "foo bar")
        }
        if (true && false) {
            fmt.Println(v2)
            switch -405 {
            case 9223372036854775807:
                for fi103 := 0; fi103 < 6; fi103++ {
                    var v104 uint64 = 626
                    _ = fi103
                    _ = v104
                }
            case 803:
                if (0 >= 42) {
                    v1 += v1
                    for fi105 := 0; fi105 < 20; fi105++ {
                        v4[(-(fi105))] = ""
                        _ = fi105
                    }
                } else {
                    v2 = (^(v2))
                    var v106 int16 = 1000
                    var v107 int8 = -128
                    _ = v106
                    _ = v107
                }
            case 923:
                v2 = v2
            }
        } else {
            if (((("foo bar" + "foo bar") + ("" + "hello")) + "foo bar") != "foo bar") {
                v2 = v2
            } else {
                v2 = v2
                v108 := GenS11{F0: "", F1: 1}
                v1 = v1
                _ = v108
            }
        }
    } else {
        v109 := GenS10{F0: -9223372036854775808, F1: 669}
        var v110 int16 = -1
        _ = v109
        _ = v110
    }
    v111 := GenS3{F0: 2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v79
}
