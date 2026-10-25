package main

import (
    "fmt"
)


func main() {
    var v1 string = ""
    var v2 uint = 65535
    var v3 float64 = 1e10
    v4 := make(map[string]uint32)
    v5 := make(map[string]int)
    fmt.Println(v1, v2, v1)
    v6 := []int{9223372036854775807, -1, -9223372036854775808}
    var v7 uint = (v2 * v2)
    v3 -= ((v3 - 3.14159) + v3)
    v2 -= v2
    v4[v1] = 4294967295
    var v8 string = "a"
    for fi9 := 0; fi9 < 3; fi9++ {
        if (-2.25 <= (((((-2.25 - 0.1) * (1e10 / (1.5 + 3.14159))) / (3.14159 + 0.0)) - 3.14159) + 0.1)) {
            delete(v5, v8)
            v2 = v2
        }
        v1 = v1
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
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
