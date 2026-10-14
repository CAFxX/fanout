package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 int16 = 1
    var v5 uint16 = 32767
    v4 = (-(v4))
    fmt.Println(v5)
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 555}
        var v9 uint32 = 1
        fmt.Println(v6, v5, v6)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v3)
    v4 -= v4
    var v10 uint32 = 4294967295
    switch "\n\t" {
    case "\n\t":
        v5 -= (^(v5))
    case "hello":
        fmt.Println(v10)
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        switch ("foo bar" + "") {
        case "hello":
            v5 = (v5 - (v5 ^ (v5 | 869)))
        }
        v12 := []int{(v1 + (v1 / ((fi11) | 1))), fi11, int(v4), -784}
        var v13 uint8 = (v6 % ((v6) | 1))
        _ = fi11
        _ = v12
        _ = v13
    }
    if (0.0 < -2.25) {
        var v14 bool = (true || false)
        _ = v14
    }
    v15 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
