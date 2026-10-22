package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint, p1 int16) uint32 {
    p0 = (p0 + p0)
    switch 237 {
    case 127:
        var v5 bool = (false || true)
        _ = v5
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    var v1 float64 = 3.14159
    var v2 int8 = -98
    var v3 uint32 = 0
    v4 := []string{"a"}
    v6 := GenS1{F0: 6}
    fmt.Println(v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= 1e10
    v7 := GenS1{F0: 2}
    var v8 string = "foo bar"
    var v9 int = 42
    var v10 uint8 = 95
    switch "a" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        var v11 int32 = 542
        _ = v11
    }
    v12 := GenS1{F0: (v10 | 255)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
