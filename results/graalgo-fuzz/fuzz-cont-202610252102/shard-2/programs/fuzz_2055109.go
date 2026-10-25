package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8; F1 string; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 int64 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 int16 }
type GenS10 struct { F0 bool; F1 uint8; F2 string }
type GenS11 struct { F0 int16 }
func genFn0(p0 int16) uint {
    if false {
        p0 = (p0 | (p0 - -146))
        v6 := []int{}
        if (-2147483648 < 2147483647) {
            var v7 uint8 = 42
            var v8 uint32 = 713
            var v9 int = 2147483647
            _ = v7
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 2; fi10++ {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 string = "hello"
                    _ = fi11
                    _ = v12
                }
                if true {
                    var v13 int16 = -1
                    _ = v13
                } else {
                    if len(v6) > 0 {
                        v6[0] = fi10
                    }
                }
                _ = fi10
            }
        }
        _ = v6
    } else {
        p0 = p0
        var v14 int32 = 0
        _ = v14
    }
    v15 := GenS1{F0: (p0 << ((42) & 63)), F1: 18446744073709551615}
    fmt.Println(p0)
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint8 = 128
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v15
    return 775
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 4294967295
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    var v5 int8 = -1
    for fi18 := 0; fi18 < 4; fi18++ {
        if (false && false) {
            if (1 >= 42) {
                v19 := GenS1{F0: -1, F1: v2}
                v20 := v1[-2147483648]
                var v21 uint = 107
                _ = v19
                _ = v20
                _ = v21
            } else {
                _ = len(v1)
                v22 := GenS2{F0: 140, F1: ("foo bar" + ("" + "hello")), F2: 64}
                fmt.Println(v5)
                _ = v22
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 string = "hello"
                if (false && (1 < 928)) {
                    v25 := v1[(fi23 + fi23)]
                    _ = v25
                } else {
                    fmt.Println(fi18, v2, fi23)
                    var v26 string = "\n\t"
                    _ = v26
                }
                if (18446744073709551615 != 18446744073709551615) {
                    var v27 uint = v2
                    v28 := []int{-2147483648}
                    v27 = v27
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v2, fi23)
                    v5 = -1
                }
                v29 := GenS3{F0: 255, F1: (fi18 * (fi23 % ((-29) | 1)))}
                var v30 float64 = 3.14159
                _ = fi23
                _ = v24
                _ = v29
                _ = v30
            }
        } else {
            v31 := GenS3{F0: 65535, F1: fi18}
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 int64 = 9223372036854775807
                if false {
                    v34 := []int{(fi32 % ((int(v2)) | 1)), fi32, (fi18 - fi18)}
                    _ = v34
                } else {
                    v33 = v33
                }
                _ = fi32
                _ = v33
            }
            _ = v31
        }
        _ = fi18
    }
    delete(v1, 4294967295)
    v2 = v2
    var v35 int32 = 1
    var v36 uint64 = 372
    v36 += (v36 - (^(v36)))
    v37 := v4[2147483648]
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 uint32 = 1
        if v37 {
            var v40 int8 = (v5 * 0)
            _ = v40
        } else {
            var v41 int32 = v35
            v39 = (^(v39))
            var v42 float32 = float32(1e10)
            _ = v41
            _ = v42
        }
        _ = fi38
        _ = v39
    }
    var v43 int8 = 0
    for fi44 := 0; fi44 < 1; fi44++ {
        v45 := []int{}
        if (genFn0(42) <= genFn0(-779)) {
            if len(v45) > 0 {
                v45[0] = (fi44 % ((2147483647) | 1))
            }
            if false {
                v37 = v37
            }
        }
        v37 = (772 > 2)
        _ = fi44
        _ = v45
    }
    v37 = (("\n\t" + "") != "\n\t")
    if (!v37) {
        if (!(!false)) {
            v36 += 2
            var v46 int8 = v43
            v2 += (v2 >> (((^(v2))) & 63))
            _ = v46
        }
        switch -2147483648 {
        case -2147483648:
            v47 := GenS3{F0: 2, F1: 2147483647}
            _ = v47
        case -711:
            v48 := GenS3{F0: 255, F1: 32768}
            _ = v48
        case 4294967295:
            v49 := GenS1{F0: 32767, F1: (v2 >> (((v2 & (v2 + (v2 / ((1) | 1))))) & 63))}
            _ = v49
        default:
            for fi50 := 0; fi50 < 9; fi50++ {
                var v51 uint = v2
                v36 = v36
                if v37 {
                    var v52 float32 = -2.25
                    v53 := []int{}
                    _ = v52
                    _ = v53
                }
                _ = fi50
                _ = v51
            }
        }
    } else {
        if (v37 && v37) {
            for fi54 := 0; fi54 < 5; fi54++ {
                v37 = false
                for fi55 := 0; fi55 < 10; fi55++ {
                    if ((3.14159 == (0.0 + ((float32((3.14159 / 1e10)) + ((((-2.25 * 3.14159) * ((0.0 / -2.25) * (3.14159 / -2.25))) / (((0.1 / 1.0) - (0.0 - float32(((float64(float32(0.0)) * 3.14159) + (-2.25 - float64(0.1)))))) + (float32((1.5 / 1e10)) * 3.14159))) / -2.25)) * (1.5 / (1.5 * (1.5 + 3.14159)))))) || (!(!(!(9223372036854775807 >= 155))))) {
                        v43 = v5
                        v56 := v3[-9223372036854775808]
                        if v37 {
                            v57 := GenS3{F0: 65535, F1: 127}
                            fmt.Println(v56)
                            fmt.Println(fi55, v2, v37)
                            _ = v57
                        }
                        _ = v56
                    } else {
                        v58 := GenS3{F0: 200, F1: (fi55 * (fi54 ^ 42))}
                        v2 = v2
                        _ = v58
                    }
                    delete(v3, fi55)
                    _ = fi55
                }
                for fi59 := 0; fi59 < 17; fi59++ {
                    v37 = (true && ((^(v43)) != (-(v43))))
                    _ = fi59
                }
                _ = fi54
            }
            for fi60 := 0; fi60 < 19; fi60++ {
                var v61 uint8 = 128
                if ((false && true) || (v61 >= (^(v61)))) {
                    _ = len(v4)
                }
                _ = fi60
                _ = v61
            }
        } else {
            if false {
                v62 := []int{201}
                _ = v62
            } else {
                delete(v4, 2)
                v63 := []int{128, 32768}
                for fi64 := 0; fi64 < 12; fi64++ {
                    if v37 {
                        if false {
                            var v65 uint64 = 1
                            _ = v65
                        } else {
                            switch (("hello" + "a") + "") {
                            case "\n\t":
                                if (v37 && (!v37)) {
                                    for fi66 := 0; fi66 < 20; fi66++ {
                                        v43 = 100
                                        if v37 {
                                            v37 = ((!(!false)) || (v37 && true))
                                            v37 = (128 != 23)
                                            if true {
                                                v37 = v37
                                            } else {
                                                v63 = append(v63, fi66)
                                                if (604 != 0) {
                                                    if (-286 >= v35) {
                                                        if (true && (!(1.5 == ((1.5 / 1.5) / 0.1)))) {
                                                            if (float32(1e10) > float32(0.0)) {
                                                                var v67 bool = (((float64(1e10) - 1e10) + (((1.5 * (3.14159 * -2.25)) - -2.25) - 3.14159)) > 3.14159)
                                                                v68 := GenS3{F0: 2, F1: (fi64 % ((-199) | 1))}
                                                                v69 := GenS4{F0: 1}
                                                                v70 := v3[(fi66 << ((v2) & 63))]
                                                                _ = v67
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            } else {
                                                                if (-1 >= 70) {
                                                                    if false {
                                                                        v36 = (^(v36))
                                                                        var v71 int8 = 42
                                                                        if v37 {
                                                                            v5 -= (v5 + (v71 * v71))
                                                                            switch fi64 {
                                                                            case 42:
                                                                                if len(v63) > 0 {
                                                                                    v63[0] = fi66
                                                                                }
                                                                            case -758:
                                                                                _ = len(v4)
                                                                            default:
                                                                                delete(v4, 65535)
                                                                            }
                                                                            fmt.Println(fi66, v35, v37)
                                                                            var v72 uint16 = 65535
                                                                            if true {
                                                                                fmt.Println(v72, v71)
                                                                                v72 = 2
                                                                            } else {
                                                                                fmt.Println(v72, v71, v2)
                                                                            }
                                                                            _ = v72
                                                                        }
                                                                        switch ("hello" + "hello") {
                                                                        case "foo bar":
                                                                            var v73 float64 = 1.5
                                                                            _ = v73
                                                                        case "hello":
                                                                            switch ((("\n\t" + (("hello" + ("" + ("hello" + "\n\t"))) + "a")) + ("foo bar" + "foo bar")) + (("\n\t" + "foo bar") + (((("" + "") + "\n\t") + "foo bar") + ("a" + "\n\t")))) {
                                                                            case "a":
                                                                                v43 = v5
                                                                            }
                                                                        case "":
                                                                            _ = len(v63)
                                                                        }
                                                                        _ = v71
                                                                    }
                                                                    v74 := GenS5{F0: v35, F1: (v43 / (((^(v43))) | 1)), F2: 1}
                                                                    v63 = append(v63, (fi64 << ((v2) & 63)))
                                                                    _ = v74
                                                                }
                                                            }
                                                            if v37 {
                                                                if v37 {
                                                                    if ((fi64 / (((fi66 - (fi64 & fi66))) | 1)) <= fi66) {
                                                                        var v75 uint16 = 126
                                                                        v76 := GenS6{F0: (v2 & v2), F1: 875}
                                                                        v77 := GenS7{F0: (v37 || v37)}
                                                                        fmt.Println(v36)
                                                                        _ = v75
                                                                        _ = v76
                                                                        _ = v77
                                                                    } else {
                                                                        v35 = 627
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi64)
                                                                    var v78 int32 = (v35 << (((v2 * v2)) & 63))
                                                                    if (2147483647 <= 957) {
                                                                        v5 = -69
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            v2 += (v2 % ((v2) | 1))
                                                                        case "\n\t":
                                                                            v2 += 0
                                                                        case "hello":
                                                                            switch "" {
                                                                            case "a":
                                                                                v79 := GenS7{F0: ((fi66 | -9223372036854775808) >= (^(fi66)))}
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v78
                                                                }
                                                            }
                                                            var v80 uint = 2
                                                            v80 = (v80 % ((65535) | 1))
                                                            v37 = true
                                                            v81 := v3[int(v5)]
                                                            v43 = (v43 * -128)
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            if v37 {
                                                                fmt.Println(v35, fi64, v2)
                                                                v2 = genFn0(-32768)
                                                            }
                                                            var v82 uint64 = (v36 & v36)
                                                            _ = v82
                                                        }
                                                        var v83 uint = (v2 + (v2 * v2))
                                                        if (!true) {
                                                            if (1e10 >= 1.5) {
                                                                fmt.Println(v5)
                                                                var v84 int64 = 9223372036854775807
                                                                v85 := GenS1{F0: -1, F1: (^(v2))}
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                            v35 = v35
                                                            v43 += 42
                                                            fmt.Println(fi64)
                                                        }
                                                        _ = v83
                                                    } else {
                                                        v86 := v3[(fi66 | fi64)]
                                                        _ = v86
                                                    }
                                                    fmt.Println(v43, fi64, v36)
                                                    if v37 {
                                                        fmt.Println(v2, fi64, v36)
                                                        v37 = true
                                                    }
                                                }
                                            }
                                        }
                                        if v37 {
                                            switch fi64 {
                                            case 4294967295:
                                                var v87 float64 = 1.5
                                                _ = v87
                                            default:
                                                if v37 {
                                                    var v88 int8 = (-(v43))
                                                    v35 -= 42
                                                    _ = v88
                                                }
                                            }
                                            v35 = v35
                                            v89 := GenS6{F0: 2, F1: 0}
                                            _ = v89
                                        } else {
                                            v90 := GenS8{F0: "\n\t"}
                                            if (!true) {
                                                v91 := GenS9{F0: v35, F1: 32767}
                                                v2 -= 1
                                                v92 := GenS9{F0: v35, F1: 702}
                                                var v93 int = -945
                                                var v94 int16 = 411
                                                _ = v91
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            }
                                            var v95 uint = v2
                                            _ = v90
                                            _ = v95
                                        }
                                        switch 32768 {
                                        case -379:
                                            _ = len(v63)
                                        case -394:
                                            v96 := GenS10{F0: (59 <= 2), F1: 128, F2: "\n\t"}
                                            _ = v96
                                        }
                                        _ = fi66
                                    }
                                    var v97 uint16 = 65535
                                    v63 = append(v63, fi64)
                                    _ = v97
                                }
                            case "a":
                                delete(v1, (-(fi64)))
                            case "":
                                for fi98 := 0; fi98 < 9; fi98++ {
                                    if false {
                                        switch ("hello" + "a") {
                                        case "\n\t":
                                            var v99 uint8 = 228
                                            _ = v99
                                        default:
                                            delete(v3, (fi64 * -1))
                                        }
                                        v100 := GenS11{F0: 32767}
                                        _ = v100
                                    } else {
                                        v101 := GenS1{F0: -381, F1: genFn0(563)}
                                        _ = v101
                                    }
                                    fmt.Println(v2, v43)
                                    var v102 int16 = 32767
                                    _ = fi98
                                    _ = v102
                                }
                            }
                        }
                        v103 := GenS11{F0: 42}
                        _ = v103
                    }
                    v5 += (^(v43))
                    var v104 string = "\n\t"
                    _ = fi64
                    _ = v104
                }
                _ = v63
            }
            var v105 int32 = -608
            v106 := []int{42, -980, 383, 32768}
            _ = v105
            _ = v106
        }
        if v37 {
            if (("" + (("hello" + ((("\n\t" + (("a" + ("foo bar" + "a")) + "hello")) + "") + "hello")) + "foo bar")) <= "a") {
                fmt.Println(v43, v5)
            }
        } else {
            if (v37 && (42 > 0)) {
                v107 := v4[1]
                for fi108 := 0; fi108 < 3; fi108++ {
                    v2 = genFn0(int16(fi108))
                    _ = fi108
                }
                _ = v107
            } else {
                var v109 uint32 = 0
                if ((v35 + v35) == v35) {
                    v110 := GenS8{F0: ("\n\t" + "")}
                    _ = v110
                }
                _ = v109
            }
            switch -215 {
            case 0:
                v2 = v2
            }
            switch -9223372036854775808 {
            case 593:
                fmt.Println(v43)
            case 32767:
                if true {
                    var v111 string = ("a" + "\n\t")
                    fmt.Println(v35, v35, v35)
                    fmt.Println(v36)
                    _ = v111
                }
            case 128:
                var v112 int64 = -9223372036854775808
                _ = v112
            default:
                v1[1] = "a"
            }
        }
    }
    v113 := GenS8{F0: ""}
    for fi114 := 0; fi114 < 18; fi114++ {
        var v115 float32 = 1.5
        switch fi114 {
        case 127:
            if v37 {
                if (519 >= 1) {
                    v35 = (v35 % (((v35 ^ v35)) | 1))
                    v116 := GenS3{F0: 65535, F1: 176}
                    _ = v116
                }
                var v117 uint32 = 65535
                _ = v117
            }
        default:
            if (0.1 == v115) {
                var v118 uint32 = 223
                v2 += (v2 ^ (v2 << ((v2) & 63)))
                v43 += v5
                _ = v118
            }
        }
        _ = fi114
        _ = v115
    }
    fmt.Println(v5, v43)
    fmt.Println(v35, v5, v43)
    switch 2147483647 {
    case 611:
        fmt.Println(v2, v36)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v113
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v43
    _ = v5
}
