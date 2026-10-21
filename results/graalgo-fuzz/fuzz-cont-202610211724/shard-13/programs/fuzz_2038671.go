package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int32; F1 int8 }
func genFn0(p0 string, p1 bool) bool {
    if p1 {
        switch -97 {
        case -1:
            var v5 int32 = -1
            _ = v5
        default:
            p1 = (p1 || p1)
        }
        var v6 uint8 = 255
        _ = v6
    }
    p0 = ((p0 + p0) + ("hello" + "hello"))
    switch (p0 + "foo bar") {
    case "hello":
        v7 := GenS1{F0: 1}
        _ = v7
    case "":
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint8 = 127
            v9 = v9
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 uint, p1 int16) uint32 {
    v10 := GenS1{F0: 1}
    for fi11 := 0; fi11 < 11; fi11++ {
        p1 = (p1 >> (((p0 + (p0 << ((p0) & 63)))) & 63))
        _ = fi11
    }
    if true {
        var v12 uint16 = 255
        _ = v12
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v10
    return 442
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 0
    var v3 uint8 = 0
    var v4 int16 = 1
    var v13 uint32 = 512
    for fi14 := 0; fi14 < 7; fi14++ {
        switch "a" {
        case "a":
            v3 -= v3
        case "hello":
            for fi15 := 0; fi15 < 13; fi15++ {
                if true {
                    v2 += (fi15 & (v2 / ((fi14) | 1)))
                    v16 := GenS2{F0: 416, F1: v3, F2: ((0 < uint16(v13)) && (true || (false && true)))}
                    var v17 bool = false
                    _ = v16
                    _ = v17
                }
                fmt.Println(v3, v2)
                if (!(!false)) {
                    switch "\n\t" {
                    case "hello":
                        switch ("foo bar" + "\n\t") {
                        case "":
                            var v18 int8 = -60
                            _ = v18
                        }
                    }
                }
                v4 += v4
                _ = fi15
            }
        }
        _ = fi14
    }
    var v19 uint64 = 42
    fmt.Println(v3)
    switch "a" {
    case "\n\t":
        v20 := []int{(v2 + (-(v2))), (v2 >> ((971) & 63)), v2}
        _ = v20
    case "foo bar":
        if (!(v2 >= 9223372036854775807)) {
            v1[genFn1(187, v4)] = true
        } else {
            v21 := GenS3{F0: 255, F1: "\n\t"}
            v22 := v1[2147483648]
            _ = v21
            _ = v22
        }
    }
    var v23 string = "hello"
    v24 := []int{961, v2, -228, (-(v2))}
    v25 := GenS2{F0: v19, F1: v3, F2: true}
    v2 -= (v2 % ((v2) | 1))
    var v26 int32 = -2147483648
    if (false && false) {
        var v27 int8 = 42
        var v28 uint32 = (v13 << ((2) & 63))
        var v29 uint32 = v13
        _ = v27
        _ = v28
        _ = v29
    } else {
        switch v23 {
        case "foo bar":
            var v30 int64 = 9223372036854775807
            _ = v30
        case "\n\t":
            var v31 int32 = v26
            _ = v31
        case "":
            if false {
                if false {
                    for fi32 := 0; fi32 < 8; fi32++ {
                        fmt.Println(v2)
                        v33 := GenS2{F0: 542, F1: 216, F2: ((v3 % ((v3) | 1)) == 79)}
                        _ = fi32
                        _ = v33
                    }
                    v34 := GenS3{F0: 1, F1: ((("foo bar" + v23) + ("hello" + v23)) + ((v23 + v23) + v23))}
                    _ = v34
                } else {
                    if true {
                        for fi35 := 0; fi35 < 6; fi35++ {
                            var v36 float32 = 1.5
                            var v37 uint32 = (v13 + (v13 & 0))
                            var v38 uint8 = v3
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 10; fi39++ {
                            switch fi39 {
                            case 4294967295:
                                v13 -= (^(v13))
                            case -9223372036854775808:
                                v26 += int32(v19)
                            }
                            for fi40 := 0; fi40 < 12; fi40++ {
                                var v41 float32 = ((1e10 / 1.0) / -2.25)
                                if true {
                                    v42 := GenS4{F0: ((float64(v41) - 1.5) <= 1e10), F1: v19, F2: 255}
                                    _ = len(v24)
                                    switch v23 {
                                    case "hello":
                                        v3 += 2
                                    default:
                                        fmt.Println(v13, v13)
                                    }
                                    _ = v42
                                } else {
                                    switch fi40 {
                                    case -576:
                                        var v43 float64 = -2.25
                                        _ = v43
                                    case 9223372036854775807:
                                        fmt.Println(v13, v13)
                                    case 32767:
                                        v44 := GenS4{F0: true, F1: v19, F2: 32767}
                                        _ = v44
                                    default:
                                        v4 -= (-(v4))
                                    }
                                    v45 := GenS1{F0: v13}
                                    v46 := GenS2{F0: (^(v19)), F1: 127, F2: true}
                                    _ = v45
                                    _ = v46
                                }
                                _ = fi40
                                _ = v41
                            }
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                            _ = fi39
                        }
                        var v47 uint16 = uint16(v13)
                        _ = v47
                    } else {
                        switch (v2 | (v2 << ((4294967295) & 63))) {
                        case 4294967295:
                            v23 = (v23 + (v23 + v23))
                        case 32767:
                            if ((0.1 - ((float32(0.0) * 3.14159) - 0.1)) != -2.25) {
                                v26 = v26
                                var v48 float32 = 1e10
                                v49 := GenS5{F0: v19, F1: (v4 >> ((627) & 63)), F2: 236}
                                _ = v48
                                _ = v49
                            } else {
                                _ = len(v1)
                            }
                        }
                    }
                }
                switch (v2 + v2) {
                case 42:
                    for fi50 := 0; fi50 < 4; fi50++ {
                        for fi51 := 0; fi51 < 18; fi51++ {
                            switch -496 {
                            case 0:
                                v52 := GenS6{F0: fi51, F1: 2}
                                _ = v52
                            case 42:
                                switch (("\n\t" + "a") + ("" + (v23 + ("\n\t" + v23)))) {
                                case "a":
                                    v53 := GenS4{F0: (false || false), F1: (v19 + 4294967295), F2: 42}
                                    _ = v53
                                default:
                                    v54 := GenS7{F0: 107, F1: 77}
                                    _ = v54
                                }
                            case 2147483647:
                                var v55 string = ""
                                _ = v55
                            default:
                                v56 := GenS3{F0: 1, F1: "a"}
                                _ = v56
                            }
                            v19 += (v19 >> ((18446744073709551615) & 63))
                            v4 -= -782
                            _ = fi51
                        }
                        v24 = append(v24, -9223372036854775808)
                        var v57 string = (v23 + v23)
                        _ = fi50
                        _ = v57
                    }
                case 128:
                    for fi58 := 0; fi58 < 10; fi58++ {
                        v3 += (^(v3))
                        for fi59 := 0; fi59 < 14; fi59++ {
                            delete(v1, 466)
                            fmt.Println(v13, fi59)
                            if ((v4 - (v4 % (((-(v4))) | 1))) > (v4 ^ 301)) {
                                v4 += (v4 % ((v4) | 1))
                            }
                            var v60 uint64 = v19
                            _ = fi59
                            _ = v60
                        }
                        v13 += 2147483648
                        _ = fi58
                    }
                }
                if true {
                    if (0 == 42) {
                        fmt.Println(v19)
                        var v61 float32 = (-2.25 + 1.5)
                        _ = v61
                    } else {
                        v62 := GenS8{F0: (v26 << ((671) & 63)), F1: 100}
                        var v63 uint64 = v19
                        if true {
                            for fi64 := 0; fi64 < 4; fi64++ {
                                switch int(v26) {
                                case 127:
                                    var v65 int64 = 0
                                    _ = v65
                                default:
                                    _ = len(v1)
                                }
                                v4 = (-(v4))
                                var v66 float64 = 1.5
                                _ = fi64
                                _ = v66
                            }
                            var v67 uint = 0
                            fmt.Println(v63, v67, v23)
                            _ = v67
                        } else {
                            v23 = "hello"
                        }
                        _ = v62
                        _ = v63
                    }
                }
            } else {
                switch v23 {
                case "a":
                    delete(v1, (v13 % ((v13) | 1)))
                default:
                    switch v23 {
                    case "foo bar":
                        v68 := GenS4{F0: (v19 != v19), F1: (v19 + v19), F2: 387}
                        _ = v68
                    case "":
                        _ = len(v24)
                    }
                }
                fmt.Println(v13, v23, v4)
                var v69 int64 = -9223372036854775808
                _ = v69
            }
        default:
            switch "\n\t" {
            case "a":
                v70 := GenS4{F0: false, F1: uint64(v13), F2: 32767}
                _ = v70
            case "":
                var v71 int = v2
                _ = v71
            default:
                v23 = v23
            }
        }
        var v72 uint = 72
        _ = v72
    }
    if false {
        for fi73 := 0; fi73 < 2; fi73++ {
            v13 = v13
            if false {
                v74 := GenS7{F0: 127, F1: 252}
                _ = v74
            }
            _ = fi73
        }
        v13 = 2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
