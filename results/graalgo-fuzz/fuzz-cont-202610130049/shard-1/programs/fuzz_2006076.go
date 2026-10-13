package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1000
    var v3 int = 32768
    v4 := []string{""}
    v5 := make(map[uint32]int)
    if false {
        v6 := GenS1{F0: 100, F1: 497, F2: 2}
        if false {
            v7 := GenS1{F0: 0, F1: 1, F2: 255}
            for fi8 := 0; fi8 < 20; fi8++ {
                switch "\n\t" {
                case "foo bar":
                    var v9 bool = false
                    _ = v9
                case "a":
                    v5[42] = v3
                case "":
                    v2 += 1
                }
                if true {
                    fmt.Println(v3)
                    var v10 uint32 = 2147483647
                    _ = v10
                } else {
                    var v11 uint8 = 0
                    if ((335 == 1) || (!((3.14159 + (-2.25 - (1e10 + -2.25))) != (((1.5 - 0.0) / 1.0) * (((-2.25 + -2.25) * (((1.5 / (0.1 + 0.1)) - ((((0.0 + ((1.5 / ((1e10 + (-2.25 * ((float64(3.14159) / float64(float32((1e10 * (1e10 - 1.5))))) * 1e10))) - 0.1)) / -2.25)) + ((-2.25 - (float64((1e10 / 1e10)) - 0.0)) / 1e10)) * 1e10) / ((((((1.5 + 1.5) + 0.0) + 0.1) * ((float64(float32(3.14159)) / 1e10) / 1e10)) + (((0.0 + -2.25) * float64(float32((((3.14159 / (float64((0.1 * 3.14159)) / ((float64((((-2.25 / (1e10 * 0.1)) / float32(-2.25)) - 3.14159)) + (0.0 * (float64(float32(0.0)) + -2.25))) / (-2.25 / -2.25)))) + ((0.1 / (1.5 / 3.14159)) - (0.0 + 3.14159))) - 1.5)))) / float64(-2.25))) - -2.25))) * (0.0 * 0.0))) + (3.14159 * 1.5)))))) {
                        v11 -= (^(v11))
                    }
                    v2 = (^(v2))
                    _ = v11
                }
                if ((!true) && ((551 >= 199) && false)) {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi12
                    }
                    v13 := GenS2{F0: 2147483647}
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v15 := GenS1{F0: 127, F1: 13, F2: 2}
                        _ = fi14
                        _ = v15
                    }
                    _ = v13
                }
                _ = fi8
            }
            _ = v7
        } else {
            var v16 int = (-(v3))
            if false {
                var v17 int = v16
                for fi18 := 0; fi18 < 6; fi18++ {
                    if (128 == 255) {
                        if (!(1.5 == (1.5 - (((0.0 / (0.1 - (3.14159 - ((3.14159 * ((0.1 / (((((((1e10 - 1e10) * 0.1) - 0.1) + 0.1) - (-2.25 + ((1.5 * 1e10) * 3.14159))) * (0.0 - 1e10)) / float64(1.5))) - 1.5)) + 0.0)))) + ((((1.5 / (1.5 + 3.14159)) / 1.5) * ((((-2.25 / (1.5 - 1e10)) * 0.1) / 1.0) - 3.14159)) + (1.5 - (float64(3.14159) - 0.1)))) / -2.25)))) {
                            fmt.Println(v2)
                            fmt.Println(v2, v17)
                            var v19 uint32 = 2
                            _ = v19
                        }
                        switch fi18 {
                        case 2147483647:
                            for fi20 := 0; fi20 < 1; fi20++ {
                                var v21 uint64 = 18446744073709551615
                                v22 := GenS1{F0: -95, F1: 2, F2: uint8(v21)}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case 32767:
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v17)
                            case "a":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                var v23 uint64 = 502
                                _ = v23
                            }
                        case 128:
                            v24 := v5[0]
                            _ = v24
                        }
                    } else {
                        _ = len(v5)
                        v2 += (v2 >> ((4294967295) & 63))
                    }
                    v25 := GenS2{F0: 1}
                    _ = fi18
                    _ = v25
                }
                _ = v17
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v3 += (v16 | (v16 / ((v16) | 1)))
                v26 := GenS2{F0: 2147483647}
                _ = v26
            }
            _ = v16
        }
        _ = v6
    }
    v3 = v3
    fmt.Println(v2)
    var v27 string = "a"
    v4 = append(v4, (v27 + (("hello" + "foo bar") + v27)))
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 float32 = 0.0
            _ = fi29
            _ = v30
        }
        var v31 int = -938
        _ = fi28
        _ = v31
    }
    if ((1 >= v3) && (-60 < 1)) {
        v2 += v2
        v32 := GenS1{F0: -20, F1: 255, F2: 55}
        var v33 float32 = -2.25
        _ = v32
        _ = v33
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v2, v3)
        var v34 int16 = (v2 ^ v2)
        _ = v34
    }
    var v35 int = v3
    switch v35 {
    case 32768:
        var v36 int32 = 42
        _ = v36
    case 42:
        if true {
            for fi37 := 0; fi37 < 2; fi37++ {
                v2 = (v2 | -1)
                v38 := GenS1{F0: 42, F1: 255, F2: 229}
                _ = fi37
                _ = v38
            }
            v39 := GenS2{F0: -678}
            v40 := GenS2{F0: -1}
            _ = v39
            _ = v40
        } else {
            switch ((v27 + v27) + "") {
            case "\n\t":
                v41 := GenS1{F0: 1, F1: 880, F2: 247}
                _ = v41
            case "a":
                v42 := GenS1{F0: -1, F1: 964, F2: 255}
                _ = v42
            default:
                var v43 uint = 18446744073709551615
                _ = v43
            }
        }
    }
    switch v35 {
    case 2147483647:
        for fi44 := 0; fi44 < 8; fi44++ {
            if (1 != 0) {
                for fi45 := 0; fi45 < 9; fi45++ {
                    var v46 int64 = 9223372036854775807
                    _ = fi45
                    _ = v46
                }
                _ = len(v5)
            } else {
                var v47 int8 = 127
                v27 = (v27 + (v27 + v27))
                _ = v47
            }
            for fi48 := 0; fi48 < 8; fi48++ {
                if (true || (0 == -1)) {
                    _ = len(v4)
                    v4 = append(v4, v27)
                }
                var v49 uint = 759
                _ = fi48
                _ = v49
            }
            _ = fi44
        }
    }
    _ = len(v5)
    var v50 int = v3
    v27 = v27
    v2 -= 437
    switch (v50 * -22) {
    case -802:
        var v51 int16 = (^(v2))
        _ = v51
    case -274:
        if ((-9223372036854775808 == 0) && ((v2 | v2) <= v2)) {
            var v52 int64 = 42
            v50 = 128
            _ = v52
        }
    case -545:
        switch "hello" {
        case "\n\t":
            fmt.Println(v3, v27)
        case "":
            v53 := GenS2{F0: 42}
            _ = v53
        case "foo bar":
            v54 := GenS3{F0: 42, F1: 644}
            _ = v54
        default:
            var v55 float64 = 1e10
            _ = v55
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        if true {
            if true {
                if false {
                    v57 := GenS4{F0: (fi56 - (-(v50)))}
                    var v58 int8 = 100
                    v3 -= v50
                    _ = v57
                    _ = v58
                } else {
                    v27 = ("" + v27)
                }
                var v59 uint32 = 198
                _ = v59
            } else {
                var v60 uint8 = 1
                if ((false || (!false)) || false) {
                    var v61 int32 = 42
                    _ = v61
                } else {
                    var v62 float64 = 1e10
                    var v63 string = v27
                    _ = v62
                    _ = v63
                }
                for fi64 := 0; fi64 < 3; fi64++ {
                    var v65 float32 = (3.14159 - float32(float64(1.5)))
                    var v66 int32 = 42
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                _ = v60
            }
            v2 = (v2 & v2)
            v27 = ("" + "a")
        } else {
            switch "a" {
            case "\n\t":
                var v67 float32 = (1.5 * (((0.1 / (-2.25 / 1.5)) * ((0.0 * float32((-2.25 / 1.5))) * 1.5)) - 3.14159))
                _ = v67
            }
        }
        v2 += v2
        _ = fi56
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v50
}
