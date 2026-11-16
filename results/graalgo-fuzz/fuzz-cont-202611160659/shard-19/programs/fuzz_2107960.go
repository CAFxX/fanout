package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int; F2 bool }

func main() {
    v1 := []uint32{591, 42, 2, 259}
    v2 := []uint32{42, 2, 1, 2147483647}
    var v3 uint64 = 1
    v4 := make(map[int]uint32)
    var v5 float32 = 3.14159
    var v6 uint = 42
    v7 := GenS1{F0: "", F1: 1}
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            v1 = append(v1, 206)
            var v9 uint8 = 2
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            if true {
                fmt.Println(fi10)
            }
            v11 := GenS1{F0: "a", F1: 520}
            var v12 uint32 = 42
            _ = fi10
            _ = v11
            _ = v12
        }
        if false {
            if (((-2.25 >= (-2.25 * ((1e10 / 1.0) - (0.0 - 0.0)))) && true) || (-2147483648 != int(v6))) {
                fmt.Println(v6)
                v6 = v6
                var v13 float32 = 0.1
                _ = v13
            }
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 uint = v6
                _ = fi14
                _ = v15
            }
            fmt.Println(v6, v6)
        }
    } else {
        v5 = ((((((-2.25 * v5) + 0.1) / -2.25) * 1e10) / v5) + (float32(1.5) - v5))
        fmt.Println(v5, v3, v5)
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    if false {
        var v16 int = -637
        v3 = 4294967295
        fmt.Println(v5)
        _ = v16
    } else {
        switch "\n\t" {
        case "foo bar":
            for fi17 := 0; fi17 < 8; fi17++ {
                switch fi17 {
                case 126:
                    var v18 string = ""
                    _ = v18
                case 0:
                    var v19 bool = (!true)
                    _ = v19
                case 1:
                    var v20 int8 = 1
                    _ = v20
                }
                _ = fi17
            }
        case "a":
            var v21 float64 = (3.14159 + float64(float32(3.14159)))
            _ = v21
        case "hello":
            if (!(370 >= 255)) {
                var v22 int = -2147483648
                _ = len(v4)
                var v23 uint = v6
                _ = v22
                _ = v23
            } else {
                v24 := GenS2{F0: 0}
                _ = v24
            }
        }
        v25 := v4[-870]
        switch "hello" {
        case "a":
            var v26 bool = (true || (32768 <= 1))
            _ = v26
        case "hello":
            v27 := GenS1{F0: "hello", F1: -430}
            _ = v27
        }
        _ = v25
    }
    if true {
        v6 -= (v6 - 18446744073709551615)
        var v28 int = -9223372036854775808
        if ((233 == 0) || false) {
            if (!((0.0 / ((v5 - v5) * v5)) <= ((v5 - (v5 / 3.14159)) + (v5 / (v5 * (1.5 - (v5 + v5))))))) {
                if true {
                    delete(v4, 122)
                    v3 -= (v3 ^ (v3 & (v3 >> (((v6 - v6)) & 63))))
                }
            } else {
                for fi29 := 0; fi29 < 3; fi29++ {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        if ((fi29 << ((v6) & 63)) > fi29) {
                            v3 += uint64(v6)
                            var v31 bool = (-21 > -1)
                            switch (fi29 ^ (fi30 + (^(fi30)))) {
                            case 128:
                                var v32 uint32 = 2
                                _ = v32
                            case -98:
                                var v33 int16 = 859
                                _ = v33
                            default:
                                if (((-2.25 - 1.5) + 0.0) > -2.25) {
                                    v5 = (float32((-2.25 + 3.14159)) + v5)
                                    var v34 string = "hello"
                                    if (1 <= -1) {
                                        var v35 uint8 = 75
                                        v3 = v3
                                        _ = v35
                                    }
                                    var v36 float64 = 0.1
                                    var v37 uint16 = 0
                                    switch (fi29 << ((v6) & 63)) {
                                    case 604:
                                        var v38 int64 = 0
                                        _ = v38
                                    default:
                                        v4[32768] = 2147483647
                                    }
                                    _ = v34
                                    _ = v36
                                    _ = v37
                                } else {
                                    v3 = v3
                                    v39 := GenS3{F0: 0, F1: (938 != 255), F2: v3}
                                    switch "a" {
                                    case "foo bar":
                                        var v40 int8 = 42
                                        _ = v40
                                    }
                                    switch fi29 {
                                    case -381:
                                        v6 -= 42
                                    case 2147483647:
                                        var v41 uint = v6
                                        _ = v41
                                    case -260:
                                        var v42 uint32 = 2147483648
                                        _ = v42
                                    default:
                                        var v43 uint64 = 2
                                        _ = v43
                                    }
                                    var v44 int8 = -128
                                    _ = v39
                                    _ = v44
                                }
                            }
                            _ = v31
                        }
                        switch (v28 + fi30) {
                        case 32768:
                            var v45 string = "hello"
                            _ = v45
                        case -2147483648:
                            var v46 int16 = -1
                            _ = v46
                        case 605:
                            fmt.Println(v28, v5, v6)
                        }
                        var v47 bool = (267 < 0)
                        _ = fi30
                        _ = v47
                    }
                    if (false && ("a" == "\n\t")) {
                        v48 := GenS4{F0: -2147483648}
                        var v49 bool = false
                        if v49 {
                            var v50 uint64 = (^(v3))
                            v51 := GenS2{F0: -1}
                            var v52 int64 = 1
                            _ = v50
                            _ = v51
                            _ = v52
                        } else {
                            var v53 string = "\n\t"
                            v28 = -1
                            _ = v53
                        }
                        _ = v48
                        _ = v49
                    } else {
                        fmt.Println(v3, fi29)
                    }
                    _ = fi29
                }
                v54 := GenS5{F0: 127}
                _ = v54
            }
            for fi55 := 0; fi55 < 14; fi55++ {
                if true {
                    fmt.Println(v28)
                }
                switch -308 {
                case 9223372036854775807:
                    var v56 uint8 = 58
                    _ = v56
                case 127:
                    var v57 int32 = -166
                    _ = v57
                }
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
                _ = fi55
            }
            v58 := GenS2{F0: 19}
            _ = v58
        }
        _ = v28
    } else {
        v5 -= v5
    }
    v6 = 255
    v6 = (v6 * 42)
    switch 9223372036854775807 {
    case -9223372036854775808:
        for fi59 := 0; fi59 < 17; fi59++ {
            v3 = v3
            var v60 int = fi59
            _ = fi59
            _ = v60
        }
    case 939:
        for fi61 := 0; fi61 < 8; fi61++ {
            var v62 int32 = 837
            if false {
                v6 = (v6 + v6)
                var v63 int8 = 0
                _ = v63
            }
            v64 := GenS6{F0: 776, F1: (fi61 - 9223372036854775807), F2: ((9223372036854775807 >= 0) || (false || (0 <= 255)))}
            _ = fi61
            _ = v62
            _ = v64
        }
    default:
        var v65 bool = (false || (!true))
        _ = v65
    }
    var v66 uint32 = 545
    if false {
        for fi67 := 0; fi67 < 17; fi67++ {
            fmt.Println(v5, fi67, v6)
            var v68 uint8 = 127
            _ = fi67
            _ = v68
        }
    } else {
        v6 = (v6 - uint(v66))
        var v69 string = "a"
        _ = len(v1)
        _ = v69
    }
    for fi70 := 0; fi70 < 3; fi70++ {
        v71 := v4[(-(fi70))]
        _ = fi70
        _ = v71
    }
    var v72 uint = (^(v6))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v7
    _ = v72
}
