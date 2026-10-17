package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 uint8; F2 int32 }
type GenS10 struct { F0 uint16; F1 uint32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 int; F2 uint64 }
type GenS14 struct { F0 string; F1 string; F2 string }
func genFn0(p0 uint64, p1 int64) uint {
    var v5 int16 = 42
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int32 = -1
            if true {
                switch "hello" {
                case "a":
                    p0 -= p0
                case "foo bar":
                    var v9 string = ""
                    _ = v9
                case "hello":
                    v10 := GenS1{F0: 0, F1: 255}
                    _ = v10
                default:
                    v5 += -249
                }
            }
            v11 := GenS2{F0: 128, F1: ("\n\t" + "hello")}
            if ((p1 == (p1 % (((p1 / (((p1 * p1)) | 1))) | 1))) || true) {
                var v12 string = "foo bar"
                _ = v12
            } else {
                if true {
                    if true {
                        v13 := GenS2{F0: 2, F1: ""}
                        _ = v13
                    }
                } else {
                    v14 := GenS2{F0: 20, F1: ("" + "a")}
                    v15 := GenS1{F0: 1000, F1: 42}
                    p0 += 18446744073709551615
                    switch (fi6 >> ((18446744073709551615) & 63)) {
                    case -104:
                        v16 := []int{int(v8), fi7, fi7, (fi6 - fi6)}
                        _ = v16
                    case 32767:
                        v17 := []int{32768}
                        _ = v17
                    }
                    p0 -= p0
                    var v18 int8 = -1
                    var v19 uint = 42
                    v8 -= 0
                    v18 += v18
                    _ = v14
                    _ = v15
                    _ = v18
                    _ = v19
                }
            }
            v5 = int16(p1)
            if (p1 > p1) {
                v8 += (v8 * (-(v8)))
                v20 := GenS2{F0: 99, F1: "hello"}
                _ = v20
            } else {
                v21 := GenS3{F0: (p1 | (p1 - 127))}
                _ = v21
            }
            var v22 int16 = (v5 >> ((0) & 63))
            fmt.Println(v22, p1)
            if true {
                fmt.Println(p0)
            }
            _ = fi7
            _ = v11
            _ = v22
            _ = v8
        }
        var v23 bool = (false || (true || ("hello" < "hello")))
        _ = fi6
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS2{F0: 2, F1: "foo bar"}
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}

func main() {
    v1 := []int{671, 2147483647}
    var v2 float32 = float32(0.1)
    v3 := make(map[int]string)
    var v4 float32 = 0.0
    switch "hello" {
    case "":
        var v26 uint = 255
        _ = v26
    default:
        switch 0 {
        case 32767:
            if (227 >= 2) {
                v2 = v4
            } else {
                v2 -= v2
                v4 -= v2
                for fi27 := 0; fi27 < 14; fi27++ {
                    v4 = v4
                    _ = fi27
                }
            }
        }
    }
    delete(v3, 4294967295)
    var v28 string = "foo bar"
    switch (v28 + (v28 + "")) {
    case "\n\t":
        v29 := GenS1{F0: 1000, F1: 572}
        _ = v29
    case "":
        var v30 float32 = v4
        _ = v30
    default:
        fmt.Println(v28)
    }
    var v31 uint8 = 2
    switch -1 {
    case 32767:
        var v32 int64 = -1
        _ = v32
    case 974:
        v31 = v31
    case 128:
        v33 := GenS2{F0: v31, F1: (("\n\t" + "\n\t") + v28)}
        _ = v33
    }
    var v34 uint8 = (v31 + (v31 * 96))
    switch 224 {
    case -611:
        switch -915 {
        case -331:
            v31 += v31
        default:
            for fi35 := 0; fi35 < 14; fi35++ {
                var v36 int64 = -9223372036854775808
                _ = fi35
                _ = v36
            }
        }
    case 0:
        v4 = v4
    default:
        v3[9223372036854775807] = v28
    }
    v37 := GenS3{F0: 1}
    for fi38 := 0; fi38 < 20; fi38++ {
        v3[fi38] = (v28 + (v28 + v28))
        _ = fi38
    }
    v39 := GenS1{F0: 342, F1: 32767}
    if (1 > 327) {
        for fi40 := 0; fi40 < 12; fi40++ {
            var v41 int = (fi40 * 127)
            v42 := GenS3{F0: 9223372036854775807}
            delete(v3, fi40)
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    switch v28 {
    case "\n\t":
        for fi43 := 0; fi43 < 8; fi43++ {
            var v44 uint16 = 1
            _ = fi43
            _ = v44
        }
    case "hello":
        if (true && (false || true)) {
            fmt.Println(v28, v28, v31)
            if (2 <= 774) {
                var v45 uint16 = 65535
                var v46 float64 = 0.0
                switch v28 {
                case "hello":
                    fmt.Println(v34, v31)
                case "":
                    v4 += (v2 * float32(v46))
                }
                _ = v45
                _ = v46
            } else {
                v47 := GenS4{F0: -32768, F1: "hello"}
                var v48 int64 = 42
                v49 := v3[-9223372036854775808]
                _ = v47
                _ = v48
                _ = v49
            }
            fmt.Println(v31, v31, v2)
        } else {
            var v50 uint32 = 2147483648
            _ = v50
        }
    case "a":
        var v51 uint16 = 2
        _ = v51
    default:
        var v52 uint8 = (v34 << ((255) & 63))
        _ = v52
    }
    if (0 <= -1) {
        v31 = (^(v34))
        if (32767 <= -32768) {
            _ = len(v3)
        } else {
            var v53 bool = false
            v34 -= (v34 % ((v34) | 1))
            _ = v53
        }
    } else {
        for fi54 := 0; fi54 < 8; fi54++ {
            if (1 != 437) {
                if (false && ((((0.0 + (0.1 / 3.14159)) / (0.0 - -2.25)) <= 3.14159) && false)) {
                    var v55 int16 = -32768
                    for fi56 := 0; fi56 < 7; fi56++ {
                        if true {
                            v57 := GenS1{F0: v55, F1: 953}
                            _ = v57
                        }
                        _ = fi56
                    }
                    for fi58 := 0; fi58 < 3; fi58++ {
                        if (4294967295 < 0) {
                            switch (((v28 + ("a" + (v28 + v28))) + v28) + "\n\t") {
                            case "hello":
                                v2 = -2.25
                            }
                        } else {
                            var v59 uint64 = 412
                            v60 := GenS5{F0: true, F1: 480}
                            var v61 float64 = float64(float32((1e10 / float64(((1.5 - ((v4 * (v4 * float32((((0.1 + (-2.25 / -2.25)) + 0.0) * 0.0)))) + ((float32((0.1 / (-2.25 * (1e10 / -2.25)))) - v2) * ((v4 + v4) / v2)))) * (v4 + (1.5 - (0.0 * 1e10))))))))
                            v1 = append(v1, (fi58 | -419))
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                        if (!(2147483647 > 42)) {
                            if ((v28 + "a") == v28) {
                                v31 -= (^(v34))
                                var v62 uint64 = 18446744073709551615
                                _ = v62
                            }
                            switch fi54 {
                            case -901:
                                v31 -= 1
                                fmt.Println(v4, v34)
                            case 9223372036854775807:
                                if ((v34 * (v31 * 18)) <= (v31 ^ (v31 & 18))) {
                                    var v63 uint8 = v34
                                    fmt.Println(fi54, fi54, v55)
                                    v55 = -32768
                                    _ = v63
                                }
                            case -9223372036854775808:
                                v34 -= 2
                            default:
                                switch ("a" + v28) {
                                case "":
                                    if false {
                                        if false {
                                            v3[0] = ("" + v28)
                                        } else {
                                            v34 -= 2
                                            v34 = (v34 >> ((42) & 63))
                                            v64 := GenS1{F0: v55, F1: 32767}
                                            _ = v64
                                        }
                                    }
                                }
                            }
                            fmt.Println(v34)
                        } else {
                            if true {
                                if ((float64(1.5) <= 3.14159) || ((v4 - ((-2.25 / v2) / (v4 * (0.0 + v4)))) == 1.5)) {
                                    v65 := GenS4{F0: (^(v55)), F1: v28}
                                    v55 = (v55 - 42)
                                    v2 -= 1.5
                                    fmt.Println(v4, fi54, v31)
                                    v28 = (v28 + "")
                                    _ = v65
                                }
                                if true {
                                    v66 := v3[fi54]
                                    v67 := GenS6{F0: (((((v2 * (float32(1.5) - v4)) + v2) + v2) / v4) <= (v2 * v2)), F1: 255, F2: 1}
                                    _ = v66
                                    _ = v67
                                } else {
                                    fmt.Println(v2)
                                    var v68 int64 = 9223372036854775807
                                    _ = v68
                                }
                                var v69 int64 = -1
                                _ = v69
                            }
                        }
                        var v70 int8 = 127
                        _ = fi58
                        _ = v70
                    }
                    _ = v55
                }
                if true {
                    switch "\n\t" {
                    case "":
                        v71 := GenS6{F0: false, F1: 1, F2: 4294967295}
                        _ = v71
                    case "\n\t":
                        if (602 < 2147483648) {
                            delete(v3, fi54)
                        }
                    case "a":
                        if (328 == 2) {
                            var v72 uint = genFn0(478, -209)
                            if true {
                                for fi73 := 0; fi73 < 7; fi73++ {
                                    if (-1 < 0) {
                                        var v74 bool = true
                                        v75 := GenS7{F0: fi54, F1: -9223372036854775808}
                                        _ = v74
                                        _ = v75
                                    } else {
                                        var v76 uint16 = 65535
                                        v76 += (v76 ^ (v76 + (v76 / ((uint16(v72)) | 1))))
                                        v77 := GenS6{F0: true, F1: v76, F2: (^(v72))}
                                        _ = v76
                                        _ = v77
                                    }
                                    var v78 int16 = 1
                                    var v79 int16 = v78
                                    _ = fi73
                                    _ = v78
                                    _ = v79
                                }
                                if len(v1) > 0 {
                                    v1[0] = fi54
                                }
                                v80 := GenS8{F0: -945}
                                _ = v80
                            } else {
                                var v81 uint32 = 2147483648
                                _ = v81
                            }
                            _ = v72
                        }
                    default:
                        switch v28 {
                        case "foo bar":
                            for fi82 := 0; fi82 < 7; fi82++ {
                                var v83 uint = 286
                                if (false || false) {
                                    switch "hello" {
                                    case "foo bar":
                                        delete(v3, fi54)
                                    }
                                    v84 := GenS9{F0: v28, F1: v31, F2: -2147483648}
                                    switch (fi54 % (((fi54 >> ((genFn0(0, -1)) & 63))) | 1)) {
                                    case -1:
                                        if (2147483647 >= 2147483647) {
                                            switch v28 {
                                            case "a":
                                                if false {
                                                    v4 = v2
                                                }
                                            case "hello":
                                                var v85 uint16 = uint16(v83)
                                                _ = v85
                                            default:
                                                v2 -= -2.25
                                            }
                                            if (float32(0.1) <= (0.1 * (1.5 / v4))) {
                                                switch fi82 {
                                                case 32768:
                                                    switch "" {
                                                    case "":
                                                        if false {
                                                            var v86 uint64 = 42
                                                            _ = len(v3)
                                                            _ = v86
                                                        }
                                                    case "a":
                                                        v4 = 3.14159
                                                    case "\n\t":
                                                        v83 += v83
                                                    }
                                                default:
                                                    v87 := GenS10{F0: 2, F1: 2147483647}
                                                    _ = v87
                                                }
                                            } else {
                                                _ = len(v1)
                                            }
                                            v28 = ("" + v28)
                                        }
                                    case 816:
                                        if (-69 == 42) {
                                            if len(v1) > 0 {
                                                v1[0] = (-(fi82))
                                            }
                                            var v88 uint16 = 32767
                                            v88 += 65535
                                            _ = v88
                                        } else {
                                            var v89 string = v28
                                            var v90 uint16 = 709
                                            fmt.Println(v28, v89, fi82)
                                            _ = v89
                                            _ = v90
                                        }
                                    default:
                                        v91 := GenS5{F0: false, F1: 4294967295}
                                        _ = v91
                                    }
                                    _ = v84
                                } else {
                                    if (v4 == 1.5) {
                                        v92 := GenS11{F0: v31}
                                        var v93 int16 = 110
                                        var v94 uint = v83
                                        if true {
                                            var v95 float64 = float64(1e10)
                                            v2 = 1e10
                                            var v96 int8 = 1
                                            _ = v95
                                            _ = v96
                                        }
                                        _ = v92
                                        _ = v93
                                        _ = v94
                                    } else {
                                        v97 := GenS12{F0: -437}
                                        _ = v97
                                    }
                                    var v98 uint32 = 42
                                    v1 = append(v1, (fi82 & 341))
                                    var v99 int16 = 42
                                    _ = v98
                                    _ = v99
                                }
                                var v100 int8 = -128
                                _ = fi82
                                _ = v100
                                _ = v83
                            }
                        case "hello":
                            var v101 uint64 = 227
                            _ = v101
                        case "\n\t":
                            v102 := GenS1{F0: 1000, F1: 516}
                            _ = v102
                        default:
                            v31 = 1
                        }
                    }
                    v3[(fi54 % ((fi54) | 1))] = ((v28 + (v28 + v28)) + v28)
                }
            } else {
                switch "a" {
                case "a":
                    var v103 int32 = -87
                    _ = v103
                case "":
                    var v104 string = ""
                    _ = v104
                default:
                    switch "foo bar" {
                    case "foo bar":
                        var v105 int64 = -473
                        _ = v105
                    case "\n\t":
                        for fi106 := 0; fi106 < 8; fi106++ {
                            _ = len(v1)
                            _ = fi106
                        }
                    case "":
                        if true {
                            v34 = 42
                            for fi107 := 0; fi107 < 17; fi107++ {
                                v2 = v4
                                _ = fi107
                            }
                            if (0 != -2147483648) {
                                v28 = "a"
                                v3[(-(fi54))] = (v28 + ("foo bar" + ""))
                                for fi108 := 0; fi108 < 1; fi108++ {
                                    v4 += 0.1
                                    _ = fi108
                                }
                            }
                        }
                    }
                }
            }
            for fi109 := 0; fi109 < 1; fi109++ {
                if true {
                    if (127 > 84) {
                        fmt.Println(v31)
                        _ = len(v1)
                        v4 += float32(0.1)
                    } else {
                        v34 = v34
                        fmt.Println(v28, v28, v31)
                        var v110 float32 = 0.1
                        _ = v110
                    }
                    v111 := GenS13{F0: -899, F1: 128, F2: 42}
                    switch v28 {
                    case "\n\t":
                        fmt.Println(fi109)
                    }
                    if false {
                        if true {
                            v28 = (("" + v28) + (v28 + v28))
                            var v112 int16 = -32768
                            _ = v112
                        }
                    } else {
                        fmt.Println(v2)
                        v2 += (1.5 + (3.14159 - (v2 + v2)))
                        var v113 int = fi54
                        v114 := GenS14{F0: "hello", F1: v28, F2: "\n\t"}
                        _ = v113
                        _ = v114
                    }
                    v34 -= (^(v34))
                    _ = v111
                }
                _ = fi109
            }
            v115 := GenS4{F0: 32767, F1: ("a" + v28)}
            _ = fi54
            _ = v115
        }
    }
    if (true && (100 <= 0)) {
        v1 = append(v1, 127)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v34
    _ = v37
    _ = v39
    _ = v4
}
