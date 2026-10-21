package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 int = 648
    switch 32768 {
    case -2147483648:
        var v4 float64 = 1e10
        _ = v4
    case 580:
        v3 -= (v3 | -1)
    default:
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            fmt.Println(v3)
        }
    }
    v5 := v2[860]
    fmt.Println(v3)
    fmt.Println(v3)
    var v6 uint = 59
    var v7 int32 = 96
    for fi8 := 0; fi8 < 16; fi8++ {
        switch v3 {
        case 32768:
            switch "a" {
            case "foo bar":
                var v9 uint64 = 4294967295
                _ = v9
            default:
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch 2147483647 {
                    case 42:
                        var v11 uint16 = 32767
                        _ = v11
                    }
                    _ = fi10
                }
            }
        case 4294967295:
            fmt.Println(v3)
        }
        var v12 int16 = 847
        var v13 uint = v6
        _ = fi8
        _ = v12
        _ = v13
    }
    var v14 uint64 = 4294967295
    fmt.Println(v7)
    _ = len(v1)
    v15 := GenS1{F0: 0, F1: 1000}
    v3 = v3
    v5 -= v5
    var v16 int64 = 491
    v17 := []int{v3, v3, -950}
    v5 = (v5 * v5)
    switch (^(v3)) {
    case 0:
        if (!((v7 << ((2) & 63)) != v7)) {
            _ = len(v2)
            v17 = append(v17, v3)
        }
    case 127:
        switch (v3 | (v3 + v3)) {
        case 1:
            v18 := GenS1{F0: v16, F1: 1000}
            _ = v18
        case 128:
            fmt.Println(v7, v3)
        case 708:
            var v19 uint = v6
            _ = v19
        }
    case -1:
        var v20 uint64 = 42
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
