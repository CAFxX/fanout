package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v4 int32 = 0
    var v5 float32 = 0.1
    fmt.Println(v5, v5)
    if (!((v4 - v4) <= v4)) {
        v5 = (v5 + 1.5)
        var v6 uint32 = 0
        _ = v6
    } else {
        fmt.Println(v5, v5)
    }
    fmt.Println(v4)
    v7 := GenS1{F0: 42}
    var v8 bool = false
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 int = 127
        if (!v8) {
            for fi11 := 0; fi11 < 3; fi11++ {
                _ = len(v3)
                fmt.Println(v4, fi11, fi11)
                v12 := GenS1{F0: 229}
                _ = fi11
                _ = v12
            }
            switch 9223372036854775807 {
            case 128:
                var v13 float64 = 0.1
                _ = v13
            default:
                switch "\n\t" {
                case "hello":
                    var v14 uint = 18446744073709551615
                    _ = v14
                case "\n\t":
                    if (128 == 2) {
                        var v15 uint = 255
                        _ = v15
                    }
                default:
                    v16 := GenS2{F0: 1, F1: 32767, F2: 127}
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                _ = len(v3)
                _ = fi17
            }
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v5, v8)
    v18 := v1[846]
    var v19 uint64 = 975
    var v20 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
