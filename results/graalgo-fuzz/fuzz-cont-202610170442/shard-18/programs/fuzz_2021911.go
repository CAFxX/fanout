package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 string; F1 int16 }
type GenS10 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 uint; F1 string }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool; F1 uint8 }
func genFn0(p0 string, p1 int64) uint64 {
    v6 := []int{127, -2147483648}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 int8 = -1
    p0 = "foo bar"
    for fi8 := 0; fi8 < 7; fi8++ {
        if true {
            v9 := GenS1{F0: 0, F1: (p1 << ((4294967295) & 63))}
            v10 := GenS2{F0: -9223372036854775808}
            v11 := GenS2{F0: (p1 ^ 42)}
            _ = v10
            _ = v11
            _ = v9
        } else {
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS3{F0: p0, F1: 42, F2: (v7 * 100)}
                v6 = append(v6, fi12)
                fmt.Println(fi12, p1, p0)
                _ = fi12
                _ = v13
            }
            p1 = p1
            for fi14 := 0; fi14 < 15; fi14++ {
                if true {
                    var v15 int64 = (p1 + -9223372036854775808)
                    _ = v15
                } else {
                    var v16 float32 = float32(0.0)
                    fmt.Println(p1, v16, v7)
                    if (!true) {
                        var v17 int8 = v7
                        _ = v17
                    }
                    _ = v16
                }
                _ = fi14
            }
        }
        fmt.Println(v7, p1, v7)
        if ((!((float32(1e10) / 3.14159) >= (1.5 - 3.14159))) && true) {
            v7 -= v7
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int) int16 {
    var v18 uint = 18446744073709551615
    var v19 uint8 = 1
    fmt.Println(v18)
    switch "hello" {
    case "a":
        v18 -= 145
    }
    _ = p0
    _ = v18
    _ = v19
    return 42
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 65535
    v3 := []string{"\n\t"}
    var v4 string = "foo bar"
    var v5 float64 = ((1e10 - (0.1 / float64(float32(-2.25)))) + -2.25)
    v5 -= 1e10
    for fi20 := 0; fi20 < 16; fi20++ {
        _ = len(v3)
        for fi21 := 0; fi21 < 13; fi21++ {
            if len(v3) > 0 {
                v3[0] = ""
            }
            _ = fi21
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            switch 767 {
            case 128:
                fmt.Println(v4, fi22, v4)
            case 1:
                var v23 uint8 = 2
                _ = v23
            case 102:
                fmt.Println(fi22)
            default:
                var v24 int16 = 0
                _ = v24
            }
            if false {
                v2 = (v2 ^ 255)
                switch 128 {
                case 127:
                    if ((0.1 / float32(v5)) < v1) {
                        v25 := GenS2{F0: 961}
                        if false {
                            if (42 >= -769) {
                                var v26 int16 = -880
                                switch fi20 {
                                case 32767:
                                    fmt.Println(v2, fi22)
                                    v4 = "foo bar"
                                    if (!(134 < 150)) {
                                        v3 = append(v3, "")
                                        if true {
                                            fmt.Println(v5, v1, fi20)
                                            if true {
                                                _ = len(v3)
                                                v3 = append(v3, "a")
                                                var v27 uint64 = 2
                                                _ = v27
                                            }
                                            switch (fi20 / (((-(fi20))) | 1)) {
                                            case -945:
                                                _ = len(v3)
                                            case 1:
                                                v28 := GenS1{F0: (v2 % (((v2 - v2)) | 1)), F1: 42}
                                                _ = v28
                                            default:
                                                if ((!(0 > 789)) || (42 < 2)) {
                                                    if len(v3) > 0 {
                                                        v3[0] = "hello"
                                                    }
                                                }
                                            }
                                        } else {
                                            v29 := GenS1{F0: (v2 % (((v2 * 2)) | 1)), F1: -1}
                                            var v30 uint = (v2 * 65535)
                                            _ = v29
                                            _ = v30
                                        }
                                    } else {
                                        if true {
                                            v5 -= v5
                                            var v31 uint16 = 32767
                                            _ = v31
                                        } else {
                                            var v32 string = ("a" + v4)
                                            v33 := GenS4{F0: 18446744073709551615, F1: 1, F2: fi22}
                                            switch (v32 + v32) {
                                            case "\n\t":
                                                switch 554 {
                                                case 0:
                                                    v1 = v1
                                                case 896:
                                                    var v34 bool = ((v26 % (((v26 ^ (v26 / ((v26) | 1)))) | 1)) > 1000)
                                                    _ = v34
                                                }
                                            case "foo bar":
                                                v35 := GenS5{F0: v32, F1: 42}
                                                switch "" {
                                                case "hello":
                                                    v4 = v32
                                                    var v36 float32 = v1
                                                    v4 = v32
                                                    _ = v36
                                                }
                                                _ = v35
                                            default:
                                                v5 += v5
                                            }
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                    v1 -= v1
                                    if false {
                                        v3 = append(v3, "a")
                                        v37 := GenS4{F0: 1, F1: 0, F2: 128}
                                        var v38 int16 = v26
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                v26 += v26
                                _ = len(v3)
                                if true {
                                    switch (fi22 | fi22) {
                                    case 128:
                                        if (v2 != (v2 * 2)) {
                                            fmt.Println(v26, v26)
                                            var v39 float64 = -2.25
                                            v5 += 3.14159
                                            _ = v39
                                        } else {
                                            _ = len(v3)
                                        }
                                    default:
                                        if true {
                                            v40 := GenS6{F0: v2}
                                            _ = v40
                                        }
                                    }
                                }
                                _ = v26
                            } else {
                                v41 := GenS7{F0: -63, F1: (fi20 << ((42) & 63))}
                                _ = v41
                            }
                            switch (v4 + "") {
                            case "hello":
                                v3 = append(v3, (v4 + ((v4 + "\n\t") + "foo bar")))
                            case "foo bar":
                                var v42 int16 = 32767
                                _ = v42
                            case "\n\t":
                                switch (fi22 >> ((v2) & 63)) {
                                case -9223372036854775808:
                                    v43 := GenS8{F0: -928, F1: 2}
                                    _ = v43
                                case -2147483648:
                                    v44 := GenS9{F0: v4, F1: 32767}
                                    v45 := GenS10{F0: 65535, F1: 2, F2: 42}
                                    _ = v44
                                    _ = v45
                                default:
                                    if true {
                                        if (v2 > 42) {
                                            v2 = v2
                                            v1 += ((v1 + ((v1 + 3.14159) - 0.0)) / v1)
                                            if false {
                                                if (-26 > 42) {
                                                    var v46 int16 = genFn1(fi20)
                                                    if len(v3) > 0 {
                                                        v3[0] = ("a" + v4)
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v47 uint8 = 2
                                                        _ = v47
                                                    case "a":
                                                        v48 := GenS1{F0: v2, F1: 0}
                                                        _ = v48
                                                    case "\n\t":
                                                        var v49 int16 = (v46 / ((32767) | 1))
                                                        _ = v49
                                                    default:
                                                        v50 := GenS8{F0: 591, F1: 2}
                                                        _ = v50
                                                    }
                                                    _ = v46
                                                }
                                                fmt.Println(v1)
                                                v51 := GenS10{F0: 255, F1: genFn0("", 1), F2: 75}
                                                switch fi20 {
                                                case 32768:
                                                    v5 = v5
                                                case 2147483647:
                                                    var v52 uint32 = 2
                                                    _ = v52
                                                case 42:
                                                    fmt.Println(fi22, fi20)
                                                }
                                                v3 = append(v3, (v4 + "\n\t"))
                                                _ = v51
                                            }
                                        } else {
                                            v1 += 3.14159
                                            fmt.Println(fi22, fi20, v5)
                                            v5 += (-2.25 * v5)
                                            var v53 uint = v2
                                            _ = v53
                                        }
                                    }
                                }
                                switch (fi20 + (fi22 - fi22)) {
                                case 111:
                                    switch ("hello" + "\n\t") {
                                    case "\n\t":
                                        if (((1 < 0) && (v1 == float32(3.14159))) || (0 <= 1)) {
                                            if ((((1e10 + 0.1) * v1) / ((v1 / (((v1 / v1) + 0.0) + ((0.1 + (((v1 + v1) * (0.1 - v1)) - v1)) + 0.0))) - 0.1)) >= (3.14159 - ((v1 - (v1 * v1)) / (v1 + v1)))) {
                                                if (43 == 2147483647) {
                                                    var v54 uint32 = 2
                                                    v55 := GenS11{F0: "foo bar", F1: (v2 >> ((v2) & 63))}
                                                    _ = v54
                                                    _ = v55
                                                } else {
                                                    if false {
                                                        fmt.Println(v5)
                                                    } else {
                                                        var v56 uint64 = 18446744073709551615
                                                        _ = v56
                                                    }
                                                    _ = len(v3)
                                                    if true {
                                                        if (!(false && true)) {
                                                            var v57 int32 = 42
                                                            switch -462 {
                                                            case -499:
                                                                v58 := GenS7{F0: 81, F1: 9223372036854775807}
                                                                _ = v58
                                                            case 128:
                                                                v59 := GenS5{F0: v4, F1: 859}
                                                                v60 := GenS12{F0: 42, F1: v4}
                                                                var v61 uint8 = 128
                                                                _ = v59
                                                                _ = v60
                                                                _ = v61
                                                            default:
                                                                if len(v3) > 0 {
                                                                    v3[0] = v4
                                                                }
                                                            }
                                                            fmt.Println(v4, v4, v5)
                                                            _ = v57
                                                        }
                                                        fmt.Println(v2, fi22)
                                                        if false {
                                                            if len(v3) > 0 {
                                                                v3[0] = (v4 + ((((v4 + v4) + ("hello" + ((((v4 + v4) + v4) + v4) + v4))) + v4) + "hello"))
                                                            }
                                                            v3 = append(v3, "\n\t")
                                                            v62 := GenS8{F0: 2147483647, F1: 101}
                                                            _ = v62
                                                        } else {
                                                            v5 = (-2.25 / ((v5 * 1e10) / 0.1))
                                                            v5 = -2.25
                                                        }
                                                        fmt.Println(fi20)
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = v4
                                                        }
                                                    } else {
                                                        if false {
                                                            v2 = 65535
                                                        }
                                                        v2 = 255
                                                    }
                                                }
                                                switch 0 {
                                                case -90:
                                                    v2 += 65535
                                                case 274:
                                                    switch -1 {
                                                    case -9223372036854775808:
                                                        if (!true) {
                                                            var v63 int8 = 0
                                                            v1 -= v1
                                                            _ = v63
                                                        } else {
                                                            fmt.Println(fi22, v5)
                                                            if false {
                                                                if len(v3) > 0 {
                                                                    v3[0] = "a"
                                                                }
                                                            }
                                                        }
                                                    case 923:
                                                        var v64 uint = 857
                                                        _ = v64
                                                    case 4294967295:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v65 := GenS13{F0: 1}
                                                            _ = v65
                                                        case "foo bar":
                                                            if len(v3) > 0 {
                                                                v3[0] = ""
                                                            }
                                                        }
                                                    default:
                                                        v5 -= v5
                                                    }
                                                }
                                                var v66 int16 = 0
                                                v3 = append(v3, v4)
                                                var v67 uint64 = 26
                                                _ = v66
                                                _ = v67
                                            } else {
                                                _ = len(v3)
                                            }
                                            v68 := GenS5{F0: "", F1: 2147483648}
                                            switch fi22 {
                                            case 0:
                                                switch (v4 + "a") {
                                                case "hello":
                                                    v69 := GenS11{F0: v4, F1: 42}
                                                    _ = v69
                                                case "foo bar":
                                                    fmt.Println(fi22, fi20)
                                                }
                                            case 42:
                                                if true {
                                                    v4 = v4
                                                    if true {
                                                        v70 := GenS6{F0: 470}
                                                        _ = v70
                                                    } else {
                                                        v5 = v5
                                                        v5 = v5
                                                    }
                                                }
                                            }
                                            _ = v68
                                        }
                                    case "":
                                        var v71 float32 = v1
                                        _ = v71
                                    }
                                default:
                                    var v72 int16 = genFn1(42)
                                    _ = v72
                                }
                            }
                            v5 -= (((float64((1e10 * float32((0.1 - float64(v1))))) / v5) + v5) * v5)
                        }
                        v2 += 2
                        _ = v25
                    } else {
                        var v73 uint64 = 2
                        var v74 uint32 = 2147483647
                        _ = v73
                        _ = v74
                    }
                case 0:
                    v75 := GenS2{F0: -9223372036854775808}
                    _ = v75
                }
                var v76 uint64 = 231
                _ = v76
            }
            _ = fi22
        }
        _ = fi20
    }
    for fi77 := 0; fi77 < 5; fi77++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v78 uint = (v2 - (v2 >> ((449) & 63)))
        v78 += (v2 * v78)
        _ = fi77
        _ = v78
    }
    v5 += 1e10
    var v79 float32 = ((1e10 * ((3.14159 * float32(-2.25)) * (-2.25 + (v1 - v1)))) / v1)
    for fi80 := 0; fi80 < 20; fi80++ {
        v81 := GenS5{F0: v4, F1: 482}
        _ = fi80
        _ = v81
    }
    if (v4 >= v4) {
        var v82 uint8 = 231
        _ = v82
    }
    v79 = v1
    v2 += v2
    switch 127 {
    case -1:
        v5 -= (-2.25 - v5)
    case 418:
        switch -2147483648 {
        case 522:
            fmt.Println(v4)
        case 4294967295:
            for fi83 := 0; fi83 < 12; fi83++ {
                fmt.Println(v5, v79)
                for fi84 := 0; fi84 < 20; fi84++ {
                    v3 = append(v3, v4)
                    fmt.Println(v5, fi84)
                    _ = fi84
                }
                v1 -= 0.1
                _ = fi83
            }
        default:
            v2 = (v2 ^ v2)
        }
    }
    for fi85 := 0; fi85 < 2; fi85++ {
        v86 := GenS12{F0: 255, F1: (v4 + v4)}
        _ = fi85
        _ = v86
    }
    if (128 == 127) {
        for fi87 := 0; fi87 < 18; fi87++ {
            if len(v3) > 0 {
                v3[0] = ((((v4 + v4) + v4) + v4) + ("hello" + (((((v4 + (("foo bar" + (v4 + v4)) + v4)) + ((v4 + (v4 + v4)) + v4)) + ("hello" + ((v4 + (v4 + v4)) + v4))) + v4) + "\n\t")))
            }
            fmt.Println(v5, v2)
            _ = fi87
        }
    } else {
        v88 := GenS14{F0: (!((!false) || (486 >= -2147483648))), F1: 0}
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
}
