package main

import (
    "fmt"
)


func main() {
    var v1 int = 1
    var v2 uint8 = 0
    var v3 uint8 = 128
    var v4 int = (v1 | v1)
    var v5 int16 = -178
    var v6 float32 = (0.1 + (1e10 / 1.0))
    switch "a" {
    case "a":
        var v7 string = ("" + "a")
        _ = v7
    case "\n\t":
        v4 = (v4 * v4)
    }
    v8 := []int{-9223372036854775808}
    for fi9 := 0; fi9 < 7; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v6 = v6
        _ = fi9
    }
    v6 = v6
    var v10 uint32 = 1
    v4 = (-(v4))
    fmt.Println(v3, v10)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(v3, v1, v10)
        var v12 int = (-(fi11))
        var v13 int16 = (v5 / ((int16(v3)) | 1))
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
