package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 0
    var v3 int64 = -9223372036854775808
    var v4 float64 = (0.0 * 0.1)
    var v5 float32 = 1.5
    var v6 int = 32767
    _ = len(v1)
    var v7 int = 677
    v4 += v4
    switch (v6 + 42) {
    case -771:
        fmt.Println(v6, v4)
    case 0:
        v5 = 1e10
    case 9223372036854775807:
        v4 = float64((0.1 + 0.1))
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint16 = 0
        _ = fi8
        _ = v9
    }
    v5 += 1.5
    var v10 uint16 = 42
    if true {
        if true {
            v11 := GenS1{F0: (v2 * v2)}
            _ = v11
        }
        if (("foo bar" + "hello") < ("foo bar" + ("foo bar" + "a"))) {
            v6 += v7
            v4 -= ((((1e10 * ((v4 / (-2.25 - (0.1 - v4))) / v4)) + float64(float32(0.1))) / 1e10) + v4)
        }
    }
    v5 -= v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
