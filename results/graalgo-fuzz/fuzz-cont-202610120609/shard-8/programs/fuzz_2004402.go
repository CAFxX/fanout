package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 int8; F1 int32 }

func main() {
    var v1 string = "foo bar"
    var v2 int64 = -984
    var v3 uint = 484
    var v4 uint = v3
    v5 := []int{-541, 0}
    if (42 == 57) {
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        var v6 int8 = -128
        fmt.Println(v1)
        _ = v6
    } else {
        var v7 uint32 = 0
        fmt.Println(v4, v1, v4)
        var v8 uint = v4
        _ = v7
        _ = v8
    }
    if (-1 >= -1) {
        var v9 uint16 = 32767
        var v10 float64 = 1.5
        _ = v10
        _ = v9
    } else {
        if false {
            v3 += (v4 >> (((v4 & 4294967295)) & 63))
        }
        if (1 > 1) {
            if true {
                var v11 float64 = (1.5 + 3.14159)
                var v12 float32 = 3.14159
                _ = v11
                _ = v12
            }
            v13 := GenS1{F0: 0, F1: (false && (!true))}
            _ = v13
        } else {
            if true {
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 bool = (-1 <= 180)
                    switch v1 {
                    case "a":
                        v2 = (^(v2))
                    default:
                        switch (("foo bar" + v1) + v1) {
                        case "hello":
                            v3 = v4
                        default:
                            var v16 uint8 = 255
                            _ = v16
                        }
                    }
                    _ = fi14
                    _ = v15
                }
            } else {
                for fi17 := 0; fi17 < 9; fi17++ {
                    switch (fi17 - fi17) {
                    case 1:
                        for fi18 := 0; fi18 < 17; fi18++ {
                            v4 = v3
                            var v19 int8 = 127
                            switch "foo bar" {
                            case "hello":
                                var v20 int = fi17
                                _ = v20
                            case "foo bar":
                                var v21 float32 = 3.14159
                                _ = v21
                            case "\n\t":
                                if true {
                                    var v22 float64 = -2.25
                                    v19 = v19
                                    _ = v22
                                }
                            default:
                                v19 = (^(v19))
                            }
                            _ = fi18
                            _ = v19
                        }
                    case 4294967295:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case -926:
                        var v23 int16 = 0
                        _ = v23
                    default:
                        fmt.Println(v1, v3)
                    }
                    _ = fi17
                }
                switch -160 {
                case 127:
                    if len(v5) > 0 {
                        v5[0] = 127
                    }
                case 1:
                    v1 = v1
                default:
                    var v24 uint64 = 1
                    _ = v24
                }
            }
            if (true && false) {
                var v25 string = v1
                var v26 int = 748
                _ = v25
                _ = v26
            } else {
                v5 = append(v5, 4294967295)
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 float64 = (((float64((3.14159 * ((-2.25 - 3.14159) / 0.1))) + 0.1) / (-2.25 + (float64((0.1 * 1e10)) / 1e10))) / 1e10)
                    if len(v5) > 0 {
                        v5[0] = 128
                    }
                    var v29 int64 = v2
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                fmt.Println(v1, v2)
            }
            fmt.Println(v1)
        }
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        _ = len(v5)
        for fi31 := 0; fi31 < 19; fi31++ {
            v1 = "\n\t"
            v32 := GenS2{F0: 128}
            if (v1 > v1) {
                var v33 int32 = 42
                _ = v33
            } else {
                fmt.Println(v2, fi30)
                _ = len(v5)
            }
            _ = fi31
            _ = v32
        }
        var v34 uint16 = 2
        _ = fi30
        _ = v34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 uint = v4
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        if true {
            for fi38 := 0; fi38 < 16; fi38++ {
                v39 := GenS3{F0: 2147483647, F1: 32767}
                switch fi37 {
                case 1:
                    var v40 int64 = (v2 % ((v2) | 1))
                    _ = v40
                case 9223372036854775807:
                    v1 = ((v1 + v1) + ("" + (v1 + (v1 + ("hello" + "a")))))
                case 32767:
                    v41 := GenS4{F0: -9223372036854775808, F1: 377}
                    _ = v41
                }
                _ = fi38
                _ = v39
            }
            for fi42 := 0; fi42 < 13; fi42++ {
                v5 = append(v5, fi37)
                _ = fi42
            }
            switch "a" {
            case "":
                var v43 uint64 = 571
                _ = v43
            default:
                switch "\n\t" {
                case "":
                    fmt.Println(fi37, v4)
                }
            }
        }
        _ = fi37
    }
    if true {
        v3 += 65535
    } else {
        if (v1 >= "foo bar") {
            var v44 uint16 = 988
            _ = v44
        }
        for fi45 := 0; fi45 < 3; fi45++ {
            var v46 uint64 = 575
            v47 := GenS4{F0: v2, F1: 32767}
            switch 744 {
            case 127:
                v3 = v3
            case 42:
                fmt.Println(v3)
            case 32767:
                if len(v5) > 0 {
                    v5[0] = -918
                }
            }
            _ = fi45
            _ = v46
            _ = v47
        }
    }
    var v48 uint64 = 4294967295
    var v49 uint64 = (^(v48))
    if false {
        var v50 bool = (89 != 100)
        if v50 {
            v3 += (v4 * 258)
        }
        if v50 {
            if false {
                fmt.Println(v50)
                if v50 {
                    var v51 int32 = 1
                    var v52 int = 616
                    _ = v51
                    _ = v52
                } else {
                    v48 -= 843
                    for fi53 := 0; fi53 < 14; fi53++ {
                        _ = len(v5)
                        switch -2147483648 {
                        case 592:
                            fmt.Println(fi53, v1, v48)
                        }
                        if (0 <= 0) {
                            var v54 uint16 = 2
                            for fi55 := 0; fi55 < 1; fi55++ {
                                fmt.Println(v4, fi53)
                                _ = fi55
                            }
                            _ = len(v5)
                            _ = v54
                        }
                        _ = fi53
                    }
                }
            } else {
                v56 := GenS1{F0: 127, F1: (105 < 100)}
                fmt.Println(v4)
                v57 := GenS1{F0: -128, F1: (((^(v48)) == 584) && v50)}
                _ = v56
                _ = v57
            }
            for fi58 := 0; fi58 < 17; fi58++ {
                if len(v5) > 0 {
                    v5[0] = (fi58 % ((fi58) | 1))
                }
                switch (fi58 * (fi58 << (((v3 << (((v3 >> ((18446744073709551615) & 63))) & 63))) & 63))) {
                case 42:
                    for fi59 := 0; fi59 < 10; fi59++ {
                        if ((v1 > "a") && ((^(v3)) >= (v4 ^ (v3 - (v3 / (((^(v3))) | 1)))))) {
                            v3 += 369
                        }
                        var v60 int32 = 42
                        var v61 uint32 = 741
                        v48 = (v48 & v49)
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                default:
                    v1 = "a"
                }
                _ = fi58
            }
            v62 := GenS1{F0: -128, F1: v50}
            _ = v62
        } else {
            v4 += (v3 ^ v3)
            switch v1 {
            case "foo bar":
                v63 := GenS4{F0: (v2 ^ (v2 + v2)), F1: uint16(v3)}
                _ = v63
            case "\n\t":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "a":
                v64 := GenS5{F0: (v2 - (v2 - v2)), F1: -2147483648}
                _ = v64
            }
        }
        _ = v50
    } else {
        for fi65 := 0; fi65 < 20; fi65++ {
            v66 := GenS6{F0: uint32(v3), F1: 65535}
            v3 += 1
            var v67 int32 = 431
            _ = fi65
            _ = v66
            _ = v67
        }
        _ = len(v5)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi68 := 0; fi68 < 18; fi68++ {
        for fi69 := 0; fi69 < 13; fi69++ {
            v70 := GenS4{F0: 1, F1: 588}
            var v71 float32 = -2.25
            _ = len(v5)
            v72 := GenS1{F0: 100, F1: ((-32768 <= -32768) || ((v49 << (((v3 - (v3 % ((v4) | 1)))) & 63)) < v48))}
            _ = fi69
            _ = v70
            _ = v71
            _ = v72
        }
        _ = fi68
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    if true {
        _ = len(v5)
        if false {
            var v73 int8 = 37
            fmt.Println(v49, v48, v4)
            var v74 uint64 = v49
            _ = v73
            _ = v74
        } else {
            _ = len(v5)
            for fi75 := 0; fi75 < 5; fi75++ {
                v4 = 1
                _ = fi75
            }
            switch v1 {
            case "\n\t":
                for fi76 := 0; fi76 < 12; fi76++ {
                    _ = len(v5)
                    v77 := GenS7{F0: (v3 | v4), F1: 127}
                    _ = fi76
                    _ = v77
                }
            default:
                v78 := GenS5{F0: v2, F1: 335}
                _ = v78
            }
        }
        v48 = (v49 >> ((v4) & 63))
    } else {
        v79 := GenS8{F0: 1, F1: -2147483648}
        _ = v79
    }
    v1 = v1
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
