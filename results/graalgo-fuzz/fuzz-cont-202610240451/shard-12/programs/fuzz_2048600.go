package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint32; F2 string }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint8; F1 string; F2 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint64; F1 int; F2 int16 }
func genFn0() uint64 {
    var v5 uint64 = 685
    if true {
        fmt.Println(v5)
        var v6 string = ("foo bar" + "a")
        v7 := []int{-9223372036854775808, 2147483647, -850, 116}
        _ = v6
        _ = v7
    } else {
        fmt.Println(v5)
        switch 128 {
        case 32768:
            v8 := []int{128, 128, -1, 90}
            _ = v8
        case 128:
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(fi9, fi9)
                _ = fi9
            }
        default:
            for fi10 := 0; fi10 < 4; fi10++ {
                v5 -= v5
                _ = fi10
            }
        }
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint8 = 154
        var v14 int = -1
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = v5
    return v5
}
func genFn1(p0 int16) int64 {
    p0 += (p0 * (p0 | p0))
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint8 = 42
        fmt.Println(fi15, fi15)
        _ = fi15
        _ = v16
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{87}
    var v3 uint = 767
    v4 := []string{"hello", "a", ("foo bar" + "a")}
    for fi17 := 0; fi17 < 10; fi17++ {
        if false {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float64 = 0.1
                if true {
                    delete(v1, (fi18 % (((fi17 / ((fi17) | 1))) | 1)))
                    switch "" {
                    case "a":
                        var v20 uint8 = 143
                        _ = v20
                    case "hello":
                        switch "" {
                        case "foo bar":
                            var v21 uint8 = 151
                            _ = v21
                        case "hello":
                            var v22 int8 = 1
                            _ = v22
                        case "\n\t":
                            if (!(float32((v19 / float64((-2.25 + 1e10)))) <= 1e10)) {
                                v23 := GenS1{F0: 255}
                                fmt.Println(v3, v19)
                                var v24 int16 = -576
                                _ = v23
                                _ = v24
                            }
                        default:
                            if true {
                                var v25 uint = (v3 + 956)
                                v1[fi18] = 42
                                _ = v25
                            }
                        }
                    }
                    if false {
                        var v26 int64 = -938
                        var v27 int32 = -163
                        _ = v26
                        _ = v27
                    } else {
                        v19 = (v19 - (((((v19 + v19) + (0.0 / 3.14159)) - v19) / (-2.25 - v19)) * ((float64(1e10) / 3.14159) * (1e10 + (0.1 - v19)))))
                        var v28 int32 = int32(v3)
                        v29 := GenS1{F0: v3}
                        _ = v28
                        _ = v29
                    }
                } else {
                    var v30 uint16 = 32767
                    _ = v30
                }
                if false {
                    v1[fi18] = 60
                    var v31 int8 = 0
                    _ = v31
                } else {
                    var v32 int32 = 1
                    v33 := GenS2{F0: "", F1: 2147483648, F2: "foo bar"}
                    v34 := GenS2{F0: "", F1: 2, F2: "foo bar"}
                    v19 += v19
                    if (2147483647 > 2147483648) {
                        v19 -= v19
                    } else {
                        var v35 int = fi17
                        _ = v35
                    }
                    if true {
                        fmt.Println(v3)
                        v3 -= (v3 << (((v3 << (((v3 ^ v3)) & 63))) & 63))
                        v36 := GenS1{F0: v3}
                        _ = v36
                    }
                    switch ("a" + ("\n\t" + "foo bar")) {
                    case "":
                        if (!(false || ((fi17 + (fi18 | (fi17 & fi18))) >= (fi18 | 2147483647)))) {
                            delete(v1, fi18)
                        } else {
                            v1[fi17] = 2147483648
                            v37 := GenS3{F0: 814, F1: genFn0()}
                            _ = v37
                        }
                    case "hello":
                        switch ("hello" + "") {
                        case "a":
                            if true {
                                v32 = v32
                            }
                        }
                        fmt.Println(v32, v32)
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    var v38 uint32 = 2
                    fmt.Println(v38, fi18)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v32
                    _ = v33
                    _ = v34
                    _ = v38
                }
                _ = fi18
                _ = v19
            }
            switch (fi17 | fi17) {
            case 32768:
                var v39 uint32 = 2147483647
                _ = v39
            default:
                fmt.Println(fi17)
            }
        }
        _ = fi17
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        for fi41 := 0; fi41 < 10; fi41++ {
            switch "" {
            case "":
                fmt.Println(fi40)
            default:
                v42 := GenS4{F0: (v3 - 2), F1: 255, F2: -409}
                switch "\n\t" {
                case "\n\t":
                    v43 := GenS5{F0: 42, F1: ("" + "foo bar"), F2: 1}
                    _ = v43
                default:
                    var v44 uint32 = 536
                    _ = v44
                }
                switch fi41 {
                case 2147483647:
                    if (-2147483648 <= 2147483647) {
                        v45 := GenS3{F0: (v3 ^ v3), F1: 0}
                        var v46 int16 = -32768
                        fmt.Println(fi41, fi40, v46)
                        _ = v45
                        _ = v46
                    } else {
                        var v47 uint64 = 4294967295
                        if false {
                            v48 := v1[fi41]
                            v47 += (v47 / (((v47 | (^(v47)))) | 1))
                            var v49 uint64 = (^(v47))
                            _ = v48
                            _ = v49
                        }
                        switch "a" {
                        case "hello":
                            v50 := GenS4{F0: v3, F1: 45, F2: genFn1(32767)}
                            _ = v50
                        }
                        fmt.Println(v47)
                        _ = v47
                    }
                case 32768:
                    switch (-(fi41)) {
                    case 964:
                        v51 := GenS4{F0: v3, F1: 128, F2: genFn1(-30)}
                        _ = v51
                    case 32767:
                        v52 := GenS5{F0: 113, F1: ("a" + "foo bar"), F2: 387}
                        _ = v52
                    case 9223372036854775807:
                        v3 = v3
                    default:
                        var v53 float64 = 0.0
                        _ = v53
                    }
                case -1:
                    v3 = v3
                }
                if (true && (!(-477 > genFn1(-957)))) {
                    v54 := GenS6{F0: -2147483648}
                    v55 := GenS6{F0: 2147483647}
                    _ = v54
                    _ = v55
                } else {
                    if true {
                        switch "\n\t" {
                        case "":
                            switch fi41 {
                            case 9223372036854775807:
                                v3 = (v3 & (v3 + 1))
                            }
                        case "a":
                            v56 := GenS7{F0: 1, F1: 42}
                            _ = v56
                        case "hello":
                            _ = len(v1)
                        }
                        var v57 float32 = float32(((1.5 * 0.1) * 3.14159))
                        v57 -= (v57 + v57)
                        _ = v57
                    } else {
                        var v58 string = "a"
                        v58 = ((v58 + (v58 + v58)) + ("a" + v58))
                        var v59 float32 = 1.5
                        v1[fi41] = 411
                        v60 := GenS1{F0: 42}
                        v58 = v58
                        var v61 bool = ((v58 + (v58 + "\n\t")) >= ("\n\t" + ((v58 + v58) + v58)))
                        var v62 bool = v61
                        _ = v58
                        _ = v59
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    switch int(v3) {
                    case 127:
                        v63 := GenS8{F0: 44, F1: 2147483647, F2: 42}
                        _ = v63
                    }
                    if (0.1 != 1e10) {
                        v3 = (^(v3))
                        v3 = (v3 * v3)
                        fmt.Println(v3, v3, fi41)
                    }
                }
                _ = v42
            }
            var v64 float64 = float64(0.1)
            _ = fi41
            _ = v64
        }
        var v65 int = fi40
        _ = fi40
        _ = v65
    }
    v3 += 803
    v66 := GenS6{F0: 184}
    v67 := v1[32767]
    v67 -= 2147483648
    v67 -= 0
    switch ("a" + (("foo bar" + "hello") + "")) {
    case "":
        for fi68 := 0; fi68 < 20; fi68++ {
            v3 = v3
            _ = fi68
        }
    case "a":
        var v69 int = -157
        _ = v69
    }
    if true {
        if ("hello" < "") {
            v70 := GenS2{F0: "foo bar", F1: v67, F2: "hello"}
            _ = v70
        } else {
            for fi71 := 0; fi71 < 17; fi71++ {
                for fi72 := 0; fi72 < 1; fi72++ {
                    v2 = append(v2, v67)
                    if (false || false) {
                        fmt.Println(fi71)
                        v73 := GenS3{F0: uint(v67), F1: 4294967295}
                        _ = v73
                    } else {
                        v67 += (^(v67))
                    }
                    v74 := GenS5{F0: 1, F1: "", F2: (v67 ^ 65535)}
                    _ = fi72
                    _ = v74
                }
                for fi75 := 0; fi75 < 13; fi75++ {
                    fmt.Println(fi71)
                    v76 := GenS8{F0: 0, F1: (fi75 << ((v3) & 63)), F2: 114}
                    _ = fi75
                    _ = v76
                }
                v3 = v3
                _ = fi71
            }
            v77 := GenS8{F0: uint64(v67), F1: -9223372036854775808, F2: -628}
            _ = v77
        }
        for fi78 := 0; fi78 < 15; fi78++ {
            var v79 int8 = 1
            for fi80 := 0; fi80 < 9; fi80++ {
                fmt.Println(fi78, fi80)
                var v81 uint32 = v67
                switch 0 {
                case 1:
                    switch 9223372036854775807 {
                    case -605:
                        var v82 float64 = 1.5
                        _ = v82
                    case 516:
                        v3 = (v3 >> ((v3) & 63))
                    default:
                        if (1000 != -1) {
                            var v83 int8 = v79
                            var v84 uint16 = 32767
                            v79 = v83
                            _ = v83
                            _ = v84
                        }
                    }
                case 716:
                    _ = len(v1)
                default:
                    var v85 int = int(v79)
                    _ = v85
                }
                _ = fi80
                _ = v81
            }
            var v86 uint = v3
            _ = fi78
            _ = v79
            _ = v86
        }
    }
    v3 += v3
    switch 4294967295 {
    case -2147483648:
        var v87 float32 = -2.25
        _ = v87
    case -99:
        var v88 int = -370
        _ = v88
    default:
        _ = len(v4)
    }
    var v89 int64 = 9223372036854775807
    delete(v1, 577)
    var v90 int32 = 0
    if false {
        for fi91 := 0; fi91 < 10; fi91++ {
            v89 = v89
            v92 := GenS5{F0: 1, F1: "hello", F2: v67}
            _ = fi91
            _ = v92
        }
    }
    _ = len(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v67:", v67)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v66
    _ = v67
    _ = v89
    _ = v90
}
