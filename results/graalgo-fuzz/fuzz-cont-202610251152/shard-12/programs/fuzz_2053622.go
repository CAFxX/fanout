package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 int16, p2 int) uint8 {
    switch "a" {
    case "\n\t":
        var v7 float32 = ((-2.25 - float32(1e10)) - (3.14159 - (float32(1.5) * 0.1)))
        _ = v7
    }
    v8 := GenS1{F0: ("" + "\n\t"), F1: 525, F2: 42}
    var v9 float32 = float32(0.0)
    var v10 uint8 = 141
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return (v10 | 255)
}

func main() {
    var v1 int16 = 1000
    var v2 float64 = 1e10
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 uint16 = 2
    v6 := make(map[uint32]int)
    var v11 string = ("hello" + "hello")
    v12 := GenS1{F0: "foo bar", F1: 4294967295, F2: 42}
    v13 := []int{127, 32768, -9223372036854775808}
    var v14 int16 = (v1 ^ (v1 | v1))
    var v15 int16 = v14
    if true {
        var v16 uint16 = v5
        _ = v16
    }
    if (!true) {
        v17 := GenS2{F0: v11}
        fmt.Println(v11)
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 uint64 = 4294967295
            _ = len(v6)
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    if (127 <= 42) {
        v6[861] = -1
    } else {
        var v20 string = v11
        var v21 bool = (!false)
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v5 += (v5 / (((v5 % ((v5) | 1))) | 1))
        v23 := GenS1{F0: ("" + v11), F1: 1, F2: 533}
        _ = fi22
        _ = v23
    }
    v13 = append(v13, 9223372036854775807)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
