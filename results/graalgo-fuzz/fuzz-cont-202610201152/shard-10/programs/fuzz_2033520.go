package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }
func genFn0() uint16 {
    var v4 int8 = 1
    if (0.0 == float64(0.1)) {
        fmt.Println(v4)
    }
    if true {
        v5 := GenS1{F0: (v4 % (((v4 >> ((0) & 63))) | 1))}
        _ = v5
    }
    _ = v4
    return 319
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]bool)
    var v3 int8 = 1
    for fi6 := 0; fi6 < 5; fi6++ {
        v1 -= v1
        _ = fi6
    }
    fmt.Println(v3, v3)
    v7 := GenS2{F0: -44}
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: 100}
        v10 := GenS2{F0: 497}
        if true {
            if (65535 == 331) {
                _ = len(v2)
                for fi11 := 0; fi11 < 8; fi11++ {
                    if true {
                        var v12 uint8 = 1
                        v1 += v1
                        var v13 float32 = v1
                        _ = v12
                        _ = v13
                    } else {
                        v1 += ((v1 - v1) * v1)
                        v14 := GenS2{F0: 2147483647}
                        var v15 uint32 = 1
                        _ = v14
                        _ = v15
                    }
                    _ = fi11
                }
            }
            switch (fi8 ^ (fi8 & -302)) {
            case 299:
                var v16 uint16 = 65535
                _ = v16
            default:
                var v17 string = "\n\t"
                _ = v17
            }
            switch "hello" {
            case "":
                if true {
                    var v18 int8 = v3
                    fmt.Println(fi8)
                    if (!true) {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            fmt.Println(fi19, v18)
                            if false {
                                var v20 int = (fi19 / ((fi8) | 1))
                                _ = v20
                            } else {
                                v3 = (v18 * v3)
                                var v21 uint16 = 42
                                _ = v21
                            }
                            _ = fi19
                        }
                        fmt.Println(v18)
                        for fi22 := 0; fi22 < 20; fi22++ {
                            v23 := GenS2{F0: -1}
                            if (true || (true || ((!true) && true))) {
                                v3 = v3
                            } else {
                                v3 += (-(v3))
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = v18
                }
            case "foo bar":
                v3 += (v3 ^ (v3 >> ((2) & 63)))
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = -100
    v24 := GenS3{F0: -1, F1: 18446744073709551615, F2: (true || true)}
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := []int{fi25, fi25, 9223372036854775807, 127}
        v1 -= (v1 + v1)
        _ = fi25
        _ = v26
    }
    v1 -= -2.25
    var v27 uint8 = 102
    v28 := []int{0, -453, -276, 852}
    var v29 uint8 = v27
    fmt.Println(v1, v29, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
}
