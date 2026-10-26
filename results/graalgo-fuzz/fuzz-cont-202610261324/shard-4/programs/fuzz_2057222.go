package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint) uint16 {
    p0 = (^(p0))
    v6 := []int{-202, -239}
    if (!(!true)) {
        p0 += (^(p0))
        p0 = (p0 >> (((p0 | p0)) & 63))
        var v7 uint = p0
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = v6
    _ = v8
    return 32767
}

func main() {
    var v1 uint64 = 1
    var v2 int64 = -148
    v3 := make(map[string]int)
    var v4 uint8 = 2
    var v5 int8 = -1
    v2 = (v2 | v2)
    v1 = 4294967295
    fmt.Println(v5, v1)
    fmt.Println(v4, v4)
    var v9 float32 = ((-2.25 / 1.0) + -2.25)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint8 = (v4 ^ uint8(v1))
        v2 += 80
        delete(v3, "\n\t")
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: v4, F1: v4}
    v13 := GenS1{F0: 127, F1: v4}
    if true {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int16 = int16(v5)
            _ = fi14
            _ = v15
        }
    } else {
        delete(v3, "foo bar")
        fmt.Println(v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
