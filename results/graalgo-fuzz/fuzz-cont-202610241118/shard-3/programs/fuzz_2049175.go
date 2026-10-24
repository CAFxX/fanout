package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }

func main() {
    var v1 uint8 = 42
    var v2 uint8 = (^(v1))
    var v3 int = 127
    var v4 float64 = 1e10
    var v5 uint32 = 1
    for fi6 := 0; fi6 < 15; fi6++ {
        v3 = -66
        _ = fi6
    }
    fmt.Println(v1, v1)
    v7 := []int{}
    var v8 int = (v3 % (((v3 ^ (v3 ^ (v3 << ((65535) & 63))))) | 1))
    var v9 float32 = (((1e10 / (((1e10 + (0.0 / 1e10)) + -2.25) - float32(v4))) - 0.1) + 0.1)
    if (v9 >= v9) {
        _ = len(v7)
    }
    v8 -= v3
    if (true && (!false)) {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v4 -= float64((0.0 / (0.0 - v9)))
        var v10 uint = 771
        _ = v10
    }
    fmt.Println(v3, v5, v9)
    _ = len(v7)
    if false {
        v7 = append(v7, 982)
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: 1, F1: ("hello" + (("a" + ("a" + "hello")) + "\n\t")), F2: -1}
            _ = fi11
            _ = v12
        }
        v3 = (v3 % ((v3) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
