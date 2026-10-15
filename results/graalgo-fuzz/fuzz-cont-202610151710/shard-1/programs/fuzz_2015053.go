package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint8 = 2
    var v2 string = ("a" + "a")
    v3 := make(map[string]int)
    v4 := GenS1{F0: 2, F1: false}
    fmt.Println(v1, v2)
    v1 += 127
    switch (v2 + v2) {
    case "\n\t":
        if (0 >= 2) {
            for fi5 := 0; fi5 < 12; fi5++ {
                if true {
                    v1 = 127
                }
                v6 := GenS1{F0: (v1 | (v1 - v1)), F1: (42 <= 2)}
                v1 -= (v1 / (((^(v1))) | 1))
                _ = fi5
                _ = v6
            }
            v7 := GenS2{F0: (v2 + ("foo bar" + "\n\t")), F1: v1}
            v8 := GenS2{F0: v2, F1: (v1 % ((128) | 1))}
            _ = v7
            _ = v8
        }
    }
    v9 := GenS1{F0: (^(v1)), F1: (687 < 255)}
    v10 := []int{}
    var v11 float32 = 0.1
    fmt.Println(v1)
    if ((false && (-9223372036854775808 > 785)) || ((!false) || false)) {
        fmt.Println(v1, v1, v1)
        var v12 int32 = -1
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
