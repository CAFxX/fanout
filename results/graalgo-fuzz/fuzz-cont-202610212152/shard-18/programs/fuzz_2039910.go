package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 int32 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int32; F1 uint; F2 string }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 uint64; F1 int64; F2 int64 }
func genFn0(p0 string, p1 int8, p2 string) int8 {
    v6 := GenS1{F0: 127, F1: 303}
    var v7 int = 278
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint64 {
    if (0 > 65535) {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int32 = 900
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: 0, F1: 255, F2: 0}
        _ = v10
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 4; fi12++ {
        switch -2147483648 {
        case -1:
            var v13 uint8 = 0
            _ = v13
        case 141:
            if true {
                fmt.Println(fi12)
                v14 := GenS1{F0: (^(fi12)), F1: 42}
                _ = v14
            } else {
                var v15 int16 = 1000
                v16 := GenS3{F0: -1, F1: (fi12 >> ((255) & 63)), F2: (-1 != (fi12 * -527))}
                _ = v15
                _ = v16
            }
        default:
            switch fi12 {
            case 32767:
                if (2147483648 == 2) {
                    var v17 uint32 = 0
                    switch (fi12 ^ fi12) {
                    case -2147483648:
                        if len(v11) > 0 {
                            v11[0] = 9223372036854775807
                        }
                    default:
                        if true {
                            v11 = append(v11, fi12)
                        }
                    }
                    _ = v17
                }
            default:
                if false {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi12)
                    }
                } else {
                    var v18 int8 = 42
                    var v19 int = (^(fi12))
                    _ = v18
                    _ = v19
                }
            }
        }
        var v20 uint32 = 42
        fmt.Println(v20, fi12)
        _ = fi12
        _ = v20
    }
    if len(v11) > 0 {
        v11[0] = -9223372036854775808
    }
    switch 338 {
    case 4294967295:
        var v21 int32 = 1
        _ = v21
    case 798:
        if true {
            switch -1 {
            case -491:
                switch 833 {
                case -9223372036854775808:
                    if ((1 <= 752) && ((105 > 42) && (255 <= 0))) {
                        var v22 string = "foo bar"
                        v22 = v22
                        _ = v22
                    } else {
                        var v23 int = 9223372036854775807
                        if (!false) {
                            var v24 uint64 = 717
                            if ((0.1 <= (1e10 + (0.0 * (1e10 * 0.1)))) || ((3.14159 - 3.14159) == 3.14159)) {
                                var v25 float64 = ((0.1 / float64(1e10)) - (((float64(-2.25) * ((0.0 * (((1.5 * -2.25) + 1.5) / ((0.1 / (float64((-2.25 * 3.14159)) / (1e10 * -2.25))) + 3.14159))) + 1e10)) * 0.0) - 0.0))
                                v26 := GenS2{F0: 103, F1: 163, F2: -1}
                                v23 = v23
                                _ = v25
                                _ = v26
                            }
                            _ = v24
                        } else {
                            v11 = append(v11, 209)
                            var v27 float64 = 1.5
                            v27 = (v27 + v27)
                            _ = v27
                        }
                        switch ("foo bar" + "\n\t") {
                        case "":
                            fmt.Println(v23)
                        case "a":
                            v28 := GenS3{F0: 0, F1: (v23 % ((v23) | 1)), F2: true}
                            _ = v28
                        case "\n\t":
                            var v29 float32 = (3.14159 / (((1.5 + 1.5) / 1.0) - ((-2.25 * 0.1) * -2.25)))
                            _ = v29
                        }
                        _ = v23
                    }
                case 2147483647:
                    var v30 float32 = (1.5 + -2.25)
                    _ = v30
                default:
                    var v31 int32 = -855
                    _ = v31
                }
            default:
                for fi32 := 0; fi32 < 4; fi32++ {
                    if true {
                        if true {
                            switch ("a" + "hello") {
                            case "\n\t":
                                fmt.Println(fi32)
                            case "foo bar":
                                _ = len(v11)
                            case "hello":
                                var v33 uint8 = 0
                                _ = v33
                            default:
                                for fi34 := 0; fi34 < 14; fi34++ {
                                    if ((4 < 1) || false) {
                                        var v35 uint = 4294967295
                                        var v36 uint16 = 1
                                        _ = v35
                                        _ = v36
                                    }
                                    v37 := GenS4{F0: "hello", F1: 18446744073709551615}
                                    _ = fi34
                                    _ = v37
                                }
                            }
                            v38 := GenS5{F0: -1, F1: 221, F2: ((("foo bar" + ("\n\t" + "hello")) + "\n\t") + "a")}
                            var v39 float32 = 1e10
                            _ = v38
                            _ = v39
                        }
                    }
                    var v40 uint8 = 127
                    _ = fi32
                    _ = v40
                }
            }
            var v41 float64 = ((3.14159 - 0.1) * float64((1e10 * (1e10 - 1e10))))
            var v42 int64 = 993
            _ = v41
            _ = v42
        }
    case -2147483648:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    default:
        for fi43 := 0; fi43 < 9; fi43++ {
            if ((("a" + (("" + ("hello" + ("hello" + (("a" + "") + "a")))) + "a")) + "") > "\n\t") {
                var v44 int = fi43
                var v45 uint32 = 2147483648
                _ = v44
                _ = v45
            } else {
                v11 = append(v11, fi43)
                if false {
                    if true {
                        var v46 uint32 = 541
                        if (42 > 42) {
                            fmt.Println(fi43, fi43)
                            v47 := GenS2{F0: 1, F1: 528, F2: 42}
                            var v48 int = -9223372036854775808
                            _ = v47
                            _ = v48
                        }
                        _ = v46
                    }
                    var v49 int8 = -1
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 18; fi50++ {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        var v51 float64 = (((0.0 + 1.5) - 3.14159) - (1e10 / -2.25))
                        _ = fi50
                        _ = v51
                    }
                }
                var v52 uint = 18446744073709551615
                _ = v52
            }
            var v53 float64 = float64(-2.25)
            _ = fi43
            _ = v53
        }
    }
    _ = v11
    return 18446744073709551615
}

func main() {
    var v1 int = 915
    var v2 uint32 = 2147483648
    var v3 uint = 65535
    var v4 int = v1
    v5 := []int{v4, (v4 / ((v1) | 1))}
    switch (v1 | v1) {
    case 4294967295:
        v54 := GenS3{F0: -795, F1: v4, F2: false}
        _ = v54
    case 2147483647:
        v1 = v1
    }
    v4 = v4
    var v55 uint32 = v2
    var v56 bool = (80 < 16)
    if false {
        if v56 {
            v57 := GenS6{F0: v1, F1: (^(v3))}
            _ = v57
        } else {
            for fi58 := 0; fi58 < 11; fi58++ {
                for fi59 := 0; fi59 < 20; fi59++ {
                    if v56 {
                        var v60 uint = v3
                        v56 = (v4 != fi59)
                        v61 := GenS7{F0: -1, F1: -2147483648}
                        _ = v60
                        _ = v61
                    }
                    _ = fi59
                }
                v56 = ((v56 || (42 < 18446744073709551615)) && (42 > -449))
                if (v56 && v56) {
                    var v62 int = 1
                    _ = v62
                }
                _ = fi58
            }
        }
        switch v1 {
        case 32767:
            var v63 string = "hello"
            _ = v63
        case 2147483647:
            var v64 int16 = 1
            _ = v64
        case 4294967295:
            for fi65 := 0; fi65 < 2; fi65++ {
                var v66 bool = false
                var v67 int16 = 32767
                var v68 float64 = (-2.25 / 3.14159)
                _ = fi65
                _ = v66
                _ = v67
                _ = v68
            }
        }
        var v69 string = "\n\t"
        _ = v69
    }
    if v56 {
        switch "hello" {
        case "foo bar":
            var v70 uint8 = uint8(v3)
            _ = v70
        case "":
            if (204 <= -96) {
                if (!(0 == -465)) {
                    fmt.Println(v55)
                    if true {
                        for fi71 := 0; fi71 < 18; fi71++ {
                            v4 += v4
                            for fi72 := 0; fi72 < 1; fi72++ {
                                if (("\n\t" + "\n\t") >= ("a" + "")) {
                                    v5 = append(v5, fi72)
                                    v56 = true
                                }
                                v73 := GenS7{F0: 797, F1: -929}
                                var v74 string = "a"
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            v75 := GenS6{F0: v1, F1: (v3 + (^(v3)))}
                            _ = fi71
                            _ = v75
                        }
                        for fi76 := 0; fi76 < 15; fi76++ {
                            v4 -= 32767
                            v77 := GenS4{F0: "a", F1: v3}
                            v4 += (fi76 ^ v1)
                            _ = fi76
                            _ = v77
                        }
                        var v78 float64 = ((1.5 / float64(((0.1 + -2.25) / (0.1 * 1.5)))) * (float64(-2.25) * -2.25))
                        _ = v78
                    } else {
                        _ = len(v5)
                    }
                    var v79 uint = (^(v3))
                    _ = v79
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = len(v5)
            }
        case "a":
            var v80 uint16 = 1
            _ = v80
        }
        for fi81 := 0; fi81 < 19; fi81++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi82 := 0; fi82 < 7; fi82++ {
                var v83 uint8 = 0
                v1 = v1
                _ = fi82
                _ = v83
            }
            _ = fi81
        }
        var v84 uint64 = genFn1()
        _ = v84
    }
    var v85 int = (-(v1))
    if v56 {
        for fi86 := 0; fi86 < 6; fi86++ {
            if len(v5) > 0 {
                v5[0] = 127
            }
            switch (("hello" + "foo bar") + "a") {
            case "":
                v85 = (-(v4))
            case "hello":
                var v87 bool = ((!(42 > 647)) && v56)
                _ = v87
            case "\n\t":
                for fi88 := 0; fi88 < 5; fi88++ {
                    fmt.Println(v4, fi86, v3)
                    if (0 >= 1) {
                        var v89 float64 = 0.0
                        _ = v89
                    }
                    v5 = append(v5, (v4 - (v85 % ((4294967295) | 1))))
                    _ = fi88
                }
            }
            var v90 uint16 = 0
            _ = fi86
            _ = v90
        }
        if (0 > 303) {
            var v91 uint8 = 6
            fmt.Println(v55, v56)
            v92 := GenS4{F0: "", F1: v3}
            _ = v91
            _ = v92
        }
        var v93 int32 = 2147483647
        _ = v93
    } else {
        v94 := GenS8{F0: v2, F1: -1}
        var v95 uint = (v3 | (v3 | (v3 & v3)))
        _ = v94
        _ = v95
    }
    var v96 float32 = (1.5 * 1.5)
    var v97 int8 = 100
    var v98 int64 = 565
    for fi99 := 0; fi99 < 12; fi99++ {
        v100 := GenS1{F0: v1, F1: v2}
        for fi101 := 0; fi101 < 20; fi101++ {
            switch ("a" + "\n\t") {
            case "a":
                v56 = (0.1 < (0.0 + 1e10))
            }
            _ = fi101
        }
        for fi102 := 0; fi102 < 12; fi102++ {
            fmt.Println(v98, v98, v85)
            _ = fi102
        }
        _ = fi99
        _ = v100
    }
    v55 = (v55 | v2)
    if (!v56) {
        var v103 int8 = int8(v98)
        v104 := GenS6{F0: (v1 + (^(v4))), F1: (v3 ^ v3)}
        _ = v103
        _ = v104
    } else {
        fmt.Println(v98, v98)
        if ((true || ((0.0 - (3.14159 * 1.5)) <= float64(float32(float64(v96))))) || false) {
            _ = len(v5)
            var v105 string = "a"
            _ = v105
        }
    }
    v106 := GenS4{F0: "a", F1: v3}
    v107 := GenS9{F0: 4294967295, F1: -1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v85:", v85)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v85
    _ = v96
    _ = v97
    _ = v98
}
