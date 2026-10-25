package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 684
    v3 := []uint32{}
    v4 := []int{-944}
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]string)
    switch 9223372036854775807 {
    case 4294967295:
        if len(v3) > 0 {
            v3[0] = 42
        }
    case -1:
        fmt.Println(v2)
    case 9223372036854775807:
        var v7 int = -362
        _ = v7
    default:
        if true {
            delete(v1, 32768)
        }
    }
    v2 = v2
    v2 += (^(v2))
    switch "hello" {
    case "":
        v8 := GenS1{F0: 9223372036854775807, F1: -1, F2: 1}
        _ = v8
    case "\n\t":
        v2 = v2
    default:
        v2 -= (^(v2))
    }
    _ = len(v6)
    var v9 int8 = 127
    if false {
        switch 2147483647 {
        case 128:
            v10 := GenS1{F0: 644, F1: (v9 | (v9 << ((350) & 63))), F2: 155}
            _ = v10
        case 0:
            v11 := GenS1{F0: -1, F1: (v9 + v9), F2: -1}
            _ = v11
        default:
            switch 4294967295 {
            case -306:
                switch "" {
                case "a":
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                case "hello":
                    if (2147483647 >= 4294967295) {
                        v9 = -1
                    } else {
                        v9 = (v9 + v9)
                    }
                default:
                    v12 := GenS1{F0: -9223372036854775808, F1: (v9 ^ (v9 << ((467) & 63))), F2: 1}
                    _ = v12
                }
            }
        }
        v5[2147483647] = 982
    }
    fmt.Println(v2, v9)
    for fi13 := 0; fi13 < 15; fi13++ {
        v9 = (v9 & v9)
        fmt.Println(v2, v9, v9)
        _ = fi13
    }
    if ((("\n\t" + "hello") + "") == ("" + "foo bar")) {
        v2 = v2
        fmt.Println(v9)
        for fi14 := 0; fi14 < 8; fi14++ {
            v2 = 1
            _ = fi14
        }
    } else {
        var v15 uint = 415
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        _ = len(v4)
        v17 := GenS2{F0: 1}
        if true {
            v18 := GenS1{F0: 395, F1: (v9 * v9), F2: -1}
            var v19 int16 = 1
            _ = v18
            _ = v19
        } else {
            var v20 int32 = 513
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v9 += v9
    if false {
        var v21 int8 = 127
        for fi22 := 0; fi22 < 19; fi22++ {
            if ((^(v2)) <= (^(v2))) {
                var v23 uint16 = (v2 >> ((1) & 63))
                _ = v23
            } else {
                v9 += (v9 << ((821) & 63))
                var v24 int = fi22
                for fi25 := 0; fi25 < 3; fi25++ {
                    var v26 uint = 255
                    if false {
                        v26 += 520
                        v27 := GenS1{F0: -9223372036854775808, F1: (^(v21)), F2: 9223372036854775807}
                        v28 := GenS1{F0: 42, F1: -128, F2: -849}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 uint = v26
                        var v30 int = 1
                        if false {
                            v30 = 825
                            v29 = (v26 ^ v29)
                        }
                        _ = v29
                        _ = v30
                    }
                    var v31 uint8 = 239
                    _ = fi25
                    _ = v26
                    _ = v31
                }
                _ = v24
            }
            _ = fi22
        }
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
