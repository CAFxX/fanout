package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int; F1 uint64; F2 int }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint32 }
func genFn0() bool {
    v6 := GenS1{F0: 255}
    var v7 uint8 = 128
    if false {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := []int{(fi8 % ((fi8) | 1)), (fi8 - fi8), 9223372036854775807}
            _ = fi8
            _ = v9
        }
        var v10 int64 = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if (-128 != -1) {
            var v12 uint64 = 0
            var v13 int8 = 127
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    v14 := []int{9223372036854775807, -853, 68}
    _ = v14
    _ = v6
    _ = v7
    return ((9223372036854775807 <= 1) || false)
}

func main() {
    var v1 uint64 = 841
    var v2 uint32 = 2147483648
    var v3 uint = 42
    var v4 uint16 = 531
    var v5 float64 = (3.14159 + 0.0)
    v15 := []int{int(v1), 127, 42}
    var v16 uint8 = 42
    v2 -= v2
    v2 = (v2 - v2)
    switch "" {
    case "":
        var v17 uint16 = (v4 | v4)
        _ = v17
    default:
        if (!false) {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        } else {
            if len(v15) > 0 {
                v15[0] = -1
            }
        }
    }
    var v18 uint8 = 1
    if false {
        var v19 string = "\n\t"
        v20 := GenS1{F0: (^(v3))}
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int16 = 32767
            for fi23 := 0; fi23 < 14; fi23++ {
                switch v19 {
                case "\n\t":
                    if false {
                        var v24 int32 = 132
                        var v25 int = 42
                        var v26 int = (v25 & 32768)
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                case "a":
                    if len(v15) > 0 {
                        v15[0] = -259
                    }
                default:
                    var v27 uint = v3
                    _ = v27
                }
                var v28 float64 = 0.1
                switch (v19 + "foo bar") {
                case "":
                    v16 -= uint8(v3)
                case "foo bar":
                    var v29 bool = false
                    _ = v29
                }
                _ = fi23
                _ = v28
            }
            v30 := GenS1{F0: (v3 - v3)}
            _ = fi21
            _ = v22
            _ = v30
        }
        _ = v19
        _ = v20
    }
    if true {
        if true {
            v4 = (v4 << ((v3) & 63))
        } else {
            fmt.Println(v18, v16, v4)
            var v31 int32 = 740
            _ = v31
        }
    }
    v2 -= (v2 | 65535)
    var v32 int64 = 868
    fmt.Println(v4, v1)
    for fi33 := 0; fi33 < 3; fi33++ {
        v34 := GenS2{F0: (-(fi33)), F1: 4294967295}
        for fi35 := 0; fi35 < 5; fi35++ {
            if true {
                if len(v15) > 0 {
                    v15[0] = 1
                }
                var v36 int64 = (v32 + int64(fi33))
                var v37 uint64 = 2
                v38 := GenS1{F0: v3}
                v36 = (v32 - int64(fi33))
                _ = v36
                _ = v37
                _ = v38
            } else {
                v16 -= v16
                var v39 float64 = v5
                var v40 float64 = -2.25
                _ = v39
                _ = v40
            }
            _ = fi35
        }
        for fi41 := 0; fi41 < 15; fi41++ {
            if ((v2 + (v2 + v2)) >= uint32(v3)) {
                v42 := GenS2{F0: fi33, F1: fi41}
                _ = len(v15)
                _ = v42
            } else {
                v1 -= (^(v1))
            }
            switch "" {
            case "hello":
                v15 = append(v15, fi41)
            case "\n\t":
                v4 += (v4 - 65535)
                if (!false) {
                    v4 -= 32767
                    if (-2.25 < float32(float64((-2.25 + 0.0)))) {
                        switch fi33 {
                        case -373:
                            v18 = (v16 + 2)
                        }
                    }
                    var v43 uint8 = v16
                    v44 := GenS3{F0: 565, F1: 2, F2: (^(v3))}
                    _ = v43
                    _ = v44
                }
                v45 := GenS1{F0: v3}
                if true {
                    v4 -= v4
                    if (v18 >= uint8(v4)) {
                        if (-104 != 0) {
                            var v46 int8 = 0
                            v32 = int64(v4)
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            _ = v46
                        }
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        if (("foo bar" > "hello") || ((fi41 - 32767) >= fi41)) {
                            v47 := GenS1{F0: v3}
                            var v48 uint8 = (^(v18))
                            v15 = append(v15, int(v4))
                            _ = v47
                            _ = v48
                        } else {
                            v49 := GenS2{F0: (fi33 | fi33), F1: (fi33 | fi41)}
                            v32 -= (v32 << ((v3) & 63))
                            v5 -= ((float64(float32(v5)) - float64(0.1)) / float64(float32((float64(1e10) - (0.0 / v5)))))
                            _ = v49
                        }
                    } else {
                        switch "hello" {
                        case "\n\t":
                            v50 := GenS2{F0: 32768, F1: fi33}
                            v51 := GenS1{F0: v3}
                            fmt.Println(v16, v1)
                            fmt.Println(v5, v5)
                            var v52 uint16 = 32767
                            var v53 uint = (v3 / (((v3 >> ((v3) & 63))) | 1))
                            switch "\n\t" {
                            case "foo bar":
                                if true {
                                    if false {
                                        v3 = (v3 | (v3 + (v53 + (v53 * v53))))
                                    }
                                    if (-128 == 0) {
                                        v54 := GenS4{F0: (v4 | 2), F1: (v1 / (((^(v1))) | 1)), F2: 18446744073709551615}
                                        _ = v54
                                    } else {
                                        var v55 int = (fi33 ^ -2147483648)
                                        if true {
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                        } else {
                                            var v56 uint = v3
                                            switch fi41 {
                                            case 127:
                                                _ = len(v15)
                                            }
                                            fmt.Println(v55)
                                            _ = v56
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            var v57 int16 = -1
                                            _ = v57
                                        case "":
                                            if len(v15) > 0 {
                                                v15[0] = v55
                                            }
                                        }
                                        _ = v55
                                    }
                                    var v58 float64 = v5
                                    _ = v58
                                }
                            case "":
                                v59 := GenS5{F0: (fi33 - fi33), F1: 862, F2: int(v32)}
                                _ = v59
                            case "hello":
                                switch "foo bar" {
                                case "foo bar":
                                    if true {
                                        var v60 int = (^(fi33))
                                        v60 += 356
                                        var v61 float32 = 1.5
                                        fmt.Println(v53, v18, fi41)
                                        _ = v60
                                        _ = v61
                                    } else {
                                        if ((((-2.25 / (v5 + v5)) - (v5 / (v5 + 1.5))) > v5) || (v52 == v52)) {
                                            if false {
                                                if (0.1 > (1e10 * 0.1)) {
                                                    if true {
                                                        _ = len(v15)
                                                        var v62 uint16 = v4
                                                        v63 := GenS5{F0: fi41, F1: v1, F2: fi41}
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        var v64 int8 = 100
                                                        var v65 int16 = 510
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    if (true || true) {
                                                        var v66 string = "hello"
                                                        var v67 uint32 = 478
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                }
                                                if (v52 < 42) {
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "hello":
                                                    _ = len(v15)
                                                case "foo bar":
                                                    v68 := GenS6{F0: (^(v53)), F1: ("\n\t" + "\n\t")}
                                                    _ = v68
                                                case "":
                                                    if len(v15) > 0 {
                                                        v15[0] = int(v32)
                                                    }
                                                }
                                            } else {
                                                v18 = (v18 ^ 42)
                                                v52 = (^(v52))
                                                if true {
                                                    var v69 uint64 = v1
                                                    var v70 uint8 = v18
                                                    _ = v69
                                                    _ = v70
                                                }
                                            }
                                        }
                                    }
                                case "a":
                                    _ = len(v15)
                                case "":
                                    var v71 uint64 = v1
                                    _ = v71
                                default:
                                    v18 = (^(v18))
                                }
                            default:
                                var v72 uint32 = v2
                                _ = v72
                            }
                            var v73 int16 = 0
                            switch "a" {
                            case "foo bar":
                                switch "a" {
                                case "hello":
                                    v2 = uint32(v16)
                                case "a":
                                    var v74 int16 = (v73 + -113)
                                    _ = v74
                                default:
                                    if len(v15) > 0 {
                                        v15[0] = (fi41 * 127)
                                    }
                                }
                            case "hello":
                                v4 = 972
                            default:
                                switch ((("a" + "\n\t") + ("\n\t" + "\n\t")) + "a") {
                                case "foo bar":
                                    switch "a" {
                                    case "":
                                        fmt.Println(v2, fi33)
                                    }
                                    var v75 float32 = 1.5
                                    _ = v75
                                case "":
                                    if false {
                                        v2 -= (^(v2))
                                    }
                                case "\n\t":
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                }
                            }
                            _ = v50
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v73
                        case "hello":
                            if len(v15) > 0 {
                                v15[0] = fi41
                            }
                        }
                        v3 -= uint(v4)
                    }
                }
                _ = v45
            default:
                if (((-2.25 + (v5 + v5)) + v5) < 1.5) {
                    if (!false) {
                        v76 := GenS6{F0: uint(v1), F1: "\n\t"}
                        v5 -= float64(((((1.5 + 1.5) + (0.1 / 1.0)) - 0.0) - ((1.5 * 0.0) - (-2.25 / (3.14159 - 1.5)))))
                        v1 += (v1 % ((v1) | 1))
                        switch fi41 {
                        case 366:
                            if true {
                                v1 -= (v1 | (v1 | v1))
                                if true {
                                    if false {
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                        var v77 float64 = v5
                                        _ = v77
                                    } else {
                                        switch 1 {
                                        case 9223372036854775807:
                                            var v78 float32 = ((1.5 + 0.0) + ((-2.25 * -2.25) + 3.14159))
                                            _ = v78
                                        case 321:
                                            if false {
                                                v79 := GenS1{F0: v3}
                                                _ = v79
                                            } else {
                                                var v80 uint32 = uint32(v18)
                                                if false {
                                                    v81 := GenS1{F0: v3}
                                                    fmt.Println(v80, v3)
                                                    v3 = v3
                                                    var v82 uint32 = 2147483647
                                                    if false {
                                                        var v83 uint = (v3 % (((v3 | v3)) | 1))
                                                        _ = v83
                                                    } else {
                                                        var v84 uint = (v3 % (((v3 % ((v3) | 1))) | 1))
                                                        var v85 string = "a"
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    _ = v81
                                                    _ = v82
                                                }
                                                _ = len(v15)
                                                var v86 int16 = 32767
                                                _ = v80
                                                _ = v86
                                            }
                                        case -391:
                                            fmt.Println(fi33)
                                        }
                                    }
                                    v18 += (v18 * v16)
                                } else {
                                    v5 = (v5 / v5)
                                    v2 = (v2 | (^(v2)))
                                    var v87 int8 = -128
                                    var v88 int64 = (v32 >> ((v3) & 63))
                                    var v89 int = (-(fi41))
                                    fmt.Println(v87, v16, v87)
                                    _ = v87
                                    _ = v88
                                    _ = v89
                                }
                                if true {
                                    fmt.Println(v1, v1)
                                    v3 -= (^(v3))
                                    var v90 uint32 = (v2 % ((65535) | 1))
                                    var v91 int64 = -475
                                    v92 := GenS7{F0: 1, F1: 0}
                                    switch fi41 {
                                    case 32767:
                                        var v93 uint16 = v4
                                        _ = v93
                                    case -1:
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    case 4294967295:
                                        v4 -= (v4 - (v4 * (v4 | v4)))
                                    default:
                                        v90 += v90
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch (fi33 % (((^(fi33))) | 1)) {
                                        case 9223372036854775807:
                                            v94 := GenS5{F0: 127, F1: v1, F2: fi41}
                                            _ = v94
                                        case 694:
                                            v1 -= (^(v1))
                                        case -2147483648:
                                            var v95 float64 = v5
                                            _ = v95
                                        default:
                                            if false {
                                                v96 := GenS3{F0: v1, F1: v16, F2: v3}
                                                _ = v96
                                            } else {
                                                var v97 uint8 = 2
                                                var v98 uint32 = 350
                                                _ = v97
                                                _ = v98
                                            }
                                            v99 := GenS8{F0: 32767}
                                            _ = v99
                                        }
                                    case "a":
                                        var v100 uint8 = (v16 ^ v18)
                                        _ = v100
                                    default:
                                        var v101 string = (("a" + "") + "")
                                        _ = v101
                                    }
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    v102 := GenS5{F0: (fi33 >> ((v3) & 63)), F1: v1, F2: fi41}
                                    _ = len(v15)
                                    v103 := GenS3{F0: uint64(v90), F1: (^(v18)), F2: v3}
                                    _ = v102
                                    _ = v103
                                    _ = v90
                                    _ = v91
                                    _ = v92
                                }
                            } else {
                                var v104 string = "foo bar"
                                var v105 int = fi33
                                _ = v104
                                _ = v105
                            }
                        case 4294967295:
                            var v106 uint = uint(v2)
                            _ = v106
                        default:
                            if (int16(v18) >= 42) {
                                if len(v15) > 0 {
                                    v15[0] = (-(fi33))
                                }
                                var v107 int32 = 2147483647
                                _ = v107
                            } else {
                                if len(v15) > 0 {
                                    v15[0] = (fi33 / ((fi41) | 1))
                                }
                                var v108 uint = 255
                                _ = len(v15)
                                _ = v108
                            }
                        }
                        _ = v76
                    }
                }
            }
            v109 := GenS9{F0: v2}
            _ = fi41
            _ = v109
        }
        _ = fi33
        _ = v34
    }
    v16 += v16
    v3 = v3
    var v110 int64 = (v32 ^ -378)
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v110
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
