package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 bool }
type GenS4 struct { F0 int8; F1 int; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 float32 = ((-2.25 + 3.14159) * 3.14159)
    var v4 float64 = (-2.25 / 3.14159)
    var v5 string = "hello"
    v6 := make(map[string]int)
    var v7 uint16 = 688
    fmt.Println(v4, v7, v7)
    if (1 < v7) {
        if true {
            v4 += (0.1 / v4)
            if true {
                if false {
                    switch 0 {
                    case -58:
                        v3 = (v3 / ((-2.25 * v3) - ((((3.14159 - (v3 + v3)) / 1.0) * v3) + 3.14159)))
                    case -960:
                        switch 0 {
                        case 32768:
                            for fi8 := 0; fi8 < 20; fi8++ {
                                if (v4 > (3.14159 + float64(float32((v4 - 0.0))))) {
                                    v9 := GenS1{F0: 2, F1: false}
                                    v5 = v5
                                    _ = v9
                                }
                                _ = fi8
                            }
                        default:
                            var v10 int16 = 1000
                            _ = v10
                        }
                    default:
                        var v11 int8 = 0
                        _ = v11
                    }
                    if true {
                        var v12 int8 = -128
                        if (!true) {
                            switch 905 {
                            case 78:
                                for fi13 := 0; fi13 < 1; fi13++ {
                                    var v14 bool = true
                                    if false {
                                        var v15 string = ("foo bar" + ("hello" + "a"))
                                        var v16 uint16 = v7
                                        var v17 uint32 = 2147483648
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    } else {
                                        for fi18 := 0; fi18 < 5; fi18++ {
                                            v7 += v7
                                            _ = fi18
                                        }
                                        switch fi13 {
                                        case 476:
                                            var v19 uint64 = 759
                                            _ = v19
                                        case 2147483647:
                                            var v20 uint64 = 159
                                            _ = v20
                                        case -234:
                                            for fi21 := 0; fi21 < 14; fi21++ {
                                                fmt.Println(v4)
                                                _ = fi21
                                            }
                                        }
                                    }
                                    var v22 int8 = (^(v12))
                                    _ = fi13
                                    _ = v14
                                    _ = v22
                                }
                            case 707:
                                for fi23 := 0; fi23 < 12; fi23++ {
                                    if (true || (2147483647 > -1)) {
                                        if true {
                                            v4 -= 0.0
                                            var v24 uint64 = 1
                                            v25 := []int{fi23, (^(fi23)), fi23, (fi23 % ((42) | 1))}
                                            _ = v24
                                            _ = v25
                                        } else {
                                            v26 := []int{fi23, (fi23 + 1), fi23, (fi23 * 621)}
                                            var v27 uint64 = uint64(v7)
                                            _ = v26
                                            _ = v27
                                        }
                                    } else {
                                        v12 = v12
                                        var v28 int = fi23
                                        _ = v28
                                    }
                                    _ = fi23
                                }
                            case 423:
                                v29 := []int{501}
                                _ = v29
                            default:
                                var v30 int64 = -77
                                _ = v30
                            }
                            v31 := GenS2{F0: (v7 * (v7 | v7))}
                            fmt.Println(v3, v7)
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 20; fi32++ {
                                v7 += v7
                                v12 += v12
                                v33 := v2[(fi32 + (-(fi32)))]
                                _ = fi32
                                _ = v33
                            }
                            var v34 float32 = float32(v4)
                            _ = v34
                        }
                        _ = v12
                    }
                    if (!false) {
                        v3 += 0.1
                    } else {
                        v35 := GenS3{F0: 9223372036854775807, F1: (v7 | (v7 >> ((42) & 63))), F2: false}
                        var v36 int64 = 433
                        _ = v35
                        _ = v36
                    }
                }
            } else {
                var v37 float32 = ((v3 + (v3 * (v3 * v3))) * (float32(v4) - v3))
                _ = len(v6)
                _ = v37
            }
        } else {
            var v38 uint64 = 42
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 bool = ((true && true) && ("" > v5))
                _ = fi39
                _ = v40
            }
            _ = v38
        }
        if true {
            var v41 uint16 = v7
            v42 := GenS4{F0: -61, F1: 0, F2: 9223372036854775807}
            var v43 uint8 = 42
            _ = v41
            _ = v42
            _ = v43
        } else {
            switch 32768 {
            case -1:
                v44 := GenS5{F0: v7, F1: 2, F2: 42}
                _ = v44
            default:
                if (false || (v4 < v4)) {
                    v7 = v7
                    v7 -= v7
                }
            }
            for fi45 := 0; fi45 < 19; fi45++ {
                var v46 bool = false
                v7 += 2
                v7 = (v7 << ((65535) & 63))
                _ = fi45
                _ = v46
            }
        }
        if true {
            for fi47 := 0; fi47 < 11; fi47++ {
                fmt.Println(v7, v7)
                var v48 float64 = (v4 + v4)
                _ = fi47
                _ = v48
            }
            if false {
                v49 := GenS4{F0: 100, F1: -1, F2: -504}
                var v50 int8 = 24
                v3 = (((float32(float64(v3)) + float32(v4)) * float32(float64(v3))) + v3)
                _ = v49
                _ = v50
            } else {
                var v51 int8 = 100
                v52 := GenS1{F0: 255, F1: (false && (true && (!false)))}
                if true {
                    var v53 uint = 4294967295
                    _ = v53
                } else {
                    for fi54 := 0; fi54 < 10; fi54++ {
                        fmt.Println(fi54)
                        for fi55 := 0; fi55 < 4; fi55++ {
                            delete(v1, 486)
                            switch fi54 {
                            case 127:
                                var v56 int64 = -1
                                _ = v56
                            default:
                                v57 := GenS1{F0: 0, F1: (1 >= 127)}
                                _ = v57
                            }
                            _ = fi55
                        }
                        var v58 uint8 = 127
                        _ = fi54
                        _ = v58
                    }
                    var v59 float64 = (v4 + ((0.1 + v4) * v4))
                    var v60 int32 = 42
                    _ = v59
                    _ = v60
                }
                _ = v51
                _ = v52
            }
        }
    }
    var v61 uint64 = 394
    fmt.Println(v61)
    switch v5 {
    case "":
        fmt.Println(v7)
    case "hello":
        var v62 string = (("a" + "\n\t") + "\n\t")
        _ = v62
    default:
        v63 := GenS5{F0: v7, F1: 18446744073709551615, F2: -1}
        _ = v63
    }
    var v64 uint64 = v61
    var v65 int = 0
    v66 := GenS4{F0: 100, F1: v65, F2: 9223372036854775807}
    v64 = (v64 ^ v61)
    var v67 int = v65
    switch 42 {
    case -9223372036854775808:
        for fi68 := 0; fi68 < 16; fi68++ {
            var v69 uint16 = (v7 + v7)
            _ = fi68
            _ = v69
        }
    default:
        if true {
            for fi70 := 0; fi70 < 3; fi70++ {
                for fi71 := 0; fi71 < 8; fi71++ {
                    var v72 int = v67
                    v73 := []int{fi71}
                    _ = fi71
                    _ = v72
                    _ = v73
                }
                for fi74 := 0; fi74 < 14; fi74++ {
                    if (193 >= 2147483647) {
                        var v75 uint = 1
                        switch "a" {
                        case "a":
                            v61 -= uint64(v7)
                        case "\n\t":
                            fmt.Println(v7)
                        }
                        _ = v75
                    } else {
                        v76 := GenS2{F0: v7}
                        fmt.Println(v67, v67)
                        _ = v76
                    }
                    _ = fi74
                }
                _ = fi70
            }
            var v77 uint16 = v7
            if (0 > 18446744073709551615) {
                v61 = 42
                v78 := GenS2{F0: (^(v77))}
                v79 := GenS6{F0: 65535, F1: 658}
                _ = v78
                _ = v79
            } else {
                var v80 uint64 = (v64 / ((0) | 1))
                _ = v80
            }
            _ = v77
        } else {
            v3 = v3
        }
    }
    for fi81 := 0; fi81 < 6; fi81++ {
        var v82 uint64 = (v64 - v64)
        for fi83 := 0; fi83 < 19; fi83++ {
            var v84 int8 = -1
            _ = fi83
            _ = v84
        }
        var v85 int8 = -18
        _ = fi81
        _ = v82
        _ = v85
    }
    switch v5 {
    case "":
        if (!(v65 < v65)) {
            var v86 bool = ((v61 & 307) != 18446744073709551615)
            _ = v86
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v7
}
