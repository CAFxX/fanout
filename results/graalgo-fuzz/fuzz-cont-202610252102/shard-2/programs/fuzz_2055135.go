package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int32; F1 string; F2 int }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
type GenS4 struct { F0 int32; F1 int64; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint8; F1 string; F2 int }
type GenS7 struct { F0 uint32; F1 uint16; F2 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 bool = false
    v4 := make(map[int]uint32)
    var v5 uint8 = 42
    var v6 uint32 = 2147483647
    fmt.Println(v6)
    for fi7 := 0; fi7 < 9; fi7++ {
        v3 = (!v3)
        if true {
            switch "foo bar" {
            case "foo bar":
                var v8 int16 = 0
                _ = v8
            case "a":
                var v9 float64 = 3.14159
                _ = v9
            case "\n\t":
                v10 := v4[-2147483648]
                _ = v10
            default:
                v11 := GenS1{F0: -9223372036854775808, F1: "hello"}
                _ = v11
            }
            var v12 int64 = 609
            delete(v1, (v6 << ((147) & 63)))
            _ = v12
        }
        _ = fi7
    }
    v4[-729] = (v6 | (v6 - 324))
    var v13 uint32 = (v6 ^ (v6 + (v6 >> ((30) & 63))))
    for fi14 := 0; fi14 < 7; fi14++ {
        switch "a" {
        case "hello":
            v15 := []int{(fi14 * fi14), (fi14 >> ((42) & 63))}
            _ = v15
        }
        switch (fi14 % ((fi14) | 1)) {
        case 1:
            fmt.Println(v13, v3)
        case 0:
            v6 += (v13 & v6)
        }
        _ = fi14
    }
    v16 := GenS1{F0: -273, F1: "\n\t"}
    if v3 {
        if false {
            switch 4294967295 {
            case 42:
                v17 := GenS1{F0: 9223372036854775807, F1: ("" + "foo bar")}
                _ = v17
            default:
                var v18 int16 = 1000
                _ = v18
            }
            switch "\n\t" {
            case "a":
                v5 = (^(v5))
            case "\n\t":
                switch 32768 {
                case -1:
                    var v19 bool = v3
                    _ = v19
                default:
                    v6 = v6
                }
            case "hello":
                var v20 uint = 0
                _ = v20
            default:
                v21 := GenS1{F0: 1, F1: "foo bar"}
                _ = v21
            }
            delete(v4, 32768)
        }
        switch "hello" {
        case "\n\t":
            if v3 {
                var v22 bool = (1 >= 65535)
                fmt.Println(v13)
                for fi23 := 0; fi23 < 1; fi23++ {
                    if (((fi23 ^ (fi23 * (fi23 ^ fi23))) >= (fi23 | (fi23 & -1))) && ((((uint32(v5) >= (v6 + (v13 >> ((362) & 63)))) || v22) || false) && (-504 < 597))) {
                        v22 = v3
                    } else {
                        _ = len(v2)
                        v13 = v6
                        switch (fi23 + fi23) {
                        case 127:
                            v24 := GenS2{F0: int32(fi23), F1: "hello", F2: (-(fi23))}
                            _ = v24
                        }
                    }
                    v5 = (v5 % (((v5 / ((uint8(v6)) | 1))) | 1))
                    if v3 {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 float64 = (3.14159 * 0.0)
                            _ = fi25
                            _ = v26
                        }
                        fmt.Println(v22, v3)
                    } else {
                        switch (fi23 & -880) {
                        case 32768:
                            if v22 {
                                var v27 int64 = 0
                                var v28 int = (-(fi23))
                                _ = v27
                                _ = v28
                            } else {
                                v29 := GenS2{F0: -216, F1: "hello", F2: (-(fi23))}
                                var v30 int16 = -32768
                                if (!((-9223372036854775808 == 688) || v3)) {
                                    v22 = v22
                                    v30 = v30
                                } else {
                                    for fi31 := 0; fi31 < 12; fi31++ {
                                        v32 := GenS3{F0: -1, F1: v5, F2: 42}
                                        _ = fi31
                                        _ = v32
                                    }
                                    var v33 uint8 = 2
                                    _ = v33
                                }
                                _ = v29
                                _ = v30
                            }
                        case 202:
                            v34 := GenS4{F0: -1, F1: 42, F2: 50}
                            _ = v34
                        default:
                            v35 := GenS5{F0: v22, F1: 4294967295}
                            _ = v35
                        }
                        fmt.Println(v22, v13, v5)
                        v6 = (v13 % ((v13) | 1))
                    }
                    _ = fi23
                }
                _ = v22
            }
        case "hello":
            if false {
                v36 := v4[714]
                if (1 != v36) {
                    v6 += v36
                    fmt.Println(v6)
                    var v37 bool = v3
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 9; fi38++ {
                        if ((fi38 & (fi38 * -1)) >= 0) {
                            v6 += (v13 << ((255) & 63))
                            switch (-(fi38)) {
                            case -9223372036854775808:
                                var v39 float32 = (0.1 - 1.5)
                                _ = v39
                            }
                        }
                        _ = fi38
                    }
                    fmt.Println(v36, v6)
                }
                for fi40 := 0; fi40 < 2; fi40++ {
                    if (643 > 0) {
                        v6 = 531
                    } else {
                        v36 -= v13
                        v41 := GenS3{F0: 0, F1: (v5 << ((997) & 63)), F2: 4294967295}
                        _ = v41
                    }
                    if true {
                        switch "hello" {
                        case "":
                            var v42 bool = false
                            _ = v42
                        case "\n\t":
                            for fi43 := 0; fi43 < 18; fi43++ {
                                v44 := v4[(fi40 & fi40)]
                                var v45 float64 = 3.14159
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        case "hello":
                            v46 := GenS4{F0: 2147483647, F1: -647, F2: (v13 & (v13 ^ (v13 + (v13 << ((0) & 63)))))}
                            _ = v46
                        }
                        if (!false) {
                            for fi47 := 0; fi47 < 6; fi47++ {
                                var v48 float32 = -2.25
                                v49 := GenS6{F0: 1, F1: "", F2: -2147483648}
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                            switch "" {
                            case "\n\t":
                                v50 := []int{183, (fi40 % ((fi40) | 1))}
                                _ = v50
                            case "":
                                if v3 {
                                    delete(v2, v6)
                                    var v51 uint8 = v5
                                    _ = v51
                                } else {
                                    for fi52 := 0; fi52 < 18; fi52++ {
                                        v36 += v6
                                        if (0.1 != 1.5) {
                                            v53 := GenS7{F0: v6, F1: 182, F2: (fi52 + fi40)}
                                            fmt.Println(fi52, v5)
                                            _ = v53
                                        }
                                        _ = fi52
                                    }
                                    for fi54 := 0; fi54 < 1; fi54++ {
                                        v55 := []int{}
                                        switch (fi54 + (fi54 >> ((42) & 63))) {
                                        case -2147483648:
                                            v1[(^(v36))] = (fi54 % ((fi54) | 1))
                                        default:
                                            var v56 bool = false
                                            if v3 {
                                                delete(v2, v36)
                                            } else {
                                                switch (fi40 / ((9223372036854775807) | 1)) {
                                                case 4294967295:
                                                    v57 := GenS3{F0: 32767, F1: v5, F2: uint(v6)}
                                                    v56 = (18446744073709551615 > 943)
                                                    v36 += (v13 % ((42) | 1))
                                                    v6 += (^(v13))
                                                    _ = v57
                                                }
                                                var v58 float64 = (1e10 * (-2.25 + 0.1))
                                                _ = v58
                                            }
                                            _ = len(v55)
                                            switch "a" {
                                            case "foo bar":
                                                v59 := GenS7{F0: v6, F1: 0, F2: (fi40 | 0)}
                                                _ = v59
                                            case "hello":
                                                v60 := GenS6{F0: (v5 | 128), F1: ("foo bar" + ("" + ("\n\t" + "hello"))), F2: (-(fi54))}
                                                _ = v60
                                            }
                                            v61 := GenS8{F0: -41}
                                            fmt.Println(v36, fi54)
                                            _ = v56
                                            _ = v61
                                        }
                                        var v62 int16 = -32768
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v63 uint = 905
                                            _ = v63
                                        case "\n\t":
                                            var v64 int16 = -32768
                                            _ = v64
                                        case "hello":
                                            v62 -= (v62 >> ((uint(v36)) & 63))
                                        }
                                        _ = fi54
                                        _ = v55
                                        _ = v62
                                    }
                                    v5 = uint8(v36)
                                }
                            }
                            var v65 uint16 = 32767
                            _ = v65
                        } else {
                            v1[v6] = 807
                            delete(v2, v13)
                        }
                    } else {
                        v66 := GenS9{F0: -32768, F1: ""}
                        _ = v66
                    }
                    v4[(fi40 % ((4294967295) | 1))] = v13
                    _ = fi40
                }
                _ = v36
            } else {
                if (0 <= 42) {
                    var v67 string = "foo bar"
                    _ = v67
                } else {
                    v5 = (v5 + v5)
                }
                var v68 float64 = 0.1
                _ = v68
            }
        case "a":
            var v69 uint32 = 0
            _ = v69
        }
        var v70 uint16 = 65535
        _ = v70
    } else {
        if v3 {
            v71 := []int{}
            switch -9223372036854775808 {
            case -130:
                var v72 string = "foo bar"
                _ = v72
            case -9223372036854775808:
                switch "" {
                case "a":
                    _ = len(v4)
                }
            }
            _ = v71
        }
        switch 128 {
        case -1:
            if true {
                var v73 float32 = 0.0
                fmt.Println(v3, v5, v5)
                _ = v73
            }
        case 0:
            v13 = (v6 >> ((42) & 63))
        case -473:
            var v74 float32 = ((1e10 + 0.0) - 0.1)
            _ = v74
        }
        v75 := []int{-559, -1, 4294967295, 127}
        _ = v75
    }
    v76 := []int{-529, 128, -546}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v76:", len(v76), v76)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
