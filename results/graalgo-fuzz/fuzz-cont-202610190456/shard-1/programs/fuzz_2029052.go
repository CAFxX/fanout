package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8; F1 int; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
type GenS5 struct { F0 uint32; F1 uint16; F2 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int; F1 int }
func genFn0(p0 int32, p1 uint16) int8 {
    if ((true || false) || true) {
        p0 = p0
    }
    v6 := []int{-1}
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 uint, p1 uint8) string {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint32 = 2
        p0 = (p0 << (((p0 << ((p0) & 63))) & 63))
        var v9 float32 = 3.14159
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = (0.0 + -2.25)
    var v11 bool = (178 > 87)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return "a"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 float32 = -2.25
    v4 := make(map[uint32]bool)
    var v5 int16 = 460
    v5 += (v5 / (((v5 * v5)) | 1))
    if (0 < 979) {
        switch (("hello" + ("foo bar" + "hello")) + "") {
        case "foo bar":
            var v12 int = 2147483647
            _ = v12
        case "\n\t":
            switch ("a" + "hello") {
            case "\n\t":
                var v13 int8 = 80
                _ = v13
            case "":
                _ = len(v2)
            }
        default:
            var v14 float32 = ((1e10 / v3) * v3)
            _ = v14
        }
    } else {
        v15 := []int{-462}
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v5)
            v17 := GenS1{F0: false, F1: "a"}
            fmt.Println(fi16, v3)
            _ = fi16
            _ = v17
        }
        v5 += (v5 - (v5 + (v5 << ((2) & 63))))
        _ = v15
    }
    v18 := GenS2{F0: 9223372036854775807, F1: 281}
    v3 = 3.14159
    v5 += (v5 % ((1) | 1))
    var v19 int = 476
    v20 := GenS2{F0: 42, F1: 765}
    if true {
        for fi21 := 0; fi21 < 18; fi21++ {
            if ((2147483647 != -109) || true) {
                var v22 uint64 = 2
                fmt.Println(fi21)
                if true {
                    if ((400 >= 994) || (false || true)) {
                        if true {
                            v23 := GenS3{F0: 216, F1: (v19 << ((399) & 63)), F2: 1}
                            if (true || (!false)) {
                                delete(v2, "hello")
                                var v24 uint32 = 42
                                _ = v24
                            }
                            _ = v23
                        }
                        fmt.Println(v19, v19)
                    } else {
                        v25 := GenS1{F0: (!(!(true || false))), F1: "a"}
                        _ = v25
                    }
                    if (false && (1 != -936)) {
                        v3 += (v3 - (v3 + v3))
                    } else {
                        fmt.Println(v3)
                    }
                } else {
                    v19 -= (v19 - (v19 / ((42) | 1)))
                    v22 = (v22 % (((^(v22))) | 1))
                    if (2 >= 0) {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            if (true || ((fi21 + fi21) == fi21)) {
                                switch 724 {
                                case 9223372036854775807:
                                    v5 += 772
                                }
                            } else {
                                var v27 float64 = (3.14159 * ((-2.25 + float64(v3)) - 1.5))
                                v28 := []int{(fi26 + -2147483648), fi21, (v19 | (fi26 >> ((981) & 63)))}
                                v29 := v1[("hello" + "")]
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            if (true && false) {
                                if (true || (0 <= 42)) {
                                    v30 := GenS1{F0: (128 > 1), F1: "a"}
                                    switch 509 {
                                    case 9223372036854775807:
                                        v31 := GenS2{F0: 1, F1: 255}
                                        _ = v31
                                    case 347:
                                        v3 -= (0.0 - v3)
                                    case -1:
                                        if (1 <= 65535) {
                                            var v32 uint16 = 42
                                            v33 := []int{}
                                            if (!false) {
                                                v33 = append(v33, 42)
                                                var v34 int8 = -65
                                                fmt.Println(fi21)
                                                _ = v34
                                            }
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v19 += (fi21 ^ (fi26 & fi21))
                                            if (!false) {
                                                if true {
                                                    v35 := GenS4{F0: 255, F1: v22, F2: ""}
                                                    var v36 bool = false
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    fmt.Println(v22, v19, v19)
                                                }
                                            }
                                            v1[("foo bar" + (("" + "\n\t") + "\n\t"))] = "a"
                                            fmt.Println(v19, fi26)
                                        }
                                    }
                                    _ = v30
                                } else {
                                    if false {
                                        var v37 uint16 = 1
                                        var v38 int64 = 0
                                        if false {
                                            var v39 int8 = 42
                                            _ = v39
                                        }
                                        v22 -= v22
                                        var v40 uint32 = 1
                                        if (false && (true && true)) {
                                            var v41 int16 = (v5 / ((1) | 1))
                                            _ = v41
                                        }
                                        v40 -= v40
                                        if (v37 == (^(v37))) {
                                            var v42 uint = 1
                                            var v43 int16 = v5
                                            switch (v19 << ((1) & 63)) {
                                            case -2147483648:
                                                fmt.Println(v19, v22)
                                            case 146:
                                                v44 := []int{(v19 << (((v42 << ((18446744073709551615) & 63))) & 63))}
                                                fmt.Println(v42, v19)
                                                v38 = -23
                                                _ = v44
                                            }
                                            _ = v42
                                            _ = v43
                                        }
                                        if (!false) {
                                            v38 += (v38 + (v38 >> ((65535) & 63)))
                                            v40 += v40
                                        } else {
                                            if (!false) {
                                                var v45 uint8 = 0
                                                fmt.Println(v38, v5)
                                                var v46 int64 = 1
                                                if true {
                                                    switch (v19 << ((4294967295) & 63)) {
                                                    case 127:
                                                        v47 := GenS5{F0: v40, F1: 32767, F2: 1}
                                                        _ = v47
                                                    case 895:
                                                        var v48 int64 = (v46 - v46)
                                                        _ = v48
                                                    case 768:
                                                        v22 = v22
                                                    }
                                                } else {
                                                    var v49 uint8 = (^(v45))
                                                    var v50 int32 = 1
                                                    _ = v49
                                                    _ = v50
                                                }
                                                v37 += v37
                                                v51 := GenS5{F0: v40, F1: v37, F2: (fi26 / ((32767) | 1))}
                                                fmt.Println(v3)
                                                if (42 <= 42) {
                                                    v37 = v37
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v22, v37)
                                                    case "foo bar":
                                                        v52 := []int{fi26, fi21, (v19 ^ fi21)}
                                                        _ = v52
                                                    }
                                                }
                                                v22 -= uint64(v37)
                                                _ = v45
                                                _ = v46
                                                _ = v51
                                            } else {
                                                v5 = (v5 & (v5 / ((0) | 1)))
                                                if true {
                                                    if (-9223372036854775808 != (v38 << ((42) & 63))) {
                                                        fmt.Println(v40, v37)
                                                        if false {
                                                            if true {
                                                                var v53 string = "\n\t"
                                                                _ = v53
                                                            }
                                                            v19 += (v19 / (((fi26 ^ v19)) | 1))
                                                        } else {
                                                            v54 := []int{128, (fi26 - (v19 + fi21)), -209, 32768}
                                                            _ = v54
                                                        }
                                                        var v55 int32 = -1
                                                        _ = v55
                                                    } else {
                                                        if true {
                                                            v40 -= uint32(v22)
                                                        } else {
                                                            var v56 int16 = v5
                                                            v22 += v22
                                                            fmt.Println(v3, v19)
                                                            fmt.Println(v3, v5)
                                                            if false {
                                                                v4[v40] = (v38 < -1)
                                                            }
                                                            _ = v56
                                                        }
                                                        v57 := GenS6{F0: 100}
                                                        _ = v57
                                                    }
                                                    var v58 uint16 = (v37 % (((v37 * (v37 << ((0) & 63)))) | 1))
                                                    _ = v58
                                                }
                                                v38 = 232
                                            }
                                            v59 := GenS4{F0: (^(v37)), F1: v22, F2: "\n\t"}
                                            switch v19 {
                                            case 0:
                                                v38 += (v38 - v38)
                                            case 658:
                                                if true {
                                                    switch ("hello" + "foo bar") {
                                                    case "":
                                                        if (-2.25 == 1.5) {
                                                            fmt.Println(v37, v5)
                                                            var v60 int8 = genFn0(869, v37)
                                                            v22 -= 2
                                                            _ = v60
                                                        } else {
                                                            if true {
                                                                v61 := GenS5{F0: (v40 >> ((255) & 63)), F1: v37, F2: v19}
                                                                fmt.Println(v22, fi21, v5)
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    fmt.Println(v40, v38, fi26)
                                                                case "hello":
                                                                    var v62 int32 = 19
                                                                    _ = v62
                                                                case "a":
                                                                    var v63 string = ("hello" + ("a" + ("" + "\n\t")))
                                                                    _ = v63
                                                                default:
                                                                    fmt.Println(v40, v3)
                                                                }
                                                                var v64 int8 = 1
                                                                _ = v61
                                                                _ = v64
                                                            } else {
                                                                v65 := GenS3{F0: 127, F1: fi26, F2: 41}
                                                                v66 := GenS7{F0: 42, F1: v22}
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            var v67 int64 = v38
                                                            var v68 uint = 464
                                                            delete(v1, "hello")
                                                            var v69 bool = false
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    default:
                                                        if false {
                                                            v40 = v40
                                                            if false {
                                                                v70 := GenS8{F0: fi26, F1: v19}
                                                                v71 := []int{}
                                                                v4[(v40 >> ((946) & 63))] = true
                                                                _ = v70
                                                                _ = v71
                                                            } else {
                                                                v72 := GenS4{F0: v37, F1: (^(v22)), F2: "\n\t"}
                                                                _ = v72
                                                            }
                                                        }
                                                        v38 -= (v38 / (((v38 * 0)) | 1))
                                                        var v73 uint16 = 854
                                                        _ = v73
                                                    }
                                                }
                                                var v74 int = (fi26 % ((v19) | 1))
                                                if (!(0 >= uint(v40))) {
                                                    v74 = fi21
                                                    v4[v40] = false
                                                } else {
                                                    v40 += v40
                                                    var v75 uint64 = 781
                                                    var v76 int16 = v5
                                                    _ = v75
                                                    _ = v76
                                                }
                                                _ = v74
                                            case 42:
                                                if (false || (false && ((65535 >= 18446744073709551615) && true))) {
                                                    if (9223372036854775807 > (-(v19))) {
                                                        fmt.Println(v22, fi21, fi26)
                                                    }
                                                } else {
                                                    if false {
                                                        var v77 uint16 = v37
                                                        var v78 uint = 4294967295
                                                        v38 -= v38
                                                        _ = v77
                                                        _ = v78
                                                    } else {
                                                        if (!(false && true)) {
                                                            v38 += 97
                                                            fmt.Println(v3, v40, fi26)
                                                            if (false || false) {
                                                                var v79 int32 = 216
                                                                v37 = (v37 + (v37 | v37))
                                                                var v80 float64 = 0.0
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                        } else {
                                                            var v81 bool = false
                                                            _ = v81
                                                        }
                                                        v3 += ((v3 - 1e10) / v3)
                                                    }
                                                    fmt.Println(v38, v22)
                                                }
                                            default:
                                                var v82 float64 = 3.14159
                                                _ = v82
                                            }
                                            _ = v59
                                        }
                                        _ = v37
                                        _ = v38
                                        _ = v40
                                    } else {
                                        v3 = float32((0.0 + 1e10))
                                        v22 += v22
                                        var v83 uint8 = 128
                                        _ = v83
                                    }
                                    var v84 float32 = v3
                                    _ = v84
                                }
                                if true {
                                    v22 = (v22 | 71)
                                    v85 := GenS6{F0: 127}
                                    v86 := []int{-728, (fi21 ^ -237)}
                                    _ = v85
                                    _ = v86
                                } else {
                                    v22 += (^(v22))
                                    v3 -= v3
                                    v5 = (^(v5))
                                    v87 := []int{v19, -9223372036854775808, 128}
                                    _ = v87
                                }
                            }
                            _ = fi26
                        }
                        for fi88 := 0; fi88 < 1; fi88++ {
                            var v89 uint8 = uint8(v22)
                            _ = fi88
                            _ = v89
                        }
                    } else {
                        var v90 int16 = v5
                        _ = v90
                    }
                }
                _ = v22
            }
            _ = fi21
        }
        v19 += v19
    }
    v91 := []int{v19}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v91:", len(v91), v91)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v91
}
