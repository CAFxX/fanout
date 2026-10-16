package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 255
    var v3 int16 = -30
    var v4 bool = (42 <= 42)
    v5 := []uint32{2, 1, 0, 748}
    var v6 string = ("" + ("" + "hello"))
    _ = len(v1)
    var v7 int64 = 428
    switch (v6 + (((("a" + v6) + v6) + "foo bar") + "foo bar")) {
    case "\n\t":
        var v8 uint16 = v2
        _ = v8
    case "a":
        fmt.Println(v4, v2, v2)
    case "hello":
        var v9 float64 = float64((0.0 / 1.5))
        _ = v9
    default:
        var v10 int64 = 1
        _ = v10
    }
    if ((!((678 >= 0) || (3.14159 == (3.14159 - (((0.0 + 0.1) / -2.25) * (((-2.25 / 1.5) + 0.0) - 0.1)))))) && ((255 != 18446744073709551615) && (v6 <= "a"))) {
        if true {
            fmt.Println(v7, v4)
            v4 = (-1 >= 32767)
            switch -329 {
            case -1:
                var v11 int8 = -128
                _ = v11
            case -2147483648:
                if v4 {
                    v4 = v4
                    fmt.Println(v6, v6, v2)
                    var v12 uint64 = 18446744073709551615
                    _ = v12
                }
            case 9223372036854775807:
                v13 := GenS1{F0: 42, F1: 42}
                _ = v13
            }
        } else {
            v14 := GenS1{F0: 63, F1: -2147483648}
            _ = v14
        }
    }
    fmt.Println(v4, v4, v7)
    v2 = 255
    var v15 uint64 = 347
    fmt.Println(v3, v2)
    v1[804] = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
