package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool) int16 {
    var v6 int = 32768
    v7 := []int{}
    _ = len(v7)
    _ = p0
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    v4 := []uint32{2}
    var v5 uint64 = 904
    for fi8 := 0; fi8 < 17; fi8++ {
        v4 = append(v4, 383)
        var v9 int8 = 100
        for fi10 := 0; fi10 < 20; fi10++ {
            if ((1.5 != float64((-2.25 + 0.1))) || (v9 < v9)) {
                v11 := GenS1{F0: -32768, F1: 65535, F2: v3}
                _ = v11
            }
            var v12 int16 = 1000
            _ = fi10
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v13 := v2[1]
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: 921, F1: 2, F2: v3}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: 2147483648, F2: 116}
    for fi17 := 0; fi17 < 7; fi17++ {
        switch "hello" {
        case "foo bar":
            for fi18 := 0; fi18 < 9; fi18++ {
                fmt.Println(v3, v13)
                var v19 uint64 = 187
                _ = fi18
                _ = v19
            }
        case "\n\t":
            fmt.Println(v13)
        case "hello":
            if ((v3 << ((396) & 63)) >= v3) {
                v13 = (v13 + ("" + v13))
                for fi20 := 0; fi20 < 6; fi20++ {
                    var v21 uint32 = 42
                    _ = fi20
                    _ = v21
                }
                _ = len(v2)
            }
        default:
            v5 = 916
        }
        for fi22 := 0; fi22 < 2; fi22++ {
            if (!((1.5 * 0.0) < -2.25)) {
                var v23 uint16 = 255
                var v24 int = (fi17 << ((1) & 63))
                delete(v2, (fi22 / (((fi17 & (v24 ^ (fi17 * 1)))) | 1)))
                _ = v23
                _ = v24
            }
            var v25 int8 = -57
            _ = fi22
            _ = v25
        }
        _ = fi17
    }
    var v26 uint32 = 499
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v26 = v26
    switch v13 {
    case "foo bar":
        if (v5 <= v5) {
            var v27 int32 = (v3 + v3)
            v28 := GenS1{F0: int16(v26), F1: 2147483647, F2: v3}
            if (int16(v5) == 1) {
                for fi29 := 0; fi29 < 2; fi29++ {
                    if ((fi29 & (fi29 % ((fi29) | 1))) < fi29) {
                        for fi30 := 0; fi30 < 1; fi30++ {
                            if (v27 > v27) {
                                var v31 int64 = 0
                                var v32 uint64 = (v5 & v5)
                                _ = v31
                                _ = v32
                            }
                            _ = fi30
                        }
                        v4 = append(v4, 1)
                    } else {
                        v26 += (^(v26))
                    }
                    var v33 uint16 = 65535
                    var v34 bool = (fi29 >= int(v27))
                    _ = fi29
                    _ = v33
                    _ = v34
                }
                fmt.Println(v27, v5)
                var v35 uint = 42
                _ = v35
            }
            _ = v27
            _ = v28
        } else {
            var v36 uint16 = 911
            _ = v36
        }
    case "hello":
        for fi37 := 0; fi37 < 3; fi37++ {
            if ((v26 % (((v26 / ((v26) | 1))) | 1)) != (v26 << ((1) & 63))) {
                v5 = (v5 & 18446744073709551615)
            } else {
                if false {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 string = "\n\t"
                        _ = fi38
                        _ = v39
                    }
                    _ = len(v2)
                    if true {
                        fmt.Println(v5, v13)
                        v5 -= (v5 % ((1) | 1))
                    }
                } else {
                    var v40 uint16 = 255
                    _ = v40
                }
                var v41 float64 = (float64(-2.25) - float64(1e10))
                for fi42 := 0; fi42 < 19; fi42++ {
                    v5 = 320
                    switch ((v13 + v13) + (v13 + v13)) {
                    case "\n\t":
                        if true {
                            v43 := v2[(fi37 + fi37)]
                            v44 := GenS2{F0: 0}
                            v41 = v41
                            v41 = ((3.14159 + v41) / 3.14159)
                            var v45 int = 4294967295
                            _ = v43
                            _ = v44
                            _ = v45
                        } else {
                            var v46 string = v13
                            _ = v46
                        }
                    case "":
                        if true {
                            switch "a" {
                            case "":
                                switch v13 {
                                case "\n\t":
                                    v41 -= v41
                                }
                            case "a":
                                v41 -= v41
                            default:
                                if (-30 >= -9223372036854775808) {
                                    var v47 uint64 = v5
                                    v47 -= (^(v5))
                                    var v48 uint8 = 68
                                    _ = v47
                                    _ = v48
                                } else {
                                    v49 := GenS3{F0: "", F1: 255}
                                    _ = v49
                                }
                            }
                            fmt.Println(v5, v3, v26)
                        }
                    default:
                        if true {
                            v13 = v13
                        }
                    }
                    _ = fi42
                }
                _ = v41
            }
            switch v13 {
            case "\n\t":
                v5 -= (^(v5))
            case "foo bar":
                fmt.Println(v5, v3, fi37)
            default:
                switch v13 {
                case "\n\t":
                    v3 = 42
                case "hello":
                    v26 = v26
                default:
                    switch 32767 {
                    case -9223372036854775808:
                        switch (fi37 ^ fi37) {
                        case 9223372036854775807:
                            var v50 uint8 = 127
                            _ = v50
                        default:
                            for fi51 := 0; fi51 < 14; fi51++ {
                                var v52 int8 = -128
                                v2[fi37] = (v13 + (("hello" + v13) + v13))
                                _ = fi51
                                _ = v52
                            }
                        }
                    case 705:
                        var v53 uint = 612
                        _ = v53
                    }
                }
            }
            var v54 int32 = (v3 + (v3 - 0))
            _ = fi37
            _ = v54
        }
    default:
        v13 = v13
    }
    var v55 string = ""
    v56 := GenS4{F0: -1}
    switch 42 {
    case -9223372036854775808:
        for fi57 := 0; fi57 < 18; fi57++ {
            var v58 int32 = 2147483647
            v58 = -2147483648
            fmt.Println(v58)
            _ = fi57
            _ = v58
        }
    case 42:
        switch -34 {
        case 128:
            var v59 uint = 685
            _ = v59
        }
    default:
        v60 := GenS2{F0: -1}
        _ = v60
    }
    if false {
        var v61 bool = ((9223372036854775807 == 334) || false)
        fmt.Println(v55, v3)
        switch ("hello" + ("a" + ("a" + "\n\t"))) {
        case "hello":
            switch "\n\t" {
            case "":
                v13 = v13
            case "foo bar":
                v13 = ((((((("\n\t" + ("foo bar" + v55)) + ((v13 + v13) + v55)) + v13) + v55) + v13) + ("hello" + v13)) + "")
            case "\n\t":
                var v62 string = "foo bar"
                _ = v62
            }
        case "":
            var v63 uint = 255
            _ = v63
        }
        _ = v61
    } else {
        var v64 int8 = -94
        _ = v64
    }
    var v65 int32 = v3
    for fi66 := 0; fi66 < 12; fi66++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v67 uint64 = 1
        _ = fi66
        _ = v67
    }
    if (65535 != uint(v5)) {
        var v68 float32 = -2.25
        for fi69 := 0; fi69 < 8; fi69++ {
            v5 -= v5
            var v70 bool = true
            _ = fi69
            _ = v70
        }
        for fi71 := 0; fi71 < 1; fi71++ {
            fmt.Println(v68)
            switch fi71 {
            case 2147483647:
                for fi72 := 0; fi72 < 3; fi72++ {
                    var v73 float32 = v68
                    var v74 int16 = 1
                    _ = fi72
                    _ = v73
                    _ = v74
                }
            }
            _ = fi71
        }
        _ = v68
    }
    fmt.Println(v65, v65, v65)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v65
}
