package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 int; F2 uint16 }
type GenS5 struct { F0 int8; F1 int8; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int }
func genFn0(p0 int, p1 uint64, p2 int) string {
    var v7 int = -523
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{}
        var v10 uint = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}
func genFn1(p0 int8, p1 int, p2 uint8) string {
    p1 += (p1 - p1)
    v11 := GenS1{F0: 1, F1: 505}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return ("\n\t" + "a")
}

func main() {
    var v1 int32 = 801
    v2 := make(map[int]int)
    var v3 uint8 = 42
    var v4 uint8 = (v3 / ((1) | 1))
    var v5 float32 = 1.5
    v6 := []uint32{1, 4294967295, 530}
    var v12 int16 = 42
    fmt.Println(v4)
    v5 += float32(1.5)
    var v13 float64 = 3.14159
    var v14 string = "hello"
    for fi15 := 0; fi15 < 3; fi15++ {
        if false {
            if true {
                fmt.Println(v3, v12)
                fmt.Println(v4, v3, v1)
                v1 = (v1 & v1)
            } else {
                switch v14 {
                case "":
                    var v16 uint32 = 578
                    _ = v16
                case "foo bar":
                    fmt.Println(v4)
                case "\n\t":
                    v13 = v13
                default:
                    v12 += v12
                }
            }
            var v17 uint8 = 128
            _ = v17
        } else {
            v14 = v14
            if true {
                var v18 int16 = (v12 | -309)
                _ = v18
            } else {
                fmt.Println(fi15)
            }
            v19 := GenS1{F0: 1, F1: 2}
            _ = v19
        }
        v20 := GenS2{F0: 255, F1: 42}
        if ((v12 % (((v12 & (v12 & (v12 << ((1) & 63))))) | 1)) <= (v12 | -1)) {
            switch "hello" {
            case "\n\t":
                for fi21 := 0; fi21 < 3; fi21++ {
                    switch fi21 {
                    case 2147483647:
                        if false {
                            v12 += 32767
                            if true {
                                v22 := GenS1{F0: 0, F1: 2}
                                v23 := GenS2{F0: 1, F1: 17}
                                if true {
                                    var v24 uint32 = 315
                                    _ = v24
                                }
                                v13 += v13
                                v6 = append(v6, 42)
                                if false {
                                    var v25 int16 = (v12 % (((-(v12))) | 1))
                                    var v26 float32 = (float32((0.0 - 1e10)) + v5)
                                    _ = v25
                                    _ = v26
                                }
                                v3 += (v4 & (v3 / ((v4) | 1)))
                                _ = v22
                                _ = v23
                            } else {
                                fmt.Println(v14, fi15, v4)
                                v3 = 166
                                v1 = (v1 + (^(v1)))
                            }
                        } else {
                            var v27 int16 = 32767
                            var v28 string = "foo bar"
                            var v29 int64 = -9223372036854775808
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    case 42:
                        v13 = 1.5
                    case -433:
                        if ((789 <= 133) || true) {
                            var v30 float64 = v13
                            _ = v30
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v4 += v3
                            switch (fi21 << ((0) & 63)) {
                            case 32767:
                                v14 = v14
                            case 912:
                                v31 := GenS3{F0: false, F1: v4, F2: 1}
                                _ = v31
                            case -9223372036854775808:
                                v5 -= ((3.14159 - v5) / v5)
                                v13 = 0.0
                            }
                        }
                    default:
                        v6 = append(v6, 2)
                    }
                    if (!false) {
                        switch fi15 {
                        case -9223372036854775808:
                            if (-128 != 100) {
                                if true {
                                    v32 := GenS1{F0: 0, F1: uint16(v4)}
                                    v33 := GenS4{F0: 255, F1: fi21, F2: 1}
                                    if false {
                                        v4 -= v3
                                    } else {
                                        var v34 uint = 373
                                        _ = len(v6)
                                        _ = v34
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                v4 -= (v3 + v3)
                            } else {
                                switch v14 {
                                case "a":
                                    var v35 uint16 = 2
                                    _ = v35
                                case "foo bar":
                                    if false {
                                        switch 32768 {
                                        case 0:
                                            var v36 bool = false
                                            v37 := GenS5{F0: -128, F1: 42, F2: 1}
                                            _ = v36
                                            _ = v37
                                        case 9223372036854775807:
                                            if true {
                                                if (3.14159 < -2.25) {
                                                    var v38 int8 = -18
                                                    var v39 uint8 = (v4 + (v4 - v4))
                                                    v40 := GenS2{F0: 255, F1: (-(v38))}
                                                    v41 := GenS1{F0: 65535, F1: 2}
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                }
                                                v42 := GenS6{F0: v14}
                                                switch 32768 {
                                                case 985:
                                                    v43 := GenS4{F0: 2, F1: fi21, F2: 32767}
                                                    _ = v43
                                                default:
                                                    v4 = v3
                                                }
                                                _ = v42
                                            }
                                        case -423:
                                            var v44 uint = 2
                                            v45 := GenS5{F0: -82, F1: -52, F2: -74}
                                            v46 := v2[2147483647]
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        default:
                                            v3 = (^(v3))
                                        }
                                        v47 := GenS5{F0: -77, F1: 100, F2: 0}
                                        _ = v47
                                    }
                                case "":
                                    v48 := GenS2{F0: 65535, F1: -128}
                                    _ = v48
                                default:
                                    if true {
                                        v5 = (v5 + 1e10)
                                        fmt.Println(v4, v3, v4)
                                        switch v14 {
                                        case "hello":
                                            var v49 float64 = (0.1 / v13)
                                            _ = v49
                                        default:
                                            v12 = (v12 * -108)
                                        }
                                        v50 := GenS6{F0: ((v14 + v14) + v14)}
                                        v12 += v12
                                        _ = v50
                                    } else {
                                        v2[fi15] = fi21
                                    }
                                }
                                switch (fi21 & 32767) {
                                case 0:
                                    var v51 uint = 2
                                    _ = v51
                                case 128:
                                    if (42 == 27) {
                                        var v52 float32 = ((3.14159 + (float32((3.14159 - v13)) - v5)) + v5)
                                        _ = v52
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v53 uint64 = 116
                                        _ = v53
                                    }
                                default:
                                    delete(v2, 434)
                                }
                            }
                        case -716:
                            v3 -= v3
                        case -160:
                            v54 := GenS5{F0: 1, F1: 42, F2: 42}
                            _ = v54
                        }
                        var v55 int64 = 1
                        v14 = "hello"
                        var v56 int64 = -9223372036854775808
                        v6 = append(v6, 0)
                        _ = v55
                        _ = v56
                    }
                    var v57 uint8 = 239
                    v12 += (v12 + v12)
                    v5 = (v5 * v5)
                    v3 = (v3 & v3)
                    switch v14 {
                    case "\n\t":
                        v58 := GenS6{F0: ("" + "\n\t")}
                        _ = v58
                    case "":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        switch v14 {
                        case "\n\t":
                            v3 = v3
                        case "hello":
                            var v59 string = v14
                            _ = v59
                        case "a":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            switch (v14 + "a") {
                            case "":
                                fmt.Println(fi15, v5, v3)
                            default:
                                var v60 int = -77
                                v57 += 203
                                switch 127 {
                                case 4294967295:
                                    v57 += v3
                                case -1:
                                    var v61 int8 = 106
                                    _ = v61
                                default:
                                    var v62 uint16 = 638
                                    _ = v62
                                }
                                _ = v60
                            }
                        }
                    }
                    fmt.Println(v14, v4)
                    _ = fi21
                    _ = v57
                }
            case "a":
                if (!(!(1000 < (v12 / ((v12) | 1))))) {
                    var v63 int8 = 102
                    _ = v63
                }
            case "":
                v64 := v2[fi15]
                _ = v64
            default:
                if (!false) {
                    if false {
                        if true {
                            v3 = v3
                            v12 += (v12 ^ int16(v3))
                        }
                        fmt.Println(v5, v12)
                    } else {
                        for fi65 := 0; fi65 < 16; fi65++ {
                            if (!false) {
                                var v66 string = "hello"
                                switch fi65 {
                                case 458:
                                    fmt.Println(v66)
                                case 0:
                                    _ = len(v2)
                                }
                                var v67 uint64 = 835
                                var v68 bool = (42 <= 0)
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                            v4 = (v4 ^ v4)
                            switch ((v14 + v14) + "foo bar") {
                            case "":
                                var v69 uint64 = 42
                                _ = v69
                            case "hello":
                                fmt.Println(v13)
                                v6 = append(v6, 2)
                            case "\n\t":
                                v5 = ((v5 - (float32((v13 * v13)) + v5)) * 1e10)
                            }
                            _ = fi65
                        }
                        _ = len(v6)
                        var v70 bool = true
                        _ = v70
                    }
                    var v71 int16 = v12
                    _ = v71
                }
            }
            for fi72 := 0; fi72 < 19; fi72++ {
                if false {
                    delete(v2, fi15)
                    v3 += (v3 & (v3 - 50))
                }
                _ = len(v2)
                _ = fi72
            }
            v13 += v13
        } else {
            var v73 uint64 = 1
            var v74 string = v14
            fmt.Println(fi15, v4)
            _ = v73
            _ = v74
        }
        _ = fi15
        _ = v20
    }
    var v75 int64 = -908
    fmt.Println(v3, v75)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v76 := GenS1{F0: 42, F1: 0}
    var v77 int32 = -334
    var v78 int = 320
    v14 = ("a" + "foo bar")
    for fi79 := 0; fi79 < 3; fi79++ {
        fmt.Println(v75, v4)
        v80 := GenS4{F0: 665, F1: 9223372036854775807, F2: 1}
        v2[fi79] = (v78 & 711)
        _ = fi79
        _ = v80
    }
    for fi81 := 0; fi81 < 14; fi81++ {
        v2[v78] = fi81
        v82 := GenS7{F0: (fi81 << ((42) & 63))}
        _ = fi81
        _ = v82
    }
    var v83 int8 = 93
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v83
}
