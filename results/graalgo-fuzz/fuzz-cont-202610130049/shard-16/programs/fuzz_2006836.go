package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 bool; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint; F2 uint8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 uint }
type GenS13 struct { F0 int32; F1 uint16; F2 string }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 string = "\n\t"
    var v5 uint64 = 18446744073709551615
    switch v4 {
    case "":
        for fi6 := 0; fi6 < 15; fi6++ {
            delete(v3, fi6)
            _ = fi6
        }
    case "foo bar":
        v5 -= v5
    case "hello":
        var v7 int32 = 525
        _ = v7
    }
    v4 = (v4 + ("a" + v4))
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch fi8 {
            case 4294967295:
                v10 := GenS1{F0: "a", F1: 0}
                _ = v10
            case -816:
                v11 := GenS2{F0: 255, F1: 32767, F2: 0}
                _ = v11
            case -677:
                var v12 int32 = int32(v5)
                _ = v12
            default:
                v5 = 45
            }
            v5 = 4294967295
            v13 := []int{127, fi9, -627}
            _ = fi9
            _ = v13
        }
        if (-2.25 != float64(float32(-2.25))) {
            v14 := v2[(v4 + v4)]
            v4 = v4
            v15 := []int{-9223372036854775808, fi8, (fi8 | (fi8 | 9223372036854775807)), fi8}
            _ = v14
            _ = v15
        } else {
            fmt.Println(v5, v4, fi8)
            if (0 >= 0) {
                v16 := GenS2{F0: 1, F1: 690, F2: -2147483648}
                v17 := GenS1{F0: v4, F1: -1}
                v4 = ""
                _ = v16
                _ = v17
            }
        }
        _ = fi8
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(fi18, v4)
        _ = fi18
    }
    v5 = v5
    fmt.Println(v1)
    v19 := GenS3{F0: 0, F1: -32768, F2: ((v4 + v4) + (v4 + (v4 + v4)))}
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := []int{2147483647, fi20, (fi20 | (fi20 >> ((4294967295) & 63))), 777}
        v4 = "hello"
        v22 := GenS2{F0: 127, F1: 2, F2: 0}
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint32 = 788
    v24 := v3[9223372036854775807]
    var v25 uint16 = 255
    v26 := GenS2{F0: 104, F1: 2, F2: -1}
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(v23)
        for fi28 := 0; fi28 < 9; fi28++ {
            if (int64(v5) <= 9223372036854775807) {
                fmt.Println(v24, v23)
                var v29 int8 = 100
                if true {
                    if false {
                        v30 := GenS2{F0: 115, F1: 1, F2: 2147483647}
                        var v31 string = "hello"
                        switch (fi28 << ((816) & 63)) {
                        case 2147483647:
                            if (v24 > (v23 << ((42) & 63))) {
                                v4 = "a"
                                if (!false) {
                                    fmt.Println(v31, v23, v25)
                                    var v32 uint = 1
                                    v33 := GenS2{F0: 2, F1: 42, F2: -1}
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 int32 = -2147483648
                                    v4 = v4
                                    _ = v34
                                }
                                switch "foo bar" {
                                case "":
                                    v29 = (v29 - (v29 >> ((uint(v25)) & 63)))
                                default:
                                    var v35 float64 = 0.0
                                    _ = v35
                                }
                                var v36 uint8 = 128
                                if false {
                                    v25 = 1
                                } else {
                                    if ((^(v29)) >= v29) {
                                        v37 := []int{4294967295, (fi27 - 0), (^(fi28))}
                                        _ = v37
                                    } else {
                                        var v38 float64 = 1e10
                                        var v39 uint32 = v24
                                        v25 = (v25 + 65535)
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                _ = v36
                            } else {
                                if true {
                                    if true {
                                        v40 := []int{(-(fi27)), 495, fi27, 32767}
                                        _ = v40
                                    } else {
                                        fmt.Println(v25, v24)
                                        if false {
                                            v5 = (v5 ^ 42)
                                            v5 += v5
                                            v41 := GenS4{F0: 0}
                                            _ = v41
                                        } else {
                                            switch (((("\n\t" + "") + v31) + v4) + v4) {
                                            case "":
                                                v31 = (v4 + ("foo bar" + (v31 + ("\n\t" + ("" + v4)))))
                                            case "\n\t":
                                                _ = len(v2)
                                            case "foo bar":
                                                v23 = v23
                                                var v42 int16 = 600
                                                var v43 int = (fi28 % ((127) | 1))
                                                var v44 float32 = v1
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            default:
                                                switch (v31 + (((v4 + v4) + v31) + v31)) {
                                                case "hello":
                                                    v2["\n\t"] = (1 != 2)
                                                case "foo bar":
                                                    if (float32((float64(v1) / (((1e10 - (-2.25 / 1e10)) * -2.25) + (-2.25 - 0.1)))) == 0.0) {
                                                        if true {
                                                            var v45 uint64 = (v5 & v5)
                                                            fmt.Println(v31, v29)
                                                            var v46 uint8 = 2
                                                            _ = v45
                                                            _ = v46
                                                        } else {
                                                            fmt.Println(v23)
                                                            var v47 int8 = (v29 / ((-128) | 1))
                                                            _ = v47
                                                        }
                                                        var v48 string = (v4 + v31)
                                                        v49 := []int{(fi27 >> ((586) & 63)), fi27}
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                }
                                            }
                                            v24 = v24
                                        }
                                    }
                                    var v50 float64 = 1.5
                                    var v51 int32 = -610
                                    _ = v50
                                    _ = v51
                                } else {
                                    if (v5 >= v5) {
                                        if false {
                                            if (true || false) {
                                                var v52 uint = 0
                                                _ = v52
                                            } else {
                                                v3[(fi28 - fi28)] = 32
                                            }
                                            v53 := GenS4{F0: 366}
                                            v54 := v3[fi27]
                                            _ = v53
                                            _ = v54
                                        }
                                        if (!(true && true)) {
                                            v55 := GenS2{F0: 54, F1: v25, F2: -717}
                                            v56 := GenS1{F0: v4, F1: -1}
                                            _ = v55
                                            _ = v56
                                        }
                                    } else {
                                        switch (^(fi28)) {
                                        case 913:
                                            v2[v4] = true
                                        case 32768:
                                            var v57 uint16 = (v25 + v25)
                                            _ = v57
                                        case 32767:
                                            fmt.Println(fi28, v29, fi28)
                                        default:
                                            v58 := GenS5{F0: 0}
                                            _ = v58
                                        }
                                    }
                                    switch (-(fi28)) {
                                    case 128:
                                        v24 = 0
                                    }
                                    v25 = v25
                                    if false {
                                        fmt.Println(v29)
                                        var v59 bool = false
                                        _ = v59
                                    }
                                    v60 := v2["a"]
                                    if false {
                                        var v61 int16 = 42
                                        _ = v61
                                    } else {
                                        v62 := []int{(fi27 << ((4294967295) & 63))}
                                        var v63 uint8 = 2
                                        if (!false) {
                                            v60 = ((v24 > 2147483647) && false)
                                        } else {
                                            v64 := GenS6{F0: 128, F1: 0}
                                            var v65 uint8 = (v63 % ((v63) | 1))
                                            _ = v64
                                            _ = v65
                                        }
                                        _ = v62
                                        _ = v63
                                    }
                                    v66 := GenS7{F0: (v5 ^ (v5 & (^(v5))))}
                                    _ = v60
                                    _ = v66
                                }
                                var v67 uint64 = v5
                                _ = v67
                            }
                        case 32767:
                            var v68 uint = 500
                            _ = v68
                        default:
                            var v69 float32 = (v1 - -2.25)
                            fmt.Println(fi27, v25, fi27)
                            _ = v69
                        }
                        _ = v30
                        _ = v31
                    }
                    switch fi27 {
                    case 2147483647:
                        v5 += v5
                    case 453:
                        if (42 == -253) {
                            v5 += v5
                            v70 := GenS7{F0: v5}
                            _ = v70
                        }
                    case -358:
                        v29 = v29
                    }
                    if true {
                        v1 = (float32(float64(((1e10 / v1) / 0.1))) + 1.5)
                        v71 := GenS8{F0: 2147483647}
                        _ = v71
                    } else {
                        var v72 bool = true
                        v73 := GenS9{F0: "hello", F1: ((76 > 4294967295) && true), F2: v5}
                        _ = v72
                        _ = v73
                    }
                    v74 := []int{fi28, fi27}
                    if false {
                        var v75 string = ((((v4 + "hello") + v4) + "") + v4)
                        switch "foo bar" {
                        case "a":
                            var v76 int16 = 1
                            _ = v76
                        case "":
                            var v77 uint8 = 42
                            _ = v77
                        case "\n\t":
                            var v78 bool = true
                            fmt.Println(v5, fi27, v23)
                            switch (fi28 ^ 595) {
                            case 1:
                                var v79 int8 = (v29 % (((^(v29))) | 1))
                                _ = v79
                            case -630:
                                fmt.Println(v25)
                            case 127:
                                v74 = append(v74, 32768)
                                v23 = (v24 << ((1) & 63))
                                v80 := GenS6{F0: -688, F1: 64}
                                _ = v80
                            }
                            _ = v78
                        }
                        v4 = v75
                        _ = v75
                    } else {
                        if false {
                            v81 := GenS3{F0: uint8(v23), F1: 32767, F2: v4}
                            v24 += v23
                            _ = v81
                        } else {
                            v24 = (^(v23))
                        }
                    }
                    _ = v74
                } else {
                    v82 := GenS10{F0: 9223372036854775807, F1: 0, F2: 176}
                    _ = v82
                }
                _ = v29
            } else {
                var v83 float32 = v1
                var v84 uint = uint(v25)
                switch v4 {
                case "\n\t":
                    switch (v4 + ("" + ("\n\t" + (v4 + ("a" + v4))))) {
                    case "hello":
                        switch v4 {
                        case "foo bar":
                            if false {
                                if (!true) {
                                    switch -2147483648 {
                                    case -268:
                                        var v85 uint32 = (v23 * v23)
                                        _ = v85
                                    case 32767:
                                        _ = len(v3)
                                    case 0:
                                        var v86 bool = true
                                        _ = v86
                                    }
                                    fmt.Println(fi28, v23, v24)
                                } else {
                                    if true {
                                        v87 := GenS11{F0: -2147483648}
                                        _ = v87
                                    } else {
                                        var v88 int16 = 32767
                                        if false {
                                            var v89 bool = ((v88 % (((v88 * v88)) | 1)) > (v88 ^ 1000))
                                            _ = v89
                                        }
                                        _ = v88
                                    }
                                    v83 -= v83
                                }
                                v90 := GenS5{F0: 1}
                                _ = v90
                            }
                            switch v4 {
                            case "hello":
                                var v91 uint64 = 0
                                _ = v91
                            case "a":
                                v3[fi28] = 2147483647
                            case "\n\t":
                                v92 := GenS12{F0: 42, F1: 42}
                                _ = v92
                            default:
                                v5 = (v5 << (((^(v84))) & 63))
                            }
                        }
                    }
                case "a":
                    v83 += (v83 * v83)
                }
                _ = v83
                _ = v84
            }
            v93 := GenS13{F0: 0, F1: 42, F2: v4}
            v94 := []int{}
            _ = fi28
            _ = v93
            _ = v94
        }
        v95 := []int{}
        _ = fi27
        _ = v95
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
