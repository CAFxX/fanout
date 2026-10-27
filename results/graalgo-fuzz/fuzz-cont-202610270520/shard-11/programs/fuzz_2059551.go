package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
func genFn0(p0 int8, p1 uint16) int {
    var v7 float32 = 3.14159
    p1 = p1
    var v8 uint32 = 120
    switch 127 {
    case 32767:
        var v9 float32 = float32((3.14159 / ((1e10 * float64(v7)) / float64(1e10))))
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1() uint8 {
    var v10 int16 = 140
    v11 := []int{-13, 32768, 606}
    _ = v10
    _ = v11
    return 128
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = (3.14159 / 3.14159)
    var v3 bool = ((true || (false || (false || (v1 > 2)))) || false)
    v4 := make(map[uint32]int)
    var v5 uint = 1
    var v6 uint = (v5 >> ((uint(v1)) & 63))
    var v12 int32 = 2147483647
    v3 = false
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint8 = 2
            _ = fi14
            _ = v15
        }
        var v16 uint16 = v1
        _ = fi13
        _ = v16
    }
    fmt.Println(v6, v6)
    v3 = (!((((v2 * ((-2.25 + 3.14159) + v2)) != ((v2 - v2) - 1.5)) && (true || v3)) && v3))
    var v17 int8 = 1
    _ = len(v4)
    v18 := GenS1{F0: "foo bar", F1: 42, F2: v5}
    switch 2147483647 {
    case 2147483647:
        v19 := GenS1{F0: "", F1: 9223372036854775807, F2: (^(v6))}
        _ = v19
    case 1:
        v2 = (3.14159 * (((-2.25 * ((v2 * v2) - ((v2 / 3.14159) / v2))) * (v2 / v2)) * v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
