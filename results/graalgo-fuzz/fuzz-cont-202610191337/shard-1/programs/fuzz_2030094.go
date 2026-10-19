package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32; F1 bool; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int16; F1 int64; F2 int32 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 uint8; F2 uint16 }
type GenS10 struct { F0 string; F1 int8; F2 uint32 }
type GenS11 struct { F0 bool; F1 uint16; F2 int64 }
type GenS12 struct { F0 string; F1 uint8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint64 }
func genFn0(p0 bool) int8 {
    if true {
        v5 := GenS1{F0: 911, F1: -46, F2: 42}
        _ = v5
    } else {
        var v6 uint8 = 42
        _ = v6
    }
    p0 = p0
    v7 := []int{-317}
    _ = p0
    _ = v7
    return 127
}
func genFn1(p0 bool, p1 uint32, p2 string) int8 {
    switch "foo bar" {
    case "hello":
        p2 = p2
    case "foo bar":
        if (650 > 2147483648) {
            v8 := []int{4294967295}
            v9 := GenS2{F0: (p1 | (p1 - 2147483648)), F1: p0, F2: "foo bar"}
            switch "a" {
            case "hello":
                var v10 string = p2
                _ = v10
            case "\n\t":
                var v11 int32 = 2147483647
                _ = v11
            case "":
                if (!(!(!false))) {
                    v12 := GenS3{F0: (float64(0.0) >= ((float64((1e10 / 3.14159)) * 0.1) + -2.25))}
                    for fi13 := 0; fi13 < 4; fi13++ {
                        var v14 int = fi13
                        for fi15 := 0; fi15 < 19; fi15++ {
                            fmt.Println(fi15, fi15, p0)
                            var v16 int8 = 1
                            _ = fi15
                            _ = v16
                        }
                        v17 := GenS2{F0: (p1 >> ((65535) & 63)), F1: p0, F2: "a"}
                        _ = fi13
                        _ = v14
                        _ = v17
                    }
                    p0 = (!(p0 && (0 <= 127)))
                    _ = v12
                } else {
                    var v18 int16 = 1
                    if p0 {
                        var v19 float32 = 1e10
                        for fi20 := 0; fi20 < 11; fi20++ {
                            var v21 uint = 255
                            p1 = 336
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        var v22 int16 = v18
                        for fi23 := 0; fi23 < 1; fi23++ {
                            var v24 uint64 = 67
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    }
                    _ = v18
                }
            }
            _ = v8
            _ = v9
        } else {
            p0 = (42 <= p1)
        }
    }
    var v25 string = "hello"
    var v26 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return v26
}

func main() {
    var v1 uint8 = 127
    var v2 int8 = -7
    v3 := make(map[int]bool)
    var v4 uint = 65535
    v1 += v1
    if false {
        v1 -= (v1 >> ((v4) & 63))
    } else {
        v2 = v2
    }
    if true {
        v1 += (v1 | 2)
        for fi27 := 0; fi27 < 4; fi27++ {
            if (((-602 <= 42) || false) || false) {
                for fi28 := 0; fi28 < 5; fi28++ {
                    v29 := GenS1{F0: 42, F1: 60, F2: 707}
                    v2 = -30
                    var v30 int32 = -2147483648
                    var v31 uint8 = v1
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                var v32 uint8 = (v1 / ((v1) | 1))
                if false {
                    switch (fi27 - fi27) {
                    case 60:
                        if ((0 == 4294967295) && false) {
                            v33 := GenS4{F0: 2, F1: 42}
                            if true {
                                for fi34 := 0; fi34 < 14; fi34++ {
                                    var v35 int64 = 9223372036854775807
                                    _ = fi34
                                    _ = v35
                                }
                                switch "foo bar" {
                                case "hello":
                                    v1 += (^(v32))
                                case "":
                                    v32 = v1
                                case "foo bar":
                                    v32 = v32
                                default:
                                    if true {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v36 := GenS5{F0: 42, F1: 14, F2: -1}
                                            _ = v36
                                        case "hello":
                                            for fi37 := 0; fi37 < 8; fi37++ {
                                                _ = len(v3)
                                                v38 := GenS6{F0: 4294967295, F1: fi37}
                                                _ = fi37
                                                _ = v38
                                            }
                                        case "a":
                                            v39 := GenS2{F0: 1, F1: true, F2: ("a" + "a")}
                                            _ = v39
                                        }
                                        if false {
                                            for fi40 := 0; fi40 < 8; fi40++ {
                                                if (1e10 >= 1.5) {
                                                    var v41 int = int(v1)
                                                    v42 := []int{(v41 * (-(fi40))), fi27, v41, fi27}
                                                    v43 := GenS2{F0: 65535, F1: (false || ("a" <= "a")), F2: ""}
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    v44 := GenS4{F0: 42, F1: 42}
                                                    v2 = (v2 >> (((v4 & 1)) & 63))
                                                    _ = v44
                                                }
                                                _ = fi40
                                            }
                                            fmt.Println(fi27, v2, v1)
                                            switch "a" {
                                            case "a":
                                                var v45 float32 = 3.14159
                                                _ = v45
                                            case "\n\t":
                                                v3[(fi27 + -1)] = false
                                            default:
                                                v3[fi27] = true
                                            }
                                        }
                                    } else {
                                        v4 += v4
                                    }
                                }
                                var v46 string = "foo bar"
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 17; fi47++ {
                                v1 = (^(v32))
                                if (!false) {
                                    var v48 float64 = float64(0.0)
                                    _ = v48
                                }
                                fmt.Println(v32, fi27)
                                var v49 int8 = int8(fi47)
                                v2 -= 102
                                v2 += (v2 / (((v49 | (v49 / ((0) | 1)))) | 1))
                                v4 = (^(v4))
                                switch "a" {
                                case "foo bar":
                                    switch ("hello" + ((("hello" + "foo bar") + ("hello" + "a")) + "a")) {
                                    case "a":
                                        v2 -= (v49 << ((v4) & 63))
                                    }
                                case "hello":
                                    v1 = v1
                                }
                                _ = fi47
                                _ = v49
                            }
                            _ = v33
                        }
                    case 9223372036854775807:
                        var v50 bool = false
                        _ = v50
                    default:
                        _ = len(v3)
                    }
                    for fi51 := 0; fi51 < 4; fi51++ {
                        var v52 uint64 = 699
                        _ = fi51
                        _ = v52
                    }
                }
                _ = v32
            }
            _ = fi27
        }
        switch 9223372036854775807 {
        case -9223372036854775808:
            v1 = (^(v1))
        case -727:
            v53 := GenS7{F0: 1}
            _ = v53
        case -2147483648:
            v54 := GenS8{F0: 2147483647}
            _ = v54
        }
    } else {
        if ((((-1 <= genFn0((1 == 42))) || (true || true)) && true) || (true || true)) {
            v55 := v3[9223372036854775807]
            _ = v55
        } else {
            _ = len(v3)
            if false {
                if false {
                    v4 = (v4 >> (((v4 / ((v4) | 1))) & 63))
                    for fi56 := 0; fi56 < 2; fi56++ {
                        v57 := GenS5{F0: 444, F1: 42, F2: 1}
                        v4 = (v4 + 0)
                        v4 = (^(v4))
                        _ = fi56
                        _ = v57
                    }
                }
            } else {
                var v58 uint8 = 255
                _ = v58
            }
        }
        var v59 int16 = -639
        _ = v59
    }
    for fi60 := 0; fi60 < 2; fi60++ {
        var v61 int32 = 467
        var v62 int32 = v61
        v63 := GenS4{F0: 1, F1: 32767}
        _ = fi60
        _ = v61
        _ = v62
        _ = v63
    }
    if (false && true) {
        switch ((((("hello" + (("a" + "\n\t") + ("\n\t" + "foo bar"))) + "hello") + (("a" + "") + "\n\t")) + "hello") + "") {
        case "\n\t":
            switch "foo bar" {
            case "foo bar":
                if true {
                    v64 := []int{4294967295}
                    _ = v64
                } else {
                    v2 = (v2 | 42)
                    for fi65 := 0; fi65 < 9; fi65++ {
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            switch "" {
                            case "a":
                                fmt.Println(v1, v4)
                            case "hello":
                                if false {
                                    delete(v3, (fi65 * fi65))
                                    v1 = v1
                                    var v66 int = fi65
                                    _ = v66
                                }
                            default:
                                v1 = v1
                            }
                        default:
                            var v67 uint64 = 93
                            _ = v67
                        }
                        var v68 uint16 = 991
                        fmt.Println(v1, v1)
                        _ = fi65
                        _ = v68
                    }
                    for fi69 := 0; fi69 < 2; fi69++ {
                        v2 = (v2 * (^(v2)))
                        _ = fi69
                    }
                }
            default:
                for fi70 := 0; fi70 < 16; fi70++ {
                    v4 = v4
                    v71 := []int{int(v1), 127, fi70}
                    if (true && (42 != 909)) {
                        fmt.Println(v2, v4)
                        switch "a" {
                        case "hello":
                            if (!false) {
                                fmt.Println(v1, v4, v4)
                                switch (fi70 | fi70) {
                                case 1:
                                    switch -9223372036854775808 {
                                    case 128:
                                        if ((v4 >> ((724) & 63)) < (v4 * v4)) {
                                            fmt.Println(fi70, v1, v1)
                                        }
                                    }
                                default:
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v4, v4, v2)
                                    case "hello":
                                        _ = len(v71)
                                    default:
                                        for fi72 := 0; fi72 < 18; fi72++ {
                                            v2 += genFn1(true, 748, "")
                                            v73 := GenS9{F0: 128, F1: 128, F2: 32767}
                                            if false {
                                                v4 -= 936
                                            }
                                            _ = fi72
                                            _ = v73
                                        }
                                    }
                                }
                            } else {
                                var v74 float64 = 1.5
                                if ((false && (!true)) && ("\n\t" == ("a" + ("" + "")))) {
                                    if (false && (!true)) {
                                        var v75 bool = false
                                        var v76 int = fi70
                                        _ = v75
                                        _ = v76
                                    }
                                    v1 -= v1
                                    switch (fi70 | (fi70 * 1)) {
                                    case 1:
                                        var v77 int16 = 982
                                        _ = v77
                                    case 127:
                                        fmt.Println(v2, fi70, v74)
                                    }
                                }
                                _ = v74
                            }
                        default:
                            var v78 int8 = v2
                            _ = v78
                        }
                    }
                    _ = fi70
                    _ = v71
                }
            }
        case "a":
            var v79 int64 = -628
            _ = v79
        }
        fmt.Println(v1)
        switch "a" {
        case "foo bar":
            for fi80 := 0; fi80 < 5; fi80++ {
                var v81 float32 = (0.1 + 1.5)
                var v82 uint = v4
                if true {
                    if true {
                        var v83 int16 = -1
                        _ = v83
                    } else {
                        for fi84 := 0; fi84 < 16; fi84++ {
                            v4 -= (v82 % (((v4 - (v4 + (v4 ^ 1)))) | 1))
                            _ = fi84
                        }
                    }
                }
                _ = fi80
                _ = v81
                _ = v82
            }
        case "a":
            v85 := GenS10{F0: "\n\t", F1: 42, F2: 0}
            _ = v85
        case "\n\t":
            switch 881 {
            case -9223372036854775808:
                fmt.Println(v2)
            default:
                _ = len(v3)
            }
        }
    }
    if (65535 <= 32767) {
        var v86 uint16 = 255
        v87 := GenS11{F0: true, F1: 838, F2: -465}
        _ = v86
        _ = v87
    }
    v2 = v2
    v88 := GenS12{F0: "a", F1: (v1 - (v1 / ((123) | 1)))}
    var v89 float32 = -2.25
    var v90 string = "a"
    v91 := GenS6{F0: 2147483647, F1: 127}
    for fi92 := 0; fi92 < 3; fi92++ {
        if (1 < -9223372036854775808) {
            var v93 float64 = 0.0
            if ((-724 >= 0) && true) {
                var v94 uint8 = 128
                _ = v94
            } else {
                switch (fi92 - (fi92 % ((1) | 1))) {
                case 128:
                    fmt.Println(v93, v4, v89)
                case 895:
                    for fi95 := 0; fi95 < 6; fi95++ {
                        var v96 int16 = -397
                        _ = fi95
                        _ = v96
                    }
                default:
                    fmt.Println(fi92, v90, v89)
                }
                var v97 uint = v4
                v98 := GenS13{F0: "foo bar"}
                _ = v97
                _ = v98
            }
            _ = v93
        }
        v99 := GenS14{F0: 378}
        var v100 uint64 = 0
        _ = fi92
        _ = v100
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
