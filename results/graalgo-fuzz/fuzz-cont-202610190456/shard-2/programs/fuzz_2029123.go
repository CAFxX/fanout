package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
func genFn0(p0 string) uint32 {
    var v6 uint = 215
    v6 -= (v6 >> ((v6) & 63))
    v7 := GenS1{F0: -117, F1: -2147483648}
    _ = p0
    _ = v6
    _ = v7
    return 907
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 65535
    v3 := []int{-779, -1, -2147483648, 128}
    var v4 bool = true
    var v5 float64 = 0.0
    v4 = v4
    if (((1 == 255) && (v4 && (!v4))) || v4) {
        fmt.Println(v4, v5, v4)
    } else {
        switch 2147483647 {
        case 127:
            v3 = append(v3, -171)
        }
    }
    v5 -= 1e10
    if v4 {
        v8 := GenS1{F0: -52, F1: -2147483648}
        _ = v8
    } else {
        var v9 float32 = ((1.5 - ((((1.5 * (1e10 / 1.0)) * (((((1e10 + 0.0) + (0.1 * float32(v5))) - 0.0) * 1.5) - (0.0 / 0.1))) - float32(((v5 / v5) / (1e10 + v5)))) + (float32(1e10) - (0.1 * 3.14159)))) / (((0.1 * float32(1e10)) / 1.0) - (0.0 - 1.5)))
        _ = v9
    }
    switch -34 {
    case 42:
        fmt.Println(v4, v2, v4)
    case 188:
        v10 := GenS1{F0: 11, F1: 1}
        _ = v10
    default:
        fmt.Println(v4)
    }
    switch ((("" + "\n\t") + (("hello" + ("a" + "hello")) + "\n\t")) + ("" + ("\n\t" + "foo bar"))) {
    case "\n\t":
        _ = len(v1)
    }
    var v11 int32 = -889
    var v12 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
