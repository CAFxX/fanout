package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int; F1 string }

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    v4 := GenS1{F0: 419, F1: 255, F2: 0}
    var v5 float64 = ((((-2.25 / 3.14159) * float64(1e10)) / -2.25) - (((1e10 * 1e10) + 1e10) - (1e10 * 0.0)))
    var v6 int32 = 666
    switch "hello" {
    case "a":
        delete(v3, ("a" + ((("\n\t" + ("hello" + "\n\t")) + "\n\t") + "\n\t")))
    }
    v7 := GenS2{F0: 2147483648, F1: 156}
    for fi8 := 0; fi8 < 13; fi8++ {
        v5 = v5
        _ = fi8
    }
    var v9 bool = false
    var v10 int32 = 1
    v11 := GenS3{F0: 128, F1: "hello"}
    switch 127 {
    case 9223372036854775807:
        for fi12 := 0; fi12 < 19; fi12++ {
            v3[((("hello" + ("hello" + ("hello" + "a"))) + "hello") + "a")] = fi12
            fmt.Println(v5, v9)
            v13 := GenS2{F0: 0, F1: 255}
            _ = fi12
            _ = v13
        }
    }
    var v14 uint8 = 0
    v15 := []int{}
    fmt.Println(v14, v6, v14)
    v5 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
