package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 int16) uint32 {
    v6 := GenS1{F0: true}
    p0 -= -289
    _ = p0
    _ = v6
    return 372
}

func main() {
    var v1 int32 = -1
    var v2 uint8 = 127
    var v3 uint8 = 0
    var v4 int64 = 850
    v5 := make(map[int]int)
    delete(v5, 128)
    v4 = (v4 - -9223372036854775808)
    v7 := GenS1{F0: true}
    var v8 uint8 = 42
    v9 := []int{0}
    var v10 int64 = (v4 << ((255) & 63))
    v10 -= 9223372036854775807
    v1 -= -524
    switch -167 {
    case -663:
        v11 := GenS2{F0: 609, F1: false}
        _ = v11
    case 0:
        _ = len(v9)
    default:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        if true {
            fmt.Println(v2)
        }
        _ = len(v9)
        fmt.Println(v2, v4, fi12)
        _ = fi12
    }
    v3 += 128
    if false {
        _ = len(v5)
        v10 += int64(v1)
    } else {
        v13 := v5[-1]
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := GenS1{F0: (true || (!true))}
            var v16 uint16 = 65535
            var v17 int = (fi14 + (fi14 % (((^(fi14))) | 1)))
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
