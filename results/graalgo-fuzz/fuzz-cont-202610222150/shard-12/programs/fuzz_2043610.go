package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() int32 {
    v7 := []int{0, -2147483648, 42, 32768}
    if len(v7) > 0 {
        v7[0] = 32767
    }
    var v8 int16 = 0
    _ = v7
    _ = v8
    return 0
}
func genFn1() int16 {
    var v9 uint = 2
    v9 -= (v9 ^ 860)
    v10 := []int{4294967295, 883}
    _ = v10
    _ = v9
    return 390
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-650, -2147483648}
    var v3 int8 = 127
    v4 := make(map[int]int)
    var v5 uint32 = 2147483647
    var v6 uint = 2
    v3 = 127
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int32 = 278
        _ = fi11
        _ = v12
    }
    v5 += v5
    delete(v1, 1)
    for fi13 := 0; fi13 < 16; fi13++ {
        v5 = v5
        _ = fi13
    }
    var v14 uint16 = 169
    switch -2147483648 {
    case 128:
        if (((0.0 * 1e10) <= float32(-2.25)) && (28 <= 148)) {
            v3 = v3
            var v15 uint32 = (^(v5))
            _ = v15
        } else {
            var v16 float32 = (1.5 / 1.0)
            var v17 int8 = (v3 - v3)
            v14 += 32767
            _ = v16
            _ = v17
        }
    case -702:
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int8 = -1
            var v20 float64 = 3.14159
            _ = fi18
            _ = v19
            _ = v20
        }
    case 9223372036854775807:
        fmt.Println(v3, v5)
    }
    v21 := GenS1{F0: (("\n\t" + "hello") + "a")}
    var v22 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
