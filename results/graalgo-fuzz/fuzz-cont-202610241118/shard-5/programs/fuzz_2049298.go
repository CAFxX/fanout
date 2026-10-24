package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string }

func main() {
    var v1 int64 = 42
    var v2 int16 = 1000
    v3 := []uint32{2}
    if false {
        switch 0 {
        case 32768:
            switch (("foo bar" + "foo bar") + "a") {
            case "a":
                var v4 uint8 = 127
                _ = v4
            default:
                var v5 float32 = float32(0.0)
                _ = v5
            }
        default:
            v2 = (^(v2))
        }
    } else {
        if (v1 == 42) {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint8 = 45
                _ = fi6
                _ = v7
            }
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: (v1 << ((499) & 63)), F1: (!false), F2: (fi8 & (fi8 << ((420) & 63)))}
                if len(v3) > 0 {
                    v3[0] = 738
                }
                _ = fi8
                _ = v9
            }
        }
        fmt.Println(v2)
        v10 := GenS2{F0: 2147483647, F1: 0}
        _ = v10
    }
    v3 = append(v3, 2147483647)
    if true {
        if false {
            if false {
                v11 := GenS1{F0: 1, F1: (2147483648 != 554), F2: 1}
                _ = v11
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch 0 {
                case -949:
                    var v12 int = 32767
                    _ = v12
                case 42:
                    v1 += -1
                default:
                    v1 = (v1 ^ v1)
                }
                var v13 string = ("hello" + ("a" + ("hello" + "")))
                _ = v13
            }
        }
        var v14 bool = (-117 < v1)
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int8 = -95
            switch fi15 {
            case -2147483648:
                if ((fi15 != int(v16)) && (2147483647 < 42)) {
                    fmt.Println(v16, fi15, v2)
                    for fi17 := 0; fi17 < 1; fi17++ {
                        switch 111 {
                        case -384:
                            v18 := GenS2{F0: -1, F1: 0}
                            _ = v18
                        case 128:
                            v16 = v16
                        default:
                            switch "" {
                            case "foo bar":
                                v3 = append(v3, 218)
                            case "a":
                                v2 = (v2 - v2)
                            case "\n\t":
                                var v19 string = "a"
                                _ = v19
                            default:
                                var v20 int8 = 0
                                _ = v20
                            }
                        }
                        v1 += (v1 | -1)
                        var v21 int16 = 1000
                        _ = fi17
                        _ = v21
                    }
                }
            default:
                if v14 {
                    var v22 int64 = v1
                    if (v14 && (false || (906 >= 65535))) {
                        switch 32768 {
                        case 2147483647:
                            for fi23 := 0; fi23 < 12; fi23++ {
                                var v24 int8 = v16
                                fmt.Println(v1, v16)
                                _ = fi23
                                _ = v24
                            }
                        case -868:
                            var v25 int16 = v2
                            _ = v25
                        default:
                            fmt.Println(v1, v1)
                        }
                        v26 := GenS2{F0: 0, F1: 43}
                        if true {
                            v16 = v16
                            v16 += v16
                        }
                        _ = v26
                    }
                    _ = v22
                } else {
                    if (v14 && v14) {
                        var v27 uint16 = 0
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 4; fi28++ {
                            var v29 int64 = (v1 % (((v1 ^ 100)) | 1))
                            var v30 float32 = (0.0 / 1.5)
                            switch 931 {
                            case -2147483648:
                                var v31 uint32 = 2147483647
                                var v32 bool = (659 == uint16(v31))
                                v33 := GenS2{F0: 0, F1: (v31 | (v31 / ((340) | 1)))}
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        v34 := GenS1{F0: v1, F1: (!(0 <= -556)), F2: (-(fi15))}
                        var v35 uint32 = 2147483647
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS2{F0: int32(v1), F1: 906}
                    if v14 {
                        var v37 float32 = (1.5 / (-2.25 - float32(1.5)))
                        _ = v37
                    } else {
                        v2 -= (v2 ^ v2)
                        v38 := GenS1{F0: v1, F1: (true || (1 == 1)), F2: fi15}
                        _ = v38
                    }
                    _ = v36
                }
            }
            v1 = v1
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    var v39 int32 = 42
    for fi40 := 0; fi40 < 4; fi40++ {
        v1 += v1
        for fi41 := 0; fi41 < 5; fi41++ {
            v42 := GenS1{F0: 9223372036854775807, F1: (!(!(3.14159 >= float64((0.0 - (1.5 * -2.25)))))), F2: 127}
            fmt.Println(v39, fi41)
            fmt.Println(v1)
            _ = fi41
            _ = v42
        }
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = fi40
    }
    v1 += (^(v1))
    for fi43 := 0; fi43 < 20; fi43++ {
        fmt.Println(fi43, v1, v1)
        for fi44 := 0; fi44 < 19; fi44++ {
            v1 -= (-(v1))
            fmt.Println(v1, fi43, v39)
            _ = fi44
        }
        _ = fi43
    }
    _ = len(v3)
    v39 = 513
    var v45 float64 = 0.0
    for fi46 := 0; fi46 < 16; fi46++ {
        fmt.Println(fi46, v39, v39)
        if false {
            fmt.Println(v45)
            switch "" {
            case "hello":
                for fi47 := 0; fi47 < 17; fi47++ {
                    if (!false) {
                        var v48 uint8 = 188
                        fmt.Println(fi47, v45, v45)
                        var v49 int32 = int32(v1)
                        _ = v48
                        _ = v49
                    } else {
                        fmt.Println(v1, fi47, v39)
                    }
                    _ = fi47
                }
            case "foo bar":
                switch fi46 {
                case 404:
                    var v50 float64 = 0.0
                    _ = v50
                case 894:
                    var v51 int16 = (^(v2))
                    _ = v51
                case -399:
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                }
            case "\n\t":
                for fi52 := 0; fi52 < 12; fi52++ {
                    if ((-(v39)) > (^(v39))) {
                        switch 285 {
                        case -276:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case -145:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v39 = v39
                            v53 := GenS3{F0: (2 != 2)}
                            _ = v53
                        case -1:
                            _ = len(v3)
                        default:
                            switch "hello" {
                            case "a":
                                v54 := GenS4{F0: 528}
                                _ = v54
                            case "hello":
                                switch fi52 {
                                case -883:
                                    v1 = (v1 % ((v1) | 1))
                                case 1:
                                    v55 := GenS5{F0: 0}
                                    _ = v55
                                default:
                                    var v56 uint16 = 65535
                                    _ = v56
                                }
                            case "":
                                v1 += 894
                            default:
                                if true {
                                    fmt.Println(v39, v39, v39)
                                    v2 += (^(v2))
                                    if ("a" <= ("foo bar" + (("foo bar" + "hello") + "hello"))) {
                                        switch (((("hello" + "a") + "hello") + (("hello" + "hello") + ("a" + "hello"))) + "hello") {
                                        case "\n\t":
                                            _ = len(v3)
                                        case "a":
                                            v57 := GenS6{F0: ("foo bar" + "hello")}
                                            _ = v57
                                        case "hello":
                                            var v58 float32 = (0.0 + -2.25)
                                            _ = v58
                                        default:
                                            v59 := GenS4{F0: 894}
                                            _ = v59
                                        }
                                        v2 += int16(fi46)
                                    }
                                    var v60 uint8 = 36
                                    var v61 int32 = 0
                                    v2 -= 42
                                    fmt.Println(v39, v1)
                                    _ = v60
                                    _ = v61
                                }
                            }
                        }
                        var v62 int = (fi46 | 4294967295)
                        var v63 bool = (float32(v45) <= 1e10)
                        _ = v62
                        _ = v63
                    }
                    _ = fi52
                }
            default:
                if (v2 > (^(v2))) {
                    if (134 > 64) {
                        _ = len(v3)
                    }
                    var v64 uint64 = 1
                    _ = v64
                }
            }
            var v65 int = (fi46 / (((-(fi46))) | 1))
            _ = v65
        }
        var v66 int64 = v1
        _ = fi46
        _ = v66
    }
    for fi67 := 0; fi67 < 11; fi67++ {
        v2 = v2
        _ = fi67
    }
    if true {
        fmt.Println(v1, v2, v45)
        v68 := GenS6{F0: ""}
        v39 += v39
        _ = v68
    }
    var v69 uint = 4294967295
    v1 = 9223372036854775807
    var v70 uint8 = 127
    fmt.Println(v1, v39, v45)
    switch -9223372036854775808 {
    case -9223372036854775808:
        v69 += v69
    case 127:
        v39 = v39
    case -560:
        for fi71 := 0; fi71 < 1; fi71++ {
            var v72 uint16 = 255
            if (!false) {
                v2 -= int16(v70)
            } else {
                switch int(v69) {
                case 398:
                    var v73 int32 = 0
                    _ = v73
                case 514:
                    switch (fi71 * fi71) {
                    case 128:
                        var v74 int32 = (v39 ^ v39)
                        _ = v74
                    }
                case 32768:
                    v1 = -831
                }
                var v75 uint = (v69 | uint(v72))
                _ = v75
            }
            _ = fi71
            _ = v72
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v45:", v45)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v45
    _ = v69
    _ = v70
}
