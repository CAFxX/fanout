package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    var v1 uint16 = 255
    var v2 uint64 = 42
    v3 := make(map[string]bool)
    var v4 float64 = 0.1
    v5 := []int{}
    for fi6 := 0; fi6 < 11; fi6++ {
        v1 -= (v1 << ((255) & 63))
        _ = fi6
    }
    var v7 float64 = ((0.1 / ((v4 - 3.14159) + v4)) / v4)
    var v8 uint16 = 0
    _ = len(v3)
    v8 -= (v8 - (v1 >> ((4294967295) & 63)))
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int8 = 1
        if true {
            var v12 uint = uint(v2)
            v13 := GenS1{F0: 255}
            var v14 uint8 = 1
            _ = v12
            _ = v13
            _ = v14
        }
        switch "foo bar" {
        case "\n\t":
            if true {
                if (32767 < 665) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v15 int32 = 352
                    _ = v15
                }
            }
        case "":
            v16 := GenS2{F0: 1, F1: "foo bar"}
            _ = v16
        default:
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 string = ""
                var v19 int64 = -75
                if true {
                    v1 += (v8 | 0)
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v1, v9, v4)
    v1 -= 246
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
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
