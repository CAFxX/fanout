package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0(p0 uint) string {
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 bool = false
            v10 := GenS1{F0: fi8}
            fmt.Println(v6, v6, p0)
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            if ((((!true) && true) || false) && ((v6 - 234) == v6)) {
                p0 -= p0
                p0 += (p0 / ((p0) | 1))
                v12 := []int{}
                _ = v12
            }
            _ = fi11
        }
        _ = fi7
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return ""
}
func genFn1() uint32 {
    var v13 float32 = 0.0
    v13 = v13
    if (!false) {
        v14 := GenS1{F0: -224}
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 uint16 = 42
        var v17 uint16 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v13
    return 2147483647
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 0
    var v3 uint64 = 4294967295
    var v4 int64 = (v2 * 9223372036854775807)
    v5 := make(map[int]string)
    switch "foo bar" {
    case "\n\t":
        switch 42 {
        case 0:
            fmt.Println(v1)
        case 348:
            if (4294967295 >= 0) {
                if (942 < 65535) {
                    var v18 uint64 = v3
                    _ = v18
                }
                v19 := GenS1{F0: 1}
                if true {
                    switch -517 {
                    case -1:
                        v2 -= (-(v4))
                    default:
                        v1 = 489
                    }
                    var v20 uint = 42
                    if false {
                        v2 = v2
                    } else {
                        fmt.Println(v20, v2)
                        v2 = 662
                    }
                    _ = v20
                } else {
                    var v21 int8 = -1
                    for fi22 := 0; fi22 < 1; fi22++ {
                        if (32767 != 2) {
                            fmt.Println(v2, v1, fi22)
                        } else {
                            if true {
                                var v23 uint16 = 308
                                v24 := v5[9223372036854775807]
                                _ = v23
                                _ = v24
                            } else {
                                v3 -= v3
                                var v25 uint = 609
                                switch "hello" {
                                case "foo bar":
                                    v26 := v5[32768]
                                    _ = v26
                                case "":
                                    v25 += v25
                                case "\n\t":
                                    v27 := []int{}
                                    _ = v27
                                default:
                                    fmt.Println(fi22)
                                }
                                _ = v25
                            }
                            for fi28 := 0; fi28 < 14; fi28++ {
                                var v29 uint32 = genFn1()
                                var v30 string = ("a" + "foo bar")
                                if (!((v21 ^ (v21 / ((v21) | 1))) != (v21 << ((0) & 63)))) {
                                    v1 += v1
                                } else {
                                    var v31 uint32 = v29
                                    var v32 uint16 = 0
                                    if true {
                                        fmt.Println(fi28, v32)
                                        if (-2.25 >= 1.5) {
                                            var v33 int8 = -84
                                            _ = v33
                                        } else {
                                            if true {
                                                delete(v5, (^(fi28)))
                                                var v34 int16 = 32767
                                                _ = v34
                                            }
                                            fmt.Println(v29, v1, v31)
                                            if false {
                                                v31 = v29
                                                v32 -= (^(v32))
                                            }
                                        }
                                    } else {
                                        fmt.Println(v32)
                                    }
                                    _ = v31
                                    _ = v32
                                }
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            for fi35 := 0; fi35 < 14; fi35++ {
                                switch (fi22 ^ fi22) {
                                case -916:
                                    v1 += (v1 * genFn1())
                                }
                                v4 -= -1
                                _ = fi35
                            }
                        }
                        _ = fi22
                    }
                    _ = v21
                }
                _ = v19
            } else {
                if (1 < 32767) {
                    switch "\n\t" {
                    case "\n\t":
                        v36 := GenS1{F0: 1}
                        _ = v36
                    case "":
                        switch 42 {
                        case 32767:
                            var v37 int8 = 101
                            _ = v37
                        case 950:
                            if (0.0 >= -2.25) {
                                var v38 float32 = -2.25
                                v1 = (v1 ^ (v1 << ((409) & 63)))
                                _ = v38
                            }
                        case -546:
                            v39 := GenS2{F0: 127, F1: 1, F2: 0}
                            _ = v39
                        default:
                            v40 := GenS3{F0: 127}
                            _ = v40
                        }
                    default:
                        delete(v5, 127)
                    }
                }
            }
        }
    }
    v41 := GenS2{F0: 1, F1: 255, F2: 0}
    v42 := GenS1{F0: 2147483647}
    if false {
        v43 := GenS4{F0: 0, F1: 26}
        v44 := GenS2{F0: 255, F1: 0, F2: uint16(v1)}
        _ = v43
        _ = v44
    }
    v45 := GenS5{F0: "hello", F1: (("foo bar" + "") + "hello")}
    var v46 bool = ((v2 | (v4 / (((v4 + (^(v4)))) | 1))) > 823)
    fmt.Println(v1, v46)
    if ((^(v3)) != 4294967295) {
        if (true && ((((!false) && v46) || (!(255 >= 1))) && v46)) {
            fmt.Println(v4, v46, v4)
            var v47 int = 2147483647
            _ = v47
        }
    } else {
        v48 := []int{626, 4294967295}
        for fi49 := 0; fi49 < 13; fi49++ {
            var v50 string = "a"
            if len(v48) >= 2 {
                v48 = v48[1:len(v48)]
            }
            switch (v50 + "") {
            case "hello":
                switch fi49 {
                case 5:
                    var v51 bool = ((337 < (^(v1))) && v46)
                    _ = v51
                case 708:
                    if (255 > 42) {
                        if v46 {
                            if (-612 == 2147483647) {
                                _ = len(v48)
                                v5[(fi49 & (fi49 / (((fi49 | (fi49 & (fi49 - (fi49 - fi49))))) | 1)))] = v50
                            } else {
                                fmt.Println(v50, v50, v46)
                                if v46 {
                                    for fi52 := 0; fi52 < 9; fi52++ {
                                        var v53 int8 = -30
                                        _ = fi52
                                        _ = v53
                                    }
                                    fmt.Println(fi49)
                                }
                            }
                            fmt.Println(v3, v46)
                        } else {
                            if (!(!v46)) {
                                var v54 int = fi49
                                fmt.Println(v50)
                                var v55 int = (v54 >> ((2) & 63))
                                _ = v54
                                _ = v55
                            } else {
                                var v56 uint16 = 144
                                if true {
                                    if v46 {
                                        _ = len(v48)
                                        v5[fi49] = (v50 + v50)
                                    }
                                    var v57 float64 = 0.1
                                    var v58 int8 = 1
                                    _ = v57
                                    _ = v58
                                }
                                switch fi49 {
                                case -354:
                                    v4 = int64(v1)
                                case 32767:
                                    v3 -= 873
                                }
                                _ = v56
                            }
                        }
                        if len(v48) >= 2 {
                            v48 = v48[1:len(v48)]
                        }
                        var v59 uint64 = 0
                        _ = v59
                    }
                default:
                    v50 = ("a" + ("foo bar" + ("" + "hello")))
                }
            case "":
                v48 = append(v48, fi49)
            }
            _ = fi49
            _ = v50
        }
        var v60 float64 = 1.5
        _ = v48
        _ = v60
    }
    v2 = v2
    v61 := []int{9223372036854775807, 9223372036854775807, 4294967295}
    var v62 uint = 65535
    fmt.Println(v46, v4)
    var v63 float64 = (1.5 * float64(float32(-2.25)))
    for fi64 := 0; fi64 < 17; fi64++ {
        if v46 {
            v1 += 353
            v3 = v3
            if v46 {
                if v46 {
                    for fi65 := 0; fi65 < 4; fi65++ {
                        fmt.Println(fi65, v63)
                        fmt.Println(v46, v3)
                        if false {
                            v4 -= 745
                            v66 := GenS4{F0: -1, F1: 95}
                            v67 := GenS5{F0: ("a" + ""), F1: "hello"}
                            switch "" {
                            case "foo bar":
                                v68 := GenS6{F0: v46, F1: v3, F2: -1}
                                _ = v68
                            case "":
                                var v69 float32 = (-2.25 * 0.1)
                                _ = v69
                            case "hello":
                                if true {
                                    v63 = v63
                                } else {
                                    var v70 uint8 = 233
                                    v71 := v5[(-(fi64))]
                                    _ = v70
                                    _ = v71
                                }
                            }
                            _ = v66
                            _ = v67
                        }
                        v3 = (^(v3))
                        v62 -= 255
                        if ((v46 && true) && v46) {
                            switch ("hello" + ("hello" + "\n\t")) {
                            case "a":
                                v72 := GenS3{F0: fi65}
                                _ = v72
                            }
                            var v73 float32 = 0.1
                            if len(v61) > 0 {
                                v61[0] = (fi64 << ((v62) & 63))
                            }
                            _ = v73
                        }
                        _ = fi65
                    }
                }
            }
        } else {
            v46 = (65535 > (v1 ^ 2))
            if (true || v46) {
                if v46 {
                    _ = len(v5)
                }
            } else {
                if ((!(v46 && ((v2 >> (((v62 ^ v62)) & 63)) < (v2 - (^(v2)))))) || ((^(fi64)) <= (fi64 - fi64))) {
                    fmt.Println(v46, v46)
                    v3 = (v3 << (((v62 >> ((v62) & 63))) & 63))
                }
                v1 -= (v1 >> (((v62 | 4294967295)) & 63))
                var v74 uint = uint(v1)
                _ = v74
            }
        }
        fmt.Println(v4)
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v5
    _ = v61
    _ = v62
    _ = v63
}
