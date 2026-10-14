package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int16; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint8; F2 string }
type GenS6 struct { F0 int32; F1 int64; F2 int }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 int16; F1 int32; F2 int16 }
type GenS12 struct { F0 int8; F1 int; F2 uint64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint; F1 int64 }
type GenS16 struct { F0 int64; F1 uint8 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 uint32, p1 int64, p2 string) int64 {
    p0 = (p0 * p0)
    v5 := []int{}
    var v6 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 + (p1 / ((p1) | 1)))
}
func genFn1(p0 int8, p1 uint, p2 uint) int8 {
    v7 := []int{}
    switch 1 {
    case -9223372036854775808:
        p2 = 65535
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 ^ (p0 | p0))
}

func main() {
    v1 := []uint32{65535, 470}
    var v2 string = ""
    var v3 uint32 = 0
    var v4 int64 = -210
    v3 -= (^(v3))
    if (false || false) {
        switch 32767 {
        case 0:
            var v8 uint32 = (v3 / (((v3 / ((1) | 1))) | 1))
            _ = v8
        case 943:
            switch -339 {
            case 32767:
                if false {
                    var v9 bool = true
                    var v10 uint32 = (v3 | (^(v3)))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v10
                    _ = v9
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v11 uint32 = (v3 >> ((2) & 63))
                    _ = v11
                }
            case -2147483648:
                if true {
                    if false {
                        for fi12 := 0; fi12 < 9; fi12++ {
                            if true {
                                var v13 uint32 = 341
                                v2 = ("" + (v2 + ((v2 + v2) + ("\n\t" + ("" + ((((("foo bar" + ((v2 + v2) + (v2 + v2))) + v2) + "a") + v2) + (v2 + v2)))))))
                                _ = v13
                            } else {
                                if true {
                                    fmt.Println(v3)
                                    v3 = (^(v3))
                                }
                                var v14 uint = 2
                                _ = len(v1)
                                _ = v14
                            }
                            _ = fi12
                        }
                        switch 32767 {
                        case 32768:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case -380:
                            for fi15 := 0; fi15 < 5; fi15++ {
                                v16 := GenS1{F0: v4, F1: true, F2: 255}
                                _ = fi15
                                _ = v16
                            }
                        case 2147483647:
                            var v17 string = (v2 + v2)
                            _ = v17
                        }
                    }
                    for fi18 := 0; fi18 < 5; fi18++ {
                        var v19 uint64 = 4294967295
                        for fi20 := 0; fi20 < 10; fi20++ {
                            var v21 float64 = 1.5
                            _ = fi20
                            _ = v21
                        }
                        fmt.Println(v2)
                        _ = fi18
                        _ = v19
                    }
                }
            }
        case 9223372036854775807:
            switch (v2 + ("foo bar" + v2)) {
            case "hello":
                fmt.Println(v4, v3, v2)
            default:
                var v22 int16 = 1
                _ = v22
            }
        }
    } else {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint = 887
            var v25 bool = true
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v26 := GenS2{F0: -396}
    v2 = ""
    v27 := GenS3{F0: 89, F1: 1, F2: 195}
    var v28 int = 0
    _ = len(v1)
    var v29 int64 = v4
    if (72 != 127) {
        if (!(!true)) {
            var v30 uint16 = 65535
            var v31 int16 = 32767
            _ = v30
            _ = v31
        } else {
            var v32 int8 = 42
            for fi33 := 0; fi33 < 18; fi33++ {
                if false {
                    var v34 uint8 = 2
                    _ = v34
                } else {
                    var v35 float64 = ((-2.25 + 1.5) - (float64(1.5) * 1e10))
                    if len(v1) > 0 {
                        v1[0] = (v3 - 1)
                    }
                    fmt.Println(v28, v3)
                    _ = v35
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v28 = v28
                _ = fi33
            }
            _ = v32
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            v29 += (v4 & (v4 * v4))
            _ = fi36
        }
        if true {
            var v37 int16 = 0
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 int8 = -75
        for fi40 := 0; fi40 < 15; fi40++ {
            v41 := GenS4{F0: (!(0 != 190))}
            if len(v1) > 0 {
                v1[0] = v3
            }
            if (!((!false) && (-189 <= 42))) {
                v1 = append(v1, v3)
                var v42 uint8 = 245
                _ = v42
            } else {
                switch v2 {
                case "foo bar":
                    if true {
                        var v43 uint8 = uint8(v3)
                        fmt.Println(v4)
                        fmt.Println(v28, v39)
                        v29 += -128
                        _ = v43
                    }
                default:
                    v44 := GenS1{F0: v4, F1: (true || true), F2: 432}
                    _ = v44
                }
                switch fi38 {
                case -42:
                    v29 = genFn0(v3, (v29 ^ -335), v2)
                case 107:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                }
            }
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    if (127 == 127) {
        for fi45 := 0; fi45 < 2; fi45++ {
            var v46 int = int(v3)
            v47 := GenS5{F0: 42, F1: 128, F2: (v2 + "\n\t")}
            for fi48 := 0; fi48 < 13; fi48++ {
                if (1 > -128) {
                    v46 -= (fi45 & (^(fi45)))
                    v49 := GenS6{F0: 963, F1: (^(v29)), F2: v28}
                    _ = v49
                }
                var v50 bool = ((v3 - (^(v3))) >= (v3 + v3))
                switch v2 {
                case "\n\t":
                    fmt.Println(v4, v3, v2)
                case "hello":
                    switch (fi48 ^ 127) {
                    case -885:
                        v2 = ((v2 + v2) + (v2 + (((v2 + v2) + v2) + v2)))
                    case 42:
                        var v51 uint8 = 42
                        var v52 string = v2
                        switch "\n\t" {
                        case "":
                            var v53 int8 = 1
                            _ = v53
                        case "\n\t":
                            var v54 bool = (!false)
                            _ = v54
                        case "hello":
                            var v55 float32 = 3.14159
                            _ = v55
                        }
                        _ = v51
                        _ = v52
                    case 190:
                        fmt.Println(v2)
                    }
                default:
                    if (0.0 <= -2.25) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v56 uint16 = 1
                        if v50 {
                            _ = len(v1)
                        } else {
                            switch ((v2 + "a") + v2) {
                            case "":
                                if true {
                                    v29 -= genFn0((^(v3)), genFn0(v3, genFn0((v3 + v3), -9223372036854775808, v2), v2), (v2 + "foo bar"))
                                    if (!true) {
                                        fmt.Println(fi48, v2, v46)
                                        v56 = 65535
                                    } else {
                                        if true {
                                            var v57 int16 = 42
                                            var v58 int8 = 100
                                            switch fi45 {
                                            case 2147483647:
                                                var v59 uint32 = (v3 % ((65535) | 1))
                                                _ = v59
                                            default:
                                                if v50 {
                                                    v60 := GenS4{F0: v50}
                                                    fmt.Println(v50, fi45)
                                                    if (!(v2 >= "a")) {
                                                        v61 := GenS5{F0: -1, F1: 127, F2: v2}
                                                        switch v28 {
                                                        case 0:
                                                            fmt.Println(v56, v4)
                                                        case -466:
                                                            v2 = v2
                                                        case -904:
                                                            fmt.Println(v50, v2, fi45)
                                                        default:
                                                            switch ("\n\t" + (v2 + ("" + "a"))) {
                                                            case "\n\t":
                                                                if (((v4 << ((1) & 63)) != (^(v29))) && v50) {
                                                                    var v62 int64 = -1
                                                                    _ = len(v1)
                                                                    var v63 uint32 = 991
                                                                    _ = v62
                                                                    _ = v63
                                                                } else {
                                                                    var v64 int64 = -530
                                                                    switch fi45 {
                                                                    case 453:
                                                                        var v65 int = (v28 << ((18446744073709551615) & 63))
                                                                        _ = v65
                                                                    case 32768:
                                                                        v66 := GenS7{F0: v2, F1: int32(v3)}
                                                                        _ = v66
                                                                    default:
                                                                        v56 += 413
                                                                        if ((v50 && true) || ((v58 - -65) <= v58)) {
                                                                            switch 32768 {
                                                                            case 42:
                                                                                v4 += genFn0((v3 << ((543) & 63)), v64, (("foo bar" + v2) + v2))
                                                                            case 953:
                                                                                v28 = (fi48 << ((65535) & 63))
                                                                            default:
                                                                                v56 = (^(v56))
                                                                            }
                                                                            if v50 {
                                                                                var v67 int16 = 0
                                                                                v68 := GenS4{F0: ((((float32(0.1) - (float32(3.14159) + 0.0)) * 0.1) * (((3.14159 + 1.5) - (0.0 * 3.14159)) + (float32((3.14159 - -2.25)) * (3.14159 / 1e10)))) == (((1e10 - 1e10) * 3.14159) + (float32((0.1 * float64(float32((0.0 + (3.14159 / 1.0)))))) - (-2.25 / 3.14159))))}
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        } else {
                                                                            switch v2 {
                                                                            case "hello":
                                                                                _ = len(v1)
                                                                            default:
                                                                                v29 = 1
                                                                            }
                                                                        }
                                                                        if v50 {
                                                                            if v50 {
                                                                                switch (v28 >> ((463) & 63)) {
                                                                                case -326:
                                                                                    fmt.Println(v58, v2)
                                                                                    v1 = append(v1, v3)
                                                                                    var v69 string = "\n\t"
                                                                                    _ = v69
                                                                                case -373:
                                                                                    var v70 int32 = 0
                                                                                    _ = v70
                                                                                }
                                                                                if (!v50) {
                                                                                    v56 = v56
                                                                                }
                                                                                if v50 {
                                                                                    var v71 string = ((("\n\t" + ("\n\t" + v2)) + ((v2 + v2) + "foo bar")) + ((v2 + v2) + (v2 + (v2 + v2))))
                                                                                    var v72 uint = 4294967295
                                                                                    if v50 {
                                                                                        v28 = (v28 + fi45)
                                                                                        v71 = (((v2 + (v2 + v2)) + v2) + (v71 + v2))
                                                                                    } else {
                                                                                        switch (^(fi48)) {
                                                                                        case 2147483647:
                                                                                            fmt.Println(v64, v2, v29)
                                                                                        case 4294967295:
                                                                                            v73 := GenS8{F0: 0, F1: -384}
                                                                                            _ = v73
                                                                                        case -979:
                                                                                            var v74 int8 = (v58 | -128)
                                                                                            _ = v74
                                                                                        default:
                                                                                            if (!v50) {
                                                                                                v28 = 1
                                                                                                if (v50 && true) {
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                    if v50 {
                                                                                                        switch (fi48 + (v28 * (v28 ^ v46))) {
                                                                                                        case 4294967295:
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(fi48, v3, fi48)
                                                                                                        }
                                                                                                        v1 = append(v1, v3)
                                                                                                        _ = len(v1)
                                                                                                    } else {
                                                                                                        if (false || (18446744073709551615 <= 4294967295)) {
                                                                                                            if v50 {
                                                                                                                switch (v71 + v71) {
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v71)
                                                                                                                case "a":
                                                                                                                    v56 += v56
                                                                                                                    var v75 int = fi48
                                                                                                                    _ = v75
                                                                                                                default:
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                }
                                                                                                                var v76 float32 = (3.14159 - 0.0)
                                                                                                                _ = v76
                                                                                                            } else {
                                                                                                                switch (-(v28)) {
                                                                                                                case -9223372036854775808:
                                                                                                                    v77 := GenS5{F0: 0, F1: 128, F2: v2}
                                                                                                                    _ = v77
                                                                                                                default:
                                                                                                                    if v50 {
                                                                                                                        var v78 int8 = genFn1(v58, uint(v3), (v72 * 255))
                                                                                                                        var v79 int32 = -1
                                                                                                                        v50 = true
                                                                                                                        _ = v78
                                                                                                                        _ = v79
                                                                                                                    } else {
                                                                                                                        if v50 {
                                                                                                                            if v50 {
                                                                                                                                var v80 int64 = -1
                                                                                                                                switch fi45 {
                                                                                                                                case 9223372036854775807:
                                                                                                                                    v72 -= v72
                                                                                                                                default:
                                                                                                                                    var v81 string = v2
                                                                                                                                    _ = v81
                                                                                                                                }
                                                                                                                                _ = v80
                                                                                                                            } else {
                                                                                                                                _ = len(v1)
                                                                                                                            }
                                                                                                                            var v82 uint64 = 1
                                                                                                                            v58 += genFn1((v58 / ((v58) | 1)), (^(v72)), (v72 % ((v72) | 1)))
                                                                                                                            v71 = v2
                                                                                                                            _ = v82
                                                                                                                        }
                                                                                                                        var v83 int64 = genFn0(v3, (^(v4)), v2)
                                                                                                                        _ = v83
                                                                                                                    }
                                                                                                                }
                                                                                                                var v84 string = v71
                                                                                                                var v85 float64 = (1e10 / (0.1 / 3.14159))
                                                                                                                var v86 uint8 = 244
                                                                                                                if len(v1) >= 2 {
                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                }
                                                                                                                _ = v84
                                                                                                                _ = v85
                                                                                                                _ = v86
                                                                                                            }
                                                                                                            if (18446744073709551615 >= 18446744073709551615) {
                                                                                                                if false {
                                                                                                                    v87 := GenS6{F0: -1, F1: v4, F2: (v28 - (fi48 - -1))}
                                                                                                                    if v50 {
                                                                                                                        fmt.Println(v2)
                                                                                                                    }
                                                                                                                    if (0.0 >= (1.5 - 3.14159)) {
                                                                                                                        _ = len(v1)
                                                                                                                        v50 = v50
                                                                                                                    } else {
                                                                                                                        if (((v3 + v3) >= (v3 | 2147483647)) || ((v72 * v72) == (v72 / ((v72) | 1)))) {
                                                                                                                            var v88 string = "a"
                                                                                                                            var v89 int8 = -128
                                                                                                                            v4 = v64
                                                                                                                            _ = v88
                                                                                                                            _ = v89
                                                                                                                        }
                                                                                                                        if ((v72 != (v72 >> ((v72) & 63))) && v50) {
                                                                                                                            v90 := GenS9{F0: 42, F1: (v56 / ((v56) | 1))}
                                                                                                                            var v91 uint16 = (v56 - v56)
                                                                                                                            _ = v90
                                                                                                                            _ = v91
                                                                                                                        } else {
                                                                                                                            if true {
                                                                                                                                var v92 int = fi48
                                                                                                                                v1 = append(v1, 2147483648)
                                                                                                                                _ = v92
                                                                                                                            }
                                                                                                                            var v93 int32 = -1
                                                                                                                            _ = v93
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v94 int32 = 633
                                                                                                                    if v50 {
                                                                                                                        v1 = append(v1, 4294967295)
                                                                                                                    } else {
                                                                                                                        v72 = v72
                                                                                                                    }
                                                                                                                    _ = v87
                                                                                                                    _ = v94
                                                                                                                } else {
                                                                                                                    v29 += (v64 % ((v4) | 1))
                                                                                                                    switch -2147483648 {
                                                                                                                    case -2147483648:
                                                                                                                        var v95 int = fi45
                                                                                                                        _ = v95
                                                                                                                    default:
                                                                                                                        v64 = v29
                                                                                                                    }
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v96 string = "foo bar"
                                                                                                                v1 = append(v1, (v3 | (v3 % ((42) | 1))))
                                                                                                                _ = v96
                                                                                                            }
                                                                                                        } else {
                                                                                                            v4 += genFn0((v3 + v3), v29, "a")
                                                                                                            fmt.Println(v56, v46, fi45)
                                                                                                            v97 := GenS8{F0: 255, F1: -2147483648}
                                                                                                            _ = v97
                                                                                                        }
                                                                                                        fmt.Println(fi45, v4)
                                                                                                        switch (v2 + (v2 + v2)) {
                                                                                                        case "\n\t":
                                                                                                            v57 = (v57 % (((v57 * (v57 - v57))) | 1))
                                                                                                        case "":
                                                                                                            v98 := GenS10{F0: 2147483647, F1: 42}
                                                                                                            _ = v98
                                                                                                        case "hello":
                                                                                                            if (v71 < (((((v71 + v2) + v71) + "foo bar") + ((v2 + (v71 + v2)) + v71)) + "")) {
                                                                                                                v1 = append(v1, 252)
                                                                                                                switch (((v71 + (v71 + (("\n\t" + (v2 + (v71 + v71))) + v71))) + "hello") + ("a" + (v2 + (("a" + v2) + v71)))) {
                                                                                                                case "":
                                                                                                                    fmt.Println(v2, v72)
                                                                                                                case "hello":
                                                                                                                    v71 = "a"
                                                                                                                case "a":
                                                                                                                    if v50 {
                                                                                                                        var v99 uint8 = 11
                                                                                                                        _ = v99
                                                                                                                    } else {
                                                                                                                        v57 += (v57 / (((v57 << ((v72) & 63))) | 1))
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v3 = 0
                                                                                                                }
                                                                                                            }
                                                                                                            var v100 bool = (!false)
                                                                                                            _ = v100
                                                                                                        default:
                                                                                                            v72 = (v72 | (v72 % ((65535) | 1)))
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v28 = (v46 | fi45)
                                                                                            } else {
                                                                                                if (!((1e10 > (3.14159 * -2.25)) && ((!false) && (2147483647 < -2147483648)))) {
                                                                                                    v3 -= v3
                                                                                                }
                                                                                                v101 := GenS2{F0: 9223372036854775807}
                                                                                                fmt.Println(v71)
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = v3
                                                                                        }
                                                                                    }
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v102 int8 = (v58 | (v58 % ((v58) | 1)))
                                                                            _ = v102
                                                                        }
                                                                    }
                                                                    v103 := GenS11{F0: v57, F1: 0, F2: v57}
                                                                    if (0.0 < -2.25) {
                                                                        var v104 string = v2
                                                                        _ = v104
                                                                    }
                                                                    v105 := GenS12{F0: v58, F1: (v28 ^ -613), F2: 898}
                                                                    _ = v103
                                                                    _ = v105
                                                                    _ = v64
                                                                }
                                                            case "foo bar":
                                                                v2 = v2
                                                            case "":
                                                                v106 := GenS13{F0: v28}
                                                                _ = v106
                                                            }
                                                        }
                                                        _ = v61
                                                    }
                                                    _ = v60
                                                }
                                            }
                                            _ = v57
                                            _ = v58
                                        }
                                    }
                                    switch "foo bar" {
                                    case "":
                                        var v107 int16 = 1
                                        _ = v107
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (v3 - (v3 & (v3 ^ (v3 % (((^(v3))) | 1)))))
                                    }
                                    _ = len(v1)
                                }
                            case "a":
                                var v108 int = 127
                                _ = v108
                            case "foo bar":
                                v109 := GenS14{F0: 65535}
                                switch "\n\t" {
                                case "":
                                    fmt.Println(fi48)
                                }
                                var v110 int64 = (v29 | (v4 << ((18446744073709551615) & 63)))
                                _ = v109
                                _ = v110
                            }
                            switch v2 {
                            case "foo bar":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "hello":
                                if (((((v4 % ((v4) | 1)) == v4) || v50) || v50) || v50) {
                                    var v111 int = (v28 * v46)
                                    var v112 float32 = 0.0
                                    if (genFn0(v3, (v29 % ((v4) | 1)), (v2 + ("\n\t" + v2))) == -209) {
                                        var v113 uint64 = 0
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = len(v1)
                                        _ = v113
                                    } else {
                                        var v114 bool = (v50 || ((v3 | v3) == 2147483648))
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v114
                                    }
                                    v1 = append(v1, (v3 << ((4294967295) & 63)))
                                    if true {
                                        var v115 uint8 = 0
                                        var v116 uint32 = (v3 - uint32(v115))
                                        _ = v115
                                        _ = v116
                                    }
                                    _ = v111
                                    _ = v112
                                } else {
                                    v3 += v3
                                }
                            default:
                                var v117 uint8 = 2
                                _ = v117
                            }
                            v29 += (v29 << ((588) & 63))
                            v2 = ""
                        }
                        _ = v56
                    }
                }
                _ = fi48
                _ = v50
            }
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    switch ((v2 + (v2 + ("foo bar" + (v2 + v2)))) + "\n\t") {
    case "\n\t":
        if (42 > 0) {
            if (!false) {
                for fi118 := 0; fi118 < 11; fi118++ {
                    switch v2 {
                    case "hello":
                        for fi119 := 0; fi119 < 16; fi119++ {
                            _ = len(v1)
                            v4 = (v29 / ((v29) | 1))
                            switch (v2 + v2) {
                            case "\n\t":
                                if (0 == (v29 | v29)) {
                                    v3 = v3
                                    switch "a" {
                                    case "":
                                        v1 = append(v1, 65535)
                                    case "foo bar":
                                        var v120 float64 = 3.14159
                                        var v121 uint16 = uint16(v3)
                                        _ = v120
                                        _ = v121
                                    case "a":
                                        v29 -= v29
                                    }
                                    v2 = ("foo bar" + "")
                                }
                            default:
                                v122 := GenS15{F0: 65535, F1: genFn0(4294967295, (v29 >> ((578) & 63)), v2)}
                                _ = v122
                            }
                            var v123 int16 = 0
                            _ = fi119
                            _ = v123
                        }
                    case "\n\t":
                        _ = len(v1)
                    default:
                        if (v2 <= v2) {
                            fmt.Println(v2, v3, fi118)
                            v4 += (v4 ^ (v4 - 42))
                        } else {
                            v124 := GenS16{F0: genFn0(4294967295, (-(v4)), "hello"), F1: 255}
                            fmt.Println(v2, v29, v3)
                            _ = v124
                        }
                    }
                    var v125 uint32 = v3
                    var v126 uint32 = (v3 + 437)
                    _ = fi118
                    _ = v125
                    _ = v126
                }
            } else {
                v127 := GenS16{F0: (^(v4)), F1: 1}
                if true {
                    v28 += v28
                    if true {
                        var v128 uint = 42
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v28)
                        _ = v128
                    }
                }
                fmt.Println(v29, v29, v28)
                _ = v127
            }
        }
    case "hello":
        v129 := GenS17{F0: 1}
        _ = v129
    case "":
        switch 1 {
        case 2147483647:
            if false {
                if (!false) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v4 = v4
                } else {
                    for fi130 := 0; fi130 < 18; fi130++ {
                        var v131 bool = ((false && (float32((0.1 * 1e10)) == 0.1)) && (false && false))
                        _ = fi130
                        _ = v131
                    }
                }
                v132 := GenS18{F0: 1, F1: 32767, F2: 156}
                fmt.Println(v3, v4)
                _ = v132
            } else {
                var v133 string = "hello"
                _ = v133
            }
        case 0:
            v28 += 32767
        case -488:
            switch (v2 + "hello") {
            case "a":
                if len(v1) > 0 {
                    v1[0] = (v3 >> ((2) & 63))
                }
            case "\n\t":
                v1 = append(v1, v3)
            case "":
                v1 = append(v1, (v3 * 840))
            }
        default:
            var v134 uint16 = 2
            _ = v134
        }
    default:
        for fi135 := 0; fi135 < 3; fi135++ {
            var v136 int16 = 158
            var v137 float32 = (-2.25 / (float32((0.1 + 1.5)) * 1.5))
            switch (^(fi135)) {
            case -638:
                v1 = append(v1, 2147483648)
            case 2147483647:
                if false {
                    var v138 string = (v2 + ("" + "\n\t"))
                    for fi139 := 0; fi139 < 7; fi139++ {
                        v1 = append(v1, (v3 >> ((42) & 63)))
                        if true {
                            var v140 int = (-(v28))
                            v141 := GenS1{F0: (v29 | genFn0(4294967295, v4, v138)), F1: ((!false) || true), F2: 2}
                            v2 = "foo bar"
                            _ = v140
                            _ = v141
                        } else {
                            v29 = 945
                            var v142 float32 = float32(-2.25)
                            switch (fi135 + fi139) {
                            case 42:
                                v143 := GenS15{F0: 4294967295, F1: 1}
                                _ = v143
                            case 1:
                                v1 = append(v1, v3)
                            case 643:
                                if (42 == 852) {
                                    v29 = genFn0(v3, int64(v136), "a")
                                    var v144 int = 128
                                    _ = v144
                                } else {
                                    if false {
                                        v3 = (v3 % ((42) | 1))
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = v3
                                        }
                                        var v145 int64 = v4
                                        if ((v136 | v136) <= v136) {
                                            fmt.Println(v136, v138, v29)
                                            if (9223372036854775807 < -9223372036854775808) {
                                                v146 := GenS15{F0: 65535, F1: (-(v4))}
                                                fmt.Println(v138)
                                                _ = v146
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = (v3 | (v3 / ((v3) | 1)))
                                                }
                                                var v147 string = v138
                                                v29 += (v4 >> ((255) & 63))
                                                _ = v147
                                            }
                                        }
                                        _ = v145
                                    }
                                    v142 -= 1e10
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            _ = v142
                        }
                        _ = fi139
                    }
                    _ = v138
                }
            case 4294967295:
                _ = len(v1)
            }
            _ = fi135
            _ = v136
            _ = v137
        }
    }
    v2 = ((v2 + "") + "\n\t")
    v148 := GenS10{F0: v28, F1: 42}
    var v149 uint32 = (^(v3))
    v28 -= v28
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v148
    _ = v149
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
