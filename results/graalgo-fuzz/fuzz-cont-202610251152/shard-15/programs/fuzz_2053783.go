package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 580
    v3 := []uint32{42, 0, 1}
    v4 := []uint32{0, 2147483648}
    var v5 int = 2147483647
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 0, F1: v1}
        v5 += (v5 << ((870) & 63))
        _ = fi6
        _ = v7
    }
    v1 += v1
    v4 = append(v4, 0)
    v1 += 85
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 951, F1: 2}
        v1 = v1
        _ = fi8
        _ = v9
    }
    var v10 int8 = 0
    v11 := GenS1{F0: 42, F1: (v1 >> ((541) & 63))}
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS1{F0: -798, F1: (v1 >> ((2) & 63))}
        fmt.Println(v10, fi12)
        v5 += (fi12 | v5)
        _ = fi12
        _ = v13
    }
    switch "a" {
    case "\n\t":
        switch v5 {
        case 4294967295:
            _ = len(v4)
        }
    case "":
        fmt.Println(v5)
    case "a":
        switch (^(v5)) {
        case -2147483648:
            var v14 int32 = -1
            _ = v14
        case 128:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
