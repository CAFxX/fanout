package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 int16; F2 int64 }
type GenS3 struct { F0 int16; F1 string; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 int8; F1 uint; F2 uint32 }
func genFn0(p0 bool, p1 uint16, p2 int16) int64 {
    if p0 {
        switch "\n\t" {
        case "a":
            v5 := []int{42, 1}
            _ = v5
        default:
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := []int{fi6}
                v8 := GenS1{F0: fi6, F1: ("foo bar" + "\n\t")}
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint32 = 2147483648
            v11 := []int{fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    switch 127 {
    case 127:
        var v12 float32 = ((1e10 - float32(1e10)) + 0.0)
        _ = v12
    case -2147483648:
        var v13 uint16 = 255
        _ = v13
    case 2147483647:
        var v14 uint8 = 127
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: (fi15 ^ (fi15 + fi15)), F1: ("\n\t" + "foo bar")}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 265
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 uint16 = 1
    v4 := make(map[int]bool)
    switch "foo bar" {
    case "a":
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := GenS1{F0: (^(fi17)), F1: "hello"}
            switch "hello" {
            case "foo bar":
                if false {
                    v3 = v3
                    v19 := GenS2{F0: 1, F1: -1, F2: genFn0((((false || (!true)) || (!(255 != 128))) && true), 2, 1000)}
                    _ = v19
                }
            case "\n\t":
                v3 = (v3 % (((v3 + (v3 % ((v3) | 1)))) | 1))
            case "a":
                fmt.Println(v3, fi17)
            }
            _ = fi17
            _ = v18
        }
    case "foo bar":
        if true {
            v3 -= (v3 / (((^(v3))) | 1))
            fmt.Println(v3)
            var v20 int8 = 100
            _ = v20
        } else {
            v2[42] = "\n\t"
            var v21 int32 = 600
            fmt.Println(v3, v3)
            _ = v21
        }
    case "hello":
        var v22 int = 4294967295
        _ = v22
    default:
        fmt.Println(v3)
    }
    if (9223372036854775807 > 0) {
        v23 := []int{0, 42, -1}
        _ = len(v1)
        switch 740 {
        case 128:
            v24 := GenS2{F0: 526, F1: 476, F2: -9223372036854775808}
            _ = v24
        case 127:
            for fi25 := 0; fi25 < 20; fi25++ {
                if false {
                    fmt.Println(v3, fi25)
                    v26 := v1[505]
                    _ = v26
                } else {
                    if true {
                        var v27 string = ("hello" + "foo bar")
                        v27 = ""
                        _ = v27
                    } else {
                        var v28 float32 = -2.25
                        _ = v28
                    }
                }
                v29 := v4[fi25]
                _ = fi25
                _ = v29
            }
        case -678:
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 int64 = 262
                switch ("\n\t" + ("a" + "")) {
                case "\n\t":
                    v32 := GenS2{F0: 65535, F1: 0, F2: (v31 ^ v31)}
                    _ = v32
                }
                v33 := GenS2{F0: 304, F1: 830, F2: v31}
                _ = fi30
                _ = v31
                _ = v33
            }
        }
        _ = v23
    }
    if (((32767 == 2147483647) && false) && (0 == -128)) {
        if true {
            if true {
                v34 := []int{1, 2147483647, -58, -9223372036854775808}
                v1[660] = (true && true)
                v35 := GenS1{F0: -597, F1: "\n\t"}
                _ = v34
                _ = v35
            }
            if false {
                v3 = v3
                v36 := GenS3{F0: 177, F1: ("\n\t" + "\n\t"), F2: 2}
                fmt.Println(v3)
                _ = v36
            } else {
                var v37 int64 = 0
                _ = v37
            }
        }
    }
    if true {
        if ((860 == -9223372036854775808) && (true && true)) {
            v3 += 1
            var v38 int64 = -1
            for fi39 := 0; fi39 < 14; fi39++ {
                switch "a" {
                case "":
                    for fi40 := 0; fi40 < 1; fi40++ {
                        var v41 uint64 = 2
                        var v42 uint32 = 2
                        var v43 int16 = -1
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                }
                _ = len(v4)
                if false {
                    v44 := GenS2{F0: 1, F1: -1, F2: 1}
                    switch ((("" + "") + "\n\t") + "hello") {
                    case "foo bar":
                        if true {
                            v45 := GenS1{F0: fi39, F1: "a"}
                            _ = v45
                        } else {
                            for fi46 := 0; fi46 < 19; fi46++ {
                                v3 -= v3
                                var v47 float32 = -2.25
                                v47 = v47
                                v48 := []int{(fi46 * -1)}
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            v49 := GenS4{F0: 829, F1: 65535}
                            v3 = (v3 | v3)
                            _ = v49
                        }
                    default:
                        if false {
                            fmt.Println(v38, v38)
                            if (59 != -1) {
                                var v50 bool = true
                                var v51 uint16 = 32767
                                switch ((("a" + "hello") + "foo bar") + ("a" + "")) {
                                case "foo bar":
                                    if (242 != 32767) {
                                        v50 = (-1 <= -1)
                                        for fi52 := 0; fi52 < 19; fi52++ {
                                            var v53 int8 = 100
                                            if false {
                                                v38 = (-(v38))
                                                var v54 uint = 406
                                                fmt.Println(v3, v51, v3)
                                                _ = v54
                                            }
                                            v55 := GenS5{F0: 296, F1: v3}
                                            var v56 int = 893
                                            v53 += v53
                                            fmt.Println(fi39, fi39, v51)
                                            _ = fi52
                                            _ = v53
                                            _ = v55
                                            _ = v56
                                        }
                                        v50 = v50
                                    }
                                }
                                _ = v50
                                _ = v51
                            }
                            var v57 uint64 = 352
                            _ = v57
                        }
                    }
                    var v58 string = "a"
                    _ = v44
                    _ = v58
                } else {
                    var v59 uint = 2
                    _ = v59
                }
                _ = fi39
            }
            _ = v38
        } else {
            for fi60 := 0; fi60 < 7; fi60++ {
                var v61 bool = false
                for fi62 := 0; fi62 < 11; fi62++ {
                    v61 = true
                    fmt.Println(fi60, fi62, fi62)
                    v3 = v3
                    if false {
                        var v63 int16 = 0
                        _ = v63
                    }
                    _ = fi62
                }
                _ = fi60
                _ = v61
            }
            v64 := GenS5{F0: 2147483648, F1: (v3 * (v3 << ((65535) & 63)))}
            _ = v64
        }
        var v65 int8 = 127
        if true {
            fmt.Println(v3)
            fmt.Println(v65)
            for fi66 := 0; fi66 < 6; fi66++ {
                v67 := []int{(fi66 << ((335) & 63))}
                _ = fi66
                _ = v67
            }
        }
        _ = v65
    } else {
        if (((((-2.25 / ((-2.25 - (-2.25 / (1.5 * (-2.25 - 1.5)))) + 1e10)) / ((1.5 + ((0.1 + float32((0.1 * 0.0))) * 1.5)) * (float32(1e10) * 0.1))) / ((0.1 + (0.0 * 1.5)) - (1.5 - -2.25))) - 0.1) == float32(((-2.25 / ((float64(float32(3.14159)) / 1e10) - 1.5)) - 1e10))) {
            v68 := GenS3{F0: 0, F1: "a", F2: 2}
            _ = len(v1)
            _ = v68
        }
        fmt.Println(v3)
        if true {
            switch ("a" + "") {
            case "hello":
                v69 := []int{128}
                _ = v69
            case "\n\t":
                v3 = (v3 << ((2) & 63))
            case "a":
                for fi70 := 0; fi70 < 1; fi70++ {
                    for fi71 := 0; fi71 < 9; fi71++ {
                        v72 := v2[654]
                        _ = fi71
                        _ = v72
                    }
                    switch "\n\t" {
                    case "\n\t":
                        v73 := []int{9223372036854775807, -2147483648, fi70, (fi70 * fi70)}
                        _ = v73
                    case "hello":
                        v74 := []int{fi70}
                        _ = v74
                    case "a":
                        if ((42 == 4294967295) || true) {
                            v3 = v3
                            switch 0 {
                            case 127:
                                v3 -= v3
                            }
                            if (false && true) {
                                v75 := []int{fi70, (fi70 - (^(fi70)))}
                                _ = v75
                            }
                        }
                    default:
                        if (false || (!(!true))) {
                            for fi76 := 0; fi76 < 11; fi76++ {
                                v3 -= 609
                                fmt.Println(fi70, v3)
                                v77 := GenS6{F0: 780}
                                _ = fi76
                                _ = v77
                            }
                            v78 := GenS6{F0: 9223372036854775807}
                            _ = v78
                        }
                    }
                    _ = fi70
                }
            }
            delete(v4, 9223372036854775807)
            v79 := []int{}
            _ = v79
        }
    }
    if (0 > 42) {
        var v80 string = "hello"
        if (2147483647 > 128) {
            switch 823 {
            case -1:
                v3 -= v3
            case -57:
                if true {
                    switch 42 {
                    case -247:
                        v81 := GenS6{F0: 9223372036854775807}
                        _ = v81
                    }
                }
            case -894:
                v80 = "a"
            }
            v82 := GenS7{F0: 0, F1: (genFn0((true && ("hello" >= (v80 + "\n\t"))), (v3 >> ((42) & 63)), -32768) == -164)}
            _ = v82
        }
        _ = v80
    }
    v83 := []int{-1, 32767, -425, 42}
    var v84 float32 = (0.0 / (-2.25 * float32(-2.25)))
    v85 := GenS4{F0: 697, F1: 65535}
    v3 = 42
    v86 := GenS8{F0: 125, F1: 234, F2: 1}
    v87 := GenS1{F0: -254, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v87
}
