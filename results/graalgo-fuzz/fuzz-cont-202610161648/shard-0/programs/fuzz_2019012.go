package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []string{("a" + ""), "\n\t", "hello"}
    var v2 uint16 = 566
    var v3 int8 = 42
    var v4 bool = false
    var v5 int64 = 9223372036854775807
    switch "a" {
    case "hello":
        v4 = (!(18446744073709551615 != 1))
    default:
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    fmt.Println(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    for fi7 := 0; fi7 < 6; fi7++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch -2147483648 {
        case 2147483647:
            var v8 uint64 = 297
            _ = v8
        case 127:
            if (fi7 < fi7) {
                var v9 int = fi7
                _ = v9
            } else {
                fmt.Println(v5)
            }
        }
        var v10 float32 = float32((-2.25 - float64(3.14159)))
        _ = fi7
        _ = v10
    }
    v4 = v4
    var v11 uint64 = 2
    v12 := GenS1{F0: 128}
    fmt.Println(v11)
    v3 -= int8(v11)
    fmt.Println(v4, v5, v11)
    var v13 float64 = (0.0 + 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
