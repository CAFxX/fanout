package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    var v3 int16 = 1
    for fi4 := 0; fi4 < 2; fi4++ {
        fmt.Println(v1, v3, v3)
        _ = fi4
    }
    switch 228 {
    case 811:
        v3 += -32768
    default:
        for fi5 := 0; fi5 < 10; fi5++ {
            switch fi5 {
            case -9223372036854775808:
                var v6 uint8 = 10
                _ = v6
            default:
                v7 := v2[(fi5 - fi5)]
                _ = v7
            }
            var v8 uint32 = 2
            _ = fi5
            _ = v8
        }
    }
    v9 := GenS1{F0: 1}
    if true {
        v10 := GenS2{F0: 348, F1: 486}
        _ = v10
    } else {
        var v11 uint8 = 2
        _ = v11
    }
    fmt.Println(v1)
    v12 := []int{}
    fmt.Println(v1, v1)
    v13 := GenS3{F0: (!(false && (2 >= 2147483647))), F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
