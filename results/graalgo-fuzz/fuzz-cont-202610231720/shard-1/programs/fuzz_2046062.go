package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 807
    var v2 uint = (v1 / ((312) | 1))
    v3 := make(map[uint32]int)
    var v4 int32 = -903
    v5 := make(map[string]bool)
    var v6 uint32 = 983
    v7 := GenS1{F0: "\n\t", F1: 153}
    if (0 > 2147483647) {
        fmt.Println(v6)
        v8 := []int{}
        v9 := GenS1{F0: "hello", F1: 42}
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: "hello", F1: 42}
        _ = v10
    }
    var v11 uint16 = 816
    if (0.0 != (0.0 - 3.14159)) {
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        v14 := GenS2{F0: 945}
        _ = v14
    }
    var v15 int32 = -632
    for fi16 := 0; fi16 < 19; fi16++ {
        v11 += 2
        _ = fi16
    }
    v15 += (-(v4))
    var v17 uint = v1
    v6 = 840
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 string = ""
        v20 := GenS1{F0: ((v19 + v19) + ("\n\t" + v19)), F1: 42}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
