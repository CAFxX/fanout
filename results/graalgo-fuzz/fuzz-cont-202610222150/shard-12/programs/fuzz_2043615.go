package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 string = ""
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 728
    v4 := make(map[uint32]uint32)
    var v5 uint64 = (v3 ^ v3)
    v6 := []int{}
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int32 = 0
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        delete(v4, 0)
        _ = fi10
    }
    _ = len(v4)
    _ = len(v6)
    for fi11 := 0; fi11 < 18; fi11++ {
        _ = len(v4)
        fmt.Println(v1)
        switch fi11 {
        case -620:
            var v12 bool = (65535 > 65535)
            _ = v12
        case 0:
            switch "" {
            case "a":
                v13 := GenS1{F0: 61, F1: 2147483648}
                _ = v13
            }
        case 4294967295:
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint16 = 7
                _ = fi14
                _ = v15
            }
        default:
            v2 = v2
        }
        _ = fi11
    }
    if false {
        var v16 uint32 = 2147483647
        _ = v16
    } else {
        v3 = 815
        v17 := GenS1{F0: 42, F1: 2147483647}
        _ = v17
    }
    var v18 int64 = v2
    v3 = v3
    fmt.Println(v3, v1)
    var v19 int32 = -845
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
