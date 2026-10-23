package main

import (
    "fmt"
)

func genFn0(p0 int16, p1 int32) uint32 {
    p0 = 42
    p1 = p1
    var v4 int = -709
    var v5 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 317
}

func main() {
    v1 := []uint32{65535, 561}
    var v2 float64 = float64(float32(0.1))
    var v3 uint16 = 42
    v2 += (v2 + (v2 + 1e10))
    var v6 uint64 = 1
    var v7 int = 4294967295
    var v8 int8 = int8(v3)
    var v9 uint32 = 675
    var v10 float64 = float64(3.14159)
    switch ("\n\t" + "a") {
    case "foo bar":
        var v11 string = "foo bar"
        _ = v11
    case "":
        if true {
            fmt.Println(v6, v8, v7)
            fmt.Println(v10, v2)
        }
    default:
        fmt.Println(v10, v10, v2)
    }
    if ((-2.25 != 3.14159) || false) {
        v7 = (v7 & v7)
    }
    var v12 uint16 = (v3 / ((v3) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
