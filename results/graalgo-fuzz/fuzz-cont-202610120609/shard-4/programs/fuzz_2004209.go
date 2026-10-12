package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS9 struct { F0 uint32; F1 uint; F2 bool }
type GenS10 struct { F0 string; F1 bool; F2 int8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint }
func genFn0() uint16 {
    var v7 int64 = -1
    var v8 uint8 = 128
    v8 += 1
    switch "foo bar" {
    case "hello":
        if false {
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 float64 = (float64(-2.25) + -2.25)
                if false {
                    v10 -= 3.14159
                    var v11 uint = 673
                    _ = v11
                }
                v8 -= v8
                _ = fi9
                _ = v10
            }
        }
    case "foo bar":
        for fi12 := 0; fi12 < 3; fi12++ {
            v7 = 956
            _ = fi12
        }
    }
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint, p1 int64, p2 int) int64 {
    v13 := []int{}
    var v14 int = p2
    if false {
        var v15 int8 = -128
        _ = v15
    }
    if len(v13) > 0 {
        v13[0] = (p2 << (((p0 - 65535)) & 63))
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        _ = len(v13)
        for fi17 := 0; fi17 < 14; fi17++ {
            switch ("hello" + "a") {
            case "hello":
                var v18 string = ("" + ("" + "\n\t"))
                _ = v18
            case "foo bar":
                switch ("a" + (("" + "") + "")) {
                case "hello":
                    p2 -= v14
                }
            default:
                var v19 uint = p0
                _ = v19
            }
            if ((-(v14)) == 4294967295) {
                v20 := GenS1{F0: 922, F1: 18446744073709551615}
                _ = v20
            } else {
                var v21 uint8 = 42
                v22 := GenS2{F0: v21, F1: 255, F2: 4294967295}
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                _ = v21
                _ = v22
            }
            fmt.Println(p0, v14, fi17)
            _ = fi17
        }
        if (825 < uint16(p0)) {
            if true {
                v14 = (fi16 >> ((p0) & 63))
                if true {
                    v23 := GenS3{F0: 42, F1: genFn0(), F2: false}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        fmt.Println(fi16, fi16)
                        var v25 int8 = -94
                        _ = fi24
                        _ = v25
                    }
                    v14 = (p2 % ((128) | 1))
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    v14 -= fi26
                    var v27 uint = 0
                    var v28 uint64 = 2
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                p1 = -893
            }
            if true {
                p0 = p0
                p2 += p2
                for fi29 := 0; fi29 < 11; fi29++ {
                    switch fi29 {
                    case 423:
                        switch -296 {
                        case 2147483647:
                            v30 := GenS4{F0: -1}
                            _ = v30
                        case 384:
                            if (!(!true)) {
                                switch "\n\t" {
                                case "a":
                                    _ = len(v13)
                                case "\n\t":
                                    if (549 == 1) {
                                        var v31 int64 = (p1 - p1)
                                        _ = v31
                                    } else {
                                        if (uint16(p0) != 42) {
                                            var v32 int16 = -302
                                            _ = v32
                                        }
                                        fmt.Println(fi29, p2)
                                        _ = len(v13)
                                    }
                                default:
                                    if ((p1 * 1) >= p1) {
                                        v33 := GenS2{F0: 255, F1: 264, F2: 2}
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        var v34 string = ""
                                        v35 := GenS5{F0: (v34 + (v34 + v34)), F1: 1}
                                        v36 := GenS2{F0: 170, F1: 305, F2: 1}
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    } else {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                    }
                                }
                                v37 := GenS4{F0: -1}
                                p1 += p1
                                p0 = (p0 + (p0 >> (((p0 * 427)) & 63)))
                                fmt.Println(p2, p2, p0)
                                _ = v37
                            } else {
                                switch "a" {
                                case "a":
                                    if (42 == 18446744073709551615) {
                                        if ("\n\t" < ("a" + "\n\t")) {
                                            var v38 uint32 = 2147483647
                                            _ = v38
                                        } else {
                                            v14 = -811
                                            var v39 float64 = (1.5 - 1e10)
                                            var v40 uint8 = 188
                                            _ = v39
                                            _ = v40
                                        }
                                        v41 := GenS3{F0: -2147483648, F1: uint16(p0), F2: (!(!(false || true)))}
                                        if (!false) {
                                            v42 := GenS2{F0: 42, F1: 677, F2: 4294967295}
                                            p1 -= (-(p1))
                                            p2 = 128
                                            p0 += p0
                                            _ = v42
                                        } else {
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                            switch ((((((("foo bar" + "foo bar") + ("" + "foo bar")) + ("foo bar" + "")) + "hello") + "\n\t") + "") + "\n\t") {
                                            case "foo bar":
                                                v13 = append(v13, fi29)
                                            case "a":
                                                p1 = (p1 & 9223372036854775807)
                                            case "":
                                                v43 := GenS1{F0: 32767, F1: 4294967295}
                                                _ = v43
                                            default:
                                                v13 = append(v13, (fi16 + 981))
                                            }
                                            var v44 float32 = -2.25
                                            _ = v44
                                        }
                                        _ = v41
                                    } else {
                                        p2 -= (fi16 + -480)
                                    }
                                    p0 += p0
                                    var v45 uint8 = 66
                                    var v46 int8 = 100
                                    var v47 int64 = 9223372036854775807
                                    var v48 uint = 4294967295
                                    fmt.Println(p2, v14)
                                    if false {
                                        if false {
                                            v13 = append(v13, (fi16 << (((^(p0))) & 63)))
                                            p2 = -1
                                        }
                                    } else {
                                        switch "hello" {
                                        case "foo bar":
                                            var v49 int8 = -1
                                            _ = v49
                                        case "hello":
                                            var v50 float64 = 1e10
                                            _ = v50
                                        default:
                                            v47 = v47
                                        }
                                    }
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                case "hello":
                                    var v51 uint32 = 2
                                    _ = v51
                                }
                                p2 = p2
                            }
                            fmt.Println(v14, v14, p2)
                            if (!false) {
                                switch (fi29 >> ((p0) & 63)) {
                                case 42:
                                    fmt.Println(p0, v14, fi16)
                                case -660:
                                    v52 := GenS6{F0: fi16, F1: -426}
                                    _ = v52
                                }
                                if true {
                                    var v53 int8 = 42
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    v53 -= (v53 ^ (v53 * v53))
                                    _ = v53
                                } else {
                                    var v54 bool = ("hello" == "foo bar")
                                    p1 = 42
                                    v13 = append(v13, -658)
                                    _ = v54
                                }
                                var v55 uint32 = 518
                                _ = v55
                            }
                        case -461:
                            switch "hello" {
                            case "":
                                v56 := GenS1{F0: 98, F1: 0}
                                _ = v56
                            default:
                                var v57 uint64 = 0
                                _ = v57
                            }
                        default:
                            var v58 float64 = (3.14159 - ((0.1 - -2.25) * float64(float32((1e10 - 1e10)))))
                            _ = v58
                        }
                    default:
                        p0 += p0
                    }
                    var v59 int16 = 751
                    _ = fi29
                    _ = v59
                }
            } else {
                for fi60 := 0; fi60 < 13; fi60++ {
                    var v61 uint16 = 32767
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = fi60
                    _ = v61
                }
                var v62 int = (-(v14))
                _ = v62
            }
        } else {
            if (1.5 > float64(0.1)) {
                if true {
                    var v63 uint64 = 18446744073709551615
                    p0 = 42
                    var v64 uint64 = (v63 & (v63 / (((v63 + (v63 | v63))) | 1)))
                    _ = v63
                    _ = v64
                } else {
                    fmt.Println(p0)
                    for fi65 := 0; fi65 < 6; fi65++ {
                        var v66 int32 = 0
                        _ = fi65
                        _ = v66
                    }
                }
                switch (-(v14)) {
                case 0:
                    if false {
                        if (!true) {
                            p0 = p0
                        } else {
                            var v67 int64 = 1
                            p1 = (^(v67))
                            v14 += int(p0)
                            _ = v67
                        }
                        var v68 int32 = -629
                        var v69 int64 = int64(p2)
                        _ = v68
                        _ = v69
                    } else {
                        _ = len(v13)
                    }
                }
            } else {
                v70 := GenS7{F0: (p1 << ((p0) & 63)), F1: -2147483648}
                v71 := GenS4{F0: -1}
                v72 := GenS8{F0: genFn0(), F1: 127, F2: 127}
                _ = v70
                _ = v71
                _ = v72
            }
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    v3 := make(map[uint32]bool)
    var v4 uint64 = 4294967295
    var v5 string = ("a" + "")
    v6 := []string{"\n\t", v5, "a", "foo bar"}
    switch v5 {
    case "":
        v4 = v4
    case "\n\t":
        fmt.Println(v4, v4)
    case "a":
        v73 := GenS6{F0: 623, F1: 61}
        _ = v73
    default:
        var v74 int64 = 677
        _ = v74
    }
    fmt.Println(v4)
    switch "" {
    case "a":
        var v75 int8 = -128
        _ = v75
    case "hello":
        v4 = 0
    case "":
        var v76 bool = false
        _ = v76
    default:
        for fi77 := 0; fi77 < 18; fi77++ {
            v4 -= (v4 + (v4 >> ((255) & 63)))
            if (311 > (v4 * 42)) {
                v78 := v3[uint32(v4)]
                _ = v78
            } else {
                var v79 uint8 = 255
                var v80 int16 = 0
                for fi81 := 0; fi81 < 14; fi81++ {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    _ = fi81
                }
                _ = v79
                _ = v80
            }
            if (v4 > (v4 * v4)) {
                switch fi77 {
                case -1:
                    var v82 bool = false
                    _ = v82
                case 128:
                    switch "foo bar" {
                    case "foo bar":
                        v4 = (^(v4))
                    case "a":
                        v83 := GenS9{F0: 0, F1: 4294967295, F2: (-629 == -1)}
                        _ = v83
                    case "hello":
                        delete(v1, ((("a" + v5) + "foo bar") + (v5 + (("hello" + "foo bar") + ("\n\t" + v5)))))
                    default:
                        for fi84 := 0; fi84 < 15; fi84++ {
                            v5 = v5
                            fmt.Println(fi77, fi77, v5)
                            fmt.Println(v5, v4)
                            _ = fi84
                        }
                    }
                case -221:
                    v85 := GenS10{F0: v5, F1: (255 == 964), F2: -1}
                    _ = v85
                }
                for fi86 := 0; fi86 < 13; fi86++ {
                    v5 = v5
                    var v87 int32 = -1
                    v87 += v87
                    _ = fi86
                    _ = v87
                }
            } else {
                _ = len(v2)
                v5 = "foo bar"
                v4 -= v4
            }
            _ = fi77
        }
    }
    v6 = append(v6, v5)
    var v88 bool = true
    fmt.Println(v5)
    switch 361 {
    case 32768:
        if (0 > 0) {
            v88 = v88
            fmt.Println(v4, v5, v5)
        }
    }
    if false {
        var v89 uint16 = uint16(v4)
        v4 -= (v4 << ((65535) & 63))
        _ = len(v3)
        _ = v89
    }
    v90 := GenS7{F0: genFn1(813, -165, 32767), F1: -1}
    switch "\n\t" {
    case "\n\t":
        v91 := GenS11{F0: 42}
        _ = v91
    case "":
        v88 = v88
    case "a":
        if (100 > -1) {
            v92 := GenS12{F0: 859}
            v93 := GenS2{F0: 100, F1: 65535, F2: 42}
            fmt.Println(v5, v5)
            _ = v92
            _ = v93
        }
    default:
        var v94 uint8 = 204
        _ = v94
    }
    v95 := GenS8{F0: 0, F1: 42, F2: 0}
    if v88 {
        var v96 bool = (1 > -420)
        if (!(127 <= -117)) {
            var v97 bool = (100 == -1)
            _ = v97
        }
        _ = v96
    }
    v88 = (!((((v5 + (v5 + "foo bar")) + "foo bar") + v5) != ("foo bar" + ("a" + (("" + "") + v5)))))
    fmt.Println(v5)
    v4 -= (v4 / (((v4 % ((v4) | 1))) | 1))
    v88 = (!v88)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
    _ = v90
    _ = v95
}
