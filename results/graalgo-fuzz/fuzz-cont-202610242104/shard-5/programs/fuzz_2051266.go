package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 int8; F2 int }
type GenS9 struct { F0 string; F1 uint8; F2 int }
type GenS10 struct { F0 int16 }
func genFn0() uint64 {
    if (!true) {
        for fi4 := 0; fi4 < 9; fi4++ {
            var v5 uint = 894
            var v6 float64 = float64(1.5)
            _ = fi4
            _ = v5
            _ = v6
        }
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := []int{42}
                var v10 float32 = (float32(float64(1e10)) - 3.14159)
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint32 = 72
            fmt.Println(v11, fi7)
            _ = fi7
            _ = v11
        }
        if true {
            var v12 uint64 = 2
            for fi13 := 0; fi13 < 10; fi13++ {
                if (32767 < 255) {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 float32 = 3.14159
                        v16 := GenS1{F0: false, F1: 128, F2: 65535}
                        var v17 string = (("hello" + (("foo bar" + (("foo bar" + ("\n\t" + ("foo bar" + (("a" + "\n\t") + ((("foo bar" + "a") + "foo bar") + ("" + ("" + ""))))))) + "\n\t")) + "foo bar")) + ("foo bar" + "a"))
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 11; fi18++ {
                        switch "a" {
                        case "\n\t":
                            var v19 uint8 = 184
                            _ = v19
                        default:
                            var v20 uint32 = 4294967295
                            _ = v20
                        }
                        if false {
                            if (false || (-2.25 < (-2.25 / 0.1))) {
                                v12 += v12
                            }
                        } else {
                            switch ("foo bar" + ("\n\t" + "foo bar")) {
                            case "a":
                                if (-133 >= -9223372036854775808) {
                                    if true {
                                        v12 = v12
                                    } else {
                                        v21 := GenS1{F0: (0 != 2147483647), F1: 2, F2: 18446744073709551615}
                                        v12 -= v12
                                        _ = v21
                                    }
                                    var v22 int16 = -164
                                    _ = v22
                                } else {
                                    v23 := GenS1{F0: (true && false), F1: 128, F2: 245}
                                    fmt.Println(v12)
                                    v24 := GenS2{F0: fi18, F1: 42}
                                    _ = v23
                                    _ = v24
                                }
                            case "\n\t":
                                if ((0 >= 65535) && (548 == 4294967295)) {
                                    v25 := []int{0}
                                    var v26 uint32 = 42
                                    _ = v25
                                    _ = v26
                                }
                            }
                            v12 += 185
                            fmt.Println(v12, fi18, fi18)
                            var v27 uint64 = (v12 + (v12 >> ((392) & 63)))
                            var v28 float64 = -2.25
                            if true {
                                var v29 uint16 = 32767
                                v30 := GenS3{F0: 42, F1: 127, F2: 65535}
                                var v31 uint16 = (v29 * 538)
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            if (238 <= 183) {
                                switch 9223372036854775807 {
                                case 4294967295:
                                    if false {
                                        var v32 int8 = 0
                                        _ = v32
                                    }
                                }
                            } else {
                                v33 := GenS4{F0: 42, F1: 42, F2: 1}
                                var v34 int = (fi13 / ((fi13) | 1))
                                var v35 uint8 = 47
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            _ = v27
                            _ = v28
                        }
                        if (false && (float32(1e10) < -2.25)) {
                            if false {
                                switch "foo bar" {
                                case "hello":
                                    v36 := []int{-26, fi18, fi13}
                                    _ = v36
                                case "foo bar":
                                    var v37 float64 = 1.5
                                    _ = v37
                                default:
                                    fmt.Println(fi13)
                                }
                            } else {
                                if true {
                                    var v38 int = 127
                                    var v39 uint = 65535
                                    fmt.Println(fi18)
                                    _ = v38
                                    _ = v39
                                } else {
                                    var v40 int64 = -9223372036854775808
                                    v41 := GenS5{F0: 1}
                                    _ = v40
                                    _ = v41
                                }
                                switch "" {
                                case "":
                                    if false {
                                        v42 := []int{(-(fi18))}
                                        v12 = (v12 - (v12 * v12))
                                        var v43 int8 = 1
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                switch (fi18 / (((fi18 << ((uint(v12)) & 63))) | 1)) {
                                case 453:
                                    if (42 < v12) {
                                        v44 := []int{fi13, (^(fi13)), (fi18 >> ((0) & 63))}
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                        var v45 int = 772
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v46 uint16 = 236
                                        v12 = v12
                                        switch fi18 {
                                        case 32767:
                                            v46 = (v46 / ((32767) | 1))
                                        case -306:
                                            v12 -= v12
                                            fmt.Println(v12, v12, fi13)
                                            switch (fi13 & fi18) {
                                            case 128:
                                                v46 = (v46 >> ((uint(v12)) & 63))
                                            case 391:
                                                fmt.Println(v46)
                                            case -334:
                                                v12 -= (v12 ^ (v12 / ((109) | 1)))
                                                v12 = (^(v12))
                                                v12 = v12
                                                v47 := GenS6{F0: (float32(0.0) > 0.1)}
                                                v48 := []int{(fi13 + fi13), (fi18 >> ((667) & 63)), (fi18 + (fi13 ^ int(v46))), fi13}
                                                var v49 uint32 = 4294967295
                                                if false {
                                                    v48 = append(v48, fi13)
                                                    fmt.Println(v49)
                                                    var v50 uint = 42
                                                    _ = v50
                                                } else {
                                                    fmt.Println(fi13)
                                                }
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            default:
                                                v12 += v12
                                                fmt.Println(v12)
                                            }
                                        default:
                                            var v51 int = (fi18 - fi13)
                                            _ = v51
                                        }
                                        _ = v46
                                    }
                                case 4294967295:
                                    if (!(false && ((v12 << ((18446744073709551615) & 63)) != 39))) {
                                        var v52 float32 = ((-2.25 * (-2.25 / (1.5 / (0.1 / 1e10)))) + 1e10)
                                        v53 := []int{(fi13 % (((fi13 / ((fi18) | 1))) | 1)), (-(fi13)), 127, (fi18 << ((139) & 63))}
                                        v52 += v52
                                        _ = v52
                                        _ = v53
                                    } else {
                                        v54 := []int{(^(fi18)), fi18}
                                        var v55 int8 = 1
                                        fmt.Println(v12)
                                        _ = v54
                                        _ = v55
                                    }
                                case 496:
                                    switch fi18 {
                                    case 32767:
                                        var v56 bool = (v12 > 1)
                                        _ = v56
                                    case 891:
                                        var v57 uint64 = (v12 / (((v12 | v12)) | 1))
                                        v58 := []int{fi18, fi13, (-(fi18)), 4294967295}
                                        _ = v57
                                        _ = v58
                                    default:
                                        v59 := []int{fi18}
                                        _ = v59
                                    }
                                }
                            }
                        }
                        _ = fi18
                    }
                    v60 := []int{(^(fi13)), 42}
                    _ = v60
                }
                v61 := GenS1{F0: false, F1: 0, F2: 4294967295}
                v62 := []int{(fi13 / ((fi13) | 1))}
                _ = fi13
                _ = v61
                _ = v62
            }
            _ = v12
        } else {
            var v63 float64 = (float64(3.14159) + (0.1 + float64(0.1)))
            _ = v63
        }
    } else {
        if true {
            switch 42 {
            case 9223372036854775807:
                var v64 float64 = 0.0
                _ = v64
            case 1:
                var v65 int32 = -351
                _ = v65
            case -518:
                var v66 uint64 = 18446744073709551615
                _ = v66
            default:
                if (0 > 13) {
                    if true {
                        var v67 int32 = 2147483647
                        v67 += v67
                        _ = v67
                    } else {
                        for fi68 := 0; fi68 < 12; fi68++ {
                            v69 := GenS4{F0: 255, F1: 0, F2: 1}
                            _ = fi68
                            _ = v69
                        }
                        var v70 uint = 42
                        _ = v70
                    }
                }
            }
            var v71 uint = 538
            v72 := []int{-1}
            _ = v71
            _ = v72
        }
        if (false || true) {
            v73 := []int{127, 4294967295}
            for fi74 := 0; fi74 < 4; fi74++ {
                fmt.Println(fi74)
                var v75 bool = false
                _ = len(v73)
                _ = fi74
                _ = v75
            }
            if false {
                switch 128 {
                case 4294967295:
                    var v76 uint64 = 42
                    _ = v76
                default:
                    var v77 float32 = (1e10 + (float32((float64(((0.0 * (1.5 * 1e10)) / (1.5 - 3.14159))) / 3.14159)) * -2.25))
                    _ = v77
                }
                _ = len(v73)
            }
            _ = v73
        } else {
            v78 := []int{42, -189, -1}
            var v79 bool = (0 < 0)
            var v80 float64 = float64(1.5)
            _ = v78
            _ = v79
            _ = v80
        }
    }
    v81 := GenS7{F0: 0}
    switch -692 {
    case -211:
        var v82 uint64 = 42
        _ = v82
    case 194:
        if (4294967295 < 2) {
            v83 := []int{4294967295}
            var v84 uint16 = 732
            _ = v83
            _ = v84
        }
    case 32767:
        var v85 uint32 = 1
        _ = v85
    default:
        if (!false) {
            v86 := GenS7{F0: 2147483647}
            _ = v86
        }
    }
    _ = v81
    return 1
}
func genFn1(p0 string) string {
    fmt.Println(p0)
    p0 = ((p0 + ("foo bar" + (p0 + ((p0 + (p0 + (((((p0 + p0) + (p0 + p0)) + "\n\t") + (("\n\t" + (("" + ((((("\n\t" + ((p0 + p0) + "")) + p0) + p0) + "\n\t") + "foo bar")) + p0)) + p0)) + p0))) + (p0 + p0))))) + ("foo bar" + (p0 + p0)))
    if false {
        if false {
            fmt.Println(p0)
        }
        for fi87 := 0; fi87 < 12; fi87++ {
            p0 = ("a" + "a")
            _ = fi87
        }
    }
    _ = p0
    return (p0 + p0)
}

func main() {
    v1 := []string{("a" + ("" + "hello")), (((("a" + "foo bar") + ("a" + "a")) + ("" + "foo bar")) + "foo bar"), "\n\t", "a"}
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]bool)
    for fi88 := 0; fi88 < 7; fi88++ {
        var v89 uint = 1
        v89 = (v89 / (((v89 ^ (^(v89)))) | 1))
        _ = fi88
        _ = v89
    }
    for fi90 := 0; fi90 < 20; fi90++ {
        v91 := GenS8{F0: ((int32(fi90) == -1) || true), F1: 0, F2: 870}
        _ = fi90
        _ = v91
    }
    if ((v2 & v2) <= 212) {
        fmt.Println(v2)
        var v92 string = "foo bar"
        v93 := v3[(("hello" + (((v92 + (v92 + "foo bar")) + v92) + (v92 + v92))) + ((("foo bar" + (v92 + v92)) + "foo bar") + (("a" + ((v92 + v92) + ((v92 + "hello") + "hello"))) + (v92 + ""))))]
        _ = v92
        _ = v93
    }
    var v94 uint8 = 127
    var v95 uint64 = 0
    v94 += v94
    for fi96 := 0; fi96 < 19; fi96++ {
        v2 += (v2 / (((v2 | 669)) | 1))
        v95 += genFn0()
        _ = fi96
    }
    if ((((-2.25 - 0.1) - 0.1) * 0.1) >= 3.14159) {
        v97 := GenS5{F0: 32767}
        var v98 int8 = 40
        v98 += v98
        _ = v97
        _ = v98
    } else {
        for fi99 := 0; fi99 < 20; fi99++ {
            for fi100 := 0; fi100 < 18; fi100++ {
                fmt.Println(v95)
                var v101 uint64 = v95
                if (!true) {
                    if true {
                        v95 -= (v95 | genFn0())
                    }
                    v2 -= v2
                } else {
                    v94 = (v94 * (^(v94)))
                    fmt.Println(v101, fi100)
                    switch fi99 {
                    case -1:
                        switch -445 {
                        case 1:
                            v102 := GenS9{F0: "", F1: v94, F2: 32768}
                            _ = v102
                        case 32767:
                            v1 = append(v1, "foo bar")
                        default:
                            v101 -= v101
                        }
                    case 141:
                        var v103 int = fi99
                        v104 := GenS10{F0: 1000}
                        v105 := GenS5{F0: 32767}
                        _ = v103
                        _ = v104
                        _ = v105
                    case 32768:
                        v2 = v2
                    }
                }
                _ = fi100
                _ = v101
            }
            switch (fi99 / ((1) | 1)) {
            case -2147483648:
                var v106 uint8 = v94
                _ = v106
            case 128:
                var v107 bool = ((!(1 < 802)) && (float64(1e10) == 1.5))
                _ = v107
            default:
                v94 += (^(v94))
            }
            if false {
                switch (^(fi99)) {
                case 1:
                    var v108 int = fi99
                    _ = v108
                default:
                    var v109 bool = ((v2 + v2) >= (v2 | (v2 - (v2 * 129))))
                    _ = v109
                }
            }
            _ = fi99
        }
    }
    v110 := GenS4{F0: 701, F1: 66, F2: 327}
    var v111 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v110
    _ = v111
    _ = v2
    _ = v3
    _ = v94
    _ = v95
}
