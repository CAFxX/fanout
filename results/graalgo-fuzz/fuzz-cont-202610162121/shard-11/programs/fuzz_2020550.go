package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8 }
func genFn0() uint16 {
    v7 := []int{776, 2147483647}
    if ((410 == 745) && false) {
        var v8 uint = 693
        v8 += (v8 * v8)
        _ = v8
    } else {
        var v9 uint64 = 1
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 uint32 = 2147483647
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    var v12 int = 32768
    v12 -= 1
    var v13 int64 = 0
    _ = v12
    _ = v13
    _ = v7
    return 1
}
func genFn1(p0 uint, p1 uint32, p2 int16) int {
    v14 := GenS1{F0: genFn0()}
    for fi15 := 0; fi15 < 14; fi15++ {
        switch "a" {
        case "foo bar":
            v16 := GenS1{F0: 2}
            _ = v16
        }
        _ = fi15
    }
    if true {
        if (((3.14159 / 1.0) - ((float64(float32(0.0)) / -2.25) + -2.25)) < (3.14159 * 0.0)) {
            fmt.Println(p2)
            var v17 uint16 = genFn0()
            _ = v17
        } else {
            if (42 >= -680) {
                fmt.Println(p2, p2)
                for fi18 := 0; fi18 < 19; fi18++ {
                    p2 = (-(p2))
                    _ = fi18
                }
                var v19 string = "foo bar"
                _ = v19
            }
            v20 := []int{42, 0, 34, 32768}
            v21 := GenS2{F0: 0, F1: (p0 / ((p0) | 1))}
            _ = v20
            _ = v21
        }
    }
    var v22 uint16 = 0
    v23 := GenS3{F0: int8(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v22
    _ = v23
    return -142
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 int16 = 32767
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    var v6 int = 4294967295
    v24 := GenS3{F0: 0}
    fmt.Println(v3)
    var v25 float64 = 1e10
    switch genFn1(646, 808, -32768) {
    case 2147483647:
        var v26 string = "foo bar"
        _ = v26
    }
    var v27 bool = (-9223372036854775808 == 42)
    switch (v6 * v6) {
    case -988:
        if (!false) {
            if v27 {
                if (42 >= 100) {
                    v6 -= (v6 * 32768)
                    v6 += v6
                    fmt.Println(v6)
                } else {
                    if false {
                        for fi28 := 0; fi28 < 1; fi28++ {
                            v29 := []int{(fi28 & (^(fi28)))}
                            fmt.Println(v27, v2, v3)
                            if v27 {
                                v30 := v4[v6]
                                var v31 bool = false
                                if v31 {
                                    v27 = (0.1 == (0.1 + (-2.25 / 3.14159)))
                                    for fi32 := 0; fi32 < 16; fi32++ {
                                        var v33 uint32 = 2147483648
                                        if (65535 == 4294967295) {
                                            v30 = (v30 + "foo bar")
                                            var v34 uint64 = (v2 % ((v2) | 1))
                                            switch genFn1(1, v33, v3) {
                                            case -1:
                                                v33 = (v33 >> ((255) & 63))
                                            case 42:
                                                if (1 == 1) {
                                                    var v35 float32 = 3.14159
                                                    var v36 uint8 = 1
                                                    v27 = (!true)
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    if v27 {
                                                        var v37 uint8 = 128
                                                        _ = v37
                                                    }
                                                }
                                            }
                                            _ = v34
                                        } else {
                                            var v38 float32 = 1e10
                                            var v39 int8 = -128
                                            switch fi28 {
                                            case 128:
                                                v2 = uint64(v33)
                                            case 688:
                                                v40 := GenS3{F0: -34}
                                                _ = v40
                                            case 0:
                                                v30 = "a"
                                            }
                                            switch "hello" {
                                            case "a":
                                                v6 = (v6 * (fi32 * fi28))
                                            case "":
                                                v6 = v6
                                            case "hello":
                                                var v41 bool = v31
                                                _ = v41
                                            }
                                            var v42 uint8 = 216
                                            if (false || (genFn1(488, 0, v3) >= genFn1(255, (^(v33)), 1000))) {
                                                v33 += v33
                                                switch -796 {
                                                case 127:
                                                    v42 = 228
                                                case 32768:
                                                    v38 += float32(float64((0.1 - v38)))
                                                    var v43 uint8 = (^(v42))
                                                    _ = v43
                                                case -9223372036854775808:
                                                    v30 = "foo bar"
                                                }
                                            } else {
                                                var v44 string = ("a" + v30)
                                                switch genFn1(458, 2147483648, (^(v3))) {
                                                case 32767:
                                                    switch fi32 {
                                                    case 127:
                                                        var v45 uint16 = 92
                                                        _ = v45
                                                    case 4294967295:
                                                        v46 := v4[-2147483648]
                                                        _ = v46
                                                    default:
                                                        switch (v44 + (("\n\t" + (v30 + v30)) + "\n\t")) {
                                                        case "":
                                                            fmt.Println(fi28, fi32)
                                                            fmt.Println(fi28, v27)
                                                            v33 += (^(v33))
                                                        case "\n\t":
                                                            var v47 int16 = 42
                                                            _ = v47
                                                        case "hello":
                                                            var v48 float64 = (v25 + 1e10)
                                                            _ = v48
                                                        default:
                                                            v49 := GenS3{F0: (v39 ^ 42)}
                                                            v6 = (v6 ^ genFn1(130, v33, 0))
                                                            _ = v49
                                                        }
                                                    }
                                                    var v50 int32 = 257
                                                    switch v30 {
                                                    case "":
                                                        if (v2 != v2) {
                                                            v30 = v30
                                                            var v51 int32 = v50
                                                            fmt.Println(v42, v25)
                                                            switch ("" + v30) {
                                                            case "hello":
                                                                v52 := GenS4{F0: v39, F1: (v39 | v39)}
                                                                _ = v52
                                                            case "foo bar":
                                                                v53 := GenS4{F0: (v39 + 0), F1: -68}
                                                                _ = v53
                                                            }
                                                            _ = v51
                                                        } else {
                                                            v31 = true
                                                        }
                                                        switch (^(v6)) {
                                                        case -2147483648:
                                                            switch fi28 {
                                                            case 117:
                                                                v54 := GenS3{F0: (v39 / ((int8(v50)) | 1))}
                                                                _ = v54
                                                            }
                                                        }
                                                        delete(v4, (fi28 & (v6 & 2147483647)))
                                                    case "a":
                                                        var v55 float64 = 0.1
                                                        _ = v55
                                                    case "hello":
                                                        v29 = append(v29, (v6 & (-(fi32))))
                                                    default:
                                                        if ((v3 >= v3) && v31) {
                                                            v56 := GenS3{F0: (v39 ^ v39)}
                                                            var v57 uint32 = (v33 | (v33 >> ((0) & 63)))
                                                            var v58 int8 = (v39 ^ (-(v39)))
                                                            if (v27 || (-1 <= 42)) {
                                                                _ = len(v29)
                                                                v2 -= 42
                                                            } else {
                                                                if (7 > 601) {
                                                                    v6 += (v6 + (v6 * v6))
                                                                    v59 := GenS2{F0: 609, F1: 2}
                                                                    v60 := GenS5{F0: 2, F1: (v2 >> ((0) & 63))}
                                                                    var v61 int32 = v50
                                                                    var v62 uint8 = (v42 | v42)
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    v63 := GenS5{F0: uint16(v42), F1: (v2 & (v2 & (v2 % ((v2) | 1))))}
                                                                    v64 := GenS4{F0: (v39 >> ((670) & 63)), F1: (v39 ^ -9)}
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                                if (!false) {
                                                                    fmt.Println(fi32, v31)
                                                                    v29 = append(v29, (v6 - fi32))
                                                                }
                                                                v65 := v1["foo bar"]
                                                                fmt.Println(v2, v25)
                                                                _ = v65
                                                            }
                                                            v66 := v4[(fi28 & fi28)]
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                            _ = v66
                                                        }
                                                    }
                                                    _ = v50
                                                case -2147483648:
                                                    v31 = v31
                                                case 128:
                                                    switch v30 {
                                                    case "hello":
                                                        var v67 int16 = (v3 & v3)
                                                        _ = v67
                                                    case "\n\t":
                                                        var v68 float32 = v38
                                                        _ = v68
                                                    case "a":
                                                        fmt.Println(fi32, fi28, fi28)
                                                    default:
                                                        v4[(fi28 >> ((0) & 63))] = v44
                                                    }
                                                default:
                                                    v69 := GenS2{F0: -9223372036854775808, F1: 392}
                                                    _ = v69
                                                }
                                                if false {
                                                    v4[int(v3)] = v30
                                                    if (!v31) {
                                                        if ((v27 && false) && v27) {
                                                            v27 = (42 < (v2 * (v2 + uint64(v42))))
                                                            v42 += (v42 >> ((65535) & 63))
                                                        } else {
                                                            if (668 != -1) {
                                                                v33 -= v33
                                                                v70 := GenS6{F0: 2, F1: (v2 + (^(v2))), F2: 255}
                                                                v29 = append(v29, fi32)
                                                                if v31 {
                                                                    switch v30 {
                                                                    case "foo bar":
                                                                        if (int8(v6) == (v39 + (v39 ^ 14))) {
                                                                            fmt.Println(v30, v33, fi28)
                                                                            v71 := GenS1{F0: 42}
                                                                            switch -64 {
                                                                            case -1:
                                                                                if (v38 < v38) {
                                                                                    v44 = v30
                                                                                    v5[int(v39)] = (!false)
                                                                                    v72 := GenS3{F0: (-(v39))}
                                                                                    v27 = v27
                                                                                    _ = v72
                                                                                } else {
                                                                                    fmt.Println(fi32, fi28, fi28)
                                                                                }
                                                                            case 127:
                                                                                _ = len(v4)
                                                                            }
                                                                            _ = v71
                                                                        } else {
                                                                            var v73 uint64 = (^(v2))
                                                                            v30 = v44
                                                                            var v74 int8 = v39
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                    case "hello":
                                                                        _ = len(v4)
                                                                    default:
                                                                        var v75 float64 = v25
                                                                        _ = v75
                                                                    }
                                                                    switch genFn1(782, (v33 >> ((818) & 63)), v3) {
                                                                    case -1:
                                                                        v39 += -36
                                                                    case 32768:
                                                                        fmt.Println(v30, fi32)
                                                                    case 279:
                                                                        if ((!((false && (((v44 + v30) + "foo bar") > "\n\t")) && false)) || v31) {
                                                                            v76 := GenS7{F0: genFn1(582, 765, (v3 / ((v3) | 1)))}
                                                                            v4[(fi32 % ((-2147483648) | 1))] = v30
                                                                            if (!v31) {
                                                                                v77 := GenS8{F0: (v42 | 42)}
                                                                                fmt.Println(fi28, v39, v25)
                                                                                _ = v77
                                                                            }
                                                                            _ = v76
                                                                        } else {
                                                                            v1[("a" + (v30 + v44))] = 4294967295
                                                                        }
                                                                        var v78 string = ("" + ("" + "\n\t"))
                                                                        _ = v78
                                                                    }
                                                                } else {
                                                                    var v79 int64 = 1
                                                                    if v27 {
                                                                        v33 -= 2147483648
                                                                    }
                                                                    var v80 uint64 = 4294967295
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                var v81 uint32 = (^(v33))
                                                                _ = v70
                                                                _ = v81
                                                            }
                                                        }
                                                        var v82 int32 = -238
                                                        _ = v82
                                                    }
                                                }
                                                v27 = (v42 != v42)
                                                _ = v44
                                            }
                                            _ = v38
                                            _ = v39
                                            _ = v42
                                        }
                                        if len(v29) > 0 {
                                            v29[0] = int(v3)
                                        }
                                        v5[fi32] = (!v27)
                                        v25 = float64(1.5)
                                        var v83 int = fi32
                                        switch "a" {
                                        case "foo bar":
                                            var v84 int8 = -1
                                            _ = v84
                                        case "hello":
                                            var v85 uint8 = 0
                                            _ = v85
                                        case "\n\t":
                                            var v86 uint = 4294967295
                                            _ = v86
                                        default:
                                            fmt.Println(fi32, fi28)
                                        }
                                        _ = fi32
                                        _ = v33
                                        _ = v83
                                    }
                                }
                                _ = v30
                                _ = v31
                            } else {
                                v25 += ((-2.25 + (((float64((3.14159 / (1.5 / 1.0))) - 0.1) + (((float64(3.14159) * float64(1.5)) - -2.25) + v25)) + (1e10 - (1e10 - ((3.14159 * 3.14159) + v25))))) / v25)
                                var v87 int64 = 216
                                _ = v87
                            }
                            _ = fi28
                            _ = v29
                        }
                        v3 = (v3 | 32767)
                        v88 := v1["\n\t"]
                        _ = v88
                    }
                    var v89 int8 = 0
                    _ = v89
                }
            } else {
                switch "hello" {
                case "hello":
                    if (!v27) {
                        delete(v1, (("foo bar" + "\n\t") + ""))
                        var v90 int64 = 0
                        v91 := []int{(v6 >> ((42) & 63)), int(v90), v6, v6}
                        _ = v90
                        _ = v91
                    }
                case "a":
                    v2 = (v2 ^ v2)
                case "foo bar":
                    v1["foo bar"] = (v6 ^ (v6 | (v6 % ((v6) | 1))))
                }
                fmt.Println(v27)
                v2 = v2
            }
            fmt.Println(v6)
            if v27 {
                switch "hello" {
                case "foo bar":
                    for fi92 := 0; fi92 < 15; fi92++ {
                        var v93 int = (v6 % (((fi92 ^ 127)) | 1))
                        _ = fi92
                        _ = v93
                    }
                }
                for fi94 := 0; fi94 < 8; fi94++ {
                    var v95 float64 = (v25 - v25)
                    v25 += v25
                    _ = fi94
                    _ = v95
                }
                v25 -= 3.14159
            }
        }
    }
    for fi96 := 0; fi96 < 6; fi96++ {
        v2 += v2
        var v97 uint32 = 314
        var v98 uint = uint(v2)
        _ = fi96
        _ = v97
        _ = v98
    }
    if v27 {
        v3 = v3
        if (!(!(-140 <= 0))) {
            v5[1] = ((true && (3.14159 != 0.1)) || (1 >= 1))
            for fi99 := 0; fi99 < 5; fi99++ {
                v100 := []int{v6, -1, v6, fi99}
                if (v27 || true) {
                    v3 = -464
                    fmt.Println(fi99, v25, fi99)
                } else {
                    v25 -= v25
                }
                _ = fi99
                _ = v100
            }
            v6 = v6
        }
    } else {
        v101 := []int{}
        for fi102 := 0; fi102 < 14; fi102++ {
            v3 -= (v3 - 552)
            fmt.Println(fi102)
            _ = fi102
        }
        _ = v101
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
