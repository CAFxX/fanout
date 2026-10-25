package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int }
type GenS2 struct { F0 int16; F1 uint; F2 bool }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 uint16; F2 int }
type GenS9 struct { F0 uint16; F1 int; F2 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint16; F1 int64 }
func genFn0(p0 int64) uint16 {
    if true {
        v7 := []int{42}
        var v8 uint8 = 42
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        if (!(!true)) {
            var v10 int64 = (p0 | (p0 / ((p0) | 1)))
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            p0 = p0
            v12 := []int{fi11, fi9, (-(fi9)), (fi9 | 42)}
            p0 = (p0 & (p0 % ((1) | 1)))
            var v13 uint8 = 0
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    _ = p0
    return 255
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]int)
    var v3 int8 = 127
    v4 := make(map[uint32]bool)
    var v5 uint8 = 158
    v6 := make(map[uint32]uint32)
    if ((1000 <= 0) && (true || false)) {
        v3 = (v3 / (((v3 + (-(v3)))) | 1))
        var v14 uint8 = v5
        _ = v14
    } else {
        v1 = -2.25
        if (v5 >= (v5 & v5)) {
            var v15 uint8 = (v5 % ((127) | 1))
            v5 += (v15 ^ (^(v5)))
            v16 := GenS1{F0: 42, F1: 32767, F2: 32767}
            _ = v15
            _ = v16
        } else {
            v3 += 0
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            if true {
                v19 := GenS1{F0: 4294967295, F1: 383, F2: 127}
                v1 = v1
                _ = v19
            } else {
                if true {
                    delete(v2, uint32(v5))
                } else {
                    delete(v6, 449)
                    switch fi18 {
                    case 9223372036854775807:
                        switch fi18 {
                        case 32767:
                            if false {
                                var v20 string = "hello"
                                v21 := GenS1{F0: 904, F1: -1, F2: -2147483648}
                                _ = v20
                                _ = v21
                            } else {
                                var v22 uint = 42
                                switch "\n\t" {
                                case "\n\t":
                                    if ((1000 >= 42) || false) {
                                        var v23 bool = false
                                        fmt.Println(v1)
                                        var v24 int64 = -1
                                        v5 = (v5 << (((v22 - 4294967295)) & 63))
                                        _ = v23
                                        _ = v24
                                    } else {
                                        _ = len(v6)
                                        var v25 bool = ((280 < 332) && false)
                                        fmt.Println(v25, v1)
                                        v26 := []int{}
                                        _ = v25
                                        _ = v26
                                    }
                                case "foo bar":
                                    v2[2147483648] = 127
                                default:
                                    v5 -= (^(v5))
                                }
                                _ = v22
                            }
                        case -1:
                            if true {
                                v27 := GenS1{F0: 42, F1: 42, F2: 730}
                                v28 := GenS1{F0: 0, F1: -793, F2: 4294967295}
                                _ = v27
                                _ = v28
                            }
                        case 9223372036854775807:
                            v29 := []int{fi17, (fi17 | fi17)}
                            _ = v29
                        }
                    default:
                        v3 -= (v3 / (((v3 << ((2) & 63))) | 1))
                    }
                    fmt.Println(fi17, fi18)
                }
                var v30 float64 = (((((1e10 - (float64(float32((float64((1e10 / float32(1.5))) + 1e10))) - 1.5)) / (3.14159 + float64((1e10 / (v1 - (v1 * 0.0)))))) - ((1e10 / ((float64((float32(-2.25) / v1)) - (0.1 * (float64((((v1 / v1) * -2.25) / (v1 / v1))) - (float64((((v1 + 0.0) / 3.14159) * v1)) / 1e10)))) - 0.0)) * ((0.0 * (float64(-2.25) - 0.0)) + (-2.25 * 1.5)))) * ((0.1 / 3.14159) + (3.14159 + (1e10 * 0.1)))) / 1e10)
                _ = v30
            }
            _ = len(v4)
            v31 := GenS2{F0: 1, F1: 255, F2: false}
            _ = fi18
            _ = v31
        }
        for fi32 := 0; fi32 < 8; fi32++ {
            v33 := GenS3{F0: -9223372036854775808, F1: ("a" + "")}
            var v34 string = (("\n\t" + (("hello" + "") + "a")) + (("a" + "") + "a"))
            fmt.Println(v1)
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi17
    }
    v1 -= v1
    var v35 uint32 = 65535
    switch ("" + (("hello" + ("\n\t" + "hello")) + "\n\t")) {
    case "":
        if true {
            v36 := []int{-1, 42, 42, 336}
            fmt.Println(v1, v3, v1)
            v5 -= 128
            _ = v36
        } else {
            fmt.Println(v3, v35)
            v37 := []int{-621, -470}
            _ = v37
        }
    case "hello":
        v38 := v2[2147483647]
        _ = v38
    default:
        switch 2147483647 {
        case 2147483647:
            v39 := GenS4{F0: 365}
            _ = v39
        }
    }
    var v40 bool = false
    if v40 {
        v41 := GenS5{F0: 1000}
        if v40 {
            var v42 int64 = 376
            if v40 {
                for fi43 := 0; fi43 < 3; fi43++ {
                    v35 += 65535
                    _ = fi43
                }
            } else {
                for fi44 := 0; fi44 < 9; fi44++ {
                    v4[v35] = v40
                    v1 -= v1
                    v45 := GenS6{F0: "a"}
                    _ = fi44
                    _ = v45
                }
                var v46 uint64 = 506
                _ = v46
            }
            var v47 int = -334
            _ = v42
            _ = v47
        } else {
            if v40 {
                var v48 float64 = (0.1 / (3.14159 / 3.14159))
                _ = v48
            } else {
                for fi49 := 0; fi49 < 13; fi49++ {
                    for fi50 := 0; fi50 < 4; fi50++ {
                        v35 += (v35 + v35)
                        switch fi49 {
                        case 127:
                            v51 := GenS7{F0: v35}
                            _ = v51
                        }
                        var v52 int16 = 1000
                        var v53 uint16 = 727
                        _ = fi50
                        _ = v52
                        _ = v53
                    }
                    _ = fi49
                }
            }
        }
        _ = v41
    }
    if (-3 < -1) {
        if false {
            switch 460 {
            case 2147483647:
                v54 := GenS7{F0: (v35 % ((v35) | 1))}
                _ = v54
            case 330:
                v55 := GenS8{F0: "hello", F1: 910, F2: 127}
                _ = v55
            }
            if (-1 > 215) {
                v1 += v1
                for fi56 := 0; fi56 < 9; fi56++ {
                    v5 += (^(v5))
                    fmt.Println(v5, v35)
                    _ = fi56
                }
            }
        }
        fmt.Println(v35, v35)
    }
    v1 = 0.0
    var v57 int64 = 42
    for fi58 := 0; fi58 < 14; fi58++ {
        v6[v35] = 608
        _ = fi58
    }
    v40 = (2147483647 <= 2147483647)
    for fi59 := 0; fi59 < 2; fi59++ {
        if (false || v40) {
            for fi60 := 0; fi60 < 15; fi60++ {
                switch (^(fi59)) {
                case 127:
                    if v40 {
                        if v40 {
                            var v61 uint32 = (v35 - 2147483648)
                            v5 += (^(v5))
                            v35 = v35
                            _ = v61
                        } else {
                            v62 := GenS7{F0: (v35 + v35)}
                            _ = v62
                        }
                        var v63 int = -2147483648
                        _ = v63
                    }
                case -2147483648:
                    v57 -= (^(v57))
                }
                v64 := GenS2{F0: int16(v35), F1: 0, F2: v40}
                fmt.Println(v1, v35)
                v65 := []int{9223372036854775807, (^(fi60)), fi60, fi60}
                v35 = v35
                _ = fi60
                _ = v64
                _ = v65
            }
            var v66 int32 = 2147483647
            switch 127 {
            case 32767:
                var v67 uint64 = 4294967295
                _ = v67
            case -1:
                var v68 bool = v40
                _ = v68
            case -942:
                v69 := GenS5{F0: -327}
                _ = v69
            }
            _ = v66
        } else {
            if ("" < "\n\t") {
                switch (fi59 & fi59) {
                case 760:
                    v70 := v4[v35]
                    _ = v70
                case 32767:
                    if (!(!true)) {
                        v71 := []int{fi59, fi59, (fi59 | 967)}
                        v72 := GenS6{F0: "a"}
                        _ = v71
                        _ = v72
                    } else {
                        var v73 int = 9223372036854775807
                        if (-2147483648 > 2147483647) {
                            v74 := GenS9{F0: 255, F1: int(v35), F2: fi59}
                            _ = v74
                        } else {
                            if (1.5 <= 0.0) {
                                for fi75 := 0; fi75 < 1; fi75++ {
                                    v40 = true
                                    v76 := []int{(v73 + fi75), 334, (fi59 << ((4294967295) & 63))}
                                    var v77 int64 = v57
                                    _ = fi75
                                    _ = v76
                                    _ = v77
                                }
                                for fi78 := 0; fi78 < 14; fi78++ {
                                    v79 := GenS1{F0: 0, F1: 0, F2: (v73 ^ int(v3))}
                                    v80 := GenS10{F0: 32767}
                                    _ = fi78
                                    _ = v79
                                    _ = v80
                                }
                            } else {
                                v81 := GenS3{F0: v57, F1: ""}
                                v82 := GenS2{F0: -360, F1: 255, F2: (int32(fi59) == 42)}
                                _ = v81
                                _ = v82
                            }
                            var v83 int = 42
                            _ = v83
                        }
                        if v40 {
                            for fi84 := 0; fi84 < 3; fi84++ {
                                switch ("foo bar" + ("foo bar" + "hello")) {
                                case "":
                                    fmt.Println(v40, v73, v35)
                                case "foo bar":
                                    switch (-(v73)) {
                                    case -9223372036854775808:
                                        v73 = fi84
                                    case 752:
                                        v85 := []int{}
                                        _ = v85
                                    }
                                case "hello":
                                    v86 := []int{}
                                    _ = v86
                                default:
                                    var v87 bool = (42 < (v35 ^ v35))
                                    _ = v87
                                }
                                switch ("a" + "\n\t") {
                                case "a":
                                    var v88 uint64 = 0
                                    _ = v88
                                case "\n\t":
                                    v57 = (v57 * v57)
                                case "":
                                    v89 := []int{v73, v73, fi84, v73}
                                    _ = v89
                                default:
                                    if (!(uint(v35) <= 42)) {
                                        fmt.Println(v73, v5, v35)
                                        var v90 uint8 = v5
                                        _ = v90
                                    }
                                }
                                var v91 uint64 = 901
                                _ = fi84
                                _ = v91
                            }
                        } else {
                            v92 := GenS11{F0: 459, F1: -9223372036854775808}
                            _ = v92
                        }
                        _ = v73
                    }
                case 32768:
                    v93 := GenS9{F0: 684, F1: 127, F2: fi59}
                    _ = v93
                }
                v4[0] = (!v40)
            } else {
                fmt.Println(v5, v35, v40)
            }
        }
        _ = fi59
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v57
    _ = v6
}
