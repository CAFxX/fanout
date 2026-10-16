package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint64; F1 int32; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 string; F2 int8 }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]bool)
    v3 := []int{32768}
    v4 := []uint32{2147483648, 78, 42, 42}
    var v5 int = -439
    var v6 uint = 563
    v5 -= v5
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 81, F1: 964, F2: 0}
            for fi9 := 0; fi9 < 19; fi9++ {
                if (true && (!false)) {
                    var v10 int = -1
                    switch ("hello" + ("" + ("" + ("\n\t" + (("hello" + ("foo bar" + ("\n\t" + "hello"))) + ("\n\t" + (("a" + "") + "a"))))))) {
                    case "a":
                        _ = len(v4)
                    default:
                        v1 -= v1
                    }
                    v11 := GenS1{F0: 128, F1: -171, F2: 116}
                    var v12 int64 = 264
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v13 := GenS2{F0: 1, F1: 381}
                _ = fi9
                _ = v13
            }
            _ = fi7
            _ = v8
        }
        var v14 uint16 = (v1 | 32767)
        switch (v5 << ((1) & 63)) {
        case 127:
            if ((float64(-2.25) - float64(0.1)) == 0.0) {
                v15 := GenS3{F0: 1}
                var v16 int = v5
                switch "\n\t" {
                case "foo bar":
                    v6 = 0
                case "":
                    v17 := GenS4{F0: 18446744073709551615, F1: int64(v1)}
                    _ = v17
                }
                _ = v15
                _ = v16
            }
        case -68:
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint = v6
                switch (^(fi18)) {
                case -2147483648:
                    v20 := GenS2{F0: 2147483647, F1: -1}
                    _ = v20
                case 9223372036854775807:
                    _ = len(v3)
                case -12:
                    switch (fi18 * (^(v5))) {
                    case -9223372036854775808:
                        if false {
                            v6 = (v19 | 0)
                            v21 := GenS5{F0: 0, F1: 2147483647, F2: 32767}
                            var v22 int32 = 2147483647
                            _ = v21
                            _ = v22
                        } else {
                            v23 := GenS3{F0: 63}
                            for fi24 := 0; fi24 < 3; fi24++ {
                                if (true && false) {
                                    _ = len(v2)
                                    v25 := GenS6{F0: -1}
                                    v1 = v14
                                    _ = v25
                                }
                                _ = fi24
                            }
                            _ = v23
                        }
                    default:
                        fmt.Println(v1)
                    }
                default:
                    v14 = (v1 & v1)
                }
                _ = fi18
                _ = v19
            }
        case 32767:
            if false {
                if false {
                    var v26 float64 = 0.1
                    _ = v26
                } else {
                    var v27 int16 = 32767
                    _ = v27
                }
                v14 = 42
            }
        }
        _ = v14
    }
    switch (^(v5)) {
    case 32767:
        v1 = v1
    case 32768:
        fmt.Println(v6, v6, v1)
    case 42:
        for fi28 := 0; fi28 < 16; fi28++ {
            for fi29 := 0; fi29 < 3; fi29++ {
                fmt.Println(fi29, v6)
                var v30 uint32 = 739
                if len(v3) > 0 {
                    v3[0] = int(v30)
                }
                _ = fi29
                _ = v30
            }
            fmt.Println(v5, v6, v1)
            _ = fi28
        }
    }
    if (455 > 0) {
        fmt.Println(v1)
        v2[42] = false
        var v31 int = v5
        _ = v31
    } else {
        for fi32 := 0; fi32 < 3; fi32++ {
            switch ("hello" + ("foo bar" + ("a" + "foo bar"))) {
            case "foo bar":
                v6 -= (v6 / ((v6) | 1))
            case "hello":
                v33 := GenS2{F0: 42, F1: 1000}
                _ = v33
            default:
                var v34 int16 = 929
                _ = v34
            }
            v5 = v5
            switch 32768 {
            case 0:
                switch 42 {
                case 1:
                    switch (v5 + v5) {
                    case -423:
                        v6 = 676
                    }
                case 2147483647:
                    v35 := v2[0]
                    _ = v35
                default:
                    switch v5 {
                    case 127:
                        var v36 float64 = -2.25
                        _ = v36
                    case -501:
                        if false {
                            if (true || (((true && (true && true)) || true) || false)) {
                                for fi37 := 0; fi37 < 4; fi37++ {
                                    v5 -= -2147483648
                                    if len(v4) > 0 {
                                        v4[0] = 2
                                    }
                                    v6 = (^(v6))
                                    switch "hello" {
                                    case "a":
                                        v38 := v2[4294967295]
                                        _ = v38
                                    }
                                    v39 := GenS4{F0: 2, F1: 0}
                                    _ = fi37
                                    _ = v39
                                }
                            } else {
                                for fi40 := 0; fi40 < 6; fi40++ {
                                    v1 -= (v1 - v1)
                                    v41 := GenS4{F0: 0, F1: -245}
                                    _ = fi40
                                    _ = v41
                                }
                                if false {
                                    var v42 int = fi32
                                    _ = v42
                                }
                                v1 = v1
                            }
                        }
                    case 2147483647:
                        var v43 int8 = 0
                        _ = v43
                    default:
                        fmt.Println(fi32)
                    }
                }
            }
            _ = fi32
        }
        var v44 string = "a"
        fmt.Println(v6, v1, v1)
        _ = v44
    }
    if (0.1 < (0.0 - 3.14159)) {
        for fi45 := 0; fi45 < 3; fi45++ {
            v5 += (-(fi45))
            for fi46 := 0; fi46 < 3; fi46++ {
                var v47 float64 = 0.0
                v48 := GenS5{F0: 4294967295, F1: -1, F2: 1}
                _ = fi46
                _ = v47
                _ = v48
            }
            _ = len(v2)
            _ = fi45
        }
        v5 = 2147483647
        v49 := GenS7{F0: -2147483648, F1: "\n\t", F2: 127}
        _ = v49
    } else {
        fmt.Println(v6, v5)
        var v50 uint64 = 4294967295
        switch ("foo bar" + ("\n\t" + "a")) {
        case "\n\t":
            var v51 float32 = (-2.25 * 0.1)
            _ = v51
        default:
            if false {
                var v52 bool = (((false || (128 > v5)) || (0 < 65535)) || (!false))
                _ = v52
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v53 int8 = -28
                _ = v53
            }
        }
        _ = v50
    }
    switch "a" {
    case "a":
        fmt.Println(v5, v6, v6)
    default:
        v6 += (v6 + (v6 + (^(v6))))
    }
    fmt.Println(v1)
    for fi54 := 0; fi54 < 20; fi54++ {
        for fi55 := 0; fi55 < 3; fi55++ {
            fmt.Println(fi54, v6)
            _ = fi55
        }
        for fi56 := 0; fi56 < 13; fi56++ {
            v57 := GenS2{F0: int32(v6), F1: 1000}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v58 uint64 = 42
            _ = fi56
            _ = v57
            _ = v58
        }
        v1 -= (v1 ^ v1)
        _ = fi54
    }
    var v59 uint16 = (v1 & (v1 | v1))
    v1 += (v59 % (((^(v1))) | 1))
    switch (("foo bar" + "") + "hello") {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 127
        }
    case "hello":
        for fi60 := 0; fi60 < 20; fi60++ {
            v61 := GenS1{F0: 105, F1: 132, F2: 42}
            for fi62 := 0; fi62 < 14; fi62++ {
                if ("\n\t" != "\n\t") {
                    v5 = fi60
                    _ = len(v3)
                    _ = len(v2)
                } else {
                    if true {
                        fmt.Println(fi62)
                    }
                    delete(v2, 2147483648)
                    fmt.Println(v1)
                }
                delete(v2, 958)
                _ = fi62
            }
            v5 -= (fi60 | (^(fi60)))
            _ = fi60
            _ = v61
        }
    case "":
        var v63 int8 = -101
        _ = v63
    }
    delete(v2, uint32(v6))
    v2[2147483648] = true
    var v64 int64 = 9223372036854775807
    v64 += (v64 << ((v6) & 63))
    if ((!(-1 >= -826)) && (!("\n\t" > ((("hello" + "a") + "") + "a")))) {
        if true {
            for fi65 := 0; fi65 < 9; fi65++ {
                for fi66 := 0; fi66 < 13; fi66++ {
                    fmt.Println(v59, v1, v5)
                    _ = fi66
                }
                _ = fi65
            }
            v6 = 18446744073709551615
        }
        fmt.Println(v64)
        switch 1 {
        case -1:
            if (!(387 < 674)) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v67 uint32 = 427
                _ = v67
            }
        case 42:
            v68 := GenS8{F0: 42}
            _ = v68
        case 487:
            var v69 uint = (^(v6))
            _ = v69
        default:
            var v70 int64 = (v64 ^ (v64 * -9223372036854775808))
            _ = v70
        }
    }
    v71 := GenS7{F0: 955, F1: "a", F2: int8(v64)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v64
    _ = v71
}
