package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint8 = 1
    var v2 bool = false
    v3 := make(map[int]bool)
    var v4 uint16 = 42
    v5 := []int{1, 977, 0, 4294967295}
    v6 := GenS1{F0: 1, F1: 569}
    fmt.Println(v2, v4, v4)
    if (4294967295 <= 127) {
        switch -9223372036854775808 {
        case -457:
            for fi7 := 0; fi7 < 6; fi7++ {
                for fi8 := 0; fi8 < 15; fi8++ {
                    fmt.Println(fi7, fi7)
                    var v9 uint8 = (v1 & v1)
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
        case 795:
            v2 = (32767 >= 42)
        case -2147483648:
            var v10 int16 = 162
            _ = v10
        default:
            v5 = append(v5, 1)
        }
    } else {
        v2 = true
        var v11 string = "\n\t"
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint16 = 720
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int8 = 41
        _ = fi14
        _ = v15
    }
    var v16 bool = v2
    if true {
        var v17 uint = 65535
        _ = v17
    } else {
        var v18 uint8 = v1
        fmt.Println(v1, v18, v2)
        fmt.Println(v16, v1)
        _ = v18
    }
    v5 = append(v5, -980)
    fmt.Println(v1)
    if v16 {
        var v19 uint = 255
        v16 = ((2147483647 <= 42) && v16)
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS2{F0: 4294967295, F1: 1000}
            var v22 int = 32767
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = v19
    }
    if (!v2) {
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 int16 = -32768
            _ = fi23
            _ = v24
        }
    }
    v2 = v16
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    v2 = true
    var v25 uint16 = (v4 >> ((255) & 63))
    var v26 float64 = 0.1
    var v27 float32 = (1e10 + -2.25)
    var v28 float64 = float64(0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
