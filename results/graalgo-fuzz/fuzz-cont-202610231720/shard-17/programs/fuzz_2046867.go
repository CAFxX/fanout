package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }

func main() {
    var v1 float64 = (1e10 / 3.14159)
    var v2 int64 = 0
    var v3 uint8 = 231
    var v4 int = 128
    v5 := make(map[string]bool)
    var v6 uint64 = 1
    v6 = (v6 - (v6 | (v6 & v6)))
    v7 := GenS1{F0: 1, F1: v4}
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            v5[(("a" + (("\n\t" + ("hello" + (("foo bar" + (("" + "hello") + (("" + "foo bar") + "\n\t"))) + ""))) + "")) + ((((("hello" + "\n\t") + "hello") + "") + "hello") + ("hello" + "")))] = false
            _ = fi9
        }
        _ = fi8
    }
    v3 += v3
    if false {
        var v10 int32 = 0
        _ = v10
    } else {
        v6 -= 42
        v11 := GenS2{F0: 1000, F1: uint64(v3), F2: 42}
        _ = v11
    }
    v5["\n\t"] = true
    var v12 int8 = -128
    _ = len(v5)
    fmt.Println(v6, v2, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
