package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int }

func main() {
    var v1 bool = false
    var v2 int = 127
    v3 := make(map[uint32]string)
    if (((0.0 == 0.1) && v1) || v1) {
        v2 = v2
        var v4 float64 = 1e10
        _ = v4
    }
    var v5 int16 = 0
    v6 := GenS1{F0: v2, F1: (42 != 18446744073709551615), F2: -492}
    for fi7 := 0; fi7 < 5; fi7++ {
        if (!((v2 | (fi7 << ((345) & 63))) <= fi7)) {
            var v8 int = (fi7 << ((375) & 63))
            _ = v8
        }
        v5 -= 42
        v9 := []int{fi7, (^(v2)), v2, (v2 - fi7)}
        _ = fi7
        _ = v9
    }
    delete(v3, 1)
    v10 := []int{v2, v2}
    switch "hello" {
    case "\n\t":
        v5 = v5
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println(v1, v2)
    v3[2] = "hello"
    v5 += 1
    var v11 int = -2147483648
    v12 := v3[4294967295]
    for fi13 := 0; fi13 < 17; fi13++ {
        fmt.Println(fi13, v12, v11)
        for fi14 := 0; fi14 < 15; fi14++ {
            if false {
                var v15 string = (v12 + (v12 + v12))
                v16 := GenS1{F0: 42, F1: (float64((-2.25 * (((1e10 + 1e10) / (3.14159 + -2.25)) * (3.14159 / ((((1.5 / 1.0) * -2.25) / 0.1) / 1.0))))) > 0.0), F2: fi14}
                _ = v15
                _ = v16
            }
            v12 = "hello"
            v2 = fi14
            _ = fi14
        }
        _ = fi13
    }
    v17 := GenS1{F0: 32768, F1: v1, F2: (-(v11))}
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := v3[953]
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
