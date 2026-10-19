package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = -190
    var v3 uint = 0
    var v4 int16 = -657
    v1 -= (v1 - -9223372036854775808)
    var v5 uint16 = 42
    v1 = v1
    v6 := []int{v2}
    for fi7 := 0; fi7 < 20; fi7++ {
        v2 -= (fi7 | (v2 + -2147483648))
        _ = fi7
    }
    switch "foo bar" {
    case "hello":
        v1 += 9223372036854775807
    case "a":
        v4 += (-(v4))
    }
    switch "\n\t" {
    case "foo bar":
        var v8 int = v2
        _ = v8
    case "a":
        var v9 int16 = (v4 / ((v4) | 1))
        _ = v9
    }
    fmt.Println(v4)
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int = fi10
            var v13 uint64 = 18446744073709551615
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            v1 = v1
            v15 := GenS1{F0: (v1 - 1), F1: v5}
            var v16 int16 = v4
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
