package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint16; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 int8; F2 string }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 uint64; F1 int64; F2 string }
func genFn0(p0 int, p1 int32, p2 int16) int32 {
    p1 += p1
    p0 += 9223372036854775807
    p0 = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    return (p1 + p1)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 0
    var v3 int64 = 0
    v4 := []uint32{}
    var v5 int16 = -1
    v5 = v5
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint32 = 65535
        _ = fi6
        _ = v7
    }
    switch "a" {
    case "a":
        if ((((3.14159 - -2.25) / float64((1.5 / float32(1e10)))) / (0.1 - ((0.1 * 0.1) + 0.0))) > -2.25) {
            v2 += v5
        }
    default:
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(v3, v2, fi9)
                _ = fi9
            }
            var v10 int = fi8
            _ = fi8
            _ = v10
        }
    }
    fmt.Println(v5)
    var v11 float64 = 0.0
    v12 := GenS1{F0: -9223372036854775808, F1: -2147483648}
    switch 127 {
    case 4294967295:
        for fi13 := 0; fi13 < 14; fi13++ {
            if (42 != 23) {
                var v14 int32 = -1
                var v15 float64 = 0.0
                var v16 int32 = v14
                _ = v14
                _ = v15
                _ = v16
            }
            if (((2 < 2147483647) && (!true)) || (!false)) {
                v17 := GenS2{F0: 2}
                for fi18 := 0; fi18 < 17; fi18++ {
                    if true {
                        v2 -= (v5 + (v2 * (v5 % (((v2 | -1)) | 1))))
                    }
                    var v19 uint64 = 603
                    v5 += v2
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                var v20 int8 = 1
                if (!true) {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if false {
                            var v22 float32 = (float32(((v11 / v11) / (v11 / (1e10 * 1e10)))) * 0.0)
                            _ = len(v1)
                            v20 -= v20
                            _ = v22
                        } else {
                            if false {
                                if false {
                                    if (!false) {
                                        if false {
                                            v2 = v5
                                            var v23 uint32 = 0
                                            fmt.Println(v20, v11, v20)
                                            v11 = 0.1
                                            v3 += v3
                                            _ = v23
                                        } else {
                                            var v24 string = ("a" + "")
                                            v2 += (v5 % ((-32768) | 1))
                                            fmt.Println(v5, v2)
                                            _ = v24
                                        }
                                        if ((3.14159 + v11) < (v11 / 1e10)) {
                                            if ((!(false || true)) || (v11 == v11)) {
                                                v11 += 1.5
                                                v25 := GenS3{F0: "hello"}
                                                _ = v25
                                            }
                                            v1[(fi13 * fi13)] = ((true && (genFn0(973, 42, (v5 | (v5 >> ((2) & 63)))) > genFn0(fi13, genFn0(fi13, genFn0(fi21, 1, (v5 ^ v5)), (v5 >> ((646) & 63))), 1))) && false)
                                            if (!(!(844 == 0))) {
                                                fmt.Println(v2, v11, v11)
                                                var v26 float64 = v11
                                                _ = v26
                                            } else {
                                                v27 := GenS4{F0: (v5 | 32767)}
                                                var v28 uint = 2
                                                fmt.Println(v28)
                                                _ = v27
                                                _ = v28
                                            }
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = 623
                                            }
                                        }
                                        fmt.Println(v5, v2)
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = 1
                                        }
                                        var v29 float32 = 0.0
                                        if true {
                                            _ = len(v1)
                                        }
                                        _ = v29
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if true {
                                        v30 := GenS3{F0: ""}
                                        var v31 int64 = v3
                                        switch (("foo bar" + ("hello" + "foo bar")) + (("foo bar" + "hello") + "hello")) {
                                        case "\n\t":
                                            var v32 int32 = 2147483647
                                            _ = v32
                                        case "":
                                            fmt.Println(v11, v5, v11)
                                        }
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            } else {
                                if (42 > -658) {
                                    var v33 uint32 = 42
                                    v5 -= (v2 ^ 1000)
                                    if (127 > (^(fi21))) {
                                        var v34 float32 = 1.5
                                        if (4294967295 >= 2) {
                                            v35 := GenS3{F0: "a"}
                                            v5 = (v2 | (v5 ^ 42))
                                            _ = v35
                                        }
                                        var v36 float32 = (((((v34 + v34) / float32(v11)) * (v34 - -2.25)) / float32(-2.25)) - 0.0)
                                        _ = v34
                                        _ = v36
                                    }
                                    _ = v33
                                } else {
                                    var v37 int32 = -2147483648
                                    var v38 string = ("\n\t" + ("a" + ("a" + "foo bar")))
                                    _ = v37
                                    _ = v38
                                }
                                fmt.Println(v11)
                                _ = len(v1)
                            }
                            v39 := GenS3{F0: "a"}
                            _ = v39
                        }
                        switch (fi13 - 4294967295) {
                        case 1:
                            v2 += (v2 >> ((48) & 63))
                        default:
                            var v40 uint32 = 2147483648
                            _ = v40
                        }
                        switch fi21 {
                        case 1:
                            v41 := GenS3{F0: "foo bar"}
                            _ = v41
                        case 32768:
                            var v42 uint16 = 42
                            _ = v42
                        case 2147483647:
                            switch "a" {
                            case "foo bar":
                                var v43 float32 = (1e10 - 3.14159)
                                _ = v43
                            case "\n\t":
                                var v44 bool = true
                                v3 += -1
                                _ = v44
                            case "a":
                                _ = len(v4)
                            }
                            fmt.Println(v5)
                            delete(v1, int(v2))
                        }
                        _ = fi21
                    }
                }
                if true {
                    delete(v1, -20)
                } else {
                    if true {
                        var v45 int = 42
                        if (255 == 4294967295) {
                            v46 := GenS3{F0: "a"}
                            fmt.Println(v2, v3)
                            v45 += (v45 - fi13)
                            _ = v46
                        } else {
                            switch "a" {
                            case "hello":
                                fmt.Println(v2, v11)
                            case "foo bar":
                                var v47 uint32 = 1
                                _ = v47
                            case "\n\t":
                                var v48 int32 = -2147483648
                                _ = v48
                            default:
                                if (("\n\t" + "foo bar") <= (("\n\t" + ("hello" + "a")) + "foo bar")) {
                                    if (false || (!((!true) && (float64(1.5) != v11)))) {
                                        _ = len(v1)
                                        v3 += v3
                                    } else {
                                        for fi49 := 0; fi49 < 2; fi49++ {
                                            v11 = 0.0
                                            fmt.Println(v3, v3, v45)
                                            _ = fi49
                                        }
                                    }
                                } else {
                                    switch "a" {
                                    case "hello":
                                        v50 := GenS5{F0: 386, F1: 301, F2: "hello"}
                                        _ = v50
                                    case "":
                                        var v51 int8 = 42
                                        _ = v51
                                    case "a":
                                        var v52 string = "\n\t"
                                        _ = v52
                                    }
                                    for fi53 := 0; fi53 < 11; fi53++ {
                                        var v54 float64 = (v11 * v11)
                                        if (true || true) {
                                            _ = len(v1)
                                            if ((v20 == (v20 / ((100) | 1))) && false) {
                                                fmt.Println(fi13, v20, v11)
                                                _ = len(v1)
                                                v55 := GenS2{F0: 42}
                                                _ = v55
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch fi13 {
                                                case 715:
                                                    var v56 uint64 = 42
                                                    _ = v56
                                                }
                                                var v57 int16 = (v2 + -32768)
                                                if true {
                                                    if true {
                                                        if false {
                                                            var v58 bool = false
                                                            v57 += v57
                                                            var v59 float32 = 3.14159
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                fmt.Println(v11, v45)
                                                            default:
                                                                var v60 bool = (!false)
                                                                _ = v60
                                                            }
                                                            if false {
                                                                if (!(!(!(!false)))) {
                                                                    if true {
                                                                        var v61 uint8 = 19
                                                                        _ = v61
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 2
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v3 = 42
                                                                        case "\n\t":
                                                                            if true {
                                                                                v57 = 42
                                                                            }
                                                                        }
                                                                        switch 0 {
                                                                        case 4294967295:
                                                                            var v62 int16 = v5
                                                                            _ = v62
                                                                        case 127:
                                                                            var v63 uint64 = 1
                                                                            _ = v63
                                                                        case -9223372036854775808:
                                                                            v64 := GenS3{F0: "a"}
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    _ = len(v1)
                                                                } else {
                                                                    v65 := GenS4{F0: (-(v5))}
                                                                    v54 -= (float64(((0.0 - (-2.25 - 0.1)) + (1e10 * 1e10))) + float64(-2.25))
                                                                    var v66 int64 = (v3 ^ (v3 * (v3 & v3)))
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                            } else {
                                                                v67 := GenS6{F0: 231}
                                                                var v68 uint64 = 397
                                                                v69 := v1[-1]
                                                                v2 += -1
                                                                _ = v67
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                        }
                                                        var v70 uint16 = 255
                                                        _ = v70
                                                    }
                                                    v54 = v54
                                                } else {
                                                    switch fi13 {
                                                    case -364:
                                                        var v71 int32 = 966
                                                        _ = v71
                                                    case 652:
                                                        var v72 int32 = 2147483647
                                                        _ = v72
                                                    }
                                                    fmt.Println(fi13)
                                                    var v73 int8 = (v20 % (((v20 << ((593) & 63))) | 1))
                                                    _ = v73
                                                }
                                                _ = v57
                                            }
                                            v74 := GenS7{F0: v45}
                                            _ = v74
                                        }
                                        fmt.Println(fi53, fi13, v45)
                                        _ = fi53
                                        _ = v54
                                    }
                                    if (0 >= 2) {
                                        var v75 int16 = 500
                                        fmt.Println(v20, v5)
                                        _ = v75
                                    }
                                }
                            }
                            if (-2.25 != v11) {
                                v1[v45] = (-2147483648 > fi13)
                                if (false || (!(!true))) {
                                    delete(v1, (fi13 | (^(v45))))
                                    for fi76 := 0; fi76 < 9; fi76++ {
                                        fmt.Println(fi13, v11)
                                        var v77 bool = true
                                        _ = fi76
                                        _ = v77
                                    }
                                }
                            } else {
                                var v78 uint64 = 4294967295
                                for fi79 := 0; fi79 < 8; fi79++ {
                                    fmt.Println(v11, fi79)
                                    var v80 uint = 370
                                    _ = fi79
                                    _ = v80
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v78
                            }
                            for fi81 := 0; fi81 < 12; fi81++ {
                                v82 := GenS8{F0: 2, F1: 100, F2: "\n\t"}
                                if (503 != 1) {
                                    var v83 uint16 = 0
                                    var v84 uint64 = 125
                                    _ = len(v1)
                                    if (!true) {
                                        v20 = (v20 | (-(v20)))
                                        if true {
                                            v5 -= -629
                                            var v85 int16 = v2
                                            _ = len(v1)
                                            _ = v85
                                        }
                                        if (v20 < 127) {
                                            if len(v4) > 0 {
                                                v4[0] = 961
                                            }
                                        }
                                    } else {
                                        var v86 float32 = (1e10 * 3.14159)
                                        _ = v86
                                    }
                                    _ = v83
                                    _ = v84
                                } else {
                                    if true {
                                        var v87 uint8 = 241
                                        _ = v87
                                    } else {
                                        _ = len(v1)
                                        var v88 int16 = v2
                                        var v89 float64 = v11
                                        _ = v88
                                        _ = v89
                                    }
                                }
                                v90 := GenS2{F0: 42}
                                var v91 float32 = (3.14159 + 0.0)
                                _ = fi81
                                _ = v82
                                _ = v90
                                _ = v91
                            }
                        }
                        for fi92 := 0; fi92 < 15; fi92++ {
                            v93 := GenS5{F0: 0, F1: 341, F2: "\n\t"}
                            var v94 string = ("hello" + "\n\t")
                            var v95 uint8 = 102
                            _ = fi92
                            _ = v93
                            _ = v94
                            _ = v95
                        }
                        _ = v45
                    } else {
                        v96 := GenS1{F0: 9223372036854775807, F1: 313}
                        v3 -= 9223372036854775807
                        _ = v96
                    }
                }
                _ = v20
            }
            v97 := GenS9{F0: 18446744073709551615, F1: 1}
            _ = fi13
            _ = v97
        }
    case 134:
        var v98 uint16 = 948
        _ = v98
    }
    v99 := GenS10{F0: 0, F1: 0}
    v100 := GenS11{F0: 69, F1: 252, F2: "foo bar"}
    var v101 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v100
    _ = v101
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
