package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int16; F1 bool }

func main() {
    var v1 int64 = 42
    v2 := []int{int(v1), 0, -300, 1}
    var v3 int32 = 0
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]uint32)
    if false {
        v3 = (v3 + (v3 << ((18446744073709551615) & 63)))
        var v7 uint32 = 706
        var v8 string = "\n\t"
        _ = v7
        _ = v8
    } else {
        var v9 int8 = 127
        switch "a" {
        case "\n\t":
            v4[-2147483648] = (18446744073709551615 == 973)
        case "a":
            var v10 int8 = v9
            _ = v10
        case "hello":
            var v11 uint16 = 882
            _ = v11
        }
        if (0 >= 2147483647) {
            switch -661 {
            case 983:
                if len(v2) > 0 {
                    v2[0] = -1
                }
            default:
                v5[890] = 1
            }
            var v12 int = 42
            v13 := GenS1{F0: int8(v12), F1: v3}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 9; fi14++ {
                if false {
                    switch ("hello" + "hello") {
                    case "hello":
                        fmt.Println(v9, fi14)
                    default:
                        v4[int(v1)] = (false || false)
                    }
                    if true {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            if true {
                                var v16 string = "hello"
                                v17 := GenS1{F0: v9, F1: v3}
                                _ = v16
                                _ = v17
                            }
                            _ = fi15
                        }
                    } else {
                        switch -644 {
                        case 9223372036854775807:
                            switch "" {
                            case "hello":
                                for fi18 := 0; fi18 < 15; fi18++ {
                                    var v19 int64 = (v1 + (^(v1)))
                                    switch fi18 {
                                    case 127:
                                        v20 := GenS2{F0: 356, F1: 588}
                                        _ = v20
                                    case -2147483648:
                                        fmt.Println(fi18, v3)
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = fi14
                                        }
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                            case "a":
                                v9 = (v9 + 42)
                            }
                        case -9223372036854775808:
                            v9 = int8(fi14)
                        case 42:
                            var v21 float64 = ((-2.25 - 3.14159) + float64(1.5))
                            _ = v21
                        }
                        v1 = -9223372036854775808
                    }
                    _ = len(v6)
                }
                _ = fi14
            }
            for fi22 := 0; fi22 < 19; fi22++ {
                v23 := GenS1{F0: (v9 ^ (v9 & (v9 << ((2) & 63)))), F1: v3}
                v24 := GenS3{F0: (fi22 * fi22), F1: 255}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        _ = v9
    }
    fmt.Println(v3, v1)
    var v25 uint16 = 144
    fmt.Println(v3, v1, v25)
    v26 := GenS1{F0: 1, F1: (v3 >> ((65535) & 63))}
    var v27 uint64 = 18446744073709551615
    var v28 int = 353
    v1 = (v1 * int64(v27))
    v29 := GenS4{F0: 1, F1: true}
    fmt.Println(v27, v3, v3)
    var v30 uint16 = (v25 / ((v25) | 1))
    v31 := v6[("hello" + ("hello" + ""))]
    v5[(v31 * (v31 << ((2) & 63)))] = v31
    v5[v31] = v31
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
