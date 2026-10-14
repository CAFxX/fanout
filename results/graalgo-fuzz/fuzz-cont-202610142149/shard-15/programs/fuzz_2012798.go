package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 int; F2 uint8 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 uint; F1 uint16; F2 bool }

func main() {
    v1 := []int{-226, 32768, 4294967295}
    var v2 int16 = 42
    var v3 int8 = 1
    var v4 string = ((("hello" + "") + "\n\t") + "hello")
    v1 = append(v1, -336)
    if false {
        v1 = append(v1, 4294967295)
        switch (((v4 + (v4 + ("" + (v4 + v4)))) + v4) + ((v4 + "\n\t") + v4)) {
        case "\n\t":
            v2 += (v2 / (((v2 >> ((255) & 63))) | 1))
        case "a":
            var v5 uint8 = 2
            _ = v5
        case "foo bar":
            var v6 int = 2147483647
            _ = v6
        default:
            fmt.Println(v2, v3, v3)
        }
    }
    v2 -= v2
    v3 -= 100
    var v7 uint32 = 2147483647
    v4 = v4
    v4 = v4
    switch -2147483648 {
    case -2147483648:
        switch v4 {
        case "":
            var v8 bool = (-1 > 239)
            _ = v8
        case "a":
            if (!true) {
                var v9 bool = true
                var v10 string = ("a" + v4)
                _ = v10
                _ = v9
            }
        case "\n\t":
            v4 = ""
        default:
            v11 := GenS1{F0: -1, F1: 0, F2: 587}
            _ = v11
        }
    case 0:
        var v12 float32 = ((3.14159 + 1.5) * float32(0.0))
        _ = v12
    case 128:
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint8 = 1
            for fi15 := 0; fi15 < 6; fi15++ {
                v1 = append(v1, fi13)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    default:
        if (0.0 <= 0.1) {
            v2 = (v2 | (^(v2)))
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS2{F0: 409, F1: (v2 % ((v2) | 1)), F2: fi16}
                _ = fi16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        if true {
            var v19 bool = (1 <= 42)
            if len(v1) > 0 {
                v1[0] = (fi18 + (fi18 & fi18))
            }
            if (42 <= 255) {
                switch ("foo bar" + (v4 + v4)) {
                case "foo bar":
                    var v20 int64 = 92
                    _ = v20
                case "a":
                    switch v4 {
                    case "foo bar":
                        var v21 int = fi18
                        _ = v21
                    }
                case "hello":
                    v22 := GenS3{F0: v7}
                    _ = v22
                }
                switch v4 {
                case "a":
                    var v23 uint32 = v7
                    _ = v23
                default:
                    if v19 {
                        if v19 {
                            _ = len(v1)
                            _ = len(v1)
                            if ((!v19) && v19) {
                                var v24 int8 = (v3 | (-(v3)))
                                switch fi18 {
                                case -1:
                                    if false {
                                        var v25 int = -1
                                        _ = v25
                                    }
                                case -623:
                                    fmt.Println(v4, v24, v19)
                                case 1:
                                    if true {
                                        v2 += (v2 * v2)
                                        fmt.Println(v4, v24)
                                        v3 = -128
                                    } else {
                                        if (!false) {
                                            v24 = v24
                                            v26 := GenS4{F0: 1000, F1: -32768}
                                            _ = v26
                                        }
                                        v27 := GenS5{F0: 2147483647, F1: v3}
                                        _ = v27
                                    }
                                default:
                                    if v19 {
                                        v4 = v4
                                        fmt.Println(v19)
                                        for fi28 := 0; fi28 < 1; fi28++ {
                                            v24 = -128
                                            _ = fi28
                                        }
                                    } else {
                                        v3 = 100
                                        v3 += v24
                                        for fi29 := 0; fi29 < 17; fi29++ {
                                            var v30 uint64 = 11
                                            var v31 int = fi18
                                            if v19 {
                                                v32 := GenS6{F0: (v30 * v30)}
                                                switch ("hello" + ((v4 + v4) + (((((v4 + ("foo bar" + ("a" + "foo bar"))) + v4) + "foo bar") + ((v4 + v4) + "hello")) + v4))) {
                                                case "\n\t":
                                                    var v33 int = (fi18 % (((fi29 + (v31 + fi29))) | 1))
                                                    _ = v33
                                                case "":
                                                    var v34 float64 = 0.0
                                                    _ = v34
                                                case "hello":
                                                    switch fi18 {
                                                    case 0:
                                                        var v35 int32 = 0
                                                        _ = v35
                                                    case -2147483648:
                                                        var v36 uint64 = v30
                                                        _ = v36
                                                    default:
                                                        var v37 string = v4
                                                        _ = v37
                                                    }
                                                default:
                                                    if (68 < 255) {
                                                        v19 = (v7 <= v7)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    } else {
                                                        v3 -= (v24 | (v24 / ((v3) | 1)))
                                                        var v38 uint8 = 2
                                                        _ = v38
                                                    }
                                                }
                                                _ = v32
                                            }
                                            v39 := GenS4{F0: v2, F1: 1000}
                                            v30 = 667
                                            _ = fi29
                                            _ = v30
                                            _ = v31
                                            _ = v39
                                        }
                                    }
                                }
                                _ = v24
                            } else {
                                v19 = (3.14159 > 3.14159)
                                fmt.Println(v7, v4)
                            }
                        }
                        var v40 int64 = 9223372036854775807
                        v41 := GenS6{F0: 18446744073709551615}
                        _ = v40
                        _ = v41
                    } else {
                        v19 = true
                    }
                }
            } else {
                v3 -= v3
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v19
        } else {
            v7 -= v7
            v3 += v3
            v3 = -1
        }
        _ = fi18
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        if ((0 == 1) && (true || false)) {
            for fi43 := 0; fi43 < 17; fi43++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi43
            }
            for fi44 := 0; fi44 < 5; fi44++ {
                if true {
                    v7 -= 42
                } else {
                    _ = len(v1)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if true {
                    v2 = 0
                    if false {
                        v4 = (v4 + v4)
                        v45 := GenS7{F0: 42, F1: (fi42 & (fi42 - (fi42 * fi42))), F2: 127}
                        var v46 int16 = (v2 % ((v2) | 1))
                        if true {
                            if (true && (2 == 42)) {
                                switch "\n\t" {
                                case "foo bar":
                                    v46 = v46
                                case "\n\t":
                                    v2 += (v2 / (((v46 << ((2) & 63))) | 1))
                                default:
                                    var v47 uint16 = 707
                                    _ = v47
                                }
                                v3 -= v3
                                var v48 bool = false
                                _ = v48
                            } else {
                                var v49 string = (v4 + "foo bar")
                                var v50 int64 = -755
                                v3 -= 1
                                v3 -= v3
                                var v51 int8 = (-(v3))
                                fmt.Println(v49)
                                v52 := GenS2{F0: 480, F1: (v46 >> ((1) & 63)), F2: fi42}
                                var v53 uint = 498
                                if len(v1) > 0 {
                                    v1[0] = (^(fi44))
                                }
                                var v54 uint32 = (^(v7))
                                fmt.Println(v3)
                                v3 = v3
                                _ = v49
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        }
                        var v55 int32 = -2147483648
                        _ = v45
                        _ = v46
                        _ = v55
                    }
                    if false {
                        fmt.Println(fi42, v7, fi44)
                        v2 -= (v2 << ((6) & 63))
                        v7 -= (v7 >> ((0) & 63))
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 777
                }
                _ = fi44
            }
        } else {
            if len(v1) > 0 {
                v1[0] = fi42
            }
            var v56 float64 = float64(0.0)
            _ = v56
        }
        for fi57 := 0; fi57 < 12; fi57++ {
            switch ("hello" + v4) {
            case "a":
                if len(v1) > 0 {
                    v1[0] = fi42
                }
            case "hello":
                v3 = v3
            default:
                v58 := GenS4{F0: (v2 << ((255) & 63)), F1: 0}
                _ = v58
            }
            v2 -= v2
            switch (("" + ((((v4 + v4) + "foo bar") + v4) + ("hello" + ""))) + (v4 + "a")) {
            case "foo bar":
                if true {
                    var v59 int = (fi57 + 521)
                    v60 := GenS3{F0: v7}
                    _ = v59
                    _ = v60
                } else {
                    fmt.Println(v4, fi57, fi57)
                }
            case "\n\t":
                if (v4 > (v4 + "a")) {
                    v4 = ""
                    fmt.Println(v7)
                    var v61 int16 = v2
                    _ = v61
                } else {
                    v2 += v2
                    fmt.Println(fi57, v3, v2)
                }
            default:
                v62 := GenS1{F0: -3, F1: 18446744073709551615, F2: 0}
                _ = v62
            }
            _ = fi57
        }
        v2 += v2
        _ = fi42
    }
    v2 -= (v2 / ((v2) | 1))
    var v63 float32 = 1.5
    for fi64 := 0; fi64 < 10; fi64++ {
        if (("hello" + ("" + "foo bar")) == "") {
            v65 := GenS8{F0: v2, F1: v2}
            v66 := GenS7{F0: 0, F1: fi64, F2: 42}
            _ = v65
            _ = v66
        }
        _ = fi64
    }
    var v67 int32 = -464
    var v68 int16 = v2
    fmt.Println(v67, v7, v67)
    switch "" {
    case "a":
        for fi69 := 0; fi69 < 20; fi69++ {
            for fi70 := 0; fi70 < 14; fi70++ {
                var v71 uint32 = v7
                _ = fi70
                _ = v71
            }
            var v72 int8 = (v3 << ((97) & 63))
            for fi73 := 0; fi73 < 17; fi73++ {
                var v74 uint = 255
                if len(v1) > 0 {
                    v1[0] = (fi73 ^ int(v7))
                }
                v7 -= (v7 ^ 1)
                _ = fi73
                _ = v74
            }
            _ = fi69
            _ = v72
        }
    case "hello":
        for fi75 := 0; fi75 < 5; fi75++ {
            if (0 >= 846) {
                v76 := GenS9{F0: 65535, F1: 848, F2: false}
                _ = v76
            }
            v68 += (v2 / ((32767) | 1))
            if false {
                v7 -= (v7 << ((18446744073709551615) & 63))
            }
            _ = fi75
        }
    default:
        fmt.Println(v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v67
    _ = v68
    _ = v7
}
