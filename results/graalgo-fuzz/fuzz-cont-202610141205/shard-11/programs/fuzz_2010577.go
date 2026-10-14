package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() string {
    if (!(0 <= 2)) {
        switch 42 {
        case 485:
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 float64 = -2.25
                var v7 int = 952
                _ = fi5
                _ = v6
                _ = v7
            }
        }
        var v8 uint = 4294967295
        v9 := []int{-9223372036854775808, 42}
        _ = v8
        _ = v9
    } else {
        var v10 float64 = (((float64(1e10) + -2.25) + 3.14159) / ((1.5 / 1.0) - (1e10 - (-2.25 - float64(0.1)))))
        v11 := GenS1{F0: 278}
        var v12 uint = 4294967295
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 float32 = ((1e10 / ((0.0 + float32((float64(0.1) * (float64(1e10) + 0.0)))) - (3.14159 * 1e10))) + 1.5)
        var v16 int = (fi14 / ((127) | 1))
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 string = "foo bar"
    _ = v13
    _ = v17
    return "\n\t"
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    var v4 int32 = 2147483647
    v4 = -591
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int64 = 0
        v20 := GenS1{F0: 0}
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        v4 -= (-(v4))
        v4 -= -1
    }
    v1 = (^(v1))
    v4 += int32(v1)
    if (-32768 >= 11) {
        var v21 uint64 = 42
        var v22 float32 = (0.0 / float32(3.14159))
        _ = v21
        _ = v22
    }
    fmt.Println(v4)
    v1 -= v1
    var v23 uint = 18446744073709551615
    v24 := []int{(v1 | v1), (v1 / (((v1 % ((v1) | 1))) | 1)), (-(v1))}
    var v25 float64 = float64((1e10 + 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
