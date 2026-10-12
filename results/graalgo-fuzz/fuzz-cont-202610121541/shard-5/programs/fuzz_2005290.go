package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 bool }
type GenS2 struct { F0 int; F1 int64; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 uint64; F1 int; F2 int32 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 uint16; F1 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 int16 = 772
    var v4 int8 = -92
    v5 := make(map[uint32]string)
    v6 := make(map[int]string)
    fmt.Println(v4, v3)
    v4 -= (v4 & (v4 + -109))
    var v7 uint64 = 132
    var v8 uint16 = 65535
    switch "hello" {
    case "foo bar":
        var v9 float64 = 1e10
        _ = v9
    case "a":
        fmt.Println(v4, v4)
    }
    v8 = 65535
    var v10 string = (("foo bar" + "foo bar") + ("hello" + "a"))
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v4)
            fmt.Println(v10, v8, fi11)
            _ = fi12
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 bool = ((-246 <= 362) || (-701 > 1))
            v15 := []int{(fi13 % ((fi13) | 1)), fi11, (fi13 ^ -81), (fi13 & fi11)}
            var v16 float32 = 1.5
            fmt.Println(v7)
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi11
    }
    var v17 int8 = v4
    if true {
        v3 += (v3 & (v3 & (v3 / ((1000) | 1))))
    } else {
        fmt.Println(v4, v7, v10)
    }
    switch "hello" {
    case "":
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 int32 = -232
            v7 -= (v7 ^ v7)
            _ = fi18
            _ = v19
        }
    case "hello":
        var v20 uint32 = 707
        _ = v20
    case "\n\t":
        var v21 int32 = -442
        _ = v21
    }
    v22 := []int{957, -145, 127, 32767}
    v8 = (v8 - (v8 % ((2) | 1)))
    for fi23 := 0; fi23 < 2; fi23++ {
        switch v10 {
        case "hello":
            if false {
                v24 := GenS1{F0: (v8 - 708), F1: v10, F2: false}
                if (!true) {
                    v17 -= (v17 / ((42) | 1))
                    var v25 int64 = -9223372036854775808
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        var v27 int8 = v17
                        if false {
                            var v28 float64 = -2.25
                            var v29 float64 = 1e10
                            switch fi23 {
                            case 0:
                                if (!(2147483647 == 2147483647)) {
                                    v30 := GenS1{F0: v8, F1: ((v10 + v10) + "a"), F2: (fi26 != (-(fi23)))}
                                    _ = v30
                                } else {
                                    fmt.Println(v17)
                                    if true {
                                        _ = len(v22)
                                        v31 := v1[(^(fi26))]
                                        _ = v31
                                    }
                                    v32 := GenS1{F0: uint16(v7), F1: (v10 + v10), F2: false}
                                    _ = v32
                                }
                            case -1:
                                if true {
                                    fmt.Println(fi23)
                                    switch ((v10 + (v10 + ((v10 + v10) + "a"))) + v10) {
                                    case "hello":
                                        var v33 int8 = (^(v17))
                                        _ = v33
                                    case "foo bar":
                                        switch -680 {
                                        case 77:
                                            if (uint(v7) >= 42) {
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v3)
                                                }
                                                var v34 int16 = (v3 << ((4294967295) & 63))
                                                var v35 uint16 = (v8 << ((65535) & 63))
                                                _ = v34
                                                _ = v35
                                            } else {
                                                var v36 uint8 = 255
                                                if (false || false) {
                                                    var v37 uint64 = (v7 - v7)
                                                    v7 += 0
                                                    _ = v37
                                                } else {
                                                    if (0.1 >= 1.5) {
                                                        if (!true) {
                                                            v36 = 241
                                                        }
                                                        var v38 uint64 = 4294967295
                                                        v39 := GenS2{F0: (^(fi23)), F1: -300, F2: (v10 + "")}
                                                        _ = v38
                                                        _ = v39
                                                    } else {
                                                        v17 += -58
                                                        var v40 uint8 = 51
                                                        switch v10 {
                                                        case "":
                                                            _ = len(v1)
                                                        case "hello":
                                                            v36 += 2
                                                            v3 = v3
                                                            v28 -= v28
                                                        case "\n\t":
                                                            fmt.Println(v29, v17)
                                                        }
                                                        if (!(9223372036854775807 == int64(v3))) {
                                                            var v41 int64 = -931
                                                            _ = v41
                                                        }
                                                        _ = v40
                                                    }
                                                    if len(v22) > 0 {
                                                        v22[0] = int(v8)
                                                    }
                                                }
                                                if (v7 == 309) {
                                                    var v42 uint64 = (v7 & (^(v7)))
                                                    _ = v42
                                                }
                                                _ = v36
                                            }
                                        case -279:
                                            var v43 float32 = -2.25
                                            _ = v43
                                        default:
                                            if true {
                                                v44 := GenS3{F0: (v4 % ((v17) | 1))}
                                                _ = v44
                                            } else {
                                                fmt.Println(v17)
                                            }
                                        }
                                    }
                                }
                            case 42:
                                fmt.Println(v8, fi26, v7)
                            default:
                                v2[987] = v10
                            }
                            _ = v28
                            _ = v29
                        }
                        if false {
                            v45 := GenS4{F0: -1, F1: "foo bar", F2: 4}
                            v8 = 0
                            fmt.Println(v17, v27, v4)
                            _ = v45
                        } else {
                            v27 = (-(v4))
                            _ = len(v22)
                            v8 += v8
                            switch "hello" {
                            case "":
                                if true {
                                    v2[172] = v10
                                } else {
                                    v46 := GenS1{F0: 32767, F1: (("a" + "a") + v10), F2: false}
                                    fmt.Println(v17, v3, v8)
                                    if false {
                                        var v47 bool = (!false)
                                        _ = v47
                                    } else {
                                        var v48 string = v10
                                        _ = v48
                                    }
                                    _ = v46
                                }
                                var v49 bool = ((1000 > v3) && true)
                                _ = v49
                            case "a":
                                if false {
                                    v17 = (v27 | (v27 | v17))
                                } else {
                                    if true {
                                        switch 1 {
                                        case 828:
                                            switch fi26 {
                                            case -9223372036854775808:
                                                switch (v10 + v10) {
                                                case "foo bar":
                                                    switch (v10 + v10) {
                                                    case "hello":
                                                        var v50 int = -506
                                                        _ = v50
                                                    }
                                                default:
                                                    _ = len(v22)
                                                }
                                            case 4294967295:
                                                if ((2147483647 == -866) && (true && (true || ((v8 ^ uint16(v7)) <= 42)))) {
                                                    switch ("hello" + v10) {
                                                    case "hello":
                                                        v51 := GenS5{F0: int64(v8), F1: v7}
                                                        _ = v51
                                                    case "foo bar":
                                                        v17 = 127
                                                    }
                                                } else {
                                                    delete(v5, 0)
                                                    var v52 uint64 = 1
                                                    _ = v52
                                                }
                                            case -1:
                                                v53 := GenS4{F0: -2147483648, F1: "", F2: 81}
                                                _ = v53
                                            default:
                                                v8 -= (^(v8))
                                                fmt.Println(v3, v10, v27)
                                                if true {
                                                    v22 = append(v22, fi26)
                                                    var v54 int = (fi26 / ((fi26) | 1))
                                                    var v55 int = (fi23 & (-(fi23)))
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    v56 := GenS1{F0: v8, F1: ("foo bar" + "a"), F2: ((true || (false || true)) || (v7 != v7))}
                                                    v4 += int8(v3)
                                                    if true {
                                                        v57 := v6[-623]
                                                        v5[621] = (v57 + "hello")
                                                        v58 := GenS2{F0: (fi26 | 32768), F1: 42, F2: v10}
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        var v59 int32 = -188
                                                        v60 := GenS2{F0: (fi23 % (((^(fi23))) | 1)), F1: -144, F2: ("a" + v10)}
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    if false {
                                                        if (v7 != (v7 & v7)) {
                                                            v7 += 1
                                                            switch v10 {
                                                            case "foo bar":
                                                                fmt.Println(v3, v8)
                                                            case "hello":
                                                                var v61 int16 = int16(v7)
                                                                _ = v61
                                                            case "a":
                                                                if len(v22) > 0 {
                                                                    v22[0] = fi23
                                                                }
                                                            default:
                                                                var v62 int = int(v8)
                                                                _ = v62
                                                            }
                                                            var v63 int64 = -373
                                                            v64 := GenS4{F0: -2147483648, F1: v10, F2: 2}
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                    }
                                                    var v65 int64 = 0
                                                    if false {
                                                        v10 = v10
                                                        var v66 uint = 2
                                                        _ = v66
                                                    }
                                                    _ = v56
                                                    _ = v65
                                                }
                                            }
                                        }
                                        switch ("a" + (v10 + v10)) {
                                        case "foo bar":
                                            var v67 float32 = 1.5
                                            _ = v67
                                        case "a":
                                            v1[fi26] = (57 >= 42)
                                            var v68 int16 = v3
                                            _ = len(v1)
                                            if (true && (((0.0 + 1.5) > 0.1) || ((-2.25 - ((1e10 + 0.1) * 0.0)) <= 0.1))) {
                                                var v69 float32 = 0.0
                                                v70 := GenS2{F0: (fi26 & fi23), F1: -1, F2: ((v10 + v10) + (v10 + ""))}
                                                _ = v69
                                                _ = v70
                                            }
                                            v71 := GenS5{F0: 1, F1: v7}
                                            _ = v68
                                            _ = v71
                                        case "\n\t":
                                            if (!(false && (115 < (v17 * (v4 ^ (v17 | v17)))))) {
                                                v4 = v4
                                                v72 := GenS6{F0: (v4 & v27)}
                                                _ = v72
                                            }
                                        default:
                                            var v73 uint8 = 215
                                            _ = v73
                                        }
                                    }
                                    v7 = (v7 & (v7 - v7))
                                    switch (fi26 & fi26) {
                                    case 496:
                                        if false {
                                            v3 -= (-(v3))
                                            var v74 uint64 = v7
                                            if false {
                                                v75 := GenS1{F0: (v8 >> ((uint(v74)) & 63)), F1: v10, F2: (65535 >= 576)}
                                                var v76 float32 = float32(float64(((1.5 / 0.1) - 0.0)))
                                                v3 = 0
                                                _ = v75
                                                _ = v76
                                            } else {
                                                v27 += -128
                                            }
                                            v22 = append(v22, fi23)
                                            v77 := GenS7{F0: 0, F1: 1}
                                            _ = v74
                                            _ = v77
                                        } else {
                                            var v78 float32 = -2.25
                                            v10 = v10
                                            switch "hello" {
                                            case "hello":
                                                v79 := GenS8{F0: 2, F1: fi26, F2: -459}
                                                _ = v79
                                            case "foo bar":
                                                fmt.Println(v78, v4)
                                            }
                                            _ = v78
                                        }
                                    default:
                                        switch fi23 {
                                        case 128:
                                            v80 := GenS3{F0: (^(v17))}
                                            _ = v80
                                        }
                                    }
                                }
                            }
                            v81 := GenS9{F0: 1, F1: v17}
                            _ = v81
                        }
                        _ = fi26
                        _ = v27
                    }
                    if false {
                        v82 := GenS7{F0: 2147483648, F1: -201}
                        v10 = "foo bar"
                        _ = v82
                    }
                }
                _ = v24
            }
        default:
            fmt.Println(v4, v10, v3)
        }
        if false {
            var v83 int8 = -79
            delete(v6, fi23)
            v8 = v8
            _ = v83
        } else {
            v84 := GenS2{F0: 1, F1: 0, F2: "a"}
            v3 = (v3 << ((4294967295) & 63))
            for fi85 := 0; fi85 < 9; fi85++ {
                v4 -= v17
                _ = fi85
            }
            _ = v84
        }
        for fi86 := 0; fi86 < 8; fi86++ {
            v6[fi23] = ("hello" + "hello")
            switch ((v10 + v10) + v10) {
            case "foo bar":
                var v87 float32 = (1e10 - 3.14159)
                _ = v87
            case "":
                if (!true) {
                    _ = len(v5)
                } else {
                    v88 := GenS3{F0: v4}
                    if true {
                        v89 := v6[-584]
                        if (65535 != 297) {
                            fmt.Println(v4, v17)
                            switch "" {
                            case "a":
                                var v90 int = 9223372036854775807
                                _ = v90
                            case "":
                                v22 = append(v22, (fi86 + fi86))
                            default:
                                var v91 int64 = -9223372036854775808
                                _ = v91
                            }
                        } else {
                            var v92 float64 = (float64((float32(0.0) / 3.14159)) - 1.5)
                            fmt.Println(v3, v89)
                            fmt.Println(v8, v92, v7)
                            if (42 <= 1) {
                                if false {
                                    var v93 uint32 = 65535
                                    _ = v93
                                } else {
                                    var v94 float64 = float64(((3.14159 + 1e10) - 0.0))
                                    fmt.Println(v3)
                                    _ = v94
                                }
                                v4 -= (v4 - (v17 * int8(v8)))
                            } else {
                                switch v89 {
                                case "foo bar":
                                    switch (v89 + v10) {
                                    case "foo bar":
                                        var v95 uint16 = (v8 & v8)
                                        _ = v95
                                    case "":
                                        if (false && ((v8 & v8) < 521)) {
                                            var v96 int = fi23
                                            v97 := GenS3{F0: (v4 * (v4 + (v4 + v17)))}
                                            fmt.Println(v4, fi23)
                                            v8 += 475
                                            _ = v96
                                            _ = v97
                                        } else {
                                            switch ("foo bar" + v89) {
                                            case "hello":
                                                switch -1 {
                                                case 127:
                                                    v10 = ("hello" + (v10 + v10))
                                                }
                                            case "foo bar":
                                                var v98 int8 = (v17 + (v4 - v17))
                                                _ = v98
                                            case "":
                                                var v99 float32 = (0.0 * 0.0)
                                                _ = v99
                                            }
                                        }
                                        v3 += (v3 % (((v3 >> ((18446744073709551615) & 63))) | 1))
                                    case "hello":
                                        fmt.Println(v92, v8, v89)
                                    }
                                case "\n\t":
                                    fmt.Println(v92, fi86, v89)
                                default:
                                    var v100 uint16 = 545
                                    v7 += 83
                                    fmt.Println(v17, v100)
                                    _ = v100
                                }
                            }
                            _ = v92
                        }
                        _ = v89
                    } else {
                        v4 -= (v17 / ((v17) | 1))
                        v3 = (v3 + v3)
                        if (false && true) {
                            var v101 uint32 = 4294967295
                            var v102 int = 4294967295
                            _ = v101
                            _ = v102
                        }
                    }
                    v103 := GenS10{F0: v8, F1: (float32(-2.25) != 1e10)}
                    var v104 float64 = 3.14159
                    _ = v103
                    _ = v104
                    _ = v88
                }
            case "hello":
                v6[(^(fi23))] = (v10 + "a")
            }
            v5[2147483648] = ((v10 + v10) + "a")
            _ = fi86
        }
        _ = fi23
    }
    var v105 uint16 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v105:", v105)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v105
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
