package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 uint64; F1 uint; F2 bool }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 int8; F1 uint64 }
func genFn0(p0 int8, p1 uint8, p2 int8) uint8 {
    v4 := GenS1{F0: 4294967295, F1: 65535}
    for fi5 := 0; fi5 < 13; fi5++ {
        p1 -= p1
        _ = fi5
    }
    fmt.Println(p0)
    switch "hello" {
    case "hello":
        for fi6 := 0; fi6 < 15; fi6++ {
            p2 = p2
            var v7 int16 = -487
            _ = fi6
            _ = v7
        }
    case "":
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint8 = (p1 - (p1 * p1))
            _ = fi8
            _ = v9
        }
    default:
        v10 := []int{9223372036854775807, -501, 0, -9223372036854775808}
        _ = v10
    }
    v11 := GenS1{F0: 396, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v4
    return 2
}
func genFn1() uint64 {
    v12 := []int{42, -1, -487, 0}
    switch ("\n\t" + ("" + "foo bar")) {
    case "a":
        for fi13 := 0; fi13 < 2; fi13++ {
            if false {
                _ = len(v12)
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS2{F0: 65535}
                    _ = fi14
                    _ = v15
                }
            }
            var v16 uint64 = 149
            var v17 uint8 = 0
            _ = fi13
            _ = v16
            _ = v17
        }
    case "hello":
        v18 := GenS3{F0: (391 > 309), F1: 2147483648}
        _ = v18
    default:
        var v19 uint = 65535
        _ = v19
    }
    _ = v12
    return 900
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 59
    v3 := make(map[string]bool)
    for fi20 := 0; fi20 < 14; fi20++ {
        v1 = (v1 >> ((18446744073709551615) & 63))
        if ((!true) || true) {
            v2 -= v2
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            if false {
                if false {
                    switch 32768 {
                    case 9223372036854775807:
                        v2 -= genFn0(-24, 197, 42)
                    default:
                        var v22 int16 = -1
                        v22 += v22
                        _ = v22
                    }
                } else {
                    v23 := GenS2{F0: 604}
                    if true {
                        var v24 uint16 = 681
                        var v25 float32 = -2.25
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS4{F0: v1}
                        _ = v26
                    }
                    var v27 float32 = float32((-2.25 + 0.0))
                    if (!((-2.25 - (float32((3.14159 - (1.5 * 0.1))) / ((v27 - ((v27 * (0.0 / (1e10 - 0.0))) * 0.1)) * 1.5))) >= 3.14159)) {
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v1, v2, fi20)
                        case "a":
                            v1 -= (v1 % ((v1) | 1))
                        }
                        v28 := []int{(fi21 + (fi20 * fi21)), -2147483648}
                        _ = v28
                    } else {
                        if true {
                            v2 = (v2 * (v2 - 2))
                        }
                        if (!true) {
                            _ = len(v3)
                            v29 := GenS1{F0: 0, F1: 42}
                            _ = v29
                        }
                    }
                    v27 = v27
                    _ = v23
                    _ = v27
                }
                v30 := GenS2{F0: 18446744073709551615}
                _ = v30
            } else {
                if true {
                    fmt.Println(v1, fi20)
                } else {
                    if true {
                        var v31 int = fi20
                        if false {
                            fmt.Println(fi21, v1, v1)
                            v32 := []int{(fi20 >> ((181) & 63)), fi20, -261, fi20}
                            v33 := GenS1{F0: 2147483647, F1: 2147483647}
                            var v34 string = (("hello" + "\n\t") + "\n\t")
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        v2 -= 128
                        _ = v31
                    }
                    if (v2 <= (v2 * (v2 >> ((255) & 63)))) {
                        if false {
                            v3[(("a" + ("foo bar" + "")) + "")] = false
                        } else {
                            v2 = v2
                        }
                    } else {
                        fmt.Println(fi20)
                    }
                    v1 -= (v1 << ((65535) & 63))
                }
            }
            _ = fi21
        }
        _ = fi20
    }
    v3[(("\n\t" + "a") + ((((("hello" + ("foo bar" + ((("\n\t" + "foo bar") + ("" + "hello")) + "hello"))) + "foo bar") + "a") + ((("" + ("a" + "foo bar")) + "\n\t") + ("\n\t" + "a"))) + ("a" + "foo bar")))] = (-2.25 < -2.25)
    if false {
        fmt.Println(v1, v2)
        for fi35 := 0; fi35 < 2; fi35++ {
            if (1e10 <= 0.0) {
                if (3.14159 < 0.1) {
                    v3[""] = ((!((1e10 / 1e10) >= -2.25)) || true)
                    for fi36 := 0; fi36 < 5; fi36++ {
                        switch (("hello" + "") + "hello") {
                        case "hello":
                            var v37 float64 = 3.14159
                            _ = v37
                        case "\n\t":
                            if true {
                                var v38 float32 = 0.1
                                var v39 float64 = 0.0
                                fmt.Println(v39, v38, v38)
                                switch fi35 {
                                case 708:
                                    var v40 uint8 = (v2 % ((v2) | 1))
                                    _ = v40
                                default:
                                    switch ("hello" + ("foo bar" + ("" + (("hello" + (("\n\t" + ("\n\t" + "foo bar")) + "a")) + "foo bar")))) {
                                    case "a":
                                        v41 := GenS5{F0: 42, F1: 2147483648}
                                        _ = v41
                                    default:
                                        v38 -= v38
                                    }
                                    var v42 int = 127
                                    _ = v42
                                }
                                _ = v38
                                _ = v39
                            }
                        case "foo bar":
                            fmt.Println(v1, fi35)
                        default:
                            var v43 uint64 = 572
                            _ = v43
                        }
                        if true {
                            if (((!true) || true) && true) {
                                if (false && ((^(v2)) != (v2 ^ (^(v2))))) {
                                    v1 += v1
                                    v2 += v2
                                } else {
                                    var v44 int8 = -128
                                    _ = v44
                                }
                                var v45 uint64 = 0
                                _ = v45
                            }
                        }
                        _ = fi36
                    }
                }
                v46 := GenS6{F0: 0, F1: -1}
                v47 := []int{(fi35 * (fi35 << ((350) & 63))), fi35, (fi35 + (fi35 << ((803) & 63))), (fi35 >> ((780) & 63))}
                _ = v46
                _ = v47
            } else {
                switch (^(fi35)) {
                case -876:
                    v2 -= (v2 * genFn0(int8(fi35), (v2 << ((uint(v2)) & 63)), -1))
                case 128:
                    var v48 uint8 = v2
                    _ = v48
                case -157:
                    var v49 int8 = -1
                    _ = v49
                }
                v1 = (v1 & v1)
                fmt.Println(v2, v1, v1)
            }
            _ = fi35
        }
        v50 := []int{4294967295, -655}
        _ = v50
    }
    v51 := GenS4{F0: (v1 % ((v1) | 1))}
    var v52 int = 2147483647
    v53 := GenS2{F0: 231}
    v54 := GenS7{F0: 170, F1: 65535, F2: false}
    var v55 int16 = 1
    fmt.Println(v52, v55)
    var v56 string = ("" + ("hello" + "foo bar"))
    if (((-2.25 / float64((((-2.25 - 1e10) + 1e10) / float32(1e10)))) * 0.1) >= 1.5) {
        var v57 int16 = v55
        _ = v57
    } else {
        if false {
            for fi58 := 0; fi58 < 7; fi58++ {
                if (v56 != "\n\t") {
                    for fi59 := 0; fi59 < 9; fi59++ {
                        var v60 uint16 = 2
                        switch v56 {
                        case "":
                            if true {
                                v1 = v1
                                v1 -= (v1 << ((42) & 63))
                                v55 = v55
                            }
                        case "foo bar":
                            var v61 float64 = 3.14159
                            _ = v61
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                v1 -= (v1 >> ((377) & 63))
                            }
                        }
                        v56 = "hello"
                        _ = fi59
                        _ = v60
                    }
                } else {
                    for fi62 := 0; fi62 < 8; fi62++ {
                        switch v52 {
                        case -305:
                            switch (v56 + (("foo bar" + v56) + v56)) {
                            case "a":
                                v3[(v56 + ("" + ("hello" + "")))] = true
                            case "hello":
                                v63 := GenS5{F0: (v1 * v1), F1: 42}
                                _ = v63
                            case "":
                                var v64 uint64 = 878
                                _ = v64
                            }
                        case 32767:
                            v65 := GenS8{F0: 675, F1: v2}
                            _ = v65
                        case 128:
                            var v66 bool = (32767 != uint16(v2))
                            _ = v66
                        }
                        v52 = fi58
                        v52 = fi62
                        v67 := GenS9{F0: -9223372036854775808, F1: (v52 & (v52 << ((224) & 63)))}
                        v56 = v56
                        _ = fi62
                        _ = v67
                    }
                }
                for fi68 := 0; fi68 < 19; fi68++ {
                    var v69 uint16 = 65535
                    _ = fi68
                    _ = v69
                }
                v2 += v2
                _ = fi58
            }
            for fi70 := 0; fi70 < 7; fi70++ {
                for fi71 := 0; fi71 < 17; fi71++ {
                    v72 := v3["hello"]
                    _ = fi71
                    _ = v72
                }
                v52 -= 0
                if true {
                    if true {
                        v73 := []int{v52, -890}
                        _ = v73
                    }
                }
                _ = fi70
            }
        }
    }
    if (true && (true || (18446744073709551615 > 4294967295))) {
        fmt.Println(v52, v52, v56)
        fmt.Println(v55, v2)
    } else {
        v74 := []int{v52, -9223372036854775808, v52, 128}
        switch -7 {
        case 32767:
            var v75 int8 = 127
            _ = v75
        case 128:
            v76 := GenS10{F0: 0, F1: 18446744073709551615}
            _ = v76
        default:
            v55 = 1000
        }
        v56 = (v56 + "\n\t")
        _ = v74
    }
    fmt.Println(v56)
    var v77 float64 = float64(((3.14159 + 3.14159) - float32(-2.25)))
    if true {
        if (2 >= 855) {
            v55 -= v55
            switch v52 {
            case -9223372036854775808:
                for fi78 := 0; fi78 < 5; fi78++ {
                    for fi79 := 0; fi79 < 14; fi79++ {
                        var v80 int64 = 505
                        var v81 uint = 255
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                    _ = fi78
                }
            case 128:
                if true {
                    var v82 float64 = 1e10
                    _ = v82
                } else {
                    var v83 int32 = v1
                    _ = v83
                }
            case -2147483648:
                v84 := v3[(v56 + ("a" + v56))]
                _ = v84
            }
            var v85 uint32 = 428
            _ = v85
        } else {
            var v86 int16 = -940
            v1 = (v1 - (^(v1)))
            _ = v86
        }
    }
    var v87 uint32 = 990
    for fi88 := 0; fi88 < 8; fi88++ {
        v55 -= -32768
        v3[v56] = (!true)
        var v89 uint64 = 42
        _ = fi88
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v77:", v77)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v77
    _ = v87
}
