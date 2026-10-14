package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{2, 127}
    var v2 int8 = 53
    var v3 uint8 = 2
    v4 := []uint32{}
    v5 := make(map[int]uint32)
    v3 = (^(v3))
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint8 = v3
            fmt.Println(v3)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    switch "\n\t" {
    case "hello":
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS1{F0: 255, F1: false, F2: 42}
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: 4294967295, F1: (false && (false || false)), F2: 0}
            _ = fi9
            _ = v12
        }
    case "\n\t":
        var v13 uint = 255
        _ = v13
    case "":
        var v14 int = 32768
        _ = v14
    }
    v5[32768] = 2
    var v15 uint = 65535
    v3 += (v3 >> (((v15 << ((v15) & 63))) & 63))
    v3 += v3
    v16 := GenS2{F0: v3}
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v15)
        _ = fi17
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint = (v15 ^ v15)
        _ = fi18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
