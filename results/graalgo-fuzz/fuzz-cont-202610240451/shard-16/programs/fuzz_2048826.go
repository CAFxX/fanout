package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint; F1 string }

func main() {
    var v1 float64 = 0.0
    var v2 string = "foo bar"
    v3 := make(map[string]int)
    v4 := []string{}
    var v5 int8 = 100
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: 42, F1: 65535, F2: -31}
    v8 := v6[v2]
    v9 := GenS2{F0: 42, F1: v2}
    v2 = (v2 + ("hello" + v2))
    for fi10 := 0; fi10 < 17; fi10++ {
        v8 = v8
        switch fi10 {
        case 32767:
            v1 += (v1 - (v1 / 1.0))
        case 101:
            v11 := v6[v2]
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v2, v2, v5)
    v2 = v2
    v1 = (3.14159 - v1)
    delete(v6, v2)
    var v12 int8 = v5
    if (446 != 32767) {
        v5 -= (v5 >> ((uint(v8)) & 63))
        v13 := GenS2{F0: 508, F1: v2}
        _ = v13
    } else {
        var v14 bool = false
        _ = v14
    }
    v1 -= v1
    v12 += (v5 ^ (v12 << ((606) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
