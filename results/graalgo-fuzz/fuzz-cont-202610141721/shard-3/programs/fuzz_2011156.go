package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint32 }
func genFn0() int16 {
    v4 := GenS1{F0: 42, F1: 177, F2: 42}
    v5 := []int{128}
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 string = "\n\t"
    var v6 bool = (float32(float64((float32(3.14159) - 0.1))) == float32(((float64(float32(float64(((0.1 - (3.14159 / 1e10)) * float32(1e10))))) - 1.5) * -2.25)))
    var v7 uint32 = 533
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint = 1
        v10 := GenS1{F0: (v9 & (v9 << ((255) & 63))), F1: v7, F2: 65535}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v1[(v7 ^ v7)] = v6
        switch fi11 {
        case 127:
            for fi12 := 0; fi12 < 6; fi12++ {
                v7 += (^(v7))
                _ = fi12
            }
        case 32767:
            var v13 float64 = (1e10 - ((float64(3.14159) / 1e10) - float64(0.1)))
            _ = v13
        default:
            var v14 int64 = -150
            _ = v14
        }
        _ = fi11
    }
    var v15 int = 852
    if ((false || v6) && v6) {
        var v16 string = "foo bar"
        _ = v16
    } else {
        v17 := []int{(^(v15)), v15}
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        delete(v2, v7)
        var v19 int = fi18
        if (43 != 127) {
            fmt.Println(fi18, v7, v3)
            v20 := []int{(v19 + (v19 / ((fi18) | 1))), (fi18 + 2147483647), -703}
            delete(v1, 2147483647)
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    v7 = (v7 ^ 2147483648)
    v7 -= 948
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := []int{(v15 << ((42) & 63)), v15}
        var v23 float32 = float32((1.5 - (1.5 / 1e10)))
        _ = fi21
        _ = v22
        _ = v23
    }
    switch (v15 ^ (v15 | (^(v15)))) {
    case 42:
        v1[(^(v7))] = (!(v3 != v3))
    default:
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 uint = 833
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v15, v7, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
