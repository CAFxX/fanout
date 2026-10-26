package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 string }
type GenS5 struct { F0 int8; F1 uint32; F2 string }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int; F1 uint8; F2 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string }

func main() {
    var v1 int64 = 0
    var v2 bool = true
    var v3 uint16 = 0
    switch -2147483648 {
    case 32767:
        var v4 int = -9223372036854775808
        _ = v4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        if (0.0 <= 1e10) {
            var v6 int = (-(fi5))
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 float32 = -2.25
                v9 := []int{(v6 / ((4294967295) | 1)), (fi7 * fi7)}
                var v10 uint = 42
                if v2 {
                    if ((2147483648 <= 60) && (0 >= 937)) {
                        v10 -= (v10 & (v10 + v10))
                    }
                    v8 = v8
                } else {
                    var v11 float32 = v8
                    _ = v11
                }
                var v12 string = "hello"
                var v13 int64 = (v1 | v1)
                v3 = v3
                var v14 int = fi7
                switch (fi5 | (fi5 & (fi7 + fi7))) {
                case -2147483648:
                    var v15 uint32 = 65535
                    _ = v15
                default:
                    var v16 bool = (0.0 != ((v8 - v8) * (1e10 * v8)))
                    _ = v16
                }
                var v17 int = 127
                v13 = v1
                switch "a" {
                case "a":
                    if false {
                        v14 += fi7
                        if len(v9) > 0 {
                            v9[0] = -622
                        }
                        v10 -= (v10 / (((v10 << ((v10) & 63))) | 1))
                        var v18 int8 = 0
                        if (165 > 255) {
                            var v19 int = v14
                            var v20 uint32 = 1
                            _ = v19
                            _ = v20
                        } else {
                            if (!false) {
                                v21 := GenS1{F0: uint32(v3), F1: v1}
                                v13 += v1
                                v22 := GenS2{F0: (v18 & (v18 | v18)), F1: -1, F2: ""}
                                _ = v21
                                _ = v22
                            }
                            if v2 {
                                fmt.Println(v18, v6)
                                if true {
                                    if v2 {
                                        fmt.Println(v14)
                                    }
                                    if (!true) {
                                        var v23 int = (fi5 * (v14 / ((v14) | 1)))
                                        if ((-128 != int8(v10)) || ((v6 % ((32767) | 1)) > (fi7 & 42))) {
                                            var v24 int16 = 502
                                            var v25 int16 = 0
                                            _ = v24
                                            _ = v25
                                        }
                                        _ = v23
                                    } else {
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    }
                                    if v2 {
                                        if false {
                                            v13 += int64(v10)
                                            if false {
                                                fmt.Println(v8)
                                            }
                                        } else {
                                            switch "\n\t" {
                                            case "a":
                                                var v26 int16 = -32768
                                                fmt.Println(v3)
                                                var v27 int32 = 1
                                                switch v12 {
                                                case "hello":
                                                    if (42 <= 2) {
                                                        switch -9223372036854775808 {
                                                        case 856:
                                                            var v28 uint64 = 2
                                                            _ = v28
                                                        case 970:
                                                            v6 += v14
                                                        case -576:
                                                            var v29 int16 = 0
                                                            _ = v29
                                                        default:
                                                            if true {
                                                                v30 := GenS3{F0: (v3 * (^(v3)))}
                                                                if (18446744073709551615 > 4294967295) {
                                                                    v31 := GenS4{F0: v13, F1: v26, F2: v12}
                                                                    _ = v31
                                                                }
                                                                if v2 {
                                                                    var v32 uint64 = 807
                                                                    _ = v32
                                                                } else {
                                                                    v33 := GenS1{F0: 4294967295, F1: (v1 - v13)}
                                                                    var v34 int = (v14 % ((9223372036854775807) | 1))
                                                                    switch fi5 {
                                                                    case 9223372036854775807:
                                                                        switch (^(v34)) {
                                                                        case 695:
                                                                            v35 := GenS4{F0: (^(v13)), F1: (^(v26)), F2: "a"}
                                                                            var v36 int8 = v18
                                                                            _ = v35
                                                                            _ = v36
                                                                        default:
                                                                            fmt.Println(v1, v3, v12)
                                                                            v3 += v3
                                                                            if v2 {
                                                                                if v2 {
                                                                                    v37 := GenS4{F0: (v1 + (-(v13))), F1: v26, F2: v12}
                                                                                    v9 = append(v9, (fi7 >> ((v10) & 63)))
                                                                                    v1 += (v1 << ((v10) & 63))
                                                                                    _ = v37
                                                                                } else {
                                                                                    if len(v9) >= 2 {
                                                                                        v9 = v9[1:len(v9)]
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        v38 := GenS4{F0: 42, F1: v26, F2: ((((("a" + v12) + v12) + (("foo bar" + v12) + (("a" + v12) + ("foo bar" + ("hello" + "\n\t"))))) + (v12 + v12)) + "a")}
                                                                        _ = v38
                                                                    }
                                                                    var v39 uint64 = 0
                                                                    switch ((("\n\t" + (("foo bar" + ("" + v12)) + "")) + v12) + v12) {
                                                                    case "\n\t":
                                                                        v2 = (v2 || (v13 >= 9223372036854775807))
                                                                    case "hello":
                                                                        var v40 float32 = float32(3.14159)
                                                                        _ = v40
                                                                    default:
                                                                        v39 += v39
                                                                    }
                                                                    _ = v33
                                                                    _ = v34
                                                                    _ = v39
                                                                }
                                                                var v41 uint8 = 191
                                                                var v42 int32 = (v27 - (-(v27)))
                                                                _ = v30
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                        }
                                                    }
                                                case "":
                                                    fmt.Println(v14, v12)
                                                case "\n\t":
                                                    if true {
                                                        if (0.1 == 1.5) {
                                                            fmt.Println(v13, v10)
                                                        }
                                                        var v43 int = fi7
                                                        v43 -= -1
                                                        fmt.Println(v2, fi5)
                                                        _ = v43
                                                    }
                                                }
                                                v44 := GenS3{F0: v3}
                                                fmt.Println(v1, v3)
                                                var v45 int8 = (v18 - v18)
                                                var v46 uint32 = 2147483647
                                                _ = v26
                                                _ = v27
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            case "\n\t":
                                                fmt.Println(v14, v14)
                                            default:
                                                if true {
                                                    var v47 int8 = v18
                                                    switch v12 {
                                                    case "hello":
                                                        v6 = (-(v6))
                                                    }
                                                    if v2 {
                                                        var v48 int8 = v18
                                                        v2 = (v8 >= v8)
                                                        var v49 int = v6
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        var v50 int32 = 1
                                                        var v51 uint32 = 2147483648
                                                        v9 = append(v9, fi5)
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    _ = v47
                                                }
                                            }
                                            v8 += 0.0
                                            switch fi7 {
                                            case 42:
                                                v3 += v3
                                            case 32767:
                                                v2 = v2
                                            default:
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                            }
                                            v6 += v6
                                            var v52 bool = (v8 == float32(1.5))
                                            _ = v52
                                        }
                                        v53 := GenS5{F0: v18, F1: 4294967295, F2: v12}
                                        var v54 uint64 = 1
                                        _ = v53
                                        _ = v54
                                    }
                                    if false {
                                        if len(v9) > 0 {
                                            v9[0] = 32768
                                        }
                                        var v55 int16 = -32768
                                        switch (v12 + ((("" + v12) + v12) + (v12 + v12))) {
                                        case "hello":
                                            switch v6 {
                                            case -1:
                                                var v56 int32 = 0
                                                _ = v56
                                            }
                                        case "a":
                                            v12 = v12
                                        case "foo bar":
                                            var v57 int = v17
                                            _ = v57
                                        }
                                        _ = v55
                                    } else {
                                        var v58 int8 = (v18 / ((-128) | 1))
                                        _ = v58
                                    }
                                    if len(v9) > 0 {
                                        v9[0] = 32768
                                    }
                                    switch (v17 + -382) {
                                    case 495:
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    case -9223372036854775808:
                                        if (true && (!true)) {
                                            if (v2 || (2147483648 >= 262)) {
                                                if v2 {
                                                    v18 += (v18 % (((^(v18))) | 1))
                                                    if (!(42 == -516)) {
                                                        v1 = 0
                                                        v13 += 0
                                                        v17 = -2147483648
                                                    }
                                                    v59 := GenS2{F0: v18, F1: -1, F2: ((v12 + "\n\t") + (v12 + ("a" + "foo bar")))}
                                                    var v60 uint8 = 127
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    var v61 int16 = -1
                                                    _ = v61
                                                }
                                                if (fi5 > fi7) {
                                                    switch (fi5 >> (((v10 ^ (v10 ^ v10))) & 63)) {
                                                    case -9223372036854775808:
                                                        if (!(!(0.0 < 3.14159))) {
                                                            v9 = append(v9, (fi7 / (((-(v17))) | 1)))
                                                            if ((1 <= -382) || (!v2)) {
                                                                v62 := GenS6{F0: v12, F1: 1}
                                                                v13 += 1
                                                                _ = v62
                                                            } else {
                                                                if len(v9) >= 2 {
                                                                    v9 = v9[1:len(v9)]
                                                                }
                                                                v8 = (v8 / v8)
                                                                v14 = -9223372036854775808
                                                                var v63 float64 = 0.1
                                                                var v64 uint16 = v3
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                            switch v12 {
                                                            case "foo bar":
                                                                v9 = append(v9, fi5)
                                                            case "hello":
                                                                var v65 int64 = v13
                                                                _ = v65
                                                            }
                                                        }
                                                    case 128:
                                                        switch v12 {
                                                        case "hello":
                                                            fmt.Println(v1)
                                                            v3 = v3
                                                        case "a":
                                                            fmt.Println(v14, v2, v17)
                                                        case "\n\t":
                                                            v66 := GenS7{F0: v3}
                                                            _ = v66
                                                        default:
                                                            if false {
                                                                _ = len(v9)
                                                                var v67 uint32 = 622
                                                                switch int(v1) {
                                                                case -1:
                                                                    var v68 float32 = v8
                                                                    _ = v68
                                                                case -9223372036854775808:
                                                                    if (!(!v2)) {
                                                                        v69 := GenS8{F0: fi5}
                                                                        if len(v9) > 0 {
                                                                            v9[0] = (fi7 - (fi7 - (-(v6))))
                                                                        }
                                                                        _ = v69
                                                                    } else {
                                                                        if false {
                                                                            v70 := GenS4{F0: (v1 >> ((v10) & 63)), F1: int16(v3), F2: v12}
                                                                            var v71 uint64 = 1
                                                                            _ = v70
                                                                            _ = v71
                                                                        } else {
                                                                            if (v10 == v10) {
                                                                                var v72 int8 = v18
                                                                                if v2 {
                                                                                    var v73 uint64 = 337
                                                                                    fmt.Println(v67)
                                                                                    fmt.Println(v2, v72, v3)
                                                                                    v10 = (v10 - v10)
                                                                                    _ = v73
                                                                                }
                                                                                if (((v1 / ((v1) | 1)) == (v13 % ((v13) | 1))) && (((float32(1.5) * v8) - v8) != v8)) {
                                                                                    v1 -= v1
                                                                                    var v74 float32 = v8
                                                                                    _ = v74
                                                                                }
                                                                                _ = v72
                                                                            } else {
                                                                                fmt.Println(v6)
                                                                                var v75 bool = v2
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                    }
                                                                case 51:
                                                                    v76 := GenS1{F0: v67, F1: v13}
                                                                    _ = v76
                                                                }
                                                                v77 := GenS9{F0: -1}
                                                                _ = v67
                                                                _ = v77
                                                            } else {
                                                                var v78 uint32 = 2147483647
                                                                _ = v78
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v1 = v13
                                                    _ = len(v9)
                                                    v14 = (fi5 << ((v10) & 63))
                                                }
                                                v1 -= -1
                                            }
                                        } else {
                                            var v79 uint32 = 0
                                            if (1 > -691) {
                                                v6 = v17
                                                v80 := GenS4{F0: (v13 / ((v1) | 1)), F1: -1, F2: v12}
                                                v81 := GenS5{F0: v18, F1: v79, F2: v12}
                                                _ = v80
                                                _ = v81
                                            }
                                            v18 = (v18 / ((v18) | 1))
                                            _ = v79
                                        }
                                    case 291:
                                        v17 -= v17
                                    default:
                                        v82 := GenS10{F0: -926, F1: uint8(v10), F2: fi5}
                                        _ = v82
                                    }
                                } else {
                                    switch v14 {
                                    case -1:
                                        if v2 {
                                            v83 := GenS11{F0: 319}
                                            _ = v83
                                        } else {
                                            v84 := GenS3{F0: 625}
                                            _ = v84
                                        }
                                    default:
                                        v85 := GenS5{F0: (v18 & (v18 | (-(v18)))), F1: 840, F2: (v12 + v12)}
                                        _ = v85
                                    }
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                    var v86 int64 = (v1 >> (((v10 * uint(v3))) & 63))
                                    v3 = 32767
                                    _ = v86
                                }
                                v87 := GenS12{F0: 2}
                                _ = v87
                            }
                            v6 = v14
                            fmt.Println(fi7, v14)
                            v18 = -62
                        }
                        _ = v18
                    } else {
                        var v88 int8 = int8(v1)
                        var v89 uint8 = 32
                        fmt.Println(v1, fi5)
                        _ = v88
                        _ = v89
                    }
                case "":
                    var v90 float32 = 0.1
                    v1 += (v1 / ((9223372036854775807) | 1))
                    _ = v90
                case "\n\t":
                    if len(v9) > 0 {
                        v9[0] = fi5
                    }
                }
                _ = fi7
                _ = v10
                _ = v12
                _ = v13
                _ = v14
                _ = v17
                _ = v8
                _ = v9
            }
            _ = v6
        }
        v91 := GenS12{F0: 137}
        v3 = (v3 / ((v3) | 1))
        _ = fi5
        _ = v91
    }
    for fi92 := 0; fi92 < 12; fi92++ {
        v93 := []int{(fi92 - fi92), fi92, fi92, (-(fi92))}
        v93 = append(v93, (fi92 ^ (fi92 % ((-895) | 1))))
        _ = fi92
        _ = v93
    }
    v3 = (^(v3))
    v3 += 934
    var v94 bool = v2
    for fi95 := 0; fi95 < 11; fi95++ {
        switch "a" {
        case "hello":
            var v96 uint8 = 145
            _ = v96
        case "foo bar":
            v97 := GenS1{F0: 2147483648, F1: (v1 * (v1 | v1))}
            _ = v97
        }
        for fi98 := 0; fi98 < 12; fi98++ {
            if (345 > 65535) {
                switch fi98 {
                case 1:
                    var v99 uint32 = 65535
                    _ = v99
                default:
                    v100 := GenS5{F0: 100, F1: 631, F2: ((((("hello" + ("\n\t" + "a")) + "\n\t") + "hello") + "\n\t") + "a")}
                    _ = v100
                }
                fmt.Println(v3)
                v1 = 1
            }
            v101 := []int{(fi98 ^ -914)}
            _ = fi98
            _ = v101
        }
        _ = fi95
    }
    v102 := GenS6{F0: "", F1: 0}
    v103 := GenS13{F0: "hello"}
    fmt.Println(v94, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v94
}
