package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 int32 = 42
    var v2 int32 = (v1 / ((v1) | 1))
    var v3 int = 369
    var v4 uint32 = 103
    fmt.Println(v1, v4, v1)
    var v5 int = (-(v3))
    var v6 uint = 18446744073709551615
    v5 -= -383
    v7 := []int{(v5 - v3), 479, v5, v3}
    v4 -= (v4 / (((v4 << ((v6) & 63))) | 1))
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = 1
        v4 += 65535
        var v10 uint16 = 255
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(v1, v4, fi11)
        _ = fi11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint8 = 2
        v1 += v1
        _ = fi12
        _ = v13
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v2 += 2147483647
        _ = fi14
    }
    var v15 string = ((("foo bar" + "foo bar") + "foo bar") + "hello")
    v16 := GenS1{F0: true, F1: (v5 ^ 32767)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
