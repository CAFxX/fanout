package main

import (
    "fmt"
)

func genFn0(p0 uint16) int16 {
    p0 += (^(p0))
    if false {
        v5 := []int{}
        _ = v5
    }
    switch 0 {
    case 128:
        var v6 uint16 = (p0 * (p0 & p0))
        _ = v6
    case -662:
        fmt.Println(p0)
    default:
        var v7 uint8 = 128
        _ = v7
    }
    p0 = p0
    _ = p0
    return 362
}

func main() {
    var v1 int16 = 42
    var v2 int = -440
    var v3 float32 = -2.25
    var v4 int64 = 0
    if (false && (false && true)) {
        v8 := []int{300}
        _ = v8
    } else {
        v9 := []int{v2, v2}
        var v10 float64 = 1e10
        v4 = v4
        _ = v10
        _ = v9
    }
    if true {
        v2 = (v2 % (((v2 / (((v2 >> ((255) & 63))) | 1))) | 1))
        var v11 uint8 = 255
        _ = v11
    } else {
        var v12 int64 = -9223372036854775808
        _ = v12
    }
    v1 = (v1 / (((v1 * (-(v1)))) | 1))
    v4 = v4
    var v13 string = ("foo bar" + "")
    v14 := []int{(v2 / ((1) | 1)), v2}
    v13 = v13
    var v15 uint64 = 18446744073709551615
    var v16 uint8 = 0
    var v17 float32 = float32((1.5 - -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
