package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0() uint64 {
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: 1, F1: 0}
    v6 += v6
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 float64 = 1e10
    var v4 string = ("hello" + "")
    v5 := []string{"hello", ("\n\t" + v4)}
    var v8 float64 = v3
    _ = len(v2)
    if len(v5) > 0 {
        v5[0] = ""
    }
    fmt.Println(v8, v8)
    v8 += ((0.1 + (0.1 * (1.5 + (1.5 * float64(float32(0.1)))))) * (((3.14159 / float64(3.14159)) / (v3 / 3.14159)) / 1.5))
    switch v4 {
    case "\n\t":
        v5 = append(v5, (v4 + ("foo bar" + v4)))
    }
    v3 = v3
    for fi9 := 0; fi9 < 4; fi9++ {
        v3 -= 3.14159
        _ = fi9
    }
    v4 = ("hello" + v4)
    var v10 int16 = 0
    v8 = (v3 / v3)
    v3 += v8
    v5 = append(v5, v4)
    var v11 uint8 = 2
    v10 -= (-(v10))
    var v12 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
