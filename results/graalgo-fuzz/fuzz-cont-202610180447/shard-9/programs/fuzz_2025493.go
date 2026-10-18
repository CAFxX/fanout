package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int8, p1 string) string {
    var v4 uint = 38
    fmt.Println(v4, p1)
    v4 -= 42
    if (!(!(1 >= 42))) {
        var v5 int32 = 647
        fmt.Println(p1)
        _ = v5
    } else {
        v6 := GenS1{F0: 0, F1: 2, F2: 2147483647}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return p1
}
func genFn1(p0 uint32, p1 int8) uint64 {
    v7 := GenS1{F0: 2, F1: (p0 | (p0 % (((p0 % ((p0) | 1))) | 1))), F2: (p0 % (((^(p0))) | 1))}
    var v8 float64 = 0.0
    v9 := GenS1{F0: 292, F1: (p0 % ((p0) | 1)), F2: 2147483648}
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 float32 = -2.25
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 458
}

func main() {
    v1 := []int{-583, 32767, 914}
    var v2 bool = true
    var v3 int = 32768
    for fi12 := 0; fi12 < 7; fi12++ {
        v3 += (fi12 >> ((42) & 63))
        v1 = append(v1, (fi12 + -2147483648))
        for fi13 := 0; fi13 < 12; fi13++ {
            v3 -= (v3 - (fi12 * 2147483647))
            if true {
                v14 := GenS1{F0: 65535, F1: 2, F2: 4294967295}
                _ = v14
            } else {
                if true {
                    if (153 >= 128) {
                        v15 := GenS1{F0: 255, F1: 408, F2: 65535}
                        _ = v15
                    }
                    var v16 int32 = 2147483647
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v16
                } else {
                    v2 = false
                }
            }
            v2 = (((19 > 2147483648) || v2) && false)
            var v17 uint32 = 2147483647
            _ = fi13
            _ = v17
        }
        _ = fi12
    }
    switch "hello" {
    case "\n\t":
        if (1 <= 4294967295) {
            switch "\n\t" {
            case "\n\t":
                var v18 int = (v3 >> ((1) & 63))
                _ = v18
            case "a":
                var v19 int64 = 1
                _ = v19
            case "hello":
                v2 = (!(!v2))
            default:
                if len(v1) > 0 {
                    v1[0] = v3
                }
            }
        } else {
            for fi20 := 0; fi20 < 4; fi20++ {
                switch 42 {
                case -960:
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v2 = v2
                        switch 128 {
                        case 128:
                            if true {
                                switch 32768 {
                                case 185:
                                    var v22 uint16 = 327
                                    _ = v22
                                }
                            }
                        case 0:
                            v3 -= (v3 & (fi21 * (fi21 >> ((1) & 63))))
                        }
                        _ = fi21
                    }
                case 199:
                    if (-2147483648 <= 1) {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            var v24 uint = 4294967295
                            switch ("foo bar" + "foo bar") {
                            case "":
                                switch fi20 {
                                case 32768:
                                    var v25 string = ""
                                    _ = v25
                                case 0:
                                    _ = len(v1)
                                case 4294967295:
                                    if (-159 <= 986) {
                                        var v26 uint = v24
                                        v27 := GenS1{F0: 0, F1: 2147483647, F2: 444}
                                        var v28 uint32 = 2147483648
                                        v1 = append(v1, fi23)
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                default:
                                    v29 := GenS1{F0: 0, F1: 65535, F2: 2147483648}
                                    _ = v29
                                }
                            case "foo bar":
                                if ((^(fi20)) <= fi23) {
                                    if v2 {
                                        switch -503 {
                                        case -9223372036854775808:
                                            switch v3 {
                                            case -421:
                                                v2 = v2
                                            case 777:
                                                var v30 uint16 = 97
                                                _ = v30
                                            case 127:
                                                if false {
                                                    _ = len(v1)
                                                    v24 += 1
                                                } else {
                                                    v24 = v24
                                                }
                                            }
                                        case 185:
                                            v1 = append(v1, v3)
                                            var v31 float32 = 0.0
                                            v32 := GenS2{F0: v24, F1: 255}
                                            if (!(!((42 < 2) || true))) {
                                                v33 := GenS2{F0: 763, F1: 1}
                                                fmt.Println(v24, fi20)
                                                _ = v33
                                            } else {
                                                v2 = ((v31 * 0.1) <= float32(1e10))
                                                var v34 int64 = -190
                                                v24 += v24
                                                _ = v34
                                            }
                                            if (float64(v31) <= 0.1) {
                                                _ = len(v1)
                                                v35 := GenS3{F0: 176, F1: -9223372036854775808, F2: 2}
                                                _ = v35
                                            }
                                            var v36 uint32 = 42
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v3 -= (fi23 - (v3 - fi23))
                                            var v37 int16 = 42
                                            _ = v31
                                            _ = v32
                                            _ = v36
                                            _ = v37
                                        default:
                                            switch ("\n\t" + "foo bar") {
                                            case "\n\t":
                                                switch -144 {
                                                case 938:
                                                    fmt.Println(v2, v3, fi23)
                                                case 341:
                                                    fmt.Println(v24, fi23)
                                                }
                                            case "hello":
                                                v3 = (fi23 & fi23)
                                            }
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = -639
                                        }
                                    }
                                    v38 := GenS1{F0: 255, F1: 2147483647, F2: 4294967295}
                                    _ = len(v1)
                                    _ = v38
                                }
                            case "\n\t":
                                fmt.Println(v24)
                            default:
                                var v39 int32 = 0
                                _ = v39
                            }
                            _ = fi23
                            _ = v24
                        }
                    }
                default:
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 bool = (2 > 18446744073709551615)
                        switch 0 {
                        case -268:
                            _ = len(v1)
                        case -718:
                            v42 := GenS4{F0: (255 >= 25)}
                            _ = v42
                        case 2147483647:
                            var v43 int64 = -9223372036854775808
                            _ = v43
                        default:
                            switch (fi40 >> ((2) & 63)) {
                            case 128:
                                fmt.Println(v2)
                            case -499:
                                v44 := GenS1{F0: 897, F1: 2147483647, F2: 4294967295}
                                _ = v44
                            case 4294967295:
                                v2 = v41
                                var v45 uint = 2
                                v46 := GenS3{F0: 255, F1: -9223372036854775808, F2: 0}
                                _ = v45
                                _ = v46
                            default:
                                v47 := GenS3{F0: 102, F1: 0, F2: 1}
                                _ = v47
                            }
                        }
                        _ = fi40
                        _ = v41
                    }
                }
                _ = fi20
            }
            switch v3 {
            case 151:
                switch "" {
                case "\n\t":
                    v3 -= 128
                case "hello":
                    var v48 int32 = -2147483648
                    _ = v48
                case "foo bar":
                    var v49 float64 = ((0.1 + float64(-2.25)) / 3.14159)
                    _ = v49
                }
            }
            v50 := GenS3{F0: 183, F1: 0, F2: 105}
            _ = v50
        }
    case "a":
        switch "" {
        case "\n\t":
            var v51 float64 = 1e10
            _ = v51
        case "hello":
            var v52 uint64 = 4294967295
            _ = v52
        case "foo bar":
            v2 = true
        default:
            for fi53 := 0; fi53 < 9; fi53++ {
                fmt.Println(fi53)
                var v54 uint8 = 127
                _ = fi53
                _ = v54
            }
        }
    case "":
        v3 -= v3
    default:
        v55 := GenS5{F0: 1}
        _ = v55
    }
    for fi56 := 0; fi56 < 1; fi56++ {
        var v57 string = ("a" + "")
        v2 = (!(0 > 45))
        v58 := GenS1{F0: 1, F1: 146, F2: 2147483648}
        _ = fi56
        _ = v57
        _ = v58
    }
    v59 := GenS3{F0: 2, F1: 9223372036854775807, F2: 1}
    var v60 int8 = 100
    v1 = append(v1, (v3 % (((v3 + 128)) | 1)))
    if (1 >= -798) {
        v3 += 42
        v3 += (v3 ^ (v3 / ((1) | 1)))
    }
    var v61 float32 = 3.14159
    for fi62 := 0; fi62 < 5; fi62++ {
        var v63 uint32 = 742
        var v64 uint32 = v63
        _ = fi62
        _ = v63
        _ = v64
    }
    v65 := GenS5{F0: 665}
    for fi66 := 0; fi66 < 15; fi66++ {
        if v2 {
            var v67 uint32 = 4294967295
            _ = v67
        }
        _ = fi66
    }
    if v2 {
        var v68 uint64 = 2
        var v69 uint64 = (v68 & 0)
        _ = v68
        _ = v69
    } else {
        v70 := GenS1{F0: 255, F1: 2, F2: 0}
        v71 := GenS1{F0: 255, F1: 1, F2: 590}
        _ = v70
        _ = v71
    }
    v60 = v60
    if (42 <= -9223372036854775808) {
        var v72 int32 = -1
        _ = v72
    } else {
        for fi73 := 0; fi73 < 14; fi73++ {
            var v74 string = ""
            _ = fi73
            _ = v74
        }
        v75 := GenS5{F0: 139}
        v76 := GenS3{F0: 42, F1: -9223372036854775808, F2: 1}
        _ = v75
        _ = v76
    }
    for fi77 := 0; fi77 < 17; fi77++ {
        v78 := GenS6{F0: 65535}
        v2 = (!true)
        switch v3 {
        case 128:
            if v2 {
                _ = len(v1)
            } else {
                v61 += float32(0.0)
                fmt.Println(v60)
            }
        }
        _ = fi77
        _ = v78
    }
    fmt.Println(v61, v3)
    for fi79 := 0; fi79 < 17; fi79++ {
        for fi80 := 0; fi80 < 12; fi80++ {
            v81 := GenS5{F0: -756}
            if (!v2) {
                v3 += (fi80 ^ 128)
                v82 := GenS3{F0: 170, F1: -876, F2: 42}
                var v83 uint64 = genFn1(42, (^(v60)))
                _ = v82
                _ = v83
            }
            _ = fi80
            _ = v81
        }
        v84 := GenS7{F0: 402}
        _ = fi79
        _ = v84
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v60
    _ = v61
    _ = v65
}
