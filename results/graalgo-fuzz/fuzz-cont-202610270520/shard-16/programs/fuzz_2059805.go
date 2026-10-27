package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 uint; F2 uint8 }

func main() {
    var v1 uint32 = 42
    var v2 float32 = 3.14159
    var v3 bool = false
    v4 := GenS1{F0: "hello", F1: (v1 + (v1 % ((v1) | 1))), F2: 82}
    fmt.Println(v3, v3)
    v5 := []int{2147483647}
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int16 = 32767
        _ = fi6
        _ = v7
    }
    fmt.Println(v1, v3, v3)
    if v3 {
        var v8 int = 2147483647
        v2 += float32((float64(1.5) + 1.5))
        var v9 int32 = 382
        _ = v8
        _ = v9
    } else {
        fmt.Println(v1, v2, v1)
        _ = len(v5)
        var v10 float32 = 1e10
        _ = v10
    }
    if true {
        v3 = (-932 <= 32767)
    }
    var v11 int32 = -1
    v2 = 1.5
    v2 += v2
    if true {
        switch "hello" {
        case "foo bar":
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: "\n\t", F1: v1, F2: 1}
                _ = fi12
                _ = v13
            }
        case "":
            var v14 uint = 365
            _ = v14
        case "\n\t":
            v3 = (1 <= 127)
        default:
            v5 = append(v5, 32768)
        }
        v15 := GenS2{F0: "foo bar", F1: 684, F2: 42}
        _ = v15
    } else {
        v5 = append(v5, 1)
    }
    var v16 float64 = (3.14159 * 3.14159)
    var v17 float64 = 1.5
    var v18 float32 = (0.1 / ((float32((float64(1.5) / v17)) / float32(1e10)) * ((v2 - ((v2 + 3.14159) * -2.25)) * v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
