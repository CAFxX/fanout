package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint; F1 int16; F2 int8 }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }
type GenS5 struct { F0 int16; F1 int64; F2 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64; F1 int64 }
func genFn0(p0 int32, p1 string, p2 string) int {
    if true {
        v4 := GenS1{F0: (p2 + p1)}
        _ = v4
    }
    switch 128 {
    case -2147483648:
        for fi5 := 0; fi5 < 17; fi5++ {
            p1 = (("hello" + "foo bar") + p1)
            for fi6 := 0; fi6 < 12; fi6++ {
                v7 := GenS1{F0: p2}
                switch fi6 {
                case 42:
                    if (0 <= 65535) {
                        if (-1 == 1) {
                            p0 = 0
                            var v8 int32 = (p0 >> ((2) & 63))
                            var v9 string = (p2 + "hello")
                            _ = v8
                            _ = v9
                        } else {
                            switch -1 {
                            case -104:
                                if (121 >= 0) {
                                    v10 := GenS1{F0: p1}
                                    var v11 uint64 = 1
                                    _ = v10
                                    _ = v11
                                } else {
                                    var v12 int8 = int8(fi6)
                                    if ((!((410 >= 255) && (938 <= 0))) || (-535 < -460)) {
                                        if (((p1 > p2) || true) || (0.0 == (-2.25 * (float32(0.1) - 0.1)))) {
                                            var v13 int8 = (-(v12))
                                            var v14 float64 = (float64(((0.0 + (3.14159 * 0.0)) + 1e10)) * 1e10)
                                            _ = v13
                                            _ = v14
                                        } else {
                                            var v15 int16 = 32767
                                            if (p1 > (((("" + (((("hello" + "hello") + p1) + p1) + "\n\t")) + "hello") + p2) + "a")) {
                                                v16 := GenS2{F0: false, F1: (v12 - -128)}
                                                v17 := []int{fi5, (fi5 * int(p0)), (fi5 << ((18446744073709551615) & 63)), fi6}
                                                _ = v16
                                                _ = v17
                                            }
                                            fmt.Println(fi5, p1, v12)
                                            v18 := GenS1{F0: (p2 + ((p2 + ("hello" + p2)) + ((p2 + "hello") + ("foo bar" + p1))))}
                                            var v19 int32 = (p0 + p0)
                                            _ = v15
                                            _ = v18
                                            _ = v19
                                        }
                                        v20 := GenS2{F0: ((!true) && (42 <= 769)), F1: v12}
                                        v21 := GenS1{F0: "foo bar"}
                                        v22 := GenS3{F0: 42, F1: 0, F2: v12}
                                        p0 = (^(p0))
                                        if (2 < 462) {
                                            if (true || false) {
                                                var v23 uint64 = 18446744073709551615
                                                v24 := []int{}
                                                v24 = append(v24, (fi5 & 32767))
                                                _ = v23
                                                _ = v24
                                            } else {
                                                var v25 uint16 = 1
                                                p0 += 42
                                                _ = v25
                                            }
                                            if false {
                                                v26 := []int{32768, fi5}
                                                var v27 uint32 = 79
                                                var v28 int8 = v12
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                            var v29 int16 = 32767
                                            _ = v29
                                        } else {
                                            if (3.14159 != 1e10) {
                                                switch fi5 {
                                                case 128:
                                                    if false {
                                                        p2 = "foo bar"
                                                    } else {
                                                        if (475 >= 1) {
                                                            if true {
                                                                fmt.Println(v12, p0, fi5)
                                                                if (255 <= 65535) {
                                                                    v30 := GenS3{F0: 0, F1: 32767, F2: v12}
                                                                    fmt.Println(fi5, fi6)
                                                                    if false {
                                                                        v31 := GenS1{F0: ((("foo bar" + ((p1 + p2) + p2)) + (p1 + p2)) + "hello")}
                                                                        v12 += (v12 | v12)
                                                                        var v32 uint32 = 42
                                                                        _ = v31
                                                                        _ = v32
                                                                    } else {
                                                                        v33 := []int{fi6}
                                                                        _ = v33
                                                                    }
                                                                    _ = v30
                                                                } else {
                                                                    p1 = ""
                                                                    if ((3.14159 * (1e10 / float64((-2.25 / 3.14159)))) >= (-2.25 * 1.5)) {
                                                                        var v34 string = ((p1 + (("a" + ("" + p2)) + ("\n\t" + p1))) + (("a" + (p1 + "foo bar")) + ("hello" + "hello")))
                                                                        _ = v34
                                                                    } else {
                                                                        if ((833 < 914) && false) {
                                                                            p2 = p1
                                                                            v35 := GenS1{F0: p2}
                                                                            fmt.Println(p1, p2, p2)
                                                                            _ = v35
                                                                        }
                                                                        switch p2 {
                                                                        case "foo bar":
                                                                            fmt.Println(p1)
                                                                        case "hello":
                                                                            var v36 string = p2
                                                                            _ = v36
                                                                        case "\n\t":
                                                                            fmt.Println(fi6, v12)
                                                                            var v37 int64 = 42
                                                                            switch 469 {
                                                                            case 4294967295:
                                                                                var v38 uint16 = 2
                                                                                _ = v38
                                                                            case 2147483647:
                                                                                var v39 bool = true
                                                                                _ = v39
                                                                            }
                                                                            _ = v37
                                                                        default:
                                                                            var v40 int = (fi5 & fi6)
                                                                            _ = v40
                                                                        }
                                                                        v41 := []int{(fi5 / ((fi6) | 1)), fi6, (-(fi5)), (-(fi6))}
                                                                        _ = v41
                                                                    }
                                                                }
                                                                switch p2 {
                                                                case "":
                                                                    v42 := []int{(fi6 ^ (^(fi5))), fi5, fi5, (fi6 * 456)}
                                                                    _ = v42
                                                                case "a":
                                                                    if false {
                                                                        v43 := GenS3{F0: 0, F1: -1, F2: (v12 & v12)}
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 string = ((p2 + (("a" + (p1 + ((((p2 + p2) + p1) + p1) + p1))) + ("" + p1))) + "foo bar")
                                                                        switch ((v44 + v44) + p1) {
                                                                        case "a":
                                                                            v45 := []int{32767, (fi6 | (fi6 | fi5)), 2147483647}
                                                                            var v46 float64 = 1.5
                                                                            var v47 uint = 372
                                                                            _ = v45
                                                                            _ = v46
                                                                            _ = v47
                                                                        case "\n\t":
                                                                            var v48 string = "hello"
                                                                            _ = v48
                                                                        case "hello":
                                                                            fmt.Println(p0)
                                                                        default:
                                                                            if (false && true) {
                                                                                switch fi5 {
                                                                                case 128:
                                                                                    v49 := []int{fi6, 32767}
                                                                                    _ = v49
                                                                                case 1:
                                                                                    var v50 int8 = 1
                                                                                    _ = v50
                                                                                case 775:
                                                                                    v51 := GenS1{F0: p1}
                                                                                    _ = v51
                                                                                default:
                                                                                    p0 += p0
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v44
                                                                    }
                                                                case "\n\t":
                                                                    v52 := GenS3{F0: 2, F1: 0, F2: v12}
                                                                    _ = v52
                                                                }
                                                            } else {
                                                                v53 := GenS3{F0: 65535, F1: 42, F2: v12}
                                                                p2 = (("hello" + p1) + "a")
                                                                var v54 uint16 = 32767
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                        } else {
                                                            v55 := GenS1{F0: ("a" + "a")}
                                                            _ = v55
                                                        }
                                                        v56 := GenS2{F0: false, F1: v12}
                                                        if (685 < 2147483647) {
                                                            p2 = p2
                                                        }
                                                        _ = v56
                                                    }
                                                case 32768:
                                                    v57 := []int{fi6, (fi5 ^ fi6), (fi5 ^ fi6), fi6}
                                                    _ = v57
                                                }
                                                p2 = "hello"
                                                var v58 uint = 18446744073709551615
                                                _ = v58
                                            }
                                            p0 -= 0
                                        }
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    } else {
                                        v59 := GenS2{F0: false, F1: v12}
                                        _ = v59
                                    }
                                    _ = v12
                                }
                            case 9223372036854775807:
                                switch p2 {
                                case "foo bar":
                                    p2 = "foo bar"
                                    p1 = "hello"
                                case "":
                                    if (true && (426 != 453)) {
                                        fmt.Println(p1)
                                    }
                                case "\n\t":
                                    fmt.Println(p2, fi6, p1)
                                default:
                                    fmt.Println(p2, p0, p1)
                                    p0 = (p0 % (((p0 << ((433) & 63))) | 1))
                                    v60 := []int{(fi5 + fi5), (fi5 * fi6), int(p0)}
                                    _ = v60
                                }
                            }
                            v61 := GenS2{F0: ((811 == 65535) || ((^(fi6)) == (fi5 & -2147483648))), F1: -1}
                            fmt.Println(fi5, p1)
                            _ = v61
                        }
                        v62 := []int{(fi5 & fi5), (^(fi5))}
                        p1 = p2
                        _ = v62
                    } else {
                        p0 -= p0
                        p2 = (p2 + (("hello" + ("hello" + "\n\t")) + p2))
                        v63 := []int{fi6}
                        _ = v63
                    }
                case 4294967295:
                    p2 = p2
                default:
                    v64 := GenS3{F0: 255, F1: 42, F2: 127}
                    _ = v64
                }
                _ = fi6
                _ = v7
            }
            if false {
                p1 = (p1 + p1)
                v65 := GenS4{F0: 1, F1: p0, F2: 1}
                _ = v65
            }
            _ = fi5
        }
    case 128:
        var v66 uint = 568
        _ = v66
    }
    v67 := GenS4{F0: -1, F1: 1, F2: 631}
    v68 := GenS5{F0: 1, F1: 42, F2: p1}
    if false {
        switch "" {
        case "foo bar":
            var v69 uint16 = 2
            _ = v69
        }
        if false {
            v70 := GenS6{F0: "a"}
            if (255 <= 42) {
                for fi71 := 0; fi71 < 6; fi71++ {
                    if true {
                        p0 = 323
                        if false {
                            var v72 float32 = -2.25
                            _ = v72
                        } else {
                            for fi73 := 0; fi73 < 5; fi73++ {
                                switch fi73 {
                                case 4294967295:
                                    fmt.Println(p1, p2, fi73)
                                case 32767:
                                    switch 128 {
                                    case 0:
                                        v74 := GenS5{F0: 1, F1: -276, F2: "hello"}
                                        _ = v74
                                    case 6:
                                        v75 := GenS5{F0: 980, F1: 609, F2: ((p1 + "") + p1)}
                                        _ = v75
                                    case 42:
                                        var v76 uint64 = 1
                                        _ = v76
                                    }
                                case -9223372036854775808:
                                    var v77 int64 = -9223372036854775808
                                    _ = v77
                                default:
                                    var v78 uint = 2
                                    _ = v78
                                }
                                v79 := GenS5{F0: -486, F1: 9223372036854775807, F2: p2}
                                p0 -= (p0 >> ((59) & 63))
                                _ = fi73
                                _ = v79
                            }
                            v80 := []int{(-(fi71)), fi71}
                            var v81 float64 = 0.0
                            _ = v80
                            _ = v81
                        }
                        if (!(!false)) {
                            var v82 uint32 = 42
                            var v83 int32 = (p0 | (p0 * (p0 ^ (p0 * -693))))
                            _ = v82
                            _ = v83
                        }
                    }
                    var v84 string = "hello"
                    v85 := GenS7{F0: 0}
                    _ = fi71
                    _ = v84
                    _ = v85
                }
                var v86 int32 = 1
                v87 := GenS8{F0: 42, F1: 0, F2: 255}
                _ = v86
                _ = v87
            }
            for fi88 := 0; fi88 < 3; fi88++ {
                var v89 int64 = 376
                switch (fi88 + fi88) {
                case 65:
                    var v90 bool = (!false)
                    _ = v90
                }
                _ = fi88
                _ = v89
            }
            _ = v70
        } else {
            v91 := GenS9{F0: 42}
            _ = v91
        }
    } else {
        var v92 uint64 = 4294967295
        v93 := []int{2147483647, 0, -556, 42}
        if (!false) {
            v92 += v92
        } else {
            v94 := GenS9{F0: v92}
            if len(v93) >= 2 {
                v93 = v93[1:len(v93)]
            }
            _ = v94
        }
        _ = v92
        _ = v93
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v67
    _ = v68
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 0
    var v3 bool = false
    v95 := GenS1{F0: "a"}
    v3 = true
    for fi96 := 0; fi96 < 11; fi96++ {
        var v97 int64 = -720
        v3 = v3
        v98 := GenS8{F0: 1, F1: (v2 / ((v2) | 1)), F2: 65535}
        _ = fi96
        _ = v97
        _ = v98
    }
    var v99 float32 = float32(3.14159)
    v100 := GenS10{F0: 2}
    v101 := GenS7{F0: 1}
    var v102 int32 = -16
    var v103 float64 = (((3.14159 / (float64((3.14159 * float32(3.14159))) + float64(v99))) + 3.14159) * 0.0)
    if v3 {
        for fi104 := 0; fi104 < 1; fi104++ {
            if (954 <= 958) {
                for fi105 := 0; fi105 < 14; fi105++ {
                    var v106 int64 = 42
                    var v107 int64 = v106
                    var v108 uint = 210
                    _ = fi105
                    _ = v106
                    _ = v107
                    _ = v108
                }
                if (v103 != 0.0) {
                    var v109 int16 = (v2 / ((v2) | 1))
                    var v110 uint64 = 740
                    _ = v109
                    _ = v110
                }
            }
            _ = fi104
        }
    } else {
        var v111 int16 = v2
        _ = v111
    }
    fmt.Println(v103, v103)
    for fi112 := 0; fi112 < 20; fi112++ {
        fmt.Println(v102, v2, v103)
        var v113 string = "\n\t"
        _ = fi112
        _ = v113
    }
    v114 := GenS5{F0: 32767, F1: 103, F2: "a"}
    v103 = v103
    var v115 int = 647
    for fi116 := 0; fi116 < 18; fi116++ {
        v117 := GenS11{F0: -9223372036854775808}
        v115 += v115
        var v118 float32 = (v99 + (float32(v103) - v99))
        _ = fi116
        _ = v117
        _ = v118
    }
    if v3 {
        fmt.Println(v3, v3, v2)
        _ = len(v1)
        fmt.Println(v2)
    }
    var v119 uint = 449
    v120 := GenS12{F0: -57, F1: 602}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v115:", v115)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v103
    _ = v114
    _ = v115
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v95
    _ = v99
}
