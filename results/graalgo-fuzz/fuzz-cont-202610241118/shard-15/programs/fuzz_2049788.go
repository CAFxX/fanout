package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int16 = 42
    var v2 float64 = 0.0
    var v3 int8 = 114
    var v4 bool = (1 < 255)
    v5 := make(map[int]int)
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 4; fi7++ {
        v3 = (v3 >> ((1) & 63))
        delete(v5, fi7)
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(v2, v4, v4)
            v9 := []int{}
            v10 := v5[fi8]
            v6 = (v6 + v6)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := []int{}
    v3 -= v3
    v5[-9223372036854775808] = 32767
    v2 -= 1.5
    if len(v11) > 0 {
        v11[0] = 32767
    }
    v2 -= 0.0
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: -1}
        switch ((v6 + "a") + (v6 + "")) {
        case "a":
            for fi14 := 0; fi14 < 11; fi14++ {
                v1 = 42
                v15 := GenS1{F0: 217}
                _ = fi14
                _ = v15
            }
        }
        fmt.Println(v6)
        _ = fi12
        _ = v13
    }
    v3 = (v3 % (((-(v3))) | 1))
    v5[-109] = -429
    v16 := GenS1{F0: -403}
    if true {
        v17 := GenS2{F0: 18446744073709551615, F1: -244}
        var v18 uint = 65535
        _ = v17
        _ = v18
    } else {
        v3 -= v3
        v6 = v6
    }
    delete(v5, 42)
    _ = len(v5)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v19 int8 = (v3 / ((-1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
