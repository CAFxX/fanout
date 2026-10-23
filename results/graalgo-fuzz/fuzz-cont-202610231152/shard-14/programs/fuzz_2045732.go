package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 bool = (false || (!(1 == 4294967295)))
    var v2 float32 = ((((1e10 / (3.14159 * 0.1)) - (1.5 - 1.5)) - (-2.25 - (3.14159 + (1.5 / 1.0)))) / float32((((1.5 * (1e10 / 1.5)) + 1.5) / 1.0)))
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    var v5 uint64 = 4294967295
    v6 := make(map[string]bool)
    fmt.Println(v5, v5, v1)
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 float64 = ((0.1 / float64(v2)) - (0.0 / 1.0))
        _ = fi7
        _ = v8
    }
    if (3.14159 == (-2.25 + 1e10)) {
        v2 = (float32(-2.25) / v2)
        var v9 int8 = 127
        fmt.Println(v1, v5, v9)
        _ = v9
    }
    v10 := []int{42, 127, -2147483648, -2147483648}
    var v11 string = "foo bar"
    fmt.Println(v11, v1, v11)
    v6[v11] = true
    v5 += v5
    v12 := GenS1{F0: 4294967295, F1: (v1 && ((-1 == -2147483648) && v1))}
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            if true {
                v2 -= ((v2 / ((v2 + (((1.5 + 1.5) * ((float32(0.0) - -2.25) * 1.5)) / v2)) * 1e10)) + float32(0.0))
                switch fi14 {
                case 9223372036854775807:
                    fmt.Println(fi13, v5, v1)
                default:
                    _ = len(v10)
                }
                fmt.Println(v5)
            }
            v11 = (v11 + "\n\t")
            if (0 >= -369) {
                v2 -= -2.25
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi14
        }
        fmt.Println(v11)
        _ = fi13
    }
    v5 += (^(v5))
    if v1 {
        if (0.1 < v2) {
            for fi15 := 0; fi15 < 20; fi15++ {
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(fi16, fi15)
                    switch -2147483648 {
                    case 128:
                        v2 += 0.1
                    case 127:
                        v17 := GenS1{F0: 790, F1: ((true || (128 < 2)) && ((((float64(0.1) + 1e10) + (0.1 / 1.0)) * 0.1) > -2.25))}
                        _ = v17
                    case 32768:
                        v11 = v11
                    }
                    _ = fi16
                }
                _ = fi15
            }
        } else {
            v11 = "hello"
            v18 := GenS2{F0: 127}
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(v1, fi19, v11)
            if len(v10) > 0 {
                v10[0] = 0
            }
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
