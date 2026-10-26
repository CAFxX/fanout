package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 uint32; F1 uint; F2 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 1
    var v3 string = "foo bar"
    var v4 int8 = 1
    var v5 int = 32768
    var v6 float64 = 1.5
    v7 := []int{94, (v5 + v5)}
    _ = len(v1)
    if (!(18446744073709551615 != 18446744073709551615)) {
        fmt.Println(v3, v3)
        v7 = append(v7, v5)
        fmt.Println(v3, v4, v4)
    }
    v8 := GenS1{F0: (-(v5)), F1: 0, F2: (v3 + "\n\t")}
    v9 := GenS2{F0: 42, F1: 4294967295, F2: v2}
    var v10 float32 = 1.5
    v4 += int8(v2)
    v2 -= 1
    var v11 float64 = (v6 + (v6 - (1.5 * ((-2.25 - v6) - (0.1 * (v6 - v6))))))
    fmt.Println(v4, v10, v6)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v12 int32 = -162
    var v13 int32 = 2147483647
    v12 = -2147483648
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v4)
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
