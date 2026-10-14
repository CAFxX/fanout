package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint16, p1 int) string {
    if true {
        var v7 int64 = 386
        var v8 uint8 = 116
        _ = v7
        _ = v8
    } else {
        v9 := []int{334}
        _ = v9
    }
    v10 := GenS1{F0: 1}
    p0 = 0
    v11 := GenS1{F0: 2}
    v12 := []int{376, (p1 + (p1 - (p1 + p1))), p1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return "a"
}

func main() {
    v1 := []string{}
    v2 := []string{("\n\t" + ("hello" + "hello")), "", "foo bar"}
    v3 := make(map[uint32]bool)
    v4 := []string{}
    var v5 uint32 = 2147483647
    var v6 bool = false
    fmt.Println(v5, v5)
    var v13 float32 = (0.0 * ((1.5 + -2.25) / 3.14159))
    var v14 uint64 = 184
    switch -400 {
    case 32767:
        v6 = (-9223372036854775808 > 0)
    default:
        v15 := GenS1{F0: uint8(v14)}
        _ = v15
    }
    v16 := GenS2{F0: 255, F1: (^(v14))}
    var v17 uint64 = 73
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int64 = -9223372036854775808
        var v20 uint8 = 135
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int16 = 0
    v14 += 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 int64 = 0
    v3[42] = false
    var v23 uint64 = v14
    v24 := GenS2{F0: 7, F1: 42}
    v25 := GenS2{F0: 0, F1: v17}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
