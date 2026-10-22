package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -128
    var v2 string = "foo bar"
    var v3 int64 = 0
    var v4 int32 = 0
    var v5 int = -9223372036854775808
    var v6 bool = (42 != -24)
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(v5 % ((fi7) | 1)), -2147483648, v5}
        var v9 int64 = (v3 + 1)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v3 = (v3 % ((0) | 1))
        _ = fi10
    }
    var v11 float64 = 3.14159
    v11 += (((1e10 - v11) * 0.0) * (0.1 + 3.14159))
    v12 := []int{v5}
    v11 += 1e10
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            v11 -= v11
            fmt.Println(v5)
            v12 = append(v12, -9223372036854775808)
            _ = fi14
        }
        _ = fi13
    }
    var v15 uint32 = 65535
    var v16 string = ("foo bar" + ("" + v2))
    switch "" {
    case "":
        v12 = append(v12, (v5 + 2147483647))
    case "a":
        v15 += v15
    case "foo bar":
        var v17 float32 = ((-2.25 * ((0.1 / (1.5 + -2.25)) * (0.0 * 0.0))) * 0.1)
        _ = v17
    default:
        var v18 uint32 = 2147483648
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
