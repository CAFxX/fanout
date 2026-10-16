package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 int16; F2 int16 }
type GenS8 struct { F0 uint32; F1 int32; F2 uint }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64; F1 int8 }
type GenS11 struct { F0 int64; F1 int8; F2 int32 }
type GenS12 struct { F0 uint; F1 int32 }
func genFn0() int {
    var v5 uint64 = 18446744073709551615
    if false {
        fmt.Println(v5)
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := []int{(^(fi6))}
            for fi8 := 0; fi8 < 20; fi8++ {
                if false {
                    if true {
                        if (("foo bar" != ("foo bar" + "hello")) && (false && true)) {
                            switch fi6 {
                            case -162:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            case -730:
                                var v9 int8 = -128
                                _ = v9
                            case 32767:
                                v10 := GenS1{F0: false, F1: 2147483647}
                                _ = v10
                            default:
                                var v11 string = "a"
                                _ = v11
                            }
                            _ = len(v7)
                            var v12 int = (^(fi8))
                            var v13 int32 = 1
                            var v14 uint = 0
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        var v15 float32 = (1.5 * -2.25)
                        v5 = v5
                        _ = v15
                    }
                    var v16 bool = true
                    _ = v16
                }
                var v17 int = fi6
                var v18 int8 = 126
                _ = fi8
                _ = v17
                _ = v18
            }
            _ = fi6
            _ = v7
        }
        v19 := GenS1{F0: true, F1: -1}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 9; fi20++ {
            switch (-(fi20)) {
            case -2147483648:
                if ((uint32(v5) >= 866) || ((128 <= 181) && ((505 > 1) || false))) {
                    v21 := GenS2{F0: 0}
                    v22 := GenS3{F0: -1}
                    if false {
                        var v23 uint64 = v5
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                } else {
                    var v24 float32 = 0.0
                    _ = v24
                }
            case -872:
                switch (-(fi20)) {
                case -115:
                    v5 = v5
                case 32768:
                    switch (fi20 >> ((0) & 63)) {
                    case 4294967295:
                        v25 := GenS4{F0: 696, F1: 1}
                        _ = v25
                    case 128:
                        v26 := []int{}
                        _ = v26
                    case 9223372036854775807:
                        var v27 int = fi20
                        _ = v27
                    default:
                        v28 := []int{fi20}
                        _ = v28
                    }
                default:
                    fmt.Println(fi20, fi20)
                }
            case 819:
                v29 := GenS5{F0: uint16(v5), F1: -654}
                _ = v29
            }
            var v30 int64 = 407
            _ = fi20
            _ = v30
        }
        if false {
            v31 := GenS4{F0: -435, F1: 127}
            switch -9223372036854775808 {
            case 127:
                v5 -= (^(v5))
            }
            _ = v31
        }
        v5 -= v5
    }
    v32 := []int{128}
    _ = v32
    _ = v5
    return 0
}
func genFn1(p0 uint8, p1 int) uint {
    p0 = 42
    switch (p1 % ((p1) | 1)) {
    case 1:
        switch (p1 >> ((uint(p0)) & 63)) {
        case -2147483648:
            for fi33 := 0; fi33 < 9; fi33++ {
                v34 := []int{(p1 - (fi33 % (((fi33 >> ((2) & 63))) | 1))), (^(fi33)), genFn0()}
                switch 128 {
                case 127:
                    p0 += (p0 >> ((4294967295) & 63))
                }
                var v35 uint = 2
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    case -9223372036854775808:
        p1 -= (p1 << ((0) & 63))
    case 552:
        if false {
            var v36 bool = (true && false)
            var v37 uint8 = (p0 - 1)
            v36 = v36
            _ = v36
            _ = v37
        } else {
            if true {
                var v38 int16 = 0
                _ = v38
            } else {
                p1 = (p1 & (p1 * 0))
                v39 := []int{p1, (p1 >> ((255) & 63))}
                _ = v39
            }
        }
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        var v41 int8 = 64
        if ((("foo bar" + ("hello" + "foo bar")) > ("a" + "\n\t")) || ((1.5 + 3.14159) <= (0.0 / float32((1e10 - (-2.25 - -2.25)))))) {
            v42 := GenS4{F0: (p1 * genFn0()), F1: -90}
            p1 += (^(p1))
            var v43 bool = false
            _ = v42
            _ = v43
        } else {
            fmt.Println(p0)
            if (-2.25 < (3.14159 - (0.1 + 0.0))) {
                v44 := GenS6{F0: 32767}
                v45 := GenS7{F0: "a", F1: -1, F2: 1000}
                v46 := []int{0, 32, int(p0), (p1 | p1)}
                _ = v44
                _ = v45
                _ = v46
            } else {
                var v47 int16 = 42
                for fi48 := 0; fi48 < 20; fi48++ {
                    p0 += p0
                    if true {
                        v49 := GenS8{F0: 0, F1: -717, F2: 150}
                        _ = v49
                    } else {
                        v50 := GenS8{F0: 109, F1: -186, F2: 471}
                        var v51 string = "foo bar"
                        fmt.Println(v41)
                        _ = v50
                        _ = v51
                    }
                    _ = fi48
                }
                if (18446744073709551615 <= 18446744073709551615) {
                    v52 := GenS5{F0: 193, F1: 42}
                    p1 = 791
                    _ = v52
                }
                _ = v47
            }
            var v53 uint8 = 171
            _ = v53
        }
        _ = fi40
        _ = v41
    }
    v54 := GenS4{F0: (p1 >> ((18446744073709551615) & 63)), F1: -1}
    _ = p0
    _ = p1
    _ = v54
    return 294
}

func main() {
    var v1 int8 = -98
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 int8 = (^(v1))
    if (genFn1(38, 623) < 121) {
        v55 := GenS9{F0: 42}
        v56 := []int{32767, 127, 2147483647, 230}
        if (!false) {
            fmt.Println(v4, v1)
            fmt.Println(v4, v1)
        }
        _ = v55
        _ = v56
    } else {
        if (0.0 > (1.5 - ((float32(float64(1e10)) * float32(0.1)) * (-2.25 - 0.0)))) {
            v3[951] = "hello"
            v4 -= v4
        }
        for fi57 := 0; fi57 < 8; fi57++ {
            v1 = 91
            _ = fi57
        }
    }
    fmt.Println(v1)
    switch 2147483647 {
    case 127:
        fmt.Println(v1, v1)
    case 42:
        v4 = (v4 | -1)
    default:
        var v58 int8 = (v4 % ((v4) | 1))
        _ = v58
    }
    v59 := GenS9{F0: 536}
    v60 := v2[-2147483648]
    var v61 uint64 = 583
    v62 := []int{(v60 * (v60 ^ (v60 << ((1) & 63)))), 32767, v60, int(v61)}
    var v63 float64 = float64(-2.25)
    if false {
        for fi64 := 0; fi64 < 1; fi64++ {
            var v65 int32 = -1
            var v66 uint8 = 42
            var v67 float32 = ((0.1 - 1e10) + (3.14159 / 1.0))
            _ = fi64
            _ = v65
            _ = v66
            _ = v67
        }
    } else {
        v61 = 682
        for fi68 := 0; fi68 < 4; fi68++ {
            for fi69 := 0; fi69 < 11; fi69++ {
                v70 := GenS5{F0: 2, F1: 520}
                v63 -= v63
                _ = fi69
                _ = v70
            }
            var v71 string = "\n\t"
            _ = fi68
            _ = v71
        }
        if false {
            var v72 uint = 824
            _ = v72
        } else {
            v73 := GenS10{F0: 42, F1: (v1 | -115)}
            v1 = v1
            var v74 int8 = 42
            _ = v73
            _ = v74
        }
    }
    v75 := v2[v60]
    var v76 uint = 4294967295
    v4 += (-(v4))
    if (!true) {
        delete(v2, 128)
        v4 -= (v1 | v1)
        var v77 float32 = ((((3.14159 + float32(v63)) / -2.25) * (3.14159 + -2.25)) / float32((-2.25 - v63)))
        _ = v77
    } else {
        _ = len(v62)
        v78 := GenS9{F0: 0}
        _ = v78
    }
    for fi79 := 0; fi79 < 7; fi79++ {
        for fi80 := 0; fi80 < 19; fi80++ {
            switch int(v61) {
            case -644:
                var v81 float64 = float64(3.14159)
                _ = v81
            case -9223372036854775808:
                v63 = v63
            case 566:
                v82 := GenS11{F0: -9223372036854775808, F1: v1, F2: 1}
                _ = v82
            default:
                if true {
                    v76 = genFn1(3, genFn0())
                    switch (v60 | (v75 & (v75 | fi79))) {
                    case 72:
                        v83 := GenS7{F0: (("hello" + "") + ("a" + (((("hello" + "hello") + "hello") + ("\n\t" + "a")) + ""))), F1: -635, F2: 1}
                        _ = v83
                    case 814:
                        if true {
                            var v84 float32 = (float32(v63) / (-2.25 / ((-2.25 * 3.14159) + -2.25)))
                            v85 := GenS9{F0: 17}
                            _ = v84
                            _ = v85
                        } else {
                            if false {
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    var v86 float32 = 0.1
                                    _ = v86
                                case "foo bar":
                                    if len(v62) >= 2 {
                                        v62 = v62[1:len(v62)]
                                    }
                                }
                                v87 := GenS2{F0: 2147483647}
                                var v88 string = ""
                                _ = v87
                                _ = v88
                            }
                            _ = len(v3)
                        }
                    case -1:
                        fmt.Println(v60)
                    }
                }
                _ = len(v3)
                delete(v3, 2147483648)
                var v89 string = (("foo bar" + "\n\t") + "foo bar")
                v90 := GenS11{F0: 9223372036854775807, F1: 0, F2: 0}
                _ = v89
                _ = v90
            }
            _ = fi80
        }
        fmt.Println(v63, v60, v63)
        fmt.Println(v75)
        _ = fi79
    }
    v91 := GenS1{F0: false, F1: -1}
    switch (v60 ^ (v75 / ((9223372036854775807) | 1))) {
    case -1:
        if (!(175 > 138)) {
            if len(v62) >= 2 {
                v62 = v62[1:len(v62)]
            }
        } else {
            var v92 int32 = 565
            switch ("a" + "hello") {
            case "foo bar":
                switch 2147483647 {
                case 0:
                    if (42 != 127) {
                        v93 := GenS12{F0: v76, F1: (v92 * (v92 | (v92 | v92)))}
                        fmt.Println(v76, v61, v1)
                        _ = v93
                    }
                case 2147483647:
                    for fi94 := 0; fi94 < 17; fi94++ {
                        var v95 int = v75
                        var v96 int64 = -117
                        _ = fi94
                        _ = v95
                        _ = v96
                    }
                }
            case "":
                v2[v75] = 1
            }
            _ = v92
        }
    case 663:
        fmt.Println(v1, v4, v63)
    default:
        fmt.Println(v75)
    }
    fmt.Println(v4, v4, v60)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    fmt.Println("v63:", v63)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v75
    _ = v76
    _ = v91
}
