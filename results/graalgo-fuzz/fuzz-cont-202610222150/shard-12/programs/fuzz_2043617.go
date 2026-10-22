package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0() string {
    var v5 bool = (((1 < -2147483648) || (!false)) && true)
    var v6 bool = ((!v5) && v5)
    fmt.Println(v5, v5)
    v6 = (32767 > 192)
    var v7 int32 = -2147483648
    _ = v5
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = v1
    var v3 uint8 = 0
    var v4 uint8 = 105
    switch "" {
    case "":
        fmt.Println(v3, v2)
    case "hello":
        var v8 bool = ("\n\t" <= "a")
        _ = v8
    case "\n\t":
        if ((0.1 * (1.5 + 0.1)) > ((float64(1.5) / 1.5) + float64(float32(float64(3.14159))))) {
            v2 -= v1
        }
    default:
        var v9 uint8 = v4
        _ = v9
    }
    v2 += v2
    fmt.Println(v2)
    var v10 int16 = 32767
    v11 := []int{397, 1, 9223372036854775807}
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint = (v2 - 42)
        var v14 uint64 = 4294967295
        v2 = (v1 % ((v13) | 1))
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: (127 == 127), F1: 245}
    v2 = (v2 * (v1 | (v2 / ((255) | 1))))
    v3 += (^(v3))
    var v16 int8 = 0
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int16 = (v10 ^ v10)
        fmt.Println(v18, v10)
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v3)
    v10 += (v10 ^ (-(v10)))
    v10 += v10
    _ = len(v11)
    for fi19 := 0; fi19 < 1; fi19++ {
        v4 = v4
        var v20 uint64 = 159
        v20 = (v20 / ((uint64(v1)) | 1))
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
