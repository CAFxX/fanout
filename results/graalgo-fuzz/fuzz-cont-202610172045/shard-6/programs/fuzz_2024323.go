package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int32; F1 int; F2 int }
type GenS6 struct { F0 int16 }
func genFn0(p0 int64, p1 uint8) int64 {
    v6 := []int{}
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int = fi8
        for fi10 := 0; fi10 < 7; fi10++ {
            if true {
                p0 += p0
                v11 := GenS1{F0: -128, F1: 637}
                _ = v11
            } else {
                v12 := GenS1{F0: 42, F1: -1}
                v7 += v7
                var v13 float64 = 1.5
                _ = v12
                _ = v13
            }
            v7 -= v7
            v9 = 42
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint8, p1 int, p2 bool) uint16 {
    var v14 int32 = 42
    p1 -= 4294967295
    switch "\n\t" {
    case "hello":
        p1 = (p1 % ((127) | 1))
    case "a":
        var v15 uint = 4294967295
        _ = v15
    default:
        var v16 uint16 = 42
        _ = v16
    }
    v17 := []int{(p1 | p1), 32768, (p1 | (p1 * 9223372036854775807))}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v17
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 4294967295
    var v3 bool = false
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    var v18 uint32 = 2147483648
    v19 := []int{32767, 2147483647, -2147483648, 562}
    fmt.Println(v1)
    v18 += (v18 << ((505) & 63))
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint32 = v18
        fmt.Println(v21)
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 int = (fi22 % (((fi22 & (fi22 - (fi22 >> (((v2 & (v2 + v2))) & 63))))) | 1))
        var v24 int16 = 0
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 1, F1: 1000}
    _ = len(v5)
    switch "foo bar" {
    case "\n\t":
        switch -346 {
        case 4294967295:
            fmt.Println(v18, v2, v3)
        case 857:
            v26 := GenS2{F0: -128}
            _ = v26
        case -9223372036854775808:
            var v27 uint64 = v1
            _ = v27
        }
    default:
        switch 42 {
        case 194:
            v4[-9223372036854775808] = 451
        case 9223372036854775807:
            var v28 uint8 = 51
            _ = v28
        case 645:
            if len(v19) > 0 {
                v19[0] = 4294967295
            }
        default:
            if true {
                var v29 uint64 = 763
                _ = v29
            }
        }
    }
    v30 := GenS2{F0: 0}
    for fi31 := 0; fi31 < 8; fi31++ {
        fmt.Println(v18)
        var v32 int = (fi31 - 32767)
        _ = fi31
        _ = v32
    }
    fmt.Println(v18)
    for fi33 := 0; fi33 < 1; fi33++ {
        if v3 {
            v19 = append(v19, fi33)
            v34 := v4[fi33]
            if true {
                switch (fi33 & -490) {
                case -196:
                    for fi35 := 0; fi35 < 17; fi35++ {
                        v36 := v5[(fi33 - int(v2))]
                        if len(v19) > 0 {
                            v19[0] = (^(fi35))
                        }
                        var v37 int16 = int16(v18)
                        v3 = (!false)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                }
            }
            _ = v34
        } else {
            var v38 float64 = 0.0
            v39 := GenS2{F0: -64}
            _ = v38
            _ = v39
        }
        v40 := GenS1{F0: -15, F1: 1}
        if (!v3) {
            switch "" {
            case "\n\t":
                v3 = true
            case "hello":
                if v3 {
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v42 := GenS3{F0: 127, F1: -1}
                        if (-1 != -363) {
                            var v43 string = "hello"
                            v3 = ((v18 >> ((v2) & 63)) < v18)
                            var v44 uint8 = 255
                            _ = v43
                            _ = v44
                        } else {
                            fmt.Println(fi33)
                        }
                        fmt.Println(v3, v3)
                        fmt.Println(v1)
                        fmt.Println(v18, v18)
                        _ = fi41
                        _ = v42
                    }
                } else {
                    var v45 float64 = 3.14159
                    if (32767 <= -32768) {
                        if v3 {
                            v45 -= float64(-2.25)
                            var v46 int32 = -1
                            fmt.Println(v1)
                            _ = v46
                        }
                        if ((^(v18)) != 42) {
                            var v47 uint8 = 1
                            for fi48 := 0; fi48 < 11; fi48++ {
                                var v49 uint32 = (v18 ^ v18)
                                if (!(292 != 32767)) {
                                    v47 += 128
                                    fmt.Println(v3, fi48)
                                    var v50 bool = false
                                    _ = v50
                                }
                                _ = fi48
                                _ = v49
                            }
                            _ = v47
                        } else {
                            v45 += (((-2.25 * float64(1.5)) * v45) * (0.0 - v45))
                        }
                        for fi51 := 0; fi51 < 10; fi51++ {
                            if v3 {
                                switch (fi33 & 128) {
                                case 2147483647:
                                    v1 -= (v1 - 4294967295)
                                case 4294967295:
                                    v52 := GenS3{F0: 0, F1: -684}
                                    _ = v52
                                case 128:
                                    if false {
                                        if true {
                                            if v3 {
                                                switch "" {
                                                case "a":
                                                    if (42 != fi51) {
                                                        var v53 int = (fi51 << ((v2) & 63))
                                                        _ = v53
                                                    } else {
                                                        if (v45 < v45) {
                                                            v54 := v5[fi33]
                                                            _ = v54
                                                        } else {
                                                            switch -2147483648 {
                                                            case 38:
                                                                var v55 uint32 = 0
                                                                _ = v55
                                                            case -1:
                                                                v18 = 42
                                                            default:
                                                                fmt.Println(v1, fi33)
                                                            }
                                                            if (!(-370 >= -778)) {
                                                                if v3 {
                                                                    v56 := GenS2{F0: 0}
                                                                    var v57 float64 = v45
                                                                    if v3 {
                                                                        var v58 uint64 = v1
                                                                        v59 := GenS4{F0: v58, F1: v58}
                                                                        v60 := GenS3{F0: int8(v18), F1: -533}
                                                                        _ = v58
                                                                        _ = v59
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 float64 = ((float64((0.1 - float32((v45 / 0.1)))) + 0.1) * float64(1.5))
                                                                        switch 2147483647 {
                                                                        case 128:
                                                                            v1 = 4294967295
                                                                        case 9223372036854775807:
                                                                            if len(v19) > 0 {
                                                                                v19[0] = 43
                                                                            }
                                                                        case 336:
                                                                            var v62 uint64 = v1
                                                                            _ = v62
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    var v63 float32 = float32(v45)
                                                                    fmt.Println(v57, v2, v18)
                                                                    _ = v56
                                                                    _ = v57
                                                                    _ = v63
                                                                }
                                                                var v64 int16 = 1
                                                                v2 += 147
                                                                _ = v64
                                                            } else {
                                                                if (v45 <= v45) {
                                                                    if (1 <= 127) {
                                                                        var v65 string = ("a" + "foo bar")
                                                                        switch 42 {
                                                                        case -2147483648:
                                                                            if v3 {
                                                                                v5[fi51] = (v65 + v65)
                                                                            }
                                                                        case 2147483647:
                                                                            switch (v65 + (("foo bar" + (("\n\t" + ((("a" + "") + v65) + (v65 + (v65 + "foo bar")))) + v65)) + (("foo bar" + "foo bar") + v65))) {
                                                                            case "\n\t":
                                                                                if len(v19) >= 2 {
                                                                                    v19 = v19[1:len(v19)]
                                                                                }
                                                                            default:
                                                                                var v66 int = 4294967295
                                                                                _ = v66
                                                                            }
                                                                        default:
                                                                            var v67 int64 = 9223372036854775807
                                                                            _ = v67
                                                                        }
                                                                        var v68 float32 = (1e10 - 0.1)
                                                                        _ = v65
                                                                        _ = v68
                                                                    }
                                                                    if (fi33 > (fi51 + -13)) {
                                                                        v2 -= (^(v2))
                                                                    }
                                                                    fmt.Println(v2, v18)
                                                                }
                                                                var v69 int16 = 1000
                                                                var v70 int8 = -128
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            v3 = v3
                                                            if len(v19) > 0 {
                                                                v19[0] = fi51
                                                            }
                                                            v71 := GenS1{F0: -1, F1: -32768}
                                                            var v72 int = fi51
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    }
                                                case "foo bar":
                                                    v73 := GenS5{F0: 42, F1: fi51, F2: fi33}
                                                    _ = v73
                                                case "\n\t":
                                                    v74 := GenS6{F0: 32767}
                                                    _ = v74
                                                default:
                                                    v18 += (^(v18))
                                                }
                                            } else {
                                                var v75 int16 = -482
                                                _ = v75
                                            }
                                            var v76 uint8 = 127
                                            _ = v76
                                        } else {
                                            v1 += (v1 / (((v1 >> ((v2) & 63))) | 1))
                                        }
                                        v77 := GenS5{F0: -67, F1: fi51, F2: 0}
                                        _ = v77
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v45 += 1.5
                                        case "a":
                                            _ = len(v4)
                                        default:
                                            switch ("" + "") {
                                            case "foo bar":
                                                fmt.Println(v2)
                                            }
                                        }
                                        v3 = v3
                                    }
                                }
                                switch 4294967295 {
                                case 1:
                                    v3 = (!(!(!v3)))
                                case 32768:
                                    var v78 float32 = 1e10
                                    _ = v78
                                }
                                v18 = 2147483648
                            } else {
                                v2 = v2
                                var v79 bool = (!(-9223372036854775808 <= 9223372036854775807))
                                _ = v79
                            }
                            var v80 int64 = 0
                            _ = fi51
                            _ = v80
                        }
                    }
                    fmt.Println(v45)
                    _ = v45
                }
            case "a":
                var v81 float32 = 1e10
                _ = v81
            }
            fmt.Println(v2)
            var v82 bool = (!(v1 >= 0))
            _ = v82
        } else {
            var v83 int = 1
            _ = v83
        }
        _ = fi33
        _ = v40
    }
    var v84 string = "hello"
    for fi85 := 0; fi85 < 15; fi85++ {
        v3 = (!true)
        if (!true) {
            var v86 int = (fi85 * 128)
            _ = v86
        }
        v18 += (^(v18))
        _ = fi85
    }
    delete(v4, -2147483648)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v84:", v84)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v84
}
