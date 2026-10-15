package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float32 = float32(0.0)
    v2 := []uint32{231, 2147483647, 0, 76}
    var v3 float64 = 0.0
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    v1 = ((v1 * v1) + v1)
    v1 += -2.25
    v3 -= v3
    var v6 int8 = 42
    var v7 string = ("\n\t" + (((("foo bar" + "a") + "hello") + "hello") + ("" + "")))
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: fi8, F1: 1}
        if (true && false) {
            v7 = v7
            for fi10 := 0; fi10 < 16; fi10++ {
                v2 = append(v2, 4294967295)
                _ = fi10
            }
        }
        _ = fi8
        _ = v9
    }
    if true {
        v11 := GenS1{F0: 128, F1: 2}
        v1 -= float32(-2.25)
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v3 -= v3
        var v13 int64 = 9223372036854775807
        _ = fi12
        _ = v13
    }
    switch v7 {
    case "foo bar":
        delete(v4, 4294967295)
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS2{F0: 42}
        _ = fi14
        _ = v15
    }
    _ = len(v5)
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 int64 = 775
        var v18 float32 = ((0.1 + v1) * (1.5 * v1))
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = len(v4)
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 int16 = 32767
        fmt.Println(v20, v7)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint64 = 42
        fmt.Println(v22, v6)
        _ = fi21
        _ = v22
    }
    v6 = v6
    if false {
        v6 += 100
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
