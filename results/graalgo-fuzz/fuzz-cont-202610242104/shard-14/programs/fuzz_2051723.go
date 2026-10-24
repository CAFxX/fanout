package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }

func main() {
    var v1 int8 = 42
    var v2 string = "a"
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 18; fi4++ {
        v1 -= (v1 << ((2) & 63))
        _ = fi4
    }
    var v5 float64 = 1e10
    v6 := GenS1{F0: 176, F1: v1, F2: v1}
    if (false || (0.0 > (((1e10 * 0.1) / (1e10 + -2.25)) / ((1e10 - (-2.25 * float32(v5))) - ((3.14159 * float32(v5)) - 0.1))))) {
        v7 := []int{4294967295, 42, 127}
        _ = v7
    }
    v3[127] = 2147483647
    v1 = (^(v1))
    v5 = 1.5
    delete(v3, -9223372036854775808)
    v2 = (v2 + "foo bar")
    var v8 bool = ((((v5 * (v5 + (float64(1e10) + 1e10))) != (1.5 + 3.14159)) || (!true)) || (-1 <= 248))
    v1 = 25
    v3[128] = 2
    v9 := GenS1{F0: 2, F1: -1, F2: (v1 - v1)}
    v10 := GenS2{F0: 517, F1: v2, F2: 1}
    switch (((v2 + ("hello" + "a")) + "") + "") {
    case "":
        if true {
            v11 := v3[-1]
            fmt.Println(v11, v2, v5)
            _ = v11
        }
    case "a":
        var v12 uint32 = 42
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
