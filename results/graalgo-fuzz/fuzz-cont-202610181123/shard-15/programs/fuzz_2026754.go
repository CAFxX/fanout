package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32; F1 bool; F2 uint }
type GenS10 struct { F0 uint64; F1 uint }
type GenS11 struct { F0 string; F1 int16; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 127
    var v3 uint32 = 42
    v4 := make(map[uint32]bool)
    var v5 bool = (4294967295 < 4294967295)
    v6 := v1["foo bar"]
    v3 = (v3 + (v3 ^ (v3 & v3)))
    if (-9223372036854775808 == 32767) {
        v7 := []int{0}
        v8 := GenS1{F0: v6, F1: 32767}
        switch 127 {
        case -618:
            v6 = v5
        case 394:
            var v9 uint64 = 964
            _ = v9
        default:
            v5 = v6
        }
        _ = v7
        _ = v8
    } else {
        if true {
            var v10 uint64 = 1
            v3 -= uint32(v2)
            v11 := GenS1{F0: v6, F1: 1000}
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            v5 = false
            v6 = (1.5 >= float32(3.14159))
            for fi13 := 0; fi13 < 14; fi13++ {
                switch (^(fi12)) {
                case 32768:
                    if (9223372036854775807 <= 0) {
                        v14 := v4[2]
                        _ = len(v1)
                        _ = v14
                    }
                default:
                    v2 = (v2 % ((128) | 1))
                }
                switch "" {
                case "":
                    v1["hello"] = true
                }
                v15 := GenS1{F0: v6, F1: -1}
                _ = fi13
                _ = v15
            }
            _ = fi12
        }
        v16 := []int{32767, 0}
        _ = v16
    }
    if false {
        v5 = v6
        var v17 int64 = 1
        _ = v17
    } else {
        if v6 {
            var v18 uint16 = 65535
            _ = v18
        } else {
            switch 42 {
            case -1:
                if (1 <= 42) {
                    v19 := []int{32767}
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 string = "a"
                        fmt.Println(fi20, v3)
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v6 = v5
                            _ = fi22
                        }
                        _ = fi20
                        _ = v21
                    }
                    delete(v1, "hello")
                }
            case 127:
                switch "a" {
                case "":
                    var v23 float64 = 1e10
                    _ = v23
                case "hello":
                    v6 = (-2147483648 > 0)
                }
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                if v6 {
                    v25 := GenS1{F0: ((-2.25 != 1.5) || ((-(fi24)) >= fi24)), F1: -222}
                    fmt.Println(fi24, v3)
                    if v5 {
                        if (v6 || true) {
                            v6 = false
                            for fi26 := 0; fi26 < 6; fi26++ {
                                switch fi24 {
                                case 593:
                                    if (158 != 684) {
                                        var v27 int64 = 9223372036854775807
                                        if false {
                                            fmt.Println(v27)
                                            if (true || false) {
                                                if (-2147483648 >= int32(v27)) {
                                                    v27 = (v27 % (((v27 + v27)) | 1))
                                                    v28 := GenS1{F0: (18446744073709551615 <= 42), F1: -32768}
                                                    if v5 {
                                                        var v29 bool = v5
                                                        _ = v29
                                                    }
                                                    var v30 float64 = (-2.25 + (((-2.25 - 1.5) * (((1e10 - 1.5) - 1.5) - (1e10 + -2.25))) + 1e10))
                                                    v2 = (v2 % ((127) | 1))
                                                    _ = v28
                                                    _ = v30
                                                } else {
                                                    v2 += 41
                                                    if (1 != -9223372036854775808) {
                                                        switch fi24 {
                                                        case 127:
                                                            var v31 uint64 = 42
                                                            _ = v31
                                                        case 618:
                                                            v32 := GenS2{F0: v3, F1: 0, F2: 42}
                                                            _ = v32
                                                        default:
                                                            var v33 int64 = (v27 + (v27 - -470))
                                                            _ = v33
                                                        }
                                                        var v34 float32 = ((1e10 * 0.1) - 1e10)
                                                        _ = v34
                                                    } else {
                                                        _ = len(v4)
                                                        v35 := GenS2{F0: (v3 % ((v3) | 1)), F1: 69, F2: -671}
                                                        _ = v35
                                                    }
                                                }
                                                if (v3 == v3) {
                                                    fmt.Println(v6, fi26, v3)
                                                    v36 := GenS1{F0: v5, F1: 0}
                                                    var v37 bool = false
                                                    _ = v36
                                                    _ = v37
                                                }
                                                var v38 uint8 = v2
                                                fmt.Println(v6)
                                                v2 = v38
                                                _ = v38
                                            }
                                            var v39 float32 = 0.1
                                            v40 := GenS2{F0: (^(v3)), F1: 127, F2: -731}
                                            _ = v39
                                            _ = v40
                                        }
                                        _ = v27
                                    } else {
                                        switch ("a" + ("a" + "")) {
                                        case "":
                                            v6 = false
                                        case "\n\t":
                                            v41 := []int{4294967295}
                                            _ = v41
                                        case "hello":
                                            v6 = v5
                                            v42 := GenS2{F0: v3, F1: 127, F2: 625}
                                            if false {
                                                v43 := v1["\n\t"]
                                                switch ("\n\t" + "\n\t") {
                                                case "":
                                                    v44 := []int{fi26, fi26}
                                                    _ = v44
                                                case "foo bar":
                                                    switch (fi24 * fi24) {
                                                    case 715:
                                                        var v45 float64 = (0.0 - (((1e10 * 0.1) * 1.5) * (0.0 * -2.25)))
                                                        _ = v45
                                                    case -2147483648:
                                                        v46 := GenS3{F0: 1000}
                                                        _ = v46
                                                    case 0:
                                                        var v47 int16 = -636
                                                        v48 := GenS4{F0: -128, F1: 419}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                default:
                                                    var v49 uint32 = (v3 & v3)
                                                    _ = v49
                                                }
                                                _ = v43
                                            } else {
                                                v5 = false
                                            }
                                            _ = v42
                                        default:
                                            if (!((((v3 << ((65535) & 63)) > 2147483647) || true) && v6)) {
                                                v3 = (v3 + (^(v3)))
                                            } else {
                                                switch fi24 {
                                                case 2147483647:
                                                    v5 = ((77 != 304) && v5)
                                                    v3 -= v3
                                                    var v50 int64 = 1
                                                    if v5 {
                                                        if v5 {
                                                            switch "hello" {
                                                            case "hello":
                                                                var v51 int64 = (v50 + 844)
                                                                _ = v51
                                                            case "":
                                                                if v5 {
                                                                    if false {
                                                                        if (1.5 <= 0.1) {
                                                                            _ = len(v4)
                                                                            var v52 float32 = 0.0
                                                                            v53 := v4[(v3 + v3)]
                                                                            _ = v52
                                                                            _ = v53
                                                                        } else {
                                                                            switch -297 {
                                                                            case 42:
                                                                                var v54 int32 = -2147483648
                                                                                if true {
                                                                                    fmt.Println(v54)
                                                                                    v55 := GenS5{F0: (v3 / ((4294967295) | 1)), F1: ""}
                                                                                    _ = v55
                                                                                } else {
                                                                                    v56 := []int{fi24, (fi26 - int(v3)), -119, int(v3)}
                                                                                    _ = v56
                                                                                }
                                                                                _ = v54
                                                                            default:
                                                                                v57 := []int{fi26, fi26, (fi26 / (((^(fi24))) | 1)), fi24}
                                                                                _ = v57
                                                                            }
                                                                            v4[(^(v3))] = ((^(v2)) != 2)
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v58 string = "foo bar"
                                                                    if v5 {
                                                                        delete(v4, v3)
                                                                        delete(v1, (("foo bar" + v58) + ""))
                                                                        v6 = v5
                                                                    }
                                                                    v6 = (!(((v50 - v50) != (v50 & 0)) || true))
                                                                    _ = v58
                                                                }
                                                            case "a":
                                                                fmt.Println(v2, v5, v6)
                                                            default:
                                                                if (-887 < int(v50)) {
                                                                    v50 = v50
                                                                    var v59 int64 = (v50 & (v50 + v50))
                                                                    v2 = 229
                                                                    _ = v59
                                                                }
                                                            }
                                                            var v60 uint32 = (v3 << ((0) & 63))
                                                            _ = v60
                                                        }
                                                        var v61 int16 = 33
                                                        _ = v61
                                                    } else {
                                                        if (-683 >= -1) {
                                                            var v62 string = ("a" + "foo bar")
                                                            var v63 uint16 = uint16(v3)
                                                            v6 = ((^(v3)) == v3)
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            delete(v1, ("" + "hello"))
                                                        }
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            fmt.Println(v5, v3)
                                                        case "":
                                                            fmt.Println(v5, v6)
                                                        default:
                                                            var v64 uint32 = v3
                                                            switch (("\n\t" + ("hello" + ("foo bar" + "\n\t"))) + ("a" + "\n\t")) {
                                                            case "\n\t":
                                                                var v65 int16 = int16(v50)
                                                                _ = v65
                                                            }
                                                            if true {
                                                                v66 := GenS3{F0: -456}
                                                                _ = v66
                                                            }
                                                            _ = v64
                                                        }
                                                        v4[(v3 + v3)] = v6
                                                        v67 := GenS4{F0: 42, F1: v3}
                                                        switch (fi26 ^ (fi24 & (fi26 << ((1) & 63)))) {
                                                        case -9223372036854775808:
                                                            v5 = ((198 > (v2 >> ((57) & 63))) && (v6 || ("" != ("\n\t" + (((("" + ("a" + ("a" + "a"))) + "\n\t") + "\n\t") + "a")))))
                                                        case 128:
                                                            switch "" {
                                                            case "foo bar":
                                                                switch fi24 {
                                                                case -937:
                                                                    var v68 bool = true
                                                                    _ = v68
                                                                case -9223372036854775808:
                                                                    switch ("a" + (("foo bar" + ("\n\t" + (("hello" + "foo bar") + "a"))) + "hello")) {
                                                                    case "hello":
                                                                        var v69 uint16 = 191
                                                                        _ = v69
                                                                    case "\n\t":
                                                                        v70 := GenS1{F0: (v3 == (v3 & 804)), F1: -1}
                                                                        fmt.Println(v50, fi26, v3)
                                                                        _ = v70
                                                                    default:
                                                                        var v71 uint = 630
                                                                        _ = v71
                                                                    }
                                                                case 4294967295:
                                                                    v2 += 42
                                                                }
                                                            default:
                                                                if v6 {
                                                                    v72 := GenS6{F0: 1000, F1: v50}
                                                                    v73 := GenS7{F0: 748}
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            }
                                                        default:
                                                            var v74 uint = 2
                                                            _ = v74
                                                        }
                                                        _ = v67
                                                    }
                                                    _ = v50
                                                case 789:
                                                    var v75 uint32 = 2147483648
                                                    _ = v75
                                                }
                                                fmt.Println(fi24, v3)
                                            }
                                        }
                                    }
                                case -9223372036854775808:
                                    var v76 bool = (uint8(v3) < v2)
                                    _ = v76
                                case -1:
                                    v77 := GenS6{F0: 112, F1: int64(v2)}
                                    _ = v77
                                }
                                var v78 int = fi24
                                var v79 uint8 = (^(v2))
                                _ = fi26
                                _ = v78
                                _ = v79
                            }
                        } else {
                            for fi80 := 0; fi80 < 20; fi80++ {
                                fmt.Println(v3, v2, v6)
                                v81 := GenS3{F0: -32768}
                                var v82 uint8 = (v2 | 2)
                                _ = fi80
                                _ = v81
                                _ = v82
                            }
                            var v83 bool = (255 > 42)
                            _ = v83
                        }
                        v5 = true
                    }
                    _ = v25
                }
                _ = fi24
            }
        }
        v2 -= v2
        switch 128 {
        case -144:
            v6 = true
        case -2147483648:
            v84 := GenS8{F0: (!v6)}
            _ = v84
        case 0:
            _ = len(v4)
        }
    }
    v85 := GenS9{F0: v3, F1: v5, F2: 18446744073709551615}
    for fi86 := 0; fi86 < 11; fi86++ {
        if false {
            if (v5 || false) {
                v5 = v6
            } else {
                v3 -= (^(v3))
            }
        } else {
            var v87 int = (fi86 << ((0) & 63))
            v87 = v87
            v88 := GenS10{F0: 234, F1: 18446744073709551615}
            _ = v87
            _ = v88
        }
        v89 := []int{}
        _ = fi86
        _ = v89
    }
    switch 340 {
    case 2147483647:
        switch 727 {
        case 42:
            var v90 int32 = 42
            _ = v90
        case -774:
            v91 := []int{2147483647, 32768, 658, 128}
            _ = v91
        default:
            v92 := GenS11{F0: ("" + ("\n\t" + "\n\t")), F1: 0, F2: (32767 == 794)}
            _ = v92
        }
    case -9223372036854775808:
        var v93 int16 = -402
        _ = v93
    }
    fmt.Println(v6, v3, v3)
    v5 = ("a" < (("hello" + "a") + "hello"))
    var v94 bool = (-2.25 >= (1e10 - 3.14159))
    v95 := GenS2{F0: 2, F1: -128, F2: -494}
    v96 := []int{762, 32768}
    var v97 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v94:", v94)
    fmt.Println("slicelen v96:", len(v96), v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
    _ = v94
    _ = v95
    _ = v96
    _ = v97
}
