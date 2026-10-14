package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 string; F1 int32; F2 uint64 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int32 }
func genFn0(p0 bool, p1 bool) uint64 {
    switch 1 {
    case 304:
        v7 := GenS1{F0: 554, F1: 0, F2: -746}
        _ = v7
    case -9223372036854775808:
        p1 = p0
    default:
        var v8 string = ""
        _ = v8
    }
    fmt.Println(p1, p1)
    var v9 string = "hello"
    _ = p0
    _ = p1
    _ = v9
    return 683
}

func main() {
    var v1 bool = false
    var v2 uint8 = 128
    var v3 uint8 = 127
    var v4 uint8 = 2
    var v5 int8 = -1
    var v6 float32 = ((0.1 - -2.25) / ((float32(3.14159) * 3.14159) / 1.0))
    if v1 {
        v10 := GenS1{F0: 258, F1: 0, F2: 42}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            if ((((42 == 2147483648) || false) && (1e10 != float32((-2.25 * 0.1)))) && (42 != 32767)) {
                v12 := GenS1{F0: -1, F1: 9223372036854775807, F2: 2147483647}
                fmt.Println(v1, v5)
                _ = v12
            } else {
                fmt.Println(v6)
            }
            _ = fi11
        }
        v13 := GenS1{F0: 32767, F1: 42, F2: int32(v2)}
        _ = v13
    }
    switch 127 {
    case 398:
        v14 := GenS2{F0: 18446744073709551615, F1: v4}
        _ = v14
    case 1:
        v1 = (((42 <= 32767) && v1) && false)
    }
    switch ("" + ("hello" + ("a" + "a"))) {
    case "foo bar":
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := []int{(fi15 >> ((18446744073709551615) & 63))}
            var v17 string = ("a" + ("\n\t" + "hello"))
            _ = fi15
            _ = v16
            _ = v17
        }
    case "\n\t":
        v18 := GenS1{F0: -633, F1: 159, F2: -2147483648}
        _ = v18
    case "a":
        v5 = v5
    }
    if (false && v1) {
        v5 += 0
        v4 += v3
    } else {
        switch (("foo bar" + "a") + "") {
        case "foo bar":
            v19 := []int{128, 368, -600}
            _ = v19
        case "\n\t":
            var v20 int32 = -2147483648
            _ = v20
        default:
            v21 := GenS2{F0: 2, F1: v4}
            _ = v21
        }
        if v1 {
            var v22 string = "\n\t"
            v4 = 0
            switch ("a" + (v22 + v22)) {
            case "hello":
                for fi23 := 0; fi23 < 8; fi23++ {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        if (!(1 == 42)) {
                            v25 := GenS2{F0: 295, F1: (^(v3))}
                            _ = v25
                        }
                        v26 := []int{(fi23 / ((fi24) | 1)), 127, fi23, (fi23 / (((fi23 - fi23)) | 1))}
                        _ = fi24
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 4; fi27++ {
                        fmt.Println(v3)
                        var v28 int64 = -899
                        _ = fi27
                        _ = v28
                    }
                    _ = fi23
                }
            default:
                v6 += v6
            }
            _ = v22
        }
        v4 += v2
    }
    if (-2147483648 < -2147483648) {
        var v29 string = "a"
        var v30 int = 128
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: 1, F1: v4}
    v32 := GenS3{F0: -345, F1: 1, F2: 2}
    v1 = v1
    if (4294967295 == 282) {
        var v33 uint64 = 42
        v34 := []int{-2147483648}
        var v35 int16 = 120
        _ = v33
        _ = v34
        _ = v35
    } else {
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := GenS1{F0: 42, F1: 1, F2: -1}
            if (v1 || v1) {
                fmt.Println(v4)
                if v1 {
                    var v38 int8 = v5
                    fmt.Println(v1, v6, v4)
                    v39 := []int{}
                    _ = v38
                    _ = v39
                } else {
                    v1 = v1
                    v40 := []int{}
                    var v41 float32 = (v6 * 0.1)
                    _ = v40
                    _ = v41
                }
                var v42 int64 = -879
                _ = v42
            }
            _ = fi36
            _ = v37
        }
        v3 = 42
    }
    v43 := GenS1{F0: 791, F1: 1, F2: 1}
    if (!true) {
        if v1 {
            v44 := []int{2147483647, 128, -9223372036854775808, 9223372036854775807}
            fmt.Println(v1, v6, v4)
            _ = v44
        } else {
            if v1 {
                if true {
                    v45 := []int{217}
                    fmt.Println(v2, v1, v2)
                    _ = v45
                } else {
                    var v46 float64 = (-2.25 - (3.14159 - -2.25))
                    if v1 {
                        for fi47 := 0; fi47 < 18; fi47++ {
                            v48 := []int{-1, fi47, (^(fi47)), fi47}
                            _ = fi47
                            _ = v48
                        }
                    } else {
                        v49 := GenS4{F0: ("hello" + ("foo bar" + "\n\t"))}
                        v50 := []int{32768, -1}
                        v46 = v46
                        _ = v49
                        _ = v50
                    }
                    var v51 uint32 = 0
                    _ = v46
                    _ = v51
                }
                switch "a" {
                case "":
                    fmt.Println(v5, v5, v3)
                case "a":
                    var v52 uint16 = 145
                    _ = v52
                default:
                    if v1 {
                        v5 = v5
                    } else {
                        for fi53 := 0; fi53 < 9; fi53++ {
                            v3 = v2
                            var v54 float64 = (((0.0 - 1e10) * (0.0 + float64(1e10))) - -2.25)
                            var v55 float64 = 0.1
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        v6 = (((v6 - -2.25) + v6) * v6)
                    }
                }
            }
            v1 = ((!v1) || v1)
            var v56 uint = 496
            _ = v56
        }
    } else {
        var v57 int64 = 42
        if (1 != 1) {
            v6 += (0.0 * float32(1.5))
        }
        for fi58 := 0; fi58 < 2; fi58++ {
            v2 = v4
            if ((((9223372036854775807 < -2147483648) && (((478 > 262) && (!true)) && v1)) || false) || (32767 == -1)) {
                if ((true && (77 <= (v5 * v5))) || (!v1)) {
                    v59 := GenS5{F0: (1 < -527)}
                    for fi60 := 0; fi60 < 11; fi60++ {
                        if false {
                            var v61 float64 = (3.14159 * -2.25)
                            v62 := []int{fi58}
                            v63 := GenS6{F0: "hello", F1: 101}
                            if len(v62) >= 2 {
                                v62 = v62[1:len(v62)]
                            }
                            v62 = append(v62, fi60)
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        var v64 float32 = (float32((0.1 / 1e10)) * (float32(3.14159) * v6))
                        var v65 int8 = -64
                        _ = fi60
                        _ = v64
                        _ = v65
                    }
                    _ = v59
                } else {
                    for fi66 := 0; fi66 < 10; fi66++ {
                        v4 = (v2 | v4)
                        v4 += (v4 | 42)
                        _ = fi66
                    }
                    for fi67 := 0; fi67 < 13; fi67++ {
                        if true {
                            v68 := GenS4{F0: "a"}
                            _ = v68
                        }
                        _ = fi67
                    }
                }
                if (!((2 <= 4294967295) || v1)) {
                    for fi69 := 0; fi69 < 20; fi69++ {
                        if v1 {
                            if true {
                                var v70 int = (fi69 / ((127) | 1))
                                if v1 {
                                    v71 := []int{32767}
                                    if v1 {
                                        if (v3 < 165) {
                                            var v72 uint16 = 255
                                            fmt.Println(v57, fi58)
                                            _ = v72
                                        } else {
                                            _ = len(v71)
                                        }
                                        switch 9223372036854775807 {
                                        case 42:
                                            switch 32768 {
                                            case 1:
                                                if len(v71) >= 2 {
                                                    v71 = v71[1:len(v71)]
                                                }
                                            default:
                                                _ = len(v71)
                                            }
                                        case 0:
                                            fmt.Println(v5, fi58, v57)
                                        case 616:
                                            fmt.Println(v57)
                                            if true {
                                                v4 = (^(v3))
                                                v5 -= (^(v5))
                                            }
                                        }
                                    } else {
                                        v73 := GenS3{F0: 0, F1: v57, F2: v4}
                                        if ((-2.25 - 1.5) != (3.14159 + 3.14159)) {
                                            var v74 int8 = (v5 * 127)
                                            var v75 uint32 = 989
                                            switch v70 {
                                            case -2147483648:
                                                if (2 != 972) {
                                                    var v76 uint32 = 4294967295
                                                    var v77 int64 = (^(v57))
                                                    fmt.Println(v70)
                                                    _ = v76
                                                    _ = v77
                                                }
                                            default:
                                                v78 := GenS7{F0: ("\n\t" + "\n\t"), F1: -979, F2: 652}
                                                _ = v78
                                            }
                                            if len(v71) > 0 {
                                                v71[0] = fi69
                                            }
                                            _ = v74
                                            _ = v75
                                        } else {
                                            if ("" > ("a" + "a")) {
                                                fmt.Println(v5)
                                                v4 -= (^(v2))
                                            }
                                            v3 -= 255
                                        }
                                        v79 := GenS6{F0: (("a" + "a") + "a"), F1: (-(v5))}
                                        _ = v73
                                        _ = v79
                                    }
                                    var v80 string = "\n\t"
                                    _ = v71
                                    _ = v80
                                } else {
                                    v57 = 0
                                    v81 := GenS8{F0: 2147483648, F1: (v70 - (fi69 * (v70 * v70)))}
                                    if (v1 || v1) {
                                        switch v70 {
                                        case 2147483647:
                                            v82 := GenS9{F0: genFn0((0 >= 273), (0 > -2147483648)), F1: 999}
                                            _ = v82
                                        case 748:
                                            switch fi69 {
                                            case 4294967295:
                                                v2 = 115
                                            }
                                        case 1:
                                            if v1 {
                                                switch (fi69 % ((9223372036854775807) | 1)) {
                                                case 0:
                                                    v3 -= (v3 / (((v4 * v4)) | 1))
                                                case 32767:
                                                    fmt.Println(v57, v57)
                                                }
                                                var v83 bool = (65535 >= 42)
                                                _ = v83
                                            }
                                        }
                                        v84 := GenS7{F0: "foo bar", F1: 42, F2: 42}
                                        switch (fi69 * -140) {
                                        case -723:
                                            var v85 int = (fi58 * (v70 + (-(fi69))))
                                            _ = v85
                                        case 200:
                                            v4 -= (v3 ^ 148)
                                        default:
                                            v86 := []int{(fi69 - -153)}
                                            _ = v86
                                        }
                                        _ = v84
                                    }
                                    switch "a" {
                                    case "a":
                                        v4 += (v4 / (((v3 % ((127) | 1))) | 1))
                                    }
                                    if (2 >= 4294967295) {
                                        v87 := []int{v70, (fi69 - 0)}
                                        v88 := GenS3{F0: fi58, F1: v57, F2: 127}
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(v4, v5, v4)
                                            if (975 >= 18446744073709551615) {
                                                var v89 uint32 = 42
                                                _ = v89
                                            } else {
                                                if len(v87) >= 2 {
                                                    v87 = v87[1:len(v87)]
                                                }
                                            }
                                            var v90 uint = 333
                                            _ = v90
                                        case "\n\t":
                                            if (!(!true)) {
                                                v91 := GenS5{F0: (!v1)}
                                                if len(v87) > 0 {
                                                    v87[0] = fi58
                                                }
                                                v2 -= 42
                                                fmt.Println(v5, fi58, v70)
                                                if len(v87) >= 2 {
                                                    v87 = v87[1:len(v87)]
                                                }
                                                v92 := GenS4{F0: "\n\t"}
                                                var v93 bool = (fi69 == int(v5))
                                                _ = v91
                                                _ = v92
                                                _ = v93
                                            }
                                            v94 := GenS9{F0: 0, F1: 0}
                                            if false {
                                                _ = len(v87)
                                                fmt.Println(v1)
                                                v95 := GenS8{F0: 42, F1: (v70 & -260)}
                                                _ = v95
                                            }
                                            _ = v94
                                        }
                                        _ = v87
                                        _ = v88
                                    }
                                    _ = v81
                                }
                                _ = v70
                            }
                        } else {
                            if v1 {
                                var v96 int = fi58
                                var v97 int32 = -2147483648
                                if (v97 < (v97 ^ (v97 >> ((947) & 63)))) {
                                    var v98 int8 = (v5 * -8)
                                    var v99 int64 = (v57 + v57)
                                    _ = v98
                                    _ = v99
                                } else {
                                    v100 := []int{(-(fi58)), (v96 ^ (^(v96))), -525}
                                    _ = v100
                                }
                                fmt.Println(v3, fi58)
                                if (((0.1 / 1e10) + 1e10) < 1e10) {
                                    v57 -= 1
                                    var v101 int = fi69
                                    v2 = 184
                                    switch ("" + "hello") {
                                    case "hello":
                                        v5 += 107
                                    default:
                                        switch fi58 {
                                        case 128:
                                            v102 := []int{(fi69 & v101), (fi69 / ((1) | 1))}
                                            _ = v102
                                        case 440:
                                            fmt.Println(v96)
                                        default:
                                            v103 := []int{}
                                            _ = v103
                                        }
                                    }
                                    _ = v101
                                }
                                _ = v96
                                _ = v97
                            } else {
                                v1 = v1
                            }
                            var v104 int = (fi69 ^ 128)
                            _ = v104
                        }
                        _ = fi69
                    }
                    for fi105 := 0; fi105 < 1; fi105++ {
                        v57 = -1
                        _ = fi105
                    }
                    var v106 uint64 = 4294967295
                    _ = v106
                }
                var v107 float64 = float64((3.14159 - v6))
                _ = v107
            }
            v108 := []int{2147483647, fi58}
            _ = fi58
            _ = v108
        }
        _ = v57
    }
    v109 := GenS10{F0: 42}
    v110 := []int{}
    if (32768 > 4294967295) {
        if (float32(1.5) >= (v6 / 1.0)) {
            var v111 uint = 1
            _ = v111
        }
        v5 = v5
    } else {
        var v112 string = "hello"
        switch "\n\t" {
        case "a":
            v112 = v112
        }
        v3 = v2
        _ = v112
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v110:", len(v110), v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
