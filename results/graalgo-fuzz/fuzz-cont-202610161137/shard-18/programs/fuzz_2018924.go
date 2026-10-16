package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8) int32 {
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 int16 = 32767
            p0 += (p0 ^ (p0 + 42))
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS1{F0: 885, F1: 65535, F2: 1}
    var v8 int32 = 2147483647
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int32) uint8 {
    var v9 string = ""
    v10 := []int{32767}
    if len(v10) > 0 {
        v10[0] = 19
    }
    fmt.Println(v9)
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int64 = 952
    v2 := make(map[uint32]int)
    var v3 int32 = -861
    _ = len(v2)
    switch "a" {
    case "":
        var v11 uint8 = 42
        _ = v11
    }
    var v12 int32 = (v3 | v3)
    if (2 == 255) {
        if (-517 > -32768) {
            v13 := []int{-708}
            v13 = append(v13, 4294967295)
            v14 := GenS1{F0: 622, F1: 0, F2: 813}
            _ = v13
            _ = v14
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 float32 = 1.5
                v17 := GenS1{F0: 1, F1: 573, F2: 214}
                _ = fi15
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                v3 = genFn0(42)
                switch "hello" {
                case "hello":
                    var v19 uint = 4294967295
                    _ = v19
                case "a":
                    fmt.Println(v3, v3, v1)
                case "foo bar":
                    v20 := GenS1{F0: 1, F1: 412, F2: 0}
                    _ = v20
                }
                _ = fi18
            }
        }
    } else {
        fmt.Println(v12, v1, v1)
        v1 = (v1 << ((4294967295) & 63))
        v21 := GenS1{F0: 124, F1: 1, F2: 1}
        _ = v21
    }
    if true {
        if (true && ((int(v1) == 982) && (!(-128 > -128)))) {
            delete(v2, 95)
        } else {
            var v22 int8 = 100
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 12; fi23++ {
            if (!(0 < 138)) {
                for fi24 := 0; fi24 < 17; fi24++ {
                    switch ("\n\t" + ("foo bar" + "hello")) {
                    case "hello":
                        if true {
                            v25 := GenS1{F0: 42, F1: 65535, F2: 65535}
                            _ = v25
                        }
                    default:
                        v3 = v3
                    }
                    switch (fi24 - 1) {
                    case 127:
                        if (((fi23 & fi24) > fi24) && (2 > 42)) {
                            var v26 int16 = -98
                            switch ("a" + "hello") {
                            case "\n\t":
                                var v27 string = "foo bar"
                                _ = v27
                            default:
                                v28 := GenS1{F0: 121, F1: 42, F2: 2}
                                _ = v28
                            }
                            v3 -= v12
                            _ = v26
                        }
                    case -9223372036854775808:
                        switch ("a" + (("\n\t" + ("a" + "\n\t")) + "foo bar")) {
                        case "hello":
                            switch "a" {
                            case "a":
                                v1 = v1
                            default:
                                if true {
                                    v29 := GenS1{F0: 594, F1: 183, F2: 42}
                                    _ = v29
                                } else {
                                    v3 = (v12 >> ((0) & 63))
                                    fmt.Println(v3, fi24, fi24)
                                    v30 := GenS1{F0: 42, F1: 32767, F2: 4294967295}
                                    v31 := []int{fi24, fi24, fi23, fi24}
                                    _ = v30
                                    _ = v31
                                }
                            }
                        }
                    case 1:
                        v1 += 9223372036854775807
                    default:
                        var v32 uint8 = 1
                        _ = v32
                    }
                    _ = fi24
                }
                for fi33 := 0; fi33 < 5; fi33++ {
                    v12 = (v12 & v3)
                    _ = fi33
                }
            }
            _ = fi23
        }
        var v34 uint32 = 4294967295
        if false {
            for fi35 := 0; fi35 < 17; fi35++ {
                if true {
                    var v36 int32 = v3
                    _ = v36
                }
                var v37 uint = 964
                _ = fi35
                _ = v37
            }
            for fi38 := 0; fi38 < 17; fi38++ {
                v1 = (v1 | -9223372036854775808)
                v39 := GenS1{F0: 0, F1: 783, F2: (v34 << ((255) & 63))}
                _ = fi38
                _ = v39
            }
            v1 -= v1
        }
        _ = v34
    }
    if false {
        v1 += (v1 << ((291) & 63))
    }
    if false {
        for fi40 := 0; fi40 < 20; fi40++ {
            for fi41 := 0; fi41 < 7; fi41++ {
                var v42 float32 = (((1e10 + 0.1) - (-2.25 / (0.0 + 0.1))) + -2.25)
                v43 := []int{(-(fi40)), fi41, fi40}
                v44 := GenS2{F0: 0}
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
            var v45 int = fi40
            _ = fi40
            _ = v45
        }
        fmt.Println(v1, v1)
        var v46 uint8 = 1
        _ = v46
    }
    v12 = 42
    fmt.Println(v3, v12)
    fmt.Println(v3, v3, v1)
    var v47 uint = 42
    var v48 bool = (!(genFn1((v12 >> (((v47 / (((v47 >> ((v47) & 63))) | 1))) & 63))) != 1))
    var v49 float64 = 3.14159
    for fi50 := 0; fi50 < 3; fi50++ {
        for fi51 := 0; fi51 < 14; fi51++ {
            var v52 int32 = (v3 / (((v3 ^ (v12 - -535))) | 1))
            v47 = v47
            _ = fi51
            _ = v52
        }
        switch "\n\t" {
        case "":
            for fi53 := 0; fi53 < 17; fi53++ {
                v48 = (v49 != v49)
                switch "foo bar" {
                case "hello":
                    var v54 uint = (v47 << ((v47) & 63))
                    _ = v54
                case "\n\t":
                    v55 := v2[2147483647]
                    var v56 bool = (18446744073709551615 > 947)
                    _ = v55
                    _ = v56
                case "a":
                    if (true || (v48 && (false && (!v48)))) {
                        if (!v48) {
                            var v57 uint64 = 42
                            if v48 {
                                v12 -= (v3 & (-(v12)))
                                if (fi53 < 32768) {
                                    v58 := []int{fi53, -676, -913, (fi50 - 4294967295)}
                                    var v59 uint32 = 1
                                    var v60 int32 = (-(v12))
                                    fmt.Println(v47, v59)
                                    v60 -= (^(v12))
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                if true {
                                    var v61 float64 = float64(1.5)
                                    if (-892 != 42) {
                                        v12 = v12
                                        switch "\n\t" {
                                        case "hello":
                                            delete(v2, 166)
                                        case "a":
                                            var v62 uint8 = 127
                                            _ = v62
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            switch 128 {
                                            case 42:
                                                v63 := []int{2147483647, fi53, 2147483647, (fi53 - 32767)}
                                                _ = v63
                                            case 0:
                                                var v64 int8 = 127
                                                _ = v64
                                            case -737:
                                                var v65 int16 = 600
                                                _ = v65
                                            }
                                        case "":
                                            v66 := GenS2{F0: 997}
                                            switch int(v47) {
                                            case 42:
                                                switch ("\n\t" + "hello") {
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "hello":
                                                        v67 := []int{(-(fi50))}
                                                        _ = v67
                                                    case "a":
                                                        delete(v2, 974)
                                                    default:
                                                        v48 = v48
                                                    }
                                                default:
                                                    var v68 int8 = 100
                                                    v69 := []int{fi50, 128}
                                                    fmt.Println(v1, v61, fi50)
                                                    if v48 {
                                                        var v70 int8 = 0
                                                        _ = v70
                                                    } else {
                                                        v47 = (v47 * v47)
                                                        v71 := GenS3{F0: 0}
                                                        switch -138 {
                                                        case 9223372036854775807:
                                                            var v72 uint8 = genFn1((v3 | (-(v3))))
                                                            _ = v72
                                                        case 32768:
                                                            v73 := v2[4294967295]
                                                            _ = v73
                                                        }
                                                        _ = v71
                                                    }
                                                    _ = v68
                                                    _ = v69
                                                }
                                            case 125:
                                                var v74 int64 = v1
                                                _ = v74
                                            case 389:
                                                v75 := GenS2{F0: 2}
                                                _ = v75
                                            default:
                                                var v76 uint16 = 743
                                                _ = v76
                                            }
                                            _ = v66
                                        case "\n\t":
                                            v77 := []int{fi53, fi53, 2147483647}
                                            _ = v77
                                        }
                                    } else {
                                        var v78 uint16 = 1
                                        v79 := GenS4{F0: 2, F1: (true && false)}
                                        _ = v78
                                        _ = v79
                                    }
                                    v80 := []int{(fi50 << ((uint(v57)) & 63)), fi50, fi50}
                                    fmt.Println(v57)
                                    v80 = append(v80, fi53)
                                    _ = v61
                                    _ = v80
                                }
                            }
                            var v81 uint8 = 188
                            _ = v57
                            _ = v81
                        }
                        v1 += (v1 * 0)
                    }
                }
                var v82 int = (fi53 / (((fi50 << (((^(v47))) & 63))) | 1))
                _ = fi53
                _ = v82
            }
        default:
            var v83 uint16 = 0
            _ = v83
        }
        v84 := GenS5{F0: 1}
        _ = fi50
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
}
