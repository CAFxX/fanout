package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 bool; F1 int32; F2 uint }
type GenS5 struct { F0 string; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int32; F1 int; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{42, 2147483648, 79, 1}
    var v3 int32 = -2147483648
    var v4 string = "foo bar"
    var v5 int = 128
    var v6 int8 = 127
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 870, F1: 245}
        v9 := GenS1{F0: 105, F1: 65535}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int8 = (v6 - (v6 ^ v6))
    if (!false) {
        var v11 int = (v5 >> ((2) & 63))
        _ = v11
    } else {
        var v12 int = (v5 >> ((255) & 63))
        if false {
            switch 1 {
            case 4294967295:
                fmt.Println(v4, v6)
            case -520:
                _ = len(v2)
            case 128:
                if len(v2) > 0 {
                    v2[0] = 173
                }
            default:
                for fi13 := 0; fi13 < 11; fi13++ {
                    if (true || (0.0 == float32(0.1))) {
                        var v14 int = v12
                        var v15 bool = (0 < 18446744073709551615)
                        _ = v14
                        _ = v15
                    }
                    _ = fi13
                }
            }
            switch (v5 * (v12 | v12)) {
            case 9223372036854775807:
                var v16 float64 = (((1e10 * (1.5 * 3.14159)) * 1e10) / 1.0)
                _ = v16
            case 32768:
                v2 = append(v2, 247)
            case 1:
                var v17 uint = 420
                _ = v17
            default:
                if (2 > 2) {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        fmt.Println(v12, v10, v5)
                        _ = fi18
                    }
                } else {
                    switch (v12 ^ v5) {
                    case 127:
                        var v19 int32 = (v3 & int32(v5))
                        _ = v19
                    case -1:
                        v6 -= (v6 << ((515) & 63))
                    case -9223372036854775808:
                        v10 -= v6
                    }
                    if (270 == -1) {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v12 += (-(v5))
                            v12 = -334
                            v21 := GenS1{F0: 42, F1: 669}
                            _ = fi20
                            _ = v21
                        }
                        _ = len(v2)
                    } else {
                        if (!false) {
                            for fi22 := 0; fi22 < 10; fi22++ {
                                _ = len(v2)
                                if (false && (((!false) || (651 <= 751)) || true)) {
                                    v23 := GenS2{F0: 42, F1: -2147483648}
                                    var v24 float64 = 0.0
                                    _ = v23
                                    _ = v24
                                } else {
                                    v1[4294967295] = 801
                                }
                                _ = fi22
                            }
                        }
                        if false {
                            for fi25 := 0; fi25 < 6; fi25++ {
                                v5 -= int(v6)
                                _ = fi25
                            }
                            _ = len(v1)
                        } else {
                            switch "foo bar" {
                            case "hello":
                                v12 -= v5
                            case "foo bar":
                                v26 := GenS1{F0: 722, F1: 990}
                                _ = v26
                            }
                            switch (v12 * (v12 % ((212) | 1))) {
                            case -1:
                                fmt.Println(v12, v10, v5)
                            case 4294967295:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case -2147483648:
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    v28 := GenS3{F0: int(v6), F1: 1}
                                    var v29 int16 = -620
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                var v30 int64 = 1
                                _ = v30
                            }
                        }
                    }
                    var v31 uint = 2
                    _ = v31
                }
            }
        } else {
            v1[2] = 391
            var v32 uint16 = 42
            _ = v32
        }
        fmt.Println(v5, v4, v4)
        _ = v12
    }
    v33 := GenS1{F0: 42, F1: 166}
    if false {
        var v34 uint8 = 50
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if (((3.14159 - (1e10 + 0.0)) + (1e10 / 0.1)) < (1e10 / 1.5)) {
                if false {
                    delete(v1, 4294967295)
                } else {
                    v1[2147483647] = 128
                    switch (v4 + v4) {
                    case "a":
                        switch v4 {
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "\n\t":
                            v10 = (v10 + (-(v10)))
                        case "foo bar":
                            switch v4 {
                            case "":
                                switch 9223372036854775807 {
                                case 754:
                                    switch v5 {
                                    case -783:
                                        v35 := GenS4{F0: (v5 != (^(v5))), F1: (v3 & (v3 - -542)), F2: 197}
                                        _ = v35
                                    case 127:
                                        v4 = "hello"
                                    default:
                                        v36 := GenS5{F0: ((("foo bar" + "\n\t") + v4) + v4), F1: false, F2: 158}
                                        _ = v36
                                    }
                                case 42:
                                    if true {
                                        v37 := v1[65535]
                                        if ((42 != 154) && true) {
                                            fmt.Println(v6, v37, v34)
                                            for fi38 := 0; fi38 < 18; fi38++ {
                                                v6 -= (v10 << ((0) & 63))
                                                _ = fi38
                                            }
                                            v37 = (v5 / ((v37) | 1))
                                        }
                                        v34 = v34
                                        _ = v37
                                    }
                                default:
                                    var v39 bool = (v4 != (("\n\t" + ((("" + (v4 + v4)) + (v4 + v4)) + ("\n\t" + ""))) + v4))
                                    _ = v39
                                }
                            case "a":
                                v40 := GenS3{F0: v5, F1: (v34 - v34)}
                                _ = v40
                            }
                        }
                    case "\n\t":
                        v6 += (v10 >> ((42) & 63))
                    case "hello":
                        v34 += (v34 | v34)
                    default:
                        v6 = (v6 + (v6 / ((v6) | 1)))
                    }
                    v2 = append(v2, 2147483648)
                }
                var v41 uint32 = 4294967295
                _ = v41
            } else {
                v42 := GenS6{F0: 42, F1: v34}
                v5 = 0
                _ = v42
            }
        }
        _ = v34
    }
    var v43 float32 = 1.5
    switch "a" {
    case "hello":
        fmt.Println(v3, v5, v6)
    case "a":
        _ = len(v2)
    }
    v44 := GenS7{F0: (v3 | v3), F1: v5, F2: 9223372036854775807}
    v45 := GenS3{F0: (-(v5)), F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
