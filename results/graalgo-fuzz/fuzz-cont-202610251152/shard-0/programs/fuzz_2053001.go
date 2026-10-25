package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 int) uint32 {
    var v7 int8 = -60
    var v8 uint16 = 839
    if ((v8 / ((2) | 1)) >= 255) {
        v9 := []int{p0, -9223372036854775808, p0, p0}
        for fi10 := 0; fi10 < 13; fi10++ {
            switch (fi10 ^ 127) {
            case -9223372036854775808:
                v8 = v8
            default:
                _ = len(v9)
            }
            _ = fi10
        }
        _ = v9
    }
    fmt.Println(v8, v7, p0)
    var v11 float64 = (0.0 + 0.0)
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 uint64 = 1
    var v2 int8 = 100
    var v3 string = "a"
    var v4 int8 = v2
    var v5 int8 = (-(v4))
    var v6 uint64 = 18446744073709551615
    var v12 int = -896
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            v5 -= -1
            _ = fi14
        }
        v15 := []int{(fi13 >> ((458) & 63)), 1, v12}
        _ = fi13
        _ = v15
    }
    switch -2147483648 {
    case -37:
        var v16 int16 = int16(v12)
        _ = v16
    case -901:
        var v17 uint16 = 32767
        _ = v17
    case 9223372036854775807:
        v18 := []int{v12, 9223372036854775807, 180}
        _ = v18
    default:
        v19 := GenS1{F0: 127, F1: 42}
        _ = v19
    }
    v5 = (v2 ^ v5)
    v20 := GenS1{F0: 42, F1: -1}
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 float64 = ((float64(((1e10 * 0.0) * float32(3.14159))) - 1.5) - 0.0)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := []int{v12, (v12 + v12), fi23}
        _ = fi23
        _ = v24
    }
    var v25 int32 = -1
    v26 := GenS1{F0: 146, F1: -1}
    v27 := []int{v12}
    v2 = v2
    v27 = append(v27, (v12 + v12))
    v25 -= (v25 + (v25 + v25))
    var v28 bool = (525 < 1000)
    switch (((v3 + ((((((v3 + "\n\t") + v3) + (v3 + v3)) + ((v3 + v3) + ("a" + "\n\t"))) + (((v3 + v3) + ("hello" + ((v3 + v3) + v3))) + v3)) + "foo bar")) + v3) + v3) {
    case "":
        var v29 int64 = 9223372036854775807
        _ = v29
    }
    v27 = append(v27, (v12 / ((v12) | 1)))
    v25 += v25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
