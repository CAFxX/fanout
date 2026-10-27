package main

import (
    "fmt"
)


func main() {
    var v1 float64 = float64(((((float32((0.0 / ((float64(0.1) * (1e10 + -2.25)) - float64(float32(-2.25))))) - float32(1.5)) * -2.25) - float32(0.0)) + 1.5))
    var v2 uint8 = 42
    v3 := []uint32{}
    var v4 string = "foo bar"
    fmt.Println(v1)
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int64 = 272
        var v7 int32 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    v1 = v1
    var v9 int16 = 1000
    v1 -= -2.25
    v4 = v4
    v4 = ((((v4 + "hello") + ("a" + v4)) + v4) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
