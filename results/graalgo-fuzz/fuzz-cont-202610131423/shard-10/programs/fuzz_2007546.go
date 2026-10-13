package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    var v2 int = -900
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    var v5 int16 = -604
    var v6 float32 = (-2.25 + (1e10 + (0.1 - ((3.14159 + (((0.1 / ((3.14159 + 0.0) / 0.1)) + 0.1) * (-2.25 / 1.5))) + 1e10))))
    v5 -= 1
    for fi7 := 0; fi7 < 15; fi7++ {
        v2 += (v2 - fi7)
        var v8 int = 9223372036854775807
        v2 += (fi7 % ((0) | 1))
        _ = fi7
        _ = v8
    }
    v5 -= (v5 + (v5 >> ((255) & 63)))
    if false {
        v2 = (^(v2))
        fmt.Println(v5)
    } else {
        fmt.Println(v2, v2)
        var v9 uint8 = 0
        _ = v9
    }
    v2 = (v2 & 288)
    fmt.Println(v6, v5, v5)
    if false {
        fmt.Println(v2, v6, v5)
        var v10 uint16 = 32767
        var v11 int16 = v5
        _ = v10
        _ = v11
    }
    v12 := []int{72, -975, (v2 | (v2 | v2)), int(v5)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
