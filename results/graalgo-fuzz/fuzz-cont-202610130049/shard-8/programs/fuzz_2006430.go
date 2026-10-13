package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS3 struct { F0 int; F1 int32; F2 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 int64 }
type GenS5 struct { F0 int64; F1 uint64; F2 int64 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 int8 }

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]int)
    var v3 int64 = 0
    v4 := make(map[string]int)
    var v5 int16 = 32767
    if true {
        if (-2147483648 <= -2147483648) {
            v6 := GenS1{F0: 42}
            _ = v6
        }
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v3 += v3
                v9 := GenS1{F0: -9223372036854775808}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        v5 += v5
    }
    var v10 string = "a"
    var v11 int = -2147483648
    fmt.Println(v10, v3)
    var v12 float64 = 0.1
    v3 += v3
    v13 := []int{v11, v11, (^(v11)), 4294967295}
    switch "a" {
    case "a":
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 int = (-(fi14))
            if true {
                switch v11 {
                case 9223372036854775807:
                    var v16 int = v15
                    _ = v16
                case 43:
                    if true {
                        if (3.14159 > 3.14159) {
                            if ((v11 % (((v15 * (^(v15)))) | 1)) == -2147483648) {
                                v17 := GenS1{F0: v15}
                                _ = v17
                            }
                            v1 = -2.25
                            switch (v10 + (v10 + "")) {
                            case "\n\t":
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    var v19 uint16 = 2
                                    fmt.Println(v1, v12, v3)
                                    _ = fi18
                                    _ = v19
                                }
                            case "foo bar":
                                var v20 int8 = -1
                                _ = v20
                            default:
                                v11 -= (v11 + 15)
                            }
                        } else {
                            v10 = ((("\n\t" + v10) + v10) + v10)
                            if (42 != 0) {
                                v15 = (^(v15))
                            }
                        }
                        _ = len(v13)
                        fmt.Println(v15, fi14)
                    }
                case 32768:
                    v21 := GenS1{F0: (v11 ^ (^(fi14)))}
                    _ = v21
                default:
                    v22 := GenS1{F0: -1}
                    _ = v22
                }
                v15 = (v15 - 127)
                if len(v13) > 0 {
                    v13[0] = (fi14 + v15)
                }
            } else {
                v23 := v4[""]
                var v24 uint64 = 0
                var v25 uint = 255
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi14
            _ = v15
        }
    case "hello":
        v26 := GenS2{F0: 1, F1: 42, F2: 1}
        _ = v26
    case "":
        if true {
            if (!false) {
                v27 := GenS1{F0: v11}
                var v28 int64 = v3
                _ = v27
                _ = v28
            } else {
                v29 := GenS1{F0: v11}
                for fi30 := 0; fi30 < 4; fi30++ {
                    if (491 != 2) {
                        v31 := GenS3{F0: 2147483647, F1: -395, F2: 1}
                        fmt.Println(v3, v11, fi30)
                        v5 = 32767
                        _ = v31
                    }
                    _ = fi30
                }
                _ = v29
            }
        } else {
            fmt.Println(v3)
            fmt.Println(v11, v12)
            var v32 float32 = (0.0 / 1.0)
            _ = v32
        }
    }
    v13 = append(v13, v11)
    v5 = v5
    fmt.Println(v10)
    v33 := GenS2{F0: 42, F1: 255, F2: -813}
    _ = len(v2)
    if true {
        v34 := GenS1{F0: (-(v11))}
        fmt.Println(v10, v1, v3)
        switch v10 {
        case "foo bar":
            if (false || ((!false) || (!true))) {
                for fi35 := 0; fi35 < 2; fi35++ {
                    var v36 int = -268
                    _ = fi35
                    _ = v36
                }
                switch v10 {
                case "a":
                    var v37 int8 = 104
                    _ = v37
                case "foo bar":
                    for fi38 := 0; fi38 < 10; fi38++ {
                        if true {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        } else {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        }
                        var v39 int32 = -322
                        _ = fi38
                        _ = v39
                    }
                default:
                    var v40 uint32 = 987
                    _ = v40
                }
                v10 = "a"
            } else {
                v10 = "hello"
                var v41 float64 = v12
                v5 = v5
                _ = v41
            }
        default:
            if false {
                if true {
                    if true {
                        v13 = append(v13, 32767)
                        if (65535 <= 0) {
                            delete(v2, v10)
                            var v42 int8 = -128
                            _ = v42
                        }
                    }
                } else {
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        var v43 string = ("a" + v10)
                        _ = v43
                    case 4294967295:
                        var v44 float32 = (((1e10 - float32((v12 / -2.25))) - 1.5) + 1.5)
                        _ = v44
                    default:
                        var v45 uint16 = 255
                        _ = v45
                    }
                    var v46 bool = ((true || (1 > 4294967295)) || (!false))
                    _ = v46
                }
            } else {
                if (!(65535 < 152)) {
                    if (2147483647 > 984) {
                        switch v10 {
                        case "\n\t":
                            for fi47 := 0; fi47 < 17; fi47++ {
                                var v48 int = (v11 + (v11 / (((v11 / ((v11) | 1))) | 1)))
                                for fi49 := 0; fi49 < 15; fi49++ {
                                    v50 := GenS4{F0: "hello", F1: 255, F2: (-(v3))}
                                    _ = len(v13)
                                    var v51 string = v10
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                }
                                _ = fi47
                                _ = v48
                            }
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                if true {
                                    v52 := v4[(("\n\t" + v10) + "hello")]
                                    _ = v52
                                }
                            case "hello":
                                if true {
                                    v3 = v3
                                    if (2147483647 < 641) {
                                        delete(v4, (v10 + ((((v10 + "foo bar") + v10) + "foo bar") + v10)))
                                        v5 = v5
                                    } else {
                                        fmt.Println(v3, v11)
                                        var v53 uint8 = 46
                                        if true {
                                            var v54 uint8 = 127
                                            v55 := GenS2{F0: 4294967295, F1: 64, F2: -260}
                                            _ = v54
                                            _ = v55
                                        }
                                        _ = v53
                                    }
                                    fmt.Println(v3, v12)
                                }
                            case "foo bar":
                                var v56 string = v10
                                _ = v56
                            default:
                                fmt.Println(v1, v12, v3)
                            }
                        }
                        if false {
                            if (-2147483648 < 1) {
                                v5 -= (v5 & int16(v11))
                                fmt.Println(v11, v12)
                                for fi57 := 0; fi57 < 9; fi57++ {
                                    for fi58 := 0; fi58 < 12; fi58++ {
                                        v5 = v5
                                        _ = fi58
                                    }
                                    _ = fi57
                                }
                            }
                            switch ("hello" + v10) {
                            case "":
                                if (!(65535 != 859)) {
                                    fmt.Println(v12, v1)
                                    v11 = v11
                                    v59 := GenS5{F0: (v3 - (v3 & v3)), F1: 890, F2: v3}
                                    _ = v59
                                }
                            case "foo bar":
                                v60 := GenS6{F0: (("a" + v10) > v10), F1: v11}
                                _ = v60
                            }
                            var v61 uint64 = 640
                            _ = v61
                        } else {
                            fmt.Println(v3, v1)
                        }
                    } else {
                        for fi62 := 0; fi62 < 9; fi62++ {
                            switch v11 {
                            case -1:
                                if false {
                                    var v63 int16 = (v5 % (((v5 & v5)) | 1))
                                    _ = v63
                                } else {
                                    switch 32767 {
                                    case 42:
                                        v64 := GenS3{F0: fi62, F1: 1, F2: (v5 << ((0) & 63))}
                                        _ = v64
                                    case 9223372036854775807:
                                        if (128 > 255) {
                                            v10 = v10
                                        }
                                    default:
                                        fmt.Println(v1)
                                    }
                                    var v65 uint16 = 2
                                    _ = v65
                                }
                            case 2147483647:
                                delete(v2, ("\n\t" + (v10 + (v10 + v10))))
                            default:
                                v66 := GenS7{F0: 42}
                                _ = v66
                            }
                            delete(v4, "\n\t")
                            v11 -= v11
                            _ = fi62
                        }
                    }
                } else {
                    switch (-(v11)) {
                    case -2147483648:
                        fmt.Println(v5)
                    case 4294967295:
                        for fi67 := 0; fi67 < 19; fi67++ {
                            _ = len(v4)
                            for fi68 := 0; fi68 < 20; fi68++ {
                                v1 = ((1e10 / v12) - 1.5)
                                _ = fi68
                            }
                            _ = fi67
                        }
                    case 1:
                        v13 = append(v13, v11)
                    default:
                        v10 = ""
                    }
                    v5 = v5
                }
                v69 := GenS1{F0: int(v5)}
                _ = v69
            }
        }
        _ = v34
    }
    switch "foo bar" {
    case "hello":
        v3 -= (v3 % ((-1) | 1))
    case "\n\t":
        var v70 int64 = (v3 ^ (v3 + (v3 >> ((2) & 63))))
        _ = v70
    case "a":
        var v71 int16 = (v5 >> ((4294967295) & 63))
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
