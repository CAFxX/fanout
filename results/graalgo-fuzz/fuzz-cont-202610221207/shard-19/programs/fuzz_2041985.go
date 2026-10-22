package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint; F1 uint; F2 int }
type GenS7 struct { F0 bool; F1 int64; F2 int64 }
type GenS8 struct { F0 uint; F1 int32; F2 int32 }
func genFn0() uint16 {
    v4 := []int{482, -270, 2147483647, 32767}
    v4 = append(v4, 3)
    var v5 uint8 = 163
    var v6 uint64 = 42
    _ = v4
    _ = v5
    _ = v6
    return 457
}

func main() {
    var v1 float32 = (((float32(0.0) / -2.25) / 0.1) + (float32(0.0) / 1.0))
    v2 := make(map[int]string)
    var v3 uint32 = 1
    var v7 int16 = 32767
    fmt.Println(v1)
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{fi8, (fi8 * (fi8 - (fi8 ^ fi8))), -2147483648, (-(fi8))}
        _ = fi8
        _ = v9
    }
    v1 = (v1 - v1)
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: (v3 - v3), F1: 482, F2: -1}
        if false {
            if (1 > 1) {
                v12 := []int{(fi10 | 817), (fi10 * -132)}
                _ = v12
            } else {
                v13 := GenS2{F0: 0, F1: 18446744073709551615}
                v14 := []int{32768}
                _ = v13
                _ = v14
            }
            if false {
                if (1 > 42) {
                    switch "hello" {
                    case "foo bar":
                        switch fi10 {
                        case 2147483647:
                            delete(v2, 87)
                        case 0:
                            if (!true) {
                                var v15 int16 = (v7 | v7)
                                v16 := []int{4294967295, fi10}
                                v17 := GenS3{F0: (v3 % ((2147483648) | 1)), F1: v3}
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        case -206:
                            var v18 int16 = (v7 >> ((42) & 63))
                            _ = v18
                        }
                    }
                    if false {
                        for fi19 := 0; fi19 < 6; fi19++ {
                            var v20 int64 = 9223372036854775807
                            _ = fi19
                            _ = v20
                        }
                        v3 = v3
                        v2[(fi10 * fi10)] = "a"
                    }
                    v21 := GenS2{F0: 1, F1: 62}
                    _ = v21
                } else {
                    var v22 int16 = int16(fi10)
                    var v23 int32 = 842
                    v23 = v23
                    _ = v22
                    _ = v23
                }
                switch (("a" + "\n\t") + "") {
                case "hello":
                    for fi24 := 0; fi24 < 3; fi24++ {
                        var v25 uint16 = 32767
                        _ = fi24
                        _ = v25
                    }
                }
                fmt.Println(fi10, fi10, v7)
            }
            v26 := GenS4{F0: 2}
            _ = v26
        } else {
            if false {
                var v27 int16 = (v7 + v7)
                var v28 uint = 195
                fmt.Println(v3, v1)
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 18; fi29++ {
                    v3 = v3
                    var v30 uint32 = v3
                    _ = fi29
                    _ = v30
                }
            }
            var v31 uint32 = (v3 | v3)
            for fi32 := 0; fi32 < 8; fi32++ {
                var v33 bool = (246 < 127)
                if v33 {
                    v34 := GenS5{F0: 24, F1: 907}
                    switch "" {
                    case "a":
                        v35 := []int{0}
                        _ = v35
                    default:
                        if false {
                            v33 = (!true)
                            if (677 != 42) {
                                var v36 int64 = 1
                                var v37 uint = 255
                                if (v1 < v1) {
                                    v38 := GenS2{F0: 2, F1: 473}
                                    v2[(-(fi10))] = ("foo bar" + "")
                                    _ = v38
                                }
                                if v33 {
                                    v39 := []int{fi10}
                                    switch -988 {
                                    case -522:
                                        v36 += v36
                                    }
                                    _ = v39
                                } else {
                                    var v40 float64 = 0.1
                                    _ = v40
                                }
                                var v41 int32 = int32(v37)
                                _ = v36
                                _ = v37
                                _ = v41
                            }
                        }
                    }
                    _ = v34
                } else {
                    switch 9223372036854775807 {
                    case -993:
                        var v42 uint = 255
                        _ = v42
                    case -80:
                        fmt.Println(v1, v33, v7)
                    case 91:
                        var v43 int32 = 0
                        _ = v43
                    default:
                        fmt.Println(v3, v1)
                    }
                    v7 = v7
                }
                fmt.Println(fi10, v33)
                v44 := v2[(fi10 + fi32)]
                var v45 int16 = (v7 - v7)
                _ = len(v2)
                _ = fi32
                _ = v33
                _ = v44
                _ = v45
            }
            _ = v31
        }
        _ = fi10
        _ = v11
    }
    var v46 bool = (0 != 255)
    v3 = v3
    if v46 {
        switch -860 {
        case 128:
            if (574 < 42) {
                if v46 {
                    var v47 uint = 2
                    fmt.Println(v47)
                    fmt.Println(v7)
                    _ = v47
                }
            } else {
                v48 := GenS3{F0: v3, F1: v3}
                v7 += 866
                for fi49 := 0; fi49 < 4; fi49++ {
                    v50 := []int{fi49}
                    for fi51 := 0; fi51 < 17; fi51++ {
                        switch "" {
                        case "":
                            if len(v50) >= 2 {
                                v50 = v50[1:len(v50)]
                            }
                        }
                        if (!(-28 != fi49)) {
                            v50 = append(v50, (fi51 / ((fi51) | 1)))
                            switch ("foo bar" + "") {
                            case "hello":
                                var v52 uint8 = 128
                                _ = v52
                            case "\n\t":
                                var v53 uint8 = 255
                                _ = v53
                            case "a":
                                v3 = v3
                            default:
                                v1 -= 0.0
                            }
                            fmt.Println(v1, v7, v3)
                            v46 = true
                        } else {
                            var v54 int16 = v7
                            v55 := GenS2{F0: 65535, F1: 247}
                            if len(v50) > 0 {
                                v50[0] = fi51
                            }
                            _ = v54
                            _ = v55
                        }
                        if v46 {
                            v7 = v7
                            v56 := GenS6{F0: 42, F1: 42, F2: fi49}
                            _ = v56
                        }
                        _ = fi51
                    }
                    for fi57 := 0; fi57 < 5; fi57++ {
                        if ("a" == (("a" + "foo bar") + "a")) {
                            var v58 int = -946
                            v59 := GenS7{F0: v46, F1: -80, F2: -1}
                            v60 := GenS2{F0: 242, F1: 194}
                            v46 = (!(-53 <= -9223372036854775808))
                            var v61 uint8 = 138
                            _ = v58
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                        if v46 {
                            var v62 int16 = (v7 >> ((255) & 63))
                            if (false || v46) {
                                v63 := GenS2{F0: 255, F1: 2}
                                v62 -= v7
                                _ = v63
                            }
                            v64 := GenS8{F0: 119, F1: 1, F2: -681}
                            _ = v62
                            _ = v64
                        } else {
                            v65 := GenS4{F0: 2}
                            if (!v46) {
                                v1 -= ((v1 + (float32(0.1) / (((0.0 * (1e10 * v1)) / float32(1.5)) * (v1 + (((float32(0.1) * 1e10) * v1) - (float32(1.5) - v1)))))) + v1)
                            } else {
                                fmt.Println(v7, fi57)
                                v66 := GenS5{F0: 1, F1: 0}
                                _ = v66
                            }
                            v46 = (81 >= 18446744073709551615)
                            _ = v65
                        }
                        _ = fi57
                    }
                    _ = fi49
                    _ = v50
                }
                _ = v48
            }
        default:
            for fi67 := 0; fi67 < 6; fi67++ {
                switch "foo bar" {
                case "hello":
                    switch fi67 {
                    case 387:
                        var v68 int8 = -128
                        _ = v68
                    case -616:
                        var v69 string = ("hello" + "\n\t")
                        _ = v69
                    }
                case "foo bar":
                    v7 -= v7
                case "a":
                    v1 += 1e10
                default:
                    for fi70 := 0; fi70 < 20; fi70++ {
                        var v71 string = ("foo bar" + ("hello" + "foo bar"))
                        _ = fi70
                        _ = v71
                    }
                }
                v46 = v46
                for fi72 := 0; fi72 < 12; fi72++ {
                    delete(v2, (fi67 >> ((18446744073709551615) & 63)))
                    v3 = (v3 % ((v3) | 1))
                    if v46 {
                        var v73 uint64 = 777
                        if (v73 != v73) {
                            _ = len(v2)
                            v73 = (v73 << ((255) & 63))
                        }
                        var v74 uint = 18446744073709551615
                        _ = v73
                        _ = v74
                    }
                    if v46 {
                        v46 = (!v46)
                        v46 = (false && false)
                        if v46 {
                            v3 -= (v3 << ((0) & 63))
                            if v46 {
                                v1 = 3.14159
                            } else {
                                fmt.Println(v1, v3, v7)
                                var v75 int = fi67
                                _ = v75
                            }
                            var v76 int8 = 0
                            _ = v76
                        } else {
                            var v77 float64 = 0.0
                            var v78 uint32 = (v3 / (((^(v3))) | 1))
                            _ = v77
                            _ = v78
                        }
                    } else {
                        v79 := []int{}
                        var v80 uint16 = genFn0()
                        v46 = v46
                        _ = v79
                        _ = v80
                    }
                    v1 -= v1
                    _ = fi72
                }
                _ = fi67
            }
        }
    }
    var v81 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v7:", v7)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v7
    _ = v81
}
