package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint32) int8 {
    v5 := []int{}
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 5; fi7++ {
        p0 -= 0
        v8 := GenS2{F0: fi7, F1: 65535, F2: 1}
        _ = fi7
        _ = v8
    }
    var v9 int16 = 1
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint16 = 801
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return -1
}

func main() {
    var v1 uint8 = 0
    var v2 int16 = -352
    var v3 uint32 = 0
    var v4 uint16 = 2
    if true {
        fmt.Println(v4, v1)
        switch 867 {
        case 2147483647:
            var v12 uint64 = 0
            _ = v12
        default:
            v3 = 727
        }
        fmt.Println(v3, v4, v2)
    }
    v13 := []int{-1, -900, 0, 127}
    fmt.Println(v4)
    v1 -= uint8(v3)
    v1 = (^(v1))
    v13 = append(v13, -535)
    v14 := GenS3{F0: (^(v2))}
    for fi15 := 0; fi15 < 2; fi15++ {
        if (0.0 >= 0.1) {
            v1 = 127
            fmt.Println(v4, v3)
        }
        _ = fi15
    }
    var v16 uint8 = (v1 + (v1 ^ v1))
    fmt.Println(v3)
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 bool = (0 <= 42)
        v4 = (v4 * (v4 - (^(v4))))
        fmt.Println(v2, fi17)
        _ = fi17
        _ = v18
    }
    v4 -= (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
