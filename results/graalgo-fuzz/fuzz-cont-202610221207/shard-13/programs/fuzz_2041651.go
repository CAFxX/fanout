package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 50
    v3 := []int{983, 9223372036854775807, -2147483648}
    v4 := make(map[string]uint32)
    v2 = v2
    v5 := v1[""]
    v6 := GenS1{F0: 1, F1: 18446744073709551615}
    fmt.Println(v5, v2)
    fmt.Println(v2, v2)
    fmt.Println(v5, v5)
    v7 := v1["foo bar"]
    var v8 uint16 = 42
    fmt.Println(v2, v7)
    if (!(!(-2.25 != (float64(0.1) * (0.1 + (((-2.25 - 0.0) / (1e10 * ((-2.25 + 0.1) * -2.25))) + 1e10)))))) {
        v5 = ((v7 + v5) + v7)
        v4[v7] = 320
        delete(v1, ("a" + ""))
    }
    if ((0 > 1) && false) {
        fmt.Println(v7, v5, v8)
        v8 -= v8
    } else {
        switch v5 {
        case "hello":
            v9 := GenS1{F0: -1, F1: 2}
            _ = v9
        default:
            var v10 float64 = 0.0
            _ = v10
        }
        if false {
            v5 = (v7 + ("hello" + "foo bar"))
            v11 := v1[v5]
            var v12 int32 = -2147483648
            _ = v11
            _ = v12
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
