package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
func genFn0(p0 string, p1 bool, p2 int32) uint8 {
    fmt.Println(p2, p1, p2)
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 bool = p1
            var v9 int8 = -3
            var v10 int32 = -2147483648
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        v11 := []int{(fi6 - fi6), fi6}
        _ = fi6
        _ = v11
    }
    v12 := GenS1{F0: 65535, F1: 1, F2: -553}
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint64 = 485
            v16 := []int{(fi13 ^ fi13), fi13, fi14, (fi13 - 127)}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    var v17 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v17
    return 2
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 65535
    var v3 float32 = (0.0 + -2.25)
    var v4 int16 = 42
    var v5 float64 = (((((-2.25 * ((3.14159 * (0.1 / 1.0)) - -2.25)) / (1.5 * 0.1)) / -2.25) + 1.5) / 3.14159)
    v3 = (v3 * 0.0)
    fmt.Println(v3)
    var v18 uint64 = uint64(v2)
    for fi19 := 0; fi19 < 9; fi19++ {
        v2 = (v2 << ((731) & 63))
        var v20 bool = false
        _ = fi19
        _ = v20
    }
    v21 := []int{42, -899}
    v22 := GenS1{F0: (v2 << ((42) & 63)), F1: 289, F2: -300}
    var v23 int16 = (v4 * (-(v4)))
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
