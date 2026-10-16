package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0() int16 {
    var v6 uint = 374
    v6 += (v6 & v6)
    var v7 float64 = (((float64(-2.25) + 0.0) - ((((0.0 / 3.14159) - 3.14159) * 1.5) + 0.1)) * (0.0 * -2.25))
    _ = v6
    _ = v7
    return -585
}

func main() {
    var v1 int = 32767
    var v2 float64 = 3.14159
    v3 := []int{(v1 * 664), -297}
    var v4 uint8 = 1
    var v5 float64 = (0.0 + (1.5 * ((v2 / v2) + (((1.5 / 3.14159) - (float64(((((3.14159 - 0.0) / 0.1) / (3.14159 * (-2.25 + float32((((float64(((3.14159 + ((0.1 / 1.0) * (-2.25 - (0.1 - 3.14159)))) * 1e10)) * float64(float32(3.14159))) - ((0.0 * (float64(float32(v2)) + v2)) * float64(1.5))) + v2))))) * 1.5)) * float64(0.0))) + v2))))
    if ((-2.25 - (1.5 / 1.0)) <= 0.0) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 float64 = ((v2 - float64((1e10 - (((1e10 / float32(0.1)) / (float32(v2) - 0.1)) + 1.5)))) + 0.0)
            if false {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 uint = 42
                    _ = fi10
                    _ = v11
                }
                switch (fi8 * fi8) {
                case 2147483647:
                    v12 := GenS1{F0: 868, F1: 0}
                    _ = v12
                case 604:
                    v3 = append(v3, 0)
                case 4294967295:
                    v3 = append(v3, fi8)
                default:
                    var v13 uint16 = 32767
                    _ = v13
                }
            }
            _ = fi8
            _ = v9
        }
    } else {
        v14 := GenS1{F0: genFn0(), F1: -9223372036854775808}
        if ((!(!(!false))) && false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v14
    }
    switch v1 {
    case -564:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch 42 {
        case 9223372036854775807:
            var v15 float32 = 0.1
            _ = v15
        }
    }
    v16 := GenS1{F0: 1, F1: 42}
    v17 := GenS1{F0: 1, F1: 9223372036854775807}
    v2 = v5
    fmt.Println(v2)
    v2 = v5
    v18 := GenS2{F0: 1, F1: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
