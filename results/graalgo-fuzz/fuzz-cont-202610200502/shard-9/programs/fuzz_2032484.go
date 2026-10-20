package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint32; F1 int }

func main() {
    var v1 uint64 = 263
    v2 := make(map[int]int)
    var v3 string = (("hello" + "hello") + "")
    fmt.Println(v3, v1)
    var v4 uint64 = (v1 & 398)
    if (!false) {
        v5 := GenS1{F0: -1, F1: -1}
        switch 127 {
        case -328:
            for fi6 := 0; fi6 < 2; fi6++ {
                switch v3 {
                case "foo bar":
                    for fi7 := 0; fi7 < 20; fi7++ {
                        v1 -= 236
                        v8 := GenS2{F0: (fi6 << ((0) & 63)), F1: fi6}
                        var v9 float64 = 1.5
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                default:
                    var v10 float32 = ((3.14159 + (1.5 - (float32(1.5) / 0.1))) / 1.0)
                    _ = v10
                }
                _ = fi6
            }
        case 1:
            v11 := GenS3{F0: 4294967295, F1: 9223372036854775807}
            _ = v11
        }
        _ = v5
    }
    var v12 uint16 = 0
    fmt.Println(v4, v3, v1)
    var v13 int64 = -9223372036854775808
    v4 = v4
    var v14 float64 = (1e10 - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
