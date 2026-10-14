package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 127
    v5 := GenS1{F0: 32768, F1: 1}
    var v6 uint16 = 1
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: (fi7 >> ((509) & 63)), F1: (v6 >> ((1) & 63))}
        v9 := v2[(fi7 % (((fi7 - fi7)) | 1))]
        _ = fi7
        _ = v8
        _ = v9
    }
    delete(v3, 687)
    v6 = 2
    var v10 float64 = -2.25
    _ = len(v3)
    var v11 float64 = v10
    switch "a" {
    case "a":
        switch 4294967295 {
        case 718:
            if true {
                var v12 float64 = ((1e10 - v11) * ((float64(1e10) * 0.1) + v10))
                v13 := GenS1{F0: -2147483648, F1: v6}
                _ = v12
                _ = v13
            }
        case 2147483647:
            switch 570 {
            case 4294967295:
                var v14 uint64 = 0
                _ = v14
            case 42:
                var v15 uint64 = 18446744073709551615
                _ = v15
            case 128:
                v16 := []int{-947, -9223372036854775808, 103}
                _ = v16
            }
        case -2147483648:
            v17 := []int{42, 127, 217}
            _ = v17
        default:
            v10 = (v11 * v10)
        }
    case "":
        var v18 uint16 = 65535
        _ = v18
    default:
        v4 = (v4 << ((42) & 63))
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v10 -= v10
        _ = fi19
    }
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
