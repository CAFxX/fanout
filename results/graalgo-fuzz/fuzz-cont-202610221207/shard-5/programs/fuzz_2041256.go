package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0(p0 int8, p1 int, p2 int8) uint32 {
    var v7 float32 = 1.5
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: p0, F1: 42}
        _ = fi8
        _ = v9
    }
    var v10 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 451
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 2147483647
    v3 := []uint32{2147483648, 349, 362, 2147483648}
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]uint32)
    v3 = append(v3, 1)
    var v11 float64 = (float64(3.14159) + 0.1)
    v3 = append(v3, 65535)
    for fi12 := 0; fi12 < 3; fi12++ {
        v2 = (v2 | (-(v2)))
        v13 := GenS1{F0: 85, F1: v2}
        _ = fi12
        _ = v13
    }
    _ = len(v4)
    if (false || (65535 < 65535)) {
        v14 := GenS1{F0: 100, F1: (v2 << ((644) & 63))}
        if true {
            var v15 uint64 = 0
            _ = v15
        }
        _ = v14
    } else {
        for fi16 := 0; fi16 < 10; fi16++ {
            if true {
                if ((fi16 < 2147483647) || ("a" > "")) {
                    if false {
                        if (4294967295 < 128) {
                            fmt.Println(fi16, fi16, fi16)
                            v2 += 32768
                        } else {
                            var v17 float64 = v11
                            fmt.Println(v17)
                            for fi18 := 0; fi18 < 2; fi18++ {
                                var v19 uint = 0
                                if false {
                                    _ = len(v4)
                                    var v20 uint16 = 908
                                    _ = v20
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                            _ = v17
                        }
                        var v21 int8 = -1
                        var v22 uint64 = 18446744073709551615
                        _ = v21
                        _ = v22
                    }
                } else {
                    v11 += (v11 - v11)
                    v2 = (v2 % ((-2147483648) | 1))
                    v11 -= v11
                }
                var v23 uint16 = 672
                var v24 int64 = -9223372036854775808
                _ = v23
                _ = v24
            } else {
                switch 9223372036854775807 {
                case 2147483647:
                    var v25 int8 = 127
                    _ = v25
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v26 uint = 863
            var v27 int32 = -2147483648
            _ = fi16
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            if true {
                if (v11 == v11) {
                    var v29 float32 = 0.1
                    _ = v29
                }
                _ = len(v6)
            } else {
                v30 := GenS2{F0: false}
                var v31 float32 = (((0.0 / 1.0) * float32(v11)) / 3.14159)
                v32 := GenS3{F0: 432}
                _ = v30
                _ = v31
                _ = v32
            }
            var v33 uint = 65535
            var v34 uint64 = 42
            _ = fi28
            _ = v33
            _ = v34
        }
    }
    switch (v2 % (((^(v2))) | 1)) {
    case 2147483647:
        switch ("" + (("" + "") + "a")) {
        case "foo bar":
            if false {
                var v35 int16 = 42
                if (false && (4294967295 >= 42)) {
                    v2 -= 127
                }
                _ = v35
            } else {
                fmt.Println(v2)
                v36 := GenS4{F0: ("foo bar" + (("foo bar" + "foo bar") + "foo bar")), F1: 4294967295, F2: -1}
                _ = v36
            }
        case "":
            if (((342 != 42) && false) && true) {
                v37 := GenS2{F0: (genFn0(-1, (^(v2)), -69) != 1)}
                fmt.Println(v2, v2)
                fmt.Println(v2, v2)
                _ = v37
            }
        default:
            for fi38 := 0; fi38 < 3; fi38++ {
                for fi39 := 0; fi39 < 1; fi39++ {
                    var v40 uint64 = 0
                    var v41 float64 = ((v11 + (v11 + v11)) - (0.0 - v11))
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                _ = fi38
            }
        }
    }
    v11 = (0.0 + (v11 / v11))
    switch (^(v2)) {
    case 127:
        fmt.Println(v11, v2)
    case -2147483648:
        fmt.Println(v2, v2)
    case 1:
        var v42 float64 = v11
        _ = v42
    }
    _ = len(v4)
    v2 = v2
    if true {
        var v43 int64 = 191
        v43 = (v43 % ((v43) | 1))
        _ = v43
    }
    switch -1 {
    case 2147483647:
        if true {
            v44 := GenS5{F0: 1, F1: 26}
            if (2 <= 128) {
                if (1e10 < float64(0.1)) {
                    switch (("\n\t" + (("" + "a") + "foo bar")) + "hello") {
                    case "":
                        v11 = v11
                    case "\n\t":
                        var v45 uint8 = 167
                        _ = v45
                    }
                    if false {
                        v2 = (v2 / (((v2 >> ((255) & 63))) | 1))
                        if ((true && true) && false) {
                            v3 = append(v3, 2147483647)
                            v1[(v2 << ((0) & 63))] = "a"
                            v11 = v11
                        } else {
                            v11 += 1.5
                            v46 := GenS6{F0: -862, F1: 1, F2: 0}
                            _ = v46
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                    }
                    var v47 int8 = 15
                    _ = v47
                }
                v2 = (v2 * (v2 - 127))
            } else {
                switch "" {
                case "a":
                    for fi48 := 0; fi48 < 8; fi48++ {
                        for fi49 := 0; fi49 < 14; fi49++ {
                            var v50 uint32 = 832
                            v11 = ((float64(1e10) - v11) - 0.0)
                            fmt.Println(v2, v11)
                            var v51 bool = (0 != -128)
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        _ = fi48
                    }
                case "":
                    switch v2 {
                    case 9223372036854775807:
                        v11 = (float64(1.5) / 3.14159)
                    case -917:
                        v11 += (v11 * (v11 - float64(3.14159)))
                    }
                default:
                    fmt.Println(v2, v11)
                }
                v52 := v4[("\n\t" + "foo bar")]
                for fi53 := 0; fi53 < 6; fi53++ {
                    var v54 uint64 = 274
                    switch fi53 {
                    case 1:
                        var v55 int32 = 2147483647
                        _ = v55
                    case 616:
                        for fi56 := 0; fi56 < 18; fi56++ {
                            var v57 int8 = 0
                            _ = fi56
                            _ = v57
                        }
                    case 0:
                        v52 = (v2 ^ 9223372036854775807)
                    default:
                        v52 = (v2 % ((v2) | 1))
                    }
                    v54 = 42
                    _ = fi53
                    _ = v54
                }
                _ = v52
            }
            v2 = v2
            _ = v44
        }
    }
    for fi58 := 0; fi58 < 11; fi58++ {
        if (false && false) {
            v2 -= -2147483648
        } else {
            for fi59 := 0; fi59 < 20; fi59++ {
                v2 = 9223372036854775807
                var v60 uint64 = 108
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                _ = fi59
                _ = v60
            }
        }
        switch ("\n\t" + "a") {
        case "a":
            fmt.Println(v11, v2, v11)
        case "hello":
            v11 += 0.1
        }
        v11 = (0.1 - v11)
        _ = fi58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
