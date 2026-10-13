package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    var v1 uint = 42
    var v2 float32 = (3.14159 - (0.0 / (0.1 / 1.5)))
    var v3 float64 = 1.5
    if (32768 <= 128) {
        v3 -= v3
        v1 -= v1
        fmt.Println(v1)
    }
    v4 := GenS1{F0: -849, F1: 592}
    fmt.Println(v2, v1)
    var v5 uint16 = 2
    v6 := []int{0, -2147483648}
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v3, v1)
        fmt.Println(v1, v5, v5)
        v3 += (-2.25 - 1.5)
        _ = fi7
    }
    v1 = 42
    _ = len(v6)
    var v8 uint = v1
    switch "hello" {
    case "\n\t":
        v5 += (^(v5))
    default:
        if (("" + "foo bar") == ("a" + (("a" + "") + "hello"))) {
            v2 += 1.5
        }
    }
    var v9 uint8 = 128
    if true {
        v5 -= v5
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v5)
            _ = fi10
        }
        v2 -= (float32(v3) / v2)
    } else {
        var v11 uint16 = v5
        _ = v11
    }
    v1 = 18446744073709551615
    if true {
        _ = len(v6)
    } else {
        switch ("hello" + "hello") {
        case "a":
            if (!true) {
                v12 := GenS1{F0: 32767, F1: (^(v5))}
                fmt.Println(v3, v9, v5)
                _ = v12
            }
        case "hello":
            fmt.Println(v2)
        case "foo bar":
            for fi13 := 0; fi13 < 14; fi13++ {
                v9 = (v9 & 127)
                _ = fi13
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
