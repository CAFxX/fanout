package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 string }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint64; F2 int }
type GenS6 struct { F0 uint16; F1 uint64; F2 string }
func genFn0(p0 bool, p1 int16) uint {
    v7 := GenS1{F0: -9223372036854775808}
    if ((3.14159 - 3.14159) != (1.5 / 1.0)) {
        if p0 {
            v8 := []int{127}
            _ = v8
        } else {
            if true {
                v9 := GenS1{F0: 42}
                v10 := GenS1{F0: 42}
                for fi11 := 0; fi11 < 8; fi11++ {
                    if p0 {
                        v12 := []int{fi11}
                        p1 = 1000
                        _ = v12
                    }
                    v13 := GenS1{F0: 690}
                    _ = fi11
                    _ = v13
                }
                _ = v10
                _ = v9
            }
        }
        p0 = (258 < 2)
    } else {
        v14 := GenS2{F0: 42, F1: "foo bar"}
        var v15 uint8 = 42
        _ = v14
        _ = v15
    }
    switch "hello" {
    case "hello":
        if ((p1 - (-(p1))) >= (p1 % ((p1) | 1))) {
            var v16 int32 = int32(p1)
            v16 += (v16 ^ (v16 >> ((255) & 63)))
            _ = v16
        }
    }
    var v17 float32 = 0.0
    v17 = float32(1e10)
    _ = p0
    _ = p1
    _ = v17
    _ = v7
    return 7
}
func genFn1(p0 uint8, p1 uint64) int64 {
    fmt.Println(p1, p0)
    if (false || (true || (42 >= 0))) {
        v18 := []int{-1}
        v19 := GenS2{F0: (p1 & p1), F1: (("\n\t" + "") + "\n\t")}
        var v20 int32 = -553
        _ = v18
        _ = v19
        _ = v20
    }
    p1 = 42
    _ = p0
    _ = p1
    return 553
}

func main() {
    var v1 int16 = -32768
    var v2 float64 = ((1.5 / 0.1) - 0.1)
    var v3 int8 = 42
    var v4 bool = (585 == 488)
    var v5 int8 = v3
    var v6 int = 32768
    if false {
        v4 = true
    } else {
        var v21 int32 = -394
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 float64 = (v2 * 1.5)
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    var v24 int64 = 1
    if (-128 <= (v5 | (v3 % (((v3 & v5)) | 1)))) {
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 string = "hello"
            var v27 int64 = (-(v24))
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := []int{}
                _ = fi28
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        v24 -= v24
    }
    v30 := GenS1{F0: 9223372036854775807}
    switch (("a" + ("foo bar" + "a")) + ("hello" + ("" + "foo bar"))) {
    case "":
        var v31 int64 = 1
        _ = v31
    }
    if (18446744073709551615 >= 582) {
        if (false && v4) {
            v32 := []int{v6, v6, (v6 << ((255) & 63))}
            for fi33 := 0; fi33 < 9; fi33++ {
                v34 := GenS1{F0: (v24 % ((9223372036854775807) | 1))}
                v3 -= -128
                if (false && (2 <= 1)) {
                    v3 = v3
                }
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := GenS3{F0: 848, F1: v4, F2: ("a" + "a")}
            if (95 > 127) {
                if v4 {
                    var v37 uint8 = 1
                    _ = v37
                } else {
                    switch (("hello" + "hello") + "foo bar") {
                    case "a":
                        var v38 int64 = (v24 - 0)
                        _ = v38
                    default:
                        fmt.Println(v2)
                    }
                    v24 -= v24
                    if (float32(0.1) >= 1.5) {
                        switch "foo bar" {
                        case "foo bar":
                            v2 -= (((v2 + (float64(-2.25) + -2.25)) + 0.1) - v2)
                        case "":
                            fmt.Println(v6)
                        default:
                            v39 := []int{2147483647, (^(fi35)), v6, 9223372036854775807}
                            _ = v39
                        }
                        v4 = v4
                        for fi40 := 0; fi40 < 1; fi40++ {
                            v41 := []int{104, (v6 >> ((716) & 63)), (fi35 ^ (fi40 ^ (^(fi35)))), 42}
                            var v42 int64 = genFn1(0, 0)
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    } else {
                        switch (fi35 + fi35) {
                        case -1:
                            fmt.Println(v24)
                        case -2147483648:
                            v43 := GenS1{F0: (v24 * v24)}
                            _ = v43
                        default:
                            for fi44 := 0; fi44 < 19; fi44++ {
                                v45 := []int{(v6 - 4294967295), (fi35 % (((^(v6))) | 1)), (fi35 - (fi44 | fi35))}
                                v46 := GenS3{F0: 774, F1: (!v4), F2: "foo bar"}
                                var v47 int64 = genFn1(2, 42)
                                _ = fi44
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        }
                    }
                }
                if ((true || (!false)) || (!true)) {
                    v48 := GenS2{F0: 474, F1: ("hello" + "")}
                    var v49 uint = 42
                    v2 = (1.5 * (v2 / v2))
                    _ = v48
                    _ = v49
                }
                v50 := []int{}
                _ = v50
            } else {
                for fi51 := 0; fi51 < 4; fi51++ {
                    fmt.Println(v4)
                    if ((v1 << ((65535) & 63)) > -925) {
                        fmt.Println(fi35, v1, v3)
                        v52 := []int{fi35}
                        fmt.Println(v3, v2, fi51)
                        _ = v52
                    } else {
                        var v53 int8 = 1
                        v54 := GenS4{F0: v53, F1: 32767}
                        if true {
                            switch "hello" {
                            case "a":
                                v53 = v53
                            }
                            v55 := GenS3{F0: 2147483648, F1: ("\n\t" >= ""), F2: ("hello" + "hello")}
                            _ = v55
                        } else {
                            v56 := []int{1}
                            v3 = (v5 << ((42) & 63))
                            var v57 float64 = 0.1
                            _ = v56
                            _ = v57
                        }
                        _ = v53
                        _ = v54
                    }
                    _ = fi51
                }
                fmt.Println(fi35)
            }
            v58 := GenS2{F0: 572, F1: "hello"}
            _ = fi35
            _ = v36
            _ = v58
        }
        for fi59 := 0; fi59 < 19; fi59++ {
            v2 = v2
            fmt.Println(v4)
            var v60 int8 = (v5 ^ (v5 >> ((genFn0(v4, v1)) & 63)))
            _ = fi59
            _ = v60
        }
    } else {
        fmt.Println(v2)
        v61 := []int{}
        for fi62 := 0; fi62 < 20; fi62++ {
            fmt.Println(fi62, v2)
            _ = fi62
        }
        _ = v61
    }
    for fi63 := 0; fi63 < 13; fi63++ {
        var v64 float64 = v2
        _ = fi63
        _ = v64
    }
    if ((314 <= 2) && v4) {
        v6 += (-(v6))
        for fi65 := 0; fi65 < 12; fi65++ {
            v4 = (226 == -327)
            _ = fi65
        }
        var v66 float64 = 3.14159
        _ = v66
    } else {
        for fi67 := 0; fi67 < 11; fi67++ {
            var v68 float64 = 0.0
            _ = fi67
            _ = v68
        }
    }
    fmt.Println(v2, v3)
    for fi69 := 0; fi69 < 14; fi69++ {
        v70 := GenS5{F0: 100, F1: 518, F2: 0}
        for fi71 := 0; fi71 < 10; fi71++ {
            if (!(false && (true || (false && (("\n\t" + "") == ("\n\t" + ("hello" + ""))))))) {
                var v72 bool = (float64((((0.1 - ((1e10 - 0.0) / 1e10)) - -2.25) / 3.14159)) != v2)
                if ((v5 << ((0) & 63)) != -77) {
                    v4 = false
                    if (42 >= 0) {
                        var v73 uint64 = 0
                        var v74 int = -1
                        _ = v73
                        _ = v74
                    } else {
                        var v75 bool = (0 > 1)
                        v75 = true
                        _ = v75
                    }
                }
                v3 += -58
                _ = v72
            }
            if v4 {
                if (!(!v4)) {
                    var v76 string = ("\n\t" + (("hello" + "foo bar") + "foo bar"))
                    v77 := GenS4{F0: v5, F1: 903}
                    v78 := GenS6{F0: 65535, F1: 2, F2: ("\n\t" + "")}
                    _ = v76
                    _ = v77
                    _ = v78
                } else {
                    fmt.Println(v4, v6)
                    var v79 float64 = v2
                    _ = v79
                }
                v80 := []int{(^(fi69)), v6, (fi69 % (((fi69 ^ int(v5))) | 1))}
                _ = v80
            }
            if ((18446744073709551615 != 42) && (!v4)) {
                if (!((v24 + -1) >= v24)) {
                    var v81 int32 = -2147483648
                    if (v6 < (-(fi69))) {
                        var v82 uint32 = 65535
                        v5 = v3
                        if ((!(!(1 < 1))) || v4) {
                            v6 -= (fi71 - fi69)
                        }
                        _ = v82
                    }
                    v83 := GenS3{F0: 384, F1: ((("foo bar" + ("\n\t" + "hello")) + "foo bar") >= "a"), F2: "a"}
                    _ = v81
                    _ = v83
                } else {
                    var v84 int8 = v3
                    v6 = fi71
                    v4 = (fi69 != fi71)
                    _ = v84
                }
            }
            v2 = (v2 / -2.25)
            switch "hello" {
            case "\n\t":
                var v85 uint8 = 2
                _ = v85
            }
            if (558 < 2) {
                v86 := []int{int(v1), v6, v6, (^(v6))}
                v24 = genFn1(91, 2)
                var v87 uint32 = 785
                _ = v86
                _ = v87
            }
            _ = fi71
        }
        var v88 int8 = (v5 + v5)
        _ = fi69
        _ = v70
        _ = v88
    }
    v89 := GenS2{F0: 18446744073709551615, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v89
}
