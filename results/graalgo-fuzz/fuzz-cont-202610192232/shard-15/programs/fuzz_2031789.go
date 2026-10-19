package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 127
    var v3 int = 128
    v4 := make(map[uint32]uint32)
    var v5 int = (v3 & -8)
    var v6 string = ("" + "\n\t")
    if (0.0 < (((((3.14159 - 3.14159) + float32((3.14159 * 0.0))) + 3.14159) + 1e10) + 0.0)) {
        var v7 int16 = -1
        var v8 float32 = 0.1
        var v9 int = (v5 & (^(v5)))
        _ = v7
        _ = v8
        _ = v9
    }
    v2 = v2
    v3 -= (v3 >> ((430) & 63))
    if (!(0 > 42)) {
        var v10 uint32 = 42
        if ((9223372036854775807 > -1) || false) {
            if ((1 < 1) && false) {
                fmt.Println(v3)
                v11 := GenS1{F0: 42, F1: -2147483648, F2: 42}
                _ = v11
            } else {
                v12 := []int{v3, (-(v5)), v5}
                v13 := GenS2{F0: (v10 * 0), F1: -1, F2: -9223372036854775808}
                _ = v12
                _ = v13
            }
            v4[(^(v10))] = (v10 - v10)
        }
        _ = len(v4)
        _ = v10
    }
    v6 = "foo bar"
    var v14 float64 = (((0.1 / ((1.5 * 1.5) - (float64(0.0) / ((float64(1e10) / 1e10) * ((1.5 + 1.5) * 1e10))))) / 1e10) / 1.0)
    v14 = ((v14 / ((float64(1.5) * v14) / v14)) / 0.1)
    v14 = (v14 + float64(3.14159))
    for fi15 := 0; fi15 < 4; fi15++ {
        if (!(true || false)) {
            fmt.Println(v6, v5, v6)
            switch v5 {
            case -2147483648:
                v4[247] = 939
            default:
                v14 = v14
            }
            var v16 int64 = 1
            _ = v16
        } else {
            v17 := GenS1{F0: -1, F1: -984, F2: 205}
            var v18 int64 = -1
            _ = v17
            _ = v18
        }
        v19 := GenS2{F0: 0, F1: -526, F2: int64(v2)}
        var v20 uint32 = 65535
        _ = fi15
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: 140}
    switch "a" {
    case "a":
        switch v6 {
        case "hello":
            switch v3 {
            case 963:
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (1e10 == (v14 + float64(float32(v14)))) {
                        v23 := GenS4{F0: 65535, F1: 1, F2: ((v6 + v6) + "")}
                        v24 := GenS4{F0: 1, F1: 0, F2: v6}
                        fmt.Println(v6)
                        _ = v23
                        _ = v24
                    } else {
                        v1[v6] = (!false)
                    }
                    _ = fi22
                }
            default:
                for fi25 := 0; fi25 < 7; fi25++ {
                    v26 := GenS5{F0: v6, F1: 2}
                    if (false && (4294967295 >= 701)) {
                        var v27 uint32 = 0
                        var v28 uint32 = 1
                        switch ((v6 + v6) + ((((((v6 + v6) + ("foo bar" + (v6 + ""))) + v6) + ((v6 + v6) + v6)) + "") + "")) {
                        case "":
                            for fi29 := 0; fi29 < 8; fi29++ {
                                v28 = (^(v27))
                                fmt.Println(fi29)
                                _ = fi29
                            }
                        case "\n\t":
                            var v30 string = (v6 + "\n\t")
                            _ = v30
                        }
                        _ = v27
                        _ = v28
                    }
                    if true {
                        v31 := []int{(fi25 | v5)}
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        v14 = v14
                        _ = v31
                    }
                    _ = fi25
                    _ = v26
                }
            }
        case "a":
            switch v6 {
            case "hello":
                var v32 int64 = 0
                _ = v32
            }
        case "":
            v33 := GenS1{F0: 0, F1: -508, F2: 255}
            _ = v33
        }
    case "hello":
        switch (v5 | 9223372036854775807) {
        case 2147483647:
            fmt.Println(v2, v14)
        case -1:
            v6 = v6
        }
    case "\n\t":
        v2 += -112
    }
    switch v3 {
    case -2147483648:
        switch v6 {
        case "\n\t":
            v34 := GenS6{F0: 9223372036854775807, F1: v6}
            _ = v34
        case "":
            v35 := GenS5{F0: "a", F1: 698}
            _ = v35
        }
    case 32768:
        switch v6 {
        case "foo bar":
            v36 := GenS5{F0: v6, F1: 42}
            _ = v36
        case "hello":
            switch "a" {
            case "foo bar":
                switch v6 {
                case "\n\t":
                    v37 := []int{}
                    _ = v37
                case "foo bar":
                    switch (v3 / ((v5) | 1)) {
                    case -606:
                        for fi38 := 0; fi38 < 19; fi38++ {
                            for fi39 := 0; fi39 < 16; fi39++ {
                                fmt.Println(fi39, v3)
                                fmt.Println(v5, v5, v6)
                                if true {
                                    if false {
                                        switch "" {
                                        case "\n\t":
                                            var v40 int8 = v2
                                            _ = v40
                                        case "a":
                                            var v41 uint64 = 42
                                            _ = v41
                                        case "foo bar":
                                            fmt.Println(v2)
                                        default:
                                            v5 += 32768
                                        }
                                        v42 := GenS7{F0: 874}
                                        switch ("foo bar" + (v6 + (v6 + ("foo bar" + (((v6 + v6) + v6) + "\n\t"))))) {
                                        case "a":
                                            v43 := GenS2{F0: 65535, F1: -32768, F2: 42}
                                            _ = v43
                                        }
                                        var v44 uint64 = 116
                                        _ = v42
                                        _ = v44
                                    } else {
                                        v45 := GenS2{F0: 2147483648, F1: int16(v5), F2: 613}
                                        v2 += 127
                                        _ = v45
                                    }
                                    var v46 uint8 = 0
                                    v5 = v3
                                    _ = v46
                                }
                                _ = fi39
                            }
                            _ = fi38
                        }
                    case 128:
                        v1[v6] = (-128 < 127)
                    case 32767:
                        v2 += (-(v2))
                    default:
                        var v47 int8 = 1
                        _ = v47
                    }
                default:
                    v3 = 1
                }
            case "hello":
                if (831 > 32767) {
                    for fi48 := 0; fi48 < 10; fi48++ {
                        v49 := GenS3{F0: 127}
                        var v50 float32 = ((0.1 + (1.5 - 0.1)) + ((1e10 / 1e10) * 1e10))
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                }
            }
        case "\n\t":
            var v51 bool = (42 > 255)
            _ = v51
        default:
            for fi52 := 0; fi52 < 11; fi52++ {
                fmt.Println(v14, v14, v3)
                v6 = v6
                var v53 bool = false
                _ = fi52
                _ = v53
            }
        }
    case -974:
        var v54 float32 = (-2.25 - (0.0 / -2.25))
        _ = v54
    default:
        switch (v3 << ((25) & 63)) {
        case 9223372036854775807:
            var v55 int32 = 257
            _ = v55
        case 2147483647:
            var v56 bool = (true && (32767 != 0))
            _ = v56
        case 1:
            for fi57 := 0; fi57 < 7; fi57++ {
                if (v2 == (v2 + v2)) {
                    v4[2] = 4294967295
                    if ((((false || false) || false) && true) || false) {
                        _ = len(v4)
                        v58 := GenS6{F0: 9223372036854775807, F1: "hello"}
                        if false {
                            v1[("\n\t" + (v6 + "foo bar"))] = true
                        }
                        _ = v58
                    }
                } else {
                    for fi59 := 0; fi59 < 17; fi59++ {
                        v60 := v4[1000]
                        v61 := GenS4{F0: 42, F1: 18446744073709551615, F2: (v6 + (v6 + v6))}
                        if false {
                            v62 := GenS2{F0: (v60 * 0), F1: int16(v3), F2: 42}
                            if true {
                                fmt.Println(v5, fi57, fi59)
                                switch (-(fi57)) {
                                case -2147483648:
                                    v2 = (-(v2))
                                case 2147483647:
                                    v14 -= v14
                                case 4294967295:
                                    var v63 int64 = -1
                                    _ = v63
                                default:
                                    v64 := GenS1{F0: 517, F1: -621, F2: 23}
                                    _ = v64
                                }
                                switch (v6 + "hello") {
                                case "":
                                    v65 := []int{103, v5, fi59}
                                    _ = v65
                                case "\n\t":
                                    var v66 uint = 413
                                    _ = v66
                                case "a":
                                    fmt.Println(v5)
                                }
                            }
                            v67 := GenS3{F0: uint8(v60)}
                            _ = v62
                            _ = v67
                        }
                        var v68 float32 = 3.14159
                        _ = fi59
                        _ = v60
                        _ = v61
                        _ = v68
                    }
                }
                _ = fi57
            }
        }
    }
    for fi69 := 0; fi69 < 2; fi69++ {
        v14 += float64(1e10)
        _ = fi69
    }
    var v70 int64 = 9223372036854775807
    var v71 uint16 = 42
    var v72 bool = (1 != 2)
    var v73 float64 = (((0.0 / v14) * 1e10) + (v14 / v14))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v72
    _ = v73
}
