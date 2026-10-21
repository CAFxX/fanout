package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint32; F1 bool; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int }
func genFn0(p0 int32, p1 uint) int8 {
    v5 := []int{}
    p0 = 2147483647
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 -= (p0 % ((p0) | 1))
        v7 := GenS1{F0: (p1 + p1), F1: "", F2: 0}
        if false {
            if (!true) {
                for fi8 := 0; fi8 < 15; fi8++ {
                    if (1 < -9223372036854775808) {
                        var v9 int8 = -71
                        _ = v9
                    } else {
                        v10 := GenS1{F0: p1, F1: "foo bar", F2: 1000}
                        if len(v5) > 0 {
                            v5[0] = fi8
                        }
                        p0 += (p0 | (p0 * 0))
                        _ = v10
                    }
                    var v11 uint32 = 17
                    var v12 float64 = (-2.25 * 0.0)
                    v13 := GenS1{F0: (^(p1)), F1: (("a" + ("a" + ("hello" + ("" + ("foo bar" + ("a" + "\n\t")))))) + ("hello" + ("" + "hello"))), F2: -32768}
                    _ = fi8
                    _ = v11
                    _ = v12
                    _ = v13
                }
            } else {
                p1 = p1
                _ = len(v5)
                switch fi6 {
                case 42:
                    if true {
                        p1 = (p1 / (((p1 << ((900) & 63))) | 1))
                        p1 += (p1 | p1)
                    }
                case -347:
                    switch "foo bar" {
                    case "\n\t":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                default:
                    var v14 float32 = -2.25
                    _ = v14
                }
            }
        } else {
            fmt.Println(fi6, p1, p0)
        }
        _ = fi6
        _ = v7
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        p1 = (p1 / ((p1) | 1))
        switch (((("a" + "hello") + "hello") + "a") + ("a" + "foo bar")) {
        case "a":
            fmt.Println(fi15, p0)
        }
        var v16 uint8 = 128
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 float32 = ((1e10 + 0.0) / (0.1 + 3.14159))
    var v4 int32 = 0
    var v17 uint16 = 883
    v18 := GenS1{F0: 644, F1: "foo bar", F2: 1}
    v3 = -2.25
    switch "a" {
    case "hello":
        if (-874 > -1) {
            var v19 uint8 = 42
            _ = v19
        }
    default:
        if (255 <= 0) {
            for fi20 := 0; fi20 < 20; fi20++ {
                if (3.14159 < -2.25) {
                    v4 = int32(v17)
                    if false {
                        fmt.Println(fi20, fi20, v3)
                    } else {
                        switch fi20 {
                        case 972:
                            v21 := GenS2{F0: (^(v17)), F1: ("\n\t" + "hello"), F2: -105}
                            _ = v21
                        case 485:
                            switch (fi20 + (fi20 | -828)) {
                            case 4294967295:
                                var v22 uint64 = 0
                                _ = v22
                            case 32767:
                                var v23 int64 = -9223372036854775808
                                _ = v23
                            case -2147483648:
                                var v24 int64 = 9223372036854775807
                                _ = v24
                            default:
                                switch int(v4) {
                                case -290:
                                    switch "foo bar" {
                                    case "a":
                                        v25 := GenS3{F0: (fi20 | fi20), F1: -29}
                                        _ = v25
                                    case "foo bar":
                                        v26 := GenS4{F0: (fi20 % ((-1) | 1)), F1: 65535}
                                        _ = v26
                                    }
                                case -9223372036854775808:
                                    var v27 int16 = 42
                                    _ = v27
                                }
                            }
                        }
                    }
                } else {
                    if (100 == 1) {
                        v28 := GenS3{F0: (-(fi20)), F1: genFn0((v4 - v4), 211)}
                        _ = v28
                    }
                }
                switch "foo bar" {
                case "":
                    var v29 float32 = v3
                    _ = v29
                }
                _ = fi20
            }
            v30 := GenS4{F0: -504, F1: 4294967295}
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 int64 = 0
        if (0 != 711) {
            v2["a"] = true
            switch "foo bar" {
            case "\n\t":
                if ((!true) || (42 < v17)) {
                    if ((v17 ^ (v17 | 182)) >= v17) {
                        if false {
                            var v33 uint32 = 2147483648
                            fmt.Println(fi31, fi31, v32)
                            fmt.Println(v3)
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 18; fi34++ {
                            switch -9223372036854775808 {
                            case 0:
                                var v35 int16 = -600
                                _ = v35
                            case 207:
                                v32 = (v32 / (((v32 & 1)) | 1))
                            default:
                                var v36 int32 = (v4 - v4)
                                _ = v36
                            }
                            if false {
                                v4 = (-(v4))
                                var v37 uint64 = 18446744073709551615
                                if false {
                                    var v38 int32 = v4
                                    if (false || (!false)) {
                                        v39 := GenS5{F0: 1, F1: false, F2: (v38 ^ 42)}
                                        var v40 uint16 = v17
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v4 += v38
                                    }
                                    if (!(!true)) {
                                        v3 = -2.25
                                        fmt.Println(v4)
                                    } else {
                                        switch 0 {
                                        case 32768:
                                            var v41 float32 = ((v3 * (0.0 / (v3 - (float32((3.14159 - ((0.1 + (3.14159 / (1.5 + (3.14159 + 0.1)))) / 1.0))) / v3)))) / v3)
                                            _ = v41
                                        case 32767:
                                            delete(v2, "foo bar")
                                        }
                                        v42 := v2["foo bar"]
                                        _ = v42
                                    }
                                    _ = v38
                                } else {
                                    var v43 bool = false
                                    fmt.Println(v3, v43, v43)
                                    var v44 uint64 = v37
                                    _ = v43
                                    _ = v44
                                }
                                if ("a" < "hello") {
                                    fmt.Println(fi34)
                                    switch int(v37) {
                                    case -532:
                                        v32 = v32
                                    default:
                                        v45 := GenS3{F0: (fi34 / ((fi31) | 1)), F1: 127}
                                        _ = v45
                                    }
                                    _ = len(v1)
                                    v46 := GenS2{F0: (v17 / ((811) | 1)), F1: "\n\t", F2: -1}
                                    v17 = v17
                                    if (!false) {
                                        var v47 int32 = int32(v37)
                                        v48 := GenS4{F0: (fi31 >> ((42) & 63)), F1: 65535}
                                        _ = v47
                                        _ = v48
                                    } else {
                                        var v49 uint16 = v17
                                        fmt.Println(v3, v37)
                                        if (true && (204 != 32767)) {
                                            if true {
                                                v17 += v49
                                                switch (fi34 ^ (fi34 | 1)) {
                                                case 678:
                                                    _ = len(v1)
                                                case 127:
                                                    var v50 float32 = (float32(0.1) - v3)
                                                    _ = v50
                                                }
                                                if (v17 < 1) {
                                                    v49 = 2
                                                    if false {
                                                        var v51 float64 = (1.5 - 0.0)
                                                        _ = v51
                                                    }
                                                    var v52 float32 = (1e10 * 0.0)
                                                    v49 += 32767
                                                    var v53 uint = 1
                                                    v17 = v17
                                                    _ = v52
                                                    _ = v53
                                                }
                                            } else {
                                                v17 = (v17 | v49)
                                                v32 = -337
                                            }
                                        } else {
                                            var v54 bool = ((!(!(18446744073709551615 == 579))) && true)
                                            _ = v54
                                        }
                                        _ = v49
                                    }
                                    var v55 string = (("a" + "\n\t") + "\n\t")
                                    _ = v46
                                    _ = v55
                                } else {
                                    fmt.Println(v32, v4, v3)
                                    if true {
                                        v4 = (v4 & 378)
                                        switch "" {
                                        case "hello":
                                            var v56 uint16 = (v17 ^ v17)
                                            _ = v56
                                        case "\n\t":
                                            if ((v37 << ((4294967295) & 63)) < uint64(v17)) {
                                                if (true && false) {
                                                    if false {
                                                        var v57 uint8 = 2
                                                        v58 := v2[""]
                                                        switch "\n\t" {
                                                        case "hello":
                                                            switch "hello" {
                                                            case "a":
                                                                v59 := []int{(fi34 << ((42) & 63)), (fi31 | 659)}
                                                                _ = v59
                                                            case "":
                                                                var v60 int32 = (v4 / (((v4 & v4)) | 1))
                                                                _ = v60
                                                            }
                                                        case "foo bar":
                                                            v4 = v4
                                                        }
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        v3 -= 0.1
                                                        fmt.Println(fi34)
                                                        switch (fi34 * fi34) {
                                                        case -2147483648:
                                                            v61 := []int{fi31}
                                                            _ = v61
                                                        case 928:
                                                            if (false && ("\n\t" <= "\n\t")) {
                                                                var v62 uint = 42
                                                                _ = v62
                                                            } else {
                                                                v32 -= 0
                                                                v17 = (v17 >> ((18446744073709551615) & 63))
                                                            }
                                                        default:
                                                            fmt.Println(fi31, v17)
                                                        }
                                                    }
                                                    v37 += 1
                                                    switch fi34 {
                                                    case 4294967295:
                                                        var v63 int32 = (v4 ^ 1)
                                                        _ = v63
                                                    case -9223372036854775808:
                                                        v64 := GenS2{F0: (v17 ^ (v17 | v17)), F1: "a", F2: 41}
                                                        _ = v64
                                                    default:
                                                        fmt.Println(fi34, fi34, v37)
                                                    }
                                                } else {
                                                    fmt.Println(fi31)
                                                }
                                                v3 += ((v3 + -2.25) * float32((0.0 - float64(v3))))
                                            }
                                        case "":
                                            var v65 uint16 = 320
                                            _ = v65
                                        }
                                    } else {
                                        var v66 string = ("\n\t" + ("hello" + "\n\t"))
                                        v32 -= -1
                                        switch fi31 {
                                        case 4294967295:
                                            v37 += (v37 / (((^(v37))) | 1))
                                        case 128:
                                            var v67 int = -1
                                            _ = v67
                                        }
                                        _ = v66
                                    }
                                    var v68 int64 = (v32 * -678)
                                    v68 -= (v32 + v32)
                                    _ = v68
                                }
                                _ = v37
                            } else {
                                var v69 uint16 = v17
                                var v70 uint32 = 65535
                                if false {
                                    if (3.14159 != 0.1) {
                                        v3 = (v3 / v3)
                                        var v71 int = fi34
                                        _ = v71
                                    } else {
                                        fmt.Println(fi34, v69, v69)
                                    }
                                    v69 += v69
                                    v72 := GenS6{F0: 45}
                                    _ = v72
                                }
                                _ = v69
                                _ = v70
                            }
                            var v73 int = 32767
                            _ = fi34
                            _ = v73
                        }
                        if ((true || true) && false) {
                            var v74 uint64 = 4294967295
                            delete(v1, 0)
                            v75 := GenS3{F0: -9223372036854775808, F1: 1}
                            _ = v74
                            _ = v75
                        }
                    } else {
                        if true {
                            if true {
                                var v76 float64 = (float64(float32((1e10 - -2.25))) - 0.0)
                                v32 = (-(v32))
                                for fi77 := 0; fi77 < 16; fi77++ {
                                    v78 := []int{fi77, (^(fi31))}
                                    _ = fi77
                                    _ = v78
                                }
                                _ = v76
                            } else {
                                for fi79 := 0; fi79 < 11; fi79++ {
                                    v17 = (v17 | v17)
                                    v17 = (^(v17))
                                    _ = fi79
                                }
                                v3 = v3
                                v2[(("hello" + "foo bar") + ("" + ("a" + "a")))] = false
                            }
                        } else {
                            if true {
                                if (-985 >= 2147483647) {
                                    if false {
                                        for fi80 := 0; fi80 < 3; fi80++ {
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v17)
                                            case "hello":
                                                var v81 uint16 = (v17 % (((v17 * 542)) | 1))
                                                _ = v81
                                            case "":
                                                var v82 int32 = 839
                                                _ = v82
                                            default:
                                                fmt.Println(fi80)
                                            }
                                            if true {
                                                var v83 int32 = v4
                                                var v84 uint16 = v17
                                                v17 -= (^(v84))
                                                _ = v83
                                                _ = v84
                                            }
                                            _ = fi80
                                        }
                                        if ((2 < 0) || ("hello" >= "a")) {
                                            v85 := GenS7{F0: (^(fi31)), F1: (-(fi31))}
                                            var v86 int16 = 1000
                                            _ = v85
                                            _ = v86
                                        }
                                    }
                                } else {
                                    v3 = v3
                                }
                                for fi87 := 0; fi87 < 7; fi87++ {
                                    var v88 int = fi87
                                    v32 = v32
                                    var v89 uint64 = uint64(v17)
                                    _ = fi87
                                    _ = v88
                                    _ = v89
                                }
                            } else {
                                for fi90 := 0; fi90 < 14; fi90++ {
                                    var v91 int64 = 1
                                    v32 -= v32
                                    _ = fi90
                                    _ = v91
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                var v92 uint32 = 143
                                _ = v92
                            case "\n\t":
                                v93 := GenS3{F0: (fi31 + (fi31 & fi31)), F1: -6}
                                _ = v93
                            }
                            if false {
                                if ((44 < 799) || false) {
                                    var v94 uint16 = 42
                                    v95 := GenS1{F0: 0, F1: "\n\t", F2: -1}
                                    _ = v94
                                    _ = v95
                                } else {
                                    var v96 string = ("\n\t" + ((("hello" + "\n\t") + (("a" + "\n\t") + (("a" + "") + ""))) + ""))
                                    fmt.Println(v3)
                                    fmt.Println(v4, v17, v3)
                                    _ = v96
                                }
                                if false {
                                    for fi97 := 0; fi97 < 10; fi97++ {
                                        v98 := GenS1{F0: 255, F1: ("" + ("" + ("foo bar" + ("\n\t" + "")))), F2: 933}
                                        _ = fi97
                                        _ = v98
                                    }
                                }
                            }
                        }
                        if ((v17 % ((801) | 1)) > (v17 & v17)) {
                            var v99 int = fi31
                            for fi100 := 0; fi100 < 1; fi100++ {
                                var v101 uint32 = 2147483648
                                _ = fi100
                                _ = v101
                            }
                            _ = v99
                        } else {
                            for fi102 := 0; fi102 < 12; fi102++ {
                                var v103 float64 = ((1.5 * 3.14159) - 1e10)
                                _ = fi102
                                _ = v103
                            }
                        }
                    }
                    fmt.Println(v17, v32)
                    v104 := GenS5{F0: 65535, F1: ((29 < 1) || false), F2: v4}
                    _ = v104
                }
            case "foo bar":
                var v105 float32 = 1e10
                _ = v105
            case "":
                v106 := GenS6{F0: -1}
                _ = v106
            }
        }
        _ = fi31
        _ = v32
    }
    if (!((32767 >= -1) || false)) {
        fmt.Println(v17, v3)
        v3 -= 0.0
    }
    v3 += float32(-2.25)
    v107 := []int{-9223372036854775808, 42}
    v108 := GenS4{F0: 2147483647, F1: 892}
    var v109 uint16 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v107:", len(v107), v107)
    fmt.Println("v109:", v109)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
