package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint; F2 int64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 bool }

func main() {
    v1 := []int{-696, -9223372036854775808, 1}
    var v2 float64 = 0.1
    var v3 int8 = -27
    var v4 uint16 = 916
    var v5 float32 = -2.25
    v6 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 4; fi7++ {
        v4 = v4
        v5 = v5
        _ = fi7
    }
    v2 += v2
    _ = len(v6)
    v8 := GenS1{F0: -479}
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    fmt.Println(v2, v4)
    v5 += v5
    v10 := GenS1{F0: 32767}
    switch "\n\t" {
    case "":
        var v11 int = 127
        _ = v11
    default:
        var v12 uint32 = 1
        _ = v12
    }
    v4 = (v4 + v4)
    var v13 int8 = (v3 << ((uint(v4)) & 63))
    if (!true) {
        if true {
            fmt.Println(v13, v3, v3)
            var v14 uint8 = 207
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v14
        } else {
            v13 = 100
            v13 -= v3
            var v15 float32 = 1e10
            _ = v15
        }
    } else {
        v5 += v5
        fmt.Println(v5, v2)
    }
    v2 += v2
    v4 -= v4
    v2 = (0.1 / 0.1)
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int8 = (v13 & (v13 % ((1) | 1)))
            switch (("hello" + "") + "foo bar") {
            case "a":
                var v19 uint64 = 4294967295
                _ = v19
            case "hello":
                var v20 uint32 = 231
                _ = v20
            case "":
                var v21 float32 = v5
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if ((v4 / ((2) | 1)) <= v4) {
                    v6[0] = true
                }
                _ = v21
            default:
                var v22 uint8 = 70
                _ = v22
            }
            _ = len(v6)
            _ = fi17
            _ = v18
        }
        v23 := GenS1{F0: fi16}
        _ = len(v1)
        _ = fi16
        _ = v23
    }
    if false {
        var v24 uint64 = 18446744073709551615
        _ = v24
    } else {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 int16 = 0
            _ = len(v6)
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 11; fi27++ {
            var v28 uint = 0
            v1 = append(v1, 127)
            v6[4294967295] = true
            _ = fi27
            _ = v28
        }
        v29 := GenS1{F0: -2147483648}
        _ = v29
    }
    if (18446744073709551615 >= 4294967295) {
        v13 += v3
        delete(v6, 1)
    } else {
        v30 := GenS1{F0: int(v3)}
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := GenS2{F0: -2147483648, F1: 4294967295}
            for fi33 := 0; fi33 < 2; fi33++ {
                v4 += 800
                var v34 uint8 = 42
                fmt.Println(v4, v3, v3)
                var v35 int = 32768
                var v36 int8 = v3
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            if false {
                if (!true) {
                    var v37 int64 = -242
                    if (42 >= 42) {
                        switch fi31 {
                        case 32767:
                            v6[2147483648] = true
                        default:
                            _ = len(v1)
                        }
                        var v38 int64 = 1
                        _ = v38
                    } else {
                        var v39 int8 = (v13 % ((v3) | 1))
                        fmt.Println(v2, v2)
                        _ = v39
                    }
                    _ = len(v1)
                    _ = v37
                } else {
                    switch (("" + "foo bar") + "") {
                    case "hello":
                        v40 := GenS3{F0: 42, F1: 65535, F2: -430}
                        _ = v40
                    case "foo bar":
                        if true {
                            if (328 == 4294967295) {
                                switch -599 {
                                case 128:
                                    delete(v6, 2147483647)
                                case -625:
                                    v13 = (v3 + v3)
                                case 667:
                                    v41 := GenS4{F0: 142, F1: 127}
                                    _ = v41
                                }
                                if false {
                                    if true {
                                        v42 := GenS3{F0: 722, F1: 549, F2: 0}
                                        switch fi31 {
                                        case 32768:
                                            var v43 uint16 = (v4 << ((255) & 63))
                                            _ = v43
                                        case 127:
                                            for fi44 := 0; fi44 < 15; fi44++ {
                                                if (false || ((((v5 - (0.0 / -2.25)) * ((v5 - (v5 - v5)) * 3.14159)) != v5) || (88 < 1))) {
                                                    fmt.Println(v2, fi44, v4)
                                                    v45 := GenS1{F0: fi31}
                                                    var v46 uint32 = 2147483647
                                                    _ = v45
                                                    _ = v46
                                                }
                                                var v47 string = "foo bar"
                                                _ = fi44
                                                _ = v47
                                            }
                                        default:
                                            for fi48 := 0; fi48 < 17; fi48++ {
                                                v1 = append(v1, fi31)
                                                _ = fi48
                                            }
                                        }
                                        var v49 uint16 = 32767
                                        _ = v42
                                        _ = v49
                                    }
                                    switch fi31 {
                                    case 13:
                                        var v50 uint = 0
                                        _ = v50
                                    case 127:
                                        delete(v6, 2147483648)
                                    default:
                                        var v51 float32 = (float32(1e10) + v5)
                                        _ = v51
                                    }
                                    if true {
                                        var v52 uint8 = 128
                                        v4 += 220
                                        var v53 float32 = 0.0
                                        _ = v52
                                        _ = v53
                                    } else {
                                        var v54 int32 = 2147483647
                                        _ = v54
                                    }
                                } else {
                                    var v55 float64 = 0.1
                                    _ = v55
                                }
                            } else {
                                v13 -= v3
                                for fi56 := 0; fi56 < 10; fi56++ {
                                    _ = len(v6)
                                    v57 := GenS5{F0: (v3 * v3), F1: "\n\t"}
                                    v58 := GenS2{F0: 382, F1: 42}
                                    var v59 int32 = 415
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                var v60 int32 = -703
                                _ = v60
                            }
                            var v61 uint = 18446744073709551615
                            _ = v61
                        } else {
                            var v62 uint = 255
                            _ = v62
                        }
                    case "":
                        if true {
                            var v63 float64 = v2
                            v63 = (((v2 / ((0.0 * 0.1) + (v63 * 3.14159))) - float64(v5)) / 1.0)
                            var v64 bool = false
                            _ = v63
                            _ = v64
                        } else {
                            v3 = -1
                            for fi65 := 0; fi65 < 19; fi65++ {
                                fmt.Println(fi65, v13)
                                var v66 float64 = -2.25
                                if false {
                                    var v67 int = 1
                                    if ((!true) || true) {
                                        v5 -= (((v5 / (v5 + float32(float64(1.5)))) / (v5 * float32(v2))) + 1e10)
                                    } else {
                                        v3 -= v3
                                    }
                                    _ = v67
                                }
                                var v68 int8 = (v13 ^ (v3 << ((4294967295) & 63)))
                                _ = fi65
                                _ = v66
                                _ = v68
                            }
                        }
                    default:
                        var v69 int8 = 89
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 9; fi70++ {
                        fmt.Println(fi31)
                        v71 := GenS6{F0: (9223372036854775807 > 1)}
                        v6[489] = false
                        _ = fi70
                        _ = v71
                    }
                    for fi72 := 0; fi72 < 20; fi72++ {
                        switch ("hello" + "hello") {
                        case "a":
                            v1 = append(v1, fi72)
                        case "\n\t":
                            switch (fi72 + (fi31 | 0)) {
                            case -166:
                                switch (fi31 | (^(fi72))) {
                                case 0:
                                    if (true && ((((v2 - 1e10) + v2) * ((v2 * (((v2 - v2) * 3.14159) / 1.5)) * ((v2 - v2) - float64(float32(float64(1e10)))))) != (float64(v5) * (v2 * v2)))) {
                                        v13 = v13
                                        if (18446744073709551615 > 65535) {
                                            v73 := GenS3{F0: 294, F1: 255, F2: 1}
                                            _ = v73
                                        } else {
                                            v5 = (float32((v2 + v2)) - -2.25)
                                            if true {
                                                v3 = 91
                                            }
                                            if ((float32((v2 / v2)) - (1e10 * 0.1)) <= 1.5) {
                                                v74 := GenS4{F0: 0, F1: 2}
                                                var v75 int8 = v3
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                        v2 = float64((v5 - float32(3.14159)))
                                    }
                                    var v76 uint64 = 0
                                    var v77 int32 = 0
                                    _ = v76
                                    _ = v77
                                case 14:
                                    var v78 uint8 = 0
                                    _ = v78
                                default:
                                    var v79 int8 = v13
                                    _ = v79
                                }
                            case 0:
                                if (!(v3 != -64)) {
                                    var v80 uint32 = 42
                                    _ = len(v1)
                                    _ = v80
                                }
                            default:
                                v4 = (v4 | (v4 + (^(v4))))
                            }
                        default:
                            delete(v6, 1)
                        }
                        _ = fi72
                    }
                }
                v1 = append(v1, 42)
                for fi81 := 0; fi81 < 8; fi81++ {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v13 += (v3 | v3)
                    _ = fi81
                }
            }
            _ = fi31
            _ = v32
        }
        v3 = (v3 - (v13 % ((-128) | 1)))
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
