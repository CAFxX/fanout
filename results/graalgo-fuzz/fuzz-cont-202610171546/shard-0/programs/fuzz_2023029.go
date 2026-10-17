package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int32; F1 int32; F2 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 int16; F1 int8; F2 uint }
type GenS10 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS11 struct { F0 int16 }
func genFn0() bool {
    var v6 uint = 4294967295
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint8 = 1
        switch fi7 {
        case -9223372036854775808:
            v9 := GenS1{F0: 2147483647, F1: 1}
            _ = v9
        default:
            for fi10 := 0; fi10 < 14; fi10++ {
                fmt.Println(v8)
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    switch (fi10 & 1) {
                    case 32767:
                        if false {
                            var v11 uint = (v6 * v6)
                            fmt.Println(v6)
                            var v12 float64 = 1e10
                            v11 += v6
                            v13 := []int{fi7}
                            _ = v11
                            _ = v12
                            _ = v13
                        } else {
                            if (false && true) {
                                v14 := GenS1{F0: 1, F1: -913}
                                var v15 uint64 = uint64(v6)
                                v8 = (^(v8))
                                _ = v14
                                _ = v15
                            } else {
                                var v16 int16 = 146
                                v8 += (v8 % ((uint8(v6)) | 1))
                                _ = v16
                            }
                            v17 := GenS1{F0: 225, F1: 0}
                            _ = v17
                        }
                    case -90:
                        switch fi7 {
                        case 1:
                            var v18 float32 = 0.0
                            _ = v18
                        case 4294967295:
                            var v19 uint16 = 328
                            _ = v19
                        }
                    default:
                        v20 := []int{(fi10 + (fi10 ^ fi10)), (fi10 ^ fi7), 453}
                        _ = v20
                    }
                }
                var v21 int32 = -2147483648
                _ = fi10
                _ = v21
            }
        }
        _ = fi7
        _ = v8
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := []int{594}
        var v24 float32 = 0.0
        if (false || false) {
            var v25 uint32 = 2147483648
            var v26 bool = (true || (!((fi22 % (((fi22 % ((fi22) | 1))) | 1)) == 32767)))
            switch "hello" {
            case "\n\t":
                switch fi22 {
                case 127:
                    v25 = (v25 << (((v6 ^ 1)) & 63))
                case 128:
                    if v26 {
                        v23 = append(v23, fi22)
                    }
                default:
                    var v27 int8 = 127
                    _ = v27
                }
            default:
                v28 := GenS2{F0: 32767}
                _ = v28
            }
            _ = v25
            _ = v26
        } else {
            var v29 int16 = -32768
            if (713 < 442) {
                fmt.Println(v6, v29)
                var v30 int32 = 2147483647
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = v30
            } else {
                for fi31 := 0; fi31 < 3; fi31++ {
                    v6 = (v6 * 1)
                    v32 := GenS3{F0: 42, F1: 1}
                    _ = len(v23)
                    v33 := GenS4{F0: -13, F1: 384}
                    v23 = append(v23, 32768)
                    fmt.Println(v29)
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                switch ("\n\t" + ("\n\t" + "")) {
                case "\n\t":
                    switch "hello" {
                    case "foo bar":
                        if (0 != 2) {
                            if (-1 == 0) {
                                v6 += 160
                                if false {
                                    switch fi22 {
                                    case -2147483648:
                                        var v34 float32 = float32(3.14159)
                                        _ = v34
                                    case 4294967295:
                                        switch "" {
                                        case "\n\t":
                                            var v35 int = fi22
                                            _ = v35
                                        default:
                                            var v36 uint16 = 42
                                            _ = v36
                                        }
                                    default:
                                        if len(v23) > 0 {
                                            v23[0] = (-(fi22))
                                        }
                                    }
                                } else {
                                    switch fi22 {
                                    case 32767:
                                        v23 = append(v23, fi22)
                                    case 0:
                                        v6 -= 54
                                    case -2147483648:
                                        v37 := GenS5{F0: (v6 & v6), F1: 172, F2: 128}
                                        _ = v37
                                    }
                                    v38 := GenS6{F0: -1, F1: -2147483648, F2: (v6 * 1)}
                                    v39 := GenS1{F0: 2147483647, F1: -198}
                                    _ = v38
                                    _ = v39
                                }
                            } else {
                                for fi40 := 0; fi40 < 18; fi40++ {
                                    v29 = -1
                                    v6 -= 1
                                    v41 := GenS2{F0: 1000}
                                    switch (fi40 - (fi22 + fi40)) {
                                    case 0:
                                        v29 += v29
                                    case 32768:
                                        var v42 int = (fi40 | (fi22 ^ 9223372036854775807))
                                        _ = v42
                                    case 42:
                                        var v43 uint = v6
                                        var v44 string = "\n\t"
                                        var v45 int16 = v29
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    default:
                                        v23 = append(v23, (fi22 & (-(fi22))))
                                    }
                                    var v46 uint64 = 198
                                    v46 += v46
                                    _ = fi40
                                    _ = v41
                                    _ = v46
                                }
                            }
                            _ = len(v23)
                        } else {
                            _ = len(v23)
                            if len(v23) > 0 {
                                v23[0] = (fi22 - fi22)
                            }
                            v47 := GenS5{F0: v6, F1: 42, F2: 16}
                            _ = v47
                        }
                    }
                default:
                    _ = len(v23)
                }
            }
            _ = v29
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = v6
    return (!false)
}
func genFn1() uint {
    var v48 float32 = 3.14159
    if (32767 <= 836) {
        var v49 int = 127
        _ = v49
    } else {
        var v50 uint64 = 18446744073709551615
        _ = v50
    }
    fmt.Println(v48)
    var v51 uint32 = 0
    fmt.Println(v48, v51)
    _ = v48
    _ = v51
    return 663
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 0
    v3 := make(map[string]bool)
    v4 := []int{-191, 32768, -809}
    v5 := make(map[uint32]bool)
    if (32767 != 65535) {
        for fi52 := 0; fi52 < 14; fi52++ {
            var v53 float32 = (1e10 - (0.1 / 3.14159))
            if len(v4) > 0 {
                v4[0] = -1
            }
            _ = fi52
            _ = v53
        }
        switch "" {
        case "hello":
            switch 127 {
            case 0:
                switch "a" {
                case "a":
                    v54 := GenS7{F0: -1}
                    _ = v54
                case "":
                    v4 = append(v4, 2147483647)
                case "hello":
                    var v55 uint32 = 478
                    _ = v55
                default:
                    v1 -= (v1 / 1.0)
                }
            case 259:
                if true {
                    fmt.Println(v1)
                } else {
                    delete(v5, 2147483647)
                }
            }
        }
        for fi56 := 0; fi56 < 11; fi56++ {
            var v57 int64 = 1
            if true {
                switch fi56 {
                case 32768:
                    _ = len(v4)
                case 32767:
                    switch "\n\t" {
                    case "\n\t":
                        for fi58 := 0; fi58 < 13; fi58++ {
                            var v59 bool = (!(!(1000 >= -328)))
                            var v60 string = "foo bar"
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                    case "foo bar":
                        fmt.Println(v2)
                    case "a":
                        var v61 int8 = -41
                        _ = v61
                    default:
                        switch 42 {
                        case 127:
                            for fi62 := 0; fi62 < 19; fi62++ {
                                var v63 float64 = 1.5
                                _ = fi62
                                _ = v63
                            }
                        case -9223372036854775808:
                            if true {
                                var v64 float64 = v1
                                if len(v4) > 0 {
                                    v4[0] = 481
                                }
                                if false {
                                    var v65 int16 = 1
                                    for fi66 := 0; fi66 < 18; fi66++ {
                                        if false {
                                            v2 = 385
                                        } else {
                                            fmt.Println(v2, fi66)
                                        }
                                        if ((1e10 >= 1e10) || (32767 > 0)) {
                                            v67 := GenS6{F0: 2147483647, F1: int32(fi56), F2: 42}
                                            var v68 int32 = 2147483647
                                            v57 += (v57 & v57)
                                            var v69 uint16 = 2
                                            if true {
                                                fmt.Println(fi56)
                                                v68 += 42
                                            }
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        }
                                        fmt.Println(v64)
                                        _ = fi66
                                    }
                                    _ = v65
                                }
                                _ = v64
                            } else {
                                switch ("foo bar" + "hello") {
                                case "":
                                    v70 := GenS1{F0: 2147483647, F1: -1}
                                    _ = v70
                                default:
                                    var v71 uint64 = 4294967295
                                    _ = v71
                                }
                            }
                        default:
                            v72 := GenS2{F0: 980}
                            _ = v72
                        }
                    }
                case -645:
                    var v73 int64 = 42
                    _ = v73
                }
            }
            fmt.Println(v57)
            _ = fi56
            _ = v57
        }
    }
    if true {
        fmt.Println(v1, v2)
        var v74 int32 = 42
        for fi75 := 0; fi75 < 13; fi75++ {
            fmt.Println(v2)
            v76 := GenS3{F0: -1, F1: 42}
            if (((!false) || true) && (0 < 1)) {
                if (127 > 2) {
                    delete(v3, "")
                    var v77 int16 = 0
                    v74 += (v74 - 42)
                    _ = v77
                } else {
                    if false {
                        var v78 string = "foo bar"
                        _ = v78
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v79 := GenS5{F0: v2, F1: 882, F2: 128}
                        var v80 uint64 = 285
                        _ = v79
                        _ = v80
                    }
                    v81 := v3["hello"]
                    _ = len(v5)
                    _ = v81
                }
                if (1 > 107) {
                    var v82 int = (fi75 & -2147483648)
                    v4 = append(v4, (fi75 + v82))
                    _ = v82
                }
            }
            _ = fi75
            _ = v76
        }
        _ = v74
    }
    delete(v5, 4294967295)
    v83 := GenS8{F0: 0, F1: "hello"}
    for fi84 := 0; fi84 < 15; fi84++ {
        fmt.Println(v1)
        v2 += (v2 - v2)
        if (-2147483648 <= 0) {
            if len(v4) > 0 {
                v4[0] = int(v2)
            }
        }
        _ = fi84
    }
    var v85 uint = v2
    v86 := GenS3{F0: -9223372036854775808, F1: 44}
    v87 := GenS9{F0: 1, F1: 127, F2: (v85 - (v2 * (v2 - v85)))}
    v88 := GenS10{F0: 4294967295, F1: 127, F2: 896}
    if ("hello" > ("" + ("" + "\n\t"))) {
        if true {
            var v89 uint32 = 413
            v89 += (v89 ^ (^(v89)))
            switch 127 {
            case 32768:
                v2 -= (^(v85))
            case 42:
                fmt.Println(v85, v1, v2)
            default:
                v90 := GenS6{F0: 1, F1: 0, F2: (v2 & v85)}
                _ = v90
            }
            _ = v89
        }
        v91 := GenS11{F0: 1000}
        var v92 string = ("" + ("\n\t" + ""))
        _ = v91
        _ = v92
    } else {
        switch 5 {
        case 567:
            for fi93 := 0; fi93 < 18; fi93++ {
                fmt.Println(v1, v2, v1)
                var v94 uint32 = 968
                _ = fi93
                _ = v94
            }
        case 32768:
            for fi95 := 0; fi95 < 9; fi95++ {
                for fi96 := 0; fi96 < 5; fi96++ {
                    if (!true) {
                        if true {
                            fmt.Println(v1, v85, fi95)
                            var v97 float64 = ((v1 + float64(3.14159)) * v1)
                            _ = v97
                        }
                        v4 = append(v4, (^(fi95)))
                    } else {
                        var v98 uint32 = 1
                        _ = v98
                    }
                    _ = fi96
                }
                var v99 uint64 = 2
                _ = fi95
                _ = v99
            }
        default:
            var v100 uint64 = 18446744073709551615
            _ = v100
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v85
    _ = v86
    _ = v87
    _ = v88
}
