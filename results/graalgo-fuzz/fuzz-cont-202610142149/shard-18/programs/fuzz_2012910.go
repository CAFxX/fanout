package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }
func genFn0() uint {
    var v4 int64 = -249
    fmt.Println(v4)
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{1, (fi5 % ((-9223372036854775808) | 1))}
        if len(v6) > 0 {
            v6[0] = fi5
        }
        _ = fi5
        _ = v6
    }
    _ = v4
    return 792
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v7 uint16 = 757
    v8 := GenS1{F0: 165, F1: -9223372036854775808, F2: 42}
    fmt.Println(v7)
    var v9 int64 = 254
    var v10 int16 = -32768
    var v11 int16 = (v10 >> ((4294967295) & 63))
    v11 += v11
    var v12 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
