package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int16; F2 int16 }
type GenS4 struct { F0 string; F1 uint16; F2 int32 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 string; F1 uint; F2 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16; F1 int64 }
func genFn0() uint {
    var v6 int64 = -9223372036854775808
    v7 := []int{32768, 32768, -222}
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{912, 2}
    var v2 string = ("hello" + "hello")
    var v3 float32 = 1e10
    var v4 float32 = (1.5 / float32(((1e10 + -2.25) * 0.1)))
    v5 := []uint32{2147483648}
    var v8 int16 = 1000
    v9 := GenS1{F0: 2147483647, F1: 403}
    v1 = append(v1, 806)
    v4 += 0.1
    if ((0.1 + -2.25) <= (0.1 / (0.0 + -2.25))) {
        v10 := GenS2{F0: 32768}
        v8 -= (v8 + (v8 & (^(v8))))
        _ = v10
    } else {
        switch 32768 {
        case -2147483648:
            v8 = v8
        }
    }
    v8 = -32768
    v11 := GenS3{F0: 1, F1: v8, F2: (v8 / ((v8) | 1))}
    var v12 uint32 = 0
    var v13 float64 = 3.14159
    fmt.Println(v2)
    switch "hello" {
    case "a":
        for fi14 := 0; fi14 < 12; fi14++ {
            if (100 != 100) {
                switch "\n\t" {
                case "foo bar":
                    v15 := GenS3{F0: 18446744073709551615, F1: 1000, F2: v8}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v5 = append(v5, v12)
                        v17 := GenS4{F0: "a", F1: 1, F2: 2147483647}
                        if len(v1) > 0 {
                            v1[0] = v12
                        }
                        var v18 uint64 = 1
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                v19 := GenS5{F0: "", F1: 152}
                _ = v19
            } else {
                if false {
                    var v20 string = v2
                    for fi21 := 0; fi21 < 7; fi21++ {
                        if (!(true && true)) {
                            v22 := GenS3{F0: 476, F1: v8, F2: v8}
                            if len(v5) > 0 {
                                v5[0] = (v12 * v12)
                            }
                            _ = v22
                        } else {
                            var v23 float64 = float64(0.1)
                            v5 = append(v5, 2147483647)
                            _ = v23
                        }
                        if false {
                            v24 := GenS2{F0: (^(fi21))}
                            var v25 uint8 = 255
                            var v26 uint16 = 255
                            v25 += 2
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        switch "\n\t" {
                        case "hello":
                            switch (^(fi14)) {
                            case 32767:
                                if false {
                                    fmt.Println(v8, v20, fi14)
                                    switch (v20 + "") {
                                    case "hello":
                                        v20 = (v20 + (("hello" + ("\n\t" + v20)) + "a"))
                                    case "":
                                        var v27 uint8 = 0
                                        _ = v27
                                    case "\n\t":
                                        var v28 float64 = float64(((-2.25 * v4) + ((float32(v13) - ((float32((3.14159 + 1e10)) / (0.0 - v4)) + v3)) + float32(v13))))
                                        _ = v28
                                    default:
                                        if ((^(v12)) > v12) {
                                            if false {
                                                v3 -= (float32(v13) / 1.5)
                                                var v29 int32 = 2147483647
                                                fmt.Println(v3, v20, fi21)
                                                v3 += v4
                                                var v30 int64 = 42
                                                _ = v29
                                                _ = v30
                                            } else {
                                                v31 := GenS2{F0: fi21}
                                                _ = v31
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 189
                                            }
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            var v32 string = v2
                                            _ = v32
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = (^(v12))
                                            }
                                        }
                                    }
                                    var v33 uint64 = 0
                                    _ = v33
                                } else {
                                    _ = len(v5)
                                    var v34 uint16 = 0
                                    v13 = (1e10 * (v13 * 1.5))
                                    _ = v34
                                }
                            case -2147483648:
                                var v35 uint8 = 2
                                _ = v35
                            case 128:
                                v8 += int16(v12)
                            default:
                                var v36 int32 = 1
                                _ = v36
                            }
                        case "":
                            var v37 int64 = 0
                            _ = v37
                        }
                        _ = fi21
                    }
                    v20 = "a"
                    _ = v20
                } else {
                    fmt.Println(fi14)
                }
            }
            var v38 float32 = ((3.14159 + float32(v13)) / v3)
            _ = fi14
            _ = v38
        }
    case "hello":
        for fi39 := 0; fi39 < 5; fi39++ {
            if true {
                v8 -= -1
                for fi40 := 0; fi40 < 12; fi40++ {
                    var v41 int = fi39
                    switch 1 {
                    case 707:
                        if (!true) {
                            if true {
                                var v42 uint16 = 1
                                var v43 int32 = 0
                                _ = v42
                                _ = v43
                            } else {
                                v1 = append(v1, v12)
                                v2 = v2
                            }
                            fmt.Println(v13)
                            var v44 float64 = float64(float32(v13))
                            _ = v44
                        } else {
                            var v45 uint32 = (v12 * v12)
                            _ = v45
                        }
                    case 2147483647:
                        fmt.Println(v4)
                    case -160:
                        var v46 uint32 = (^(v12))
                        _ = v46
                    }
                    if false {
                        if false {
                            fmt.Println(fi40, fi39, v12)
                            v47 := GenS1{F0: 42, F1: int16(fi40)}
                            var v48 float32 = v3
                            v5 = append(v5, (v12 << ((963) & 63)))
                            _ = v47
                            _ = v48
                        } else {
                            switch (v2 + "\n\t") {
                            case "":
                                fmt.Println(v12, v8, v13)
                                if (true && true) {
                                    var v49 int16 = v8
                                    var v50 float32 = (v3 * v3)
                                    fmt.Println(fi40, v8)
                                    _ = v49
                                    _ = v50
                                } else {
                                    var v51 uint32 = (v12 / (((v12 + (v12 % ((6) | 1)))) | 1))
                                    v52 := GenS3{F0: 133, F1: v8, F2: 422}
                                    _ = v51
                                    _ = v52
                                }
                            case "\n\t":
                                switch v2 {
                                case "hello":
                                    var v53 int32 = 681
                                    _ = v53
                                default:
                                    v8 = v8
                                }
                            }
                            v4 = 1e10
                            fmt.Println(v41, v4)
                        }
                        switch v41 {
                        case -581:
                            switch ("a" + v2) {
                            case "a":
                                v2 = ""
                            }
                        case 2147483647:
                            var v54 float64 = (v13 * ((-2.25 / ((1e10 - ((((v13 * 1e10) / -2.25) + 1e10) / float64(((0.1 + 3.14159) / (0.0 * float32(v13)))))) - v13)) + (v13 - 3.14159)))
                            _ = v54
                        case 9:
                            switch "" {
                            case "":
                                if (0 == 42) {
                                    switch fi39 {
                                    case 32768:
                                        switch (fi40 + 9223372036854775807) {
                                        case 2147483647:
                                            v12 += 4294967295
                                        default:
                                            var v55 int8 = 1
                                            _ = v55
                                        }
                                    case -9223372036854775808:
                                        v56 := GenS3{F0: 646, F1: (v8 % ((42) | 1)), F2: 32767}
                                        _ = v56
                                    default:
                                        v57 := GenS6{F0: (!((float64(v4) / v13) != v13)), F1: 255}
                                        _ = v57
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = v12
                                    }
                                    v1 = append(v1, (v12 + v12))
                                    if (127 > 0) {
                                        v58 := GenS2{F0: fi40}
                                        var v59 float64 = v13
                                        switch v41 {
                                        case 42:
                                            if (70 > 2) {
                                                var v60 int64 = 271
                                                v12 += 2147483647
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v60
                                            } else {
                                                var v61 int64 = 1
                                                _ = v61
                                            }
                                        case 128:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            switch "\n\t" {
                                            case "":
                                                if false {
                                                    v62 := GenS6{F0: true, F1: 0}
                                                    var v63 float64 = v59
                                                    _ = v62
                                                    _ = v63
                                                }
                                            case "a":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "foo bar":
                                                var v64 uint16 = 65535
                                                _ = v64
                                            }
                                        default:
                                            v4 -= float32(v13)
                                        }
                                        _ = v58
                                        _ = v59
                                    } else {
                                        v1 = append(v1, v12)
                                    }
                                }
                            case "foo bar":
                                switch v2 {
                                case "":
                                    var v65 float64 = 3.14159
                                    _ = v65
                                case "\n\t":
                                    v3 = (v4 + (v4 - 0.1))
                                default:
                                    var v66 uint64 = 2
                                    _ = v66
                                }
                            default:
                                if (0 == 128) {
                                    switch "\n\t" {
                                    case "a":
                                        switch (v2 + (v2 + "hello")) {
                                        case "\n\t":
                                            v67 := GenS7{F0: ("\n\t" + (v2 + "foo bar")), F1: 91, F2: 32767}
                                            _ = v67
                                        default:
                                            fmt.Println(v13, v8, fi39)
                                        }
                                    case "foo bar":
                                        if false {
                                            var v68 int64 = 42
                                            _ = v68
                                        } else {
                                            v12 -= (v12 >> ((genFn0()) & 63))
                                            if ((434 == 32767) || (-502 <= 1)) {
                                                v69 := GenS8{F0: (963 != 248)}
                                                v70 := GenS9{F0: 996}
                                                _ = v69
                                                _ = v70
                                            }
                                        }
                                    }
                                    if (false || (v8 >= (-(v8)))) {
                                        v41 += 127
                                        var v71 uint32 = (v12 | v12)
                                        if (234 != 0) {
                                            v4 = float32(float64(((v4 + v4) - v3)))
                                            v2 = ("hello" + v2)
                                            v5 = append(v5, 288)
                                        } else {
                                            fmt.Println(v4, v12)
                                            v72 := GenS9{F0: 42}
                                            fmt.Println(v13)
                                            fmt.Println(v4, v12)
                                            switch (v2 + "a") {
                                            case "hello":
                                                v1 = append(v1, v12)
                                            }
                                            _ = v72
                                        }
                                        var v73 uint64 = 4294967295
                                        var v74 int8 = 100
                                        var v75 int8 = (^(v74))
                                        _ = v71
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    }
                                    var v76 int32 = -203
                                    _ = v76
                                }
                            }
                        }
                    } else {
                        var v77 int64 = 1
                        fmt.Println(v77, fi40)
                        _ = v77
                    }
                    _ = fi40
                    _ = v41
                }
                v5 = append(v5, 42)
            } else {
                var v78 uint64 = 0
                var v79 uint32 = (v12 >> ((330) & 63))
                if true {
                    if (255 >= 255) {
                        v3 = -2.25
                    } else {
                        var v80 int32 = 1
                        var v81 uint = 65535
                        for fi82 := 0; fi82 < 16; fi82++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi82
                        }
                        _ = v80
                        _ = v81
                    }
                } else {
                    fmt.Println(fi39, v78)
                    if ((258 != 1) && true) {
                        for fi83 := 0; fi83 < 7; fi83++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi83
                        }
                        v78 = (v78 % (((v78 | (v78 + v78))) | 1))
                    } else {
                        fmt.Println(v3, v2)
                        if false {
                            if true {
                                if false {
                                    if true {
                                        for fi84 := 0; fi84 < 15; fi84++ {
                                            if false {
                                                fmt.Println(v8, v3, v12)
                                            } else {
                                                _ = len(v1)
                                                fmt.Println(v8)
                                                v4 = 0.1
                                            }
                                            if (141 != 267) {
                                                var v85 int16 = -32768
                                                v86 := GenS10{F0: (-(v8)), F1: -830}
                                                var v87 string = ((v2 + v2) + v2)
                                                v3 -= (v3 - (0.0 - v3))
                                                if (175 > 18446744073709551615) {
                                                    v1 = append(v1, 2147483648)
                                                } else {
                                                    v4 += float32(3.14159)
                                                }
                                                _ = v85
                                                _ = v86
                                                _ = v87
                                            }
                                            _ = fi84
                                        }
                                        switch ((((v2 + ("foo bar" + v2)) + v2) + v2) + "a") {
                                        case "a":
                                            if ((!(!((v12 - v12) == v12))) && true) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        case "\n\t":
                                            v78 = (v78 ^ (v78 << ((2) & 63)))
                                        case "foo bar":
                                            v8 += -624
                                        }
                                        var v88 uint8 = 42
                                        _ = v88
                                    } else {
                                        var v89 uint64 = v78
                                        v1 = append(v1, v79)
                                        _ = v89
                                    }
                                    for fi90 := 0; fi90 < 20; fi90++ {
                                        var v91 int32 = 0
                                        if len(v1) > 0 {
                                            v1[0] = 1
                                        }
                                        var v92 int8 = 0
                                        _ = fi90
                                        _ = v91
                                        _ = v92
                                    }
                                } else {
                                    v1 = append(v1, v79)
                                }
                            }
                            v93 := GenS9{F0: 0}
                            _ = v93
                        } else {
                            var v94 int8 = 37
                            _ = v94
                        }
                    }
                    v95 := GenS5{F0: v2, F1: 882}
                    _ = v95
                }
                _ = v78
                _ = v79
            }
            _ = fi39
        }
    default:
        var v96 uint = 42
        _ = v96
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
