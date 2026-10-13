package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 4294967295
    var v3 float32 = (1.5 * float32(3.14159))
    v4 := make(map[uint32]bool)
    v5 := []string{("\n\t" + "\n\t"), "a", ("foo bar" + "\n\t"), ("hello" + ((("hello" + ("\n\t" + "\n\t")) + "a") + "foo bar"))}
    var v6 int64 = 9223372036854775807
    fmt.Println(v2, v1)
    switch "hello" {
    case "hello":
        v1 += 831
    case "\n\t":
        fmt.Println(v1, v1, v1)
    }
    var v7 float64 = (0.0 / float64(1e10))
    if false {
        for fi8 := 0; fi8 < 16; fi8++ {
            v3 -= ((v3 / (-2.25 / v3)) + 3.14159)
            for fi9 := 0; fi9 < 13; fi9++ {
                v2 -= (v2 % (((^(v2))) | 1))
                _ = fi9
            }
            _ = fi8
        }
        fmt.Println(v1)
    } else {
        var v10 int8 = -128
        v7 -= (1.5 * v7)
        var v11 int8 = 100
        _ = v10
        _ = v11
    }
    var v12 float32 = (((-2.25 / (((0.1 + v3) / v3) / float32((v7 + -2.25)))) / 1.0) * v3)
    var v13 uint32 = (v2 / ((v2) | 1))
    var v14 uint32 = v2
    var v15 uint = 22
    v16 := GenS1{F0: ("" + ("" + "hello"))}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
