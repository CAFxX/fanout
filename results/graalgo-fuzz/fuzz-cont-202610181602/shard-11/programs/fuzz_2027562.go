package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 uint8; F2 int8 }
func genFn0(p0 string, p1 int64) int {
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = ("hello" + p0)
        _ = fi6
    }
    v7 := []int{int(p1)}
    _ = p0
    _ = p1
    _ = v7
    return 32768
}
func genFn1(p0 int8, p1 int8, p2 uint16) bool {
    v8 := GenS1{F0: 65535, F1: -889}
    var v9 uint16 = 255
    p1 = 34
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (0 >= 765)
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 uint16 = 2
    var v5 float64 = (3.14159 / 1.0)
    if (((float64(1.5) - v5) * float64(3.14159)) < 0.0) {
        if (0 < 0) {
            if ((!(true && ((32767 == (v4 - (v4 << ((4294967295) & 63)))) && true))) || (12 <= 255)) {
                v5 = 1e10
            }
            switch "hello" {
            case "\n\t":
                switch -1 {
                case -75:
                    v10 := GenS2{F0: 519, F1: 424}
                    _ = v10
                case 1:
                    v5 += v5
                case 2147483647:
                    var v11 int64 = 9223372036854775807
                    _ = v11
                }
            case "hello":
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS3{F0: ("" + "foo bar"), F1: 510}
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 float32 = 0.0
                        _ = fi14
                        _ = v15
                    }
                    var v16 uint32 = 42
                    _ = fi12
                    _ = v13
                    _ = v16
                }
            case "":
                var v17 uint16 = v4
                _ = v17
            }
            if true {
                for fi18 := 0; fi18 < 14; fi18++ {
                    if false {
                        v1 = v1
                        v5 = v5
                        v19 := GenS3{F0: "a", F1: (^(v4))}
                        _ = v19
                    } else {
                        v1 -= 0
                        var v20 string = ("hello" + (((((("foo bar" + "\n\t") + "\n\t") + "hello") + "\n\t") + "foo bar") + ("foo bar" + "hello")))
                        fmt.Println(v5, v5, v1)
                        _ = v20
                    }
                    _ = fi18
                }
                var v21 string = (("a" + "\n\t") + "foo bar")
                _ = v21
            } else {
                v22 := []int{32767}
                if len(v22) > 0 {
                    v22[0] = 9223372036854775807
                }
                v23 := GenS2{F0: 2, F1: 255}
                _ = v22
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := v3["hello"]
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 uint32 = 2147483647
                    if (1.5 <= 1e10) {
                        fmt.Println(v27, v4)
                    }
                    if (!(0.1 < 1.5)) {
                        var v28 string = (v25 + "hello")
                        _ = v28
                    } else {
                        switch v25 {
                        case "a":
                            var v29 string = v25
                            _ = v29
                        case "foo bar":
                            if (3.14159 <= v5) {
                                fmt.Println(v27, v1)
                                var v30 int = (fi26 << ((131) & 63))
                                v25 = v25
                                _ = v30
                            } else {
                                v25 = ("" + "\n\t")
                            }
                        }
                        if (!(((float64(0.1) - (v5 * float64(3.14159))) / 3.14159) != v5)) {
                            var v31 int32 = -35
                            v31 = v31
                            _ = v31
                        } else {
                            var v32 int64 = int64(fi26)
                            v33 := GenS1{F0: 65535, F1: -1}
                            _ = v32
                            _ = v33
                        }
                        var v34 float32 = 0.0
                        _ = v34
                    }
                    _ = fi26
                    _ = v27
                }
                if ((1 < -128) || ((v25 + (v25 + "hello")) > (v25 + ("\n\t" + ("\n\t" + v25))))) {
                    var v35 string = v25
                    _ = v35
                } else {
                    v1 += 317
                }
                _ = fi24
                _ = v25
            }
        }
        var v36 int = 889
        _ = v36
    } else {
        v37 := GenS3{F0: "", F1: 0}
        _ = v37
    }
    v3["hello"] = ""
    _ = len(v2)
    if (v1 != v1) {
        for fi38 := 0; fi38 < 17; fi38++ {
            var v39 uint = 18446744073709551615
            v40 := v3[("foo bar" + "\n\t")]
            v4 -= v4
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        if false {
            var v42 uint64 = 2
            v43 := GenS4{F0: (v1 | 0), F1: 100}
            v1 = v1
            _ = v42
            _ = v43
        } else {
            var v44 uint32 = 4294967295
            _ = v44
        }
        _ = fi41
    }
    if (-9223372036854775808 > 2147483647) {
        var v45 int8 = -18
        if true {
            switch -1 {
            case 2147483647:
                if (42 < -677) {
                    v45 += (v45 | (^(v45)))
                    if (v1 == v1) {
                        for fi46 := 0; fi46 < 4; fi46++ {
                            var v47 int64 = 0
                            v45 = v45
                            _ = fi46
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 4; fi48++ {
                            var v49 bool = (((1e10 / 3.14159) + ((float32(v5) / 0.1) / 0.1)) < 1.5)
                            v50 := GenS4{F0: v1, F1: (-(v45))}
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        fmt.Println(v45)
                    }
                } else {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        if (2 <= 2) {
                            if true {
                                fmt.Println(v45)
                                v5 = v5
                                v5 -= (v5 - v5)
                            }
                            v4 -= 0
                            v52 := []int{genFn0((("a" + "hello") + "hello"), -9223372036854775808), genFn0("", 42), 32768}
                            _ = v52
                        }
                        v53 := []int{(-(fi51))}
                        _ = fi51
                        _ = v53
                    }
                }
            case -2147483648:
                v45 -= (v45 >> ((1) & 63))
            default:
                v1 = v1
            }
            for fi54 := 0; fi54 < 18; fi54++ {
                if true {
                    var v55 int32 = 1
                    _ = v55
                } else {
                    v5 = 1.5
                    var v56 int = fi54
                    v57 := GenS1{F0: 2147483647, F1: -9223372036854775808}
                    _ = v56
                    _ = v57
                }
                _ = fi54
            }
            if true {
                var v58 int64 = 916
                v59 := GenS1{F0: 728, F1: (v58 & -1)}
                if (1.5 > -2.25) {
                    v1 += -1
                    switch ((("hello" + "hello") + (("" + "foo bar") + "\n\t")) + ("" + "")) {
                    case "":
                        v60 := GenS5{F0: -538, F1: v1}
                        _ = v60
                    case "\n\t":
                        var v61 int16 = (v1 / (((v1 >> ((65535) & 63))) | 1))
                        _ = v61
                    }
                    if (42 > 2) {
                        v62 := GenS5{F0: v58, F1: v1}
                        v3[(("" + ("hello" + "foo bar")) + "hello")] = (("hello" + "a") + ("hello" + ("\n\t" + "foo bar")))
                        var v63 uint = 390
                        _ = v62
                        _ = v63
                    } else {
                        var v64 uint32 = 2147483648
                        _ = v64
                    }
                }
                _ = v58
                _ = v59
            } else {
                var v65 int8 = -72
                for fi66 := 0; fi66 < 19; fi66++ {
                    fmt.Println(v65)
                    v65 = v65
                    for fi67 := 0; fi67 < 7; fi67++ {
                        if ("\n\t" == ("a" + ((("hello" + (("a" + "") + (("foo bar" + "") + "foo bar"))) + (("\n\t" + "\n\t") + "a")) + "a"))) {
                            fmt.Println(v45, fi66, v4)
                            v68 := GenS1{F0: 42, F1: -1}
                            v69 := GenS4{F0: v1, F1: (v45 - 1)}
                            _ = v68
                            _ = v69
                        } else {
                            v5 += float64(1e10)
                        }
                        v70 := GenS6{F0: 1}
                        v5 = v5
                        _ = fi67
                        _ = v70
                    }
                    _ = fi66
                }
                _ = v65
            }
        }
        _ = v45
    } else {
        v1 = (v1 ^ v1)
        fmt.Println(v1, v5, v1)
        v71 := []int{}
        _ = v71
    }
    var v72 int8 = 1
    var v73 int16 = -32768
    v74 := GenS1{F0: 65535, F1: -764}
    if (false && (!((v72 | v72) < (v72 - -1)))) {
        for fi75 := 0; fi75 < 18; fi75++ {
            var v76 bool = true
            switch (fi75 / (((-(fi75))) | 1)) {
            case -9223372036854775808:
                switch "foo bar" {
                case "hello":
                    if (546 != 1) {
                        var v77 float64 = float64(((((float32((v5 + 0.0)) / 0.1) * 0.0) * ((1.5 * float32(0.0)) / float32((((v5 + (v5 / (3.14159 - (v5 - v5)))) / (float64(1e10) + ((0.1 / (v5 / ((((v5 - (3.14159 * (v5 * 3.14159))) / 1.0) * v5) + 0.0))) + (3.14159 + float64(float32((v5 + 1e10))))))) - v5)))) + (-2.25 - -2.25)))
                        _ = v77
                    } else {
                        v78 := GenS3{F0: "", F1: 0}
                        v79 := GenS5{F0: 315, F1: v73}
                        _ = len(v2)
                        _ = v78
                        _ = v79
                    }
                case "a":
                    v1 = (v73 - (v73 | v73))
                case "\n\t":
                    var v80 uint8 = 112
                    _ = v80
                default:
                    for fi81 := 0; fi81 < 8; fi81++ {
                        var v82 uint = 42
                        _ = fi81
                        _ = v82
                    }
                }
            case 9223372036854775807:
                var v83 int32 = 1
                _ = v83
            case 32767:
                if (456 != (v73 >> ((uint(v4)) & 63))) {
                    switch "\n\t" {
                    case "":
                        var v84 bool = true
                        _ = v84
                    case "a":
                        switch "foo bar" {
                        case "a":
                            switch (fi75 ^ (fi75 << ((935) & 63))) {
                            case -1:
                                v85 := GenS4{F0: 1, F1: int8(fi75)}
                                _ = v85
                            default:
                                if (v76 && (v5 != v5)) {
                                    v1 = (v1 % ((32767) | 1))
                                    v86 := GenS7{F0: "\n\t", F1: 0, F2: 42}
                                    if false {
                                        v1 = -32768
                                        if (v76 && v76) {
                                            switch (fi75 ^ (fi75 & fi75)) {
                                            case -2147483648:
                                                if false {
                                                    v87 := v2[(fi75 & (fi75 & fi75))]
                                                    if v76 {
                                                        var v88 uint = 632
                                                        if ((v88 >> ((v88) & 63)) > v88) {
                                                            v89 := []int{fi75, genFn0("hello", -1), fi75}
                                                            fmt.Println(v72, v1, v76)
                                                            var v90 int32 = -2147483648
                                                            _ = v89
                                                            _ = v90
                                                        } else {
                                                            v91 := v2[genFn0((("\n\t" + ("hello" + "hello")) + ((("hello" + "foo bar") + ((("hello" + ("" + "foo bar")) + "foo bar") + "a")) + "\n\t")), int64(fi75))]
                                                            _ = v91
                                                        }
                                                        v73 = int16(fi75)
                                                        _ = v88
                                                    } else {
                                                        fmt.Println(v76, v76, v1)
                                                        for fi92 := 0; fi92 < 7; fi92++ {
                                                            if (1 != 512) {
                                                                if true {
                                                                    if ((v87 * (v87 >> ((65535) & 63))) >= v87) {
                                                                        v93 := []int{genFn0("a", 680)}
                                                                        _ = v93
                                                                    }
                                                                }
                                                            }
                                                            v94 := GenS3{F0: "a", F1: v4}
                                                            _ = fi92
                                                            _ = v94
                                                        }
                                                        fmt.Println(v73, v72, v4)
                                                    }
                                                    v95 := GenS6{F0: 255}
                                                    _ = v87
                                                    _ = v95
                                                } else {
                                                    switch (("hello" + "foo bar") + "a") {
                                                    case "a":
                                                        for fi96 := 0; fi96 < 15; fi96++ {
                                                            v1 = v1
                                                            var v97 uint = 4294967295
                                                            var v98 float32 = (-2.25 - 1.5)
                                                            _ = fi96
                                                            _ = v97
                                                            _ = v98
                                                        }
                                                    case "hello":
                                                        fmt.Println(v5, fi75)
                                                    case "foo bar":
                                                        var v99 bool = false
                                                        _ = v99
                                                    }
                                                    for fi100 := 0; fi100 < 15; fi100++ {
                                                        var v101 int8 = (v72 | v72)
                                                        _ = fi100
                                                        _ = v101
                                                    }
                                                }
                                            }
                                        }
                                        _ = len(v3)
                                    } else {
                                        v4 += v4
                                    }
                                    _ = v86
                                }
                            }
                        case "\n\t":
                            v1 -= v73
                        case "hello":
                            v1 = (v73 ^ -32768)
                        }
                    case "foo bar":
                        fmt.Println(v5, fi75)
                    }
                    var v102 bool = v76
                    _ = v102
                } else {
                    v103 := GenS6{F0: 127}
                    var v104 uint16 = (v4 >> ((4294967295) & 63))
                    _ = v103
                    _ = v104
                }
            }
            _ = fi75
            _ = v76
        }
    }
    v4 = v4
    for fi105 := 0; fi105 < 15; fi105++ {
        fmt.Println(fi105)
        var v106 uint8 = 2
        delete(v2, fi105)
        _ = fi105
        _ = v106
    }
    v107 := []int{42, genFn0("", -391), 9223372036854775807, 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v107:", len(v107), v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v74
}
