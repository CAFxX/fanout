package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0() uint {
    v6 := GenS1{F0: 67, F1: 0}
    var v7 float32 = -2.25
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 uint32 = 420
    var v2 uint16 = 2
    var v3 uint16 = v2
    v4 := make(map[int]string)
    var v5 int32 = -600
    if (false && false) {
        fmt.Println(v2, v2, v1)
        _ = len(v4)
    } else {
        v2 -= v3
    }
    if true {
        v3 = v3
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint32 = (v1 << ((133) & 63))
        var v10 int64 = 76
        var v11 uint = 255
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v1 -= (v1 / ((1) | 1))
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v5)
        _ = fi12
    }
    if false {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int16 = 462
            v15 := []int{(^(fi13))}
            var v16 int64 = -9223372036854775808
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        v17 := GenS2{F0: (v1 << ((219) & 63)), F1: -1}
        _ = v17
    } else {
        if true {
            v3 += v3
            for fi18 := 0; fi18 < 19; fi18++ {
                v19 := []int{(-(fi18)), -490, (fi18 / (((fi18 / ((fi18) | 1))) | 1))}
                _ = fi18
                _ = v19
            }
        }
        v20 := []int{127}
        v20 = append(v20, -110)
        _ = v20
    }
    var v21 uint8 = 68
    var v22 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
