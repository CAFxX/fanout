package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
func genFn0(p0 uint16) int64 {
    v6 := GenS1{F0: 343, F1: 18446744073709551615, F2: 9223372036854775807}
    var v7 int64 = 42
    var v8 float32 = -2.25
    var v9 bool = (!false)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v7 * (v7 ^ (v7 * v7)))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = 0.0
    v3 := make(map[string]int)
    var v4 float64 = v2
    var v5 string = ("a" + "a")
    v2 -= v2
    var v10 int32 = -2147483648
    v5 = v5
    var v11 uint = 4294967295
    var v12 uint32 = uint32(v11)
    var v13 float32 = ((((1.5 * 0.1) + 3.14159) - float32(float64(1e10))) - 0.0)
    var v14 uint64 = 4294967295
    var v15 int8 = 63
    var v16 int16 = -1
    v2 -= float64((v13 / 1e10))
    v14 = 0
    fmt.Println(v16)
    var v17 float32 = v13
    v18 := GenS1{F0: 42, F1: (v14 * (v14 >> ((v11) & 63))), F2: 32768}
    v2 -= 3.14159
    fmt.Println(v12, v15)
    var v19 uint8 = 188
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
