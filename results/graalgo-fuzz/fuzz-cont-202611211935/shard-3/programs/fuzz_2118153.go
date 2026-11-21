package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
type GenS4 struct { F0 string; F1 int; F2 int64 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 uint; F1 int16; F2 uint64 }
type GenS10 struct { F0 int64 }

func main() {
    var v1 uint = 42
    v2 := []int{127, 32767}
    var v3 uint32 = 858
    v4 := []string{"a", "\n\t", ""}
    v5 := []string{((("hello" + "\n\t") + "a") + "a"), (((("\n\t" + "foo bar") + ("hello" + "\n\t")) + "\n\t") + "\n\t")}
    for fi6 := 0; fi6 < 1; fi6++ {
        if (3.14159 < (0.0 - float64(0.0))) {
            var v7 uint16 = 429
            v7 += 255
            _ = v7
        } else {
            if false {
                for fi8 := 0; fi8 < 14; fi8++ {
                    switch ("" + "a") {
                    case "\n\t":
                        var v9 string = "\n\t"
                        _ = v9
                    case "":
                        if true {
                            v10 := GenS1{F0: 2, F1: 1}
                            v11 := GenS2{F0: 0, F1: (v3 * (v3 ^ (v3 << ((v1) & 63)))), F2: int64(v3)}
                            _ = v10
                            _ = v11
                        }
                    case "hello":
                        v4 = append(v4, "\n\t")
                    }
                    _ = len(v5)
                    v1 -= 577
                    _ = fi8
                }
                _ = len(v2)
            }
            if false {
                if true {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                        _ = fi12
                    }
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v1 = (v1 ^ v1)
                        fmt.Println(fi6, fi13)
                        v5 = append(v5, "foo bar")
                        _ = fi13
                    }
                    var v14 float64 = ((0.0 * 3.14159) - float64(3.14159))
                    _ = v14
                } else {
                    if true {
                        switch "\n\t" {
                        case "foo bar":
                            v1 += (v1 % (((v1 >> ((v1) & 63))) | 1))
                        case "\n\t":
                            v3 = 528
                        case "hello":
                            v15 := GenS2{F0: 42, F1: 4294967295, F2: -1}
                            _ = v15
                        default:
                            v16 := GenS3{F0: fi6, F1: 0, F2: 1}
                            _ = v16
                        }
                        v1 -= (^(v1))
                    } else {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                v2 = append(v2, 128)
                            }
                            _ = fi17
                        }
                        switch "\n\t" {
                        case "":
                            v3 = v3
                        case "\n\t":
                            v18 := GenS4{F0: ("a" + ""), F1: (fi6 - (fi6 ^ (fi6 + (fi6 % (((fi6 + (fi6 | (fi6 >> ((uint(v3)) & 63))))) | 1))))), F2: -181}
                            _ = v18
                        case "foo bar":
                            for fi19 := 0; fi19 < 12; fi19++ {
                                v20 := GenS5{F0: 664, F1: false}
                                var v21 bool = false
                                switch "" {
                                case "a":
                                    v21 = true
                                default:
                                    var v22 uint = (v1 - v1)
                                    _ = v22
                                }
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                        default:
                            if (1 == 42) {
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    var v24 int16 = -1
                                    fmt.Println(v24)
                                    _ = fi23
                                    _ = v24
                                }
                                v25 := GenS5{F0: 666, F1: ((v1 - 0) <= v1)}
                                _ = v25
                            } else {
                                var v26 uint8 = 91
                                _ = v26
                            }
                        }
                        for fi27 := 0; fi27 < 2; fi27++ {
                            fmt.Println(v3, v3)
                            _ = fi27
                        }
                    }
                    fmt.Println(v1, v3)
                    var v28 uint8 = 0
                    _ = v28
                }
            }
        }
        v5 = append(v5, ("\n\t" + "hello"))
        _ = fi6
    }
    var v29 uint16 = 0
    v3 -= 42
    if false {
        var v30 float64 = ((float64(0.1) / 1.0) * -2.25)
        _ = v30
    } else {
        v3 += v3
        for fi31 := 0; fi31 < 10; fi31++ {
            fmt.Println(v3, fi31, v3)
            var v32 uint64 = 1
            _ = fi31
            _ = v32
        }
    }
    switch "" {
    case "hello":
        if false {
            var v33 int8 = 127
            v34 := GenS2{F0: 117, F1: v3, F2: 0}
            _ = len(v5)
            _ = v33
            _ = v34
        } else {
            switch "a" {
            case "hello":
                v29 = uint16(v3)
            case "foo bar":
                if true {
                    switch "a" {
                    case "":
                        v3 = (^(v3))
                    case "hello":
                        var v35 uint32 = (v3 | v3)
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 17; fi36++ {
                        v1 = (v1 / ((601) | 1))
                        v37 := GenS3{F0: fi36, F1: 498, F2: -59}
                        _ = fi36
                        _ = v37
                    }
                } else {
                    var v38 int16 = 147
                    if false {
                        for fi39 := 0; fi39 < 16; fi39++ {
                            var v40 int = (fi39 + (-(fi39)))
                            _ = fi39
                            _ = v40
                        }
                        v41 := GenS2{F0: 127, F1: v3, F2: -1}
                        var v42 int = 2147483647
                        _ = v41
                        _ = v42
                    } else {
                        fmt.Println(v38, v29)
                        if ((v3 % ((uint32(v1)) | 1)) == (^(v3))) {
                            fmt.Println(v38, v1)
                        } else {
                            v43 := GenS6{F0: v29}
                            if true {
                                v29 = 1
                                if (false && true) {
                                    var v44 float32 = (0.0 + (1.5 / 1e10))
                                    v29 -= 488
                                    switch 42 {
                                    case 127:
                                        if false {
                                            v29 -= (v29 >> ((v1) & 63))
                                            if (1 > v1) {
                                                for fi45 := 0; fi45 < 4; fi45++ {
                                                    var v46 uint64 = 177
                                                    var v47 int16 = (v38 << ((4294967295) & 63))
                                                    v29 -= 2
                                                    _ = fi45
                                                    _ = v46
                                                    _ = v47
                                                }
                                                v1 = (v1 + v1)
                                                var v48 uint16 = (^(v29))
                                                _ = v48
                                            }
                                        } else {
                                            var v49 int64 = -1
                                            _ = v49
                                        }
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v44
                                } else {
                                    v1 -= v1
                                    v29 -= (v29 + (v29 | 410))
                                }
                                if true {
                                    for fi50 := 0; fi50 < 1; fi50++ {
                                        switch ("" + "\n\t") {
                                        case "a":
                                            if (-308 <= -2147483648) {
                                                for fi51 := 0; fi51 < 8; fi51++ {
                                                    switch (^(fi51)) {
                                                    case 842:
                                                        if false {
                                                            switch 32767 {
                                                            case 32768:
                                                                v52 := GenS7{F0: -2147483648}
                                                                _ = v52
                                                            case 4294967295:
                                                                var v53 float64 = 0.0
                                                                _ = v53
                                                            }
                                                        } else {
                                                            var v54 int8 = 100
                                                            fmt.Println(v29, v54)
                                                            var v55 uint32 = (^(v3))
                                                            if false {
                                                                fmt.Println(v1, v3)
                                                            } else {
                                                                var v56 uint = (v1 & v1)
                                                                if len(v5) > 0 {
                                                                    v5[0] = ("hello" + "foo bar")
                                                                }
                                                                var v57 bool = false
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            var v58 uint32 = v3
                                                            if (-128 <= v54) {
                                                                v38 = v38
                                                                fmt.Println(v55)
                                                            } else {
                                                                var v59 string = ("a" + ("hello" + ("\n\t" + "")))
                                                                _ = v59
                                                            }
                                                            v38 = (^(v38))
                                                            _ = v54
                                                            _ = v55
                                                            _ = v58
                                                        }
                                                    }
                                                    switch -546 {
                                                    case 2147483647:
                                                        var v60 uint = uint(v3)
                                                        _ = v60
                                                    default:
                                                        var v61 float32 = ((0.1 - 3.14159) / 1.0)
                                                        _ = v61
                                                    }
                                                    _ = fi51
                                                }
                                            }
                                        case "":
                                            switch "a" {
                                            case "":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            case "\n\t":
                                                v29 = v29
                                            case "hello":
                                                v62 := GenS6{F0: (v29 * 1)}
                                                _ = v62
                                            }
                                        case "foo bar":
                                            var v63 int32 = -1
                                            _ = v63
                                        }
                                        var v64 uint8 = 234
                                        v65 := GenS8{F0: 42, F1: v64}
                                        _ = fi50
                                        _ = v64
                                        _ = v65
                                    }
                                }
                            } else {
                                for fi66 := 0; fi66 < 10; fi66++ {
                                    v38 += v38
                                    var v67 float32 = float32(1.5)
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                    _ = fi66
                                    _ = v67
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v3 = (v3 % (((v3 % (((^(v3))) | 1))) | 1))
                            }
                            var v68 uint64 = 370
                            _ = v43
                            _ = v68
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v38
                }
            case "a":
                v1 += v1
            }
        }
    case "a":
        var v69 int16 = -1
        _ = v69
    case "\n\t":
        v29 = v29
    }
    v70 := GenS9{F0: v1, F1: 1, F2: 1}
    for fi71 := 0; fi71 < 7; fi71++ {
        for fi72 := 0; fi72 < 17; fi72++ {
            var v73 int16 = 1
            _ = fi72
            _ = v73
        }
        _ = fi71
    }
    v74 := GenS10{F0: -565}
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v74
}
