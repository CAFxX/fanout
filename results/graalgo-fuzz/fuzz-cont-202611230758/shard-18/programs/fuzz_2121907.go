package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
type GenS4 struct { F0 int8; F1 uint16; F2 bool }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int64; F1 int16; F2 int16 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint32, p1 uint8, p2 bool) int8 {
    if (32768 >= 9223372036854775807) {
        v7 := []int{}
        _ = v7
    }
    if true {
        fmt.Println(p2, p2, p0)
    } else {
        var v8 int8 = 0
        if ((350 < int64(p0)) || p2) {
            v9 := GenS1{F0: -564, F1: 1, F2: p0}
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := []int{32768, 1, fi10}
                p2 = true
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        var v12 uint32 = 2
        _ = v12
        _ = v8
    }
    var v13 uint = 65535
    if (0 > 0) {
        fmt.Println(p1, p2, p2)
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int = fi14
            _ = fi14
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 9223372036854775807, F1: 414, F2: (p0 | (p0 | 1))}
        var v17 int64 = 0
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 string = "\n\t"
            if p2 {
                for fi20 := 0; fi20 < 19; fi20++ {
                    switch v19 {
                    case "\n\t":
                        p2 = true
                    case "hello":
                        v13 += (v13 ^ v13)
                    }
                    _ = fi20
                }
                p0 += p0
                v17 -= 0
            }
            _ = fi18
            _ = v19
        }
        _ = v16
        _ = v17
    }
    v21 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v21
    return 65
}
func genFn1(p0 uint8) uint32 {
    fmt.Println(p0)
    v22 := GenS2{F0: 360, F1: 65535}
    if (1 >= uint64(p0)) {
        var v23 bool = ((true && false) || true)
        v24 := GenS3{F0: -1, F1: 2, F2: 0}
        _ = v23
        _ = v24
    }
    switch 469 {
    case 599:
        var v25 uint16 = 19
        _ = v25
    case 9223372036854775807:
        if (4294967295 <= 2) {
            fmt.Println(p0)
        }
    default:
        p0 = 153
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        if (!true) {
            p0 += 1
            var v27 float64 = -2.25
            for fi28 := 0; fi28 < 11; fi28++ {
                p0 = p0
                _ = fi28
            }
            _ = v27
        } else {
            if true {
                v29 := []int{fi26, (^(fi26))}
                v30 := GenS2{F0: 1, F1: 457}
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 13; fi31++ {
                    var v32 uint8 = p0
                    _ = fi31
                    _ = v32
                }
                var v33 uint = 255
                _ = v33
            }
        }
        var v34 string = "a"
        v35 := []int{fi26, fi26}
        _ = fi26
        _ = v34
        _ = v35
    }
    _ = p0
    _ = v22
    return 1
}

func main() {
    v1 := []string{("hello" + "a"), "a", "\n\t", "foo bar"}
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := []uint32{65535, 542}
    var v5 uint64 = 2
    v6 := []int{}
    if true {
        var v36 uint = 1
        v5 -= (v5 >> (((v36 >> (((v36 ^ 927)) & 63))) & 63))
        _ = v36
    }
    fmt.Println(v5)
    _ = len(v2)
    var v37 int8 = 1
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 uint64 = (v5 & 942)
        var v40 int8 = v37
        switch (fi38 & (fi38 - fi38)) {
        case 2147483647:
            for fi41 := 0; fi41 < 8; fi41++ {
                switch 1 {
                case 28:
                    v40 = v37
                default:
                    if true {
                        v5 = v5
                        if len(v6) > 0 {
                            v6[0] = int(v37)
                        }
                        var v42 int64 = 1
                        _ = v42
                    }
                }
                _ = fi41
            }
        case 9223372036854775807:
            for fi43 := 0; fi43 < 6; fi43++ {
                switch ("" + "") {
                case "a":
                    v44 := GenS2{F0: -9223372036854775808, F1: 65535}
                    _ = v44
                }
                if false {
                    fmt.Println(v40)
                    v37 -= v40
                    v45 := GenS3{F0: 42, F1: 745, F2: 898}
                    _ = v45
                } else {
                    v39 = (v5 << ((2) & 63))
                    if false {
                        switch fi43 {
                        case 127:
                            if (true && true) {
                                var v46 uint32 = 2147483648
                                _ = v46
                            } else {
                                fmt.Println(v5, v39, fi43)
                            }
                        case 128:
                            v47 := GenS2{F0: 162, F1: genFn1(37)}
                            _ = v47
                        }
                        if false {
                            v48 := GenS4{F0: (v40 % ((0) | 1)), F1: 1, F2: true}
                            var v49 string = "a"
                            _ = v48
                            _ = v49
                        } else {
                            var v50 uint16 = 975
                            switch -2147483648 {
                            case 1:
                                var v51 uint64 = 4294967295
                                _ = v51
                            case -552:
                                _ = len(v2)
                            case 4294967295:
                                v5 = v39
                            default:
                                if true {
                                    fmt.Println(v50)
                                } else {
                                    v52 := GenS5{F0: (fi38 / ((fi43) | 1)), F1: 2}
                                    var v53 uint64 = 2
                                    if true {
                                        var v54 int32 = 2147483647
                                        if false {
                                            var v55 int = 32767
                                            switch 774 {
                                            case 4294967295:
                                                if true {
                                                    v56 := v3[(fi38 >> ((343) & 63))]
                                                    _ = v56
                                                }
                                            case 32767:
                                                if false {
                                                    if false {
                                                        v57 := GenS3{F0: int64(v37), F1: v50, F2: 10}
                                                        var v58 uint32 = 1
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                } else {
                                                    var v59 uint8 = 227
                                                    _ = v59
                                                }
                                            }
                                            _ = v55
                                        } else {
                                            switch (fi38 ^ fi38) {
                                            case 42:
                                                v5 = 804
                                            }
                                            if true {
                                                var v60 int64 = 365
                                                _ = v60
                                            } else {
                                                v50 -= v50
                                                v54 = -2147483648
                                                switch (fi43 ^ 32768) {
                                                case -9223372036854775808:
                                                    v4 = append(v4, 1)
                                                }
                                            }
                                            var v61 int64 = 1
                                            _ = v61
                                        }
                                        if (1 < 2) {
                                            fmt.Println(fi43, v5, v37)
                                            fmt.Println(fi38)
                                        }
                                        _ = v54
                                    }
                                    _ = v52
                                    _ = v53
                                }
                            }
                            var v62 uint64 = 0
                            _ = v50
                            _ = v62
                        }
                        switch fi38 {
                        case 42:
                            v40 = (^(v37))
                        case 4294967295:
                            var v63 int64 = -1
                            _ = v63
                        case -247:
                            v1 = append(v1, "hello")
                        default:
                            var v64 float32 = 1.5
                            _ = v64
                        }
                        v65 := GenS6{F0: 42, F1: 1000, F2: 1}
                        switch fi38 {
                        case -888:
                            v37 += genFn0(1, 1, (65535 > 2))
                        case -177:
                            fmt.Println(v37, v40, fi43)
                        case 42:
                            fmt.Println(v37)
                        default:
                            v40 -= v40
                        }
                        _ = v65
                    }
                    v66 := GenS5{F0: (fi38 ^ -366), F1: 1}
                    _ = v66
                }
                v67 := GenS5{F0: (^(fi38)), F1: 42}
                _ = fi43
                _ = v67
            }
        case 32767:
            v39 -= (^(v39))
        default:
            v37 += v37
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    if len(v6) > 0 {
        v6[0] = -376
    }
    v68 := GenS7{F0: v37, F1: 1}
    if true {
        if (!(true || ((-463 <= 1) && (4294967295 <= 627)))) {
            v37 -= (^(v37))
            fmt.Println(v37)
        }
        v69 := GenS8{F0: 118}
        _ = v69
    } else {
        fmt.Println(v5, v37)
    }
    v70 := GenS7{F0: (v37 << ((65535) & 63)), F1: -263}
    var v71 uint8 = 255
    v37 = (v37 + v37)
    if (-9223372036854775808 >= -480) {
        for fi72 := 0; fi72 < 18; fi72++ {
            fmt.Println(fi72)
            _ = fi72
        }
    }
    delete(v3, 4294967295)
    v71 -= v71
    var v73 bool = (!(42 < 1))
    for fi74 := 0; fi74 < 10; fi74++ {
        switch "hello" {
        case "\n\t":
            v73 = (true && ((fi74 >> ((283) & 63)) >= (fi74 >> ((715) & 63))))
        default:
            for fi75 := 0; fi75 < 4; fi75++ {
                fmt.Println(v5, v5)
                if v73 {
                    var v76 uint = 65535
                    if v73 {
                        v77 := GenS2{F0: -9223372036854775808, F1: genFn1(128)}
                        if ("a" > "") {
                            v73 = (!v73)
                            switch "hello" {
                            case "\n\t":
                                if (557 < 998) {
                                    v37 -= v37
                                    var v78 uint32 = 420
                                    _ = v78
                                } else {
                                    fmt.Println(v71)
                                }
                            case "foo bar":
                                switch "" {
                                case "hello":
                                    v79 := GenS6{F0: -226, F1: 1, F2: -281}
                                    _ = v79
                                case "a":
                                    var v80 uint8 = (v71 + v71)
                                    v37 += v37
                                    _ = v80
                                case "":
                                    if v73 {
                                        v81 := GenS2{F0: -281, F1: 1}
                                        var v82 int8 = (v37 - (v37 + -75))
                                        var v83 uint64 = (^(v5))
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                    }
                                default:
                                    v4 = append(v4, 2147483647)
                                }
                            case "hello":
                                v76 = 230
                            }
                        } else {
                            v84 := GenS8{F0: v71}
                            fmt.Println(v5, v37, fi74)
                            _ = v84
                        }
                        if v73 {
                            var v85 int8 = genFn0(2147483647, (v71 % ((v71) | 1)), v73)
                            v86 := GenS3{F0: -1, F1: 2, F2: -1}
                            _ = v85
                            _ = v86
                        } else {
                            var v87 float32 = -2.25
                            if false {
                                v88 := GenS3{F0: 0, F1: 780, F2: 1}
                                var v89 uint8 = v71
                                _ = v88
                                _ = v89
                            }
                            if ((1 > -32768) || (42 > 4294967295)) {
                                v73 = (!(!(((v71 - uint8(v76)) >= v71) && v73)))
                            }
                            _ = v87
                        }
                        _ = v77
                    } else {
                        v90 := GenS6{F0: 1, F1: -434, F2: int16(v71)}
                        _ = v90
                    }
                    v73 = ((1e10 / 1e10) <= 1e10)
                    _ = v76
                } else {
                    if len(v6) > 0 {
                        v6[0] = (fi74 * fi74)
                    }
                }
                _ = fi75
            }
        }
        var v91 float32 = float32(0.0)
        _ = fi74
        _ = v91
    }
    if (42 <= 1) {
        for fi92 := 0; fi92 < 1; fi92++ {
            var v93 uint = 18446744073709551615
            var v94 int32 = -2147483648
            var v95 int32 = (v94 / ((v94) | 1))
            _ = fi92
            _ = v93
            _ = v94
            _ = v95
        }
    } else {
        if v73 {
            var v96 float64 = 1.5
            v73 = ((!true) && (v71 >= (v71 / (((v71 << ((uint(v5)) & 63))) | 1))))
            var v97 int = 0
            _ = v96
            _ = v97
        }
        v98 := GenS1{F0: 9223372036854775807, F1: 2, F2: 1}
        v4 = append(v4, 1)
        _ = v98
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v70
    _ = v71
    _ = v73
}
