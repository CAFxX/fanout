package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
func genFn0(p0 int32, p1 int) int {
    if (true && false) {
        var v6 float32 = (0.1 + (((1.5 * (1.5 * 0.1)) / 0.1) / 1.0))
        _ = v6
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return (p1 & 128)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 65535
    var v3 int32 = 2147483647
    v4 := []string{("foo bar" + ("\n\t" + ("foo bar" + "hello"))), "foo bar"}
    v5 := make(map[uint32]int)
    if false {
        v3 = (v3 & -2147483648)
        fmt.Println(v2, v2)
    } else {
        var v7 uint32 = 0
        _ = v7
    }
    v3 += (v3 >> ((v2) & 63))
    var v8 bool = (v3 <= (^(v3)))
    var v9 float64 = 0.1
    fmt.Println(v8)
    fmt.Println(v3, v8, v9)
    v2 = v2
    var v10 uint64 = 635
    v4 = append(v4, "foo bar")
    var v11 float64 = v9
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 bool = v8
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: -136, F1: (v3 / ((0) | 1)), F2: 255}
    var v15 uint8 = 255
    var v16 float64 = v11
    var v17 uint8 = (v15 << (((v2 - (^(v2)))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
