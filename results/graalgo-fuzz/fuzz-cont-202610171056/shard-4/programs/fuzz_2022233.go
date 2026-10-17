package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }

func main() {
    var v1 int32 = 1
    var v2 int64 = -580
    v3 := make(map[int]int)
    var v4 int64 = 188
    var v5 float32 = (float32(3.14159) + (1.5 - 0.0))
    var v6 bool = true
    v7 := []int{127, -452, 501, 42}
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    var v8 uint32 = 2
    v4 += v4
    v9 := GenS1{F0: -1, F1: "foo bar"}
    v10 := GenS1{F0: int8(v1), F1: "a"}
    _ = len(v3)
    if (0 >= (v1 | (v1 / ((v1) | 1)))) {
        var v11 float32 = ((float32(1.5) + 1e10) - v5)
        if len(v7) > 0 {
            v7[0] = 127
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            if v6 {
                for fi13 := 0; fi13 < 16; fi13++ {
                    if true {
                        v14 := v3[int(v1)]
                        v6 = (!(!((v8 / ((42) | 1)) <= 4294967295)))
                        _ = v14
                    } else {
                        v4 += (v4 & (v4 % ((-1) | 1)))
                        fmt.Println(v8, v11, v1)
                        var v15 int8 = 100
                        v7 = append(v7, (fi12 >> ((322) & 63)))
                        _ = v15
                    }
                    v16 := GenS2{F0: (v2 >> ((0) & 63)), F1: 498, F2: v8}
                    _ = fi13
                    _ = v16
                }
            }
            _ = fi12
        }
        _ = v11
    }
    v1 += 147
    if (v2 > (^(v2))) {
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(v2, v1, v2)
            _ = fi17
        }
        var v18 int8 = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v4 -= v2
        var v20 float64 = 1e10
        fmt.Println(v6, v4, v6)
        _ = fi19
        _ = v20
    }
    v1 += -523
    v1 = (v1 * 2147483647)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
