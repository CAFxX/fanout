package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    var v1 int = 2147483647
    var v2 bool = false
    v3 := make(map[int]string)
    var v4 uint8 = 0
    var v5 int64 = 43
    fmt.Println(v4, v2)
    var v6 uint = 772
    var v7 int16 = 700
    var v8 string = ("" + "hello")
    v6 -= v6
    v1 = (v1 / ((v1) | 1))
    var v9 float32 = ((((float32(-2.25) / 1.0) * (((3.14159 * float32(0.0)) + 0.0) + 3.14159)) + float32(0.1)) / 1.0)
    fmt.Println(v9)
    var v10 uint64 = 18446744073709551615
    v8 = (v8 + (v8 + "a"))
    fmt.Println(v1, v1, v9)
    var v11 int = (v1 + (v1 / ((-173) | 1)))
    var v12 uint64 = v10
    if true {
        v13 := GenS1{F0: 2147483647, F1: (v6 >> (((v6 % ((607) | 1))) & 63))}
        v14 := []int{}
        _ = v13
        _ = v14
    } else {
        switch v8 {
        case "":
            if v2 {
                var v15 int32 = 42
                var v16 uint16 = 255
                _ = v15
                _ = v16
            }
        case "\n\t":
            fmt.Println(v12)
        case "hello":
            switch "" {
            case "":
                v17 := []int{}
                _ = v17
            default:
                v5 = (v5 >> ((v6) & 63))
            }
        }
        var v18 int64 = v5
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
