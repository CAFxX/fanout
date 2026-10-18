package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint32 = 0
    var v2 string = "\n\t"
    var v3 float64 = (0.1 / 0.1)
    var v4 int = 42
    var v5 int = (v4 - (v4 >> ((0) & 63)))
    v6 := []int{}
    var v7 int16 = 32767
    v6 = append(v6, (v4 << ((1) & 63)))
    v8 := GenS1{F0: 350, F1: v2}
    for fi9 := 0; fi9 < 4; fi9++ {
        v1 = v1
        v4 -= v4
        _ = fi9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        fmt.Println(fi10, v4)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi10
    }
    var v11 string = (v2 + v2)
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 uint64 = 376
        if (v5 > v5) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v14 float64 = (1.5 - 1.5)
            fmt.Println(v2)
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    var v15 int = int(v7)
    v2 = ((v2 + "a") + ((((("foo bar" + v2) + (v2 + v2)) + (v11 + ("\n\t" + v11))) + v2) + ((v2 + ("foo bar" + (v11 + v11))) + "")))
    var v16 int8 = 1
    v17 := GenS2{F0: v1}
    if (4294967295 != v5) {
        fmt.Println(v5, v3, v3)
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
