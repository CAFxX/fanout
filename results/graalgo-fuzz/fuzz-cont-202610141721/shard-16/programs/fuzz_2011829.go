package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 uint = 255
    var v3 uint32 = 42
    v4 := []string{"foo bar", "", ""}
    v5 := []uint32{(v3 | 564), (v3 << ((4294967295) & 63)), (v3 ^ (v3 / ((0) | 1)))}
    var v6 uint16 = 471
    delete(v1, 2147483647)
    v6 = 740
    var v7 float64 = (1.5 / (((1e10 / ((1e10 - 0.0) + 1.5)) * 1e10) / (((3.14159 / 1.5) + (1e10 - float64(((-2.25 / 1.0) + 0.0)))) * 1.5)))
    fmt.Println(v7, v3, v7)
    var v8 uint = v2
    if (!true) {
        v7 += ((1e10 * v7) - 0.0)
    }
    delete(v1, -781)
    switch "\n\t" {
    case "foo bar":
        for fi9 := 0; fi9 < 20; fi9++ {
            v6 = (v6 + (v6 | v6))
            var v10 float64 = (v7 / v7)
            _ = fi9
            _ = v10
        }
    default:
        v3 -= v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
