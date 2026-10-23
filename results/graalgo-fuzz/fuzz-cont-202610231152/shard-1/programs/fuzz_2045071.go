package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    var v3 uint16 = 65535
    v4 := []int{}
    if v2 {
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 int = fi5
            _ = fi5
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: -482, F1: 384}
        if true {
            switch fi7 {
            case 42:
                v9 := GenS2{F0: 235}
                _ = v9
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v10 := GenS3{F0: 2}
            _ = v10
        }
        if len(v4) > 0 {
            v4[0] = (fi7 ^ (fi7 - 0))
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v3)
    var v11 int64 = 0
    _ = len(v1)
    v3 -= v3
    v12 := GenS3{F0: 238}
    v11 = (v11 * (^(v11)))
    var v13 int64 = (v11 | 42)
    v14 := GenS4{F0: 4294967295}
    v11 -= v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
