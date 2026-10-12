package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int, p1 uint32) uint8 {
    v5 := []int{(p0 & 32767)}
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float32 = float32(float64(1.5))
        if false {
            if len(v5) > 0 {
                v5[0] = (fi6 * (-(fi6)))
            }
        } else {
            fmt.Println(fi6)
        }
        var v8 int64 = 9223372036854775807
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0)
    if false {
        v9 := GenS1{F0: ("a" + (("" + (("hello" + "\n\t") + "")) + "\n\t")), F1: 365}
        var v10 bool = (true && true)
        var v11 uint32 = (p1 & (^(p1)))
        _ = v10
        _ = v11
        _ = v9
    } else {
        v5 = append(v5, (p0 * (p0 >> ((1) & 63))))
        var v12 uint32 = p1
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1() uint32 {
    v13 := GenS2{F0: 1}
    v14 := []int{}
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 string = ""
    v2 := make(map[int]int)
    var v3 float64 = (0.1 - 3.14159)
    v4 := []uint32{1, 317, 585, 542}
    v3 = (0.1 + 0.1)
    if (v3 <= v3) {
        fmt.Println(v1)
        var v15 uint32 = 1
        _ = v15
    } else {
        fmt.Println(v3)
    }
    fmt.Println(v3)
    fmt.Println(v1)
    fmt.Println(v1)
    var v16 int64 = -1
    if false {
        var v17 uint64 = 18446744073709551615
        if false {
            v3 = (float64(((0.0 * 0.1) - 1.5)) / 1.0)
            if (100 <= 42) {
                v4 = append(v4, 1)
                v4 = append(v4, 135)
                v4 = append(v4, 663)
            } else {
                v17 -= (v17 << ((2) & 63))
                v3 -= 0.0
                for fi18 := 0; fi18 < 14; fi18++ {
                    fmt.Println(v3, v1)
                    switch "a" {
                    case "a":
                        if (true && false) {
                            for fi19 := 0; fi19 < 11; fi19++ {
                                v20 := GenS2{F0: -2147483648}
                                v16 += v16
                                _ = fi19
                                _ = v20
                            }
                        }
                    case "hello":
                        v21 := v2[4294967295]
                        _ = v21
                    }
                    v17 = (v17 % ((0) | 1))
                    _ = fi18
                }
            }
        } else {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 float64 = float64(0.0)
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 int16 = 0
                    var v26 uint64 = (v17 | (v17 ^ 480))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v27 := GenS1{F0: v1, F1: 0}
                _ = fi22
                _ = v23
                _ = v27
            }
            var v28 string = "foo bar"
            _ = v28
        }
        if (42 <= 132) {
            var v29 uint8 = 42
            for fi30 := 0; fi30 < 14; fi30++ {
                v3 -= (v3 - 0.1)
                _ = fi30
            }
            _ = v29
        }
        _ = v17
    }
    v3 += float64(0.1)
    if false {
        fmt.Println(v3, v1)
        var v31 int64 = (v16 | v16)
        v16 = 881
        _ = v31
    } else {
        v16 = v16
        fmt.Println(v1, v3)
        v4 = append(v4, 4294967295)
    }
    var v32 int16 = -562
    v16 += v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
