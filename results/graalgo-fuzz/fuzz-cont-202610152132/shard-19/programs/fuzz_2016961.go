package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint8; F2 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool; F1 uint16; F2 string }
func genFn0(p0 int16, p1 int16, p2 string) int64 {
    v7 := []int{42, 4294967295}
    for fi8 := 0; fi8 < 12; fi8++ {
        p2 = p2
        for fi9 := 0; fi9 < 9; fi9++ {
            if (0 <= 1) {
                if true {
                    switch (fi9 * fi9) {
                    case 127:
                        v10 := GenS1{F0: false, F1: -35, F2: 1}
                        _ = v10
                    case 885:
                        var v11 int = fi8
                        _ = v11
                    case 0:
                        var v12 float32 = float32(3.14159)
                        _ = v12
                    }
                    if ((1.5 / ((1.5 + float64((float32(0.1) - 0.0))) + 3.14159)) == ((1e10 / (1e10 / 0.1)) + 1e10)) {
                        var v13 uint = 0
                        _ = v13
                    } else {
                        if true {
                            p2 = ("hello" + "")
                            v7 = append(v7, (fi8 / ((42) | 1)))
                            var v14 uint = 964
                            _ = v14
                        } else {
                            var v15 bool = true
                            v7 = append(v7, fi8)
                            if v15 {
                                v16 := GenS1{F0: false, F1: 1, F2: 66}
                                _ = len(v7)
                                v17 := GenS2{F0: 42, F1: -795}
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                _ = v16
                                _ = v17
                            }
                            switch "" {
                            case "foo bar":
                                if ((!v15) || (((1 < 1) && false) && v15)) {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    switch p2 {
                                    case "foo bar":
                                        var v18 float32 = ((0.1 / 1.0) / 1.0)
                                        _ = v18
                                    case "\n\t":
                                        fmt.Println(v15, p1)
                                    case "a":
                                        p1 += (^(p0))
                                    }
                                    fmt.Println(p2, v15, p2)
                                } else {
                                    v15 = v15
                                    var v19 int = (fi8 >> ((378) & 63))
                                    v15 = true
                                    _ = v19
                                }
                            case "hello":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            p1 = p0
                            _ = v15
                        }
                    }
                    p1 += -572
                    v20 := GenS1{F0: (0.1 < (float64((3.14159 + 0.0)) / (-2.25 + 1e10))), F1: 0, F2: 252}
                    _ = v20
                }
                p2 = p2
                if len(v7) > 0 {
                    v7[0] = fi9
                }
                if (-220 > 1) {
                    v21 := GenS3{F0: p2, F1: (^(p1))}
                    v22 := GenS1{F0: false, F1: 1, F2: 255}
                    _ = v21
                    _ = v22
                } else {
                    v23 := GenS3{F0: (p2 + p2), F1: (^(p1))}
                    _ = v23
                }
                var v24 int32 = 42
                _ = v24
            } else {
                var v25 float64 = (1e10 + (float64(0.1) - 0.1))
                _ = v25
            }
            if true {
                p0 = (-(p1))
                var v26 uint32 = 2147483647
                _ = v26
            }
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 uint, p1 bool) string {
    v27 := []int{-141, 4294967295, 127}
    var v28 int8 = 100
    v28 = v28
    if (65535 < 135) {
        var v29 uint = 255
        _ = len(v27)
        _ = v29
    } else {
        if (-9223372036854775808 == 42) {
            for fi30 := 0; fi30 < 15; fi30++ {
                v31 := GenS4{F0: p1, F1: 2}
                fmt.Println(v28)
                _ = fi30
                _ = v31
            }
            fmt.Println(p0, p0, p1)
        } else {
            var v32 int32 = -1
            _ = v32
        }
        var v33 string = "hello"
        _ = v33
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return ""
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 uint = 1
    v4 := []int{-289, 32768, 32768}
    v5 := make(map[int]uint32)
    v6 := []string{"", "foo bar", "a", "foo bar"}
    var v34 int16 = 1000
    fmt.Println(v34, v3)
    var v35 int64 = 1
    var v36 string = ("foo bar" + (("hello" + "hello") + ""))
    if (18446744073709551615 == 2) {
        for fi37 := 0; fi37 < 8; fi37++ {
            v3 = v3
            var v38 uint = (v3 >> (((^(v3))) & 63))
            v3 = (v38 & 826)
            _ = fi37
            _ = v38
        }
        var v39 int8 = 0
        v39 -= (-(v39))
        _ = v39
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 int32 = -2147483648
            for fi42 := 0; fi42 < 4; fi42++ {
                fmt.Println(fi42, v41, v41)
                v36 = ("hello" + (v36 + v36))
                _ = fi42
            }
            _ = fi40
            _ = v41
        }
        v43 := GenS3{F0: v36, F1: (v34 - v34)}
        _ = v43
    }
    v44 := v5[4294967295]
    for fi45 := 0; fi45 < 1; fi45++ {
        var v46 string = "\n\t"
        v34 = v34
        _ = fi45
        _ = v46
    }
    _ = len(v6)
    v47 := GenS5{F0: 32767}
    switch "" {
    case "hello":
        for fi48 := 0; fi48 < 17; fi48++ {
            switch fi48 {
            case 32768:
                var v49 int32 = 116
                _ = v49
            case 2147483647:
                v50 := GenS5{F0: 1}
                _ = v50
            case -2147483648:
                for fi51 := 0; fi51 < 20; fi51++ {
                    var v52 int64 = (v35 % ((v35) | 1))
                    _ = fi51
                    _ = v52
                }
            default:
                _ = len(v2)
            }
            _ = fi48
        }
    default:
        switch 0 {
        case 604:
            switch "a" {
            case "\n\t":
                var v53 uint32 = 2147483647
                _ = v53
            case "":
                _ = len(v4)
            default:
                switch 32768 {
                case -381:
                    for fi54 := 0; fi54 < 8; fi54++ {
                        var v55 int64 = (v35 * (v35 / ((v35) | 1)))
                        if ((fi54 + fi54) <= -633) {
                            switch ("\n\t" + (v36 + ("foo bar" + ("foo bar" + (v36 + v36))))) {
                            case "foo bar":
                                v56 := GenS6{F0: 2147483647, F1: 201, F2: (!(v35 >= v55))}
                                _ = v56
                            case "":
                                switch v36 {
                                case "hello":
                                    v57 := GenS7{F0: v34}
                                    _ = v57
                                case "\n\t":
                                    delete(v2, (fi54 / (((^(fi54))) | 1)))
                                case "":
                                    fmt.Println(v55, v44)
                                }
                            }
                            var v58 uint8 = 2
                            _ = v58
                        }
                        var v59 uint16 = 926
                        _ = fi54
                        _ = v55
                        _ = v59
                    }
                case 127:
                    if len(v4) > 0 {
                        v4[0] = 32768
                    }
                case 481:
                    if ((1e10 != (0.0 + -2.25)) || ((0 > 373) && false)) {
                        for fi60 := 0; fi60 < 12; fi60++ {
                            v36 = "hello"
                            for fi61 := 0; fi61 < 9; fi61++ {
                                v34 += int16(v3)
                                _ = len(v5)
                                _ = fi61
                            }
                            var v62 uint32 = v44
                            _ = fi60
                            _ = v62
                        }
                        var v63 int16 = v34
                        _ = v63
                    } else {
                        if false {
                            fmt.Println(v34, v36)
                            for fi64 := 0; fi64 < 6; fi64++ {
                                for fi65 := 0; fi65 < 15; fi65++ {
                                    fmt.Println(v3, v3, fi65)
                                    _ = fi65
                                }
                                _ = fi64
                            }
                        }
                        var v66 int8 = 0
                        _ = v66
                    }
                default:
                    if true {
                        var v67 bool = false
                        v1["hello"] = (!v67)
                        v68 := GenS8{F0: false, F1: uint16(v44), F2: ""}
                        _ = v67
                        _ = v68
                    } else {
                        var v69 int16 = v34
                        v44 += v44
                        if true {
                            v35 = (v35 | genFn0(v69, (v34 % ((v69) | 1)), v36))
                            var v70 float64 = 1e10
                            _ = v70
                        } else {
                            v44 = (^(v44))
                            if ((v3 < 18446744073709551615) && true) {
                                v71 := GenS2{F0: 128, F1: -2147483648}
                                v35 = (v35 * 173)
                                _ = v71
                            }
                        }
                        _ = v69
                    }
                }
            }
        default:
            v34 -= 1000
        }
    }
    var v72 float32 = (float32(float64(-2.25)) * ((((0.0 * (0.1 * 3.14159)) / 1.0) + float32((-2.25 - 0.0))) + float32((-2.25 + 0.1))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v47
    _ = v5
    _ = v6
    _ = v72
}
