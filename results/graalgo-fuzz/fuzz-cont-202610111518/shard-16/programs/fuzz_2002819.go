package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint8 = 100
    v2 := make(map[string]uint32)
    var v3 bool = (!false)
    v4 := []int{706, 42, 127, 622}
    var v5 int = 9223372036854775807
    v3 = v3
    fmt.Println(v3)
    v5 -= v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 float64 = 1e10
    switch "foo bar" {
    case "hello":
        if (-1 > 9223372036854775807) {
            if v3 {
                v5 -= v5
                v7 := GenS1{F0: "\n\t"}
                _ = v7
            } else {
                var v8 float32 = 0.1
                _ = v8
            }
        }
    default:
        fmt.Println(v6, v6, v3)
    }
    v3 = ((-2147483648 < 42) && (2147483647 >= 4294967295))
    v1 = v1
    fmt.Println(v3)
    fmt.Println(v1, v1)
    v9 := GenS1{F0: ""}
    var v10 uint32 = 1
    fmt.Println(v6, v10, v10)
    switch (-(v5)) {
    case -142:
        switch 982 {
        case 32768:
            v11 := GenS1{F0: "\n\t"}
            _ = v11
        case 42:
            v6 = 1e10
        case 801:
            var v12 uint = 42
            _ = v12
        default:
            v3 = (int8(v5) >= 42)
        }
    case 347:
        var v13 string = "foo bar"
        _ = v13
    default:
        v5 = (v5 % ((v5) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
