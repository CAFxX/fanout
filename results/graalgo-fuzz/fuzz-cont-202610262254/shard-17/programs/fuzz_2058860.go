package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
func genFn0(p0 int32, p1 uint) int16 {
    var v5 uint = 760
    if false {
        var v6 uint8 = 0
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return -32768
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    switch "" {
    case "hello":
        v7 := GenS1{F0: "\n\t"}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 float64 = -2.25
            v10 := GenS1{F0: ("\n\t" + ("\n\t" + "foo bar"))}
            v9 -= (((((float64(0.1) + (v9 * (float64(-2.25) - v9))) + ((v9 / v9) + v9)) * (float64(1e10) - -2.25)) - (v9 * (v9 * 0.0))) / v9)
            _ = fi8
            _ = v10
            _ = v9
        }
    case "a":
        var v11 bool = (true || (false || (-42 < -128)))
        _ = v11
    }
    var v12 int = 32767
    v12 -= (v12 * v12)
    switch (("" + ("foo bar" + "foo bar")) + "hello") {
    case "hello":
        var v13 int32 = -2147483648
        _ = v13
    case "\n\t":
        v12 = v12
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS2{F0: 1, F1: 42, F2: 1}
        if false {
            v12 = (fi14 & 821)
        }
        v12 = (v12 % (((fi14 ^ (v12 / ((128) | 1)))) | 1))
        _ = fi14
        _ = v15
    }
    fmt.Println(v12)
    v12 = (-(v12))
    v12 = (v12 << ((2) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
