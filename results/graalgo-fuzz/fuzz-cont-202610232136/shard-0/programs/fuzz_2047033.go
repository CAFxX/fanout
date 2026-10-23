package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0(p0 uint16, p1 int64) int8 {
    if true {
        if ((0.0 / 0.1) > 0.0) {
            p0 = p0
        }
        v4 := []int{}
        _ = v4
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []string{("\n\t" + "\n\t"), "\n\t"}
    var v2 int16 = 262
    var v3 float32 = float32((float64(-2.25) / 1e10))
    v2 = v2
    var v5 float64 = (3.14159 / ((-2.25 + (-2.25 / (((((1.5 - 3.14159) - float64((v3 + 3.14159))) - (3.14159 / -2.25)) + 0.1) * float64(v3)))) + 0.0))
    v2 += (v2 / ((v2) | 1))
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if true {
        var v6 bool = true
        _ = v6
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch (fi7 % (((fi7 % (((fi7 - (fi7 | 2147483647))) | 1))) | 1)) {
            case 0:
                v8 := GenS1{F0: 0, F1: 647, F2: 68}
                _ = v8
            case 1:
                fmt.Println(v2, fi7)
            case -237:
                var v9 uint = 42
                _ = v9
            default:
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
            _ = fi7
        }
    }
    var v10 uint = 654
    var v11 uint64 = 690
    var v12 float64 = (v5 / v5)
    v13 := GenS2{F0: -940, F1: 100}
    var v14 int64 = 864
    fmt.Println(v10, v2)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v5
}
