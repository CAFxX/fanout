package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := []int{-9223372036854775808, 9223372036854775807}
    v2 := []string{}
    var v3 uint8 = 42
    var v4 uint64 = 1
    var v5 uint16 = 755
    _ = len(v1)
    v6 := GenS1{F0: 42, F1: v5}
    v1 = append(v1, -9223372036854775808)
    var v7 uint32 = 2
    v5 -= uint16(v7)
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: (v7 - v7), F1: v5}
        v4 += v4
        _ = fi8
        _ = v9
    }
    var v10 uint16 = v5
    v3 = v3
    switch 32768 {
    case 203:
        v10 = v10
    }
    if false {
        var v11 uint8 = 255
        v5 = uint16(v7)
        var v12 int = -1
        _ = v11
        _ = v12
    } else {
        v3 = v3
        var v13 int = -1
        var v14 int8 = 16
        _ = v13
        _ = v14
    }
    fmt.Println(v10, v4)
    var v15 float64 = 0.0
    if false {
        v4 += (v4 ^ v4)
        switch 684 {
        case -9223372036854775808:
            v3 -= v3
        case 4294967295:
            if false {
                var v16 uint16 = 1
                _ = v16
            } else {
                for fi17 := 0; fi17 < 19; fi17++ {
                    var v18 string = "a"
                    v2 = append(v2, ("hello" + v18))
                    _ = fi17
                    _ = v18
                }
            }
        case 959:
            v19 := GenS1{F0: 2147483647, F1: 113}
            _ = v19
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
