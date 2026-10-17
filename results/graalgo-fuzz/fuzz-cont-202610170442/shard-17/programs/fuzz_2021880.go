package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 string = ("a" + ((("a" + "") + "a") + ("foo bar" + "foo bar")))
    v2 := make(map[uint32]bool)
    var v3 uint32 = 677
    var v4 float64 = float64((1.5 / 3.14159))
    var v5 int16 = -1
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float32 = 3.14159
        _ = fi6
        _ = v7
    }
    var v8 int8 = -85
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 42}
        v3 += v3
        _ = fi9
        _ = v10
    }
    fmt.Println(v3, v1)
    var v11 float32 = float32((v4 * (3.14159 * v4)))
    var v12 uint32 = v3
    v2[65535] = (!true)
    var v13 int16 = v5
    if (!((4294967295 < 325) && false)) {
        var v14 uint64 = 18446744073709551615
        _ = v14
    }
    var v15 int16 = 1000
    switch "a" {
    case "a":
        var v16 uint64 = 2
        _ = v16
    case "hello":
        var v17 int32 = 0
        _ = v17
    case "foo bar":
        v3 = v12
    default:
        var v18 bool = ((v8 * (v8 | v8)) != v8)
        _ = v18
    }
    var v19 int16 = (^(v15))
    fmt.Println(v11, v13, v11)
    var v20 bool = (uint(v3) != 42)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
