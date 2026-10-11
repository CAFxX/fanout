package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 int64 }
func genFn0(p0 uint64, p1 uint8, p2 uint64) int8 {
    var v5 string = ((("\n\t" + "a") + "") + "a")
    var v6 uint8 = p1
    var v7 float32 = 3.14159
    var v8 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1() bool {
    v9 := []int{}
    var v10 bool = true
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v10 = ((!(v10 || v10)) && (!v10))
            if len(v9) > 0 {
                v9[0] = 0
            }
            v13 := GenS1{F0: -49, F1: 1}
            var v14 string = ""
            _ = fi12
            _ = v13
            _ = v14
        }
        v10 = (-2147483648 >= 2147483647)
        fmt.Println(fi11, v10)
        _ = fi11
    }
    var v15 uint32 = 625
    v16 := GenS1{F0: 100, F1: 823}
    _ = v10
    _ = v15
    _ = v16
    _ = v9
    return (float64(3.14159) == ((0.0 / 1.5) - 1.5))
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint32 = 42
    var v3 uint32 = (v2 / (((v2 << ((4294967295) & 63))) | 1))
    var v4 int = -1
    fmt.Println(v4)
    v17 := GenS2{F0: 42, F1: 0, F2: 885}
    v2 += (v3 & 181)
    var v18 float64 = (((0.1 * 0.1) - (1.5 + 0.1)) / ((1e10 / 1e10) / (((1.5 - 0.1) * (3.14159 * -2.25)) - 0.0)))
    v4 = 794
    var v19 uint32 = (v2 / ((2147483647) | 1))
    var v20 uint32 = (^(v3))
    if (int(v1) > (v4 << ((18446744073709551615) & 63))) {
        v21 := []int{(v4 ^ v4), v4, (v4 - -9223372036854775808)}
        fmt.Println(v18, v19)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
