package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 uint; F2 int16 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 string; F1 uint16; F2 uint }
type GenS10 struct { F0 int16 }
func genFn0() uint64 {
    v7 := []int{127, 1, 834, 104}
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 387}
        switch (fi8 & fi8) {
        case 9223372036854775807:
            var v10 int16 = -1
            _ = v10
        case 438:
            var v11 uint32 = 1
            _ = v11
        }
        var v12 int = fi8
        _ = fi8
        _ = v12
        _ = v9
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        switch "hello" {
        case "foo bar":
            var v14 int32 = 122
            _ = v14
        default:
            if false {
                var v15 int64 = 9223372036854775807
                _ = v15
            } else {
                v16 := GenS1{F0: 916}
                if (42 > -128) {
                    var v17 uint16 = 0
                    v17 = 976
                    _ = v17
                }
                if (897 == fi13) {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        var v19 int8 = -128
                        _ = fi18
                        _ = v19
                    }
                }
                _ = v16
            }
        }
        var v20 string = "foo bar"
        _ = fi13
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 int8 = 26
        v23 := GenS1{F0: 0}
        _ = fi21
        _ = v22
        _ = v23
    }
    if (true || false) {
        if true {
            v7 = append(v7, -1)
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            switch (("a" + "foo bar") + ((("hello" + "a") + "a") + "hello")) {
            case "hello":
                var v25 uint64 = 4294967295
                _ = v25
            case "foo bar":
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 uint64 = 0
                    fmt.Println(v27, fi24, fi26)
                    switch fi24 {
                    case -630:
                        if true {
                            v28 := GenS1{F0: 788}
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v7 = append(v7, 752)
                            _ = v28
                        }
                    case 263:
                        fmt.Println(fi24, fi24)
                        if false {
                            if len(v7) > 0 {
                                v7[0] = (fi24 * 1)
                            }
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        } else {
                            var v29 string = (("" + "") + ("" + "foo bar"))
                            if false {
                                if false {
                                    if false {
                                        v30 := GenS1{F0: 42}
                                        fmt.Println(v29, fi24, fi26)
                                        var v31 bool = (42 < 1)
                                        _ = v30
                                        _ = v31
                                    }
                                    if true {
                                        v27 -= v27
                                    } else {
                                        var v32 uint64 = v27
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        _ = v32
                                    }
                                    if len(v7) > 0 {
                                        v7[0] = (fi26 & fi26)
                                    }
                                    fmt.Println(fi24, v27, v27)
                                } else {
                                    v29 = ((("" + "hello") + (v29 + (v29 + v29))) + ((v29 + "foo bar") + v29))
                                }
                                var v33 int16 = 1
                                if false {
                                    v34 := GenS2{F0: 42, F1: (-(fi24)), F2: 83}
                                    v7 = append(v7, fi24)
                                    if (0.0 < 1e10) {
                                        fmt.Println(v27, fi24)
                                        switch (fi24 - fi26) {
                                        case 1:
                                            switch fi24 {
                                            case 42:
                                                if false {
                                                    fmt.Println(fi24, v33, v29)
                                                }
                                            case -493:
                                                if (true && (!false)) {
                                                    var v35 uint32 = 274
                                                    if false {
                                                        v36 := GenS1{F0: -926}
                                                        fmt.Println(v27, v35)
                                                        _ = v36
                                                    } else {
                                                        var v37 uint64 = v27
                                                        _ = v37
                                                    }
                                                    _ = v35
                                                } else {
                                                    switch "foo bar" {
                                                    case "a":
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi26 >> ((255) & 63))
                                                        }
                                                    case "\n\t":
                                                        if len(v7) > 0 {
                                                            v7[0] = fi26
                                                        }
                                                    case "hello":
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi24 - fi26)
                                                        }
                                                    default:
                                                        var v38 uint = 65535
                                                        _ = v38
                                                    }
                                                    v39 := GenS2{F0: 2, F1: (^(fi26)), F2: 0}
                                                    fmt.Println(v29, v27, v29)
                                                    _ = v39
                                                }
                                            }
                                        case -2147483648:
                                            v40 := GenS1{F0: 1}
                                            _ = v40
                                        }
                                        v41 := GenS3{F0: 2, F1: (fi26 * fi26), F2: 2}
                                        _ = v41
                                    } else {
                                        var v42 int16 = -32768
                                        _ = v42
                                    }
                                    _ = v34
                                } else {
                                    switch fi26 {
                                    case 2147483647:
                                        v33 -= (v33 | (v33 - v33))
                                    default:
                                        v43 := GenS4{F0: (!true)}
                                        _ = v43
                                    }
                                }
                                var v44 string = "a"
                                _ = v33
                                _ = v44
                            }
                            if true {
                                var v45 uint = 570
                                switch fi26 {
                                case 4294967295:
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                case 42:
                                    switch v29 {
                                    case "a":
                                        switch -1 {
                                        case -2147483648:
                                            var v46 uint = v45
                                            _ = v46
                                        case -480:
                                            var v47 int32 = 1
                                            _ = v47
                                        case 32767:
                                            v48 := GenS3{F0: (v27 - (v27 / ((v27) | 1))), F1: fi26, F2: 32767}
                                            _ = v48
                                        default:
                                            if len(v7) > 0 {
                                                v7[0] = fi26
                                            }
                                        }
                                    case "foo bar":
                                        v27 = (v27 - (v27 - 18446744073709551615))
                                    default:
                                        if true {
                                            v45 = 1
                                            var v49 string = "foo bar"
                                            if false {
                                                var v50 uint64 = (v27 ^ 42)
                                                _ = v50
                                            }
                                            if false {
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                v29 = v29
                                            }
                                            _ = v49
                                        } else {
                                            v27 = (v27 + v27)
                                            v29 = ""
                                            var v51 uint32 = 653
                                            _ = v51
                                        }
                                    }
                                default:
                                    v29 = v29
                                }
                                _ = len(v7)
                                var v52 uint64 = (v27 & v27)
                                fmt.Println(v45)
                                _ = v45
                                _ = v52
                            } else {
                                switch v29 {
                                case "hello":
                                    var v53 bool = (81 <= 128)
                                    _ = v53
                                }
                                switch (-(fi26)) {
                                case 9223372036854775807:
                                    var v54 uint = 0
                                    _ = v54
                                default:
                                    v29 = ((v29 + "") + v29)
                                }
                            }
                            v55 := GenS5{F0: v27}
                            v56 := GenS6{F0: 9223372036854775807, F1: 0, F2: 1}
                            _ = v29
                            _ = v55
                            _ = v56
                        }
                    case 128:
                        _ = len(v7)
                    default:
                        v57 := GenS7{F0: ("a" + ""), F1: 0}
                        _ = v57
                    }
                    _ = fi26
                    _ = v27
                }
            default:
                var v58 int64 = 0
                _ = v58
            }
            var v59 bool = (628 <= 1)
            _ = fi24
            _ = v59
        }
    } else {
        if true {
            if (574 < 982) {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    var v60 uint64 = 42
                    _ = v60
                default:
                    var v61 uint = 42
                    _ = v61
                }
            } else {
                var v62 uint16 = 393
                _ = len(v7)
                _ = v62
            }
            var v63 int32 = 303
            _ = v63
        }
        if (503 == 0) {
            for fi64 := 0; fi64 < 8; fi64++ {
                if true {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    if true {
                        var v65 int32 = 42
                        fmt.Println(fi64, fi64)
                        v66 := GenS8{F0: 1, F1: 186}
                        _ = v65
                        _ = v66
                    } else {
                        var v67 uint = 255
                        v67 = v67
                        _ = len(v7)
                        _ = v67
                    }
                }
                _ = fi64
            }
            v68 := GenS9{F0: "hello", F1: 32767, F2: 1}
            _ = v68
        } else {
            v69 := GenS5{F0: 4294967295}
            if (true && false) {
                for fi70 := 0; fi70 < 9; fi70++ {
                    var v71 int = fi70
                    fmt.Println(fi70)
                    _ = fi70
                    _ = v71
                }
                var v72 int64 = 1
                _ = v72
            } else {
                var v73 int8 = 125
                if false {
                    if true {
                        var v74 int16 = 42
                        var v75 uint16 = 0
                        _ = v74
                        _ = v75
                    }
                    var v76 uint64 = 1
                    _ = v76
                } else {
                    v73 = v73
                }
                _ = v73
            }
            _ = v69
        }
        var v77 string = "a"
        _ = v77
    }
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", (("hello" + ("foo bar" + "")) + ""), ("a" + ("hello" + "")), ""}
    var v3 bool = false
    var v4 int = 541
    v5 := make(map[string]bool)
    var v6 int32 = 42
    if v3 {
        var v78 uint = 42
        var v79 bool = (!v3)
        if v3 {
            v80 := GenS10{F0: 32767}
            fmt.Println(v79, v79, v4)
            _ = v80
        }
        _ = v78
        _ = v79
    }
    v3 = (1 > 1)
    if v3 {
        var v81 bool = ((^(v4)) > -2147483648)
        v82 := GenS4{F0: false}
        var v83 uint32 = 42
        _ = v81
        _ = v82
        _ = v83
    } else {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    switch 32767 {
    case -29:
        v4 -= (v4 + 592)
    }
    fmt.Println(v3)
    v6 = (v6 % ((v6) | 1))
    v4 -= v4
    var v84 int32 = (v6 % (((v6 << ((2) & 63))) | 1))
    v4 = v4
    v3 = (-392 != v6)
    if (!(1 == 42)) {
        _ = len(v1)
        switch 76 {
        case 9223372036854775807:
            if v3 {
                if true {
                    var v85 int = v4
                    for fi86 := 0; fi86 < 15; fi86++ {
                        switch ("\n\t" + "\n\t") {
                        case "":
                            fmt.Println(v4, fi86)
                        case "a":
                            var v87 uint16 = 1
                            _ = v87
                        default:
                            v84 = 451
                        }
                        var v88 bool = v3
                        v6 -= (v84 | v84)
                        _ = fi86
                        _ = v88
                    }
                    _ = v85
                }
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                _ = len(v2)
            }
        case 1:
            var v89 uint64 = 562
            _ = v89
        case 515:
            v4 += v4
        }
        var v90 int32 = v6
        _ = v90
    }
    v3 = (422 > 42)
    var v91 bool = ((1 < 2) && v3)
    var v92 uint32 = 716
    if ((v3 && (v91 || v3)) || ((v4 << ((2) & 63)) != (v4 * (v4 | v4)))) {
        for fi93 := 0; fi93 < 14; fi93++ {
            fmt.Println(fi93, v84, v4)
            var v94 float32 = float32(0.1)
            v84 = 0
            _ = fi93
            _ = v94
        }
        for fi95 := 0; fi95 < 7; fi95++ {
            v96 := v1[("\n\t" + ("hello" + "a"))]
            switch (("" + "a") + "a") {
            case "\n\t":
                v4 = -2147483648
            }
            for fi97 := 0; fi97 < 16; fi97++ {
                v98 := GenS10{F0: 1000}
                v99 := GenS6{F0: -1, F1: 792, F2: 705}
                v100 := GenS7{F0: "a", F1: 42}
                _ = fi97
                _ = v100
                _ = v98
                _ = v99
            }
            _ = fi95
            _ = v96
        }
        var v101 uint8 = 1
        _ = v101
    } else {
        v92 += v92
        fmt.Println(v91, v84, v6)
        switch "\n\t" {
        case "a":
            if v3 {
                v102 := GenS10{F0: -866}
                var v103 float32 = 0.0
                var v104 float32 = ((0.1 + v103) / (((3.14159 - ((v103 - v103) - 0.0)) / v103) * v103))
                _ = v102
                _ = v103
                _ = v104
            }
        case "hello":
            fmt.Println(v92, v91)
        }
    }
    var v105 int8 = 90
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v84
    _ = v91
    _ = v92
}
