package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int64; F1 uint; F2 uint64 }
type GenS7 struct { F0 bool; F1 uint16; F2 string }
type GenS8 struct { F0 int; F1 bool; F2 int16 }
type GenS9 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS10 struct { F0 int; F1 int16; F2 uint8 }
type GenS11 struct { F0 int; F1 int64 }
type GenS12 struct { F0 int16 }
func genFn0(p0 int16, p1 int64) int32 {
    p1 += 1
    v4 := []int{985, int(p0), 128}
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1(p0 int8, p1 int64) int8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint = 2
        v7 := []int{fi5, (-(fi5)), (fi5 % ((fi5) | 1))}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p1, p1)
    switch "" {
    case "hello":
        v8 := []int{32767}
        _ = v8
    case "foo bar":
        v9 := []int{9223372036854775807, 4294967295, 32768, 9223372036854775807}
        _ = v9
    }
    switch 4294967295 {
    case 32767:
        var v10 bool = false
        _ = v10
    case 42:
        fmt.Println(p0, p0)
    default:
        fmt.Println(p1, p0)
    }
    p1 += p1
    _ = p0
    _ = p1
    return p0
}

func main() {
    v1 := []string{"a", (("a" + (("\n\t" + "hello") + (("foo bar" + "a") + (("\n\t" + "foo bar") + "a")))) + "hello"), (("" + (("\n\t" + "a") + "\n\t")) + ((("" + "foo bar") + "foo bar") + ""))}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    for fi11 := 0; fi11 < 9; fi11++ {
        if (false && (false || true)) {
            if true {
                if (true && true) {
                    v12 := v3[2147483647]
                    for fi13 := 0; fi13 < 18; fi13++ {
                        if false {
                            if true {
                                fmt.Println(fi13, v12, v12)
                                v12 = v12
                            }
                            switch (-(fi13)) {
                            case 9223372036854775807:
                                var v14 uint16 = 255
                                _ = v14
                            case 1:
                                if true {
                                    if true {
                                        switch fi13 {
                                        case -660:
                                            if (42 >= fi13) {
                                                var v15 uint8 = 0
                                                _ = v15
                                            }
                                        case 127:
                                            var v16 uint64 = 4294967295
                                            _ = v16
                                        default:
                                            v12 = v12
                                        }
                                    } else {
                                        if false {
                                            var v17 uint32 = 42
                                            v18 := GenS1{F0: 2, F1: 2, F2: fi13}
                                            if false {
                                                fmt.Println(fi11, v17, fi13)
                                                fmt.Println(v12, fi11, v12)
                                            } else {
                                                switch (v12 + "a") {
                                                case "foo bar":
                                                    var v19 uint64 = 2
                                                    _ = v19
                                                case "\n\t":
                                                    switch (v12 + v12) {
                                                    case "\n\t":
                                                        v20 := GenS2{F0: v12, F1: v17}
                                                        _ = v20
                                                    case "foo bar":
                                                        if (128 > uint8(v17)) {
                                                            v12 = v12
                                                            v17 = 2147483648
                                                        } else {
                                                            v17 = v17
                                                        }
                                                        v17 -= 623
                                                    case "a":
                                                        fmt.Println(v12, v17)
                                                    }
                                                }
                                                v21 := GenS3{F0: 4294967295, F1: -1}
                                                if ((8 > 255) || true) {
                                                    var v22 float64 = (0.0 / (((float64(0.1) - 3.14159) + float64(3.14159)) - float64(1e10)))
                                                    if true {
                                                        var v23 uint16 = 65535
                                                        var v24 uint64 = 1
                                                        _ = v23
                                                        _ = v24
                                                    } else {
                                                        var v25 uint8 = 128
                                                        if true {
                                                            if (fi11 != fi13) {
                                                                v26 := GenS4{F0: 395}
                                                                _ = v26
                                                            }
                                                            v27 := GenS2{F0: v12, F1: v17}
                                                            if ((v17 - (v17 + (v17 & uint32(v25)))) >= v17) {
                                                                v17 = 2147483648
                                                                fmt.Println(v22, v17, v17)
                                                                v12 = "\n\t"
                                                                v28 := GenS4{F0: -174}
                                                                _ = v28
                                                            }
                                                            _ = v27
                                                        } else {
                                                            switch v12 {
                                                            case "\n\t":
                                                                fmt.Println(v25, fi13)
                                                            case "":
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    switch v12 {
                                                                    case "foo bar":
                                                                        v29 := GenS2{F0: v12, F1: 1}
                                                                        _ = v29
                                                                    case "\n\t":
                                                                        var v30 uint64 = 42
                                                                        _ = v30
                                                                    }
                                                                case "":
                                                                    if false {
                                                                        v12 = (v12 + v12)
                                                                        if true {
                                                                            switch v12 {
                                                                            case "\n\t":
                                                                                if (0 <= 65535) {
                                                                                    var v31 uint16 = 923
                                                                                    switch -9223372036854775808 {
                                                                                    case 32768:
                                                                                        var v32 float32 = float32(1.5)
                                                                                        _ = v32
                                                                                    case -523:
                                                                                        fmt.Println(fi11)
                                                                                    default:
                                                                                        switch v12 {
                                                                                        case "hello":
                                                                                            v33 := GenS3{F0: 0, F1: -128}
                                                                                            _ = v33
                                                                                        case "a":
                                                                                            v34 := GenS4{F0: -810}
                                                                                            _ = v34
                                                                                        }
                                                                                    }
                                                                                    _ = len(v1)
                                                                                    _ = v31
                                                                                }
                                                                                v25 = (v25 / ((v25) | 1))
                                                                                var v35 int8 = -128
                                                                                _ = v35
                                                                            case "hello":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ""
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v36 int = fi11
                                                                        _ = v36
                                                                    }
                                                                case "hello":
                                                                    if (float64(((0.0 / 3.14159) / ((0.1 + (0.0 + 0.1)) - 1e10))) <= 3.14159) {
                                                                        v12 = v12
                                                                        if false {
                                                                            v3[(v17 ^ v17)] = v12
                                                                        } else {
                                                                            if true {
                                                                                v22 += v22
                                                                                if false {
                                                                                    v22 = float64((0.0 / 1.5))
                                                                                    var v37 int64 = 42
                                                                                    fmt.Println(v25, v25, fi13)
                                                                                    var v38 int = fi11
                                                                                    var v39 float64 = ((((v22 / (v22 - (1.5 - 1e10))) - (v22 / 1.5)) * (0.1 - 3.14159)) + v22)
                                                                                    _ = v37
                                                                                    _ = v38
                                                                                    _ = v39
                                                                                }
                                                                                var v40 bool = (!true)
                                                                                _ = v40
                                                                            } else {
                                                                                v25 -= (v25 - 128)
                                                                                v22 += (v22 * (float64(0.1) + v22))
                                                                            }
                                                                            if (-30 != -110) {
                                                                                fmt.Println(v22, fi11, v17)
                                                                                if ((v22 / 3.14159) < (1e10 + v22)) {
                                                                                    v17 += v17
                                                                                    v25 = (v25 | v25)
                                                                                } else {
                                                                                    v22 -= float64(float32((v22 * (((((v22 + v22) * (1.5 / (1.5 / ((1.5 + 0.1) + (0.1 + (v22 - (0.1 + float64(-2.25)))))))) - v22) - v22) - v22))))
                                                                                    v22 = 0.0
                                                                                }
                                                                                var v41 float32 = (0.0 + -2.25)
                                                                                _ = v41
                                                                            }
                                                                            fmt.Println(v22, fi11, v25)
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (v12 + v12)
                                                                        }
                                                                    }
                                                                }
                                                            case "hello":
                                                                if (-833 >= 1) {
                                                                    _ = len(v2)
                                                                } else {
                                                                    var v42 int8 = 80
                                                                    v12 = v12
                                                                    v42 = 84
                                                                    _ = v42
                                                                }
                                                                if false {
                                                                    fmt.Println(v12, v25)
                                                                }
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = ("\n\t" + v12)
                                                            }
                                                        }
                                                        var v43 int = 128
                                                        var v44 bool = (!false)
                                                        if true {
                                                            _ = len(v2)
                                                        } else {
                                                            switch "" {
                                                            case "foo bar":
                                                                v43 = -613
                                                            default:
                                                                v1 = append(v1, v12)
                                                            }
                                                            v45 := GenS4{F0: -722}
                                                            v46 := v2[17]
                                                            fmt.Println(v17, fi13)
                                                            _ = len(v1)
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        _ = v25
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    _ = v22
                                                }
                                                _ = v21
                                            }
                                            _ = v17
                                            _ = v18
                                        } else {
                                            var v47 uint16 = 958
                                            v48 := GenS5{F0: (-(fi11)), F1: -9223372036854775808}
                                            _ = v47
                                            _ = v48
                                        }
                                        v49 := GenS6{F0: -213, F1: 1, F2: 1}
                                        _ = v49
                                    }
                                } else {
                                    var v50 string = v12
                                    v51 := GenS7{F0: true, F1: 338, F2: v50}
                                    _ = v50
                                    _ = v51
                                }
                            default:
                                var v52 int = 128
                                _ = v52
                            }
                            if false {
                                if false {
                                    v12 = ("\n\t" + ((v12 + "foo bar") + "\n\t"))
                                    v12 = "a"
                                    fmt.Println(fi13)
                                    var v53 uint16 = 2
                                    v12 = "\n\t"
                                    if (1 != (fi11 / (((fi11 / ((fi13) | 1))) | 1))) {
                                        if true {
                                            v12 = "\n\t"
                                            v53 = 323
                                            switch "a" {
                                            case "hello":
                                                v53 -= (v53 & (v53 - v53))
                                            case "foo bar":
                                                v54 := GenS8{F0: 42, F1: true, F2: -1}
                                                _ = v54
                                            case "\n\t":
                                                if false {
                                                    if true {
                                                        v53 = v53
                                                        v55 := GenS2{F0: v12, F1: 63}
                                                        _ = v55
                                                    }
                                                    var v56 uint = 18446744073709551615
                                                    fmt.Println(fi11)
                                                    _ = v56
                                                } else {
                                                    v53 -= v53
                                                    fmt.Println(v12)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            default:
                                                var v57 float64 = 1.5
                                                _ = v57
                                            }
                                            v58 := GenS6{F0: 324, F1: 2, F2: 370}
                                            _ = v58
                                        }
                                        v59 := GenS5{F0: fi11, F1: 9223372036854775807}
                                        _ = v59
                                    }
                                    _ = v53
                                } else {
                                    var v60 bool = true
                                    var v61 int64 = -266
                                    _ = v60
                                    _ = v61
                                }
                                v62 := GenS1{F0: 127, F1: 4294967295, F2: fi11}
                                var v63 bool = true
                                _ = v62
                                _ = v63
                            } else {
                                v64 := GenS1{F0: 128, F1: 523, F2: (-(fi11))}
                                var v65 float32 = 0.1
                                v65 = float32((3.14159 - 0.0))
                                _ = v64
                                _ = v65
                            }
                        } else {
                            v66 := GenS9{F0: 199, F1: 0, F2: 65535}
                            _ = v66
                        }
                        v12 = v12
                        fmt.Println(fi13, fi11)
                        _ = fi13
                    }
                    v12 = v12
                    _ = v12
                }
            } else {
                v67 := GenS6{F0: -701, F1: 141, F2: 153}
                var v68 int32 = 0
                _ = v67
                _ = v68
            }
            v69 := v2[0]
            v69 = (float32((0.0 * (0.1 / 3.14159))) >= 1.5)
            _ = v69
        } else {
            if true {
                var v70 uint16 = 941
                var v71 uint = 4294967295
                if (2 > 1) {
                    if (false || true) {
                        v72 := GenS2{F0: "", F1: 0}
                        v2[1] = (genFn1(100, 0) != 0)
                        fmt.Println(fi11, v71)
                        _ = v72
                    } else {
                        v71 += uint(v70)
                    }
                    fmt.Println(fi11, fi11)
                } else {
                    switch ((((("\n\t" + ("\n\t" + ("a" + ""))) + ("hello" + "hello")) + ("foo bar" + "")) + "") + "foo bar") {
                    case "a":
                        v71 = v71
                    case "":
                        if true {
                            v73 := GenS6{F0: 205, F1: (^(v71)), F2: 0}
                            fmt.Println(v70)
                            _ = v73
                        } else {
                            v70 = 483
                            for fi74 := 0; fi74 < 18; fi74++ {
                                var v75 int8 = -117
                                fmt.Println(fi74, v75, v71)
                                _ = fi74
                                _ = v75
                            }
                        }
                    case "\n\t":
                        v71 = 4294967295
                    }
                }
                _ = v70
                _ = v71
            }
        }
        var v76 int8 = -128
        var v77 uint = 690
        _ = fi11
        _ = v76
        _ = v77
    }
    var v78 float32 = ((((-2.25 + 0.1) - 1e10) / -2.25) * 3.14159)
    if true {
        v78 = v78
        v78 += float32(0.0)
    }
    fmt.Println(v78)
    switch ("hello" + ("" + (("foo bar" + ("foo bar" + (("foo bar" + "\n\t") + "hello"))) + ""))) {
    case "hello":
        v1 = append(v1, "\n\t")
    case "a":
        v79 := GenS9{F0: 0, F1: 0, F2: 2}
        _ = v79
    case "foo bar":
        if ((42 != -556) || (9223372036854775807 < 42)) {
            v80 := v2[4294967295]
            for fi81 := 0; fi81 < 14; fi81++ {
                v80 = (338 == -484)
                v78 = v78
                _ = fi81
            }
            for fi82 := 0; fi82 < 3; fi82++ {
                if false {
                    if (!(!(true && v80))) {
                        var v83 string = ("foo bar" + "a")
                        var v84 bool = v80
                        v83 = (v83 + v83)
                        _ = v83
                        _ = v84
                    }
                    if v80 {
                        v80 = (fi82 != 127)
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v85 int8 = 42
                        var v86 string = "a"
                        _ = v85
                        _ = v86
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    v87 := GenS10{F0: (fi82 & (fi82 + 2147483647)), F1: 32767, F2: 128}
                    v80 = (float32((-2.25 * 0.0)) < v78)
                    _ = v87
                }
                v78 += v78
                _ = fi82
            }
            _ = v80
        } else {
            var v88 int64 = 1
            var v89 int64 = (v88 % ((1) | 1))
            _ = v88
            _ = v89
        }
    }
    if false {
        for fi90 := 0; fi90 < 19; fi90++ {
            v91 := GenS10{F0: (fi90 * 1), F1: -32768, F2: 127}
            v78 = v78
            if (1.5 > v78) {
                if len(v1) > 0 {
                    v1[0] = ""
                }
            } else {
                v92 := GenS11{F0: -1, F1: 42}
                v78 += float32(0.0)
                _ = v92
            }
            _ = fi90
            _ = v91
        }
        var v93 int16 = 0
        v94 := GenS12{F0: (-(v93))}
        _ = v93
        _ = v94
    } else {
        v95 := GenS8{F0: -2147483648, F1: false, F2: 976}
        _ = v95
    }
    for fi96 := 0; fi96 < 13; fi96++ {
        var v97 uint16 = 964
        _ = fi96
        _ = v97
    }
    _ = len(v1)
    v78 += v78
    var v98 float64 = (0.0 + float64((((-2.25 + (v78 - v78)) + v78) - v78)))
    v98 += (v98 - v98)
    var v99 uint16 = 0
    var v100 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v78:", v78)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v78
    _ = v98
    _ = v99
}
