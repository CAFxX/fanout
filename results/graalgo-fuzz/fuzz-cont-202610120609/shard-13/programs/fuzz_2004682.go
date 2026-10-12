package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0(p0 int32, p1 bool) uint64 {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: 2, F1: 0}
            p0 = p0
            p0 = (^(p0))
            _ = fi7
            _ = v8
        }
        fmt.Println(p1, p0)
        p0 -= p0
        _ = fi6
    }
    v9 := GenS1{F0: 1, F1: 1}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 11
    v2 := make(map[string]uint32)
    var v3 int64 = -635
    var v4 uint16 = 65535
    var v5 uint8 = v1
    var v10 int64 = (v3 + v3)
    var v11 float64 = (((0.1 + ((1.5 - 3.14159) / (float64(-2.25) / 1.5))) / 3.14159) / ((float64(((-2.25 / -2.25) / -2.25)) * 0.1) + float64(float32(1.5))))
    var v12 float64 = 0.1
    var v13 string = "\n\t"
    if false {
        switch -1 {
        case 32768:
            v14 := GenS1{F0: 208, F1: 150}
            _ = v14
        }
    } else {
        v13 = v13
    }
    if (v4 < 1) {
        var v15 uint32 = 0
        v5 -= (v1 / ((127) | 1))
        _ = v15
    }
    var v16 float32 = ((0.0 - 0.0) * -2.25)
    v17 := []int{}
    for fi18 := 0; fi18 < 7; fi18++ {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = len(v17)
        if true {
            fmt.Println(v1, v16)
            var v19 uint = 2
            _ = v19
        }
        _ = fi18
    }
    v20 := v2[(v13 + v13)]
    v21 := GenS1{F0: 4294967295, F1: 0}
    v11 -= 0.1
    fmt.Println(v11)
    var v22 int32 = 980
    v23 := GenS1{F0: genFn0((v22 << ((255) & 63)), false), F1: 4294967295}
    delete(v2, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
