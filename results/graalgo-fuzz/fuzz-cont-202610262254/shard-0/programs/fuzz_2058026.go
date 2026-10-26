package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int32; F1 int; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 int8; F1 int16; F2 int8 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 uint32) uint32 {
    for fi5 := 0; fi5 < 12; fi5++ {
        switch fi5 {
        case -1:
            if (720 == 789) {
                v6 := GenS1{F0: 2, F1: 2, F2: (^(p0))}
                p0 = p0
                var v7 bool = ((-959 != 347) || (42 <= 9223372036854775807))
                _ = v6
                _ = v7
            }
        case -633:
            fmt.Println(p0, fi5)
        case 32768:
            fmt.Println(p0, p0)
        }
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p0)
        case "hello":
            for fi8 := 0; fi8 < 15; fi8++ {
                switch (fi5 & (fi8 ^ -862)) {
                case 42:
                    v9 := []int{-426, fi5, (fi8 >> ((664) & 63))}
                    _ = v9
                }
                var v10 int8 = -1
                v11 := []int{fi8, fi8, fi8, (fi8 & 1)}
                _ = fi8
                _ = v10
                _ = v11
            }
        case "foo bar":
            switch (-(fi5)) {
            case -9223372036854775808:
                v12 := []int{-9223372036854775808, (fi5 >> ((1) & 63)), (fi5 ^ fi5)}
                _ = v12
            case 2147483647:
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 int32 = -2147483648
                    var v15 uint = 108
                    var v16 int8 = 42
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            case -454:
                if false {
                    if (true && (255 >= 236)) {
                        var v17 int = (fi5 & (fi5 >> ((159) & 63)))
                        _ = v17
                    } else {
                        var v18 uint = 4294967295
                        v19 := GenS1{F0: 42, F1: 2, F2: (p0 | p0)}
                        v20 := GenS1{F0: 1, F1: 0, F2: p0}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    switch "\n\t" {
                    case "":
                        v21 := GenS1{F0: 0, F1: 1, F2: p0}
                        _ = v21
                    case "hello":
                        var v22 string = "hello"
                        _ = v22
                    case "foo bar":
                        v23 := GenS1{F0: 192, F1: 184, F2: (^(p0))}
                        _ = v23
                    }
                    v24 := []int{fi5, fi5}
                    _ = v24
                } else {
                    var v25 uint16 = 42
                    var v26 uint32 = (p0 - (p0 ^ p0))
                    _ = v25
                    _ = v26
                }
            default:
                p0 = 2147483648
            }
        default:
            if true {
                var v27 bool = false
                v27 = false
                if (-329 == 857) {
                    if (fi5 == (fi5 & (fi5 / (((fi5 ^ fi5)) | 1)))) {
                        v27 = (!(53 < 100))
                        var v28 uint32 = (p0 % ((p0) | 1))
                        v28 = p0
                        _ = v28
                    }
                    v29 := GenS1{F0: 2, F1: 63, F2: p0}
                    switch 155 {
                    case 127:
                        fmt.Println(p0, v27, fi5)
                    }
                    _ = v29
                }
                _ = v27
            } else {
                p0 += (p0 + (p0 + (p0 & (p0 & p0))))
                if (65535 != 42) {
                    var v30 int8 = 0
                    if false {
                        v31 := GenS2{F0: 289, F1: -32768}
                        switch ("a" + ("hello" + ((("a" + (("a" + ("hello" + "")) + "\n\t")) + "") + ("hello" + "a")))) {
                        case "a":
                            var v32 uint8 = 171
                            _ = v32
                        default:
                            v33 := []int{fi5, fi5, (fi5 + (fi5 - (fi5 << ((405) & 63))))}
                            _ = v33
                        }
                        _ = v31
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(p0, p0)
                    default:
                        switch fi5 {
                        case -6:
                            v34 := []int{fi5}
                            _ = v34
                        }
                    }
                    _ = v30
                } else {
                    if (4294967295 < 4294967295) {
                        p0 += p0
                    } else {
                        var v35 uint = 63
                        var v36 bool = (true && false)
                        _ = v35
                        _ = v36
                    }
                }
                var v37 float32 = (1.5 + 0.1)
                _ = v37
            }
        }
        _ = fi5
    }
    var v38 float32 = 1e10
    _ = p0
    _ = v38
    return p0
}

func main() {
    var v1 uint8 = 127
    v2 := []int{-2147483648, -404, 127, -878}
    var v3 int64 = 0
    var v4 uint8 = 61
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    var v39 uint = 4294967295
    var v40 uint16 = 65535
    v3 = v3
    var v41 uint32 = 2
    if (0.0 == (float32((1e10 * ((0.0 - 0.1) / 1.5))) / (0.1 / -2.25))) {
        switch "\n\t" {
        case "":
            v42 := GenS3{F0: 2147483647, F1: -9223372036854775808, F2: (v41 % ((v41) | 1))}
            _ = v42
        case "\n\t":
            if ((-(v3)) != v3) {
                switch "foo bar" {
                case "a":
                    switch 42 {
                    case 9223372036854775807:
                        v43 := GenS4{F0: v39}
                        _ = v43
                    default:
                        v44 := GenS5{F0: 0, F1: v3}
                        _ = v44
                    }
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    for fi45 := 0; fi45 < 10; fi45++ {
                        if true {
                            var v46 int16 = -32768
                            v2 = append(v2, (fi45 ^ fi45))
                            if false {
                                var v47 int8 = 127
                                _ = v47
                            }
                            _ = v46
                        } else {
                            var v48 int16 = -1
                            v49 := GenS6{F0: -796, F1: 72}
                            _ = v48
                            _ = v49
                        }
                        if (38 < 42) {
                            fmt.Println(v3, v41)
                            fmt.Println(v40, fi45)
                            if (true && (3.14159 == 0.1)) {
                                var v50 uint16 = (^(v40))
                                v51 := GenS7{F0: 681}
                                for fi52 := 0; fi52 < 16; fi52++ {
                                    v53 := GenS8{F0: (!false), F1: uint32(v50)}
                                    _ = fi52
                                    _ = v53
                                }
                                _ = v50
                                _ = v51
                            } else {
                                var v54 int8 = 0
                                _ = v54
                            }
                        }
                        var v55 uint64 = 2
                        _ = fi45
                        _ = v55
                    }
                }
                var v56 uint64 = 1
                v57 := GenS3{F0: 0, F1: -9223372036854775808, F2: v41}
                _ = v56
                _ = v57
            }
        case "foo bar":
            v1 = v1
        }
        switch "foo bar" {
        case "a":
            fmt.Println(v4)
        }
        for fi58 := 0; fi58 < 14; fi58++ {
            v59 := GenS9{F0: -84, F1: -32768, F2: -1}
            for fi60 := 0; fi60 < 15; fi60++ {
                v4 -= (^(v4))
                _ = fi60
            }
            _ = fi58
            _ = v59
        }
    }
    for fi61 := 0; fi61 < 20; fi61++ {
        v3 += (v3 - v3)
        _ = fi61
    }
    _ = len(v2)
    fmt.Println(v41, v4, v3)
    v41 += 4294967295
    v40 = v40
    v4 += (^(v1))
    for fi62 := 0; fi62 < 17; fi62++ {
        var v63 uint16 = 32767
        _ = fi62
        _ = v63
    }
    for fi64 := 0; fi64 < 18; fi64++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(v40)
        case "foo bar":
            v65 := GenS6{F0: 0, F1: (^(v39))}
            _ = v65
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            if (!false) {
                v2 = append(v2, fi64)
                if (!(v1 == v4)) {
                    v4 = 2
                }
                if (((fi64 >> ((956) & 63)) <= 726) && true) {
                    v66 := GenS10{F0: v40}
                    var v67 int32 = 775
                    _ = v66
                    _ = v67
                }
            }
        }
        v2 = append(v2, (fi64 & fi64))
        _ = fi64
    }
    v41 -= v41
    v68 := GenS3{F0: 42, F1: 669, F2: v41}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v68
}
