package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    var v1 float32 = 1.5
    v2 := []int{42, 42, 684}
    v3 := make(map[int]string)
    v4 := []uint32{294}
    v5 := make(map[uint32]int)
    v1 -= 0.0
    var v6 int64 = 779
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int32 = 0
        for fi9 := 0; fi9 < 5; fi9++ {
            v4 = append(v4, 172)
            fmt.Println(v6, fi7)
            var v10 int8 = 1
            _ = fi9
            _ = v10
        }
        v6 -= (v6 >> ((255) & 63))
        _ = fi7
        _ = v8
    }
    var v11 bool = false
    switch 127 {
    case 1:
        v11 = v11
    case 127:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        delete(v3, 9223372036854775807)
    }
    v11 = v11
    if v11 {
        v12 := v5[0]
        v13 := GenS1{F0: (!(v11 || (!(18446744073709551615 >= 314)))), F1: (v12 - (v12 - -404))}
        _ = v12
        _ = v13
    } else {
        v5[300] = 103
        _ = len(v5)
        switch "hello" {
        case "\n\t":
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 uint64 = 927
                var v16 float64 = 0.1
                _ = fi14
                _ = v15
                _ = v16
            }
        case "a":
            v17 := GenS1{F0: false, F1: 1}
            _ = v17
        }
    }
    v18 := GenS2{F0: 0, F1: 0}
    v11 = (65535 > 219)
    var v19 int32 = -1
    var v20 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
