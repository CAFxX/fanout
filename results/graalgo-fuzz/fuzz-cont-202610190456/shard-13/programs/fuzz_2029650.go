package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8) uint64 {
    if (!false) {
        var v7 int8 = -1
        v8 := []int{42, -2147483648}
        fmt.Println(p0)
        _ = v7
        _ = v8
    }
    var v9 float64 = 3.14159
    var v10 int32 = 2147483647
    _ = p0
    _ = v10
    _ = v9
    return 42
}
func genFn1() uint8 {
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int = (fi11 & fi11)
        fmt.Println(fi11, v12)
        v13 := GenS1{F0: 65535, F1: 1, F2: -916}
        _ = fi11
        _ = v12
        _ = v13
    }
    switch 565 {
    case 9223372036854775807:
        var v14 int8 = -32
        _ = v14
    default:
        var v15 uint = 1
        _ = v15
    }
    v16 := GenS1{F0: 2147483647, F1: 128, F2: 32767}
    if (((3.14159 * (((-2.25 - 1e10) * float64(3.14159)) + 1e10)) * (1e10 - (0.1 + (-2.25 - 0.1)))) != 0.0) {
        v17 := GenS1{F0: 75, F1: 255, F2: -1}
        _ = v17
    }
    var v18 uint8 = 185
    _ = v16
    _ = v18
    return (v18 >> ((255) & 63))
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{}
    var v3 int32 = v1
    var v4 int64 = -9223372036854775808
    v5 := []int{1, 528, 127}
    var v6 int16 = 703
    v3 -= (v1 << ((757) & 63))
    v19 := GenS1{F0: 931, F1: 2, F2: (v6 ^ (v6 << ((4294967295) & 63)))}
    var v20 int32 = (v3 << ((18446744073709551615) & 63))
    var v21 int32 = 2147483647
    var v22 uint16 = 329
    if false {
        var v23 float64 = 3.14159
        v6 += (v6 & v6)
        _ = v23
    }
    v21 = -2147483648
    var v24 string = "a"
    v6 -= (v6 / ((0) | 1))
    switch "foo bar" {
    case "hello":
        var v25 int8 = int8(v6)
        _ = v25
    }
    v26 := GenS2{F0: -919}
    var v27 uint8 = 42
    switch (v24 + "hello") {
    case "":
        v3 = (v3 % ((v20) | 1))
    case "\n\t":
        v22 = (v22 + 255)
    case "foo bar":
        v2 = append(v2, 2147483647)
    }
    v28 := GenS2{F0: (v6 / ((v6) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
