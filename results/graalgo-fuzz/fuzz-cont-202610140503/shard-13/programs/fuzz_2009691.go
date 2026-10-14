package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -1
    v3 := []string{"", ("a" + "hello")}
    v2 = (v2 | 9223372036854775807)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v4 bool = ((((((3.14159 * float32(3.14159)) + float32(1.5)) / 1.0) - 3.14159) - 1.5) > (((((-2.25 * 0.1) - (0.0 * 1.5)) + float32(0.0)) * (0.1 + 1e10)) / 1.5))
    var v5 int16 = 42
    if false {
        switch "a" {
        case "hello":
            _ = len(v3)
        default:
            _ = len(v1)
        }
        if false {
            var v6 uint16 = 2
            fmt.Println(v6)
            _ = v6
        } else {
            var v7 uint32 = 2
            _ = v7
        }
    }
    v2 -= v2
    var v8 float32 = ((1.5 * (float32(-2.25) / 1.5)) - 0.0)
    fmt.Println(v4)
    v3 = append(v3, ("\n\t" + ("a" + "")))
    v9 := GenS1{F0: 1, F1: 294, F2: 4294967295}
    v5 += (^(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
