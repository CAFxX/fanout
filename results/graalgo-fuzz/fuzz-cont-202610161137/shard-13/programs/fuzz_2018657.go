package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int; F1 uint; F2 string }
type GenS7 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 0
    var v3 int = -9223372036854775808
    var v4 int16 = 1000
    v5 := make(map[string]uint32)
    var v6 int8 = 127
    var v7 uint16 = 2
    if (v2 != v2) {
        switch (v3 & int(v6)) {
        case 484:
            v5[("a" + "foo bar")] = 726
        case 32767:
            var v8 string = (("hello" + "\n\t") + "\n\t")
            _ = v8
        case 625:
            v6 = (v6 | 42)
        default:
            switch ("foo bar" + ("a" + "foo bar")) {
            case "\n\t":
                v6 -= (v6 << ((18446744073709551615) & 63))
            case "a":
                var v9 string = (("foo bar" + "a") + "a")
                _ = v9
            case "hello":
                var v10 float32 = (3.14159 - 3.14159)
                _ = v10
            default:
                for fi11 := 0; fi11 < 17; fi11++ {
                    v2 -= v2
                    for fi12 := 0; fi12 < 1; fi12++ {
                        delete(v5, ("hello" + (("" + "") + ("" + "a"))))
                        var v13 uint8 = 128
                        var v14 uint32 = 97
                        fmt.Println(fi12, v4)
                        fmt.Println(v7, v14)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    if true {
                        for fi15 := 0; fi15 < 12; fi15++ {
                            var v16 int = (-(v3))
                            var v17 bool = (9223372036854775807 != 32767)
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        v7 = (v7 & 34)
                    }
                    _ = fi11
                }
            }
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS1{F0: 42}
            _ = fi18
            _ = v19
        }
    }
    if (!true) {
        fmt.Println(v3, v2, v6)
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 float64 = (((0.0 / float64(((1.5 - (1e10 * (3.14159 * 1.5))) + float32((-2.25 * 1.5))))) + 0.0) * -2.25)
            var v22 uint16 = 454
            v23 := []int{(^(v3))}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        var v24 uint16 = v7
        v2 -= v2
        _ = v24
    }
    v25 := GenS1{F0: 862}
    var v26 uint64 = 920
    for fi27 := 0; fi27 < 16; fi27++ {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v6, v26)
        }
        _ = fi27
    }
    v2 -= v2
    fmt.Println(v2, v6)
    v7 += v7
    v28 := GenS1{F0: (v7 - 385)}
    if (1.5 >= (3.14159 - -2.25)) {
        if true {
            switch "hello" {
            case "\n\t":
                fmt.Println(v7)
            case "hello":
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 float64 = 0.0
                    _ = fi29
                    _ = v30
                }
            }
            v31 := v1[2]
            if false {
                if (true && (v6 == (v6 % (((v6 / ((-105) | 1))) | 1)))) {
                    fmt.Println(v31, v31, v26)
                    for fi32 := 0; fi32 < 15; fi32++ {
                        v33 := []int{4294967295, (fi32 / ((32767) | 1)), fi32}
                        _ = fi32
                        _ = v33
                    }
                    if (!false) {
                        var v34 float32 = (3.14159 + 0.1)
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 19; fi35++ {
                            var v36 int16 = v4
                            if false {
                                v2 = v2
                                if false {
                                    v37 := v1[uint32(v7)]
                                    delete(v1, v37)
                                    for fi38 := 0; fi38 < 18; fi38++ {
                                        v39 := GenS1{F0: 255}
                                        _ = fi38
                                        _ = v39
                                    }
                                    _ = v37
                                }
                                v40 := GenS2{F0: v2, F1: 127}
                                _ = v40
                            } else {
                                v41 := []int{v3, (v3 % ((int(v4)) | 1))}
                                v42 := GenS1{F0: 2}
                                if true {
                                    for fi43 := 0; fi43 < 14; fi43++ {
                                        if (false && (true && false)) {
                                            var v44 int8 = -128
                                            var v45 int8 = (^(v6))
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = fi43
                                    }
                                    if (!false) {
                                        v6 += 1
                                        if ((74 >= 128) && false) {
                                            v46 := GenS2{F0: v2, F1: int(v31)}
                                            v47 := v5["hello"]
                                            v48 := GenS1{F0: (^(v7))}
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        } else {
                                            fmt.Println(fi35, v4)
                                            v49 := v1[(v31 + (v31 - v31))]
                                            v50 := GenS3{F0: 4294967295, F1: "a"}
                                            _ = v49
                                            _ = v50
                                        }
                                        fmt.Println(v4)
                                    } else {
                                        var v51 int32 = (v2 | (v2 | v2))
                                        v51 -= (v2 ^ 42)
                                        if ((v26 * (v26 + 4294967295)) < 735) {
                                            v52 := GenS1{F0: v7}
                                            var v53 string = "a"
                                            if false {
                                                if len(v41) > 0 {
                                                    v41[0] = fi35
                                                }
                                            }
                                            _ = v52
                                            _ = v53
                                        } else {
                                            if true {
                                                v26 = (^(v26))
                                                var v54 uint32 = 65535
                                                v7 = (v7 + v7)
                                                _ = v54
                                            } else {
                                                v36 -= (v36 ^ v36)
                                                var v55 uint64 = 2
                                                _ = v55
                                            }
                                            var v56 uint16 = (v7 - 961)
                                            _ = v56
                                        }
                                        _ = v51
                                    }
                                    fmt.Println(v36, v36)
                                }
                                _ = v41
                                _ = v42
                            }
                            _ = fi35
                            _ = v36
                        }
                        fmt.Println(v6, v2)
                        if ((v31 << ((65535) & 63)) < v31) {
                            v57 := []int{(v3 * v3), v3, 127}
                            if false {
                                var v58 uint64 = v26
                                _ = v58
                            } else {
                                v31 = (v31 - (v31 | 4294967295))
                            }
                            _ = v57
                        } else {
                            switch v3 {
                            case 128:
                                v7 = v7
                            case -833:
                                v59 := GenS4{F0: -2147483648}
                                _ = v59
                            }
                            for fi60 := 0; fi60 < 19; fi60++ {
                                fmt.Println(v7, v7)
                                _ = fi60
                            }
                        }
                    }
                }
            }
            _ = v31
        }
        v4 = v4
    } else {
        if false {
            var v61 uint64 = uint64(v7)
            if (false || true) {
                if (!false) {
                    switch "" {
                    case "a":
                        v61 -= v26
                    default:
                        if (-32768 >= (v4 ^ 42)) {
                            switch (("foo bar" + "\n\t") + ("\n\t" + "foo bar")) {
                            case "hello":
                                v26 = (v26 % ((v26) | 1))
                            case "":
                                var v62 uint32 = 2147483647
                                _ = v62
                            }
                            _ = len(v5)
                        }
                    }
                    var v63 uint8 = 255
                    _ = v63
                } else {
                    if true {
                        var v64 int8 = -9
                        _ = v64
                    }
                    var v65 float64 = 3.14159
                    for fi66 := 0; fi66 < 8; fi66++ {
                        if (false || (true || true)) {
                            _ = len(v5)
                        } else {
                            v6 = v6
                            v6 += v6
                            fmt.Println(v4)
                        }
                        v3 = fi66
                        v2 += (v2 + (v2 % ((v2) | 1)))
                        _ = fi66
                    }
                    _ = v65
                }
            }
            if ((v2 & (v2 | 2147483647)) >= v2) {
                v67 := GenS5{F0: v7, F1: "a"}
                fmt.Println(v7, v3, v2)
                _ = v67
            }
            _ = v61
        } else {
            for fi68 := 0; fi68 < 7; fi68++ {
                switch "\n\t" {
                case "hello":
                    v69 := GenS5{F0: 0, F1: ("hello" + ("\n\t" + "foo bar"))}
                    _ = v69
                case "a":
                    v6 += (v6 % ((v6) | 1))
                default:
                    if true {
                        v2 += v2
                        v70 := v1[366]
                        v26 += 18446744073709551615
                        _ = v70
                    } else {
                        if ((!false) || false) {
                            switch "\n\t" {
                            case "\n\t":
                                v71 := v5["hello"]
                                _ = v71
                            }
                            if false {
                                v3 = v3
                            } else {
                                v72 := GenS2{F0: 1, F1: (fi68 << ((454) & 63))}
                                switch v3 {
                                case 4294967295:
                                    switch (v3 ^ fi68) {
                                    case 4294967295:
                                        var v73 uint16 = v7
                                        _ = v73
                                    }
                                case 821:
                                    v7 = (v7 + v7)
                                default:
                                    switch (fi68 * fi68) {
                                    case 432:
                                        var v74 bool = ("foo bar" <= "")
                                        _ = v74
                                    case 0:
                                        for fi75 := 0; fi75 < 16; fi75++ {
                                            var v76 int8 = int8(v4)
                                            if (false && (86 > 176)) {
                                                v7 = v7
                                                v77 := GenS6{F0: (fi75 % ((int(v26)) | 1)), F1: 42, F2: ("a" + "hello")}
                                                var v78 int32 = v2
                                                _ = v77
                                                _ = v78
                                            } else {
                                                if true {
                                                    v4 += (v4 >> ((2) & 63))
                                                } else {
                                                    var v79 uint = 255
                                                    if false {
                                                        if true {
                                                            delete(v5, ("foo bar" + ("hello" + ((("\n\t" + "a") + "foo bar") + ("\n\t" + "hello")))))
                                                        }
                                                        v26 = uint64(v7)
                                                        if (3.14159 > ((1e10 - (-2.25 / 1e10)) * (1.5 * (1e10 / (1e10 - (1.5 / (-2.25 * 1e10))))))) {
                                                            v26 = 0
                                                            if false {
                                                                fmt.Println(fi68)
                                                                var v80 uint64 = v26
                                                                _ = v80
                                                            } else {
                                                                v2 = int32(v26)
                                                                var v81 uint64 = v26
                                                                switch fi75 {
                                                                case -9223372036854775808:
                                                                    v82 := []int{(v3 - (-(fi75))), v3, (fi68 % ((fi75) | 1))}
                                                                    _ = v82
                                                                case 9223372036854775807:
                                                                    fmt.Println(v26, v6, fi68)
                                                                case 32768:
                                                                    if ((-2.25 - ((1e10 - 1.5) / 0.1)) < (((1e10 / 3.14159) - 0.1) - 0.0)) {
                                                                        var v83 int16 = (v4 / ((v4) | 1))
                                                                        v84 := v5["foo bar"]
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                default:
                                                                    fmt.Println(fi75, v7)
                                                                }
                                                                var v85 float32 = -2.25
                                                                _ = v81
                                                                _ = v85
                                                            }
                                                            fmt.Println(v76, fi75)
                                                        }
                                                    }
                                                    var v86 bool = true
                                                    _ = v79
                                                    _ = v86
                                                }
                                                var v87 float64 = (-2.25 * ((3.14159 + 1.5) - 1e10))
                                                _ = v87
                                            }
                                            _ = fi75
                                            _ = v76
                                        }
                                    case 127:
                                        v88 := GenS5{F0: v7, F1: "hello"}
                                        _ = v88
                                    default:
                                        v89 := GenS5{F0: (^(v7)), F1: "hello"}
                                        _ = v89
                                    }
                                }
                                _ = v72
                            }
                        } else {
                            v7 = (v7 & (v7 | v7))
                            for fi90 := 0; fi90 < 2; fi90++ {
                                v91 := GenS7{F0: (^(v26)), F1: (v26 << ((uint(v7)) & 63))}
                                var v92 float64 = 1.5
                                v93 := v1[65535]
                                _ = fi90
                                _ = v91
                                _ = v92
                                _ = v93
                            }
                        }
                    }
                }
                fmt.Println(fi68, v7)
                _ = fi68
            }
            if (-589 > -599) {
                for fi94 := 0; fi94 < 5; fi94++ {
                    for fi95 := 0; fi95 < 1; fi95++ {
                        var v96 int8 = (v6 / ((v6) | 1))
                        switch (((("a" + "\n\t") + (("\n\t" + "foo bar") + (("\n\t" + "hello") + ("" + "")))) + ("foo bar" + "\n\t")) + ("\n\t" + "\n\t")) {
                        case "hello":
                            v4 = (v4 >> ((18446744073709551615) & 63))
                        case "a":
                            switch "\n\t" {
                            case "foo bar":
                                v3 = fi94
                            case "\n\t":
                                fmt.Println(v26, v3)
                            case "":
                                if false {
                                    v96 = v6
                                    v97 := GenS4{F0: v2}
                                    v98 := GenS3{F0: 478, F1: "\n\t"}
                                    _ = v97
                                    _ = v98
                                }
                            default:
                                var v99 uint64 = 689
                                _ = v99
                            }
                        }
                        var v100 float64 = ((3.14159 + float64(1e10)) * (0.0 * (3.14159 - (-2.25 + (0.1 / 1.0)))))
                        _ = fi95
                        _ = v100
                        _ = v96
                    }
                    var v101 uint32 = 42
                    fmt.Println(v7, v101)
                    _ = fi94
                    _ = v101
                }
            }
            var v102 bool = false
            _ = v102
        }
        v103 := []int{(v3 | v3), v3, (v3 >> ((2) & 63)), (v3 & int(v7))}
        _ = v103
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
