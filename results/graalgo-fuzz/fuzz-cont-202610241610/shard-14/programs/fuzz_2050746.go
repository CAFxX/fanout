package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint32 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    v4 := GenS1{F0: 2147483647, F1: 267}
    var v5 uint16 = 287
    v5 += v5
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 65535, F1: 0}
            _ = fi7
            _ = v8
        }
        v9 := []int{(-(fi6)), fi6, fi6}
        _ = fi6
        _ = v9
    }
    v10 := GenS2{F0: (v5 + (v5 ^ v5)), F1: (251 == 255), F2: 2147483647}
    fmt.Println(v5, v5)
    fmt.Println(v1, v1)
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
