package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
func genFn0() int {
    var v5 uint16 = 962
    if (-32768 >= 42) {
        var v6 uint32 = 1
        v6 = (v6 >> ((1) & 63))
        _ = v6
    }
    _ = v5
    return int(v5)
}

func main() {
    var v1 int64 = 0
    v2 := []int{}
    var v3 uint = 42
    v4 := make(map[uint32]bool)
    var v7 int32 = -1
    switch -1 {
    case 42:
        for fi8 := 0; fi8 < 20; fi8++ {
            switch (fi8 | fi8) {
            case -2147483648:
                v3 = v3
            }
            var v9 uint64 = 1
            _ = fi8
            _ = v9
        }
    default:
        v10 := GenS1{F0: int32(v3), F1: 255}
        _ = v10
    }
    v11 := GenS1{F0: v7, F1: 65535}
    v2 = append(v2, 128)
    v1 = v1
    v12 := GenS2{F0: (((("" + "a") != "hello") || (false && true)) || true)}
    if (!(false || true)) {
        v1 += 9223372036854775807
    } else {
        if ((v3 - v3) <= 350) {
            v7 -= 42
            v7 -= 2147483647
        } else {
            v13 := GenS3{F0: 1}
            v3 += 255
            fmt.Println(v3)
            _ = v13
        }
    }
    v7 -= 282
    fmt.Println(v7, v1, v7)
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint16 = 65535
        v1 = (v1 - (-(v1)))
        if (!false) {
            var v16 bool = false
            var v17 float32 = -2.25
            _ = v16
            _ = v17
        } else {
            v15 = 42
        }
        _ = fi14
        _ = v15
    }
    if ((42 > -128) && ("a" <= ("hello" + (("" + "") + ("" + "a"))))) {
        var v18 uint32 = 2147483648
        v19 := GenS4{F0: v18}
        _ = v18
        _ = v19
    }
    _ = len(v2)
    v20 := GenS5{F0: true}
    if true {
        var v21 int64 = v1
        _ = v21
    }
    if (0 > 2) {
        if (-9223372036854775808 != -418) {
            var v22 bool = true
            switch "" {
            case "foo bar":
                v7 += -626
            case "\n\t":
                if v22 {
                    if false {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            fmt.Println(v1, v1, v7)
                            v22 = (1000 >= 1)
                            _ = fi23
                        }
                        if (0 == -54) {
                            var v24 uint8 = 128
                            _ = v24
                        }
                    }
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 float32 = (float32(1e10) / float32(1.5))
                        _ = fi25
                        _ = v26
                    }
                    v3 = 950
                } else {
                    v3 += v3
                    for fi27 := 0; fi27 < 2; fi27++ {
                        v3 += v3
                        _ = len(v2)
                        _ = fi27
                    }
                }
            case "hello":
                if (889 > 255) {
                    if true {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            switch fi28 {
                            case 2147483647:
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    v2 = append(v2, (fi28 & fi28))
                                    _ = fi29
                                }
                            default:
                                var v30 uint = (v3 << (((v3 >> (((v3 | (v3 & (v3 * (^(v3)))))) & 63))) & 63))
                                _ = v30
                            }
                            var v31 uint16 = 861
                            _ = fi28
                            _ = v31
                        }
                    } else {
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v22 = (!(4294967295 <= 1))
                            v33 := GenS2{F0: (!false)}
                            for fi34 := 0; fi34 < 8; fi34++ {
                                var v35 uint16 = 505
                                var v36 float64 = -2.25
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                    fmt.Println(v22, v22, v22)
                }
            }
            _ = v22
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            var v38 uint = (v3 * (v3 ^ (v3 << (((^(v3))) & 63))))
            fmt.Println(v7, fi37)
            _ = fi37
            _ = v38
        }
        v3 = (v3 / ((0) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
}
