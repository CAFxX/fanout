package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    var v1 uint16 = 567
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 bool = false
    v5 := []int{-1, 2147483647}
    v1 = (v1 ^ v1)
    v6 := GenS1{F0: (((3.14159 - 0.1) + 0.0) <= (-2.25 - (0.0 - 1.5)))}
    v4 = v4
    v1 = v1
    if (-1 >= -1) {
        for fi7 := 0; fi7 < 12; fi7++ {
            v4 = (4294967295 > 167)
            var v8 int32 = 1
            v9 := GenS1{F0: (3.14159 <= 0.1)}
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(v1)
        var v10 uint64 = 434
        _ = v10
    } else {
        if v4 {
            for fi11 := 0; fi11 < 16; fi11++ {
                if true {
                    v1 = (v1 % (((v1 - (v1 & v1))) | 1))
                    if (!v4) {
                        v4 = (v1 < 802)
                        fmt.Println(v4, fi11, v1)
                    } else {
                        v1 = 2
                        fmt.Println(v4)
                        v1 -= (v1 ^ (v1 + v1))
                    }
                    fmt.Println(v4)
                } else {
                    fmt.Println(v1, v4)
                }
                _ = fi11
            }
        } else {
            if (v4 || (!(false || (v4 || false)))) {
                v12 := GenS1{F0: false}
                for fi13 := 0; fi13 < 3; fi13++ {
                    v1 -= (v1 / (((v1 >> ((65535) & 63))) | 1))
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v4 = (128 == 127)
                        if (!v4) {
                            v5 = append(v5, fi14)
                            v15 := GenS1{F0: (18446744073709551615 < 0)}
                            _ = v15
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 290
                            }
                        }
                        var v16 int = (fi14 / (((^(fi14))) | 1))
                        _ = fi14
                        _ = v16
                    }
                    _ = fi13
                }
                _ = v12
            }
            var v17 uint = 255
            if v4 {
                var v18 uint64 = 101
                _ = v18
            } else {
                v19 := GenS1{F0: false}
                var v20 int8 = -1
                _ = v19
                _ = v20
            }
            _ = v17
        }
        v1 = 1
        if ((v1 - 65535) <= (^(v1))) {
            _ = len(v3)
        }
    }
    v1 = 432
    var v21 float64 = -2.25
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := GenS2{F0: "foo bar", F1: 1}
        _ = fi22
        _ = v23
    }
    v21 -= v21
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
