package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 int8; F1 int8 }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]int)
    var v3 int16 = 32767
    v4 := make(map[int]int)
    var v5 bool = ((42 == -2147483648) || true)
    var v6 int16 = v3
    for fi7 := 0; fi7 < 5; fi7++ {
        switch fi7 {
        case 127:
            for fi8 := 0; fi8 < 9; fi8++ {
                if (0.0 != ((-2.25 + float32(((0.1 * (-2.25 * 0.1)) * -2.25))) - 0.1)) {
                    v3 = v3
                }
                v5 = (42 == 255)
                v9 := GenS1{F0: (1 > 0), F1: ((45 <= 441) && (!(176 >= 65535))), F2: 292}
                _ = fi8
                _ = v9
            }
        default:
            var v10 int32 = 627
            _ = v10
        }
        fmt.Println(v1, fi7, v6)
        v6 -= -325
        _ = fi7
    }
    fmt.Println(v3, v3)
    var v11 int32 = 2147483647
    var v12 uint8 = 1
    var v13 uint = 0
    v11 = (-(v11))
    v3 = -575
    v14 := GenS1{F0: false, F1: v5, F2: (v13 >> ((v13) & 63))}
    if v5 {
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(v5, v11, v11)
            _ = fi15
        }
        switch v1 {
        case -550:
            if v5 {
                v16 := []int{(v1 << (((v13 | 2)) & 63)), 387, (^(v1))}
                switch v1 {
                case -773:
                    fmt.Println(v3)
                case 32768:
                    v17 := GenS1{F0: true, F1: true, F2: (^(v13))}
                    _ = v17
                case 1:
                    v12 = (v12 & 135)
                }
                _ = v16
            } else {
                _ = len(v4)
                var v18 bool = (2 != 203)
                _ = v18
            }
        default:
            v13 -= uint(v12)
        }
        v3 -= int16(v13)
    }
    var v19 uint8 = 255
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 uint16 = 1
        if (v21 == v21) {
            for fi22 := 0; fi22 < 9; fi22++ {
                switch fi20 {
                case -2147483648:
                    var v23 uint8 = 187
                    _ = v23
                case 32767:
                    v24 := []int{v1, v1}
                    _ = v24
                case 4294967295:
                    v25 := GenS2{F0: "\n\t", F1: 4294967295}
                    _ = v25
                default:
                    if (("" + (("hello" + "a") + "")) != "foo bar") {
                        v21 += (v21 ^ (v21 / ((212) | 1)))
                        if v5 {
                            v6 = (^(v6))
                            v26 := GenS3{F0: v11, F1: v5}
                            fmt.Println(v12, v11, v21)
                            v1 = fi22
                            if (1 != (v13 << (((v13 * v13)) & 63))) {
                                if (v5 || ((1.5 < 3.14159) && v5)) {
                                    fmt.Println(v12, v6)
                                } else {
                                    v13 += (v13 + (v13 - 4294967295))
                                    var v27 float32 = -2.25
                                    fmt.Println(v3)
                                    v13 = v13
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch ((("a" + (("hello" + "hello") + ("foo bar" + ""))) + "\n\t") + "\n\t") {
                                        case "a":
                                            v28 := GenS3{F0: (v11 % ((42) | 1)), F1: v5}
                                            _ = v28
                                        case "":
                                            v29 := GenS3{F0: v11, F1: v5}
                                            _ = v29
                                        case "foo bar":
                                            var v30 uint = (v13 & 622)
                                            _ = v30
                                        }
                                    case "foo bar":
                                        var v31 bool = true
                                        _ = v31
                                    case "hello":
                                        v32 := GenS4{F0: 4294967295, F1: -1}
                                        _ = v32
                                    }
                                    v33 := v4[32768]
                                    if (v5 || (!(false || true))) {
                                        fmt.Println(v6)
                                        v34 := []int{v1, (v1 / ((-9223372036854775808) | 1)), fi20}
                                        _ = v34
                                    } else {
                                        fmt.Println(v1)
                                        fmt.Println(fi20)
                                    }
                                    _ = v27
                                    _ = v33
                                }
                            } else {
                                var v35 uint64 = 1
                                v2[0] = (v1 / ((fi22) | 1))
                                if (!v5) {
                                    v21 = (v21 * v21)
                                    v36 := GenS1{F0: false, F1: (v5 && false), F2: v13}
                                    v12 = (v19 * 128)
                                    _ = v36
                                } else {
                                    v21 = (v21 & 324)
                                    var v37 int16 = (v6 & v6)
                                    var v38 int64 = -627
                                    _ = v37
                                    _ = v38
                                }
                                _ = v35
                            }
                            _ = v26
                        }
                        var v39 int16 = (^(v3))
                        _ = v39
                    } else {
                        v40 := GenS5{F0: 4294967295}
                        v41 := v2[0]
                        v42 := v2[880]
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
                var v43 int32 = 2147483647
                _ = fi22
                _ = v43
            }
        } else {
            switch ("" + ((((("" + ((("hello" + ("hello" + "a")) + ((("\n\t" + "\n\t") + "foo bar") + ("foo bar" + "\n\t"))) + "\n\t")) + "hello") + (((("hello" + "a") + ("" + "hello")) + "hello") + "hello")) + "") + ("foo bar" + "foo bar"))) {
            case "":
                v2[2147483647] = (fi20 / ((v1) | 1))
            case "a":
                for fi44 := 0; fi44 < 6; fi44++ {
                    var v45 bool = false
                    v4[383] = fi20
                    _ = fi44
                    _ = v45
                }
            case "foo bar":
                v46 := []int{int(v19), (fi20 ^ fi20)}
                _ = v46
            default:
                if v5 {
                    v3 -= (v3 % ((32767) | 1))
                    for fi47 := 0; fi47 < 8; fi47++ {
                        v48 := []int{(fi20 & fi47), (fi20 >> ((v13) & 63))}
                        v12 -= v19
                        switch "hello" {
                        case "a":
                            if (v5 && ((("hello" + ("" + "")) + "") <= ("a" + "a"))) {
                                switch "" {
                                case "":
                                    v1 += (v1 << (((v13 % (((^(v13))) | 1))) & 63))
                                case "hello":
                                    fmt.Println(fi20, v13, v13)
                                case "foo bar":
                                    v5 = v5
                                }
                                var v49 uint64 = 157
                                var v50 bool = false
                                v13 -= v13
                                _ = v49
                                _ = v50
                            }
                        }
                        if false {
                            var v51 int32 = v11
                            v21 -= v21
                            if ((v5 && (9223372036854775807 <= 643)) && v5) {
                                v11 = v11
                                if ((v21 << ((342) & 63)) > 65535) {
                                    var v52 bool = (0 > v19)
                                    var v53 uint32 = 2147483648
                                    _ = len(v4)
                                    var v54 string = ""
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                } else {
                                    fmt.Println(v5)
                                    var v55 int8 = -52
                                    _ = v55
                                }
                            }
                            _ = v51
                        }
                        _ = fi47
                        _ = v48
                    }
                    v13 = 2
                }
            }
            v4[fi20] = v1
            v56 := GenS6{F0: (v1 | (v1 % ((v1) | 1))), F1: (v21 << (((v13 / ((v13) | 1))) & 63))}
            _ = v56
        }
        _ = fi20
        _ = v21
    }
    v57 := GenS7{F0: 127, F1: 1}
    var v58 string = ("\n\t" + "")
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
