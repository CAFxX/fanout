package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS10 struct { F0 bool; F1 int64; F2 uint16 }
type GenS11 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 bool = true
    var v2 float64 = 1e10
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    delete(v3, 230)
    var v5 int64 = 9223372036854775807
    switch -2147483648 {
    case 4294967295:
        var v6 uint64 = 915
        _ = v6
    case 127:
        if false {
            v7 := GenS1{F0: (42 < 255), F1: 168, F2: 616}
            _ = v7
        }
    case 632:
        v4 = v4
    }
    var v8 int8 = -1
    v9 := GenS2{F0: v8, F1: 348}
    var v10 uint32 = 2147483647
    v11 := GenS3{F0: (v4 + v4), F1: uint16(v10), F2: 4294967295}
    var v12 int8 = 0
    var v13 bool = (!true)
    if (v4 == v4) {
        if v1 {
            if v1 {
                v14 := []int{}
                if true {
                    v15 := GenS1{F0: (v13 || false), F1: 18446744073709551615, F2: (v10 & v10)}
                    var v16 uint = 2
                    _ = v15
                    _ = v16
                } else {
                    var v17 int64 = (v5 | (v5 % ((v5) | 1)))
                    for fi18 := 0; fi18 < 7; fi18++ {
                        fmt.Println(v17)
                        _ = fi18
                    }
                    _ = v17
                }
                v19 := GenS4{F0: 196, F1: 127}
                _ = v14
                _ = v19
            } else {
                var v20 int = 797
                _ = v20
            }
            var v21 int32 = 1
            _ = v21
        } else {
            if v1 {
                switch -800 {
                case 345:
                    v2 += 3.14159
                case 127:
                    fmt.Println(v10)
                default:
                    switch "foo bar" {
                    case "foo bar":
                        var v22 int = 128
                        _ = v22
                    case "a":
                        v13 = (!true)
                    default:
                        var v23 uint64 = 42
                        _ = v23
                    }
                }
                if true {
                    v24 := GenS5{F0: -740, F1: (^(v5))}
                    _ = v24
                }
                var v25 bool = false
                _ = v25
            } else {
                for fi26 := 0; fi26 < 16; fi26++ {
                    v8 = 127
                    var v27 bool = (true || true)
                    v1 = (4294967295 <= (v10 * 65535))
                    _ = fi26
                    _ = v27
                }
            }
        }
        v28 := []int{128, 291}
        v2 += v2
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        switch v4 {
        case "\n\t":
            v30 := []int{}
            _ = v30
        case "a":
            for fi31 := 0; fi31 < 19; fi31++ {
                switch 4294967295 {
                case 0:
                    if (127 >= 88) {
                        switch fi29 {
                        case 417:
                            v32 := GenS6{F0: 2, F1: 0}
                            _ = v32
                        case -2147483648:
                            v2 = float64((0.0 * -2.25))
                        case 32767:
                            switch ((v4 + "hello") + (("a" + v4) + "a")) {
                            case "hello":
                                var v33 uint32 = 42
                                _ = v33
                            case "":
                                v34 := []int{fi31, (fi31 / ((-753) | 1)), fi31, fi29}
                                _ = v34
                            }
                        }
                        switch fi29 {
                        case 128:
                            var v35 uint8 = 127
                            _ = v35
                        case -621:
                            var v36 int = (^(fi31))
                            _ = v36
                        }
                    } else {
                        if true {
                            v37 := GenS7{F0: (v13 && false)}
                            v13 = v13
                            if false {
                                switch v4 {
                                case "hello":
                                    var v38 uint16 = 0
                                    _ = v38
                                case "a":
                                    v39 := GenS8{F0: 0}
                                    _ = v39
                                }
                            }
                            v40 := []int{}
                            _ = v37
                            _ = v40
                        }
                        fmt.Println(v12, v2)
                        v5 += v5
                        switch v4 {
                        case "a":
                            if (1.5 != -2.25) {
                                v5 -= (v5 & v5)
                                if (v5 < (v5 & -174)) {
                                    switch v4 {
                                    case "":
                                        switch v4 {
                                        case "a":
                                            v8 -= v8
                                        }
                                        var v41 uint64 = 18446744073709551615
                                        _ = v41
                                    case "hello":
                                        v8 += (v12 + 127)
                                        fmt.Println(v2, v10, v5)
                                        v42 := GenS9{F0: (v1 && v13), F1: uint16(v10), F2: 32767}
                                        _ = v42
                                    default:
                                        fmt.Println(v4, v8)
                                    }
                                    fmt.Println(fi29)
                                    switch v4 {
                                    case "foo bar":
                                        v3[v10] = (("foo bar" + v4) >= v4)
                                    case "a":
                                        v43 := GenS6{F0: 32767, F1: 358}
                                        _ = v43
                                    }
                                } else {
                                    v13 = (false && (v1 && false))
                                    _ = len(v3)
                                }
                            } else {
                                v1 = false
                                v12 -= (v12 | (v8 ^ (v8 % ((v8) | 1))))
                                var v44 float32 = float32((v2 / (v2 - v2)))
                                var v45 uint32 = 208
                                v10 -= (v45 | (^(v10)))
                                var v46 int8 = v12
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        case "":
                            v47 := []int{-519, fi31, fi31, fi31}
                            _ = v47
                        }
                        v12 += v8
                    }
                case 42:
                    v13 = v1
                case 1:
                    if true {
                        var v48 int = fi29
                        v10 += v10
                        _ = v48
                    }
                }
                if false {
                    v4 = "hello"
                } else {
                    v13 = true
                    if (!((uint(v10) < 618) && v1)) {
                        v49 := GenS5{F0: (fi31 / ((767) | 1)), F1: (v5 % (((-(v5))) | 1))}
                        if v13 {
                            var v50 uint = 2
                            var v51 int64 = (v5 - v5)
                            _ = v50
                            _ = v51
                        } else {
                            v52 := GenS2{F0: 100, F1: v10}
                            v53 := GenS10{F0: (!((!((v10 - 2147483648) > v10)) && v13)), F1: v5, F2: 202}
                            _ = v52
                            _ = v53
                        }
                        v5 = (-(v5))
                        _ = v49
                    } else {
                        var v54 uint32 = (v10 / ((v10) | 1))
                        v12 = (-(v12))
                        var v55 bool = (((-2.25 / 1.5) / v2) > (v2 / v2))
                        _ = v54
                        _ = v55
                    }
                }
                if (411 > 1) {
                    var v56 int = (^(fi29))
                    var v57 uint16 = 255
                    v58 := []int{(fi31 / ((fi29) | 1)), (^(fi29)), 128, (v56 >> ((4) & 63))}
                    _ = v56
                    _ = v57
                    _ = v58
                }
                _ = fi31
            }
        case "":
            var v59 uint8 = 17
            _ = v59
        default:
            for fi60 := 0; fi60 < 8; fi60++ {
                var v61 float64 = 0.1
                fmt.Println(v10, v12)
                v1 = (v13 || false)
                _ = fi60
                _ = v61
            }
        }
        _ = fi29
    }
    var v62 int8 = (v12 << ((4294967295) & 63))
    v63 := []int{}
    v8 = 100
    var v64 string = ((v4 + v4) + v4)
    for fi65 := 0; fi65 < 18; fi65++ {
        var v66 bool = ((v62 - 0) != -65)
        if (true || v66) {
            v4 = v64
            v67 := GenS6{F0: 270, F1: 4294967295}
            if (42 >= 2) {
                fmt.Println(v64, v8)
                if (-2147483648 != 2147483647) {
                    _ = len(v63)
                    for fi68 := 0; fi68 < 13; fi68++ {
                        var v69 uint32 = v10
                        if true {
                            v10 = v69
                            v13 = false
                            delete(v3, 486)
                        }
                        if len(v63) > 0 {
                            v63[0] = -187
                        }
                        _ = fi68
                        _ = v69
                    }
                }
            }
            _ = v67
        } else {
            v4 = v4
            for fi70 := 0; fi70 < 7; fi70++ {
                _ = len(v3)
                switch (v4 + ((v64 + "a") + v64)) {
                case "hello":
                    if (v4 <= v64) {
                        v1 = v13
                        if ((((v64 + v4) + ("" + v4)) + v64) >= "foo bar") {
                            if ((0.1 > 1.5) && v66) {
                                fmt.Println(v1, fi70, v1)
                                var v71 uint32 = (v10 - (v10 | v10))
                                var v72 uint16 = 825
                                switch 42 {
                                case 42:
                                    _ = len(v63)
                                case 430:
                                    if v13 {
                                        fmt.Println(v66)
                                        v8 -= (-(v8))
                                        var v73 int64 = 9223372036854775807
                                        _ = v73
                                    } else {
                                        v2 = (1e10 / 1.5)
                                        v1 = (-1 == 2147483647)
                                    }
                                }
                                var v74 uint = uint(v72)
                                if ("" >= v64) {
                                    v12 -= v8
                                }
                                switch "" {
                                case "\n\t":
                                    v75 := GenS9{F0: (!v13), F1: v72, F2: (v72 + (v72 + (v72 & 2)))}
                                    _ = v75
                                default:
                                    var v76 int = fi65
                                    _ = v76
                                }
                                _ = v71
                                _ = v72
                                _ = v74
                            }
                            switch fi65 {
                            case 114:
                                if (!(!(2 >= 0))) {
                                    var v77 int8 = 0
                                    if (42 <= 235) {
                                        var v78 uint32 = 42
                                        if ((3.14159 - 1.5) >= float32(v2)) {
                                            var v79 int16 = -373
                                            var v80 uint16 = 32767
                                            _ = v79
                                            _ = v80
                                        }
                                        _ = len(v3)
                                        _ = v78
                                    } else {
                                        if ((true && v66) || v1) {
                                            v66 = (!false)
                                            _ = len(v63)
                                            v62 = v77
                                        }
                                        switch (((((v4 + v4) + v4) + v4) + ("hello" + "\n\t")) + v64) {
                                        case "":
                                            if (v2 <= 3.14159) {
                                                var v81 uint = 199
                                                _ = v81
                                            } else {
                                                var v82 int64 = v5
                                                v64 = (((v4 + ("" + v4)) + "a") + "\n\t")
                                                _ = v82
                                            }
                                        case "\n\t":
                                            var v83 uint32 = 0
                                            _ = v83
                                        case "hello":
                                            v84 := GenS11{F0: 4294967295, F1: 18446744073709551615}
                                            _ = v84
                                        }
                                    }
                                    delete(v3, (v10 / (((^(v10))) | 1)))
                                    _ = v77
                                }
                            default:
                                switch "hello" {
                                case "\n\t":
                                    var v85 float32 = 1.5
                                    _ = v85
                                }
                            }
                            if v1 {
                                fmt.Println(v2, v5)
                                var v86 bool = v13
                                delete(v3, 1)
                                var v87 int64 = 9223372036854775807
                                v88 := v3[v10]
                                _ = v86
                                _ = v87
                                _ = v88
                            }
                        } else {
                            fmt.Println(v8, v2)
                            v89 := GenS2{F0: 1, F1: v10}
                            _ = v89
                        }
                    }
                case "":
                    fmt.Println(v10)
                case "foo bar":
                    fmt.Println(v5, v66, v62)
                }
                switch -566 {
                case 32768:
                    v13 = false
                case -1:
                    v1 = ((fi70 / (((fi65 | fi70)) | 1)) < (fi70 & -336))
                case 127:
                    var v90 int = (^(fi65))
                    _ = v90
                default:
                    v2 -= float64((((0.0 + -2.25) - 0.0) - -2.25))
                    var v91 float32 = (3.14159 / 0.1)
                    _ = v91
                }
                _ = fi70
            }
        }
        _ = fi65
        _ = v66
    }
    v4 = (v4 + "")
    var v92 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    fmt.Println("v8:", v8)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v8
    _ = v9
    _ = v92
}
