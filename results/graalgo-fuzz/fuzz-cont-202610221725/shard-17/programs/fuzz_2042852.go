package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{(((("hello" + "a") + (("" + "") + "")) + "\n\t") + "foo bar"), "foo bar", "foo bar", ""}
    var v3 int16 = -32768
    var v4 float64 = 0.0
    var v5 uint8 = 128
    for fi6 := 0; fi6 < 15; fi6++ {
        v5 = 1
        v4 += (1.5 * ((v4 + 3.14159) - v4))
        var v7 float32 = 3.14159
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 42, F1: 0, F2: -52}
    _ = len(v1)
    var v9 int64 = 42
    v9 = v9
    v5 = (v5 >> ((4294967295) & 63))
    var v10 uint8 = v5
    var v11 uint16 = 2
    var v12 int64 = (v9 % (((v9 | int64(v3))) | 1))
    for fi13 := 0; fi13 < 7; fi13++ {
        v5 -= (^(v5))
        _ = fi13
    }
    v14 := GenS1{F0: v12, F1: 0, F2: -7}
    if (!(float64(float32(v4)) > ((0.1 * ((v4 * v4) - 1.5)) + v4))) {
        _ = len(v2)
        var v15 uint32 = 4294967295
        _ = v15
    } else {
        var v16 bool = (("\n\t" + ("" + "a")) > ("foo bar" + "\n\t"))
        var v17 int16 = (v3 - v3)
        _ = v16
        _ = v17
    }
    var v18 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
