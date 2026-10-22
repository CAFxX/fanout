package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint = 2
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    var v6 int8 = 127
    if false {
        var v7 float32 = 0.0
        _ = v7
    } else {
        var v8 uint16 = 2
        v9 := GenS1{F0: -372}
        _ = v8
        _ = v9
    }
    v1 = (v1 / (((v1 << ((572) & 63))) | 1))
    v1 += (v1 >> ((1) & 63))
    delete(v2, 850)
    switch (("hello" + "foo bar") + "foo bar") {
    case "\n\t":
        v10 := GenS2{F0: 2, F1: (!(!(0 > 1000)))}
        _ = v10
    }
    switch 512 {
    case 32768:
        if true {
            var v11 uint64 = 1
            fmt.Println(v11)
            _ = v11
        } else {
            var v12 int = 0
            _ = v12
        }
    }
    v1 = v1
    v6 -= (^(v6))
    fmt.Println(v6, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
