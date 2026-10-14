package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS4 struct { F0 bool; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS9 struct { F0 bool; F1 bool; F2 string }
func genFn0(p0 int8, p1 uint) int {
    v5 := []int{}
    if ((("a" + ("foo bar" + ((((("hello" + ("hello" + "a")) + "a") + "hello") + (("foo bar" + ("foo bar" + "\n\t")) + ((("\n\t" + (("" + ("" + "foo bar")) + ("\n\t" + ""))) + "foo bar") + "a"))) + "\n\t"))) > "a") || (false && true)) {
        v6 := GenS1{F0: (p1 * p1), F1: 612, F2: "hello"}
        p1 += p1
        _ = v6
    }
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            if len(v5) > 0 {
                v5[0] = 0
            }
            switch ("foo bar" + "a") {
            case "":
                if false {
                    v5 = append(v5, fi7)
                    p1 = p1
                } else {
                    fmt.Println(p0, fi7, p1)
                    p0 = (p0 / ((p0) | 1))
                    for fi8 := 0; fi8 < 3; fi8++ {
                        p1 = (p1 & 762)
                        var v9 int16 = 1
                        v10 := GenS2{F0: v9}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
            }
            var v11 int = 2147483647
            _ = fi7
            _ = v11
        }
        p1 += p1
    } else {
        for fi12 := 0; fi12 < 14; fi12++ {
            if false {
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 int = 662
                    var v15 int64 = -9223372036854775808
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 uint64 = 18446744073709551615
                switch "foo bar" {
                case "hello":
                    v5 = append(v5, 32768)
                case "":
                    fmt.Println(p0)
                }
                _ = v16
            } else {
                p1 -= (p1 ^ p1)
                var v17 uint8 = 255
                v17 = uint8(p1)
                _ = v17
            }
            _ = fi12
        }
        switch -16 {
        case 42:
            if (p1 > (p1 & 353)) {
                v18 := GenS1{F0: (p1 << (((p1 % ((p1) | 1))) & 63)), F1: 42, F2: ("hello" + "\n\t")}
                for fi19 := 0; fi19 < 16; fi19++ {
                    var v20 uint = p1
                    switch "a" {
                    case "\n\t":
                        var v21 uint = 0
                        _ = v21
                    case "":
                        var v22 int8 = (p0 & (p0 | p0))
                        _ = v22
                    case "a":
                        if ((1000 < -696) || (0 >= 42)) {
                            p1 = (p1 & 2)
                        }
                    }
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v24 := GenS3{F0: 68, F1: (v20 >> ((p1) & 63)), F2: 18446744073709551615}
                        p0 += (p0 % ((p0) | 1))
                        switch "foo bar" {
                        case "hello":
                            var v25 uint32 = 602
                            _ = v25
                        case "a":
                            switch fi23 {
                            case 42:
                                switch "hello" {
                                case "a":
                                    if (2 <= 0) {
                                        if false {
                                            fmt.Println(p0, v20, p1)
                                        } else {
                                            switch fi19 {
                                            case 1:
                                                var v26 bool = true
                                                _ = v26
                                            }
                                            var v27 uint64 = uint64(v20)
                                            var v28 uint64 = v27
                                            _ = v27
                                            _ = v28
                                        }
                                        v29 := GenS4{F0: false, F1: 1000, F2: 301}
                                        v20 = 750
                                        _ = v29
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                default:
                                    fmt.Println(p1, fi19, v20)
                                }
                            case 9223372036854775807:
                                fmt.Println(fi19, fi19)
                            default:
                                fmt.Println(fi23, v20, p0)
                            }
                        case "\n\t":
                            v30 := GenS5{F0: 2147483647, F1: 2147483647, F2: 0}
                            v31 := GenS6{F0: 65535, F1: -1}
                            var v32 uint = 1
                            _ = v30
                            _ = v31
                            _ = v32
                        default:
                            switch "hello" {
                            case "a":
                                _ = len(v5)
                            }
                        }
                        var v33 int16 = -1
                        v20 += (v20 * (p1 % (((p1 | 725)) | 1)))
                        _ = fi23
                        _ = v24
                        _ = v33
                    }
                    _ = fi19
                    _ = v20
                }
                p1 = (p1 ^ (p1 % (((p1 ^ (^(p1)))) | 1)))
                _ = v18
            } else {
                v34 := GenS7{F0: ((-1 > 42) && true), F1: (p0 / (((p0 / ((91) | 1))) | 1))}
                _ = v34
            }
        default:
            var v35 uint32 = 955
            _ = v35
        }
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1() int32 {
    if true {
        v36 := GenS7{F0: true, F1: 1}
        if true {
            var v37 int = 123
            var v38 uint32 = 2147483647
            for fi39 := 0; fi39 < 2; fi39++ {
                v40 := []int{fi39}
                _ = fi39
                _ = v40
            }
            _ = v37
            _ = v38
        }
        for fi41 := 0; fi41 < 4; fi41++ {
            fmt.Println(fi41)
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 float32 = 0.0
                fmt.Println(v43, fi41)
                v43 -= ((float32(0.0) + 0.0) / ((v43 - float32(3.14159)) / (float32(float64(float32((1e10 / ((3.14159 + 3.14159) / 1.5))))) / v43)))
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
        _ = v36
    } else {
        var v44 int16 = -32768
        switch "hello" {
        case "":
            fmt.Println(v44)
        case "\n\t":
            v45 := []int{4294967295, 32767, genFn0(42, 553)}
            _ = v45
        case "a":
            v46 := GenS8{F0: 127, F1: (v44 + (v44 << ((0) & 63))), F2: 65535}
            _ = v46
        }
        _ = v44
    }
    if (((0.0 / float64((-2.25 - float32(float64(1.5))))) / (1e10 / 1e10)) > float64(3.14159)) {
        if (2 != 991) {
            v47 := []int{2147483647, genFn0(117, 1), 303, 32767}
            if len(v47) > 0 {
                v47[0] = -9223372036854775808
            }
            for fi48 := 0; fi48 < 2; fi48++ {
                var v49 bool = false
                for fi50 := 0; fi50 < 9; fi50++ {
                    v49 = (1 == 39)
                    var v51 int = genFn0(1, 42)
                    _ = fi50
                    _ = v51
                }
                fmt.Println(fi48, v49)
                _ = fi48
                _ = v49
            }
            _ = v47
        } else {
            if true {
                if false {
                    v52 := GenS4{F0: (-1 <= 1), F1: 1000, F2: 65535}
                    if true {
                        v53 := GenS9{F0: false, F1: (true && true), F2: "foo bar"}
                        _ = v53
                    } else {
                        var v54 uint = 4294967295
                        fmt.Println(v54)
                        _ = v54
                    }
                    _ = v52
                } else {
                    var v55 int8 = 100
                    if ((!(134 <= 128)) || true) {
                        var v56 string = "a"
                        var v57 int = 0
                        v57 = 9223372036854775807
                        _ = v56
                        _ = v57
                    }
                    _ = v55
                }
                v58 := GenS1{F0: 0, F1: 18446744073709551615, F2: ""}
                _ = v58
            }
            var v59 bool = false
            _ = v59
        }
    }
    v60 := []int{9223372036854775807, -9223372036854775808, 128}
    _ = v60
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{0, 1}
    v3 := []int{32767}
    var v4 float64 = float64(0.1)
    for fi61 := 0; fi61 < 9; fi61++ {
        v62 := GenS2{F0: 32767}
        _ = fi61
        _ = v62
    }
    v4 = (v4 * (((-2.25 + 1.5) + v4) * (v4 / (v4 + v4))))
    v4 += v4
    fmt.Println(v4)
    for fi63 := 0; fi63 < 9; fi63++ {
        for fi64 := 0; fi64 < 9; fi64++ {
            v1[0] = (^(fi64))
            fmt.Println(fi64)
            _ = fi64
        }
        v2 = append(v2, 655)
        var v65 string = ("" + "a")
        _ = fi63
        _ = v65
    }
    var v66 int16 = 1000
    var v67 uint = 494
    v4 += ((0.0 * (float64((-2.25 * (1.5 / 0.1))) - v4)) * (v4 * 1.5))
    fmt.Println(v66)
    for fi68 := 0; fi68 < 8; fi68++ {
        fmt.Println(v66, fi68)
        v67 += 890
        _ = fi68
    }
    v4 += v4
    var v69 int32 = -437
    if (0.0 >= 3.14159) {
        fmt.Println(v66, v66, v69)
        fmt.Println(v69, v69, v66)
    } else {
        if false {
            v66 += v66
            var v70 int = 2147483647
            _ = v70
        } else {
            var v71 int = -9223372036854775808
            _ = v71
        }
        var v72 float32 = (((1e10 * 1.5) - -2.25) * (1.5 - (float32(float64(1.5)) + 1.5)))
        v67 += v67
        _ = v72
    }
    if (false || (!false)) {
        if len(v3) > 0 {
            v3[0] = 128
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v69
}
