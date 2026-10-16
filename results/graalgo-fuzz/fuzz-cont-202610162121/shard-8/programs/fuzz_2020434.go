package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int; F2 int32 }
type GenS6 struct { F0 int64; F1 int; F2 int }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 string }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint16) string {
    if true {
        p0 = p0
    }
    var v4 float32 = (1.5 / 1.0)
    v5 := GenS1{F0: 2, F1: 75}
    v6 := GenS1{F0: (p0 >> ((255) & 63)), F1: 65535}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return "a"
}

func main() {
    var v1 int16 = 229
    var v2 int = 4294967295
    var v3 bool = false
    v7 := GenS2{F0: 0, F1: v1}
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v1)
        switch "a" {
        case "":
            v9 := []int{(fi8 << ((18446744073709551615) & 63)), fi8}
            _ = v9
        case "a":
            var v10 int = (-(v2))
            _ = v10
        default:
            if v3 {
                v11 := GenS2{F0: 43, F1: v1}
                var v12 uint16 = 736
                _ = v11
                _ = v12
            }
        }
        _ = fi8
    }
    var v13 int8 = -112
    fmt.Println(v1, v3)
    var v14 int64 = -248
    fmt.Println(v1, v1, v2)
    switch "hello" {
    case "hello":
        var v15 int = (^(v2))
        _ = v15
    case "foo bar":
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := []int{fi16, fi16, (v2 / ((42) | 1))}
            if true {
                var v18 uint32 = 252
                v19 := GenS2{F0: 18446744073709551615, F1: (v1 / ((v1) | 1))}
                _ = v18
                _ = v19
            } else {
                switch "foo bar" {
                case "hello":
                    if v3 {
                        var v20 int16 = int16(v13)
                        for fi21 := 0; fi21 < 14; fi21++ {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            _ = fi21
                        }
                        _ = v20
                    }
                case "":
                    for fi22 := 0; fi22 < 17; fi22++ {
                        fmt.Println(fi22, v3, v14)
                        _ = fi22
                    }
                default:
                    fmt.Println(v13)
                }
            }
            var v23 int32 = -2147483648
            _ = fi16
            _ = v17
            _ = v23
        }
    case "a":
        v3 = v3
    default:
        v3 = true
    }
    v24 := []int{v2}
    v25 := GenS3{F0: (v1 / (((v1 / (((v1 << ((2) & 63))) | 1))) | 1)), F1: 1, F2: -1}
    v13 = v13
    v26 := GenS2{F0: 2, F1: v1}
    var v27 int8 = (^(v13))
    fmt.Println(v2)
    v14 -= 489
    var v28 int = (v2 >> ((0) & 63))
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 20; fi30++ {
            if len(v24) > 0 {
                v24[0] = 4294967295
            }
            v13 -= 17
            if v3 {
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi29, fi29)
                case "\n\t":
                    if false {
                        var v31 int32 = 2147483647
                        switch v2 {
                        case 144:
                            var v32 uint16 = 65535
                            _ = v32
                        case -679:
                            var v33 string = ("foo bar" + "a")
                            _ = v33
                        }
                        switch fi30 {
                        case -472:
                            v34 := GenS1{F0: 979, F1: 65535}
                            _ = v34
                        case 1:
                            v27 += (v13 >> ((42) & 63))
                            if (880 != 245) {
                                v35 := GenS1{F0: 65535, F1: 255}
                                v36 := GenS1{F0: 42, F1: 107}
                                if true {
                                    switch ("foo bar" + "") {
                                    case "a":
                                        v37 := GenS4{F0: (v14 % ((int64(v31)) | 1)), F1: 931, F2: 732}
                                        _ = v37
                                    default:
                                        v27 -= v13
                                    }
                                }
                                _ = v35
                                _ = v36
                            }
                            if false {
                                var v38 uint8 = 71
                                _ = v38
                            }
                        default:
                            switch v28 {
                            case 2147483647:
                                var v39 string = "foo bar"
                                _ = v39
                            case 9223372036854775807:
                                v40 := GenS4{F0: v14, F1: (v14 << ((65535) & 63)), F2: 634}
                                _ = v40
                            default:
                                v41 := GenS3{F0: (v1 % ((v1) | 1)), F1: v13, F2: (v27 + (v13 ^ 1))}
                                _ = v41
                            }
                            var v42 bool = false
                            v13 = (v13 / (((^(v13))) | 1))
                            _ = v42
                        }
                        if (0 != 562) {
                            _ = len(v24)
                            var v43 int = (fi30 * int(v31))
                            switch fi30 {
                            case 0:
                                var v44 int = (v2 << ((42) & 63))
                                _ = v44
                            }
                            _ = v43
                        } else {
                            v3 = false
                            switch "foo bar" {
                            case "":
                                var v45 uint32 = 330
                                _ = v45
                            case "\n\t":
                                if true {
                                    if (v3 && v3) {
                                        var v46 int16 = v1
                                        if len(v24) > 0 {
                                            v24[0] = 4294967295
                                        }
                                        if (!(255 < 1)) {
                                            if len(v24) > 0 {
                                                v24[0] = fi30
                                            }
                                            v24 = append(v24, fi29)
                                        }
                                        _ = v46
                                    }
                                    var v47 int64 = v14
                                    fmt.Println(v31)
                                    _ = v47
                                }
                            }
                            var v48 int = 127
                            _ = v48
                        }
                        _ = v31
                    }
                default:
                    v28 += (fi29 * (fi29 & (fi30 - (fi29 - -1))))
                }
                var v49 uint16 = 1
                var v50 float64 = float64(-2.25)
                switch ("\n\t" + "") {
                case "":
                    if v3 {
                        if len(v24) > 0 {
                            v24[0] = (fi30 << ((240) & 63))
                        }
                        var v51 int = (v2 / ((-250) | 1))
                        if (450 >= 716) {
                            switch "hello" {
                            case "hello":
                                var v52 int64 = (-(v14))
                                _ = v52
                            case "\n\t":
                                switch (((("a" + ("\n\t" + (("\n\t" + "foo bar") + ("a" + (((("a" + "hello") + "hello") + "a") + "a"))))) + "foo bar") + "\n\t") + ("hello" + "")) {
                                case "\n\t":
                                    var v53 uint = 1
                                    _ = v53
                                case "foo bar":
                                    if len(v24) > 0 {
                                        v24[0] = v28
                                    }
                                case "hello":
                                    var v54 string = (("foo bar" + "foo bar") + "hello")
                                    _ = v54
                                }
                            case "a":
                                v28 -= v2
                            default:
                                var v55 int32 = 761
                                _ = v55
                            }
                            if len(v24) > 0 {
                                v24[0] = (v2 ^ -1)
                            }
                            v27 -= v27
                        } else {
                            if false {
                                var v56 uint64 = 18446744073709551615
                                _ = v56
                            }
                        }
                        v24 = append(v24, fi30)
                        if ((^(v49)) < (v49 / (((^(v49))) | 1))) {
                            if (!(155 != 2)) {
                                if (18 > 0) {
                                    v50 = v50
                                } else {
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    if true {
                                        var v57 uint8 = 127
                                        switch "foo bar" {
                                        case "foo bar":
                                            v58 := GenS1{F0: (v49 ^ (v49 >> ((42) & 63))), F1: 4294967295}
                                            _ = v58
                                        case "hello":
                                            v59 := GenS4{F0: v14, F1: (v14 + v14), F2: (^(v49))}
                                            _ = v59
                                        case "":
                                            v60 := GenS3{F0: v1, F1: (v13 >> ((909) & 63)), F2: 42}
                                            _ = v60
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            var v61 int64 = (v14 % (((v14 % (((^(v14))) | 1))) | 1))
                                            var v62 bool = v3
                                            _ = v61
                                            _ = v62
                                        default:
                                            var v63 uint8 = uint8(v49)
                                            _ = v63
                                        }
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            if (3.14159 <= v50) {
                                                var v64 int32 = 2147483647
                                                v24 = append(v24, (fi30 ^ (v2 >> ((uint(v57)) & 63))))
                                                _ = v64
                                            } else {
                                                v65 := GenS5{F0: 2, F1: (fi29 << ((139) & 63)), F2: -313}
                                                switch v2 {
                                                case 128:
                                                    v13 = v27
                                                case 32768:
                                                    v57 -= 42
                                                }
                                                _ = len(v24)
                                                _ = v65
                                            }
                                        case "foo bar":
                                            v66 := GenS3{F0: v1, F1: 100, F2: int8(v1)}
                                            _ = v66
                                        case "hello":
                                            v3 = (((1e10 > 1.5) || ((-(v13)) < v27)) && v3)
                                        }
                                        v67 := GenS6{F0: v14, F1: fi29, F2: v28}
                                        _ = v57
                                        _ = v67
                                    } else {
                                        fmt.Println(v14, fi30)
                                        v68 := GenS7{F0: 2, F1: 0}
                                        switch v51 {
                                        case 127:
                                            if ((v49 - 701) != (v49 - (v49 ^ v49))) {
                                                if false {
                                                    var v69 int32 = -1
                                                    if false {
                                                        v1 -= (v1 + int16(v69))
                                                    }
                                                    fmt.Println(v2)
                                                    _ = v69
                                                } else {
                                                    v49 = 0
                                                    fmt.Println(fi30, v50, v14)
                                                    var v70 int32 = -113
                                                    var v71 float32 = (1.5 / 1.0)
                                                    _ = v70
                                                    _ = v71
                                                }
                                                fmt.Println(fi30)
                                            }
                                        case 0:
                                            var v72 uint16 = v49
                                            _ = v72
                                        case 40:
                                            v14 += (-(v14))
                                        }
                                        _ = v68
                                    }
                                }
                                v27 += (-(v13))
                            } else {
                                v28 -= (fi30 & 128)
                            }
                            v73 := GenS2{F0: 898, F1: int16(v49)}
                            switch "\n\t" {
                            case "foo bar":
                                var v74 int32 = -385
                                _ = v74
                            case "":
                                v51 += v51
                            }
                            if v3 {
                                v13 = (^(v27))
                                v3 = ((v49 | v49) >= (v49 * 255))
                            } else {
                                var v75 string = (("" + "a") + "\n\t")
                                _ = v75
                            }
                            _ = v73
                        }
                        _ = len(v24)
                        v1 = 32767
                        v24 = append(v24, v28)
                        _ = v51
                    } else {
                        v76 := GenS2{F0: 822, F1: -32768}
                        v2 = (v28 << ((975) & 63))
                        fmt.Println(v13, v3, v28)
                        _ = v76
                    }
                case "foo bar":
                    if len(v24) > 0 {
                        v24[0] = (v2 / (((v2 / ((2147483647) | 1))) | 1))
                    }
                case "a":
                    var v77 float32 = float32(v50)
                    _ = v77
                default:
                    var v78 string = ""
                    _ = v78
                }
                switch "a" {
                case "":
                    switch v28 {
                    case -950:
                        var v79 uint8 = 0
                        _ = v79
                    case 312:
                        v13 -= (v13 << ((18446744073709551615) & 63))
                    case -9223372036854775808:
                        v80 := GenS8{F0: 449}
                        _ = v80
                    }
                default:
                    var v81 uint64 = 2
                    _ = v81
                }
                _ = v49
                _ = v50
            }
            _ = fi30
        }
        _ = fi29
    }
    v13 = v13
    if v3 {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        v82 := GenS9{F0: (v14 | (v14 * 1)), F1: "a"}
        _ = v82
    } else {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        if (!v3) {
            if (-2.25 >= 3.14159) {
                var v83 uint8 = 42
                _ = v83
            }
            switch ("" + "") {
            case "a":
                if (1 >= 255) {
                    if false {
                        v84 := GenS10{F0: 82}
                        _ = v84
                    } else {
                        v14 = v14
                        var v85 int = 0
                        _ = v85
                    }
                }
            case "hello":
                if len(v24) > 0 {
                    v24[0] = 0
                }
            default:
                var v86 uint64 = 18446744073709551615
                _ = v86
            }
        }
        var v87 uint32 = 255
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v7
}
