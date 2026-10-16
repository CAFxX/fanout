package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int16 = -635
    var v2 uint64 = 18446744073709551615
    var v3 int16 = -32768
    v4 := make(map[string]int)
    v5 := []int{32767, 873, -2147483648, -1}
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: v3}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: v3}
    _ = len(v5)
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 uint64 = (v2 / ((v2) | 1))
            v12 := GenS2{F0: 2}
            _ = fi10
            _ = v11
            _ = v12
        }
        v3 = (v3 - v1)
        v1 = (^(v1))
        _ = fi9
    }
    v2 -= v2
    var v13 string = "\n\t"
    switch 108 {
    case 127:
        v5 = append(v5, 870)
    case -2147483648:
        var v14 float32 = 0.1
        _ = v14
    }
    v1 = (v3 & (-(v3)))
    v15 := GenS3{F0: (^(v2))}
    switch v13 {
    case "\n\t":
        var v16 uint32 = 430
        _ = v16
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        var v17 string = v13
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
