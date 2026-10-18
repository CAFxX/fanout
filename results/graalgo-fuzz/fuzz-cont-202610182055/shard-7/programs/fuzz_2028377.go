package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint8, p1 uint16) int {
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 4; fi6++ {
        p0 += (^(p0))
        _ = fi6
    }
    v7 := []int{-739, 32768, 127}
    p0 += 168
    var v8 float64 = (-2.25 + ((1.5 * ((-2.25 * 1e10) / 1.5)) + (float64(1.5) * 3.14159)))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 42
    v3 := []string{"a", ("a" + "foo bar"), "\n\t"}
    var v4 uint8 = 42
    var v5 int64 = 51
    v2 += (v2 * (v2 % (((v2 + (-(v2)))) | 1)))
    v4 = (v4 >> ((0) & 63))
    var v9 int16 = -665
    fmt.Println(v2, v4)
    var v10 uint16 = 65535
    var v11 uint8 = (v4 % (((v4 % ((v4) | 1))) | 1))
    var v12 float64 = ((1.5 + ((-2.25 + 3.14159) * 1e10)) * 3.14159)
    var v13 int16 = (v9 << ((378) & 63))
    v14 := GenS1{F0: (v13 & 383), F1: (^(v4))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
