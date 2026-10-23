package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int8, p1 int) uint8 {
    fmt.Println(p0, p0)
    v7 := []int{}
    var v8 float64 = 1.5
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 157
}

func main() {
    v1 := []int{9223372036854775807, -832, -496}
    var v2 uint16 = 1
    var v3 float64 = (0.0 + (((0.1 * 0.1) * 1.5) - 0.0))
    var v4 int64 = 9223372036854775807
    v5 := make(map[uint32]int)
    v6 := []string{"hello", "a", "a", "foo bar"}
    fmt.Println(v3, v3)
    switch 32768 {
    case 32767:
        fmt.Println(v2)
    }
    var v9 uint8 = 1
    v9 += v9
    var v10 string = "\n\t"
    delete(v5, 714)
    v3 = v3
    if true {
        v11 := GenS1{F0: 42, F1: (("hello" + (v10 + "")) + ("" + "foo bar"))}
        v9 -= v9
        _ = v11
    } else {
        var v12 int64 = v4
        v4 += (v12 >> ((1) & 63))
        var v13 int16 = -1
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        _ = len(v5)
        v15 := GenS2{F0: -927, F1: 65535}
        _ = fi14
        _ = v15
    }
    v16 := v5[365]
    v3 -= ((v3 + v3) - (3.14159 - v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
