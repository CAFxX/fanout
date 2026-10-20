package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 string; F1 int8; F2 string }
type GenS6 struct { F0 int32; F1 int8; F2 uint16 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 int64; F1 uint16 }
type GenS11 struct { F0 int16; F1 string; F2 uint8 }
type GenS12 struct { F0 uint; F1 uint; F2 bool }
type GenS13 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5)
        var v6 int = fi5
        _ = fi5
        _ = v6
    }
    var v7 int8 = -1
    if false {
        fmt.Println(v7)
    }
    _ = v7
    return 127
}
func genFn1(p0 int16) uint8 {
    fmt.Println(p0)
    if (!(32767 < 4294967295)) {
        p0 = 0
        if true {
            if (false && true) {
                switch 32768 {
                case 4294967295:
                    fmt.Println(p0)
                case 128:
                    if true {
                        fmt.Println(p0)
                        v8 := GenS1{F0: (("hello" + "foo bar") + "")}
                        _ = v8
                    }
                case 32768:
                    p0 -= -32768
                }
            } else {
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 int16 = 0
                    _ = fi9
                    _ = v10
                }
                var v11 uint32 = 261
                fmt.Println(v11, v11)
                _ = v11
            }
            var v12 int64 = 1
            var v13 int = -2147483648
            _ = v12
            _ = v13
        }
        var v14 string = "a"
        _ = v14
    }
    p0 += (p0 | p0)
    _ = p0
    return 160
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2147483647, 42}
    var v3 int16 = 799
    var v4 float64 = 3.14159
    var v15 uint = 4294967295
    fmt.Println(v4)
    if (733 > 2147483648) {
        v15 = v15
        v15 -= 0
        var v16 uint64 = 1
        _ = v16
    }
    v17 := GenS2{F0: (v3 % ((v3) | 1))}
    if true {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS1{F0: "\n\t"}
            if true {
                fmt.Println(v4, fi18, v3)
                var v20 float32 = 1.5
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 float64 = -2.25
                    if ((false && (!(4294967295 >= 509))) && ((!(1000 == -32768)) && true)) {
                        fmt.Println(v3)
                        v22 -= -2.25
                    } else {
                        fmt.Println(fi18, v4, v4)
                        if (!(-42 <= -56)) {
                            delete(v1, 4294967295)
                            v23 := v1[2147483647]
                            v15 += (v15 % ((v15) | 1))
                            _ = v23
                        } else {
                            var v24 int = fi21
                            _ = v24
                        }
                        if true {
                            switch "a" {
                            case "hello":
                                var v25 uint = (v15 * 0)
                                _ = v25
                            case "a":
                                v26 := GenS3{F0: fi18, F1: 0, F2: (^(fi21))}
                                _ = v26
                            case "\n\t":
                                var v27 uint16 = 333
                                _ = v27
                            }
                            v22 -= v4
                            var v28 float32 = v20
                            _ = v28
                        } else {
                            v22 -= float64(float32(0.1))
                            if ("foo bar" <= "foo bar") {
                                fmt.Println(v3, v20, v3)
                                var v29 uint32 = 2147483647
                                _ = v29
                            } else {
                                var v30 uint = v15
                                if (65535 <= 33) {
                                    var v31 uint64 = 2
                                    var v32 int8 = -128
                                    v30 -= 41
                                    _ = v31
                                    _ = v32
                                }
                                _ = v30
                            }
                        }
                    }
                    if ((false && ((!false) && false)) || false) {
                        if true {
                            v15 = (v15 + (v15 * 18446744073709551615))
                            var v33 int32 = 42
                            v34 := GenS2{F0: v3}
                            _ = v33
                            _ = v34
                        }
                    } else {
                        fmt.Println(v4, v22)
                        v35 := GenS4{F0: 42, F1: (fi21 - 9223372036854775807)}
                        var v36 uint32 = 0
                        var v37 int64 = 0
                        if true {
                            v36 -= 65535
                            delete(v1, v36)
                        } else {
                            if false {
                                v36 -= (v36 - v36)
                            } else {
                                if ((((v20 - (((1.5 * v20) / v20) - v20)) / 1.0) - v20) == (float32((float64(v20) + float64(v20))) + v20)) {
                                    switch "foo bar" {
                                    case "hello":
                                        v1[462] = ("hello" + ("\n\t" + "hello"))
                                    case "\n\t":
                                        var v38 float32 = v20
                                        fmt.Println(v38)
                                        if false {
                                            switch ("" + ("foo bar" + "foo bar")) {
                                            case "a":
                                                var v39 int8 = -28
                                                _ = v39
                                            case "":
                                                v40 := GenS3{F0: (-(fi18)), F1: (^(v37)), F2: 42}
                                                _ = v40
                                            case "\n\t":
                                                v22 = (v22 * (v4 / ((3.14159 - (v22 / -2.25)) + ((v22 - 0.1) + (float64(1.5) / float64(v20))))))
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v3 -= (v3 % ((394) | 1))
                                            case "":
                                                fmt.Println(fi18, v37)
                                            }
                                            v41 := GenS5{F0: ("a" + ""), F1: 125, F2: "foo bar"}
                                            if (v15 == v15) {
                                                v36 -= (v36 >> ((42) & 63))
                                                if false {
                                                    v36 += (v36 ^ 2)
                                                    switch (^(fi21)) {
                                                    case 845:
                                                        if true {
                                                            switch fi21 {
                                                            case 2147483647:
                                                                fmt.Println(v3, fi18, v22)
                                                            case -506:
                                                                fmt.Println(v37, v36, v38)
                                                            case 32767:
                                                                var v42 bool = true
                                                                _ = v42
                                                            }
                                                            fmt.Println(fi21, v38, v38)
                                                        }
                                                        v20 -= -2.25
                                                    case 2147483647:
                                                        var v43 uint32 = 2
                                                        _ = v43
                                                    case -738:
                                                        var v44 uint64 = 490
                                                        _ = v44
                                                    }
                                                    if true {
                                                        var v45 int = (^(fi18))
                                                        _ = v45
                                                    }
                                                    v2 = append(v2, v36)
                                                    var v46 uint32 = (v36 & v36)
                                                    _ = v46
                                                } else {
                                                    switch "\n\t" {
                                                    case "":
                                                        fmt.Println(v4, v20, fi18)
                                                    }
                                                    v20 += v20
                                                    fmt.Println(v22, fi18)
                                                }
                                                var v47 int8 = 93
                                                _ = v47
                                            } else {
                                                v48 := GenS2{F0: v3}
                                                if true {
                                                    if ((^(v3)) < v3) {
                                                        fmt.Println(v20)
                                                        fmt.Println(v36)
                                                        var v49 uint64 = 245
                                                        _ = v49
                                                    }
                                                } else {
                                                    v50 := GenS6{F0: -670, F1: 70, F2: 682}
                                                    fmt.Println(v37, fi18, fi21)
                                                    fmt.Println(v36)
                                                    _ = v50
                                                }
                                                _ = v48
                                            }
                                            fmt.Println(v38)
                                            _ = v41
                                        }
                                        _ = v38
                                    case "foo bar":
                                        var v51 uint16 = 831
                                        _ = v51
                                    }
                                }
                                v1[v36] = "foo bar"
                            }
                        }
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    v15 += (v15 - 18446744073709551615)
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                v4 = (((((v4 * float64(1.5)) / v4) * (1.5 - v4)) - v4) - (v4 - v4))
            }
            _ = fi18
            _ = v19
        }
        for fi52 := 0; fi52 < 19; fi52++ {
            var v53 float32 = (3.14159 * 0.1)
            if false {
                v54 := GenS7{F0: genFn0(), F1: fi52}
                _ = v54
            } else {
                v15 -= (v15 & (v15 & v15))
                for fi55 := 0; fi55 < 14; fi55++ {
                    v56 := GenS8{F0: 861, F1: v15}
                    _ = fi55
                    _ = v56
                }
            }
            var v57 uint8 = 200
            _ = fi52
            _ = v53
            _ = v57
        }
    } else {
        if false {
            var v58 bool = (!(50 != 85))
            _ = v58
        } else {
            v59 := GenS9{F0: "foo bar", F1: 227}
            if (-704 == -2147483648) {
                switch "foo bar" {
                case "a":
                    v15 = v15
                case "\n\t":
                    v60 := GenS1{F0: "hello"}
                    _ = v60
                case "foo bar":
                    switch -2147483648 {
                    case 9223372036854775807:
                        fmt.Println(v3, v15, v15)
                    case 0:
                        if (-128 >= -116) {
                            var v61 int = 4294967295
                            v4 += (float64(1e10) / v4)
                            v2 = append(v2, 65535)
                            _ = v61
                        }
                    }
                }
                for fi62 := 0; fi62 < 6; fi62++ {
                    v63 := GenS10{F0: -666, F1: 65535}
                    if (!true) {
                        var v64 float32 = (0.1 + (3.14159 / 0.1))
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 6; fi65++ {
                        fmt.Println(fi62, v3, v3)
                        _ = fi65
                    }
                    _ = fi62
                    _ = v63
                }
            }
            var v66 uint32 = 65535
            _ = v59
            _ = v66
        }
        var v67 int8 = 47
        _ = v67
    }
    if (707 > 42) {
        fmt.Println(v15, v15)
    } else {
        if (617 >= 445) {
            for fi68 := 0; fi68 < 19; fi68++ {
                fmt.Println(v15)
                _ = fi68
            }
            v4 += v4
            if (true || ("" != "\n\t")) {
                v4 -= (1e10 * -2.25)
            } else {
                var v69 uint16 = 42
                switch 9223372036854775807 {
                case -1:
                    v70 := GenS11{F0: v3, F1: "\n\t", F2: 127}
                    _ = v70
                default:
                    for fi71 := 0; fi71 < 15; fi71++ {
                        fmt.Println(v3, v4)
                        switch "" {
                        case "foo bar":
                            if (0.0 != -2.25) {
                                switch "\n\t" {
                                case "a":
                                    v69 = (v69 ^ v69)
                                case "hello":
                                    for fi72 := 0; fi72 < 5; fi72++ {
                                        v73 := GenS12{F0: (^(v15)), F1: (^(v15)), F2: ("\n\t" == "hello")}
                                        _ = len(v1)
                                        _ = fi72
                                        _ = v73
                                    }
                                }
                                var v74 int32 = -744
                                _ = v74
                            } else {
                                v15 -= v15
                            }
                        default:
                            switch ("\n\t" + (("\n\t" + "foo bar") + "")) {
                            case "foo bar":
                                v4 += v4
                            }
                        }
                        _ = fi71
                    }
                }
                _ = v69
            }
        }
    }
    var v75 int32 = -639
    if false {
        switch "foo bar" {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            if (int(v75) != 953) {
                for fi76 := 0; fi76 < 12; fi76++ {
                    v77 := GenS13{F0: "hello", F1: 0, F2: 498}
                    v2 = append(v2, 1)
                    v75 += 42
                    _ = fi76
                    _ = v77
                }
            }
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v78 int = 9223372036854775807
        _ = v78
    }
    var v79 int16 = v3
    v3 -= 32767
    fmt.Println(v4, v79)
    switch "a" {
    case "\n\t":
        v80 := GenS7{F0: 42, F1: 0}
        _ = v80
    case "foo bar":
        switch ("\n\t" + "") {
        case "foo bar":
            var v81 uint8 = 1
            _ = v81
        case "hello":
            if (0 <= 18446744073709551615) {
                switch 1 {
                case 32767:
                    for fi82 := 0; fi82 < 17; fi82++ {
                        var v83 uint16 = 1
                        var v84 uint32 = 2147483648
                        _ = fi82
                        _ = v83
                        _ = v84
                    }
                }
            } else {
                switch "foo bar" {
                case "a":
                    if (("a" >= "foo bar") && true) {
                        var v85 uint = (v15 * (v15 << ((65535) & 63)))
                        _ = v85
                    }
                case "foo bar":
                    fmt.Println(v75)
                case "":
                    fmt.Println(v79, v15, v75)
                }
                v86 := GenS4{F0: 102, F1: 787}
                _ = v86
            }
        }
    case "":
        fmt.Println(v79, v75)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v75
    _ = v79
}
