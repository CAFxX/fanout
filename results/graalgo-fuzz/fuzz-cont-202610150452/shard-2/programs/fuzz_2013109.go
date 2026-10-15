package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS7 struct { F0 string; F1 int32; F2 string }
type GenS8 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 bool; F1 bool; F2 int64 }
type GenS11 struct { F0 uint64; F1 int8; F2 int64 }
func genFn0(p0 int16, p1 int) int16 {
    switch (p1 + p1) {
    case 1:
        var v6 uint8 = 144
        _ = v6
    case -2147483648:
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                p1 -= p1
                var v9 uint32 = 2147483648
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case 42:
        var v10 int16 = (p0 >> ((255) & 63))
        _ = v10
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            p1 = fi11
            v12 := []int{(fi11 & 844), (fi11 >> ((4294967295) & 63))}
            v13 := GenS1{F0: 42, F1: "", F2: 2}
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: -497, F1: ("\n\t" + "\n\t"), F2: 65535}
        var v16 int64 = 0
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 42, F1: ("\n\t" + (("hello" + ("a" + ("\n\t" + "foo bar"))) + (("\n\t" + "foo bar") + "\n\t"))), F2: 0}
    v18 := []int{p1, p1, p1, p1}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return p0
}

func main() {
    var v1 int16 = -32768
    var v2 uint64 = 0
    var v3 bool = (1 <= 100)
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v19 int32 = 1
    switch ("a" + "") {
    case "a":
        v3 = v3
    }
    v5[""] = v3
    v4["foo bar"] = 441
    switch "" {
    case "":
        switch "hello" {
        case "\n\t":
            switch 97 {
            case -1:
                for fi20 := 0; fi20 < 4; fi20++ {
                    if v3 {
                        v3 = (128 >= 2)
                        switch fi20 {
                        case 1:
                            var v21 float64 = float64(1.5)
                            _ = v21
                        }
                        v19 = v19
                    } else {
                        v22 := GenS1{F0: (^(v19)), F1: ("a" + ("hello" + ("a" + "\n\t"))), F2: 4294967295}
                        _ = v22
                    }
                    v23 := GenS2{F0: 664, F1: 147, F2: 65535}
                    if (2147483647 != fi20) {
                        if v3 {
                            switch -2147483648 {
                            case -18:
                                for fi24 := 0; fi24 < 13; fi24++ {
                                    var v25 uint8 = 1
                                    var v26 uint8 = (^(v25))
                                    var v27 int32 = (v19 ^ v19)
                                    var v28 float64 = 0.0
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            case 32768:
                                var v29 uint32 = 2
                                _ = v29
                            case 594:
                                var v30 uint16 = 2
                                _ = v30
                            default:
                                switch ("\n\t" + ("\n\t" + "")) {
                                case "hello":
                                    for fi31 := 0; fi31 < 10; fi31++ {
                                        var v32 float32 = (1e10 - (1.5 + (float32(float64(float32((3.14159 - 0.1)))) / (((0.1 / 3.14159) + (float32((3.14159 + 0.0)) / (1.5 * -2.25))) / 1e10))))
                                        v33 := GenS1{F0: 1, F1: (("\n\t" + "") + "a"), F2: 18446744073709551615}
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                case "":
                                    for fi34 := 0; fi34 < 16; fi34++ {
                                        v19 = 265
                                        v19 = (v19 >> ((4294967295) & 63))
                                        _ = fi34
                                    }
                                case "\n\t":
                                    v35 := GenS3{F0: "", F1: 1}
                                    _ = v35
                                }
                            }
                            switch 42 {
                            case 128:
                                var v36 uint32 = 2147483648
                                _ = v36
                            case 2147483647:
                                switch (("hello" + "a") + "a") {
                                case "\n\t":
                                    fmt.Println(v3, v19)
                                case "":
                                    v19 = v19
                                case "foo bar":
                                    var v37 uint8 = 255
                                    _ = v37
                                }
                            case -929:
                                fmt.Println(v1, fi20)
                            default:
                                for fi38 := 0; fi38 < 2; fi38++ {
                                    v3 = v3
                                    delete(v5, ("hello" + (("" + "foo bar") + ("\n\t" + (("foo bar" + "") + ("foo bar" + ("" + (("a" + "") + "\n\t"))))))))
                                    _ = fi38
                                }
                            }
                        }
                        var v39 float64 = 3.14159
                        switch "hello" {
                        case "":
                            if (9223372036854775807 > 42) {
                                if false {
                                    var v40 uint16 = 672
                                    _ = v40
                                } else {
                                    v41 := GenS2{F0: 4294967295, F1: 42, F2: 694}
                                    _ = v41
                                }
                                if v3 {
                                    fmt.Println(v2)
                                    var v42 int32 = -2147483648
                                    var v43 uint64 = (v2 << ((1) & 63))
                                    _ = v42
                                    _ = v43
                                }
                            }
                        }
                        _ = v39
                    }
                    _ = fi20
                    _ = v23
                }
            case 9223372036854775807:
                var v44 bool = v3
                _ = v44
            }
        case "":
            for fi45 := 0; fi45 < 14; fi45++ {
                v1 = 0
                for fi46 := 0; fi46 < 20; fi46++ {
                    switch "a" {
                    case "":
                        if (!(16 <= 183)) {
                            v19 -= v19
                            switch "foo bar" {
                            case "foo bar":
                                if v3 {
                                    var v47 float64 = (0.1 * float64(0.0))
                                    if (v3 || ((("" + ("a" + "foo bar")) + "") > ("\n\t" + ""))) {
                                        v2 -= 2
                                    }
                                    v19 = v19
                                    _ = v47
                                } else {
                                    switch ("a" + ("a" + "\n\t")) {
                                    case "":
                                        var v48 float32 = (-2.25 + 1e10)
                                        _ = v48
                                    }
                                }
                            case "a":
                                v49 := v4[("\n\t" + "foo bar")]
                                _ = v49
                            default:
                                if (4294967295 == 2) {
                                    if (187 < 216) {
                                        v19 = v19
                                        var v50 int16 = v1
                                        if (42 <= 1) {
                                            v51 := GenS4{F0: 4294967295}
                                            v52 := GenS5{F0: -128}
                                            _ = v51
                                            _ = v52
                                        } else {
                                            fmt.Println(v1, v3, v3)
                                        }
                                        _ = v50
                                    } else {
                                        if (4294967295 != 0) {
                                            fmt.Println(fi45)
                                        }
                                        if v3 {
                                            if true {
                                                if false {
                                                    var v53 uint64 = (v2 | v2)
                                                    var v54 uint32 = 55
                                                    var v55 bool = v3
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                }
                                                fmt.Println(v2, fi45, v19)
                                            }
                                        } else {
                                            if v3 {
                                                var v56 int16 = genFn0(32767, (fi46 * -2147483648))
                                                var v57 bool = (0 > 0)
                                                _ = v56
                                                _ = v57
                                            }
                                        }
                                        var v58 uint = 881
                                        if v3 {
                                            if true {
                                                v3 = (27 > 0)
                                            }
                                            v59 := []int{fi45, fi45}
                                            if v3 {
                                                var v60 int16 = v1
                                                _ = v60
                                            }
                                            v61 := GenS4{F0: 2147483647}
                                            var v62 float64 = (3.14159 / ((1.5 / 1.0) * 1.5))
                                            var v63 uint16 = 42
                                            if (128 >= 0) {
                                                var v64 string = "\n\t"
                                                if (1 > 0) {
                                                    v2 = 4294967295
                                                    v4[v64] = (^(fi46))
                                                }
                                                var v65 float64 = 1e10
                                                var v66 string = v64
                                                v1 = 926
                                                v64 = v64
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            } else {
                                                v67 := GenS6{F0: 1, F1: 127, F2: -32768}
                                                var v68 int = 702
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = v59
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        }
                                        _ = v58
                                    }
                                    switch fi46 {
                                    case 853:
                                        v69 := v5["foo bar"]
                                        _ = v69
                                    case 9223372036854775807:
                                        v70 := GenS7{F0: ("a" + ""), F1: (v19 % (((v19 ^ int32(fi46))) | 1)), F2: "hello"}
                                        _ = v70
                                    default:
                                        v71 := v5["\n\t"]
                                        _ = v71
                                    }
                                    var v72 uint32 = 2147483647
                                    _ = v72
                                } else {
                                    var v73 int = (fi45 / ((fi46) | 1))
                                    _ = v73
                                }
                            }
                        } else {
                            v5["foo bar"] = (128 < 255)
                        }
                    case "a":
                        var v74 int = fi45
                        _ = v74
                    case "foo bar":
                        v75 := GenS3{F0: "foo bar", F1: -9223372036854775808}
                        _ = v75
                    }
                    v76 := GenS7{F0: "foo bar", F1: v19, F2: "hello"}
                    fmt.Println(fi45, v2, v19)
                    _ = fi46
                    _ = v76
                }
                _ = fi45
            }
        case "hello":
            v77 := []int{-961, -9223372036854775808}
            _ = v77
        }
    case "a":
        var v78 int32 = v19
        _ = v78
    case "\n\t":
        if v3 {
            v79 := GenS1{F0: v19, F1: "\n\t", F2: 65535}
            _ = v79
        } else {
            if true {
                switch 1 {
                case 127:
                    if v3 {
                        var v80 bool = true
                        _ = v80
                    } else {
                        switch 0 {
                        case -737:
                            v19 += -566
                        case 32768:
                            v81 := GenS8{F0: 59, F1: (v19 >> ((42) & 63)), F2: 255}
                            _ = v81
                        default:
                            fmt.Println(v1, v2)
                        }
                        for fi82 := 0; fi82 < 16; fi82++ {
                            for fi83 := 0; fi83 < 17; fi83++ {
                                var v84 int32 = 42
                                fmt.Println(v84, v84, v1)
                                _ = fi83
                                _ = v84
                            }
                            _ = fi82
                        }
                        if true {
                            var v85 uint64 = 0
                            var v86 int64 = -1
                            _ = v85
                            _ = v86
                        }
                    }
                case 42:
                    var v87 float32 = (0.1 - (0.1 - -2.25))
                    _ = v87
                case -632:
                    v88 := GenS9{F0: (("" + "") + "\n\t"), F1: "\n\t"}
                    _ = v88
                }
                if (2147483648 == 750) {
                    var v89 bool = true
                    if (!false) {
                        v90 := GenS5{F0: 0}
                        for fi91 := 0; fi91 < 2; fi91++ {
                            v3 = v89
                            var v92 int16 = v1
                            v93 := []int{(fi91 % (((^(fi91))) | 1)), (fi91 ^ (fi91 * fi91)), fi91, (fi91 << ((42) & 63))}
                            _ = fi91
                            _ = v92
                            _ = v93
                        }
                        var v94 uint16 = 0
                        _ = v90
                        _ = v94
                    } else {
                        switch 9223372036854775807 {
                        case 42:
                            v95 := GenS10{F0: false, F1: false, F2: -396}
                            _ = v95
                        case 0:
                            v96 := GenS1{F0: v19, F1: ("a" + "hello"), F2: 0}
                            _ = v96
                        default:
                            if ((v2 * v2) < (v2 >> ((255) & 63))) {
                                v97 := GenS10{F0: (false && false), F1: ((v2 | v2) >= (v2 >> ((591) & 63))), F2: -725}
                                var v98 bool = v89
                                if false {
                                    v3 = v98
                                } else {
                                    v99 := GenS5{F0: 100}
                                    v100 := GenS2{F0: 265, F1: 1, F2: 250}
                                    var v101 int64 = 9223372036854775807
                                    _ = v100
                                    _ = v101
                                    _ = v99
                                }
                                _ = v97
                                _ = v98
                            } else {
                                v102 := []int{127}
                                if true {
                                    v103 := GenS5{F0: 1}
                                    v2 = v2
                                    _ = v103
                                }
                                v19 = 2147483647
                                _ = v102
                            }
                        }
                        v89 = v89
                    }
                    _ = v89
                } else {
                    var v104 bool = true
                    v105 := GenS11{F0: v2, F1: 1, F2: 9223372036854775807}
                    _ = v104
                    _ = v105
                }
            }
            var v106 uint = 65535
            _ = v106
        }
    default:
        fmt.Println(v3, v3)
    }
    v2 = 18446744073709551615
    switch 706 {
    case -482:
        delete(v5, (("a" + "") + (("a" + (("\n\t" + ((("hello" + "\n\t") + "foo bar") + ("a" + "hello"))) + ("" + ("\n\t" + "")))) + "")))
    }
    var v107 uint16 = 644
    v3 = true
    fmt.Println(v1, v107)
    v108 := GenS11{F0: (v2 * (v2 ^ v2)), F1: 1, F2: 688}
    var v109 uint8 = 42
    v110 := []int{-268, 379}
    if v3 {
        var v111 int64 = -1
        _ = v111
    }
    var v112 int64 = -9223372036854775808
    var v113 uint32 = 2147483648
    if v3 {
        var v114 uint64 = 4294967295
        _ = v114
    } else {
        v115 := GenS11{F0: 616, F1: -128, F2: (v112 ^ (v112 / ((-491) | 1)))}
        if false {
            v19 += 42
        }
        v109 += (v109 % ((93) | 1))
        _ = v115
    }
    switch ("a" + ("" + ("hello" + "foo bar"))) {
    case "hello":
        for fi116 := 0; fi116 < 4; fi116++ {
            var v117 bool = ((v109 < v109) && (255 == 192))
            v112 = v112
            _ = fi116
            _ = v117
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v109:", v109)
    fmt.Println("slicelen v110:", len(v110), v110)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v110
    _ = v112
    _ = v113
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
