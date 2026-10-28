package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int8; F1 int8; F2 uint }
type GenS7 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS8 struct { F0 bool; F1 uint64; F2 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint8 }
func genFn0(p0 uint) uint {
    p0 += 65535
    p0 = (p0 * 5)
    var v4 int8 = 80
    var v5 uint = 434
    _ = p0
    _ = v4
    _ = v5
    return (p0 | (^(p0)))
}
func genFn1(p0 uint8, p1 int32, p2 int8) bool {
    p0 -= p0
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint16 = 0
        _ = fi6
        _ = v7
    }
    v8 := []int{-459, -2147483648}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ((p0 >> ((932) & 63)) != (p0 + 160))
}

func main() {
    var v1 int8 = 91
    var v2 uint64 = 4294967295
    var v3 int32 = 0
    if (32767 < 32767) {
        v9 := GenS1{F0: 42, F1: v3}
        var v10 uint32 = 2
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: (int32(v2) < v3), F1: 42, F2: (4294967295 <= 42)}
    var v12 int16 = -1
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS3{F0: false, F1: 2147483648, F2: "a"}
        v15 := []int{-1, fi13, fi13, (fi13 << ((18446744073709551615) & 63))}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v2, v3, v2)
    v12 = int16(v3)
    for fi16 := 0; fi16 < 19; fi16++ {
        if (!false) {
            if ((v12 * -378) <= v12) {
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 int64 = 776
                    _ = fi17
                    _ = v18
                }
                for fi19 := 0; fi19 < 2; fi19++ {
                    v1 -= 1
                    v20 := GenS3{F0: false, F1: 48, F2: "\n\t"}
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 4; fi21++ {
                    if (177 != 420) {
                        v22 := []int{}
                        var v23 uint16 = 32767
                        if true {
                            if true {
                                v3 -= 2147483647
                                switch "a" {
                                case "\n\t":
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                case "hello":
                                    v24 := GenS1{F0: 0, F1: (v3 | 1)}
                                    _ = v24
                                default:
                                    v3 = v3
                                }
                                v25 := GenS1{F0: (-(v1)), F1: (v3 + v3)}
                                _ = v25
                            }
                            _ = len(v22)
                            v12 = -1
                        }
                        var v26 int = (fi16 ^ 42)
                        _ = v22
                        _ = v23
                        _ = v26
                    }
                    v1 += int8(fi21)
                    _ = fi21
                }
            } else {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v12 = v12
                    _ = fi27
                }
            }
        } else {
            var v28 int = 2147483647
            _ = v28
        }
        _ = fi16
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        switch 32768 {
        case 9223372036854775807:
            v3 = -1
        default:
            fmt.Println(v1, v2)
        }
        v30 := []int{}
        switch (fi29 % ((1) | 1)) {
        case 359:
            var v31 uint32 = 278
            _ = v31
        case 657:
            switch "" {
            case "a":
                var v32 int8 = (v1 % (((v1 & (v1 >> ((42) & 63)))) | 1))
                _ = v32
            default:
                var v33 uint64 = (v2 >> ((65535) & 63))
                _ = v33
            }
        }
        _ = fi29
        _ = v30
    }
    if (false && (!((!false) && true))) {
        var v34 float64 = 0.1
        v3 = 2147483647
        _ = v34
    } else {
        var v35 uint64 = (^(v2))
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 string = (("\n\t" + (("hello" + "hello") + (("hello" + ("a" + "a")) + ""))) + "hello")
            fmt.Println(v1, v37)
            for fi38 := 0; fi38 < 1; fi38++ {
                if true {
                    switch "a" {
                    case "hello":
                        v39 := []int{}
                        v40 := GenS1{F0: v1, F1: (^(v3))}
                        _ = v39
                        _ = v40
                    case "":
                        v3 = v3
                    }
                    v41 := []int{fi38, fi38, (fi36 - (fi38 | fi36))}
                    switch -9223372036854775808 {
                    case -116:
                        fmt.Println(fi38, v35)
                    case 32768:
                        v37 = "a"
                    case -9223372036854775808:
                        v42 := GenS1{F0: (-(v1)), F1: -2147483648}
                        _ = v42
                    }
                    _ = v41
                }
                v43 := []int{26, 4294967295, (-(fi36))}
                _ = fi38
                _ = v43
            }
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    if (true || (2147483647 == -438)) {
        switch "hello" {
        case "a":
            v12 += (v12 + v12)
        case "\n\t":
            v1 += int8(v12)
        case "hello":
            fmt.Println(v3)
        default:
            fmt.Println(v1, v3)
        }
        v44 := []int{304, int(v1), 127, 127}
        _ = v44
    } else {
        if (v3 >= (v3 % ((v3) | 1))) {
            if false {
                v1 += (v1 + 100)
                if (100 >= (v1 << ((genFn0(genFn0(genFn0(826)))) & 63))) {
                    for fi45 := 0; fi45 < 16; fi45++ {
                        v46 := GenS3{F0: (-1 >= -32768), F1: 65535, F2: "foo bar"}
                        if true {
                            var v47 bool = (false || ((((3.14159 <= 1e10) || true) || false) && ((v12 + (v12 - -32768)) <= v12)))
                            v48 := GenS4{F0: -32768}
                            _ = v47
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 12; fi49++ {
                            if false {
                                var v50 bool = (true || (!true))
                                v51 := GenS2{F0: (v50 && (753 < 1)), F1: 1, F2: (!(1.5 < 1.5))}
                                v52 := GenS5{F0: 0, F1: 67}
                                if false {
                                    v53 := []int{127, fi49}
                                    v1 -= (-(v1))
                                    _ = v53
                                }
                                var v54 int = fi49
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v54
                            } else {
                                var v55 uint32 = 2
                                var v56 uint16 = 255
                                switch fi45 {
                                case 2147483647:
                                    var v57 float64 = (-2.25 * 3.14159)
                                    _ = v57
                                default:
                                    v58 := []int{(fi49 / ((fi49) | 1)), (fi49 + (fi49 >> ((18446744073709551615) & 63))), (fi49 + (fi49 >> ((1) & 63)))}
                                    _ = v58
                                }
                                _ = v55
                                _ = v56
                            }
                            v59 := GenS6{F0: (^(v1)), F1: v1, F2: 457}
                            fmt.Println(fi45, fi49)
                            _ = fi49
                            _ = v59
                        }
                        _ = fi45
                        _ = v46
                    }
                    var v60 uint32 = 2147483648
                    var v61 float64 = 1e10
                    _ = v60
                    _ = v61
                }
            } else {
                if (255 < 255) {
                    for fi62 := 0; fi62 < 10; fi62++ {
                        if false {
                            v1 += (v1 - 1)
                            for fi63 := 0; fi63 < 14; fi63++ {
                                if false {
                                    var v64 int8 = (v1 / ((v1) | 1))
                                    v65 := GenS1{F0: (^(v1)), F1: v3}
                                    _ = v64
                                    _ = v65
                                }
                                _ = fi63
                            }
                            fmt.Println(v2, v3, v2)
                        } else {
                            v66 := GenS7{F0: 2147483648, F1: 1, F2: 42}
                            v3 -= v3
                            _ = v66
                        }
                        var v67 int16 = 1000
                        v68 := []int{fi62, -9223372036854775808, fi62, -181}
                        _ = fi62
                        _ = v67
                        _ = v68
                    }
                    v12 += v12
                }
            }
            if false {
                for fi69 := 0; fi69 < 19; fi69++ {
                    if false {
                        if (1 < 42) {
                            v70 := GenS7{F0: 2147483647, F1: 663, F2: 341}
                            _ = v70
                        }
                        if (!false) {
                            v71 := GenS3{F0: false, F1: 0, F2: "foo bar"}
                            fmt.Println(v1, v1, v12)
                            for fi72 := 0; fi72 < 4; fi72++ {
                                switch fi72 {
                                case 228:
                                    v73 := []int{(fi69 ^ -582), fi69}
                                    _ = v73
                                case -9223372036854775808:
                                    var v74 int16 = v12
                                    _ = v74
                                default:
                                    v12 = v12
                                }
                                v75 := GenS3{F0: ((255 == 42) && (2147483647 == (fi72 / ((127) | 1)))), F1: 2147483648, F2: ("foo bar" + "a")}
                                _ = fi72
                                _ = v75
                            }
                            _ = v71
                        } else {
                            var v76 int32 = v3
                            fmt.Println(v12)
                            _ = v76
                        }
                        switch "a" {
                        case "hello":
                            v77 := GenS6{F0: (v1 + -1), F1: (v1 / ((v1) | 1)), F2: 18446744073709551615}
                            _ = v77
                        case "foo bar":
                            for fi78 := 0; fi78 < 2; fi78++ {
                                if (false && (!(4294967295 > 0))) {
                                    v1 -= -1
                                    var v79 int64 = 1
                                    var v80 uint16 = 2
                                    _ = v79
                                    _ = v80
                                } else {
                                    v81 := GenS5{F0: 65535, F1: 42}
                                    v12 = v12
                                    v82 := GenS8{F0: true, F1: v2, F2: 65535}
                                    _ = v81
                                    _ = v82
                                }
                                _ = fi78
                            }
                        }
                    }
                    fmt.Println(v12, v2, v12)
                    _ = fi69
                }
                if false {
                    var v83 int = 32767
                    fmt.Println(v3)
                    var v84 int64 = 218
                    _ = v83
                    _ = v84
                }
            }
            var v85 uint = 2
            _ = v85
        } else {
            v2 -= (^(v2))
            switch ("a" + "foo bar") {
            case "foo bar":
                fmt.Println(v2)
            case "hello":
                if (-565 <= -9223372036854775808) {
                    var v86 uint8 = 1
                    v3 = (v3 << ((0) & 63))
                    var v87 int64 = 548
                    _ = v86
                    _ = v87
                }
            case "\n\t":
                for fi88 := 0; fi88 < 20; fi88++ {
                    v12 += v12
                    for fi89 := 0; fi89 < 2; fi89++ {
                        v90 := GenS9{F0: -9223372036854775808}
                        _ = fi89
                        _ = v90
                    }
                    _ = fi88
                }
            }
            var v91 int32 = v3
            _ = v91
        }
        for fi92 := 0; fi92 < 12; fi92++ {
            fmt.Println(v2, fi92, v12)
            _ = fi92
        }
    }
    var v93 uint64 = 684
    switch "\n\t" {
    case "foo bar":
        switch "a" {
        case "\n\t":
            v94 := []int{1}
            _ = v94
        default:
            if ((!((-280 == -196) && true)) || (2 != v2)) {
                v1 = 127
                v95 := GenS10{F0: 188}
                _ = v95
            }
        }
    }
    v96 := []int{42}
    _ = len(v96)
    fmt.Println(v3)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v93:", v93)
    fmt.Println("slicelen v96:", len(v96), v96)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v93
    _ = v96
}
