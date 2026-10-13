package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint8 = 128
    var v2 int = 128
    v3 := []uint32{}
    fmt.Println(v2, v2)
    v1 = 1
    v3 = append(v3, 2147483648)
    v4 := GenS1{F0: 1}
    for fi5 := 0; fi5 < 3; fi5++ {
        v2 = 0
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi5
    }
    fmt.Println(v1)
    v6 := GenS2{F0: 1, F1: 4294967295}
    _ = len(v3)
    v3 = append(v3, 2)
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS3{F0: 42}
        switch "hello" {
        case "foo bar":
            _ = len(v3)
        }
        _ = fi7
        _ = v8
    }
    v1 -= (^(v1))
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
