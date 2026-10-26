package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{4294967295}
    var v3 uint = 1
    var v4 int16 = 0
    v5 := GenS1{F0: 2147483647}
    v3 = (v3 - v3)
    v6 := GenS1{F0: 42}
    if (((1e10 + 1.5) != ((0.1 / 0.1) + 1.5)) && (!(127 < 32767))) {
        fmt.Println(v3, v4)
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int32 = 882
            _ = fi7
            _ = v8
        }
    }
    if (!((v4 >> (((v3 ^ (v3 - (v3 / (((^(v3))) | 1))))) & 63)) != v4)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v9 float32 = ((0.0 + 0.1) - 1e10)
    if (!(false && false)) {
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(v9)
            _ = fi10
        }
    } else {
        v3 -= v3
    }
    v11 := GenS2{F0: v4, F1: ("a" + ("hello" + ""))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
