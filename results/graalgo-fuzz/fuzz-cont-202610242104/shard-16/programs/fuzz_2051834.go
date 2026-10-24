package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 bool = (-2.25 >= 3.14159)
    v5 := make(map[int]uint32)
    var v6 uint64 = 18446744073709551615
    fmt.Println(v6, v6)
    fmt.Println(v1, v6, v4)
    v1 = ((v1 + v1) + v1)
    var v7 int8 = 100
    var v8 int = int(v6)
    if v4 {
        v6 += (^(v6))
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: v7, F1: v4}
            _ = fi9
            _ = v10
        }
    }
    if v4 {
        v8 -= v8
        v4 = false
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if false {
            var v12 int16 = -167
            _ = v12
        }
        _ = fi11
    }
    switch "a" {
    case "foo bar":
        switch v8 {
        case 127:
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(v1, v1, v4)
                if (v1 >= v1) {
                    switch int(v6) {
                    case 2147483647:
                        var v14 int32 = -655
                        _ = v14
                    default:
                        v15 := GenS1{F0: 127, F1: (true || (1 != v7))}
                        _ = v15
                    }
                    var v16 int64 = 42
                    _ = v16
                } else {
                    v1 = ((v1 + "hello") + v1)
                }
                var v17 uint8 = 128
                _ = fi13
                _ = v17
            }
        case 961:
            v7 += v7
        case -716:
            v18 := GenS2{F0: 42, F1: (v7 / (((v7 % ((v7) | 1))) | 1))}
            _ = v18
        }
    case "a":
        var v19 bool = ((v4 && (v1 >= "\n\t")) || ((-(v8)) < v8))
        _ = v19
    case "":
        var v20 uint8 = 1
        _ = v20
    default:
        v8 = v8
    }
    fmt.Println(v4)
    v1 = (((v1 + v1) + "hello") + ("foo bar" + "hello"))
    var v21 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
