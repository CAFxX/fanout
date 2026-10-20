package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint) string {
    var v4 float32 = ((-2.25 / 1.5) * 1e10)
    fmt.Println(v4, p0)
    var v5 uint8 = 42
    _ = p0
    _ = v4
    _ = v5
    return "\n\t"
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]bool)
    var v3 int32 = -117
    var v6 bool = (false && false)
    if v6 {
        v1 = 42
    } else {
        var v7 uint64 = 1
        fmt.Println(v7)
        _ = v7
    }
    var v8 uint32 = 42
    var v9 int = 177
    v10 := v2["\n\t"]
    switch "\n\t" {
    case "":
        var v11 float64 = 1.5
        _ = v11
    case "hello":
        fmt.Println(v3, v1, v9)
    case "a":
        for fi12 := 0; fi12 < 15; fi12++ {
            delete(v2, "a")
            var v13 int64 = -1
            var v14 int = (fi12 - v9)
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        fmt.Println(v3, v8)
    }
    v15 := []int{v9, -9223372036854775808, 32768}
    if (-1 != -1) {
        fmt.Println(v10, v9)
        var v16 float32 = 0.0
        _ = len(v15)
        _ = v16
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS1{F0: v1, F1: v8}
            fmt.Println(v8, v1, v9)
            _ = fi17
            _ = v18
        }
        v3 = 967
    }
    var v19 int32 = -2147483648
    v19 += v19
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
