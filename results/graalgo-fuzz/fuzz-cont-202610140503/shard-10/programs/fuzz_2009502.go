package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint; F2 int64 }
type GenS5 struct { F0 uint16; F1 int32; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS8 struct { F0 uint64; F1 int32 }
func genFn0() uint64 {
    if false {
        var v5 int64 = 265
        _ = v5
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            fmt.Println(fi6)
        } else {
            var v7 int32 = 2147483647
            v8 := []int{fi6, (fi6 | fi6), (fi6 % (((fi6 + (fi6 - fi6))) | 1)), (fi6 * 1)}
            _ = v7
            _ = v8
        }
        var v9 uint64 = 276
        _ = fi6
        _ = v9
    }
    var v10 bool = false
    _ = v10
    return 2
}
func genFn1(p0 uint64, p1 bool, p2 int64) int64 {
    if ((-2147483648 < 0) && ("\n\t" <= "\n\t")) {
        var v11 bool = ((float64(1e10) - -2.25) > (0.0 / 1.0))
        _ = v11
    } else {
        v12 := GenS1{F0: 127, F1: 42}
        _ = v12
    }
    var v13 uint = 0
    var v14 uint64 = (p0 >> ((42) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]string)
    var v4 uint32 = 553
    for fi15 := 0; fi15 < 18; fi15++ {
        if (true && true) {
            fmt.Println(v1)
            v2 -= v2
        } else {
            var v16 int16 = 109
            switch 4294967295 {
            case 127:
                var v17 int8 = -128
                _ = v17
            case 9223372036854775807:
                v18 := []int{-2147483648}
                _ = v18
            }
            _ = v16
        }
        _ = fi15
    }
    if true {
        v2 -= (v2 * v2)
        if ((v1 - (v1 * v1)) >= v1) {
            v2 = 398
        }
    }
    if true {
        if (false && false) {
            var v19 bool = (!true)
            _ = v19
        } else {
            v20 := []int{312, int(v2), 4294967295}
            switch -714 {
            case 2147483647:
                switch (("hello" + (("" + "hello") + "a")) + "foo bar") {
                case "":
                    v4 = 2
                case "foo bar":
                    if (v1 < (v1 ^ (v1 << ((1) & 63)))) {
                        v21 := GenS1{F0: 100, F1: 128}
                        _ = v21
                    }
                case "hello":
                    if (((!false) || (872 > (v4 ^ (v4 >> ((1) & 63))))) || true) {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            for fi23 := 0; fi23 < 5; fi23++ {
                                v4 = (v4 - (v4 - (v4 % (((v4 / (((v4 + (v4 / (((v4 >> ((65535) & 63))) | 1)))) | 1))) | 1))))
                                if (2 == 199) {
                                    v24 := GenS1{F0: 1, F1: 127}
                                    fmt.Println(fi22)
                                    _ = v24
                                }
                                _ = fi23
                            }
                            fmt.Println(fi22, v2, v1)
                            _ = fi22
                        }
                    } else {
                        v2 -= (v2 + (v2 ^ genFn1(4294967295, (127 < 127), 0)))
                        fmt.Println(v4, v1, v1)
                    }
                }
            case 128:
                _ = len(v20)
            case 613:
                v1 += (v1 - (v1 % ((v1) | 1)))
            }
            var v25 int8 = 1
            _ = v20
            _ = v25
        }
        v2 -= genFn1(308, true, (^(v2)))
    } else {
        v3[-2147483648] = ""
    }
    v26 := GenS2{F0: "a"}
    switch "foo bar" {
    case "":
        v27 := []int{42, 0, 32767}
        _ = v27
    case "a":
        v3[582] = "a"
    case "\n\t":
        v28 := []int{32767}
        _ = v28
    }
    var v29 uint = 211
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 bool = (v2 != (v2 % (((^(v2))) | 1)))
        if (v4 == (v4 >> ((v29) & 63))) {
            var v32 int16 = 42
            if false {
                switch "\n\t" {
                case "":
                    v2 -= (v2 >> ((v29) & 63))
                case "a":
                    var v33 int = (fi30 * int(v1))
                    _ = v33
                }
                if v31 {
                    v4 = (v4 - v4)
                    var v34 uint32 = v4
                    _ = v34
                }
            } else {
                for fi35 := 0; fi35 < 20; fi35++ {
                    v36 := GenS2{F0: (("\n\t" + (((("hello" + "a") + "\n\t") + "a") + "hello")) + ("foo bar" + "foo bar"))}
                    _ = fi35
                    _ = v36
                }
                var v37 bool = (v31 && (100 == 79))
                fmt.Println(v2, v4)
                _ = v37
            }
            _ = v32
        } else {
            var v38 bool = (false || (1e10 < ((float64((1e10 + 3.14159)) + (1.5 * 3.14159)) - (1e10 / ((1e10 + (1e10 / (0.1 * (0.1 / 1.5)))) / 1.5)))))
            fmt.Println(v4, fi30)
            for fi39 := 0; fi39 < 4; fi39++ {
                v40 := GenS1{F0: -128, F1: 0}
                v41 := v3[715]
                var v42 float32 = -2.25
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
            _ = v38
        }
        v29 = (v29 % (((v29 * v29)) | 1))
        _ = fi30
        _ = v31
    }
    if (!((!false) || (!true))) {
        switch "a" {
        case "\n\t":
            if false {
                switch -1 {
                case 0:
                    if (3.14159 != (0.0 / (1.5 - 0.1))) {
                        if false {
                            for fi43 := 0; fi43 < 11; fi43++ {
                                v29 += (^(v29))
                                for fi44 := 0; fi44 < 7; fi44++ {
                                    switch "a" {
                                    case "hello":
                                        if (!false) {
                                            if true {
                                                v45 := []int{fi43, -2147483648, (fi44 / ((fi43) | 1))}
                                                var v46 int32 = (-(v1))
                                                v46 += v46
                                                _ = v45
                                                _ = v46
                                            }
                                            v47 := GenS1{F0: 127, F1: 127}
                                            fmt.Println(v2)
                                            v1 = int32(v2)
                                            if true {
                                                v48 := []int{fi43, (fi43 + fi44)}
                                                v1 -= v1
                                                _ = v48
                                            }
                                            _ = v47
                                        } else {
                                            var v49 int32 = 42
                                            _ = v49
                                        }
                                        var v50 uint32 = v4
                                        fmt.Println(v50, fi44, v29)
                                        _ = v50
                                    case "":
                                        var v51 uint8 = 1
                                        if (42 < 1000) {
                                            v51 -= v51
                                            v1 = -835
                                            v29 = v29
                                        }
                                        switch 32768 {
                                        case 83:
                                            var v52 int64 = v2
                                            var v53 float64 = ((3.14159 / 1.5) * (0.1 - float64(float32(-2.25))))
                                            _ = v52
                                            _ = v53
                                        case -804:
                                            switch 42 {
                                            case 128:
                                                v54 := GenS3{F0: 32767, F1: (fi43 % ((int(v29)) | 1)), F2: 1}
                                                _ = v54
                                            case 1:
                                                v55 := GenS2{F0: "hello"}
                                                v1 += v1
                                                _ = v55
                                            case 9223372036854775807:
                                                v56 := []int{(fi43 & (fi44 + fi44))}
                                                _ = v56
                                            default:
                                                fmt.Println(v4, v29)
                                            }
                                        }
                                        fmt.Println(v29)
                                        _ = v51
                                    }
                                    switch "hello" {
                                    case "hello":
                                        switch ("hello" + (("hello" + "\n\t") + "")) {
                                        case "hello":
                                            var v57 int32 = 1
                                            _ = v57
                                        case "":
                                            v4 = (v4 & (v4 % (((v4 << ((v29) & 63))) | 1)))
                                        default:
                                            delete(v3, fi43)
                                        }
                                    case "foo bar":
                                        v1 = v1
                                        switch "\n\t" {
                                        case "foo bar":
                                            v58 := []int{}
                                            _ = v58
                                        }
                                    case "\n\t":
                                        fmt.Println(v4, fi43)
                                    }
                                    if (-128 <= 0) {
                                        v4 -= v4
                                    } else {
                                        v59 := GenS2{F0: ("hello" + "hello")}
                                        _ = v59
                                    }
                                    v60 := GenS1{F0: 42, F1: 255}
                                    v1 = (v1 - (v1 / (((v1 ^ v1)) | 1)))
                                    _ = fi44
                                    _ = v60
                                }
                                _ = fi43
                            }
                        } else {
                            v61 := []int{9223372036854775807, 32767, 42, -29}
                            switch "a" {
                            case "foo bar":
                                v62 := GenS3{F0: 143, F1: 9223372036854775807, F2: 255}
                                _ = v62
                            case "a":
                                var v63 uint8 = 0
                                _ = v63
                            }
                            var v64 uint32 = v4
                            _ = v61
                            _ = v64
                        }
                    } else {
                        var v65 int64 = (v2 / (((-(v2))) | 1))
                        v66 := []int{128}
                        var v67 int8 = 100
                        _ = v65
                        _ = v66
                        _ = v67
                    }
                case -1:
                    v4 = 499
                case -791:
                    var v68 float64 = (1e10 * (0.0 - float64(1e10)))
                    _ = v68
                default:
                    if ((float64(1e10) + -2.25) > 0.1) {
                        switch ((("" + "a") + "") + "\n\t") {
                        case "foo bar":
                            var v69 int8 = -106
                            _ = v69
                        case "a":
                            for fi70 := 0; fi70 < 10; fi70++ {
                                v71 := GenS3{F0: 65535, F1: fi70, F2: 0}
                                switch (fi70 / (((fi70 << ((2) & 63))) | 1)) {
                                case -924:
                                    var v72 int16 = -32768
                                    _ = v72
                                case 32768:
                                    v73 := GenS3{F0: 32767, F1: (fi70 / ((fi70) | 1)), F2: 65535}
                                    _ = v73
                                default:
                                    switch "a" {
                                    case "a":
                                        var v74 uint32 = (v4 - (v4 << ((v29) & 63)))
                                        _ = v74
                                    case "foo bar":
                                        for fi75 := 0; fi75 < 19; fi75++ {
                                            v2 -= genFn1(42, true, genFn1(671, true, v2))
                                            v76 := GenS4{F0: 64, F1: v29, F2: (v2 << ((v29) & 63))}
                                            _ = fi75
                                            _ = v76
                                        }
                                    }
                                }
                                _ = fi70
                                _ = v71
                            }
                        case "\n\t":
                            v77 := GenS3{F0: 42, F1: 175, F2: 2}
                            _ = v77
                        default:
                            var v78 uint8 = 185
                            _ = v78
                        }
                    } else {
                        var v79 bool = (true && false)
                        var v80 uint64 = 4294967295
                        v81 := v3[2147483647]
                        _ = v79
                        _ = v80
                        _ = v81
                    }
                }
                for fi82 := 0; fi82 < 5; fi82++ {
                    if true {
                        v83 := GenS5{F0: 0, F1: -2147483648, F2: 0}
                        switch "" {
                        case "hello":
                            switch "" {
                            case "a":
                                var v84 uint = v29
                                _ = v84
                            case "\n\t":
                                var v85 uint16 = 1
                                _ = v85
                            }
                        case "foo bar":
                            v2 -= (^(v2))
                        case "":
                            v86 := GenS6{F0: (false && ((false && (!(0.1 < (1e10 - 3.14159)))) || true))}
                            _ = v86
                        }
                        _ = v83
                    }
                    _ = fi82
                }
            } else {
                v29 += 693
                var v87 uint64 = 18446744073709551615
                switch "hello" {
                case "a":
                    v87 -= (v87 >> ((v29) & 63))
                case "":
                    v88 := []int{127, 442, -327, 42}
                    _ = v88
                case "\n\t":
                    var v89 bool = ((1e10 - (3.14159 / 1e10)) >= (-2.25 * (float64((0.0 * float32(1.5))) / 0.1)))
                    _ = v89
                default:
                    if false {
                        if false {
                            v90 := []int{32767, -369}
                            if true {
                                for fi91 := 0; fi91 < 8; fi91++ {
                                    var v92 uint = v29
                                    _ = fi91
                                    _ = v92
                                }
                            }
                            _ = v90
                        }
                        var v93 float64 = 3.14159
                        _ = v93
                    }
                }
                _ = v87
            }
        default:
            var v94 string = ("" + "\n\t")
            _ = v94
        }
    } else {
        for fi95 := 0; fi95 < 11; fi95++ {
            if (42 != 32767) {
                var v96 uint8 = 0
                v97 := GenS7{F0: v4, F1: v96, F2: 42}
                _ = v96
                _ = v97
            } else {
                switch (fi95 ^ -828) {
                case -2147483648:
                    var v98 int16 = -32768
                    _ = v98
                case 513:
                    var v99 float64 = 0.0
                    _ = v99
                case 9223372036854775807:
                    v100 := []int{9223372036854775807}
                    _ = v100
                default:
                    var v101 uint16 = 171
                    _ = v101
                }
            }
            _ = fi95
        }
    }
    var v102 uint64 = 666
    var v103 int32 = v1
    v102 -= v102
    v104 := GenS8{F0: (v102 & (v102 >> ((4) & 63))), F1: -960}
    v1 += 1
    fmt.Println(v29, v4)
    var v105 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v105
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
