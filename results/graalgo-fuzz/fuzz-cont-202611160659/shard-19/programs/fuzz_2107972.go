package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 694
    var v3 int = 32767
    v2 = 252
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 int32 = 1
        v3 = (fi4 * v3)
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 1}
    v3 -= v3
    var v7 int32 = 2147483647
    v3 = 441
    if (576 >= 178) {
        if true {
            if false {
                delete(v1, v3)
                v7 += v7
                if false {
                    v8 := GenS2{F0: v2, F1: 213, F2: -614}
                    delete(v1, 2147483647)
                    _ = v8
                }
            }
            fmt.Println(v3)
            fmt.Println(v3, v7)
        }
        v7 = v7
    } else {
        switch int(v7) {
        case 42:
            v2 += v2
        }
        fmt.Println(v7, v7, v2)
    }
    if ((!true) && false) {
        v7 = int32(v3)
    }
    v9 := GenS1{F0: (v2 | -9)}
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS1{F0: (v2 ^ v2)}
        var v12 int32 = (v7 << ((255) & 63))
        _ = fi10
        _ = v11
        _ = v12
    }
    v7 = v7
    var v13 uint8 = 231
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
