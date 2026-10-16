package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
func genFn0() string {
    var v4 float64 = 3.14159
    v4 = 0.1
    v4 -= 3.14159
    _ = v4
    return "\n\t"
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 0
    var v3 int32 = -1
    var v5 int32 = (v3 / (((v3 >> ((2) & 63))) | 1))
    v2 += v2
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(v3, v5, v2)
        v5 = (v3 * v3)
        _ = fi6
    }
    _ = len(v1)
    if (!((float64(((1e10 * float32((1e10 - (0.0 + ((float64((1e10 - 0.1)) / (-2.25 / 3.14159)) - (-2.25 + 3.14159)))))) * 1e10)) == 1e10) || false)) {
        for fi7 := 0; fi7 < 20; fi7++ {
            if (false && (!false)) {
                switch "\n\t" {
                case "hello":
                    for fi8 := 0; fi8 < 8; fi8++ {
                        fmt.Println(fi8)
                        v9 := GenS1{F0: (false || (fi7 >= 333)), F1: v3, F2: 1000}
                        v3 = v3
                        _ = fi8
                        _ = v9
                    }
                default:
                    var v10 int8 = -128
                    _ = v10
                }
                v11 := GenS1{F0: ((v5 < v3) || (0.1 <= 1e10)), F1: v3, F2: -32768}
                _ = v11
            } else {
                v5 = -1
                v5 += v5
            }
            v12 := GenS1{F0: true, F1: v3, F2: 0}
            for fi13 := 0; fi13 < 20; fi13++ {
                v5 = v3
                v3 += (v5 * (v3 << ((0) & 63)))
                _ = fi13
            }
            _ = fi7
            _ = v12
        }
    } else {
        fmt.Println(v3)
    }
    if false {
        var v14 int = 932
        _ = v14
    } else {
        var v15 bool = true
        if v15 {
            var v16 int = 0
            v15 = v15
            v2 += 42
            _ = v16
        }
        if v15 {
            v5 += 0
            fmt.Println(v2, v3)
        }
        _ = v15
    }
    v17 := []int{}
    var v18 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
}
