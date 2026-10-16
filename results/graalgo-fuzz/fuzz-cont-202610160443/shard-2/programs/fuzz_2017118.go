package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int; F2 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint8 }
func genFn0() bool {
    var v7 int32 = 566
    fmt.Println(v7)
    if true {
        var v8 float32 = 1.5
        v9 := []int{-9223372036854775808, -696, 0, 128}
        _ = v8
        _ = v9
    } else {
        var v10 bool = ((!((1.5 * 3.14159) > float64(1.5))) || (32767 == 1))
        v7 += (v7 - v7)
        if (2 != 2) {
            switch 0 {
            case 1:
                if (83 >= 127) {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v12 := []int{}
                        _ = fi11
                        _ = v12
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        if v10 {
                            v10 = v10
                        } else {
                            if false {
                                v13 := GenS1{F0: 2}
                                _ = v13
                            }
                            if (0 < 65535) {
                                for fi14 := 0; fi14 < 18; fi14++ {
                                    var v15 uint8 = 42
                                    v16 := []int{32768}
                                    v15 -= (^(v15))
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                                switch 4294967295 {
                                case 9223372036854775807:
                                    v7 += (v7 ^ v7)
                                }
                            } else {
                                var v17 int = 32768
                                _ = v17
                            }
                            fmt.Println(v10, v7)
                        }
                    case "hello":
                        for fi18 := 0; fi18 < 4; fi18++ {
                            switch ("a" + "a") {
                            case "a":
                                v19 := GenS2{F0: 483}
                                _ = v19
                            }
                            v20 := GenS1{F0: 2}
                            var v21 uint16 = 0
                            _ = fi18
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 uint16 = 533
                    v23 := GenS3{F0: 18446744073709551615, F1: v22, F2: -761}
                    _ = v22
                    _ = v23
                }
            case -1:
                v10 = v10
            case 4294967295:
                var v24 uint64 = 669
                _ = v24
            default:
                v7 = v7
            }
            var v25 int8 = 127
            var v26 uint8 = 0
            _ = v25
            _ = v26
        } else {
            v7 = (v7 & (-(v7)))
        }
        _ = v10
    }
    _ = v7
    return false
}
func genFn1() int16 {
    switch "" {
    case "hello":
        if false {
            v27 := []int{}
            _ = v27
        }
    case "\n\t":
        switch -1 {
        case 32767:
            var v28 uint64 = 42
            _ = v28
        case 9223372036854775807:
            v29 := []int{2147483647, -538, 909, -1}
            _ = v29
        }
    default:
        var v30 int8 = 122
        _ = v30
    }
    switch 32767 {
    case 0:
        switch "foo bar" {
        case "hello":
            var v31 int16 = 1
            _ = v31
        }
    case 4294967295:
        if false {
            var v32 int16 = 339
            for fi33 := 0; fi33 < 17; fi33++ {
                v34 := GenS3{F0: 18446744073709551615, F1: 1, F2: 135}
                _ = fi33
                _ = v34
            }
            v35 := GenS4{F0: 128}
            _ = v32
            _ = v35
        } else {
            var v36 int = -522
            v37 := GenS3{F0: 780, F1: 0, F2: 792}
            v36 = v36
            _ = v36
            _ = v37
        }
    case -513:
        v38 := GenS3{F0: 18446744073709551615, F1: 0, F2: 51}
        _ = v38
    default:
        v39 := GenS5{F0: false, F1: -1, F2: 932}
        _ = v39
    }
    var v40 int64 = -873
    v41 := GenS2{F0: 543}
    _ = v40
    _ = v41
    return -1
}

func main() {
    v1 := []uint32{957, 0, 2147483647, 1}
    var v2 int8 = 0
    var v3 int16 = -32768
    var v4 int64 = 42
    var v5 string = ("" + "")
    var v6 float64 = (((0.1 - 1.5) - 0.0) - (((float64(0.1) - 3.14159) * (0.0 * 1.5)) / 3.14159))
    switch 1 {
    case 127:
        if true {
            fmt.Println(v5, v3)
            if len(v1) > 0 {
                v1[0] = 525
            }
        }
    case 461:
        var v42 float32 = 1e10
        _ = v42
    case -9223372036854775808:
        fmt.Println(v6)
    }
    if (!true) {
        if false {
            if true {
                for fi43 := 0; fi43 < 14; fi43++ {
                    if (!((true && (v6 != (v6 / float64((3.14159 / 3.14159))))) || (2 > 255))) {
                        var v44 float32 = ((1e10 / 1.0) + float32((3.14159 * v6)))
                        v5 = v5
                        _ = v44
                    }
                    _ = fi43
                }
                if (!true) {
                    var v45 int8 = 0
                    fmt.Println(v45)
                    v5 = ((v5 + ((("" + (v5 + "")) + v5) + (v5 + v5))) + v5)
                    _ = v45
                } else {
                    v3 = v3
                }
                switch 42 {
                case 996:
                    var v46 uint8 = 128
                    _ = v46
                case 388:
                    if false {
                        v2 += v2
                        v1 = append(v1, 0)
                    } else {
                        v4 = v4
                        if (!(((v5 != v5) || (67 <= 2)) || (!(2 < 128)))) {
                            for fi47 := 0; fi47 < 9; fi47++ {
                                for fi48 := 0; fi48 < 14; fi48++ {
                                    v3 = (-(v3))
                                    fmt.Println(v2, v6)
                                    _ = fi48
                                }
                                if (v4 >= int64(v2)) {
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        var v50 uint16 = 831
                                        var v51 uint16 = (v50 & (v50 >> ((110) & 63)))
                                        switch fi49 {
                                        case 42:
                                            v50 = v50
                                        }
                                        v52 := GenS2{F0: 2}
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    for fi53 := 0; fi53 < 2; fi53++ {
                                        v54 := GenS6{F0: "\n\t"}
                                        v2 = 31
                                        v55 := GenS4{F0: 42}
                                        _ = fi53
                                        _ = v54
                                        _ = v55
                                    }
                                } else {
                                    for fi56 := 0; fi56 < 2; fi56++ {
                                        switch (fi56 ^ fi47) {
                                        case 4294967295:
                                            switch (fi56 % ((128) | 1)) {
                                            case -1:
                                                v57 := GenS3{F0: 1, F1: 65535, F2: v4}
                                                _ = v57
                                            default:
                                                v5 = v5
                                            }
                                        case 128:
                                            v58 := GenS6{F0: (v5 + v5)}
                                            _ = v58
                                        default:
                                            v2 -= (v2 ^ v2)
                                        }
                                        if (128 >= 42) {
                                            v2 = v2
                                        }
                                        _ = fi56
                                    }
                                }
                                v4 = v4
                                _ = fi47
                            }
                            for fi59 := 0; fi59 < 8; fi59++ {
                                fmt.Println(fi59)
                                fmt.Println(fi59, v4)
                                _ = fi59
                            }
                        } else {
                            var v60 int16 = 1
                            v3 = v60
                            _ = v60
                        }
                        fmt.Println(v5)
                    }
                case -663:
                    _ = len(v1)
                default:
                    v2 -= v2
                }
            }
        } else {
            fmt.Println(v5, v2)
            v5 = "foo bar"
            switch 2147483647 {
            case 2147483647:
                v61 := GenS3{F0: 288, F1: 0, F2: v4}
                _ = v61
            }
        }
        if len(v1) > 0 {
            v1[0] = 750
        }
    } else {
        switch "hello" {
        case "\n\t":
            var v62 bool = false
            _ = v62
        case "foo bar":
            v63 := GenS5{F0: (v2 == (v2 % ((v2) | 1))), F1: -489, F2: 0}
            _ = v63
        case "hello":
            var v64 uint32 = 601
            _ = v64
        }
        var v65 int32 = 948
        if ((^(v3)) <= 32767) {
            v1 = append(v1, 2)
            switch (v5 + "foo bar") {
            case "":
                var v66 float32 = 1.5
                _ = v66
            case "hello":
                for fi67 := 0; fi67 < 6; fi67++ {
                    v4 += v4
                    v65 = (-(v65))
                    _ = fi67
                }
            default:
                _ = len(v1)
            }
            v2 = (v2 + v2)
        }
        _ = v65
    }
    var v68 uint = 392
    var v69 uint16 = 2
    var v70 int = 1
    for fi71 := 0; fi71 < 16; fi71++ {
        var v72 float64 = 1e10
        if false {
            v73 := GenS7{F0: v68, F1: 1}
            var v74 uint = (v68 + (v68 * v68))
            _ = v73
            _ = v74
        } else {
            var v75 int8 = v2
            _ = v75
        }
        _ = fi71
        _ = v72
    }
    var v76 uint64 = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 -= v6
    fmt.Println(v2, v6)
    for fi77 := 0; fi77 < 15; fi77++ {
        v76 = (v76 % ((v76) | 1))
        _ = fi77
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v70
    _ = v76
}
