package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint8; F1 bool; F2 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
func genFn0() uint8 {
    v7 := []int{128, 9223372036854775807, 42, 127}
    _ = len(v7)
    v8 := GenS1{F0: 744, F1: 42, F2: 1}
    v7 = append(v7, 4294967295)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v7
    _ = v8
    return 84
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 2147483647
    var v3 uint64 = 104
    var v4 uint64 = 0
    v5 := make(map[int]bool)
    v6 := make(map[string]uint32)
    switch ("a" + "foo bar") {
    case "hello":
        v3 = v4
    case "foo bar":
        var v9 uint = 272
        _ = v9
    case "a":
        if true {
            switch -9223372036854775808 {
            case 32768:
                if true {
                    var v10 int8 = 1
                    v11 := GenS1{F0: 42, F1: v2, F2: 42}
                    _ = v10
                    _ = v11
                }
            case 1:
                fmt.Println(v2, v3, v4)
            case -810:
                if (true && (!(0 > 1))) {
                    v12 := []int{}
                    _ = v12
                }
            }
            var v13 int8 = 100
            var v14 uint64 = 0
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v2)
    var v15 uint = 4294967295
    fmt.Println(v4)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int64 = 9223372036854775807
        var v18 int32 = (v1 >> ((v15) & 63))
        v19 := []int{(^(fi16)), (fi16 ^ (fi16 + (fi16 / ((2147483647) | 1))))}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(v1)
    v3 = (v3 << (((v15 + v15)) & 63))
    switch -9223372036854775808 {
    case 638:
        var v20 float64 = float64(1.5)
        _ = v20
    case -463:
        if true {
            fmt.Println(v15)
            v21 := []int{-9223372036854775808, 128, -1, 952}
            _ = v21
        } else {
            v15 -= (^(v15))
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 uint32 = 65535
                if (-577 >= 653) {
                    v2 -= uint32(v15)
                    fmt.Println(v2, fi22)
                    delete(v5, fi22)
                }
                _ = len(v5)
                _ = fi22
                _ = v23
            }
            v24 := GenS1{F0: 110, F1: (v2 % ((v2) | 1)), F2: 127}
            _ = v24
        }
    default:
        if (!true) {
            v4 += v3
            for fi25 := 0; fi25 < 11; fi25++ {
                if (!true) {
                    v2 -= v2
                    v26 := GenS2{F0: v1, F1: fi25}
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v2 += (v2 % ((v2) | 1))
                        var v28 float64 = 1.5
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                } else {
                    var v29 int16 = -1
                    v30 := GenS3{F0: 1, F1: (v2 / (((v2 - v2)) | 1)), F2: (v3 * (v3 * v3))}
                    var v31 float32 = -2.25
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi25
            }
            for fi32 := 0; fi32 < 8; fi32++ {
                var v33 int = (fi32 * (fi32 & fi32))
                var v34 float64 = (-2.25 - -2.25)
                _ = fi32
                _ = v33
                _ = v34
            }
        }
    }
    var v35 bool = false
    if (105 == genFn0()) {
        if ((v2 >> (((v15 >> ((v15) & 63))) & 63)) > 295) {
            if (!false) {
                v2 -= v2
                switch 32767 {
                case 128:
                    v15 += 1
                case -574:
                    var v36 int32 = (-(v1))
                    _ = v36
                case 32767:
                    var v37 uint32 = uint32(v15)
                    _ = v37
                default:
                    v4 -= 83
                }
                for fi38 := 0; fi38 < 10; fi38++ {
                    if v35 {
                        var v39 float64 = ((3.14159 / (0.1 + 1.5)) / float64((0.1 * 0.1)))
                        for fi40 := 0; fi40 < 1; fi40++ {
                            if false {
                                v2 += (^(v2))
                                v4 = uint64(v2)
                                var v41 uint64 = v4
                                _ = v41
                            }
                            _ = fi40
                        }
                        if v35 {
                            switch (^(fi38)) {
                            case 4294967295:
                                v15 = (v15 * v15)
                            }
                            if (v2 <= 2147483647) {
                                v1 -= (v1 * v1)
                            } else {
                                v42 := []int{}
                                var v43 int16 = 32767
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS4{F0: -2147483648}
                            _ = v44
                        }
                        _ = v39
                    } else {
                        v3 = v4
                    }
                    if v35 {
                        v45 := []int{fi38, (fi38 * (fi38 / ((fi38) | 1))), (fi38 & fi38)}
                        for fi46 := 0; fi46 < 2; fi46++ {
                            var v47 int = 0
                            _ = fi46
                            _ = v47
                        }
                        _ = v45
                    }
                    v48 := []int{(fi38 / ((1) | 1)), 42, fi38}
                    _ = fi38
                    _ = v48
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    for fi49 := 0; fi49 < 9; fi49++ {
                        switch fi49 {
                        case -1:
                            if v35 {
                                var v50 int = 42
                                for fi51 := 0; fi51 < 8; fi51++ {
                                    var v52 int32 = v1
                                    v53 := GenS3{F0: (v3 - 984), F1: v2, F2: 761}
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                                if true {
                                    v15 = (v15 / (((^(v15))) | 1))
                                    switch -2147483648 {
                                    case 98:
                                        var v54 uint32 = 668
                                        _ = v54
                                    default:
                                        var v55 uint64 = v4
                                        _ = v55
                                    }
                                }
                                _ = v50
                            }
                        case 1:
                            var v56 int8 = 44
                            _ = v56
                        default:
                            fmt.Println(v35, v3, v3)
                        }
                        _ = fi49
                    }
                case "hello":
                    for fi57 := 0; fi57 < 2; fi57++ {
                        v15 -= v15
                        v3 = v4
                        _ = fi57
                    }
                case "":
                    for fi58 := 0; fi58 < 10; fi58++ {
                        for fi59 := 0; fi59 < 10; fi59++ {
                            if (("" != ("\n\t" + "a")) || v35) {
                                v60 := []int{fi59}
                                v61 := GenS5{F0: -235, F1: 816}
                                if (fi58 > (fi58 * fi59)) {
                                    if ((fi58 & (-(fi58))) <= fi59) {
                                        v5[fi59] = v35
                                        if v35 {
                                            var v62 bool = v35
                                            var v63 int = (fi58 & 42)
                                            fmt.Println(v62, v1, fi58)
                                            _ = v62
                                            _ = v63
                                        }
                                    } else {
                                        v64 := v6["a"]
                                        v1 = v1
                                        _ = v64
                                    }
                                }
                                _ = v60
                                _ = v61
                            } else {
                                v65 := GenS1{F0: 255, F1: v2, F2: uint8(v2)}
                                _ = v65
                            }
                            _ = fi59
                        }
                        for fi66 := 0; fi66 < 4; fi66++ {
                            var v67 uint = (v15 + v15)
                            v68 := GenS6{F0: 128, F1: (65535 >= 2147483647), F2: ("foo bar" + (("" + ("\n\t" + "hello")) + ""))}
                            switch "a" {
                            case "a":
                                v69 := GenS7{F0: 423}
                                _ = v69
                            case "":
                                if ((v35 || (((v1 * v1) >= v1) || v35)) || (42 <= 440)) {
                                    v70 := GenS7{F0: -32768}
                                    _ = v70
                                }
                            default:
                                var v71 uint64 = (v3 - (v4 % (((v3 & 2)) | 1)))
                                _ = v71
                            }
                            v6["hello"] = 4294967295
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                        _ = fi58
                    }
                }
                switch 127 {
                case 421:
                    switch 0 {
                    case -137:
                        v72 := GenS2{F0: v1, F1: 750}
                        _ = v72
                    case 727:
                        var v73 uint8 = 166
                        _ = v73
                    }
                case -1:
                    v2 = 4294967295
                }
            }
            v6[("" + "hello")] = (v2 >> ((v15) & 63))
        }
        fmt.Println(v4, v4)
    } else {
        if (!(!(v35 && false))) {
            var v74 int16 = 776
            var v75 uint16 = 1
            _ = v74
            _ = v75
        } else {
            delete(v6, "a")
            v4 += (v4 - v4)
            var v76 string = ("foo bar" + "a")
            _ = v76
        }
        delete(v6, "a")
    }
    v77 := GenS8{F0: -1}
    switch 4294967295 {
    case 9223372036854775807:
        v78 := []int{9223372036854775807, 127, -842, 0}
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v77
}
