package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float32 = 1e10
    var v2 int8 = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    var v5 uint64 = 548
    fmt.Println(v2, v2)
    v6 := GenS1{F0: -1, F1: (!(255 != 42))}
    if false {
        var v7 int8 = v2
        _ = v7
    } else {
        fmt.Println(v2, v1, v2)
    }
    if true {
        if (!(!true)) {
            fmt.Println(v1)
        }
        if true {
            for fi8 := 0; fi8 < 5; fi8++ {
                for fi9 := 0; fi9 < 7; fi9++ {
                    v1 = (v1 / v1)
                    if false {
                        var v10 float32 = -2.25
                        switch "hello" {
                        case "hello":
                            fmt.Println(v5)
                        case "\n\t":
                            if (772 != 0) {
                                v1 = v10
                                fmt.Println(v5, fi8, v10)
                                if true {
                                    v11 := []int{fi9, (-(fi8))}
                                    _ = v11
                                }
                                var v12 int16 = 455
                                _ = v12
                            }
                            if false {
                                if (128 >= 128) {
                                    var v13 float64 = (0.0 / (1.5 + (1e10 - ((0.0 - (0.0 * 0.1)) * (float64(v1) - float64(v1))))))
                                    _ = v13
                                } else {
                                    v5 += (^(v5))
                                    if (!((true && false) || ((fi8 % (((fi9 / ((fi8) | 1))) | 1)) >= 221))) {
                                        v14 := GenS1{F0: 1, F1: true}
                                        _ = v14
                                    }
                                }
                            } else {
                                v5 = v5
                            }
                        case "a":
                            v15 := GenS1{F0: 946, F1: (65535 < 65535)}
                            _ = v15
                        default:
                            var v16 uint32 = 2147483648
                            _ = v16
                        }
                        _ = v10
                    }
                    var v17 float32 = (v1 - v1)
                    _ = fi9
                    _ = v17
                }
                _ = fi8
            }
            v1 = ((-2.25 - (1e10 + v1)) + v1)
        }
    } else {
        v1 += v1
        v18 := GenS2{F0: v5}
        var v19 int32 = 0
        _ = v18
        _ = v19
    }
    if false {
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint64 = (v5 ^ (v5 ^ v5))
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 uint8 = 185
                if false {
                    v24 := GenS1{F0: -2147483648, F1: false}
                    fmt.Println(v1, v1, v23)
                    v1 = ((v1 / -2.25) / v1)
                    _ = v24
                } else {
                    v5 = (v5 % ((v5) | 1))
                }
                fmt.Println(v2)
                _ = fi22
                _ = v23
            }
            var v25 uint64 = 0
            _ = fi20
            _ = v21
            _ = v25
        }
        v5 = v5
    }
    v2 -= (v2 ^ 104)
    if true {
        fmt.Println(v5, v2)
        fmt.Println(v1, v5)
    }
    var v26 int64 = 0
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := []int{(fi27 % ((fi27) | 1)), (fi27 << ((65535) & 63)), 0}
        var v29 uint32 = 181
        v3[2] = (v29 << ((512) & 63))
        _ = fi27
        _ = v28
        _ = v29
    }
    v1 += v1
    v26 += v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
