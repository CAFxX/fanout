package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 bool }
func genFn0(p0 uint8) int64 {
    switch -552 {
    case 747:
        v6 := []int{-9223372036854775808}
        _ = v6
    case 4294967295:
        v7 := GenS1{F0: 2, F1: 42, F2: 32767}
        _ = v7
    }
    switch -2147483648 {
    case 127:
        v8 := GenS2{F0: 290, F1: true}
        _ = v8
    default:
        v9 := GenS1{F0: 4294967295, F1: -1, F2: 384}
        _ = v9
    }
    v10 := []int{0, -9223372036854775808, 327, 214}
    fmt.Println(p0)
    _ = p0
    _ = v10
    return -1
}
func genFn1() int16 {
    var v11 int64 = -570
    v11 += (v11 % ((v11) | 1))
    _ = v11
    return 42
}

func main() {
    var v1 uint64 = 195
    var v2 uint32 = 0
    var v3 int32 = -1
    var v4 float64 = 3.14159
    var v5 int8 = 127
    if ((v3 - (v3 / ((42) | 1))) <= 2147483647) {
        fmt.Println(v2)
        v12 := []int{527, 128, 32767, 127}
        var v13 int64 = 421
        _ = v12
        _ = v13
    }
    var v14 int8 = (v5 ^ v5)
    v15 := GenS1{F0: v1, F1: 0, F2: 42}
    var v16 uint64 = v1
    var v17 float64 = (v4 - float64(0.0))
    var v18 uint32 = 4294967295
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 uint64 = (v1 / (((v1 * (^(v1)))) | 1))
        _ = fi19
        _ = v20
    }
    var v21 int16 = 0
    var v22 int = 4294967295
    v22 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
