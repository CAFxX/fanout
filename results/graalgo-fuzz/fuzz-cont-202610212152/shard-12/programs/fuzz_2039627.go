package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int32; F1 string }
func genFn0(p0 int16, p1 int, p2 int64) uint64 {
    switch p1 {
    case 4294967295:
        v7 := GenS1{F0: true}
        _ = v7
    }
    if true {
        switch "hello" {
        case "foo bar":
            v8 := []int{(^(p1)), 32768, p1}
            _ = v8
        case "hello":
            if (!((!(!(!false))) && false)) {
                p0 = (p0 << ((214) & 63))
            } else {
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 int16 = (p0 / (((p0 << ((448) & 63))) | 1))
                    var v11 int8 = -119
                    var v12 string = ("a" + "a")
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
        case "":
            var v13 int = p1
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() uint16 {
    if (("a" + "a") < ("" + "foo bar")) {
        if true {
            var v14 uint64 = 1
            for fi15 := 0; fi15 < 8; fi15++ {
                for fi16 := 0; fi16 < 20; fi16++ {
                    v14 += v14
                    v14 += v14
                    if (((-2.25 - (1e10 * -2.25)) != 1.5) && false) {
                        v17 := []int{fi16, -9223372036854775808}
                        fmt.Println(fi15, fi16, fi16)
                        _ = v17
                    }
                    var v18 uint8 = 229
                    _ = fi16
                    _ = v18
                }
                var v19 uint16 = 678
                _ = fi15
                _ = v19
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                switch (fi20 * fi20) {
                case 0:
                    v14 -= v14
                case 4294967295:
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v22 := []int{(fi20 - -373), fi21, fi20}
                        fmt.Println(fi21)
                        _ = fi21
                        _ = v22
                    }
                case 376:
                    var v23 bool = (42 < 575)
                    _ = v23
                default:
                    switch 9223372036854775807 {
                    case -1:
                        for fi24 := 0; fi24 < 16; fi24++ {
                            var v25 int32 = 1
                            var v26 uint32 = 724
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                }
                _ = fi20
            }
            _ = v14
        }
    } else {
        switch 32768 {
        case -9223372036854775808:
            var v27 float32 = float32(0.1)
            _ = v27
        default:
            var v28 int = -9223372036854775808
            _ = v28
        }
        for fi29 := 0; fi29 < 7; fi29++ {
            v30 := GenS1{F0: (!false)}
            if true {
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 int8 = -112
                    _ = fi31
                    _ = v32
                }
            } else {
                switch fi29 {
                case 42:
                    v33 := GenS1{F0: (!(1 >= 42))}
                    _ = v33
                case 4294967295:
                    v34 := GenS2{F0: 18446744073709551615, F1: (0 >= 1), F2: 207}
                    _ = v34
                case -9223372036854775808:
                    if false {
                        if true {
                            var v35 int16 = -591
                            v36 := GenS3{F0: -1, F1: 0}
                            _ = v35
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 17; fi37++ {
                            var v38 bool = (((1e10 - (3.14159 / 0.1)) + 1.5) <= 0.1)
                            v38 = false
                            fmt.Println(fi37)
                            var v39 float64 = float64(0.0)
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        var v40 int16 = -1
                        _ = v40
                    }
                default:
                    switch fi29 {
                    case -465:
                        var v41 uint64 = 4294967295
                        _ = v41
                    case 32767:
                        v42 := []int{fi29, fi29, (fi29 ^ 42)}
                        _ = v42
                    case -9223372036854775808:
                        var v43 int = -901
                        _ = v43
                    default:
                        var v44 uint32 = 2
                        _ = v44
                    }
                }
                for fi45 := 0; fi45 < 18; fi45++ {
                    v46 := []int{}
                    _ = fi45
                    _ = v46
                }
                v47 := []int{(fi29 >> ((340) & 63)), fi29, 9223372036854775807}
                _ = v47
            }
            _ = fi29
            _ = v30
        }
        for fi48 := 0; fi48 < 10; fi48++ {
            v49 := []int{42, -9223372036854775808, (fi48 << ((129) & 63))}
            switch 4294967295 {
            case 4294967295:
                var v50 uint8 = 0
                _ = v50
            case 42:
                var v51 int64 = 9223372036854775807
                _ = v51
            case 127:
                var v52 string = "hello"
                _ = v52
            }
            fmt.Println(fi48)
            _ = fi48
            _ = v49
        }
    }
    var v53 string = "hello"
    fmt.Println(v53)
    v54 := []int{42}
    v53 = "\n\t"
    _ = v53
    _ = v54
    return 38
}

func main() {
    var v1 int16 = 417
    var v2 uint = 924
    v3 := []string{}
    var v4 uint32 = 584
    v5 := make(map[int]uint32)
    var v6 int16 = 1000
    v2 = (v2 >> ((v2) & 63))
    for fi55 := 0; fi55 < 1; fi55++ {
        var v56 int8 = 42
        for fi57 := 0; fi57 < 5; fi57++ {
            v58 := v5[(fi57 % (((fi57 >> (((v2 ^ v2)) & 63))) | 1))]
            var v59 bool = (fi55 > (fi57 & (fi57 + fi55)))
            var v60 int8 = (v56 % ((v56) | 1))
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
        }
        for fi61 := 0; fi61 < 17; fi61++ {
            if (v4 >= v4) {
                var v62 uint8 = 127
                fmt.Println(v1)
                v56 = (v56 + (v56 >> ((v2) & 63)))
                _ = v62
            }
            v6 -= (v6 << (((^(v2))) & 63))
            var v63 int32 = -1
            _ = fi61
            _ = v63
        }
        _ = fi55
        _ = v56
    }
    if false {
        fmt.Println(v1, v4, v4)
        for fi64 := 0; fi64 < 8; fi64++ {
            if true {
                var v65 int8 = 1
                _ = v65
            }
            if (!(!((true && ((v1 % (((v6 | (v1 | 1000))) | 1)) < -32768)) && (4294967295 <= 0)))) {
                v66 := GenS2{F0: 2, F1: true, F2: 42}
                var v67 int = (fi64 & (fi64 * (^(fi64))))
                for fi68 := 0; fi68 < 1; fi68++ {
                    var v69 int8 = -1
                    v70 := GenS4{F0: -856, F1: uint64(v4), F2: 109}
                    _ = len(v5)
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                _ = v66
                _ = v67
            } else {
                if true {
                    for fi71 := 0; fi71 < 4; fi71++ {
                        if true {
                            v72 := GenS4{F0: 1, F1: 0, F2: 128}
                            if false {
                                v73 := GenS4{F0: 1, F1: 4294967295, F2: 128}
                                var v74 int16 = 0
                                _ = v73
                                _ = v74
                            } else {
                                var v75 int64 = 9223372036854775807
                                switch 32767 {
                                case -516:
                                    if false {
                                        switch "a" {
                                        case "\n\t":
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(v75, v1)
                                            default:
                                                v76 := GenS4{F0: 0, F1: 18446744073709551615, F2: 1}
                                                _ = v76
                                            }
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if (-1 == 42) {
                                            v1 = -1
                                            var v77 int = (fi64 + (fi64 % ((-915) | 1)))
                                            _ = v77
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    } else {
                                        v75 += v75
                                        _ = len(v3)
                                        switch fi64 {
                                        case 4294967295:
                                            v75 -= v75
                                        case 981:
                                            var v78 uint32 = (^(v4))
                                            _ = v78
                                        }
                                    }
                                }
                                var v79 uint8 = 128
                                _ = v75
                                _ = v79
                            }
                            var v80 string = "hello"
                            _ = v72
                            _ = v80
                        }
                        _ = fi71
                    }
                    v81 := v5[fi64]
                    if ((fi64 != fi64) || true) {
                        var v82 int64 = 279
                        v2 = 42
                        _ = v82
                    } else {
                        v81 += (v4 / ((v81) | 1))
                        v2 -= 4294967295
                        switch ("" + ((("a" + "hello") + "hello") + ("\n\t" + ((("hello" + "a") + "a") + "a")))) {
                        case "\n\t":
                            v4 = v4
                        case "foo bar":
                            v1 = 42
                        case "hello":
                            v1 += (^(v6))
                        default:
                            var v83 uint32 = (v81 | v4)
                            _ = v83
                        }
                    }
                    _ = v81
                }
                var v84 uint8 = 255
                for fi85 := 0; fi85 < 12; fi85++ {
                    if false {
                        v4 = (^(v4))
                    }
                    fmt.Println(v84)
                    fmt.Println(v1, v1)
                    _ = fi85
                }
                _ = v84
            }
            _ = fi64
        }
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v2 = (v2 | v2)
    delete(v5, 2147483647)
    for fi86 := 0; fi86 < 5; fi86++ {
        v6 += v6
        _ = fi86
    }
    v4 = (v4 ^ (v4 * v4))
    var v87 int64 = -9223372036854775808
    if (!false) {
        v2 += (v2 ^ v2)
    } else {
        v6 = v1
        for fi88 := 0; fi88 < 17; fi88++ {
            fmt.Println(v4)
            var v89 int64 = v87
            _ = fi88
            _ = v89
        }
        switch ("" + "") {
        case "foo bar":
            var v90 int8 = -52
            _ = v90
        case "a":
            for fi91 := 0; fi91 < 3; fi91++ {
                if false {
                    if false {
                        for fi92 := 0; fi92 < 6; fi92++ {
                            var v93 int16 = (v1 - v6)
                            v94 := GenS5{F0: -2147483648, F1: (("foo bar" + "a") + "")}
                            if len(v3) > 0 {
                                v3[0] = (("" + (("a" + "hello") + "")) + "foo bar")
                            }
                            _ = fi92
                            _ = v93
                            _ = v94
                        }
                        switch "hello" {
                        case "":
                            var v95 float32 = float32((0.1 - ((1e10 - 1.5) / 1.5)))
                            _ = v95
                        }
                        var v96 uint64 = 2
                        _ = v96
                    } else {
                        fmt.Println(v1)
                    }
                    fmt.Println(v2, v1)
                } else {
                    if true {
                        switch (fi91 / (((fi91 / (((^(fi91))) | 1))) | 1)) {
                        case -889:
                            for fi97 := 0; fi97 < 15; fi97++ {
                                fmt.Println(v4, v87)
                                _ = fi97
                            }
                        }
                    } else {
                        switch "hello" {
                        case "foo bar":
                            v98 := v5[int(v2)]
                            _ = v98
                        case "":
                            switch ("\n\t" + "a") {
                            case "a":
                                for fi99 := 0; fi99 < 1; fi99++ {
                                    v5[(fi99 ^ fi99)] = (v4 ^ (^(v4)))
                                    _ = fi99
                                }
                            default:
                                v4 = (v4 ^ (v4 | v4))
                            }
                        case "a":
                            if false {
                                v2 -= v2
                            }
                        }
                        var v100 string = ""
                        var v101 int = 785
                        _ = v100
                        _ = v101
                    }
                }
                _ = fi91
            }
        case "":
            v5[-643] = 691
        }
    }
    if (-1 <= -728) {
        fmt.Println(v2, v2, v2)
    } else {
        v87 = (-(v87))
    }
    v87 += (v87 - (^(v87)))
    v102 := GenS1{F0: true}
    var v103 float32 = -2.25
    v6 = (v6 / (((^(v6))) | 1))
    fmt.Println(v2, v6, v4)
    var v104 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v87
}
