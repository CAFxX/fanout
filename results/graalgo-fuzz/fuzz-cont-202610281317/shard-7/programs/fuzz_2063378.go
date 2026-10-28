package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 uint32; F2 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint64; F1 uint32; F2 string }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 bool; F1 uint32 }
type GenS14 struct { F0 bool; F1 uint; F2 int }
func genFn0(p0 int8, p1 uint, p2 uint64) uint {
    switch "a" {
    case "a":
        switch 1 {
        case 377:
            var v7 uint32 = 1
            _ = v7
        case 94:
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 bool = false
                _ = fi8
                _ = v9
            }
        case 485:
            var v10 int32 = 0
            _ = v10
        }
    case "":
        var v11 uint = p1
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int = 32767
            var v14 uint32 = 65535
            fmt.Println(p0, v13)
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    if ("\n\t" != ("foo bar" + "\n\t")) {
        var v15 int = -661
        _ = v15
    }
    var v16 uint64 = (p2 / ((0) | 1))
    v17 := []int{127, 32767, 9223372036854775807}
    var v18 uint8 = 17
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return p1
}
func genFn1(p0 uint32) bool {
    var v19 uint16 = 445
    var v20 int64 = 9223372036854775807
    v21 := GenS1{F0: "hello"}
    v22 := GenS2{F0: "a", F1: int16(v20)}
    v23 := GenS3{F0: -9223372036854775808, F1: "foo bar"}
    _ = p0
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    return ((1 <= uint64(v19)) && ("" != "foo bar"))
}

func main() {
    var v1 string = "\n\t"
    var v2 int = -2147483648
    v3 := make(map[uint32]int)
    var v4 string = v1
    var v5 float32 = -2.25
    var v6 int8 = 0
    var v24 int = v2
    v25 := []int{(v24 - (^(v24))), v24, v24, v2}
    v26 := GenS4{F0: 42, F1: -2147483648}
    if len(v25) > 0 {
        v25[0] = v24
    }
    v27 := GenS5{F0: v24}
    var v28 int16 = -32768
    for fi29 := 0; fi29 < 20; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int32 = 0
            v6 += (-(v6))
            _ = len(v3)
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 9; fi32++ {
            fmt.Println(fi29)
            fmt.Println(v1, v5)
            _ = fi32
        }
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        _ = fi29
    }
    if false {
        if (!true) {
            if true {
                for fi33 := 0; fi33 < 17; fi33++ {
                    v34 := GenS2{F0: "", F1: (^(v28))}
                    for fi35 := 0; fi35 < 9; fi35++ {
                        v36 := GenS6{F0: 238}
                        v37 := GenS6{F0: 453}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    v5 += v5
                    _ = fi33
                    _ = v34
                }
                for fi38 := 0; fi38 < 16; fi38++ {
                    v39 := GenS7{F0: (fi38 - v2)}
                    _ = fi38
                    _ = v39
                }
                for fi40 := 0; fi40 < 9; fi40++ {
                    switch (v4 + v1) {
                    case "":
                        if false {
                            if true {
                                switch v1 {
                                case "a":
                                    var v41 uint8 = 255
                                    _ = v41
                                case "foo bar":
                                    if true {
                                        v42 := GenS7{F0: fi40}
                                        var v43 int16 = v28
                                        if false {
                                            var v44 float32 = (v5 + -2.25)
                                            if true {
                                                fmt.Println(v24, v4, v24)
                                                fmt.Println(v24)
                                            }
                                            if true {
                                                for fi45 := 0; fi45 < 10; fi45++ {
                                                    v46 := v3[2147483648]
                                                    _ = fi45
                                                    _ = v46
                                                }
                                                v4 = "foo bar"
                                                if (true && (42 < 32767)) {
                                                    if true {
                                                        v28 = (v43 * int16(v6))
                                                        for fi47 := 0; fi47 < 12; fi47++ {
                                                            var v48 int8 = 89
                                                            fmt.Println(v5, v5)
                                                            if (42 < 0) {
                                                                switch (v2 ^ fi47) {
                                                                case 128:
                                                                    var v49 bool = false
                                                                    _ = v49
                                                                case -824:
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v50 uint64 = 2
                                                                        _ = v50
                                                                    case "\n\t":
                                                                        var v51 int = v24
                                                                        _ = v51
                                                                    case "hello":
                                                                        var v52 float64 = 1e10
                                                                        _ = v52
                                                                    }
                                                                case 32767:
                                                                    v53 := GenS8{F0: (fi40 | (^(fi47))), F1: 1, F2: 2}
                                                                    _ = v53
                                                                default:
                                                                    if false {
                                                                        var v54 uint = 65535
                                                                        _ = v54
                                                                    }
                                                                }
                                                                v4 = v4
                                                                switch v24 {
                                                                case 32768:
                                                                    if true {
                                                                        switch (-(v24)) {
                                                                        case -205:
                                                                            v24 += 1
                                                                        case 0:
                                                                            v1 = v1
                                                                        default:
                                                                            v43 += -462
                                                                        }
                                                                        if (1 < 42) {
                                                                            if (255 == 255) {
                                                                                var v55 int16 = (v43 >> ((2) & 63))
                                                                                v2 += (v24 << ((18446744073709551615) & 63))
                                                                                _ = v55
                                                                            }
                                                                            var v56 float32 = float32(1.5)
                                                                            var v57 uint32 = 4294967295
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                        fmt.Println(v2)
                                                                    }
                                                                    v58 := GenS2{F0: v4, F1: v28}
                                                                    switch 0 {
                                                                    case -2147483648:
                                                                        if (true && true) {
                                                                            switch (v24 & (fi40 & 283)) {
                                                                            case -2147483648:
                                                                                v59 := GenS1{F0: v1}
                                                                                _ = v59
                                                                            default:
                                                                                v6 = 127
                                                                            }
                                                                            v25 = append(v25, fi40)
                                                                            var v60 int = v2
                                                                            switch v4 {
                                                                            case "a":
                                                                                v3[4294967295] = 4294967295
                                                                            }
                                                                            _ = v60
                                                                        } else {
                                                                            var v61 int16 = (v28 ^ (v43 << ((2) & 63)))
                                                                            v62 := GenS9{F0: 178}
                                                                            var v63 int = v2
                                                                            _ = v61
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                    }
                                                                    _ = v58
                                                                default:
                                                                    fmt.Println(v5)
                                                                }
                                                            }
                                                            _ = fi47
                                                            _ = v48
                                                        }
                                                        v64 := v3[42]
                                                        _ = v64
                                                    } else {
                                                        for fi65 := 0; fi65 < 2; fi65++ {
                                                            var v66 int64 = 1
                                                            _ = fi65
                                                            _ = v66
                                                        }
                                                        fmt.Println(v4, v44)
                                                        v43 = 589
                                                    }
                                                }
                                            } else {
                                                v67 := GenS6{F0: 0}
                                                v43 -= (v43 - 42)
                                                var v68 float32 = ((0.1 / v5) / v44)
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = v44
                                        } else {
                                            fmt.Println(v2)
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                v69 := v3[591]
                                _ = v69
                            } else {
                                v70 := GenS9{F0: 2147483647}
                                v3[535] = 32767
                                _ = v70
                            }
                        } else {
                            if false {
                                var v71 int64 = -1
                                for fi72 := 0; fi72 < 11; fi72++ {
                                    v73 := GenS2{F0: ("foo bar" + ((v4 + "a") + (((v4 + "foo bar") + v1) + "\n\t"))), F1: v28}
                                    switch fi72 {
                                    case 507:
                                        var v74 int = -320
                                        _ = v74
                                    default:
                                        fmt.Println(v5)
                                    }
                                    var v75 int32 = 2147483647
                                    _ = fi72
                                    _ = v73
                                    _ = v75
                                }
                                v3[42] = v2
                                _ = v71
                            }
                            v76 := GenS10{F0: 2, F1: 42, F2: v1}
                            _ = v76
                        }
                    case "hello":
                        v77 := GenS4{F0: 128, F1: -1}
                        _ = v77
                    default:
                        var v78 int64 = 9223372036854775807
                        _ = v78
                    }
                    var v79 int = v24
                    _ = fi40
                    _ = v79
                }
            } else {
                switch (v4 + v4) {
                case "foo bar":
                    var v80 int16 = (v28 / (((v28 - (^(v28)))) | 1))
                    _ = v80
                case "a":
                    fmt.Println(v1)
                case "":
                    v1 = v1
                }
                _ = len(v25)
            }
            v81 := GenS11{F0: (-(v6))}
            v82 := GenS11{F0: v6}
            _ = v81
            _ = v82
        } else {
            v83 := GenS12{F0: 2147483647}
            _ = v83
        }
        v84 := GenS2{F0: ((v1 + "foo bar") + v4), F1: v28}
        _ = v84
    } else {
        switch v2 {
        case -9223372036854775808:
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
        default:
            var v85 int8 = v6
            _ = v85
        }
        var v86 int64 = 0
        for fi87 := 0; fi87 < 4; fi87++ {
            v88 := GenS13{F0: (!false), F1: 4294967295}
            v89 := GenS14{F0: false, F1: 2, F2: -2147483648}
            var v90 int16 = v28
            _ = fi87
            _ = v88
            _ = v89
            _ = v90
        }
        _ = v86
    }
    var v91 string = v4
    switch v24 {
    case 32767:
        fmt.Println(v2, v91)
    case 0:
        fmt.Println(v28, v28)
    }
    v2 = 1
    if true {
        fmt.Println(v5)
        if (0 != 924) {
            var v92 uint64 = 18446744073709551615
            if len(v25) > 0 {
                v25[0] = v24
            }
            _ = v92
        }
    } else {
        fmt.Println(v4, v24)
        fmt.Println(v24, v5)
    }
    var v93 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v91
    _ = v93
}
