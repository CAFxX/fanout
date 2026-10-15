package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 int8; F2 string }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 int16 }
type GenS9 struct { F0 int; F1 uint8; F2 uint32 }
func genFn0(p0 bool, p1 uint) int32 {
    var v5 int32 = 42
    v6 := GenS1{F0: (("foo bar" + "foo bar") + "a")}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint32 = 2147483647
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(v5, fi9, p0)
        p0 = p0
        v10 := GenS2{F0: 32767, F1: 2147483648}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 float64 = 0.1
    var v2 int8 = 127
    var v3 int32 = 42
    var v4 int8 = (v2 - 42)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: "\n\t"}
        switch fi11 {
        case 32767:
            v3 -= (v3 & (v3 % (((v3 + v3)) | 1)))
        default:
            v13 := GenS2{F0: 1000, F1: 462}
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    if true {
        var v14 float64 = (0.1 - v1)
        var v15 float64 = 3.14159
        var v16 float32 = 3.14159
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint64 = 2
    for fi18 := 0; fi18 < 5; fi18++ {
        if false {
            for fi19 := 0; fi19 < 20; fi19++ {
                if (!true) {
                    v20 := GenS1{F0: "foo bar"}
                    if false {
                        v21 := GenS3{F0: true, F1: fi18}
                        fmt.Println(v4, v2, fi19)
                        fmt.Println(v2)
                        _ = v21
                    }
                    _ = v20
                } else {
                    var v22 bool = (true && (true || (1 > 2147483648)))
                    _ = v22
                }
                var v23 uint8 = 128
                if false {
                    v24 := []int{fi19, (fi18 << ((4294967295) & 63)), fi19}
                    v4 += 127
                    var v25 float64 = float64(3.14159)
                    _ = v24
                    _ = v25
                } else {
                    var v26 float64 = float64(float32(0.0))
                    if false {
                        v23 -= 128
                        if false {
                            v2 = v4
                            switch "a" {
                            case "hello":
                                if (65535 < 791) {
                                    v27 := GenS3{F0: true, F1: (-(fi18))}
                                    var v28 bool = true
                                    _ = v27
                                    _ = v28
                                }
                                v29 := GenS1{F0: "foo bar"}
                                if false {
                                    fmt.Println(fi19, v4, v17)
                                    v30 := []int{fi19, (fi19 % ((-9223372036854775808) | 1)), 346, fi18}
                                    _ = len(v30)
                                    _ = v30
                                }
                                _ = v29
                            case "a":
                                var v31 bool = false
                                _ = v31
                            }
                            switch "a" {
                            case "foo bar":
                                v23 -= 127
                            case "a":
                                v32 := []int{fi18, -530, fi19, (^(fi18))}
                                _ = v32
                            case "hello":
                                v33 := GenS4{F0: v3}
                                _ = v33
                            default:
                                if false {
                                    v34 := []int{(fi18 - -220), fi19, (fi19 << ((1) & 63)), fi18}
                                    if ((fi18 * 2147483647) > (-(fi19))) {
                                        v35 := GenS5{F0: ("a" + ("a" + (((((("a" + "a") + ("\n\t" + ("" + (("" + ("hello" + "\n\t")) + ("foo bar" + "foo bar"))))) + ("hello" + ("" + "\n\t"))) + "\n\t") + ((("" + "a") + "") + (("hello" + "\n\t") + "hello"))) + ("foo bar" + "hello")))), F1: (^(v4)), F2: "foo bar"}
                                        v3 -= genFn0(((v2 % ((v4) | 1)) == v2), 39)
                                        switch "hello" {
                                        case "foo bar":
                                            var v36 uint8 = v23
                                            _ = v36
                                        case "":
                                            var v37 int8 = (v2 | (v2 + (v4 % ((v4) | 1))))
                                            _ = v37
                                        case "\n\t":
                                            var v38 int32 = (v3 % ((v3) | 1))
                                            _ = v38
                                        default:
                                            v34 = append(v34, fi18)
                                        }
                                        _ = v35
                                    } else {
                                        var v39 int = (fi19 | 794)
                                        if ((^(v3)) < v3) {
                                            v26 += (v1 + v26)
                                            if false {
                                                v26 = (1.5 * ((0.0 + 0.0) + -2.25))
                                                switch "a" {
                                                case "foo bar":
                                                    if true {
                                                        var v40 int16 = 347
                                                        _ = v40
                                                    } else {
                                                        if true {
                                                            fmt.Println(v26, v23, v26)
                                                            v26 += v26
                                                            var v41 int8 = (-(v2))
                                                            _ = v41
                                                        }
                                                        v26 = v26
                                                        fmt.Println(v1, v4, fi18)
                                                    }
                                                case "\n\t":
                                                    var v42 int16 = 1
                                                    _ = v42
                                                default:
                                                    var v43 uint = 76
                                                    _ = v43
                                                }
                                            } else {
                                                _ = len(v34)
                                                v2 -= v2
                                                var v44 int = fi18
                                                _ = v44
                                            }
                                        } else {
                                            var v45 int64 = 400
                                            _ = v45
                                        }
                                        v4 = (^(v2))
                                        _ = v39
                                    }
                                    _ = v34
                                } else {
                                    var v46 uint32 = uint32(v23)
                                    fmt.Println(v26)
                                    var v47 uint = 171
                                    _ = v46
                                    _ = v47
                                }
                            }
                        } else {
                            v4 -= -78
                            v48 := GenS4{F0: (v3 / (((v3 << ((415) & 63))) | 1))}
                            fmt.Println(v17, fi18)
                            _ = v48
                        }
                        fmt.Println(v4, v23)
                        v17 = 42
                        fmt.Println(v3, v2)
                    } else {
                        fmt.Println(v2, v26, v23)
                        var v49 string = "\n\t"
                        _ = v49
                    }
                    var v50 float64 = v26
                    switch "\n\t" {
                    case "\n\t":
                        v51 := []int{(fi18 % ((fi19) | 1)), 4294967295, fi19}
                        _ = v51
                    case "":
                        var v52 int64 = 1
                        _ = v52
                    default:
                        if false {
                            if true {
                                var v53 float32 = 3.14159
                                v17 = 2
                                if false {
                                    v54 := []int{(fi19 ^ 0), fi18, int(v2)}
                                    var v55 int64 = -1
                                    _ = v54
                                    _ = v55
                                }
                                _ = v53
                            } else {
                                v4 = v4
                                if true {
                                    switch fi18 {
                                    case -9223372036854775808:
                                        var v56 float32 = 3.14159
                                        _ = v56
                                    case 32767:
                                        if (!false) {
                                            if false {
                                                v57 := GenS6{F0: v17, F1: (v23 * 229)}
                                                v23 = 255
                                                fmt.Println(v4)
                                                _ = v57
                                            }
                                            if (570 != uint(v23)) {
                                                v4 += (v2 / ((v2) | 1))
                                            } else {
                                                fmt.Println(fi19, v17)
                                                var v58 int32 = (v3 + genFn0(true, 570))
                                                v1 += ((v1 / ((v50 - v26) + v50)) / v50)
                                                _ = v58
                                            }
                                        } else {
                                            v59 := []int{fi19, fi18, fi19}
                                            switch "\n\t" {
                                            case "a":
                                                var v60 uint8 = (v23 + (v23 - (v23 << ((65535) & 63))))
                                                _ = v60
                                            case "hello":
                                                v59 = append(v59, fi19)
                                            }
                                            switch "" {
                                            case "":
                                                if false {
                                                    if false {
                                                        v3 += (v3 ^ (v3 % ((genFn0((false || (!false)), 0)) | 1)))
                                                        var v61 int8 = (v4 - v2)
                                                        _ = v61
                                                    } else {
                                                        var v62 uint = 4294967295
                                                        _ = v62
                                                    }
                                                }
                                            case "\n\t":
                                                v63 := GenS5{F0: "a", F1: v4, F2: ""}
                                                _ = v63
                                            case "a":
                                                v1 = (v1 + v50)
                                            default:
                                                fmt.Println(v23, fi18)
                                            }
                                            _ = v59
                                        }
                                    default:
                                        fmt.Println(v4, v23)
                                    }
                                    fmt.Println(v1)
                                    v23 = v23
                                }
                                if false {
                                    fmt.Println(v1, fi19, v1)
                                    if false {
                                        var v64 int16 = 32767
                                        v1 -= ((v50 * (v50 + v26)) - ((float64((1.5 / 1.0)) * (1e10 - v26)) * (0.1 / (-2.25 / v26))))
                                        _ = v64
                                    }
                                } else {
                                    var v65 string = ""
                                    fmt.Println(v50)
                                    _ = v65
                                }
                            }
                            v50 -= (1.5 * 0.1)
                        } else {
                            v66 := GenS6{F0: (v17 % ((uint64(v23)) | 1)), F1: v23}
                            if false {
                                var v67 uint8 = 145
                                v50 += -2.25
                                v68 := []int{(fi19 * (fi18 << ((42) & 63))), 32768, (fi18 | fi18), fi19}
                                _ = v67
                                _ = v68
                            } else {
                                switch 9223372036854775807 {
                                case 42:
                                    if false {
                                        v69 := []int{(fi18 + fi19), -622, 693, (fi19 % (((fi18 - fi19)) | 1))}
                                        _ = v69
                                    } else {
                                        if ((132 < (v17 / ((581) | 1))) && false) {
                                            v70 := GenS1{F0: "hello"}
                                            v71 := GenS1{F0: "hello"}
                                            _ = v70
                                            _ = v71
                                        } else {
                                            if ((true && false) && false) {
                                                v3 = (^(v3))
                                                v17 = v17
                                            } else {
                                                v72 := []int{fi19}
                                                v2 += (v2 * v2)
                                                fmt.Println(v4, v17, v4)
                                                _ = v72
                                            }
                                            v4 += v4
                                        }
                                        v2 = (v4 ^ -128)
                                        var v73 uint32 = 645
                                        _ = v73
                                    }
                                case 4294967295:
                                    v74 := []int{}
                                    _ = v74
                                case 1:
                                    if true {
                                        switch "\n\t" {
                                        case "a":
                                            v75 := GenS3{F0: ((3.14159 - (float32((v1 * (float64((3.14159 + 1.5)) - (v26 * ((v50 * v50) * 0.1))))) + 0.0)) == (1.5 / 1.5)), F1: fi19}
                                            _ = v75
                                        case "":
                                            switch (fi18 - int(v3)) {
                                            case 280:
                                                v76 := GenS2{F0: 1000, F1: 1}
                                                _ = v76
                                            case -161:
                                                fmt.Println(fi19, v2)
                                            case 32768:
                                                v4 -= int8(v23)
                                            }
                                        default:
                                            v77 := []int{(fi19 ^ -2147483648), fi19, 0}
                                            _ = v77
                                        }
                                        v78 := GenS7{F0: 704}
                                        _ = v78
                                    } else {
                                        var v79 uint32 = 191
                                        var v80 uint16 = 1
                                        _ = v79
                                        _ = v80
                                    }
                                }
                                var v81 float32 = 1.5
                                _ = v81
                            }
                            v2 -= v4
                            var v82 float32 = (3.14159 + 1e10)
                            _ = v66
                            _ = v82
                        }
                    }
                    _ = v26
                    _ = v50
                }
                _ = fi19
                _ = v23
            }
        }
        v4 += v4
        v83 := GenS8{F0: genFn0(((v4 <= (v4 << ((0) & 63))) || (255 <= 18446744073709551615)), 2), F1: 0}
        _ = fi18
        _ = v83
    }
    if true {
        v84 := GenS4{F0: genFn0(true, 255)}
        fmt.Println(v4, v4, v4)
        _ = v84
    }
    var v85 int32 = (v3 + genFn0(false, 915))
    if false {
        if false {
            v86 := GenS9{F0: 0, F1: 2, F2: 2147483647}
            _ = v86
        } else {
            v85 += (v85 + v3)
        }
    }
    if true {
        switch (("\n\t" + ("hello" + "foo bar")) + "\n\t") {
        case "":
            v85 -= v85
        case "a":
            fmt.Println(v1)
        default:
            if false {
                var v87 int64 = 0
                v88 := []int{127}
                _ = v87
                _ = v88
            }
        }
    }
    if (true || (((((0.1 - 3.14159) - v1) / v1) * float64(0.1)) >= (0.1 + (((v1 + ((-2.25 * float64(-2.25)) + 0.1)) - (-2.25 * v1)) / v1)))) {
        for fi89 := 0; fi89 < 4; fi89++ {
            var v90 int = fi89
            var v91 uint32 = 457
            if true {
                fmt.Println(v85, v85)
                v92 := []int{}
                _ = v92
            } else {
                v91 = (^(v91))
                v93 := []int{(v90 >> ((2) & 63)), v90}
                _ = v93
            }
            _ = fi89
            _ = v90
            _ = v91
        }
    } else {
        var v94 float64 = ((v1 - v1) - (float64((float32((1e10 * v1)) + ((-2.25 / float32((3.14159 + ((v1 / (v1 - 1.5)) * ((((1e10 - (v1 / float64(-2.25))) - v1) / float64(3.14159)) / (0.0 + v1)))))) * (-2.25 - (-2.25 / 1.0))))) - float64(1e10)))
        switch (("foo bar" + "a") + "\n\t") {
        case "":
            v95 := []int{4294967295, 32768}
            _ = v95
        default:
            var v96 int32 = (-(v3))
            _ = v96
        }
        _ = v94
    }
    for fi97 := 0; fi97 < 16; fi97++ {
        switch fi97 {
        case 42:
            var v98 int = fi97
            _ = v98
        }
        _ = fi97
    }
    fmt.Println(v85)
    v4 += v2
    fmt.Println(v17, v4)
    v17 = v17
    if (((!false) && (542 <= 65535)) && false) {
        if (v2 <= 100) {
            v99 := []int{1, -533}
            _ = v99
        }
        if (4294967295 <= 236) {
            v17 += (v17 >> ((643) & 63))
            if true {
                fmt.Println(v17)
                v100 := []int{0, 32768, 42, 0}
                v85 += genFn0((99 < 127), 4294967295)
                _ = v100
            } else {
                switch 32767 {
                case 32768:
                    var v101 float64 = v1
                    _ = v101
                case 128:
                    var v102 uint = 0
                    _ = v102
                case 9223372036854775807:
                    var v103 int32 = -499
                    _ = v103
                default:
                    v104 := []int{}
                    _ = v104
                }
            }
            var v105 uint8 = 90
            _ = v105
        }
        var v106 uint8 = 0
        _ = v106
    }
    switch (("a" + "") + "\n\t") {
    case "hello":
        if false {
            v17 -= v17
            var v107 float64 = 1.5
            for fi108 := 0; fi108 < 12; fi108++ {
                var v109 uint = 699
                _ = fi108
                _ = v109
            }
            _ = v107
        }
    case "foo bar":
        for fi110 := 0; fi110 < 11; fi110++ {
            switch (fi110 % (((^(fi110))) | 1)) {
            case 32767:
                var v111 uint8 = 1
                _ = v111
            default:
                v112 := GenS7{F0: 42}
                _ = v112
            }
            if ("\n\t" == "") {
                v113 := []int{(fi110 >> ((0) & 63)), fi110, fi110, (fi110 >> ((998) & 63))}
                _ = v113
            }
            _ = fi110
        }
    default:
        v114 := []int{-72, 4294967295, 2147483647}
        _ = v114
    }
    switch ("foo bar" + "\n\t") {
    case "":
        v17 = (v17 >> ((0) & 63))
    case "foo bar":
        v1 = 0.0
    case "hello":
        switch -809 {
        case -2147483648:
            v4 += 127
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v85
}
