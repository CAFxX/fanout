package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    var v1 float32 = float32(float64((1.5 * 1e10)))
    var v2 int16 = 0
    v3 := make(map[string]bool)
    var v4 int = 212
    var v5 float64 = 0.0
    v6 := GenS1{F0: (-(v2)), F1: (^(v4))}
    v4 = (v4 % ((-816) | 1))
    v2 += (-(v2))
    for fi7 := 0; fi7 < 20; fi7++ {
        v2 -= (v2 ^ (v2 & v2))
        v1 += 0.0
        if true {
            v8 := GenS1{F0: (-(v2)), F1: (fi7 >> ((0) & 63))}
            _ = v8
        }
        _ = fi7
    }
    fmt.Println(v4, v1)
    v1 += 1e10
    v5 -= (1e10 - float64(v1))
    for fi9 := 0; fi9 < 10; fi9++ {
        switch "\n\t" {
        case "":
            v1 = v1
        case "a":
            v2 = 0
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v4 -= v4
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v5, v5, fi10)
            v4 += (-(fi11))
            fmt.Println(v1, v2, v5)
            _ = fi11
        }
        _ = fi10
    }
    v3[(("foo bar" + "") + "foo bar")] = ("hello" != "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
