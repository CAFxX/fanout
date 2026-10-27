package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS4 struct { F0 int64; F1 string; F2 int }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int16 }

func main() {
    v1 := []int{9223372036854775807, 32768, 127, 0}
    var v2 uint64 = 1
    var v3 uint = 65535
    var v4 string = "hello"
    v4 = "foo bar"
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint64 = v2
        for fi7 := 0; fi7 < 9; fi7++ {
            v6 = v6
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    v3 += v3
    switch 127 {
    case 0:
        switch -2147483648 {
        case 894:
            v8 := GenS1{F0: v3, F1: 383, F2: 184}
            _ = v8
        case 509:
            var v9 int64 = -9223372036854775808
            _ = v9
        case -9223372036854775808:
            var v10 uint64 = 2
            _ = v10
        default:
            fmt.Println(v4, v4, v2)
        }
    case 32768:
        v3 += v3
    case -1:
        v2 -= (v2 & (v2 & (v2 / (((v2 + (v2 % (((^(v2))) | 1)))) | 1))))
    default:
        v4 = "\n\t"
    }
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            switch 42 {
            case 32767:
                if len(v1) > 0 {
                    v1[0] = fi11
                }
            case 0:
                if false {
                    var v12 int16 = -1
                    fmt.Println(fi11, v2)
                    fmt.Println(v12)
                    _ = v12
                }
            }
            v13 := GenS1{F0: (v3 ^ (v3 >> ((116) & 63))), F1: 4294967295, F2: 1000}
            _ = fi11
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 11; fi14++ {
            if true {
                var v15 uint64 = 1
                fmt.Println(v2)
                _ = v15
            }
            if (-128 > 1) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v16 int16 = 977
                for fi17 := 0; fi17 < 7; fi17++ {
                    if (v4 < v4) {
                        if true {
                            if ((fi17 | fi14) == 2147483647) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch "hello" {
                                case "foo bar":
                                    v1 = append(v1, (-(fi17)))
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = fi17
                                    }
                                case "a":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    fmt.Println(fi17)
                                }
                                v16 = (-(v16))
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi17 ^ (fi14 % ((fi14) | 1)))
                            }
                        } else {
                            var v18 uint = v3
                            _ = v18
                        }
                    } else {
                        v4 = "\n\t"
                        v2 = v2
                    }
                    v19 := GenS2{F0: 127}
                    _ = len(v1)
                    _ = len(v1)
                    fmt.Println(v16)
                    _ = fi17
                    _ = v19
                }
                _ = v16
            }
            _ = fi14
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            switch (fi20 & (fi20 >> (((v3 / (((v3 >> ((v3) & 63))) | 1))) & 63))) {
            case -2147483648:
                v4 = ((v4 + v4) + v4)
            case -1:
                if (2147483648 <= 65535) {
                    switch fi20 {
                    case 127:
                        v1 = append(v1, fi20)
                    case 42:
                        v2 += 42
                    default:
                        for fi21 := 0; fi21 < 12; fi21++ {
                            v22 := GenS2{F0: 100}
                            _ = len(v1)
                            _ = fi21
                            _ = v22
                        }
                    }
                }
            case 0:
                fmt.Println(v4, v4, fi20)
            }
            if ((((("foo bar" + "hello") + "") + (v4 + (v4 + v4))) < "a") && false) {
                for fi23 := 0; fi23 < 6; fi23++ {
                    fmt.Println(v4, v4, fi23)
                    if len(v1) > 0 {
                        v1[0] = (fi20 << ((v3) & 63))
                    }
                    if (true && false) {
                        v4 = ""
                        var v24 string = (v4 + "\n\t")
                        _ = v24
                    }
                    _ = fi23
                }
                v2 -= v2
            } else {
                if ((fi20 * fi20) <= fi20) {
                    v2 += (^(v2))
                } else {
                    var v25 float32 = 1.5
                    for fi26 := 0; fi26 < 11; fi26++ {
                        v27 := GenS3{F0: v2, F1: -743, F2: 124}
                        _ = fi26
                        _ = v27
                    }
                    if len(v1) > 0 {
                        v1[0] = -801
                    }
                    _ = v25
                }
                v4 = v4
            }
            _ = fi20
        }
        var v28 uint = uint(v2)
        _ = v28
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        if (true && false) {
            var v30 string = (v4 + v4)
            if ((fi29 / ((fi29) | 1)) > 9223372036854775807) {
                v31 := GenS2{F0: -128}
                _ = v31
            } else {
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 uint8 = 197
                    _ = len(v1)
                    _ = fi32
                    _ = v33
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v34 := GenS3{F0: (^(v2)), F1: 32767, F2: 2}
            _ = v30
            _ = v34
        }
        if false {
            var v35 uint = v3
            v4 = "foo bar"
            v36 := GenS4{F0: 42, F1: v4, F2: (fi29 * -149)}
            _ = v35
            _ = v36
        } else {
            v37 := GenS5{F0: ("foo bar" + ("a" + v4)), F1: v4}
            for fi38 := 0; fi38 < 13; fi38++ {
                fmt.Println(fi29, v2, v4)
                switch 1 {
                case 2147483647:
                    var v39 uint = (v3 / ((v3) | 1))
                    _ = v39
                }
                _ = fi38
            }
            _ = v37
        }
        _ = fi29
    }
    if (42 > 51) {
        fmt.Println(v4, v2)
    } else {
        v4 = v4
        var v40 uint = (v3 << (((v3 & v3)) & 63))
        v4 = v4
        _ = v40
    }
    var v41 float32 = -2.25
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 int8 = 0
        v4 = "\n\t"
        _ = fi42
        _ = v43
    }
    fmt.Println(v4, v41)
    if false {
        v44 := GenS6{F0: 577}
        _ = v44
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            if (!false) {
                v41 = 3.14159
                fmt.Println(v4, v2, v4)
            } else {
                v41 += v41
            }
            v2 = v2
            var v46 string = "foo bar"
            _ = fi45
            _ = v46
        }
        var v47 uint32 = 756
        var v48 int64 = 0
        _ = v47
        _ = v48
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        var v50 int8 = 0
        fmt.Println(v3)
        _ = fi49
        _ = v50
    }
    _ = len(v1)
    fmt.Println(v41)
    v4 = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
}
