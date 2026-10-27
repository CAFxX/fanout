package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }

func main() {
    var v1 float32 = (1e10 - (0.1 - (3.14159 - 0.0)))
    var v2 int = 500
    var v3 int = (v2 % (((v2 - 2147483647)) | 1))
    var v4 uint = 255
    v1 = 0.1
    v2 -= -588
    fmt.Println(v4, v2)
    v5 := GenS1{F0: 221, F1: v4}
    v6 := GenS1{F0: 1, F1: 669}
    switch "" {
    case "\n\t":
        var v7 int32 = -125
        _ = v7
    case "a":
        fmt.Println(v1, v3, v2)
    case "hello":
        var v8 string = ("hello" + "hello")
        _ = v8
    default:
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int = (v3 | (v3 ^ (v3 + (v3 - (-(v3))))))
            _ = fi9
            _ = v10
        }
    }
    var v11 uint64 = 1
    var v12 string = ("foo bar" + "hello")
    if true {
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := []int{fi13, 9223372036854775807, (fi13 << ((v4) & 63))}
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: 191, F1: (^(v4))}
        v16 := []int{1}
        _ = v15
        _ = v16
    }
    v3 += 128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
