package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 bool; F2 int }
type GenS3 struct { F0 int8; F1 int32; F2 int32 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 int16; F2 int32 }

func main() {
    var v1 int = 4294967295
    v2 := make(map[int]string)
    var v3 uint64 = 707
    var v4 bool = false
    v5 := make(map[int]uint32)
    v4 = (-1 >= 1)
    v6 := v2[(v1 & 42)]
    v7 := GenS1{F0: v1}
    for fi8 := 0; fi8 < 9; fi8++ {
        v1 -= -1
        var v9 uint16 = 2
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        if (true || (true || (true || (!v4)))) {
            if true {
                if (42 > 65535) {
                    switch v6 {
                    case "a":
                        if true {
                            v6 = v6
                            switch (v1 / ((v1) | 1)) {
                            case 4294967295:
                                var v11 uint16 = 633
                                _ = v11
                            }
                            var v12 int64 = 42
                            _ = v12
                        }
                    case "":
                        _ = len(v2)
                    }
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v4 = (((!(42 > -1)) || v4) && (!((127 < -57) && v4)))
                        switch (-(fi10)) {
                        case -1:
                            if (!true) {
                                v3 -= (v3 ^ 42)
                            }
                        case -9223372036854775808:
                            var v14 float32 = 1e10
                            _ = v14
                        case 752:
                            v15 := GenS1{F0: fi10}
                            _ = v15
                        default:
                            var v16 float32 = -2.25
                            v16 = float32(0.1)
                            v1 = -952
                            _ = v16
                        }
                        switch fi13 {
                        case -70:
                            v17 := GenS1{F0: (fi13 / ((4294967295) | 1))}
                            _ = v17
                        case -2147483648:
                            var v18 float32 = 0.0
                            _ = v18
                        default:
                            if (!(!v4)) {
                                var v19 int32 = 583
                                var v20 uint16 = 1
                                v2[(v1 & (^(fi10)))] = v6
                                _ = v19
                                _ = v20
                            }
                        }
                        _ = fi13
                    }
                    v21 := GenS1{F0: (v1 / ((fi10) | 1))}
                    _ = v21
                }
                if v4 {
                    switch (v6 + "") {
                    case "hello":
                        var v22 int64 = 0
                        _ = v22
                    case "a":
                        var v23 float64 = (0.0 + -2.25)
                        _ = v23
                    }
                    fmt.Println(v6, v1)
                    v6 = v6
                }
            } else {
                var v24 uint = 1
                fmt.Println(v24, v24)
                _ = v24
            }
        } else {
            v1 -= v1
            if (!v4) {
                delete(v2, (fi10 * int(v3)))
            }
            var v25 int64 = 0
            _ = v25
        }
        _ = fi10
    }
    var v26 uint = 65535
    var v27 int64 = int64(v3)
    var v28 int64 = v27
    if (!v4) {
        if true {
            for fi29 := 0; fi29 < 5; fi29++ {
                v27 -= (v28 - (v28 >> (((v26 & 993)) & 63)))
                var v30 uint16 = 392
                _ = fi29
                _ = v30
            }
            v27 = (v27 * (v27 | v27))
        }
        if (v3 != (v3 ^ 418)) {
            if true {
                v31 := []int{}
                v32 := GenS1{F0: (v1 % (((v1 + int(v26))) | 1))}
                _ = v31
                _ = v32
            }
        }
    } else {
        if (((1000 <= 1000) && v4) || (87 <= 0)) {
            for fi33 := 0; fi33 < 1; fi33++ {
                switch (v1 % (((^(fi33))) | 1)) {
                case 2147483647:
                    if (164 != 74) {
                        v34 := []int{v1, 579, fi33, fi33}
                        _ = v34
                    }
                case -1:
                    v35 := []int{fi33}
                    _ = v35
                default:
                    var v36 uint32 = 682
                    _ = v36
                }
                var v37 int8 = 21
                var v38 uint32 = 2
                _ = fi33
                _ = v37
                _ = v38
            }
            var v39 uint8 = 42
            _ = v39
        }
    }
    if (v4 && v4) {
        if (!(3.14159 > 0.0)) {
            var v40 uint16 = uint16(v26)
            _ = v40
        }
        var v41 float64 = -2.25
        v41 -= (float64(float32(float64(0.1))) - v41)
        _ = v41
    } else {
        v26 += (^(v26))
        var v42 int8 = 1
        var v43 string = (("\n\t" + v6) + ("foo bar" + ((v6 + "a") + (v6 + v6))))
        _ = v42
        _ = v43
    }
    switch v6 {
    case "foo bar":
        for fi44 := 0; fi44 < 8; fi44++ {
            v28 += (v27 >> ((v26) & 63))
            var v45 string = ("\n\t" + v6)
            v46 := []int{-9223372036854775808, fi44, int(v27), v1}
            _ = fi44
            _ = v45
            _ = v46
        }
    default:
        for fi47 := 0; fi47 < 6; fi47++ {
            for fi48 := 0; fi48 < 7; fi48++ {
                fmt.Println(v26, v3)
                _ = fi48
            }
            _ = fi47
        }
    }
    v49 := []int{(^(v1)), v1}
    delete(v5, 127)
    v50 := GenS1{F0: 42}
    var v51 uint8 = 245
    var v52 bool = v4
    switch v1 {
    case 32767:
        for fi53 := 0; fi53 < 11; fi53++ {
            v54 := GenS2{F0: fi53, F1: (v51 <= (^(v51))), F2: (-(fi53))}
            if (!(65535 == 4294967295)) {
                v52 = false
                _ = len(v49)
                for fi55 := 0; fi55 < 9; fi55++ {
                    var v56 uint8 = 0
                    if (v28 >= 1) {
                        if v4 {
                            if v52 {
                                v57 := GenS3{F0: 1, F1: 0, F2: -2147483648}
                                v58 := GenS3{F0: 100, F1: 0, F2: 396}
                                var v59 uint16 = 255
                                _ = v57
                                _ = v58
                                _ = v59
                            } else {
                                if (!false) {
                                    if len(v49) >= 2 {
                                        v49 = v49[1:len(v49)]
                                    }
                                    v6 = v6
                                    _ = len(v49)
                                    if (v52 && v52) {
                                        switch v6 {
                                        case "a":
                                            fmt.Println(v28, v1)
                                        case "":
                                            if (2 <= 255) {
                                                var v60 int8 = 42
                                                _ = v60
                                            } else {
                                                _ = len(v49)
                                            }
                                        case "hello":
                                            v51 -= 27
                                        }
                                        v61 := GenS1{F0: fi55}
                                        var v62 int = fi53
                                        switch v1 {
                                        case 127:
                                            if (false || (v62 != v1)) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if v4 {
                                                        fmt.Println(v56)
                                                        var v63 bool = false
                                                        _ = v63
                                                    } else {
                                                        if ((^(v3)) >= v3) {
                                                            if len(v49) >= 2 {
                                                                v49 = v49[1:len(v49)]
                                                            }
                                                            v52 = (!v52)
                                                            switch 898 {
                                                            case 42:
                                                                switch (fi55 >> ((v26) & 63)) {
                                                                case 2147483647:
                                                                    var v64 uint16 = 0
                                                                    _ = v64
                                                                case -305:
                                                                    v1 = fi53
                                                                case 32768:
                                                                    v62 = (fi53 ^ 134)
                                                                }
                                                            }
                                                        } else {
                                                            var v65 int32 = 2147483647
                                                            _ = v65
                                                        }
                                                    }
                                                }
                                                switch v6 {
                                                case "hello":
                                                    v66 := GenS4{F0: 2147483647, F1: v4}
                                                    _ = v66
                                                case "":
                                                    delete(v2, v62)
                                                }
                                            } else {
                                                if ("foo bar" == ("\n\t" + ("\n\t" + v6))) {
                                                    v49 = append(v49, fi55)
                                                    fmt.Println(v51, v56, v62)
                                                } else {
                                                    if (2147483647 != 2) {
                                                        v27 -= (-(v27))
                                                        if true {
                                                            switch v6 {
                                                            case "\n\t":
                                                                var v67 int16 = -1
                                                                _ = v67
                                                            case "hello":
                                                                fmt.Println(v3)
                                                            case "":
                                                                if (false && v4) {
                                                                    delete(v2, 1)
                                                                    v28 -= (v28 ^ 9223372036854775807)
                                                                    if ("hello" == "foo bar") {
                                                                        var v68 int8 = -1
                                                                        v69 := GenS5{F0: 2}
                                                                        v68 += 100
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                }
                                                            }
                                                            var v70 int32 = 0
                                                            fmt.Println(fi55, v26)
                                                            _ = v70
                                                        }
                                                        delete(v5, 814)
                                                    } else {
                                                        fmt.Println(fi55, v56, v4)
                                                        v27 += (-(v27))
                                                        if v4 {
                                                            fmt.Println(v27, v6)
                                                            delete(v5, v1)
                                                            v4 = ((v52 || v52) && v4)
                                                            switch ("hello" + v6) {
                                                            case "hello":
                                                                v3 = uint64(v26)
                                                            case "":
                                                                fmt.Println(fi53, v28)
                                                            default:
                                                                fmt.Println(v4, v51, fi55)
                                                            }
                                                        } else {
                                                            if false {
                                                                v51 -= uint8(v26)
                                                            } else {
                                                                var v71 string = ""
                                                                v71 = ""
                                                                _ = v71
                                                            }
                                                            var v72 int64 = v28
                                                            v26 = (^(v26))
                                                            v72 = -805
                                                            if v4 {
                                                                if (v4 || ((v1 ^ fi55) != fi53)) {
                                                                    fmt.Println(v3, v3)
                                                                    v27 = (v28 + v72)
                                                                    v3 = 1
                                                                    v73 := GenS6{F0: -2147483648, F1: 32767, F2: 660}
                                                                    _ = v73
                                                                } else {
                                                                    fmt.Println(v1, v51, v26)
                                                                }
                                                                var v74 uint8 = v56
                                                                _ = v74
                                                            }
                                                            _ = v72
                                                        }
                                                    }
                                                }
                                            }
                                            v51 += (v51 - 255)
                                            v56 -= (v56 & (v56 >> ((2) & 63)))
                                        case 9223372036854775807:
                                            fmt.Println(v26, fi53, v51)
                                        case -585:
                                            var v75 int8 = 127
                                            var v76 float64 = 0.0
                                            var v77 float64 = 3.14159
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        default:
                                            fmt.Println(v26, v1)
                                        }
                                        var v78 int = (fi53 >> (((v26 >> ((v26) & 63))) & 63))
                                        if v52 {
                                            var v79 float32 = float32((float64((3.14159 + -2.25)) * 3.14159))
                                            _ = v79
                                        }
                                        var v80 uint32 = uint32(v26)
                                        _ = v61
                                        _ = v62
                                        _ = v78
                                        _ = v80
                                    }
                                    v26 -= 1
                                }
                                v26 = (v26 & v26)
                                if v52 {
                                    v1 = (^(fi53))
                                    var v81 uint16 = 269
                                    fmt.Println(v6, v81)
                                    _ = v81
                                }
                            }
                        }
                    }
                    v56 -= v56
                    _ = fi55
                    _ = v56
                }
            }
            _ = fi53
            _ = v54
        }
    case 2147483647:
        v82 := GenS2{F0: v1, F1: v4, F2: (v1 % (((v1 / (((v1 % ((v1) | 1))) | 1))) | 1))}
        _ = v82
    case 546:
        var v83 uint32 = uint32(v51)
        _ = v83
    default:
        if ((!v4) || v52) {
            if (127 < -1) {
                var v84 int8 = 1
                if (false && false) {
                    v84 += (-(v84))
                    var v85 int8 = v84
                    var v86 uint32 = 649
                    _ = v85
                    _ = v86
                }
                _ = v84
            }
            fmt.Println(v52, v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v7
}
