package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 int64) int8 {
    if (128 < 128) {
        v5 := []int{}
        _ = v5
    } else {
        var v6 int16 = -178
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    return 3
}

func main() {
    var v1 bool = (0.0 >= ((1e10 * (-2.25 - 1e10)) * -2.25))
    var v2 float64 = (3.14159 - (((1.5 - 3.14159) - 1.5) + 1.5))
    var v3 uint = 869
    v4 := []int{127}
    for fi7 := 0; fi7 < 1; fi7++ {
        v2 = float64(0.1)
        switch 128 {
        case -576:
            var v8 int16 = 0
            _ = v8
        }
        _ = fi7
    }
    var v9 int16 = 42
    if ((100 == 1) || (42 != 864)) {
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
        fmt.Println(v1)
    }
    var v10 float32 = 1.5
    if ((0 > 0) || (0 >= 0)) {
        fmt.Println(v2, v10)
        switch 203 {
        case 128:
            for fi11 := 0; fi11 < 10; fi11++ {
                for fi12 := 0; fi12 < 7; fi12++ {
                    fmt.Println(v9)
                    v13 := GenS1{F0: 2, F1: (v3 | v3)}
                    fmt.Println(v10, fi12, fi12)
                    _ = fi12
                    _ = v13
                }
                v14 := GenS1{F0: 194, F1: (^(v3))}
                _ = fi11
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 16; fi15++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi15
            }
        }
    }
    if (1 != int32(v9)) {
        fmt.Println(v3, v10, v3)
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int16 = -394
            _ = fi16
            _ = v17
        }
    }
    v2 -= (((float64(v10) * 1e10) + (v2 + ((0.0 - v2) + ((1.5 + (0.1 * ((float64(v10) * v2) + ((v2 / 0.1) + (v2 / ((1e10 * v2) + 1e10)))))) / -2.25)))) / 1e10)
    var v18 int64 = 941
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
