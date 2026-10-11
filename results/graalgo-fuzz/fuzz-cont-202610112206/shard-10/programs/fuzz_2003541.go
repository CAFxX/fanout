package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -1
    var v3 uint = 255
    v4 := []int{}
    var v5 string = ("foo bar" + "a")
    if true {
        v2 -= v2
        if false {
            v1[65535] = (-(v2))
        }
        var v6 uint = (v3 >> (((v3 >> ((v3) & 63))) & 63))
        _ = v6
    } else {
        var v7 bool = ((v3 * v3) == (v3 / ((0) | 1)))
        for fi8 := 0; fi8 < 20; fi8++ {
            if true {
                v9 := GenS1{F0: 0, F1: (v3 / (((v3 * v3)) | 1)), F2: (!(0 < -97))}
                _ = v9
            }
            fmt.Println(v2, v7)
            _ = fi8
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            v7 = (1 <= 0)
            switch (v5 + v5) {
            case "a":
                var v11 int64 = int64(fi10)
                _ = v11
            default:
                fmt.Println(fi10, fi10, fi10)
            }
            _ = fi10
        }
        _ = v7
    }
    if false {
        fmt.Println(v3, v3)
        if false {
            switch ("hello" + (("hello" + (v5 + ("hello" + "hello"))) + (v5 + (v5 + "\n\t")))) {
            case "foo bar":
                v2 = v2
            }
            fmt.Println(v2, v2, v2)
        }
    }
    v12 := GenS1{F0: 233, F1: 380, F2: false}
    v3 = 857
    v3 = (^(v3))
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 uint8 = 255
        _ = fi13
        _ = v14
    }
    var v15 int16 = 1
    v16 := GenS2{F0: 0}
    var v17 float32 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
