package main

import (
    "fmt"
)

func genFn0(p0 bool) uint64 {
    var v6 uint64 = 186
    var v7 uint8 = 140
    _ = p0
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 float32 = (0.1 - -2.25)
    var v2 uint8 = 125
    v3 := []int{-795, 128, -74, -2147483648}
    var v4 int16 = 0
    var v5 uint16 = 42
    fmt.Println(v1)
    var v8 uint8 = (v2 + v2)
    switch ("foo bar" + (("hello" + "a") + (("\n\t" + ("\n\t" + "")) + "hello"))) {
    case "hello":
        _ = len(v3)
    case "\n\t":
        v5 -= 125
    case "":
        var v9 int = 4294967295
        _ = v9
    default:
        fmt.Println(v1)
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v4 += v4
        var v11 uint8 = (v2 * (v8 - uint8(v5)))
        _ = fi10
        _ = v11
    }
    fmt.Println(v4)
    var v12 int64 = 42
    for fi13 := 0; fi13 < 5; fi13++ {
        v12 -= -420
        var v14 float64 = (3.14159 * (0.0 + ((((3.14159 - float64((((0.1 - v1) - 1.5) + (v1 + v1)))) * (1.5 / 1e10)) - -2.25) - (3.14159 + (float64(v1) / 1.0)))))
        _ = fi13
        _ = v14
    }
    if len(v3) > 0 {
        v3[0] = 127
    }
    v8 = (v2 - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
