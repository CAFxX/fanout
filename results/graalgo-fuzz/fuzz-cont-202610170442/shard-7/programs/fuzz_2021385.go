package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 string; F2 string }
func genFn0() int64 {
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint = 4294967295
        v10 := []int{(-(fi8)), fi8, (fi8 + (-(fi8))), fi8}
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 uint32 = (v7 ^ 42)
            v13 := GenS1{F0: fi11}
            if len(v10) > 0 {
                v10[0] = (fi8 | fi11)
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v14 := GenS1{F0: -1}
    switch "hello" {
    case "":
        if (!true) {
            v15 := []int{9223372036854775807, 4294967295}
            if (!false) {
                var v16 int = 32767
                for fi17 := 0; fi17 < 16; fi17++ {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    _ = fi17
                }
                for fi18 := 0; fi18 < 6; fi18++ {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if ((1000 > int16(v16)) && true) {
                            var v20 uint64 = 4294967295
                            switch "hello" {
                            case "":
                                if false {
                                    var v21 int = fi19
                                    if false {
                                        v22 := GenS1{F0: v16}
                                        v20 -= (v20 - 42)
                                        if len(v15) > 0 {
                                            v15[0] = v21
                                        }
                                        _ = v22
                                    } else {
                                        fmt.Println(v7, v20, v16)
                                    }
                                    var v23 uint8 = 255
                                    fmt.Println(v20, v16)
                                    _ = v21
                                    _ = v23
                                } else {
                                    if (false && ((^(v20)) < (v20 >> ((2) & 63)))) {
                                        var v24 uint = 4294967295
                                        _ = v24
                                    } else {
                                        switch ("foo bar" + "foo bar") {
                                        case "foo bar":
                                            if true {
                                                _ = len(v15)
                                                if (false && true) {
                                                    if len(v15) > 0 {
                                                        v15[0] = fi19
                                                    }
                                                    fmt.Println(fi19, fi18, fi19)
                                                    var v25 uint8 = 113
                                                    _ = v25
                                                }
                                                if true {
                                                    v26 := GenS1{F0: fi19}
                                                    v27 := GenS1{F0: (fi18 + 32767)}
                                                    _ = v26
                                                    _ = v27
                                                }
                                                var v28 uint64 = (v20 * (v20 & v20))
                                                _ = v28
                                            } else {
                                                var v29 uint8 = 2
                                                _ = v29
                                            }
                                        case "\n\t":
                                            v15 = append(v15, fi18)
                                        case "a":
                                            if true {
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(v16)
                                                }
                                            } else {
                                                v20 = (v20 % ((v20) | 1))
                                            }
                                            v20 = (v20 + 18446744073709551615)
                                            v30 := GenS2{F0: 263, F1: 1, F2: v20}
                                            _ = v30
                                        default:
                                            v20 -= v20
                                            if true {
                                                v31 := GenS2{F0: v20, F1: -642, F2: v20}
                                                _ = v31
                                            }
                                        }
                                        v16 -= fi19
                                        switch (^(v16)) {
                                        case 32768:
                                            var v32 uint32 = (v7 << ((65535) & 63))
                                            _ = v32
                                        case 497:
                                            v33 := GenS3{F0: 0, F1: v20, F2: 543}
                                            _ = v33
                                        case 4294967295:
                                            switch "a" {
                                            case "":
                                                if (true || true) {
                                                    var v34 int32 = 0
                                                    var v35 int32 = 1
                                                    if false {
                                                        v16 -= fi18
                                                        var v36 int64 = -366
                                                        _ = v36
                                                    }
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    switch ("\n\t" + "a") {
                                                    case "hello":
                                                        if false {
                                                            if false {
                                                                v37 := GenS1{F0: fi18}
                                                                v38 := GenS2{F0: (v20 << ((42) & 63)), F1: 826, F2: v20}
                                                                _ = v37
                                                                _ = v38
                                                            } else {
                                                                _ = len(v15)
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        fmt.Println(v16)
                                                    default:
                                                        v39 := GenS3{F0: 127, F1: 0, F2: -2147483648}
                                                        _ = v39
                                                    }
                                                }
                                            default:
                                                _ = len(v15)
                                            }
                                        default:
                                            var v40 uint16 = 2
                                            _ = v40
                                        }
                                    }
                                }
                            case "hello":
                                var v41 float32 = (1e10 - (((1.5 - ((0.0 - -2.25) / 3.14159)) - ((1e10 - 0.1) / 3.14159)) * 3.14159))
                                _ = v41
                            }
                            _ = v20
                        }
                        _ = fi19
                    }
                    fmt.Println(fi18)
                    v42 := GenS4{F0: -219}
                    _ = fi18
                    _ = v42
                }
                _ = v16
            } else {
                var v43 uint32 = 1
                var v44 int64 = 42
                switch "a" {
                case "\n\t":
                    var v45 float32 = (1e10 + (0.1 - 0.0))
                    _ = v45
                case "hello":
                    fmt.Println(v44)
                default:
                    v7 = v7
                }
                _ = v43
                _ = v44
            }
            _ = v15
        } else {
            v7 += v7
        }
    default:
        for fi46 := 0; fi46 < 7; fi46++ {
            v47 := GenS5{F0: 824, F1: true}
            if true {
                v48 := []int{(^(fi46)), (fi46 - (fi46 * fi46)), -44}
                for fi49 := 0; fi49 < 13; fi49++ {
                    v7 = (v7 / ((v7) | 1))
                    if false {
                        v50 := GenS4{F0: 70}
                        var v51 int16 = 0
                        var v52 uint = 18446744073709551615
                        v53 := GenS6{F0: "", F1: 0}
                        var v54 uint64 = 4294967295
                        _ = v50
                        _ = v51
                        _ = v52
                        _ = v53
                        _ = v54
                    } else {
                        if (107 != 126) {
                            fmt.Println(v7)
                            v7 = 42
                        }
                    }
                    switch 127 {
                    case 0:
                        switch "hello" {
                        case "":
                            v55 := GenS7{F0: 0}
                            _ = v55
                        case "hello":
                            if len(v48) >= 2 {
                                v48 = v48[1:len(v48)]
                            }
                            var v56 uint64 = 0
                            _ = v56
                        case "a":
                            v57 := GenS5{F0: 65535, F1: (2 < 128)}
                            _ = v57
                        }
                    case 32768:
                        if (42 != 42) {
                            var v58 float64 = float64(-2.25)
                            _ = v58
                        }
                    case -9223372036854775808:
                        var v59 uint64 = 976
                        _ = v59
                    }
                    _ = fi49
                }
                for fi60 := 0; fi60 < 9; fi60++ {
                    if len(v48) > 0 {
                        v48[0] = fi60
                    }
                    fmt.Println(v7)
                    fmt.Println(fi46, v7, fi46)
                    _ = fi60
                }
                _ = v48
            } else {
                if false {
                    var v61 uint8 = 9
                    var v62 uint = 392
                    _ = v61
                    _ = v62
                } else {
                    if false {
                        for fi63 := 0; fi63 < 1; fi63++ {
                            v64 := []int{9223372036854775807, fi63}
                            if len(v64) > 0 {
                                v64[0] = fi63
                            }
                            _ = fi63
                            _ = v64
                        }
                        v65 := GenS8{F0: (v7 ^ (v7 >> ((65535) & 63))), F1: "hello", F2: ("a" + ("foo bar" + ("\n\t" + (("a" + ("\n\t" + "")) + "a"))))}
                        _ = v65
                    } else {
                        v66 := []int{128, fi46, -1, (fi46 - (fi46 << ((18446744073709551615) & 63)))}
                        _ = v66
                    }
                    v7 = v7
                    v7 += (v7 % ((530) | 1))
                }
            }
            _ = fi46
            _ = v47
        }
    }
    if false {
        var v67 uint32 = v7
        v68 := GenS2{F0: 49, F1: 0, F2: 18446744073709551615}
        for fi69 := 0; fi69 < 3; fi69++ {
            v70 := GenS4{F0: 0}
            _ = fi69
            _ = v70
        }
        _ = v67
        _ = v68
    } else {
        switch -313 {
        case 4294967295:
            v71 := []int{-3, -889, 25}
            _ = v71
        case 326:
            v72 := []int{4294967295, 32768, -578, -493}
            _ = v72
        default:
            var v73 float32 = (float32((-2.25 / 1.0)) * (3.14159 / ((0.1 * 3.14159) / 1.5)))
            _ = v73
        }
    }
    _ = v14
    _ = v7
    return -1
}

func main() {
    var v1 float64 = (1e10 / 1.0)
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 uint8 = 127
    v5 := []string{((("hello" + "a") + "hello") + "foo bar"), "foo bar", "a", "a"}
    var v6 int32 = 0
    if false {
        var v74 string = "foo bar"
        var v75 uint8 = v4
        for fi76 := 0; fi76 < 3; fi76++ {
            v75 += 127
            _ = fi76
        }
        _ = v74
        _ = v75
    }
    for fi77 := 0; fi77 < 1; fi77++ {
        var v78 float64 = (-2.25 / v1)
        _ = fi77
        _ = v78
    }
    var v79 string = ""
    var v80 int8 = 1
    var v81 int32 = (v6 * 2147483647)
    var v82 int64 = 42
    fmt.Println(v79, v3, v80)
    if (!(((3.14159 + 1.5) < -2.25) && v3)) {
        for fi83 := 0; fi83 < 5; fi83++ {
            v84 := GenS8{F0: 42, F1: v79, F2: v79}
            _ = fi83
            _ = v84
        }
    }
    var v85 int16 = 32767
    v86 := GenS7{F0: 501}
    v87 := GenS8{F0: 2147483648, F1: v79, F2: "foo bar"}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v80
    _ = v81
    _ = v82
    _ = v85
    _ = v86
    _ = v87
}
