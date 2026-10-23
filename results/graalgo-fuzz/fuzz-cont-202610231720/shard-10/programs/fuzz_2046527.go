package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int64; F1 int; F2 bool }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 bool }
func genFn0(p0 uint32, p1 uint) int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{-1, 42, fi7}
        var v9 float32 = (3.14159 - (0.0 * 1e10))
        var v10 bool = (true && true)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch 66 {
    case 4294967295:
        if (1e10 > 1.5) {
            v11 := GenS1{F0: -1, F1: 375}
            var v12 int32 = 0
            _ = v11
            _ = v12
        }
    case -9223372036854775808:
        if (-158 == -32768) {
            v13 := []int{-999}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = v13
        } else {
            v14 := GenS1{F0: 511, F1: -904}
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 float64 = -2.25
                var v17 uint32 = (p0 >> ((uint(p0)) & 63))
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 int8 = 127
            _ = v14
            _ = v18
        }
    case -818:
        var v19 uint16 = 908
        _ = v19
    }
    v20 := []int{128, 2147483647, -1}
    _ = p0
    _ = p1
    _ = v20
    return 61
}

func main() {
    var v1 int64 = 219
    var v2 uint32 = 0
    var v3 uint = 0
    var v4 bool = true
    var v5 int8 = -1
    v6 := []uint32{(^(v2)), 42, 442}
    v3 = v3
    v6 = append(v6, v2)
    if (v4 || v4) {
        if (113 < 0) {
            var v21 string = "a"
            v22 := GenS1{F0: 1, F1: (-(v1))}
            var v23 uint8 = 24
            _ = v21
            _ = v22
            _ = v23
        } else {
            var v24 int32 = -512
            var v25 uint32 = 42
            var v26 string = "foo bar"
            _ = v24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 uint64 = 972
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := GenS1{F0: 934, F1: (v1 << ((65535) & 63))}
                switch "foo bar" {
                case "a":
                    if (v4 || v4) {
                        v31 := GenS2{F0: 123}
                        v5 -= v5
                        var v32 uint32 = v2
                        var v33 int = fi29
                        v28 += v28
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                default:
                    if (255 != 2) {
                        if len(v6) > 0 {
                            v6[0] = 65535
                        }
                    } else {
                        if v4 {
                            _ = len(v6)
                            var v34 uint = 301
                            _ = v34
                        } else {
                            switch fi29 {
                            case 835:
                                v35 := GenS2{F0: 0}
                                _ = v35
                            case 32767:
                                var v36 string = "a"
                                _ = v36
                            case 146:
                                v1 = v1
                            default:
                                switch "" {
                                case "\n\t":
                                    if (!v4) {
                                        var v37 uint = v3
                                        v37 = (v3 + 542)
                                        _ = v37
                                    } else {
                                        if len(v6) > 0 {
                                            v6[0] = uint32(v28)
                                        }
                                    }
                                case "hello":
                                    var v38 uint64 = v28
                                    _ = v38
                                }
                            }
                            v28 = v28
                        }
                    }
                }
                if (2 == 1) {
                    v28 -= v28
                    v39 := GenS2{F0: (v2 ^ (v2 << (((v3 * (v3 * (v3 | 180)))) & 63)))}
                    _ = v39
                } else {
                    if v4 {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    } else {
                        var v40 uint64 = v28
                        _ = v40
                    }
                }
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        v5 = (v5 % (((-(v5))) | 1))
    } else {
        if true {
            if v4 {
                if v4 {
                    for fi41 := 0; fi41 < 13; fi41++ {
                        if true {
                            _ = len(v6)
                            fmt.Println(v1, v5)
                        } else {
                            v6 = append(v6, uint32(v3))
                        }
                        _ = fi41
                    }
                    var v42 uint = (v3 & v3)
                    _ = v42
                }
                for fi43 := 0; fi43 < 3; fi43++ {
                    for fi44 := 0; fi44 < 1; fi44++ {
                        var v45 bool = (65535 != v2)
                        if v45 {
                            var v46 string = "hello"
                            v6 = append(v6, v2)
                            v6 = append(v6, v2)
                            var v47 int8 = (v5 >> ((576) & 63))
                            _ = v46
                            _ = v47
                        }
                        var v48 uint32 = 65535
                        v5 = 1
                        _ = fi44
                        _ = v45
                        _ = v48
                    }
                    v49 := GenS2{F0: 1}
                    switch ((("foo bar" + "foo bar") + "hello") + "foo bar") {
                    case "hello":
                        if (v4 || true) {
                            v50 := GenS3{F0: fi43, F1: 42}
                            _ = v50
                        }
                    case "a":
                        switch "" {
                        case "foo bar":
                            v51 := GenS2{F0: 83}
                            _ = v51
                        }
                    default:
                        v3 -= v3
                    }
                    _ = fi43
                    _ = v49
                }
                var v52 float32 = ((float32(1.5) / (float32(-2.25) - 3.14159)) + -2.25)
                _ = v52
            }
            var v53 int8 = v5
            _ = v53
        } else {
            var v54 int8 = (v5 << ((222) & 63))
            _ = v54
        }
    }
    if (!(936 != -32768)) {
        var v55 int8 = v5
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch -870 {
        case 9223372036854775807:
            if false {
                v56 := GenS3{F0: 127, F1: 0}
                if true {
                    v3 = v3
                } else {
                    v2 += 4294967295
                    fmt.Println(v55)
                }
                _ = v56
            }
        }
        _ = v55
    } else {
        for fi57 := 0; fi57 < 19; fi57++ {
            for fi58 := 0; fi58 < 8; fi58++ {
                v3 = v3
                _ = fi58
            }
            if len(v6) > 0 {
                v6[0] = (v2 / (((v2 + (v2 << (((v3 - (^(v3)))) & 63)))) | 1))
            }
            for fi59 := 0; fi59 < 14; fi59++ {
                v4 = v4
                _ = fi59
            }
            _ = fi57
        }
        v60 := GenS1{F0: -1, F1: 528}
        _ = v60
    }
    v61 := GenS4{F0: -692, F1: -781, F2: v4}
    var v62 bool = (2147483647 <= -2147483648)
    v63 := GenS4{F0: -382, F1: 32767, F2: (-261 >= 2147483647)}
    if v62 {
        v2 -= (v2 << ((v3) & 63))
    } else {
        v64 := GenS5{F0: 2, F1: 127}
        var v65 float64 = (0.1 - 3.14159)
        _ = v64
        _ = v65
    }
    if false {
        v66 := GenS2{F0: 42}
        var v67 uint32 = (^(v2))
        _ = v66
        _ = v67
    } else {
        fmt.Println(v2, v62, v4)
        if len(v6) > 0 {
            v6[0] = (v2 & 0)
        }
    }
    if v4 {
        var v68 float64 = (-2.25 * (-2.25 / 1e10))
        var v69 int32 = 42
        _ = v68
        _ = v69
    } else {
        var v70 float64 = (-2.25 + (0.1 + ((-2.25 + (1e10 - (-2.25 + 0.1))) - (((0.0 - (((1.5 - 1.5) * 1e10) / -2.25)) / ((3.14159 - 0.1) - 1.5)) + (float64(1.5) / 1e10)))))
        switch -316 {
        case 0:
            var v71 int8 = genFn0((v2 / ((v2) | 1)), (^(v3)))
            _ = v71
        case 284:
            for fi72 := 0; fi72 < 10; fi72++ {
                var v73 bool = true
                v74 := GenS6{F0: fi72}
                if (fi72 == (fi72 * (fi72 % ((fi72) | 1)))) {
                    v75 := GenS6{F0: fi72}
                    _ = v75
                }
                _ = fi72
                _ = v73
                _ = v74
            }
        default:
            switch 0 {
            case 127:
                if (-9223372036854775808 == -9223372036854775808) {
                    var v76 uint32 = 4294967295
                    v77 := GenS6{F0: -1}
                    v78 := GenS7{F0: 0}
                    _ = v76
                    _ = v77
                    _ = v78
                } else {
                    switch 469 {
                    case 128:
                        for fi79 := 0; fi79 < 18; fi79++ {
                            fmt.Println(fi79)
                            _ = len(v6)
                            _ = fi79
                        }
                    }
                }
            case 741:
                var v80 int8 = (^(v5))
                _ = v80
            }
        }
        _ = v70
    }
    for fi81 := 0; fi81 < 17; fi81++ {
        fmt.Println(v2, v2, v1)
        _ = len(v6)
        for fi82 := 0; fi82 < 1; fi82++ {
            fmt.Println(v1, v1)
            var v83 uint = (v3 / (((^(v3))) | 1))
            _ = len(v6)
            _ = fi82
            _ = v83
        }
        _ = fi81
    }
    for fi84 := 0; fi84 < 14; fi84++ {
        v85 := GenS8{F0: v5, F1: v4}
        var v86 float32 = (3.14159 / 1.0)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi84
        _ = v85
        _ = v86
    }
    for fi87 := 0; fi87 < 15; fi87++ {
        if ((fi87 != (fi87 ^ -339)) || true) {
            v88 := GenS8{F0: (v5 & genFn0(2147483648, (v3 | (v3 & 1)))), F1: (!(v1 >= (v1 % ((-82) | 1))))}
            fmt.Println(v2, v1, v1)
            v5 = genFn0((v2 << (((v3 & 42)) & 63)), v3)
            _ = v88
        } else {
            if (v5 > v5) {
                v3 -= 91
                v89 := GenS3{F0: -1, F1: 1}
                _ = v89
            }
            for fi90 := 0; fi90 < 7; fi90++ {
                v91 := GenS6{F0: fi87}
                _ = fi90
                _ = v91
            }
        }
        var v92 int8 = (v5 >> ((v3) & 63))
        _ = fi87
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v63
}
