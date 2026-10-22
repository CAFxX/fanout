package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool; F1 int8 }
type GenS11 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS12 struct { F0 uint; F1 uint16; F2 int }
type GenS13 struct { F0 uint; F1 int8; F2 int16 }
type GenS14 struct { F0 uint64; F1 int16 }
type GenS15 struct { F0 int16 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]bool)
    var v3 int64 = 42
    v4 := make(map[int]bool)
    var v5 uint64 = 1
    v6 := v2[0]
    v7 := GenS1{F0: 1, F1: -1, F2: v6}
    fmt.Println(v6, v1)
    delete(v2, -1)
    var v8 int8 = 42
    switch 128 {
    case 705:
        var v9 bool = v6
        _ = v9
    case 42:
        switch "foo bar" {
        case "":
            v3 += v3
        case "foo bar":
            var v10 uint8 = 2
            _ = v10
        case "hello":
            v11 := GenS1{F0: 127, F1: 132, F2: false}
            _ = v11
        }
    default:
        v12 := GenS1{F0: -2147483648, F1: 1, F2: (-2147483648 < 42)}
        _ = v12
    }
    if ((v6 && v6) && true) {
        v1 = v1
        v13 := GenS2{F0: 70}
        var v14 int32 = 0
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 180
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if (4294967295 == v5) {
            v17 := []int{(fi16 / (((fi16 + fi16)) | 1)), (fi16 * (^(fi16))), (-(fi16))}
            if len(v17) > 0 {
                v17[0] = (fi16 & (-(fi16)))
            }
            if (v6 || v6) {
                v18 := GenS3{F0: 42}
                fmt.Println(v3, fi16, v1)
                switch (fi16 + -2147483648) {
                case 32768:
                    v5 = (v5 % ((4294967295) | 1))
                case 1:
                    for fi19 := 0; fi19 < 15; fi19++ {
                        fmt.Println(fi19)
                        _ = fi19
                    }
                case -9223372036854775808:
                    switch "\n\t" {
                    case "\n\t":
                        var v20 uint16 = v1
                        _ = v20
                    case "hello":
                        for fi21 := 0; fi21 < 10; fi21++ {
                            if (fi21 != 2147483647) {
                                v4[fi16] = v6
                                var v22 int64 = (v3 + (-(v3)))
                                _ = v22
                            } else {
                                v6 = v6
                                delete(v4, fi16)
                                if len(v17) > 0 {
                                    v17[0] = (fi16 << ((255) & 63))
                                }
                            }
                            _ = fi21
                        }
                    case "":
                        switch "foo bar" {
                        case "":
                            v8 = v8
                        case "foo bar":
                            var v23 string = ("hello" + "")
                            _ = v23
                        case "a":
                            if (true || (!(v6 && (!false)))) {
                                _ = len(v17)
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    v25 := GenS4{F0: "a", F1: 963}
                                    _ = fi24
                                    _ = v25
                                }
                            }
                        default:
                            var v26 float32 = 1.5
                            _ = v26
                        }
                    default:
                        switch "foo bar" {
                        case "":
                            if len(v17) > 0 {
                                v17[0] = (fi16 & (fi16 * 2147483647))
                            }
                        default:
                            var v27 string = "hello"
                            _ = v27
                        }
                    }
                }
                _ = v18
            } else {
                var v28 uint16 = (v1 << ((591) & 63))
                _ = v28
            }
            _ = v17
        } else {
            v29 := []int{}
            _ = v29
        }
        if (v5 <= (v5 % (((^(v5))) | 1))) {
            var v30 float32 = (1.5 / (float32(-2.25) - ((float32(3.14159) * 3.14159) - 1.5)))
            _ = v30
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            v6 = (613 == 32767)
            _ = fi31
        }
        _ = fi16
    }
    var v32 uint16 = v1
    for fi33 := 0; fi33 < 11; fi33++ {
        if (!false) {
            v8 += -128
            var v34 float64 = 1e10
            var v35 int64 = -1
            _ = v34
            _ = v35
        } else {
            var v36 float64 = (float64(-2.25) * 1e10)
            switch "a" {
            case "a":
                v5 -= v5
            case "":
                switch "" {
                case "a":
                    var v37 uint32 = 604
                    _ = v37
                case "hello":
                    for fi38 := 0; fi38 < 13; fi38++ {
                        var v39 float64 = v36
                        fmt.Println(v6, v1, fi33)
                        if false {
                            if (409 < 2) {
                                v40 := GenS4{F0: ((("a" + "foo bar") + "") + "hello"), F1: 2147483647}
                                v41 := GenS1{F0: (fi33 + 1), F1: 716, F2: true}
                                _ = v40
                                _ = v41
                            } else {
                                switch "hello" {
                                case "\n\t":
                                    v5 = 4294967295
                                case "a":
                                    v42 := GenS5{F0: (v5 % (((^(v5))) | 1)), F1: (v5 ^ (v5 | 1)), F2: 1}
                                    _ = v42
                                default:
                                    v8 -= v8
                                }
                                _ = len(v2)
                                v1 += 241
                            }
                            v43 := v4[fi33]
                            v44 := GenS6{F0: ("hello" > "a"), F1: 2}
                            _ = v43
                            _ = v44
                        } else {
                            v5 = 889
                            var v45 uint64 = v5
                            fmt.Println(v45)
                            _ = v45
                        }
                        _ = fi38
                        _ = v39
                    }
                case "":
                    for fi46 := 0; fi46 < 17; fi46++ {
                        var v47 int64 = v3
                        _ = fi46
                        _ = v47
                    }
                default:
                    switch fi33 {
                    case 1:
                        if (false || (true || true)) {
                            switch 2147483647 {
                            case -634:
                                switch -2147483648 {
                                case 9223372036854775807:
                                    var v48 uint16 = 0
                                    _ = v48
                                default:
                                    var v49 int64 = 681
                                    _ = v49
                                }
                            case 127:
                                var v50 uint8 = 128
                                _ = v50
                            }
                        }
                    case 213:
                        for fi51 := 0; fi51 < 20; fi51++ {
                            v52 := []int{(fi33 ^ fi33)}
                            var v53 int16 = -383
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    }
                }
            default:
                var v54 int64 = 9223372036854775807
                _ = v54
            }
            _ = v36
        }
        if false {
            var v55 int32 = -697
            if (!(!(96 != 42))) {
                v56 := GenS6{F0: v6, F1: 65535}
                fmt.Println(v1, v3)
                _ = v56
            } else {
                var v57 string = ("hello" + "")
                for fi58 := 0; fi58 < 13; fi58++ {
                    switch "\n\t" {
                    case "a":
                        var v59 uint64 = 630
                        _ = v59
                    case "\n\t":
                        var v60 int64 = v3
                        _ = v60
                    case "":
                        if v6 {
                            v4[32767] = (v55 != (v55 | v55))
                            switch -9223372036854775808 {
                            case 39:
                                fmt.Println(fi58, v8)
                            case -9223372036854775808:
                                var v61 uint64 = (v5 / (((v5 | (v5 >> ((0) & 63)))) | 1))
                                _ = v61
                            case 128:
                                v62 := GenS7{F0: -1, F1: 0}
                                _ = v62
                            default:
                                v63 := []int{fi33, (fi33 ^ -375)}
                                _ = v63
                            }
                            var v64 int16 = 37
                            _ = v64
                        }
                    default:
                        if false {
                            var v65 float64 = 0.1
                            if (!(v65 < (v65 - (v65 * (v65 * (3.14159 / v65)))))) {
                                v6 = v6
                                fmt.Println(fi33)
                                var v66 uint64 = v5
                                v67 := GenS2{F0: 0}
                                _ = v66
                                _ = v67
                            } else {
                                switch v57 {
                                case "hello":
                                    fmt.Println(fi33)
                                default:
                                    v32 -= v32
                                }
                            }
                            v68 := GenS5{F0: (v5 & v5), F1: v5, F2: 1}
                            _ = v65
                            _ = v68
                        } else {
                            if (-2.25 == 3.14159) {
                                var v69 int8 = (v8 & v8)
                                _ = v69
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    v6 = false
                                case "a":
                                    var v70 uint16 = v1
                                    _ = v70
                                case "\n\t":
                                    switch fi58 {
                                    case 460:
                                        v57 = (v57 + (v57 + v57))
                                    case -425:
                                        if ((0.0 / 1.0) >= -2.25) {
                                            switch "" {
                                            case "":
                                                switch (^(fi58)) {
                                                case 32768:
                                                    v6 = (((0.0 * (float32(1e10) * (float32(-2.25) * -2.25))) / ((1e10 - -2.25) / ((float32(0.1) + (0.1 * (0.1 + 1.5))) - (((1e10 - (1.5 + (0.0 / float32(-2.25)))) * (((3.14159 * (1e10 + 0.0)) * -2.25) * 0.0)) - 0.0)))) >= 1.5)
                                                case -591:
                                                    v57 = (((v57 + v57) + v57) + (((((v57 + v57) + v57) + v57) + "a") + v57))
                                                case -342:
                                                    fmt.Println(fi33, fi33)
                                                }
                                                if ((float32(3.14159) - 0.0) < float32((1.5 + float64(1e10)))) {
                                                    switch (v57 + "\n\t") {
                                                    case "":
                                                        v71 := v2[-404]
                                                        _ = v71
                                                    case "a":
                                                        v72 := GenS8{F0: 183, F1: v57}
                                                        v73 := GenS9{F0: -112}
                                                        if v6 {
                                                            v74 := GenS10{F0: false, F1: v8}
                                                            if v6 {
                                                                v6 = true
                                                            }
                                                            _ = v74
                                                        }
                                                        _ = v72
                                                        _ = v73
                                                    case "\n\t":
                                                        v5 = v5
                                                    }
                                                } else {
                                                    v75 := GenS11{F0: 2, F1: -1, F2: 127}
                                                    if v6 {
                                                        v3 -= -554
                                                        var v76 uint64 = v5
                                                        _ = v76
                                                    } else {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v77 := GenS3{F0: uint8(v5)}
                                                            var v78 int32 = v55
                                                            v79 := GenS12{F0: 65535, F1: (v32 % (((v32 & 255)) | 1)), F2: fi33}
                                                            v57 = (("hello" + v57) + "a")
                                                            _ = v77
                                                            _ = v78
                                                            _ = v79
                                                        }
                                                        v6 = (v3 >= (v3 | int64(v55)))
                                                    }
                                                    var v80 int32 = v55
                                                    _ = v75
                                                    _ = v80
                                                }
                                            }
                                            var v81 uint32 = 2147483648
                                            _ = v81
                                        }
                                    case 42:
                                        v8 += (v8 & v8)
                                    }
                                default:
                                    var v82 uint32 = 42
                                    _ = v82
                                }
                                var v83 int = (fi33 + 42)
                                var v84 uint8 = 2
                                _ = v83
                                _ = v84
                            }
                            switch v57 {
                            case "a":
                                delete(v2, (fi58 / ((int(v55)) | 1)))
                            default:
                                var v85 uint64 = v5
                                _ = v85
                            }
                            v5 += (v5 ^ (v5 + v5))
                        }
                    }
                    v86 := GenS13{F0: 42, F1: v8, F2: 32767}
                    _ = fi58
                    _ = v86
                }
                var v87 int64 = v3
                _ = v57
                _ = v87
            }
            if false {
                fmt.Println(fi33, v3)
                v88 := GenS14{F0: 18446744073709551615, F1: 1}
                v55 = v55
                _ = v88
            }
            _ = v55
        }
        _ = fi33
    }
    var v89 string = ""
    v90 := GenS6{F0: (v6 || v6), F1: v1}
    if (32767 <= v32) {
        if v6 {
            for fi91 := 0; fi91 < 2; fi91++ {
                v8 += 85
                v92 := GenS6{F0: false, F1: 255}
                fmt.Println(v8)
                _ = fi91
                _ = v92
            }
            var v93 uint = 676
            _ = v93
        } else {
            var v94 float64 = 1.5
            v4[9223372036854775807] = (v5 == (v5 / ((986) | 1)))
            _ = v94
        }
        for fi95 := 0; fi95 < 14; fi95++ {
            v96 := GenS15{F0: 356}
            _ = fi95
            _ = v96
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v89
    _ = v90
}
