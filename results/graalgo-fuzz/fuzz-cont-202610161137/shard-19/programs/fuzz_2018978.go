package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int64; F1 string; F2 uint32 }
func genFn0() int16 {
    var v7 uint8 = 128
    v8 := []int{-855, 9223372036854775807, 1, -9223372036854775808}
    switch 128 {
    case 9223372036854775807:
        switch 32767 {
        case -9223372036854775808:
            v7 += (v7 << ((0) & 63))
        case 1:
            if (625 != -147) {
                v8 = append(v8, 399)
            } else {
                for fi9 := 0; fi9 < 14; fi9++ {
                    var v10 uint16 = 0
                    v11 := GenS1{F0: 1, F1: 255, F2: 0}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                if (true && (2 >= 807)) {
                    if len(v8) > 0 {
                        v8[0] = -651
                    }
                } else {
                    switch 1 {
                    case 2147483647:
                        if (!false) {
                            for fi12 := 0; fi12 < 9; fi12++ {
                                v7 = (v7 + (^(v7)))
                                _ = fi12
                            }
                        }
                    case 1:
                        v7 -= v7
                    default:
                        if (-814 >= 2147483647) {
                            for fi13 := 0; fi13 < 10; fi13++ {
                                for fi14 := 0; fi14 < 1; fi14++ {
                                    if false {
                                        fmt.Println(fi14, fi14, v7)
                                        var v15 int64 = 960
                                        if false {
                                            var v16 uint16 = 1
                                            var v17 int32 = 1
                                            _ = v16
                                            _ = v17
                                        }
                                        _ = v15
                                    } else {
                                        var v18 int8 = 0
                                        _ = v18
                                    }
                                    v7 = (v7 & v7)
                                    _ = fi14
                                }
                                _ = fi13
                            }
                        }
                    }
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 int8 = -128
                        _ = fi19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 1; fi21++ {
                        if (uint(v7) != 595) {
                            v8 = append(v8, (fi21 * fi21))
                        } else {
                            var v22 uint8 = v7
                            for fi23 := 0; fi23 < 1; fi23++ {
                                v7 -= v7
                                var v24 string = (((("" + "foo bar") + "") + "\n\t") + "a")
                                _ = fi23
                                _ = v24
                            }
                            _ = v22
                        }
                        v7 += 1
                        var v25 int16 = -1
                        _ = fi21
                        _ = v25
                    }
                }
                v7 += 128
            }
        case 52:
            var v26 uint64 = 297
            _ = v26
        }
    case 1:
        for fi27 := 0; fi27 < 2; fi27++ {
            v8 = append(v8, fi27)
            _ = fi27
        }
    case 515:
        for fi28 := 0; fi28 < 9; fi28++ {
            v7 += (v7 >> ((2) & 63))
            fmt.Println(v7, fi28)
            fmt.Println(v7, v7)
            _ = fi28
        }
    default:
        fmt.Println(v7)
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v7
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{837, -322, 435, 128}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 1
    v6 := make(map[int]int)
    v5 += 2
    fmt.Println(v5)
    if false {
        v6[9223372036854775807] = 32767
    } else {
        var v29 int32 = 905
        v5 = (v5 * (v5 >> ((933) & 63)))
        _ = v29
    }
    v5 -= (^(v5))
    fmt.Println(v5)
    v5 -= (v5 * v5)
    v5 -= (v5 + v5)
    for fi30 := 0; fi30 < 3; fi30++ {
        v5 = (v5 << ((0) & 63))
        _ = fi30
    }
    delete(v3, 104)
    if true {
        v31 := GenS2{F0: 2, F1: 42}
        _ = v31
    }
    switch 42 {
    case -9223372036854775808:
        var v32 uint8 = v5
        _ = v32
    case 1:
        fmt.Println(v5)
    case 4294967295:
        var v33 uint = 912
        _ = v33
    default:
        v34 := GenS3{F0: -1, F1: (((("" + ("hello" + ("" + ""))) + "hello") + ("\n\t" + "a")) + "a"), F2: 65535}
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
