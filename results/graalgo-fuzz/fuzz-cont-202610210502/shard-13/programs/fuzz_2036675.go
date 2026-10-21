package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string; F1 int8; F2 uint8 }
type GenS4 struct { F0 int; F1 bool; F2 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 uint; F2 int32 }
type GenS10 struct { F0 int16; F1 string }
type GenS11 struct { F0 uint8; F1 uint32 }
type GenS12 struct { F0 int8 }
func genFn0(p0 uint16) uint16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        if true {
            p0 = 640
        }
        _ = fi5
    }
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 19; fi7++ {
        if true {
            if false {
                if (-1 >= -9223372036854775808) {
                    p0 = p0
                }
                if true {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        var v9 uint64 = 2
                        v10 := []int{(fi8 % (((fi8 / ((fi8) | 1))) | 1)), (fi7 ^ int(p0))}
                        v9 = v9
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    p0 += (^(p0))
                } else {
                    p0 += (p0 << ((1) & 63))
                    fmt.Println(fi7)
                }
            } else {
                v11 := []int{(fi7 + 9223372036854775807)}
                _ = v11
            }
        } else {
            var v12 float32 = float32(0.0)
            if (!(-2147483648 <= -2147483648)) {
                if (0 <= 127) {
                    var v13 uint16 = 0
                    fmt.Println(fi7, v12)
                    _ = v13
                } else {
                    if false {
                        if (!(!(!true))) {
                            var v14 uint16 = 209
                            fmt.Println(fi7)
                            var v15 int8 = 1
                            _ = v14
                            _ = v15
                        }
                        p0 -= 2
                    }
                    v16 := []int{(fi7 % ((-491) | 1)), fi7, 9223372036854775807}
                    v17 := GenS1{F0: -9223372036854775808}
                    _ = v16
                    _ = v17
                }
                fmt.Println(p0)
                fmt.Println(p0, v12, p0)
            } else {
                v18 := GenS1{F0: -150}
                _ = v18
            }
            v19 := GenS2{F0: 42, F1: 136, F2: (p0 >> ((1) & 63))}
            _ = v12
            _ = v19
        }
        switch fi7 {
        case -9223372036854775808:
            if false {
                switch fi7 {
                case 32767:
                    var v20 int32 = -421
                    _ = v20
                case -9223372036854775808:
                    p0 = p0
                }
                p0 = (p0 - (p0 * p0))
            }
        case -1:
            if true {
                v21 := []int{fi7, 4294967295, (fi7 ^ fi7), (^(fi7))}
                p0 = (p0 + (p0 % ((p0) | 1)))
                _ = v21
            }
        default:
            v22 := GenS1{F0: 1}
            _ = v22
        }
        p0 = p0
        _ = fi7
    }
    var v23 float64 = 1.5
    if (0.0 == 1e10) {
        switch "" {
        case "foo bar":
            v23 = v23
        case "hello":
            for fi24 := 0; fi24 < 15; fi24++ {
                p0 += (^(p0))
                if ((false || true) && (-1 <= 1)) {
                    var v25 uint8 = 43
                    v25 += 26
                    v25 = v25
                    _ = v25
                } else {
                    var v26 int = (fi24 / ((-2147483648) | 1))
                    switch "a" {
                    case "foo bar":
                        var v27 int32 = 42
                        _ = v27
                    default:
                        if ((^(p0)) != p0) {
                            var v28 uint8 = 95
                            var v29 int32 = 2147483647
                            if false {
                                var v30 int16 = 32767
                                var v31 int32 = v29
                                _ = v30
                                _ = v31
                            }
                            _ = v28
                            _ = v29
                        } else {
                            v26 += (v26 >> ((uint(p0)) & 63))
                            var v32 int16 = 42
                            _ = v32
                        }
                    }
                    fmt.Println(fi24)
                    _ = v26
                }
                _ = fi24
            }
        case "\n\t":
            var v33 string = (("a" + "\n\t") + "\n\t")
            _ = v33
        default:
            switch ("foo bar" + ("a" + "a")) {
            case "hello":
                p0 += 65535
            case "foo bar":
                var v34 bool = (-1 <= -1)
                _ = v34
            default:
                p0 -= 65535
            }
        }
        var v35 uint = 255
        _ = v35
    }
    _ = p0
    _ = v23
    _ = v6
    return 42
}

func main() {
    v1 := []string{("a" + "hello")}
    var v2 uint64 = 2
    var v3 int16 = 1
    var v4 float64 = float64(float32(((float64(3.14159) * float64(0.1)) * float64(3.14159))))
    v36 := GenS2{F0: 2147483648, F1: 143, F2: 255}
    switch 4294967295 {
    case 32767:
        v37 := GenS2{F0: 2, F1: 255, F2: 689}
        _ = v37
    case -1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -2147483648:
        for fi38 := 0; fi38 < 12; fi38++ {
            v39 := GenS3{F0: "a", F1: -1, F2: 21}
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            v4 = v4
            _ = fi38
            _ = v39
        }
    }
    switch "\n\t" {
    case "":
        fmt.Println(v3, v4, v3)
    case "hello":
        var v40 uint64 = (^(v2))
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        if false {
            var v42 uint8 = 255
            v3 -= v3
            _ = v42
        }
        _ = fi41
    }
    v3 = (v3 / ((-101) | 1))
    for fi43 := 0; fi43 < 9; fi43++ {
        if true {
            v2 = 834
            for fi44 := 0; fi44 < 6; fi44++ {
                v45 := GenS4{F0: -767, F1: false, F2: (1 == 536)}
                _ = fi44
                _ = v45
            }
        } else {
            fmt.Println(v3, v3)
        }
        _ = fi43
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        v4 = v4
        switch "hello" {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            switch ("" + "") {
            case "\n\t":
                var v47 float32 = (0.1 * 1.5)
                _ = v47
            case "":
                if true {
                    v3 -= v3
                } else {
                    for fi48 := 0; fi48 < 3; fi48++ {
                        v4 -= ((v4 - (float64((-2.25 / 1e10)) / float64(float32(((1.5 - v4) * (((v4 - 1.5) / 1e10) / float64(1.5))))))) - v4)
                        if true {
                            switch "hello" {
                            case "":
                                v49 := GenS5{F0: ((v3 % ((v3) | 1)) <= v3), F1: v3}
                                _ = v49
                            case "foo bar":
                                v50 := GenS1{F0: 956}
                                _ = v50
                            case "\n\t":
                                v51 := GenS4{F0: 317, F1: true, F2: (true && ("hello" > ("\n\t" + ("a" + ""))))}
                                _ = v51
                            default:
                                var v52 int8 = 0
                                _ = v52
                            }
                            if (("hello" > "a") && (2 >= 42)) {
                                var v53 int = fi48
                                var v54 int = (-(v53))
                                _ = v53
                                _ = v54
                            }
                            var v55 string = ("" + "hello")
                            _ = v55
                        } else {
                            v2 -= (v2 | 42)
                            v56 := GenS3{F0: ("" + "hello"), F1: 100, F2: 0}
                            v4 = v4
                            _ = v56
                        }
                        switch "a" {
                        case "\n\t":
                            var v57 bool = true
                            _ = v57
                        default:
                            var v58 int32 = 965
                            _ = v58
                        }
                        var v59 int32 = 180
                        switch (fi46 - fi48) {
                        case -2147483648:
                            if true {
                                v60 := GenS2{F0: 42, F1: 46, F2: genFn0(32767)}
                                if true {
                                    v61 := GenS6{F0: 37}
                                    _ = v61
                                } else {
                                    _ = len(v1)
                                    v59 = (v59 % ((v59) | 1))
                                }
                                _ = v60
                            }
                        default:
                            switch "" {
                            case "hello":
                                v2 -= v2
                            case "foo bar":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                _ = len(v1)
                            }
                        }
                        if (true && false) {
                            var v62 int16 = v3
                            var v63 uint16 = 253
                            _ = v62
                            _ = v63
                        } else {
                            v59 = (v59 & -2147483648)
                        }
                        v64 := GenS7{F0: fi46}
                        _ = fi48
                        _ = v59
                        _ = v64
                    }
                    v2 = (v2 << ((2) & 63))
                    for fi65 := 0; fi65 < 10; fi65++ {
                        var v66 uint8 = 93
                        _ = fi65
                        _ = v66
                    }
                }
            case "foo bar":
                v3 += v3
            }
        case "\n\t":
            if true {
                v2 = v2
            } else {
                fmt.Println(v4)
            }
        }
        _ = fi46
    }
    v4 += v4
    if false {
        v67 := GenS8{F0: "\n\t"}
        var v68 uint32 = 0
        fmt.Println(v4, v2, v3)
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 9; fi69++ {
        fmt.Println(fi69)
        _ = fi69
    }
    switch "foo bar" {
    case "hello":
        var v70 uint16 = 382
        _ = v70
    case "a":
        var v71 uint64 = 0
        _ = v71
    case "\n\t":
        _ = len(v1)
    default:
        for fi72 := 0; fi72 < 7; fi72++ {
            if false {
                v73 := GenS6{F0: 719}
                _ = v73
            } else {
                v74 := GenS9{F0: 261, F1: 680, F2: 2147483647}
                var v75 uint8 = 151
                switch "a" {
                case "a":
                    var v76 float64 = (v4 * (v4 + v4))
                    _ = v76
                case "":
                    switch ("hello" + ("" + "foo bar")) {
                    case "hello":
                        v77 := GenS3{F0: "\n\t", F1: 25, F2: (^(v75))}
                        _ = v77
                    default:
                        v78 := GenS9{F0: 0, F1: 4294967295, F2: 756}
                        _ = v78
                    }
                case "hello":
                    for fi79 := 0; fi79 < 11; fi79++ {
                        v2 -= (v2 * v2)
                        var v80 uint32 = 2
                        _ = fi79
                        _ = v80
                    }
                default:
                    for fi81 := 0; fi81 < 14; fi81++ {
                        v82 := GenS3{F0: "foo bar", F1: 1, F2: 212}
                        _ = fi81
                        _ = v82
                    }
                }
                _ = v74
                _ = v75
            }
            fmt.Println(fi72)
            switch ("foo bar" + ("" + (("foo bar" + "\n\t") + "hello"))) {
            case "":
                var v83 float64 = 3.14159
                _ = v83
            }
            _ = fi72
        }
    }
    switch "foo bar" {
    case "hello":
        switch -1 {
        case -9223372036854775808:
            v3 += (v3 ^ 1000)
        case 127:
            v4 = v4
        case 4294967295:
            var v84 uint32 = 654
            _ = v84
        }
    case "\n\t":
        if (127 != 255) {
            switch "hello" {
            case "a":
                var v85 int = -825
                _ = v85
            case "foo bar":
                var v86 int = 128
                _ = v86
            case "":
                fmt.Println(v4, v2)
            default:
                v87 := GenS10{F0: (^(v3)), F1: "hello"}
                _ = v87
            }
        }
    case "a":
        v3 += v3
    }
    v3 = (v3 << ((564) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v4)
    for fi88 := 0; fi88 < 8; fi88++ {
        v89 := GenS8{F0: (("\n\t" + ("\n\t" + "a")) + "")}
        v4 += 1.5
        _ = fi88
        _ = v89
    }
    if true {
        if false {
            if len(v1) > 0 {
                v1[0] = "hello"
            }
        }
        if (true && (-445 != -1)) {
            for fi90 := 0; fi90 < 2; fi90++ {
                v91 := GenS8{F0: ""}
                v92 := GenS11{F0: 255, F1: 2147483648}
                v3 = 32767
                _ = fi90
                _ = v91
                _ = v92
            }
        } else {
            var v93 float64 = v4
            _ = len(v1)
            v94 := GenS12{F0: int8(v2)}
            _ = v93
            _ = v94
        }
        v4 = float64(1.5)
    }
    var v95 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v95
}
