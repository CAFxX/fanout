package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []uint32{}
    var v2 uint16 = 436
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]int)
    if (true && (((^(v3)) != -1) && (-1 < 1))) {
        var v5 bool = true
        _ = v5
    } else {
        var v6 uint64 = 42
        var v7 int8 = 1
        _ = v6
        _ = v7
    }
    var v8 uint32 = 4294967295
    v9 := GenS1{F0: 127}
    _ = len(v4)
    v8 = v8
    if true {
        _ = len(v4)
    } else {
        _ = len(v4)
    }
    _ = len(v4)
    if (0.0 > (-2.25 - 0.0)) {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
    } else {
        v2 -= (v2 + v2)
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v3, v8, v2)
    case "hello":
        switch 4294967295 {
        case -1:
            v11 := GenS1{F0: 2147483647}
            _ = v11
        }
    case "a":
        v3 += v3
    default:
        v2 = (^(v2))
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v3 = v3
        for fi13 := 0; fi13 < 18; fi13++ {
            if (v3 > v3) {
                v3 -= -808
                fmt.Println(fi12, fi12, v2)
                var v14 float64 = (-2.25 - (-2.25 / float64(1.5)))
                _ = v14
            }
            fmt.Println(v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v1)
            _ = fi13
        }
        _ = fi12
    }
    v8 = v8
    var v15 bool = ((0 == 160) || true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
