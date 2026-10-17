package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int16; F2 string }
type GenS9 struct { F0 uint8; F1 uint64; F2 uint }
type GenS10 struct { F0 int; F1 int8; F2 bool }
func genFn0() bool {
    switch 0 {
    case 9223372036854775807:
        v7 := []int{358, 127, 42}
        _ = v7
    case -9223372036854775808:
        for fi8 := 0; fi8 < 15; fi8++ {
            switch fi8 {
            case 32767:
                v9 := GenS1{F0: -1, F1: 2147483647, F2: ("foo bar" + "foo bar")}
                _ = v9
            case -2147483648:
                if (true || (-791 <= 2147483647)) {
                    if (!false) {
                        for fi10 := 0; fi10 < 14; fi10++ {
                            var v11 uint = 255
                            _ = fi10
                            _ = v11
                        }
                        if false {
                            v12 := []int{}
                            if true {
                                for fi13 := 0; fi13 < 16; fi13++ {
                                    fmt.Println(fi13)
                                    _ = fi13
                                }
                            } else {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                if true {
                                    switch ("a" + (("" + "hello") + "hello")) {
                                    case "":
                                        for fi14 := 0; fi14 < 6; fi14++ {
                                            var v15 uint32 = 4294967295
                                            _ = fi14
                                            _ = v15
                                        }
                                    case "hello":
                                        v16 := GenS1{F0: 42, F1: 922, F2: "hello"}
                                        _ = v16
                                    case "foo bar":
                                        var v17 string = (("hello" + (("hello" + "") + "foo bar")) + "foo bar")
                                        _ = v17
                                    }
                                    var v18 int = 0
                                    _ = v18
                                }
                                switch (("hello" + "foo bar") + ("foo bar" + "")) {
                                case "foo bar":
                                    v19 := GenS2{F0: 4294967295, F1: 375, F2: 1}
                                    _ = v19
                                case "hello":
                                    fmt.Println(fi8)
                                case "a":
                                    var v20 int = fi8
                                    _ = v20
                                default:
                                    var v21 float64 = 0.1
                                    _ = v21
                                }
                            }
                            for fi22 := 0; fi22 < 4; fi22++ {
                                var v23 uint32 = 42
                                v24 := GenS2{F0: (v23 - 774), F1: (^(v23)), F2: 1}
                                fmt.Println(fi8)
                                if true {
                                    var v25 string = ("a" + ("\n\t" + "hello"))
                                    if true {
                                        v25 = ((("" + "a") + v25) + ("\n\t" + v25))
                                    } else {
                                        var v26 float32 = 3.14159
                                        var v27 float32 = v26
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v25
                                } else {
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        fmt.Println(v23)
                                    case "":
                                        v28 := GenS1{F0: 42, F1: 330, F2: "hello"}
                                        _ = v28
                                    case "foo bar":
                                        v29 := GenS3{F0: 721}
                                        _ = v29
                                    default:
                                        v23 -= v23
                                    }
                                }
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            _ = v12
                        } else {
                            if false {
                                var v30 int32 = 780
                                var v31 uint32 = 2147483648
                                if true {
                                    if (false || true) {
                                        fmt.Println(v30)
                                        if true {
                                            switch fi8 {
                                            case -1:
                                                for fi32 := 0; fi32 < 5; fi32++ {
                                                    switch ("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + "")) {
                                                    case "a":
                                                        v30 += v30
                                                    case "foo bar":
                                                        v31 = v31
                                                    }
                                                    fmt.Println(fi8)
                                                    if (("hello" + "foo bar") >= "hello") {
                                                        var v33 int16 = -563
                                                        var v34 float64 = ((0.0 + 3.14159) + 0.1)
                                                        v34 -= (3.14159 * float64((0.0 + float32(-2.25))))
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    _ = fi32
                                                }
                                            case 9223372036854775807:
                                                v35 := []int{fi8, 24, (fi8 / ((fi8) | 1)), fi8}
                                                _ = v35
                                            case 918:
                                                switch (fi8 / ((fi8) | 1)) {
                                                case -9223372036854775808:
                                                    switch fi8 {
                                                    case 128:
                                                        for fi36 := 0; fi36 < 4; fi36++ {
                                                            fmt.Println(fi36)
                                                            v37 := []int{0, (fi8 & (^(fi8))), fi36}
                                                            v37 = append(v37, (fi8 ^ fi8))
                                                            _ = fi36
                                                            _ = v37
                                                        }
                                                    case -2147483648:
                                                        v38 := GenS4{F0: 820, F1: 2}
                                                        _ = v38
                                                    }
                                                case -80:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v30 = (v30 | -2147483648)
                                                    }
                                                default:
                                                    v30 += (v30 ^ (^(v30)))
                                                }
                                            }
                                        } else {
                                            v39 := []int{4294967295, 32767, (fi8 * fi8), (fi8 ^ (fi8 & fi8))}
                                            for fi40 := 0; fi40 < 9; fi40++ {
                                                v41 := GenS5{F0: (v30 + v30), F1: "a"}
                                                _ = fi40
                                                _ = v41
                                            }
                                            var v42 int8 = 42
                                            _ = v39
                                            _ = v42
                                        }
                                        v43 := GenS6{F0: 1}
                                        _ = v43
                                    } else {
                                        var v44 int16 = int16(v30)
                                        var v45 uint64 = 391
                                        for fi46 := 0; fi46 < 15; fi46++ {
                                            fmt.Println(fi46)
                                            _ = fi46
                                        }
                                        _ = v44
                                        _ = v45
                                    }
                                    var v47 int32 = (^(v30))
                                    v31 -= (v31 - 1)
                                    _ = v47
                                }
                                _ = v30
                                _ = v31
                            } else {
                                var v48 int = fi8
                                _ = v48
                            }
                            switch fi8 {
                            case 128:
                                for fi49 := 0; fi49 < 14; fi49++ {
                                    v50 := GenS7{F0: -1}
                                    fmt.Println(fi8, fi8)
                                    if false {
                                        var v51 float64 = 1e10
                                        _ = v51
                                    }
                                    _ = fi49
                                    _ = v50
                                }
                            case 131:
                                fmt.Println(fi8)
                            }
                            var v52 float64 = 1e10
                            _ = v52
                        }
                        var v53 int32 = 2147483647
                        _ = v53
                    } else {
                        v54 := []int{fi8, (-(fi8))}
                        _ = v54
                    }
                    var v55 uint8 = 1
                    _ = v55
                } else {
                    var v56 uint8 = 1
                    fmt.Println(fi8)
                    _ = v56
                }
            case 405:
                fmt.Println(fi8)
            default:
                for fi57 := 0; fi57 < 13; fi57++ {
                    var v58 uint32 = 396
                    _ = fi57
                    _ = v58
                }
            }
            v59 := GenS2{F0: 2, F1: 210, F2: -549}
            fmt.Println(fi8)
            _ = fi8
            _ = v59
        }
    case 32767:
        for fi60 := 0; fi60 < 16; fi60++ {
            var v61 int = fi60
            for fi62 := 0; fi62 < 16; fi62++ {
                if ((1000 == -32768) && true) {
                    switch ("a" + ("a" + "foo bar")) {
                    case "foo bar":
                        var v63 float32 = 0.0
                        _ = v63
                    default:
                        v64 := []int{32768, 32767, (fi62 >> ((0) & 63))}
                        _ = v64
                    }
                    v65 := GenS8{F0: 1, F1: -32768, F2: "foo bar"}
                    var v66 uint = 0
                    _ = v65
                    _ = v66
                }
                switch "hello" {
                case "foo bar":
                    if true {
                        v61 -= fi60
                        v61 -= v61
                    }
                case "":
                    switch -9223372036854775808 {
                    case 0:
                        var v67 bool = true
                        _ = v67
                    case 32768:
                        if (!(false || false)) {
                            switch "" {
                            case "a":
                                v68 := GenS4{F0: 1, F1: 246}
                                _ = v68
                            case "foo bar":
                                var v69 uint32 = 2147483648
                                _ = v69
                            }
                            v70 := []int{(fi62 * -1), (v61 + v61), 901, v61}
                            if len(v70) >= 2 {
                                v70 = v70[1:len(v70)]
                            }
                            _ = v70
                        } else {
                            fmt.Println(fi60, v61, v61)
                            v71 := []int{(fi60 / ((fi62) | 1)), (v61 & (v61 >> ((506) & 63))), -185, fi62}
                            _ = v71
                        }
                    default:
                        if ((!false) || (184 <= 4294967295)) {
                            v72 := []int{1, (fi60 * v61), (fi60 % ((fi62) | 1))}
                            v61 = (v61 * 127)
                            if false {
                                var v73 int = 4294967295
                                _ = v73
                            } else {
                                var v74 uint16 = 255
                                v75 := GenS2{F0: 523, F1: 708, F2: 42}
                                var v76 int64 = 839
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                            _ = v72
                        }
                    }
                case "a":
                    v61 -= fi62
                default:
                    v61 = v61
                }
                _ = fi62
            }
            _ = fi60
            _ = v61
        }
    }
    switch -1 {
    case -910:
        var v77 uint64 = 18446744073709551615
        _ = v77
    default:
        v78 := GenS5{F0: -2147483648, F1: "hello"}
        _ = v78
    }
    var v79 int16 = 1000
    v80 := GenS9{F0: 1, F1: 42, F2: 18446744073709551615}
    _ = v79
    _ = v80
    return (-1 != 1)
}

func main() {
    var v1 int16 = 365
    v2 := make(map[int]bool)
    var v3 int32 = 1
    v4 := make(map[uint32]int)
    var v5 uint32 = 651
    v6 := make(map[int]uint32)
    fmt.Println(v1, v5)
    for fi81 := 0; fi81 < 15; fi81++ {
        var v82 int8 = -1
        v6[127] = v5
        fmt.Println(fi81)
        _ = fi81
        _ = v82
    }
    v3 = (^(v3))
    v1 -= (v1 * (v1 >> ((65535) & 63)))
    for fi83 := 0; fi83 < 19; fi83++ {
        fmt.Println(fi83, v1)
        _ = fi83
    }
    if false {
        v1 = v1
        switch 765 {
        case 128:
            v3 = v3
        case -2147483648:
            var v84 string = "hello"
            _ = v84
        case 32768:
            var v85 string = "foo bar"
            _ = v85
        }
    }
    for fi86 := 0; fi86 < 16; fi86++ {
        var v87 int = (fi86 % ((fi86) | 1))
        _ = fi86
        _ = v87
    }
    v5 = (v5 & (^(v5)))
    var v88 int = -209
    v88 -= (v88 * (v88 * 539))
    switch (v88 * v88) {
    case -916:
        if false {
            var v89 float64 = ((1.5 - 1e10) - 0.0)
            fmt.Println(v89)
            _ = v89
        } else {
            v90 := v2[v88]
            _ = v90
        }
    default:
        v91 := GenS10{F0: (v88 | -1), F1: -128, F2: (32767 != 42)}
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
}
