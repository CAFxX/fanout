package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 bool = (1.5 < 0.1)
    v2 := []int{32767, -341, -9223372036854775808}
    var v3 int32 = -2147483648
    fmt.Println(v1)
    v1 = (255 > 255)
    v2 = append(v2, -819)
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: "a", F1: -101}
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 uint8 = 1
            v3 = v3
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v8 := GenS2{F0: (0 <= -1)}
    v1 = false
    fmt.Println(v3, v3)
    v1 = false
    var v9 uint64 = 42
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
