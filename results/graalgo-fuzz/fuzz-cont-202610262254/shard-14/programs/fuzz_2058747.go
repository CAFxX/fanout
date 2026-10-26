package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 1
    var v3 float32 = 1e10
    v4 := []string{"foo bar"}
    var v5 int32 = 42
    var v6 uint32 = 80
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 = v5
    var v7 string = ("a" + ("\n\t" + ""))
    if (float64(v3) > 1e10) {
        _ = len(v1)
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v1[v2] = 1
        _ = fi8
    }
    v9 := GenS1{F0: (v2 ^ v6), F1: -604}
    fmt.Println(v7)
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS2{F0: 0}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (v2 % ((v6) | 1)), F1: 0}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
