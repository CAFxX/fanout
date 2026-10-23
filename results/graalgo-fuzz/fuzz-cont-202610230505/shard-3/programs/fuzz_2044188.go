package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 int64 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint }
func genFn0(p0 uint16, p1 int32, p2 uint8) uint8 {
    v7 := []int{}
    if false {
        if (!true) {
            var v8 uint = 1
            p1 = -857
            p0 = p0
            _ = v8
        }
        p2 -= 133
        if (-1 > 81) {
            p1 = 1
            switch "" {
            case "hello":
                if false {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        fmt.Println(p2, fi9, p1)
                        _ = fi9
                    }
                    fmt.Println(p0, p1)
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            default:
                p0 = 13
            }
        } else {
            v7 = append(v7, -493)
            if false {
                switch 2147483647 {
                case -641:
                    fmt.Println(p1, p0)
                case 125:
                    v10 := GenS1{F0: "hello", F1: (p0 % ((p0) | 1))}
                    _ = v10
                default:
                    switch (((("\n\t" + "a") + "foo bar") + ((("" + "\n\t") + ("\n\t" + ("" + ((((("foo bar" + "") + "\n\t") + ("a" + ("a" + "\n\t"))) + "foo bar") + "foo bar")))) + "\n\t")) + "") {
                    case "a":
                        p2 += 202
                    case "hello":
                        var v11 uint = 2
                        _ = v11
                    }
                }
            }
            v7 = append(v7, 4294967295)
        }
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: "a", F1: (p0 & 65535)}
        if false {
            v14 := GenS2{F0: -899, F1: (0 >= 4294967295), F2: 1}
            _ = v14
        }
        fmt.Println(fi12, fi12)
        _ = fi12
        _ = v13
    }
    if false {
        var v15 uint = 174
        var v16 uint8 = 0
        v17 := GenS2{F0: 0, F1: false, F2: -961}
        _ = v15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 225
}
func genFn1(p0 int64) int32 {
    v18 := []int{32767, 4294967295, 9223372036854775807, -9223372036854775808}
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS1{F0: "foo bar", F1: 934}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            if ((fi20 - (fi19 ^ fi19)) != fi20) {
                p0 = (p0 >> ((4294967295) & 63))
                if ((((0.0 * 1.5) == 1e10) && (true || ((fi19 + (^(fi20))) > -2147483648))) || false) {
                    fmt.Println(fi20, fi19)
                    var v22 int16 = -32768
                    if false {
                        fmt.Println(fi20)
                        fmt.Println(p0, fi19, fi19)
                        v22 = 913
                    }
                    _ = v22
                }
            } else {
                var v23 string = ""
                _ = len(v18)
                v24 := GenS2{F0: 32767, F1: (((false && true) || false) && false), F2: 42}
                v23 = v23
                fmt.Println(p0, p0, fi19)
                if true {
                    v23 = ((v23 + (v23 + v23)) + v23)
                } else {
                    if (!(861 != 0)) {
                        var v25 uint64 = 0
                        v23 = v23
                        _ = v25
                    } else {
                        var v26 int32 = 677
                        fmt.Println(v23, fi19, v26)
                        var v27 float64 = -2.25
                        v26 = (v26 | v26)
                        v28 := GenS3{F0: ("a" > (v23 + ("hello" + v23))), F1: 0}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
                var v29 uint8 = 42
                switch "\n\t" {
                case "":
                    var v30 int16 = 1
                    _ = v30
                case "hello":
                    var v31 uint = 476
                    _ = v31
                default:
                    v29 = genFn0(255, 0, 255)
                }
                var v32 uint8 = genFn0(42, 1, (v29 >> ((255) & 63)))
                v23 = v23
                var v33 int8 = 66
                _ = v23
                _ = v24
                _ = v29
                _ = v32
                _ = v33
            }
            _ = fi20
            _ = v21
        }
        p0 += (p0 / ((9223372036854775807) | 1))
        _ = fi19
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        switch fi34 {
        case 128:
            var v35 bool = (-106 < 127)
            _ = v35
        case -2147483648:
            var v36 float64 = 0.0
            _ = v36
        case -375:
            v18 = append(v18, fi34)
        default:
            var v37 int32 = 295
            _ = v37
        }
        if (!(true && (!true))) {
            if ((fi34 % (((fi34 / ((fi34) | 1))) | 1)) <= 128) {
                fmt.Println(fi34, fi34)
                p0 += -446
                p0 = p0
            }
        }
        _ = fi34
    }
    fmt.Println(p0)
    _ = p0
    _ = v18
    return 1
}

func main() {
    v1 := []string{"\n\t", "foo bar", ("\n\t" + "a"), "\n\t"}
    var v2 int = 472
    var v3 string = "hello"
    var v4 int32 = 241
    var v5 int64 = -962
    v6 := []int{(v2 << ((888) & 63)), v2}
    fmt.Println(v2, v3)
    switch v3 {
    case "hello":
        v2 += v2
    case "foo bar":
        v2 = v2
    default:
        v38 := GenS4{F0: -9223372036854775808, F1: -60}
        _ = v38
    }
    v4 += 674
    if (true && true) {
        _ = len(v1)
        switch "\n\t" {
        case "a":
            for fi39 := 0; fi39 < 10; fi39++ {
                var v40 int64 = v5
                var v41 uint = 255
                _ = fi39
                _ = v40
                _ = v41
            }
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                switch "hello" {
                case "hello":
                    v42 := GenS5{F0: 4294967295}
                    _ = v42
                }
            case "a":
                if (!(!(((((("hello" + "\n\t") + (v3 + ("foo bar" + ""))) + "") + v3) + (("foo bar" + "\n\t") + ("\n\t" + ("a" + v3)))) < (("\n\t" + v3) + v3)))) {
                    v4 = (v4 - 2147483647)
                } else {
                    if true {
                        if false {
                            if true {
                                if true {
                                    for fi43 := 0; fi43 < 8; fi43++ {
                                        switch -610 {
                                        case 32767:
                                            v4 -= v4
                                        case 715:
                                            if ((42 <= 127) && false) {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            } else {
                                                fmt.Println(v5, v3, v5)
                                                var v44 int32 = (v4 + 42)
                                                fmt.Println(v44)
                                                _ = v44
                                            }
                                        default:
                                            var v45 uint64 = 1
                                            _ = v45
                                        }
                                        v6 = append(v6, (v2 << ((174) & 63)))
                                        _ = fi43
                                    }
                                    switch (v3 + (("\n\t" + v3) + "\n\t")) {
                                    case "\n\t":
                                        v5 = v5
                                    case "a":
                                        if false {
                                            if ((3.14159 * (0.0 / 3.14159)) >= 1.5) {
                                                for fi46 := 0; fi46 < 1; fi46++ {
                                                    var v47 int8 = 24
                                                    if len(v1) > 0 {
                                                        v1[0] = v3
                                                    }
                                                    v48 := GenS6{F0: -1}
                                                    _ = fi46
                                                    _ = v47
                                                    _ = v48
                                                }
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                var v49 uint = 65535
                                                var v50 int32 = (v4 >> ((0) & 63))
                                                _ = v49
                                                _ = v50
                                            }
                                            var v51 uint = 429
                                            _ = v51
                                        }
                                    case "":
                                        v2 += int(v5)
                                    default:
                                        var v52 uint = 1
                                        _ = v52
                                    }
                                    v53 := GenS7{F0: (false && true)}
                                    _ = v53
                                }
                                v3 = ""
                                fmt.Println(v3, v2)
                            } else {
                                v54 := GenS6{F0: v5}
                                var v55 uint16 = 32767
                                v56 := GenS2{F0: 24, F1: false, F2: v5}
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                        } else {
                            v5 -= (v5 | -9223372036854775808)
                        }
                        v5 = v5
                    }
                    if false {
                        v3 = "foo bar"
                        var v57 float64 = 0.1
                        v58 := GenS8{F0: 128, F1: 1}
                        _ = v57
                        _ = v58
                    } else {
                        v3 = (v3 + ("a" + "hello"))
                        for fi59 := 0; fi59 < 9; fi59++ {
                            var v60 bool = (32767 != 1)
                            var v61 float32 = 1.5
                            switch v3 {
                            case "a":
                                v4 += (v4 ^ (v4 % (((v4 ^ (v4 << ((99) & 63)))) | 1)))
                            case "":
                                for fi62 := 0; fi62 < 16; fi62++ {
                                    fmt.Println(v2)
                                    _ = fi62
                                }
                            case "hello":
                                fmt.Println(v61)
                            default:
                                if (65535 <= 571) {
                                    if v60 {
                                        if (561 <= 2) {
                                            for fi63 := 0; fi63 < 13; fi63++ {
                                                v2 -= v2
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                fmt.Println(v61, v3, v61)
                                                var v64 uint32 = 34
                                                _ = fi63
                                                _ = v64
                                            }
                                            _ = len(v6)
                                        } else {
                                            v60 = (0.0 > float32(float64(3.14159)))
                                            v65 := GenS9{F0: -610}
                                            var v66 int16 = 779
                                            _ = v65
                                            _ = v66
                                        }
                                        var v67 string = v3
                                        _ = v67
                                    }
                                } else {
                                    fmt.Println(v61, v3, v2)
                                    v3 = ""
                                }
                            }
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        if (((v3 + v3) + "hello") != v3) {
                            v68 := GenS1{F0: (v3 + (((v3 + "") + v3) + v3)), F1: 1}
                            if true {
                                fmt.Println(v2)
                            }
                            _ = v68
                        }
                    }
                }
            case "":
                switch v3 {
                case "\n\t":
                    switch v2 {
                    case 814:
                        if (false || (416 >= 2)) {
                            switch "" {
                            case "a":
                                var v69 bool = true
                                _ = v69
                            case "hello":
                                switch v2 {
                                case 805:
                                    fmt.Println(v2)
                                case 127:
                                    v4 = -1
                                case 664:
                                    var v70 float32 = 1.5
                                    _ = v70
                                }
                            case "":
                                if (!true) {
                                    if len(v6) > 0 {
                                        v6[0] = v2
                                    }
                                    v5 = (v5 & v5)
                                }
                            default:
                                fmt.Println(v2)
                            }
                            if len(v1) > 0 {
                                v1[0] = (v3 + v3)
                            }
                            v71 := GenS5{F0: 18446744073709551615}
                            _ = v71
                        }
                    case -1:
                        v1 = append(v1, (v3 + v3))
                    default:
                        if true {
                            var v72 uint64 = 0
                            _ = v72
                        }
                    }
                case "hello":
                    var v73 int64 = v5
                    _ = v73
                }
            }
        case "\n\t":
            var v74 uint64 = 294
            _ = v74
        default:
            var v75 int16 = 94
            _ = v75
        }
    }
    if false {
        v2 -= (v2 ^ (^(v2)))
        if false {
            v2 -= int(v5)
            var v76 string = "hello"
            switch 1 {
            case -9223372036854775808:
                var v77 uint16 = 65535
                _ = v77
            case 199:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 1:
                v3 = v76
            }
            _ = v76
        }
        for fi78 := 0; fi78 < 16; fi78++ {
            v2 += 43
            _ = fi78
        }
    } else {
        v2 = (^(v2))
        v4 = v4
    }
    v79 := GenS10{F0: 16}
    v3 = (v3 + v3)
    v4 = (-(v4))
    fmt.Println(v5)
    if (-1 > 1) {
        fmt.Println(v3)
        for fi80 := 0; fi80 < 11; fi80++ {
            for fi81 := 0; fi81 < 8; fi81++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v82 int8 = 100
                _ = fi81
                _ = v82
            }
            _ = fi80
        }
    }
    var v83 uint8 = 163
    fmt.Println(v5, v83, v83)
    if true {
        if (-32768 == -32768) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v84 uint64 = 163
            _ = v84
        }
    } else {
        var v85 string = v3
        v3 = ""
        _ = v85
    }
    v4 -= 231
    var v86 uint8 = v83
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v83:", v83)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v83
    _ = v86
}
