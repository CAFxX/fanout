package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
func genFn0(p0 uint16, p1 int8, p2 bool) int64 {
    var v4 string = "\n\t"
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: (fi5 ^ fi5), F1: 0}
        v7 := []int{fi5}
        p0 -= (p0 - (p0 * 255))
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{}
    fmt.Println(p2, p1)
    v8 = append(v8, 170)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 218
}

func main() {
    var v1 float32 = float32(float64(float32(1e10)))
    v2 := make(map[int]uint32)
    var v3 float32 = 0.0
    if (65535 > 93) {
        v1 += v3
        v9 := GenS1{F0: 127, F1: 63}
        _ = v9
    } else {
        v10 := GenS2{F0: 0, F1: false, F2: 42}
        if (!((float32(3.14159) / v1) == v3)) {
            v1 += ((v3 / v1) + float32(float64(v1)))
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 uint32 = 0
                _ = fi11
                _ = v12
            }
        }
        var v13 int8 = 127
        _ = v10
        _ = v13
    }
    var v14 uint32 = 2
    v1 += v3
    if ((!(true || (0 <= 32767))) || false) {
        v15 := []int{-249, 4294967295, 127}
        _ = v15
    }
    var v16 int = 128
    var v17 uint64 = 936
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint = 519
        var v20 uint16 = 32767
        var v21 uint16 = v20
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v14 -= 2147483647
    var v22 int16 = 42
    delete(v2, v16)
    var v23 uint8 = 38
    var v24 uint = 65535
    v3 -= 1.5
    var v25 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
