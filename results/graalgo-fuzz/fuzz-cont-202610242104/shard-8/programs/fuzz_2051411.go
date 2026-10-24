package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 string; F1 uint }

func main() {
    var v1 float32 = (0.0 / (float32(3.14159) - 0.1))
    v2 := make(map[string]int)
    var v3 bool = (249 > 0)
    fmt.Println(v1)
    switch 128 {
    case -858:
        var v4 float32 = ((v1 - 0.1) * (v1 - 0.0))
        _ = v4
    case 235:
        if (float64(-2.25) <= 1e10) {
            var v5 uint64 = 379
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int = (fi6 << ((1) & 63))
        v8 := GenS1{F0: 666, F1: 378}
        if (128 == 2) {
            v9 := []int{2147483647, fi6, 829, (v7 | v7)}
            fmt.Println(v7, v7)
            fmt.Println(v7, fi6, v7)
            _ = v9
        } else {
            var v10 float32 = (v1 + v1)
            v11 := []int{-2147483648}
            if v3 {
                switch "hello" {
                case "\n\t":
                    v12 := GenS1{F0: 1, F1: 0}
                    _ = v12
                case "foo bar":
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v2[""] = fi13
                        _ = fi13
                    }
                default:
                    fmt.Println(v7, v1)
                }
                v1 = v10
            }
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 128 {
    case -496:
        if (42 == 113) {
            v3 = (127 > -118)
            if v3 {
                fmt.Println(v1)
            }
        }
    default:
        v2["hello"] = 127
    }
    v3 = (!(!((!(100 < 63)) && v3)))
    switch -547 {
    case 32768:
        v1 = v1
    case 1:
        var v14 uint = 4294967295
        _ = v14
    default:
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                v3 = ((true || ((v3 || v3) && (1.5 <= float64(((v1 * 1e10) + ((3.14159 + (float32((0.0 + 0.0)) / (v1 * v1))) - (((v1 / v1) + v1) + -2.25))))))) && true)
                v17 := GenS2{F0: false, F1: 128, F2: 127}
                var v18 uint16 = 758
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := GenS3{F0: 23}
            for fi20 := 0; fi20 < 18; fi20++ {
                switch "a" {
                case "hello":
                    v21 := []int{-9223372036854775808, 32768, (fi20 ^ (-(fi15))), fi15}
                    _ = v21
                case "":
                    v3 = v3
                }
                v1 += v1
                var v22 int8 = -128
                _ = fi20
                _ = v22
            }
            _ = fi15
            _ = v19
        }
    }
    v3 = false
    if false {
        delete(v2, "hello")
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS4{F0: 2, F1: 9223372036854775807, F2: 98}
            for fi25 := 0; fi25 < 7; fi25++ {
                v26 := []int{}
                switch (-(fi23)) {
                case 1:
                    var v27 int8 = -87
                    _ = v27
                default:
                    switch ("" + "\n\t") {
                    case "hello":
                        v3 = true
                    default:
                        _ = len(v26)
                    }
                }
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = fi25
                _ = v26
            }
            for fi28 := 0; fi28 < 19; fi28++ {
                if true {
                    v29 := []int{-302, (fi23 << ((0) & 63)), (fi23 - (-(fi23))), 2147483647}
                    if v3 {
                        var v30 float32 = v1
                        _ = v30
                    } else {
                        v3 = ((((float32((-2.25 + 0.1)) + 1e10) * v1) + (-2.25 * 1e10)) > v1)
                        if v3 {
                            v3 = ((("" + "hello") < "foo bar") && (v3 || true))
                        } else {
                            _ = len(v2)
                            v2[("\n\t" + "")] = fi23
                            v3 = v3
                            var v31 uint32 = 2
                            var v32 int8 = -1
                            _ = v31
                            _ = v32
                        }
                    }
                    _ = v29
                } else {
                    v33 := GenS2{F0: v3, F1: 1, F2: 0}
                    v34 := GenS4{F0: 42, F1: 9223372036854775807, F2: 0}
                    _ = v33
                    _ = v34
                }
                _ = fi28
            }
            _ = fi23
            _ = v24
        }
    } else {
        switch 2147483647 {
        case 833:
            v1 = ((v1 - (v1 / 1.5)) * (v1 * (v1 * 0.0)))
        case 1:
            v3 = false
        }
        fmt.Println(v1, v1)
        for fi35 := 0; fi35 < 13; fi35++ {
            for fi36 := 0; fi36 < 9; fi36++ {
                v3 = ((!((-2.25 > v1) || (fi36 != fi36))) && (127 >= 1))
                switch "foo bar" {
                case "":
                    switch 32767 {
                    case -9223372036854775808:
                        v3 = false
                    case -2147483648:
                        var v37 int32 = 1
                        _ = v37
                    default:
                        var v38 int8 = 127
                        _ = v38
                    }
                case "foo bar":
                    v39 := []int{fi36, fi36, fi35}
                    _ = v39
                default:
                    v1 -= v1
                }
                _ = fi36
            }
            switch "" {
            case "hello":
                v40 := []int{}
                _ = v40
            case "":
                var v41 uint16 = 760
                _ = v41
            }
            fmt.Println(v3, v1)
            _ = fi35
        }
    }
    fmt.Println(v1)
    v42 := GenS1{F0: 4294967295, F1: 759}
    fmt.Println(v1, v1)
    v43 := GenS2{F0: (65535 <= 65535), F1: 126, F2: 1}
    v44 := GenS5{F0: 1, F1: 18446744073709551615}
    for fi45 := 0; fi45 < 6; fi45++ {
        if false {
            fmt.Println(v3, v3, fi45)
            if v3 {
                fmt.Println(v1)
            } else {
                var v46 string = "\n\t"
                if ((0 != 255) || (666 > 0)) {
                    fmt.Println(v3, v3, v3)
                    for fi47 := 0; fi47 < 20; fi47++ {
                        v48 := v2[(("" + (v46 + v46)) + v46)]
                        if (127 == -128) {
                            var v49 int16 = 0
                            v49 = -173
                            switch (("" + v46) + v46) {
                            case "a":
                                var v50 string = v46
                                _ = v50
                            }
                            _ = v49
                        } else {
                            v51 := GenS3{F0: 491}
                            var v52 int8 = -108
                            v53 := GenS6{F0: 42, F1: 18446744073709551615}
                            if (42 == 828) {
                                if v3 {
                                    switch v46 {
                                    case "hello":
                                        _ = len(v2)
                                    }
                                }
                                v54 := []int{fi47, (v48 / ((fi47) | 1))}
                                var v55 uint16 = 32767
                                v56 := GenS4{F0: 642, F1: -1, F2: 863}
                                if v3 {
                                    var v57 string = (v46 + (v46 + ""))
                                    if v3 {
                                        var v58 int = (^(fi47))
                                        var v59 string = ("" + (v46 + v46))
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v57
                                }
                                v60 := GenS2{F0: false, F1: 142, F2: (v52 / ((42) | 1))}
                                if (!(true && v3)) {
                                    v61 := GenS4{F0: 0, F1: 452, F2: 1}
                                    switch (fi45 & fi45) {
                                    case 9223372036854775807:
                                        fmt.Println(v3, v46, v1)
                                    case 4294967295:
                                        v62 := GenS4{F0: 31, F1: -9223372036854775808, F2: 509}
                                        _ = v62
                                    case 127:
                                        v55 = (v55 / ((v55) | 1))
                                    default:
                                        var v63 int8 = (-(v52))
                                        _ = v63
                                    }
                                    _ = v61
                                }
                                _ = v54
                                _ = v55
                                _ = v56
                                _ = v60
                            }
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        v1 = (v1 * float32((3.14159 + 3.14159)))
                        _ = fi47
                        _ = v48
                    }
                    switch (("hello" + ((v46 + v46) + "a")) + v46) {
                    case "hello":
                        v64 := GenS6{F0: 828, F1: 18446744073709551615}
                        _ = v64
                    case "\n\t":
                        if (v3 && (3.14159 == -2.25)) {
                            var v65 uint32 = 1
                            var v66 uint64 = 814
                            if (-1 <= -2147483648) {
                                var v67 float32 = (v1 - v1)
                                var v68 string = ((v46 + "foo bar") + v46)
                                if (float32((1e10 + -2.25)) <= v67) {
                                    for fi69 := 0; fi69 < 11; fi69++ {
                                        v70 := GenS7{F0: v46, F1: 255}
                                        var v71 int16 = 1000
                                        v72 := GenS5{F0: 2147483647, F1: 103}
                                        _ = fi69
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    for fi73 := 0; fi73 < 7; fi73++ {
                                        if (v68 > v68) {
                                            delete(v2, ("hello" + v68))
                                            v74 := []int{(fi73 + fi45)}
                                            fmt.Println(v67)
                                            _ = v74
                                        }
                                        v75 := GenS6{F0: 30, F1: 42}
                                        _ = fi73
                                        _ = v75
                                    }
                                }
                                _ = v67
                                _ = v68
                            } else {
                                var v76 string = ((v46 + ((v46 + "a") + (v46 + v46))) + (("hello" + v46) + v46))
                                var v77 uint = 4294967295
                                v78 := []int{(fi45 / (((fi45 ^ (fi45 + fi45))) | 1))}
                                _ = v76
                                _ = v77
                                _ = v78
                            }
                            _ = v65
                            _ = v66
                        }
                    default:
                        v46 = "a"
                    }
                } else {
                    var v79 int8 = -128
                    v79 = (v79 & (v79 ^ (v79 | -128)))
                    v46 = v46
                    _ = v79
                }
                _ = v46
            }
            v80 := GenS5{F0: 42, F1: 65535}
            _ = v80
        }
        v1 -= 0.1
        _ = fi45
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v44
}
