package main

import (
    "fmt"
)

func genFn0(p0 uint16) uint8 {
    var v4 uint16 = (p0 | (p0 * 317))
    v5 := []int{2147483647, 0, 4294967295, 32768}
    switch "\n\t" {
    case "a":
        v5 = append(v5, 32768)
    default:
        p0 += (v4 + 1)
    }
    var v6 int64 = 843
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 38
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int8 = 42
    var v3 int = 127
    for fi7 := 0; fi7 < 3; fi7++ {
        v3 -= v3
        fmt.Println(v1, fi7, fi7)
        _ = fi7
    }
    v2 = v2
    var v8 int = (v3 & v3)
    var v9 uint32 = v1
    v10 := []int{127, (v8 << ((255) & 63)), v3}
    var v11 int64 = 39
    var v12 string = "a"
    if (!true) {
        var v13 float64 = 0.0
        _ = v13
    }
    fmt.Println(v2)
    fmt.Println(v1)
    var v14 uint32 = v1
    fmt.Println(v2, v8, v8)
    switch v12 {
    case "hello":
        v3 -= v3
    }
    v8 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
