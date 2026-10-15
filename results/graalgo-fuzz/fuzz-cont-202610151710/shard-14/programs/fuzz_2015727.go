package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 string; F2 string }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS7 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint64, p1 int32, p2 int16) string {
    fmt.Println(p1)
    var v7 uint16 = 763
    v8 := GenS1{F0: ("hello" + ""), F1: ("\n\t" + "foo bar")}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := []uint32{4294967295, 419, 0}
    var v2 int16 = 1
    v3 := []int{1, 29, 0, 1}
    var v4 bool = false
    var v5 int32 = 1
    var v6 int8 = 42
    var v9 uint8 = 50
    var v10 float32 = (3.14159 - 1.5)
    v11 := GenS2{F0: ((-(v5)) != (^(v5))), F1: -9223372036854775808, F2: 1}
    v2 += (v2 & v2)
    v6 = v6
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 string = "\n\t"
    switch 2147483647 {
    case 0:
        fmt.Println(v12, v4)
    case 9223372036854775807:
        var v13 int = -303
        _ = v13
    }
    if (float64((v10 - 1.5)) >= -2.25) {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    } else {
        if (v4 || (!true)) {
            v12 = v12
        }
        v14 := GenS2{F0: (2 < 530), F1: -1, F2: 278}
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v2 -= (v2 - (-(v2)))
        var v16 string = ((v12 + ("hello" + "hello")) + v12)
        _ = fi15
        _ = v16
    }
    switch "a" {
    case "foo bar":
        for fi17 := 0; fi17 < 12; fi17++ {
            if v4 {
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
            _ = fi17
        }
    case "a":
        if (v2 > (v2 % (((v2 | (^(v2)))) | 1))) {
            fmt.Println(v5)
            fmt.Println(v9, v5)
            v5 = 157
        }
    default:
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(fi18, v10)
            var v19 uint32 = 1
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        if (4294967295 < 42) {
            for fi21 := 0; fi21 < 10; fi21++ {
                fmt.Println(v5, fi20, v9)
                fmt.Println(v12, fi20)
                _ = fi21
            }
            v4 = (32768 == 1)
        } else {
            fmt.Println(v4, v4, v10)
            v6 = (-(v6))
            _ = len(v3)
        }
        switch (^(fi20)) {
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 757:
            v5 = (v5 / (((v5 | -1)) | 1))
        case 850:
            if ((v9 + v9) >= (v9 ^ v9)) {
                for fi22 := 0; fi22 < 8; fi22++ {
                    switch 0 {
                    case -2147483648:
                        fmt.Println(v9)
                    case 1:
                        v4 = (255 == 687)
                        fmt.Println(v5, fi20)
                        switch (v12 + (v12 + "")) {
                        case "\n\t":
                            _ = len(v3)
                        }
                        var v23 string = ("a" + (v12 + v12))
                        _ = v23
                    }
                    _ = fi22
                }
            } else {
                if false {
                    var v24 float32 = 0.0
                    _ = v24
                }
                if v4 {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v2, v12)
                        v26 := GenS3{F0: 9223372036854775807}
                        _ = fi25
                        _ = v26
                    }
                    v6 = v6
                }
                if v4 {
                    if (fi20 != 42) {
                        fmt.Println(fi20, v9)
                    }
                    switch (-(fi20)) {
                    case -578:
                        v10 = 0.0
                    case 42:
                        switch v12 {
                        case "hello":
                            if true {
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    switch v12 {
                                    case "hello":
                                        v6 = (v6 & v6)
                                    case "a":
                                        fmt.Println(v10)
                                    case "foo bar":
                                        switch "" {
                                        case "hello":
                                            if v4 {
                                                v28 := GenS2{F0: false, F1: fi27, F2: 255}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v28
                                            }
                                        case "a":
                                            var v29 uint64 = 0
                                            _ = v29
                                        default:
                                            v30 := GenS1{F0: (((((v12 + v12) + "foo bar") + v12) + (v12 + ("a" + (((v12 + v12) + "a") + v12)))) + ((v12 + ((v12 + v12) + v12)) + ((v12 + "") + v12))), F1: v12}
                                            _ = v30
                                        }
                                    }
                                    _ = fi27
                                }
                                v9 = 2
                                switch ("" + v12) {
                                case "\n\t":
                                    var v31 string = "hello"
                                    _ = v31
                                }
                            }
                        case "a":
                            v32 := GenS1{F0: v12, F1: ("" + v12)}
                            _ = v32
                        case "\n\t":
                            v4 = (((((!v4) && v4) && ((^(v9)) < 127)) || true) && v4)
                        default:
                            if v4 {
                                fmt.Println(v6, v2, v6)
                                var v33 uint8 = 217
                                _ = v33
                            } else {
                                v34 := GenS4{F0: -951, F1: (v12 + v12), F2: v12}
                                var v35 int16 = v2
                                var v36 uint = 1
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                        }
                    }
                    if v4 {
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                    } else {
                        for fi37 := 0; fi37 < 15; fi37++ {
                            var v38 int32 = -79
                            if v4 {
                                if (!v4) {
                                    var v39 float32 = 1.5
                                    v38 = (v38 * v5)
                                    _ = len(v1)
                                    _ = v39
                                }
                            } else {
                                v40 := GenS5{F0: 42, F1: (v6 < v6)}
                                switch v12 {
                                case "a":
                                    var v41 uint = 0
                                    _ = v41
                                case "":
                                    v42 := GenS2{F0: (true || true), F1: fi37, F2: 2}
                                    _ = v42
                                default:
                                    v38 += v5
                                }
                                v43 := GenS5{F0: v38, F1: v4}
                                _ = v40
                                _ = v43
                            }
                            _ = fi37
                            _ = v38
                        }
                    }
                }
            }
        default:
            switch 1 {
            case 4294967295:
                for fi44 := 0; fi44 < 17; fi44++ {
                    v45 := GenS1{F0: (v12 + ""), F1: ((v12 + "") + ("\n\t" + v12))}
                    var v46 uint16 = 2
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            default:
                var v47 uint16 = 765
                _ = v47
            }
        }
        v48 := GenS1{F0: "hello", F1: v12}
        _ = fi20
        _ = v48
    }
    var v49 int64 = 517
    if (676 <= 69) {
        v5 = 494
        v3 = append(v3, 128)
        switch 2147483647 {
        case 141:
            fmt.Println(v4, v12)
        case 42:
            if (127 <= 1) {
                v49 = (v49 & v49)
            } else {
                if (!v4) {
                    fmt.Println(v49, v5)
                } else {
                    v4 = (true && v4)
                    v49 -= 0
                    var v50 string = v12
                    _ = v50
                }
            }
        case 600:
            var v51 uint32 = 2
            _ = v51
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (0 == 32768) {
        if len(v1) > 0 {
            v1[0] = 142
        }
    } else {
        var v52 float64 = 0.1
        _ = len(v1)
        switch 32768 {
        case 42:
            if len(v1) > 0 {
                v1[0] = 0
            }
        case 13:
            v53 := GenS6{F0: v9, F1: 42, F2: v9}
            _ = v53
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v52
    }
    v54 := GenS5{F0: 2147483647, F1: v4}
    for fi55 := 0; fi55 < 17; fi55++ {
        if v4 {
            v6 += (-(v6))
            var v56 int8 = v6
            _ = v56
        } else {
            switch "" {
            case "hello":
                var v57 float64 = (-2.25 * (1e10 - float64(v10)))
                _ = v57
            default:
                if (4294967295 < 1) {
                    v4 = v4
                } else {
                    switch ("hello" + ((("foo bar" + v12) + "") + "")) {
                    case "hello":
                        for fi58 := 0; fi58 < 12; fi58++ {
                            var v59 uint16 = 255
                            _ = fi58
                            _ = v59
                        }
                    case "a":
                        var v60 float64 = 1.5
                        _ = v60
                    case "foo bar":
                        if true {
                            for fi61 := 0; fi61 < 17; fi61++ {
                                v62 := GenS4{F0: -1, F1: "foo bar", F2: v12}
                                var v63 int8 = v6
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                            v1 = append(v1, 871)
                            v12 = v12
                        }
                    default:
                        v64 := GenS7{F0: (v9 / ((v9) | 1)), F1: 1}
                        _ = v64
                    }
                    var v65 int = fi55
                    _ = v65
                }
            }
            if v4 {
                v1 = append(v1, 65535)
                var v66 uint32 = 211
                _ = v66
            } else {
                for fi67 := 0; fi67 < 14; fi67++ {
                    fmt.Println(v49)
                    _ = fi67
                }
            }
        }
        _ = fi55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v54
    _ = v6
    _ = v9
}
