package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint64 }

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]bool)
    var v3 bool = ((3.14159 / 0.1) < float64(1.5))
    var v4 int = -1
    v5 := make(map[string]uint32)
    v6 := []int{}
    v7 := GenS1{F0: ((v1 % (((v1 | (v1 - v1))) | 1)) >= (v1 << ((v1) & 63))), F1: v3, F2: 0}
    if v3 {
        v3 = (!true)
        fmt.Println(v3, v3, v1)
        v8 := GenS1{F0: (true || (int16(v4) != 32767)), F1: v3, F2: 42}
        _ = v8
    }
    var v9 float32 = 1e10
    var v10 uint64 = 18446744073709551615
    v4 += (v4 / (((v4 * v4)) | 1))
    v5["foo bar"] = 4294967295
    v11 := GenS1{F0: (((v9 + ((0.1 * ((float32((0.0 - 1.5)) / 1.5) - -2.25)) + 3.14159)) - v9) == v9), F1: true, F2: (v10 ^ (^(v10)))}
    var v12 string = ("\n\t" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
