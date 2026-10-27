package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 int32; F1 uint8; F2 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int64; F1 int64; F2 int16 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int; F1 int64; F2 int32 }
func genFn0(p0 int16, p1 int64, p2 bool) int {
    v4 := []int{}
    v5 := GenS1{F0: -1, F1: 127, F2: "hello"}
    p0 = 0
    var v6 uint8 = 255
    var v7 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 int8 = 1
    var v2 uint16 = 65535
    var v3 uint64 = 1
    var v8 uint16 = 798
    for fi9 := 0; fi9 < 6; fi9++ {
        switch ("" + (("" + "\n\t") + ("a" + "a"))) {
        case "":
            for fi10 := 0; fi10 < 5; fi10++ {
                switch "hello" {
                case "":
                    if (!true) {
                        v8 = v2
                        if true {
                            var v11 uint16 = v8
                            fmt.Println(fi9)
                            if false {
                                if true {
                                    v8 += 32767
                                    var v12 int16 = 1
                                    _ = v12
                                } else {
                                    v8 -= (v2 - v8)
                                    var v13 int8 = (-(v1))
                                    v13 = v1
                                    _ = v13
                                }
                                v14 := GenS1{F0: v1, F1: 1, F2: "hello"}
                                _ = v14
                            }
                            _ = v11
                        }
                        v8 -= v2
                    } else {
                        fmt.Println(fi10, fi10)
                        v15 := GenS1{F0: 97, F1: 57, F2: "foo bar"}
                        if ((!false) && (false && true)) {
                            var v16 float64 = (float64((3.14159 + (0.1 * (float32(1.5) + 0.1)))) / -2.25)
                            _ = v16
                        }
                        v17 := []int{(-(fi10)), genFn0(-1, 9223372036854775807, true)}
                        _ = v15
                        _ = v17
                    }
                case "foo bar":
                    var v18 int32 = -344
                    _ = v18
                case "\n\t":
                    v19 := []int{(fi9 ^ (fi10 % (((-(fi9))) | 1))), 1, 298, genFn0(0, -1, (18446744073709551615 != 255))}
                    _ = v19
                default:
                    var v20 uint64 = (v3 ^ 865)
                    _ = v20
                }
                v1 += v1
                _ = fi10
            }
        case "foo bar":
            if (!(!true)) {
                v3 = (v3 ^ (v3 & (v3 | (v3 & 0))))
                v8 = v2
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 uint = 255
                    fmt.Println(v2, v8, v1)
                    _ = fi21
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: v1, F1: 2, F2: "foo bar"}
                _ = v23
            }
        default:
            var v24 string = "\n\t"
            _ = v24
        }
        var v25 bool = true
        fmt.Println(v1)
        _ = fi9
        _ = v25
    }
    if false {
        var v26 uint32 = 42
        _ = v26
    }
    var v27 int32 = -2147483648
    if (true || false) {
        v28 := GenS2{F0: -324, F1: 1, F2: genFn0(32767, -808, ((float32((3.14159 - 0.0)) != float32(3.14159)) || (("" + "hello") == (("hello" + "foo bar") + (((("a" + "\n\t") + ("a" + "a")) + "foo bar") + "foo bar")))))}
        v29 := []int{923, 651, 32767, 42}
        _ = v28
        _ = v29
    }
    v27 = int32(v2)
    v30 := []int{9223372036854775807}
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 6; fi32++ {
            switch "\n\t" {
            case "a":
                v33 := GenS3{F0: "", F1: v1}
                switch (fi31 << ((804) & 63)) {
                case -9223372036854775808:
                    var v34 int32 = v27
                    _ = v34
                case 42:
                    _ = len(v30)
                    if false {
                        v27 += v27
                        v35 := GenS2{F0: (-(v27)), F1: 42, F2: 127}
                        _ = len(v30)
                        switch -9223372036854775808 {
                        case 4294967295:
                            var v36 uint64 = (^(v3))
                            if (false || true) {
                                if false {
                                    var v37 float32 = -2.25
                                    _ = v37
                                } else {
                                    if len(v30) > 0 {
                                        v30[0] = (fi31 & (fi32 / (((-(fi32))) | 1)))
                                    }
                                    v1 -= (v1 | v1)
                                    switch -61 {
                                    case 0:
                                        var v38 float32 = 0.1
                                        _ = v38
                                    case -2147483648:
                                        if false {
                                            var v39 uint16 = 65535
                                            fmt.Println(v36, v1)
                                            fmt.Println(fi32)
                                            _ = v39
                                        }
                                    case 32768:
                                        if false {
                                            v3 -= 205
                                            var v40 uint64 = (v36 ^ (v3 | (v36 & v36)))
                                            v40 = v3
                                            v27 += (v27 * v27)
                                            v41 := GenS4{F0: 1000}
                                            v36 += (^(v40))
                                            switch "hello" {
                                            case "hello":
                                                var v42 string = "foo bar"
                                                _ = v42
                                            case "":
                                                fmt.Println(fi32)
                                            case "foo bar":
                                                if len(v30) >= 2 {
                                                    v30 = v30[1:len(v30)]
                                                }
                                            default:
                                                v1 += (v1 % (((v1 * (v1 & 1))) | 1))
                                            }
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v2 = v2
                                        }
                                    }
                                }
                                if false {
                                    fmt.Println(v2)
                                    var v43 int32 = (-(v27))
                                    _ = v43
                                }
                                switch 9223372036854775807 {
                                case 42:
                                    v27 += 12
                                case 32767:
                                    fmt.Println(fi31, fi32, v3)
                                }
                            }
                            _ = v36
                        }
                        var v44 int32 = (v27 - (v27 + v27))
                        _ = v35
                        _ = v44
                    }
                case 9223372036854775807:
                    var v45 uint = 449
                    _ = v45
                }
                _ = v33
            default:
                var v46 int32 = v27
                _ = v46
            }
            v1 = 0
            var v47 uint32 = 65535
            _ = fi32
            _ = v47
        }
        _ = fi31
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        if (!(1.5 != 3.14159)) {
            fmt.Println(fi48)
        } else {
            switch fi48 {
            case 709:
                for fi49 := 0; fi49 < 6; fi49++ {
                    switch (^(fi48)) {
                    case 9223372036854775807:
                        if true {
                            switch "\n\t" {
                            case "hello":
                                if (128 != 42) {
                                    var v50 uint8 = 255
                                    fmt.Println(fi49, fi49, fi48)
                                    _ = v50
                                } else {
                                    fmt.Println(v3, v2, v1)
                                }
                            case "a":
                                v2 -= v2
                            case "":
                                var v51 int16 = 32767
                                _ = v51
                            default:
                                v27 = v27
                            }
                            fmt.Println(v3, v8)
                        } else {
                            if ((!(255 <= 201)) || (!false)) {
                                fmt.Println(v8)
                                v1 = v1
                                var v52 uint32 = 506
                                _ = v52
                            } else {
                                v3 -= v3
                                v3 = 4294967295
                            }
                            var v53 string = "\n\t"
                            _ = v53
                        }
                    default:
                        v54 := GenS5{F0: 42}
                        _ = v54
                    }
                    if len(v30) > 0 {
                        v30[0] = 127
                    }
                    _ = fi49
                }
            case 127:
                v55 := GenS4{F0: -32768}
                _ = v55
            case -9223372036854775808:
                var v56 bool = false
                _ = v56
            }
            var v57 int64 = -501
            _ = v57
        }
        var v58 bool = ((v8 | (v8 % (((v2 & v8)) | 1))) <= (v8 << ((0) & 63)))
        var v59 int32 = 108
        _ = fi48
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 15; fi60++ {
        if ((v8 + v2) != (v8 >> ((305) & 63))) {
            v2 = v8
        } else {
            for fi61 := 0; fi61 < 10; fi61++ {
                fmt.Println(v27, v3, v2)
                if true {
                    switch fi60 {
                    case 377:
                        switch (fi60 & 4294967295) {
                        case 980:
                            var v62 uint = 1
                            _ = v62
                        case 32768:
                            fmt.Println(v27, fi60)
                        default:
                            var v63 uint16 = v2
                            fmt.Println(v63, fi60, v1)
                            _ = v63
                        }
                    case 32768:
                        v8 += (v8 ^ v2)
                    case 9223372036854775807:
                        fmt.Println(v27)
                    default:
                        var v64 float32 = 1e10
                        _ = v64
                    }
                    v65 := GenS1{F0: (^(v1)), F1: 76, F2: ""}
                    switch "" {
                    case "\n\t":
                        v66 := GenS5{F0: (v27 << ((18446744073709551615) & 63))}
                        _ = v66
                    case "a":
                        v27 = 1
                        fmt.Println(fi60)
                    case "":
                        if (2 <= 0) {
                            var v67 int32 = v27
                            _ = v67
                        } else {
                            var v68 int32 = v27
                            var v69 bool = (!true)
                            switch fi60 {
                            case -549:
                                fmt.Println(v69)
                            case 1:
                                v70 := GenS3{F0: "a", F1: v1}
                                _ = v70
                            }
                            _ = v68
                            _ = v69
                        }
                    default:
                        if false {
                            if true {
                                v71 := GenS5{F0: (v27 * (^(v27)))}
                                v27 = v27
                                _ = v71
                            } else {
                                var v72 uint16 = v8
                                _ = v72
                            }
                            if len(v30) > 0 {
                                v30[0] = 127
                            }
                        } else {
                            if (462 > 4294967295) {
                                var v73 uint64 = 18446744073709551615
                                _ = v73
                            } else {
                                switch "a" {
                                case "hello":
                                    fmt.Println(v3, v27, v1)
                                case "":
                                    v27 = v27
                                case "\n\t":
                                    var v74 float32 = -2.25
                                    switch 42 {
                                    case 1:
                                        v3 = (v3 - v3)
                                    case -33:
                                        if len(v30) > 0 {
                                            v30[0] = genFn0(1000, 0, (220 == 4294967295))
                                        }
                                    case -168:
                                        var v75 uint = 859
                                        var v76 int16 = 0
                                        v27 = v27
                                        _ = v75
                                        _ = v76
                                    default:
                                        v77 := GenS6{F0: false, F1: 2}
                                        _ = v77
                                    }
                                    var v78 uint8 = 2
                                    _ = v74
                                    _ = v78
                                default:
                                    v3 -= 1
                                }
                                if ((v27 << ((42) & 63)) >= (^(v27))) {
                                    switch (fi61 >> ((65535) & 63)) {
                                    case 128:
                                        if len(v30) > 0 {
                                            v30[0] = (fi60 / ((genFn0(1000, -745, (!false))) | 1))
                                        }
                                        var v79 string = (("a" + ("" + "hello")) + ("" + (((("\n\t" + ("foo bar" + "a")) + "hello") + (("" + "foo bar") + ("a" + ("\n\t" + "a")))) + ("" + "\n\t"))))
                                        _ = v79
                                    case 9223372036854775807:
                                        var v80 uint = 903
                                        _ = v80
                                    default:
                                        v2 -= v2
                                    }
                                    if false {
                                        if ((!(-9223372036854775808 == 0)) || true) {
                                            fmt.Println(fi61)
                                            var v81 uint8 = 128
                                            fmt.Println(v3)
                                            _ = v81
                                        }
                                        if len(v30) > 0 {
                                            v30[0] = fi61
                                        }
                                        if ((fi60 | fi60) <= genFn0(1000, 0, false)) {
                                            var v82 string = (("hello" + ("foo bar" + "")) + ("hello" + "a"))
                                            if (!(v27 == -2147483648)) {
                                                switch genFn0(966, 0, (true || true)) {
                                                case 1:
                                                    v2 += 1
                                                case 396:
                                                    var v83 bool = ((!false) || false)
                                                    _ = v83
                                                case -9223372036854775808:
                                                    v30 = append(v30, -9223372036854775808)
                                                default:
                                                    v82 = v82
                                                }
                                                v2 -= (v2 ^ (v8 / ((v8) | 1)))
                                                v84 := GenS7{F0: -258, F1: 548, F2: -742}
                                                _ = len(v30)
                                                _ = v84
                                            }
                                            _ = v82
                                        } else {
                                            var v85 uint = 498
                                            fmt.Println(fi60, v3, v1)
                                            v86 := GenS8{F0: 2, F1: (v1 % (((v1 % (((v1 ^ -128)) | 1))) | 1))}
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                }
                            }
                            v87 := GenS3{F0: (("\n\t" + "\n\t") + (("a" + ("foo bar" + "\n\t")) + "")), F1: (v1 >> ((785) & 63))}
                            v8 -= uint16(v3)
                            _ = v87
                        }
                    }
                    _ = v65
                } else {
                    v8 = v8
                }
                v88 := GenS6{F0: false, F1: 661}
                _ = fi61
                _ = v88
            }
            if ((0 > 255) && (42 < 84)) {
                if false {
                    v1 -= -1
                    v89 := GenS9{F0: -32768}
                    switch fi60 {
                    case 492:
                        _ = len(v30)
                    case 2147483647:
                        var v90 uint64 = 4294967295
                        _ = v90
                    case 128:
                        v91 := GenS10{F0: fi60, F1: 42, F2: (v27 + (v27 >> ((18446744073709551615) & 63)))}
                        _ = v91
                    default:
                        _ = len(v30)
                    }
                    _ = v89
                } else {
                    _ = len(v30)
                    v27 = v27
                }
            } else {
                var v92 int8 = v1
                _ = v92
            }
            v93 := GenS9{F0: 32767}
            _ = v93
        }
        v2 = (v2 / ((v8) | 1))
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            if (-128 <= int8(v3)) {
                var v94 int32 = (-(v27))
                switch "foo bar" {
                case "foo bar":
                    v1 += 1
                case "hello":
                    v1 += (v1 / ((v1) | 1))
                case "":
                    v27 += v27
                }
                _ = v94
            } else {
                v95 := GenS10{F0: -1, F1: -475, F2: 0}
                _ = len(v30)
                _ = v95
            }
        case "hello":
            for fi96 := 0; fi96 < 20; fi96++ {
                switch "\n\t" {
                case "hello":
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                default:
                    fmt.Println(fi60)
                }
                _ = fi96
            }
        case "foo bar":
            var v97 int64 = -52
            _ = v97
        default:
            var v98 int = fi60
            _ = v98
        }
        _ = fi60
    }
    var v99 float64 = ((-2.25 * (0.1 * (((1.5 + float64((3.14159 + 1.5))) - 1.5) + ((3.14159 / ((float64(1.5) * 1e10) / ((1.5 - 1e10) / 1e10))) * ((0.0 - (float64(0.1) * 1e10)) + ((float64(1e10) / 3.14159) + ((((((0.0 + float64((((float32(float64(1.5)) * ((float32((((3.14159 - 0.0) - (3.14159 / 1.5)) * (-2.25 - 0.0))) * 3.14159) * ((-2.25 / 1.0) * 1e10))) - 1.5) / 1.5))) + (1.5 + (float64((1e10 + 1e10)) + 3.14159))) + ((((3.14159 - float64(1.5)) / 1.0) + (1e10 * -2.25)) * 0.1)) * (float64((1e10 / (1.5 + (float32((0.1 + -2.25)) + -2.25)))) / -2.25)) * 3.14159) + -2.25))))))) / ((0.1 / (1.5 + 1.5)) * float64(float32(((1.5 * 1e10) - ((1.5 + 1e10) / 0.1))))))
    var v100 uint16 = (^(v8))
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v8:", v8)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v8
    _ = v99
}
