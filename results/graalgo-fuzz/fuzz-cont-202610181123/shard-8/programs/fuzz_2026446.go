package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 string }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int64; F2 string }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 int }
func genFn0(p0 int64, p1 int32) int64 {
    p1 = 42
    p1 = p1
    if false {
        if (-2.25 > (float32(0.0) + 0.1)) {
            fmt.Println(p1, p0)
            for fi6 := 0; fi6 < 11; fi6++ {
                p0 = (p0 / ((p0) | 1))
                if (-649 < 32767) {
                    p1 = p1
                } else {
                    p1 -= p1
                    v7 := []int{int(p1), 127}
                    fmt.Println(p1, fi6)
                    _ = v7
                }
                var v8 float32 = ((float32(1e10) + 0.0) / 1.5)
                _ = fi6
                _ = v8
            }
            var v9 string = "a"
            _ = v9
        } else {
            var v10 float32 = 0.1
            fmt.Println(p0, p1, p1)
            _ = v10
        }
        p0 = 9223372036854775807
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 591
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    var v11 bool = false
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(v1)
        var v13 int32 = v2
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 0}
    var v15 uint64 = 823
    if (!((float32((0.1 + (1.5 - (3.14159 - ((3.14159 - float64(((0.1 / 3.14159) / 1.0))) / float64(-2.25)))))) / float32(3.14159)) > 0.0)) {
        if v11 {
            v2 -= v2
        }
        if true {
            v4[65535] = v11
            for fi16 := 0; fi16 < 4; fi16++ {
                if (-128 == 124) {
                    v17 := []int{(-(fi16)), fi16}
                    _ = v17
                }
                _ = fi16
            }
            if ("\n\t" <= "a") {
                for fi18 := 0; fi18 < 8; fi18++ {
                    delete(v5, (fi18 / (((fi18 << (((v3 % (((^(v3))) | 1))) & 63))) | 1)))
                    v11 = (((v3 <= (v3 | 376)) && true) && true)
                    v1 -= (v1 / (((v1 - (-(v1)))) | 1))
                    _ = fi18
                }
                var v19 bool = v11
                if false {
                    v2 = v2
                    delete(v4, 886)
                    v19 = (!(!v11))
                } else {
                    var v20 uint = v3
                    fmt.Println(v3, v3)
                    _ = v20
                }
                _ = v19
            }
        }
    }
    if v11 {
        var v21 string = "hello"
        var v22 int64 = (v1 + -800)
        for fi23 := 0; fi23 < 3; fi23++ {
            v3 += v3
            if (!false) {
                var v24 int16 = 631
                _ = v24
            }
            v25 := GenS1{F0: 2}
            _ = fi23
            _ = v25
        }
        _ = v21
        _ = v22
    }
    switch 2147483647 {
    case -720:
        var v26 int32 = v2
        _ = v26
    case 4294967295:
        if false {
            if false {
                v27 := []int{127, 4294967295, 42}
                fmt.Println(v15, v3)
                for fi28 := 0; fi28 < 6; fi28++ {
                    switch "" {
                    case "\n\t":
                        delete(v5, 32767)
                    }
                    switch "foo bar" {
                    case "hello":
                        switch (fi28 >> ((v3) & 63)) {
                        case -1:
                            v29 := GenS2{F0: fi28, F1: "hello", F2: 379}
                            _ = v29
                        case -9223372036854775808:
                            var v30 uint8 = uint8(v15)
                            _ = v30
                        }
                    }
                    _ = fi28
                }
                _ = v27
            }
        }
    }
    switch "hello" {
    case "":
        for fi31 := 0; fi31 < 15; fi31++ {
            v11 = v11
            for fi32 := 0; fi32 < 18; fi32++ {
                switch ("a" + "foo bar") {
                case "a":
                    fmt.Println(fi31)
                    switch (-(fi31)) {
                    case -10:
                        fmt.Println(v3)
                    }
                    fmt.Println(v1, v2)
                case "hello":
                    var v33 int64 = v1
                    v34 := []int{128, fi32}
                    _ = len(v5)
                    var v35 int32 = v2
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = fi32
            }
            v36 := GenS1{F0: 61}
            _ = fi31
            _ = v36
        }
    case "a":
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 int = fi37
            _ = fi37
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 11; fi39++ {
            v2 = (v2 << ((v3) & 63))
            v11 = v11
            v40 := GenS3{F0: 1, F1: v15, F2: int16(v2)}
            _ = fi39
            _ = v40
        }
    }
    if v11 {
        var v41 int32 = (v2 >> ((v3) & 63))
        var v42 float32 = (0.0 * 1.5)
        fmt.Println(v11, v42)
        _ = v41
        _ = v42
    } else {
        var v43 uint16 = 497
        if v11 {
            v44 := GenS4{F0: 56}
            v1 += -37
            for fi45 := 0; fi45 < 7; fi45++ {
                fmt.Println(v3, v43)
                _ = fi45
            }
            _ = v44
        }
        var v46 int64 = genFn0(v1, v2)
        _ = v43
        _ = v46
    }
    v47 := []int{32767, -536}
    for fi48 := 0; fi48 < 16; fi48++ {
        var v49 uint32 = 1
        v5[32767] = v49
        v1 += 0
        _ = fi48
        _ = v49
    }
    switch "hello" {
    case "foo bar":
        var v50 uint8 = 0
        _ = v50
    case "hello":
        switch 127 {
        case 42:
            var v51 uint32 = 4294967295
            _ = v51
        case -1:
            var v52 int = -1
            _ = v52
        case 127:
            var v53 uint16 = 65535
            _ = v53
        default:
            if v11 {
                for fi54 := 0; fi54 < 16; fi54++ {
                    if v11 {
                        v55 := GenS5{F0: 42, F1: (v2 % (((v2 & v2)) | 1)), F2: "hello"}
                        var v56 int16 = 0
                        _ = v55
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 15; fi57++ {
                        var v58 bool = (32767 < -32768)
                        _ = fi57
                        _ = v58
                    }
                    if len(v47) > 0 {
                        v47[0] = int(v3)
                    }
                    _ = fi54
                }
            } else {
                if true {
                    var v59 uint16 = 1
                    _ = v59
                }
                if (v11 && true) {
                    if (1.5 >= 1.5) {
                        v60 := GenS6{F0: 1, F1: 2}
                        _ = v60
                    }
                    if false {
                        switch "hello" {
                        case "":
                            v61 := GenS5{F0: 593, F1: -2147483648, F2: ("a" + "a")}
                            _ = v61
                        case "hello":
                            v2 -= 1
                        default:
                            var v62 int = 1
                            _ = v62
                        }
                        for fi63 := 0; fi63 < 5; fi63++ {
                            v15 += v15
                            var v64 uint64 = (^(v15))
                            v15 = 1
                            _ = fi63
                            _ = v64
                        }
                        if true {
                            v65 := GenS1{F0: 127}
                            v66 := GenS2{F0: -2147483648, F1: "\n\t", F2: (-(v1))}
                            fmt.Println(v2, v11, v11)
                            _ = v65
                            _ = v66
                        } else {
                            switch 346 {
                            case 380:
                                if len(v47) > 0 {
                                    v47[0] = -9223372036854775808
                                }
                            case -1:
                                v4[65535] = (v1 == genFn0(v1, 586))
                            default:
                                v67 := GenS7{F0: 4294967295, F1: (v1 + -534), F2: ("a" + "foo bar")}
                                _ = v67
                            }
                            fmt.Println(v1)
                        }
                    }
                }
            }
        }
    default:
        var v68 string = ""
        _ = v68
    }
    switch 672 {
    case 2147483647:
        var v69 int16 = 1
        _ = v69
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            if true {
                for fi70 := 0; fi70 < 18; fi70++ {
                    var v71 uint64 = v15
                    _ = fi70
                    _ = v71
                }
                switch "" {
                case "hello":
                    var v72 float64 = ((0.0 - (0.1 + 3.14159)) / (1e10 * -2.25))
                    _ = v72
                }
            }
            v73 := GenS8{F0: 2147483647}
            _ = v73
        }
    }
    switch 32768 {
    case 128:
        v3 += v3
    case -9223372036854775808:
        var v74 string = ((("a" + ("foo bar" + ("" + "a"))) + (("foo bar" + ((("\n\t" + "a") + ("" + ("\n\t" + "foo bar"))) + ("a" + "a"))) + "")) + "\n\t")
        _ = v74
    case -2147483648:
        if v11 {
            v75 := v5[-2147483648]
            if (v11 && (0.1 == (1e10 - 3.14159))) {
                for fi76 := 0; fi76 < 9; fi76++ {
                    v77 := GenS9{F0: 1000}
                    if ((v3 | 255) != (v3 / ((289) | 1))) {
                        if v11 {
                            var v78 uint32 = 2147483648
                            switch fi76 {
                            case -499:
                                switch ("foo bar" + "") {
                                case "\n\t":
                                    var v79 uint = 945
                                    _ = v79
                                }
                            case 0:
                                for fi80 := 0; fi80 < 17; fi80++ {
                                    if false {
                                        v81 := GenS10{F0: v11, F1: (v1 - 0)}
                                        switch ("hello" + (((("foo bar" + "foo bar") + "a") + "foo bar") + ("a" + (("\n\t" + ("" + ("a" + ("hello" + "\n\t")))) + ("\n\t" + ""))))) {
                                        case "foo bar":
                                            v82 := GenS11{F0: -9223372036854775808}
                                            _ = v82
                                        case "hello":
                                            v83 := GenS7{F0: (v75 - (v78 * v78)), F1: (v1 * -1), F2: "\n\t"}
                                            _ = v83
                                        }
                                        if true {
                                            v75 -= v75
                                            v84 := v4[v75]
                                            _ = v84
                                        }
                                        _ = v81
                                    } else {
                                        fmt.Println(v78, v1)
                                    }
                                    v15 = v15
                                    v85 := GenS6{F0: 0, F1: 0}
                                    var v86 float64 = 1e10
                                    _ = fi80
                                    _ = v85
                                    _ = v86
                                }
                            case -111:
                                var v87 uint8 = 127
                                _ = v87
                            default:
                                v15 = (v15 * (v15 | (v15 & (v15 - (v15 | (v15 + (v15 >> ((v3) & 63))))))))
                            }
                            _ = v78
                        } else {
                            if ((^(fi76)) >= -1) {
                                fmt.Println(v75, v1)
                            }
                            switch 542 {
                            case 0:
                                var v88 int16 = 1000
                                _ = v88
                            case 42:
                                v89 := v5[904]
                                _ = v89
                            default:
                                var v90 int8 = 1
                                _ = v90
                            }
                            v15 -= v15
                        }
                        switch "a" {
                        case "":
                            v1 += (v1 & (v1 + int64(v3)))
                        case "foo bar":
                            var v91 int32 = 324
                            _ = v91
                        }
                    } else {
                        v92 := GenS11{F0: 42}
                        _ = v92
                    }
                    for fi93 := 0; fi93 < 8; fi93++ {
                        var v94 uint32 = (v75 & 65535)
                        v95 := v4[v75]
                        _ = fi93
                        _ = v94
                        _ = v95
                    }
                    _ = fi76
                    _ = v77
                }
                var v96 uint = v3
                _ = v96
            } else {
                fmt.Println(v11)
            }
            v3 -= (v3 >> (((v3 % ((v3) | 1))) & 63))
            _ = v75
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
