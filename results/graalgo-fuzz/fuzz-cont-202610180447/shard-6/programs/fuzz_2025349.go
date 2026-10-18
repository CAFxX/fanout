package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := []string{"foo bar", "a", "hello", "a"}
    v4 := make(map[uint32]bool)
    var v5 int16 = -1
    v4[2147483648] = false
    v5 -= -977
    v5 = (v5 + (-(v5)))
    v5 = (^(v5))
    var v6 bool = (true && (-666 == 250))
    var v7 float32 = (((1.5 - (((0.1 - 1.5) / (((0.1 - (1e10 / (float32(-2.25) / 1.0))) / 1e10) * (((0.1 - -2.25) * -2.25) - float32(1.5)))) * float32(float64(1.5)))) / 3.14159) / (1.5 * (0.1 + 1.5)))
    v6 = (0 != -1)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 string = ("\n\t" + "hello")
        _ = fi8
        _ = v9
    }
    fmt.Println(v7)
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
