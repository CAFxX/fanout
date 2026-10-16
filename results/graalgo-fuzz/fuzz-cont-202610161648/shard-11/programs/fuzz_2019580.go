package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32; F1 uint; F2 int8 }
type GenS10 struct { F0 int32; F1 uint64 }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 string = v1
    v6 := []int{505, -2147483648, 128, 128}
    switch v5 {
    case "":
        v7 := v4[(v5 + ((v1 + "hello") + v1))]
        _ = v7
    case "\n\t":
        var v8 float64 = 0.1
        _ = v8
    default:
        v1 = (v5 + "")
    }
    if (3.14159 > -2.25) {
        for fi9 := 0; fi9 < 18; fi9++ {
            if len(v6) > 0 {
                v6[0] = -699
            }
            v10 := GenS1{F0: 2147483648, F1: 2}
            _ = fi9
            _ = v10
        }
        fmt.Println(v1, v1)
        var v11 float32 = -2.25
        _ = v11
    }
    v12 := GenS1{F0: 65535, F1: 225}
    if ((!true) && true) {
        v1 = v1
        v13 := GenS1{F0: 2147483647, F1: 0}
        v5 = "hello"
        _ = v13
    } else {
        v1 = v5
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v5 = v5
        _ = fi14
    }
    _ = len(v6)
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 bool = true
        fmt.Println(fi15, v5, v1)
        v17 := GenS1{F0: 190, F1: 791}
        _ = fi15
        _ = v16
        _ = v17
    }
    if false {
        v18 := GenS2{F0: 155}
        if true {
            if false {
                var v19 float64 = (1e10 + (float64(1.5) - 0.1))
                _ = v19
            }
            if false {
                v1 = v5
                v1 = (v1 + (("a" + v1) + v5))
            }
        }
        if (false || (float64((((1.5 * ((-2.25 + 1e10) / (0.1 + -2.25))) + (1e10 - -2.25)) + -2.25)) < 0.0)) {
            var v20 uint64 = 18446744073709551615
            var v21 string = v1
            if (!false) {
                v20 += v20
                v22 := GenS1{F0: 701, F1: (v20 % (((v20 << ((4294967295) & 63))) | 1))}
                v23 := GenS3{F0: 4294967295, F1: 4294967295}
                _ = v22
                _ = v23
            } else {
                v4[((v1 + v21) + (v21 + (v21 + (v21 + v1))))] = -784
                switch 4294967295 {
                case 128:
                    if true {
                        switch 691 {
                        case 127:
                            if (42 > 255) {
                                v24 := GenS2{F0: 255}
                                _ = v24
                            }
                        }
                        v2[368] = 537
                        v21 = (v5 + (v1 + (("hello" + v21) + "hello")))
                    }
                }
                var v25 float64 = -2.25
                _ = v25
            }
            _ = v20
            _ = v21
        } else {
            switch 4294967295 {
            case 0:
                v26 := GenS3{F0: 944, F1: -1}
                _ = v26
            }
            fmt.Println(v5, v5)
        }
        _ = v18
    }
    v27 := GenS1{F0: 2147483647, F1: 2}
    if len(v6) > 0 {
        v6[0] = 32767
    }
    if true {
        switch "\n\t" {
        case "":
            if (127 <= 20) {
                if false {
                    switch "" {
                    case "hello":
                        var v28 uint = 0
                        _ = v28
                    case "a":
                        var v29 uint64 = 968
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 10; fi30++ {
                        var v31 uint32 = 646
                        if ((false || true) && true) {
                            var v32 uint8 = 0
                            for fi33 := 0; fi33 < 19; fi33++ {
                                switch ("hello" + v5) {
                                case "a":
                                    v34 := GenS3{F0: 319, F1: fi30}
                                    _ = v34
                                case "\n\t":
                                    v32 = v32
                                    v32 -= 1
                                    var v35 int8 = 127
                                    _ = v35
                                }
                                v3[(v5 + ((("\n\t" + "") + v5) + (v1 + (v1 + "a"))))] = ""
                                _ = fi33
                            }
                            for fi36 := 0; fi36 < 18; fi36++ {
                                switch ("\n\t" + "foo bar") {
                                case "hello":
                                    v37 := GenS4{F0: 0, F1: 135}
                                    _ = v37
                                default:
                                    var v38 int16 = 0
                                    _ = v38
                                }
                                v3[v5] = v1
                                _ = fi36
                            }
                            _ = v32
                        }
                        for fi39 := 0; fi39 < 15; fi39++ {
                            v5 = v1
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            switch ((v5 + (((v1 + "hello") + "\n\t") + "hello")) + "hello") {
                            case "hello":
                                v5 = v5
                            case "a":
                                _ = len(v2)
                            case "foo bar":
                                v1 = v5
                            default:
                                if false {
                                    v31 = (v31 << ((42) & 63))
                                }
                            }
                            _ = fi39
                        }
                        _ = fi30
                        _ = v31
                    }
                    switch ("hello" + v1) {
                    case "hello":
                        v1 = ("" + "a")
                    default:
                        v5 = (v5 + v5)
                    }
                } else {
                    v5 = (("" + v5) + v1)
                    var v40 float64 = 1e10
                    v40 -= ((v40 / float64(0.1)) * 1e10)
                    _ = v40
                }
                for fi41 := 0; fi41 < 2; fi41++ {
                    v42 := GenS3{F0: 1, F1: fi41}
                    for fi43 := 0; fi43 < 9; fi43++ {
                        v44 := GenS5{F0: 1, F1: 42}
                        fmt.Println(v1, fi43, v1)
                        _ = fi43
                        _ = v44
                    }
                    _ = fi41
                    _ = v42
                }
            }
        case "a":
            _ = len(v4)
        case "hello":
            var v45 int64 = -9223372036854775808
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 int8 = 127
        if true {
            fmt.Println(v47, v47)
            var v48 int64 = 577
            _ = v48
        } else {
            if false {
                v47 = v47
                v49 := GenS6{F0: v5}
                var v50 float64 = (float64(0.0) / ((0.0 + float64(float32(0.0))) + (((1.5 + 3.14159) - -2.25) / 1.5)))
                _ = v49
                _ = v50
            } else {
                if false {
                    var v51 uint = 2
                    var v52 int16 = 1
                    _ = v51
                    _ = v52
                }
                var v53 int16 = 1
                var v54 int16 = v53
                _ = v53
                _ = v54
            }
            v6 = append(v6, (fi46 ^ 127))
            for fi55 := 0; fi55 < 1; fi55++ {
                v56 := GenS4{F0: 2, F1: 42}
                fmt.Println(v47, v5, v5)
                var v57 int8 = -1
                if len(v6) > 0 {
                    v6[0] = (fi46 / ((fi55) | 1))
                }
                _ = fi55
                _ = v56
                _ = v57
            }
        }
        if true {
            var v58 uint16 = 314
            for fi59 := 0; fi59 < 1; fi59++ {
                v60 := GenS7{F0: 18446744073709551615}
                v61 := GenS2{F0: 0}
                if (65535 > 930) {
                    v1 = ""
                    var v62 uint = 102
                    v63 := GenS8{F0: -464}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v64 := GenS9{F0: 42, F1: v62, F2: (v47 & (-(v47)))}
                    _ = v62
                    _ = v63
                    _ = v64
                }
                var v65 int16 = -32768
                v47 += (v47 - (-(v47)))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v66 int = -678
                var v67 uint16 = (v58 >> ((18446744073709551615) & 63))
                _ = fi59
                _ = v60
                _ = v61
                _ = v65
                _ = v66
                _ = v67
            }
            var v68 int32 = 463
            _ = v58
            _ = v68
        } else {
            var v69 uint32 = 2147483647
            _ = v69
        }
        _ = fi46
        _ = v47
    }
    switch 128 {
    case 32768:
        switch v1 {
        case "foo bar":
            var v70 string = v1
            _ = v70
        case "a":
            _ = len(v3)
        default:
            fmt.Println(v1, v1)
        }
    case -521:
        v71 := GenS10{F0: 42, F1: 18446744073709551615}
        _ = v71
    case 0:
        v72 := v4[v1]
        _ = v72
    }
    v73 := GenS8{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
