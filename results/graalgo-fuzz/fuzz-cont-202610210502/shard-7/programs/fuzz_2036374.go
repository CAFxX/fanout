package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16; F1 int32; F2 int16 }
type GenS4 struct { F0 uint; F1 string; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 int64; F2 int16 }
type GenS9 struct { F0 int8 }
func genFn0(p0 uint8, p1 bool, p2 uint32) int16 {
    if (18446744073709551615 <= 105) {
        var v4 string = "a"
        _ = v4
    }
    v5 := GenS1{F0: (p2 << ((101) & 63)), F1: 360}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = uint32(p0)
        if p1 {
            v8 := GenS1{F0: 2147483648, F1: fi6}
            if true {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 int8 = 42
                    if (-32768 <= -844) {
                        v11 := []int{(fi6 << ((657) & 63)), (fi9 ^ (fi6 << ((240) & 63)))}
                        fmt.Println(p0)
                        _ = v11
                    }
                    _ = fi9
                    _ = v10
                }
                var v12 int64 = 1
                v13 := []int{127, fi6}
                _ = v12
                _ = v13
            } else {
                v7 = p2
                p0 = (p0 << ((65535) & 63))
            }
            v14 := GenS1{F0: 4294967295, F1: fi6}
            _ = v14
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 758
}
func genFn1(p0 int32) bool {
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: 7, F1: fi15}
        _ = fi15
        _ = v16
    }
    var v17 uint8 = 2
    p0 = (p0 & 1)
    if false {
        v18 := []int{-9223372036854775808, 42}
        fmt.Println(v17, p0)
        _ = v18
    } else {
        fmt.Println(p0, v17)
        v19 := []int{109}
        _ = v19
    }
    var v20 int16 = genFn0(v17, false, 301)
    _ = p0
    _ = v17
    _ = v20
    return true
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (255 != 1)
    var v3 uint64 = 2
    fmt.Println(v2, v2)
    if false {
        v21 := []int{-976, 4294967295, 32767}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 7; fi22++ {
            v2 = v2
            fmt.Println(fi22, fi22)
            var v23 uint = 255
            _ = fi22
            _ = v23
        }
        var v24 float64 = (0.1 + 1e10)
        _ = v24
    }
    var v25 string = "a"
    var v26 float32 = -2.25
    v27 := GenS1{F0: 691, F1: 127}
    for fi28 := 0; fi28 < 19; fi28++ {
        v3 -= (^(v3))
        _ = fi28
    }
    if false {
        var v29 int8 = 89
        v30 := []int{}
        if (1.5 != v26) {
            v31 := GenS2{F0: 9223372036854775807, F1: -1, F2: -2147483648}
            v26 = (v26 - v26)
            _ = v31
        } else {
            if v2 {
                switch "hello" {
                case "hello":
                    if (244 >= 128) {
                        var v32 float64 = -2.25
                        for fi33 := 0; fi33 < 16; fi33++ {
                            for fi34 := 0; fi34 < 15; fi34++ {
                                v29 = (v29 << ((353) & 63))
                                v26 -= (v26 / 1.0)
                                _ = fi34
                            }
                            for fi35 := 0; fi35 < 17; fi35++ {
                                if len(v30) > 0 {
                                    v30[0] = 209
                                }
                                v3 -= v3
                                switch ("foo bar" + v25) {
                                case "\n\t":
                                    _ = len(v30)
                                default:
                                    v36 := GenS3{F0: 1, F1: 42, F2: 0}
                                    _ = v36
                                }
                                _ = fi35
                            }
                            _ = fi33
                        }
                        _ = v32
                    } else {
                        var v37 string = ""
                        for fi38 := 0; fi38 < 17; fi38++ {
                            if v2 {
                                v39 := GenS2{F0: 124, F1: 2147483647, F2: 0}
                                var v40 uint16 = 1
                                var v41 uint = 4294967295
                                _ = v39
                                _ = v40
                                _ = v41
                            } else {
                                v2 = ("foo bar" == v37)
                                if v2 {
                                    var v42 int16 = 1000
                                    var v43 uint32 = 31
                                    for fi44 := 0; fi44 < 8; fi44++ {
                                        v1[(v43 + v43)] = ((((v37 + (v37 + ((v25 + "") + (("a" + v37) + (v25 + "\n\t"))))) + (((v25 + v25) + v37) + v37)) + ((v25 + (v37 + v25)) + ((v37 + ((((v37 + (((v25 + v37) + "hello") + v25)) + (v25 + v37)) + v37) + v37)) + v25))) + v25)
                                        var v45 uint64 = (^(v3))
                                        switch (fi44 >> ((42) & 63)) {
                                        case 43:
                                            v46 := GenS2{F0: int64(v42), F1: 0, F2: 0}
                                            _ = v46
                                        default:
                                            var v47 uint = 1
                                            _ = v47
                                        }
                                        v25 = v25
                                        _ = fi44
                                        _ = v45
                                    }
                                    _ = v42
                                    _ = v43
                                } else {
                                    if v2 {
                                        v3 -= v3
                                    } else {
                                        var v48 int16 = 756
                                        var v49 int = fi38
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                v50 := GenS4{F0: 344, F1: v25, F2: 0}
                                _ = v50
                            }
                            v51 := GenS5{F0: 2147483647}
                            var v52 bool = v2
                            _ = fi38
                            _ = v51
                            _ = v52
                        }
                        if (v25 < v25) {
                            v53 := GenS2{F0: 1, F1: 0, F2: 2147483647}
                            _ = v53
                        } else {
                            _ = len(v30)
                            if ((v37 + ("hello" + "foo bar")) == ((v37 + "foo bar") + v25)) {
                                v3 -= (v3 | v3)
                                v3 -= v3
                                if len(v30) >= 2 {
                                    v30 = v30[1:len(v30)]
                                }
                            }
                        }
                        _ = v37
                    }
                }
                fmt.Println(v2)
                fmt.Println(v29, v29, v29)
            }
            fmt.Println(v3)
        }
        _ = v29
        _ = v30
    }
    for fi54 := 0; fi54 < 7; fi54++ {
        fmt.Println(v2, v2, v25)
        _ = fi54
    }
    if true {
        if false {
            v55 := v1[16]
            var v56 bool = false
            switch -1 {
            case 366:
                for fi57 := 0; fi57 < 8; fi57++ {
                    switch (((v55 + (v25 + v25)) + "a") + v55) {
                    case "\n\t":
                        for fi58 := 0; fi58 < 17; fi58++ {
                            var v59 bool = true
                            if (false && (!(0 != 2147483647))) {
                                v56 = (65535 <= 2)
                            } else {
                                v60 := GenS2{F0: 42, F1: int32(fi57), F2: 148}
                                if true {
                                    v61 := GenS1{F0: 42, F1: fi57}
                                    v56 = true
                                    _ = v61
                                }
                                var v62 string = "\n\t"
                                fmt.Println(fi58, v59)
                                _ = v60
                                _ = v62
                            }
                            v63 := GenS2{F0: 9223372036854775807, F1: 42, F2: 0}
                            fmt.Println(fi58)
                            _ = fi58
                            _ = v59
                            _ = v63
                        }
                    case "foo bar":
                        for fi64 := 0; fi64 < 10; fi64++ {
                            fmt.Println(v55)
                            _ = fi64
                        }
                    }
                    _ = fi57
                }
            case 32767:
                for fi65 := 0; fi65 < 5; fi65++ {
                    v66 := []int{}
                    v25 = ""
                    v67 := GenS5{F0: 2147483647}
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            }
            _ = v55
            _ = v56
        } else {
            for fi68 := 0; fi68 < 16; fi68++ {
                v69 := GenS6{F0: v2}
                if (-342 >= 42) {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        v71 := GenS7{F0: v2}
                        v72 := []int{(fi70 / (((fi68 % ((42) | 1))) | 1)), (fi70 - (fi68 * fi68)), (^(fi70))}
                        if (!(!true)) {
                            var v73 float64 = 0.0
                            var v74 uint32 = 4294967295
                            _ = v73
                            _ = v74
                        }
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                    var v75 int64 = -305
                    _ = v75
                }
                var v76 float64 = float64(0.1)
                _ = fi68
                _ = v69
                _ = v76
            }
        }
        var v77 uint32 = 2147483648
        _ = v77
    }
    var v78 int8 = 127
    fmt.Println(v26, v3, v25)
    if v2 {
        v79 := []int{}
        if ((true && (v2 || false)) && v2) {
            var v80 uint8 = 128
            _ = v80
        }
        v81 := GenS8{F0: v78, F1: 40, F2: -447}
        _ = v79
        _ = v81
    }
    var v82 string = ("" + (v25 + (((v25 + (v25 + "a")) + (v25 + v25)) + ("" + "\n\t"))))
    if (v25 >= v25) {
        v2 = v2
        if v2 {
            v83 := GenS1{F0: 2, F1: 594}
            v25 = v82
            _ = v83
        }
        if v2 {
            switch 1 {
            case 9223372036854775807:
                for fi84 := 0; fi84 < 19; fi84++ {
                    var v85 int16 = -516
                    if true {
                        switch (fi84 % ((fi84) | 1)) {
                        case -9223372036854775808:
                            switch (fi84 % ((0) | 1)) {
                            case 32767:
                                var v86 float64 = 0.1
                                _ = v86
                            }
                        case -1:
                            v78 += -1
                        default:
                            v87 := GenS1{F0: 1, F1: (fi84 - -2147483648)}
                            _ = v87
                        }
                        var v88 uint32 = 4294967295
                        v89 := []int{}
                        _ = v88
                        _ = v89
                    } else {
                        v82 = "foo bar"
                        var v90 float32 = 0.1
                        _ = v90
                    }
                    _ = fi84
                    _ = v85
                }
            case 4294967295:
                if v2 {
                    var v91 int16 = int16(v3)
                    v91 -= v91
                    var v92 bool = false
                    _ = v91
                    _ = v92
                }
            default:
                var v93 uint = 0
                _ = v93
            }
        } else {
            v94 := []int{127}
            for fi95 := 0; fi95 < 3; fi95++ {
                v2 = true
                _ = fi95
            }
            _ = v94
        }
    } else {
        if (((1.5 * -2.25) + 1e10) == (0.0 - float64(((1.5 / v26) / v26)))) {
            for fi96 := 0; fi96 < 6; fi96++ {
                var v97 uint8 = 42
                _ = fi96
                _ = v97
            }
        } else {
            if (-297 != 941) {
                v98 := GenS9{F0: v78}
                _ = v98
            } else {
                for fi99 := 0; fi99 < 18; fi99++ {
                    if v2 {
                        var v100 uint8 = 125
                        fmt.Println(v25, v78, v100)
                        v101 := []int{(fi99 | (fi99 / ((703) | 1))), fi99, (-(fi99)), (fi99 ^ (^(fi99)))}
                        _ = v100
                        _ = v101
                    }
                    var v102 string = "\n\t"
                    _ = fi99
                    _ = v102
                }
            }
            delete(v1, 930)
        }
        v103 := []int{32768}
        _ = v103
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v78
    _ = v82
}
