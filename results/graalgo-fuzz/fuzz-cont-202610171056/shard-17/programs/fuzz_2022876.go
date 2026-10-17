package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }

func main() {
    var v1 float64 = 3.14159
    v2 := []int{295, 384, 2147483647}
    v3 := make(map[string]int)
    v4 := []int{809, -846}
    var v5 int64 = 9223372036854775807
    v6 := []int{}
    v6 = append(v6, 393)
    var v7 string = (("hello" + ("" + "foo bar")) + ("foo bar" + ((("\n\t" + "") + "") + ("" + ("hello" + (("" + "a") + "foo bar"))))))
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 13, F1: 1, F2: 42}
        v10 := GenS1{F0: 2, F1: 954, F2: 255}
        _ = fi8
        _ = v10
        _ = v9
    }
    v5 = (v5 << ((1) & 63))
    var v11 uint8 = 49
    if len(v6) > 0 {
        v6[0] = 127
    }
    if (1 > 1) {
        v12 := GenS1{F0: v11, F1: 1, F2: 2}
        _ = v12
    } else {
        var v13 float64 = (v1 * v1)
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(fi14)
            _ = fi14
        }
        fmt.Println(v11, v5, v11)
        _ = v13
    }
    switch (v7 + ("a" + (v7 + (v7 + (v7 + (("\n\t" + v7) + v7)))))) {
    case "\n\t":
        var v15 float64 = (0.0 - v1)
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint = 18446744073709551615
            _ = fi16
            _ = v17
        }
    }
    var v18 int = -668
    fmt.Println(v1, v5, v11)
    v1 = (v1 / 0.1)
    delete(v3, ("" + v7))
    v7 = v7
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
