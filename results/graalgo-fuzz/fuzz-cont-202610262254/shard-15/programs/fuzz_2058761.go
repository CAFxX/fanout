package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 127
    v3 := make(map[int]int)
    v4 := []uint32{}
    delete(v3, 4294967295)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(fi5, v2)
        _ = fi5
    }
    v2 = v2
    _ = len(v3)
    if ((v2 & v2) <= (v2 + v2)) {
        v6 := GenS1{F0: 78, F1: 4294967295}
        _ = v6
    } else {
        _ = len(v4)
        v2 = (v2 % (((v2 << ((255) & 63))) | 1))
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    v2 = 42
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = 1
        if true {
            var v9 uint32 = 42
            _ = v9
        } else {
            var v10 int16 = -32768
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v2)
    _ = len(v3)
    var v11 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
