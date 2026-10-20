package main

import (
    "fmt"
)


func main() {
    v1 := []string{"hello", "\n\t", ((("a" + "\n\t") + "") + "\n\t"), ""}
    var v2 uint16 = 1
    var v3 uint8 = 42
    v3 -= v3
    v2 -= v2
    v1 = append(v1, "foo bar")
    var v4 int8 = 42
    if (true || (false && (true || true))) {
        v4 = -89
        switch "" {
        case "\n\t":
            var v5 bool = (((v4 << ((932) & 63)) != 1) && (1e10 >= float32(0.0)))
            _ = v5
        case "a":
            v2 = (v2 & (v2 - v2))
        case "":
            var v6 int64 = -772
            _ = v6
        default:
            var v7 uint8 = (v3 * 127)
            _ = v7
        }
        fmt.Println(v3, v3, v4)
    } else {
        var v8 float32 = (0.1 - 3.14159)
        _ = v8
    }
    v2 += v2
    v2 -= (^(v2))
    _ = len(v1)
    fmt.Println(v4)
    var v9 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
