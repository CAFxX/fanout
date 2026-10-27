package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 string; F1 uint64; F2 int }
type GenS8 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 uint; F1 int32; F2 int }
func genFn0(p0 string, p1 uint64, p2 uint8) int16 {
    if (true && (127 <= 92)) {
        p0 = ""
    } else {
        var v7 uint64 = 4294967295
        var v8 float32 = 0.0
        for fi9 := 0; fi9 < 1; fi9++ {
            if true {
                p1 = v7
                var v10 string = "foo bar"
                _ = v10
            } else {
                p2 += (^(p2))
                var v11 int = (^(fi9))
                _ = v11
            }
            p0 = p0
            if (uint16(p1) == 1) {
                fmt.Println(v7)
                v12 := GenS1{F0: 42, F1: (v7 & 2), F2: -9223372036854775808}
                _ = v12
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p0)
    v13 := []int{-1}
    var v14 float32 = 1e10
    v15 := GenS1{F0: 127, F1: (p1 << ((uint(p2)) & 63)), F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return -1
}
func genFn1() int8 {
    var v16 string = (((("\n\t" + "\n\t") + "\n\t") + ("hello" + "hello")) + ("hello" + ("\n\t" + "a")))
    if false {
        if (255 > 2) {
            if (42 < 112) {
                var v17 float64 = (((1.5 * 1e10) + (float64(0.0) - 0.0)) / (1.5 + 1.5))
                var v18 uint64 = 449
                _ = v17
                _ = v18
            }
        }
        switch 32767 {
        case 0:
            fmt.Println(v16)
        case -949:
            var v19 uint64 = 0
            _ = v19
        default:
            var v20 int = 2147483647
            _ = v20
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            fmt.Println(fi21)
            var v22 int64 = 0
            fmt.Println(v22, fi21, fi21)
            _ = fi21
            _ = v22
        }
    } else {
        fmt.Println(v16)
        var v23 float32 = 1.5
        _ = v23
    }
    var v24 float64 = float64((0.0 + (float32(1e10) + (((0.0 * 1.5) * 0.1) / (1.5 / (0.0 - (3.14159 - 1e10)))))))
    if (!true) {
        v25 := []int{-611, 1, 128, 1}
        var v26 int16 = 32767
        _ = v25
        _ = v26
    }
    _ = v16
    _ = v24
    return 100
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = -116
    var v3 int16 = 1
    var v4 float32 = 1.5
    v5 := []int{}
    v6 := make(map[uint32]int)
    fmt.Println(v3, v1)
    for fi27 := 0; fi27 < 13; fi27++ {
        if len(v5) > 0 {
            v5[0] = (fi27 - (-(fi27)))
        }
        var v28 uint64 = (v1 ^ v1)
        _ = fi27
        _ = v28
    }
    if (v1 == 2) {
        var v29 int16 = 247
        _ = v29
    }
    if (42 != 0) {
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 int32 = -2147483648
            v32 := GenS1{F0: v2, F1: (^(v1)), F2: -1}
            if (!(!(("\n\t" != "a") || false))) {
                switch fi30 {
                case -743:
                    v33 := GenS2{F0: -778}
                    _ = v33
                case -817:
                    var v34 int16 = v3
                    _ = v34
                }
                for fi35 := 0; fi35 < 15; fi35++ {
                    fmt.Println(v2, fi30, v31)
                    _ = len(v5)
                    _ = fi35
                }
            }
            _ = fi30
            _ = v31
            _ = v32
        }
        v36 := GenS1{F0: 43, F1: v1, F2: -37}
        fmt.Println(v4)
        _ = v36
    }
    var v37 int32 = 2147483647
    v38 := GenS3{F0: ("foo bar" >= "foo bar")}
    for fi39 := 0; fi39 < 12; fi39++ {
        if (26 <= v2) {
            for fi40 := 0; fi40 < 3; fi40++ {
                v41 := GenS2{F0: 695}
                switch int(v2) {
                case 1:
                    v42 := GenS2{F0: -9223372036854775808}
                    _ = v42
                case 127:
                    if (v37 == (v37 & (^(v37)))) {
                        switch "" {
                        case "":
                            if true {
                                v43 := v6[2]
                                var v44 int16 = int16(v37)
                                v45 := GenS3{F0: (true || false)}
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                if true {
                                    v46 := GenS2{F0: 9223372036854775807}
                                    _ = v46
                                }
                            }
                            switch ("" + "foo bar") {
                            case "\n\t":
                                fmt.Println(fi40, fi39)
                            case "hello":
                                if true {
                                    v47 := GenS3{F0: (802 < 264)}
                                    _ = v47
                                } else {
                                    v2 = 67
                                    v48 := GenS3{F0: ((2 <= 255) && (fi40 >= 145))}
                                    fmt.Println(v4, fi40)
                                    fmt.Println(fi40)
                                    fmt.Println(v4, fi39)
                                    _ = v48
                                }
                            }
                            if (-677 < 9223372036854775807) {
                                var v49 float64 = ((0.0 + 3.14159) / 1.0)
                                if false {
                                    var v50 int16 = int16(v1)
                                    var v51 uint16 = 2
                                    switch "foo bar" {
                                    case "foo bar":
                                        v49 -= (v49 - (1.5 * (v49 / 0.1)))
                                    case "\n\t":
                                        var v52 uint32 = 4294967295
                                        _ = v52
                                    }
                                    v1 -= (v1 >> ((823) & 63))
                                    if false {
                                        v4 = v4
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = (fi39 % ((fi39) | 1))
                                        }
                                    }
                                    _ = v50
                                    _ = v51
                                }
                                _ = v49
                            } else {
                                var v53 uint8 = 109
                                v2 = 127
                                v37 = v37
                                var v54 bool = true
                                _ = v53
                                _ = v54
                            }
                            if false {
                                fmt.Println(v2)
                                var v55 uint64 = (^(v1))
                                v56 := GenS4{F0: fi39}
                                _ = v55
                                _ = v56
                            }
                        case "hello":
                            delete(v6, 0)
                        case "\n\t":
                            switch "\n\t" {
                            case "hello":
                                var v57 float64 = -2.25
                                _ = v57
                            }
                        default:
                            v58 := GenS4{F0: fi40}
                            _ = v58
                        }
                        v59 := GenS5{F0: 2}
                        _ = v59
                    } else {
                        if (!(false || (1.5 < (0.0 * ((0.0 * -2.25) - -2.25))))) {
                            if false {
                                switch "hello" {
                                case "hello":
                                    var v60 uint64 = (v1 ^ (v1 - v1))
                                    _ = v60
                                case "a":
                                    v6[1] = fi40
                                case "":
                                    delete(v6, 65535)
                                }
                            } else {
                                delete(v6, 640)
                                if (fi40 != (fi39 | 403)) {
                                    var v61 uint32 = 2
                                    v5 = append(v5, (-(fi39)))
                                    _ = v61
                                } else {
                                    v1 += v1
                                }
                            }
                            if true {
                                v3 = v3
                            } else {
                                fmt.Println(v3, v3)
                            }
                        } else {
                            v3 = 405
                            if true {
                                var v62 uint64 = (v1 ^ (v1 << ((4294967295) & 63)))
                                v4 -= 1.5
                                v62 += v1
                                _ = v62
                            } else {
                                v1 -= (^(v1))
                                v63 := GenS2{F0: -1}
                                _ = v63
                            }
                        }
                        v2 += genFn1()
                        _ = len(v6)
                    }
                case -9223372036854775808:
                    v64 := GenS6{F0: genFn0(("\n\t" + ("hello" + "a")), (v1 + (^(v1))), 42), F1: (v3 / ((1000) | 1))}
                    _ = v64
                default:
                    if (-128 > (v2 - (v2 | v2))) {
                        fmt.Println(v2, v3)
                        if true {
                            v6[347] = (fi39 | (fi40 * fi39))
                            v65 := GenS1{F0: v2, F1: (v1 / (((v1 % ((742) | 1))) | 1)), F2: -952}
                            var v66 int8 = v2
                            _ = v65
                            _ = v66
                        }
                    } else {
                        v67 := GenS7{F0: "a", F1: 4294967295, F2: (fi39 / (((fi39 * 4294967295)) | 1))}
                        v68 := GenS1{F0: (v2 % (((-(v2))) | 1)), F1: 2, F2: 42}
                        v4 = ((float32(0.1) / v4) / 1.0)
                        _ = v67
                        _ = v68
                    }
                }
                fmt.Println(v3, v2)
                var v69 int32 = (v37 >> ((0) & 63))
                _ = fi40
                _ = v41
                _ = v69
            }
        } else {
            v1 += 561
            v5 = append(v5, 2147483647)
        }
        _ = fi39
    }
    switch ("\n\t" + ("a" + ("foo bar" + ("" + "a")))) {
    case "":
        switch ("hello" + (("a" + "hello") + ("hello" + "\n\t"))) {
        case "a":
            var v70 int16 = (v3 >> ((18446744073709551615) & 63))
            _ = v70
        case "\n\t":
            for fi71 := 0; fi71 < 9; fi71++ {
                if (v37 <= 42) {
                    for fi72 := 0; fi72 < 2; fi72++ {
                        fmt.Println(fi72, v4, v2)
                        v73 := GenS7{F0: "foo bar", F1: 252, F2: (^(fi72))}
                        if (v3 < -410) {
                            if false {
                                if true {
                                    if (!true) {
                                        if true {
                                            v37 = v37
                                            var v74 float64 = 1.5
                                            if ((("" + "hello") + "foo bar") != (("a" + "hello") + "")) {
                                                if false {
                                                    v75 := GenS6{F0: v3, F1: genFn0(("hello" + "foo bar"), 394, 255)}
                                                    delete(v6, 65535)
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v75
                                                } else {
                                                    v4 += v4
                                                    v6[2] = (^(fi71))
                                                    if (!true) {
                                                        v2 += (v2 % ((0) | 1))
                                                        var v76 bool = true
                                                        _ = v76
                                                    }
                                                }
                                            }
                                            _ = v74
                                        } else {
                                            switch ("a" + "") {
                                            case "foo bar":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            case "":
                                                var v77 int16 = genFn0("", v1, 42)
                                                _ = v77
                                            case "\n\t":
                                                var v78 int64 = -65
                                                _ = v78
                                            }
                                            v37 = (-(v37))
                                        }
                                        v4 += (1.5 - 1e10)
                                        v3 += genFn0(("hello" + ("\n\t" + "hello")), (v1 + 18446744073709551615), 128)
                                    }
                                    v6[1] = fi72
                                }
                                v79 := GenS1{F0: v2, F1: v1, F2: 0}
                                _ = v79
                            } else {
                                v80 := GenS8{F0: 89, F1: 2147483647, F2: 255}
                                switch ("hello" + "hello") {
                                case "a":
                                    var v81 float32 = float32((0.0 / (1e10 * -2.25)))
                                    _ = v81
                                default:
                                    if (true || false) {
                                        var v82 float64 = ((float64(float32(((float64(v4) + 0.1) * (float64(0.0) - (-2.25 / 1.0))))) - 0.1) + -2.25)
                                        _ = v82
                                    }
                                }
                                _ = v80
                            }
                            if (!true) {
                                _ = len(v6)
                                v2 += genFn1()
                                fmt.Println(fi72, fi72, v1)
                                v3 += v3
                                var v83 uint32 = 1
                                _ = v83
                            }
                        } else {
                            var v84 uint = 1
                            var v85 uint = (v84 * 65535)
                            v86 := v6[42]
                            _ = v84
                            _ = v85
                            _ = v86
                        }
                        if (-941 >= 42) {
                            var v87 uint = 835
                            _ = v87
                        }
                        var v88 uint8 = 1
                        _ = fi72
                        _ = v73
                        _ = v88
                    }
                    v89 := GenS3{F0: true}
                    _ = v89
                } else {
                    var v90 uint64 = (v1 << ((18446744073709551615) & 63))
                    _ = v90
                }
                var v91 int8 = genFn1()
                for fi92 := 0; fi92 < 5; fi92++ {
                    v4 += (v4 / v4)
                    if len(v5) > 0 {
                        v5[0] = (fi92 * (fi92 ^ (fi92 % ((int(v2)) | 1))))
                    }
                    v4 -= v4
                    _ = fi92
                }
                _ = fi71
                _ = v91
            }
        }
    case "hello":
        for fi93 := 0; fi93 < 18; fi93++ {
            v4 = -2.25
            var v94 uint = 255
            _ = fi93
            _ = v94
        }
    default:
        v95 := GenS9{F0: 42, F1: 1}
        _ = v95
    }
    for fi96 := 0; fi96 < 5; fi96++ {
        v97 := GenS10{F0: 568, F1: 634}
        for fi98 := 0; fi98 < 5; fi98++ {
            _ = len(v6)
            var v99 uint8 = 10
            var v100 string = "hello"
            _ = fi98
            _ = v100
            _ = v99
        }
        v4 += (3.14159 * float32(0.1))
        _ = fi96
        _ = v97
    }
    if false {
        delete(v6, 42)
    }
    if false {
        for fi101 := 0; fi101 < 1; fi101++ {
            v6[0] = fi101
            var v102 float64 = 3.14159
            fmt.Println(v2, fi101)
            _ = fi101
            _ = v102
        }
        v103 := GenS11{F0: 520, F1: (-(v37)), F2: 127}
        _ = len(v6)
        _ = v103
    }
    v3 -= (v3 + genFn0("\n\t", (v1 / (((v1 % (((v1 >> ((1) & 63))) | 1))) | 1)), 0))
    var v104 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
