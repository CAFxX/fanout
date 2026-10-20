package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int }

func main() {
    v1 := []string{"hello", "a", (((("foo bar" + "foo bar") + "hello") + ("hello" + ("" + ("\n\t" + "hello")))) + "")}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 float32 = (1.5 / 1.5)
    if true {
        v4 = (v4 + (((((float32(0.0) - ((0.0 * (1.5 / (float32(-2.25) * v4))) * (v4 * (float32((0.1 - 0.0)) * float32(1.5))))) * (float32(0.0) * v4)) - v4) + (v4 - (float32(3.14159) * v4))) + (v4 * v4)))
        var v5 uint8 = 2
        var v6 float64 = (1e10 + 0.1)
        _ = v5
        _ = v6
    } else {
        var v7 bool = true
        fmt.Println(v7)
        _ = v7
    }
    v4 += v4
    switch (("a" + (("" + "\n\t") + "hello")) + (("foo bar" + "hello") + "a")) {
    case "a":
        fmt.Println(v4)
    case "foo bar":
        for fi8 := 0; fi8 < 13; fi8++ {
            _ = len(v3)
            var v9 int = (fi8 * (fi8 / (((fi8 * (-(fi8)))) | 1)))
            switch ("a" + ("hello" + "hello")) {
            case "hello":
                if (!false) {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        v9 -= -1
                        _ = fi10
                    }
                } else {
                    v9 = v9
                }
            }
            _ = fi8
            _ = v9
        }
    case "\n\t":
        v11 := GenS1{F0: 65535, F1: 181, F2: 260}
        _ = v11
    }
    v12 := v3[2147483647]
    v4 = 0.1
    var v13 int64 = 0
    fmt.Println(v4)
    fmt.Println(v4, v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
