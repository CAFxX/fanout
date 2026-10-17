package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    v1 := []int{9223372036854775807}
    var v2 int = 127
    var v3 int16 = 32767
    v4 := GenS1{F0: -1}
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(v2)
        fmt.Println(v3)
        v2 -= -251
        _ = fi5
    }
    v3 += 32767
    v2 = (^(v2))
    var v6 float64 = 3.14159
    var v7 uint64 = 1
    v8 := GenS1{F0: -1}
    var v9 uint = 2
    v10 := GenS2{F0: 100, F1: 42}
    v3 -= v3
    if false {
        fmt.Println(v7, v6, v7)
    } else {
        switch v2 {
        case 32767:
            _ = len(v1)
        case -1:
            if ((v3 % (((v3 + v3)) | 1)) >= 32767) {
                fmt.Println(v2, v9)
            } else {
                var v11 uint32 = 1
                v3 -= v3
                _ = v11
            }
        }
        _ = len(v1)
        fmt.Println(v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
