package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 int8; F2 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 bool; F2 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int; F1 uint32; F2 int8 }
type GenS14 struct { F0 int64; F1 bool; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 9223372036854775807
    v3 := make(map[string]uint32)
    var v4 bool = false
    v5 := GenS1{F0: 255, F1: v2, F2: 2}
    v6 := GenS1{F0: 1, F1: (^(v2)), F2: 42}
    v2 = 1
    switch "\n\t" {
    case "foo bar":
        v7 := GenS1{F0: 255, F1: (v2 - -2147483648), F2: 4294967295}
        _ = v7
    case "a":
        switch 9223372036854775807 {
        case 904:
            for fi8 := 0; fi8 < 10; fi8++ {
                if v4 {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        v2 = (fi9 % (((fi9 ^ (fi8 | fi9))) | 1))
                        v10 := v1["foo bar"]
                        if false {
                            v2 += fi8
                        } else {
                            if v4 {
                                switch v2 {
                                case -893:
                                    fmt.Println(fi8)
                                default:
                                    v11 := []int{}
                                    _ = v11
                                }
                                switch "\n\t" {
                                case "a":
                                    v12 := GenS1{F0: 0, F1: fi9, F2: v10}
                                    _ = v12
                                default:
                                    if v4 {
                                        if (791 > (v10 - (v10 >> ((42) & 63)))) {
                                            fmt.Println(v4, v10, v10)
                                            v4 = true
                                        }
                                    } else {
                                        var v13 int64 = -9223372036854775808
                                        _ = v13
                                    }
                                }
                            }
                            var v14 uint8 = 42
                            _ = v14
                        }
                        _ = fi9
                        _ = v10
                    }
                }
                _ = fi8
            }
        case 0:
            if (2147483647 < 2147483647) {
                var v15 int = -369
                _ = v15
            } else {
                v4 = (-1 != 42)
                var v16 int = 2147483647
                _ = v16
            }
        default:
            v4 = v4
        }
    }
    if v4 {
        if (!(0 < 690)) {
            v17 := GenS2{F0: "foo bar", F1: 42}
            v18 := []int{(v2 | 127), v2}
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: 42, F1: v2, F2: 943}
            _ = v19
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            if (!(18446744073709551615 > 42)) {
                v21 := GenS3{F0: 853, F1: 515}
                v2 -= (v2 * 1)
                _ = v21
            } else {
                var v22 string = "a"
                v4 = v4
                var v23 uint64 = 0
                _ = v22
                _ = v23
            }
            if ((747 == 298) || v4) {
                if ((32767 != (fi20 + 0)) || true) {
                    v24 := []int{926, fi20, (-(fi20))}
                    v2 = (fi20 | 2147483647)
                    switch -9223372036854775808 {
                    case -304:
                        fmt.Println(v4)
                    case -1:
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v2, fi20)
                        case "a":
                            v2 += 32768
                        case "":
                            for fi25 := 0; fi25 < 1; fi25++ {
                                var v26 bool = false
                                if v26 {
                                    fmt.Println(v26, fi25)
                                    switch (-(v2)) {
                                    case -1:
                                        var v27 int32 = -1
                                        if false {
                                            switch "a" {
                                            case "hello":
                                                switch ("foo bar" + ("" + "a")) {
                                                case "a":
                                                    v28 := GenS4{F0: 42}
                                                    _ = v28
                                                case "":
                                                    v29 := GenS5{F0: (("foo bar" + ("a" + ("\n\t" + "hello"))) + ("hello" + "\n\t"))}
                                                    _ = v29
                                                case "hello":
                                                    fmt.Println(fi20, fi20, v2)
                                                }
                                                if (!(v4 || true)) {
                                                    v27 += (v27 >> ((0) & 63))
                                                } else {
                                                    v27 = (v27 ^ v27)
                                                    if len(v24) > 0 {
                                                        v24[0] = (v2 / ((v2) | 1))
                                                    }
                                                    v30 := v3[(("a" + "foo bar") + "foo bar")]
                                                    _ = v30
                                                }
                                            case "":
                                                v4 = (v4 && (!(true || (18446744073709551615 > 18446744073709551615))))
                                            }
                                            var v31 int = -2147483648
                                            switch -9223372036854775808 {
                                            case -1:
                                                v27 = (v27 & (v27 - v27))
                                            case 9223372036854775807:
                                                fmt.Println(v4, v27, v26)
                                            default:
                                                v26 = (1 <= 25)
                                            }
                                            _ = v31
                                        }
                                        _ = len(v3)
                                        _ = v27
                                    }
                                } else {
                                    v32 := GenS5{F0: ("hello" + ("hello" + "foo bar"))}
                                    if v26 {
                                        v33 := GenS5{F0: "\n\t"}
                                        _ = v33
                                    }
                                    _ = v32
                                }
                                v34 := GenS6{F0: (true && (!v26)), F1: 57, F2: 612}
                                var v35 int32 = 2147483647
                                _ = fi25
                                _ = v26
                                _ = v34
                                _ = v35
                            }
                        default:
                            fmt.Println(v2, v2)
                        }
                    default:
                        if v4 {
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                        } else {
                            for fi36 := 0; fi36 < 2; fi36++ {
                                var v37 uint = 18446744073709551615
                                switch "hello" {
                                case "":
                                    if len(v24) > 0 {
                                        v24[0] = (fi36 - (fi36 ^ (v2 % ((-437) | 1))))
                                    }
                                }
                                v37 = 746
                                _ = fi36
                                _ = v37
                            }
                            v38 := v3[("a" + ("\n\t" + "foo bar"))]
                            _ = v38
                        }
                    }
                    _ = v24
                }
                v2 = fi20
            } else {
                v39 := GenS7{F0: 9223372036854775807}
                fmt.Println(v2, v2)
                _ = v39
            }
            _ = fi20
        }
    }
    if true {
        if v4 {
            var v40 float64 = ((((0.0 * 1e10) + (3.14159 + 1.5)) * 0.0) - (0.1 + (1e10 - float64(0.1))))
            var v41 uint64 = 4294967295
            switch (-(v2)) {
            case 810:
                v41 -= 2
            default:
                v4 = (v4 || v4)
            }
            _ = v40
            _ = v41
        } else {
            if (0 == 128) {
                if v4 {
                    v4 = ("hello" == ("" + "foo bar"))
                }
                switch 502 {
                case 2147483647:
                    switch "\n\t" {
                    case "a":
                        switch "a" {
                        case "hello":
                            var v42 int = (-(v2))
                            _ = v42
                        case "":
                            var v43 uint64 = 463
                            _ = v43
                        case "\n\t":
                            v44 := []int{(v2 - (v2 | (v2 % ((v2) | 1)))), 128, (v2 - (v2 >> ((2) & 63))), 208}
                            _ = v44
                        }
                    }
                default:
                    for fi45 := 0; fi45 < 1; fi45++ {
                        if true {
                            for fi46 := 0; fi46 < 17; fi46++ {
                                switch 128 {
                                case 0:
                                    if false {
                                        switch "foo bar" {
                                        case "hello":
                                            v47 := []int{(^(fi45)), (fi46 % (((fi45 ^ v2)) | 1))}
                                            _ = v47
                                        case "":
                                            v48 := GenS6{F0: false, F1: 0, F2: 255}
                                            _ = v48
                                        }
                                        _ = len(v3)
                                        v2 -= (-(v2))
                                    } else {
                                        v4 = (v2 > v2)
                                    }
                                case 4294967295:
                                    v49 := GenS8{F0: 2147483647, F1: -2147483648}
                                    _ = v49
                                }
                                var v50 int16 = 42
                                switch (("foo bar" + (("" + "\n\t") + "foo bar")) + "") {
                                case "foo bar":
                                    v3[("foo bar" + "\n\t")] = 42
                                    v4 = v4
                                }
                                _ = fi46
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 8; fi51++ {
                                var v52 int = fi45
                                if v4 {
                                    if ((v4 && true) || false) {
                                        fmt.Println(v2, v4, fi51)
                                        if v4 {
                                            v4 = (1e10 < 3.14159)
                                            v53 := []int{}
                                            switch (v2 * (v52 + fi51)) {
                                            case 1:
                                                switch v52 {
                                                case 32767:
                                                    switch ("foo bar" + "\n\t") {
                                                    case "foo bar":
                                                        v54 := GenS3{F0: 4294967295, F1: 4294967295}
                                                        _ = v54
                                                    default:
                                                        v2 = 4294967295
                                                    }
                                                case -179:
                                                    var v55 uint16 = 1
                                                    _ = v55
                                                case -502:
                                                    var v56 uint = 255
                                                    _ = v56
                                                default:
                                                    var v57 uint16 = 65535
                                                    _ = v57
                                                }
                                            default:
                                                v58 := GenS9{F0: 861, F1: v52}
                                                _ = v58
                                            }
                                            _ = v53
                                        }
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v2)
                                        case "\n\t":
                                            v59 := GenS10{F0: 1000}
                                            _ = v59
                                        default:
                                            var v60 int8 = 5
                                            v2 -= fi45
                                            v61 := []int{}
                                            var v62 float32 = float32(-2.25)
                                            _ = v60
                                            _ = v61
                                            _ = v62
                                        }
                                    } else {
                                        if (!false) {
                                            var v63 uint8 = 132
                                            _ = v63
                                        } else {
                                            switch ("" + ("a" + "a")) {
                                            case "hello":
                                                v64 := GenS6{F0: ((!v4) && v4), F1: int8(v52), F2: 663}
                                                _ = v64
                                            }
                                            fmt.Println(v2, v4, fi45)
                                            v52 = (^(fi45))
                                            v65 := GenS11{F0: 1, F1: v4, F2: 164}
                                            var v66 uint16 = 20
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                    if ((!v4) || (fi51 == 127)) {
                                        v67 := v1[""]
                                        v68 := []int{-969, 436}
                                        _ = v67
                                        _ = v68
                                    }
                                    fmt.Println(v2, fi45, v4)
                                }
                                _ = fi51
                                _ = v52
                            }
                            switch (((("\n\t" + "a") + "hello") + "") + "\n\t") {
                            case "a":
                                v69 := []int{(v2 ^ 42), fi45}
                                _ = v69
                            default:
                                switch "hello" {
                                case "a":
                                    v2 = 846
                                default:
                                    switch fi45 {
                                    case 127:
                                        var v70 uint = 18446744073709551615
                                        _ = v70
                                    }
                                }
                            }
                        } else {
                            for fi71 := 0; fi71 < 11; fi71++ {
                                v72 := v3["a"]
                                v4 = (false && false)
                                var v73 float64 = -2.25
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                            v74 := GenS1{F0: 1, F1: (^(v2)), F2: 4294967295}
                            _ = v74
                        }
                        for fi75 := 0; fi75 < 8; fi75++ {
                            if (255 == 0) {
                                if ((!true) && v4) {
                                    if v4 {
                                        var v76 string = "a"
                                        _ = v76
                                    }
                                } else {
                                    var v77 string = "foo bar"
                                    v4 = (!((^(fi75)) <= (fi45 >> ((341) & 63))))
                                    var v78 uint64 = 538
                                    _ = v77
                                    _ = v78
                                }
                            }
                            var v79 float64 = 1e10
                            var v80 uint64 = 0
                            _ = fi75
                            _ = v79
                            _ = v80
                        }
                        _ = fi45
                    }
                }
            }
            switch v2 {
            case -213:
                if (42 < -1) {
                    v2 = v2
                    fmt.Println(v4)
                } else {
                    v81 := GenS12{F0: -695}
                    var v82 uint = 255
                    v2 -= v2
                    _ = v81
                    _ = v82
                }
            case -470:
                v83 := GenS13{F0: v2, F1: 2, F2: -1}
                _ = v83
            case -742:
                var v84 uint16 = 42
                _ = v84
            }
            var v85 uint8 = 0
            _ = v85
        }
        if ((-(v2)) != 127) {
            v2 = v2
            v2 = (v2 | v2)
        } else {
            if false {
                if true {
                    var v86 bool = v4
                    if false {
                        if false {
                            v2 -= -938
                            var v87 int16 = -32768
                            var v88 uint8 = 128
                            _ = v87
                            _ = v88
                        }
                        v89 := []int{774, v2, 4294967295}
                        _ = v89
                    } else {
                        v2 = (v2 | -2147483648)
                        var v90 int8 = 100
                        v91 := GenS14{F0: -440, F1: v4, F2: 1}
                        _ = v90
                        _ = v91
                    }
                    _ = v86
                } else {
                    switch "foo bar" {
                    case "":
                        delete(v3, "a")
                    case "a":
                        if v4 {
                            fmt.Println(v4)
                            v4 = true
                        }
                    }
                    fmt.Println(v4)
                    var v92 bool = ((v4 && (-1 <= 515)) || false)
                    _ = v92
                }
            } else {
                if v4 {
                    v4 = v4
                    var v93 string = ""
                    fmt.Println(v4, v2)
                    _ = v93
                }
                fmt.Println(v2, v4)
                _ = len(v1)
            }
            v94 := GenS13{F0: 9223372036854775807, F1: 2147483647, F2: -55}
            fmt.Println(v4, v2)
            _ = v94
        }
        var v95 uint64 = 756
        _ = v95
    } else {
        v4 = v4
        switch v2 {
        case 127:
            for fi96 := 0; fi96 < 1; fi96++ {
                v4 = true
                v97 := GenS6{F0: (2147483647 >= 1), F1: -128, F2: 65535}
                v98 := []int{}
                _ = fi96
                _ = v97
                _ = v98
            }
        }
    }
    v99 := []int{v2, v2, (v2 << ((1) & 63)), v2}
    v2 = (-(v2))
    var v100 int = v2
    if v4 {
        if v4 {
            v2 = v100
            v101 := v3["\n\t"]
            var v102 uint64 = 1
            _ = v101
            _ = v102
        } else {
            var v103 uint64 = 2
            _ = v103
        }
        if v4 {
            v104 := GenS4{F0: 36}
            var v105 float32 = 1e10
            _ = v104
            _ = v105
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v99:", len(v99), v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
