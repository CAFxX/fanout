package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int; F1 uint32; F2 int }
type GenS6 struct { F0 int8; F1 uint16; F2 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 uint; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    var v4 uint = 65535
    var v5 int16 = -208
    v6 := GenS1{F0: 127, F1: 130}
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{0, -1}
            _ = fi8
            _ = v9
        }
        switch 32767 {
        case 4294967295:
            v10 := GenS1{F0: 128, F1: 0}
            _ = v10
        default:
            switch (fi7 % (((fi7 << ((v4) & 63))) | 1)) {
            case 128:
                v4 += 0
            case 0:
                var v11 int8 = 34
                _ = v11
            default:
                var v12 uint32 = 2147483647
                _ = v12
            }
        }
        _ = fi7
    }
    v5 = v5
    var v13 uint64 = 1
    switch "" {
    case "\n\t":
        switch 1 {
        case -584:
            var v14 float64 = 0.0
            _ = v14
        case 1:
            v15 := []int{128, 42, 129, -832}
            _ = v15
        default:
            v5 = 722
        }
    default:
        v16 := []int{933, 42}
        _ = v16
    }
    fmt.Println(v4)
    v17 := GenS1{F0: 255, F1: 57}
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS1{F0: 125, F1: 127}
        switch "foo bar" {
        case "foo bar":
            v20 := GenS1{F0: 0, F1: 0}
            _ = v20
        case "":
            switch "\n\t" {
            case "":
                fmt.Println(v4, v5)
            case "foo bar":
                var v21 int = (fi18 + (fi18 + (-(fi18))))
                _ = v21
            case "hello":
                if (128 < 2) {
                    if true {
                        fmt.Println(v13)
                        v5 -= v5
                    }
                }
            default:
                switch fi18 {
                case 122:
                    for fi22 := 0; fi22 < 14; fi22++ {
                        if (2 > 255) {
                            if true {
                                v23 := GenS2{F0: true}
                                _ = v23
                            }
                            switch -9223372036854775808 {
                            case -686:
                                if false {
                                    v4 = 2
                                    if false {
                                        switch "" {
                                        case "a":
                                            switch "hello" {
                                            case "":
                                                fmt.Println(fi18)
                                            }
                                        default:
                                            if false {
                                                var v24 uint64 = 204
                                                v25 := GenS3{F0: -899, F1: (fi22 >> (((^(v4))) & 63))}
                                                fmt.Println(v24)
                                                _ = v24
                                                _ = v25
                                            } else {
                                                v26 := GenS4{F0: 42, F1: (fi22 <= 42)}
                                                var v27 int64 = -9223372036854775808
                                                v28 := GenS5{F0: (fi18 * fi18), F1: 2, F2: 42}
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                        var v29 int8 = -80
                                        v30 := []int{-253, (fi22 >> ((184) & 63)), fi22}
                                        _ = v29
                                        _ = v30
                                    }
                                    if ("a" == "foo bar") {
                                        v4 += 2
                                        v4 -= v4
                                    } else {
                                        v31 := GenS5{F0: (fi18 + (fi18 >> (((v4 << ((2) & 63))) & 63))), F1: 4294967295, F2: (fi22 | (fi22 % (((fi18 - fi18)) | 1)))}
                                        var v32 int32 = -696
                                        _ = v31
                                        _ = v32
                                    }
                                    fmt.Println(v5, v4)
                                } else {
                                    v1[65535] = 65535
                                    var v33 int32 = -2147483648
                                    v4 = (^(v4))
                                    _ = v33
                                }
                            case 584:
                                v4 = v4
                            case 4294967295:
                                v34 := []int{(fi18 ^ (fi18 % ((fi22) | 1))), (^(fi18))}
                                switch (-(fi22)) {
                                case 4294967295:
                                    var v35 float64 = (1e10 + 1.5)
                                    _ = v35
                                case 127:
                                    var v36 int16 = (-(v5))
                                    if (1.5 == -2.25) {
                                        var v37 int16 = v36
                                        v38 := GenS1{F0: 127, F1: 0}
                                        _ = v37
                                        _ = v38
                                    }
                                    if (float64(0.0) <= (-2.25 - 1.5)) {
                                        fmt.Println(v36, v4, v36)
                                        var v39 float32 = 1e10
                                        v4 += v4
                                        var v40 uint32 = 121
                                        var v41 int32 = -702
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    } else {
                                        var v42 uint = v4
                                        var v43 uint16 = 42
                                        if (v43 > (v43 + (v43 / ((v43) | 1)))) {
                                            var v44 uint64 = 42
                                            var v45 float32 = -2.25
                                            if true {
                                                if ("hello" >= "a") {
                                                    if (0 >= 2147483647) {
                                                        var v46 uint = (v4 * v4)
                                                        _ = v46
                                                    } else {
                                                        if false {
                                                            _ = len(v2)
                                                            _ = len(v34)
                                                            if len(v34) >= 2 {
                                                                v34 = v34[1:len(v34)]
                                                            }
                                                        } else {
                                                            if (true && (63 < 0)) {
                                                                v47 := v2["hello"]
                                                                _ = v47
                                                            } else {
                                                                switch ("foo bar" + ("" + "\n\t")) {
                                                                case "\n\t":
                                                                    v48 := GenS6{F0: -116, F1: (^(v43)), F2: fi18}
                                                                    _ = v48
                                                                }
                                                                _ = len(v34)
                                                            }
                                                            var v49 float64 = 3.14159
                                                            v2["\n\t"] = "a"
                                                            _ = v49
                                                        }
                                                        var v50 float32 = ((float32(float64(v45)) / (v45 + v45)) - (0.0 * (float32(3.14159) + v45)))
                                                        delete(v1, 42)
                                                        fmt.Println(v5, v43)
                                                        v51 := GenS2{F0: ((v5 * (^(v5))) != -1)}
                                                        if (-54 != 101) {
                                                            var v52 uint8 = 2
                                                            _ = v52
                                                        }
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                } else {
                                                    v43 = (v43 % ((v43) | 1))
                                                }
                                            } else {
                                                var v53 float64 = 0.0
                                                fmt.Println(v13, v43, fi18)
                                                v45 += v45
                                                _ = v53
                                            }
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = v36
                                case 2147483647:
                                    v54 := GenS5{F0: -2147483648, F1: 1, F2: (^(fi18))}
                                    _ = v54
                                default:
                                    var v55 uint16 = 1
                                    _ = v55
                                }
                                v13 -= 4294967295
                                _ = v34
                            default:
                                v4 = (v4 + 65535)
                            }
                            v13 -= (v13 - v13)
                        }
                        _ = fi22
                    }
                default:
                    _ = len(v2)
                }
            }
        }
        for fi56 := 0; fi56 < 10; fi56++ {
            fmt.Println(fi56, v4)
            switch "" {
            case "foo bar":
                switch (-(fi18)) {
                case 9223372036854775807:
                    fmt.Println(fi18)
                case 127:
                    switch fi18 {
                    case 32768:
                        if false {
                            v5 += 42
                            var v57 uint = 0
                            _ = v57
                        } else {
                            v58 := GenS4{F0: -1, F1: (-806 <= (^(v5)))}
                            var v59 int = 19
                            _ = v58
                            _ = v59
                        }
                    case 32767:
                        v60 := GenS5{F0: fi56, F1: 1, F2: fi56}
                        _ = v60
                    case 42:
                        v5 -= -1
                    default:
                        if false {
                            var v61 float32 = 1e10
                            v62 := []int{fi56, (fi18 ^ fi56), fi18, fi18}
                            if (2 <= 253) {
                                if len(v62) > 0 {
                                    v62[0] = (fi18 / (((^(fi18))) | 1))
                                }
                            } else {
                                v2["a"] = "a"
                                var v63 int32 = 2147483647
                                var v64 float64 = float64(0.1)
                                _ = len(v62)
                                var v65 uint8 = 151
                                var v66 uint32 = 258
                                v1[v66] = (^(v66))
                                if false {
                                    var v67 uint64 = (v13 + 4294967295)
                                    fmt.Println(v65)
                                    var v68 int8 = 127
                                    _ = v67
                                    _ = v68
                                }
                                v64 -= float64(v61)
                                var v69 float32 = v61
                                _ = v63
                                _ = v64
                                _ = v65
                                _ = v66
                                _ = v69
                            }
                            fmt.Println(v5)
                            switch "hello" {
                            case "foo bar":
                                v70 := GenS4{F0: -9223372036854775808, F1: (v5 <= (-(v5)))}
                                _ = v70
                            case "hello":
                                fmt.Println(v13, v5)
                            default:
                                var v71 int64 = 9223372036854775807
                                _ = v71
                            }
                            var v72 int8 = 127
                            _ = len(v2)
                            if (1.5 <= (1e10 - (0.0 / ((1.5 + 0.0) - 3.14159)))) {
                                v73 := GenS1{F0: 2, F1: 92}
                                if true {
                                    if ((!false) && (!(false && (!(1.5 <= 0.1))))) {
                                        _ = len(v2)
                                        v4 -= v4
                                        var v74 float64 = 1e10
                                        _ = v74
                                    }
                                    if false {
                                        switch "" {
                                        case "":
                                            if false {
                                                var v75 uint = v4
                                                _ = v75
                                            }
                                        case "a":
                                            var v76 bool = (!((^(fi56)) < (^(fi56))))
                                            var v77 string = "\n\t"
                                            _ = v76
                                            _ = v77
                                        default:
                                            var v78 int8 = 15
                                            _ = v78
                                        }
                                    } else {
                                        v72 = (v72 / ((-55) | 1))
                                    }
                                } else {
                                    var v79 int32 = -284
                                    if (1 <= -1) {
                                        var v80 int32 = v79
                                        _ = v80
                                    }
                                    _ = v79
                                }
                                switch "" {
                                case "\n\t":
                                    fmt.Println(fi18, fi56, v5)
                                case "":
                                    v5 = v5
                                case "hello":
                                    _ = len(v1)
                                default:
                                    var v81 uint16 = 103
                                    _ = v81
                                }
                                fmt.Println(v5)
                                v82 := GenS7{F0: 42}
                                var v83 float64 = 1e10
                                _ = v73
                                _ = v82
                                _ = v83
                            } else {
                                if true {
                                    delete(v1, 4294967295)
                                    if len(v62) >= 2 {
                                        v62 = v62[1:len(v62)]
                                    }
                                    var v84 uint64 = 4294967295
                                    _ = v84
                                } else {
                                    v72 -= (v72 ^ -74)
                                    v85 := GenS6{F0: (v72 / ((42) | 1)), F1: 437, F2: (fi56 % ((1) | 1))}
                                    v72 -= v72
                                    _ = v85
                                }
                                fmt.Println(v4)
                            }
                            v86 := GenS2{F0: false}
                            _ = v61
                            _ = v62
                            _ = v72
                            _ = v86
                        } else {
                            v87 := GenS4{F0: 9223372036854775807, F1: (uint64(v4) >= 18446744073709551615)}
                            var v88 int64 = 1
                            _ = v87
                            _ = v88
                        }
                    }
                case 128:
                    if true {
                        switch ("foo bar" + "foo bar") {
                        case "":
                            v13 = v13
                        case "\n\t":
                            if false {
                                var v89 int64 = 779
                                _ = v89
                            }
                        case "hello":
                            var v90 uint8 = 2
                            _ = v90
                        default:
                            if ((v4 == 340) || ((((v5 != (v5 >> (((v4 + v4)) & 63))) && ((0.0 + (3.14159 * 1.5)) == (0.0 / 1.0))) && true) && false)) {
                                switch int(v4) {
                                case 2147483647:
                                    fmt.Println(fi56, v5)
                                default:
                                    v13 = (v13 * v13)
                                }
                                v91 := GenS8{F0: int(v5), F1: (v4 & 896), F2: (v5 | v5)}
                                _ = v91
                            }
                            if (!true) {
                                var v92 uint64 = 246
                                fmt.Println(v5, v13, v13)
                                v13 = 0
                                if false {
                                    fmt.Println(v4, fi56)
                                    var v93 int8 = -128
                                    v3[65535] = ((^(fi18)) > fi56)
                                    _ = v93
                                }
                                _ = v92
                            }
                            v94 := []int{fi18}
                            _ = v94
                        }
                    }
                }
            case "":
                var v95 int64 = 0
                _ = v95
            default:
                if true {
                    var v96 int16 = v5
                    _ = len(v1)
                    fmt.Println(v13)
                    switch (fi18 / (((fi18 >> (((v4 + 291)) & 63))) | 1)) {
                    case 32768:
                        v13 = (^(v13))
                    case -715:
                        v97 := []int{-1, fi56, 42}
                        fmt.Println(fi18)
                        var v98 uint64 = 2
                        v99 := GenS4{F0: 1, F1: true}
                        _ = v97
                        _ = v98
                        _ = v99
                    case 2147483647:
                        var v100 float64 = 1e10
                        _ = v100
                    }
                    _ = v96
                }
            }
            v5 += v5
            v101 := GenS8{F0: (fi56 << ((v4) & 63)), F1: (v4 | (v4 * 2)), F2: (v5 * (v5 + (v5 % ((0) | 1))))}
            if ("\n\t" >= "hello") {
                if true {
                    _ = len(v1)
                }
            }
            _ = fi56
            _ = v101
        }
        _ = fi18
        _ = v19
    }
    if false {
        var v102 int16 = -32768
        v102 = 1
        _ = v102
    }
    var v103 uint64 = (v13 % (((v13 % (((v13 & v13)) | 1))) | 1))
    v13 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v103
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
