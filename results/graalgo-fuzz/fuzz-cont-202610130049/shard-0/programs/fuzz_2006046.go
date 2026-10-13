package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0(p0 int8) string {
    v4 := GenS1{F0: 651, F1: 128, F2: 9223372036854775807}
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{}
        var v7 uint8 = 255
        var v8 uint64 = 18446744073709551615
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v4
    return "\n\t"
}

func main() {
    var v1 float32 = 0.0
    var v2 uint64 = 4294967295
    v3 := []int{}
    if true {
        var v9 int16 = 42
        v9 -= 1
        _ = v9
    }
    fmt.Println(v1)
    _ = len(v3)
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 bool = (42 != 2)
        v11 = (v11 && v11)
        var v12 float32 = (((v1 - 1.5) / 0.1) + (((v1 - float32(0.0)) - 3.14159) - (((float32(0.1) / 1.5) * v1) * v1)))
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            switch "" {
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                var v15 int = 479
                _ = v15
            case "\n\t":
                var v16 uint = 65535
                _ = v16
            }
            fmt.Println(fi14)
            var v17 int32 = 1
            v1 -= v1
            _ = fi14
            _ = v17
        }
        switch "foo bar" {
        case "\n\t":
            v2 = 787
        case "hello":
            var v18 int8 = -33
            _ = v18
        case "a":
            v2 = v2
        }
        _ = fi13
    }
    if true {
        if ((!((1 <= -1) && (-1 < 42))) || (!false)) {
            fmt.Println(v2, v1)
            var v19 uint = 646
            var v20 int32 = -508
            _ = v19
            _ = v20
        } else {
            var v21 uint16 = 510
            if false {
                v22 := GenS2{F0: v2}
                switch "a" {
                case "":
                    if (true || ("a" < ("" + "hello"))) {
                        if true {
                            v3 = append(v3, 128)
                        } else {
                            var v23 float64 = (1.5 + 3.14159)
                            _ = v23
                        }
                    }
                }
                if false {
                    if (!false) {
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v25 := GenS2{F0: (v2 | v2)}
                            v26 := GenS2{F0: v2}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    } else {
                        var v27 uint = 1
                        if (!(("\n\t" + ("" + "a")) > "foo bar")) {
                            if len(v3) > 0 {
                                v3[0] = -910
                            }
                            for fi28 := 0; fi28 < 12; fi28++ {
                                for fi29 := 0; fi29 < 5; fi29++ {
                                    var v30 int = (fi28 << ((v27) & 63))
                                    if (!false) {
                                        var v31 bool = ((-2.25 / 1.5) > v1)
                                        _ = v31
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                                _ = fi28
                            }
                            var v32 float64 = 1.5
                            _ = v32
                        }
                        fmt.Println(v2, v2)
                        _ = v27
                    }
                    v2 -= (v2 + (^(v2)))
                } else {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        for fi34 := 0; fi34 < 9; fi34++ {
                            v2 += (v2 >> ((53) & 63))
                            if true {
                                v2 -= (v2 - 605)
                            } else {
                                var v35 uint8 = 145
                                _ = v35
                            }
                            if len(v3) > 0 {
                                v3[0] = fi33
                            }
                            _ = fi34
                        }
                        switch fi33 {
                        case 9223372036854775807:
                            if (!((fi33 << ((222) & 63)) == fi33)) {
                                v36 := GenS1{F0: 42, F1: 255, F2: 578}
                                v2 = v2
                                v21 -= (v21 % ((v21) | 1))
                                _ = v36
                            } else {
                                for fi37 := 0; fi37 < 2; fi37++ {
                                    var v38 int8 = 22
                                    _ = fi37
                                    _ = v38
                                }
                            }
                        case 32768:
                            v1 = ((0.1 - v1) * float32(-2.25))
                        default:
                            switch "hello" {
                            case "":
                                switch ("" + (("a" + (("foo bar" + "\n\t") + ("\n\t" + ("foo bar" + "hello")))) + "\n\t")) {
                                case "":
                                    v1 = ((((v1 + (v1 / (v1 / 1.5))) - v1) * float32((-2.25 - 0.0))) * float32(0.1))
                                case "hello":
                                    v39 := GenS2{F0: v2}
                                    _ = v39
                                default:
                                    if false {
                                        switch -277 {
                                        case -829:
                                            v40 := GenS1{F0: int16(v2), F1: 0, F2: 72}
                                            _ = v40
                                        default:
                                            fmt.Println(v1, fi33, v21)
                                        }
                                        var v41 float32 = v1
                                        _ = v41
                                    }
                                }
                            }
                        }
                        v42 := GenS3{F0: fi33}
                        _ = fi33
                        _ = v42
                    }
                    fmt.Println(v21, v1)
                    fmt.Println(v2)
                }
                _ = v22
            }
            _ = v21
        }
        v2 = (^(v2))
    }
    fmt.Println(v2, v2)
    var v43 int = 2147483647
    v44 := GenS2{F0: 0}
    if ((false || true) && (0.0 != 1.5)) {
        for fi45 := 0; fi45 < 9; fi45++ {
            if len(v3) > 0 {
                v3[0] = v43
            }
            for fi46 := 0; fi46 < 6; fi46++ {
                var v47 int = v43
                _ = fi46
                _ = v47
            }
            _ = fi45
        }
    } else {
        switch ("a" + "foo bar") {
        case "foo bar":
            if false {
                var v48 float32 = v1
                _ = v48
            } else {
                var v49 int16 = 32767
                _ = v49
            }
        case "":
            for fi50 := 0; fi50 < 13; fi50++ {
                v51 := GenS4{F0: -1}
                if false {
                    if len(v3) > 0 {
                        v3[0] = 9223372036854775807
                    }
                    v1 = v1
                    var v52 uint64 = (v2 >> ((687) & 63))
                    _ = v52
                } else {
                    var v53 int16 = -1
                    _ = v53
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi50
                _ = v51
            }
        case "hello":
            v54 := GenS5{F0: 70, F1: 255, F2: uint32(v2)}
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        if (2147483648 <= 475) {
            if false {
                switch (fi55 & fi55) {
                case -1:
                    if true {
                        v56 := GenS1{F0: -32768, F1: 1, F2: 1}
                        _ = len(v3)
                        v1 += v1
                        _ = v56
                    } else {
                        for fi57 := 0; fi57 < 11; fi57++ {
                            var v58 bool = (((v1 + 1.5) * (((3.14159 + (1.5 - (0.1 - v1))) / v1) - 1e10)) <= 1.5)
                            _ = fi57
                            _ = v58
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 = (v2 - v2)
                    }
                }
            }
            var v59 uint16 = 2
            _ = v59
        } else {
            v3 = append(v3, v43)
        }
        for fi60 := 0; fi60 < 19; fi60++ {
            switch fi55 {
            case 600:
                var v61 uint64 = (v2 & (v2 >> ((2) & 63)))
                v62 := GenS4{F0: 127}
                var v63 uint = 375
                _ = v61
                _ = v62
                _ = v63
            case 42:
                var v64 uint16 = 387
                _ = v64
            default:
                switch "foo bar" {
                case "hello":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 = (^(v2))
                default:
                    if (!true) {
                        var v65 string = "foo bar"
                        _ = v65
                    } else {
                        v1 = (v1 * v1)
                    }
                }
            }
            _ = fi60
        }
        _ = fi55
    }
    fmt.Println(v1, v43, v2)
    switch "foo bar" {
    case "foo bar":
        v43 = (v43 / ((957) | 1))
    case "hello":
        var v66 uint = 0
        _ = v66
    case "a":
        for fi67 := 0; fi67 < 13; fi67++ {
            v43 = fi67
            var v68 float32 = -2.25
            if (v43 < -1) {
                switch 467 {
                case 32768:
                    switch ("foo bar" + (("a" + "hello") + "")) {
                    case "hello":
                        var v69 uint8 = 0
                        _ = v69
                    case "\n\t":
                        var v70 float32 = ((v1 * v1) / (v68 - (1.5 - float32(float64((v1 - (-2.25 / 0.1)))))))
                        _ = v70
                    default:
                        v71 := GenS5{F0: 51, F1: 2, F2: 719}
                        _ = v71
                    }
                case -1:
                    if true {
                        _ = len(v3)
                        v72 := GenS4{F0: 42}
                        switch (fi67 * v43) {
                        case 127:
                            switch ("hello" + "foo bar") {
                            case "a":
                                v73 := GenS4{F0: -1}
                                _ = v73
                            case "":
                                v1 = 3.14159
                            default:
                                if false {
                                    v74 := GenS2{F0: (v2 + (v2 >> ((75) & 63)))}
                                    if len(v3) > 0 {
                                        v3[0] = (^(fi67))
                                    }
                                    _ = v74
                                }
                            }
                        case 1:
                            var v75 int32 = 0
                            _ = v75
                        default:
                            if true {
                                v68 -= 3.14159
                                v43 += (v43 * 4294967295)
                            } else {
                                var v76 uint32 = 2147483647
                                v1 = 0.1
                                _ = v76
                            }
                        }
                        _ = v72
                    }
                case -2147483648:
                    var v77 int16 = 1
                    _ = v77
                default:
                    v1 = ((v68 / (float32((1e10 * 0.0)) / v1)) / ((v68 / v68) + 1.5))
                }
                v78 := GenS1{F0: 275, F1: 255, F2: 601}
                fmt.Println(v43, v43, v43)
                _ = v78
            }
            _ = fi67
            _ = v68
        }
    default:
        var v79 int64 = 9223372036854775807
        _ = v79
    }
    v1 -= (v1 / (v1 - v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
}
