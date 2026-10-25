package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
func genFn0(p0 uint64, p1 uint32, p2 int32) uint {
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint32 = p1
        for fi6 := 0; fi6 < 14; fi6++ {
            v5 -= (^(v5))
            _ = fi6
        }
        _ = fi4
        _ = v5
    }
    v7 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: int(p2)}
    var v8 int = 42
    v9 := GenS1{F0: 79, F1: p2, F2: v8}
    p1 += uint32(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 bool = false
    var v2 float32 = 0.1
    var v3 float64 = 0.0
    v10 := []int{32767, 351, 2147483647, 9223372036854775807}
    var v11 int16 = -32768
    if len(v10) > 0 {
        v10[0] = -361
    }
    v12 := GenS1{F0: genFn0(18446744073709551615, 543, -2147483648), F1: 2147483647, F2: 9223372036854775807}
    for fi13 := 0; fi13 < 1; fi13++ {
        v3 -= 1.5
        fmt.Println(fi13, fi13, v3)
        _ = fi13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v1 = (v11 != (-(v11)))
        _ = fi14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint32 = 442
        _ = fi15
        _ = v16
    }
    var v17 float64 = v3
    v3 += (-2.25 / ((0.1 * ((0.0 + 1e10) - v3)) + float64(-2.25)))
    switch (("" + "") + (("hello" + (("hello" + "\n\t") + ((("hello" + ("\n\t" + "foo bar")) + "\n\t") + "a"))) + "\n\t")) {
    case "a":
        var v18 int32 = -2147483648
        _ = v18
    }
    var v19 uint16 = 2
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v19)
        _ = fi20
    }
    var v21 uint32 = uint32(v19)
    fmt.Println(v2, v19)
    var v22 int32 = -705
    var v23 uint16 = 32767
    if len(v10) > 0 {
        v10[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
