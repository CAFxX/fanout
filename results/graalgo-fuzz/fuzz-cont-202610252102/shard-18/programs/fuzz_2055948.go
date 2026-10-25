package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint64; F2 int }
type GenS4 struct { F0 uint16; F1 uint32; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 int; F2 int16 }
type GenS8 struct { F0 int32 }
func genFn0(p0 string) bool {
    v6 := GenS1{F0: 42, F1: 42, F2: ("foo bar" + ("" + (p0 + p0)))}
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 100, F1: fi7, F2: p0}
            var v10 int64 = -9223372036854775808
            if true {
                p0 = (p0 + (p0 + "hello"))
                p0 = p0
                fmt.Println(p0)
            } else {
                if true {
                    var v11 uint16 = 0
                    _ = v11
                }
                if false {
                    if true {
                        v12 := GenS1{F0: 0, F1: 32768, F2: ((p0 + "\n\t") + "a")}
                        if (3.14159 < 0.1) {
                            v10 = v10
                        }
                        fmt.Println(v10, fi8, p0)
                        var v13 int64 = (v10 - (v10 | -739))
                        var v14 uint64 = 745
                        var v15 bool = (p0 > ((p0 + ("" + "hello")) + ((("foo bar" + p0) + (("\n\t" + (p0 + ((p0 + "a") + p0))) + "foo bar")) + "")))
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    v16 := []int{fi8, fi8, fi8, int(v10)}
                    v10 -= v10
                    _ = v16
                } else {
                    v17 := []int{(fi8 - fi8), fi8, (fi8 | 127)}
                    v10 += v10
                    switch fi8 {
                    case 0:
                        if true {
                            var v18 int = fi8
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            var v19 int = (fi8 * (fi8 ^ (fi8 >> ((1) & 63))))
                            v10 = (^(v10))
                            _ = v18
                            _ = v19
                        }
                    default:
                        if false {
                            v10 = 9223372036854775807
                            v10 += (-(v10))
                        } else {
                            fmt.Println(fi7, fi7, v10)
                            p0 = p0
                        }
                    }
                    _ = v17
                }
                p0 = (p0 + p0)
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    p0 = p0
    p0 = (("a" + "hello") + p0)
    fmt.Println(p0)
    _ = p0
    _ = v6
    return true
}
func genFn1(p0 int32, p1 uint32, p2 uint8) uint {
    p1 -= p1
    p1 -= p1
    fmt.Println(p1, p0)
    fmt.Println(p2)
    if false {
        for fi20 := 0; fi20 < 6; fi20++ {
            if false {
                var v21 uint64 = 351
                if ((42 < p1) || (36 == v21)) {
                    v22 := GenS1{F0: 1, F1: fi20, F2: (("hello" + "a") + ("hello" + ("a" + "foo bar")))}
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if false {
                            p2 = (^(p2))
                            var v24 uint32 = 42
                            p2 += (p2 | p2)
                            if ("a" >= "foo bar") {
                                if true {
                                    var v25 uint8 = 136
                                    _ = v25
                                }
                                p2 += 2
                                fmt.Println(p2, p1)
                                v26 := []int{4294967295, (fi23 | (fi20 % (((fi20 & (fi20 - (fi20 << ((1) & 63))))) | 1))), fi23, -9223372036854775808}
                                _ = v26
                            }
                            _ = v24
                        } else {
                            var v27 int64 = 65
                            switch (fi23 >> ((uint(p2)) & 63)) {
                            case 127:
                                if (-24 != 0) {
                                    var v28 uint8 = p2
                                    p1 = (p1 >> ((976) & 63))
                                    _ = v28
                                } else {
                                    if (!false) {
                                        switch fi23 {
                                        case -1:
                                            var v29 uint32 = (^(p1))
                                            _ = v29
                                        case -2147483648:
                                            switch "" {
                                            case "\n\t":
                                                var v30 int32 = 2147483647
                                                _ = v30
                                            default:
                                                var v31 uint8 = (p2 % ((p2) | 1))
                                                _ = v31
                                            }
                                        default:
                                            var v32 int32 = (p0 >> ((865) & 63))
                                            _ = v32
                                        }
                                        v33 := GenS1{F0: -128, F1: (fi20 ^ fi20), F2: "foo bar"}
                                        _ = v33
                                    }
                                }
                            case -9223372036854775808:
                                fmt.Println(p2, v27, p1)
                            case 412:
                                v34 := []int{(fi23 + (fi20 + (fi20 | fi20)))}
                                _ = v34
                            default:
                                var v35 bool = true
                                _ = v35
                            }
                            _ = v27
                        }
                        v21 += (v21 << ((929) & 63))
                        p2 -= 228
                        _ = fi23
                    }
                    _ = v22
                }
                _ = v21
            }
            v36 := GenS1{F0: 100, F1: (-(fi20)), F2: "a"}
            _ = fi20
            _ = v36
        }
        p0 = -1
        v37 := []int{-892, 9223372036854775807}
        _ = v37
    } else {
        p2 = p2
        v38 := GenS1{F0: -128, F1: 127, F2: ("hello" + ("foo bar" + "a"))}
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := []int{fi39, fi39, fi39}
            if (882 >= 730) {
                for fi41 := 0; fi41 < 15; fi41++ {
                    v42 := GenS2{F0: "hello"}
                    v43 := GenS3{F0: (^(fi39)), F1: 859, F2: (fi39 & 4294967295)}
                    if (42 < 0) {
                        var v44 uint16 = 255
                        var v45 uint64 = 2
                        var v46 int64 = -94
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = len(v40)
            }
            if len(v40) >= 2 {
                v40 = v40[1:len(v40)]
            }
            _ = fi39
            _ = v40
        }
        _ = v38
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 100
    var v3 uint16 = 32767
    var v4 bool = true
    var v5 uint = 0
    if v4 {
        v47 := []int{0, 32768, 32768, 972}
        _ = v47
    }
    switch (v1 + v1) {
    case "\n\t":
        switch ("foo bar" + ((("" + ("a" + ((v1 + (((v1 + "") + v1) + v1)) + v1))) + ("\n\t" + (v1 + v1))) + v1)) {
        case "\n\t":
            v4 = (198 == 0)
        case "a":
            v48 := []int{}
            _ = v48
        case "hello":
            switch (v1 + "") {
            case "a":
                var v49 int16 = 42
                _ = v49
            }
        default:
            if (451 > 95) {
                v50 := []int{32768, 127, 397}
                fmt.Println(v2, v3)
                _ = v50
            } else {
                var v51 uint16 = v3
                v1 = v1
                for fi52 := 0; fi52 < 1; fi52++ {
                    v53 := []int{32768, fi52, (fi52 >> ((genFn1(1, 240, 127)) & 63))}
                    if len(v53) >= 2 {
                        v53 = v53[1:len(v53)]
                    }
                    _ = fi52
                    _ = v53
                }
                _ = v51
            }
        }
    case "a":
        v54 := []int{2147483647}
        _ = v54
    case "":
        if ((1 < 1) || true) {
            switch ((v1 + (("a" + v1) + v1)) + v1) {
            case "a":
                if (false || (true && v4)) {
                    v5 = 2
                }
            case "hello":
                fmt.Println(v4, v4, v2)
            case "\n\t":
                v55 := GenS4{F0: (v3 / ((255) | 1)), F1: 4294967295, F2: v4}
                _ = v55
            }
            v1 = v1
            for fi56 := 0; fi56 < 9; fi56++ {
                if (v4 || ((!(v4 || (((!v4) || v4) && v4))) && ((1.5 / (((1e10 - (3.14159 * 0.1)) / (1e10 / 3.14159)) * 0.1)) <= (-2.25 * (float32(float64(float32((0.0 * 1e10)))) / 1e10))))) {
                    v57 := GenS2{F0: (v1 + "a")}
                    v2 = v2
                    v58 := GenS1{F0: (-(v2)), F1: (fi56 / ((fi56) | 1)), F2: ((v1 + (v1 + (v1 + ("" + v1)))) + v1)}
                    _ = v57
                    _ = v58
                }
                var v59 int8 = v2
                _ = fi56
                _ = v59
            }
        } else {
            v60 := GenS5{F0: 221}
            _ = v60
        }
    default:
        v61 := []int{32767, 32768}
        _ = v61
    }
    switch -853 {
    case 32768:
        switch 9223372036854775807 {
        case 4294967295:
            switch -795 {
            case -9223372036854775808:
                fmt.Println(v5, v4)
            case -2147483648:
                v62 := []int{4294967295, 42, -6, 127}
                _ = v62
            case 32768:
                switch 737 {
                case -2147483648:
                    v63 := []int{32767, 2147483647, 938}
                    _ = v63
                default:
                    for fi64 := 0; fi64 < 19; fi64++ {
                        v65 := GenS6{F0: v5}
                        _ = fi64
                        _ = v65
                    }
                }
            default:
                fmt.Println(v1)
            }
        case -556:
            v66 := GenS4{F0: (v3 % ((255) | 1)), F1: 2, F2: ((v5 % ((genFn1(-1, 42, 2)) | 1)) >= 18446744073709551615)}
            _ = v66
        case 9223372036854775807:
            v67 := GenS5{F0: 1}
            _ = v67
        }
    case 2147483647:
        switch 0 {
        case -195:
            for fi68 := 0; fi68 < 15; fi68++ {
                var v69 uint16 = v3
                v70 := []int{(^(fi68)), 1, 319, fi68}
                _ = fi68
                _ = v69
                _ = v70
            }
        case 32768:
            for fi71 := 0; fi71 < 12; fi71++ {
                v72 := []int{}
                var v73 int32 = 346
                _ = fi71
                _ = v72
                _ = v73
            }
        default:
            switch (v1 + ("a" + ((v1 + ("a" + ("" + "\n\t"))) + v1))) {
            case "a":
                var v74 string = ((((v1 + (("hello" + ("a" + v1)) + "a")) + (v1 + ((v1 + (v1 + (v1 + ("" + ("a" + ""))))) + v1))) + ("\n\t" + "\n\t")) + ((v1 + "\n\t") + ("foo bar" + v1)))
                _ = v74
            default:
                v2 = v2
            }
        }
    case 128:
        var v75 uint16 = v3
        _ = v75
    default:
        for fi76 := 0; fi76 < 9; fi76++ {
            if v4 {
                if (v4 && v4) {
                    v1 = "a"
                } else {
                    v77 := GenS6{F0: v5}
                    _ = v77
                }
                if v4 {
                    if true {
                        if (169 <= -774) {
                            switch int(v5) {
                            case 127:
                                v78 := GenS7{F0: 866, F1: (fi76 >> ((v5) & 63)), F2: -1}
                                _ = v78
                            }
                        }
                        if ("\n\t" == v1) {
                            var v79 int32 = 350
                            v1 = "foo bar"
                            _ = v79
                        }
                        var v80 uint8 = 1
                        _ = v80
                    } else {
                        v81 := GenS8{F0: -1}
                        _ = v81
                    }
                }
            }
            _ = fi76
        }
    }
    fmt.Println(v4, v2, v5)
    for fi82 := 0; fi82 < 13; fi82++ {
        var v83 uint8 = 139
        for fi84 := 0; fi84 < 18; fi84++ {
            fmt.Println(v1, v2, v83)
            v5 -= uint(v3)
            v85 := []int{fi82, int(v2), fi82, 127}
            _ = fi84
            _ = v85
        }
        v86 := []int{(fi82 * fi82), (^(fi82))}
        _ = fi82
        _ = v83
        _ = v86
    }
    switch v1 {
    case "\n\t":
        var v87 uint8 = 0
        _ = v87
    case "foo bar":
        var v88 int64 = -376
        _ = v88
    }
    fmt.Println(v2, v3, v3)
    for fi89 := 0; fi89 < 19; fi89++ {
        v4 = v4
        var v90 bool = (-1 >= 0)
        _ = fi89
        _ = v90
    }
    v91 := GenS4{F0: (v3 + 448), F1: 685, F2: false}
    v5 = (v5 & (^(v5)))
    v4 = (726 >= 1)
    v92 := []int{9223372036854775807, 127}
    v2 -= v2
    switch 579 {
    case 9223372036854775807:
        if v4 {
            v92 = append(v92, -2147483648)
            var v93 int16 = 0
            _ = v93
        }
    case -365:
        fmt.Println(v1)
    }
    var v94 int64 = 781
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v92:", len(v92), v92)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
    _ = v92
    _ = v94
}
