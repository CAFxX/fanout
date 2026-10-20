package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 uint64) int16 {
    fmt.Println(p2, p1, p0)
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 float32 = -2.25
        if (fi6 > (fi6 - (fi6 / ((fi6) | 1)))) {
            v8 := []int{}
            p0 = (p0 / (((p0 - (p0 - uint8(p2)))) | 1))
            v9 := GenS1{F0: 65535, F1: (p0 % ((42) | 1))}
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 823, F1: (^(p0))}
            var v11 int64 = -9223372036854775808
            _ = v10
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 197
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{187, 128}
    v3 := make(map[int]uint32)
    var v4 float32 = 0.0
    v5 := []uint32{1, 718, 589}
    v4 -= (v4 / v4)
    v4 += 1e10
    v4 = (0.1 / (((float32(3.14159) - v4) / v4) - v4))
    var v12 float32 = v4
    v4 += v4
    v4 -= 3.14159
    for fi13 := 0; fi13 < 16; fi13++ {
        if false {
            _ = len(v1)
            if len(v2) > 0 {
                v2[0] = fi13
            }
        }
        if true {
            fmt.Println(v4)
            if true {
                var v14 string = ("a" + "hello")
                switch fi13 {
                case -2147483648:
                    v14 = "\n\t"
                case 1:
                    v4 -= 0.1
                }
                _ = len(v1)
                _ = v14
            } else {
                for fi15 := 0; fi15 < 16; fi15++ {
                    if (9223372036854775807 <= 0) {
                        fmt.Println(fi13, v4, v4)
                        if (!false) {
                            switch ("a" + "foo bar") {
                            case "foo bar":
                                fmt.Println(fi15, fi13)
                            }
                            v4 += (v12 - (v12 * ((v12 - float32(1.5)) - v4)))
                            v16 := GenS2{F0: 1}
                            _ = v16
                        } else {
                            var v17 float64 = 0.1
                            if true {
                                fmt.Println(v12, v17)
                            } else {
                                fmt.Println(v4, fi13, v17)
                                if len(v5) > 0 {
                                    v5[0] = 2
                                }
                            }
                            _ = v17
                        }
                    } else {
                        switch ("" + "hello") {
                        case "a":
                            var v18 float64 = 1.5
                            _ = v18
                        }
                    }
                    _ = fi15
                }
                v19 := GenS1{F0: 1, F1: 255}
                _ = v19
            }
        }
        fmt.Println(v4)
        _ = fi13
    }
    v4 -= v12
    v20 := GenS1{F0: 42, F1: 1}
    v21 := GenS2{F0: 42}
    v22 := GenS1{F0: 42, F1: 255}
    v4 = (v12 - (v4 * v12))
    var v23 int16 = 1
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 bool = (1.5 > 3.14159)
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
