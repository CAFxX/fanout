package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 174
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    var v5 int64 = 0
    v6 := []int{v2}
    var v7 float32 = 3.14159
    switch -122 {
    case 32:
        switch "" {
        case "a":
            v5 = v5
        case "":
            v7 += 0.0
        case "\n\t":
            v8 := v3[391]
            _ = v8
        }
    default:
        v7 = 1e10
    }
    var v9 float64 = (1e10 * float64(v7))
    v6 = append(v6, (v2 / (((v2 % ((v2) | 1))) | 1)))
    for fi10 := 0; fi10 < 15; fi10++ {
        v7 -= (float32(v9) + -2.25)
        fmt.Println(v2, v9)
        if true {
            v1[(("hello" + "a") + ("hello" + ("hello" + "a")))] = (127 >= (v2 / ((v2) | 1)))
        }
        _ = fi10
    }
    v11 := v1["hello"]
    var v12 int32 = 706
    v9 -= 1e10
    var v13 int = 128
    v11 = (true && v11)
    v12 = (v12 ^ -2147483648)
    var v14 uint = 18446744073709551615
    for fi15 := 0; fi15 < 12; fi15++ {
        v13 = -9223372036854775808
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v16 := GenS1{F0: 18, F1: 975, F2: 42}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 49, F1: 42, F2: 1000}
    v1[""] = true
    v18 := GenS1{F0: -128, F1: 18446744073709551615, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
