package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }

func main() {
    var v1 uint32 = 819
    var v2 int8 = -63
    var v3 float64 = 1e10
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := []int{-46}
            fmt.Println(v1)
            _ = fi5
            _ = v6
        }
        fmt.Println(fi4)
        _ = fi4
    }
    v7 := GenS1{F0: (v2 * (v2 + (v2 | (v2 | 42)))), F1: 1, F2: (("\n\t" + (("" + "\n\t") + "\n\t")) + "foo bar")}
    v8 := GenS1{F0: v2, F1: 1, F2: ("" + "a")}
    v1 -= v1
    v3 += v3
    var v9 int8 = (v2 + (-(v2)))
    v9 += -30
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 float32 = (0.1 * 1.5)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v1 = v1
        v13 := []int{128, (fi12 % ((fi12) | 1)), fi12}
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
