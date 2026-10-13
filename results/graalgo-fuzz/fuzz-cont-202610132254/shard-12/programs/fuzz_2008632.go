package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }

func main() {
    var v1 int = -964
    var v2 uint8 = 2
    v3 := make(map[int]string)
    var v4 bool = (false || (!((1.5 > ((0.0 / -2.25) + ((((3.14159 * float64(-2.25)) + (0.1 / 1e10)) * float64(((-2.25 * float32((((float64(float32(-2.25)) - (0.0 / 1e10)) + 0.1) * 0.1))) / (float32(0.1) - (float32((1.5 + 0.0)) + 1.5))))) - float64(-2.25)))) && true)))
    v2 += (^(v2))
    fmt.Println(v2, v2)
    fmt.Println(v4)
    if (!((65535 < 42) && (!true))) {
        var v5 int64 = 42
        var v6 int64 = (v5 << ((42) & 63))
        _ = v5
        _ = v6
    }
    v3[(v1 & (v1 % (((^(v1))) | 1)))] = ("a" + "\n\t")
    if (4294967295 > 144) {
        v7 := []int{(v1 << ((428) & 63)), (v1 << ((470) & 63)), (v1 + -1), v1}
        v8 := GenS1{F0: 96, F1: v1, F2: (-2147483648 != -2147483648)}
        _ = v7
        _ = v8
    } else {
        fmt.Println(v1, v4)
    }
    fmt.Println(v1, v1, v2)
    switch v1 {
    case -1:
        v4 = v4
    case 318:
        v9 := GenS1{F0: 65535, F1: v1, F2: (911 >= 4294967295)}
        _ = v9
    case 768:
        if (18446744073709551615 != 18446744073709551615) {
            if false {
                v10 := []int{(v1 + (v1 << ((65535) & 63))), (v1 * (-(v1)))}
                _ = v10
            } else {
                fmt.Println(v2)
                v1 = (v1 ^ v1)
            }
            v11 := GenS1{F0: 0, F1: v1, F2: v4}
            _ = v11
        }
    default:
        var v12 string = "a"
        _ = v12
    }
    var v13 uint16 = 1
    v14 := GenS1{F0: v13, F1: (v1 / ((-9223372036854775808) | 1)), F2: (!((!(!true)) || (1 <= 2)))}
    v1 = 127
    v13 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
